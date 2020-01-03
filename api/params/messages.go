package params // import "github.com/SevereCloud/vksdk/api/params"

import (
	"github.com/SevereCloud/vksdk/api"
)

// MessagesAddChatUserBulder builder
//
// Adds a new user to a chat.
//
// https://vk.com/dev/messages.addChatUser
type MessagesAddChatUserBulder struct {
	api.Params
}

// NewMessagesAddChatUserBulder func
func NewMessagesAddChatUserBulder() *MessagesAddChatUserBulder {
	return &MessagesAddChatUserBulder{api.Params{}}
}

// ChatID Chat ID.
func (b *MessagesAddChatUserBulder) ChatID(v int) {
	b.Params["chat_id"] = v
}

// UserID ID of the user to be added to the chat.
func (b *MessagesAddChatUserBulder) UserID(v int) {
	b.Params["user_id"] = v
}

// MessagesAllowMessagesFromGroupBulder builder
//
// Allows sending messages from community to the current user.
//
// https://vk.com/dev/messages.allowMessagesFromGroup
type MessagesAllowMessagesFromGroupBulder struct {
	api.Params
}

// NewMessagesAllowMessagesFromGroupBulder func
func NewMessagesAllowMessagesFromGroupBulder() *MessagesAllowMessagesFromGroupBulder {
	return &MessagesAllowMessagesFromGroupBulder{api.Params{}}
}

// GroupID Group ID.
func (b *MessagesAllowMessagesFromGroupBulder) GroupID(v int) {
	b.Params["group_id"] = v
}

// Key parameter
func (b *MessagesAllowMessagesFromGroupBulder) Key(v string) {
	b.Params["key"] = v
}

// MessagesCreateChatBulder builder
//
// Creates a chat with several participants.
//
// https://vk.com/dev/messages.createChat
type MessagesCreateChatBulder struct {
	api.Params
}

// NewMessagesCreateChatBulder func
func NewMessagesCreateChatBulder() *MessagesCreateChatBulder {
	return &MessagesCreateChatBulder{api.Params{}}
}

// UserIDs IDs of the users to be added to the chat.
func (b *MessagesCreateChatBulder) UserIDs(v []int) {
	b.Params["user_ids"] = v
}

// Title Chat title.
func (b *MessagesCreateChatBulder) Title(v string) {
	b.Params["title"] = v
}

// MessagesDeleteBulder builder
//
// Deletes one or more messages.
//
// https://vk.com/dev/messages.delete
type MessagesDeleteBulder struct {
	api.Params
}

// NewMessagesDeleteBulder func
func NewMessagesDeleteBulder() *MessagesDeleteBulder {
	return &MessagesDeleteBulder{api.Params{}}
}

// MessageIDs Message IDs.
func (b *MessagesDeleteBulder) MessageIDs(v []int) {
	b.Params["message_ids"] = v
}

// Spam '1' — to mark message as spam.
func (b *MessagesDeleteBulder) Spam(v bool) {
	b.Params["spam"] = v
}

// GroupID Group ID (for group messages with user access token)
func (b *MessagesDeleteBulder) GroupID(v int) {
	b.Params["group_id"] = v
}

// DeleteForAll '1' — delete message for for all.
func (b *MessagesDeleteBulder) DeleteForAll(v bool) {
	b.Params["delete_for_all"] = v
}

// MessagesDeleteChatPhotoBulder builder
//
// Deletes a chat's cover picture.
//
// https://vk.com/dev/messages.deleteChatPhoto
type MessagesDeleteChatPhotoBulder struct {
	api.Params
}

// NewMessagesDeleteChatPhotoBulder func
func NewMessagesDeleteChatPhotoBulder() *MessagesDeleteChatPhotoBulder {
	return &MessagesDeleteChatPhotoBulder{api.Params{}}
}

// ChatID Chat ID.
func (b *MessagesDeleteChatPhotoBulder) ChatID(v int) {
	b.Params["chat_id"] = v
}

// GroupID parameter
func (b *MessagesDeleteChatPhotoBulder) GroupID(v int) {
	b.Params["group_id"] = v
}

// MessagesDeleteConversationBulder builder
//
// Deletes all private messages in a conversation.
//
// https://vk.com/dev/messages.deleteConversation
type MessagesDeleteConversationBulder struct {
	api.Params
}

// NewMessagesDeleteConversationBulder func
func NewMessagesDeleteConversationBulder() *MessagesDeleteConversationBulder {
	return &MessagesDeleteConversationBulder{api.Params{}}
}

// UserID User ID. To clear a chat history use 'chat_id'
func (b *MessagesDeleteConversationBulder) UserID(v int) {
	b.Params["user_id"] = v
}

// PeerID Destination ID. "For user: 'User ID', e.g. '12345'. For chat: '2000000000' + 'chat_id', e.g. '2000000001'. For community: '- community ID', e.g. '-12345'. "
func (b *MessagesDeleteConversationBulder) PeerID(v int) {
	b.Params["peer_id"] = v
}

// GroupID Group ID (for group messages with user access token)
func (b *MessagesDeleteConversationBulder) GroupID(v int) {
	b.Params["group_id"] = v
}

// MessagesDenyMessagesFromGroupBulder builder
//
// Denies sending message from community to the current user.
//
// https://vk.com/dev/messages.denyMessagesFromGroup
type MessagesDenyMessagesFromGroupBulder struct {
	api.Params
}

// NewMessagesDenyMessagesFromGroupBulder func
func NewMessagesDenyMessagesFromGroupBulder() *MessagesDenyMessagesFromGroupBulder {
	return &MessagesDenyMessagesFromGroupBulder{api.Params{}}
}

// GroupID Group ID.
func (b *MessagesDenyMessagesFromGroupBulder) GroupID(v int) {
	b.Params["group_id"] = v
}

// MessagesEditBulder builder
//
// Edits the message.
//
// https://vk.com/dev/messages.edit
type MessagesEditBulder struct {
	api.Params
}

// NewMessagesEditBulder func
func NewMessagesEditBulder() *MessagesEditBulder {
	return &MessagesEditBulder{api.Params{}}
}

// PeerID Destination ID. "For user: 'User ID', e.g. '12345'. For chat: '2000000000' + 'chat_id', e.g. '2000000001'. For community: '- community ID', e.g. '-12345'. "
func (b *MessagesEditBulder) PeerID(v int) {
	b.Params["peer_id"] = v
}

// Message (Required if 'attachments' is not set.) Text of the message.
func (b *MessagesEditBulder) Message(v string) {
	b.Params["message"] = v
}

// MessageID parameter
func (b *MessagesEditBulder) MessageID(v int) {
	b.Params["message_id"] = v
}

// Lat Geographical latitude of a check-in, in degrees (from -90 to 90).
func (b *MessagesEditBulder) Lat(v float64) {
	b.Params["lat"] = v
}

// Long Geographical longitude of a check-in, in degrees (from -180 to 180).
func (b *MessagesEditBulder) Long(v float64) {
	b.Params["long"] = v
}

// Attachment (Required if 'message' is not set.) List of objects attached to the message, separated by commas, in the following format: "<owner_id>_<media_id>", '' — Type of media attachment: 'photo' — photo, 'video' — video, 'audio' — audio, 'doc' — document, 'wall' — wall post, '<owner_id>' — ID of the media attachment owner. '<media_id>' — media attachment ID. Example: "photo100172_166443618"
func (b *MessagesEditBulder) Attachment(v string) {
	b.Params["attachment"] = v
}

// KeepForwardMessages '1' — to keep forwarded, messages.
func (b *MessagesEditBulder) KeepForwardMessages(v bool) {
	b.Params["keep_forward_messages"] = v
}

// KeepSnippets '1' — to keep attached snippets.
func (b *MessagesEditBulder) KeepSnippets(v bool) {
	b.Params["keep_snippets"] = v
}

// GroupID Group ID (for group messages with user access token)
func (b *MessagesEditBulder) GroupID(v int) {
	b.Params["group_id"] = v
}

// DontParseLinks parameter
func (b *MessagesEditBulder) DontParseLinks(v bool) {
	b.Params["dont_parse_links"] = v
}

// MessagesEditChatBulder builder
//
// Edits the title of a chat.
//
// https://vk.com/dev/messages.editChat
type MessagesEditChatBulder struct {
	api.Params
}

// NewMessagesEditChatBulder func
func NewMessagesEditChatBulder() *MessagesEditChatBulder {
	return &MessagesEditChatBulder{api.Params{}}
}

// ChatID Chat ID.
func (b *MessagesEditChatBulder) ChatID(v int) {
	b.Params["chat_id"] = v
}

// Title New title of the chat.
func (b *MessagesEditChatBulder) Title(v string) {
	b.Params["title"] = v
}

// MessagesGetByConversationMessageIDBulder builder
//
// Returns messages by their IDs within the conversation.
//
// https://vk.com/dev/messages.getByConversationMessageId
type MessagesGetByConversationMessageIDBulder struct {
	api.Params
}

// NewMessagesGetByConversationMessageIDBulder func
func NewMessagesGetByConversationMessageIDBulder() *MessagesGetByConversationMessageIDBulder {
	return &MessagesGetByConversationMessageIDBulder{api.Params{}}
}

// PeerID Destination ID. "For user: 'User ID', e.g. '12345'. For chat: '2000000000' + 'chat_id', e.g. '2000000001'. For community: '- community ID', e.g. '-12345'. "
func (b *MessagesGetByConversationMessageIDBulder) PeerID(v int) {
	b.Params["peer_id"] = v
}

// ConversationMessageIDs Conversation message IDs.
func (b *MessagesGetByConversationMessageIDBulder) ConversationMessageIDs(v []int) {
	b.Params["conversation_message_ids"] = v
}

// Extended Information whether the response should be extended
func (b *MessagesGetByConversationMessageIDBulder) Extended(v bool) {
	b.Params["extended"] = v
}

// Fields Profile fields to return.
func (b *MessagesGetByConversationMessageIDBulder) Fields(v []string) {
	b.Params["fields"] = v
}

// GroupID Group ID (for group messages with group access token)
func (b *MessagesGetByConversationMessageIDBulder) GroupID(v int) {
	b.Params["group_id"] = v
}

// MessagesGetByIDBulder builder
//
// Returns messages by their IDs.
//
// https://vk.com/dev/messages.getById
type MessagesGetByIDBulder struct {
	api.Params
}

// NewMessagesGetByIDBulder func
func NewMessagesGetByIDBulder() *MessagesGetByIDBulder {
	return &MessagesGetByIDBulder{api.Params{}}
}

// MessageIDs Message IDs.
func (b *MessagesGetByIDBulder) MessageIDs(v []int) {
	b.Params["message_ids"] = v
}

// PreviewLength Number of characters after which to truncate a previewed message. To preview the full message, specify '0'. "NOTE: Messages are not truncated by default. Messages are truncated by words."
func (b *MessagesGetByIDBulder) PreviewLength(v int) {
	b.Params["preview_length"] = v
}

// Extended Information whether the response should be extended
func (b *MessagesGetByIDBulder) Extended(v bool) {
	b.Params["extended"] = v
}

// Fields Profile fields to return.
func (b *MessagesGetByIDBulder) Fields(v []string) {
	b.Params["fields"] = v
}

// GroupID Group ID (for group messages with group access token)
func (b *MessagesGetByIDBulder) GroupID(v int) {
	b.Params["group_id"] = v
}

// MessagesGetChatPreviewBulder builder
//
// https://vk.com/dev/messages.getChatPreview
type MessagesGetChatPreviewBulder struct {
	api.Params
}

// NewMessagesGetChatPreviewBulder func
func NewMessagesGetChatPreviewBulder() *MessagesGetChatPreviewBulder {
	return &MessagesGetChatPreviewBulder{api.Params{}}
}

// PeerID parameter
func (b *MessagesGetChatPreviewBulder) PeerID(v int) {
	b.Params["peer_id"] = v
}

// Link Invitation link.
func (b *MessagesGetChatPreviewBulder) Link(v string) {
	b.Params["link"] = v
}

// Fields Profile fields to return.
func (b *MessagesGetChatPreviewBulder) Fields(v []string) {
	b.Params["fields"] = v
}

// MessagesGetConversationMembersBulder builder
//
// Returns a list of IDs of users participating in a chat.
//
// https://vk.com/dev/messages.getConversationMembers
type MessagesGetConversationMembersBulder struct {
	api.Params
}

// NewMessagesGetConversationMembersBulder func
func NewMessagesGetConversationMembersBulder() *MessagesGetConversationMembersBulder {
	return &MessagesGetConversationMembersBulder{api.Params{}}
}

// PeerID Peer ID.
func (b *MessagesGetConversationMembersBulder) PeerID(v int) {
	b.Params["peer_id"] = v
}

// Fields Profile fields to return.
func (b *MessagesGetConversationMembersBulder) Fields(v []string) {
	b.Params["fields"] = v
}

// GroupID Group ID (for group messages with group access token)
func (b *MessagesGetConversationMembersBulder) GroupID(v int) {
	b.Params["group_id"] = v
}

// MessagesGetConversationsBulder builder
//
// Returns a list of the current user's conversations.
//
// https://vk.com/dev/messages.getConversations
type MessagesGetConversationsBulder struct {
	api.Params
}

// NewMessagesGetConversationsBulder func
func NewMessagesGetConversationsBulder() *MessagesGetConversationsBulder {
	return &MessagesGetConversationsBulder{api.Params{}}
}

// Offset Offset needed to return a specific subset of conversations.
func (b *MessagesGetConversationsBulder) Offset(v int) {
	b.Params["offset"] = v
}

// Count Number of conversations to return.
func (b *MessagesGetConversationsBulder) Count(v int) {
	b.Params["count"] = v
}

// Filter Filter to apply: 'all' — all conversations, 'unread' — conversations with unread messages, 'important' — conversations, marked as important (only for community messages), 'unanswered' — conversations, marked as unanswered (only for community messages)
func (b *MessagesGetConversationsBulder) Filter(v string) {
	b.Params["filter"] = v
}

// Extended '1' — return extra information about users and communities
func (b *MessagesGetConversationsBulder) Extended(v bool) {
	b.Params["extended"] = v
}

// StartMessageID ID of the message from what to return dialogs.
func (b *MessagesGetConversationsBulder) StartMessageID(v int) {
	b.Params["start_message_id"] = v
}

// Fields Profile and communities fields to return.
func (b *MessagesGetConversationsBulder) Fields(v []string) {
	b.Params["fields"] = v
}

// GroupID Group ID (for group messages with group access token)
func (b *MessagesGetConversationsBulder) GroupID(v int) {
	b.Params["group_id"] = v
}

// MessagesGetConversationsByIDBulder builder
//
// Returns conversations by their IDs
//
// https://vk.com/dev/messages.getConversationsById
type MessagesGetConversationsByIDBulder struct {
	api.Params
}

// NewMessagesGetConversationsByIDBulder func
func NewMessagesGetConversationsByIDBulder() *MessagesGetConversationsByIDBulder {
	return &MessagesGetConversationsByIDBulder{api.Params{}}
}

// PeerIDs Destination IDs. "For user: 'User ID', e.g. '12345'. For chat: '2000000000' + 'chat_id', e.g. '2000000001'. For community: '- community ID', e.g. '-12345'. "
func (b *MessagesGetConversationsByIDBulder) PeerIDs(v []int) {
	b.Params["peer_ids"] = v
}

// Extended Return extended properties
func (b *MessagesGetConversationsByIDBulder) Extended(v bool) {
	b.Params["extended"] = v
}

// Fields Profile and communities fields to return.
func (b *MessagesGetConversationsByIDBulder) Fields(v []string) {
	b.Params["fields"] = v
}

// GroupID Group ID (for group messages with group access token)
func (b *MessagesGetConversationsByIDBulder) GroupID(v int) {
	b.Params["group_id"] = v
}

// MessagesGetHistoryBulder builder
//
// Returns message history for the specified user or group chat.
//
// https://vk.com/dev/messages.getHistory
type MessagesGetHistoryBulder struct {
	api.Params
}

// NewMessagesGetHistoryBulder func
func NewMessagesGetHistoryBulder() *MessagesGetHistoryBulder {
	return &MessagesGetHistoryBulder{api.Params{}}
}

// Offset Offset needed to return a specific subset of messages.
func (b *MessagesGetHistoryBulder) Offset(v int) {
	b.Params["offset"] = v
}

// Count Number of messages to return.
func (b *MessagesGetHistoryBulder) Count(v int) {
	b.Params["count"] = v
}

// UserID ID of the user whose message history you want to return.
func (b *MessagesGetHistoryBulder) UserID(v int) {
	b.Params["user_id"] = v
}

// PeerID parameter
func (b *MessagesGetHistoryBulder) PeerID(v int) {
	b.Params["peer_id"] = v
}

// StartMessageID Starting message ID from which to return history.
func (b *MessagesGetHistoryBulder) StartMessageID(v int) {
	b.Params["start_message_id"] = v
}

// Rev Sort order: '1' — return messages in chronological order. '0' — return messages in reverse chronological order.
func (b *MessagesGetHistoryBulder) Rev(v int) {
	b.Params["rev"] = v
}

// Extended Information whether the response should be extended
func (b *MessagesGetHistoryBulder) Extended(v bool) {
	b.Params["extended"] = v
}

// Fields Profile fields to return.
func (b *MessagesGetHistoryBulder) Fields(v []string) {
	b.Params["fields"] = v
}

// GroupID Group ID (for group messages with group access token)
func (b *MessagesGetHistoryBulder) GroupID(v int) {
	b.Params["group_id"] = v
}

// MessagesGetHistoryAttachmentsBulder builder
//
// Returns media files from the dialog or group chat.
//
// https://vk.com/dev/messages.getHistoryAttachments
type MessagesGetHistoryAttachmentsBulder struct {
	api.Params
}

// NewMessagesGetHistoryAttachmentsBulder func
func NewMessagesGetHistoryAttachmentsBulder() *MessagesGetHistoryAttachmentsBulder {
	return &MessagesGetHistoryAttachmentsBulder{api.Params{}}
}

// PeerID Peer ID. ", For group chat: '2000000000 + chat ID' , , For community: '-community ID'"
func (b *MessagesGetHistoryAttachmentsBulder) PeerID(v int) {
	b.Params["peer_id"] = v
}

// MediaType Type of media files to return: *'photo',, *'video',, *'audio',, *'doc',, *'link'.,*'market'.,*'wall'.,*'share'
func (b *MessagesGetHistoryAttachmentsBulder) MediaType(v string) {
	b.Params["media_type"] = v
}

// StartFrom Message ID to start return results from.
func (b *MessagesGetHistoryAttachmentsBulder) StartFrom(v string) {
	b.Params["start_from"] = v
}

// Count Number of objects to return.
func (b *MessagesGetHistoryAttachmentsBulder) Count(v int) {
	b.Params["count"] = v
}

// PhotoSizes '1' — to return photo sizes in a
func (b *MessagesGetHistoryAttachmentsBulder) PhotoSizes(v bool) {
	b.Params["photo_sizes"] = v
}

// Fields Additional profile [vk.com/dev/fields|fields] to return.
func (b *MessagesGetHistoryAttachmentsBulder) Fields(v []string) {
	b.Params["fields"] = v
}

// GroupID Group ID (for group messages with group access token)
func (b *MessagesGetHistoryAttachmentsBulder) GroupID(v int) {
	b.Params["group_id"] = v
}

// PreserveOrder parameter
func (b *MessagesGetHistoryAttachmentsBulder) PreserveOrder(v bool) {
	b.Params["preserve_order"] = v
}

// MaxForwardsLevel parameter
func (b *MessagesGetHistoryAttachmentsBulder) MaxForwardsLevel(v int) {
	b.Params["max_forwards_level"] = v
}

// MessagesGetInviteLinkBulder builder
//
// https://vk.com/dev/messages.getInviteLink
type MessagesGetInviteLinkBulder struct {
	api.Params
}

// NewMessagesGetInviteLinkBulder func
func NewMessagesGetInviteLinkBulder() *MessagesGetInviteLinkBulder {
	return &MessagesGetInviteLinkBulder{api.Params{}}
}

// PeerID Destination ID.
func (b *MessagesGetInviteLinkBulder) PeerID(v int) {
	b.Params["peer_id"] = v
}

// Reset 1 — to generate new link (revoke previous), 0 — to return previous link.
func (b *MessagesGetInviteLinkBulder) Reset(v bool) {
	b.Params["reset"] = v
}

// GroupID Group ID
func (b *MessagesGetInviteLinkBulder) GroupID(v int) {
	b.Params["group_id"] = v
}

// MessagesGetLastActivityBulder builder
//
// Returns a user's current status and date of last activity.
//
// https://vk.com/dev/messages.getLastActivity
type MessagesGetLastActivityBulder struct {
	api.Params
}

// NewMessagesGetLastActivityBulder func
func NewMessagesGetLastActivityBulder() *MessagesGetLastActivityBulder {
	return &MessagesGetLastActivityBulder{api.Params{}}
}

// UserID User ID.
func (b *MessagesGetLastActivityBulder) UserID(v int) {
	b.Params["user_id"] = v
}

// MessagesGetLongPollHistoryBulder builder
//
// Returns updates in user's private messages.
//
// https://vk.com/dev/messages.getLongPollHistory
type MessagesGetLongPollHistoryBulder struct {
	api.Params
}

// NewMessagesGetLongPollHistoryBulder func
func NewMessagesGetLongPollHistoryBulder() *MessagesGetLongPollHistoryBulder {
	return &MessagesGetLongPollHistoryBulder{api.Params{}}
}

// Ts Last value of the 'ts' parameter returned from the Long Poll server or by using [vk.com/dev/messages.getLongPollHistory|messages.getLongPollHistory] method.
func (b *MessagesGetLongPollHistoryBulder) Ts(v int) {
	b.Params["ts"] = v
}

// Pts Lsat value of 'pts' parameter returned from the Long Poll server or by using [vk.com/dev/messages.getLongPollHistory|messages.getLongPollHistory] method.
func (b *MessagesGetLongPollHistoryBulder) Pts(v int) {
	b.Params["pts"] = v
}

// PreviewLength Number of characters after which to truncate a previewed message. To preview the full message, specify '0'. "NOTE: Messages are not truncated by default. Messages are truncated by words."
func (b *MessagesGetLongPollHistoryBulder) PreviewLength(v int) {
	b.Params["preview_length"] = v
}

// Onlines '1' — to return history with online users only.
func (b *MessagesGetLongPollHistoryBulder) Onlines(v bool) {
	b.Params["onlines"] = v
}

// Fields Additional profile [vk.com/dev/fields|fields] to return.
func (b *MessagesGetLongPollHistoryBulder) Fields(v []string) {
	b.Params["fields"] = v
}

// EventsLimit Maximum number of events to return.
func (b *MessagesGetLongPollHistoryBulder) EventsLimit(v int) {
	b.Params["events_limit"] = v
}

// MsgsLimit Maximum number of messages to return.
func (b *MessagesGetLongPollHistoryBulder) MsgsLimit(v int) {
	b.Params["msgs_limit"] = v
}

// MaxMsgID Maximum ID of the message among existing ones in the local copy. Both messages received with API methods (for example, , ), and data received from a Long Poll server (events with code 4) are taken into account.
func (b *MessagesGetLongPollHistoryBulder) MaxMsgID(v int) {
	b.Params["max_msg_id"] = v
}

// GroupID Group ID (for group messages with user access token)
func (b *MessagesGetLongPollHistoryBulder) GroupID(v int) {
	b.Params["group_id"] = v
}

// LpVersion parameter
func (b *MessagesGetLongPollHistoryBulder) LpVersion(v int) {
	b.Params["lp_version"] = v
}

// LastN parameter
func (b *MessagesGetLongPollHistoryBulder) LastN(v int) {
	b.Params["last_n"] = v
}

// Credentials parameter
func (b *MessagesGetLongPollHistoryBulder) Credentials(v bool) {
	b.Params["credentials"] = v
}

// MessagesGetLongPollServerBulder builder
//
// Returns data required for connection to a Long Poll server.
//
// https://vk.com/dev/messages.getLongPollServer
type MessagesGetLongPollServerBulder struct {
	api.Params
}

// NewMessagesGetLongPollServerBulder func
func NewMessagesGetLongPollServerBulder() *MessagesGetLongPollServerBulder {
	return &MessagesGetLongPollServerBulder{api.Params{}}
}

// NeedPts '1' — to return the 'pts' field, needed for the [vk.com/dev/messages.getLongPollHistory|messages.getLongPollHistory] method.
func (b *MessagesGetLongPollServerBulder) NeedPts(v bool) {
	b.Params["need_pts"] = v
}

// GroupID Group ID (for group messages with user access token)
func (b *MessagesGetLongPollServerBulder) GroupID(v int) {
	b.Params["group_id"] = v
}

// LpVersion Long poll version
func (b *MessagesGetLongPollServerBulder) LpVersion(v int) {
	b.Params["lp_version"] = v
}

// MessagesIsMessagesFromGroupAllowedBulder builder
//
// Returns information whether sending messages from the community to current user is allowed.
//
// https://vk.com/dev/messages.isMessagesFromGroupAllowed
type MessagesIsMessagesFromGroupAllowedBulder struct {
	api.Params
}

// NewMessagesIsMessagesFromGroupAllowedBulder func
func NewMessagesIsMessagesFromGroupAllowedBulder() *MessagesIsMessagesFromGroupAllowedBulder {
	return &MessagesIsMessagesFromGroupAllowedBulder{api.Params{}}
}

// GroupID Group ID.
func (b *MessagesIsMessagesFromGroupAllowedBulder) GroupID(v int) {
	b.Params["group_id"] = v
}

// UserID User ID.
func (b *MessagesIsMessagesFromGroupAllowedBulder) UserID(v int) {
	b.Params["user_id"] = v
}

// MessagesJoinChatByInviteLinkBulder builder
//
// https://vk.com/dev/messages.joinChatByInviteLink
type MessagesJoinChatByInviteLinkBulder struct {
	api.Params
}

// NewMessagesJoinChatByInviteLinkBulder func
func NewMessagesJoinChatByInviteLinkBulder() *MessagesJoinChatByInviteLinkBulder {
	return &MessagesJoinChatByInviteLinkBulder{api.Params{}}
}

// Link Invitation link.
func (b *MessagesJoinChatByInviteLinkBulder) Link(v string) {
	b.Params["link"] = v
}

// MessagesMarkAsAnsweredConversationBulder builder
//
// Marks and unmarks conversations as unanswered.
//
// https://vk.com/dev/messages.markAsAnsweredConversation
type MessagesMarkAsAnsweredConversationBulder struct {
	api.Params
}

// NewMessagesMarkAsAnsweredConversationBulder func
func NewMessagesMarkAsAnsweredConversationBulder() *MessagesMarkAsAnsweredConversationBulder {
	return &MessagesMarkAsAnsweredConversationBulder{api.Params{}}
}

// PeerID ID of conversation to mark as important.
func (b *MessagesMarkAsAnsweredConversationBulder) PeerID(v int) {
	b.Params["peer_id"] = v
}

// Answered '1' — to mark as answered, '0' — to remove the mark
func (b *MessagesMarkAsAnsweredConversationBulder) Answered(v bool) {
	b.Params["answered"] = v
}

// GroupID Group ID (for group messages with group access token)
func (b *MessagesMarkAsAnsweredConversationBulder) GroupID(v int) {
	b.Params["group_id"] = v
}

// MessagesMarkAsImportantBulder builder
//
// Marks and unmarks messages as important (starred).
//
// https://vk.com/dev/messages.markAsImportant
type MessagesMarkAsImportantBulder struct {
	api.Params
}

// NewMessagesMarkAsImportantBulder func
func NewMessagesMarkAsImportantBulder() *MessagesMarkAsImportantBulder {
	return &MessagesMarkAsImportantBulder{api.Params{}}
}

// MessageIDs IDs of messages to mark as important.
func (b *MessagesMarkAsImportantBulder) MessageIDs(v []int) {
	b.Params["message_ids"] = v
}

// Important '1' — to add a star (mark as important), '0' — to remove the star
func (b *MessagesMarkAsImportantBulder) Important(v int) {
	b.Params["important"] = v
}

// MessagesMarkAsImportantConversationBulder builder
//
// Marks and unmarks conversations as important.
//
// https://vk.com/dev/messages.markAsImportantConversation
type MessagesMarkAsImportantConversationBulder struct {
	api.Params
}

// NewMessagesMarkAsImportantConversationBulder func
func NewMessagesMarkAsImportantConversationBulder() *MessagesMarkAsImportantConversationBulder {
	return &MessagesMarkAsImportantConversationBulder{api.Params{}}
}

// PeerID ID of conversation to mark as important.
func (b *MessagesMarkAsImportantConversationBulder) PeerID(v int) {
	b.Params["peer_id"] = v
}

// Important '1' — to add a star (mark as important), '0' — to remove the star
func (b *MessagesMarkAsImportantConversationBulder) Important(v bool) {
	b.Params["important"] = v
}

// GroupID Group ID (for group messages with group access token)
func (b *MessagesMarkAsImportantConversationBulder) GroupID(v int) {
	b.Params["group_id"] = v
}

// MessagesMarkAsReadBulder builder
//
// Marks messages as read.
//
// https://vk.com/dev/messages.markAsRead
type MessagesMarkAsReadBulder struct {
	api.Params
}

// NewMessagesMarkAsReadBulder func
func NewMessagesMarkAsReadBulder() *MessagesMarkAsReadBulder {
	return &MessagesMarkAsReadBulder{api.Params{}}
}

// MessageIDs IDs of messages to mark as read.
func (b *MessagesMarkAsReadBulder) MessageIDs(v []int) {
	b.Params["message_ids"] = v
}

// PeerID Destination ID. "For user: 'User ID', e.g. '12345'. For chat: '2000000000' + 'chat_id', e.g. '2000000001'. For community: '- community ID', e.g. '-12345'. "
func (b *MessagesMarkAsReadBulder) PeerID(v int) {
	b.Params["peer_id"] = v
}

// StartMessageID Message ID to start from.
func (b *MessagesMarkAsReadBulder) StartMessageID(v int) {
	b.Params["start_message_id"] = v
}

// GroupID Group ID (for group messages with user access token)
func (b *MessagesMarkAsReadBulder) GroupID(v int) {
	b.Params["group_id"] = v
}

// MessagesPinBulder builder
//
// Pin a message.
//
// https://vk.com/dev/messages.pin
type MessagesPinBulder struct {
	api.Params
}

// NewMessagesPinBulder func
func NewMessagesPinBulder() *MessagesPinBulder {
	return &MessagesPinBulder{api.Params{}}
}

// PeerID Destination ID. "For user: 'User ID', e.g. '12345'. For chat: '2000000000' + 'Chat ID', e.g. '2000000001'. For community: '- Community ID', e.g. '-12345'. "
func (b *MessagesPinBulder) PeerID(v int) {
	b.Params["peer_id"] = v
}

// MessageID parameter
func (b *MessagesPinBulder) MessageID(v int) {
	b.Params["message_id"] = v
}

// MessagesRemoveChatUserBulder builder
//
// Allows the current user to leave a chat or, if the current user started the chat, allows the user to remove another user from the chat.
//
// https://vk.com/dev/messages.removeChatUser
type MessagesRemoveChatUserBulder struct {
	api.Params
}

// NewMessagesRemoveChatUserBulder func
func NewMessagesRemoveChatUserBulder() *MessagesRemoveChatUserBulder {
	return &MessagesRemoveChatUserBulder{api.Params{}}
}

// ChatID Chat ID.
func (b *MessagesRemoveChatUserBulder) ChatID(v int) {
	b.Params["chat_id"] = v
}

// UserID ID of the user to be removed from the chat.
func (b *MessagesRemoveChatUserBulder) UserID(v int) {
	b.Params["user_id"] = v
}

// MemberID parameter
func (b *MessagesRemoveChatUserBulder) MemberID(v int) {
	b.Params["member_id"] = v
}

// MessagesRestoreBulder builder
//
// Restores a deleted message.
//
// https://vk.com/dev/messages.restore
type MessagesRestoreBulder struct {
	api.Params
}

// NewMessagesRestoreBulder func
func NewMessagesRestoreBulder() *MessagesRestoreBulder {
	return &MessagesRestoreBulder{api.Params{}}
}

// MessageID ID of a previously-deleted message to restore.
func (b *MessagesRestoreBulder) MessageID(v int) {
	b.Params["message_id"] = v
}

// GroupID Group ID (for group messages with user access token)
func (b *MessagesRestoreBulder) GroupID(v int) {
	b.Params["group_id"] = v
}

// MessagesSearchBulder builder
//
// Returns a list of the current user's private messages that match search criteria.
//
// https://vk.com/dev/messages.search
type MessagesSearchBulder struct {
	api.Params
}

// NewMessagesSearchBulder func
func NewMessagesSearchBulder() *MessagesSearchBulder {
	return &MessagesSearchBulder{api.Params{}}
}

// Q Search query string.
func (b *MessagesSearchBulder) Q(v string) {
	b.Params["q"] = v
}

// PeerID Destination ID. "For user: 'User ID', e.g. '12345'. For chat: '2000000000' + 'chat_id', e.g. '2000000001'. For community: '- community ID', e.g. '-12345'. "
func (b *MessagesSearchBulder) PeerID(v int) {
	b.Params["peer_id"] = v
}

// Date Date to search message before in Unixtime.
func (b *MessagesSearchBulder) Date(v int) {
	b.Params["date"] = v
}

// PreviewLength Number of characters after which to truncate a previewed message. To preview the full message, specify '0'. "NOTE: Messages are not truncated by default. Messages are truncated by words."
func (b *MessagesSearchBulder) PreviewLength(v int) {
	b.Params["preview_length"] = v
}

// Offset Offset needed to return a specific subset of messages.
func (b *MessagesSearchBulder) Offset(v int) {
	b.Params["offset"] = v
}

// Count Number of messages to return.
func (b *MessagesSearchBulder) Count(v int) {
	b.Params["count"] = v
}

// Extended parameter
func (b *MessagesSearchBulder) Extended(v bool) {
	b.Params["extended"] = v
}

// Fields parameter
func (b *MessagesSearchBulder) Fields(v []string) {
	b.Params["fields"] = v
}

// GroupID Group ID (for group messages with group access token)
func (b *MessagesSearchBulder) GroupID(v int) {
	b.Params["group_id"] = v
}

// MessagesSearchConversationsBulder builder
//
// Returns a list of the current user's conversations that match search criteria.
//
// https://vk.com/dev/messages.searchConversations
type MessagesSearchConversationsBulder struct {
	api.Params
}

// NewMessagesSearchConversationsBulder func
func NewMessagesSearchConversationsBulder() *MessagesSearchConversationsBulder {
	return &MessagesSearchConversationsBulder{api.Params{}}
}

// Q Search query string.
func (b *MessagesSearchConversationsBulder) Q(v string) {
	b.Params["q"] = v
}

// Count Maximum number of results.
func (b *MessagesSearchConversationsBulder) Count(v int) {
	b.Params["count"] = v
}

// Extended '1' — return extra information about users and communities
func (b *MessagesSearchConversationsBulder) Extended(v bool) {
	b.Params["extended"] = v
}

// Fields Profile fields to return.
func (b *MessagesSearchConversationsBulder) Fields(v []string) {
	b.Params["fields"] = v
}

// GroupID Group ID (for group messages with user access token)
func (b *MessagesSearchConversationsBulder) GroupID(v int) {
	b.Params["group_id"] = v
}

// MessagesSendBulder builder
//
// Sends a message.
//
// https://vk.com/dev/messages.send
type MessagesSendBulder struct {
	api.Params
}

// NewMessagesSendBulder func
func NewMessagesSendBulder() *MessagesSendBulder {
	return &MessagesSendBulder{api.Params{}}
}

// UserID User ID (by default — current user).
func (b *MessagesSendBulder) UserID(v int) {
	b.Params["user_id"] = v
}

// RandomID Unique identifier to avoid resending the message.
func (b *MessagesSendBulder) RandomID(v int) {
	b.Params["random_id"] = v
}

// PeerID Destination ID. "For user: 'User ID', e.g. '12345'. For chat: '2000000000' + 'chat_id', e.g. '2000000001'. For community: '- community ID', e.g. '-12345'. "
func (b *MessagesSendBulder) PeerID(v int) {
	b.Params["peer_id"] = v
}

// Domain User's short address (for example, 'illarionov').
func (b *MessagesSendBulder) Domain(v string) {
	b.Params["domain"] = v
}

// ChatID ID of conversation the message will relate to.
func (b *MessagesSendBulder) ChatID(v int) {
	b.Params["chat_id"] = v
}

// UserIDs IDs of message recipients (if new conversation shall be started).
func (b *MessagesSendBulder) UserIDs(v []int) {
	b.Params["user_ids"] = v
}

// Message (Required if 'attachments' is not set.) Text of the message.
func (b *MessagesSendBulder) Message(v string) {
	b.Params["message"] = v
}

// Lat Geographical latitude of a check-in, in degrees (from -90 to 90).
func (b *MessagesSendBulder) Lat(v float64) {
	b.Params["lat"] = v
}

// Long Geographical longitude of a check-in, in degrees (from -180 to 180).
func (b *MessagesSendBulder) Long(v float64) {
	b.Params["long"] = v
}

// Attachment (Required if 'message' is not set.) List of objects attached to the message, separated by commas, in the following format: "<owner_id>_<media_id>", '' — Type of media attachment: 'photo' — photo, 'video' — video, 'audio' — audio, 'doc' — document, 'wall' — wall post, '<owner_id>' — ID of the media attachment owner. '<media_id>' — media attachment ID. Example: "photo100172_166443618"
func (b *MessagesSendBulder) Attachment(v string) {
	b.Params["attachment"] = v
}

// ReplyTo parameter
func (b *MessagesSendBulder) ReplyTo(v int) {
	b.Params["reply_to"] = v
}

// ForwardMessages ID of forwarded messages, separated with a comma. Listed messages of the sender will be shown in the message body at the recipient's. Example: "123,431,544"
func (b *MessagesSendBulder) ForwardMessages(v []int) {
	b.Params["forward_messages"] = v
}

// Forward parameter
func (b *MessagesSendBulder) Forward(v string) {
	b.Params["forward"] = v
}

// StickerID Sticker id.
func (b *MessagesSendBulder) StickerID(v int) {
	b.Params["sticker_id"] = v
}

// GroupID Group ID (for group messages with group access token)
func (b *MessagesSendBulder) GroupID(v int) {
	b.Params["group_id"] = v
}

// Keyboard parameter
func (b *MessagesSendBulder) Keyboard(v string) {
	b.Params["keyboard"] = v
}

// Payload parameter
func (b *MessagesSendBulder) Payload(v string) {
	b.Params["payload"] = v
}

// DontParseLinks parameter
func (b *MessagesSendBulder) DontParseLinks(v bool) {
	b.Params["dont_parse_links"] = v
}

// DisableMentions parameter
func (b *MessagesSendBulder) DisableMentions(v bool) {
	b.Params["disable_mentions"] = v
}

// MessagesSetActivityBulder builder
//
// Changes the status of a user as typing in a conversation.
//
// https://vk.com/dev/messages.setActivity
type MessagesSetActivityBulder struct {
	api.Params
}

// NewMessagesSetActivityBulder func
func NewMessagesSetActivityBulder() *MessagesSetActivityBulder {
	return &MessagesSetActivityBulder{api.Params{}}
}

// UserID User ID.
func (b *MessagesSetActivityBulder) UserID(v int) {
	b.Params["user_id"] = v
}

// Type 'typing' — user has started to type.
func (b *MessagesSetActivityBulder) Type(v string) {
	b.Params["type"] = v
}

// PeerID Destination ID. "For user: 'User ID', e.g. '12345'. For chat: '2000000000' + 'chat_id', e.g. '2000000001'. For community: '- community ID', e.g. '-12345'. "
func (b *MessagesSetActivityBulder) PeerID(v int) {
	b.Params["peer_id"] = v
}

// GroupID Group ID (for group messages with group access token)
func (b *MessagesSetActivityBulder) GroupID(v int) {
	b.Params["group_id"] = v
}

// MessagesSetChatPhotoBulder builder
//
// Sets a previously-uploaded picture as the cover picture of a chat.
//
// https://vk.com/dev/messages.setChatPhoto
type MessagesSetChatPhotoBulder struct {
	api.Params
}

// NewMessagesSetChatPhotoBulder func
func NewMessagesSetChatPhotoBulder() *MessagesSetChatPhotoBulder {
	return &MessagesSetChatPhotoBulder{api.Params{}}
}

// File Upload URL from the 'response' field returned by the [vk.com/dev/photos.getChatUploadServer|photos.getChatUploadServer] method upon successfully uploading an image.
func (b *MessagesSetChatPhotoBulder) File(v string) {
	b.Params["file"] = v
}

// MessagesUnpinBulder builder
//
// https://vk.com/dev/messages.unpin
type MessagesUnpinBulder struct {
	api.Params
}

// NewMessagesUnpinBulder func
func NewMessagesUnpinBulder() *MessagesUnpinBulder {
	return &MessagesUnpinBulder{api.Params{}}
}

// PeerID parameter
func (b *MessagesUnpinBulder) PeerID(v int) {
	b.Params["peer_id"] = v
}

// GroupID parameter
func (b *MessagesUnpinBulder) GroupID(v int) {
	b.Params["group_id"] = v
}
