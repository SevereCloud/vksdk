package api // import "github.com/SevereCloud/vksdk/5.92/api"

import (
	"github.com/SevereCloud/vksdk/5.92/object"
)

// MessagesAddChatUserResponse struct
type MessagesAddChatUserResponse int

// MessagesAddChatUser adds a new user to a chat.
//
// https://vk.com/dev/messages.addChatUser
func (vk *VK) MessagesAddChatUser(params map[string]string) (response MessagesAddChatUserResponse, vkErr Error) {
	vk.RequestUnmarshal("messages.addChatUser", params, &response, &vkErr)
	return
}

// MessagesAllowMessagesFromGroupResponse struct
type MessagesAllowMessagesFromGroupResponse int

// MessagesAllowMessagesFromGroup allows sending messages from community to the current user.
//
// https://vk.com/dev/messages.allowMessagesFromGroup
func (vk *VK) MessagesAllowMessagesFromGroup(params map[string]string) (response MessagesAllowMessagesFromGroupResponse, vkErr Error) {
	vk.RequestUnmarshal("messages.allowMessagesFromGroup", params, &response, &vkErr)
	return
}

// MessagesCreateChatResponse struct
type MessagesCreateChatResponse int

// MessagesCreateChat creates a chat with several participants.
//
// https://vk.com/dev/messages.createChat
func (vk *VK) MessagesCreateChat(params map[string]string) (response int, vkErr Error) {
	vk.RequestUnmarshal("messages.createChat", params, &response, &vkErr)
	return
}

// MessagesDeleteResponse struct
type MessagesDeleteResponse map[string]int

// MessagesDelete deletes one or more messages.
//
// https://vk.com/dev/messages.delete
func (vk *VK) MessagesDelete(params map[string]string) (response MessagesDeleteResponse, vkErr Error) {
	vk.RequestUnmarshal("messages.delete", params, &response, &vkErr)
	return
}

// MessagesDeleteChatPhotoResponse struct
type MessagesDeleteChatPhotoResponse struct {
	MessageID int                 `json:"message_id"`
	Chat      object.MessagesChat `json:"chat"`
}

// MessagesDeleteChatPhoto deletes a chat's cover picture.
//
// https://vk.com/dev/messages.deleteChatPhoto
func (vk *VK) MessagesDeleteChatPhoto(params map[string]string) (response MessagesDeleteChatPhotoResponse, vkErr Error) {
	vk.RequestUnmarshal("messages.deleteChatPhoto", params, &response, &vkErr)
	return
}

// MessagesDeleteConversationResponse struct
type MessagesDeleteConversationResponse int

// MessagesDeleteConversation deletes private messages in a conversation.
//
// https://vk.com/dev/messages.deleteConversation
func (vk *VK) MessagesDeleteConversation(params map[string]string) (response MessagesDeleteConversationResponse, vkErr Error) {
	vk.RequestUnmarshal("messages.deleteConversation", params, &response, &vkErr)
	return
}

// MessagesDenyMessagesFromGroupResponse struct
type MessagesDenyMessagesFromGroupResponse int

// MessagesDenyMessagesFromGroup denies sending message from community to the current user.
//
// https://vk.com/dev/messages.denyMessagesFromGroup
func (vk *VK) MessagesDenyMessagesFromGroup(params map[string]string) (response MessagesDenyMessagesFromGroupResponse, vkErr Error) {
	vk.RequestUnmarshal("messages.denyMessagesFromGroup", params, &response, &vkErr)
	return
}

// MessagesEditResponse struct
type MessagesEditResponse int

// MessagesEdit edits the message.
//
// https://vk.com/dev/messages.edit
func (vk *VK) MessagesEdit(params map[string]string) (response MessagesEditResponse, vkErr Error) {
	vk.RequestUnmarshal("messages.edit", params, &response, &vkErr)
	return
}

// MessagesEditChatResponse struct
type MessagesEditChatResponse int

// MessagesEditChat edits the title of a chat.
//
// https://vk.com/dev/messages.editChat
func (vk *VK) MessagesEditChat(params map[string]string) (response MessagesEditChatResponse, vkErr Error) {
	vk.RequestUnmarshal("messages.editChat", params, &response, &vkErr)
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
//
// https://vk.com/dev/messages.getByConversationMessageId
func (vk *VK) MessagesGetByConversationMessageID(params map[string]string) (response MessagesGetByConversationMessageIDResponse, vkErr Error) {
	vk.RequestUnmarshal("messages.getByConversationMessageId", params, &response, &vkErr)
	return
}

// MessagesGetByIDResponse struct
type MessagesGetByIDResponse struct {
	Count int                      `json:"count"`
	Items []object.MessagesMessage `json:"items"`
}

// MessagesGetByID returns messages by their IDs.
//
// extended=0
//
// https://vk.com/dev/messages.getById
func (vk *VK) MessagesGetByID(params map[string]string) (response MessagesGetByIDResponse, vkErr Error) {
	params["extended"] = "0"
	vk.RequestUnmarshal("messages.getById", params, &response, &vkErr)
	return
}

// MessagesGetByIDExtendedResponse struct
type MessagesGetByIDExtendedResponse struct {
	Count    int                      `json:"count"`
	Items    []object.MessagesMessage `json:"items"`
	Profiles []object.UsersUser       `json:"profiles"`
	Groups   []object.GroupsGroup     `json:"groups"`
}

// MessagesGetByIDExtended returns messages by their IDs.
//
// extended=1
//
// https://vk.com/dev/messages.getById
func (vk *VK) MessagesGetByIDExtended(params map[string]string) (response MessagesGetByIDExtendedResponse, vkErr Error) {
	params["extended"] = "1"
	vk.RequestUnmarshal("messages.getById", params, &response, &vkErr)
	return
}

// MessagesGetChatResponse struct
type MessagesGetChatResponse object.MessagesChat

// MessagesGetChat returns information about a chat.
//
// https://vk.com/dev/messages.getChat
func (vk *VK) MessagesGetChat(params map[string]string) (response MessagesGetChatResponse, vkErr Error) {
	vk.RequestUnmarshal("messages.getChat", params, &response, &vkErr)
	return
}

// MessagesGetChatChatIDsResponse struct
type MessagesGetChatChatIDsResponse []object.MessagesChat

// MessagesGetChatChatIDs returns information about a chat.
//
// https://vk.com/dev/messages.getChat
func (vk *VK) MessagesGetChatChatIDs(params map[string]string) (response MessagesGetChatChatIDsResponse, vkErr Error) {
	vk.RequestUnmarshal("messages.getChat", params, &response, &vkErr)
	return
}

// MessagesGetChatPreviewResponse struct
type MessagesGetChatPreviewResponse struct {
	Preview  object.MessagesChat  `json:"preview"`
	Profiles []object.UsersUser   `json:"profiles"`
	Groups   []object.GroupsGroup `json:"groups"`
}

// MessagesGetChatPreview allows to receive chat preview by the invitation link.
//
// https://vk.com/dev/messages.getChatPreview
func (vk *VK) MessagesGetChatPreview(params map[string]string) (response MessagesGetChatPreviewResponse, vkErr Error) {
	vk.RequestUnmarshal("messages.getChatPreview", params, &response, &vkErr)
	return
}

// MessagesGetConversationMembersResponse struct
type MessagesGetConversationMembersResponse struct {
	Items []struct {
		MemberID  int  `json:"member_id"`
		JoinDate  int  `json:"join_date"`
		InvitedBy int  `json:"invited_by"`
		IsOwner   bool `json:"is_owner,omitempty"`
		IsAdmin   bool `json:"is_admin,omitempty"`
		CanKick   bool `json:"can_kick,omitempty"`
	} `json:"items"`
	Count            int `json:"count"`
	ChatRestrictions struct {
		OnlyAdminsInvite   bool `json:"only_admins_invite"`
		OnlyAdminsEditPin  bool `json:"only_admins_edit_pin"`
		OnlyAdminsEditInfo bool `json:"only_admins_edit_info"`
		AdminsPromoteUsers bool `json:"admins_promote_users"`
	} `json:"chat_restrictions"`
	Profiles []object.UsersUser   `json:"profiles"`
	Groups   []object.GroupsGroup `json:"groups"`
}

// MessagesGetConversationMembers Returns a list of IDs of users participating in a conversation.
//
// https://vk.com/dev/messages.getConversationMembers
func (vk *VK) MessagesGetConversationMembers(params map[string]string) (response MessagesGetConversationMembersResponse, vkErr Error) {
	vk.RequestUnmarshal("messages.getConversationMembers", params, &response, &vkErr)
	return
}

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
//
// https://vk.com/dev/messages.getConversations
func (vk *VK) MessagesGetConversations(params map[string]string) (response MessagesGetConversationsResponse, vkErr Error) {
	vk.RequestUnmarshal("messages.getConversations", params, &response, &vkErr)
	return
}

// MessagesGetConversationsByIDResponse struct
type MessagesGetConversationsByIDResponse struct {
	Count int                           `json:"count"`
	Items []object.MessagesConversation `json:"items"`
}

// MessagesGetConversationsByID returns conversations by their IDs.
//
// extended=0
//
// https://vk.com/dev/messages.getConversationsById
func (vk *VK) MessagesGetConversationsByID(params map[string]string) (response MessagesGetConversationsByIDResponse, vkErr Error) {
	params["extended"] = "0"
	vk.RequestUnmarshal("messages.getConversationsById", params, &response, &vkErr)
	return
}

// MessagesGetConversationsByIDExtendedResponse struct
type MessagesGetConversationsByIDExtendedResponse struct {
	Count    int                           `json:"count"`
	Items    []object.MessagesConversation `json:"items"`
	Profiles []object.UsersUser            `json:"profiles"`
	Groups   []object.GroupsGroup          `json:"groups"`
}

// MessagesGetConversationsByIDExtended returns conversations by their IDs.
//
// extended=1
//
// https://vk.com/dev/messages.getConversationsById
func (vk *VK) MessagesGetConversationsByIDExtended(params map[string]string) (response MessagesGetConversationsByIDExtendedResponse, vkErr Error) {
	params["extended"] = "1"
	vk.RequestUnmarshal("messages.getConversationsById", params, &response, &vkErr)
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
//
// https://vk.com/dev/messages.getHistory
func (vk *VK) MessagesGetHistory(params map[string]string) (response MessagesGetHistoryResponse, vkErr Error) {
	vk.RequestUnmarshal("messages.getHistory", params, &response, &vkErr)
	return
}

// MessagesGetHistoryAttachmentsResponse struct
type MessagesGetHistoryAttachmentsResponse struct {
	Items    []object.MessagesHistoryAttachment `json:"items"`
	NextFrom string                             `json:"next_from"`
}

// MessagesGetHistoryAttachments returns media files from the dialog or group chat.
//
// https://vk.com/dev/messages.getHistoryAttachments
func (vk *VK) MessagesGetHistoryAttachments(params map[string]string) (response MessagesGetHistoryAttachmentsResponse, vkErr Error) {
	vk.RequestUnmarshal("messages.getHistoryAttachments", params, &response, &vkErr)
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
//
// https://vk.com/dev/messages.getImportantMessages
func (vk *VK) MessagesGetImportantMessages(params map[string]string) (response MessagesGetImportantMessagesResponse, vkErr Error) {
	vk.RequestUnmarshal("messages.getImportantMessages", params, &response, &vkErr)
	return
}

// MessagesGetInviteLinkResponse struct
type MessagesGetInviteLinkResponse struct {
	Link string `json:"link"`
}

// MessagesGetInviteLink receives a link to invite a user to the chat.
//
// https://vk.com/dev/messages.getInviteLink
func (vk *VK) MessagesGetInviteLink(params map[string]string) (response MessagesGetInviteLinkResponse, vkErr Error) {
	vk.RequestUnmarshal("messages.getInviteLink", params, &response, &vkErr)
	return
}

// MessagesGetLastActivityResponse struct
type MessagesGetLastActivityResponse object.MessagesLastActivity

// MessagesGetLastActivity returns a user's current status and date of last activity.
//
// https://vk.com/dev/messages.getLastActivity
func (vk *VK) MessagesGetLastActivity(params map[string]string) (response MessagesGetLastActivityResponse, vkErr Error) {
	vk.RequestUnmarshal("messages.getLastActivity", params, &response, &vkErr)
	return
}

// MessagesGetLongPollHistoryResponse struct
type MessagesGetLongPollHistoryResponse struct {
	History  [][]int              `json:"history"`
	Groups   []object.GroupsGroup `json:"groups"`
	Messages struct {
		Count int                      `json:"count"`
		Items []object.MessagesMessage `json:"items"`
	} `json:"messages"`
	Profiles []object.UsersUser `json:"profiles"`
	// Chats struct {} `json:"chats"`
	NewPTS        int                           `json:"new_pts"`
	More          bool                          `json:"chats"`
	Conversations []object.MessagesConversation `json:"conversations"`
}

// MessagesGetLongPollHistory returns updates in user's private messages.
//
// https://vk.com/dev/messages.getLongPollHistory
func (vk *VK) MessagesGetLongPollHistory(params map[string]string) (response MessagesGetLongPollHistoryResponse, vkErr Error) {
	vk.RequestUnmarshal("messages.getLongPollHistory", params, &response, &vkErr)
	return
}

// MessagesGetLongPollServerResponse struct
type MessagesGetLongPollServerResponse object.MessagesLongpollParams

// MessagesGetLongPollServer returns data required for connection to a Long Poll server.
//
// https://vk.com/dev/messages.getLongPollServer
func (vk *VK) MessagesGetLongPollServer(params map[string]string) (response MessagesGetLongPollServerResponse, vkErr Error) {
	vk.RequestUnmarshal("messages.getLongPollServer", params, &response, &vkErr)
	return
}

// MessagesIsMessagesFromGroupAllowedResponse struct
type MessagesIsMessagesFromGroupAllowedResponse struct {
	IsAllowed int `json:"is_allowed"`
}

// MessagesIsMessagesFromGroupAllowed returns information whether
// sending messages from the community to current user is allowed.
//
// https://vk.com/dev/messages.isMessagesFromGroupAllowed
func (vk *VK) MessagesIsMessagesFromGroupAllowed(params map[string]string) (response MessagesIsMessagesFromGroupAllowedResponse, vkErr Error) {
	vk.RequestUnmarshal("messages.isMessagesFromGroupAllowed", params, &response, &vkErr)
	return
}

// MessagesJoinChatByInviteLinkResponse struct
type MessagesJoinChatByInviteLinkResponse struct {
	ChatID int `json:"chat_id"`
}

// MessagesJoinChatByInviteLink allows to enter the chat by the invitation link.
//
// https://vk.com/dev/messages.joinChatByInviteLink
func (vk *VK) MessagesJoinChatByInviteLink(params map[string]string) (response MessagesJoinChatByInviteLinkResponse, vkErr Error) {
	vk.RequestUnmarshal("messages.joinChatByInviteLink", params, &response, &vkErr)
	return
}

// MessagesMarkAsAnsweredConversationResponse struct
type MessagesMarkAsAnsweredConversationResponse int

// MessagesMarkAsAnsweredConversation messages.markAsAnsweredConversation
//
// https://vk.com/dev/messages.markAsAnsweredConversation
func (vk *VK) MessagesMarkAsAnsweredConversation(params map[string]string) (response MessagesMarkAsAnsweredConversationResponse, vkErr Error) {
	vk.RequestUnmarshal("messages.markAsAnsweredConversation", params, &response, &vkErr)
	return
}

// MessagesMarkAsImportantResponse struct
type MessagesMarkAsImportantResponse []int

// MessagesMarkAsImportant marks and un marks messages as important (starred).
//
// https://vk.com/dev/messages.markAsImportant
func (vk *VK) MessagesMarkAsImportant(params map[string]string) (response MessagesMarkAsImportantResponse, vkErr Error) {
	vk.RequestUnmarshal("messages.markAsImportant", params, &response, &vkErr)
	return
}

// MessagesMarkAsImportantConversationResponse struct
type MessagesMarkAsImportantConversationResponse int

// MessagesMarkAsImportantConversation messages.markAsImportantConversation
//
// https://vk.com/dev/messages.markAsImportantConversation
func (vk *VK) MessagesMarkAsImportantConversation(params map[string]string) (response MessagesMarkAsImportantConversationResponse, vkErr Error) {
	vk.RequestUnmarshal("messages.markAsImportantConversation", params, &response, &vkErr)
	return
}

// MessagesMarkAsReadResponse struct
type MessagesMarkAsReadResponse int

// MessagesMarkAsRead marks messages as read.
//
// https://vk.com/dev/messages.markAsRead
func (vk *VK) MessagesMarkAsRead(params map[string]string) (response MessagesMarkAsReadResponse, vkErr Error) {
	vk.RequestUnmarshal("messages.markAsRead", params, &response, &vkErr)
	return
}

// MessagesPinResponse struct
type MessagesPinResponse object.MessagesMessage

// MessagesPin messages.pin
//
// https://vk.com/dev/messages.pin
func (vk *VK) MessagesPin(params map[string]string) (response MessagesPinResponse, vkErr Error) {
	vk.RequestUnmarshal("messages.pin", params, &response, &vkErr)
	return
}

// MessagesRemoveChatUserResponse struct
type MessagesRemoveChatUserResponse int

// MessagesRemoveChatUser Allows the current user to leave a chat or, if the current user started the chat, allows the user to remove another user from the chat.
//
// https://vk.com/dev/messages.removeChatUser
func (vk *VK) MessagesRemoveChatUser(params map[string]string) (response MessagesRemoveChatUserResponse, vkErr Error) {
	vk.RequestUnmarshal("messages.removeChatUser", params, &response, &vkErr)
	return
}

// MessagesRestoreResponse struct
type MessagesRestoreResponse int

// MessagesRestore restores a deleted message.
//
// https://vk.com/dev/messages.restore
func (vk *VK) MessagesRestore(params map[string]string) (response MessagesRestoreResponse, vkErr Error) {
	vk.RequestUnmarshal("messages.restore", params, &response, &vkErr)
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
//
// https://vk.com/dev/messages.search
func (vk *VK) MessagesSearch(params map[string]string) (response MessagesSearchResponse, vkErr Error) {
	vk.RequestUnmarshal("messages.search", params, &response, &vkErr)
	return
}

// MessagesSearchConversationsResponse struct
type MessagesSearchConversationsResponse struct {
	Count    int                           `json:"count"`
	Items    []object.MessagesConversation `json:"items"`
	Profiles []object.UsersUser            `json:"profiles"`
	Groups   []object.GroupsGroup          `json:"groups"`
}

// MessagesSearchConversations returns a list of conversations that match search criteria.
//
// https://vk.com/dev/messages.searchConversations
func (vk *VK) MessagesSearchConversations(params map[string]string) (response MessagesSearchConversationsResponse, vkErr Error) {
	vk.RequestUnmarshal("messages.searchConversations", params, &response, &vkErr)
	return
}

// MessagesSendResponse struct
type MessagesSendResponse int

// MessagesSend Sends a message
//
// https://vk.com/dev/messages.send
func (vk *VK) MessagesSend(params map[string]string) (response MessagesSendResponse, vkErr Error) {
	params["user_ids"] = ""
	vk.RequestUnmarshal("messages.send", params, &response, &vkErr)
	return
}

// MessagesSendUserIDsResponse struct
type MessagesSendUserIDsResponse []struct {
	PeerID    int `json:"peer_id"`
	MessageID int `json:"message_id"`
	Error     struct {
		Code        int    `json:"code"`
		Description string `json:"description"`
	} `json:"error"`
}

// MessagesSendUserIDs Sends a message
//
// need user_ids
//
// https://vk.com/dev/messages.send
func (vk *VK) MessagesSendUserIDs(params map[string]string) (response MessagesSendUserIDsResponse, vkErr Error) {
	vk.RequestUnmarshal("messages.send", params, &response, &vkErr)
	return
}

// MessagesSendStickerResponse struct
type MessagesSendStickerResponse int

// MessagesSendSticker Sends a message
//
// https://vk.com/dev/messages.sendSticker
func (vk *VK) MessagesSendSticker(params map[string]string) (response MessagesSendStickerResponse, vkErr Error) {
	params["user_ids"] = ""
	vk.RequestUnmarshal("messages.sendSticker", params, &response, &vkErr)
	return
}

// MessagesSetActivityResponse struct
type MessagesSetActivityResponse int

// MessagesSetActivity changes the status of a user as typing in a conversation.
//
// https://vk.com/dev/messages.setActivity
func (vk *VK) MessagesSetActivity(params map[string]string) (response MessagesSetActivityResponse, vkErr Error) {
	vk.RequestUnmarshal("messages.setActivity", params, &response, &vkErr)
	return
}

// MessagesSetChatPhotoResponse struct
type MessagesSetChatPhotoResponse struct {
	MessageID int                 `json:"message_id"`
	Chat      object.MessagesChat `json:"chat"`
}

// MessagesSetChatPhoto sets a previously-uploaded picture as the cover picture of a chat.
//
// https://vk.com/dev/messages.setChatPhoto
func (vk *VK) MessagesSetChatPhoto(params map[string]string) (response MessagesSetChatPhotoResponse, vkErr Error) {
	vk.RequestUnmarshal("messages.setChatPhoto", params, &response, &vkErr)
	return
}

// MessagesUnpinResponse struct
type MessagesUnpinResponse int

// MessagesUnpin messages.unpin
//
// https://vk.com/dev/messages.unpin
func (vk *VK) MessagesUnpin(params map[string]string) (response MessagesUnpinResponse, vkErr Error) {
	vk.RequestUnmarshal("messages.unpin", params, &response, &vkErr)
	return
}
