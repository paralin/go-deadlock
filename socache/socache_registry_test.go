package socache

import (
	"io"
	"sync"
	"testing"

	protobuf "github.com/aperturerobotics/protobuf-go-lite"
	"github.com/paralin/go-deadlock/cso"
	gcsm "github.com/paralin/go-deadlock/protocol"
	"github.com/sirupsen/logrus"
)

// These deliberately synthetic IDs are not Deadlock protocol mappings.
const (
	fixturePartyType cso.CSOType = 2147483001
	fixtureLobbyType cso.CSOType = 2147483002
)

func registryCache() *SOCache {
	logger := logrus.New()
	logger.SetOutput(io.Discard)
	return NewSOCache(logger)
}

func TestRegisterTypeRejectsInvalidAndReservedSchemas(t *testing.T) {
	cache := registryCache()
	for _, typ := range []cso.CSOType{-1, 0, cso.EconItem, cso.ItemRecipe, cso.EconGameAccountClient} {
		if err := cache.RegisterType(typ, &gcsm.CSOCitadelParty{}); err == nil {
			t.Fatalf("accepted invalid or reserved ID %d", typ)
		}
	}
	for _, prototype := range []protobuf.Message{nil, (*gcsm.CSOCitadelParty)(nil)} {
		if err := cache.RegisterType(fixturePartyType, prototype); err == nil {
			t.Fatal("accepted nil prototype")
		}
	}
	if err := cache.RegisterType(fixturePartyType, &gcsm.CSOCitadelParty{}); err != nil {
		t.Fatal(err)
	}
	if err := cache.RegisterType(fixturePartyType, &gcsm.CSOCitadelLobby{}); err == nil {
		t.Fatal("replaced an existing schema")
	}
	// Built-in decoding remains available and unchanged.
	if _, err := cache.GetContainerForTypeID(uint32(cso.EconItem)); err != nil {
		t.Fatal(err)
	}
}

func TestRegisterTypeIsCacheLocalAndConcurrent(t *testing.T) {
	cache, other := registryCache(), registryCache()
	var wg sync.WaitGroup
	errors := make(chan error, 16)
	for range 16 {
		wg.Go(func() { errors <- cache.RegisterType(fixturePartyType, &gcsm.CSOCitadelParty{}) })
	}
	wg.Wait()
	close(errors)
	accepted := 0
	for err := range errors {
		if err == nil {
			accepted++
		}
	}
	if accepted != 1 {
		t.Fatalf("registered one ID %d times", accepted)
	}
	if _, err := cache.GetContainerForTypeID(uint32(fixturePartyType)); err != nil {
		t.Fatal(err)
	}
	containers := make(chan *SOCacheContainer, 16)
	for range 16 {
		wg.Go(func() {
			container, err := cache.GetContainerForTypeID(uint32(fixturePartyType))
			if err != nil {
				t.Error(err)
			}
			containers <- container
		})
	}
	wg.Wait()
	close(containers)
	first := <-containers
	for container := range containers {
		if container != first {
			t.Fatal("concurrent readers received different containers")
		}
	}
	if _, err := other.GetContainerForTypeID(uint32(fixturePartyType)); err == nil {
		t.Fatal("registration leaked into another cache")
	}
	if _, err := cso.NewSharedObject(fixturePartyType); err == nil {
		t.Fatal("registration changed the global constructor table")
	}
	// Another cache may explicitly assign the same synthetic ID differently.
	if err := other.RegisterType(fixturePartyType, &gcsm.CSOCitadelLobby{}); err != nil {
		t.Fatal(err)
	}
}

func TestRegisteredPartyAndLobbyLifecycle(t *testing.T) {
	cache := registryCache()
	prototype := &gcsm.CSOCitadelParty{PartyId: new(uint64(111)), JoinCode: new(uint64(777))}
	if err := cache.RegisterType(fixturePartyType, prototype); err != nil {
		t.Fatal(err)
	}
	if err := cache.RegisterType(fixtureLobbyType, &gcsm.CSOCitadelLobby{}); err != nil {
		t.Fatal(err)
	}
	// Mutating a prototype after registration must not seed decoded packets.
	prototype.PartyId = new(uint64(999))
	partyEvents, cancelParty, err := cache.SubscribeType(fixturePartyType)
	if err != nil {
		t.Fatal(err)
	}
	defer cancelParty()
	lobbyEvents, cancelLobby, err := cache.SubscribeType(fixtureLobbyType)
	if err != nil {
		t.Fatal(err)
	}
	defer cancelLobby()
	owner := &gcsm.CMsgSOIDOwner{Id: new(uint64(73))}
	encode := func(message protobuf.Message) []byte {
		t.Helper()
		data, err := message.MarshalVT()
		if err != nil {
			t.Fatal(err)
		}
		return data
	}
	partyID := uint64(9007199254740993)
	party := &gcsm.CSOCitadelParty{PartyId: &partyID, IsPrivateLobby: new(true)}
	lobby := &gcsm.CSOCitadelLobby{LobbyId: new(uint64(9007199254740995)), MatchId: new(uint64(44))}
	if err := cache.HandleSubscribed(&gcsm.CMsgSOCacheSubscribed{
		OwnerSoid: owner,
		Objects: []*gcsm.CMsgSOCacheSubscribed_SubscribedType{
			{TypeId: new(int32(fixturePartyType)), ObjectData: [][]byte{encode(party)}},
			{TypeId: new(int32(fixtureLobbyType)), ObjectData: [][]byte{encode(lobby)}},
		},
	}); err != nil {
		t.Fatal(err)
	}
	first := <-partyEvents
	decoded := first.Object.(*gcsm.CSOCitadelParty)
	if first.EventType != EventTypeCreate || decoded.GetPartyId() != partyID || decoded.GetJoinCode() != 0 || !decoded.GetIsPrivateLobby() {
		t.Fatal("party decode lost identity or copied prototype fields")
	}
	if got := (<-lobbyEvents).Object.(*gcsm.CSOCitadelLobby); got.GetLobbyId() != lobby.GetLobbyId() || got.GetMatchId() != 44 {
		t.Fatal("lobby decode lost identity")
	}
	// Bad protobuf data must not replace or publish the retained good object.
	if err := cache.HandleUpdateMultiple(&gcsm.CMsgSOMultipleObjects{OwnerSoid: owner, ObjectsModified: []*gcsm.CMsgSOMultipleObjects_SingleObject{
		{TypeId: new(int32(fixturePartyType)), ObjectData: []byte{0xff}},
	}}); err == nil {
		t.Fatal("accepted malformed party")
	}
	select {
	case <-partyEvents:
		t.Fatal("published a malformed update")
	default:
	}
	if err := cache.HandleUpdateMultiple(&gcsm.CMsgSOMultipleObjects{OwnerSoid: owner, ObjectsModified: []*gcsm.CMsgSOMultipleObjects_SingleObject{
		{TypeId: new(int32(fixturePartyType)), ObjectData: encode(&gcsm.CSOCitadelParty{PartyId: &partyID})},
	}}); err != nil {
		t.Fatal(err)
	}
	second := <-partyEvents
	updated := second.Object.(*gcsm.CSOCitadelParty)
	if second.EventType != EventTypeUpdate || updated == decoded || updated.GetIsPrivateLobby() || !decoded.GetIsPrivateLobby() {
		t.Fatal("update reused a published message or retained absent fields")
	}
	if err := cache.HandleDestroy(&gcsm.CMsgSOSingleObject{OwnerSoid: owner, TypeId: new(int32(fixturePartyType))}); err != nil {
		t.Fatal(err)
	}
	if event := <-partyEvents; event.EventType != EventTypeDestroy || event.Object != updated {
		t.Fatal("destroy lost its last decoded object")
	}
	container, _ := cache.GetContainerForTypeID(uint32(fixturePartyType))
	if container.GetOne() != nil {
		t.Fatal("destroy retained party")
	}
}
