package params // import "github.com/SevereCloud/vksdk/api/params"

import (
	"github.com/SevereCloud/vksdk/api"
)

// BoardAddTopicBulder builder
//
// Creates a new topic on a community's discussion board.
//
// https://vk.com/dev/board.addTopic
type BoardAddTopicBulder struct {
	api.Params
}

// NewBoardAddTopicBulder func
func NewBoardAddTopicBulder() *BoardAddTopicBulder {
	return &BoardAddTopicBulder{api.Params{}}
}

// GroupID ID of the community that owns the discussion board.
func (b *BoardAddTopicBulder) GroupID(v int) {
	b.Params["group_id"] = v
}

// Title Topic title.
func (b *BoardAddTopicBulder) Title(v string) {
	b.Params["title"] = v
}

// Text Text of the topic.
func (b *BoardAddTopicBulder) Text(v string) {
	b.Params["text"] = v
}

// FromGroup For a community: '1' — to post the topic as by the community, '0' — to post the topic as by the user (default)
func (b *BoardAddTopicBulder) FromGroup(v bool) {
	b.Params["from_group"] = v
}

// Attachments List of media objects attached to the topic, in the following format: "<owner_id>_<media_id>,<owner_id>_<media_id>", '' — Type of media object: 'photo' — photo, 'video' — video, 'audio' — audio, 'doc' — document, '<owner_id>' — ID of the media owner. '<media_id>' — Media ID. Example: "photo100172_166443618,photo66748_265827614", , "NOTE: If you try to attach more than one reference, an error will be thrown.",
func (b *BoardAddTopicBulder) Attachments(v []string) {
	b.Params["attachments"] = v
}

// BoardCloseTopicBulder builder
//
// Closes a topic on a community's discussion board so that comments cannot be posted.
//
// https://vk.com/dev/board.closeTopic
type BoardCloseTopicBulder struct {
	api.Params
}

// NewBoardCloseTopicBulder func
func NewBoardCloseTopicBulder() *BoardCloseTopicBulder {
	return &BoardCloseTopicBulder{api.Params{}}
}

// GroupID ID of the community that owns the discussion board.
func (b *BoardCloseTopicBulder) GroupID(v int) {
	b.Params["group_id"] = v
}

// TopicID Topic ID.
func (b *BoardCloseTopicBulder) TopicID(v int) {
	b.Params["topic_id"] = v
}

// BoardCreateCommentBulder builder
//
// Adds a comment on a topic on a community's discussion board.
//
// https://vk.com/dev/board.createComment
type BoardCreateCommentBulder struct {
	api.Params
}

// NewBoardCreateCommentBulder func
func NewBoardCreateCommentBulder() *BoardCreateCommentBulder {
	return &BoardCreateCommentBulder{api.Params{}}
}

// GroupID ID of the community that owns the discussion board.
func (b *BoardCreateCommentBulder) GroupID(v int) {
	b.Params["group_id"] = v
}

// TopicID ID of the topic to be commented on.
func (b *BoardCreateCommentBulder) TopicID(v int) {
	b.Params["topic_id"] = v
}

// Message (Required if 'attachments' is not set.) Text of the comment.
func (b *BoardCreateCommentBulder) Message(v string) {
	b.Params["message"] = v
}

// Attachments (Required if 'text' is not set.) List of media objects attached to the comment, in the following format: "<owner_id>_<media_id>,<owner_id>_<media_id>", '' — Type of media object: 'photo' — photo, 'video' — video, 'audio' — audio, 'doc' — document, '<owner_id>' — ID of the media owner. '<media_id>' — Media ID.
func (b *BoardCreateCommentBulder) Attachments(v []string) {
	b.Params["attachments"] = v
}

// FromGroup '1' — to post the comment as by the community, '0' — to post the comment as by the user (default)
func (b *BoardCreateCommentBulder) FromGroup(v bool) {
	b.Params["from_group"] = v
}

// StickerID Sticker ID.
func (b *BoardCreateCommentBulder) StickerID(v int) {
	b.Params["sticker_id"] = v
}

// GUID Unique identifier to avoid repeated comments.
func (b *BoardCreateCommentBulder) GUID(v string) {
	b.Params["guid"] = v
}

// BoardDeleteCommentBulder builder
//
// Deletes a comment on a topic on a community's discussion board.
//
// https://vk.com/dev/board.deleteComment
type BoardDeleteCommentBulder struct {
	api.Params
}

// NewBoardDeleteCommentBulder func
func NewBoardDeleteCommentBulder() *BoardDeleteCommentBulder {
	return &BoardDeleteCommentBulder{api.Params{}}
}

// GroupID ID of the community that owns the discussion board.
func (b *BoardDeleteCommentBulder) GroupID(v int) {
	b.Params["group_id"] = v
}

// TopicID Topic ID.
func (b *BoardDeleteCommentBulder) TopicID(v int) {
	b.Params["topic_id"] = v
}

// CommentID Comment ID.
func (b *BoardDeleteCommentBulder) CommentID(v int) {
	b.Params["comment_id"] = v
}

// BoardDeleteTopicBulder builder
//
// Deletes a topic from a community's discussion board.
//
// https://vk.com/dev/board.deleteTopic
type BoardDeleteTopicBulder struct {
	api.Params
}

// NewBoardDeleteTopicBulder func
func NewBoardDeleteTopicBulder() *BoardDeleteTopicBulder {
	return &BoardDeleteTopicBulder{api.Params{}}
}

// GroupID ID of the community that owns the discussion board.
func (b *BoardDeleteTopicBulder) GroupID(v int) {
	b.Params["group_id"] = v
}

// TopicID Topic ID.
func (b *BoardDeleteTopicBulder) TopicID(v int) {
	b.Params["topic_id"] = v
}

// BoardEditCommentBulder builder
//
// Edits a comment on a topic on a community's discussion board.
//
// https://vk.com/dev/board.editComment
type BoardEditCommentBulder struct {
	api.Params
}

// NewBoardEditCommentBulder func
func NewBoardEditCommentBulder() *BoardEditCommentBulder {
	return &BoardEditCommentBulder{api.Params{}}
}

// GroupID ID of the community that owns the discussion board.
func (b *BoardEditCommentBulder) GroupID(v int) {
	b.Params["group_id"] = v
}

// TopicID Topic ID.
func (b *BoardEditCommentBulder) TopicID(v int) {
	b.Params["topic_id"] = v
}

// CommentID ID of the comment on the topic.
func (b *BoardEditCommentBulder) CommentID(v int) {
	b.Params["comment_id"] = v
}

// Message (Required if 'attachments' is not set). New comment text.
func (b *BoardEditCommentBulder) Message(v string) {
	b.Params["message"] = v
}

// Attachments (Required if 'message' is not set.) List of media objects attached to the comment, in the following format: "<owner_id>_<media_id>,<owner_id>_<media_id>", '' — Type of media object: 'photo' — photo, 'video' — video, 'audio' — audio, 'doc' — document, '<owner_id>' — ID of the media owner. '<media_id>' — Media ID. Example: "photo100172_166443618,photo66748_265827614"
func (b *BoardEditCommentBulder) Attachments(v []string) {
	b.Params["attachments"] = v
}

// BoardEditTopicBulder builder
//
// Edits the title of a topic on a community's discussion board.
//
// https://vk.com/dev/board.editTopic
type BoardEditTopicBulder struct {
	api.Params
}

// NewBoardEditTopicBulder func
func NewBoardEditTopicBulder() *BoardEditTopicBulder {
	return &BoardEditTopicBulder{api.Params{}}
}

// GroupID ID of the community that owns the discussion board.
func (b *BoardEditTopicBulder) GroupID(v int) {
	b.Params["group_id"] = v
}

// TopicID Topic ID.
func (b *BoardEditTopicBulder) TopicID(v int) {
	b.Params["topic_id"] = v
}

// Title New title of the topic.
func (b *BoardEditTopicBulder) Title(v string) {
	b.Params["title"] = v
}

// BoardFixTopicBulder builder
//
// Pins a topic (fixes its place) to the top of a community's discussion board.
//
// https://vk.com/dev/board.fixTopic
type BoardFixTopicBulder struct {
	api.Params
}

// NewBoardFixTopicBulder func
func NewBoardFixTopicBulder() *BoardFixTopicBulder {
	return &BoardFixTopicBulder{api.Params{}}
}

// GroupID ID of the community that owns the discussion board.
func (b *BoardFixTopicBulder) GroupID(v int) {
	b.Params["group_id"] = v
}

// TopicID Topic ID.
func (b *BoardFixTopicBulder) TopicID(v int) {
	b.Params["topic_id"] = v
}

// BoardGetCommentsBulder builder
//
// Returns a list of comments on a topic on a community's discussion board.
//
// https://vk.com/dev/board.getComments
type BoardGetCommentsBulder struct {
	api.Params
}

// NewBoardGetCommentsBulder func
func NewBoardGetCommentsBulder() *BoardGetCommentsBulder {
	return &BoardGetCommentsBulder{api.Params{}}
}

// GroupID ID of the community that owns the discussion board.
func (b *BoardGetCommentsBulder) GroupID(v int) {
	b.Params["group_id"] = v
}

// TopicID Topic ID.
func (b *BoardGetCommentsBulder) TopicID(v int) {
	b.Params["topic_id"] = v
}

// NeedLikes '1' — to return the 'likes' field, '0' — not to return the 'likes' field (default)
func (b *BoardGetCommentsBulder) NeedLikes(v bool) {
	b.Params["need_likes"] = v
}

// StartCommentID parameter
func (b *BoardGetCommentsBulder) StartCommentID(v int) {
	b.Params["start_comment_id"] = v
}

// Offset Offset needed to return a specific subset of comments.
func (b *BoardGetCommentsBulder) Offset(v int) {
	b.Params["offset"] = v
}

// Count Number of comments to return.
func (b *BoardGetCommentsBulder) Count(v int) {
	b.Params["count"] = v
}

// Extended '1' — to return information about users who posted comments, '0' — to return no additional fields (default)
func (b *BoardGetCommentsBulder) Extended(v bool) {
	b.Params["extended"] = v
}

// Sort Sort order: 'asc' — by creation date in chronological order, 'desc' — by creation date in reverse chronological order,
func (b *BoardGetCommentsBulder) Sort(v string) {
	b.Params["sort"] = v
}

// BoardGetTopicsBulder builder
//
// Returns a list of topics on a community's discussion board.
//
// https://vk.com/dev/board.getTopics
type BoardGetTopicsBulder struct {
	api.Params
}

// NewBoardGetTopicsBulder func
func NewBoardGetTopicsBulder() *BoardGetTopicsBulder {
	return &BoardGetTopicsBulder{api.Params{}}
}

// GroupID ID of the community that owns the discussion board.
func (b *BoardGetTopicsBulder) GroupID(v int) {
	b.Params["group_id"] = v
}

// TopicIDs IDs of topics to be returned (100 maximum). By default, all topics are returned. If this parameter is set, the 'order', 'offset', and 'count' parameters are ignored.
func (b *BoardGetTopicsBulder) TopicIDs(v []int) {
	b.Params["topic_ids"] = v
}

// Order Sort order: '1' — by date updated in reverse chronological order. '2' — by date created in reverse chronological order. '-1' — by date updated in chronological order. '-2' — by date created in chronological order. If no sort order is specified, topics are returned in the order specified by the group administrator. Pinned topics are returned first, regardless of the sorting.
func (b *BoardGetTopicsBulder) Order(v int) {
	b.Params["order"] = v
}

// Offset Offset needed to return a specific subset of topics.
func (b *BoardGetTopicsBulder) Offset(v int) {
	b.Params["offset"] = v
}

// Count Number of topics to return.
func (b *BoardGetTopicsBulder) Count(v int) {
	b.Params["count"] = v
}

// Extended '1' — to return information about users who created topics or who posted there last, '0' — to return no additional fields (default)
func (b *BoardGetTopicsBulder) Extended(v bool) {
	b.Params["extended"] = v
}

// Preview '1' — to return the first comment in each topic,, '2' — to return the last comment in each topic,, '0' — to return no comments. By default: '0'.
func (b *BoardGetTopicsBulder) Preview(v int) {
	b.Params["preview"] = v
}

// PreviewLength Number of characters after which to truncate the previewed comment. To preview the full comment, specify '0'.
func (b *BoardGetTopicsBulder) PreviewLength(v int) {
	b.Params["preview_length"] = v
}

// BoardOpenTopicBulder builder
//
// Re-opens a previously closed topic on a community's discussion board.
//
// https://vk.com/dev/board.openTopic
type BoardOpenTopicBulder struct {
	api.Params
}

// NewBoardOpenTopicBulder func
func NewBoardOpenTopicBulder() *BoardOpenTopicBulder {
	return &BoardOpenTopicBulder{api.Params{}}
}

// GroupID ID of the community that owns the discussion board.
func (b *BoardOpenTopicBulder) GroupID(v int) {
	b.Params["group_id"] = v
}

// TopicID Topic ID.
func (b *BoardOpenTopicBulder) TopicID(v int) {
	b.Params["topic_id"] = v
}

// BoardRestoreCommentBulder builder
//
// Restores a comment deleted from a topic on a community's discussion board.
//
// https://vk.com/dev/board.restoreComment
type BoardRestoreCommentBulder struct {
	api.Params
}

// NewBoardRestoreCommentBulder func
func NewBoardRestoreCommentBulder() *BoardRestoreCommentBulder {
	return &BoardRestoreCommentBulder{api.Params{}}
}

// GroupID ID of the community that owns the discussion board.
func (b *BoardRestoreCommentBulder) GroupID(v int) {
	b.Params["group_id"] = v
}

// TopicID Topic ID.
func (b *BoardRestoreCommentBulder) TopicID(v int) {
	b.Params["topic_id"] = v
}

// CommentID Comment ID.
func (b *BoardRestoreCommentBulder) CommentID(v int) {
	b.Params["comment_id"] = v
}

// BoardUnfixTopicBulder builder
//
// Unpins a pinned topic from the top of a community's discussion board.
//
// https://vk.com/dev/board.unfixTopic
type BoardUnfixTopicBulder struct {
	api.Params
}

// NewBoardUnfixTopicBulder func
func NewBoardUnfixTopicBulder() *BoardUnfixTopicBulder {
	return &BoardUnfixTopicBulder{api.Params{}}
}

// GroupID ID of the community that owns the discussion board.
func (b *BoardUnfixTopicBulder) GroupID(v int) {
	b.Params["group_id"] = v
}

// TopicID Topic ID.
func (b *BoardUnfixTopicBulder) TopicID(v int) {
	b.Params["topic_id"] = v
}
