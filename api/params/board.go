package params // import "github.com/SevereCloud/vksdk/api/params"

import (
	"github.com/SevereCloud/vksdk/api"
)

// BoardAddTopicBuilder builder
//
// Creates a new topic on a community's discussion board.
//
// https://vk.com/dev/board.addTopic
type BoardAddTopicBuilder struct {
	api.Params
}

// NewBoardAddTopicBuilder func
func NewBoardAddTopicBuilder() *BoardAddTopicBuilder {
	return &BoardAddTopicBuilder{api.Params{}}
}

// GroupID ID of the community that owns the discussion board.
func (b *BoardAddTopicBuilder) GroupID(v int) {
	b.Params["group_id"] = v
}

// Title Topic title.
func (b *BoardAddTopicBuilder) Title(v string) {
	b.Params["title"] = v
}

// Text Text of the topic.
func (b *BoardAddTopicBuilder) Text(v string) {
	b.Params["text"] = v
}

// FromGroup For a community: '1' — to post the topic as by the community, '0' — to post the topic as by the user (default)
func (b *BoardAddTopicBuilder) FromGroup(v bool) {
	b.Params["from_group"] = v
}

// Attachments List of media objects attached to the topic, in the following format: "<owner_id>_<media_id>,<owner_id>_<media_id>", '' — Type of media object: 'photo' — photo, 'video' — video, 'audio' — audio, 'doc' — document, '<owner_id>' — ID of the media owner. '<media_id>' — Media ID. Example: "photo100172_166443618,photo66748_265827614", , "NOTE: If you try to attach more than one reference, an error will be thrown.",
func (b *BoardAddTopicBuilder) Attachments(v []string) {
	b.Params["attachments"] = v
}

// BoardCloseTopicBuilder builder
//
// Closes a topic on a community's discussion board so that comments cannot be posted.
//
// https://vk.com/dev/board.closeTopic
type BoardCloseTopicBuilder struct {
	api.Params
}

// NewBoardCloseTopicBuilder func
func NewBoardCloseTopicBuilder() *BoardCloseTopicBuilder {
	return &BoardCloseTopicBuilder{api.Params{}}
}

// GroupID ID of the community that owns the discussion board.
func (b *BoardCloseTopicBuilder) GroupID(v int) {
	b.Params["group_id"] = v
}

// TopicID Topic ID.
func (b *BoardCloseTopicBuilder) TopicID(v int) {
	b.Params["topic_id"] = v
}

// BoardCreateCommentBuilder builder
//
// Adds a comment on a topic on a community's discussion board.
//
// https://vk.com/dev/board.createComment
type BoardCreateCommentBuilder struct {
	api.Params
}

// NewBoardCreateCommentBuilder func
func NewBoardCreateCommentBuilder() *BoardCreateCommentBuilder {
	return &BoardCreateCommentBuilder{api.Params{}}
}

// GroupID ID of the community that owns the discussion board.
func (b *BoardCreateCommentBuilder) GroupID(v int) {
	b.Params["group_id"] = v
}

// TopicID ID of the topic to be commented on.
func (b *BoardCreateCommentBuilder) TopicID(v int) {
	b.Params["topic_id"] = v
}

// Message (Required if 'attachments' is not set.) Text of the comment.
func (b *BoardCreateCommentBuilder) Message(v string) {
	b.Params["message"] = v
}

// Attachments (Required if 'text' is not set.) List of media objects attached to the comment, in the following format: "<owner_id>_<media_id>,<owner_id>_<media_id>", '' — Type of media object: 'photo' — photo, 'video' — video, 'audio' — audio, 'doc' — document, '<owner_id>' — ID of the media owner. '<media_id>' — Media ID.
func (b *BoardCreateCommentBuilder) Attachments(v []string) {
	b.Params["attachments"] = v
}

// FromGroup '1' — to post the comment as by the community, '0' — to post the comment as by the user (default)
func (b *BoardCreateCommentBuilder) FromGroup(v bool) {
	b.Params["from_group"] = v
}

// StickerID Sticker ID.
func (b *BoardCreateCommentBuilder) StickerID(v int) {
	b.Params["sticker_id"] = v
}

// GUID Unique identifier to avoid repeated comments.
func (b *BoardCreateCommentBuilder) GUID(v string) {
	b.Params["guid"] = v
}

// BoardDeleteCommentBuilder builder
//
// Deletes a comment on a topic on a community's discussion board.
//
// https://vk.com/dev/board.deleteComment
type BoardDeleteCommentBuilder struct {
	api.Params
}

// NewBoardDeleteCommentBuilder func
func NewBoardDeleteCommentBuilder() *BoardDeleteCommentBuilder {
	return &BoardDeleteCommentBuilder{api.Params{}}
}

// GroupID ID of the community that owns the discussion board.
func (b *BoardDeleteCommentBuilder) GroupID(v int) {
	b.Params["group_id"] = v
}

// TopicID Topic ID.
func (b *BoardDeleteCommentBuilder) TopicID(v int) {
	b.Params["topic_id"] = v
}

// CommentID Comment ID.
func (b *BoardDeleteCommentBuilder) CommentID(v int) {
	b.Params["comment_id"] = v
}

// BoardDeleteTopicBuilder builder
//
// Deletes a topic from a community's discussion board.
//
// https://vk.com/dev/board.deleteTopic
type BoardDeleteTopicBuilder struct {
	api.Params
}

// NewBoardDeleteTopicBuilder func
func NewBoardDeleteTopicBuilder() *BoardDeleteTopicBuilder {
	return &BoardDeleteTopicBuilder{api.Params{}}
}

// GroupID ID of the community that owns the discussion board.
func (b *BoardDeleteTopicBuilder) GroupID(v int) {
	b.Params["group_id"] = v
}

// TopicID Topic ID.
func (b *BoardDeleteTopicBuilder) TopicID(v int) {
	b.Params["topic_id"] = v
}

// BoardEditCommentBuilder builder
//
// Edits a comment on a topic on a community's discussion board.
//
// https://vk.com/dev/board.editComment
type BoardEditCommentBuilder struct {
	api.Params
}

// NewBoardEditCommentBuilder func
func NewBoardEditCommentBuilder() *BoardEditCommentBuilder {
	return &BoardEditCommentBuilder{api.Params{}}
}

// GroupID ID of the community that owns the discussion board.
func (b *BoardEditCommentBuilder) GroupID(v int) {
	b.Params["group_id"] = v
}

// TopicID Topic ID.
func (b *BoardEditCommentBuilder) TopicID(v int) {
	b.Params["topic_id"] = v
}

// CommentID ID of the comment on the topic.
func (b *BoardEditCommentBuilder) CommentID(v int) {
	b.Params["comment_id"] = v
}

// Message (Required if 'attachments' is not set). New comment text.
func (b *BoardEditCommentBuilder) Message(v string) {
	b.Params["message"] = v
}

// Attachments (Required if 'message' is not set.) List of media objects attached to the comment, in the following format: "<owner_id>_<media_id>,<owner_id>_<media_id>", '' — Type of media object: 'photo' — photo, 'video' — video, 'audio' — audio, 'doc' — document, '<owner_id>' — ID of the media owner. '<media_id>' — Media ID. Example: "photo100172_166443618,photo66748_265827614"
func (b *BoardEditCommentBuilder) Attachments(v []string) {
	b.Params["attachments"] = v
}

// BoardEditTopicBuilder builder
//
// Edits the title of a topic on a community's discussion board.
//
// https://vk.com/dev/board.editTopic
type BoardEditTopicBuilder struct {
	api.Params
}

// NewBoardEditTopicBuilder func
func NewBoardEditTopicBuilder() *BoardEditTopicBuilder {
	return &BoardEditTopicBuilder{api.Params{}}
}

// GroupID ID of the community that owns the discussion board.
func (b *BoardEditTopicBuilder) GroupID(v int) {
	b.Params["group_id"] = v
}

// TopicID Topic ID.
func (b *BoardEditTopicBuilder) TopicID(v int) {
	b.Params["topic_id"] = v
}

// Title New title of the topic.
func (b *BoardEditTopicBuilder) Title(v string) {
	b.Params["title"] = v
}

// BoardFixTopicBuilder builder
//
// Pins a topic (fixes its place) to the top of a community's discussion board.
//
// https://vk.com/dev/board.fixTopic
type BoardFixTopicBuilder struct {
	api.Params
}

// NewBoardFixTopicBuilder func
func NewBoardFixTopicBuilder() *BoardFixTopicBuilder {
	return &BoardFixTopicBuilder{api.Params{}}
}

// GroupID ID of the community that owns the discussion board.
func (b *BoardFixTopicBuilder) GroupID(v int) {
	b.Params["group_id"] = v
}

// TopicID Topic ID.
func (b *BoardFixTopicBuilder) TopicID(v int) {
	b.Params["topic_id"] = v
}

// BoardGetCommentsBuilder builder
//
// Returns a list of comments on a topic on a community's discussion board.
//
// https://vk.com/dev/board.getComments
type BoardGetCommentsBuilder struct {
	api.Params
}

// NewBoardGetCommentsBuilder func
func NewBoardGetCommentsBuilder() *BoardGetCommentsBuilder {
	return &BoardGetCommentsBuilder{api.Params{}}
}

// GroupID ID of the community that owns the discussion board.
func (b *BoardGetCommentsBuilder) GroupID(v int) {
	b.Params["group_id"] = v
}

// TopicID Topic ID.
func (b *BoardGetCommentsBuilder) TopicID(v int) {
	b.Params["topic_id"] = v
}

// NeedLikes '1' — to return the 'likes' field, '0' — not to return the 'likes' field (default)
func (b *BoardGetCommentsBuilder) NeedLikes(v bool) {
	b.Params["need_likes"] = v
}

// StartCommentID parameter
func (b *BoardGetCommentsBuilder) StartCommentID(v int) {
	b.Params["start_comment_id"] = v
}

// Offset Offset needed to return a specific subset of comments.
func (b *BoardGetCommentsBuilder) Offset(v int) {
	b.Params["offset"] = v
}

// Count Number of comments to return.
func (b *BoardGetCommentsBuilder) Count(v int) {
	b.Params["count"] = v
}

// Extended '1' — to return information about users who posted comments, '0' — to return no additional fields (default)
func (b *BoardGetCommentsBuilder) Extended(v bool) {
	b.Params["extended"] = v
}

// Sort Sort order: 'asc' — by creation date in chronological order, 'desc' — by creation date in reverse chronological order,
func (b *BoardGetCommentsBuilder) Sort(v string) {
	b.Params["sort"] = v
}

// BoardGetTopicsBuilder builder
//
// Returns a list of topics on a community's discussion board.
//
// https://vk.com/dev/board.getTopics
type BoardGetTopicsBuilder struct {
	api.Params
}

// NewBoardGetTopicsBuilder func
func NewBoardGetTopicsBuilder() *BoardGetTopicsBuilder {
	return &BoardGetTopicsBuilder{api.Params{}}
}

// GroupID ID of the community that owns the discussion board.
func (b *BoardGetTopicsBuilder) GroupID(v int) {
	b.Params["group_id"] = v
}

// TopicIDs IDs of topics to be returned (100 maximum). By default, all topics are returned. If this parameter is set, the 'order', 'offset', and 'count' parameters are ignored.
func (b *BoardGetTopicsBuilder) TopicIDs(v []int) {
	b.Params["topic_ids"] = v
}

// Order Sort order: '1' — by date updated in reverse chronological order. '2' — by date created in reverse chronological order. '-1' — by date updated in chronological order. '-2' — by date created in chronological order. If no sort order is specified, topics are returned in the order specified by the group administrator. Pinned topics are returned first, regardless of the sorting.
func (b *BoardGetTopicsBuilder) Order(v int) {
	b.Params["order"] = v
}

// Offset Offset needed to return a specific subset of topics.
func (b *BoardGetTopicsBuilder) Offset(v int) {
	b.Params["offset"] = v
}

// Count Number of topics to return.
func (b *BoardGetTopicsBuilder) Count(v int) {
	b.Params["count"] = v
}

// Extended '1' — to return information about users who created topics or who posted there last, '0' — to return no additional fields (default)
func (b *BoardGetTopicsBuilder) Extended(v bool) {
	b.Params["extended"] = v
}

// Preview '1' — to return the first comment in each topic,, '2' — to return the last comment in each topic,, '0' — to return no comments. By default: '0'.
func (b *BoardGetTopicsBuilder) Preview(v int) {
	b.Params["preview"] = v
}

// PreviewLength Number of characters after which to truncate the previewed comment. To preview the full comment, specify '0'.
func (b *BoardGetTopicsBuilder) PreviewLength(v int) {
	b.Params["preview_length"] = v
}

// BoardOpenTopicBuilder builder
//
// Re-opens a previously closed topic on a community's discussion board.
//
// https://vk.com/dev/board.openTopic
type BoardOpenTopicBuilder struct {
	api.Params
}

// NewBoardOpenTopicBuilder func
func NewBoardOpenTopicBuilder() *BoardOpenTopicBuilder {
	return &BoardOpenTopicBuilder{api.Params{}}
}

// GroupID ID of the community that owns the discussion board.
func (b *BoardOpenTopicBuilder) GroupID(v int) {
	b.Params["group_id"] = v
}

// TopicID Topic ID.
func (b *BoardOpenTopicBuilder) TopicID(v int) {
	b.Params["topic_id"] = v
}

// BoardRestoreCommentBuilder builder
//
// Restores a comment deleted from a topic on a community's discussion board.
//
// https://vk.com/dev/board.restoreComment
type BoardRestoreCommentBuilder struct {
	api.Params
}

// NewBoardRestoreCommentBuilder func
func NewBoardRestoreCommentBuilder() *BoardRestoreCommentBuilder {
	return &BoardRestoreCommentBuilder{api.Params{}}
}

// GroupID ID of the community that owns the discussion board.
func (b *BoardRestoreCommentBuilder) GroupID(v int) {
	b.Params["group_id"] = v
}

// TopicID Topic ID.
func (b *BoardRestoreCommentBuilder) TopicID(v int) {
	b.Params["topic_id"] = v
}

// CommentID Comment ID.
func (b *BoardRestoreCommentBuilder) CommentID(v int) {
	b.Params["comment_id"] = v
}

// BoardUnfixTopicBuilder builder
//
// Unpins a pinned topic from the top of a community's discussion board.
//
// https://vk.com/dev/board.unfixTopic
type BoardUnfixTopicBuilder struct {
	api.Params
}

// NewBoardUnfixTopicBuilder func
func NewBoardUnfixTopicBuilder() *BoardUnfixTopicBuilder {
	return &BoardUnfixTopicBuilder{api.Params{}}
}

// GroupID ID of the community that owns the discussion board.
func (b *BoardUnfixTopicBuilder) GroupID(v int) {
	b.Params["group_id"] = v
}

// TopicID Topic ID.
func (b *BoardUnfixTopicBuilder) TopicID(v int) {
	b.Params["topic_id"] = v
}
