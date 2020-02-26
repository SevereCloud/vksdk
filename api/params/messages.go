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
func (b *MessagesAddChatUserBuilder) ChatID(v int) *MessagesAddChatUserBuilder{
	b.Params["chat_id"] = v
	return b
}

// UserID ID of the user to be added to the chat.
func (b *MessagesAddChatUserBuilder) UserID(v int) *MessagesAddChatUserBuilder{
	b.Params["user_id"] = v
	return b
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
func (b *MessagesAllowMessagesFromGroupBuilder) GroupID(v int) *MessagesAllowMessagesFromGroupBuilder{
	b.Params["group_id"] = v
	return b
}

// Key parameter
func (b *MessagesAllowMessagesFromGroupBuilder) Key(v string) *MessagesAllowMessagesFromGroupBuilder{
	b.Params["key"] = v
	return b
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
func (b *MessagesCreateChatBuilder) UserIDs(v []int) *MessagesCreateChatBuilder{
	b.Params["user_ids"] = v
	return b
}

// Title Chat title.
func (b *MessagesCreateChatBuilder) Title(v string) *MessagesCreateChatBuilder{
	b.Params["title"] = v
	return b
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
func (b *MessagesDeleteBuilder) MessageIDs(v []int) *MessagesDeleteBuilder{
	b.Params["message_ids"] = v
	return b
}

// Spam parameter 
//
// * '1' — to mark message as spam.
func (b *MessagesDeleteBuilder) Spam(v bool) *MessagesDeleteBuilder{
	b.Params["spam"] = v
	return b
}

// GroupID Group ID (for group messages with user access token)
func (b *MessagesDeleteBuilder) GroupID(v int) *MessagesDeleteBuilder{
	b.Params["group_id"] = v
	return b
}

// DeleteForAll parameter 
//
// * '1' — delete message for for all.
func (b *MessagesDeleteBuilder) DeleteForAll(v bool) *MessagesDeleteBuilder{
	b.Params["delete_for_all"] = v
	return b
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
func (b *MessagesDeleteChatPhotoBuilder) ChatID(v int) *MessagesDeleteChatPhotoBuilder{
	b.Params["chat_id"] = v
	return b
}

// GroupID parameter
func (b *MessagesDeleteChatPhotoBuilder) GroupID(v int) *MessagesDeleteChatPhotoBuilder{
	b.Params["group_id"] = v
	return b
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

// UserID User ID.
// To clear a chat history use 'chat_id'
func (b *MessagesDeleteConversationBuilder) UserID(v int) *MessagesDeleteConversationBuilder{
	b.Params["user_id"] = v
	return b
}

// PeerID Destination ID.
// For user
//
// *  'User ID', e.g. '12345'.
// For chat
//
// *  '2000000000' + 'chat_id', e.g. '2000000001'.
// For community
//
// *  '- community ID', e.g. '-12345'.
func (b *MessagesDeleteConversationBuilder) PeerID(v int) *MessagesDeleteConversationBuilder{
	b.Params["peer_id"] = v
	return b
}

// GroupID Group ID (for group messages with user access token)
func (b *MessagesDeleteConversationBuilder) GroupID(v int) *MessagesDeleteConversationBuilder{
	b.Params["group_id"] = v
	return b
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
func (b *MessagesDenyMessagesFromGroupBuilder) GroupID(v int) *MessagesDenyMessagesFromGroupBuilder{
	b.Params["group_id"] = v
	return b
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

// PeerID Destination ID.
// For user
//
// *  'User ID', e.g. '12345'.
// For chat
//
// *  '2000000000' + 'chat_id', e.g. '2000000001'.
// For community
//
// *  '- community ID', e.g. '-12345'.
func (b *MessagesEditBuilder) PeerID(v int) *MessagesEditBuilder{
	b.Params["peer_id"] = v
	return b
}

// Message (Required if 'attachments' is not set.)
// Text of the message.
func (b *MessagesEditBuilder) Message(v string) *MessagesEditBuilder{
	b.Params["message"] = v
	return b
}

// MessageID parameter
func (b *MessagesEditBuilder) MessageID(v int) *MessagesEditBuilder{
	b.Params["message_id"] = v
	return b
}

// Lat Geographical latitude of a check-in, in degrees (from -90 to 90).
func (b *MessagesEditBuilder) Lat(v float64) *MessagesEditBuilder{
	b.Params["lat"] = v
	return b
}

// Long Geographical longitude of a check-in, in degrees (from -180 to 180).
func (b *MessagesEditBuilder) Long(v float64) *MessagesEditBuilder{
	b.Params["long"] = v
	return b
}

// Attachment (Required if 'message' is not set.)
// List of objects attached to the message, separated by commas, in the following format: <owner_id>_<media_id>
//
// * '' — Type of media attachment:
//
// * 'photo' — photo
//
// * 'video' — video
//
// * 'audio' — audio
//
// * 'doc' — document
//
// * 'wall' — wall post
//
// * '<owner_id>' — ID of the media attachment owner.
//
// * '<media_id>' — media attachment ID.
// Example: photo100172_166443618
func (b *MessagesEditBuilder) Attachment(v string) *MessagesEditBuilder{
	b.Params["attachment"] = v
	return b
}

// KeepForwardMessages parameter 
//
// * '1' — to keep forwarded, messages.
func (b *MessagesEditBuilder) KeepForwardMessages(v bool) *MessagesEditBuilder{
	b.Params["keep_forward_messages"] = v
	return b
}

// KeepSnippets parameter 
//
// * '1' — to keep attached snippets.
func (b *MessagesEditBuilder) KeepSnippets(v bool) *MessagesEditBuilder{
	b.Params["keep_snippets"] = v
	return b
}

// GroupID Group ID (for group messages with user access token)
func (b *MessagesEditBuilder) GroupID(v int) *MessagesEditBuilder{
	b.Params["group_id"] = v
	return b
}

// DontParseLinks parameter
func (b *MessagesEditBuilder) DontParseLinks(v bool) *MessagesEditBuilder{
	b.Params["dont_parse_links"] = v
	return b
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
func (b *MessagesEditChatBuilder) ChatID(v int) *MessagesEditChatBuilder{
	b.Params["chat_id"] = v
	return b
}

// Title New title of the chat.
func (b *MessagesEditChatBuilder) Title(v string) *MessagesEditChatBuilder{
	b.Params["title"] = v
	return b
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

// PeerID Destination ID.
// For user
//
// *  'User ID', e.g. '12345'.
// For chat
//
// *  '2000000000' + 'chat_id', e.g. '2000000001'.
// For community
//
// *  '- community ID', e.g. '-12345'.
func (b *MessagesGetByConversationMessageIDBuilder) PeerID(v int) *MessagesGetByConversationMessageIDBuilder{
	b.Params["peer_id"] = v
	return b
}

// ConversationMessageIDs Conversation message IDs.
func (b *MessagesGetByConversationMessageIDBuilder) ConversationMessageIDs(v []int) *MessagesGetByConversationMessageIDBuilder{
	b.Params["conversation_message_ids"] = v
	return b
}

// Extended Information whether the response should be extended
func (b *MessagesGetByConversationMessageIDBuilder) Extended(v bool) *MessagesGetByConversationMessageIDBuilder{
	b.Params["extended"] = v
	return b
}

// Fields Profile fields to return.
func (b *MessagesGetByConversationMessageIDBuilder) Fields(v []string) *MessagesGetByConversationMessageIDBuilder{
	b.Params["fields"] = v
	return b
}

// GroupID Group ID (for group messages with group access token)
func (b *MessagesGetByConversationMessageIDBuilder) GroupID(v int) *MessagesGetByConversationMessageIDBuilder{
	b.Params["group_id"] = v
	return b
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
func (b *MessagesGetByIDBuilder) MessageIDs(v []int) *MessagesGetByIDBuilder{
	b.Params["message_ids"] = v
	return b
}

// PreviewLength Number of characters after which to truncate a previewed message.
// To preview the full message, specify '0'.
//
// NOTE: Messages are not truncated by default.
// Messages are truncated by words.
func (b *MessagesGetByIDBuilder) PreviewLength(v int) *MessagesGetByIDBuilder{
	b.Params["preview_length"] = v
	return b
}

// Extended Information whether the response should be extended
func (b *MessagesGetByIDBuilder) Extended(v bool) *MessagesGetByIDBuilder{
	b.Params["extended"] = v
	return b
}

// Fields Profile fields to return.
func (b *MessagesGetByIDBuilder) Fields(v []string) *MessagesGetByIDBuilder{
	b.Params["fields"] = v
	return b
}

// GroupID Group ID (for group messages with group access token)
func (b *MessagesGetByIDBuilder) GroupID(v int) *MessagesGetByIDBuilder{
	b.Params["group_id"] = v
	return b
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
func (b *MessagesGetChatPreviewBuilder) PeerID(v int) *MessagesGetChatPreviewBuilder{
	b.Params["peer_id"] = v
	return b
}

// Link Invitation link.
func (b *MessagesGetChatPreviewBuilder) Link(v string) *MessagesGetChatPreviewBuilder{
	b.Params["link"] = v
	return b
}

// Fields Profile fields to return.
func (b *MessagesGetChatPreviewBuilder) Fields(v []string) *MessagesGetChatPreviewBuilder{
	b.Params["fields"] = v
	return b
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
func (b *MessagesGetConversationMembersBuilder) PeerID(v int) *MessagesGetConversationMembersBuilder{
	b.Params["peer_id"] = v
	return b
}

// Fields Profile fields to return.
func (b *MessagesGetConversationMembersBuilder) Fields(v []string) *MessagesGetConversationMembersBuilder{
	b.Params["fields"] = v
	return b
}

// GroupID Group ID (for group messages with group access token)
func (b *MessagesGetConversationMembersBuilder) GroupID(v int) *MessagesGetConversationMembersBuilder{
	b.Params["group_id"] = v
	return b
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
func (b *MessagesGetConversationsBuilder) Offset(v int) *MessagesGetConversationsBuilder{
	b.Params["offset"] = v
	return b
}

// Count Number of conversations to return.
func (b *MessagesGetConversationsBuilder) Count(v int) *MessagesGetConversationsBuilder{
	b.Params["count"] = v
	return b
}

// Filter Filter to apply:
//
// * 'all' — all conversations
//
// * 'unread' — conversations with unread messages
//
// * 'important' — conversations, marked as important (only for community messages)
//
// * 'unanswered' — conversations, marked as unanswered (only for community messages)
func (b *MessagesGetConversationsBuilder) Filter(v string) *MessagesGetConversationsBuilder{
	b.Params["filter"] = v
	return b
}

// Extended parameter 
//
// * '1' — return extra information about users and communities
func (b *MessagesGetConversationsBuilder) Extended(v bool) *MessagesGetConversationsBuilder{
	b.Params["extended"] = v
	return b
}

// StartMessageID ID of the message from what to return dialogs.
func (b *MessagesGetConversationsBuilder) StartMessageID(v int) *MessagesGetConversationsBuilder{
	b.Params["start_message_id"] = v
	return b
}

// Fields Profile and communities fields to return.
func (b *MessagesGetConversationsBuilder) Fields(v []string) *MessagesGetConversationsBuilder{
	b.Params["fields"] = v
	return b
}

// GroupID Group ID (for group messages with group access token)
func (b *MessagesGetConversationsBuilder) GroupID(v int) *MessagesGetConversationsBuilder{
	b.Params["group_id"] = v
	return b
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

// PeerIDs Destination IDs. For user
//
// *  'User ID', e.g. '12345'.
// For chat
//
// *  '2000000000' + 'chat_id', e.g. '2000000001'.
// For community
//
// *  '- community ID', e.g. '-12345'.
func (b *MessagesGetConversationsByIDBuilder) PeerIDs(v []int) *MessagesGetConversationsByIDBuilder{
	b.Params["peer_ids"] = v
	return b
}

// Extended Return extended properties
func (b *MessagesGetConversationsByIDBuilder) Extended(v bool) *MessagesGetConversationsByIDBuilder{
	b.Params["extended"] = v
	return b
}

// Fields Profile and communities fields to return.
func (b *MessagesGetConversationsByIDBuilder) Fields(v []string) *MessagesGetConversationsByIDBuilder{
	b.Params["fields"] = v
	return b
}

// GroupID Group ID (for group messages with group access token)
func (b *MessagesGetConversationsByIDBuilder) GroupID(v int) *MessagesGetConversationsByIDBuilder{
	b.Params["group_id"] = v
	return b
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
func (b *MessagesGetHistoryBuilder) Offset(v int) *MessagesGetHistoryBuilder{
	b.Params["offset"] = v
	return b
}

// Count Number of messages to return.
func (b *MessagesGetHistoryBuilder) Count(v int) *MessagesGetHistoryBuilder{
	b.Params["count"] = v
	return b
}

// UserID ID of the user whose message history you want to return.
func (b *MessagesGetHistoryBuilder) UserID(v int) *MessagesGetHistoryBuilder{
	b.Params["user_id"] = v
	return b
}

// PeerID parameter
func (b *MessagesGetHistoryBuilder) PeerID(v int) *MessagesGetHistoryBuilder{
	b.Params["peer_id"] = v
	return b
}

// StartMessageID Starting message ID from which to return history.
func (b *MessagesGetHistoryBuilder) StartMessageID(v int) *MessagesGetHistoryBuilder{
	b.Params["start_message_id"] = v
	return b
}

// Rev Sort order:
//
// * '1' — return messages in chronological order.
//
// * '0' — return messages in reverse chronological order.
func (b *MessagesGetHistoryBuilder) Rev(v int) *MessagesGetHistoryBuilder{
	b.Params["rev"] = v
	return b
}

// Extended Information whether the response should be extended
func (b *MessagesGetHistoryBuilder) Extended(v bool) *MessagesGetHistoryBuilder{
	b.Params["extended"] = v
	return b
}

// Fields Profile fields to return.
func (b *MessagesGetHistoryBuilder) Fields(v []string) *MessagesGetHistoryBuilder{
	b.Params["fields"] = v
	return b
}

// GroupID Group ID (for group messages with group access token)
func (b *MessagesGetHistoryBuilder) GroupID(v int) *MessagesGetHistoryBuilder{
	b.Params["group_id"] = v
	return b
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

// PeerID Peer ID.
// For group chat
//
// *  '2000000000 + chat ID'  For community
//
// *  '-community ID'
func (b *MessagesGetHistoryAttachmentsBuilder) PeerID(v int) *MessagesGetHistoryAttachmentsBuilder{
	b.Params["peer_id"] = v
	return b
}

// MediaType  Type of media files to return:
//
// * photo
//
// * video
//
// * audio
//
// * doc
//
// * link
//
// * market
//
// * wall
//
// * share
func (b *MessagesGetHistoryAttachmentsBuilder) MediaType(v string) *MessagesGetHistoryAttachmentsBuilder{
	b.Params["media_type"] = v
	return b
}

// StartFrom Message ID to start return results from.
func (b *MessagesGetHistoryAttachmentsBuilder) StartFrom(v string) *MessagesGetHistoryAttachmentsBuilder{
	b.Params["start_from"] = v
	return b
}

// Count Number of objects to return.
func (b *MessagesGetHistoryAttachmentsBuilder) Count(v int) *MessagesGetHistoryAttachmentsBuilder{
	b.Params["count"] = v
	return b
}

// PhotoSizes parameter 
//
// * '1' — to return photo sizes in a
func (b *MessagesGetHistoryAttachmentsBuilder) PhotoSizes(v bool) *MessagesGetHistoryAttachmentsBuilder{
	b.Params["photo_sizes"] = v
	return b
}

// Fields Additional profile [vk.com/dev/fields|fields] to return.
func (b *MessagesGetHistoryAttachmentsBuilder) Fields(v []string) *MessagesGetHistoryAttachmentsBuilder{
	b.Params["fields"] = v
	return b
}

// GroupID Group ID (for group messages with group access token)
func (b *MessagesGetHistoryAttachmentsBuilder) GroupID(v int) *MessagesGetHistoryAttachmentsBuilder{
	b.Params["group_id"] = v
	return b
}

// PreserveOrder parameter
func (b *MessagesGetHistoryAttachmentsBuilder) PreserveOrder(v bool) *MessagesGetHistoryAttachmentsBuilder{
	b.Params["preserve_order"] = v
	return b
}

// MaxForwardsLevel parameter
func (b *MessagesGetHistoryAttachmentsBuilder) MaxForwardsLevel(v int) *MessagesGetHistoryAttachmentsBuilder{
	b.Params["max_forwards_level"] = v
	return b
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
func (b *MessagesGetInviteLinkBuilder) PeerID(v int) *MessagesGetInviteLinkBuilder{
	b.Params["peer_id"] = v
	return b
}

// Reset 1 — to generate new link (revoke previous), 0 — to return previous link.
func (b *MessagesGetInviteLinkBuilder) Reset(v bool) *MessagesGetInviteLinkBuilder{
	b.Params["reset"] = v
	return b
}

// GroupID Group ID
func (b *MessagesGetInviteLinkBuilder) GroupID(v int) *MessagesGetInviteLinkBuilder{
	b.Params["group_id"] = v
	return b
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
func (b *MessagesGetLastActivityBuilder) UserID(v int) *MessagesGetLastActivityBuilder{
	b.Params["user_id"] = v
	return b
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

// Ts Last value of the 'ts' parameter returned from the Long Poll server or by using [vk.com/dev/messages.getLongPollHistory|messages.getLongPollHistory] method.
func (b *MessagesGetLongPollHistoryBuilder) Ts(v int) *MessagesGetLongPollHistoryBuilder{
	b.Params["ts"] = v
	return b
}

// Pts Lsat value of 'pts' parameter returned from the Long Poll server or by using [vk.com/dev/messages.getLongPollHistory|messages.getLongPollHistory] method.
func (b *MessagesGetLongPollHistoryBuilder) Pts(v int) *MessagesGetLongPollHistoryBuilder{
	b.Params["pts"] = v
	return b
}

// PreviewLength Number of characters after which to truncate a previewed message.
// To preview the full message, specify '0'.
//
// NOTE: Messages are not truncated by default.
// Messages are truncated by words.
func (b *MessagesGetLongPollHistoryBuilder) PreviewLength(v int) *MessagesGetLongPollHistoryBuilder{
	b.Params["preview_length"] = v
	return b
}

// Onlines parameter 
//
// * '1' — to return history with online users only.
func (b *MessagesGetLongPollHistoryBuilder) Onlines(v bool) *MessagesGetLongPollHistoryBuilder{
	b.Params["onlines"] = v
	return b
}

// Fields Additional profile [vk.com/dev/fields|fields] to return.
func (b *MessagesGetLongPollHistoryBuilder) Fields(v []string) *MessagesGetLongPollHistoryBuilder{
	b.Params["fields"] = v
	return b
}

// EventsLimit Maximum number of events to return.
func (b *MessagesGetLongPollHistoryBuilder) EventsLimit(v int) *MessagesGetLongPollHistoryBuilder{
	b.Params["events_limit"] = v
	return b
}

// MsgsLimit Maximum number of messages to return.
func (b *MessagesGetLongPollHistoryBuilder) MsgsLimit(v int) *MessagesGetLongPollHistoryBuilder{
	b.Params["msgs_limit"] = v
	return b
}

// MaxMsgID Maximum ID of the message among existing ones in the local copy.
// Both messages received with API methods (for example ), and data received from a Long Poll server (events with code 4) are taken into account.
func (b *MessagesGetLongPollHistoryBuilder) MaxMsgID(v int) *MessagesGetLongPollHistoryBuilder{
	b.Params["max_msg_id"] = v
	return b
}

// GroupID Group ID (for group messages with user access token)
func (b *MessagesGetLongPollHistoryBuilder) GroupID(v int) *MessagesGetLongPollHistoryBuilder{
	b.Params["group_id"] = v
	return b
}

// LpVersion parameter
func (b *MessagesGetLongPollHistoryBuilder) LpVersion(v int) *MessagesGetLongPollHistoryBuilder{
	b.Params["lp_version"] = v
	return b
}

// LastN parameter
func (b *MessagesGetLongPollHistoryBuilder) LastN(v int) *MessagesGetLongPollHistoryBuilder{
	b.Params["last_n"] = v
	return b
}

// Credentials parameter
func (b *MessagesGetLongPollHistoryBuilder) Credentials(v bool) *MessagesGetLongPollHistoryBuilder{
	b.Params["credentials"] = v
	return b
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

// NeedPts parameter 
//
// * '1' — to return the 'pts' field, needed for the [vk.com/dev/messages.getLongPollHistory|messages.getLongPollHistory] method.
func (b *MessagesGetLongPollServerBuilder) NeedPts(v bool) *MessagesGetLongPollServerBuilder{
	b.Params["need_pts"] = v
	return b
}

// GroupID Group ID (for group messages with user access token)
func (b *MessagesGetLongPollServerBuilder) GroupID(v int) *MessagesGetLongPollServerBuilder{
	b.Params["group_id"] = v
	return b
}

// LpVersion Long poll version
func (b *MessagesGetLongPollServerBuilder) LpVersion(v int) *MessagesGetLongPollServerBuilder{
	b.Params["lp_version"] = v
	return b
}

// MessagesIsMessagesFromGroupAllowedBuilder builder
//
// Returns information whether sending messages from the community to current user is allowed.
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
func (b *MessagesIsMessagesFromGroupAllowedBuilder) GroupID(v int) *MessagesIsMessagesFromGroupAllowedBuilder{
	b.Params["group_id"] = v
	return b
}

// UserID User ID.
func (b *MessagesIsMessagesFromGroupAllowedBuilder) UserID(v int) *MessagesIsMessagesFromGroupAllowedBuilder{
	b.Params["user_id"] = v
	return b
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
func (b *MessagesJoinChatByInviteLinkBuilder) Link(v string) *MessagesJoinChatByInviteLinkBuilder{
	b.Params["link"] = v
	return b
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
func (b *MessagesMarkAsAnsweredConversationBuilder) PeerID(v int) *MessagesMarkAsAnsweredConversationBuilder{
	b.Params["peer_id"] = v
	return b
}

// Answered parameter 
//
// * '1' — to mark as answered
//
// * '0' — to remove the mark
func (b *MessagesMarkAsAnsweredConversationBuilder) Answered(v bool) *MessagesMarkAsAnsweredConversationBuilder{
	b.Params["answered"] = v
	return b
}

// GroupID Group ID (for group messages with group access token)
func (b *MessagesMarkAsAnsweredConversationBuilder) GroupID(v int) *MessagesMarkAsAnsweredConversationBuilder{
	b.Params["group_id"] = v
	return b
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
func (b *MessagesMarkAsImportantBuilder) MessageIDs(v []int) *MessagesMarkAsImportantBuilder{
	b.Params["message_ids"] = v
	return b
}

// Important parameter 
//
// * '1' — to add a star (mark as important)
//
// * '0' — to remove the star
func (b *MessagesMarkAsImportantBuilder) Important(v int) *MessagesMarkAsImportantBuilder{
	b.Params["important"] = v
	return b
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
func (b *MessagesMarkAsImportantConversationBuilder) PeerID(v int) *MessagesMarkAsImportantConversationBuilder{
	b.Params["peer_id"] = v
	return b
}

// Important parameter 
//
// * '1' — to add a star (mark as important)
//
// * '0' — to remove the star
func (b *MessagesMarkAsImportantConversationBuilder) Important(v bool) *MessagesMarkAsImportantConversationBuilder{
	b.Params["important"] = v
	return b
}

// GroupID Group ID (for group messages with group access token)
func (b *MessagesMarkAsImportantConversationBuilder) GroupID(v int) *MessagesMarkAsImportantConversationBuilder{
	b.Params["group_id"] = v
	return b
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
func (b *MessagesMarkAsReadBuilder) MessageIDs(v []int) *MessagesMarkAsReadBuilder{
	b.Params["message_ids"] = v
	return b
}

// PeerID Destination ID.
// For user
//
// *  'User ID', e.g. '12345'.
// For chat
//
// *  '2000000000' + 'chat_id', e.g. '2000000001'.
// For community
//
// *  '- community ID', e.g. '-12345'.
func (b *MessagesMarkAsReadBuilder) PeerID(v int) *MessagesMarkAsReadBuilder{
	b.Params["peer_id"] = v
	return b
}

// StartMessageID Message ID to start from.
func (b *MessagesMarkAsReadBuilder) StartMessageID(v int) *MessagesMarkAsReadBuilder{
	b.Params["start_message_id"] = v
	return b
}

// GroupID Group ID (for group messages with user access token)
func (b *MessagesMarkAsReadBuilder) GroupID(v int) *MessagesMarkAsReadBuilder{
	b.Params["group_id"] = v
	return b
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

// PeerID Destination ID.
// For user
//
// *  'User ID', e.g. '12345'.
// For chat
//
// *  '2000000000' + 'Chat ID', e.g. '2000000001'.
// For community
//
// *  '- Community ID', e.g. '-12345'.
func (b *MessagesPinBuilder) PeerID(v int) *MessagesPinBuilder{
	b.Params["peer_id"] = v
	return b
}

// MessageID parameter
func (b *MessagesPinBuilder) MessageID(v int) *MessagesPinBuilder{
	b.Params["message_id"] = v
	return b
}

// MessagesRemoveChatUserBuilder builder
//
// Allows the current user to leave a chat or, if the current user started the chat, allows the user to remove another user from the chat.
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
func (b *MessagesRemoveChatUserBuilder) ChatID(v int) *MessagesRemoveChatUserBuilder{
	b.Params["chat_id"] = v
	return b
}

// UserID ID of the user to be removed from the chat.
func (b *MessagesRemoveChatUserBuilder) UserID(v int) *MessagesRemoveChatUserBuilder{
	b.Params["user_id"] = v
	return b
}

// MemberID parameter
func (b *MessagesRemoveChatUserBuilder) MemberID(v int) *MessagesRemoveChatUserBuilder{
	b.Params["member_id"] = v
	return b
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
func (b *MessagesRestoreBuilder) MessageID(v int) *MessagesRestoreBuilder{
	b.Params["message_id"] = v
	return b
}

// GroupID Group ID (for group messages with user access token)
func (b *MessagesRestoreBuilder) GroupID(v int) *MessagesRestoreBuilder{
	b.Params["group_id"] = v
	return b
}

// MessagesSearchBuilder builder
//
// Returns a list of the current user's private messages that match search criteria.
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
func (b *MessagesSearchBuilder) Q(v string) *MessagesSearchBuilder{
	b.Params["q"] = v
	return b
}

// PeerID Destination ID.
// For user
//
// *  'User ID', e.g. '12345'.
// For chat
//
// *  '2000000000' + 'chat_id', e.g. '2000000001'.
// For community
//
// *  '- community ID', e.g. '-12345'.
func (b *MessagesSearchBuilder) PeerID(v int) *MessagesSearchBuilder{
	b.Params["peer_id"] = v
	return b
}

// Date Date to search message before in Unixtime.
func (b *MessagesSearchBuilder) Date(v int) *MessagesSearchBuilder{
	b.Params["date"] = v
	return b
}

// PreviewLength Number of characters after which to truncate a previewed message.
// To preview the full message, specify '0'.
//
// NOTE: Messages are not truncated by default.
// Messages are truncated by words.
func (b *MessagesSearchBuilder) PreviewLength(v int) *MessagesSearchBuilder{
	b.Params["preview_length"] = v
	return b
}

// Offset Offset needed to return a specific subset of messages.
func (b *MessagesSearchBuilder) Offset(v int) *MessagesSearchBuilder{
	b.Params["offset"] = v
	return b
}

// Count Number of messages to return.
func (b *MessagesSearchBuilder) Count(v int) *MessagesSearchBuilder{
	b.Params["count"] = v
	return b
}

// Extended parameter
func (b *MessagesSearchBuilder) Extended(v bool) *MessagesSearchBuilder{
	b.Params["extended"] = v
	return b
}

// Fields parameter
func (b *MessagesSearchBuilder) Fields(v []string) *MessagesSearchBuilder{
	b.Params["fields"] = v
	return b
}

// GroupID Group ID (for group messages with group access token)
func (b *MessagesSearchBuilder) GroupID(v int) *MessagesSearchBuilder{
	b.Params["group_id"] = v
	return b
}

// MessagesSearchConversationsBuilder builder
//
// Returns a list of the current user's conversations that match search criteria.
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
func (b *MessagesSearchConversationsBuilder) Q(v string) *MessagesSearchConversationsBuilder{
	b.Params["q"] = v
	return b
}

// Count Maximum number of results.
func (b *MessagesSearchConversationsBuilder) Count(v int) *MessagesSearchConversationsBuilder{
	b.Params["count"] = v
	return b
}

// Extended parameter 
//
// * '1' — return extra information about users and communities
func (b *MessagesSearchConversationsBuilder) Extended(v bool) *MessagesSearchConversationsBuilder{
	b.Params["extended"] = v
	return b
}

// Fields Profile fields to return.
func (b *MessagesSearchConversationsBuilder) Fields(v []string) *MessagesSearchConversationsBuilder{
	b.Params["fields"] = v
	return b
}

// GroupID Group ID (for group messages with user access token)
func (b *MessagesSearchConversationsBuilder) GroupID(v int) *MessagesSearchConversationsBuilder{
	b.Params["group_id"] = v
	return b
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
func (b *MessagesSendBuilder) UserID(v int) *MessagesSendBuilder{
	b.Params["user_id"] = v
	return b
}

// RandomID Unique identifier to avoid resending the message.
func (b *MessagesSendBuilder) RandomID(v int) *MessagesSendBuilder{
	b.Params["random_id"] = v
	return b
}

// PeerID Destination ID.
// For user
//
// *  'User ID', e.g. '12345'.
// For chat
//
// *  '2000000000' + 'chat_id', e.g. '2000000001'.
// For community
//
// *  '- community ID', e.g. '-12345'.
func (b *MessagesSendBuilder) PeerID(v int) *MessagesSendBuilder{
	b.Params["peer_id"] = v
	return b
}

// Domain User's short address for example
//
// *  'illarionov'.
func (b *MessagesSendBuilder) Domain(v string) *MessagesSendBuilder{
	b.Params["domain"] = v
	return b
}

// ChatID ID of conversation the message will relate to.
func (b *MessagesSendBuilder) ChatID(v int) *MessagesSendBuilder{
	b.Params["chat_id"] = v
	return b
}

// UserIDs IDs of message recipients (if new conversation shall be started).
func (b *MessagesSendBuilder) UserIDs(v []int) *MessagesSendBuilder{
	b.Params["user_ids"] = v
	return b
}

// Message (Required if 'attachments' is not set.)
// Text of the message.
func (b *MessagesSendBuilder) Message(v string) *MessagesSendBuilder{
	b.Params["message"] = v
	return b
}

// Lat Geographical latitude of a check-in, in degrees (from -90 to 90).
func (b *MessagesSendBuilder) Lat(v float64) *MessagesSendBuilder{
	b.Params["lat"] = v
	return b
}

// Long Geographical longitude of a check-in, in degrees (from -180 to 180).
func (b *MessagesSendBuilder) Long(v float64) *MessagesSendBuilder{
	b.Params["long"] = v
	return b
}

// Attachment (Required if 'message' is not set.)
// List of objects attached to the message, separated by commas, in the following format: <owner_id>_<media_id>
//
// * '' — Type of media attachment:
//
// * 'photo' — photo
//
// * 'video' — video
//
// * 'audio' — audio
//
// * 'doc' — document
//
// * 'wall' — wall post
//
// * '<owner_id>' — ID of the media attachment owner.
//
// * '<media_id>' — media attachment ID.
// Example: photo100172_166443618
func (b *MessagesSendBuilder) Attachment(v string) *MessagesSendBuilder{
	b.Params["attachment"] = v
	return b
}

// ReplyTo parameter
func (b *MessagesSendBuilder) ReplyTo(v int) *MessagesSendBuilder{
	b.Params["reply_to"] = v
	return b
}

// ForwardMessages ID of forwarded messages, separated with a comma.
// Listed messages of the sender will be shown in the message body at the recipient's.
// Example: 123,431,544
func (b *MessagesSendBuilder) ForwardMessages(v []int) *MessagesSendBuilder{
	b.Params["forward_messages"] = v
	return b
}

// Forward parameter
func (b *MessagesSendBuilder) Forward(v string) *MessagesSendBuilder{
	b.Params["forward"] = v
	return b
}

// StickerID Sticker id.
func (b *MessagesSendBuilder) StickerID(v int) *MessagesSendBuilder{
	b.Params["sticker_id"] = v
	return b
}

// GroupID Group ID (for group messages with group access token)
func (b *MessagesSendBuilder) GroupID(v int) *MessagesSendBuilder{
	b.Params["group_id"] = v
	return b
}

// Keyboard parameter
func (b *MessagesSendBuilder) Keyboard(v string) *MessagesSendBuilder{
	b.Params["keyboard"] = v
	return b
}

// Payload parameter
func (b *MessagesSendBuilder) Payload(v string) *MessagesSendBuilder{
	b.Params["payload"] = v
	return b
}

// DontParseLinks parameter
func (b *MessagesSendBuilder) DontParseLinks(v bool) *MessagesSendBuilder{
	b.Params["dont_parse_links"] = v
	return b
}

// DisableMentions parameter
func (b *MessagesSendBuilder) DisableMentions(v bool) *MessagesSendBuilder{
	b.Params["disable_mentions"] = v
	return b
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
func (b *MessagesSetActivityBuilder) UserID(v int) *MessagesSetActivityBuilder{
	b.Params["user_id"] = v
	return b
}

// Type parameter 
//
// * 'typing' — user has started to type.
func (b *MessagesSetActivityBuilder) Type(v string) *MessagesSetActivityBuilder{
	b.Params["type"] = v
	return b
}

// PeerID Destination ID.
// For user
//
// *  'User ID', e.g. '12345'.
// For chat
//
// *  '2000000000' + 'chat_id', e.g. '2000000001'.
// For community
//
// *  '- community ID', e.g. '-12345'.
func (b *MessagesSetActivityBuilder) PeerID(v int) *MessagesSetActivityBuilder{
	b.Params["peer_id"] = v
	return b
}

// GroupID Group ID (for group messages with group access token)
func (b *MessagesSetActivityBuilder) GroupID(v int) *MessagesSetActivityBuilder{
	b.Params["group_id"] = v
	return b
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

// File Upload URL from the 'response' field returned by the [vk.com/dev/photos.getChatUploadServer|photos.getChatUploadServer] method upon successfully uploading an image.
func (b *MessagesSetChatPhotoBuilder) File(v string) *MessagesSetChatPhotoBuilder{
	b.Params["file"] = v
	return b
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
func (b *MessagesUnpinBuilder) PeerID(v int) *MessagesUnpinBuilder{
	b.Params["peer_id"] = v
	return b
}

// GroupID parameter
func (b *MessagesUnpinBuilder) GroupID(v int) *MessagesUnpinBuilder{
	b.Params["group_id"] = v
	return b
}
