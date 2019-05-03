package longpoll // import "github.com/severecloud/vksdk/5.92/longpoll-bot"

import (
	"log"
	"strconv"
	"sync/atomic"

	"github.com/severecloud/vksdk/5.92/object"

	"github.com/severecloud/vksdk/5.92/api"
	"github.com/severecloud/vksdk/5.92/handler"
)

type longpollResponse struct {
	Ts      string              `json:"ts"`
	Updates []object.GroupEvent `json:"updates"`
	Failed  int                 `json:"failed"`
}

// Longpoll struct
type Longpoll struct {
	GroupID  int
	Server   string
	Key      string
	Ts       string
	Wait     int
	VK       api.VK
	FuncList handler.FuncList

	inShutdown int32
}

// Init Longpoll
func Init(vk api.VK, groupID int) (lp Longpoll) {
	lp.VK = vk
	lp.GroupID = groupID
	lp.updateServer(true)
	return
}

func (lp *Longpoll) updateServer(updateTs bool) {
	params := map[string]string{
		"group_id": strconv.Itoa(lp.GroupID),
	}
	serverSetting, vkErr := lp.VK.GroupsGetLongPollServer(params)
	if vkErr.Code != 0 {
		log.Fatal("F")
		// TODO: return error
	}

	lp.Key = serverSetting.Key
	lp.Server = serverSetting.Server
	if updateTs {
		lp.Ts = serverSetting.Ts
	}
}

func (lp *Longpoll) check() []object.GroupEvent {
	var response longpollResponse
	// TODO: request
	// TODO: iferr
	switch response.Failed {
	case 0:
		lp.Ts = response.Ts
	case 1:
		//TODO: log it
		lp.Ts = response.Ts
	case 2:
		//TODO: log it
		lp.updateServer(false)
	case 3:
		//TODO: log it
		lp.updateServer(true)
	default:
		//TODO: log it
		lp.updateServer(true)
	}
	return response.Updates
}

// Run handler
func (lp *Longpoll) Run() {
	atomic.StoreInt32(&lp.inShutdown, 0)
	for atomic.LoadInt32(&lp.inShutdown) == 0 {
		events := lp.check()
		for _, event := range events {
			lp.FuncList.Handler(event)
		}
	}
}

// Shutdown handler
func (lp *Longpoll) Shutdown() {
	atomic.StoreInt32(&lp.inShutdown, 1)
}

// MessageNew handler
func (lp *Longpoll) MessageNew(f object.MessageNewFunc) {
	lp.FuncList.MessageNew = append(lp.FuncList.MessageNew, f)
}

// MessageReply handler
func (lp *Longpoll) MessageReply(f object.MessageReplyFunc) {
	lp.FuncList.MessageReply = append(lp.FuncList.MessageReply, f)
}

// MessageEdit handler
func (lp *Longpoll) MessageEdit(f object.MessageEditFunc) {
	lp.FuncList.MessageEdit = append(lp.FuncList.MessageEdit, f)
}

// MessageAllow handler
func (lp *Longpoll) MessageAllow(f object.MessageAllowFunc) {
	lp.FuncList.MessageAllow = append(lp.FuncList.MessageAllow, f)
}

// MessageDeny handler
func (lp *Longpoll) MessageDeny(f object.MessageDenyFunc) {
	lp.FuncList.MessageDeny = append(lp.FuncList.MessageDeny, f)
}

// MessageTypingState handler
func (lp *Longpoll) MessageTypingState(f object.MessageTypingStateFunc) {
	lp.FuncList.MessageTypingState = append(lp.FuncList.MessageTypingState, f)
}

// PhotoNew handler
func (lp *Longpoll) PhotoNew(f object.PhotoNewFunc) {
	lp.FuncList.PhotoNew = append(lp.FuncList.PhotoNew, f)
}

// PhotoCommentNew handler
func (lp *Longpoll) PhotoCommentNew(f object.PhotoCommentNewFunc) {
	lp.FuncList.PhotoCommentNew = append(lp.FuncList.PhotoCommentNew, f)
}

// PhotoCommentEdit handler
func (lp *Longpoll) PhotoCommentEdit(f object.PhotoCommentEditFunc) {
	lp.FuncList.PhotoCommentEdit = append(lp.FuncList.PhotoCommentEdit, f)
}

// PhotoCommentRestore handler
func (lp *Longpoll) PhotoCommentRestore(f object.PhotoCommentRestoreFunc) {
	lp.FuncList.PhotoCommentRestore = append(lp.FuncList.PhotoCommentRestore, f)
}

// PhotoCommentDelete handler
func (lp *Longpoll) PhotoCommentDelete(f object.PhotoCommentDeleteFunc) {
	lp.FuncList.PhotoCommentDelete = append(lp.FuncList.PhotoCommentDelete, f)
}

// AudioNew handler
func (lp *Longpoll) AudioNew(f object.AudioNewFunc) {
	lp.FuncList.AudioNew = append(lp.FuncList.AudioNew, f)
}

// VideoNew handler
func (lp *Longpoll) VideoNew(f object.VideoNewFunc) {
	lp.FuncList.VideoNew = append(lp.FuncList.VideoNew, f)
}

// VideoCommentNew handler
func (lp *Longpoll) VideoCommentNew(f object.VideoCommentNewFunc) {
	lp.FuncList.VideoCommentNew = append(lp.FuncList.VideoCommentNew, f)
}

// VideoCommentEdit handler
func (lp *Longpoll) VideoCommentEdit(f object.VideoCommentEditFunc) {
	lp.FuncList.VideoCommentEdit = append(lp.FuncList.VideoCommentEdit, f)
}

// VideoCommentRestore handler
func (lp *Longpoll) VideoCommentRestore(f object.VideoCommentRestoreFunc) {
	lp.FuncList.VideoCommentRestore = append(lp.FuncList.VideoCommentRestore, f)
}

// VideoCommentDelete handler
func (lp *Longpoll) VideoCommentDelete(f object.VideoCommentDeleteFunc) {
	lp.FuncList.VideoCommentDelete = append(lp.FuncList.VideoCommentDelete, f)
}

// WallPostNew handler
func (lp *Longpoll) WallPostNew(f object.WallPostNewFunc) {
	lp.FuncList.WallPostNew = append(lp.FuncList.WallPostNew, f)
}

// WallRepost handler
func (lp *Longpoll) WallRepost(f object.WallRepostFunc) {
	lp.FuncList.WallRepost = append(lp.FuncList.WallRepost, f)
}

// WallReplyNew handler
func (lp *Longpoll) WallReplyNew(f object.WallReplyNewFunc) {
	lp.FuncList.WallReplyNew = append(lp.FuncList.WallReplyNew, f)
}

// WallReplyEdit handler
func (lp *Longpoll) WallReplyEdit(f object.WallReplyEditFunc) {
	lp.FuncList.WallReplyEdit = append(lp.FuncList.WallReplyEdit, f)
}

// WallReplyRestore handler
func (lp *Longpoll) WallReplyRestore(f object.WallReplyRestoreFunc) {
	lp.FuncList.WallReplyRestore = append(lp.FuncList.WallReplyRestore, f)
}

// WallReplyDelete handler
func (lp *Longpoll) WallReplyDelete(f object.WallReplyDeleteFunc) {
	lp.FuncList.WallReplyDelete = append(lp.FuncList.WallReplyDelete, f)
}

// BoardPostNew handler
func (lp *Longpoll) BoardPostNew(f object.BoardPostNewFunc) {
	lp.FuncList.BoardPostNew = append(lp.FuncList.BoardPostNew, f)
}

// BoardPostEdit handler
func (lp *Longpoll) BoardPostEdit(f object.BoardPostEditFunc) {
	lp.FuncList.BoardPostEdit = append(lp.FuncList.BoardPostEdit, f)
}

// BoardPostRestore handler
func (lp *Longpoll) BoardPostRestore(f object.BoardPostRestoreFunc) {
	lp.FuncList.BoardPostRestore = append(lp.FuncList.BoardPostRestore, f)
}

// BoardPostDelete handler
func (lp *Longpoll) BoardPostDelete(f object.BoardPostDeleteFunc) {
	lp.FuncList.BoardPostDelete = append(lp.FuncList.BoardPostDelete, f)
}

// MarketCommentNew handler
func (lp *Longpoll) MarketCommentNew(f object.MarketCommentNewFunc) {
	lp.FuncList.MarketCommentNew = append(lp.FuncList.MarketCommentNew, f)
}

// MarketCommentEdit handler
func (lp *Longpoll) MarketCommentEdit(f object.MarketCommentEditFunc) {
	lp.FuncList.MarketCommentEdit = append(lp.FuncList.MarketCommentEdit, f)
}

// MarketCommentRestore handler
func (lp *Longpoll) MarketCommentRestore(f object.MarketCommentRestoreFunc) {
	lp.FuncList.MarketCommentRestore = append(lp.FuncList.MarketCommentRestore, f)
}

// MarketCommentDelete handler
func (lp *Longpoll) MarketCommentDelete(f object.MarketCommentDeleteFunc) {
	lp.FuncList.MarketCommentDelete = append(lp.FuncList.MarketCommentDelete, f)
}

// GroupLeave handler
func (lp *Longpoll) GroupLeave(f object.GroupLeaveFunc) {
	lp.FuncList.GroupLeave = append(lp.FuncList.GroupLeave, f)
}

// GroupJoin handler
func (lp *Longpoll) GroupJoin(f object.GroupJoinFunc) {
	lp.FuncList.GroupJoin = append(lp.FuncList.GroupJoin, f)
}

// UserBlock handler
func (lp *Longpoll) UserBlock(f object.UserBlockFunc) {
	lp.FuncList.UserBlock = append(lp.FuncList.UserBlock, f)
}

// UserUnblock handler
func (lp *Longpoll) UserUnblock(f object.UserUnblockFunc) {
	lp.FuncList.UserUnblock = append(lp.FuncList.UserUnblock, f)
}

// PollVoteNew handler
func (lp *Longpoll) PollVoteNew(f object.PollVoteNewFunc) {
	lp.FuncList.PollVoteNew = append(lp.FuncList.PollVoteNew, f)
}

// GroupOfficersEdit handler
func (lp *Longpoll) GroupOfficersEdit(f object.GroupOfficersEditFunc) {
	lp.FuncList.GroupOfficersEdit = append(lp.FuncList.GroupOfficersEdit, f)
}

// GroupChangeSettings handler
func (lp *Longpoll) GroupChangeSettings(f object.GroupChangeSettingsFunc) {
	lp.FuncList.GroupChangeSettings = append(lp.FuncList.GroupChangeSettings, f)
}

// GroupChangePhoto handler
func (lp *Longpoll) GroupChangePhoto(f object.GroupChangePhotoFunc) {
	lp.FuncList.GroupChangePhoto = append(lp.FuncList.GroupChangePhoto, f)
}

// VkpayTransaction handler
func (lp *Longpoll) VkpayTransaction(f object.VkpayTransactionFunc) {
	lp.FuncList.VkpayTransaction = append(lp.FuncList.VkpayTransaction, f)
}

// TODO: next version lead_forms_new handler
// TODO: next version like_add handler
// TODO: next version like_remove handler
