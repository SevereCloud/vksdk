package callback // import "github.com/severecloud/vksdk/callback"

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/severecloud/vksdk/handler"
	"github.com/severecloud/vksdk/object"
)

// Callback struct SecretKeys [GroupID]SecretKey
type Callback struct {
	ConfirmationKeys map[int]string
	ConfirmationKey  string
	SecretKeys       map[int]string
	SecretKey        string
	FuncList         handler.FuncList
}

// HandleFunc handler
func (cb Callback) HandleFunc(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)

	var e object.GroupEvent
	if err := decoder.Decode(&e); err != nil {
		return
	}

	if e.Type == "confirmation" {
		if cb.ConfirmationKeys[e.GroupID] != "" {
			fmt.Fprintf(w, cb.ConfirmationKeys[e.GroupID])
		} else {
			fmt.Fprintf(w, cb.ConfirmationKey)
		}
		return
	}

	if cb.SecretKeys[e.GroupID] != "" || cb.SecretKey != "" {
		if e.Secret != cb.SecretKeys[e.GroupID] && e.Secret != cb.SecretKey {
			fmt.Fprintf(w, "bad secret")
			return
		}
	}
	cb.FuncList.Handler(e)
	fmt.Fprintf(w, "ok")
}

// TODO: support fasthttp

// MessageNew handler
func (cb *Callback) MessageNew(f object.MessageNewFunc) {
	cb.FuncList.MessageNew = append(cb.FuncList.MessageNew, f)
}

// MessageReply handler
func (cb *Callback) MessageReply(f object.MessageReplyFunc) {
	cb.FuncList.MessageReply = append(cb.FuncList.MessageReply, f)
}

// MessageEdit handler
func (cb *Callback) MessageEdit(f object.MessageEditFunc) {
	cb.FuncList.MessageEdit = append(cb.FuncList.MessageEdit, f)
}

// MessageAllow handler
func (cb *Callback) MessageAllow(f object.MessageAllowFunc) {
	cb.FuncList.MessageAllow = append(cb.FuncList.MessageAllow, f)
}

// MessageDeny handler
func (cb *Callback) MessageDeny(f object.MessageDenyFunc) {
	cb.FuncList.MessageDeny = append(cb.FuncList.MessageDeny, f)
}

// MessageTypingState handler
func (cb *Callback) MessageTypingState(f object.MessageTypingStateFunc) {
	cb.FuncList.MessageTypingState = append(cb.FuncList.MessageTypingState, f)
}

// PhotoNew handler
func (cb *Callback) PhotoNew(f object.PhotoNewFunc) {
	cb.FuncList.PhotoNew = append(cb.FuncList.PhotoNew, f)
}

// PhotoCommentNew handler
func (cb *Callback) PhotoCommentNew(f object.PhotoCommentNewFunc) {
	cb.FuncList.PhotoCommentNew = append(cb.FuncList.PhotoCommentNew, f)
}

// PhotoCommentEdit handler
func (cb *Callback) PhotoCommentEdit(f object.PhotoCommentEditFunc) {
	cb.FuncList.PhotoCommentEdit = append(cb.FuncList.PhotoCommentEdit, f)
}

// PhotoCommentRestore handler
func (cb *Callback) PhotoCommentRestore(f object.PhotoCommentRestoreFunc) {
	cb.FuncList.PhotoCommentRestore = append(cb.FuncList.PhotoCommentRestore, f)
}

// PhotoCommentDelete handler
func (cb *Callback) PhotoCommentDelete(f object.PhotoCommentDeleteFunc) {
	cb.FuncList.PhotoCommentDelete = append(cb.FuncList.PhotoCommentDelete, f)
}

// AudioNew handler
func (cb *Callback) AudioNew(f object.AudioNewFunc) {
	cb.FuncList.AudioNew = append(cb.FuncList.AudioNew, f)
}

// VideoNew handler
func (cb *Callback) VideoNew(f object.VideoNewFunc) {
	cb.FuncList.VideoNew = append(cb.FuncList.VideoNew, f)
}

// VideoCommentNew handler
func (cb *Callback) VideoCommentNew(f object.VideoCommentNewFunc) {
	cb.FuncList.VideoCommentNew = append(cb.FuncList.VideoCommentNew, f)
}

// VideoCommentEdit handler
func (cb *Callback) VideoCommentEdit(f object.VideoCommentEditFunc) {
	cb.FuncList.VideoCommentEdit = append(cb.FuncList.VideoCommentEdit, f)
}

// VideoCommentRestore handler
func (cb *Callback) VideoCommentRestore(f object.VideoCommentRestoreFunc) {
	cb.FuncList.VideoCommentRestore = append(cb.FuncList.VideoCommentRestore, f)
}

// VideoCommentDelete handler
func (cb *Callback) VideoCommentDelete(f object.VideoCommentDeleteFunc) {
	cb.FuncList.VideoCommentDelete = append(cb.FuncList.VideoCommentDelete, f)
}

// WallPostNew handler
func (cb *Callback) WallPostNew(f object.WallPostNewFunc) {
	cb.FuncList.WallPostNew = append(cb.FuncList.WallPostNew, f)
}

// WallRepost handler
func (cb *Callback) WallRepost(f object.WallRepostFunc) {
	cb.FuncList.WallRepost = append(cb.FuncList.WallRepost, f)
}

// WallReplyNew handler
func (cb *Callback) WallReplyNew(f object.WallReplyNewFunc) {
	cb.FuncList.WallReplyNew = append(cb.FuncList.WallReplyNew, f)
}

// WallReplyEdit handler
func (cb *Callback) WallReplyEdit(f object.WallReplyEditFunc) {
	cb.FuncList.WallReplyEdit = append(cb.FuncList.WallReplyEdit, f)
}

// WallReplyRestore handler
func (cb *Callback) WallReplyRestore(f object.WallReplyRestoreFunc) {
	cb.FuncList.WallReplyRestore = append(cb.FuncList.WallReplyRestore, f)
}

// WallReplyDelete handler
func (cb *Callback) WallReplyDelete(f object.WallReplyDeleteFunc) {
	cb.FuncList.WallReplyDelete = append(cb.FuncList.WallReplyDelete, f)
}

// BoardPostNew handler
func (cb *Callback) BoardPostNew(f object.BoardPostNewFunc) {
	cb.FuncList.BoardPostNew = append(cb.FuncList.BoardPostNew, f)
}

// BoardPostEdit handler
func (cb *Callback) BoardPostEdit(f object.BoardPostEditFunc) {
	cb.FuncList.BoardPostEdit = append(cb.FuncList.BoardPostEdit, f)
}

// BoardPostRestore handler
func (cb *Callback) BoardPostRestore(f object.BoardPostRestoreFunc) {
	cb.FuncList.BoardPostRestore = append(cb.FuncList.BoardPostRestore, f)
}

// BoardPostDelete handler
func (cb *Callback) BoardPostDelete(f object.BoardPostDeleteFunc) {
	cb.FuncList.BoardPostDelete = append(cb.FuncList.BoardPostDelete, f)
}

// MarketCommentNew handler
func (cb *Callback) MarketCommentNew(f object.MarketCommentNewFunc) {
	cb.FuncList.MarketCommentNew = append(cb.FuncList.MarketCommentNew, f)
}

// MarketCommentEdit handler
func (cb *Callback) MarketCommentEdit(f object.MarketCommentEditFunc) {
	cb.FuncList.MarketCommentEdit = append(cb.FuncList.MarketCommentEdit, f)
}

// MarketCommentRestore handler
func (cb *Callback) MarketCommentRestore(f object.MarketCommentRestoreFunc) {
	cb.FuncList.MarketCommentRestore = append(cb.FuncList.MarketCommentRestore, f)
}

// MarketCommentDelete handler
func (cb *Callback) MarketCommentDelete(f object.MarketCommentDeleteFunc) {
	cb.FuncList.MarketCommentDelete = append(cb.FuncList.MarketCommentDelete, f)
}

// GroupLeave handler
func (cb *Callback) GroupLeave(f object.GroupLeaveFunc) {
	cb.FuncList.GroupLeave = append(cb.FuncList.GroupLeave, f)
}

// GroupJoin handler
func (cb *Callback) GroupJoin(f object.GroupJoinFunc) {
	cb.FuncList.GroupJoin = append(cb.FuncList.GroupJoin, f)
}

// UserBlock handler
func (cb *Callback) UserBlock(f object.UserBlockFunc) {
	cb.FuncList.UserBlock = append(cb.FuncList.UserBlock, f)
}

// UserUnblock handler
func (cb *Callback) UserUnblock(f object.UserUnblockFunc) {
	cb.FuncList.UserUnblock = append(cb.FuncList.UserUnblock, f)
}

// PollVoteNew handler
func (cb *Callback) PollVoteNew(f object.PollVoteNewFunc) {
	cb.FuncList.PollVoteNew = append(cb.FuncList.PollVoteNew, f)
}

// GroupOfficersEdit handler
func (cb *Callback) GroupOfficersEdit(f object.GroupOfficersEditFunc) {
	cb.FuncList.GroupOfficersEdit = append(cb.FuncList.GroupOfficersEdit, f)
}

// GroupChangeSettings handler
func (cb *Callback) GroupChangeSettings(f object.GroupChangeSettingsFunc) {
	cb.FuncList.GroupChangeSettings = append(cb.FuncList.GroupChangeSettings, f)
}

// GroupChangePhoto handler
func (cb *Callback) GroupChangePhoto(f object.GroupChangePhotoFunc) {
	cb.FuncList.GroupChangePhoto = append(cb.FuncList.GroupChangePhoto, f)
}

// VkpayTransaction handler
func (cb *Callback) VkpayTransaction(f object.VkpayTransactionFunc) {
	cb.FuncList.VkpayTransaction = append(cb.FuncList.VkpayTransaction, f)
}

// TODO: next version lead_forms_new handler
// TODO: next version like_add handler
// TODO: next version like_remove handler
