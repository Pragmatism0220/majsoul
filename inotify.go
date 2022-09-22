package majsoul

import "github.com/Pragmatism0220/majsoul/message"

// IFNotify
// 雀魂proto协议中缺少描述监听消息的接口
// 故添加该接口，可能会丢失某些api
type IFNotify interface {
	NotifyCaptcha(notify *message.NotifyCaptcha)
	NotifyRoomGameStart(notify *message.NotifyRoomGameStart)
	NotifyMatchGameStart(notify *message.NotifyMatchGameStart)
	NotifyRoomPlayerReady(notify *message.NotifyRoomPlayerReady)
	NotifyRoomPlayerDressing(notify *message.NotifyRoomPlayerDressing)
	NotifyRoomPlayerUpdate(notify *message.NotifyRoomPlayerUpdate)
	NotifyRoomKickOut(notify *message.NotifyRoomKickOut)
	NotifyFriendStateChange(notify *message.NotifyFriendStateChange)
	NotifyFriendViewChange(notify *message.NotifyFriendViewChange)
	NotifyFriendChange(notify *message.NotifyFriendChange)
	NotifyNewFriendApply(notify *message.NotifyNewFriendApply)
	NotifyClientMessage(notify *message.NotifyClientMessage)
	NotifyAccountUpdate(notify *message.NotifyAccountUpdate)
	NotifyAnotherLogin(notify *message.NotifyAnotherLogin)
	NotifyAccountLogout(notify *message.NotifyAccountLogout)
	NotifyAnnouncementUpdate(notify *message.NotifyAnnouncementUpdate)
	NotifyNewMail(notify *message.NotifyNewMail)
	NotifyDeleteMail(notify *message.NotifyDeleteMail)
	NotifyReviveCoinUpdate(notify *message.NotifyReviveCoinUpdate)
	NotifyDailyTaskUpdate(notify *message.NotifyDailyTaskUpdate)
	NotifyActivityTaskUpdate(notify *message.NotifyActivityTaskUpdate)
	NotifyActivityPeriodTaskUpdate(notify *message.NotifyActivityPeriodTaskUpdate)
	NotifyAccountRandomTaskUpdate(notify *message.NotifyAccountRandomTaskUpdate)
	NotifyActivitySegmentTaskUpdate(notify *message.NotifyActivitySegmentTaskUpdate)
	NotifyActivityUpdate(notify *message.NotifyActivityUpdate)
	NotifyAccountChallengeTaskUpdate(notify *message.NotifyAccountChallengeTaskUpdate)
	NotifyNewComment(notify *message.NotifyNewComment)
	NotifyRollingNotice(notify *message.NotifyRollingNotice)
	NotifyGiftSendRefresh(notify *message.NotifyGiftSendRefresh)
	NotifyShopUpdate(notify *message.NotifyShopUpdate)
	NotifyVipLevelChange(notify *message.NotifyVipLevelChange)
	NotifyServerSetting(notify *message.NotifyServerSetting)
	NotifyPayResult(notify *message.NotifyPayResult)
	NotifyCustomContestAccountMsg(notify *message.NotifyCustomContestAccountMsg)
	NotifyCustomContestSystemMsg(notify *message.NotifyCustomContestSystemMsg)
	NotifyMatchTimeout(notify *message.NotifyMatchTimeout)
	NotifyCustomContestState(notify *message.NotifyCustomContestState)
	NotifyActivityChange(notify *message.NotifyActivityChange)
	NotifyAFKResult(notify *message.NotifyAFKResult)
	NotifyGameFinishRewardV2(notify *message.NotifyGameFinishRewardV2)
	NotifyActivityRewardV2(notify *message.NotifyActivityRewardV2)
	NotifyActivityPointV2(notify *message.NotifyActivityPointV2)
	NotifyLeaderboardPointV2(notify *message.NotifyLeaderboardPointV2)
	NotifyNewGame(notify *message.NotifyNewGame)
	NotifyPlayerLoadGameReady(notify *message.NotifyPlayerLoadGameReady)
	NotifyGameBroadcast(notify *message.NotifyGameBroadcast)
	NotifyGameEndResult(notify *message.NotifyGameEndResult)
	NotifyGameTerminate(notify *message.NotifyGameTerminate)
	NotifyPlayerConnectionState(notify *message.NotifyPlayerConnectionState)
	NotifyAccountLevelChange(notify *message.NotifyAccountLevelChange)
	NotifyGameFinishReward(notify *message.NotifyGameFinishReward)
	NotifyActivityReward(notify *message.NotifyActivityReward)
	NotifyActivityPoint(notify *message.NotifyActivityPoint)
	NotifyLeaderboardPoint(notify *message.NotifyLeaderboardPoint)
	NotifyGamePause(notify *message.NotifyGamePause)
	NotifyEndGameVote(notify *message.NotifyEndGameVote)
	NotifyObserveData(notify *message.NotifyObserveData)
	NotifyRoomPlayerReady_AccountReadyState(notify *message.NotifyRoomPlayerReady_AccountReadyState)
	NotifyRoomPlayerDressing_AccountDressingState(notify *message.NotifyRoomPlayerDressing_AccountDressingState)
	NotifyAnnouncementUpdate_AnnouncementUpdate(notify *message.NotifyAnnouncementUpdate_AnnouncementUpdate)
	NotifyActivityUpdate_FeedActivityData(notify *message.NotifyActivityUpdate_FeedActivityData)
	NotifyActivityUpdate_FeedActivityData_CountWithTimeData(notify *message.NotifyActivityUpdate_FeedActivityData_CountWithTimeData)
	NotifyActivityUpdate_FeedActivityData_GiftBoxData(notify *message.NotifyActivityUpdate_FeedActivityData_GiftBoxData)
	NotifyPayResult_ResourceModify(notify *message.NotifyPayResult_ResourceModify)
	NotifyGameFinishRewardV2_LevelChange(notify *message.NotifyGameFinishRewardV2_LevelChange)
	NotifyGameFinishRewardV2_MatchChest(notify *message.NotifyGameFinishRewardV2_MatchChest)
	NotifyGameFinishRewardV2_MainCharacter(notify *message.NotifyGameFinishRewardV2_MainCharacter)
	NotifyGameFinishRewardV2_CharacterGift(notify *message.NotifyGameFinishRewardV2_CharacterGift)
	NotifyActivityRewardV2_ActivityReward(notify *message.NotifyActivityRewardV2_ActivityReward)
	NotifyActivityPointV2_ActivityPoint(notify *message.NotifyActivityPointV2_ActivityPoint)
	NotifyLeaderboardPointV2_LeaderboardPoint(notify *message.NotifyLeaderboardPointV2_LeaderboardPoint)
	NotifyGameFinishReward_LevelChange(notify *message.NotifyGameFinishReward_LevelChange)
	NotifyGameFinishReward_MatchChest(notify *message.NotifyGameFinishReward_MatchChest)
	NotifyGameFinishReward_MainCharacter(notify *message.NotifyGameFinishReward_MainCharacter)
	NotifyGameFinishReward_CharacterGift(notify *message.NotifyGameFinishReward_CharacterGift)
	NotifyActivityReward_ActivityReward(notify *message.NotifyActivityReward_ActivityReward)
	NotifyActivityPoint_ActivityPoint(notify *message.NotifyActivityPoint_ActivityPoint)
	NotifyLeaderboardPoint_LeaderboardPoint(notify *message.NotifyLeaderboardPoint_LeaderboardPoint)
	NotifyEndGameVote_VoteResult(notify *message.NotifyEndGameVote_VoteResult)
	PlayerLeaving(notify *message.PlayerLeaving)
	ActionPrototype(notify *message.ActionPrototype)
}
