package state

import (
	gcmm "github.com/paralin/go-deadlock/protocol"
)

// DeadlockState is a snapshot of the client state at a point in time.
//
// Party, Lobby and HideoutLobby follow the account's shared-object caches,
// which the GC fills after the welcome and clears when the session ends. The
// GC links a party to its match by subscribing each member to the match
// lobby's cache, not by an ID: while the account is in a party and a match,
// Lobby is that party's match lobby and MatchID names the match.
type DeadlockState struct {
	// ConnectionStatus is the status of the connection to the GC.
	ConnectionStatus gcmm.GCConnectionStatus
	// Lobby is the match lobby the account is in.
	Lobby *gcmm.CSOCitadelLobby
	// Party is the party the account is in.
	Party *gcmm.CSOCitadelParty
	// HideoutLobby is the hideout of the account's party; its PartyId names it.
	HideoutLobby *gcmm.CSOCitadelHideoutLobby
	// LastConnectionStatusUpdate is the last connection state update we received.
	LastConnectionStatusUpdate *gcmm.CMsgConnectionStatus
}

// ClearState clears everything.
func (s *DeadlockState) ClearState() {
	*s = DeadlockState{ConnectionStatus: gcmm.GCConnectionStatus_GCConnectionStatus_NO_SESSION}
}

// IsReady checks if the client is ready to receive requests.
func (s *DeadlockState) IsReady() bool {
	return s.ConnectionStatus == gcmm.GCConnectionStatus_GCConnectionStatus_HAVE_SESSION
}

// MatchID returns the ID of the account's match, or zero outside a match.
func (s *DeadlockState) MatchID() uint64 {
	return s.Lobby.GetMatchId()
}
