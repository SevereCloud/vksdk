package params // import "github.com/SevereCloud/vksdk/api/params"

import (
	"github.com/SevereCloud/vksdk/api"
)

// MessagesAddChatUserBuilder builder
//
// Adds a new user to a chat.
//
// https://vk.com/dev/messages.addChatUser
type MessagesAddChatUserBuilder struct {
	api.Params
}

// NewMessagesAddChatUserBuilder func
func NewMessagesAddChatUserBuilder() *MessagesAddChatUserBuilder {
	return &MessagesAddChatUserBuilder{api.Params{}}
}

// ChatID Chat ID.
func (b *MessagesAddChatUserBuilder) ChatID(v int) {
	b.Params["chat_id"] = v
}

// UserID ID of the user to be added to the chat.
func (b *MessagesAddChatUserBuilder) UserID(v int) {
	b.Params["user_id"] = v
}

// MessagesAllowMessagesFromGroupBuilder builder
//
// Allows sending messages from community to the current user.
//
// https://vk.com/dev/messages.allowMessagesFromGroup
type MessagesAllowMessagesFromGroupBuilder struct {
	api.Params
}

// NewMessagesAllowMessagesFromGroupBuilder func
func NewMessagesAllowMessagesFromGroupBuilder() *MessagesAllowMessagesFromGroupBuilder {
	return &MessagesAllowMessagesFromGroupBuilder{api.Params{}}
}

// GroupID Group ID.
func (b *MessagesAllowMessagesFromGroupBuilder) GroupID(v int) {
	b.Params["group_id"] = v
}

// Key parameter
func (b *MessagesAllowMessagesFromGroupBuilder) Key(v string) {
	b.Params["key"] = v
}

// MessagesCreateChatBuilder builder
//
// Creates a chat with several participants.
//
// https://vk.com/dev/messages.createChat
type MessagesCreateChatBuilder struct {
	api.Params
}

// NewMessagesCreateChatBuilder func
func NewMessagesCreateChatBuilder() *MessagesCreateChatBuilder {
	return &MessagesCreateChatBuilder{api.Params{}}
}

// UserIDs IDs of the users to be added to the chat.
func (b *MessagesCreateChatBuilder) UserIDs(v []int) {
	b.Params["user_ids"] = v
}

// Title Chat title.
func (b *MessagesCreateChatBuilder) Title(v string) {
	b.Params["title"] = v
}

// MessagesDeleteBuilder builder
//
// Deletes one or more messages.
//
// https://vk.com/dev/messages.delete
type MessagesDeleteBuilder struct {
	api.Params
}

// NewMessagesDeleteBuilder func
func NewMessagesDeleteBuilder() *MessagesDeleteBuilder {
	return &MessagesDeleteBuilder{api.Params{}}
}

// MessageIDs Message IDs.
func (b *MessagesDeleteBuilder) MessageIDs(v []int) {
	b.Params["message_ids"] = v
}

// Spam '1' — to mark message as spam.
func (b *MessagesDeleteBuilder) Spam(v bool) {
	b.Params["spam"] = v
}

// GroupID Group ID (for group messages with user access token)
func (b *MessagesDeleteBuilder) GroupID(v int) {
	b.Params["group_id"] = v
}

// DeleteForAll '1' — delete message for for all.
func (b *MessagesDeleteBuilder) DeleteForAll(v bool) {
	b.Params["delete_for_all"] = v
}

// MessagesDeleteChatPhotoBuilder builder
//
// Deletes a chat's cover picture.
//
// https://vk.com/dev/messages.deleteChatPhoto
type MessagesDeleteChatPhotoBuilder struct {
	api.Params
}

// NewMessagesDeleteChatPhotoBuilder func
func NewMessagesDeleteChatPhotoBuilder() *MessagesDeleteChatPhotoBuilder {
	return &MessagesDeleteChatPhotoBuilder{api.Params{}}
}

// ChatID Chat ID.
func (b *MessagesDeleteChatPhotoBuilder) ChatID(v int) {
	b.Params["chat_id"] = v
}

// GroupID parameter
func (b *MessagesDeleteChatPhotoBuilder) GroupID(v int) {
	b.Params["group_id"] = v
}

// MessagesDeleteConversationBuilder builder
//
// Deletes all private messages in a conversation.
//
// https://vk.com/dev/messages.deleteConversation
type MessagesDeleteConversationBuilder struct {
	api.Params
}

// NewMessagesDeleteConversationBuilder func
func NewMessagesDeleteConversationBuilder() *MessagesDeleteConversationBuilder {
	return &MessagesDeleteConversationBuilder{api.Params{}}
}

// UserID User ID. To clear a chat history use 'chat_id'
func (b *MessagesDeleteConversationBuilder) UserID(v int) {
	b.Params["user_id"] = v
}

// PeerID Destination ID. For user: 'User ID', e.g. '12345'.
// For chat: '2000000000' + 'chat_id', e.g. '2000000001'.
// For community: '- community ID', e.g. '-12345'.
func (b *MessagesDeleteConversationBuilder) PeerID(v int) {
	b.Params["peer_id"] = v
}

// GroupID Group ID (for group messages with user access token)
func (b *MessagesDeleteConversationBuilder) GroupID(v int) {
	b.Params["group_id"] = v
}

// MessagesDenyMessagesFromGroupBuilder builder
//
// Denies sending message from community to the current user.
//
// https://vk.com/dev/messages.denyMessagesFromGroup
type MessagesDenyMessagesFromGroupBuilder struct {
	api.Params
}

// NewMessagesDenyMessagesFromGroupBuilder func
func NewMessagesDenyMessagesFromGroupBuilder() *MessagesDenyMessagesFromGroupBuilder {
	return &MessagesDenyMessagesFromGroupBuilder{api.Params{}}
}

// GroupID Group ID.
func (b *MessagesDenyMessagesFromGroupBuilder) GroupID(v int) {
	b.Params["group_id"] = v
}

// MessagesEditBuilder builder
//
// Edits the message.
//
// https://vk.com/dev/messages.edit
type MessagesEditBuilder struct {
	api.Params
}

// NewMessagesEditBuilder func
func NewMessagesEditBuilder() *MessagesEditBuilder {
	return &MessagesEditBuilder{api.Params{}}
}

// PeerID Destination ID. For user: 'User ID', e.g. '12345'.
// For chat: '2000000000' + 'chat_id', e.g. '2000000001'.
// For community: '- community ID', e.g. '-12345'.
func (b *MessagesEditBuilder) PeerID(v int) {
	b.Params["peer_id"] = v
}

// Message (Required if 'attachments' is not set.) Text of the message.
func (b *MessagesEditBuilder) Message(v string) {
	b.Params["message"] = v
}

// MessageID parameter
func (b *MessagesEditBuilder) MessageID(v int) {
	b.Params["message_id"] = v
}

// Lat Geographical latitude of a check-in, in degrees (from -90 to 90).
func (b *MessagesEditBuilder) Lat(v float64) {
	b.Params["lat"] = v
}

// Long Geographical longitude of a check-in, in degrees (from -180 to 180).
func (b *MessagesEditBuilder) Long(v float64) {
	b.Params["long"] = v
}

// Attachment (Required if 'message' is not set.) List of objects attached to
// the message, separated by commas, in the following format:
// "<owner_id>_<media_id>", '' — Type of media attachment: 'photo' — photo,
// 'video' — video, 'audio' — audio, 'doc' — document, 'wall' — wall post,
// '<owner_id>' — ID of the media attachment owner. '<media_id>' — media
// attachment ID. Example: "photo100172_166443618"
func (b *MessagesEditBuilder) Attachment(v string) {
	b.Params["attachment"] = v
}

// KeepForwardMessages '1' — to keep forwarded, messages.
func (b *MessagesEditBuilder) KeepForwardMessages(v bool) {
	b.Params["keep_forward_messages"] = v
}

// KeepSnippets '1' — to keep attached snippets.
func (b *MessagesEditBuilder) KeepSnippets(v bool) {
	b.Params["keep_snippets"] = v
}

// GroupID Group ID (for group messages with user access token)
func (b *MessagesEditBuilder) GroupID(v int) {
	b.Params["group_id"] = v
}

// DontParseLinks parameter
func (b *MessagesEditBuilder) DontParseLinks(v bool) {
	b.Params["dont_parse_links"] = v
}

// MessagesEditChatBuilder builder
//
// Edits the title of a chat.
//
// https://vk.com/dev/messages.editChat
type MessagesEditChatBuilder struct {
	api.Params
}

// NewMessagesEditChatBuilder func
func NewMessagesEditChatBuilder() *MessagesEditChatBuilder {
	return &MessagesEditChatBuilder{api.Params{}}
}

// ChatID Chat ID.
func (b *MessagesEditChatBuilder) ChatID(v int) {
	b.Params["chat_id"] = v
}

// Title New title of the chat.
func (b *MessagesEditChatBuilder) Title(v string) {
	b.Params["title"] = v
}

// MessagesGetByConversationMessageIDBuilder builder
//
// Returns messages by their IDs within the conversation.
//
// https://vk.com/dev/messages.getByConversationMessageId
type MessagesGetByConversationMessageIDBuilder struct {
	api.Params
}

// NewMessagesGetByConversationMessageIDBuilder func
func NewMessagesGetByConversationMessageIDBuilder() *MessagesGetByConversationMessageIDBuilder {
	return &MessagesGetByConversationMessageIDBuilder{api.Params{}}
}

// PeerID Destination ID. For user: 'User ID', e.g. '12345'.
// For chat: '2000000000' + 'chat_id', e.g. '2000000001'.
// For community: '- community ID', e.g. '-12345'.
func (b *MessagesGetByConversationMessageIDBuilder) PeerID(v int) {
	b.Params["peer_id"] = v
}

// ConversationMessageIDs Conversation message IDs.
func (b *MessagesGetByConversationMessageIDBuilder) ConversationMessageIDs(v []int) {
	b.Params["conversation_message_ids"] = v
}

// Extended Information whether the response should be extended
func (b *MessagesGetByConversationMessageIDBuilder) Extended(v bool) {
	b.Params["extended"] = v
}

// Fields Profile fields to return.
func (b *MessagesGetByConversationMessageIDBuilder) Fields(v []string) {
	b.Params["fields"] = v
}

// GroupID Group ID (for group messages with group access token)
func (b *MessagesGetByConversationMessageIDBuilder) GroupID(v int) {
	b.Params["group_id"] = v
}

// MessagesGetByIDBuilder builder
//
// Returns messages by their IDs.
//
// https://vk.com/dev/messages.getById
type MessagesGetByIDBuilder struct {
	api.Params
}

// NewMessagesGetByIDBuilder func
func NewMessagesGetByIDBuilder() *MessagesGetByIDBuilder {
	return &MessagesGetByIDBuilder{api.Params{}}
}

// MessageIDs Message IDs.
func (b *MessagesGetByIDBuilder) MessageIDs(v []int) {
	b.Params["message_ids"] = v
}

// PreviewLength Number of characters after which to truncate a previewed
// message. To preview the full message, specify '0'.
// NOTE: Messages are not truncated by default.
// Messages are truncated by words.
func (b *MessagesGetByIDBuilder) PreviewLength(v int) {
	b.Params["preview_length"] = v
}

// Extended Information whether the response should be extended
func (b *MessagesGetByIDBuilder) Extended(v bool) {
	b.Params["extended"] = v
}

// Fields Profile fields to return.
func (b *MessagesGetByIDBuilder) Fields(v []string) {
	b.Params["fields"] = v
}

// GroupID Group ID (for group messages with group access token)
func (b *MessagesGetByIDBuilder) GroupID(v int) {
	b.Params["group_id"] = v
}

// MessagesGetChatPreviewBuilder builder
//
// https://vk.com/dev/messages.getChatPreview
type MessagesGetChatPreviewBuilder struct {
	api.Params
}

// NewMessagesGetChatPreviewBuilder func
func NewMessagesGetChatPreviewBuilder() *MessagesGetChatPreviewBuilder {
	return &MessagesGetChatPreviewBuilder{api.Params{}}
}

// PeerID parameter
func (b *MessagesGetChatPreviewBuilder) PeerID(v int) {
	b.Params["peer_id"] = v
}

// Link Invitation link.
func (b *MessagesGetChatPreviewBuilder) Link(v string) {
	b.Params["link"] = v
}

// Fields Profile fields to return.
func (b *MessagesGetChatPreviewBuilder) Fields(v []string) {
	b.Params["fields"] = v
}

// MessagesGetConversationMembersBuilder builder
//
// Returns a list of IDs of users participating in a chat.
//
// https://vk.com/dev/messages.getConversationMembers
type MessagesGetConversationMembersBuilder struct {
	api.Params
}

// NewMessagesGetConversationMembersBuilder func
func NewMessagesGetConversationMembersBuilder() *MessagesGetConversationMembersBuilder {
	return &MessagesGetConversationMembersBuilder{api.Params{}}
}

// PeerID Peer ID.
func (b *MessagesGetConversationMembersBuilder) PeerID(v int) {
	b.Params["peer_id"] = v
}

// Fields Profile fields to return.
func (b *MessagesGetConversationMembersBuilder) Fields(v []string) {
	b.Params["fields"] = v
}

// GroupID Group ID (for group messages with group access token)
func (b *MessagesGetConversationMembersBuilder) GroupID(v int) {
	b.Params["group_id"] = v
}

// MessagesGetConversationsBuilder builder
//
// Returns a list of the current user's conversations.
//
// https://vk.com/dev/messages.getConversations
type MessagesGetConversationsBuilder struct {
	api.Params
}

// NewMessagesGetConversationsBuilder func
func NewMessagesGetConversationsBuilder() *MessagesGetConversationsBuilder {
	return &MessagesGetConversationsBuilder{api.Params{}}
}

// Offset Offset needed to return a specific subset of conversations.
func (b *MessagesGetConversationsBuilder) Offset(v int) {
	b.Params["offset"] = v
}

// Count Number of conversations to return.
func (b *MessagesGetConversationsBuilder) Count(v int) {
	b.Params["count"] = v
}

// Filter Filter to apply: 'all' — all conversations, 'unread' — conversations
// with unread messages, 'important' — conversations, marked as important (only
// for community messages), 'unanswered' — conversations, marked as unanswered
// (only for community messages)
func (b *MessagesGetConversationsBuilder) Filter(v string) {
	b.Params["filter"] = v
}

// Extended '1' — return extra information about users and communities
func (b *MessagesGetConversationsBuilder) Extended(v bool) {
	b.Params["extended"] = v
}

// StartMessageID ID of the message from what to return dialogs.
func (b *MessagesGetConversationsBuilder) StartMessageID(v int) {
	b.Params["start_message_id"] = v
}

// Fields Profile and communities fields to return.
func (b *MessagesGetConversationsBuilder) Fields(v []string) {
	b.Params["fields"] = v
}

// GroupID Group ID (for group messages with group access token)
func (b *MessagesGetConversationsBuilder) GroupID(v int) {
	b.Params["group_id"] = v
}

// MessagesGetConversationsByIDBuilder builder
//
// Returns conversations by their IDs
//
// https://vk.com/dev/messages.getConversationsById
type MessagesGetConversationsByIDBuilder struct {
	api.Params
}

// NewMessagesGetConversationsByIDBuilder func
func NewMessagesGetConversationsByIDBuilder() *MessagesGetConversationsByIDBuilder {
	return &MessagesGetConversationsByIDBuilder{api.Params{}}
}

// PeerIDs Destination IDs. For user: 'User ID', e.g. '12345'.
// For chat: '2000000000' + 'chat_id', e.g. '2000000001'.
// For community: '- community ID', e.g. '-12345'.
func (b *MessagesGetConversationsByIDBuilder) PeerIDs(v []int) {
	b.Params["peer_ids"] = v
}

// Extended Return extended properties
func (b *MessagesGetConversationsByIDBuilder) Extended(v bool) {
	b.Params["extended"] = v
}

// Fields Profile and communities fields to return.
func (b *MessagesGetConversationsByIDBuilder) Fields(v []string) {
	b.Params["fields"] = v
}

// GroupID Group ID (for group messages with group access token)
func (b *MessagesGetConversationsByIDBuilder) GroupID(v int) {
	b.Params["group_id"] = v
}

// MessagesGetHistoryBuilder builder
//
// Returns message history for the specified user or group chat.
//
// https://vk.com/dev/messages.getHistory
type MessagesGetHistoryBuilder struct {
	api.Params
}

// NewMessagesGetHistoryBuilder func
func NewMessagesGetHistoryBuilder() *MessagesGetHistoryBuilder {
	return &MessagesGetHistoryBuilder{api.Params{}}
}

// Offset Offset needed to return a specific subset of messages.
func (b *MessagesGetHistoryBuilder) Offset(v int) {
	b.Params["offset"] = v
}

// Count Number of messages to return.
func (b *MessagesGetHistoryBuilder) Count(v int) {
	b.Params["count"] = v
}

// UserID ID of the user whose message history you want to return.
func (b *MessagesGetHistoryBuilder) UserID(v int) {
	b.Params["user_id"] = v
}

// PeerID parameter
func (b *MessagesGetHistoryBuilder) PeerID(v int) {
	b.Params["peer_id"] = v
}

// StartMessageID Starting message ID from which to return history.
func (b *MessagesGetHistoryBuilder) StartMessageID(v int) {
	b.Params["start_message_id"] = v
}

// Rev Sort order: '1' — return messages in chronological order.
// '0' — return messages in reverse chronological order.
func (b *MessagesGetHistoryBuilder) Rev(v int) {
	b.Params["rev"] = v
}

// Extended Information whether the response should be extended
func (b *MessagesGetHistoryBuilder) Extended(v bool) {
	b.Params["extended"] = v
}

// Fields Profile fields to return.
func (b *MessagesGetHistoryBuilder) Fields(v []string) {
	b.Params["fields"] = v
}

// GroupID Group ID (for group messages with group access token)
func (b *MessagesGetHistoryBuilder) GroupID(v int) {
	b.Params["group_id"] = v
}

// MessagesGetHistoryAttachmentsBuilder builder
//
// Returns media files from the dialog or group chat.
//
// https://vk.com/dev/messages.getHistoryAttachments
type MessagesGetHistoryAttachmentsBuilder struct {
	api.Params
}

// NewMessagesGetHistoryAttachmentsBuilder func
func NewMessagesGetHistoryAttachmentsBuilder() *MessagesGetHistoryAttachmentsBuilder {
	return &MessagesGetHistoryAttachmentsBuilder{api.Params{}}
}

// PeerID Peer ID. ", For group chat: '2000000000 + chat ID' , ,
// For community: '-community ID'"
func (b *MessagesGetHistoryAttachmentsBuilder) PeerID(v int) {
	b.Params["peer_id"] = v
}

// MediaType Type of media files to return: *'photo',, *'video',, *'audio',,
// *'doc',, *'link'.,*'market'.,*'wall'.,*'share'
func (b *MessagesGetHistoryAttachmentsBuilder) MediaType(v string) {
	b.Params["media_type"] = v
}

// StartFrom Message ID to start return results from.
func (b *MessagesGetHistoryAttachmentsBuilder) StartFrom(v string) {
	b.Params["start_from"] = v
}

// Count Number of objects to return.
func (b *MessagesGetHistoryAttachmentsBuilder) Count(v int) {
	b.Params["count"] = v
}

// PhotoSizes '1' — to return photo sizes in a
func (b *MessagesGetHistoryAttachmentsBuilder) PhotoSizes(v bool) {
	b.Params["photo_sizes"] = v
}

// Fields Additional profile [vk.com/dev/fields|fields] to return.
func (b *MessagesGetHistoryAttachmentsBuilder) Fields(v []string) {
	b.Params["fields"] = v
}

// GroupID Group ID (for group messages with group access token)
func (b *MessagesGetHistoryAttachmentsBuilder) GroupID(v int) {
	b.Params["group_id"] = v
}

// PreserveOrder parameter
func (b *MessagesGetHistoryAttachmentsBuilder) PreserveOrder(v bool) {
	b.Params["preserve_order"] = v
}

// MaxForwardsLevel parameter
func (b *MessagesGetHistoryAttachmentsBuilder) MaxForwardsLevel(v int) {
	b.Params["max_forwards_level"] = v
}

// MessagesGetInviteLinkBuilder builder
//
// https://vk.com/dev/messages.getInviteLink
type MessagesGetInviteLinkBuilder struct {
	api.Params
}

// NewMessagesGetInviteLinkBuilder func
func NewMessagesGetInviteLinkBuilder() *MessagesGetInviteLinkBuilder {
	return &MessagesGetInviteLinkBuilder{api.Params{}}
}

// PeerID Destination ID.
func (b *MessagesGetInviteLinkBuilder) PeerID(v int) {
	b.Params["peer_id"] = v
}

// Reset 1 — to generate new link (revoke previous), 0 — to return previous
// link.
func (b *MessagesGetInviteLinkBuilder) Reset(v bool) {
	b.Params["reset"] = v
}

// GroupID Group ID
func (b *MessagesGetInviteLinkBuilder) GroupID(v int) {
	b.Params["group_id"] = v
}

// MessagesGetLastActivityBuilder builder
//
// Returns a user's current status and date of last activity.
//
// https://vk.com/dev/messages.getLastActivity
type MessagesGetLastActivityBuilder struct {
	api.Params
}

// NewMessagesGetLastActivityBuilder func
func NewMessagesGetLastActivityBuilder() *MessagesGetLastActivityBuilder {
	return &MessagesGetLastActivityBuilder{api.Params{}}
}

// UserID User ID.
func (b *MessagesGetLastActivityBuilder) UserID(v int) {
	b.Params["user_id"] = v
}

// MessagesGetLongPollHistoryBuilder builder
//
// Returns updates in user's private messages.
//
// https://vk.com/dev/messages.getLongPollHistory
type MessagesGetLongPollHistoryBuilder struct {
	api.Params
}

// NewMessagesGetLongPollHistoryBuilder func
func NewMessagesGetLongPollHistoryBuilder() *MessagesGetLongPollHistoryBuilder {
	return &MessagesGetLongPollHistoryBuilder{api.Params{}}
}

// Ts Last value of the 'ts' parameter returned from the Long Poll server or
// by using [vk.com/dev/messages.getLongPollHistory|
// messages.getLongPollHistory] method.
func (b *MessagesGetLongPollHistoryBuilder) Ts(v int) {
	b.Params["ts"] = v
}

// Pts Lsat value of 'pts' parameter returned from the Long Poll server or by
// using [vk.com/dev/messages.getLongPollHistory|messages.getLongPollHistory]
// method.
func (b *MessagesGetLongPollHistoryBuilder) Pts(v int) {
	b.Params["pts"] = v
}

// PreviewLength Number of characters after which to truncate a previewed
// message. To preview the full message, specify '0'.
//
// NOTE: Messages are not truncated by default. Messages are truncated by
// words.
func (b *MessagesGetLongPollHistoryBuilder) PreviewLength(v int) {
	b.Params["preview_length"] = v
}

// Onlines '1' — to return history with online users only.
func (b *MessagesGetLongPollHistoryBuilder) Onlines(v bool) {
	b.Params["onlines"] = v
}

// Fields Additional profile [vk.com/dev/fields|fields] to return.
func (b *MessagesGetLongPollHistoryBuilder) Fields(v []string) {
	b.Params["fields"] = v
}

// EventsLimit Maximum number of events to return.
func (b *MessagesGetLongPollHistoryBuilder) EventsLimit(v int) {
	b.Params["events_limit"] = v
}

// MsgsLimit Maximum number of messages to return.
func (b *MessagesGetLongPollHistoryBuilder) MsgsLimit(v int) {
	b.Params["msgs_limit"] = v
}

// MaxMsgID Maximum ID of the message among existing ones in the local copy.
// Both messages received with API methods (for example, , ), and data
// received from a Long Poll server (events with code 4) are taken into
// account.
func (b *MessagesGetLongPollHistoryBuilder) MaxMsgID(v int) {
	b.Params["max_msg_id"] = v
}

// GroupID Group ID (for group messages with user access token)
func (b *MessagesGetLongPollHistoryBuilder) GroupID(v int) {
	b.Params["group_id"] = v
}

// LpVersion parameter
func (b *MessagesGetLongPollHistoryBuilder) LpVersion(v int) {
	b.Params["lp_version"] = v
}

// LastN parameter
func (b *MessagesGetLongPollHistoryBuilder) LastN(v int) {
	b.Params["last_n"] = v
}

// Credentials parameter
func (b *MessagesGetLongPollHistoryBuilder) Credentials(v bool) {
	b.Params["credentials"] = v
}

// MessagesGetLongPollServerBuilder builder
//
// Returns data required for connection to a Long Poll server.
//
// https://vk.com/dev/messages.getLongPollServer
type MessagesGetLongPollServerBuilder struct {
	api.Params
}

// NewMessagesGetLongPollServerBuilder func
func NewMessagesGetLongPollServerBuilder() *MessagesGetLongPollServerBuilder {
	return &MessagesGetLongPollServerBuilder{api.Params{}}
}

// NeedPts '1' — to return the 'pts' field, needed for the
// [vk.com/dev/messages.getLongPollHistory|messages.getLongPollHistory] method.
func (b *MessagesGetLongPollServerBuilder) NeedPts(v bool) {
	b.Params["need_pts"] = v
}

// GroupID Group ID (for group messages with user access token)
func (b *MessagesGetLongPollServerBuilder) GroupID(v int) {
	b.Params["group_id"] = v
}

// LpVersion Long poll version
func (b *MessagesGetLongPollServerBuilder) LpVersion(v int) {
	b.Params["lp_version"] = v
}

// MessagesIsMessagesFromGroupAllowedBuilder builder
//
// Returns information whether sending messages from the community to
// current user is allowed.
//
// https://vk.com/dev/messages.isMessagesFromGroupAllowed
type MessagesIsMessagesFromGroupAllowedBuilder struct {
	api.Params
}

// NewMessagesIsMessagesFromGroupAllowedBuilder func
func NewMessagesIsMessagesFromGroupAllowedBuilder() *MessagesIsMessagesFromGroupAllowedBuilder {
	return &MessagesIsMessagesFromGroupAllowedBuilder{api.Params{}}
}

// GroupID Group ID.
func (b *MessagesIsMessagesFromGroupAllowedBuilder) GroupID(v int) {
	b.Params["group_id"] = v
}

// UserID User ID.
func (b *MessagesIsMessagesFromGroupAllowedBuilder) UserID(v int) {
	b.Params["user_id"] = v
}

// MessagesJoinChatByInviteLinkBuilder builder
//
// https://vk.com/dev/messages.joinChatByInviteLink
type MessagesJoinChatByInviteLinkBuilder struct {
	api.Params
}

// NewMessagesJoinChatByInviteLinkBuilder func
func NewMessagesJoinChatByInviteLinkBuilder() *MessagesJoinChatByInviteLinkBuilder {
	return &MessagesJoinChatByInviteLinkBuilder{api.Params{}}
}

// Link Invitation link.
func (b *MessagesJoinChatByInviteLinkBuilder) Link(v string) {
	b.Params["link"] = v
}

// MessagesMarkAsAnsweredConversationBuilder builder
//
// Marks and unmarks conversations as unanswered.
//
// https://vk.com/dev/messages.markAsAnsweredConversation
type MessagesMarkAsAnsweredConversationBuilder struct {
	api.Params
}

// NewMessagesMarkAsAnsweredConversationBuilder func
func NewMessagesMarkAsAnsweredConversationBuilder() *MessagesMarkAsAnsweredConversationBuilder {
	return &MessagesMarkAsAnsweredConversationBuilder{api.Params{}}
}

// PeerID ID of conversation to mark as important.
func (b *MessagesMarkAsAnsweredConversationBuilder) PeerID(v int) {
	b.Params["peer_id"] = v
}

// Answered '1' — to mark as answered, '0' — to remove the mark
func (b *MessagesMarkAsAnsweredConversationBuilder) Answered(v bool) {
	b.Params["answered"] = v
}

// GroupID Group ID (for group messages with group access token)
func (b *MessagesMarkAsAnsweredConversationBuilder) GroupID(v int) {
	b.Params["group_id"] = v
}

// MessagesMarkAsImportantBuilder builder
//
// Marks and unmarks messages as important (starred).
//
// https://vk.com/dev/messages.markAsImportant
type MessagesMarkAsImportantBuilder struct {
	api.Params
}

// NewMessagesMarkAsImportantBuilder func
func NewMessagesMarkAsImportantBuilder() *MessagesMarkAsImportantBuilder {
	return &MessagesMarkAsImportantBuilder{api.Params{}}
}

// MessageIDs IDs of messages to mark as important.
func (b *MessagesMarkAsImportantBuilder) MessageIDs(v []int) {
	b.Params["message_ids"] = v
}

// Important '1' — to add a star (mark as important), '0' — to remove the star
func (b *MessagesMarkAsImportantBuilder) Important(v int) {
	b.Params["important"] = v
}

// MessagesMarkAsImportantConversationBuilder builder
//
// Marks and unmarks conversations as important.
//
// https://vk.com/dev/messages.markAsImportantConversation
type MessagesMarkAsImportantConversationBuilder struct {
	api.Params
}

// NewMessagesMarkAsImportantConversationBuilder func
func NewMessagesMarkAsImportantConversationBuilder() *MessagesMarkAsImportantConversationBuilder {
	return &MessagesMarkAsImportantConversationBuilder{api.Params{}}
}

// PeerID ID of conversation to mark as important.
func (b *MessagesMarkAsImportantConversationBuilder) PeerID(v int) {
	b.Params["peer_id"] = v
}

// Important '1' — to add a star (mark as important), '0' — to remove the star
func (b *MessagesMarkAsImportantConversationBuilder) Important(v bool) {
	b.Params["important"] = v
}

// GroupID Group ID (for group messages with group access token)
func (b *MessagesMarkAsImportantConversationBuilder) GroupID(v int) {
	b.Params["group_id"] = v
}

// MessagesMarkAsReadBuilder builder
//
// Marks messages as read.
//
// https://vk.com/dev/messages.markAsRead
type MessagesMarkAsReadBuilder struct {
	api.Params
}

// NewMessagesMarkAsReadBuilder func
func NewMessagesMarkAsReadBuilder() *MessagesMarkAsReadBuilder {
	return &MessagesMarkAsReadBuilder{api.Params{}}
}

// MessageIDs IDs of messages to mark as read.
func (b *MessagesMarkAsReadBuilder) MessageIDs(v []int) {
	b.Params["message_ids"] = v
}

// PeerID Destination ID. For user: 'User ID', e.g. '12345'.
// For chat: '2000000000' + 'chat_id', e.g. '2000000001'.
// For community: '- community ID', e.g. '-12345'.
func (b *MessagesMarkAsReadBuilder) PeerID(v int) {
	b.Params["peer_id"] = v
}

// StartMessageID Message ID to start from.
func (b *MessagesMarkAsReadBuilder) StartMessageID(v int) {
	b.Params["start_message_id"] = v
}

// GroupID Group ID (for group messages with user access token)
func (b *MessagesMarkAsReadBuilder) GroupID(v int) {
	b.Params["group_id"] = v
}

// MessagesPinBuilder builder
//
// Pin a message.
//
// https://vk.com/dev/messages.pin
type MessagesPinBuilder struct {
	api.Params
}

// NewMessagesPinBuilder func
func NewMessagesPinBuilder() *MessagesPinBuilder {
	return &MessagesPinBuilder{api.Params{}}
}

// PeerID Destination ID. For user: 'User ID', e.g. '12345'. For chat:
// '2000000000' + 'Chat ID', e.g. '2000000001'. For community:
// '- Community ID', e.g. '-12345'.
func (b *MessagesPinBuilder) PeerID(v int) {
	b.Params["peer_id"] = v
}

// MessageID parameter
func (b *MessagesPinBuilder) MessageID(v int) {
	b.Params["message_id"] = v
}

// MessagesRemoveChatUserBuilder builder
//
// Allows the current user to leave a chat or, if the current user started the
// chat, allows the user to remove another user from the chat.
//
// https://vk.com/dev/messages.removeChatUser
type MessagesRemoveChatUserBuilder struct {
	api.Params
}

// NewMessagesRemoveChatUserBuilder func
func NewMessagesRemoveChatUserBuilder() *MessagesRemoveChatUserBuilder {
	return &MessagesRemoveChatUserBuilder{api.Params{}}
}

// ChatID Chat ID.
func (b *MessagesRemoveChatUserBuilder) ChatID(v int) {
	b.Params["chat_id"] = v
}

// UserID ID of the user to be removed from the chat.
func (b *MessagesRemoveChatUserBuilder) UserID(v int) {
	b.Params["user_id"] = v
}

// MemberID parameter
func (b *MessagesRemoveChatUserBuilder) MemberID(v int) {
	b.Params["member_id"] = v
}

// MessagesRestoreBuilder builder
//
// Restores a deleted message.
//
// https://vk.com/dev/messages.restore
type MessagesRestoreBuilder struct {
	api.Params
}

// NewMessagesRestoreBuilder func
func NewMessagesRestoreBuilder() *MessagesRestoreBuilder {
	return &MessagesRestoreBuilder{api.Params{}}
}

// MessageID ID of a previously-deleted message to restore.
func (b *MessagesRestoreBuilder) MessageID(v int) {
	b.Params["message_id"] = v
}

// GroupID Group ID (for group messages with user access token)
func (b *MessagesRestoreBuilder) GroupID(v int) {
	b.Params["group_id"] = v
}

// MessagesSearchBuilder builder
//
// Returns a list of the current user's private messages that match search
// criteria.
//
// https://vk.com/dev/messages.search
type MessagesSearchBuilder struct {
	api.Params
}

// NewMessagesSearchBuilder func
func NewMessagesSearchBuilder() *MessagesSearchBuilder {
	return &MessagesSearchBuilder{api.Params{}}
}

// Q Search query string.
func (b *MessagesSearchBuilder) Q(v string) {
	b.Params["q"] = v
}

// PeerID Destination ID. For user: 'User ID', e.g. '12345'. For chat:
// '2000000000' + 'chat_id', e.g. '2000000001'. For community:
// '- community ID', e.g. '-12345'.
func (b *MessagesSearchBuilder) PeerID(v int) {
	b.Params["peer_id"] = v
}

// Date Date to search message before in Unixtime.
func (b *MessagesSearchBuilder) Date(v int) {
	b.Params["date"] = v
}

// PreviewLength Number of characters after which to truncate a previewed
// message. To preview the full message, specify '0'.
//
// NOTE: Messages are not truncated by default. Messages are truncated by
// words.
func (b *MessagesSearchBuilder) PreviewLength(v int) {
	b.Params["preview_length"] = v
}

// Offset Offset needed to return a specific subset of messages.
func (b *MessagesSearchBuilder) Offset(v int) {
	b.Params["offset"] = v
}

// Count Number of messages to return.
func (b *MessagesSearchBuilder) Count(v int) {
	b.Params["count"] = v
}

// Extended parameter
func (b *MessagesSearchBuilder) Extended(v bool) {
	b.Params["extended"] = v
}

// Fields parameter
func (b *MessagesSearchBuilder) Fields(v []string) {
	b.Params["fields"] = v
}

// GroupID Group ID (for group messages with group access token)
func (b *MessagesSearchBuilder) GroupID(v int) {
	b.Params["group_id"] = v
}

// MessagesSearchConversationsBuilder builder
//
// Returns a list of the current user's conversations that match search
// criteria.
//
// https://vk.com/dev/messages.searchConversations
type MessagesSearchConversationsBuilder struct {
	api.Params
}

// NewMessagesSearchConversationsBuilder func
func NewMessagesSearchConversationsBuilder() *MessagesSearchConversationsBuilder {
	return &MessagesSearchConversationsBuilder{api.Params{}}
}

// Q Search query string.
func (b *MessagesSearchConversationsBuilder) Q(v string) {
	b.Params["q"] = v
}

// Count Maximum number of results.
func (b *MessagesSearchConversationsBuilder) Count(v int) {
	b.Params["count"] = v
}

// Extended '1' — return extra information about users and communities
func (b *MessagesSearchConversationsBuilder) Extended(v bool) {
	b.Params["extended"] = v
}

// Fields Profile fields to return.
func (b *MessagesSearchConversationsBuilder) Fields(v []string) {
	b.Params["fields"] = v
}

// GroupID Group ID (for group messages with user access token)
func (b *MessagesSearchConversationsBuilder) GroupID(v int) {
	b.Params["group_id"] = v
}

// MessagesSendBuilder builder
//
// Sends a message.
//
// https://vk.com/dev/messages.send
type MessagesSendBuilder struct {
	api.Params
}

// NewMessagesSendBuilder func
func NewMessagesSendBuilder() *MessagesSendBuilder {
	return &MessagesSendBuilder{api.Params{}}
}

// UserID User ID (by default — current user).
func (b *MessagesSendBuilder) UserID(v int) {
	b.Params["user_id"] = v
}

// RandomID Unique identifier to avoid resending the message.
func (b *MessagesSendBuilder) RandomID(v int) {
	b.Params["random_id"] = v
}

// PeerID Destination ID. For user: 'User ID', e.g. '12345'. For chat:
// '2000000000' + 'chat_id', e.g. '2000000001'. For community:
// '- community ID', e.g. '-12345'.
func (b *MessagesSendBuilder) PeerID(v int) {
	b.Params["peer_id"] = v
}

// Domain User's short address (for example, 'illarionov').
func (b *MessagesSendBuilder) Domain(v string) {
	b.Params["domain"] = v
}

// ChatID ID of conversation the message will relate to.
func (b *MessagesSendBuilder) ChatID(v int) {
	b.Params["chat_id"] = v
}

// UserIDs IDs of message recipients (if new conversation shall be started).
func (b *MessagesSendBuilder) UserIDs(v []int) {
	b.Params["user_ids"] = v
}

// Message (Required if 'attachments' is not set.) Text of the message.
func (b *MessagesSendBuilder) Message(v string) {
	b.Params["message"] = v
}

// Lat Geographical latitude of a check-in, in degrees (from -90 to 90).
func (b *MessagesSendBuilder) Lat(v float64) {
	b.Params["lat"] = v
}

// Long Geographical longitude of a check-in, in degrees (from -180 to 180).
func (b *MessagesSendBuilder) Long(v float64) {
	b.Params["long"] = v
}

// Attachment (Required if 'message' is not set.) List of objects attached to
// the message, separated by commas, in the following format:
// "<owner_id>_<media_id>", '' — Type of media attachment: 'photo' — photo,
// 'video' — video, 'audio' — audio, 'doc' — document, 'wall' — wall post,
// '<owner_id>' — ID of the media attachment owner. '<media_id>' — media
// attachment ID. Example: "photo100172_166443618"
func (b *MessagesSendBuilder) Attachment(v interface{}) {
	b.Params["attachment"] = v
}

// ReplyTo parameter
func (b *MessagesSendBuilder) ReplyTo(v int) {
	b.Params["reply_to"] = v
}

// ForwardMessages ID of forwarded messages, separated with a comma. Listed
// messages of the sender will be shown in the message body at the
// recipient's. Example: "123,431,544"
func (b *MessagesSendBuilder) ForwardMessages(v []int) {
	b.Params["forward_messages"] = v
}

// Forward parameter
func (b *MessagesSendBuilder) Forward(v string) {
	b.Params["forward"] = v
}

// StickerID Sticker id.
func (b *MessagesSendBuilder) StickerID(v int) {
	b.Params["sticker_id"] = v
}

// GroupID Group ID (for group messages with group access token)
func (b *MessagesSendBuilder) GroupID(v int) {
	b.Params["group_id"] = v
}

// Keyboard parameter.
// https://vk.com/dev/bots_docs_3
func (b *MessagesSendBuilder) Keyboard(v interface{}) {
	b.Params["keyboard"] = v
}

// Template parameter.
// https://vk.com/dev/bot_docs_templates
func (b *MessagesSendBuilder) Template(v interface{}) {
	b.Params["template"] = v
}

// Payload parameter
func (b *MessagesSendBuilder) Payload(v string) {
	b.Params["payload"] = v
}

// DontParseLinks parameter
func (b *MessagesSendBuilder) DontParseLinks(v bool) {
	b.Params["dont_parse_links"] = v
}

// DisableMentions parameter
func (b *MessagesSendBuilder) DisableMentions(v bool) {
	b.Params["disable_mentions"] = v
}

// Intent parameter
func (b *MessagesSendBuilder) Intent(v string) {
	b.Params["intent"] = v
}

// MessagesSetActivityBuilder builder
//
// Changes the status of a user as typing in a conversation.
//
// https://vk.com/dev/messages.setActivity
type MessagesSetActivityBuilder struct {
	api.Params
}

// NewMessagesSetActivityBuilder func
func NewMessagesSetActivityBuilder() *MessagesSetActivityBuilder {
	return &MessagesSetActivityBuilder{api.Params{}}
}

// UserID User ID.
func (b *MessagesSetActivityBuilder) UserID(v int) {
	b.Params["user_id"] = v
}

// Type 'typing' — user has started to type.
func (b *MessagesSetActivityBuilder) Type(v string) {
	b.Params["type"] = v
}

// PeerID Destination ID. For user: 'User ID', e.g. '12345'.
// For chat: '2000000000' + 'chat_id', e.g. '2000000001'.
// For community: '- community ID', e.g. '-12345'.
func (b *MessagesSetActivityBuilder) PeerID(v int) {
	b.Params["peer_id"] = v
}

// GroupID Group ID (for group messages with group access token)
func (b *MessagesSetActivityBuilder) GroupID(v int) {
	b.Params["group_id"] = v
}

// MessagesSetChatPhotoBuilder builder
//
// Sets a previously-uploaded picture as the cover picture of a chat.
//
// https://vk.com/dev/messages.setChatPhoto
type MessagesSetChatPhotoBuilder struct {
	api.Params
}

// NewMessagesSetChatPhotoBuilder func
func NewMessagesSetChatPhotoBuilder() *MessagesSetChatPhotoBuilder {
	return &MessagesSetChatPhotoBuilder{api.Params{}}
}

// File Upload URL from the 'response' field returned by the
// [vk.com/dev/photos.getChatUploadServer|photos.getChatUploadServer]
// method upon successfully uploading an image.
func (b *MessagesSetChatPhotoBuilder) File(v string) {
	b.Params["file"] = v
}

// MessagesUnpinBuilder builder
//
// https://vk.com/dev/messages.unpin
type MessagesUnpinBuilder struct {
	api.Params
}

// NewMessagesUnpinBuilder func
func NewMessagesUnpinBuilder() *MessagesUnpinBuilder {
	return &MessagesUnpinBuilder{api.Params{}}
}

// PeerID parameter
func (b *MessagesUnpinBuilder) PeerID(v int) {
	b.Params["peer_id"] = v
}

// GroupID parameter
func (b *MessagesUnpinBuilder) GroupID(v int) {
	b.Params["group_id"] = v
}
