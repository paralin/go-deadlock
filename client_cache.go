package deadlock

import (
	devents "github.com/paralin/go-deadlock/events"
	gcsm "github.com/paralin/go-deadlock/protocol"
	"github.com/paralin/go-deadlock/socache"
	"github.com/paralin/go-deadlock/state"
	"github.com/paralin/go-steam/protocol/gamecoordinator"
)

// RequestCacheSubscriptionRefresh requests a subscription refresh for a specific cache ID.
func (d *Deadlock) RequestCacheSubscriptionRefresh(ownerSoid *gcsm.CMsgSOIDOwner) {
	d.write(uint32(gcsm.ESOMsg_k_ESOMsg_CacheSubscriptionRefresh), &gcsm.CMsgSOCacheSubscriptionRefresh{
		OwnerSoid: ownerSoid,
	})
}

// applyCache applies a cache change during a session, updates the state that
// follows the cache, and emits each object change and then the new state.
// Holding mtx keeps a change from landing after the session's reset.
func (d *Deadlock) applyCache(apply func(*socache.SOCache) ([]*socache.CacheEvent, error)) {
	d.mtx.Lock()
	if !d.state.IsReady() {
		d.mtx.Unlock()
		return
	}
	changes, err := apply(d.cache)
	previous := d.state
	for _, change := range changes {
		trackCacheObject(&d.state, change)
	}
	next := d.state
	d.mtx.Unlock()

	// A bad object skips only itself; the rest of the message applied.
	if err != nil {
		d.le.WithError(err).Warn("unable to apply shared objects")
	}
	for _, change := range changes {
		d.emit(change)
	}
	if previous != next {
		d.emit(devents.ClientStateChanged{OldState: previous, NewState: next})
	}
}

// trackCacheObject updates the state objects that follow a cache change.
func trackCacheObject(s *state.DeadlockState, change *socache.CacheEvent) {
	destroy := change.EventType == socache.EventTypeDestroy
	switch obj := change.Object.(type) {
	case *gcsm.CSOCitadelParty:
		s.Party = track(s.Party, obj, destroy)
	case *gcsm.CSOCitadelLobby:
		s.Lobby = track(s.Lobby, obj, destroy)
	case *gcsm.CSOCitadelHideoutLobby:
		s.HideoutLobby = track(s.HideoutLobby, obj, destroy)
	}
}

// track returns obj after it is stored, and nil after the current object is
// destroyed; destroying another object leaves current in place.
func track[T comparable](current, obj T, destroy bool) T {
	var zero T
	switch {
	case !destroy:
		return obj
	case current == obj:
		return zero
	default:
		return current
	}
}

// handleCacheSubscribed handles a CacheSubscribed packet.
func (d *Deadlock) handleCacheSubscribed(packet *gamecoordinator.GCPacket) error {
	msg := &gcsm.CMsgSOCacheSubscribed{}
	if err := d.unmarshalBody(packet, msg); err != nil {
		return err
	}
	d.applyCache(func(c *socache.SOCache) ([]*socache.CacheEvent, error) {
		return c.HandleSubscribed(msg)
	})
	return nil
}

// handleCacheUnsubscribed handles a CacheUnsubscribed packet.
func (d *Deadlock) handleCacheUnsubscribed(packet *gamecoordinator.GCPacket) error {
	msg := &gcsm.CMsgSOCacheUnsubscribed{}
	if err := d.unmarshalBody(packet, msg); err != nil {
		return err
	}
	d.applyCache(func(c *socache.SOCache) ([]*socache.CacheEvent, error) {
		return c.HandleUnsubscribed(msg), nil
	})
	return nil
}

// handleCacheUpdateMultiple handles when one or more object(s) in a cache is/are updated.
func (d *Deadlock) handleCacheUpdateMultiple(packet *gamecoordinator.GCPacket) error {
	msg := &gcsm.CMsgSOMultipleObjects{}
	if err := d.unmarshalBody(packet, msg); err != nil {
		return err
	}
	d.applyCache(func(c *socache.SOCache) ([]*socache.CacheEvent, error) {
		return c.HandleUpdateMultiple(msg)
	})
	return nil
}

// cacheSingleHandler handles a Create, Update or, with destroy set, Destroy packet.
func (d *Deadlock) cacheSingleHandler(destroy bool) func(*gamecoordinator.GCPacket) error {
	return func(packet *gamecoordinator.GCPacket) error {
		msg := &gcsm.CMsgSOSingleObject{}
		if err := d.unmarshalBody(packet, msg); err != nil {
			return err
		}
		d.applyCache(func(c *socache.SOCache) ([]*socache.CacheEvent, error) {
			return c.HandleSingle(msg, destroy)
		})
		return nil
	}
}
