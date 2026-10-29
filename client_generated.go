package deadlock

import (
	"context"

	"github.com/paralin/go-deadlock/events"
	"github.com/paralin/go-deadlock/protocol"
)

// CancelChallenge cancels one of your active challenges.
//
// The provided challenge must belong to your account. The response reports the
// result of the cancellation.
func (d *Deadlock) CancelChallenge(
	ctx context.Context,
	challengeID uint64,
) (*protocol.CMsgClientToGCChallengeCancelResponse, error) {
	req := &protocol.CMsgClientToGCChallengeCancel{
		ChallengeId: &challengeID,
	}
	resp := &protocol.CMsgClientToGCChallengeCancelResponse{}

	return resp, d.MakeRequest(
		ctx,
		uint32(protocol.EGCCitadelClientMessages_k_EMsgClientToGCChallengeCancel),
		req,
		uint32(protocol.EGCCitadelClientMessages_k_EMsgClientToGCChallengeCancelResponse),
		resp,
	)
}

// ClaimChallenge claims the reward for a completed challenge.
//
// Only challenges that have reached the completed state can be claimed. The
// response reports the result of the claim.
func (d *Deadlock) ClaimChallenge(
	ctx context.Context,
	challengeID uint64,
) (*protocol.CMsgClientToGCChallengeClaimResponse, error) {
	req := &protocol.CMsgClientToGCChallengeClaim{
		ChallengeId: &challengeID,
	}
	resp := &protocol.CMsgClientToGCChallengeClaimResponse{}

	return resp, d.MakeRequest(
		ctx,
		uint32(protocol.EGCCitadelClientMessages_k_EMsgClientToGCChallengeClaim),
		req,
		uint32(protocol.EGCCitadelClientMessages_k_EMsgClientToGCChallengeClaimResponse),
		resp,
	)
}

// CreateCabal creates a cabal.
//
// Sends the GC message k_EMsgClientToGCCabalCreate (CMsgClientToGCCabalCreate) and awaits the response k_EMsgClientToGCCabalCreateResponse,
// delivered as *CMsgClientToGCCabalCreateResponse.
func (d *Deadlock) CreateCabal(
	ctx context.Context,
) (*protocol.CMsgClientToGCCabalCreateResponse, error) {
	req := &protocol.CMsgClientToGCCabalCreate{}
	resp := &protocol.CMsgClientToGCCabalCreateResponse{}

	return resp, d.MakeRequest(
		ctx,
		uint32(protocol.EGCCitadelClientMessages_k_EMsgClientToGCCabalCreate),
		req,
		uint32(protocol.EGCCitadelClientMessages_k_EMsgClientToGCCabalCreateResponse),
		resp,
	)
}

// CreateParty creates a new party with the given matchmaking options.
//
// The party join code is returned in the response, unless party codes are
// disabled. The created party is also delivered through the SOCache as a
// CSOCitadelParty object.
func (d *Deadlock) CreateParty(
	ctx context.Context,
	req *protocol.CMsgClientToGCPartyCreate,
) (*protocol.CMsgClientToGCPartyCreateResponse, error) {
	resp := &protocol.CMsgClientToGCPartyCreateResponse{}

	return resp, d.MakeRequest(
		ctx,
		uint32(protocol.EGCCitadelClientMessages_k_EMsgClientToGCPartyCreate),
		req,
		uint32(protocol.EGCCitadelClientMessages_k_EMsgClientToGCPartyCreateResponse),
		resp,
	)
}

// FindHeroBuilds searches published hero builds.
//
// At most one filter is meaningful per call: builds by author_account_id, for
// hero_id, or matching search_text. The response lists the matching builds with
// their metadata.
func (d *Deadlock) FindHeroBuilds(
	ctx context.Context,
	authorAccountID uint32,
	heroID uint32,
	language []int32,
	searchText string,
	heroBuildID uint32,
	tags []uint32,
) (*protocol.CMsgClientToGCFindHeroBuildsResponse, error) {
	req := &protocol.CMsgClientToGCFindHeroBuilds{
		AuthorAccountId: &authorAccountID,
		HeroId:          &heroID,
		Language:        language,
		SearchText:      &searchText,
		HeroBuildId:     &heroBuildID,
		Tags:            tags,
	}
	resp := &protocol.CMsgClientToGCFindHeroBuildsResponse{}

	return resp, d.MakeRequest(
		ctx,
		uint32(protocol.EGCCitadelClientMessages_k_EMsgClientToGCFindHeroBuilds),
		req,
		uint32(protocol.EGCCitadelClientMessages_k_EMsgClientToGCFindHeroBuildsResponse),
		resp,
	)
}

// GetAccountLeaderboards returns the leaderboard placements of an account,
// one entry per ranked hero leaderboard the account appears on.
func (d *Deadlock) GetAccountLeaderboards(
	ctx context.Context,
	accountID uint32,
	leaderboardID uint32,
) (*protocol.CMsgClientToGCGetAccountLeaderboardsResponse, error) {
	req := &protocol.CMsgClientToGCGetAccountLeaderboards{
		AccountId:     &accountID,
		LeaderboardId: &leaderboardID,
	}
	resp := &protocol.CMsgClientToGCGetAccountLeaderboardsResponse{}

	return resp, d.MakeRequest(
		ctx,
		uint32(protocol.EGCCitadelClientMessages_k_EMsgClientToGCGetAccountLeaderboards),
		req,
		uint32(protocol.EGCCitadelClientMessages_k_EMsgClientToGCGetAccountLeaderboardsResponse),
		resp,
	)
}

// GetAccountMatchReports fetches the conduct report information tied to a
// match. Supply the most recent known match_id to page through older entries.
func (d *Deadlock) GetAccountMatchReports(
	ctx context.Context,
	matchID uint64,
) (*protocol.CMsgClientToGCGetAccountMatchReportsResponse, error) {
	req := &protocol.CMsgClientToGCGetAccountMatchReports{
		MatchId: &matchID,
	}
	resp := &protocol.CMsgClientToGCGetAccountMatchReportsResponse{}

	return resp, d.MakeRequest(
		ctx,
		uint32(protocol.EGCCitadelClientMessages_k_EMsgClientToGCGetAccountMatchReports),
		req,
		uint32(protocol.EGCCitadelClientMessages_k_EMsgClientToGCGetAccountMatchReportsResponse),
		resp,
	)
}

// GetAccountStats returns the career statistics of an account: wins, losses,
// kills, deaths, and per-hero aggregates. Accounts other than your own require
// the owner's privacy settings to allow it.
func (d *Deadlock) GetAccountStats(
	ctx context.Context,
	accountID uint32,
	devAccessHint bool,
	friendAccessHint bool,
) (*protocol.CMsgClientToGCGetAccountStatsResponse, error) {
	req := &protocol.CMsgClientToGCGetAccountStats{
		AccountId:        &accountID,
		DevAccessHint:    &devAccessHint,
		FriendAccessHint: &friendAccessHint,
	}
	resp := &protocol.CMsgClientToGCGetAccountStatsResponse{}

	return resp, d.MakeRequest(
		ctx,
		uint32(protocol.EGCCitadelClientMessages_k_EMsgClientToGCGetAccountStats),
		req,
		uint32(protocol.EGCCitadelClientMessages_k_EMsgClientToGCGetAccountStatsResponse),
		resp,
	)
}

// GetActiveMatches returns the matches currently in progress that involve
// your account or party.
func (d *Deadlock) GetActiveMatches(
	ctx context.Context,
) (*protocol.CMsgClientToGCGetActiveMatchesResponse, error) {
	req := &protocol.CMsgClientToGCGetActiveMatches{}
	resp := &protocol.CMsgClientToGCGetActiveMatchesResponse{}

	return resp, d.MakeRequest(
		ctx,
		uint32(protocol.EGCCitadelClientMessages_k_EMsgClientToGCGetActiveMatches),
		req,
		uint32(protocol.EGCCitadelClientMessages_k_EMsgClientToGCGetActiveMatchesResponse),
		resp,
	)
}

// GetBook returns the full state of an account's hero book, including owned
// pages and unlock progress, as a CMsgAccountBook.
func (d *Deadlock) GetBook(
	ctx context.Context,
	bookID uint32,
) (*protocol.CMsgClientToGCGetBookResponse, error) {
	req := &protocol.CMsgClientToGCGetBook{
		BookId: &bookID,
	}
	resp := &protocol.CMsgClientToGCGetBookResponse{}

	return resp, d.MakeRequest(
		ctx,
		uint32(protocol.EGCCitadelClientMessages_k_EMsgClientToGCGetBook),
		req,
		uint32(protocol.EGCCitadelClientMessages_k_EMsgClientToGCGetBookResponse),
		resp,
	)
}

// GetCabalDetails gets cabal details.
//
// Sends the GC message k_EMsgClientToGCCabalGetDetails (CMsgClientToGCCabalGetDetails) and awaits the response k_EMsgClientToGCCabalGetDetailsResponse,
// delivered as *CMsgClientToGCCabalGetDetailsResponse.
func (d *Deadlock) GetCabalDetails(
	ctx context.Context,
	cabalID uint64,
) (*protocol.CMsgClientToGCCabalGetDetailsResponse, error) {
	req := &protocol.CMsgClientToGCCabalGetDetails{
		CabalId: &cabalID,
	}
	resp := &protocol.CMsgClientToGCCabalGetDetailsResponse{}

	return resp, d.MakeRequest(
		ctx,
		uint32(protocol.EGCCitadelClientMessages_k_EMsgClientToGCCabalGetDetails),
		req,
		uint32(protocol.EGCCitadelClientMessages_k_EMsgClientToGCCabalGetDetailsResponse),
		resp,
	)
}

// GetCabalMatchHistory gets a cabal match history.
//
// Sends the GC message k_EMsgClientToGCCabalGetMatchHistory (CMsgClientToGCCabalGetMatchHistory) and awaits the response k_EMsgClientToGCCabalGetMatchHistoryResponse,
// delivered as *CMsgClientToGCCabalGetMatchHistoryResponse.
func (d *Deadlock) GetCabalMatchHistory(
	ctx context.Context,
	cabalID uint64,
	cabalMmTier protocol.ECitadelCabalMMTier,
	cursor uint64,
) (*protocol.CMsgClientToGCCabalGetMatchHistoryResponse, error) {
	req := &protocol.CMsgClientToGCCabalGetMatchHistory{
		CabalId:     &cabalID,
		CabalMmTier: &cabalMmTier,
		Cursor:      &cursor,
	}
	resp := &protocol.CMsgClientToGCCabalGetMatchHistoryResponse{}

	return resp, d.MakeRequest(
		ctx,
		uint32(protocol.EGCCitadelClientMessages_k_EMsgClientToGCCabalGetMatchHistory),
		req,
		uint32(protocol.EGCCitadelClientMessages_k_EMsgClientToGCCabalGetMatchHistoryResponse),
		resp,
	)
}

// GetDiscordLink returns a single-use URL that links the account to the
// official Deadlock Discord, valid for the returned number of hours.
func (d *Deadlock) GetDiscordLink(
	ctx context.Context,
) (*protocol.CMsgClientToGCGetDiscordLinkResponse, error) {
	req := &protocol.CMsgClientToGCGetDiscordLink{}
	resp := &protocol.CMsgClientToGCGetDiscordLinkResponse{}

	return resp, d.MakeRequest(
		ctx,
		uint32(protocol.EGCCitadelClientMessages_k_EMsgClientToGCGetDiscordLink),
		req,
		uint32(protocol.EGCCitadelClientMessages_k_EMsgClientToGCGetDiscordLinkResponse),
		resp,
	)
}

// GetFriendGameStatus returns what your Steam friends are currently doing in
// Deadlock: their game state, party, and lobby. Set include_invited to also see
// friends who have been invited to your group.
func (d *Deadlock) GetFriendGameStatus(
	ctx context.Context,
	includeInvited bool,
) (*protocol.CMsgClientToGCGetFriendGameStatusResponse, error) {
	req := &protocol.CMsgClientToGCGetFriendGameStatus{
		IncludeInvited: &includeInvited,
	}
	resp := &protocol.CMsgClientToGCGetFriendGameStatusResponse{}

	return resp, d.MakeRequest(
		ctx,
		uint32(protocol.EGCCitadelClientMessages_k_EMsgClientToGCGetFriendGameStatus),
		req,
		uint32(protocol.EGCCitadelClientMessages_k_EMsgClientToGCGetFriendGameStatusResponse),
		resp,
	)
}

// GetHeroChoice returns the hero you selected for the free hero rotation:
// the hero ID, its choice slot ID, and how many times the choice has been used.
func (d *Deadlock) GetHeroChoice(
	ctx context.Context,
) (*protocol.CMsgClientToGCGetHeroChoiceResponse, error) {
	req := &protocol.CMsgClientToGCGetHeroChoice{}
	resp := &protocol.CMsgClientToGCGetHeroChoiceResponse{}

	return resp, d.MakeRequest(
		ctx,
		uint32(protocol.EGCCitadelClientMessages_k_EMsgClientToGCGetHeroChoice),
		req,
		uint32(protocol.EGCCitadelClientMessages_k_EMsgClientToGCGetHeroChoiceResponse),
		resp,
	)
}

// GetHeroMMRRankings returns the per-hero MMR of an account, relative to the
// ranked ladder.
func (d *Deadlock) GetHeroMMRRankings(
	ctx context.Context,
	targetAccountID uint32,
) (*protocol.CMsgClientToGCGetHeroMMRRankingsResponse, error) {
	req := &protocol.CMsgClientToGCGetHeroMMRRankings{
		TargetAccountId: &targetAccountID,
	}
	resp := &protocol.CMsgClientToGCGetHeroMMRRankingsResponse{}

	return resp, d.MakeRequest(
		ctx,
		uint32(protocol.EGCCitadelClientMessages_k_EMsgClientToGCGetHeroMMRRankings),
		req,
		uint32(protocol.EGCCitadelClientMessages_k_EMsgClientToGCGetHeroMMRRankingsResponse),
		resp,
	)
}

// GetHideoutSDRTicket requests a Steam Datagram Relay ticket for connecting
// to a hideout (sandbox lobby) server. The ticket is returned as raw bytes; an
// error message is set instead if the ticket cannot be issued.
func (d *Deadlock) GetHideoutSDRTicket(
	ctx context.Context,
	lobbyID uint64,
) (*protocol.CMsgClientToGCGetHideoutSDRTicketResponse, error) {
	req := &protocol.CMsgClientToGCGetHideoutSDRTicket{
		LobbyId: &lobbyID,
	}
	resp := &protocol.CMsgClientToGCGetHideoutSDRTicketResponse{}

	return resp, d.MakeRequest(
		ctx,
		uint32(protocol.EGCCitadelClientMessages_k_EMsgClientToGCGetHideoutSDRTicket),
		req,
		uint32(protocol.EGCCitadelClientMessages_k_EMsgClientToGCGetHideoutSDRTicketResponse),
		resp,
	)
}

// GetInternalLeaderboards gets internal leaderboards.
//
// Sends the GC message k_EMsgClientToGCGetInternalLeaderboards (CMsgClientToGCGetInternalLeaderboards) and awaits the response k_EMsgClientToGCGetInternalLeaderboardsResponse,
// delivered as *CMsgClientToGCGetInternalLeaderboardsResponse.
func (d *Deadlock) GetInternalLeaderboards(
	ctx context.Context,
	accountID uint32,
) (*protocol.CMsgClientToGCGetInternalLeaderboardsResponse, error) {
	req := &protocol.CMsgClientToGCGetInternalLeaderboards{
		AccountId: &accountID,
	}
	resp := &protocol.CMsgClientToGCGetInternalLeaderboardsResponse{}

	return resp, d.MakeRequest(
		ctx,
		uint32(protocol.EGCCitadelClientMessages_k_EMsgClientToGCGetInternalLeaderboards),
		req,
		uint32(protocol.EGCCitadelClientMessages_k_EMsgClientToGCGetInternalLeaderboardsResponse),
		resp,
	)
}

// GetLeaderboard returns the calling account's placement on the leaderboard
// for the given region and hero: position, badge level, and display name.
func (d *Deadlock) GetLeaderboard(
	ctx context.Context,
	leaderboardRegion protocol.ECitadelLeaderboardRegion,
	heroID uint32,
	leaderboardID uint32,
) (*protocol.CMsgClientToGCGetLeaderboardResponse, error) {
	req := &protocol.CMsgClientToGCGetLeaderboard{
		LeaderboardRegion: &leaderboardRegion,
		HeroId:            &heroID,
		LeaderboardId:     &leaderboardID,
	}
	resp := &protocol.CMsgClientToGCGetLeaderboardResponse{}

	return resp, d.MakeRequest(
		ctx,
		uint32(protocol.EGCCitadelClientMessages_k_EMsgClientToGCGetLeaderboard),
		req,
		uint32(protocol.EGCCitadelClientMessages_k_EMsgClientToGCGetLeaderboardResponse),
		resp,
	)
}

// GetLeaderboardStatus gets leaderboard status.
//
// Sends the GC message k_EMsgClientToGCGetLeaderboardStatus (CMsgClientToGCGetLeaderboardStatus) and awaits the response k_EMsgClientToGCGetLeaderboardStatusResponse,
// delivered as *CMsgClientToGCGetLeaderboardStatusResponse.
func (d *Deadlock) GetLeaderboardStatus(
	ctx context.Context,
	leaderboardID uint32,
) (*protocol.CMsgClientToGCGetLeaderboardStatusResponse, error) {
	req := &protocol.CMsgClientToGCGetLeaderboardStatus{
		LeaderboardId: &leaderboardID,
	}
	resp := &protocol.CMsgClientToGCGetLeaderboardStatusResponse{}

	return resp, d.MakeRequest(
		ctx,
		uint32(protocol.EGCCitadelClientMessages_k_EMsgClientToGCGetLeaderboardStatus),
		req,
		uint32(protocol.EGCCitadelClientMessages_k_EMsgClientToGCGetLeaderboardStatusResponse),
		resp,
	)
}

// GetMatchHistory returns a page of the account's match history with
// per-match statistics: result, duration, K/D/A, net worth, hero level, and
// objective masks. Pass the continue_cursor from the previous response to fetch
// older matches; game_mode and match_mode narrow the results.
func (d *Deadlock) GetMatchHistory(
	ctx context.Context,
	accountID uint32,
	continueCursor uint64,
	gameMode protocol.ECitadelGameMode,
	matchMode protocol.ECitadelMatchMode,
	rankedType protocol.ECitadelRankedType,
	rankInterval uint32,
	cabalID uint64,
) (*protocol.CMsgClientToGCGetMatchHistoryResponse, error) {
	req := &protocol.CMsgClientToGCGetMatchHistory{
		AccountId:      &accountID,
		ContinueCursor: &continueCursor,
		GameMode:       &gameMode,
		MatchMode:      &matchMode,
		RankedType:     &rankedType,
		RankInterval:   &rankInterval,
		CabalId:        &cabalID,
	}
	resp := &protocol.CMsgClientToGCGetMatchHistoryResponse{}

	return resp, d.MakeRequest(
		ctx,
		uint32(protocol.EGCCitadelClientMessages_k_EMsgClientToGCGetMatchHistory),
		req,
		uint32(protocol.EGCCitadelClientMessages_k_EMsgClientToGCGetMatchHistoryResponse),
		resp,
	)
}

// GetMatchMetaData returns the replay metadata for a match, including the
// download details needed by replay tools. The metadata_salt is required for
// recent matches and is distributed with the match result.
func (d *Deadlock) GetMatchMetaData(
	ctx context.Context,
	matchID uint64,
	metadataSalt uint32,
	targetAccountID uint32,
) (*protocol.CMsgClientToGCGetMatchMetaDataResponse, error) {
	req := &protocol.CMsgClientToGCGetMatchMetaData{
		MatchId:         &matchID,
		MetadataSalt:    &metadataSalt,
		TargetAccountId: &targetAccountID,
	}
	resp := &protocol.CMsgClientToGCGetMatchMetaDataResponse{}

	return resp, d.MakeRequest(
		ctx,
		uint32(protocol.EGCCitadelClientMessages_k_EMsgClientToGCGetMatchMetaData),
		req,
		uint32(protocol.EGCCitadelClientMessages_k_EMsgClientToGCGetMatchMetaDataResponse),
		resp,
	)
}

// GetProfileCard requests the profile card of an account.
//
// The card is not returned directly; watch for a ProfileCardUpdated event after
// sending this request.
func (d *Deadlock) GetProfileCard(
	accountID uint32,
	devAccessHint bool,
	friendAccessHint bool,
) {
	req := &protocol.CMsgClientToGCGetProfileCard{
		AccountId:        &accountID,
		DevAccessHint:    &devAccessHint,
		FriendAccessHint: &friendAccessHint,
	}
	d.write(uint32(protocol.EGCCitadelClientMessages_k_EMsgClientToGCGetProfileCard), req)
}

// GetRankData returns the local account's ranked state: current rank
// confidence, calibrated rank confidence, and whether a calibration interval is
// still required.
func (d *Deadlock) GetRankData(
	ctx context.Context,
) (*protocol.CMsgGCToClientGetRankDataResponse, error) {
	req := &protocol.CMsgClientToGCGetRankData{}
	resp := &protocol.CMsgGCToClientGetRankDataResponse{}

	return resp, d.MakeRequest(
		ctx,
		uint32(protocol.EGCCitadelClientMessages_k_EMsgClientToGCGetRankData),
		req,
		uint32(protocol.EGCCitadelClientMessages_k_EMsgGCToClientGetRankDataResponse),
		resp,
	)
}

// GetSurveyQuestion returns the next survey question to show the player, if
// one is queued. Submit answers with SubmitSurvey.
func (d *Deadlock) GetSurveyQuestion(
	ctx context.Context,
	devForceNewQuestionType protocol.CMsgSurveyQuestion_EQuestionType,
) (*protocol.CMsgClientToGCGetSurveyQuestionResponse, error) {
	req := &protocol.CMsgClientToGCGetSurveyQuestion{
		DevForceNewQuestionType: &devForceNewQuestionType,
	}
	resp := &protocol.CMsgClientToGCGetSurveyQuestionResponse{}

	return resp, d.MakeRequest(
		ctx,
		uint32(protocol.EGCCitadelClientMessages_k_EMsgClientToGCGetSurveyQuestion),
		req,
		uint32(protocol.EGCCitadelClientMessages_k_EMsgClientToGCGetSurveyQuestionResponse),
		resp,
	)
}

// GrantForumAccess sends a forum access grant request bound to the given
// email address.
func (d *Deadlock) GrantForumAccess(
	ctx context.Context,
	email string,
) (*protocol.CMsgClientToGCGrantForumAccessResponse, error) {
	req := &protocol.CMsgClientToGCGrantForumAccess{
		Email: &email,
	}
	resp := &protocol.CMsgClientToGCGrantForumAccessResponse{}

	return resp, d.MakeRequest(
		ctx,
		uint32(protocol.EGCCitadelClientMessages_k_EMsgClientToGCGrantForumAccess),
		req,
		uint32(protocol.EGCCitadelClientMessages_k_EMsgClientToGCGrantForumAccessResponse),
		resp,
	)
}

// JoinParty joins an existing party by ID.
func (d *Deadlock) JoinParty(
	ctx context.Context,
	partyID uint64,
	isRejoin bool,
	partyMmInfo protocol.CMsgPartyMMInfo,
) (*protocol.CMsgClientToGCPartyJoinResponse, error) {
	req := &protocol.CMsgClientToGCPartyJoin{
		PartyId:     &partyID,
		IsRejoin:    &isRejoin,
		PartyMmInfo: &partyMmInfo,
	}
	resp := &protocol.CMsgClientToGCPartyJoinResponse{}

	return resp, d.MakeRequest(
		ctx,
		uint32(protocol.EGCCitadelClientMessages_k_EMsgClientToGCPartyJoin),
		req,
		uint32(protocol.EGCCitadelClientMessages_k_EMsgClientToGCPartyJoinResponse),
		resp,
	)
}

// JoinPartyViaCode joins a party using its numeric join code. The matchmaking
// info describes your client state for the party leader.
func (d *Deadlock) JoinPartyViaCode(
	ctx context.Context,
	joinCode uint64,
	partyMmInfo protocol.CMsgPartyMMInfo,
) (*protocol.CMsgClientToGCPartyJoinViaCodeResponse, error) {
	req := &protocol.CMsgClientToGCPartyJoinViaCode{
		JoinCode:    &joinCode,
		PartyMmInfo: &partyMmInfo,
	}
	resp := &protocol.CMsgClientToGCPartyJoinViaCodeResponse{}

	return resp, d.MakeRequest(
		ctx,
		uint32(protocol.EGCCitadelClientMessages_k_EMsgClientToGCPartyJoinViaCode),
		req,
		uint32(protocol.EGCCitadelClientMessages_k_EMsgClientToGCPartyJoinViaCodeResponse),
		resp,
	)
}

// KickCabalPlayer kicks a cabal player.
//
// Sends the GC message k_EMsgClientToGCCabalKickPlayer (CMsgClientToGCCabalKickPlayer) and awaits the response k_EMsgClientToGCCabalKickPlayerResponse,
// delivered as *CMsgClientToGCCabalKickPlayerResponse.
func (d *Deadlock) KickCabalPlayer(
	ctx context.Context,
	cabalID uint64,
	kickAccountID uint32,
) (*protocol.CMsgClientToGCCabalKickPlayerResponse, error) {
	req := &protocol.CMsgClientToGCCabalKickPlayer{
		CabalId:       &cabalID,
		KickAccountId: &kickAccountID,
	}
	resp := &protocol.CMsgClientToGCCabalKickPlayerResponse{}

	return resp, d.MakeRequest(
		ctx,
		uint32(protocol.EGCCitadelClientMessages_k_EMsgClientToGCCabalKickPlayer),
		req,
		uint32(protocol.EGCCitadelClientMessages_k_EMsgClientToGCCabalKickPlayerResponse),
		resp,
	)
}

// LeaveCabal leaves a cabal.
//
// Sends the GC message k_EMsgClientToGCCabalLeave (CMsgClientToGCCabalLeave) and awaits the response k_EMsgClientToGCCabalLeaveResponse,
// delivered as *CMsgClientToGCCabalLeaveResponse.
func (d *Deadlock) LeaveCabal(
	ctx context.Context,
	cabalID uint64,
	acknowledgedLastMember bool,
) (*protocol.CMsgClientToGCCabalLeaveResponse, error) {
	req := &protocol.CMsgClientToGCCabalLeave{
		CabalId:                &cabalID,
		AcknowledgedLastMember: &acknowledgedLastMember,
	}
	resp := &protocol.CMsgClientToGCCabalLeaveResponse{}

	return resp, d.MakeRequest(
		ctx,
		uint32(protocol.EGCCitadelClientMessages_k_EMsgClientToGCCabalLeave),
		req,
		uint32(protocol.EGCCitadelClientMessages_k_EMsgClientToGCCabalLeaveResponse),
		resp,
	)
}

// LeaveLobby leaves the current lobby, if any.
func (d *Deadlock) LeaveLobby(
	ctx context.Context,
	lobbyID uint64,
) (*protocol.CMsgClientToGCLeaveLobbyResponse, error) {
	req := &protocol.CMsgClientToGCLeaveLobby{
		LobbyId: &lobbyID,
	}
	resp := &protocol.CMsgClientToGCLeaveLobbyResponse{}

	return resp, d.MakeRequest(
		ctx,
		uint32(protocol.EGCCitadelClientMessages_k_EMsgClientToGCLeaveLobby),
		req,
		uint32(protocol.EGCCitadelClientMessages_k_EMsgClientToGCLeaveLobbyResponse),
		resp,
	)
}

// LeaveParty leaves a party.
//
// Sends the GC message k_EMsgClientToGCPartyLeave (CMsgClientToGCPartyLeave) and awaits the response k_EMsgClientToGCPartyLeaveResponse,
// delivered as *CMsgClientToGCPartyLeaveResponse.
func (d *Deadlock) LeaveParty(
	ctx context.Context,
	partyID uint64,
) (*protocol.CMsgClientToGCPartyLeaveResponse, error) {
	req := &protocol.CMsgClientToGCPartyLeave{
		PartyId: &partyID,
	}
	resp := &protocol.CMsgClientToGCPartyLeaveResponse{}

	return resp, d.MakeRequest(
		ctx,
		uint32(protocol.EGCCitadelClientMessages_k_EMsgClientToGCPartyLeave),
		req,
		uint32(protocol.EGCCitadelClientMessages_k_EMsgClientToGCPartyLeaveResponse),
		resp,
	)
}

// RecordClientEvents reports client-side account events to the GC for
// telemetry. The client_event_index and client_run_token de-duplicate retries
// within a client run.
func (d *Deadlock) RecordClientEvents(
	ctx context.Context,
	events []*protocol.CMsgClientToGCRecordClientEvents_Event,
	clientRunToken uint32,
) (*protocol.CMsgClientToGCRecordClientEventsResponse, error) {
	req := &protocol.CMsgClientToGCRecordClientEvents{
		Events:         events,
		ClientRunToken: &clientRunToken,
	}
	resp := &protocol.CMsgClientToGCRecordClientEventsResponse{}

	return resp, d.MakeRequest(
		ctx,
		uint32(protocol.EGCCitadelClientMessages_k_EMsgClientToGCRecordClientEvents),
		req,
		uint32(protocol.EGCCitadelClientMessages_k_EMsgClientToGCRecordClientEventsResponse),
		resp,
	)
}

// ReportPlayerFromMatch files a conduct report against a player from a match
// you participated in. The response confirms the report was recorded.
func (d *Deadlock) ReportPlayerFromMatch(
	ctx context.Context,
	matchID uint64,
	targetAccountID uint32,
	reportType protocol.CMsgClientToGCReportPlayerFromMatch_EReportType,
	reportText string,
) (*protocol.CMsgClientToGCReportPlayerFromMatchResponse, error) {
	req := &protocol.CMsgClientToGCReportPlayerFromMatch{
		MatchId:         &matchID,
		TargetAccountId: &targetAccountID,
		ReportType:      &reportType,
		ReportText:      &reportText,
	}
	resp := &protocol.CMsgClientToGCReportPlayerFromMatchResponse{}

	return resp, d.MakeRequest(
		ctx,
		uint32(protocol.EGCCitadelClientMessages_k_EMsgClientToGCReportPlayerFromMatch),
		req,
		uint32(protocol.EGCCitadelClientMessages_k_EMsgClientToGCReportPlayerFromMatchResponse),
		resp,
	)
}

// RequestDevCheatReports asks the GC for cheat-detection data about recent
// matches. Developer accounts only.
func (d *Deadlock) RequestDevCheatReports(
	ctx context.Context,
) (*protocol.CMsgClientToGCRequestCheatReportsResponse, error) {
	req := &protocol.CMsgClientToGCRequestCheatReports{}
	resp := &protocol.CMsgClientToGCRequestCheatReportsResponse{}

	return resp, d.MakeRequest(
		ctx,
		uint32(protocol.EGCCitadelClientMessages_k_EMsgClientToGCDevRequestCheatReports),
		req,
		uint32(protocol.EGCCitadelClientMessages_k_EMsgClientToGCDevRequestCheatReportsResponse),
		resp,
	)
}

// RequestHeroReleaseVoteTally returns the current tally of the community hero
// release vote.
func (d *Deadlock) RequestHeroReleaseVoteTally(
	voteRounds []uint32,
	isInitialRequest bool,
) {
	req := &protocol.CMsgClientToGCRequestHeroReleaseVoteTally{
		VoteRounds:       voteRounds,
		IsInitialRequest: &isInitialRequest,
	}
	d.write(uint32(protocol.EGCCitadelClientMessages_k_EMsgClientToGCRequestHeroReleaseVoteTally), req)
}

// RequestModerator submits a moderator action request against an account.
func (d *Deadlock) RequestModerator(
	ctx context.Context,
	accountID uint32,
) (*protocol.CMsgClientToGCModeratorRequestResponse, error) {
	req := &protocol.CMsgClientToGCModeratorRequest{
		AccountId: &accountID,
	}
	resp := &protocol.CMsgClientToGCModeratorRequestResponse{}

	return resp, d.MakeRequest(
		ctx,
		uint32(protocol.EGCCitadelClientMessages_k_EMsgClientToGCModeratorRequest),
		req,
		uint32(protocol.EGCCitadelClientMessages_k_EMsgClientToGCModeratorRequestResponse),
		resp,
	)
}

// RequestReporterUpdates requests reporter updates.
//
// Sends the GC message k_EMsgClientToGCRequestReporterUpdates (CMsgClientToGCRequestReporterUpdates) and awaits the response k_EMsgClientToGCRequestReporterUpdatesResponse,
// delivered as *CMsgClientToGCRequestReporterUpdatesResponse.
func (d *Deadlock) RequestReporterUpdates(
	ctx context.Context,
) (*protocol.CMsgClientToGCRequestReporterUpdatesResponse, error) {
	req := &protocol.CMsgClientToGCRequestReporterUpdates{}
	resp := &protocol.CMsgClientToGCRequestReporterUpdatesResponse{}

	return resp, d.MakeRequest(
		ctx,
		uint32(protocol.EGCCitadelClientMessages_k_EMsgClientToGCRequestReporterUpdates),
		req,
		uint32(protocol.EGCCitadelClientMessages_k_EMsgClientToGCRequestReporterUpdatesResponse),
		resp,
	)
}

// SendAcknowledgeReporterUpdates sends acknowledge reporter updates.
//
// Sends the GC message k_EMsgClientToGCAcknowledgeReporterUpdates (CMsgClientToGCAcknowledgeReporterUpdates). No response is tracked; any result
// arrives through the client event stream.
func (d *Deadlock) SendAcknowledgeReporterUpdates() {
	req := &protocol.CMsgClientToGCAcknowledgeReporterUpdates{}
	d.write(uint32(protocol.EGCCitadelClientMessages_k_EMsgClientToGCAcknowledgeReporterUpdates), req)
}

// SendBookUnlock purchases an unlock from the hero book. The expected_cost is
// checked server-side; the purchase fails if it does not match.
func (d *Deadlock) SendBookUnlock(
	ctx context.Context,
	bookID uint32,
	unlockID uint32,
	expectedCost uint32,
) (*protocol.CMsgClientToGCBookUnlockResponse, error) {
	req := &protocol.CMsgClientToGCBookUnlock{
		BookId:       &bookID,
		UnlockId:     &unlockID,
		ExpectedCost: &expectedCost,
	}
	resp := &protocol.CMsgClientToGCBookUnlockResponse{}

	return resp, d.MakeRequest(
		ctx,
		uint32(protocol.EGCCitadelClientMessages_k_EMsgClientToGCBookUnlock),
		req,
		uint32(protocol.EGCCitadelClientMessages_k_EMsgClientToGCBookUnlockResponse),
		resp,
	)
}

// SendCabalAnswerInvite sends a cabal answer invite.
//
// Sends the GC message k_EMsgClientToGCCabalAnswerInvite (CMsgClientToGCCabalAnswerInvite) and awaits the response k_EMsgClientToGCCabalAnswerInviteResponse,
// delivered as *CMsgClientToGCCabalAnswerInviteResponse.
func (d *Deadlock) SendCabalAnswerInvite(
	ctx context.Context,
	cabalID uint64,
	accepted bool,
) (*protocol.CMsgClientToGCCabalAnswerInviteResponse, error) {
	req := &protocol.CMsgClientToGCCabalAnswerInvite{
		CabalId:  &cabalID,
		Accepted: &accepted,
	}
	resp := &protocol.CMsgClientToGCCabalAnswerInviteResponse{}

	return resp, d.MakeRequest(
		ctx,
		uint32(protocol.EGCCitadelClientMessages_k_EMsgClientToGCCabalAnswerInvite),
		req,
		uint32(protocol.EGCCitadelClientMessages_k_EMsgClientToGCCabalAnswerInviteResponse),
		resp,
	)
}

// SendCabalInviteAccount sends a cabal invite account.
//
// Sends the GC message k_EMsgClientToGCCabalInviteAccount (CMsgClientToGCCabalInviteAccount) and awaits the response k_EMsgClientToGCCabalInviteAccountResponse,
// delivered as *CMsgClientToGCCabalInviteAccountResponse.
func (d *Deadlock) SendCabalInviteAccount(
	ctx context.Context,
	cabalID uint64,
	accountID uint32,
	cancelInvite bool,
) (*protocol.CMsgClientToGCCabalInviteAccountResponse, error) {
	req := &protocol.CMsgClientToGCCabalInviteAccount{
		CabalId:      &cabalID,
		AccountId:    &accountID,
		CancelInvite: &cancelInvite,
	}
	resp := &protocol.CMsgClientToGCCabalInviteAccountResponse{}

	return resp, d.MakeRequest(
		ctx,
		uint32(protocol.EGCCitadelClientMessages_k_EMsgClientToGCCabalInviteAccount),
		req,
		uint32(protocol.EGCCitadelClientMessages_k_EMsgClientToGCCabalInviteAccountResponse),
		resp,
	)
}

// SendChallengeActivate activates one of your generated challenge choices so
// it starts tracking progress.
func (d *Deadlock) SendChallengeActivate(
	ctx context.Context,
	challengeID uint64,
) (*protocol.CMsgClientToGCChallengeActivateResponse, error) {
	req := &protocol.CMsgClientToGCChallengeActivate{
		ChallengeId: &challengeID,
	}
	resp := &protocol.CMsgClientToGCChallengeActivateResponse{}

	return resp, d.MakeRequest(
		ctx,
		uint32(protocol.EGCCitadelClientMessages_k_EMsgClientToGCChallengeActivate),
		req,
		uint32(protocol.EGCCitadelClientMessages_k_EMsgClientToGCChallengeActivateResponse),
		resp,
	)
}

// SendCommendPlayerFromMatch sends a commend for a player from a match you
// played together. The GC delivers a CommendNotification event to the target.
func (d *Deadlock) SendCommendPlayerFromMatch(
	ctx context.Context,
	matchID uint64,
	targetAccountID uint32,
	commendType protocol.ECommendType,
	fakeCommendHeroID uint32,
) (*protocol.CMsgClientToGCCommendPlayerFromMatchResponse, error) {
	req := &protocol.CMsgClientToGCCommendPlayerFromMatch{
		MatchId:           &matchID,
		TargetAccountId:   &targetAccountID,
		CommendType:       &commendType,
		FakeCommendHeroId: &fakeCommendHeroID,
	}
	resp := &protocol.CMsgClientToGCCommendPlayerFromMatchResponse{}

	return resp, d.MakeRequest(
		ctx,
		uint32(protocol.EGCCitadelClientMessages_k_EMsgClientToGCCommendPlayerFromMatch),
		req,
		uint32(protocol.EGCCitadelClientMessages_k_EMsgClientToGCCommendPlayerFromMatchResponse),
		resp,
	)
}

// SendDeleteHeroBuild deletes one of your published hero builds. You must be
// the author of the build.
func (d *Deadlock) SendDeleteHeroBuild(
	ctx context.Context,
	authorAccountID uint32,
	heroBuildID uint32,
) (*protocol.CMsgClientToGCDeleteHeroBuildResponse, error) {
	req := &protocol.CMsgClientToGCDeleteHeroBuild{
		AuthorAccountId: &authorAccountID,
		HeroBuildId:     &heroBuildID,
	}
	resp := &protocol.CMsgClientToGCDeleteHeroBuildResponse{}

	return resp, d.MakeRequest(
		ctx,
		uint32(protocol.EGCCitadelClientMessages_k_EMsgClientToGCDeleteHeroBuild),
		req,
		uint32(protocol.EGCCitadelClientMessages_k_EMsgClientToGCDeleteHeroBuildResponse),
		resp,
	)
}

// SendDevAction runs a developer-only GC action. Developer accounts only.
func (d *Deadlock) SendDevAction(
	ctx context.Context,
	action protocol.CMsgClientToGCDevAction_EAction,
	accountID uint32,
	uintValue uint32,
	intValue int32,
	boolValue bool,
	strValue string,
	matchID uint64,
) (*protocol.CMsgClientToGCDevActionResponse, error) {
	req := &protocol.CMsgClientToGCDevAction{
		Action:    &action,
		AccountId: &accountID,
		UintValue: &uintValue,
		IntValue:  &intValue,
		BoolValue: &boolValue,
		StrValue:  &strValue,
		MatchId:   &matchID,
	}
	resp := &protocol.CMsgClientToGCDevActionResponse{}

	return resp, d.MakeRequest(
		ctx,
		uint32(protocol.EGCCitadelClientMessages_k_EMsgClientToGCDevAction),
		req,
		uint32(protocol.EGCCitadelClientMessages_k_EMsgClientToGCDevActionResponse),
		resp,
	)
}

// SendIsInMatchmaking checks whether the account is currently in the
// matchmaking queue. The response carries the queue state.
func (d *Deadlock) SendIsInMatchmaking(
	ctx context.Context,
) (*protocol.CMsgClientToGCIsInMatchmakingResponse, error) {
	req := &protocol.CMsgClientToGCIsInMatchmaking{}
	resp := &protocol.CMsgClientToGCIsInMatchmakingResponse{}

	return resp, d.MakeRequest(
		ctx,
		uint32(protocol.EGCCitadelClientMessages_k_EMsgClientToGCIsInMatchmaking),
		req,
		uint32(protocol.EGCCitadelClientMessages_k_EMsgClientToGCIsInMatchmakingResponse),
		resp,
	)
}

// SendPartyAction performs an action on a party member, such as kicking or
// promoting. See CMsgClientToGCPartyAction.EAction for the available actions;
// the response reports the outcome.
func (d *Deadlock) SendPartyAction(
	ctx context.Context,
	partyID uint64,
	targetAccountID uint32,
	actionID protocol.CMsgClientToGCPartyAction_EAction,
	uintValue uint64,
	boolValue bool,
	strValue string,
) (*protocol.CMsgClientToGCPartyActionResponse, error) {
	req := &protocol.CMsgClientToGCPartyAction{
		PartyId:         &partyID,
		TargetAccountId: &targetAccountID,
		ActionId:        &actionID,
		UintValue:       &uintValue,
		BoolValue:       &boolValue,
		StrValue:        &strValue,
	}
	resp := &protocol.CMsgClientToGCPartyActionResponse{}

	return resp, d.MakeRequest(
		ctx,
		uint32(protocol.EGCCitadelClientMessages_k_EMsgClientToGCPartyAction),
		req,
		uint32(protocol.EGCCitadelClientMessages_k_EMsgClientToGCPartyActionResponse),
		resp,
	)
}

// SendPartyChatMsg sends a chat line to your party members.
func (d *Deadlock) SendPartyChatMsg(
	ctx context.Context,
	partyID uint64,
	chatMsg string,
) (*protocol.CMsgClientToGCPartySendChatMsgResponse, error) {
	req := &protocol.CMsgClientToGCPartySendChatMsg{
		PartyId: &partyID,
		ChatMsg: &chatMsg,
	}
	resp := &protocol.CMsgClientToGCPartySendChatMsgResponse{}

	return resp, d.MakeRequest(
		ctx,
		uint32(protocol.EGCCitadelClientMessages_k_EMsgClientToGCPartySendChatMsg),
		req,
		uint32(protocol.EGCCitadelClientMessages_k_EMsgClientToGCPartySendChatMsgResponse),
		resp,
	)
}

// SendPartyInviteUser invites an account to your party. The invitee responds
// through the party invite flow.
func (d *Deadlock) SendPartyInviteUser(
	ctx context.Context,
	partyID uint64,
	inviteAccountID uint32,
) (*protocol.CMsgClientToGCPartyInviteUserResponse, error) {
	req := &protocol.CMsgClientToGCPartyInviteUser{
		PartyId:         &partyID,
		InviteAccountId: &inviteAccountID,
	}
	resp := &protocol.CMsgClientToGCPartyInviteUserResponse{}

	return resp, d.MakeRequest(
		ctx,
		uint32(protocol.EGCCitadelClientMessages_k_EMsgClientToGCPartyInviteUser),
		req,
		uint32(protocol.EGCCitadelClientMessages_k_EMsgClientToGCPartyInviteUserResponse),
		resp,
	)
}

// SendReplacementSDRTicket requests a replacement Steam Datagram Relay ticket
// for a match or lobby connection whose ticket expired or was lost.
func (d *Deadlock) SendReplacementSDRTicket(
	ctx context.Context,
	lobbyID uint64,
) (*protocol.CMsgClientToGCReplacementSDRTicketResponse, error) {
	req := &protocol.CMsgClientToGCReplacementSDRTicket{
		LobbyId: &lobbyID,
	}
	resp := &protocol.CMsgClientToGCReplacementSDRTicketResponse{}

	return resp, d.MakeRequest(
		ctx,
		uint32(protocol.EGCCitadelClientMessages_k_EMsgClientToGCReplacementSDRTicket),
		req,
		uint32(protocol.EGCCitadelClientMessages_k_EMsgClientToGCReplacementSDRTicketResponse),
		resp,
	)
}

// SendUnlockHero spends a hero unlock credit on the chosen hero.
func (d *Deadlock) SendUnlockHero(
	ctx context.Context,
	heroIDs []uint32,
	heroChoiceID uint32,
) (*protocol.CMsgClientToGCUnlockHeroResponse, error) {
	req := &protocol.CMsgClientToGCUnlockHero{
		HeroIds:      heroIDs,
		HeroChoiceId: &heroChoiceID,
	}
	resp := &protocol.CMsgClientToGCUnlockHeroResponse{}

	return resp, d.MakeRequest(
		ctx,
		uint32(protocol.EGCCitadelClientMessages_k_EMsgClientToGCUnlockHero),
		req,
		uint32(protocol.EGCCitadelClientMessages_k_EMsgClientToGCUnlockHeroResponse),
		resp,
	)
}

// SendUpdateAccountSync pushes updated CSOAccountSyncStorage values to the
// GC. The response confirms the sync.
func (d *Deadlock) SendUpdateAccountSync(
	ctx context.Context,
	ids []uint32,
	values []uint32,
) (*protocol.CMsgClientToGCUpdateAccountSyncResponse, error) {
	req := &protocol.CMsgClientToGCUpdateAccountSync{
		Ids:    ids,
		Values: values,
	}
	resp := &protocol.CMsgClientToGCUpdateAccountSyncResponse{}

	return resp, d.MakeRequest(
		ctx,
		uint32(protocol.EGCCitadelClientMessages_k_EMsgClientToGCUpdateAccountSync),
		req,
		uint32(protocol.EGCCitadelClientMessages_k_EMsgClientToGCUpdateAccountSyncResponse),
		resp,
	)
}

// SendUpdateHeroBuild creates or updates one of your published hero builds.
func (d *Deadlock) SendUpdateHeroBuild(
	ctx context.Context,
	heroBuild protocol.CMsgHeroBuild,
) (*protocol.CMsgClientToGCUpdateHeroBuildResponse, error) {
	req := &protocol.CMsgClientToGCUpdateHeroBuild{
		HeroBuild: &heroBuild,
	}
	resp := &protocol.CMsgClientToGCUpdateHeroBuildResponse{}

	return resp, d.MakeRequest(
		ctx,
		uint32(protocol.EGCCitadelClientMessages_k_EMsgClientToGCUpdateHeroBuild),
		req,
		uint32(protocol.EGCCitadelClientMessages_k_EMsgClientToGCUpdateHeroBuildResponse),
		resp,
	)
}

// SendUpdateHeroBuildPreference sets which published build you follow for a
// hero.
func (d *Deadlock) SendUpdateHeroBuildPreference(
	ctx context.Context,
	heroBuildID uint32,
	preference protocol.CMsgHeroBuildPreference,
) (*protocol.CMsgClientToGCUpdateHeroBuildPreferenceResponse, error) {
	req := &protocol.CMsgClientToGCUpdateHeroBuildPreference{
		HeroBuildId: &heroBuildID,
		Preference:  &preference,
	}
	resp := &protocol.CMsgClientToGCUpdateHeroBuildPreferenceResponse{}

	return resp, d.MakeRequest(
		ctx,
		uint32(protocol.EGCCitadelClientMessages_k_EMsgClientToGCUpdateHeroBuildPreference),
		req,
		uint32(protocol.EGCCitadelClientMessages_k_EMsgClientToGCUpdateHeroBuildPreferenceResponse),
		resp,
	)
}

// SendUpdateRoster updates your hero selection roster before entering a
// matchmaking queue.
func (d *Deadlock) SendUpdateRoster(
	ctx context.Context,
	heroes protocol.CMsgHeroSelectionMatchInfo,
	gameMode protocol.ECitadelGameMode,
	matchMode protocol.ECitadelMatchMode,
) (*protocol.CMsgClientToGCUpdateRosterResponse, error) {
	req := &protocol.CMsgClientToGCUpdateRoster{
		Heroes:    &heroes,
		GameMode:  &gameMode,
		MatchMode: &matchMode,
	}
	resp := &protocol.CMsgClientToGCUpdateRosterResponse{}

	return resp, d.MakeRequest(
		ctx,
		uint32(protocol.EGCCitadelClientMessages_k_EMsgClientToGCUpdateRoster),
		req,
		uint32(protocol.EGCCitadelClientMessages_k_EMsgClientToGCUpdateRosterResponse),
		resp,
	)
}

// SendUpdateSpectatorStatus tells the GC which lobby you are spectating, or
// clears it with stopped_spectating.
func (d *Deadlock) SendUpdateSpectatorStatus(
	spectatingLobbyID uint64,
	stoppedSpectating bool,
) {
	req := &protocol.CMsgClientToGCUpdateSpectatorStatus{
		SpectatingLobbyId: &spectatingLobbyID,
		StoppedSpectating: &stoppedSpectating,
	}
	d.write(uint32(protocol.EGCCitadelClientMessages_k_EMsgClientToGCUpdateSpectatorStatus), req)
}

// SendVariantItemChangeSlotStyles sends variant item change slot styles.
//
// Sends the GC message k_EMsgClientToGCVariantItemChangeSlotStyles (CMsgClientToGCVariantItemChangeSlotStyles) and awaits the response k_EMsgClientToGCVariantItemChangeSlotStylesResponse,
// delivered as *CMsgClientToGCVariantItemChangeSlotStylesResponse.
func (d *Deadlock) SendVariantItemChangeSlotStyles(
	ctx context.Context,
	itemID uint64,
	slotStyleActions []*protocol.CMsgClientToGCVariantItemChangeSlotStyles_SlotStyleAction,
) (*protocol.CMsgClientToGCVariantItemChangeSlotStylesResponse, error) {
	req := &protocol.CMsgClientToGCVariantItemChangeSlotStyles{
		ItemId:           &itemID,
		SlotStyleActions: slotStyleActions,
	}
	resp := &protocol.CMsgClientToGCVariantItemChangeSlotStylesResponse{}

	return resp, d.MakeRequest(
		ctx,
		uint32(protocol.EGCCitadelClientMessages_k_EMsgClientToGCVariantItemChangeSlotStyles),
		req,
		uint32(protocol.EGCCitadelClientMessages_k_EMsgClientToGCVariantItemChangeSlotStylesResponse),
		resp,
	)
}

// SetAccountPrivacySetting changes who can view your profile card and
// stats.
func (d *Deadlock) SetAccountPrivacySetting(
	ctx context.Context,
	privacySetting protocol.CSOGameAccountClient_EPrivacySetting,
) (*protocol.CMsgClientToGCSetAccountPrivacySettingResponse, error) {
	req := &protocol.CMsgClientToGCSetAccountPrivacySetting{
		PrivacySetting: &privacySetting,
	}
	resp := &protocol.CMsgClientToGCSetAccountPrivacySettingResponse{}

	return resp, d.MakeRequest(
		ctx,
		uint32(protocol.EGCCitadelClientMessages_k_EMsgClientToGCSetAccountPrivacySetting),
		req,
		uint32(protocol.EGCCitadelClientMessages_k_EMsgClientToGCSetAccountPrivacySettingResponse),
		resp,
	)
}

// SetChallengeGenerate generates a fresh set of challenge choices to replace
// your current ones.
func (d *Deadlock) SetChallengeGenerate(
	ctx context.Context,
) (*protocol.CMsgClientToGCGenerateChallengeSetResponse, error) {
	req := &protocol.CMsgClientToGCGenerateChallengeSet{}
	resp := &protocol.CMsgClientToGCGenerateChallengeSetResponse{}

	return resp, d.MakeRequest(
		ctx,
		uint32(protocol.EGCCitadelClientMessages_k_EMsgClientToGCGenerateChallengeSet),
		req,
		uint32(protocol.EGCCitadelClientMessages_k_EMsgClientToGCGenerateChallengeSetResponse),
		resp,
	)
}

// SetDevMMBias biases matchmaking for an account toward a skill value.
// Developer accounts only.
func (d *Deadlock) SetDevMMBias(
	accountID uint32,
	value uint32,
) {
	req := &protocol.CMsgClientToGCDevSetMMBias{
		AccountId: &accountID,
		Value:     &value,
	}
	d.write(uint32(protocol.EGCCitadelClientMessages_k_EMsgClientToGCDevSetMMBias), req)
}

// SetNewPlayerProgress marks a new-player progression milestone complete,
// such as finishing the tutorial or unlocking ranked play.
func (d *Deadlock) SetNewPlayerProgress(
	ctx context.Context,
	flag protocol.ECitadelNewPlayerProgressFlag,
) (*protocol.CMsgClientToGCSetNewPlayerProgressResponse, error) {
	req := &protocol.CMsgClientToGCSetNewPlayerProgress{
		Flag: &flag,
	}
	resp := &protocol.CMsgClientToGCSetNewPlayerProgressResponse{}

	return resp, d.MakeRequest(
		ctx,
		uint32(protocol.EGCCitadelClientMessages_k_EMsgClientToGCSetNewPlayerProgress),
		req,
		uint32(protocol.EGCCitadelClientMessages_k_EMsgClientToGCSetNewPlayerProgressResponse),
		resp,
	)
}

// SetPartyMode configures the party's matchmaking settings: match mode, game
// mode, bot difficulty, and region.
func (d *Deadlock) SetPartyMode(
	ctx context.Context,
	partyID uint64,
	matchMode protocol.ECitadelMatchMode,
	gameMode protocol.ECitadelGameMode,
	botDifficulty protocol.ECitadelBotDifficulty,
	devServerCommand string,
	regionMode protocol.ECitadelRegionMode,
) (*protocol.CMsgClientToGCPartySetModeResponse, error) {
	req := &protocol.CMsgClientToGCPartySetMode{
		PartyId:          &partyID,
		MatchMode:        &matchMode,
		GameMode:         &gameMode,
		BotDifficulty:    &botDifficulty,
		DevServerCommand: &devServerCommand,
		RegionMode:       &regionMode,
	}
	resp := &protocol.CMsgClientToGCPartySetModeResponse{}

	return resp, d.MakeRequest(
		ctx,
		uint32(protocol.EGCCitadelClientMessages_k_EMsgClientToGCPartySetMode),
		req,
		uint32(protocol.EGCCitadelClientMessages_k_EMsgClientToGCPartySetModeResponse),
		resp,
	)
}

// SetPartyReadyState marks you ready (or not ready) for the party's next
// match, along with the hero roster you intend to play.
func (d *Deadlock) SetPartyReadyState(
	ctx context.Context,
	partyID uint64,
	readyState bool,
	heroRoster protocol.CMsgHeroSelectionMatchInfo,
) (*protocol.CMsgClientToGCPartySetReadyStateResponse, error) {
	req := &protocol.CMsgClientToGCPartySetReadyState{
		PartyId:    &partyID,
		ReadyState: &readyState,
		HeroRoster: &heroRoster,
	}
	resp := &protocol.CMsgClientToGCPartySetReadyStateResponse{}

	return resp, d.MakeRequest(
		ctx,
		uint32(protocol.EGCCitadelClientMessages_k_EMsgClientToGCPartySetReadyState),
		req,
		uint32(protocol.EGCCitadelClientMessages_k_EMsgClientToGCPartySetReadyStateResponse),
		resp,
	)
}

// SetServerConVar sets a console variable on a practice lobby server.
// Developer accounts only.
func (d *Deadlock) SetServerConVar(
	ctx context.Context,
	convarName string,
	convarValue string,
	lobbyID uint64,
) (*protocol.CMsgClientToGCSetServerConVarResponse, error) {
	req := &protocol.CMsgClientToGCSetServerConVar{
		ConvarName:  &convarName,
		ConvarValue: &convarValue,
		LobbyId:     &lobbyID,
	}
	resp := &protocol.CMsgClientToGCSetServerConVarResponse{}

	return resp, d.MakeRequest(
		ctx,
		uint32(protocol.EGCCitadelClientMessages_k_EMsgClientToGCSetServerConVar),
		req,
		uint32(protocol.EGCCitadelClientMessages_k_EMsgClientToGCSetServerConVarResponse),
		resp,
	)
}

// SpectateLobby requests to spectate a lobby. The response carries the
// connection details for the spectator broadcast.
func (d *Deadlock) SpectateLobby(
	ctx context.Context,
	lobbyID uint64,
	clientPlatform protocol.EGCPlatform,
	matchID uint64,
) (*protocol.CMsgClientToGCSpectateLobbyResponse, error) {
	req := &protocol.CMsgClientToGCSpectateLobby{
		LobbyId:        &lobbyID,
		ClientPlatform: &clientPlatform,
		MatchId:        &matchID,
	}
	resp := &protocol.CMsgClientToGCSpectateLobbyResponse{}

	return resp, d.MakeRequest(
		ctx,
		uint32(protocol.EGCCitadelClientMessages_k_EMsgClientToGCSpectateLobby),
		req,
		uint32(protocol.EGCCitadelClientMessages_k_EMsgClientToGCSpectateLobbyResponse),
		resp,
	)
}

// SpectateUser requests to spectate the live game of an account. The response
// carries the connection details, or an error if no game is running.
func (d *Deadlock) SpectateUser(
	ctx context.Context,
	spectateAccountID uint32,
	clientPlatform protocol.EGCPlatform,
) (*protocol.CMsgClientToGCSpectateUserResponse, error) {
	req := &protocol.CMsgClientToGCSpectateUser{
		SpectateAccountId: &spectateAccountID,
		ClientPlatform:    &clientPlatform,
	}
	resp := &protocol.CMsgClientToGCSpectateUserResponse{}

	return resp, d.MakeRequest(
		ctx,
		uint32(protocol.EGCCitadelClientMessages_k_EMsgClientToGCSpectateUser),
		req,
		uint32(protocol.EGCCitadelClientMessages_k_EMsgClientToGCSpectateUserResponse),
		resp,
	)
}

// StartMatchmaking enters the matchmaking queue with the given match info,
// region ping times, and hero roster. The response confirms the queue state.
// Watch for MatchmakingStopped events for asynchronous cancellation.
func (d *Deadlock) StartMatchmaking(
	ctx context.Context,
	clientPlatform protocol.EGCPlatform,
	matchInfo protocol.CMsgStartFindingMatchInfo,
	pingTimes protocol.CMsgRegionPingTimesClient,
	heroes protocol.CMsgHeroSelectionMatchInfo,
	pgiVerified bool,
) (*protocol.CMsgClientToGCStartMatchmakingResponse, error) {
	req := &protocol.CMsgClientToGCStartMatchmaking{
		ClientPlatform: &clientPlatform,
		MatchInfo:      &matchInfo,
		PingTimes:      &pingTimes,
		Heroes:         &heroes,
		PgiVerified:    &pgiVerified,
	}
	resp := &protocol.CMsgClientToGCStartMatchmakingResponse{}

	return resp, d.MakeRequest(
		ctx,
		uint32(protocol.EGCCitadelClientMessages_k_EMsgClientToGCStartMatchmaking),
		req,
		uint32(protocol.EGCCitadelClientMessages_k_EMsgClientToGCStartMatchmakingResponse),
		resp,
	)
}

// StartPartyMatch starts the search for a match on behalf of the whole
// party. Party leaders only.
func (d *Deadlock) StartPartyMatch(
	ctx context.Context,
	partyID uint64,
) (*protocol.CMsgClientToGCPartyStartMatchResponse, error) {
	req := &protocol.CMsgClientToGCPartyStartMatch{
		PartyId: &partyID,
	}
	resp := &protocol.CMsgClientToGCPartyStartMatchResponse{}

	return resp, d.MakeRequest(
		ctx,
		uint32(protocol.EGCCitadelClientMessages_k_EMsgClientToGCPartyStartMatch),
		req,
		uint32(protocol.EGCCitadelClientMessages_k_EMsgClientToGCPartyStartMatchResponse),
		resp,
	)
}

// StartRankedInterval begins a ranked calibration interval. The response
// reports the new ranked state.
func (d *Deadlock) StartRankedInterval(
	ctx context.Context,
	rankType protocol.ECitadelRankedType,
	interval uint32,
) (*protocol.CMsgClientToGCStartRankedIntervalResponse, error) {
	req := &protocol.CMsgClientToGCStartRankedInterval{
		RankType: &rankType,
		Interval: &interval,
	}
	resp := &protocol.CMsgClientToGCStartRankedIntervalResponse{}

	return resp, d.MakeRequest(
		ctx,
		uint32(protocol.EGCCitadelClientMessages_k_EMsgClientToGCStartRankedInterval),
		req,
		uint32(protocol.EGCCitadelClientMessages_k_EMsgClientToGCStartRankedIntervalResponse),
		resp,
	)
}

// StopMatchmaking leaves the matchmaking queue. A MatchmakingStopped event
// confirms the exit.
func (d *Deadlock) StopMatchmaking(
	ctx context.Context,
) (*protocol.CMsgClientToGCStopMatchmakingResponse, error) {
	req := &protocol.CMsgClientToGCStopMatchmaking{}
	resp := &protocol.CMsgClientToGCStopMatchmakingResponse{}

	return resp, d.MakeRequest(
		ctx,
		uint32(protocol.EGCCitadelClientMessages_k_EMsgClientToGCStopMatchmaking),
		req,
		uint32(protocol.EGCCitadelClientMessages_k_EMsgClientToGCStopMatchmakingResponse),
		resp,
	)
}

// SubmitHeroReleaseVote casts a vote in the community hero release vote for
// the given round.
func (d *Deadlock) SubmitHeroReleaseVote(
	voteRound uint32,
	voteOption uint32,
) {
	req := &protocol.CMsgClientToGCSubmitHeroReleaseVote{
		VoteRound:  &voteRound,
		VoteOption: &voteOption,
	}
	d.write(uint32(protocol.EGCCitadelClientMessages_k_EMsgClientToGCSubmitHeroReleaseVote), req)
}

// SubmitPlaytestUser registers an account for a playtest session at the
// given location. Developer accounts only.
func (d *Deadlock) SubmitPlaytestUser(
	ctx context.Context,
	location string,
	targetAccountID uint32,
) (*protocol.CMsgClientToGCSubmitPlaytestUserResponse, error) {
	req := &protocol.CMsgClientToGCSubmitPlaytestUser{
		Location:        &location,
		TargetAccountId: &targetAccountID,
	}
	resp := &protocol.CMsgClientToGCSubmitPlaytestUserResponse{}

	return resp, d.MakeRequest(
		ctx,
		uint32(protocol.EGCCitadelClientMessages_k_EMsgClientToGCSubmitPlaytestUser),
		req,
		uint32(protocol.EGCCitadelClientMessages_k_EMsgClientToGCSubmitPlaytestUserResponse),
		resp,
	)
}

// SubmitSurvey answers a survey question previously delivered through
// GetSurveyQuestion. Set is_skip to decline without answering.
func (d *Deadlock) SubmitSurvey(
	ctx context.Context,
	questionID uint32,
	isSkip bool,
	responseValue uint32,
	devForceNewQuestionType protocol.CMsgSurveyQuestion_EQuestionType,
) (*protocol.CMsgClientToGCSubmitSurveyResponse, error) {
	req := &protocol.CMsgClientToGCSubmitSurvey{
		QuestionId:              &questionID,
		IsSkip:                  &isSkip,
		ResponseValue:           &responseValue,
		DevForceNewQuestionType: &devForceNewQuestionType,
	}
	resp := &protocol.CMsgClientToGCSubmitSurveyResponse{}

	return resp, d.MakeRequest(
		ctx,
		uint32(protocol.EGCCitadelClientMessages_k_EMsgClientToGCSubmitSurvey),
		req,
		uint32(protocol.EGCCitadelClientMessages_k_EMsgClientToGCSubmitSurveyResponse),
		resp,
	)
}

// registerGeneratedHandlers registers the auto-generated event handlers.
func (d *Deadlock) registerGeneratedHandlers() {
	d.handlers[uint32(protocol.EGCCitadelClientMessages_k_EMsgGCToClientBookUpdated)] = d.getEventEmitter(func() events.Event {
		return &events.BookUpdated{}
	})
	d.handlers[uint32(protocol.EGCCitadelClientMessages_k_EMsgGCToClientCabalUpdated)] = d.getEventEmitter(func() events.Event {
		return &events.CabalUpdated{}
	})
	d.handlers[uint32(protocol.EGCCitadelClientMessages_k_EMsgGCToClientCanRejoinParty)] = d.getEventEmitter(func() events.Event {
		return &events.CanRejoinParty{}
	})
	d.handlers[uint32(protocol.EGCCitadelClientMessages_k_EMsgGCToClientCommendNotification)] = d.getEventEmitter(func() events.Event {
		return &events.CommendNotification{}
	})
	d.handlers[uint32(protocol.EGCCitadelClientMessages_k_EMsgGCToClientDevPlaytestStatus)] = d.getEventEmitter(func() events.Event {
		return &events.DevPlaytestStatus{}
	})
	d.handlers[uint32(protocol.EGCCitadelClientMessages_k_EMsgGCToClientMatchmakingStopped)] = d.getEventEmitter(func() events.Event {
		return &events.MatchmakingStopped{}
	})
	d.handlers[uint32(protocol.EGCCitadelClientMessages_k_EMsgGCToClientPartyChatMsg)] = d.getEventEmitter(func() events.Event {
		return &events.PartyChatMsg{}
	})
	d.handlers[uint32(protocol.EGCCitadelClientMessages_k_EMsgGCToClientPartyEvent)] = d.getEventEmitter(func() events.Event {
		return &events.PartyEvent{}
	})
	d.handlers[uint32(protocol.EGCCitadelClientMessages_k_EMsgGCToClientPartyPlayerTyping)] = d.getEventEmitter(func() events.Event {
		return &events.PartyPlayerTyping{}
	})
	d.handlers[uint32(protocol.EGCCitadelClientMessages_k_EMsgGCToClientProfileCardUpdated)] = d.getEventEmitter(func() events.Event {
		return &events.ProfileCardUpdated{}
	})
	d.handlers[uint32(protocol.EGCCitadelClientMessages_k_EMsgGCToClientRankDataUpdate)] = d.getEventEmitter(func() events.Event {
		return &events.RankDataUpdate{}
	})
	d.handlers[uint32(protocol.EGCCitadelClientMessages_k_EMsgGCToClientSDRTicket)] = d.getEventEmitter(func() events.Event {
		return &events.SDRTicket{}
	})
	d.handlers[uint32(protocol.EGCCitadelClientMessages_k_EMsgGCToClientUpdateHeroReleaseVoteTally)] = d.getEventEmitter(func() events.Event {
		return &events.UpdateHeroReleaseVoteTally{}
	})
}
