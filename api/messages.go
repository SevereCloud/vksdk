package api // import "github.com/severecloud/vksdk/api"

import (
	"encoding/json"

	"github.com/severecloud/vksdk/object"
)

// MessagesAddChatUser adds a new user to a chat.
func (vk VK) MessagesAddChatUser(params map[string]string) (vkErr Error) {
	_, vkErr = vk.Request("messages.addChatUser", params)

	return
}

// MessagesAllowMessagesFromGroup allows sending messages from community to the current user.
func (vk VK) MessagesAllowMessagesFromGroup(params map[string]string) (vkErr Error) {
	_, vkErr = vk.Request("messages.allowMessagesFromGroup", params)

	return
}

// MessagesCreateChatResponse struct
type MessagesCreateChatResponse int

// MessagesCreateChat creates a chat with several participants.
func (vk VK) MessagesCreateChat(params map[string]string) (response MessagesCreateChatResponse, vkErr Error) {
	rawResponse, vkErr := vk.Request("messages.createChat", params)
	if vkErr.Code != 0 {
		return
	}

	err := json.Unmarshal(rawResponse, &response)
	if err != nil {
		panic(err)
	}

	return
}

// MessagesDeleteResponse struct
type MessagesDeleteResponse map[string]int

// MessagesDelete deletes one or more messages.
func (vk VK) MessagesDelete(params map[string]string) (response MessagesDeleteResponse, vkErr Error) {
	rawResponse, vkErr := vk.Request("messages.delete", params)
	if vkErr.Code != 0 {
		return
	}

	err := json.Unmarshal(rawResponse, &response)
	if err != nil {
		panic(err)
	}

	return
}

// MessagesDeleteChatPhotoResponse struct
type MessagesDeleteChatPhotoResponse struct {
	MessageID int                 `json:"message_id"`
	Chat      object.MessagesChat `json:"chat"`
}

// MessagesDeleteChatPhoto deletes a chat's cover picture.
func (vk VK) MessagesDeleteChatPhoto(params map[string]string) (response MessagesDeleteChatPhotoResponse, vkErr Error) {
	rawResponse, vkErr := vk.Request("messages.deleteChatPhoto", params)
	if vkErr.Code != 0 {
		return
	}

	err := json.Unmarshal(rawResponse, &response)
	if err != nil {
		panic(err)
	}

	return
}

// MessagesDeleteConversation deletes private messages in a conversation.
func (vk VK) MessagesDeleteConversation(params map[string]string) (vkErr Error) {
	_, vkErr = vk.Request("messages.deleteConversation", params)

	return
}

// MessagesDenyMessagesFromGroup denies sending message from community to the current user.
func (vk VK) MessagesDenyMessagesFromGroup(params map[string]string) (vkErr Error) {
	_, vkErr = vk.Request("messages.denyMessagesFromGroup", params)

	return
}

// MessagesEdit edits the message.
func (vk VK) MessagesEdit(params map[string]string) (vkErr Error) {
	_, vkErr = vk.Request("messages.edit", params)

	return
}

// MessagesEditChat edits the title of a chat.
func (vk VK) MessagesEditChat(params map[string]string) (vkErr Error) {
	_, vkErr = vk.Request("messages.editChat", params)

	return
}

// MessagesGetByConversationMessageIDResponse struct
type MessagesGetByConversationMessageIDResponse struct {
	Count    int                      `json:"count"`
	Items    []object.MessagesMessage `json:"items"`
	Profiles []object.UsersUser       `json:"profiles"`
	Groups   []object.GroupsGroup     `json:"groups"`
}

// MessagesGetByConversationMessageID messages.getByConversationMessageId
func (vk VK) MessagesGetByConversationMessageID(params map[string]string) (response MessagesGetByConversationMessageIDResponse, vkErr Error) {
	rawResponse, vkErr := vk.Request("messages.getByConversationMessageId", params)
	if vkErr.Code != 0 {
		return
	}

	err := json.Unmarshal(rawResponse, &response)
	if err != nil {
		panic(err)
	}

	return
}

// MessagesGetByIDResponse struct
type MessagesGetByIDResponse struct {
	Count int                      `json:"count"`
	Items []object.MessagesMessage `json:"items"`
}

// MessagesGetByID returns messages by their IDs.
func (vk VK) MessagesGetByID(params map[string]string) (response MessagesGetByIDResponse, vkErr Error) {
	rawResponse, vkErr := vk.Request("messages.getById", params)
	if vkErr.Code != 0 {
		return
	}

	err := json.Unmarshal(rawResponse, &response)
	if err != nil {
		panic(err)
	}

	return
}

// MessagesGetChatResponse struct
type MessagesGetChatResponse struct{}

// TODO: messages.getChat Returns information about a chat.

// MessagesGetChatPreviewResponse struct
type MessagesGetChatPreviewResponse struct {
	Preview  object.MessagesChat  `json:"preview"`
	Profiles []object.UsersUser   `json:"profiles"`
	Groups   []object.GroupsGroup `json:"groups"`
}

// MessagesGetChatPreview allows to receive chat preview by the invitation link.
func (vk VK) MessagesGetChatPreview(params map[string]string) (response MessagesGetChatPreviewResponse, vkErr Error) {
	rawResponse, vkErr := vk.Request("messages.getChatPreview", params)
	if vkErr.Code != 0 {
		return
	}

	err := json.Unmarshal(rawResponse, &response)
	if err != nil {
		panic(err)
	}

	return
}

// MessagesGetConversationMembersResponse struct
type MessagesGetConversationMembersResponse struct{}

// TODO: messages.getConversationMembers Returns a list of IDs of users participating in a conversation.

// MessagesGetConversationsResponse struct
type MessagesGetConversationsResponse struct {
	Count       int                                      `json:"count"`
	Items       []object.MessagesConversationWithMessage `json:"items"`
	UnreadCount int                                      `json:"unread_count"`
	Profiles    []object.UsersUser                       `json:"profiles"`
	Groups      []object.GroupsGroup                     `json:"groups"`
	//
}

// MessagesGetConversations returns a list of conversations.
func (vk VK) MessagesGetConversations(params map[string]string) (response MessagesGetConversationsResponse, vkErr Error) {
	rawResponse, vkErr := vk.Request("messages.getConversations", params)
	if vkErr.Code != 0 {
		return
	}

	err := json.Unmarshal(rawResponse, &response)
	if err != nil {
		panic(err)
	}

	return
}

// MessagesGetConversationsByIDResponse struct
type MessagesGetConversationsByIDResponse struct {
	Count    int                           `json:"count"`
	Items    []object.MessagesConversation `json:"items"`
	Profiles []object.UsersUser            `json:"profiles"`
	Groups   []object.GroupsGroup          `json:"groups"`
}

// MessagesGetConversationsByID returns conversations by their IDs.
func (vk VK) MessagesGetConversationsByID(params map[string]string) (response MessagesGetConversationsByIDResponse, vkErr Error) {
	rawResponse, vkErr := vk.Request("messages.getConversationsById", params)
	if vkErr.Code != 0 {
		return
	}

	err := json.Unmarshal(rawResponse, &response)
	if err != nil {
		panic(err)
	}

	return
}

// MessagesGetHistoryResponse struct
type MessagesGetHistoryResponse struct {
	Count   int                      `json:"count"`
	Items   []object.MessagesMessage `json:"items"`
	InRead  int                      `json:"in_read"`
	OutRead int                      `json:"out_read"`
}

// MessagesGetHistory returns message history for the specified user or group chat.
func (vk VK) MessagesGetHistory(params map[string]string) (response MessagesGetHistoryResponse, vkErr Error) {
	rawResponse, vkErr := vk.Request("messages.getHistory", params)
	if vkErr.Code != 0 {
		return
	}

	err := json.Unmarshal(rawResponse, &response)
	if err != nil {
		panic(err)
	}

	return
}

// MessagesGetHistoryAttachmentsResponse struct
type MessagesGetHistoryAttachmentsResponse struct {
	Items    []object.MessagesHistoryAttachment `json:"items"`
	NextFrom string                             `json:"next_from"`
}

// MessagesGetHistoryAttachments returns media files from the dialog or group chat.
func (vk VK) MessagesGetHistoryAttachments(params map[string]string) (response MessagesGetHistoryAttachmentsResponse, vkErr Error) {
	rawResponse, vkErr := vk.Request("messages.getHistoryAttachments", params)
	if vkErr.Code != 0 {
		return
	}

	err := json.Unmarshal(rawResponse, &response)
	if err != nil {
		panic(err)
	}

	return
}

// MessagesGetImportantMessagesResponse struct
type MessagesGetImportantMessagesResponse struct {
	Messages struct {
		Count int                      `json:"count"`
		Items []object.MessagesMessage `json:"items"`
	} `json:"messages"`
	Profiles      []object.UsersUser            `json:"profiles"`
	Groups        []object.GroupsGroup          `json:"groups"`
	Conversations []object.MessagesConversation `json:"conversations"`
}

// MessagesGetImportantMessages messages.getImportantMessages
func (vk VK) MessagesGetImportantMessages(params map[string]string) (response MessagesGetImportantMessagesResponse, vkErr Error) {
	rawResponse, vkErr := vk.Request("messages.getImportantMessages", params)
	if vkErr.Code != 0 {
		return
	}

	err := json.Unmarshal(rawResponse, &response)
	if err != nil {
		panic(err)
	}

	return
}

// MessagesGetInviteLinkResponse struct
type MessagesGetInviteLinkResponse struct {
	Link string `json:"link"`
}

// MessagesGetInviteLink receives a link to invite a user to the chat.
func (vk VK) MessagesGetInviteLink(params map[string]string) (response MessagesGetInviteLinkResponse, vkErr Error) {
	rawResponse, vkErr := vk.Request("messages.getInviteLink", params)
	if vkErr.Code != 0 {
		return
	}

	err := json.Unmarshal(rawResponse, &response)
	if err != nil {
		panic(err)
	}

	return
}

// MessagesGetLastActivityResponse struct
type MessagesGetLastActivityResponse object.MessagesLastActivity

// MessagesGetLastActivity returns a user's current status and date of last activity.
func (vk VK) MessagesGetLastActivity(params map[string]string) (response MessagesGetLastActivityResponse, vkErr Error) {
	rawResponse, vkErr := vk.Request("messages.getLastActivity", params)
	if vkErr.Code != 0 {
		return
	}

	err := json.Unmarshal(rawResponse, &response)
	if err != nil {
		panic(err)
	}

	return
}

// MessagesGetLongPollHistoryResponse struct
type MessagesGetLongPollHistoryResponse struct{}

// TODO: messages.getLongPollHistory Returns updates in user's private messages.

// MessagesGetLongPollServerResponse struct
type MessagesGetLongPollServerResponse object.MessagesLongpollParams

// MessagesGetLongPollServer returns data required for connection to a Long Poll server.
func (vk VK) MessagesGetLongPollServer(params map[string]string) (response MessagesGetLongPollServerResponse, vkErr Error) {
	rawResponse, vkErr := vk.Request("messages.getLongPollServer", params)
	if vkErr.Code != 0 {
		return
	}

	err := json.Unmarshal(rawResponse, &response)
	if err != nil {
		panic(err)
	}

	return
}

// MessagesIsMessagesFromGroupAllowedResponse struct
type MessagesIsMessagesFromGroupAllowedResponse struct {
	IsAllowed int `json:"is_allowed"`
}

// MessagesIsMessagesFromGroupAllowed returns information whether sending messages from the community to current user is allowed.
// BUG(VK): need group_id with group token vk.com/bugtracker?act=show&id=88011
func (vk VK) MessagesIsMessagesFromGroupAllowed(params map[string]string) (response MessagesIsMessagesFromGroupAllowedResponse, vkErr Error) {
	rawResponse, vkErr := vk.Request("messages.isMessagesFromGroupAllowed", params)
	if vkErr.Code != 0 {
		return
	}

	err := json.Unmarshal(rawResponse, &response)
	if err != nil {
		panic(err)
	}

	return
}

// MessagesJoinChatByInviteLinkResponse struct
type MessagesJoinChatByInviteLinkResponse struct {
	ChatID int `json:"chat_id"`
}

// MessagesJoinChatByInviteLink allows to enter the chat by the invitation link.
func (vk VK) MessagesJoinChatByInviteLink(params map[string]string) (response MessagesJoinChatByInviteLinkResponse, vkErr Error) {
	rawResponse, vkErr := vk.Request("messages.joinChatByInviteLink", params)
	if vkErr.Code != 0 {
		return
	}

	err := json.Unmarshal(rawResponse, &response)
	if err != nil {
		panic(err)
	}

	return
}

// MessagesMarkAsAnsweredConversation messages.markAsAnsweredConversation
func (vk VK) MessagesMarkAsAnsweredConversation(params map[string]string) (vkErr Error) {
	_, vkErr = vk.Request("messages.markAsAnsweredConversation", params)

	return
}

// MessagesMarkAsImportantResponse struct
type MessagesMarkAsImportantResponse []int

// MessagesMarkAsImportant marks and un marks messages as important (starred).
func (vk VK) MessagesMarkAsImportant(params map[string]string) (response MessagesMarkAsImportantResponse, vkErr Error) {
	rawResponse, vkErr := vk.Request("messages.markAsImportant", params)
	if vkErr.Code != 0 {
		return
	}

	err := json.Unmarshal(rawResponse, &response)
	if err != nil {
		panic(err)
	}

	return
}

// MessagesMarkAsImportantConversation messages.markAsImportantConversation
func (vk VK) MessagesMarkAsImportantConversation(params map[string]string) (vkErr Error) {
	_, vkErr = vk.Request("messages.markAsImportantConversation", params)

	return
}

// MessagesMarkAsRead marks messages as read.
func (vk VK) MessagesMarkAsRead(params map[string]string) (vkErr Error) {
	_, vkErr = vk.Request("messages.markAsRead", params)

	return
}

// MessagesPinResponse struct
type MessagesPinResponse object.MessagesMessage

// MessagesPin messages.pin
func (vk VK) MessagesPin(params map[string]string) (response MessagesPinResponse, vkErr Error) {
	rawResponse, vkErr := vk.Request("messages.pin", params)
	if vkErr.Code != 0 {
		return
	}

	err := json.Unmarshal(rawResponse, &response)
	if err != nil {
		panic(err)
	}

	return
}

// MessagesRemoveChatUser Allows the current user to leave a chat or, if the current user started the chat, allows the user to remove another user from the chat.
func (vk VK) MessagesRemoveChatUser(params map[string]string) (vkErr Error) {
	_, vkErr = vk.Request("messages.removeChatUser", params)

	return
}

// MessagesRestore restores a deleted message.
func (vk VK) MessagesRestore(params map[string]string) (vkErr Error) {
	_, vkErr = vk.Request("messages.restore", params)

	return
}

// MessagesSearchResponse struct
type MessagesSearchResponse struct {
	Count         int                           `json:"count"`
	Items         []object.MessagesMessage      `json:"items"`
	Profiles      []object.UsersUser            `json:"profiles"`
	Groups        []object.GroupsGroup          `json:"groups"`
	Conversations []object.MessagesConversation `json:"conversations"`
}

// MessagesSearch Returns a list of the current user's private messages that match search criteria.
func (vk VK) MessagesSearch(params map[string]string) (response MessagesSearchResponse, vkErr Error) {
	rawResponse, vkErr := vk.Request("messages.search", params)
	if vkErr.Code != 0 {
		return
	}

	err := json.Unmarshal(rawResponse, &response)
	if err != nil {
		panic(err)
	}

	return
}

// MessagesSearchConversationsResponse struct
type MessagesSearchConversationsResponse struct{}

// TODO: messages.searchConversations Returns a list of conversations that match search criteria.

// MessagesSendResponse struct
type MessagesSendResponse struct{}

// MessagesSend Sends a message
// TODO: messages.send Sends a message.
func (vk VK) MessagesSend(params map[string]string) (vkErr Error) {
	const method = "messages.send"
	// TODO: if user_ids in params

	_, vkErr = vk.Request(method, params)
	if vkErr.Code != 0 {
		return
	}

	return
}

// MessagesSetActivity changes the status of a user as typing in a conversation.
func (vk VK) MessagesSetActivity(params map[string]string) (vkErr Error) {
	_, vkErr = vk.Request("messages.setActivity", params)

	return
}

// MessagesSetChatPhotoResponse struct
type MessagesSetChatPhotoResponse struct {
	MessageID int                 `json:"message_id"`
	Chat      object.MessagesChat `json:"chat"`
}

// MessagesSetChatPhoto sets a previously-uploaded picture as the cover picture of a chat.
func (vk VK) MessagesSetChatPhoto(params map[string]string) (response MessagesSetChatPhotoResponse, vkErr Error) {
	rawResponse, vkErr := vk.Request("messages.setChatPhoto", params)
	if vkErr.Code != 0 {
		return
	}

	err := json.Unmarshal(rawResponse, &response)
	if err != nil {
		panic(err)
	}

	return
}

// MessagesUnpin messages.unpin
func (vk VK) MessagesUnpin(params map[string]string) (vkErr Error) {
	_, vkErr = vk.Request("messages.unpin", params)

	return
}
