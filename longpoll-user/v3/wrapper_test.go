package wrapper_test

import (
	"log"
	"os"
	"testing"

	"github.com/SevereCloud/vksdk/api"
	wrapper "github.com/SevereCloud/vksdk/longpoll-user/v3"
	"github.com/stretchr/testify/assert"

	"github.com/SevereCloud/vksdk/longpoll-user"
)

func needUserToken(t *testing.T) {
	t.Helper()

	if vkUser.AccessToken == "" {
		t.Skip("USER_TOKEN empty")
	}
}

func needChatID(t *testing.T) int {
	t.Helper()
	needUserToken(t)

	if vkChatID == 0 {
		var err error

		vkChatID, err = vkUser.MessagesCreateChat(api.Params{
			"title": "TestChat",
		})
		if err != nil {
			t.Skip("Get chatID", err)
		}
	}

	return vkChatID
}

var vkUser *api.VK         // nolint:gochecknoglobals
var vkUserID, vkChatID int // nolint:gochecknoglobals

func TestMain(m *testing.M) {
	vkUser = api.Init(os.Getenv("USER_TOKEN"))
	vkUser.Limit = 3

	if vkUser.AccessToken != "" {
		user, err := vkUser.UsersGet(api.Params{})
		if err != nil {
			log.Fatalf("USER_TOKEN bad: %v", err)
		}

		vkUserID = user[0].ID
	}

	runTests := m.Run()
	os.Exit(runTests)
}

func TestWrapper_2(t *testing.T) {
	needUserToken(t)
	chatID := needChatID(t)

	exit := make(chan bool, 1)

	lp, err := longpoll.Init(vkUser, 2)
	if err != nil {
		t.Fatalf("lp.Init err: %v", err)
	}

	w := wrapper.NewWrapper(&lp)

	w.OnMessageFlagsChange(func(e wrapper.MessageFlagsChange) {})
	w.OnMessageFlagsSet(func(e wrapper.MessageFlagsSet) {
		assert.NotEmpty(t, e.MessageID)
		assert.NotEmpty(t, e.Mask)
	})
	w.OnMessageFlagsReset(func(e wrapper.MessageFlagsReset) {
		assert.NotEmpty(t, e.MessageID)
		assert.NotEmpty(t, e.Mask)
	})
	w.OnNewMessage(func(e wrapper.NewMessage) {
		assert.NotEmpty(t, e.MessageID)
		assert.NotEmpty(t, e.Flags)
		assert.NotEmpty(t, e.PeerID)
		assert.NotEmpty(t, e.Timestamp)
		assert.NotEmpty(t, e.Text)
	})
	w.OnEditMessage(func(e wrapper.EditMessage) {
		assert.NotEmpty(t, e.MessageID)
		assert.NotEmpty(t, e.Flags)
		assert.NotEmpty(t, e.PeerID)
		assert.NotEmpty(t, e.Timestamp)
		assert.NotEmpty(t, e.NewText)
	})
	w.OnReadInMessages(func(e wrapper.ReadInMessages) {})
	w.OnReadOutMessages(func(e wrapper.ReadOutMessages) {})
	w.OnFriendBecameOnline(func(e wrapper.FriendBecameOnline) {})
	w.OnFriendBecameOffline(func(e wrapper.FriendBecameOffline) {})
	w.OnDialogFlagsReset(func(e wrapper.DialogFlagsReset) {})
	w.OnDialogFlagsReplace(func(e wrapper.DialogFlagsReplace) {})
	w.OnDialogsFlagsSet(func(e wrapper.DialogsFlagsSet) {})
	w.OnDeleteMessages(func(e wrapper.DeleteMessages) {})
	w.OnRestoreDeletedMessages(func(e wrapper.RestoreDeletedMessages) {})
	w.OnChatParamsChange(func(e wrapper.ChatParamsChange) {
		assert.NotEmpty(t, e.ChatID)
	})
	w.OnChatInfoChange(func(e wrapper.ChatInfoChange) {
		assert.NotEmpty(t, e.TypeID)
		assert.NotEmpty(t, e.PeerID)
		// assert.NotEmpty(t, e.Info)
	})
	w.OnUserTyping(func(e wrapper.UserTyping) {})
	w.OnUserTypingChat(func(e wrapper.UserTypingChat) {})
	w.OnUsersTyping(func(e wrapper.UsersTyping) {})
	w.OnUsersRecordingAudioMessage(func(e wrapper.UsersRecordingAudioMessage) {})
	w.OnUserCall(func(e wrapper.UserCall) {})
	w.OnCounterChange(func(e wrapper.CounterChange) {})
	w.OnNotificationSettingsChange(func(e wrapper.NotificationSettingsChange) {})

	go func() {
		lpErr := lp.Run()
		assert.NoError(t, lpErr)
		exit <- true
	}()

	msgID, err := vkUser.MessagesSend(api.Params{
		"chat_id":   chatID,
		"random_id": 0,
		"message":   "Message",
	})
	assert.NoError(t, err)

	_, err = vkUser.MessagesEdit(api.Params{
		"peer_id":    2000000000 + chatID,
		"message_id": msgID,
		"message":    "Edit message",
	})
	assert.NoError(t, err)

	_, err = vkUser.MessagesPin(api.Params{
		"peer_id":    2000000000 + chatID,
		"message_id": msgID,
	})
	assert.NoError(t, err)

	_, err = vkUser.MessagesUnpin(api.Params{
		"peer_id": 2000000000 + chatID,
	})
	assert.NoError(t, err)

	_, err = vkUser.MessagesDelete(api.Params{
		"message_ids": msgID,
	})
	assert.NoError(t, err)

	_, err = vkUser.MessagesRestore(api.Params{
		"message_id": msgID,
	})
	assert.NoError(t, err)

	lp.Shutdown()
	<-exit
}
