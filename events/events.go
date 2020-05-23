/*
Package events for community events handling.

See more https://vk.com/dev/groups_events
*/
package events // import "github.com/SevereCloud/vksdk/events"

import (
	"encoding/json"

	"github.com/SevereCloud/vksdk/object"
)

// FuncList struct.
type FuncList struct {
	messageNew           []object.MessageNewFunc
	messageReply         []object.MessageReplyFunc
	messageEdit          []object.MessageEditFunc
	messageAllow         []object.MessageAllowFunc
	messageDeny          []object.MessageDenyFunc
	messageTypingState   []object.MessageTypingStateFunc
	photoNew             []object.PhotoNewFunc
	photoCommentNew      []object.PhotoCommentNewFunc
	photoCommentEdit     []object.PhotoCommentEditFunc
	photoCommentRestore  []object.PhotoCommentRestoreFunc
	photoCommentDelete   []object.PhotoCommentDeleteFunc
	audioNew             []object.AudioNewFunc
	videoNew             []object.VideoNewFunc
	videoCommentNew      []object.VideoCommentNewFunc
	videoCommentEdit     []object.VideoCommentEditFunc
	videoCommentRestore  []object.VideoCommentRestoreFunc
	videoCommentDelete   []object.VideoCommentDeleteFunc
	wallPostNew          []object.WallPostNewFunc
	wallRepost           []object.WallRepostFunc
	wallReplyNew         []object.WallReplyNewFunc
	wallReplyEdit        []object.WallReplyEditFunc
	wallReplyRestore     []object.WallReplyRestoreFunc
	wallReplyDelete      []object.WallReplyDeleteFunc
	boardPostNew         []object.BoardPostNewFunc
	boardPostEdit        []object.BoardPostEditFunc
	boardPostRestore     []object.BoardPostRestoreFunc
	boardPostDelete      []object.BoardPostDeleteFunc
	marketCommentNew     []object.MarketCommentNewFunc
	marketCommentEdit    []object.MarketCommentEditFunc
	marketCommentRestore []object.MarketCommentRestoreFunc
	marketCommentDelete  []object.MarketCommentDeleteFunc
	groupLeave           []object.GroupLeaveFunc
	groupJoin            []object.GroupJoinFunc
	userBlock            []object.UserBlockFunc
	userUnblock          []object.UserUnblockFunc
	pollVoteNew          []object.PollVoteNewFunc
	groupOfficersEdit    []object.GroupOfficersEditFunc
	groupChangeSettings  []object.GroupChangeSettingsFunc
	groupChangePhoto     []object.GroupChangePhotoFunc
	vkpayTransaction     []object.VkpayTransactionFunc
	leadFormsNew         []object.LeadFormsNewFunc
	appPayload           []object.AppPayloadFunc
	messageRead          []object.MessageReadFunc
	likeAdd              []object.LikeAddFunc
	special              map[string][]func(object.GroupEvent)
}

// NewFuncList returns a new FuncList.
func NewFuncList() *FuncList {
	return &FuncList{
		special: make(map[string][]func(object.GroupEvent)),
	}
}

// Handler group event handler.
func (fl FuncList) Handler(e object.GroupEvent) error { // nolint:gocyclo
	if sliceFunc, ok := fl.special[e.Type]; ok {
		for _, f := range sliceFunc {
			f(e)
		}
	}

	switch e.Type {
	case object.EventMessageNew:
		var obj object.MessageNewObject
		if err := json.Unmarshal(e.Object, &obj); err != nil {
			return err
		}

		for _, f := range fl.messageNew {
			f(obj, e.GroupID)
		}
	case object.EventMessageReply:
		var obj object.MessageReplyObject
		if err := json.Unmarshal(e.Object, &obj); err != nil {
			return err
		}

		for _, f := range fl.messageReply {
			f(obj, e.GroupID)
		}
	case object.EventMessageEdit:
		var obj object.MessageEditObject
		if err := json.Unmarshal(e.Object, &obj); err != nil {
			return err
		}

		for _, f := range fl.messageEdit {
			f(obj, e.GroupID)
		}
	case object.EventMessageAllow:
		var obj object.MessageAllowObject
		if err := json.Unmarshal(e.Object, &obj); err != nil {
			return err
		}

		for _, f := range fl.messageAllow {
			f(obj, e.GroupID)
		}
	case object.EventMessageDeny:
		var obj object.MessageDenyObject
		if err := json.Unmarshal(e.Object, &obj); err != nil {
			return err
		}

		for _, f := range fl.messageDeny {
			f(obj, e.GroupID)
		}
	case object.EventMessageTypingState: // На основе ответа
		var obj object.MessageTypingStateObject
		if err := json.Unmarshal(e.Object, &obj); err != nil {
			return err
		}

		for _, f := range fl.messageTypingState {
			f(obj, e.GroupID)
		}
	case object.EventPhotoNew:
		var obj object.PhotoNewObject
		if err := json.Unmarshal(e.Object, &obj); err != nil {
			return err
		}

		for _, f := range fl.photoNew {
			f(obj, e.GroupID)
		}
	case object.EventPhotoCommentNew:
		var obj object.PhotoCommentNewObject
		if err := json.Unmarshal(e.Object, &obj); err != nil {
			return err
		}

		for _, f := range fl.photoCommentNew {
			f(obj, e.GroupID)
		}
	case object.EventPhotoCommentEdit:
		var obj object.PhotoCommentEditObject
		if err := json.Unmarshal(e.Object, &obj); err != nil {
			return err
		}

		for _, f := range fl.photoCommentEdit {
			f(obj, e.GroupID)
		}
	case object.EventPhotoCommentRestore:
		var obj object.PhotoCommentRestoreObject
		if err := json.Unmarshal(e.Object, &obj); err != nil {
			return err
		}

		for _, f := range fl.photoCommentRestore {
			f(obj, e.GroupID)
		}
	case object.EventPhotoCommentDelete:
		var obj object.PhotoCommentDeleteObject
		if err := json.Unmarshal(e.Object, &obj); err != nil {
			return err
		}

		for _, f := range fl.photoCommentDelete {
			f(obj, e.GroupID)
		}
	case object.EventAudioNew:
		var obj object.AudioNewObject
		if err := json.Unmarshal(e.Object, &obj); err != nil {
			return err
		}

		for _, f := range fl.audioNew {
			f(obj, e.GroupID)
		}
	case object.EventVideoNew:
		var obj object.VideoNewObject
		if err := json.Unmarshal(e.Object, &obj); err != nil {
			return err
		}

		for _, f := range fl.videoNew {
			f(obj, e.GroupID)
		}
	case object.EventVideoCommentNew:
		var obj object.VideoCommentNewObject
		if err := json.Unmarshal(e.Object, &obj); err != nil {
			return err
		}

		for _, f := range fl.videoCommentNew {
			f(obj, e.GroupID)
		}
	case object.EventVideoCommentEdit:
		var obj object.VideoCommentEditObject
		if err := json.Unmarshal(e.Object, &obj); err != nil {
			return err
		}

		for _, f := range fl.videoCommentEdit {
			f(obj, e.GroupID)
		}
	case object.EventVideoCommentRestore:
		var obj object.VideoCommentRestoreObject
		if err := json.Unmarshal(e.Object, &obj); err != nil {
			return err
		}

		for _, f := range fl.videoCommentRestore {
			f(obj, e.GroupID)
		}
	case object.EventVideoCommentDelete:
		var obj object.VideoCommentDeleteObject
		if err := json.Unmarshal(e.Object, &obj); err != nil {
			return err
		}

		for _, f := range fl.videoCommentDelete {
			f(obj, e.GroupID)
		}
	case object.EventWallPostNew:
		var obj object.WallPostNewObject
		if err := json.Unmarshal(e.Object, &obj); err != nil {
			return err
		}

		for _, f := range fl.wallPostNew {
			f(obj, e.GroupID)
		}
	case object.EventWallRepost:
		var obj object.WallRepostObject
		if err := json.Unmarshal(e.Object, &obj); err != nil {
			return err
		}

		for _, f := range fl.wallRepost {
			f(obj, e.GroupID)
		}
	case object.EventWallReplyNew:
		var obj object.WallReplyNewObject
		if err := json.Unmarshal(e.Object, &obj); err != nil {
			return err
		}

		for _, f := range fl.wallReplyNew {
			f(obj, e.GroupID)
		}
	case object.EventWallReplyEdit:
		var obj object.WallReplyEditObject
		if err := json.Unmarshal(e.Object, &obj); err != nil {
			return err
		}

		for _, f := range fl.wallReplyEdit {
			f(obj, e.GroupID)
		}
	case object.EventWallReplyRestore:
		var obj object.WallReplyRestoreObject
		if err := json.Unmarshal(e.Object, &obj); err != nil {
			return err
		}

		for _, f := range fl.wallReplyRestore {
			f(obj, e.GroupID)
		}
	case object.EventWallReplyDelete:
		var obj object.WallReplyDeleteObject
		if err := json.Unmarshal(e.Object, &obj); err != nil {
			return err
		}

		for _, f := range fl.wallReplyDelete {
			f(obj, e.GroupID)
		}
	case object.EventBoardPostNew:
		var obj object.BoardPostNewObject
		if err := json.Unmarshal(e.Object, &obj); err != nil {
			return err
		}

		for _, f := range fl.boardPostNew {
			f(obj, e.GroupID)
		}
	case object.EventBoardPostEdit:
		var obj object.BoardPostEditObject
		if err := json.Unmarshal(e.Object, &obj); err != nil {
			return err
		}

		for _, f := range fl.boardPostEdit {
			f(obj, e.GroupID)
		}
	case object.EventBoardPostRestore:
		var obj object.BoardPostRestoreObject
		if err := json.Unmarshal(e.Object, &obj); err != nil {
			return err
		}

		for _, f := range fl.boardPostRestore {
			f(obj, e.GroupID)
		}
	case object.EventBoardPostDelete:
		var obj object.BoardPostDeleteObject
		if err := json.Unmarshal(e.Object, &obj); err != nil {
			return err
		}

		for _, f := range fl.boardPostDelete {
			f(obj, e.GroupID)
		}
	case object.EventMarketCommentNew:
		var obj object.MarketCommentNewObject
		if err := json.Unmarshal(e.Object, &obj); err != nil {
			return err
		}

		for _, f := range fl.marketCommentNew {
			f(obj, e.GroupID)
		}
	case object.EventMarketCommentEdit:
		var obj object.MarketCommentEditObject
		if err := json.Unmarshal(e.Object, &obj); err != nil {
			return err
		}

		for _, f := range fl.marketCommentEdit {
			f(obj, e.GroupID)
		}
	case object.EventMarketCommentRestore:
		var obj object.MarketCommentRestoreObject
		if err := json.Unmarshal(e.Object, &obj); err != nil {
			return err
		}

		for _, f := range fl.marketCommentRestore {
			f(obj, e.GroupID)
		}
	case object.EventMarketCommentDelete:
		var obj object.MarketCommentDeleteObject
		if err := json.Unmarshal(e.Object, &obj); err != nil {
			return err
		}

		for _, f := range fl.marketCommentDelete {
			f(obj, e.GroupID)
		}
	case object.EventGroupLeave:
		var obj object.GroupLeaveObject
		if err := json.Unmarshal(e.Object, &obj); err != nil {
			return err
		}

		for _, f := range fl.groupLeave {
			f(obj, e.GroupID)
		}
	case object.EventGroupJoin:
		var obj object.GroupJoinObject
		if err := json.Unmarshal(e.Object, &obj); err != nil {
			return err
		}

		for _, f := range fl.groupJoin {
			f(obj, e.GroupID)
		}
	case object.EventUserBlock:
		var obj object.UserBlockObject
		if err := json.Unmarshal(e.Object, &obj); err != nil {
			return err
		}

		for _, f := range fl.userBlock {
			f(obj, e.GroupID)
		}
	case object.EventUserUnblock:
		var obj object.UserUnblockObject
		if err := json.Unmarshal(e.Object, &obj); err != nil {
			return err
		}

		for _, f := range fl.userUnblock {
			f(obj, e.GroupID)
		}
	case object.EventPollVoteNew:
		var obj object.PollVoteNewObject
		if err := json.Unmarshal(e.Object, &obj); err != nil {
			return err
		}

		for _, f := range fl.pollVoteNew {
			f(obj, e.GroupID)
		}
	case object.EventGroupOfficersEdit:
		var obj object.GroupOfficersEditObject
		if err := json.Unmarshal(e.Object, &obj); err != nil {
			return err
		}

		for _, f := range fl.groupOfficersEdit {
			f(obj, e.GroupID)
		}
	case object.EventGroupChangeSettings:
		var obj object.GroupChangeSettingsObject
		if err := json.Unmarshal(e.Object, &obj); err != nil {
			return err
		}

		for _, f := range fl.groupChangeSettings {
			f(obj, e.GroupID)
		}
	case object.EventGroupChangePhoto:
		var obj object.GroupChangePhotoObject
		if err := json.Unmarshal(e.Object, &obj); err != nil {
			return err
		}

		for _, f := range fl.groupChangePhoto {
			f(obj, e.GroupID)
		}
	case object.EventVkpayTransaction:
		var obj object.VkpayTransactionObject
		if err := json.Unmarshal(e.Object, &obj); err != nil {
			return err
		}

		for _, f := range fl.vkpayTransaction {
			f(obj, e.GroupID)
		}
	case object.EventLeadFormsNew:
		var obj object.LeadFormsNewObject
		if err := json.Unmarshal(e.Object, &obj); err != nil {
			return err
		}

		for _, f := range fl.leadFormsNew {
			f(obj, e.GroupID)
		}
	case object.EventAppPayload:
		var obj object.AppPayloadObject
		if err := json.Unmarshal(e.Object, &obj); err != nil {
			return err
		}

		for _, f := range fl.appPayload {
			f(obj, e.GroupID)
		}
	case object.EventMessageRead:
		var obj object.MessageReadObject
		if err := json.Unmarshal(e.Object, &obj); err != nil {
			return err
		}

		for _, f := range fl.messageRead {
			f(obj, e.GroupID)
		}
	case object.EventLikeAdd:
		var obj object.LikeAddObject
		if err := json.Unmarshal(e.Object, &obj); err != nil {
			return err
		}

		for _, f := range fl.likeAdd {
			f(obj, e.GroupID)
		}
	}

	return nil
}

// OnEvent handler.
func (fl *FuncList) OnEvent(eventType string, f func(object.GroupEvent)) {
	if fl.special == nil {
		fl.special = make(map[string][]func(object.GroupEvent))
	}

	fl.special[eventType] = append(fl.special[eventType], f)
}

// MessageNew handler.
func (fl *FuncList) MessageNew(f object.MessageNewFunc) {
	fl.messageNew = append(fl.messageNew, f)
}

// MessageReply handler.
func (fl *FuncList) MessageReply(f object.MessageReplyFunc) {
	fl.messageReply = append(fl.messageReply, f)
}

// MessageEdit handler.
func (fl *FuncList) MessageEdit(f object.MessageEditFunc) {
	fl.messageEdit = append(fl.messageEdit, f)
}

// MessageAllow handler.
func (fl *FuncList) MessageAllow(f object.MessageAllowFunc) {
	fl.messageAllow = append(fl.messageAllow, f)
}

// MessageDeny handler.
func (fl *FuncList) MessageDeny(f object.MessageDenyFunc) {
	fl.messageDeny = append(fl.messageDeny, f)
}

// MessageTypingState handler.
func (fl *FuncList) MessageTypingState(f object.MessageTypingStateFunc) {
	fl.messageTypingState = append(fl.messageTypingState, f)
}

// PhotoNew handler.
func (fl *FuncList) PhotoNew(f object.PhotoNewFunc) {
	fl.photoNew = append(fl.photoNew, f)
}

// PhotoCommentNew handler.
func (fl *FuncList) PhotoCommentNew(f object.PhotoCommentNewFunc) {
	fl.photoCommentNew = append(fl.photoCommentNew, f)
}

// PhotoCommentEdit handler.
func (fl *FuncList) PhotoCommentEdit(f object.PhotoCommentEditFunc) {
	fl.photoCommentEdit = append(fl.photoCommentEdit, f)
}

// PhotoCommentRestore handler.
func (fl *FuncList) PhotoCommentRestore(f object.PhotoCommentRestoreFunc) {
	fl.photoCommentRestore = append(fl.photoCommentRestore, f)
}

// PhotoCommentDelete handler.
func (fl *FuncList) PhotoCommentDelete(f object.PhotoCommentDeleteFunc) {
	fl.photoCommentDelete = append(fl.photoCommentDelete, f)
}

// AudioNew handler.
func (fl *FuncList) AudioNew(f object.AudioNewFunc) {
	fl.audioNew = append(fl.audioNew, f)
}

// VideoNew handler.
func (fl *FuncList) VideoNew(f object.VideoNewFunc) {
	fl.videoNew = append(fl.videoNew, f)
}

// VideoCommentNew handler.
func (fl *FuncList) VideoCommentNew(f object.VideoCommentNewFunc) {
	fl.videoCommentNew = append(fl.videoCommentNew, f)
}

// VideoCommentEdit handler.
func (fl *FuncList) VideoCommentEdit(f object.VideoCommentEditFunc) {
	fl.videoCommentEdit = append(fl.videoCommentEdit, f)
}

// VideoCommentRestore handler.
func (fl *FuncList) VideoCommentRestore(f object.VideoCommentRestoreFunc) {
	fl.videoCommentRestore = append(fl.videoCommentRestore, f)
}

// VideoCommentDelete handler.
func (fl *FuncList) VideoCommentDelete(f object.VideoCommentDeleteFunc) {
	fl.videoCommentDelete = append(fl.videoCommentDelete, f)
}

// WallPostNew handler.
func (fl *FuncList) WallPostNew(f object.WallPostNewFunc) {
	fl.wallPostNew = append(fl.wallPostNew, f)
}

// WallRepost handler.
func (fl *FuncList) WallRepost(f object.WallRepostFunc) {
	fl.wallRepost = append(fl.wallRepost, f)
}

// WallReplyNew handler.
func (fl *FuncList) WallReplyNew(f object.WallReplyNewFunc) {
	fl.wallReplyNew = append(fl.wallReplyNew, f)
}

// WallReplyEdit handler.
func (fl *FuncList) WallReplyEdit(f object.WallReplyEditFunc) {
	fl.wallReplyEdit = append(fl.wallReplyEdit, f)
}

// WallReplyRestore handler.
func (fl *FuncList) WallReplyRestore(f object.WallReplyRestoreFunc) {
	fl.wallReplyRestore = append(fl.wallReplyRestore, f)
}

// WallReplyDelete handler.
func (fl *FuncList) WallReplyDelete(f object.WallReplyDeleteFunc) {
	fl.wallReplyDelete = append(fl.wallReplyDelete, f)
}

// BoardPostNew handler.
func (fl *FuncList) BoardPostNew(f object.BoardPostNewFunc) {
	fl.boardPostNew = append(fl.boardPostNew, f)
}

// BoardPostEdit handler.
func (fl *FuncList) BoardPostEdit(f object.BoardPostEditFunc) {
	fl.boardPostEdit = append(fl.boardPostEdit, f)
}

// BoardPostRestore handler.
func (fl *FuncList) BoardPostRestore(f object.BoardPostRestoreFunc) {
	fl.boardPostRestore = append(fl.boardPostRestore, f)
}

// BoardPostDelete handler.
func (fl *FuncList) BoardPostDelete(f object.BoardPostDeleteFunc) {
	fl.boardPostDelete = append(fl.boardPostDelete, f)
}

// MarketCommentNew handler.
func (fl *FuncList) MarketCommentNew(f object.MarketCommentNewFunc) {
	fl.marketCommentNew = append(fl.marketCommentNew, f)
}

// MarketCommentEdit handler.
func (fl *FuncList) MarketCommentEdit(f object.MarketCommentEditFunc) {
	fl.marketCommentEdit = append(fl.marketCommentEdit, f)
}

// MarketCommentRestore handler.
func (fl *FuncList) MarketCommentRestore(f object.MarketCommentRestoreFunc) {
	fl.marketCommentRestore = append(fl.marketCommentRestore, f)
}

// MarketCommentDelete handler.
func (fl *FuncList) MarketCommentDelete(f object.MarketCommentDeleteFunc) {
	fl.marketCommentDelete = append(fl.marketCommentDelete, f)
}

// GroupLeave handler.
func (fl *FuncList) GroupLeave(f object.GroupLeaveFunc) {
	fl.groupLeave = append(fl.groupLeave, f)
}

// GroupJoin handler.
func (fl *FuncList) GroupJoin(f object.GroupJoinFunc) {
	fl.groupJoin = append(fl.groupJoin, f)
}

// UserBlock handler.
func (fl *FuncList) UserBlock(f object.UserBlockFunc) {
	fl.userBlock = append(fl.userBlock, f)
}

// UserUnblock handler.
func (fl *FuncList) UserUnblock(f object.UserUnblockFunc) {
	fl.userUnblock = append(fl.userUnblock, f)
}

// PollVoteNew handler.
func (fl *FuncList) PollVoteNew(f object.PollVoteNewFunc) {
	fl.pollVoteNew = append(fl.pollVoteNew, f)
}

// GroupOfficersEdit handler.
func (fl *FuncList) GroupOfficersEdit(f object.GroupOfficersEditFunc) {
	fl.groupOfficersEdit = append(fl.groupOfficersEdit, f)
}

// GroupChangeSettings handler.
func (fl *FuncList) GroupChangeSettings(f object.GroupChangeSettingsFunc) {
	fl.groupChangeSettings = append(fl.groupChangeSettings, f)
}

// GroupChangePhoto handler.
func (fl *FuncList) GroupChangePhoto(f object.GroupChangePhotoFunc) {
	fl.groupChangePhoto = append(fl.groupChangePhoto, f)
}

// VkpayTransaction handler.
func (fl *FuncList) VkpayTransaction(f object.VkpayTransactionFunc) {
	fl.vkpayTransaction = append(fl.vkpayTransaction, f)
}

// LeadFormsNew handler.
func (fl *FuncList) LeadFormsNew(f object.LeadFormsNewFunc) {
	fl.leadFormsNew = append(fl.leadFormsNew, f)
}

// AppPayload handler.
func (fl *FuncList) AppPayload(f object.AppPayloadFunc) {
	fl.appPayload = append(fl.appPayload, f)
}

// MessageRead handler.
func (fl *FuncList) MessageRead(f object.MessageReadFunc) {
	fl.messageRead = append(fl.messageRead, f)
}

// LikeAdd handler.
func (fl *FuncList) LikeAdd(f object.LikeAddFunc) {
	fl.likeAdd = append(fl.likeAdd, f)
}
