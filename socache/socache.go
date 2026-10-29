package socache

import (
	"errors"
	"fmt"
	"sync"

	protobuf "github.com/aperturerobotics/protobuf-go-lite"
	"github.com/paralin/go-deadlock/cso"
	gcsdkm "github.com/paralin/go-deadlock/protocol"
)

// SOID identifies a cache by the object that owns it, such as an account,
// a party or a lobby.
type SOID struct {
	// Type is the owner's type.
	Type uint32
	// ID is the owner's ID.
	ID uint64
}

// soidOf converts the GC's owner identity.
func soidOf(owner *gcsdkm.CMsgSOIDOwner) SOID {
	return SOID{Type: owner.GetType(), ID: owner.GetId()}
}

// EventType marks the type of event this is.
type EventType int

//go:generate go run golang.org/x/tools/cmd/stringer -type=EventType
const (
	// EventTypeCreate is emitted when an object is inserted into the cache.
	EventTypeCreate EventType = iota
	// EventTypeUpdate is emitted when an object is replaced in the cache.
	EventTypeUpdate
	// EventTypeDestroy is emitted when an object is removed from the cache.
	EventTypeDestroy
)

// CacheEvent reports one object change.
type CacheEvent struct {
	// EventType marks the type of event.
	EventType EventType
	// Owner identifies the cache that holds the object.
	Owner SOID
	// Type is the object's shared-object type.
	Type cso.CSOType
	// Object is the new object, or the removed one for a destroy.
	Object protobuf.Message
}

// SOCache mirrors the shared-object caches the GC subscribes a client to.
// Each cache belongs to an owner and holds its objects by type and key.
//
// Each Handle method applies one GC message and returns the changes it made,
// in order. The objects are immutable once stored, so readers may keep them.
// An object of a type without a schema is skipped.
type SOCache struct {
	// mtx guards registered and caches.
	mtx sync.Mutex
	// registered holds the schemas added with RegisterType.
	registered map[cso.CSOType]cso.Schema
	// caches holds the subscribed caches by owner.
	caches map[SOID]*cache
}

// cache is one owner's subscribed objects.
type cache struct {
	// version is the GC's version of the cache.
	version uint64
	// objects holds the objects by type and key.
	objects map[cso.CSOType]map[cso.Key]protobuf.Message
}

// NewSOCache builds an empty SOCache.
func NewSOCache() *SOCache {
	return &SOCache{
		registered: make(map[cso.CSOType]cso.Schema),
		caches:     make(map[SOID]*cache),
	}
}

// RegisterType decodes a shared-object type that has no built-in schema.
// Objects of the type received before registration are not replayed. It cannot
// replace a built-in or registered schema.
func (c *SOCache) RegisterType(typ cso.CSOType, schema cso.Schema) error {
	if schema.New == nil {
		return fmt.Errorf("shared object type %d: schema has no constructor", typ)
	}
	c.mtx.Lock()
	defer c.mtx.Unlock()
	if _, ok := c.schema(typ); ok {
		return fmt.Errorf("shared object type %d already has a schema", typ)
	}
	c.registered[typ] = schema
	return nil
}

// schema returns the schema of a type, with mtx held.
func (c *SOCache) schema(typ cso.CSOType) (cso.Schema, bool) {
	if s, ok := cso.LookupSchema(typ); ok {
		return s, true
	}
	s, ok := c.registered[typ]
	return s, ok
}

// decode decodes one object with mtx held. It returns false without an error
// for a type that has no schema.
func (c *SOCache) decode(typ cso.CSOType, data []byte) (protobuf.Message, cso.Key, bool, error) {
	s, ok := c.schema(typ)
	if !ok {
		return nil, cso.Key{}, false, nil
	}
	obj := s.New()
	if err := obj.UnmarshalVT(data); err != nil {
		return nil, cso.Key{}, false, fmt.Errorf("decode shared object type %d: %w", typ, err)
	}
	var key cso.Key
	if s.Key != nil {
		key = s.Key(obj)
	}
	return obj, key, true, nil
}

// owned returns the cache of an owner, creating it, with mtx held.
func (c *SOCache) owned(owner SOID) *cache {
	ch := c.caches[owner]
	if ch == nil {
		ch = &cache{objects: make(map[cso.CSOType]map[cso.Key]protobuf.Message)}
		c.caches[owner] = ch
	}
	return ch
}

// put stores an object and returns its change, with mtx held.
func (ch *cache) put(owner SOID, typ cso.CSOType, key cso.Key, obj protobuf.Message) *CacheEvent {
	objs := ch.objects[typ]
	if objs == nil {
		objs = make(map[cso.Key]protobuf.Message)
		ch.objects[typ] = objs
	}
	eventType := EventTypeCreate
	if _, ok := objs[key]; ok {
		eventType = EventTypeUpdate
	}
	objs[key] = obj
	return &CacheEvent{EventType: eventType, Owner: owner, Type: typ, Object: obj}
}

// remove removes an object and returns its change, or nil when the cache does
// not hold it, with mtx held.
func (ch *cache) remove(owner SOID, typ cso.CSOType, key cso.Key) *CacheEvent {
	obj, ok := ch.objects[typ][key]
	if !ok {
		return nil
	}
	delete(ch.objects[typ], key)
	if len(ch.objects[typ]) == 0 {
		delete(ch.objects, typ)
	}
	return &CacheEvent{EventType: EventTypeDestroy, Owner: owner, Type: typ, Object: obj}
}

// destroyAll appends a destroy for each object, with mtx held.
func (ch *cache) destroyAll(owner SOID, events []*CacheEvent) []*CacheEvent {
	for typ, objs := range ch.objects {
		for _, obj := range objs {
			events = append(events, &CacheEvent{EventType: EventTypeDestroy, Owner: owner, Type: typ, Object: obj})
		}
	}
	return events
}

// HandleSubscribed replaces an owner's cache with the full contents the GC
// sent, reporting what was created, updated and destroyed.
func (c *SOCache) HandleSubscribed(msg *gcsdkm.CMsgSOCacheSubscribed) ([]*CacheEvent, error) {
	owner := soidOf(msg.GetOwnerSoid())
	c.mtx.Lock()
	defer c.mtx.Unlock()

	// Store the new contents beside the old, then destroy what they lack.
	prev := c.caches[owner]
	delete(c.caches, owner)
	next := c.owned(owner)
	next.version = msg.GetVersion()
	var events []*CacheEvent
	var errs []error
	for _, set := range msg.GetObjects() {
		typ := cso.CSOType(set.GetTypeId())
		for _, data := range set.GetObjectData() {
			obj, key, ok, err := c.decode(typ, data)
			if err != nil {
				errs = append(errs, err)
			}
			if !ok {
				continue
			}
			event := next.put(owner, typ, key, obj)
			if prev != nil && prev.remove(owner, typ, key) != nil {
				event.EventType = EventTypeUpdate
			}
			events = append(events, event)
		}
	}
	if prev != nil {
		events = prev.destroyAll(owner, events)
	}
	return events, errors.Join(errs...)
}

// HandleUnsubscribed removes an owner's cache.
func (c *SOCache) HandleUnsubscribed(msg *gcsdkm.CMsgSOCacheUnsubscribed) []*CacheEvent {
	owner := soidOf(msg.GetOwnerSoid())
	c.mtx.Lock()
	defer c.mtx.Unlock()
	ch := c.caches[owner]
	if ch == nil {
		return nil
	}
	delete(c.caches, owner)
	return ch.destroyAll(owner, nil)
}

// HandleUpdateMultiple applies the objects the GC added, modified and removed.
func (c *SOCache) HandleUpdateMultiple(msg *gcsdkm.CMsgSOMultipleObjects) ([]*CacheEvent, error) {
	owner := soidOf(msg.GetOwnerSoid())
	c.mtx.Lock()
	defer c.mtx.Unlock()

	ch := c.owned(owner)
	ch.version = msg.GetVersion()
	var events []*CacheEvent
	var errs []error
	apply := func(objs []*gcsdkm.CMsgSOMultipleObjects_SingleObject, removed bool) {
		for _, o := range objs {
			typ := cso.CSOType(o.GetTypeId())
			obj, key, ok, err := c.decode(typ, o.GetObjectData())
			if err != nil {
				errs = append(errs, err)
			}
			switch {
			case !ok:
			case removed:
				if event := ch.remove(owner, typ, key); event != nil {
					events = append(events, event)
				}
			default:
				events = append(events, ch.put(owner, typ, key, obj))
			}
		}
	}
	apply(msg.GetObjectsAdded(), false)
	apply(msg.GetObjectsModified(), false)
	apply(msg.GetObjectsRemoved(), true)
	return events, errors.Join(errs...)
}

// HandleSingle applies one created, updated or, when destroy is set,
// destroyed object.
func (c *SOCache) HandleSingle(msg *gcsdkm.CMsgSOSingleObject, destroy bool) ([]*CacheEvent, error) {
	owner := soidOf(msg.GetOwnerSoid())
	typ := cso.CSOType(msg.GetTypeId())
	c.mtx.Lock()
	defer c.mtx.Unlock()

	obj, key, ok, err := c.decode(typ, msg.GetObjectData())
	if !ok {
		return nil, err
	}
	ch := c.owned(owner)
	ch.version = msg.GetVersion()
	if !destroy {
		return []*CacheEvent{ch.put(owner, typ, key, obj)}, nil
	}
	if event := ch.remove(owner, typ, key); event != nil {
		return []*CacheEvent{event}, nil
	}
	return nil, nil
}

// Reset removes every cache, as when the GC session ends.
func (c *SOCache) Reset() []*CacheEvent {
	c.mtx.Lock()
	defer c.mtx.Unlock()
	var events []*CacheEvent
	for owner, ch := range c.caches {
		events = ch.destroyAll(owner, events)
	}
	clear(c.caches)
	return events
}

// Get returns an owner's object of a type by key. A type without key fields
// uses the zero key.
func (c *SOCache) Get(owner SOID, typ cso.CSOType, key cso.Key) (protobuf.Message, bool) {
	c.mtx.Lock()
	defer c.mtx.Unlock()
	ch := c.caches[owner]
	if ch == nil {
		return nil, false
	}
	obj, ok := ch.objects[typ][key]
	return obj, ok
}

// Range calls cb with every object of a type until cb returns false. It calls
// cb without holding the cache lock, so cb may call back into the cache.
func (c *SOCache) Range(typ cso.CSOType, cb func(owner SOID, obj protobuf.Message) bool) {
	type entry struct {
		owner SOID
		obj   protobuf.Message
	}
	c.mtx.Lock()
	var entries []entry
	for owner, ch := range c.caches {
		for _, obj := range ch.objects[typ] {
			entries = append(entries, entry{owner, obj})
		}
	}
	c.mtx.Unlock()
	for _, e := range entries {
		if !cb(e.owner, e.obj) {
			return
		}
	}
}
