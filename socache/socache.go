package socache

import (
	"reflect"
	"sync"

	protobuf "github.com/aperturerobotics/protobuf-go-lite"
	"github.com/paralin/go-deadlock/cso"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"

	gcsdkm "github.com/paralin/go-deadlock/protocol"
)

// SOCache implements the shared-object cache from DOTA.
type SOCache struct {
	containerMap sync.Map
	// messageTypes holds caller-verified schemas for this cache only.
	messageTypes sync.Map
	le           logrus.FieldLogger
}

// NewSOCache builds a new SOCache.
func NewSOCache(le logrus.FieldLogger) *SOCache {
	return &SOCache{le: le}
}

// RegisterType associates a verified shared-object type ID with a message schema.
// Register before receiving packets: unsupported objects received earlier are
// not replayed. IDs must come from verified protocol evidence, not guesses.
// The prototype must be a non-nil pointer to a message struct. Only its type is
// retained; each packet is decoded into a new, empty instance of that type.
// Registrations are local to this cache and cannot replace a registered or
// built-in type. RegisterType does not request any network subscriptions.
func (c *SOCache) RegisterType(typ cso.CSOType, prototype protobuf.Message) error {
	if typ <= 0 {
		return errors.New("shared object type id must be positive")
	}
	// Reserve known built-in IDs even if their decoder is not implemented yet.
	if typ == cso.EconItem || typ == cso.ItemRecipe || typ == cso.EconGameAccountClient {
		return errors.Errorf("cannot replace built-in shared object type id: %d", typ)
	}
	t := reflect.TypeOf(prototype)
	if t == nil || t.Kind() != reflect.Pointer || t.Elem().Kind() != reflect.Struct || reflect.ValueOf(prototype).IsNil() {
		return errors.New("shared object prototype must be a non-nil pointer to a message struct")
	}
	if _, loaded := c.messageTypes.LoadOrStore(typ, t); loaded {
		return errors.Errorf("shared object type id already registered: %d", typ)
	}
	return nil
}

// GetContainerForTypeID returns the container for the type id.
func (c *SOCache) GetContainerForTypeID(id uint32) (*SOCacheContainer, error) {
	if valInter, ok := c.containerMap.Load(id); ok {
		return valInter.(*SOCacheContainer), nil
	}

	constructor := func() (protobuf.Message, error) {
		return cso.NewSharedObject(cso.CSOType(id)) //nolint:gosec
	}
	if t, ok := c.messageTypes.Load(cso.CSOType(id)); ok && id <= uint32(1<<31-1) {
		constructor = func() (protobuf.Message, error) {
			return reflect.New(t.(reflect.Type).Elem()).Interface().(protobuf.Message), nil
		}
	}
	ctr, err := newSOCacheContainer(c.le, id, constructor)
	if err != nil {
		return nil, err
	}

	valInter, loaded := c.containerMap.LoadOrStore(id, ctr)
	if loaded {
		ctr = valInter.(*SOCacheContainer)
	}
	return ctr, nil
}

// HandleSubscribed handles a subscribed message.
func (c *SOCache) HandleSubscribed(msg *gcsdkm.CMsgSOCacheSubscribed) error {
	var retErr error
	for _, obj := range msg.GetObjects() {
		if obj.GetTypeId() == 0 {
			continue
		}

		ctr, err := c.GetContainerForTypeID(uint32(obj.GetTypeId())) //nolint:gosec
		if err == nil {
			err = ctr.HandleSubscribed(msg, obj)
		}

		if retErr == nil && err != nil {
			retErr = err
		}
	}

	return retErr
}

// HandleUnsubscribed handles a subscribed message.
func (c *SOCache) HandleUnsubscribed(msg *gcsdkm.CMsgSOCacheUnsubscribed) error {
	if msg.GetOwnerSoid().GetId() == 0 {
		return errors.Errorf("object id was empty in unsubscribed")
	}

	var retErr error
	c.containerMap.Range(func(key any, value any) bool {
		val := value.(*SOCacheContainer)
		_, hasObj := val.Get(msg.GetOwnerSoid().GetId())
		if !hasObj {
			return true
		}

		retErr = val.HandleUnsubscribed(msg)
		return false
	})

	return retErr
}

// HandleUpdateMultiple handles a update multiple message.
func (c *SOCache) HandleUpdateMultiple(msg *gcsdkm.CMsgSOMultipleObjects) error {
	addUpdateObj := func(obj *gcsdkm.CMsgSOMultipleObjects_SingleObject) error {
		ctr, err := c.GetContainerForTypeID(uint32(obj.GetTypeId())) //nolint:gosec
		if err != nil {
			return nil
		}

		m := &gcsdkm.CMsgSOCacheSubscribed_SubscribedType{
			TypeId:     obj.TypeId,
			ObjectData: [][]byte{obj.ObjectData},
		}

		if err := ctr.addUpdateObject(msg.GetOwnerSoid(), m); err != nil {
			return err
		}

		return nil
	}

	for _, obj := range msg.GetObjectsAdded() {
		if err := addUpdateObj(obj); err != nil {
			return err
		}
	}

	for _, obj := range msg.GetObjectsModified() {
		if err := addUpdateObj(obj); err != nil {
			return err
		}
	}

	for _, obj := range msg.GetObjectsRemoved() {
		ctr, err := c.GetContainerForTypeID(uint32(obj.GetTypeId())) //nolint:gosec
		if err != nil {
			continue
		}

		if err := ctr.removeObject(msg.GetOwnerSoid()); err != nil {
			return err
		}
	}

	return nil
}

// HandleDestroy handles a object destroy message.
func (c *SOCache) HandleDestroy(msg *gcsdkm.CMsgSOSingleObject) error {
	ctr, err := c.GetContainerForTypeID(uint32(msg.GetTypeId())) //nolint:gosec
	if err != nil {
		return err
	}

	return ctr.HandleDestroy(msg)
}
