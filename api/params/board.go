package params // import "github.com/SevereCloud/vksdk/v2/api/params"

import (
	"github.com/SevereCloud/vksdk/v2/api"
)

// BoardAddTopicBuilder builder.
//
// Creates a new topic on a community's discussion board.
//
// https://vk.com/dev/board.addTopic
type BoardAddTopicBuilder struct {
	api.Params
}

// NewBoardAddTopicBuilder func.
func NewBoardAddTopicBuilder() *BoardAddTopicBuilder {
	return &BoardAddTopicBuilder{api.Params{}}
}

// GroupID ID of the community that owns the discussion board.
func (b *BoardAddTopicBuilder) GroupID(v int) *BoardAddTopicBuilder {
	b.Params["group_id"] = v
	return b
}

// Title Topic title.
func (b *BoardAddTopicBuilder) Title(v string) *BoardAddTopicBuilder {
	b.Params["title"] = v
	return b
}

// Text Text of the topic.
func (b *BoardAddTopicBuilder) Text(v string) *BoardAddTopicBuilder {
	b.Params["text"] = v
	return b
}

// FromGroup for a community:
//
// * 1 — to post the topic as by the community,
//
// * 0 — to post the topic as by the user (default).
func (b *BoardAddTopicBuilder) FromGroup(v bool) *BoardAddTopicBuilder {
	b.Params["from_group"] = v
	return b
}

// Attachments list of media objects attached to the topic, in the following format:
// "<owner_id>_<media_id>,<owner_id>_<media_id>", ” — Type of media object:
// 'photo' — photo, 'video' — video, 'audio' — audio, 'doc' — document, '<owner_id>' — ID of the media owner.
// '<media_id>' — Media ID. Example: "photo100172_166443618,photo66748_265827614".
//
// NOTE: If you try to attach more than one reference, an error will be thrown.
func (b *BoardAddTopicBuilder) Attachments(v interface{}) *BoardAddTopicBuilder {
	b.Params["attachments"] = v
	return b
}

// BoardCloseTopicBuilder builder.
//
// Closes a topic on a community's discussion board so that comments cannot be posted.
//
// https://vk.com/dev/board.closeTopic
type BoardCloseTopicBuilder struct {
	api.Params
}

// NewBoardCloseTopicBuilder func.
func NewBoardCloseTopicBuilder() *BoardCloseTopicBuilder {
	return &BoardCloseTopicBuilder{api.Params{}}
}

// GroupID ID of the community that owns the discussion board.
func (b *BoardCloseTopicBuilder) GroupID(v int) *BoardCloseTopicBuilder {
	b.Params["group_id"] = v
	return b
}

// TopicID parameter.
func (b *BoardCloseTopicBuilder) TopicID(v int) *BoardCloseTopicBuilder {
	b.Params["topic_id"] = v
	return b
}

// BoardCreateCommentBuilder builder.
//
// Adds a comment on a topic on a community's discussion board.
//
// https://vk.com/dev/board.createComment
type BoardCreateCommentBuilder struct {
	api.Params
}

// NewBoardCreateCommentBuilder func.
func NewBoardCreateCommentBuilder() *BoardCreateCommentBuilder {
	return &BoardCreateCommentBuilder{api.Params{}}
}

// GroupID ID of the community that owns the discussion board.
func (b *BoardCreateCommentBuilder) GroupID(v int) *BoardCreateCommentBuilder {
	b.Params["group_id"] = v
	return b
}

// TopicID ID of the topic to be commented on.
func (b *BoardCreateCommentBuilder) TopicID(v int) *BoardCreateCommentBuilder {
	b.Params["topic_id"] = v
	return b
}

// Message (Required if 'attachments' is not set.) Text of the comment.
func (b *BoardCreateCommentBuilder) Message(v string) *BoardCreateCommentBuilder {
	b.Params["message"] = v
	return b
}

// Attachments (Required if 'text' is not set.)
// List of media objects attached to the comment, in the following format:
// "<owner_id>_<media_id>,<owner_id>_<media_id>", ” — Type of media object:
// 'photo' — photo, 'video' — video, 'audio' — audio, 'doc' — document,
// '<owner_id>' — ID of the media owner. '<media_id>' — Media ID.
func (b *BoardCreateCommentBuilder) Attachments(v interface{}) *BoardCreateCommentBuilder {
	b.Params["attachments"] = v
	return b
}

// FromGroup '1' — to post the comment as by the community, '0' — to post the comment as by the user (default).
func (b *BoardCreateCommentBuilder) FromGroup(v bool) *BoardCreateCommentBuilder {
	b.Params["from_group"] = v
	return b
}

// StickerID parameter.
func (b *BoardCreateCommentBuilder) StickerID(v int) *BoardCreateCommentBuilder {
	b.Params["sticker_id"] = v
	return b
}

// GUID unique identifier to avoid repeated comments.
func (b *BoardCreateCommentBuilder) GUID(v string) *BoardCreateCommentBuilder {
	b.Params["guid"] = v
	return b
}

// BoardDeleteCommentBuilder builder.
//
// Deletes a comment on a topic on a community's discussion board.
//
// https://vk.com/dev/board.deleteComment
type BoardDeleteCommentBuilder struct {
	api.Params
}

// NewBoardDeleteCommentBuilder func.
func NewBoardDeleteCommentBuilder() *BoardDeleteCommentBuilder {
	return &BoardDeleteCommentBuilder{api.Params{}}
}

// GroupID ID of the community that owns the discussion board.
func (b *BoardDeleteCommentBuilder) GroupID(v int) *BoardDeleteCommentBuilder {
	b.Params["group_id"] = v
	return b
}

// TopicID parameter.
func (b *BoardDeleteCommentBuilder) TopicID(v int) *BoardDeleteCommentBuilder {
	b.Params["topic_id"] = v
	return b
}

// CommentID parameter.
func (b *BoardDeleteCommentBuilder) CommentID(v int) *BoardDeleteCommentBuilder {
	b.Params["comment_id"] = v
	return b
}

// BoardDeleteTopicBuilder builder.
//
// Deletes a topic from a community's discussion board.
//
// https://vk.com/dev/board.deleteTopic
type BoardDeleteTopicBuilder struct {
	api.Params
}

// NewBoardDeleteTopicBuilder func.
func NewBoardDeleteTopicBuilder() *BoardDeleteTopicBuilder {
	return &BoardDeleteTopicBuilder{api.Params{}}
}

// GroupID ID of the community that owns the discussion board.
func (b *BoardDeleteTopicBuilder) GroupID(v int) *BoardDeleteTopicBuilder {
	b.Params["group_id"] = v
	return b
}

// TopicID parameter.
func (b *BoardDeleteTopicBuilder) TopicID(v int) *BoardDeleteTopicBuilder {
	b.Params["topic_id"] = v
	return b
}

// BoardEditCommentBuilder builder.
//
// Edits a comment on a topic on a community's discussion board.
//
// https://vk.com/dev/board.editComment
type BoardEditCommentBuilder struct {
	api.Params
}

// NewBoardEditCommentBuilder func.
func NewBoardEditCommentBuilder() *BoardEditCommentBuilder {
	return &BoardEditCommentBuilder{api.Params{}}
}

// GroupID ID of the community that owns the discussion board.
func (b *BoardEditCommentBuilder) GroupID(v int) *BoardEditCommentBuilder {
	b.Params["group_id"] = v
	return b
}

// TopicID parameter.
func (b *BoardEditCommentBuilder) TopicID(v int) *BoardEditCommentBuilder {
	b.Params["topic_id"] = v
	return b
}

// CommentID ID of the comment on the topic.
func (b *BoardEditCommentBuilder) CommentID(v int) *BoardEditCommentBuilder {
	b.Params["comment_id"] = v
	return b
}

// Message (Required if 'attachments' is not set). New comment text.
func (b *BoardEditCommentBuilder) Message(v string) *BoardEditCommentBuilder {
	b.Params["message"] = v
	return b
}

// Attachments (Required if 'message' is not set.) List of media objects attached to the comment,
// in the following format: "<owner_id>_<media_id>,<owner_id>_<media_id>", ” —
// Type of media object: 'photo' — photo, 'video' — video, 'audio' — audio, 'doc' — document,
// '<owner_id>' — ID of the media owner. '<media_id>' — Media ID.
// Example: "photo100172_166443618,photo66748_265827614".
func (b *BoardEditCommentBuilder) Attachments(v interface{}) *BoardEditCommentBuilder {
	b.Params["attachments"] = v
	return b
}

// BoardEditTopicBuilder builder.
//
// Edits the title of a topic on a community's discussion board.
//
// https://vk.com/dev/board.editTopic
type BoardEditTopicBuilder struct {
	api.Params
}

// NewBoardEditTopicBuilder func.
func NewBoardEditTopicBuilder() *BoardEditTopicBuilder {
	return &BoardEditTopicBuilder{api.Params{}}
}

// GroupID ID of the community that owns the discussion board.
func (b *BoardEditTopicBuilder) GroupID(v int) *BoardEditTopicBuilder {
	b.Params["group_id"] = v
	return b
}

// TopicID parameter.
func (b *BoardEditTopicBuilder) TopicID(v int) *BoardEditTopicBuilder {
	b.Params["topic_id"] = v
	return b
}

// Title new title of the topic.
func (b *BoardEditTopicBuilder) Title(v string) *BoardEditTopicBuilder {
	b.Params["title"] = v
	return b
}

// BoardFixTopicBuilder builder.
//
// Pins a topic (fixes its place) to the top of a community's discussion board.
//
// https://vk.com/dev/board.fixTopic
type BoardFixTopicBuilder struct {
	api.Params
}

// NewBoardFixTopicBuilder func.
func NewBoardFixTopicBuilder() *BoardFixTopicBuilder {
	return &BoardFixTopicBuilder{api.Params{}}
}

// GroupID ID of the community that owns the discussion board.
func (b *BoardFixTopicBuilder) GroupID(v int) *BoardFixTopicBuilder {
	b.Params["group_id"] = v
	return b
}

// TopicID parameter.
func (b *BoardFixTopicBuilder) TopicID(v int) *BoardFixTopicBuilder {
	b.Params["topic_id"] = v
	return b
}

// BoardGetCommentsBuilder builder.
//
// Returns a list of comments on a topic on a community's discussion board.
//
// https://vk.com/dev/board.getComments
type BoardGetCommentsBuilder struct {
	api.Params
}

// NewBoardGetCommentsBuilder func.
func NewBoardGetCommentsBuilder() *BoardGetCommentsBuilder {
	return &BoardGetCommentsBuilder{api.Params{}}
}

// GroupID ID of the community that owns the discussion board.
func (b *BoardGetCommentsBuilder) GroupID(v int) *BoardGetCommentsBuilder {
	b.Params["group_id"] = v
	return b
}

// TopicID parameter.
func (b *BoardGetCommentsBuilder) TopicID(v int) *BoardGetCommentsBuilder {
	b.Params["topic_id"] = v
	return b
}

// NeedLikes '1' — to return the 'likes' field, '0' — not to return the 'likes' field (default).
func (b *BoardGetCommentsBuilder) NeedLikes(v bool) *BoardGetCommentsBuilder {
	b.Params["need_likes"] = v
	return b
}

// StartCommentID parameter.
func (b *BoardGetCommentsBuilder) StartCommentID(v int) *BoardGetCommentsBuilder {
	b.Params["start_comment_id"] = v
	return b
}

// Offset needed to return a specific subset of comments.
func (b *BoardGetCommentsBuilder) Offset(v int) *BoardGetCommentsBuilder {
	b.Params["offset"] = v
	return b
}

// Count number of comments to return.
func (b *BoardGetCommentsBuilder) Count(v int) *BoardGetCommentsBuilder {
	b.Params["count"] = v
	return b
}

// Extended '1' — to return information about users who posted comments, '0' — to return no additional fields (default).
func (b *BoardGetCommentsBuilder) Extended(v bool) *BoardGetCommentsBuilder {
	b.Params["extended"] = v
	return b
}

// Sort order:
//
// * asc — by creation date in chronological order, 'desc' — by creation date in reverse chronological order.
func (b *BoardGetCommentsBuilder) Sort(v string) *BoardGetCommentsBuilder {
	b.Params["sort"] = v
	return b
}

// BoardGetTopicsBuilder builder.
//
// Returns a list of topics on a community's discussion board.
//
// https://vk.com/dev/board.getTopics
type BoardGetTopicsBuilder struct {
	api.Params
}

// NewBoardGetTopicsBuilder func.
func NewBoardGetTopicsBuilder() *BoardGetTopicsBuilder {
	return &BoardGetTopicsBuilder{api.Params{}}
}

// GroupID ID of the community that owns the discussion board.
func (b *BoardGetTopicsBuilder) GroupID(v int) *BoardGetTopicsBuilder {
	b.Params["group_id"] = v
	return b
}

// TopicIDs IDs of topics to be returned (100 maximum). By default, all topics are returned.
// If this parameter is set, the 'order', 'offset', and 'count' parameters are ignored.
func (b *BoardGetTopicsBuilder) TopicIDs(v []int) *BoardGetTopicsBuilder {
	b.Params["topic_ids"] = v
	return b
}

// Order sort order:
//
// * 1 — by date updated in reverse chronological order.
//
// * 2 — by date created in reverse chronological order.
//
// * -1 — by date updated in chronological order.
//
// * -2 — by date created in chronological order.
//
// If no sort order is specified, topics are returned in the order specified by the group administrator.
// Pinned topics are returned first, regardless of the sorting.
func (b *BoardGetTopicsBuilder) Order(v int) *BoardGetTopicsBuilder {
	b.Params["order"] = v
	return b
}

// Offset needed to return a specific subset of topics.
func (b *BoardGetTopicsBuilder) Offset(v int) *BoardGetTopicsBuilder {
	b.Params["offset"] = v
	return b
}

// Count number of topics to return.
func (b *BoardGetTopicsBuilder) Count(v int) *BoardGetTopicsBuilder {
	b.Params["count"] = v
	return b
}

// Extended parameter.
//
// * 1 — to return information about users who created topics or who posted there last,
//
// * 0 — to return no additional fields (default).
func (b *BoardGetTopicsBuilder) Extended(v bool) *BoardGetTopicsBuilder {
	b.Params["extended"] = v
	return b
}

// Preview parameter.
//
// * 1 — to return the first comment in each topic;
//
// * 2 — to return the last comment in each topic;
//
// * 0 — to return no comments.
//
// By default: 0.
func (b *BoardGetTopicsBuilder) Preview(v int) *BoardGetTopicsBuilder {
	b.Params["preview"] = v
	return b
}

// PreviewLength number of characters after which to truncate the previewed comment.
// To preview the full comment, specify '0'.
func (b *BoardGetTopicsBuilder) PreviewLength(v int) *BoardGetTopicsBuilder {
	b.Params["preview_length"] = v
	return b
}

// BoardOpenTopicBuilder builder.
//
// Re-opens a previously closed topic on a community's discussion board.
//
// https://vk.com/dev/board.openTopic
type BoardOpenTopicBuilder struct {
	api.Params
}

// NewBoardOpenTopicBuilder func.
func NewBoardOpenTopicBuilder() *BoardOpenTopicBuilder {
	return &BoardOpenTopicBuilder{api.Params{}}
}

// GroupID ID of the community that owns the discussion board.
func (b *BoardOpenTopicBuilder) GroupID(v int) *BoardOpenTopicBuilder {
	b.Params["group_id"] = v
	return b
}

// TopicID parameter.
func (b *BoardOpenTopicBuilder) TopicID(v int) *BoardOpenTopicBuilder {
	b.Params["topic_id"] = v
	return b
}

// BoardRestoreCommentBuilder builder.
//
// Restores a comment deleted from a topic on a community's discussion board.
//
// https://vk.com/dev/board.restoreComment
type BoardRestoreCommentBuilder struct {
	api.Params
}

// NewBoardRestoreCommentBuilder func.
func NewBoardRestoreCommentBuilder() *BoardRestoreCommentBuilder {
	return &BoardRestoreCommentBuilder{api.Params{}}
}

// GroupID ID of the community that owns the discussion board.
func (b *BoardRestoreCommentBuilder) GroupID(v int) *BoardRestoreCommentBuilder {
	b.Params["group_id"] = v
	return b
}

// TopicID parameter.
func (b *BoardRestoreCommentBuilder) TopicID(v int) *BoardRestoreCommentBuilder {
	b.Params["topic_id"] = v
	return b
}

// CommentID parameter.
func (b *BoardRestoreCommentBuilder) CommentID(v int) *BoardRestoreCommentBuilder {
	b.Params["comment_id"] = v
	return b
}

// BoardUnfixTopicBuilder builder.
//
// Unpins a pinned topic from the top of a community's discussion board.
//
// https://vk.com/dev/board.unfixTopic
type BoardUnfixTopicBuilder struct {
	api.Params
}

// NewBoardUnfixTopicBuilder func.
func NewBoardUnfixTopicBuilder() *BoardUnfixTopicBuilder {
	return &BoardUnfixTopicBuilder{api.Params{}}
}

// GroupID ID of the community that owns the discussion board.
func (b *BoardUnfixTopicBuilder) GroupID(v int) *BoardUnfixTopicBuilder {
	b.Params["group_id"] = v
	return b
}

// TopicID parameter.
func (b *BoardUnfixTopicBuilder) TopicID(v int) *BoardUnfixTopicBuilder {
	b.Params["topic_id"] = v
	return b
}
