package handler

import (
	"encoding/json"

	"github.com/severecloud/vksdk/object"
)

// FuncList struct
type FuncList struct {
	MessageNew           []object.MessageNewFunc
	MessageReply         []object.MessageReplyFunc
	MessageEdit          []object.MessageEditFunc
	MessageAllow         []object.MessageAllowFunc
	MessageDeny          []object.MessageDenyFunc
	MessageTypingState   []object.MessageTypingStateFunc
	PhotoNew             []object.PhotoNewFunc
	PhotoCommentNew      []object.PhotoCommentNewFunc
	PhotoCommentEdit     []object.PhotoCommentEditFunc
	PhotoCommentRestore  []object.PhotoCommentRestoreFunc
	PhotoCommentDelete   []object.PhotoCommentDeleteFunc
	AudioNew             []object.AudioNewFunc
	VideoNew             []object.VideoNewFunc
	VideoCommentNew      []object.VideoCommentNewFunc
	VideoCommentEdit     []object.VideoCommentEditFunc
	VideoCommentRestore  []object.VideoCommentRestoreFunc
	VideoCommentDelete   []object.VideoCommentDeleteFunc
	WallPostNew          []object.WallPostNewFunc
	WallRepost           []object.WallRepostFunc
	WallReplyNew         []object.WallReplyNewFunc
	WallReplyEdit        []object.WallReplyEditFunc
	WallReplyRestore     []object.WallReplyRestoreFunc
	WallReplyDelete      []object.WallReplyDeleteFunc
	BoardPostNew         []object.BoardPostNewFunc
	BoardPostEdit        []object.BoardPostEditFunc
	BoardPostRestore     []object.BoardPostRestoreFunc
	BoardPostDelete      []object.BoardPostDeleteFunc
	MarketCommentNew     []object.MarketCommentNewFunc
	MarketCommentEdit    []object.MarketCommentEditFunc
	MarketCommentRestore []object.MarketCommentRestoreFunc
	MarketCommentDelete  []object.MarketCommentDeleteFunc
	GroupLeave           []object.GroupLeaveFunc
	GroupJoin            []object.GroupJoinFunc
	UserBlock            []object.UserBlockFunc
	UserUnblock          []object.UserUnblockFunc
	PollVoteNew          []object.PollVoteNewFunc
	GroupOfficersEdit    []object.GroupOfficersEditFunc
	GroupChangeSettings  []object.GroupChangeSettingsFunc
	GroupChangePhoto     []object.GroupChangePhotoFunc
	VkpayTransaction     []object.VkpayTransactionFunc
}

// Handler group event hundler
func (funcList FuncList) Handler(e object.GroupEvent) (err error) {
	switch e.Type {
	case "message_new":
		var obj object.MessageNewObject
		if err := json.Unmarshal(e.Object, &obj); err != nil {
			return err
		}

		for _, f := range funcList.MessageNew {
			f(obj, e.GroupID)
		}
	case "message_reply":
		var obj object.MessageReplyObject
		if err := json.Unmarshal(e.Object, &obj); err != nil {
			return err
		}

		for _, f := range funcList.MessageReply {
			f(obj, e.GroupID)
		}
	case "message_edit":
		var obj object.MessageEditObject
		if err := json.Unmarshal(e.Object, &obj); err != nil {
			return err
		}

		for _, f := range funcList.MessageEdit {
			f(obj, e.GroupID)
		}
	case "message_allow":
		var obj object.MessageAllowObject
		if err := json.Unmarshal(e.Object, &obj); err != nil {
			return err
		}

		for _, f := range funcList.MessageAllow {
			f(obj, e.GroupID)
		}
	case "messages_deny":
		var obj object.MessageDenyObject
		if err := json.Unmarshal(e.Object, &obj); err != nil {
			return err
		}

		for _, f := range funcList.MessageDeny {
			f(obj, e.GroupID)
		}
	case "message_typing_state": // На основе ответа
		var obj object.MessageTypingStateObject
		if err := json.Unmarshal(e.Object, &obj); err != nil {
			return err
		}

		for _, f := range funcList.MessageTypingState {
			f(obj, e.GroupID)
		}
	case "photo_new":
		var obj object.PhotoNewObject
		if err := json.Unmarshal(e.Object, &obj); err != nil {
			return err
		}

		for _, f := range funcList.PhotoNew {
			f(obj, e.GroupID)
		}
	case "photo_comment_new":
		var obj object.PhotoCommentNewObject
		if err := json.Unmarshal(e.Object, &obj); err != nil {
			return err
		}

		for _, f := range funcList.PhotoCommentNew {
			f(obj, e.GroupID)
		}
	case "photo_comment_edit":
		var obj object.PhotoCommentEditObject
		if err := json.Unmarshal(e.Object, &obj); err != nil {
			return err
		}

		for _, f := range funcList.PhotoCommentEdit {
			f(obj, e.GroupID)
		}
	case "photo_comment_restore":
		var obj object.PhotoCommentRestoreObject
		if err := json.Unmarshal(e.Object, &obj); err != nil {
			return err
		}

		for _, f := range funcList.PhotoCommentRestore {
			f(obj, e.GroupID)
		}
	case "photo_comment_delete":
		var obj object.PhotoCommentDeleteObject
		if err := json.Unmarshal(e.Object, &obj); err != nil {
			return err
		}

		for _, f := range funcList.PhotoCommentDelete {
			f(obj, e.GroupID)
		}
	case "audio_new":
		var obj object.AudioNewObject
		if err := json.Unmarshal(e.Object, &obj); err != nil {
			return err
		}

		for _, f := range funcList.AudioNew {
			f(obj, e.GroupID)
		}
	case "video_new":
		var obj object.VideoNewObject
		if err := json.Unmarshal(e.Object, &obj); err != nil {
			return err
		}

		for _, f := range funcList.VideoNew {
			f(obj, e.GroupID)
		}
	case "video_comment_new":
		var obj object.VideoCommentNewObject
		if err := json.Unmarshal(e.Object, &obj); err != nil {
			return err
		}

		for _, f := range funcList.VideoCommentNew {
			f(obj, e.GroupID)
		}
	case "video_comment_edit":
		var obj object.VideoCommentEditObject
		if err := json.Unmarshal(e.Object, &obj); err != nil {
			return err
		}

		for _, f := range funcList.VideoCommentEdit {
			f(obj, e.GroupID)
		}
	case "video_comment_restore":
		var obj object.VideoCommentRestoreObject
		if err := json.Unmarshal(e.Object, &obj); err != nil {
			return err
		}

		for _, f := range funcList.VideoCommentRestore {
			f(obj, e.GroupID)
		}
	case "video_comment_delete":
		var obj object.VideoCommentDeleteObject
		if err := json.Unmarshal(e.Object, &obj); err != nil {
			return err
		}

		for _, f := range funcList.VideoCommentDelete {
			f(obj, e.GroupID)
		}
	case "wall_post_new":
		var obj object.WallPostNewObject
		if err := json.Unmarshal(e.Object, &obj); err != nil {
			return err
		}

		for _, f := range funcList.WallPostNew {
			f(obj, e.GroupID)
		}
	case "wall_repost":
		var obj object.WallRepostObject
		if err := json.Unmarshal(e.Object, &obj); err != nil {
			return err
		}

		for _, f := range funcList.WallRepost {
			f(obj, e.GroupID)
		}
	case "wall_reply_new":
		var obj object.WallReplyNewObject
		if err := json.Unmarshal(e.Object, &obj); err != nil {
			return err
		}

		for _, f := range funcList.WallReplyNew {
			f(obj, e.GroupID)
		}
	case "wall_reply_edit":
		var obj object.WallReplyEditObject
		if err := json.Unmarshal(e.Object, &obj); err != nil {
			return err
		}

		for _, f := range funcList.WallReplyEdit {
			f(obj, e.GroupID)
		}
	case "wall_reply_restore":
		var obj object.WallReplyRestoreObject
		if err := json.Unmarshal(e.Object, &obj); err != nil {
			return err
		}

		for _, f := range funcList.WallReplyRestore {
			f(obj, e.GroupID)
		}
	case "wall_reply_delete":
		var obj object.WallReplyDeleteObject
		if err := json.Unmarshal(e.Object, &obj); err != nil {
			return err
		}

		for _, f := range funcList.WallReplyDelete {
			f(obj, e.GroupID)
		}
	case "board_post_new":
		var obj object.BoardPostNewObject
		if err := json.Unmarshal(e.Object, &obj); err != nil {
			return err
		}

		for _, f := range funcList.BoardPostNew {
			f(obj, e.GroupID)
		}
	case "board_post_edit":
		var obj object.BoardPostEditObject
		if err := json.Unmarshal(e.Object, &obj); err != nil {
			return err
		}

		for _, f := range funcList.BoardPostEdit {
			f(obj, e.GroupID)
		}
	case "board_post_restore":
		var obj object.BoardPostRestoreObject
		if err := json.Unmarshal(e.Object, &obj); err != nil {
			return err
		}

		for _, f := range funcList.BoardPostRestore {
			f(obj, e.GroupID)
		}
	case "board_post_delete":
		var obj object.BoardPostDeleteObject
		if err := json.Unmarshal(e.Object, &obj); err != nil {
			return err
		}

		for _, f := range funcList.BoardPostDelete {
			f(obj, e.GroupID)
		}
	case "market_comment_new":
		var obj object.MarketCommentNewObject
		if err := json.Unmarshal(e.Object, &obj); err != nil {
			return err
		}

		for _, f := range funcList.MarketCommentNew {
			f(obj, e.GroupID)
		}
	case "market_comment_edit":
		var obj object.MarketCommentEditObject
		if err := json.Unmarshal(e.Object, &obj); err != nil {
			return err
		}

		for _, f := range funcList.MarketCommentEdit {
			f(obj, e.GroupID)
		}
	case "market_comment_restore":
		var obj object.MarketCommentRestoreObject
		if err := json.Unmarshal(e.Object, &obj); err != nil {
			return err
		}

		for _, f := range funcList.MarketCommentRestore {
			f(obj, e.GroupID)
		}
	case "market_comment_delete":
		var obj object.MarketCommentDeleteObject
		if err := json.Unmarshal(e.Object, &obj); err != nil {
			return err
		}

		for _, f := range funcList.MarketCommentDelete {
			f(obj, e.GroupID)
		}
	case "group_leave":
		var obj object.GroupLeaveObject
		if err := json.Unmarshal(e.Object, &obj); err != nil {
			return err
		}

		for _, f := range funcList.GroupLeave {
			f(obj, e.GroupID)
		}
	case "group_join":
		var obj object.GroupJoinObject
		if err := json.Unmarshal(e.Object, &obj); err != nil {
			return err
		}

		for _, f := range funcList.GroupJoin {
			f(obj, e.GroupID)
		}
	case "user_block":
		var obj object.UserBlockObject
		if err := json.Unmarshal(e.Object, &obj); err != nil {
			return err
		}

		for _, f := range funcList.UserBlock {
			f(obj, e.GroupID)
		}
	case "user_unblock":
		var obj object.UserUnblockObject
		if err := json.Unmarshal(e.Object, &obj); err != nil {
			return err
		}

		for _, f := range funcList.UserUnblock {
			f(obj, e.GroupID)
		}
	case "poll_vote_new":
		var obj object.PollVoteNewObject
		if err := json.Unmarshal(e.Object, &obj); err != nil {
			return err
		}

		for _, f := range funcList.PollVoteNew {
			f(obj, e.GroupID)
		}
	case "group_officers_edit":
		var obj object.GroupOfficersEditObject
		if err := json.Unmarshal(e.Object, &obj); err != nil {
			return err
		}

		for _, f := range funcList.GroupOfficersEdit {
			f(obj, e.GroupID)
		}
	case "group_change_settings":
		var obj object.GroupChangeSettingsObject
		if err := json.Unmarshal(e.Object, &obj); err != nil {
			return err
		}

		for _, f := range funcList.GroupChangeSettings {
			f(obj, e.GroupID)
		}
	case "group_change_photo":
		var obj object.GroupChangePhotoObject
		if err := json.Unmarshal(e.Object, &obj); err != nil {
			return err
		}

		for _, f := range funcList.GroupChangePhoto {
			f(obj, e.GroupID)
		}
	case "vkpay_transaction":
		var obj object.VkpayTransactionObject
		if err := json.Unmarshal(e.Object, &obj); err != nil {
			return err
		}

		for _, f := range funcList.VkpayTransaction {
			f(obj, e.GroupID)
		}
	}
	return err
}
