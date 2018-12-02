package vksdk

// MessagesAddChatUserResponse struct
type MessagesAddChatUserResponse struct{}

// TODO messages.addChatUser Adds a new user to a chat.

// MessagesAllowMessagesFromGroupResponse struct
type MessagesAllowMessagesFromGroupResponse struct{}

// TODO messages.allowMessagesFromGroup Allows sending messages from community to the current user.

// MessagesCreateChatResponse struct
type MessagesCreateChatResponse struct{}

// TODO messages.createChat Creates a chat with several participants.

// MessagesDeleteResponse struct
type MessagesDeleteResponse struct{}

// TODO messages.delete Deletes one or more messages.

// MessagesDeleteChatPhotoResponse struct
type MessagesDeleteChatPhotoResponse struct{}

// TODO messages.deleteChatPhoto Deletes a chat's cover picture.

// MessagesDeleteConversationResponse struct
type MessagesDeleteConversationResponse struct{}

// TODO messages.deleteConversation Deletes private messages in a conversation.

// MessagesDenyMessagesFromGroupResponse struct
type MessagesDenyMessagesFromGroupResponse struct{}

// TODO messages.denyMessagesFromGroup Denies sending message from community to the current user.

// MessagesEditResponse struct
type MessagesEditResponse struct{}

// TODO messages.edit Edits the message.

// MessagesEditChatResponse struct
type MessagesEditChatResponse struct{}

// TODO messages.editChat Edits the title of a chat.

// MessagesGetByConversationMessageIDResponse struct
type MessagesGetByConversationMessageIDResponse struct{}

// TODO messages.getByConversationMessageId

// MessagesGetByIDResponse struct
type MessagesGetByIDResponse struct{}

// TODO messages.getById Returns messages by their IDs.

// MessagesGetChatResponse struct
type MessagesGetChatResponse struct{}

// TODO messages.getChat Returns information about a chat.

// MessagesGetChatPreviewResponse struct
type MessagesGetChatPreviewResponse struct{}

// TODO messages.getChatPreview Allows to receive chat preview by the invitation link.

// MessagesGetConversationMembersResponse struct
type MessagesGetConversationMembersResponse struct{}

// TODO messages.getConversationMembers Returns a list of IDs of users participating in a conversation.

// MessagesGetConversationsResponse struct
type MessagesGetConversationsResponse struct{}

// TODO messages.getConversations Returns a list of conversations.

// MessagesGetConversationsByIDResponse struct
type MessagesGetConversationsByIDResponse struct{}

// TODO messages.getConversationsById Returns conversations by their IDs.

// MessagesGetHistoryResponse struct
type MessagesGetHistoryResponse struct{}

// TODO messages.getHistory Returns message history for the specified user or group chat.

// MessagesGetHistoryAttachmentsResponse struct
type MessagesGetHistoryAttachmentsResponse struct{}

// TODO messages.getHistoryAttachments Returns media files from the dialog or group chat.

// MessagesGetImportantMessagesResponse struct
type MessagesGetImportantMessagesResponse struct{}

// TODO messages.getImportantMessages

// MessagesGetInviteLinkResponse struct
type MessagesGetInviteLinkResponse struct{}

// TODO messages.getInviteLink Receives a link to invite a user to the chat.

// MessagesGetLastActivityResponse struct
type MessagesGetLastActivityResponse struct{}

// TODO messages.getLastActivity Returns a user's current status and date of last activity.

// MessagesGetLongPollHistoryResponse struct
type MessagesGetLongPollHistoryResponse struct{}

// TODO messages.getLongPollHistory Returns updates in user's private messages.

// MessagesGetLongPollServerResponse struct
type MessagesGetLongPollServerResponse struct{}

// TODO messages.getLongPollServer Returns data required for connection to a Long Poll server.

// MessagesIsMessagesFromGroupAllowedResponse struct
type MessagesIsMessagesFromGroupAllowedResponse struct{}

// TODO messages.isMessagesFromGroupAllowed Returns information whether sending messages from the community to current user is allowed.

// MessagesJoinChatByInviteLinkResponse struct
type MessagesJoinChatByInviteLinkResponse struct{}

// TODO messages.joinChatByInviteLink Allows to enter the chat by the invitation link.

// MessagesMarkAsAnsweredConversationResponse struct
type MessagesMarkAsAnsweredConversationResponse struct{}

// TODO messages.markAsAnsweredConversation

// MessagesMarkAsImportantResponse struct
type MessagesMarkAsImportantResponse struct{}

// TODO messages.markAsImportant Marks and un marks messages as important (starred).

// MessagesMarkAsImportantConversationResponse struct
type MessagesMarkAsImportantConversationResponse struct{}

// TODO messages.markAsImportantConversation

// MessagesMarkAsReadResponse struct
type MessagesMarkAsReadResponse struct{}

// TODO messages.markAsRead Marks messages as read.

// MessagesPinResponse struct
type MessagesPinResponse struct{}

// TODO messages.pin

// MessagesRemoveChatUserResponse struct
type MessagesRemoveChatUserResponse struct{}

// TODO messages.removeChatUser Allows the current user to leave a chat or, if the current user started the chat, allows the user to remove another user from the chat.

// MessagesRestoreResponse struct
type MessagesRestoreResponse struct{}

// TODO messages.restore Restores a deleted message.

// MessagesSearchResponse struct
type MessagesSearchResponse struct{}

// TODO messages.search Returns a list of the current user's private messages that match search criteria.

// MessagesSearchConversationsResponse struct
type MessagesSearchConversationsResponse struct{}

// TODO messages.searchConversations Returns a list of conversations that match search criteria.

// MessagesSendResponse struct
type MessagesSendResponse struct{}

// MessagesSend Sends a message
// TODO messages.send Sends a message.
func (vk *VK) MessagesSend(params map[string]string) error {
	const method = "messages.send"
	// TODO if user_ids in params

	_, err := vk.Request(method, params)
	if err != nil {
		return err
	}

	return nil
}

// MessagesSetActivityResponse struct
type MessagesSetActivityResponse struct{}

// TODO messages.setActivity Changes the status of a user as typing in a conversation.

// MessagesSetChatPhotoResponse struct
type MessagesSetChatPhotoResponse struct{}

// TODO messages.setChatPhoto Sets a previously-uploaded picture as the cover picture of a chat.

// MessagesUnpinResponse struct
type MessagesUnpinResponse struct{}

// TODO messages.unpin
