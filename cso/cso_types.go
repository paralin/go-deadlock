package cso

import (
	protobuf "github.com/aperturerobotics/protobuf-go-lite"
	bgcm "github.com/paralin/go-deadlock/protocol"
)

// CSOType is a shared object type identifier.
//
//go:generate go run golang.org/x/tools/cmd/stringer -type=CSOType
type CSOType int32

// The Deadlock IDs are the client's shared-object factory registrations. The
// GC also defines server-only types (102, 106 and 111) that clients never
// receive and whose schemas the protobufs do not include.
const (
	// EconItem is an economy item.
	// This ID is shared across all Valve GC economies.
	EconItem CSOType = 1
	// ItemRecipe is an item recipe.
	ItemRecipe CSOType = 5
	// EconGameAccountClient is an economy game account client.
	EconGameAccountClient CSOType = 7
	// Lobby is the match lobby the account belongs to, CSOCitadelLobby.
	Lobby CSOType = 101
	// GameAccountClient is the account's game profile, CSOGameAccountClient.
	GameAccountClient CSOType = 104
	// Party is the party the account belongs to, CSOCitadelParty.
	Party CSOType = 105
	// AccountHeroInfo is one hero's progress, CSOAccountHeroInfo.
	AccountHeroInfo CSOType = 107
	// AccountSyncStorage is one stored account value, CSOAccountSyncStorage.
	AccountSyncStorage CSOType = 108
	// AccountChallenge is one challenge's progress, CSOAccountChallenge.
	AccountChallenge CSOType = 109
	// HideoutLobby is the party's hideout, CSOCitadelHideoutLobby.
	HideoutLobby CSOType = 110
	// RankedProgress is one ranked interval's progress, CSORankedProgress.
	RankedProgress CSOType = 112
	// AccountCabalMembership is one cabal membership, CSOAccountCabalMembership.
	AccountCabalMembership CSOType = 113
)

// Key identifies a shared object among the objects of its type in one cache.
// It holds the object's key fields in declaration order, zero padded.
type Key [3]uint64

// Schema decodes and identifies the objects of one shared-object type.
type Schema struct {
	// New returns an empty message to decode an object into.
	New func() protobuf.Message
	// Key returns the identity of a decoded object. A nil Key means the type
	// has no key fields and a cache holds at most one object of it.
	Key func(protobuf.Message) Key
}

// message constrains a type parameter to a pointer to a generated message.
type message[T any] interface {
	*T
	protobuf.Message
}

// single returns the schema of a type that has one object per cache.
func single[T any, P message[T]]() Schema {
	return Schema{New: func() protobuf.Message { return P(new(T)) }}
}

// keyed returns the schema of a type whose objects key returns identifies.
func keyed[T any, P message[T]](key func(P) Key) Schema {
	return Schema{
		New: func() protobuf.Message { return P(new(T)) },
		Key: func(obj protobuf.Message) Key { return key(obj.(P)) },
	}
}

// schemas holds the built-in shared-object types.
//
// The keys follow the protobufs' key_field options, which mark only account_id
// on the hero and challenge types and no field on sync storage. A cache holds
// one such object per hero, challenge or stored value, so those keys add the
// field that tells them apart.
var schemas = map[CSOType]Schema{
	EconItem: keyed(func(o *bgcm.CSOEconItem) Key {
		return Key{o.GetId()}
	}),
	EconGameAccountClient: single[bgcm.CSOEconGameAccountClient](),
	Lobby:                 single[bgcm.CSOCitadelLobby](),
	GameAccountClient:     single[bgcm.CSOGameAccountClient](),
	Party:                 single[bgcm.CSOCitadelParty](),
	AccountHeroInfo: keyed(func(o *bgcm.CSOAccountHeroInfo) Key {
		return Key{uint64(o.GetAccountId()), uint64(o.GetHeroId())}
	}),
	AccountSyncStorage: keyed(func(o *bgcm.CSOAccountSyncStorage) Key {
		return Key{uint64(o.GetAccountId()), uint64(o.GetId())}
	}),
	AccountChallenge: keyed(func(o *bgcm.CSOAccountChallenge) Key {
		return Key{uint64(o.GetAccountId()), o.GetChallengeId()}
	}),
	HideoutLobby: single[bgcm.CSOCitadelHideoutLobby](),
	RankedProgress: keyed(func(o *bgcm.CSORankedProgress) Key {
		return Key{uint64(o.GetAccountId()), uint64(o.GetRankType()), uint64(o.GetRankInterval())} //nolint:gosec
	}),
	AccountCabalMembership: keyed(func(o *bgcm.CSOAccountCabalMembership) Key {
		return Key{uint64(o.GetAccountId()), o.GetCabalId()}
	}),
}

// LookupSchema returns the built-in schema for a shared-object type.
func LookupSchema(typ CSOType) (Schema, bool) {
	s, ok := schemas[typ]
	return s, ok
}
