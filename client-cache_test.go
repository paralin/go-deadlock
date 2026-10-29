package deadlock

import (
	"io"
	"testing"

	protobuf "github.com/aperturerobotics/protobuf-go-lite"
	"github.com/paralin/go-deadlock/cso"
	devents "github.com/paralin/go-deadlock/events"
	gcsm "github.com/paralin/go-deadlock/protocol"
	"github.com/paralin/go-deadlock/socache"
	"github.com/paralin/go-steam/protocol/gamecoordinator"
	"github.com/sirupsen/logrus"
)

// TestCacheSession walks a party and its match lobby through the welcome,
// updates, an unsubscribe and the end of the session, checking the cache, its
// events and the state that follows it.
func TestCacheSession(t *testing.T) {
	var emitted []any
	logger := logrus.New()
	logger.SetOutput(io.Discard)
	client := NewWithCoordinator(&requestCoordinator{writes: make(chan gamecoordinator.IGCMsg, 8)}, func(event any) { emitted = append(emitted, event) }, logger)
	t.Cleanup(client.Close)
	cache := client.GetCache()

	encode := func(m protobuf.Message) []byte {
		data, err := m.MarshalVT()
		if err != nil {
			t.Fatal(err)
		}
		return data
	}
	subscribed := func(owner *gcsm.CMsgSOIDOwner, typ cso.CSOType, objs ...protobuf.Message) *gcsm.CMsgSOCacheSubscribed {
		set := &gcsm.CMsgSOCacheSubscribed_SubscribedType{TypeId: new(int32(typ))}
		for _, obj := range objs {
			set.ObjectData = append(set.ObjectData, encode(obj))
		}
		return &gcsm.CMsgSOCacheSubscribed{OwnerSoid: owner, Objects: []*gcsm.CMsgSOCacheSubscribed_SubscribedType{set}}
	}
	deliver := func(msgType gcsm.ESOMsg, body protobuf.Message) {
		emitted = nil
		client.HandleGCPacket(&gamecoordinator.GCPacket{AppId: AppID, MsgType: uint32(msgType), Body: encode(body)})
	}
	changes := func() (out []*socache.CacheEvent) {
		for _, e := range emitted {
			if change, ok := e.(*socache.CacheEvent); ok {
				out = append(out, change)
			}
		}
		return out
	}
	state := func() stateView {
		client.mtx.Lock()
		defer client.mtx.Unlock()
		return stateView{client.state.Party.GetPartyId(), client.state.Lobby.GetLobbyId(), client.state.MatchID()}
	}

	// Objects that arrive before the welcome have no session to belong to.
	party := &gcsm.CMsgSOIDOwner{Type: new(uint32(3)), Id: new(uint64(9007199254740993))}
	lobby := &gcsm.CMsgSOIDOwner{Type: new(uint32(4)), Id: new(uint64(7))}
	account := &gcsm.CMsgSOIDOwner{Type: new(uint32(1)), Id: new(uint64(73))}
	deliver(gcsm.ESOMsg_k_ESOMsg_CacheSubscribed, subscribed(party, cso.Party, &gcsm.CSOCitadelParty{PartyId: new(uint64(9007199254740993))}))
	if len(changes()) != 0 || state().Party != 0 {
		t.Fatal("cache applied an object outside a session")
	}

	// The welcome admits the session and fills the party and lobby state.
	emitted = nil
	welcome := encode(&gcsm.CMsgClientWelcome{OutofdateSubscribedCaches: []*gcsm.CMsgSOCacheSubscribed{
		subscribed(party, cso.Party, &gcsm.CSOCitadelParty{PartyId: new(uint64(9007199254740993))}),
		subscribed(lobby, cso.Lobby, &gcsm.CSOCitadelLobby{LobbyId: new(uint64(7)), MatchId: new(uint64(41))}),
	}})
	client.HandleGCPacket(&gamecoordinator.GCPacket{AppId: AppID, MsgType: uint32(gcsm.EGCBaseClientMsg_k_EMsgGCClientWelcome), Body: welcome})
	if got := state(); got != (stateView{9007199254740993, 7, 41}) {
		t.Fatalf("state after welcome %+v", got)
	}
	if _, ok := emitted[len(emitted)-1].(*devents.ClientWelcomed); !ok || len(changes()) != 2 {
		t.Fatalf("welcome emitted %#v", emitted)
	}

	// A full subscription keeps every object, keyed by its key fields, and
	// unknown types are skipped.
	heroes := subscribed(account, cso.AccountHeroInfo,
		&gcsm.CSOAccountHeroInfo{AccountId: new(uint32(73)), HeroId: new(uint32(1))},
		&gcsm.CSOAccountHeroInfo{AccountId: new(uint32(73)), HeroId: new(uint32(2))})
	heroes.Objects = append(heroes.Objects, &gcsm.CMsgSOCacheSubscribed_SubscribedType{TypeId: new(int32(999)), ObjectData: [][]byte{{1}}})
	deliver(gcsm.ESOMsg_k_ESOMsg_CacheSubscribed, heroes)
	if _, ok := cache.Get(socache.SOID{Type: 1, ID: 73}, cso.AccountHeroInfo, cso.Key{73, 2}); !ok || len(changes()) != 2 {
		t.Fatal("subscription lost a hero")
	}

	// Resubscribing replaces the set: a kept hero updates, a missing one goes.
	deliver(gcsm.ESOMsg_k_ESOMsg_CacheSubscribed, subscribed(account, cso.AccountHeroInfo,
		&gcsm.CSOAccountHeroInfo{AccountId: new(uint32(73)), HeroId: new(uint32(2)), Wins: new(uint32(5))}))
	if got := changes(); len(got) != 2 || got[0].EventType != socache.EventTypeUpdate || got[1].EventType != socache.EventTypeDestroy {
		t.Fatalf("resubscribe changes %+v", got)
	}

	// The match starting updates the lobby and its match ID.
	deliver(gcsm.ESOMsg_k_ESOMsg_UpdateMultiple, &gcsm.CMsgSOMultipleObjects{OwnerSoid: lobby, ObjectsModified: []*gcsm.CMsgSOMultipleObjects_SingleObject{
		{TypeId: new(int32(cso.Lobby)), ObjectData: encode(&gcsm.CSOCitadelLobby{LobbyId: new(uint64(7)), MatchId: new(uint64(42))})},
	}})
	if state().MatchID != 42 {
		t.Fatal("lobby update did not reach the state")
	}

	// Leaving the lobby clears only the lobby.
	deliver(gcsm.ESOMsg_k_ESOMsg_CacheUnsubscribed, &gcsm.CMsgSOCacheUnsubscribed{OwnerSoid: lobby})
	if got := state(); got != (stateView{9007199254740993, 0, 0}) {
		t.Fatalf("state after leaving the lobby %+v", got)
	}
	if _, ok := emitted[len(emitted)-1].(devents.ClientStateChanged); !ok {
		t.Fatal("leaving the lobby did not report the new state")
	}

	// The end of the session destroys every cached object.
	emitted = nil
	client.SetPlaying(false)
	if got := changes(); len(got) != 2 || state().Party != 0 {
		t.Fatalf("session end left %d changes, party %d", len(got), state().Party)
	}
	if _, ok := cache.Get(socache.SOID{Type: 1, ID: 73}, cso.AccountHeroInfo, cso.Key{73, 2}); ok {
		t.Fatal("session end kept a hero")
	}
}

// stateView is the part of the state TestCacheSession compares.
type stateView struct {
	Party, Lobby, MatchID uint64
}
