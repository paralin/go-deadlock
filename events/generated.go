package events

import (
	protobuf_go_lite "github.com/aperturerobotics/protobuf-go-lite"
	"github.com/paralin/go-deadlock/protocol"
)

// BookUpdated is emitted when the GC pushes updated hero book state, such as
// after a book purchase.
type BookUpdated struct {
	protocol.CMsgGCToClientBookUpdated
}

// GetDeadlockEventMsgID returns the deadlock message ID of the event.
func (e *BookUpdated) GetDeadlockEventMsgID() protocol.EGCCitadelClientMessages {
	return protocol.EGCCitadelClientMessages_k_EMsgGCToClientBookUpdated
}

// GetEventBody returns the event body.
func (e *BookUpdated) GetEventBody() protobuf_go_lite.Message {
	return &e.CMsgGCToClientBookUpdated
}

// GetEventName returns the event name.
func (e *BookUpdated) GetEventName() string {
	return "BookUpdated"
}

// CabalUpdated is an event delivered by the GC.
//
// Message: k_EMsgGCToClientCabalUpdated (CMsgGCToClientCabalUpdated).
type CabalUpdated struct {
	protocol.CMsgGCToClientCabalUpdated
}

// GetDeadlockEventMsgID returns the deadlock message ID of the event.
func (e *CabalUpdated) GetDeadlockEventMsgID() protocol.EGCCitadelClientMessages {
	return protocol.EGCCitadelClientMessages_k_EMsgGCToClientCabalUpdated
}

// GetEventBody returns the event body.
func (e *CabalUpdated) GetEventBody() protobuf_go_lite.Message {
	return &e.CMsgGCToClientCabalUpdated
}

// GetEventName returns the event name.
func (e *CabalUpdated) GetEventName() string {
	return "CabalUpdated"
}

// CanRejoinParty is emitted when the GC determines your account may rejoin a
// party it was disconnected from.
type CanRejoinParty struct {
	protocol.CMsgGCToClientCanRejoinParty
}

// GetDeadlockEventMsgID returns the deadlock message ID of the event.
func (e *CanRejoinParty) GetDeadlockEventMsgID() protocol.EGCCitadelClientMessages {
	return protocol.EGCCitadelClientMessages_k_EMsgGCToClientCanRejoinParty
}

// GetEventBody returns the event body.
func (e *CanRejoinParty) GetEventBody() protobuf_go_lite.Message {
	return &e.CMsgGCToClientCanRejoinParty
}

// GetEventName returns the event name.
func (e *CanRejoinParty) GetEventName() string {
	return "CanRejoinParty"
}

// CommendNotification is emitted when another player commends your account.
type CommendNotification struct {
	protocol.CMsgGCToClientCommendNotification
}

// GetDeadlockEventMsgID returns the deadlock message ID of the event.
func (e *CommendNotification) GetDeadlockEventMsgID() protocol.EGCCitadelClientMessages {
	return protocol.EGCCitadelClientMessages_k_EMsgGCToClientCommendNotification
}

// GetEventBody returns the event body.
func (e *CommendNotification) GetEventBody() protobuf_go_lite.Message {
	return &e.CMsgGCToClientCommendNotification
}

// GetEventName returns the event name.
func (e *CommendNotification) GetEventName() string {
	return "CommendNotification"
}

// DevPlaytestStatus is an event delivered by the GC.
//
// Message: k_EMsgGCToClientDevPlaytestStatus (CMsgGCToClientDevPlaytestStatus).
type DevPlaytestStatus struct {
	protocol.CMsgGCToClientDevPlaytestStatus
}

// GetDeadlockEventMsgID returns the deadlock message ID of the event.
func (e *DevPlaytestStatus) GetDeadlockEventMsgID() protocol.EGCCitadelClientMessages {
	return protocol.EGCCitadelClientMessages_k_EMsgGCToClientDevPlaytestStatus
}

// GetEventBody returns the event body.
func (e *DevPlaytestStatus) GetEventBody() protobuf_go_lite.Message {
	return &e.CMsgGCToClientDevPlaytestStatus
}

// GetEventName returns the event name.
func (e *DevPlaytestStatus) GetEventName() string {
	return "DevPlaytestStatus"
}

// MatchmakingStopped is emitted when the GC removes the account from the
// matchmaking queue, either because the client requested it or because the GC
// aborted the search.
type MatchmakingStopped struct {
	protocol.CMsgGCToClientMatchmakingStopped
}

// GetDeadlockEventMsgID returns the deadlock message ID of the event.
func (e *MatchmakingStopped) GetDeadlockEventMsgID() protocol.EGCCitadelClientMessages {
	return protocol.EGCCitadelClientMessages_k_EMsgGCToClientMatchmakingStopped
}

// GetEventBody returns the event body.
func (e *MatchmakingStopped) GetEventBody() protobuf_go_lite.Message {
	return &e.CMsgGCToClientMatchmakingStopped
}

// GetEventName returns the event name.
func (e *MatchmakingStopped) GetEventName() string {
	return "MatchmakingStopped"
}

// PartyChatMsg is an event delivered by the GC.
//
// Message: k_EMsgGCToClientPartyChatMsg (CMsgGCToClientPartyChatMsg).
type PartyChatMsg struct {
	protocol.CMsgGCToClientPartyChatMsg
}

// GetDeadlockEventMsgID returns the deadlock message ID of the event.
func (e *PartyChatMsg) GetDeadlockEventMsgID() protocol.EGCCitadelClientMessages {
	return protocol.EGCCitadelClientMessages_k_EMsgGCToClientPartyChatMsg
}

// GetEventBody returns the event body.
func (e *PartyChatMsg) GetEventBody() protobuf_go_lite.Message {
	return &e.CMsgGCToClientPartyChatMsg
}

// GetEventName returns the event name.
func (e *PartyChatMsg) GetEventName() string {
	return "PartyChatMsg"
}

// PartyEvent is emitted when something in your party changes: membership,
// readiness, or match state.
type PartyEvent struct {
	protocol.CMsgGCToClientPartyEvent
}

// GetDeadlockEventMsgID returns the deadlock message ID of the event.
func (e *PartyEvent) GetDeadlockEventMsgID() protocol.EGCCitadelClientMessages {
	return protocol.EGCCitadelClientMessages_k_EMsgGCToClientPartyEvent
}

// GetEventBody returns the event body.
func (e *PartyEvent) GetEventBody() protobuf_go_lite.Message {
	return &e.CMsgGCToClientPartyEvent
}

// GetEventName returns the event name.
func (e *PartyEvent) GetEventName() string {
	return "PartyEvent"
}

// PartyPlayerTyping is an event delivered by the GC.
//
// Message: k_EMsgGCToClientPartyPlayerTyping (CMsgGCToClientPartyPlayerTyping).
type PartyPlayerTyping struct {
	protocol.CMsgGCToClientPartyPlayerTyping
}

// GetDeadlockEventMsgID returns the deadlock message ID of the event.
func (e *PartyPlayerTyping) GetDeadlockEventMsgID() protocol.EGCCitadelClientMessages {
	return protocol.EGCCitadelClientMessages_k_EMsgGCToClientPartyPlayerTyping
}

// GetEventBody returns the event body.
func (e *PartyPlayerTyping) GetEventBody() protobuf_go_lite.Message {
	return &e.CMsgGCToClientPartyPlayerTyping
}

// GetEventName returns the event name.
func (e *PartyPlayerTyping) GetEventName() string {
	return "PartyPlayerTyping"
}

// ProfileCardUpdated is emitted with the profile card of an account, usually
// in reply to GetProfileCard.
type ProfileCardUpdated struct {
	protocol.CMsgCitadelProfileCard
}

// GetDeadlockEventMsgID returns the deadlock message ID of the event.
func (e *ProfileCardUpdated) GetDeadlockEventMsgID() protocol.EGCCitadelClientMessages {
	return protocol.EGCCitadelClientMessages_k_EMsgGCToClientProfileCardUpdated
}

// GetEventBody returns the event body.
func (e *ProfileCardUpdated) GetEventBody() protobuf_go_lite.Message {
	return &e.CMsgCitadelProfileCard
}

// GetEventName returns the event name.
func (e *ProfileCardUpdated) GetEventName() string {
	return "ProfileCardUpdated"
}

// RankDataUpdate is an event delivered by the GC.
//
// Message: k_EMsgGCToClientRankDataUpdate (CMsgGCToClientRankDataUpdate).
type RankDataUpdate struct {
	protocol.CMsgGCToClientRankDataUpdate
}

// GetDeadlockEventMsgID returns the deadlock message ID of the event.
func (e *RankDataUpdate) GetDeadlockEventMsgID() protocol.EGCCitadelClientMessages {
	return protocol.EGCCitadelClientMessages_k_EMsgGCToClientRankDataUpdate
}

// GetEventBody returns the event body.
func (e *RankDataUpdate) GetEventBody() protobuf_go_lite.Message {
	return &e.CMsgGCToClientRankDataUpdate
}

// GetEventName returns the event name.
func (e *RankDataUpdate) GetEventName() string {
	return "RankDataUpdate"
}

// SDRTicket is emitted with a Steam Datagram Relay ticket for connecting to
// a match or hideout server.
type SDRTicket struct {
	protocol.CMsgGCToClientSDRTicket
}

// GetDeadlockEventMsgID returns the deadlock message ID of the event.
func (e *SDRTicket) GetDeadlockEventMsgID() protocol.EGCCitadelClientMessages {
	return protocol.EGCCitadelClientMessages_k_EMsgGCToClientSDRTicket
}

// GetEventBody returns the event body.
func (e *SDRTicket) GetEventBody() protobuf_go_lite.Message {
	return &e.CMsgGCToClientSDRTicket
}

// GetEventName returns the event name.
func (e *SDRTicket) GetEventName() string {
	return "SDRTicket"
}

// UpdateHeroReleaseVoteTally is an event delivered by the GC.
//
// Message: k_EMsgGCToClientUpdateHeroReleaseVoteTally (CMsgGCToClientUpdateHeroReleaseVoteTally).
type UpdateHeroReleaseVoteTally struct {
	protocol.CMsgGCToClientUpdateHeroReleaseVoteTally
}

// GetDeadlockEventMsgID returns the deadlock message ID of the event.
func (e *UpdateHeroReleaseVoteTally) GetDeadlockEventMsgID() protocol.EGCCitadelClientMessages {
	return protocol.EGCCitadelClientMessages_k_EMsgGCToClientUpdateHeroReleaseVoteTally
}

// GetEventBody returns the event body.
func (e *UpdateHeroReleaseVoteTally) GetEventBody() protobuf_go_lite.Message {
	return &e.CMsgGCToClientUpdateHeroReleaseVoteTally
}

// GetEventName returns the event name.
func (e *UpdateHeroReleaseVoteTally) GetEventName() string {
	return "UpdateHeroReleaseVoteTally"
}
