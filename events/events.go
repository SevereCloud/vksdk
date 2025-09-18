/*
Package events for community events handling.

See more https://dev.vk.ru/ru/api/community-events/json-schema
*/
package events // import "github.com/SevereCloud/vksdk/v3/events"

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/SevereCloud/vksdk/v3/internal"
)

// EventType type.
type EventType string

// EventType list.
const (
	EventConfirmation                  = "confirmation"
	EventMessageNew                    = "message_new"
	EventMessageReply                  = "message_reply"
	EventMessageEdit                   = "message_edit"
	EventMessageAllow                  = "message_allow"
	EventMessageDeny                   = "message_deny"
	EventMessageTypingState            = "message_typing_state"
	EventMessageEvent                  = "message_event"
	EventPhotoNew                      = "photo_new"
	EventPhotoCommentNew               = "photo_comment_new"
	EventPhotoCommentEdit              = "photo_comment_edit"
	EventPhotoCommentRestore           = "photo_comment_restore"
	EventPhotoCommentDelete            = "photo_comment_delete"
	EventAudioNew                      = "audio_new"
	EventVideoNew                      = "video_new"
	EventVideoCommentNew               = "video_comment_new"
	EventVideoCommentEdit              = "video_comment_edit"
	EventVideoCommentRestore           = "video_comment_restore"
	EventVideoCommentDelete            = "video_comment_delete"
	EventWallPostNew                   = "wall_post_new"
	EventWallRepost                    = "wall_repost"
	EventWallReplyNew                  = "wall_reply_new"
	EventWallReplyEdit                 = "wall_reply_edit"
	EventWallReplyRestore              = "wall_reply_restore"
	EventWallReplyDelete               = "wall_reply_delete"
	EventBoardPostNew                  = "board_post_new"
	EventBoardPostEdit                 = "board_post_edit"
	EventBoardPostRestore              = "board_post_restore"
	EventBoardPostDelete               = "board_post_delete"
	EventMarketCommentNew              = "market_comment_new"
	EventMarketCommentEdit             = "market_comment_edit"
	EventMarketCommentRestore          = "market_comment_restore"
	EventMarketCommentDelete           = "market_comment_delete"
	EventMarketOrderNew                = "market_order_new"
	EventMarketOrderEdit               = "market_order_edit"
	EventGroupLeave                    = "group_leave"
	EventGroupJoin                     = "group_join"
	EventUserBlock                     = "user_block"
	EventUserUnblock                   = "user_unblock"
	EventPollVoteNew                   = "poll_vote_new"
	EventGroupOfficersEdit             = "group_officers_edit"
	EventGroupChangeSettings           = "group_change_settings"
	EventGroupChangePhoto              = "group_change_photo"
	EventVkpayTransaction              = "vkpay_transaction"
	EventLeadFormsNew                  = "lead_forms_new"
	EventAppPayload                    = "app_payload"
	EventMessageRead                   = "message_read"
	EventLikeAdd                       = "like_add"
	EventLikeRemove                    = "like_remove"
	EventDonutSubscriptionCreate       = "donut_subscription_create"
	EventDonutSubscriptionProlonged    = "donut_subscription_prolonged"
	EventDonutSubscriptionExpired      = "donut_subscription_expired"
	EventDonutSubscriptionCancelled    = "donut_subscription_cancelled"
	EventDonutSubscriptionPriceChanged = "donut_subscription_price_changed"
	EventDonutMoneyWithdraw            = "donut_money_withdraw"
	EventDonutMoneyWithdrawError       = "donut_money_withdraw_error"
)

// GroupEvent struct.
type GroupEvent struct {
	Type    EventType       `json:"type"`
	Object  json.RawMessage `json:"object"`
	GroupID int             `json:"group_id"`
	EventID string          `json:"event_id"`
	V       string          `json:"v"`
	Secret  string          `json:"secret"`
}

// FuncList struct.
type FuncList struct {
	messageNew                    []func(context.Context, MessageNewObject)
	messageReply                  []func(context.Context, MessageReplyObject)
	messageEdit                   []func(context.Context, MessageEditObject)
	messageAllow                  []func(context.Context, MessageAllowObject)
	messageDeny                   []func(context.Context, MessageDenyObject)
	messageTypingState            []func(context.Context, MessageTypingStateObject)
	messageEvent                  []func(context.Context, MessageEventObject)
	photoNew                      []func(context.Context, PhotoNewObject)
	photoCommentNew               []func(context.Context, PhotoCommentNewObject)
	photoCommentEdit              []func(context.Context, PhotoCommentEditObject)
	photoCommentRestore           []func(context.Context, PhotoCommentRestoreObject)
	photoCommentDelete            []func(context.Context, PhotoCommentDeleteObject)
	audioNew                      []func(context.Context, AudioNewObject)
	videoNew                      []func(context.Context, VideoNewObject)
	videoCommentNew               []func(context.Context, VideoCommentNewObject)
	videoCommentEdit              []func(context.Context, VideoCommentEditObject)
	videoCommentRestore           []func(context.Context, VideoCommentRestoreObject)
	videoCommentDelete            []func(context.Context, VideoCommentDeleteObject)
	wallPostNew                   []func(context.Context, WallPostNewObject)
	wallRepost                    []func(context.Context, WallRepostObject)
	wallReplyNew                  []func(context.Context, WallReplyNewObject)
	wallReplyEdit                 []func(context.Context, WallReplyEditObject)
	wallReplyRestore              []func(context.Context, WallReplyRestoreObject)
	wallReplyDelete               []func(context.Context, WallReplyDeleteObject)
	boardPostNew                  []func(context.Context, BoardPostNewObject)
	boardPostEdit                 []func(context.Context, BoardPostEditObject)
	boardPostRestore              []func(context.Context, BoardPostRestoreObject)
	boardPostDelete               []func(context.Context, BoardPostDeleteObject)
	marketCommentNew              []func(context.Context, MarketCommentNewObject)
	marketCommentEdit             []func(context.Context, MarketCommentEditObject)
	marketCommentRestore          []func(context.Context, MarketCommentRestoreObject)
	marketCommentDelete           []func(context.Context, MarketCommentDeleteObject)
	marketOrderNew                []func(context.Context, MarketOrderNewObject)
	marketOrderEdit               []func(context.Context, MarketOrderEditObject)
	groupLeave                    []func(context.Context, GroupLeaveObject)
	groupJoin                     []func(context.Context, GroupJoinObject)
	userBlock                     []func(context.Context, UserBlockObject)
	userUnblock                   []func(context.Context, UserUnblockObject)
	pollVoteNew                   []func(context.Context, PollVoteNewObject)
	groupOfficersEdit             []func(context.Context, GroupOfficersEditObject)
	groupChangeSettings           []func(context.Context, GroupChangeSettingsObject)
	groupChangePhoto              []func(context.Context, GroupChangePhotoObject)
	vkpayTransaction              []func(context.Context, VkpayTransactionObject)
	leadFormsNew                  []func(context.Context, LeadFormsNewObject)
	appPayload                    []func(context.Context, AppPayloadObject)
	messageRead                   []func(context.Context, MessageReadObject)
	likeAdd                       []func(context.Context, LikeAddObject)
	likeRemove                    []func(context.Context, LikeRemoveObject)
	donutSubscriptionCreate       []func(context.Context, DonutSubscriptionCreateObject)
	donutSubscriptionProlonged    []func(context.Context, DonutSubscriptionProlongedObject)
	donutSubscriptionExpired      []func(context.Context, DonutSubscriptionExpiredObject)
	donutSubscriptionCancelled    []func(context.Context, DonutSubscriptionCancelledObject)
	donutSubscriptionPriceChanged []func(context.Context, DonutSubscriptionPriceChangedObject)
	donutMoneyWithdraw            []func(context.Context, DonutMoneyWithdrawObject)
	donutMoneyWithdrawError       []func(context.Context, DonutMoneyWithdrawErrorObject)
	special                       map[EventType][]func(context.Context, GroupEvent)
	eventsList                    []EventType

	goroutine bool
}

// NewFuncList returns a new FuncList.
func NewFuncList() *FuncList {
	return &FuncList{
		special: make(map[EventType][]func(context.Context, GroupEvent)),
	}
}

// Handler group event handler.
func (fl *FuncList) Handler(ctx context.Context, e GroupEvent) error { //nolint:gocyclo
	ctx = context.WithValue(ctx, internal.GroupIDKey, e.GroupID)
	ctx = context.WithValue(ctx, internal.EventIDKey, e.EventID)
	ctx = context.WithValue(ctx, internal.EventVersionKey, e.V)

	if sliceFunc, ok := fl.special[e.Type]; ok {
		for _, f := range sliceFunc {
			if fl.goroutine {
				go func() { f(ctx, e) }()
			} else {
				f(ctx, e)
			}
		}
	}

	switch e.Type {
	case EventMessageNew:
		if len(fl.messageNew) == 0 {
			break
		}

		var obj MessageNewObject

		err := json.Unmarshal(e.Object, &obj)
		if err != nil {
			return fmt.Errorf("events: %w", err)
		}

		for _, f := range fl.messageNew {
			if fl.goroutine {
				go func() { f(context.WithoutCancel(ctx), obj) }()
			} else {
				f(ctx, obj)
			}
		}
	case EventMessageReply:
		if len(fl.messageReply) == 0 {
			break
		}

		var obj MessageReplyObject

		err := json.Unmarshal(e.Object, &obj)
		if err != nil {
			return fmt.Errorf("events: %w", err)
		}

		for _, f := range fl.messageReply {
			if fl.goroutine {
				go func() { f(context.WithoutCancel(ctx), obj) }()
			} else {
				f(ctx, obj)
			}
		}
	case EventMessageEdit:
		if len(fl.messageEdit) == 0 {
			break
		}

		var obj MessageEditObject

		err := json.Unmarshal(e.Object, &obj)
		if err != nil {
			return fmt.Errorf("events: %w", err)
		}

		for _, f := range fl.messageEdit {
			if fl.goroutine {
				go func() { f(context.WithoutCancel(ctx), obj) }()
			} else {
				f(ctx, obj)
			}
		}
	case EventMessageAllow:
		if len(fl.messageAllow) == 0 {
			break
		}

		var obj MessageAllowObject

		err := json.Unmarshal(e.Object, &obj)
		if err != nil {
			return fmt.Errorf("events: %w", err)
		}

		for _, f := range fl.messageAllow {
			if fl.goroutine {
				go func() { f(context.WithoutCancel(ctx), obj) }()
			} else {
				f(ctx, obj)
			}
		}
	case EventMessageDeny:
		if len(fl.messageDeny) == 0 {
			break
		}

		var obj MessageDenyObject

		err := json.Unmarshal(e.Object, &obj)
		if err != nil {
			return fmt.Errorf("events: %w", err)
		}

		for _, f := range fl.messageDeny {
			if fl.goroutine {
				go func() { f(context.WithoutCancel(ctx), obj) }()
			} else {
				f(ctx, obj)
			}
		}
	case EventMessageTypingState: // На основе ответа
		if len(fl.messageTypingState) == 0 {
			break
		}

		var obj MessageTypingStateObject

		err := json.Unmarshal(e.Object, &obj)
		if err != nil {
			return fmt.Errorf("events: %w", err)
		}

		for _, f := range fl.messageTypingState {
			if fl.goroutine {
				go func() { f(context.WithoutCancel(ctx), obj) }()
			} else {
				f(ctx, obj)
			}
		}
	case EventMessageEvent:
		if len(fl.messageEvent) == 0 {
			break
		}

		var obj MessageEventObject

		err := json.Unmarshal(e.Object, &obj)
		if err != nil {
			return fmt.Errorf("events: %w", err)
		}

		for _, f := range fl.messageEvent {
			if fl.goroutine {
				go func() { f(context.WithoutCancel(ctx), obj) }()
			} else {
				f(ctx, obj)
			}
		}
	case EventPhotoNew:
		if len(fl.photoNew) == 0 {
			break
		}

		var obj PhotoNewObject

		err := json.Unmarshal(e.Object, &obj)
		if err != nil {
			return fmt.Errorf("events: %w", err)
		}

		for _, f := range fl.photoNew {
			if fl.goroutine {
				go func() { f(context.WithoutCancel(ctx), obj) }()
			} else {
				f(ctx, obj)
			}
		}
	case EventPhotoCommentNew:
		if len(fl.photoCommentNew) == 0 {
			break
		}

		var obj PhotoCommentNewObject

		err := json.Unmarshal(e.Object, &obj)
		if err != nil {
			return fmt.Errorf("events: %w", err)
		}

		for _, f := range fl.photoCommentNew {
			if fl.goroutine {
				go func() { f(context.WithoutCancel(ctx), obj) }()
			} else {
				f(ctx, obj)
			}
		}
	case EventPhotoCommentEdit:
		if len(fl.photoCommentEdit) == 0 {
			break
		}

		var obj PhotoCommentEditObject

		err := json.Unmarshal(e.Object, &obj)
		if err != nil {
			return fmt.Errorf("events: %w", err)
		}

		for _, f := range fl.photoCommentEdit {
			if fl.goroutine {
				go func() { f(context.WithoutCancel(ctx), obj) }()
			} else {
				f(ctx, obj)
			}
		}
	case EventPhotoCommentRestore:
		if len(fl.photoCommentRestore) == 0 {
			break
		}

		var obj PhotoCommentRestoreObject

		err := json.Unmarshal(e.Object, &obj)
		if err != nil {
			return fmt.Errorf("events: %w", err)
		}

		for _, f := range fl.photoCommentRestore {
			if fl.goroutine {
				go func() { f(context.WithoutCancel(ctx), obj) }()
			} else {
				f(ctx, obj)
			}
		}
	case EventPhotoCommentDelete:
		if len(fl.photoCommentDelete) == 0 {
			break
		}

		var obj PhotoCommentDeleteObject

		err := json.Unmarshal(e.Object, &obj)
		if err != nil {
			return fmt.Errorf("events: %w", err)
		}

		for _, f := range fl.photoCommentDelete {
			if fl.goroutine {
				go func() { f(context.WithoutCancel(ctx), obj) }()
			} else {
				f(ctx, obj)
			}
		}
	case EventAudioNew:
		if len(fl.audioNew) == 0 {
			break
		}

		var obj AudioNewObject

		err := json.Unmarshal(e.Object, &obj)
		if err != nil {
			return fmt.Errorf("events: %w", err)
		}

		for _, f := range fl.audioNew {
			if fl.goroutine {
				go func() { f(context.WithoutCancel(ctx), obj) }()
			} else {
				f(ctx, obj)
			}
		}
	case EventVideoNew:
		if len(fl.videoNew) == 0 {
			break
		}

		var obj VideoNewObject

		err := json.Unmarshal(e.Object, &obj)
		if err != nil {
			return fmt.Errorf("events: %w", err)
		}

		for _, f := range fl.videoNew {
			if fl.goroutine {
				go func() { f(context.WithoutCancel(ctx), obj) }()
			} else {
				f(ctx, obj)
			}
		}
	case EventVideoCommentNew:
		if len(fl.videoCommentNew) == 0 {
			break
		}

		var obj VideoCommentNewObject

		err := json.Unmarshal(e.Object, &obj)
		if err != nil {
			return fmt.Errorf("events: %w", err)
		}

		for _, f := range fl.videoCommentNew {
			if fl.goroutine {
				go func() { f(context.WithoutCancel(ctx), obj) }()
			} else {
				f(ctx, obj)
			}
		}
	case EventVideoCommentEdit:
		if len(fl.videoCommentEdit) == 0 {
			break
		}

		var obj VideoCommentEditObject

		err := json.Unmarshal(e.Object, &obj)
		if err != nil {
			return fmt.Errorf("events: %w", err)
		}

		for _, f := range fl.videoCommentEdit {
			if fl.goroutine {
				go func() { f(context.WithoutCancel(ctx), obj) }()
			} else {
				f(ctx, obj)
			}
		}
	case EventVideoCommentRestore:
		if len(fl.videoCommentRestore) == 0 {
			break
		}

		var obj VideoCommentRestoreObject

		err := json.Unmarshal(e.Object, &obj)
		if err != nil {
			return fmt.Errorf("events: %w", err)
		}

		for _, f := range fl.videoCommentRestore {
			if fl.goroutine {
				go func() { f(context.WithoutCancel(ctx), obj) }()
			} else {
				f(ctx, obj)
			}
		}
	case EventVideoCommentDelete:
		if len(fl.videoCommentDelete) == 0 {
			break
		}

		var obj VideoCommentDeleteObject

		err := json.Unmarshal(e.Object, &obj)
		if err != nil {
			return fmt.Errorf("events: %w", err)
		}

		for _, f := range fl.videoCommentDelete {
			if fl.goroutine {
				go func() { f(context.WithoutCancel(ctx), obj) }()
			} else {
				f(ctx, obj)
			}
		}
	case EventWallPostNew:
		if len(fl.wallPostNew) == 0 {
			break
		}

		var obj WallPostNewObject

		err := json.Unmarshal(e.Object, &obj)
		if err != nil {
			return fmt.Errorf("events: %w", err)
		}

		for _, f := range fl.wallPostNew {
			if fl.goroutine {
				go func() { f(context.WithoutCancel(ctx), obj) }()
			} else {
				f(ctx, obj)
			}
		}
	case EventWallRepost:
		if len(fl.wallRepost) == 0 {
			break
		}

		var obj WallRepostObject

		err := json.Unmarshal(e.Object, &obj)
		if err != nil {
			return fmt.Errorf("events: %w", err)
		}

		for _, f := range fl.wallRepost {
			if fl.goroutine {
				go func() { f(context.WithoutCancel(ctx), obj) }()
			} else {
				f(ctx, obj)
			}
		}
	case EventWallReplyNew:
		if len(fl.wallReplyNew) == 0 {
			break
		}

		var obj WallReplyNewObject

		err := json.Unmarshal(e.Object, &obj)
		if err != nil {
			return fmt.Errorf("events: %w", err)
		}

		for _, f := range fl.wallReplyNew {
			if fl.goroutine {
				go func() { f(context.WithoutCancel(ctx), obj) }()
			} else {
				f(ctx, obj)
			}
		}
	case EventWallReplyEdit:
		if len(fl.wallReplyEdit) == 0 {
			break
		}

		var obj WallReplyEditObject

		err := json.Unmarshal(e.Object, &obj)
		if err != nil {
			return fmt.Errorf("events: %w", err)
		}

		for _, f := range fl.wallReplyEdit {
			if fl.goroutine {
				go func() { f(context.WithoutCancel(ctx), obj) }()
			} else {
				f(ctx, obj)
			}
		}
	case EventWallReplyRestore:
		if len(fl.wallReplyRestore) == 0 {
			break
		}

		var obj WallReplyRestoreObject

		err := json.Unmarshal(e.Object, &obj)
		if err != nil {
			return fmt.Errorf("events: %w", err)
		}

		for _, f := range fl.wallReplyRestore {
			if fl.goroutine {
				go func() { f(context.WithoutCancel(ctx), obj) }()
			} else {
				f(ctx, obj)
			}
		}
	case EventWallReplyDelete:
		if len(fl.wallReplyDelete) == 0 {
			break
		}

		var obj WallReplyDeleteObject

		err := json.Unmarshal(e.Object, &obj)
		if err != nil {
			return fmt.Errorf("events: %w", err)
		}

		for _, f := range fl.wallReplyDelete {
			if fl.goroutine {
				go func() { f(context.WithoutCancel(ctx), obj) }()
			} else {
				f(ctx, obj)
			}
		}
	case EventBoardPostNew:
		if len(fl.boardPostNew) == 0 {
			break
		}

		var obj BoardPostNewObject

		err := json.Unmarshal(e.Object, &obj)
		if err != nil {
			return fmt.Errorf("events: %w", err)
		}

		for _, f := range fl.boardPostNew {
			if fl.goroutine {
				go func() { f(context.WithoutCancel(ctx), obj) }()
			} else {
				f(ctx, obj)
			}
		}
	case EventBoardPostEdit:
		if len(fl.boardPostEdit) == 0 {
			break
		}

		var obj BoardPostEditObject

		err := json.Unmarshal(e.Object, &obj)
		if err != nil {
			return fmt.Errorf("events: %w", err)
		}

		for _, f := range fl.boardPostEdit {
			if fl.goroutine {
				go func() { f(context.WithoutCancel(ctx), obj) }()
			} else {
				f(ctx, obj)
			}
		}
	case EventBoardPostRestore:
		if len(fl.boardPostRestore) == 0 {
			break
		}

		var obj BoardPostRestoreObject

		err := json.Unmarshal(e.Object, &obj)
		if err != nil {
			return fmt.Errorf("events: %w", err)
		}

		for _, f := range fl.boardPostRestore {
			if fl.goroutine {
				go func() { f(context.WithoutCancel(ctx), obj) }()
			} else {
				f(ctx, obj)
			}
		}
	case EventBoardPostDelete:
		if len(fl.boardPostDelete) == 0 {
			break
		}

		var obj BoardPostDeleteObject

		err := json.Unmarshal(e.Object, &obj)
		if err != nil {
			return fmt.Errorf("events: %w", err)
		}

		for _, f := range fl.boardPostDelete {
			if fl.goroutine {
				go func() { f(context.WithoutCancel(ctx), obj) }()
			} else {
				f(ctx, obj)
			}
		}
	case EventMarketCommentNew:
		if len(fl.marketCommentNew) == 0 {
			break
		}

		var obj MarketCommentNewObject

		err := json.Unmarshal(e.Object, &obj)
		if err != nil {
			return fmt.Errorf("events: %w", err)
		}

		for _, f := range fl.marketCommentNew {
			if fl.goroutine {
				go func() { f(context.WithoutCancel(ctx), obj) }()
			} else {
				f(ctx, obj)
			}
		}
	case EventMarketCommentEdit:
		if len(fl.marketCommentEdit) == 0 {
			break
		}

		var obj MarketCommentEditObject

		err := json.Unmarshal(e.Object, &obj)
		if err != nil {
			return fmt.Errorf("events: %w", err)
		}

		for _, f := range fl.marketCommentEdit {
			if fl.goroutine {
				go func() { f(context.WithoutCancel(ctx), obj) }()
			} else {
				f(ctx, obj)
			}
		}
	case EventMarketCommentRestore:
		if len(fl.marketCommentRestore) == 0 {
			break
		}

		var obj MarketCommentRestoreObject

		err := json.Unmarshal(e.Object, &obj)
		if err != nil {
			return fmt.Errorf("events: %w", err)
		}

		for _, f := range fl.marketCommentRestore {
			if fl.goroutine {
				go func() { f(context.WithoutCancel(ctx), obj) }()
			} else {
				f(ctx, obj)
			}
		}
	case EventMarketCommentDelete:
		if len(fl.marketCommentDelete) == 0 {
			break
		}

		var obj MarketCommentDeleteObject

		err := json.Unmarshal(e.Object, &obj)
		if err != nil {
			return fmt.Errorf("events: %w", err)
		}

		for _, f := range fl.marketCommentDelete {
			if fl.goroutine {
				go func() { f(context.WithoutCancel(ctx), obj) }()
			} else {
				f(ctx, obj)
			}
		}
	case EventMarketOrderNew:
		if len(fl.marketOrderNew) == 0 {
			break
		}

		var obj MarketOrderNewObject

		err := json.Unmarshal(e.Object, &obj)
		if err != nil {
			return fmt.Errorf("events: %w", err)
		}

		for _, f := range fl.marketOrderNew {
			if fl.goroutine {
				go func() { f(context.WithoutCancel(ctx), obj) }()
			} else {
				f(ctx, obj)
			}
		}
	case EventMarketOrderEdit:
		if len(fl.marketOrderEdit) == 0 {
			break
		}

		var obj MarketOrderEditObject

		err := json.Unmarshal(e.Object, &obj)
		if err != nil {
			return fmt.Errorf("events: %w", err)
		}

		for _, f := range fl.marketOrderEdit {
			if fl.goroutine {
				go func() { f(context.WithoutCancel(ctx), obj) }()
			} else {
				f(ctx, obj)
			}
		}
	case EventGroupLeave:
		if len(fl.groupLeave) == 0 {
			break
		}

		var obj GroupLeaveObject

		err := json.Unmarshal(e.Object, &obj)
		if err != nil {
			return fmt.Errorf("events: %w", err)
		}

		for _, f := range fl.groupLeave {
			if fl.goroutine {
				go func() { f(context.WithoutCancel(ctx), obj) }()
			} else {
				f(ctx, obj)
			}
		}
	case EventGroupJoin:
		if len(fl.groupJoin) == 0 {
			break
		}

		var obj GroupJoinObject

		err := json.Unmarshal(e.Object, &obj)
		if err != nil {
			return fmt.Errorf("events: %w", err)
		}

		for _, f := range fl.groupJoin {
			if fl.goroutine {
				go func() { f(context.WithoutCancel(ctx), obj) }()
			} else {
				f(ctx, obj)
			}
		}
	case EventUserBlock:
		if len(fl.userBlock) == 0 {
			break
		}

		var obj UserBlockObject

		err := json.Unmarshal(e.Object, &obj)
		if err != nil {
			return fmt.Errorf("events: %w", err)
		}

		for _, f := range fl.userBlock {
			if fl.goroutine {
				go func() { f(context.WithoutCancel(ctx), obj) }()
			} else {
				f(ctx, obj)
			}
		}
	case EventUserUnblock:
		if len(fl.userUnblock) == 0 {
			break
		}

		var obj UserUnblockObject

		err := json.Unmarshal(e.Object, &obj)
		if err != nil {
			return fmt.Errorf("events: %w", err)
		}

		for _, f := range fl.userUnblock {
			if fl.goroutine {
				go func() { f(context.WithoutCancel(ctx), obj) }()
			} else {
				f(ctx, obj)
			}
		}
	case EventPollVoteNew:
		if len(fl.pollVoteNew) == 0 {
			break
		}

		var obj PollVoteNewObject

		err := json.Unmarshal(e.Object, &obj)
		if err != nil {
			return fmt.Errorf("events: %w", err)
		}

		for _, f := range fl.pollVoteNew {
			if fl.goroutine {
				go func() { f(context.WithoutCancel(ctx), obj) }()
			} else {
				f(ctx, obj)
			}
		}
	case EventGroupOfficersEdit:
		if len(fl.groupOfficersEdit) == 0 {
			break
		}

		var obj GroupOfficersEditObject

		err := json.Unmarshal(e.Object, &obj)
		if err != nil {
			return fmt.Errorf("events: %w", err)
		}

		for _, f := range fl.groupOfficersEdit {
			if fl.goroutine {
				go func() { f(context.WithoutCancel(ctx), obj) }()
			} else {
				f(ctx, obj)
			}
		}
	case EventGroupChangeSettings:
		if len(fl.groupChangeSettings) == 0 {
			break
		}

		var obj GroupChangeSettingsObject

		err := json.Unmarshal(e.Object, &obj)
		if err != nil {
			return fmt.Errorf("events: %w", err)
		}

		for _, f := range fl.groupChangeSettings {
			if fl.goroutine {
				go func() { f(context.WithoutCancel(ctx), obj) }()
			} else {
				f(ctx, obj)
			}
		}
	case EventGroupChangePhoto:
		if len(fl.groupChangePhoto) == 0 {
			break
		}

		var obj GroupChangePhotoObject

		err := json.Unmarshal(e.Object, &obj)
		if err != nil {
			return fmt.Errorf("events: %w", err)
		}

		for _, f := range fl.groupChangePhoto {
			if fl.goroutine {
				go func() { f(context.WithoutCancel(ctx), obj) }()
			} else {
				f(ctx, obj)
			}
		}
	case EventVkpayTransaction:
		if len(fl.vkpayTransaction) == 0 {
			break
		}

		var obj VkpayTransactionObject

		err := json.Unmarshal(e.Object, &obj)
		if err != nil {
			return fmt.Errorf("events: %w", err)
		}

		for _, f := range fl.vkpayTransaction {
			if fl.goroutine {
				go func() { f(context.WithoutCancel(ctx), obj) }()
			} else {
				f(ctx, obj)
			}
		}
	case EventLeadFormsNew:
		if len(fl.leadFormsNew) == 0 {
			break
		}

		var obj LeadFormsNewObject

		err := json.Unmarshal(e.Object, &obj)
		if err != nil {
			return fmt.Errorf("events: %w", err)
		}

		for _, f := range fl.leadFormsNew {
			if fl.goroutine {
				go func() { f(context.WithoutCancel(ctx), obj) }()
			} else {
				f(ctx, obj)
			}
		}
	case EventAppPayload:
		if len(fl.appPayload) == 0 {
			break
		}

		var obj AppPayloadObject

		err := json.Unmarshal(e.Object, &obj)
		if err != nil {
			return fmt.Errorf("events: %w", err)
		}

		for _, f := range fl.appPayload {
			if fl.goroutine {
				go func() { f(context.WithoutCancel(ctx), obj) }()
			} else {
				f(ctx, obj)
			}
		}
	case EventMessageRead:
		if len(fl.messageRead) == 0 {
			break
		}

		var obj MessageReadObject

		err := json.Unmarshal(e.Object, &obj)
		if err != nil {
			return fmt.Errorf("events: %w", err)
		}

		for _, f := range fl.messageRead {
			if fl.goroutine {
				go func() { f(context.WithoutCancel(ctx), obj) }()
			} else {
				f(ctx, obj)
			}
		}
	case EventLikeAdd:
		if len(fl.likeAdd) == 0 {
			break
		}

		var obj LikeAddObject

		err := json.Unmarshal(e.Object, &obj)
		if err != nil {
			return fmt.Errorf("events: %w", err)
		}

		for _, f := range fl.likeAdd {
			if fl.goroutine {
				go func() { f(context.WithoutCancel(ctx), obj) }()
			} else {
				f(ctx, obj)
			}
		}
	case EventLikeRemove:
		if len(fl.likeRemove) == 0 {
			break
		}

		var obj LikeRemoveObject

		err := json.Unmarshal(e.Object, &obj)
		if err != nil {
			return fmt.Errorf("events: %w", err)
		}

		for _, f := range fl.likeRemove {
			if fl.goroutine {
				go func() { f(context.WithoutCancel(ctx), obj) }()
			} else {
				f(ctx, obj)
			}
		}
	case EventDonutSubscriptionCreate:
		if len(fl.donutSubscriptionCreate) == 0 {
			break
		}

		var obj DonutSubscriptionCreateObject

		err := json.Unmarshal(e.Object, &obj)
		if err != nil {
			return fmt.Errorf("events: %w", err)
		}

		for _, f := range fl.donutSubscriptionCreate {
			if fl.goroutine {
				go func() { f(context.WithoutCancel(ctx), obj) }()
			} else {
				f(ctx, obj)
			}
		}
	case EventDonutSubscriptionProlonged:
		if len(fl.donutSubscriptionProlonged) == 0 {
			break
		}

		var obj DonutSubscriptionProlongedObject

		err := json.Unmarshal(e.Object, &obj)
		if err != nil {
			return fmt.Errorf("events: %w", err)
		}

		for _, f := range fl.donutSubscriptionProlonged {
			if fl.goroutine {
				go func() { f(context.WithoutCancel(ctx), obj) }()
			} else {
				f(ctx, obj)
			}
		}
	case EventDonutSubscriptionExpired:
		if len(fl.donutSubscriptionExpired) == 0 {
			break
		}

		var obj DonutSubscriptionExpiredObject

		err := json.Unmarshal(e.Object, &obj)
		if err != nil {
			return fmt.Errorf("events: %w", err)
		}

		for _, f := range fl.donutSubscriptionExpired {
			if fl.goroutine {
				go func() { f(context.WithoutCancel(ctx), obj) }()
			} else {
				f(ctx, obj)
			}
		}
	case EventDonutSubscriptionCancelled:
		if len(fl.donutSubscriptionCancelled) == 0 {
			break
		}

		var obj DonutSubscriptionCancelledObject

		err := json.Unmarshal(e.Object, &obj)
		if err != nil {
			return fmt.Errorf("events: %w", err)
		}

		for _, f := range fl.donutSubscriptionCancelled {
			if fl.goroutine {
				go func() { f(context.WithoutCancel(ctx), obj) }()
			} else {
				f(ctx, obj)
			}
		}
	case EventDonutSubscriptionPriceChanged:
		if len(fl.donutSubscriptionPriceChanged) == 0 {
			break
		}

		var obj DonutSubscriptionPriceChangedObject

		err := json.Unmarshal(e.Object, &obj)
		if err != nil {
			return fmt.Errorf("events: %w", err)
		}

		for _, f := range fl.donutSubscriptionPriceChanged {
			if fl.goroutine {
				go func() { f(context.WithoutCancel(ctx), obj) }()
			} else {
				f(ctx, obj)
			}
		}
	case EventDonutMoneyWithdraw:
		if len(fl.donutMoneyWithdraw) == 0 {
			break
		}

		var obj DonutMoneyWithdrawObject

		err := json.Unmarshal(e.Object, &obj)
		if err != nil {
			return fmt.Errorf("events: %w", err)
		}

		for _, f := range fl.donutMoneyWithdraw {
			if fl.goroutine {
				go func() { f(context.WithoutCancel(ctx), obj) }()
			} else {
				f(ctx, obj)
			}
		}
	case EventDonutMoneyWithdrawError:
		if len(fl.donutMoneyWithdrawError) == 0 {
			break
		}

		var obj DonutMoneyWithdrawErrorObject

		err := json.Unmarshal(e.Object, &obj)
		if err != nil {
			return fmt.Errorf("events: %w", err)
		}

		for _, f := range fl.donutMoneyWithdrawError {
			if fl.goroutine {
				go func() { f(context.WithoutCancel(ctx), obj) }()
			} else {
				f(ctx, obj)
			}
		}
	}

	return nil
}

// ListEvents return list of events.
func (fl *FuncList) ListEvents() []EventType {
	return fl.eventsList
}

// Goroutine invoke functions in a goroutine.
func (fl *FuncList) Goroutine(v bool) {
	fl.goroutine = v
}

// OnEvent handler.
func (fl *FuncList) OnEvent(eventType EventType, f func(context.Context, GroupEvent)) {
	fl.special[eventType] = append(fl.special[eventType], f)
	fl.eventsList = append(fl.eventsList, eventType)
}

// MessageNew handler.
func (fl *FuncList) MessageNew(f func(context.Context, MessageNewObject)) {
	fl.messageNew = append(fl.messageNew, f)
	fl.eventsList = append(fl.eventsList, EventMessageNew)
}

// MessageReply handler.
func (fl *FuncList) MessageReply(f func(context.Context, MessageReplyObject)) {
	fl.messageReply = append(fl.messageReply, f)
	fl.eventsList = append(fl.eventsList, EventMessageReply)
}

// MessageEdit handler.
func (fl *FuncList) MessageEdit(f func(context.Context, MessageEditObject)) {
	fl.messageEdit = append(fl.messageEdit, f)
	fl.eventsList = append(fl.eventsList, EventMessageEdit)
}

// MessageAllow handler.
func (fl *FuncList) MessageAllow(f func(context.Context, MessageAllowObject)) {
	fl.messageAllow = append(fl.messageAllow, f)
	fl.eventsList = append(fl.eventsList, EventMessageAllow)
}

// MessageDeny handler.
func (fl *FuncList) MessageDeny(f func(context.Context, MessageDenyObject)) {
	fl.messageDeny = append(fl.messageDeny, f)
	fl.eventsList = append(fl.eventsList, EventMessageDeny)
}

// MessageTypingState handler.
func (fl *FuncList) MessageTypingState(f func(context.Context, MessageTypingStateObject)) {
	fl.messageTypingState = append(fl.messageTypingState, f)
	fl.eventsList = append(fl.eventsList, EventMessageTypingState)
}

// MessageEvent handler.
func (fl *FuncList) MessageEvent(f func(context.Context, MessageEventObject)) {
	fl.messageEvent = append(fl.messageEvent, f)
	fl.eventsList = append(fl.eventsList, EventMessageEvent)
}

// PhotoNew handler.
func (fl *FuncList) PhotoNew(f func(context.Context, PhotoNewObject)) {
	fl.photoNew = append(fl.photoNew, f)
	fl.eventsList = append(fl.eventsList, EventPhotoNew)
}

// PhotoCommentNew handler.
func (fl *FuncList) PhotoCommentNew(f func(context.Context, PhotoCommentNewObject)) {
	fl.photoCommentNew = append(fl.photoCommentNew, f)
	fl.eventsList = append(fl.eventsList, EventPhotoCommentNew)
}

// PhotoCommentEdit handler.
func (fl *FuncList) PhotoCommentEdit(f func(context.Context, PhotoCommentEditObject)) {
	fl.photoCommentEdit = append(fl.photoCommentEdit, f)
	fl.eventsList = append(fl.eventsList, EventPhotoCommentEdit)
}

// PhotoCommentRestore handler.
func (fl *FuncList) PhotoCommentRestore(f func(context.Context, PhotoCommentRestoreObject)) {
	fl.photoCommentRestore = append(fl.photoCommentRestore, f)
	fl.eventsList = append(fl.eventsList, EventPhotoCommentRestore)
}

// PhotoCommentDelete handler.
func (fl *FuncList) PhotoCommentDelete(f func(context.Context, PhotoCommentDeleteObject)) {
	fl.photoCommentDelete = append(fl.photoCommentDelete, f)
	fl.eventsList = append(fl.eventsList, EventPhotoCommentDelete)
}

// AudioNew handler.
func (fl *FuncList) AudioNew(f func(context.Context, AudioNewObject)) {
	fl.audioNew = append(fl.audioNew, f)
	fl.eventsList = append(fl.eventsList, EventAudioNew)
}

// VideoNew handler.
func (fl *FuncList) VideoNew(f func(context.Context, VideoNewObject)) {
	fl.videoNew = append(fl.videoNew, f)
	fl.eventsList = append(fl.eventsList, EventVideoNew)
}

// VideoCommentNew handler.
func (fl *FuncList) VideoCommentNew(f func(context.Context, VideoCommentNewObject)) {
	fl.videoCommentNew = append(fl.videoCommentNew, f)
	fl.eventsList = append(fl.eventsList, EventVideoCommentNew)
}

// VideoCommentEdit handler.
func (fl *FuncList) VideoCommentEdit(f func(context.Context, VideoCommentEditObject)) {
	fl.videoCommentEdit = append(fl.videoCommentEdit, f)
	fl.eventsList = append(fl.eventsList, EventVideoCommentEdit)
}

// VideoCommentRestore handler.
func (fl *FuncList) VideoCommentRestore(f func(context.Context, VideoCommentRestoreObject)) {
	fl.videoCommentRestore = append(fl.videoCommentRestore, f)
	fl.eventsList = append(fl.eventsList, EventVideoCommentRestore)
}

// VideoCommentDelete handler.
func (fl *FuncList) VideoCommentDelete(f func(context.Context, VideoCommentDeleteObject)) {
	fl.videoCommentDelete = append(fl.videoCommentDelete, f)
	fl.eventsList = append(fl.eventsList, EventVideoCommentDelete)
}

// WallPostNew handler.
func (fl *FuncList) WallPostNew(f func(context.Context, WallPostNewObject)) {
	fl.wallPostNew = append(fl.wallPostNew, f)
	fl.eventsList = append(fl.eventsList, EventWallPostNew)
}

// WallRepost handler.
func (fl *FuncList) WallRepost(f func(context.Context, WallRepostObject)) {
	fl.wallRepost = append(fl.wallRepost, f)
	fl.eventsList = append(fl.eventsList, EventWallRepost)
}

// WallReplyNew handler.
func (fl *FuncList) WallReplyNew(f func(context.Context, WallReplyNewObject)) {
	fl.wallReplyNew = append(fl.wallReplyNew, f)
	fl.eventsList = append(fl.eventsList, EventWallReplyNew)
}

// WallReplyEdit handler.
func (fl *FuncList) WallReplyEdit(f func(context.Context, WallReplyEditObject)) {
	fl.wallReplyEdit = append(fl.wallReplyEdit, f)
	fl.eventsList = append(fl.eventsList, EventWallReplyEdit)
}

// WallReplyRestore handler.
func (fl *FuncList) WallReplyRestore(f func(context.Context, WallReplyRestoreObject)) {
	fl.wallReplyRestore = append(fl.wallReplyRestore, f)
	fl.eventsList = append(fl.eventsList, EventWallReplyRestore)
}

// WallReplyDelete handler.
func (fl *FuncList) WallReplyDelete(f func(context.Context, WallReplyDeleteObject)) {
	fl.wallReplyDelete = append(fl.wallReplyDelete, f)
	fl.eventsList = append(fl.eventsList, EventWallReplyDelete)
}

// BoardPostNew handler.
func (fl *FuncList) BoardPostNew(f func(context.Context, BoardPostNewObject)) {
	fl.boardPostNew = append(fl.boardPostNew, f)
	fl.eventsList = append(fl.eventsList, EventBoardPostNew)
}

// BoardPostEdit handler.
func (fl *FuncList) BoardPostEdit(f func(context.Context, BoardPostEditObject)) {
	fl.boardPostEdit = append(fl.boardPostEdit, f)
	fl.eventsList = append(fl.eventsList, EventBoardPostEdit)
}

// BoardPostRestore handler.
func (fl *FuncList) BoardPostRestore(f func(context.Context, BoardPostRestoreObject)) {
	fl.boardPostRestore = append(fl.boardPostRestore, f)
	fl.eventsList = append(fl.eventsList, EventBoardPostRestore)
}

// BoardPostDelete handler.
func (fl *FuncList) BoardPostDelete(f func(context.Context, BoardPostDeleteObject)) {
	fl.boardPostDelete = append(fl.boardPostDelete, f)
	fl.eventsList = append(fl.eventsList, EventBoardPostDelete)
}

// MarketCommentNew handler.
func (fl *FuncList) MarketCommentNew(f func(context.Context, MarketCommentNewObject)) {
	fl.marketCommentNew = append(fl.marketCommentNew, f)
	fl.eventsList = append(fl.eventsList, EventMarketCommentNew)
}

// MarketCommentEdit handler.
func (fl *FuncList) MarketCommentEdit(f func(context.Context, MarketCommentEditObject)) {
	fl.marketCommentEdit = append(fl.marketCommentEdit, f)
	fl.eventsList = append(fl.eventsList, EventMarketCommentEdit)
}

// MarketCommentRestore handler.
func (fl *FuncList) MarketCommentRestore(f func(context.Context, MarketCommentRestoreObject)) {
	fl.marketCommentRestore = append(fl.marketCommentRestore, f)
	fl.eventsList = append(fl.eventsList, EventMarketCommentRestore)
}

// MarketCommentDelete handler.
func (fl *FuncList) MarketCommentDelete(f func(context.Context, MarketCommentDeleteObject)) {
	fl.marketCommentDelete = append(fl.marketCommentDelete, f)
	fl.eventsList = append(fl.eventsList, EventMarketCommentDelete)
}

// MarketOrderNew handler.
func (fl *FuncList) MarketOrderNew(f func(context.Context, MarketOrderNewObject)) {
	fl.marketOrderNew = append(fl.marketOrderNew, f)
	fl.eventsList = append(fl.eventsList, EventMarketOrderNew)
}

// MarketOrderEdit handler.
func (fl *FuncList) MarketOrderEdit(f func(context.Context, MarketOrderEditObject)) {
	fl.marketOrderEdit = append(fl.marketOrderEdit, f)
	fl.eventsList = append(fl.eventsList, EventMarketOrderEdit)
}

// GroupLeave handler.
func (fl *FuncList) GroupLeave(f func(context.Context, GroupLeaveObject)) {
	fl.groupLeave = append(fl.groupLeave, f)
	fl.eventsList = append(fl.eventsList, EventGroupLeave)
}

// GroupJoin handler.
func (fl *FuncList) GroupJoin(f func(context.Context, GroupJoinObject)) {
	fl.groupJoin = append(fl.groupJoin, f)
	fl.eventsList = append(fl.eventsList, EventGroupJoin)
}

// UserBlock handler.
func (fl *FuncList) UserBlock(f func(context.Context, UserBlockObject)) {
	fl.userBlock = append(fl.userBlock, f)
	fl.eventsList = append(fl.eventsList, EventUserBlock)
}

// UserUnblock handler.
func (fl *FuncList) UserUnblock(f func(context.Context, UserUnblockObject)) {
	fl.userUnblock = append(fl.userUnblock, f)
	fl.eventsList = append(fl.eventsList, EventUserUnblock)
}

// PollVoteNew handler.
func (fl *FuncList) PollVoteNew(f func(context.Context, PollVoteNewObject)) {
	fl.pollVoteNew = append(fl.pollVoteNew, f)
	fl.eventsList = append(fl.eventsList, EventPollVoteNew)
}

// GroupOfficersEdit handler.
func (fl *FuncList) GroupOfficersEdit(f func(context.Context, GroupOfficersEditObject)) {
	fl.groupOfficersEdit = append(fl.groupOfficersEdit, f)
	fl.eventsList = append(fl.eventsList, EventGroupOfficersEdit)
}

// GroupChangeSettings handler.
func (fl *FuncList) GroupChangeSettings(f func(context.Context, GroupChangeSettingsObject)) {
	fl.groupChangeSettings = append(fl.groupChangeSettings, f)
	fl.eventsList = append(fl.eventsList, EventGroupChangeSettings)
}

// GroupChangePhoto handler.
func (fl *FuncList) GroupChangePhoto(f func(context.Context, GroupChangePhotoObject)) {
	fl.groupChangePhoto = append(fl.groupChangePhoto, f)
	fl.eventsList = append(fl.eventsList, EventGroupChangePhoto)
}

// VkpayTransaction handler.
func (fl *FuncList) VkpayTransaction(f func(context.Context, VkpayTransactionObject)) {
	fl.vkpayTransaction = append(fl.vkpayTransaction, f)
	fl.eventsList = append(fl.eventsList, EventVkpayTransaction)
}

// LeadFormsNew handler.
func (fl *FuncList) LeadFormsNew(f func(context.Context, LeadFormsNewObject)) {
	fl.leadFormsNew = append(fl.leadFormsNew, f)
	fl.eventsList = append(fl.eventsList, EventLeadFormsNew)
}

// AppPayload handler.
func (fl *FuncList) AppPayload(f func(context.Context, AppPayloadObject)) {
	fl.appPayload = append(fl.appPayload, f)
	fl.eventsList = append(fl.eventsList, EventAppPayload)
}

// MessageRead handler.
func (fl *FuncList) MessageRead(f func(context.Context, MessageReadObject)) {
	fl.messageRead = append(fl.messageRead, f)
	fl.eventsList = append(fl.eventsList, EventMessageRead)
}

// LikeAdd handler.
func (fl *FuncList) LikeAdd(f func(context.Context, LikeAddObject)) {
	fl.likeAdd = append(fl.likeAdd, f)
	fl.eventsList = append(fl.eventsList, EventLikeAdd)
}

// LikeRemove handler.
func (fl *FuncList) LikeRemove(f func(context.Context, LikeRemoveObject)) {
	fl.likeRemove = append(fl.likeRemove, f)
	fl.eventsList = append(fl.eventsList, EventLikeRemove)
}

// DonutSubscriptionCreate handler.
func (fl *FuncList) DonutSubscriptionCreate(f func(context.Context, DonutSubscriptionCreateObject)) {
	fl.donutSubscriptionCreate = append(fl.donutSubscriptionCreate, f)
	fl.eventsList = append(fl.eventsList, EventDonutSubscriptionCreate)
}

// DonutSubscriptionProlonged handler.
func (fl *FuncList) DonutSubscriptionProlonged(f func(context.Context, DonutSubscriptionProlongedObject)) {
	fl.donutSubscriptionProlonged = append(fl.donutSubscriptionProlonged, f)
	fl.eventsList = append(fl.eventsList, EventDonutSubscriptionProlonged)
}

// DonutSubscriptionExpired handler.
func (fl *FuncList) DonutSubscriptionExpired(f func(context.Context, DonutSubscriptionExpiredObject)) {
	fl.donutSubscriptionExpired = append(fl.donutSubscriptionExpired, f)
	fl.eventsList = append(fl.eventsList, EventDonutSubscriptionExpired)
}

// DonutSubscriptionCancelled handler.
func (fl *FuncList) DonutSubscriptionCancelled(f func(context.Context, DonutSubscriptionCancelledObject)) {
	fl.donutSubscriptionCancelled = append(fl.donutSubscriptionCancelled, f)
	fl.eventsList = append(fl.eventsList, EventDonutSubscriptionCancelled)
}

// DonutSubscriptionPriceChanged handler.
func (fl *FuncList) DonutSubscriptionPriceChanged(f func(context.Context, DonutSubscriptionPriceChangedObject)) {
	fl.donutSubscriptionPriceChanged = append(fl.donutSubscriptionPriceChanged, f)
	fl.eventsList = append(fl.eventsList, EventDonutSubscriptionPriceChanged)
}

// DonutMoneyWithdraw handler.
func (fl *FuncList) DonutMoneyWithdraw(f func(context.Context, DonutMoneyWithdrawObject)) {
	fl.donutMoneyWithdraw = append(fl.donutMoneyWithdraw, f)
	fl.eventsList = append(fl.eventsList, EventDonutMoneyWithdraw)
}

// DonutMoneyWithdrawError handler.
func (fl *FuncList) DonutMoneyWithdrawError(f func(context.Context, DonutMoneyWithdrawErrorObject)) {
	fl.donutMoneyWithdrawError = append(fl.donutMoneyWithdrawError, f)
	fl.eventsList = append(fl.eventsList, EventDonutMoneyWithdrawError)
}
