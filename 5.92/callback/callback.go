package callback // import "github.com/SevereCloud/vksdk/5.92/callback"

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/SevereCloud/vksdk/5.92/handler"
	"github.com/SevereCloud/vksdk/5.92/object"
)

// Callback struct SecretKeys [GroupID]SecretKey
type Callback struct {
	ConfirmationKeys map[int]string
	ConfirmationKey  string
	SecretKeys       map[int]string
	SecretKey        string
	funcList         handler.FuncList
}

// HandleFunc handler
func (cb Callback) HandleFunc(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)

	var e object.GroupEvent
	if err := decoder.Decode(&e); err != nil {
		log.Printf("Callback.HandleFunc: %v", err)
		// fmt.Fprintf(w, "%v", err)
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
	if err := cb.funcList.Handler(e); err != nil {
		log.Printf("Callback.HandleFunc: %v", err)
		// fmt.Fprintf(w, "%v", err)
		return
	}
	fmt.Fprintf(w, "ok")
}

// MessageNew handler
func (cb *Callback) MessageNew(f object.MessageNewFunc) {
	cb.funcList.MessageNew = append(cb.funcList.MessageNew, f)
}

// MessageReply handler
func (cb *Callback) MessageReply(f object.MessageReplyFunc) {
	cb.funcList.MessageReply = append(cb.funcList.MessageReply, f)
}

// MessageEdit handler
func (cb *Callback) MessageEdit(f object.MessageEditFunc) {
	cb.funcList.MessageEdit = append(cb.funcList.MessageEdit, f)
}

// MessageAllow handler
func (cb *Callback) MessageAllow(f object.MessageAllowFunc) {
	cb.funcList.MessageAllow = append(cb.funcList.MessageAllow, f)
}

// MessageDeny handler
func (cb *Callback) MessageDeny(f object.MessageDenyFunc) {
	cb.funcList.MessageDeny = append(cb.funcList.MessageDeny, f)
}

// MessageTypingState handler
func (cb *Callback) MessageTypingState(f object.MessageTypingStateFunc) {
	cb.funcList.MessageTypingState = append(cb.funcList.MessageTypingState, f)
}

// PhotoNew handler
func (cb *Callback) PhotoNew(f object.PhotoNewFunc) {
	cb.funcList.PhotoNew = append(cb.funcList.PhotoNew, f)
}

// PhotoCommentNew handler
func (cb *Callback) PhotoCommentNew(f object.PhotoCommentNewFunc) {
	cb.funcList.PhotoCommentNew = append(cb.funcList.PhotoCommentNew, f)
}

// PhotoCommentEdit handler
func (cb *Callback) PhotoCommentEdit(f object.PhotoCommentEditFunc) {
	cb.funcList.PhotoCommentEdit = append(cb.funcList.PhotoCommentEdit, f)
}

// PhotoCommentRestore handler
func (cb *Callback) PhotoCommentRestore(f object.PhotoCommentRestoreFunc) {
	cb.funcList.PhotoCommentRestore = append(cb.funcList.PhotoCommentRestore, f)
}

// PhotoCommentDelete handler
func (cb *Callback) PhotoCommentDelete(f object.PhotoCommentDeleteFunc) {
	cb.funcList.PhotoCommentDelete = append(cb.funcList.PhotoCommentDelete, f)
}

// AudioNew handler
func (cb *Callback) AudioNew(f object.AudioNewFunc) {
	cb.funcList.AudioNew = append(cb.funcList.AudioNew, f)
}

// VideoNew handler
func (cb *Callback) VideoNew(f object.VideoNewFunc) {
	cb.funcList.VideoNew = append(cb.funcList.VideoNew, f)
}

// VideoCommentNew handler
func (cb *Callback) VideoCommentNew(f object.VideoCommentNewFunc) {
	cb.funcList.VideoCommentNew = append(cb.funcList.VideoCommentNew, f)
}

// VideoCommentEdit handler
func (cb *Callback) VideoCommentEdit(f object.VideoCommentEditFunc) {
	cb.funcList.VideoCommentEdit = append(cb.funcList.VideoCommentEdit, f)
}

// VideoCommentRestore handler
func (cb *Callback) VideoCommentRestore(f object.VideoCommentRestoreFunc) {
	cb.funcList.VideoCommentRestore = append(cb.funcList.VideoCommentRestore, f)
}

// VideoCommentDelete handler
func (cb *Callback) VideoCommentDelete(f object.VideoCommentDeleteFunc) {
	cb.funcList.VideoCommentDelete = append(cb.funcList.VideoCommentDelete, f)
}

// WallPostNew handler
func (cb *Callback) WallPostNew(f object.WallPostNewFunc) {
	cb.funcList.WallPostNew = append(cb.funcList.WallPostNew, f)
}

// WallRepost handler
func (cb *Callback) WallRepost(f object.WallRepostFunc) {
	cb.funcList.WallRepost = append(cb.funcList.WallRepost, f)
}

// WallReplyNew handler
func (cb *Callback) WallReplyNew(f object.WallReplyNewFunc) {
	cb.funcList.WallReplyNew = append(cb.funcList.WallReplyNew, f)
}

// WallReplyEdit handler
func (cb *Callback) WallReplyEdit(f object.WallReplyEditFunc) {
	cb.funcList.WallReplyEdit = append(cb.funcList.WallReplyEdit, f)
}

// WallReplyRestore handler
func (cb *Callback) WallReplyRestore(f object.WallReplyRestoreFunc) {
	cb.funcList.WallReplyRestore = append(cb.funcList.WallReplyRestore, f)
}

// WallReplyDelete handler
func (cb *Callback) WallReplyDelete(f object.WallReplyDeleteFunc) {
	cb.funcList.WallReplyDelete = append(cb.funcList.WallReplyDelete, f)
}

// BoardPostNew handler
func (cb *Callback) BoardPostNew(f object.BoardPostNewFunc) {
	cb.funcList.BoardPostNew = append(cb.funcList.BoardPostNew, f)
}

// BoardPostEdit handler
func (cb *Callback) BoardPostEdit(f object.BoardPostEditFunc) {
	cb.funcList.BoardPostEdit = append(cb.funcList.BoardPostEdit, f)
}

// BoardPostRestore handler
func (cb *Callback) BoardPostRestore(f object.BoardPostRestoreFunc) {
	cb.funcList.BoardPostRestore = append(cb.funcList.BoardPostRestore, f)
}

// BoardPostDelete handler
func (cb *Callback) BoardPostDelete(f object.BoardPostDeleteFunc) {
	cb.funcList.BoardPostDelete = append(cb.funcList.BoardPostDelete, f)
}

// MarketCommentNew handler
func (cb *Callback) MarketCommentNew(f object.MarketCommentNewFunc) {
	cb.funcList.MarketCommentNew = append(cb.funcList.MarketCommentNew, f)
}

// MarketCommentEdit handler
func (cb *Callback) MarketCommentEdit(f object.MarketCommentEditFunc) {
	cb.funcList.MarketCommentEdit = append(cb.funcList.MarketCommentEdit, f)
}

// MarketCommentRestore handler
func (cb *Callback) MarketCommentRestore(f object.MarketCommentRestoreFunc) {
	cb.funcList.MarketCommentRestore = append(cb.funcList.MarketCommentRestore, f)
}

// MarketCommentDelete handler
func (cb *Callback) MarketCommentDelete(f object.MarketCommentDeleteFunc) {
	cb.funcList.MarketCommentDelete = append(cb.funcList.MarketCommentDelete, f)
}

// GroupLeave handler
func (cb *Callback) GroupLeave(f object.GroupLeaveFunc) {
	cb.funcList.GroupLeave = append(cb.funcList.GroupLeave, f)
}

// GroupJoin handler
func (cb *Callback) GroupJoin(f object.GroupJoinFunc) {
	cb.funcList.GroupJoin = append(cb.funcList.GroupJoin, f)
}

// UserBlock handler
func (cb *Callback) UserBlock(f object.UserBlockFunc) {
	cb.funcList.UserBlock = append(cb.funcList.UserBlock, f)
}

// UserUnblock handler
func (cb *Callback) UserUnblock(f object.UserUnblockFunc) {
	cb.funcList.UserUnblock = append(cb.funcList.UserUnblock, f)
}

// PollVoteNew handler
func (cb *Callback) PollVoteNew(f object.PollVoteNewFunc) {
	cb.funcList.PollVoteNew = append(cb.funcList.PollVoteNew, f)
}

// GroupOfficersEdit handler
func (cb *Callback) GroupOfficersEdit(f object.GroupOfficersEditFunc) {
	cb.funcList.GroupOfficersEdit = append(cb.funcList.GroupOfficersEdit, f)
}

// GroupChangeSettings handler
func (cb *Callback) GroupChangeSettings(f object.GroupChangeSettingsFunc) {
	cb.funcList.GroupChangeSettings = append(cb.funcList.GroupChangeSettings, f)
}

// GroupChangePhoto handler
func (cb *Callback) GroupChangePhoto(f object.GroupChangePhotoFunc) {
	cb.funcList.GroupChangePhoto = append(cb.funcList.GroupChangePhoto, f)
}

// VkpayTransaction handler
func (cb *Callback) VkpayTransaction(f object.VkpayTransactionFunc) {
	cb.funcList.VkpayTransaction = append(cb.funcList.VkpayTransaction, f)
}

// LeadFormsNew handler
func (cb *Callback) LeadFormsNew(f object.LeadFormsNewFunc) {
	cb.funcList.LeadFormsNew = append(cb.funcList.LeadFormsNew, f)
}

// TODO: like_add like_remove app_payload message_read
