package main

import (
	protobuf "github.com/aperturerobotics/protobuf-go-lite"

	dm "github.com/paralin/go-deadlock/protocol"
)

// The deadlock protos deviate from the Valve naming conventions in a few
// places. These maps record each deviation so the apigen heuristics can
// resolve every message ID.

// msgSenderOverrides overrides the heuristic-generated sender parties for each
// message. Messages with no protobuf definition are marked MsgSenderNone.
var msgSenderOverrides = map[dm.EGCCitadelClientMessages]MsgSender{
	dm.EGCCitadelClientMessages_k_EMsgGCToClientAccountStatsUpdated:               MsgSenderNone,
	dm.EGCCitadelClientMessages_k_EMsgGCToClientHeroLabsSchedule:                  MsgSenderNone,
	dm.EGCCitadelClientMessages_k_EMsgClientToGCDevBan:                            MsgSenderNone,
	dm.EGCCitadelClientMessages_k_EMsgClientToGCDevBanResponse:                    MsgSenderNone,
	dm.EGCCitadelClientMessages_k_EMsgClientToGCTarotGetUserData:                  MsgSenderNone,
	dm.EGCCitadelClientMessages_k_EMsgClientToGCTarotGetUserDataResponse:          MsgSenderNone,
	dm.EGCCitadelClientMessages_k_EMsgGCToClientTarotUserDataUpdated:              MsgSenderNone,
	dm.EGCCitadelClientMessages_k_EMsgClientToGCTarotPlaceTile:                    MsgSenderNone,
	dm.EGCCitadelClientMessages_k_EMsgClientToGCTarotPlaceTileResponse:            MsgSenderNone,
	dm.EGCCitadelClientMessages_k_EMsgClientToGCTarotApplyRune:                    MsgSenderNone,
	dm.EGCCitadelClientMessages_k_EMsgClientToGCTarotApplyRuneResponse:            MsgSenderNone,
	dm.EGCCitadelClientMessages_k_EMsgClientToGCTarotCompleteBoard:                MsgSenderNone,
	dm.EGCCitadelClientMessages_k_EMsgClientToGCTarotCompleteBoardResponse:        MsgSenderNone,
	dm.EGCCitadelClientMessages_k_EMsgClientToGCTarotDevCreateTile:                MsgSenderNone,
	dm.EGCCitadelClientMessages_k_EMsgClientToGCTarotDevCreateTileResponse:        MsgSenderNone,
	dm.EGCCitadelClientMessages_k_EMsgClientToGCTarotDevCreateBoard:               MsgSenderNone,
	dm.EGCCitadelClientMessages_k_EMsgClientToGCTarotDevCreateBoardResponse:       MsgSenderNone,
	dm.EGCCitadelClientMessages_k_EMsgClientToGCTarotDevGenerateTestTiles:         MsgSenderNone,
	dm.EGCCitadelClientMessages_k_EMsgClientToGCTarotDevGenerateTestTilesResponse: MsgSenderNone,
	dm.EGCCitadelClientMessages_k_EMsgClientToGCTarotDevModifyEnergy:              MsgSenderNone,
	dm.EGCCitadelClientMessages_k_EMsgClientToGCTarotDevModifyEnergyResponse:      MsgSenderNone,
	dm.EGCCitadelClientMessages_k_EMsgClientToGCPartyPlayerTyping:                 MsgSenderNone,
}

// msgMethodNameOverrides overrides the generated method name for a message.
var msgMethodNameOverrides = map[dm.EGCCitadelClientMessages]string{}

// msgResponseOverrides overrides the response message ID for a request message.
var msgResponseOverrides = map[dm.EGCCitadelClientMessages]dm.EGCCitadelClientMessages{
	// k_EMsgClientToGCGetProfileCardResponse has no protobuf definition in the
	// deadlock protos, so treat GetProfileCard as fire-and-forget.
	dm.EGCCitadelClientMessages_k_EMsgClientToGCGetProfileCard: 0,
}

// msgProtoTypeOverrides overrides the proto type used for a message ID.
var msgProtoTypeOverrides = map[dm.EGCCitadelClientMessages]protobuf.Message{
	// The profile card update reuses the profile card message.
	dm.EGCCitadelClientMessages_k_EMsgGCToClientProfileCardUpdated: &dm.CMsgCitadelProfileCard{},
	// The cheat report messages drop their Dev prefix in the proto names.
	dm.EGCCitadelClientMessages_k_EMsgClientToGCDevRequestCheatReports:         &dm.CMsgClientToGCRequestCheatReports{},
	dm.EGCCitadelClientMessages_k_EMsgClientToGCDevRequestCheatReportsResponse: &dm.CMsgClientToGCRequestCheatReportsResponse{},
}

// msgArgAsParameterOverrides forces a message argument to be passed as a parameter.
var msgArgAsParameterOverrides = map[dm.EGCCitadelClientMessages]bool{}

// msgEventNameOverrides overrides the generated event name for a message.
var msgEventNameOverrides = map[dm.EGCCitadelClientMessages]string{}

// msgDocOverrides contains handwritten godoc comments for request methods.
// An entry replaces the entire generated comment for the method.
var msgDocOverrides = map[dm.EGCCitadelClientMessages]string{
	dm.EGCCitadelClientMessages_k_EMsgClientToGCChallengeCancel: `CancelChallenge cancels one of your active challenges.

The provided challenge must belong to your account. The response reports the
result of the cancellation.`,
	dm.EGCCitadelClientMessages_k_EMsgClientToGCChallengeClaim: `ClaimChallenge claims the reward for a completed challenge.

Only challenges that have reached the completed state can be claimed. The
response reports the result of the claim.`,
	dm.EGCCitadelClientMessages_k_EMsgClientToGCPartyCreate: `CreateParty creates a new party with the given matchmaking options.

The party join code is returned in the response, unless party codes are
disabled. The created party is also delivered through the SOCache as a
CSOCitadelParty object.`,
	dm.EGCCitadelClientMessages_k_EMsgClientToGCFindHeroBuilds: `FindHeroBuilds searches published hero builds.

At most one filter is meaningful per call: builds by author_account_id, for
hero_id, or matching search_text. The response lists the matching builds with
their metadata.`,
	dm.EGCCitadelClientMessages_k_EMsgClientToGCGetAccountLeaderboards: `GetAccountLeaderboards returns the leaderboard placements of an account,
one entry per ranked hero leaderboard the account appears on.`,
	dm.EGCCitadelClientMessages_k_EMsgClientToGCGetAccountMatchReports: `GetAccountMatchReports fetches the conduct report information tied to a
match. Supply the most recent known match_id to page through older entries.`,
	dm.EGCCitadelClientMessages_k_EMsgClientToGCGetAccountStats: `GetAccountStats returns the career statistics of an account: wins, losses,
kills, deaths, and per-hero aggregates. Accounts other than your own require
the owner's privacy settings to allow it.`,
	dm.EGCCitadelClientMessages_k_EMsgClientToGCGetActiveMatches: `GetActiveMatches returns the matches currently in progress that involve
your account or party.`,
	dm.EGCCitadelClientMessages_k_EMsgClientToGCGetBook: `GetBook returns the full state of an account's hero book, including owned
pages and unlock progress, as a CMsgAccountBook.`,
	dm.EGCCitadelClientMessages_k_EMsgClientToGCGetDiscordLink: `GetDiscordLink returns a single-use URL that links the account to the
official Deadlock Discord, valid for the returned number of hours.`,
	dm.EGCCitadelClientMessages_k_EMsgClientToGCGetFriendGameStatus: `GetFriendGameStatus returns what your Steam friends are currently doing in
Deadlock: their game state, party, and lobby. Set include_invited to also see
friends who have been invited to your group.`,
	dm.EGCCitadelClientMessages_k_EMsgClientToGCGetHeroChoice: `GetHeroChoice returns the hero you selected for the free hero rotation:
the hero ID, its choice slot ID, and how many times the choice has been used.`,
	dm.EGCCitadelClientMessages_k_EMsgClientToGCGetHeroMMRRankings: `GetHeroMMRRankings returns the per-hero MMR of an account, relative to the
ranked ladder.`,
	dm.EGCCitadelClientMessages_k_EMsgClientToGCGetHideoutSDRTicket: `GetHideoutSDRTicket requests a Steam Datagram Relay ticket for connecting
to a hideout (sandbox lobby) server. The ticket is returned as raw bytes; an
error message is set instead if the ticket cannot be issued.`,
	dm.EGCCitadelClientMessages_k_EMsgClientToGCGetLeaderboard: `GetLeaderboard returns the calling account's placement on the leaderboard
for the given region and hero: position, badge level, and display name.`,
	dm.EGCCitadelClientMessages_k_EMsgClientToGCGetMatchHistory: `GetMatchHistory returns a page of the account's match history with
per-match statistics: result, duration, K/D/A, net worth, hero level, and
objective masks. Pass the continue_cursor from the previous response to fetch
older matches; game_mode and match_mode narrow the results.`,
	dm.EGCCitadelClientMessages_k_EMsgClientToGCGetMatchMetaData: `GetMatchMetaData returns the replay metadata for a match, including the
download details needed by replay tools. The metadata_salt is required for
recent matches and is distributed with the match result.`,
	dm.EGCCitadelClientMessages_k_EMsgClientToGCGetProfileCard: `GetProfileCard requests the profile card of an account.

The card is not returned directly; watch for a ProfileCardUpdated event after
sending this request.`,
	dm.EGCCitadelClientMessages_k_EMsgClientToGCGetRankData: `GetRankData returns the local account's ranked state: current rank
confidence, calibrated rank confidence, and whether a calibration interval is
still required.`,
	dm.EGCCitadelClientMessages_k_EMsgClientToGCGetSurveyQuestion: `GetSurveyQuestion returns the next survey question to show the player, if
one is queued. Submit answers with SubmitSurvey.`,
	dm.EGCCitadelClientMessages_k_EMsgClientToGCGrantForumAccess: `GrantForumAccess sends a forum access grant request bound to the given
email address.`,
	dm.EGCCitadelClientMessages_k_EMsgClientToGCPartyJoin: `JoinParty joins an existing party by ID.`,
	dm.EGCCitadelClientMessages_k_EMsgClientToGCPartyJoinViaCode: `JoinPartyViaCode joins a party using its numeric join code. The matchmaking
info describes your client state for the party leader.`,
	dm.EGCCitadelClientMessages_k_EMsgClientToGCLeaveLobby: `LeaveLobby leaves the current lobby, if any.`,
	dm.EGCCitadelClientMessages_k_EMsgClientToGCRecordClientEvents: `RecordClientEvents reports client-side account events to the GC for
telemetry. The client_event_index and client_run_token de-duplicate retries
within a client run.`,
	dm.EGCCitadelClientMessages_k_EMsgClientToGCReportPlayerFromMatch: `ReportPlayerFromMatch files a conduct report against a player from a match
you participated in. The response confirms the report was recorded.`,
	dm.EGCCitadelClientMessages_k_EMsgClientToGCDevRequestCheatReports: `RequestDevCheatReports asks the GC for cheat-detection data about recent
matches. Developer accounts only.`,
	dm.EGCCitadelClientMessages_k_EMsgClientToGCRequestHeroReleaseVoteTally: `RequestHeroReleaseVoteTally returns the current tally of the community hero
release vote.`,
	dm.EGCCitadelClientMessages_k_EMsgClientToGCModeratorRequest: `RequestModerator submits a moderator action request against an account.`,
	dm.EGCCitadelClientMessages_k_EMsgClientToGCBookUnlock: `SendBookUnlock purchases an unlock from the hero book. The expected_cost is
checked server-side; the purchase fails if it does not match.`,
	dm.EGCCitadelClientMessages_k_EMsgClientToGCChallengeActivate: `SendChallengeActivate activates one of your generated challenge choices so
it starts tracking progress.`,
	dm.EGCCitadelClientMessages_k_EMsgClientToGCCommendPlayerFromMatch: `SendCommendPlayerFromMatch sends a commend for a player from a match you
played together. The GC delivers a CommendNotification event to the target.`,
	dm.EGCCitadelClientMessages_k_EMsgClientToGCDeleteHeroBuild: `SendDeleteHeroBuild deletes one of your published hero builds. You must be
the author of the build.`,
	dm.EGCCitadelClientMessages_k_EMsgClientToGCDevAction: `SendDevAction runs a developer-only GC action. Developer accounts only.`,
	dm.EGCCitadelClientMessages_k_EMsgClientToGCIsInMatchmaking: `SendIsInMatchmaking checks whether the account is currently in the
matchmaking queue. The response carries the queue state.`,
	dm.EGCCitadelClientMessages_k_EMsgClientToGCPartyAction: `SendPartyAction performs an action on a party member, such as kicking or
promoting. See CMsgClientToGCPartyAction.EAction for the available actions;
the response reports the outcome.`,
	dm.EGCCitadelClientMessages_k_EMsgClientToGCPartySendChatMsg: `SendPartyChatMsg sends a chat line to your party members.`,
	dm.EGCCitadelClientMessages_k_EMsgClientToGCPartyInviteUser: `SendPartyInviteUser invites an account to your party. The invitee responds
through the party invite flow.`,
	dm.EGCCitadelClientMessages_k_EMsgClientToGCReplacementSDRTicket: `SendReplacementSDRTicket requests a replacement Steam Datagram Relay ticket
for a match or lobby connection whose ticket expired or was lost.`,
	dm.EGCCitadelClientMessages_k_EMsgClientToGCUnlockHero: `SendUnlockHero spends a hero unlock credit on the chosen hero.`,
	dm.EGCCitadelClientMessages_k_EMsgClientToGCUpdateAccountSync: `SendUpdateAccountSync pushes updated CSOAccountSyncStorage values to the
GC. The response confirms the sync.`,
	dm.EGCCitadelClientMessages_k_EMsgClientToGCUpdateHeroBuild: `SendUpdateHeroBuild creates or updates one of your published hero builds.`,
	dm.EGCCitadelClientMessages_k_EMsgClientToGCUpdateHeroBuildPreference: `SendUpdateHeroBuildPreference sets which published build you follow for a
hero.`,
	dm.EGCCitadelClientMessages_k_EMsgClientToGCUpdateRoster: `SendUpdateRoster updates your hero selection roster before entering a
matchmaking queue.`,
	dm.EGCCitadelClientMessages_k_EMsgClientToGCUpdateSpectatorStatus: `SendUpdateSpectatorStatus tells the GC which lobby you are spectating, or
clears it with stopped_spectating.`,
	dm.EGCCitadelClientMessages_k_EMsgClientToGCSetAccountPrivacySetting: `SetAccountPrivacySetting changes who can view your profile card and
stats.`,
	dm.EGCCitadelClientMessages_k_EMsgClientToGCGenerateChallengeSet: `SetChallengeGenerate generates a fresh set of challenge choices to replace
your current ones.`,
	dm.EGCCitadelClientMessages_k_EMsgClientToGCDevSetMMBias: `SetDevMMBias biases matchmaking for an account toward a skill value.
Developer accounts only.`,
	dm.EGCCitadelClientMessages_k_EMsgClientToGCSetNewPlayerProgress: `SetNewPlayerProgress marks a new-player progression milestone complete,
such as finishing the tutorial or unlocking ranked play.`,
	dm.EGCCitadelClientMessages_k_EMsgClientToGCPartySetMode: `SetPartyMode configures the party's matchmaking settings: match mode, game
mode, bot difficulty, and region.`,
	dm.EGCCitadelClientMessages_k_EMsgClientToGCPartySetReadyState: `SetPartyReadyState marks you ready (or not ready) for the party's next
match, along with the hero roster you intend to play.`,
	dm.EGCCitadelClientMessages_k_EMsgClientToGCSetServerConVar: `SetServerConVar sets a console variable on a practice lobby server.
Developer accounts only.`,
	dm.EGCCitadelClientMessages_k_EMsgClientToGCSpectateLobby: `SpectateLobby requests to spectate a lobby. The response carries the
connection details for the spectator broadcast.`,
	dm.EGCCitadelClientMessages_k_EMsgClientToGCSpectateUser: `SpectateUser requests to spectate the live game of an account. The response
carries the connection details, or an error if no game is running.`,
	dm.EGCCitadelClientMessages_k_EMsgClientToGCStartMatchmaking: `StartMatchmaking enters the matchmaking queue with the given match info,
region ping times, and hero roster. The response confirms the queue state.
Watch for MatchmakingStopped events for asynchronous cancellation.`,
	dm.EGCCitadelClientMessages_k_EMsgClientToGCPartyStartMatch: `StartPartyMatch starts the search for a match on behalf of the whole
party. Party leaders only.`,
	dm.EGCCitadelClientMessages_k_EMsgClientToGCStartRankedInterval: `StartRankedInterval begins a ranked calibration interval. The response
reports the new ranked state.`,
	dm.EGCCitadelClientMessages_k_EMsgClientToGCStopMatchmaking: `StopMatchmaking leaves the matchmaking queue. A MatchmakingStopped event
confirms the exit.`,
	dm.EGCCitadelClientMessages_k_EMsgClientToGCSubmitHeroReleaseVote: `SubmitHeroReleaseVote casts a vote in the community hero release vote for
the given round.`,
	dm.EGCCitadelClientMessages_k_EMsgClientToGCSubmitPlaytestUser: `SubmitPlaytestUser registers an account for a playtest session at the
given location. Developer accounts only.`,
	dm.EGCCitadelClientMessages_k_EMsgClientToGCSubmitSurvey: `SubmitSurvey answers a survey question previously delivered through
GetSurveyQuestion. Set is_skip to decline without answering.`,
}

// msgEventDocOverrides contains handwritten godoc comments for events.
// An entry replaces the entire generated comment for the event.
var msgEventDocOverrides = map[dm.EGCCitadelClientMessages]string{
	dm.EGCCitadelClientMessages_k_EMsgGCToClientMatchmakingStopped: `MatchmakingStopped is emitted when the GC removes the account from the
matchmaking queue, either because the client requested it or because the GC
aborted the search.`,
	dm.EGCCitadelClientMessages_k_EMsgGCToClientProfileCardUpdated: `ProfileCardUpdated is emitted with the profile card of an account, usually
in reply to GetProfileCard.`,
	dm.EGCCitadelClientMessages_k_EMsgGCToClientPartyEvent: `PartyEvent is emitted when something in your party changes: membership,
readiness, or match state.`,
	dm.EGCCitadelClientMessages_k_EMsgGCToClientCanRejoinParty: `CanRejoinParty is emitted when the GC determines your account may rejoin a
party it was disconnected from.`,
	dm.EGCCitadelClientMessages_k_EMsgGCToClientSDRTicket: `SDRTicket is emitted with a Steam Datagram Relay ticket for connecting to
a match or hideout server.`,
	dm.EGCCitadelClientMessages_k_EMsgGCToClientCommendNotification: `CommendNotification is emitted when another player commends your account.`,
	dm.EGCCitadelClientMessages_k_EMsgGCToClientAccountStatsUpdated: `AccountStatsUpdated is emitted when the GC pushes updated career statistics
for your account.`,
	dm.EGCCitadelClientMessages_k_EMsgGCToClientBookUpdated: `BookUpdated is emitted when the GC pushes updated hero book state, such as
after a book purchase.`,
	dm.EGCCitadelClientMessages_k_EMsgGCToClientTarotUserDataUpdated: `TarotUserDataUpdated is emitted when the GC pushes updated tarot board
state for your account.`,
	dm.EGCCitadelClientMessages_k_EMsgGCToClientHeroLabsSchedule: `HeroLabsSchedule is emitted with the schedule of upcoming hero labs
playtest sessions.`,
}
