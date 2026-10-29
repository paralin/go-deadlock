package deadlock

import (
	"context"
	"sync"

	protobuf "github.com/aperturerobotics/protobuf-go-lite"
	devents "github.com/paralin/go-deadlock/events"
	gcsm "github.com/paralin/go-deadlock/protocol"
	"github.com/paralin/go-deadlock/socache"
	"github.com/paralin/go-deadlock/state"
	"github.com/paralin/go-steam"
	steamprotocol "github.com/paralin/go-steam/protocol"
	"github.com/paralin/go-steam/protocol/gamecoordinator"
	"github.com/sirupsen/logrus"
)

// AppID is the Steam application ID for Deadlock.
const AppID = 1422450

// handlerMap is the map of message types to handler functions.
type handlerMap map[uint32]func(packet *gamecoordinator.GCPacket) error

// Deadlock owns the GC session and pending requests for one Steam connection.
type Deadlock struct {
	// le records protocol diagnostics without message bodies.
	le logrus.FieldLogger
	// coordinator transports messages to Steam.
	coordinator Coordinator
	// emit publishes events to the caller's continuously drained event stream.
	emit func(any)
	// cache retains coordinator shared objects.
	cache *socache.SOCache
	// handlers maps packet types to their decoders.
	handlers handlerMap
	// mtx guards session state, request registration and terminal closure.
	mtx sync.Mutex
	// connectionCtx lasts until the current GC session ends.
	connectionCtx context.Context
	// connectionCtxCancel interrupts requests when the session ends.
	connectionCtxCancel context.CancelFunc
	// state describes the current coordinator session.
	state state.DeadlockState
	// closed prevents late packets from reopening a disposed client.
	closed bool
	// stopped rejects late welcomes after SetPlaying(false).
	stopped bool
	// nextJobID increases across GC sessions so late replies cannot match new work.
	nextJobID steamprotocol.JobId
	// pending routes each response to its original request's receiving goroutine.
	pending map[steamprotocol.JobId]pendingRequest
}

// New registers a Deadlock handler on a Steam client before it connects.
// The caller must continuously consume client.Events and Close on disconnect.
func New(client *steam.Client, le logrus.FieldLogger) *Deadlock {
	return NewWithCoordinator(client.GC, client.Emit, le)
}

// NewWithCoordinator attaches a Deadlock session to an alternate GC transport.
// emit must continuously accept events; Close ends this handler's lifetime.
func NewWithCoordinator(coordinator Coordinator, emit func(any), le logrus.FieldLogger) *Deadlock {
	// Construct the session before registering it for incoming packets.
	c := &Deadlock{
		le:          le,
		cache:       socache.NewSOCache(),
		coordinator: coordinator,
		emit:        emit,
		pending:     make(map[steamprotocol.JobId]pendingRequest),
		state: state.DeadlockState{
			ConnectionStatus: gcsm.GCConnectionStatus_GCConnectionStatus_NO_SESSION,
		},
	}
	c.buildHandlerMap()
	coordinator.RegisterPacketHandler(c)
	return c
}

// GetCache returns the SO Cache.
func (d *Deadlock) GetCache() *socache.SOCache {
	return d.cache
}

// Close cancels requests and permanently prevents further session admission.
// It does not close the caller's Steam transport or event stream.
func (d *Deadlock) Close() {
	d.mtx.Lock()
	d.closed = true
	d.mtx.Unlock()
	d.setConnectionStatus(gcsm.GCConnectionStatus_GCConnectionStatus_NO_SESSION, nil)
}

// buildHandlerMap builds the map of bound handler functions.
func (d *Deadlock) buildHandlerMap() {
	// Bind protocol messages handled directly by the session.
	d.handlers = handlerMap{
		// Welcome and conn status
		uint32(gcsm.EGCBaseClientMsg_k_EMsgGCClientWelcome):          d.handleClientWelcome,
		uint32(gcsm.EGCBaseClientMsg_k_EMsgGCClientConnectionStatus): d.handleConnectionStatus,

		// Caching
		uint32(gcsm.ESOMsg_k_ESOMsg_CacheSubscribed):   d.handleCacheSubscribed,
		uint32(gcsm.ESOMsg_k_ESOMsg_UpdateMultiple):    d.handleCacheUpdateMultiple,
		uint32(gcsm.ESOMsg_k_ESOMsg_CacheUnsubscribed): d.handleCacheUnsubscribed,
		uint32(gcsm.ESOMsg_k_ESOMsg_Create):            d.cacheSingleHandler(false),
		uint32(gcsm.ESOMsg_k_ESOMsg_Update):            d.cacheSingleHandler(false),
		uint32(gcsm.ESOMsg_k_ESOMsg_Destroy):           d.cacheSingleHandler(true),

		// System events
		uint32(gcsm.EGCBaseClientMsg_k_EMsgGCPingRequest): d.handlePingRequest,

		// Invites
		uint32(gcsm.EGCBaseMsg_k_EMsgGCInvitationCreated): d.getEventEmitter(func() devents.Event {
			return &devents.InvitationCreated{}
		}),

		// Game updates
		uint32(gcsm.EGCItemMsg_k_EMsgGCClientVersionUpdated): d.getEventEmitter(func() devents.Event {
			return &devents.ClientVersionUpdated{}
		}),
	}

	// Add schema-generated notifications after the core handlers.
	d.registerGeneratedHandlers()
}

// write sends a message to the game coordinator.
func (d *Deadlock) write(messageType uint32, msg protobuf.Message) {
	d.coordinator.Write(gamecoordinator.NewGCMsgProtobuf(AppID, messageType, msg))
}

// unmarshalBody attempts to unmarshal a packet body.
func (d *Deadlock) unmarshalBody(packet *gamecoordinator.GCPacket, msg protobuf.Message) (parseErr error) {
	// Keep decode failures attributable to the incoming message type.
	defer func() {
		if parseErr != nil {
			d.le.WithError(parseErr).WithField("msgtype", packet.MsgType).Warn("unable to parse message")
		}
	}()

	// Reused response objects must not retain fields absent from this packet.
	msg.Reset()
	return msg.UnmarshalVT(packet.Body)
}

// HandleGCPacket handles an incoming game coordinator packet.
func (d *Deadlock) HandleGCPacket(packet *gamecoordinator.GCPacket) {
	// A shared Steam transport can deliver packets for other games.
	if packet.AppId != AppID {
		return
	}

	// Ignore packets after the caller disposes this connection's handler.
	d.mtx.Lock()
	closed := d.closed
	d.mtx.Unlock()
	if closed {
		return
	}

	// Decode protocol events before delivering a correlated response.
	le := d.le.WithField("msgtype", packet.MsgType)
	handler, ok := d.handlers[packet.MsgType]
	if ok && handler != nil {
		if err := handler(packet); err != nil {
			le.WithError(err).Warn("error handling gc msg")
			ok = false
		}
	}

	// Surface unsolicited packets without attributing them to another request.
	respHandled := d.handleResponsePacket(packet)
	if !ok && !respHandled {
		le.Debug("unhandled gc packet")
		d.emit(&devents.UnhandledGCPacket{
			Packet: packet,
		})
	}
}

// handlePingRequest handles an incoming ping request from the gc.
func (d *Deadlock) handlePingRequest(packet *gamecoordinator.GCPacket) error {
	d.write(uint32(gcsm.EGCBaseClientMsg_k_EMsgGCPingResponse), &gcsm.CMsgGCClientPing{})
	return nil
}

// getEventEmitter returns a handler that emits an event, used by the generated code.
func (d *Deadlock) getEventEmitter(ctor func() devents.Event) func(packet *gamecoordinator.GCPacket) error {
	return func(packet *gamecoordinator.GCPacket) error {
		// Decode completely before publishing the event.
		obj := ctor()
		if err := d.unmarshalBody(packet, obj.GetEventBody()); err != nil {
			return err
		}

		// Event consumers receive only successfully decoded messages.
		d.emit(obj)
		return nil
	}
}
