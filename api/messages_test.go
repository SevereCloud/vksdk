package api_test

import (
	"errors"
	"testing"
	"time"

	"github.com/SevereCloud/vksdk/v3/api"
	"github.com/SevereCloud/vksdk/v3/api/params"
	"github.com/stretchr/testify/assert"
)

// func TestVK_MessagesAddChatUser(t *testing.T) {
// TODO: write test
// }

func TestVK_MessagesAllowMessagesFromGroup(t *testing.T) {
	t.Parallel()

	needUserToken(t)

	_, err := vkUser.MessagesAllowMessagesFromGroup(api.Params{
		"group_id": 1,
	})
	noError(t, err)
}

func TestVK_MessagesDelete(t *testing.T) {
	t.Parallel()

	needUserToken(t)
	chatID := needChatID(t)

	msgID, err := vkUser.MessagesSend(api.Params{
		"chat_id":   chatID,
		"random_id": 0,
		"message":   "Message",
	})
	noErrorOrFail(t, err)

	time.Sleep(10 * time.Second)

	_, err = vkUser.MessagesEdit(api.Params{
		"peer_id":    2000000000 + chatID,
		"message_id": msgID,
		"message":    "Edit message",
	})
	if err != nil {
		if !errors.Is(err, api.ErrRequest) {
			noError(t, err)
		}
	}

	pin, err := vkUser.MessagesPin(api.Params{
		"peer_id":    2000000000 + chatID,
		"message_id": msgID,
	})
	noError(t, err)
	assert.NotEmpty(t, pin.ID)
	assert.NotEmpty(t, pin.Date)
	assert.NotEmpty(t, pin.FromID)
	assert.NotEmpty(t, pin.PeerID)
	assert.NotEmpty(t, pin.Text)
	assert.NotEmpty(t, pin.ConversationMessageID)
	assert.NotEmpty(t, pin.Keyboard)

	time.Sleep(time.Second)

	_, err = vkUser.MessagesUnpin(api.Params{
		"peer_id": 2000000000 + chatID,
	})
	noError(t, err)

	get, err := vkUser.MessagesGetByID(api.Params{
		"message_ids": msgID,
	})
	noError(t, err)
	assert.NotEmpty(t, get.Count)

	if assert.NotEmpty(t, get.Items) {
		assert.NotEmpty(t, get.Items[0].Date)
		assert.NotEmpty(t, get.Items[0].FromID)
		assert.NotEmpty(t, get.Items[0].ID)
		assert.NotEmpty(t, get.Items[0].Out)
		assert.NotEmpty(t, get.Items[0].PeerID)
		assert.NotEmpty(t, get.Items[0].Text)
		assert.NotEmpty(t, get.Items[0].ConversationMessageID)

		getConversation, err := vkUser.MessagesGetByConversationMessageID(api.Params{
			"peer_id":                  2000000000 + chatID,
			"conversation_message_ids": get.Items[0].ConversationMessageID,
		})
		noError(t, err)
		assert.NotEmpty(t, getConversation.Count)
		assert.NotEmpty(t, getConversation.Items)
	}

	getEx, err := vkUser.MessagesGetByIDExtended(api.Params{
		"message_ids": msgID,
	})
	noError(t, err)
	assert.NotEmpty(t, getEx.Profiles)

	_, err = vkUser.MessagesDelete(api.Params{
		"message_ids": msgID,
	})
	noError(t, err)

	time.Sleep(time.Second)

	_, err = vkUser.MessagesRestore(api.Params{
		"message_id": msgID,
	})
	noError(t, err)
}

// func TestVK_MessagesDeleteChatPhoto(t *testing.T) {
// TODO: write test
// }

func TestVK_MessagesDeleteConversation(t *testing.T) {
	t.Parallel()

	needUserToken(t)
	chatID := needChatID(t)

	_, err := vkUser.MessagesDeleteConversation(api.Params{
		"peer_id": 2000000000 + chatID,
	})
	noError(t, err)
}

func TestVK_MessagesDenyMessagesFromGroup(t *testing.T) {
	t.Parallel()

	needUserToken(t)

	_, err := vkUser.MessagesDenyMessagesFromGroup(api.Params{
		"group_id": 1,
	})
	noError(t, err)
}

func TestVK_MessagesEditChat(t *testing.T) {
	t.Parallel()

	needUserToken(t)
	chatID := needChatID(t)

	_, err := vkUser.MessagesEditChat(api.Params{
		"chat_id": chatID,
		"title":   "New Title",
	})
	noError(t, err)
}

func TestVK_MessagesGetChat(t *testing.T) {
	t.Parallel()

	needUserToken(t)
	chatID := needChatID(t)

	chat, err := vkUser.MessagesGetChat(api.Params{
		"chat_id": chatID,
	})
	noError(t, err)
	assert.NotEmpty(t, chat.Type)
	assert.NotEmpty(t, chat.Title)
	assert.NotEmpty(t, chat.AdminID)
	assert.NotEmpty(t, chat.MembersCount)
	assert.NotEmpty(t, chat.ID)
	assert.NotEmpty(t, chat.Users)
}

func TestVK_MessagesGetChatChatIDs(t *testing.T) {
	t.Parallel()

	needUserToken(t)
	chatID := needChatID(t)

	chat, err := vkUser.MessagesGetChatChatIDs(api.Params{
		"chat_ids": chatID,
	})
	noError(t, err)

	if assert.NotEmpty(t, chat) {
		assert.NotEmpty(t, chat[0].Type)
		assert.NotEmpty(t, chat[0].Title)
		assert.NotEmpty(t, chat[0].AdminID)
		assert.NotEmpty(t, chat[0].MembersCount)
		assert.NotEmpty(t, chat[0].ID)
		assert.NotEmpty(t, chat[0].Users)
	}
}

func TestVK_MessagesGetChatPreview(t *testing.T) {
	t.Parallel()

	needUserToken(t)
	chatID := needChatID(t)

	res, err := vkUser.MessagesGetChatPreview(api.Params{
		"peer_id": 2000000000 + chatID,
	})
	noError(t, err)
	assert.NotEmpty(t, res.Preview.AdminID)
	assert.NotEmpty(t, res.Preview.MembersCount)
	assert.NotEmpty(t, res.Preview.Members)
	assert.NotEmpty(t, res.Preview.Title)
	assert.NotEmpty(t, res.Preview.LocalID)
	assert.NotEmpty(t, res.Preview.Joined)
	assert.NotEmpty(t, res.Profiles)
}

func TestVK_MessagesGetConversationMembers(t *testing.T) {
	t.Parallel()

	needUserToken(t)
	chatID := needChatID(t)

	res, err := vkUser.MessagesGetConversationMembers(api.Params{
		"peer_id": 2000000000 + chatID,
	})
	noError(t, err)
	assert.NotEmpty(t, res.Items)
	assert.NotEmpty(t, res.Count)
	// assert.NotEmpty(t, res.ChatRestrictions)
	assert.NotEmpty(t, res.Profiles)
}

func TestVK_MessagesGetConversations(t *testing.T) {
	t.Parallel()

	needUserToken(t)

	res, err := vkUser.MessagesGetConversations(nil)
	noError(t, err)
	assert.NotEmpty(t, res.Count)

	if assert.NotEmpty(t, res.Items) {
		assert.NotEmpty(t, res.Items[0].Conversation)
		assert.NotEmpty(t, res.Items[0].LastMessage)
	}
}

func TestVK_MessagesGetConversationsByID(t *testing.T) {
	t.Parallel()

	needUserToken(t)
	chatID := needChatID(t)

	res, err := vkUser.MessagesGetConversationsByID(api.Params{
		"peer_ids": 2000000000 + chatID,
	})
	noError(t, err)
	assert.NotEmpty(t, res.Count)

	if assert.NotEmpty(t, res.Items) {
		assert.NotEmpty(t, res.Items[0].Peer)
		assert.NotEmpty(t, res.Items[0].ChatSettings.OwnerID)
		assert.NotEmpty(t, res.Items[0].ChatSettings.Title)
		assert.NotEmpty(t, res.Items[0].ChatSettings.State)
		assert.NotEmpty(t, res.Items[0].ChatSettings.ACL)
		assert.NotEmpty(t, res.Items[0].ChatSettings.MembersCount)
		assert.NotEmpty(t, res.Items[0].ChatSettings.ActiveIDs)
	}
}

func TestVK_MessagesGetConversationsByIDExtended(t *testing.T) {
	t.Parallel()

	needUserToken(t)
	chatID := needChatID(t)

	_, err := vkUser.MessagesGetConversationsByIDExtended(api.Params{
		"peer_ids": 2000000000 + chatID,
	})
	noError(t, err)
}

func TestVK_MessagesGetHistory(t *testing.T) {
	t.Parallel()

	needUserToken(t)
	chatID := needChatID(t)

	_, err := vkUser.MessagesGetHistory(api.Params{
		"peer_id": 2000000000 + chatID,
	})
	noError(t, err)
}

func TestVK_MessagesGetHistoryAttachments(t *testing.T) {
	t.Parallel()

	needUserToken(t)
	chatID := needChatID(t)

	_, err := vkUser.MessagesGetHistoryAttachments(api.Params{
		"peer_id":    2000000000 + chatID,
		"media_type": "photo",
	})
	noError(t, err)
}

func TestVK_MessagesGetImportantMessages(t *testing.T) {
	t.Parallel()

	needUserToken(t)

	_, err := vkUser.MessagesGetImportantMessages(nil)
	noError(t, err)
}

func TestVK_MessagesGetIntentUsers(t *testing.T) {
	t.Parallel()

	needGroupToken(t)

	_, err := vkGroup.MessagesGetIntentUsers(api.Params{
		"intent":       params.PromoNewsletter,
		"subscribe_id": 0,
	})
	noError(t, err)
}

func TestVK_MessagesGetInviteLink(t *testing.T) {
	t.Parallel()

	needUserToken(t)
	chatID := needChatID(t)

	res, err := vkUser.MessagesGetInviteLink(api.Params{
		"peer_id": 2000000000 + chatID,
	})
	noError(t, err)
	assert.NotEmpty(t, res.Link)
}

func TestVK_MessagesGetLastActivity(t *testing.T) {
	t.Parallel()

	needUserToken(t)

	_, err := vkUser.MessagesGetLastActivity(api.Params{
		"user_id": 6492,
	})
	noError(t, err)
}

func TestVK_MessagesGetLongPollServer(t *testing.T) {
	t.Parallel()

	needUserToken(t)

	res, err := vkUser.MessagesGetLongPollServer(api.Params{
		"need_pts":   1,
		"lp_version": 3,
	})
	noError(t, err)
	assert.NotEmpty(t, res.Key)
	assert.NotEmpty(t, res.Server)
	assert.NotEmpty(t, res.Ts)
	assert.NotEmpty(t, res.Pts)

	_, err = vkUser.MessagesGetLongPollHistory(api.Params{
		"ts":  res.Ts,
		"pts": res.Pts,
	})
	noError(t, err)
}

func TestVK_MessagesIsMessagesFromGroupAllowed(t *testing.T) {
	t.Parallel()

	needGroupToken(t)

	_, err := vkGroup.MessagesIsMessagesFromGroupAllowed(api.Params{
		"user_id":  2314852,
		"group_id": vkGroupID,
	})
	noError(t, err)
}

func TestVK_MessagesJoinChatByInviteLink(t *testing.T) {
	t.Parallel()

	needUserToken(t)

	chat, err := vkUser.MessagesJoinChatByInviteLink(api.Params{
		"link": "https://vk.me/join/AJQ1dzddEBakBaMV8L6vLbG1",
	})
	noError(t, err)

	if assert.NotEmpty(t, chat.ChatID) {
		_, err := vkUser.MessagesRemoveChatUser(api.Params{
			"chat_id": chat.ChatID,
			"user_id": vkUserID,
		})
		noError(t, err)
	}
}

// func TestVK_MessagesMarkAsAnsweredConversation(t *testing.T) {
// TODO: write test
// }

func TestVK_MessagesMarkAsImportant(t *testing.T) {
	t.Parallel()

	needUserToken(t)

	_, err := vkUser.MessagesMarkAsImportant(api.Params{
		"message_ids": 1,
		"important":   false,
	})
	noError(t, err)
}

// func TestVK_MessagesMarkAsImportantConversation(t *testing.T) {
// TODO: write test
// }

func TestVK_MessagesMarkAsRead(t *testing.T) {
	t.Parallel()

	needUserToken(t)
	chatID := needChatID(t)

	_, err := vkUser.MessagesMarkAsRead(api.Params{
		"peer_id": 2000000000 + chatID,
	})
	noError(t, err)
}

func TestVK_MessagesRemoveChatUser(t *testing.T) {
	t.Parallel()

	needUserToken(t)
	chatID := needChatID(t)

	_, err := vkUser.MessagesRemoveChatUser(api.Params{
		"chat_id": chatID,
		"user_id": vkUserID,
	})
	noError(t, err)

	_, err = vkUser.MessagesAddChatUser(api.Params{
		"chat_id": chatID,
		"user_id": vkUserID,
	})
	noError(t, err)
}

func TestVK_MessagesSearch(t *testing.T) {
	t.Parallel()

	needUserToken(t)

	_, err := vkUser.MessagesSearch(api.Params{
		"q": "test message",
	})
	noError(t, err)
}

func TestVK_MessagesSearchConversations(t *testing.T) {
	t.Parallel()

	needUserToken(t)

	_, err := vkUser.MessagesSearchConversations(nil)
	noError(t, err)
}

// func TestVK_MessagesSendUserIDs(t *testing.T) {
// TODO: write test
// }

func TestVK_MessagesSendSticker(t *testing.T) {
	t.Parallel()

	needUserToken(t)
	chatID := needChatID(t)

	_, err := vkUser.MessagesSendSticker(api.Params{
		"chat_id":    chatID,
		"random_id":  0,
		"sticker_id": 9008,
	})
	noError(t, err)
}

func TestVK_MessagesSetActivity(t *testing.T) {
	t.Parallel()

	needUserToken(t)
	chatID := needChatID(t)

	_, err := vkUser.MessagesSetActivity(api.Params{
		"peer_id": 2000000000 + chatID,
		"type":    "typing",
	})
	noError(t, err)
}

func TestVK_MessagesCalls(t *testing.T) {
	t.Parallel()

	needUserToken(t)

	startCall, err := vkUser.MessagesStartCall(nil)
	noError(t, err)

	_, err = vkUser.MessagesForceCallFinish(api.Params{
		"call_id": startCall.CallID,
	})
	noError(t, err)
}

func TestVK_MessagesUnpin(t *testing.T) {
	t.Parallel()

	needUserToken(t)
	chatID := needChatID(t)

	_, err := vkUser.MessagesUnpin(api.Params{
		"peer_id": 2000000000 + chatID,
	})
	noError(t, err)
}
