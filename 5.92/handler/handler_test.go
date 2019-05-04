package handler

import (
	"testing"

	"github.com/SevereCloud/vksdk/5.92/object"
)

const GID = 123456

func TestFuncList_Handler(t *testing.T) {
	funcList := FuncList{
		MessageNew: []object.MessageNewFunc{
			func(obj object.MessageNewObject, groupID int) {
				if groupID != GID {
					t.Errorf("FuncList.Handler() MessageNewFunc groupID != GID")
				}
				// TODO: check obj
			},
		},
		MessageReply: []object.MessageReplyFunc{
			func(obj object.MessageReplyObject, groupID int) {
				if groupID != GID {
					t.Errorf("FuncList.Handler() MessageReplyFunc groupID != GID")
				}
				// TODO: check obj
			},
		},
		MessageEdit: []object.MessageEditFunc{
			func(obj object.MessageEditObject, groupID int) {
				if groupID != GID {
					t.Errorf("FuncList.Handler() MessageEditFunc groupID != GID")
				}
				// TODO: check obj
			},
		},
		MessageAllow: []object.MessageAllowFunc{
			func(obj object.MessageAllowObject, groupID int) {
				if groupID != GID {
					t.Errorf("FuncList.Handler() MessageAllowFunc groupID != GID")
				}
				// TODO: check obj
			},
		},
		MessageDeny: []object.MessageDenyFunc{
			func(obj object.MessageDenyObject, groupID int) {
				if groupID != GID {
					t.Errorf("FuncList.Handler() MessageDenyFunc groupID != GID")
				}
				// TODO: check obj
			},
		},
		MessageTypingState: []object.MessageTypingStateFunc{
			func(obj object.MessageTypingStateObject, groupID int) {
				if groupID != GID {
					t.Errorf("FuncList.Handler() MessageTypingStateFunc groupID != GID")
				}
				// TODO: check obj
			},
		},
		PhotoNew: []object.PhotoNewFunc{
			func(obj object.PhotoNewObject, groupID int) {
				if groupID != GID {
					t.Errorf("FuncList.Handler() PhotoNewFunc groupID != GID")
				}
				// TODO: check obj
			},
		},
		PhotoCommentNew: []object.PhotoCommentNewFunc{
			func(obj object.PhotoCommentNewObject, groupID int) {
				if groupID != GID {
					t.Errorf("FuncList.Handler() PhotoCommentNewFunc groupID != GID")
				}
				// TODO: check obj
			},
		},
		PhotoCommentEdit: []object.PhotoCommentEditFunc{
			func(obj object.PhotoCommentEditObject, groupID int) {
				if groupID != GID {
					t.Errorf("FuncList.Handler() PhotoCommentEditFunc groupID != GID")
				}
				// TODO: check obj
			},
		},
		PhotoCommentRestore: []object.PhotoCommentRestoreFunc{
			func(obj object.PhotoCommentRestoreObject, groupID int) {
				if groupID != GID {
					t.Errorf("FuncList.Handler() PhotoCommentRestoreFunc groupID != GID")
				}
				// TODO: check obj
			},
		},
		PhotoCommentDelete: []object.PhotoCommentDeleteFunc{
			func(obj object.PhotoCommentDeleteObject, groupID int) {
				if groupID != GID {
					t.Errorf("FuncList.Handler() PhotoCommentDeleteFunc groupID != GID")
				}
				// TODO: check obj
			},
		},
		AudioNew: []object.AudioNewFunc{
			func(obj object.AudioNewObject, groupID int) {
				if groupID != GID {
					t.Errorf("FuncList.Handler() AudioNewFunc groupID != GID")
				}
				// TODO: check obj
			},
		},
		VideoNew: []object.VideoNewFunc{
			func(obj object.VideoNewObject, groupID int) {
				if groupID != GID {
					t.Errorf("FuncList.Handler() VideoNewFunc groupID != GID")
				}
				// TODO: check obj
			},
		},
		VideoCommentNew: []object.VideoCommentNewFunc{
			func(obj object.VideoCommentNewObject, groupID int) {
				if groupID != GID {
					t.Errorf("FuncList.Handler() VideoCommentNewFunc groupID != GID")
				}
				// TODO: check obj
			},
		},
		VideoCommentEdit: []object.VideoCommentEditFunc{
			func(obj object.VideoCommentEditObject, groupID int) {
				if groupID != GID {
					t.Errorf("FuncList.Handler() VideoCommentEditFunc groupID != GID")
				}
				// TODO: check obj
			},
		},
		VideoCommentRestore: []object.VideoCommentRestoreFunc{
			func(obj object.VideoCommentRestoreObject, groupID int) {
				if groupID != GID {
					t.Errorf("FuncList.Handler() VideoCommentRestoreFunc groupID != GID")
				}
				// TODO: check obj
			},
		},
		VideoCommentDelete: []object.VideoCommentDeleteFunc{
			func(obj object.VideoCommentDeleteObject, groupID int) {
				if groupID != GID {
					t.Errorf("FuncList.Handler() VideoCommentDeleteFunc groupID != GID")
				}
				// TODO: check obj
			},
		},
		WallPostNew: []object.WallPostNewFunc{
			func(obj object.WallPostNewObject, groupID int) {
				if groupID != GID {
					t.Errorf("FuncList.Handler() WallPostNewFunc groupID != GID")
				}
				// TODO: check obj
			},
		},
		WallRepost: []object.WallRepostFunc{
			func(obj object.WallRepostObject, groupID int) {
				if groupID != GID {
					t.Errorf("FuncList.Handler() WallRepostFunc groupID != GID")
				}
				// TODO: check obj
			},
		},
		WallReplyNew: []object.WallReplyNewFunc{
			func(obj object.WallReplyNewObject, groupID int) {
				if groupID != GID {
					t.Errorf("FuncList.Handler() WallReplyNewFunc groupID != GID")
				}
				// TODO: check obj
			},
		},
		WallReplyEdit: []object.WallReplyEditFunc{
			func(obj object.WallReplyEditObject, groupID int) {
				if groupID != GID {
					t.Errorf("FuncList.Handler() WallReplyEditFunc groupID != GID")
				}
				// TODO: check obj
			},
		},
		WallReplyRestore: []object.WallReplyRestoreFunc{
			func(obj object.WallReplyRestoreObject, groupID int) {
				if groupID != GID {
					t.Errorf("FuncList.Handler() WallReplyRestoreFunc groupID != GID")
				}
				// TODO: check obj
			},
		},
		WallReplyDelete: []object.WallReplyDeleteFunc{
			func(obj object.WallReplyDeleteObject, groupID int) {
				if groupID != GID {
					t.Errorf("FuncList.Handler() WallReplyDeleteFunc groupID != GID")
				}
				// TODO: check obj
			},
		},
		BoardPostNew: []object.BoardPostNewFunc{
			func(obj object.BoardPostNewObject, groupID int) {
				if groupID != GID {
					t.Errorf("FuncList.Handler() BoardPostNewFunc groupID != GID")
				}
				// TODO: check obj
			},
		},
		BoardPostEdit: []object.BoardPostEditFunc{
			func(obj object.BoardPostEditObject, groupID int) {
				if groupID != GID {
					t.Errorf("FuncList.Handler() BoardPostEditFunc groupID != GID")
				}
				// TODO: check obj
			},
		},
		BoardPostRestore: []object.BoardPostRestoreFunc{
			func(obj object.BoardPostRestoreObject, groupID int) {
				if groupID != GID {
					t.Errorf("FuncList.Handler() BoardPostRestoreFunc groupID != GID")
				}
				// TODO: check obj
			},
		},
		BoardPostDelete: []object.BoardPostDeleteFunc{
			func(obj object.BoardPostDeleteObject, groupID int) {
				if groupID != GID {
					t.Errorf("FuncList.Handler() BoardPostDeleteFunc groupID != GID")
				}
				// TODO: check obj
			},
		},
		MarketCommentNew: []object.MarketCommentNewFunc{
			func(obj object.MarketCommentNewObject, groupID int) {
				if groupID != GID {
					t.Errorf("FuncList.Handler() MarketCommentNewFunc groupID != GID")
				}
				// TODO: check obj
			},
		},
		MarketCommentEdit: []object.MarketCommentEditFunc{
			func(obj object.MarketCommentEditObject, groupID int) {
				if groupID != GID {
					t.Errorf("FuncList.Handler() MarketCommentEditFunc groupID != GID")
				}
				// TODO: check obj
			},
		},
		MarketCommentRestore: []object.MarketCommentRestoreFunc{
			func(obj object.MarketCommentRestoreObject, groupID int) {
				if groupID != GID {
					t.Errorf("FuncList.Handler() MarketCommentRestoreFunc groupID != GID")
				}
				// TODO: check obj
			},
		},
		MarketCommentDelete: []object.MarketCommentDeleteFunc{
			func(obj object.MarketCommentDeleteObject, groupID int) {
				if groupID != GID {
					t.Errorf("FuncList.Handler() MarketCommentDeleteFunc groupID != GID")
				}
				// TODO: check obj
			},
		},
		GroupLeave: []object.GroupLeaveFunc{
			func(obj object.GroupLeaveObject, groupID int) {
				if groupID != GID {
					t.Errorf("FuncList.Handler() GroupLeaveFunc groupID != GID")
				}
				// TODO: check obj
			},
		},
		GroupJoin: []object.GroupJoinFunc{
			func(obj object.GroupJoinObject, groupID int) {
				if groupID != GID {
					t.Errorf("FuncList.Handler() GroupJoinFunc groupID != GID")
				}
				// TODO: check obj
			},
		},
		UserBlock: []object.UserBlockFunc{
			func(obj object.UserBlockObject, groupID int) {
				if groupID != GID {
					t.Errorf("FuncList.Handler() UserBlockFunc groupID != GID")
				}
				// TODO: check obj
			},
		},
		UserUnblock: []object.UserUnblockFunc{
			func(obj object.UserUnblockObject, groupID int) {
				if groupID != GID {
					t.Errorf("FuncList.Handler() UserUnblockFunc groupID != GID")
				}
				// TODO: check obj
			},
		},
		PollVoteNew: []object.PollVoteNewFunc{
			func(obj object.PollVoteNewObject, groupID int) {
				if groupID != GID {
					t.Errorf("FuncList.Handler() PollVoteNewFunc groupID != GID")
				}
				// TODO: check obj
			},
		},
		GroupOfficersEdit: []object.GroupOfficersEditFunc{
			func(obj object.GroupOfficersEditObject, groupID int) {
				if groupID != GID {
					t.Errorf("FuncList.Handler() GroupOfficersEditFunc groupID != GID")
				}
				// TODO: check obj
			},
		},
		GroupChangeSettings: []object.GroupChangeSettingsFunc{
			func(obj object.GroupChangeSettingsObject, groupID int) {
				if groupID != GID {
					t.Errorf("FuncList.Handler() GroupChangeSettingsFunc groupID != GID")
				}
				// TODO: check obj
			},
		},
		GroupChangePhoto: []object.GroupChangePhotoFunc{
			func(obj object.GroupChangePhotoObject, groupID int) {
				if groupID != GID {
					t.Errorf("FuncList.Handler() GroupChangePhotoFunc groupID != GID")
				}
				// TODO: check obj
			},
		},
		VkpayTransaction: []object.VkpayTransactionFunc{
			func(obj object.VkpayTransactionObject, groupID int) {
				if groupID != GID {
					t.Errorf("FuncList.Handler() VkpayTransactionFunc groupID != GID")
				}
				// TODO: check obj
			},
		},
		LeadFormsNew: []object.LeadFormsNewFunc{
			func(obj object.LeadFormsNewObject, groupID int) {
				if groupID != GID {
					t.Errorf("FuncList.Handler() LeadFormsNewFunc groupID != GID")
				}
				// TODO: check obj
			},
		},
	}
	tests := []struct {
		name    string
		argE    object.GroupEvent
		wantErr bool
	}{
		{
			name: "message_new",
			argE: object.GroupEvent{
				Type:    "message_new",
				Object:  []byte("{}"),
				GroupID: GID,
			},
			wantErr: false,
		},
		{
			name: "message_new error",
			argE: object.GroupEvent{
				Type:   "message_new",
				Object: []byte(""),
			},
			wantErr: true,
		},
		{
			name: "message_reply",
			argE: object.GroupEvent{
				Type:    "message_reply",
				Object:  []byte("{}"),
				GroupID: GID,
			},
			wantErr: false,
		},
		{
			name: "message_reply error",
			argE: object.GroupEvent{
				Type:   "message_reply",
				Object: []byte(""),
			},
			wantErr: true,
		},
		{
			name: "message_edit",
			argE: object.GroupEvent{
				Type:    "message_edit",
				Object:  []byte("{}"),
				GroupID: GID,
			},
			wantErr: false,
		},
		{
			name: "message_edit error",
			argE: object.GroupEvent{
				Type:   "message_edit",
				Object: []byte(""),
			},
			wantErr: true,
		},
		{
			name: "message_allow",
			argE: object.GroupEvent{
				Type:    "message_allow",
				Object:  []byte("{}"),
				GroupID: GID,
			},
			wantErr: false,
		},
		{
			name: "message_allow error",
			argE: object.GroupEvent{
				Type:   "message_allow",
				Object: []byte(""),
			},
			wantErr: true,
		},
		{
			name: "messages_deny",
			argE: object.GroupEvent{
				Type:    "messages_deny",
				Object:  []byte("{}"),
				GroupID: GID,
			},
			wantErr: false,
		},
		{
			name: "messages_deny error",
			argE: object.GroupEvent{
				Type:   "messages_deny",
				Object: []byte(""),
			},
			wantErr: true,
		},
		{
			name: "message_typing_state",
			argE: object.GroupEvent{
				Type:    "message_typing_state",
				Object:  []byte("{}"),
				GroupID: GID,
			},
			wantErr: false,
		},
		{
			name: "message_typing_state error",
			argE: object.GroupEvent{
				Type:   "message_typing_state",
				Object: []byte(""),
			},
			wantErr: true,
		},
		{
			name: "photo_new",
			argE: object.GroupEvent{
				Type:    "photo_new",
				Object:  []byte("{}"),
				GroupID: GID,
			},
			wantErr: false,
		},
		{
			name: "photo_new error",
			argE: object.GroupEvent{
				Type:   "photo_new",
				Object: []byte(""),
			},
			wantErr: true,
		},
		{
			name: "photo_comment_new",
			argE: object.GroupEvent{
				Type:    "photo_comment_new",
				Object:  []byte("{}"),
				GroupID: GID,
			},
			wantErr: false,
		},
		{
			name: "photo_comment_new error",
			argE: object.GroupEvent{
				Type:   "photo_comment_new",
				Object: []byte(""),
			},
			wantErr: true,
		},
		{
			name: "photo_comment_edit",
			argE: object.GroupEvent{
				Type:    "photo_comment_edit",
				Object:  []byte("{}"),
				GroupID: GID,
			},
			wantErr: false,
		},
		{
			name: "photo_comment_edit error",
			argE: object.GroupEvent{
				Type:   "photo_comment_edit",
				Object: []byte(""),
			},
			wantErr: true,
		},
		{
			name: "photo_comment_restore",
			argE: object.GroupEvent{
				Type:    "photo_comment_restore",
				Object:  []byte("{}"),
				GroupID: GID,
			},
			wantErr: false,
		},
		{
			name: "photo_comment_restore error",
			argE: object.GroupEvent{
				Type:   "photo_comment_restore",
				Object: []byte(""),
			},
			wantErr: true,
		},
		{
			name: "photo_comment_delete",
			argE: object.GroupEvent{
				Type:    "photo_comment_delete",
				Object:  []byte("{}"),
				GroupID: GID,
			},
			wantErr: false,
		},
		{
			name: "photo_comment_delete error",
			argE: object.GroupEvent{
				Type:   "photo_comment_delete",
				Object: []byte(""),
			},
			wantErr: true,
		},
		{
			name: "audio_new",
			argE: object.GroupEvent{
				Type:    "audio_new",
				Object:  []byte("{}"),
				GroupID: GID,
			},
			wantErr: false,
		},
		{
			name: "audio_new error",
			argE: object.GroupEvent{
				Type:   "audio_new",
				Object: []byte(""),
			},
			wantErr: true,
		},
		{
			name: "video_new",
			argE: object.GroupEvent{
				Type:    "video_new",
				Object:  []byte("{}"),
				GroupID: GID,
			},
			wantErr: false,
		},
		{
			name: "video_new error",
			argE: object.GroupEvent{
				Type:   "video_new",
				Object: []byte(""),
			},
			wantErr: true,
		},
		{
			name: "video_comment_new",
			argE: object.GroupEvent{
				Type:    "video_comment_new",
				Object:  []byte("{}"),
				GroupID: GID,
			},
			wantErr: false,
		},
		{
			name: "video_comment_new error",
			argE: object.GroupEvent{
				Type:   "video_comment_new",
				Object: []byte(""),
			},
			wantErr: true,
		},
		{
			name: "video_comment_edit",
			argE: object.GroupEvent{
				Type:    "video_comment_edit",
				Object:  []byte("{}"),
				GroupID: GID,
			},
			wantErr: false,
		},
		{
			name: "video_comment_edit error",
			argE: object.GroupEvent{
				Type:   "video_comment_edit",
				Object: []byte(""),
			},
			wantErr: true,
		},
		{
			name: "video_comment_restore",
			argE: object.GroupEvent{
				Type:    "video_comment_restore",
				Object:  []byte("{}"),
				GroupID: GID,
			},
			wantErr: false,
		},
		{
			name: "video_comment_restore error",
			argE: object.GroupEvent{
				Type:   "video_comment_restore",
				Object: []byte(""),
			},
			wantErr: true,
		},
		{
			name: "video_comment_delete",
			argE: object.GroupEvent{
				Type:    "video_comment_delete",
				Object:  []byte("{}"),
				GroupID: GID,
			},
			wantErr: false,
		},
		{
			name: "video_comment_delete error",
			argE: object.GroupEvent{
				Type:   "video_comment_delete",
				Object: []byte(""),
			},
			wantErr: true,
		},
		{
			name: "wall_post_new",
			argE: object.GroupEvent{
				Type:    "wall_post_new",
				Object:  []byte("{}"),
				GroupID: GID,
			},
			wantErr: false,
		},
		{
			name: "wall_post_new error",
			argE: object.GroupEvent{
				Type:   "wall_post_new",
				Object: []byte(""),
			},
			wantErr: true,
		},
		{
			name: "wall_repost",
			argE: object.GroupEvent{
				Type:    "wall_repost",
				Object:  []byte("{}"),
				GroupID: GID,
			},
			wantErr: false,
		},
		{
			name: "wall_repost error",
			argE: object.GroupEvent{
				Type:   "wall_repost",
				Object: []byte(""),
			},
			wantErr: true,
		},
		{
			name: "wall_reply_new",
			argE: object.GroupEvent{
				Type:    "wall_reply_new",
				Object:  []byte("{}"),
				GroupID: GID,
			},
			wantErr: false,
		},
		{
			name: "wall_reply_new error",
			argE: object.GroupEvent{
				Type:   "wall_reply_new",
				Object: []byte(""),
			},
			wantErr: true,
		},
		{
			name: "wall_reply_edit",
			argE: object.GroupEvent{
				Type:    "wall_reply_edit",
				Object:  []byte("{}"),
				GroupID: GID,
			},
			wantErr: false,
		},
		{
			name: "wall_reply_edit error",
			argE: object.GroupEvent{
				Type:   "wall_reply_edit",
				Object: []byte(""),
			},
			wantErr: true,
		},
		{
			name: "wall_reply_restore",
			argE: object.GroupEvent{
				Type:    "wall_reply_restore",
				Object:  []byte("{}"),
				GroupID: GID,
			},
			wantErr: false,
		},
		{
			name: "wall_reply_restore error",
			argE: object.GroupEvent{
				Type:   "wall_reply_restore",
				Object: []byte(""),
			},
			wantErr: true,
		},
		{
			name: "wall_reply_delete",
			argE: object.GroupEvent{
				Type:    "wall_reply_delete",
				Object:  []byte("{}"),
				GroupID: GID,
			},
			wantErr: false,
		},
		{
			name: "wall_reply_delete error",
			argE: object.GroupEvent{
				Type:   "wall_reply_delete",
				Object: []byte(""),
			},
			wantErr: true,
		},
		{
			name: "board_post_new",
			argE: object.GroupEvent{
				Type:    "board_post_new",
				Object:  []byte("{}"),
				GroupID: GID,
			},
			wantErr: false,
		},
		{
			name: "board_post_new error",
			argE: object.GroupEvent{
				Type:   "board_post_new",
				Object: []byte(""),
			},
			wantErr: true,
		},
		{
			name: "board_post_edit",
			argE: object.GroupEvent{
				Type:    "board_post_edit",
				Object:  []byte("{}"),
				GroupID: GID,
			},
			wantErr: false,
		},
		{
			name: "board_post_edit error",
			argE: object.GroupEvent{
				Type:   "board_post_edit",
				Object: []byte(""),
			},
			wantErr: true,
		},
		{
			name: "board_post_restore",
			argE: object.GroupEvent{
				Type:    "board_post_restore",
				Object:  []byte("{}"),
				GroupID: GID,
			},
			wantErr: false,
		},
		{
			name: "board_post_restore error",
			argE: object.GroupEvent{
				Type:   "board_post_restore",
				Object: []byte(""),
			},
			wantErr: true,
		},
		{
			name: "board_post_delete",
			argE: object.GroupEvent{
				Type:    "board_post_delete",
				Object:  []byte("{}"),
				GroupID: GID,
			},
			wantErr: false,
		},
		{
			name: "board_post_delete error",
			argE: object.GroupEvent{
				Type:   "board_post_delete",
				Object: []byte(""),
			},
			wantErr: true,
		},
		{
			name: "market_comment_new",
			argE: object.GroupEvent{
				Type:    "market_comment_new",
				Object:  []byte("{}"),
				GroupID: GID,
			},
			wantErr: false,
		},
		{
			name: "market_comment_new error",
			argE: object.GroupEvent{
				Type:   "market_comment_new",
				Object: []byte(""),
			},
			wantErr: true,
		},
		{
			name: "market_comment_edit",
			argE: object.GroupEvent{
				Type:    "market_comment_edit",
				Object:  []byte("{}"),
				GroupID: GID,
			},
			wantErr: false,
		},
		{
			name: "market_comment_edit error",
			argE: object.GroupEvent{
				Type:   "market_comment_edit",
				Object: []byte(""),
			},
			wantErr: true,
		},
		{
			name: "market_comment_restore",
			argE: object.GroupEvent{
				Type:    "market_comment_restore",
				Object:  []byte("{}"),
				GroupID: GID,
			},
			wantErr: false,
		},
		{
			name: "market_comment_restore error",
			argE: object.GroupEvent{
				Type:   "market_comment_restore",
				Object: []byte(""),
			},
			wantErr: true,
		},
		{
			name: "market_comment_delete",
			argE: object.GroupEvent{
				Type:    "market_comment_delete",
				Object:  []byte("{}"),
				GroupID: GID,
			},
			wantErr: false,
		},
		{
			name: "market_comment_delete error",
			argE: object.GroupEvent{
				Type:   "market_comment_delete",
				Object: []byte(""),
			},
			wantErr: true,
		},
		{
			name: "group_leave",
			argE: object.GroupEvent{
				Type:    "group_leave",
				Object:  []byte("{}"),
				GroupID: GID,
			},
			wantErr: false,
		},
		{
			name: "group_leave error",
			argE: object.GroupEvent{
				Type:   "group_leave",
				Object: []byte(""),
			},
			wantErr: true,
		},
		{
			name: "group_join",
			argE: object.GroupEvent{
				Type:    "group_join",
				Object:  []byte("{}"),
				GroupID: GID,
			},
			wantErr: false,
		},
		{
			name: "group_join error",
			argE: object.GroupEvent{
				Type:   "group_join",
				Object: []byte(""),
			},
			wantErr: true,
		},
		{
			name: "user_block",
			argE: object.GroupEvent{
				Type:    "user_block",
				Object:  []byte("{}"),
				GroupID: GID,
			},
			wantErr: false,
		},
		{
			name: "user_block error",
			argE: object.GroupEvent{
				Type:   "user_block",
				Object: []byte(""),
			},
			wantErr: true,
		},
		{
			name: "user_unblock",
			argE: object.GroupEvent{
				Type:    "user_unblock",
				Object:  []byte("{}"),
				GroupID: GID,
			},
			wantErr: false,
		},
		{
			name: "user_unblock error",
			argE: object.GroupEvent{
				Type:   "user_unblock",
				Object: []byte(""),
			},
			wantErr: true,
		},
		{
			name: "poll_vote_new",
			argE: object.GroupEvent{
				Type:    "poll_vote_new",
				Object:  []byte("{}"),
				GroupID: GID,
			},
			wantErr: false,
		},
		{
			name: "poll_vote_new error",
			argE: object.GroupEvent{
				Type:   "poll_vote_new",
				Object: []byte(""),
			},
			wantErr: true,
		},
		{
			name: "group_officers_edit",
			argE: object.GroupEvent{
				Type:    "group_officers_edit",
				Object:  []byte("{}"),
				GroupID: GID,
			},
			wantErr: false,
		},
		{
			name: "group_officers_edit error",
			argE: object.GroupEvent{
				Type:   "group_officers_edit",
				Object: []byte(""),
			},
			wantErr: true,
		},
		{
			name: "group_change_settings",
			argE: object.GroupEvent{
				Type:    "group_change_settings",
				Object:  []byte("{}"),
				GroupID: GID,
			},
			wantErr: false,
		},
		{
			name: "group_change_settings error",
			argE: object.GroupEvent{
				Type:   "group_change_settings",
				Object: []byte(""),
			},
			wantErr: true,
		},
		{
			name: "group_change_photo",
			argE: object.GroupEvent{
				Type:    "group_change_photo",
				Object:  []byte("{}"),
				GroupID: GID,
			},
			wantErr: false,
		},
		{
			name: "group_change_photo error",
			argE: object.GroupEvent{
				Type:   "group_change_photo",
				Object: []byte(""),
			},
			wantErr: true,
		},
		{
			name: "vkpay_transaction",
			argE: object.GroupEvent{
				Type:    "vkpay_transaction",
				Object:  []byte("{}"),
				GroupID: GID,
			},
			wantErr: false,
		},
		{
			name: "vkpay_transaction error",
			argE: object.GroupEvent{
				Type:   "vkpay_transaction",
				Object: []byte(""),
			},
			wantErr: true,
		},
		{
			name: "lead_forms_new",
			argE: object.GroupEvent{
				Type:    "lead_forms_new",
				Object:  []byte("{}"),
				GroupID: GID,
			},
			wantErr: false,
		},
		{
			name: "lead_forms_new error",
			argE: object.GroupEvent{
				Type:   "lead_forms_new",
				Object: []byte(""),
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := funcList.Handler(tt.argE); (err != nil) != tt.wantErr {
				t.Errorf("FuncList.Handler() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
