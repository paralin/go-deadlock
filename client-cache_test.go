package deadlock

import (
	"io"
	"testing"

	"github.com/paralin/go-deadlock/cso"
	gcsm "github.com/paralin/go-deadlock/protocol"
	"github.com/paralin/go-steam/protocol/gamecoordinator"
	"github.com/sirupsen/logrus"
)

func TestRegisteredPartyDecoderThroughGCPacket(t *testing.T) {
	// Synthetic mapping only; this test does not assert a live protocol type ID.
	const fixturePartyType cso.CSOType = 2147483001
	transport := &requestCoordinator{writes: make(chan gamecoordinator.IGCMsg, 8)}
	logger := logrus.New()
	logger.SetOutput(io.Discard)
	client := NewWithCoordinator(transport, func(any) {}, logger)
	defer client.Close()
	cache := client.GetCache()
	if err := cache.RegisterType(fixturePartyType, &gcsm.CSOCitadelParty{}); err != nil {
		t.Fatal(err)
	}
	stream, cancel, err := cache.SubscribeType(fixturePartyType)
	if err != nil {
		t.Fatal(err)
	}
	defer cancel()
	partyData, err := (&gcsm.CSOCitadelParty{PartyId: new(uint64(9007199254740993)), IsPrivateLobby: new(true)}).MarshalVT()
	if err != nil {
		t.Fatal(err)
	}
	body, err := (&gcsm.CMsgSOCacheSubscribed{
		OwnerSoid: &gcsm.CMsgSOIDOwner{Id: new(uint64(73))},
		Objects: []*gcsm.CMsgSOCacheSubscribed_SubscribedType{
			{TypeId: new(int32(fixturePartyType)), ObjectData: [][]byte{partyData}},
		},
	}).MarshalVT()
	if err != nil {
		t.Fatal(err)
	}
	packet := &gamecoordinator.GCPacket{AppId: AppID, MsgType: uint32(gcsm.ESOMsg_k_ESOMsg_CacheSubscribed), Body: body}
	client.HandleGCPacket(packet)
	select {
	case event := <-stream:
		party := event.Object.(*gcsm.CSOCitadelParty)
		if party.GetPartyId() != 9007199254740993 || !party.GetIsPrivateLobby() {
			t.Fatal("GC handler did not preserve the configured schema")
		}
	default:
		t.Fatal("GC handler did not emit decoded party")
	}
	// Local registration/observation must not send any GC requests or joins.
	select {
	case <-transport.writes:
		t.Fatal("decoder setup or packet handling initiated a request")
	default:
	}
}
