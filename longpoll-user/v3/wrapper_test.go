package wrapper_test

import (
	"log"
	"os"
	"testing"
	"time"

	"github.com/SevereCloud/vksdk/v2/api"
	wrapper "github.com/SevereCloud/vksdk/v2/longpoll-user/v3"
	"github.com/stretchr/testify/assert"

	"github.com/SevereCloud/vksdk/v2/longpoll-user"
)

func needUserToken(t *testing.T) {
	t.Helper()

	if vkUser == nil {
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

var (
	vkUser             *api.VK //nolint:gochecknoglobals
	vkUserID, vkChatID int     //nolint:gochecknoglobals
)

func TestMain(m *testing.M) {
	time.Sleep(1 * time.Second)

	if token := os.Getenv("USER_TOKEN"); token != "" {
		vkUser = api.NewVK(token)
		vkUser.Limit = 3

		user, err := vkUser.UsersGet(nil)
		if err != nil {
			log.Fatalf("USER_TOKEN bad: %v", err)
		}

		vkUserID = user[0].ID
	}

	runTests := m.Run()
	os.Exit(runTests)
}

func TestWrapper(t *testing.T) {
	t.Parallel()

	needUserToken(t)
	chatID := needChatID(t)

	exit := make(chan bool, 1)

	lp, err := longpoll.NewLongPoll(vkUser, longpoll.ExtendedEvents+longpoll.Code8ExtraFields)
	if err != nil {
		t.Fatalf("lp.Init err: %v", err)
	}

	w := wrapper.NewWrapper(lp)

	w.OnMessageFlagsChange(func(_ wrapper.MessageFlagsChange) {})
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
	w.OnReadInMessages(func(_ wrapper.ReadInMessages) {})
	w.OnReadOutMessages(func(_ wrapper.ReadOutMessages) {})
	w.OnFriendBecameOnline(func(_ wrapper.FriendBecameOnline) {})
	w.OnFriendBecameOffline(func(_ wrapper.FriendBecameOffline) {})
	w.OnDialogFlagsReset(func(_ wrapper.DialogFlagsReset) {})
	w.OnDialogFlagsReplace(func(_ wrapper.DialogFlagsReplace) {})
	w.OnDialogsFlagsSet(func(_ wrapper.DialogsFlagsSet) {})
	w.OnDeleteMessages(func(_ wrapper.DeleteMessages) {})
	w.OnRestoreDeletedMessages(func(_ wrapper.RestoreDeletedMessages) {})
	w.OnChatParamsChange(func(e wrapper.ChatParamsChange) {
		assert.NotEmpty(t, e.ChatID)
	})
	w.OnChatInfoChange(func(e wrapper.ChatInfoChange) {
		assert.NotEmpty(t, e.TypeID)
		assert.NotEmpty(t, e.PeerID)
		lp.Shutdown()
	})
	w.OnUserTyping(func(_ wrapper.UserTyping) {})
	w.OnUserTypingChat(func(_ wrapper.UserTypingChat) {})
	w.OnUsersTyping(func(_ wrapper.UsersTyping) {})
	w.OnUsersRecordingAudioMessage(func(_ wrapper.UsersRecordingAudioMessage) {})
	w.OnUserCall(func(_ wrapper.UserCall) {})
	w.OnCounterChange(func(_ wrapper.CounterChange) {})
	w.OnNotificationSettingsChange(func(_ wrapper.NotificationSettingsChange) {})

	f := func() {
		lp := lp

		t.Log("Run")

		lpErr := lp.Run()
		assert.NoError(t, lpErr)
		exit <- true
	}

	go f()

	time.Sleep(1 * time.Second)

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

	select {
	case <-time.After(5 * time.Second):
		t.Fatal("Time out")
	case <-exit:
		t.Log("ok")
	}
}
