package longpoll // import "github.com/SevereCloud/vksdk/5.92/longpoll-bot"

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"sync/atomic"

	"github.com/SevereCloud/vksdk/5.92/object"

	"github.com/SevereCloud/vksdk/5.92/api"
	"github.com/SevereCloud/vksdk/5.92/handler"
)

// Longpoll struct
type Longpoll struct {
	GroupID int
	Server  string
	Key     string
	Ts      string
	Wait    int
	VK      *api.VK
	Client  *http.Client

	funcList             handler.FuncList
	funcFullResponseList []func(object.LongpollBotResponse)
	inShutdown           int32
}

// Init Longpoll
func Init(vk *api.VK, groupID int) (lp Longpoll, err error) {
	lp.VK = vk
	lp.GroupID = groupID
	lp.Wait = 25
	lp.Client = &http.Client{}
	err = lp.updateServer(true)
	return
}

func (lp *Longpoll) updateServer(updateTs bool) error {
	params := map[string]string{
		"group_id": strconv.Itoa(lp.GroupID),
	}
	serverSetting, vkErr := lp.VK.GroupsGetLongPollServer(params)
	if vkErr.Code != 0 {
		return fmt.Errorf(vkErr.Message)
	}

	lp.Key = serverSetting.Key
	lp.Server = serverSetting.Server
	if updateTs {
		lp.Ts = serverSetting.Ts
	}
	return nil
}

func (lp *Longpoll) check() ([]object.GroupEvent, error) {
	var response object.LongpollBotResponse
	var err error

	u := fmt.Sprintf("%s?act=a_check&key=%s&ts=%s&wait=%d", lp.Server, lp.Key, lp.Ts, lp.Wait)

	resp, err := lp.Client.Get(u)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(body, &response)
	if err != nil {
		return nil, err
	}

	err = lp.checkResponse(response)
	return response.Updates, err
}

func (lp *Longpoll) checkResponse(response object.LongpollBotResponse) (err error) {
	switch response.Failed {
	case 0:
		lp.Ts = response.Ts
	case 1:
		lp.Ts = response.Ts
	case 2:
		err = lp.updateServer(false)
	case 3:
		err = lp.updateServer(true)
	default:
		err = fmt.Errorf(`"failed":%d`, response.Failed)
	}
	return
}

// Run handler
func (lp *Longpoll) Run() error {
	atomic.StoreInt32(&lp.inShutdown, 0)
	for atomic.LoadInt32(&lp.inShutdown) == 0 {
		events, err := lp.check()
		if err != nil {
			return err
		}
		for _, event := range events {
			err = lp.funcList.Handler(event)
			if err != nil {
				return err
			}
		}
	}
	return nil
}

// Shutdown gracefully shuts down the longpoll without interrupting any active connections.
func (lp *Longpoll) Shutdown() {
	atomic.StoreInt32(&lp.inShutdown, 1)
}

// MessageNew handler
func (lp *Longpoll) MessageNew(f object.MessageNewFunc) {
	lp.funcList.MessageNew = append(lp.funcList.MessageNew, f)
}

// MessageReply handler
func (lp *Longpoll) MessageReply(f object.MessageReplyFunc) {
	lp.funcList.MessageReply = append(lp.funcList.MessageReply, f)
}

// MessageEdit handler
func (lp *Longpoll) MessageEdit(f object.MessageEditFunc) {
	lp.funcList.MessageEdit = append(lp.funcList.MessageEdit, f)
}

// MessageAllow handler
func (lp *Longpoll) MessageAllow(f object.MessageAllowFunc) {
	lp.funcList.MessageAllow = append(lp.funcList.MessageAllow, f)
}

// MessageDeny handler
func (lp *Longpoll) MessageDeny(f object.MessageDenyFunc) {
	lp.funcList.MessageDeny = append(lp.funcList.MessageDeny, f)
}

// MessageTypingState handler
func (lp *Longpoll) MessageTypingState(f object.MessageTypingStateFunc) {
	lp.funcList.MessageTypingState = append(lp.funcList.MessageTypingState, f)
}

// PhotoNew handler
func (lp *Longpoll) PhotoNew(f object.PhotoNewFunc) {
	lp.funcList.PhotoNew = append(lp.funcList.PhotoNew, f)
}

// PhotoCommentNew handler
func (lp *Longpoll) PhotoCommentNew(f object.PhotoCommentNewFunc) {
	lp.funcList.PhotoCommentNew = append(lp.funcList.PhotoCommentNew, f)
}

// PhotoCommentEdit handler
func (lp *Longpoll) PhotoCommentEdit(f object.PhotoCommentEditFunc) {
	lp.funcList.PhotoCommentEdit = append(lp.funcList.PhotoCommentEdit, f)
}

// PhotoCommentRestore handler
func (lp *Longpoll) PhotoCommentRestore(f object.PhotoCommentRestoreFunc) {
	lp.funcList.PhotoCommentRestore = append(lp.funcList.PhotoCommentRestore, f)
}

// PhotoCommentDelete handler
func (lp *Longpoll) PhotoCommentDelete(f object.PhotoCommentDeleteFunc) {
	lp.funcList.PhotoCommentDelete = append(lp.funcList.PhotoCommentDelete, f)
}

// AudioNew handler
func (lp *Longpoll) AudioNew(f object.AudioNewFunc) {
	lp.funcList.AudioNew = append(lp.funcList.AudioNew, f)
}

// VideoNew handler
func (lp *Longpoll) VideoNew(f object.VideoNewFunc) {
	lp.funcList.VideoNew = append(lp.funcList.VideoNew, f)
}

// VideoCommentNew handler
func (lp *Longpoll) VideoCommentNew(f object.VideoCommentNewFunc) {
	lp.funcList.VideoCommentNew = append(lp.funcList.VideoCommentNew, f)
}

// VideoCommentEdit handler
func (lp *Longpoll) VideoCommentEdit(f object.VideoCommentEditFunc) {
	lp.funcList.VideoCommentEdit = append(lp.funcList.VideoCommentEdit, f)
}

// VideoCommentRestore handler
func (lp *Longpoll) VideoCommentRestore(f object.VideoCommentRestoreFunc) {
	lp.funcList.VideoCommentRestore = append(lp.funcList.VideoCommentRestore, f)
}

// VideoCommentDelete handler
func (lp *Longpoll) VideoCommentDelete(f object.VideoCommentDeleteFunc) {
	lp.funcList.VideoCommentDelete = append(lp.funcList.VideoCommentDelete, f)
}

// WallPostNew handler
func (lp *Longpoll) WallPostNew(f object.WallPostNewFunc) {
	lp.funcList.WallPostNew = append(lp.funcList.WallPostNew, f)
}

// WallRepost handler
func (lp *Longpoll) WallRepost(f object.WallRepostFunc) {
	lp.funcList.WallRepost = append(lp.funcList.WallRepost, f)
}

// WallReplyNew handler
func (lp *Longpoll) WallReplyNew(f object.WallReplyNewFunc) {
	lp.funcList.WallReplyNew = append(lp.funcList.WallReplyNew, f)
}

// WallReplyEdit handler
func (lp *Longpoll) WallReplyEdit(f object.WallReplyEditFunc) {
	lp.funcList.WallReplyEdit = append(lp.funcList.WallReplyEdit, f)
}

// WallReplyRestore handler
func (lp *Longpoll) WallReplyRestore(f object.WallReplyRestoreFunc) {
	lp.funcList.WallReplyRestore = append(lp.funcList.WallReplyRestore, f)
}

// WallReplyDelete handler
func (lp *Longpoll) WallReplyDelete(f object.WallReplyDeleteFunc) {
	lp.funcList.WallReplyDelete = append(lp.funcList.WallReplyDelete, f)
}

// BoardPostNew handler
func (lp *Longpoll) BoardPostNew(f object.BoardPostNewFunc) {
	lp.funcList.BoardPostNew = append(lp.funcList.BoardPostNew, f)
}

// BoardPostEdit handler
func (lp *Longpoll) BoardPostEdit(f object.BoardPostEditFunc) {
	lp.funcList.BoardPostEdit = append(lp.funcList.BoardPostEdit, f)
}

// BoardPostRestore handler
func (lp *Longpoll) BoardPostRestore(f object.BoardPostRestoreFunc) {
	lp.funcList.BoardPostRestore = append(lp.funcList.BoardPostRestore, f)
}

// BoardPostDelete handler
func (lp *Longpoll) BoardPostDelete(f object.BoardPostDeleteFunc) {
	lp.funcList.BoardPostDelete = append(lp.funcList.BoardPostDelete, f)
}

// MarketCommentNew handler
func (lp *Longpoll) MarketCommentNew(f object.MarketCommentNewFunc) {
	lp.funcList.MarketCommentNew = append(lp.funcList.MarketCommentNew, f)
}

// MarketCommentEdit handler
func (lp *Longpoll) MarketCommentEdit(f object.MarketCommentEditFunc) {
	lp.funcList.MarketCommentEdit = append(lp.funcList.MarketCommentEdit, f)
}

// MarketCommentRestore handler
func (lp *Longpoll) MarketCommentRestore(f object.MarketCommentRestoreFunc) {
	lp.funcList.MarketCommentRestore = append(lp.funcList.MarketCommentRestore, f)
}

// MarketCommentDelete handler
func (lp *Longpoll) MarketCommentDelete(f object.MarketCommentDeleteFunc) {
	lp.funcList.MarketCommentDelete = append(lp.funcList.MarketCommentDelete, f)
}

// GroupLeave handler
func (lp *Longpoll) GroupLeave(f object.GroupLeaveFunc) {
	lp.funcList.GroupLeave = append(lp.funcList.GroupLeave, f)
}

// GroupJoin handler
func (lp *Longpoll) GroupJoin(f object.GroupJoinFunc) {
	lp.funcList.GroupJoin = append(lp.funcList.GroupJoin, f)
}

// UserBlock handler
func (lp *Longpoll) UserBlock(f object.UserBlockFunc) {
	lp.funcList.UserBlock = append(lp.funcList.UserBlock, f)
}

// UserUnblock handler
func (lp *Longpoll) UserUnblock(f object.UserUnblockFunc) {
	lp.funcList.UserUnblock = append(lp.funcList.UserUnblock, f)
}

// PollVoteNew handler
func (lp *Longpoll) PollVoteNew(f object.PollVoteNewFunc) {
	lp.funcList.PollVoteNew = append(lp.funcList.PollVoteNew, f)
}

// GroupOfficersEdit handler
func (lp *Longpoll) GroupOfficersEdit(f object.GroupOfficersEditFunc) {
	lp.funcList.GroupOfficersEdit = append(lp.funcList.GroupOfficersEdit, f)
}

// GroupChangeSettings handler
func (lp *Longpoll) GroupChangeSettings(f object.GroupChangeSettingsFunc) {
	lp.funcList.GroupChangeSettings = append(lp.funcList.GroupChangeSettings, f)
}

// GroupChangePhoto handler
func (lp *Longpoll) GroupChangePhoto(f object.GroupChangePhotoFunc) {
	lp.funcList.GroupChangePhoto = append(lp.funcList.GroupChangePhoto, f)
}

// VkpayTransaction handler
func (lp *Longpoll) VkpayTransaction(f object.VkpayTransactionFunc) {
	lp.funcList.VkpayTransaction = append(lp.funcList.VkpayTransaction, f)
}

// LeadFormsNew handler
func (lp *Longpoll) LeadFormsNew(f object.LeadFormsNewFunc) {
	lp.funcList.LeadFormsNew = append(lp.funcList.LeadFormsNew, f)
}

// TODO: like_add like_remove app_payload message_read

// FullResponse handler
func (lp *Longpoll) FullResponse(f func(object.LongpollBotResponse)) {
	lp.funcFullResponseList = append(lp.funcFullResponseList, f)
}
