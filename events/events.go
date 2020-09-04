/*
Package events for community events handling.

See more https://vk.com/dev/groups_events
*/
package events // import "github.com/SevereCloud/vksdk/events"

import (
	"context"
	"encoding/json"

	"github.com/SevereCloud/vksdk/internal"
)

// EventType type.
type EventType string

// EventType list.
const (
	EventConfirmation         = "confirmation"
	EventMessageNew           = "message_new"
	EventMessageReply         = "message_reply"
	EventMessageEdit          = "message_edit"
	EventMessageAllow         = "message_allow"
	EventMessageDeny          = "message_deny"
	EventMessageTypingState   = "message_typing_state"
	EventMessageEvent         = "message_event"
	EventPhotoNew             = "photo_new"
	EventPhotoCommentNew      = "photo_comment_new"
	EventPhotoCommentEdit     = "photo_comment_edit"
	EventPhotoCommentRestore  = "photo_comment_restore"
	EventPhotoCommentDelete   = "photo_comment_delete"
	EventAudioNew             = "audio_new"
	EventVideoNew             = "video_new"
	EventVideoCommentNew      = "video_comment_new"
	EventVideoCommentEdit     = "video_comment_edit"
	EventVideoCommentRestore  = "video_comment_restore"
	EventVideoCommentDelete   = "video_comment_delete"
	EventWallPostNew          = "wall_post_new"
	EventWallRepost           = "wall_repost"
	EventWallReplyNew         = "wall_reply_new"
	EventWallReplyEdit        = "wall_reply_edit"
	EventWallReplyRestore     = "wall_reply_restore"
	EventWallReplyDelete      = "wall_reply_delete"
	EventBoardPostNew         = "board_post_new"
	EventBoardPostEdit        = "board_post_edit"
	EventBoardPostRestore     = "board_post_restore"
	EventBoardPostDelete      = "board_post_delete"
	EventMarketCommentNew     = "market_comment_new"
	EventMarketCommentEdit    = "market_comment_edit"
	EventMarketCommentRestore = "market_comment_restore"
	EventMarketCommentDelete  = "market_comment_delete"
	EventMarketOrderNew       = "market_order_new"
	EventMarketOrderEdit      = "market_order_edit"
	EventGroupLeave           = "group_leave"
	EventGroupJoin            = "group_join"
	EventUserBlock            = "user_block"
	EventUserUnblock          = "user_unblock"
	EventPollVoteNew          = "poll_vote_new"
	EventGroupOfficersEdit    = "group_officers_edit"
	EventGroupChangeSettings  = "group_change_settings"
	EventGroupChangePhoto     = "group_change_photo"
	EventVkpayTransaction     = "vkpay_transaction"
	EventLeadFormsNew         = "lead_forms_new"
	EventAppPayload           = "app_payload"
	EventMessageRead          = "message_read"
	EventLikeAdd              = "like_add"
	EventLikeRemove           = "like_remove"
)

// GroupEvent struct.
type GroupEvent struct {
	Type    EventType       `json:"type"`
	Object  json.RawMessage `json:"object"`
	GroupID int             `json:"group_id"`
	EventID string          `json:"event_id"`
	Secret  string          `json:"secret"`
}

// FuncList struct.
type FuncList struct {
	messageNew           []MessageNewFunc
	messageReply         []MessageReplyFunc
	messageEdit          []MessageEditFunc
	messageAllow         []MessageAllowFunc
	messageDeny          []MessageDenyFunc
	messageTypingState   []MessageTypingStateFunc
	messageEvent         []MessageEventFunc
	photoNew             []PhotoNewFunc
	photoCommentNew      []PhotoCommentNewFunc
	photoCommentEdit     []PhotoCommentEditFunc
	photoCommentRestore  []PhotoCommentRestoreFunc
	photoCommentDelete   []PhotoCommentDeleteFunc
	audioNew             []AudioNewFunc
	videoNew             []VideoNewFunc
	videoCommentNew      []VideoCommentNewFunc
	videoCommentEdit     []VideoCommentEditFunc
	videoCommentRestore  []VideoCommentRestoreFunc
	videoCommentDelete   []VideoCommentDeleteFunc
	wallPostNew          []WallPostNewFunc
	wallRepost           []WallRepostFunc
	wallReplyNew         []WallReplyNewFunc
	wallReplyEdit        []WallReplyEditFunc
	wallReplyRestore     []WallReplyRestoreFunc
	wallReplyDelete      []WallReplyDeleteFunc
	boardPostNew         []BoardPostNewFunc
	boardPostEdit        []BoardPostEditFunc
	boardPostRestore     []BoardPostRestoreFunc
	boardPostDelete      []BoardPostDeleteFunc
	marketCommentNew     []MarketCommentNewFunc
	marketCommentEdit    []MarketCommentEditFunc
	marketCommentRestore []MarketCommentRestoreFunc
	marketCommentDelete  []MarketCommentDeleteFunc
	marketOrderNew       []MarketOrderNewFunc
	marketOrderEdit      []MarketOrderEditFunc
	groupLeave           []GroupLeaveFunc
	groupJoin            []GroupJoinFunc
	userBlock            []UserBlockFunc
	userUnblock          []UserUnblockFunc
	pollVoteNew          []PollVoteNewFunc
	groupOfficersEdit    []GroupOfficersEditFunc
	groupChangeSettings  []GroupChangeSettingsFunc
	groupChangePhoto     []GroupChangePhotoFunc
	vkpayTransaction     []VkpayTransactionFunc
	leadFormsNew         []LeadFormsNewFunc
	appPayload           []AppPayloadFunc
	messageRead          []MessageReadFunc
	likeAdd              []LikeAddFunc
	likeRemove           []LikeRemoveFunc
	special              map[EventType][]func(context.Context, GroupEvent)
	eventsList           []EventType
}

// NewFuncList returns a new FuncList.
func NewFuncList() *FuncList {
	return &FuncList{
		special: make(map[EventType][]func(context.Context, GroupEvent)),
	}
}

// Handler group event handler.
func (fl FuncList) Handler(ctx context.Context, e GroupEvent) error { // nolint:gocyclo
	ctx = context.WithValue(ctx, internal.GroupIDKey, e.GroupID)
	ctx = context.WithValue(ctx, internal.EventIDKey, e.EventID)

	if sliceFunc, ok := fl.special[e.Type]; ok {
		for _, f := range sliceFunc {
			f(ctx, e)
		}
	}

	switch e.Type {
	case EventMessageNew:
		var obj MessageNewObject
		if err := json.Unmarshal(e.Object, &obj); err != nil {
			return err
		}

		for _, f := range fl.messageNew {
			f(ctx, obj)
		}
	case EventMessageReply:
		var obj MessageReplyObject
		if err := json.Unmarshal(e.Object, &obj); err != nil {
			return err
		}

		for _, f := range fl.messageReply {
			f(ctx, obj)
		}
	case EventMessageEdit:
		var obj MessageEditObject
		if err := json.Unmarshal(e.Object, &obj); err != nil {
			return err
		}

		for _, f := range fl.messageEdit {
			f(ctx, obj)
		}
	case EventMessageAllow:
		var obj MessageAllowObject
		if err := json.Unmarshal(e.Object, &obj); err != nil {
			return err
		}

		for _, f := range fl.messageAllow {
			f(ctx, obj)
		}
	case EventMessageDeny:
		var obj MessageDenyObject
		if err := json.Unmarshal(e.Object, &obj); err != nil {
			return err
		}

		for _, f := range fl.messageDeny {
			f(ctx, obj)
		}
	case EventMessageTypingState: // На основе ответа
		var obj MessageTypingStateObject
		if err := json.Unmarshal(e.Object, &obj); err != nil {
			return err
		}

		for _, f := range fl.messageTypingState {
			f(ctx, obj)
		}
	case EventMessageEvent:
		var obj MessageEventObject
		if err := json.Unmarshal(e.Object, &obj); err != nil {
			return err
		}

		for _, f := range fl.messageEvent {
			f(ctx, obj)
		}
	case EventPhotoNew:
		var obj PhotoNewObject
		if err := json.Unmarshal(e.Object, &obj); err != nil {
			return err
		}

		for _, f := range fl.photoNew {
			f(ctx, obj)
		}
	case EventPhotoCommentNew:
		var obj PhotoCommentNewObject
		if err := json.Unmarshal(e.Object, &obj); err != nil {
			return err
		}

		for _, f := range fl.photoCommentNew {
			f(ctx, obj)
		}
	case EventPhotoCommentEdit:
		var obj PhotoCommentEditObject
		if err := json.Unmarshal(e.Object, &obj); err != nil {
			return err
		}

		for _, f := range fl.photoCommentEdit {
			f(ctx, obj)
		}
	case EventPhotoCommentRestore:
		var obj PhotoCommentRestoreObject
		if err := json.Unmarshal(e.Object, &obj); err != nil {
			return err
		}

		for _, f := range fl.photoCommentRestore {
			f(ctx, obj)
		}
	case EventPhotoCommentDelete:
		var obj PhotoCommentDeleteObject
		if err := json.Unmarshal(e.Object, &obj); err != nil {
			return err
		}

		for _, f := range fl.photoCommentDelete {
			f(ctx, obj)
		}
	case EventAudioNew:
		var obj AudioNewObject
		if err := json.Unmarshal(e.Object, &obj); err != nil {
			return err
		}

		for _, f := range fl.audioNew {
			f(ctx, obj)
		}
	case EventVideoNew:
		var obj VideoNewObject
		if err := json.Unmarshal(e.Object, &obj); err != nil {
			return err
		}

		for _, f := range fl.videoNew {
			f(ctx, obj)
		}
	case EventVideoCommentNew:
		var obj VideoCommentNewObject
		if err := json.Unmarshal(e.Object, &obj); err != nil {
			return err
		}

		for _, f := range fl.videoCommentNew {
			f(ctx, obj)
		}
	case EventVideoCommentEdit:
		var obj VideoCommentEditObject
		if err := json.Unmarshal(e.Object, &obj); err != nil {
			return err
		}

		for _, f := range fl.videoCommentEdit {
			f(ctx, obj)
		}
	case EventVideoCommentRestore:
		var obj VideoCommentRestoreObject
		if err := json.Unmarshal(e.Object, &obj); err != nil {
			return err
		}

		for _, f := range fl.videoCommentRestore {
			f(ctx, obj)
		}
	case EventVideoCommentDelete:
		var obj VideoCommentDeleteObject
		if err := json.Unmarshal(e.Object, &obj); err != nil {
			return err
		}

		for _, f := range fl.videoCommentDelete {
			f(ctx, obj)
		}
	case EventWallPostNew:
		var obj WallPostNewObject
		if err := json.Unmarshal(e.Object, &obj); err != nil {
			return err
		}

		for _, f := range fl.wallPostNew {
			f(ctx, obj)
		}
	case EventWallRepost:
		var obj WallRepostObject
		if err := json.Unmarshal(e.Object, &obj); err != nil {
			return err
		}

		for _, f := range fl.wallRepost {
			f(ctx, obj)
		}
	case EventWallReplyNew:
		var obj WallReplyNewObject
		if err := json.Unmarshal(e.Object, &obj); err != nil {
			return err
		}

		for _, f := range fl.wallReplyNew {
			f(ctx, obj)
		}
	case EventWallReplyEdit:
		var obj WallReplyEditObject
		if err := json.Unmarshal(e.Object, &obj); err != nil {
			return err
		}

		for _, f := range fl.wallReplyEdit {
			f(ctx, obj)
		}
	case EventWallReplyRestore:
		var obj WallReplyRestoreObject
		if err := json.Unmarshal(e.Object, &obj); err != nil {
			return err
		}

		for _, f := range fl.wallReplyRestore {
			f(ctx, obj)
		}
	case EventWallReplyDelete:
		var obj WallReplyDeleteObject
		if err := json.Unmarshal(e.Object, &obj); err != nil {
			return err
		}

		for _, f := range fl.wallReplyDelete {
			f(ctx, obj)
		}
	case EventBoardPostNew:
		var obj BoardPostNewObject
		if err := json.Unmarshal(e.Object, &obj); err != nil {
			return err
		}

		for _, f := range fl.boardPostNew {
			f(ctx, obj)
		}
	case EventBoardPostEdit:
		var obj BoardPostEditObject
		if err := json.Unmarshal(e.Object, &obj); err != nil {
			return err
		}

		for _, f := range fl.boardPostEdit {
			f(ctx, obj)
		}
	case EventBoardPostRestore:
		var obj BoardPostRestoreObject
		if err := json.Unmarshal(e.Object, &obj); err != nil {
			return err
		}

		for _, f := range fl.boardPostRestore {
			f(ctx, obj)
		}
	case EventBoardPostDelete:
		var obj BoardPostDeleteObject
		if err := json.Unmarshal(e.Object, &obj); err != nil {
			return err
		}

		for _, f := range fl.boardPostDelete {
			f(ctx, obj)
		}
	case EventMarketCommentNew:
		var obj MarketCommentNewObject
		if err := json.Unmarshal(e.Object, &obj); err != nil {
			return err
		}

		for _, f := range fl.marketCommentNew {
			f(ctx, obj)
		}
	case EventMarketCommentEdit:
		var obj MarketCommentEditObject
		if err := json.Unmarshal(e.Object, &obj); err != nil {
			return err
		}

		for _, f := range fl.marketCommentEdit {
			f(ctx, obj)
		}
	case EventMarketCommentRestore:
		var obj MarketCommentRestoreObject
		if err := json.Unmarshal(e.Object, &obj); err != nil {
			return err
		}

		for _, f := range fl.marketCommentRestore {
			f(ctx, obj)
		}
	case EventMarketCommentDelete:
		var obj MarketCommentDeleteObject
		if err := json.Unmarshal(e.Object, &obj); err != nil {
			return err
		}

		for _, f := range fl.marketCommentDelete {
			f(ctx, obj)
		}
	case EventMarketOrderNew:
		var obj MarketOrderNewObject
		if err := json.Unmarshal(e.Object, &obj); err != nil {
			return err
		}

		for _, f := range fl.marketOrderNew {
			f(ctx, obj)
		}
	case EventMarketOrderEdit:
		var obj MarketOrderEditObject
		if err := json.Unmarshal(e.Object, &obj); err != nil {
			return err
		}

		for _, f := range fl.marketOrderEdit {
			f(ctx, obj)
		}
	case EventGroupLeave:
		var obj GroupLeaveObject
		if err := json.Unmarshal(e.Object, &obj); err != nil {
			return err
		}

		for _, f := range fl.groupLeave {
			f(ctx, obj)
		}
	case EventGroupJoin:
		var obj GroupJoinObject
		if err := json.Unmarshal(e.Object, &obj); err != nil {
			return err
		}

		for _, f := range fl.groupJoin {
			f(ctx, obj)
		}
	case EventUserBlock:
		var obj UserBlockObject
		if err := json.Unmarshal(e.Object, &obj); err != nil {
			return err
		}

		for _, f := range fl.userBlock {
			f(ctx, obj)
		}
	case EventUserUnblock:
		var obj UserUnblockObject
		if err := json.Unmarshal(e.Object, &obj); err != nil {
			return err
		}

		for _, f := range fl.userUnblock {
			f(ctx, obj)
		}
	case EventPollVoteNew:
		var obj PollVoteNewObject
		if err := json.Unmarshal(e.Object, &obj); err != nil {
			return err
		}

		for _, f := range fl.pollVoteNew {
			f(ctx, obj)
		}
	case EventGroupOfficersEdit:
		var obj GroupOfficersEditObject
		if err := json.Unmarshal(e.Object, &obj); err != nil {
			return err
		}

		for _, f := range fl.groupOfficersEdit {
			f(ctx, obj)
		}
	case EventGroupChangeSettings:
		var obj GroupChangeSettingsObject
		if err := json.Unmarshal(e.Object, &obj); err != nil {
			return err
		}

		for _, f := range fl.groupChangeSettings {
			f(ctx, obj)
		}
	case EventGroupChangePhoto:
		var obj GroupChangePhotoObject
		if err := json.Unmarshal(e.Object, &obj); err != nil {
			return err
		}

		for _, f := range fl.groupChangePhoto {
			f(ctx, obj)
		}
	case EventVkpayTransaction:
		var obj VkpayTransactionObject
		if err := json.Unmarshal(e.Object, &obj); err != nil {
			return err
		}

		for _, f := range fl.vkpayTransaction {
			f(ctx, obj)
		}
	case EventLeadFormsNew:
		var obj LeadFormsNewObject
		if err := json.Unmarshal(e.Object, &obj); err != nil {
			return err
		}

		for _, f := range fl.leadFormsNew {
			f(ctx, obj)
		}
	case EventAppPayload:
		var obj AppPayloadObject
		if err := json.Unmarshal(e.Object, &obj); err != nil {
			return err
		}

		for _, f := range fl.appPayload {
			f(ctx, obj)
		}
	case EventMessageRead:
		var obj MessageReadObject
		if err := json.Unmarshal(e.Object, &obj); err != nil {
			return err
		}

		for _, f := range fl.messageRead {
			f(ctx, obj)
		}
	case EventLikeAdd:
		var obj LikeAddObject
		if err := json.Unmarshal(e.Object, &obj); err != nil {
			return err
		}

		for _, f := range fl.likeAdd {
			f(ctx, obj)
		}
	case EventLikeRemove:
		var obj LikeRemoveObject
		if err := json.Unmarshal(e.Object, &obj); err != nil {
			return err
		}

		for _, f := range fl.likeRemove {
			f(ctx, obj)
		}
	}

	return nil
}

// ListEvents return list of events.
func (fl FuncList) ListEvents() []EventType {
	return fl.eventsList
}

// OnEvent handler.
func (fl *FuncList) OnEvent(eventType EventType, f func(context.Context, GroupEvent)) {
	if fl.special == nil {
		fl.special = make(map[EventType][]func(context.Context, GroupEvent))
	}

	fl.special[eventType] = append(fl.special[eventType], f)
	fl.eventsList = append(fl.eventsList, eventType)
}

// MessageNew handler.
func (fl *FuncList) MessageNew(f MessageNewFunc) {
	fl.messageNew = append(fl.messageNew, f)
	fl.eventsList = append(fl.eventsList, EventMessageNew)
}

// MessageReply handler.
func (fl *FuncList) MessageReply(f MessageReplyFunc) {
	fl.messageReply = append(fl.messageReply, f)
	fl.eventsList = append(fl.eventsList, EventMessageReply)
}

// MessageEdit handler.
func (fl *FuncList) MessageEdit(f MessageEditFunc) {
	fl.messageEdit = append(fl.messageEdit, f)
	fl.eventsList = append(fl.eventsList, EventMessageEdit)
}

// MessageAllow handler.
func (fl *FuncList) MessageAllow(f MessageAllowFunc) {
	fl.messageAllow = append(fl.messageAllow, f)
	fl.eventsList = append(fl.eventsList, EventMessageAllow)
}

// MessageDeny handler.
func (fl *FuncList) MessageDeny(f MessageDenyFunc) {
	fl.messageDeny = append(fl.messageDeny, f)
	fl.eventsList = append(fl.eventsList, EventMessageDeny)
}

// MessageTypingState handler.
func (fl *FuncList) MessageTypingState(f MessageTypingStateFunc) {
	fl.messageTypingState = append(fl.messageTypingState, f)
	fl.eventsList = append(fl.eventsList, EventMessageTypingState)
}

// MessageEvent handler.
func (fl *FuncList) MessageEvent(f MessageEventFunc) {
	fl.messageEvent = append(fl.messageEvent, f)
	fl.eventsList = append(fl.eventsList, EventMessageEvent)
}

// PhotoNew handler.
func (fl *FuncList) PhotoNew(f PhotoNewFunc) {
	fl.photoNew = append(fl.photoNew, f)
	fl.eventsList = append(fl.eventsList, EventPhotoNew)
}

// PhotoCommentNew handler.
func (fl *FuncList) PhotoCommentNew(f PhotoCommentNewFunc) {
	fl.photoCommentNew = append(fl.photoCommentNew, f)
	fl.eventsList = append(fl.eventsList, EventPhotoCommentNew)
}

// PhotoCommentEdit handler.
func (fl *FuncList) PhotoCommentEdit(f PhotoCommentEditFunc) {
	fl.photoCommentEdit = append(fl.photoCommentEdit, f)
	fl.eventsList = append(fl.eventsList, EventPhotoCommentEdit)
}

// PhotoCommentRestore handler.
func (fl *FuncList) PhotoCommentRestore(f PhotoCommentRestoreFunc) {
	fl.photoCommentRestore = append(fl.photoCommentRestore, f)
	fl.eventsList = append(fl.eventsList, EventPhotoCommentRestore)
}

// PhotoCommentDelete handler.
func (fl *FuncList) PhotoCommentDelete(f PhotoCommentDeleteFunc) {
	fl.photoCommentDelete = append(fl.photoCommentDelete, f)
	fl.eventsList = append(fl.eventsList, EventPhotoCommentDelete)
}

// AudioNew handler.
func (fl *FuncList) AudioNew(f AudioNewFunc) {
	fl.audioNew = append(fl.audioNew, f)
	fl.eventsList = append(fl.eventsList, EventAudioNew)
}

// VideoNew handler.
func (fl *FuncList) VideoNew(f VideoNewFunc) {
	fl.videoNew = append(fl.videoNew, f)
	fl.eventsList = append(fl.eventsList, EventVideoNew)
}

// VideoCommentNew handler.
func (fl *FuncList) VideoCommentNew(f VideoCommentNewFunc) {
	fl.videoCommentNew = append(fl.videoCommentNew, f)
	fl.eventsList = append(fl.eventsList, EventVideoCommentNew)
}

// VideoCommentEdit handler.
func (fl *FuncList) VideoCommentEdit(f VideoCommentEditFunc) {
	fl.videoCommentEdit = append(fl.videoCommentEdit, f)
	fl.eventsList = append(fl.eventsList, EventVideoCommentEdit)
}

// VideoCommentRestore handler.
func (fl *FuncList) VideoCommentRestore(f VideoCommentRestoreFunc) {
	fl.videoCommentRestore = append(fl.videoCommentRestore, f)
	fl.eventsList = append(fl.eventsList, EventVideoCommentRestore)
}

// VideoCommentDelete handler.
func (fl *FuncList) VideoCommentDelete(f VideoCommentDeleteFunc) {
	fl.videoCommentDelete = append(fl.videoCommentDelete, f)
	fl.eventsList = append(fl.eventsList, EventVideoCommentDelete)
}

// WallPostNew handler.
func (fl *FuncList) WallPostNew(f WallPostNewFunc) {
	fl.wallPostNew = append(fl.wallPostNew, f)
	fl.eventsList = append(fl.eventsList, EventWallPostNew)
}

// WallRepost handler.
func (fl *FuncList) WallRepost(f WallRepostFunc) {
	fl.wallRepost = append(fl.wallRepost, f)
	fl.eventsList = append(fl.eventsList, EventWallRepost)
}

// WallReplyNew handler.
func (fl *FuncList) WallReplyNew(f WallReplyNewFunc) {
	fl.wallReplyNew = append(fl.wallReplyNew, f)
	fl.eventsList = append(fl.eventsList, EventWallReplyNew)
}

// WallReplyEdit handler.
func (fl *FuncList) WallReplyEdit(f WallReplyEditFunc) {
	fl.wallReplyEdit = append(fl.wallReplyEdit, f)
	fl.eventsList = append(fl.eventsList, EventWallReplyEdit)
}

// WallReplyRestore handler.
func (fl *FuncList) WallReplyRestore(f WallReplyRestoreFunc) {
	fl.wallReplyRestore = append(fl.wallReplyRestore, f)
	fl.eventsList = append(fl.eventsList, EventWallReplyRestore)
}

// WallReplyDelete handler.
func (fl *FuncList) WallReplyDelete(f WallReplyDeleteFunc) {
	fl.wallReplyDelete = append(fl.wallReplyDelete, f)
	fl.eventsList = append(fl.eventsList, EventWallReplyDelete)
}

// BoardPostNew handler.
func (fl *FuncList) BoardPostNew(f BoardPostNewFunc) {
	fl.boardPostNew = append(fl.boardPostNew, f)
	fl.eventsList = append(fl.eventsList, EventBoardPostNew)
}

// BoardPostEdit handler.
func (fl *FuncList) BoardPostEdit(f BoardPostEditFunc) {
	fl.boardPostEdit = append(fl.boardPostEdit, f)
	fl.eventsList = append(fl.eventsList, EventBoardPostEdit)
}

// BoardPostRestore handler.
func (fl *FuncList) BoardPostRestore(f BoardPostRestoreFunc) {
	fl.boardPostRestore = append(fl.boardPostRestore, f)
	fl.eventsList = append(fl.eventsList, EventBoardPostRestore)
}

// BoardPostDelete handler.
func (fl *FuncList) BoardPostDelete(f BoardPostDeleteFunc) {
	fl.boardPostDelete = append(fl.boardPostDelete, f)
	fl.eventsList = append(fl.eventsList, EventBoardPostDelete)
}

// MarketCommentNew handler.
func (fl *FuncList) MarketCommentNew(f MarketCommentNewFunc) {
	fl.marketCommentNew = append(fl.marketCommentNew, f)
	fl.eventsList = append(fl.eventsList, EventMarketCommentNew)
}

// MarketCommentEdit handler.
func (fl *FuncList) MarketCommentEdit(f MarketCommentEditFunc) {
	fl.marketCommentEdit = append(fl.marketCommentEdit, f)
	fl.eventsList = append(fl.eventsList, EventMarketCommentEdit)
}

// MarketCommentRestore handler.
func (fl *FuncList) MarketCommentRestore(f MarketCommentRestoreFunc) {
	fl.marketCommentRestore = append(fl.marketCommentRestore, f)
	fl.eventsList = append(fl.eventsList, EventMarketCommentRestore)
}

// MarketCommentDelete handler.
func (fl *FuncList) MarketCommentDelete(f MarketCommentDeleteFunc) {
	fl.marketCommentDelete = append(fl.marketCommentDelete, f)
	fl.eventsList = append(fl.eventsList, EventMarketCommentDelete)
}

// MarketOrderNew handler.
func (fl *FuncList) MarketOrderNew(f MarketOrderNewFunc) {
	fl.marketOrderNew = append(fl.marketOrderNew, f)
	fl.eventsList = append(fl.eventsList, EventMarketOrderNew)
}

// MarketOrderEdit handler.
func (fl *FuncList) MarketOrderEdit(f MarketOrderEditFunc) {
	fl.marketOrderEdit = append(fl.marketOrderEdit, f)
	fl.eventsList = append(fl.eventsList, EventMarketOrderEdit)
}

// GroupLeave handler.
func (fl *FuncList) GroupLeave(f GroupLeaveFunc) {
	fl.groupLeave = append(fl.groupLeave, f)
	fl.eventsList = append(fl.eventsList, EventGroupLeave)
}

// GroupJoin handler.
func (fl *FuncList) GroupJoin(f GroupJoinFunc) {
	fl.groupJoin = append(fl.groupJoin, f)
	fl.eventsList = append(fl.eventsList, EventGroupJoin)
}

// UserBlock handler.
func (fl *FuncList) UserBlock(f UserBlockFunc) {
	fl.userBlock = append(fl.userBlock, f)
	fl.eventsList = append(fl.eventsList, EventUserBlock)
}

// UserUnblock handler.
func (fl *FuncList) UserUnblock(f UserUnblockFunc) {
	fl.userUnblock = append(fl.userUnblock, f)
	fl.eventsList = append(fl.eventsList, EventUserUnblock)
}

// PollVoteNew handler.
func (fl *FuncList) PollVoteNew(f PollVoteNewFunc) {
	fl.pollVoteNew = append(fl.pollVoteNew, f)
	fl.eventsList = append(fl.eventsList, EventPollVoteNew)
}

// GroupOfficersEdit handler.
func (fl *FuncList) GroupOfficersEdit(f GroupOfficersEditFunc) {
	fl.groupOfficersEdit = append(fl.groupOfficersEdit, f)
	fl.eventsList = append(fl.eventsList, EventGroupOfficersEdit)
}

// GroupChangeSettings handler.
func (fl *FuncList) GroupChangeSettings(f GroupChangeSettingsFunc) {
	fl.groupChangeSettings = append(fl.groupChangeSettings, f)
	fl.eventsList = append(fl.eventsList, EventGroupChangeSettings)
}

// GroupChangePhoto handler.
func (fl *FuncList) GroupChangePhoto(f GroupChangePhotoFunc) {
	fl.groupChangePhoto = append(fl.groupChangePhoto, f)
	fl.eventsList = append(fl.eventsList, EventGroupChangePhoto)
}

// VkpayTransaction handler.
func (fl *FuncList) VkpayTransaction(f VkpayTransactionFunc) {
	fl.vkpayTransaction = append(fl.vkpayTransaction, f)
	fl.eventsList = append(fl.eventsList, EventVkpayTransaction)
}

// LeadFormsNew handler.
func (fl *FuncList) LeadFormsNew(f LeadFormsNewFunc) {
	fl.leadFormsNew = append(fl.leadFormsNew, f)
	fl.eventsList = append(fl.eventsList, EventLeadFormsNew)
}

// AppPayload handler.
func (fl *FuncList) AppPayload(f AppPayloadFunc) {
	fl.appPayload = append(fl.appPayload, f)
	fl.eventsList = append(fl.eventsList, EventAppPayload)
}

// MessageRead handler.
func (fl *FuncList) MessageRead(f MessageReadFunc) {
	fl.messageRead = append(fl.messageRead, f)
	fl.eventsList = append(fl.eventsList, EventMessageRead)
}

// LikeAdd handler.
func (fl *FuncList) LikeAdd(f LikeAddFunc) {
	fl.likeAdd = append(fl.likeAdd, f)
	fl.eventsList = append(fl.eventsList, EventLikeAdd)
}

// LikeRemove handler.
func (fl *FuncList) LikeRemove(f LikeRemoveFunc) {
	fl.likeRemove = append(fl.likeRemove, f)
	fl.eventsList = append(fl.eventsList, EventLikeRemove)
}
