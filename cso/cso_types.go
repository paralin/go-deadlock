package cso

import (
	protobuf "github.com/aperturerobotics/protobuf-go-lite"
	bgcm "github.com/paralin/go-deadlock/protocol"
	"github.com/pkg/errors"
)

// CSOType is a shared object type identifier.
//
//go:generate stringer -type=CSOType
type CSOType int32

const (
	// EconItem is an economy item.
	// This ID is shared across all Valve GC economies.
	EconItem CSOType = 1
	// ItemRecipe is an item recipe.
	ItemRecipe = 5
	// EconGameAccountClient is an economy game account client.
	EconGameAccountClient = 7
)

// TODO: the cache type IDs for the Deadlock-specific shared objects
// (CSOGameAccountClient, CSOAccountHeroInfo, CSOCitadelLobby, CSOCitadelParty,
// etc.) are assigned by the GC and not present in the protobufs. Capture them
// from an authorized session before adding built-in mappings. Callers with
// independently verified IDs can use SOCache.RegisterType for cache-local
// decoding without changing this process-wide constructor table.

// csoTypeCtors links type IDs to constructors.
var csoTypeCtors = map[CSOType]func() protobuf.Message{
	EconItem: func() protobuf.Message {
		return &bgcm.CSOEconItem{}
	},
	EconGameAccountClient: func() protobuf.Message {
		return &bgcm.CSOEconGameAccountClient{}
	},
}

// NewSharedObject builds a new shared object from a type ID.
func NewSharedObject(typ CSOType) (protobuf.Message, error) {
	ctor, ok := csoTypeCtors[typ]
	if !ok {
		return nil, errors.Errorf("unknown shared object type id: %d", typ)
	}

	return ctor(), nil
}
