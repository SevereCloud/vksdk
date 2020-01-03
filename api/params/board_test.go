package params_test

import (
	"testing"

	"github.com/SevereCloud/vksdk/api/params"
	"github.com/stretchr/testify/assert"
)

func TestBoardAddTopicBulder(t *testing.T) {
	b := params.NewBoardAddTopicBulder()

	b.GroupID(1)
	b.Title("text")
	b.Text("text")
	b.FromGroup(true)
	b.Attachments([]string{"text"})

	assert.Equal(t, b.Params["group_id"], 1)
	assert.Equal(t, b.Params["title"], "text")
	assert.Equal(t, b.Params["text"], "text")
	assert.Equal(t, b.Params["from_group"], true)
	assert.Equal(t, b.Params["attachments"], []string{"text"})
}

func TestBoardCloseTopicBulder(t *testing.T) {
	b := params.NewBoardCloseTopicBulder()

	b.GroupID(1)
	b.TopicID(1)

	assert.Equal(t, b.Params["group_id"], 1)
	assert.Equal(t, b.Params["topic_id"], 1)
}

func TestBoardCreateCommentBulder(t *testing.T) {
	b := params.NewBoardCreateCommentBulder()

	b.GroupID(1)
	b.TopicID(1)
	b.Message("text")
	b.Attachments([]string{"text"})
	b.FromGroup(true)
	b.StickerID(1)
	b.GUID("text")

	assert.Equal(t, b.Params["group_id"], 1)
	assert.Equal(t, b.Params["topic_id"], 1)
	assert.Equal(t, b.Params["message"], "text")
	assert.Equal(t, b.Params["attachments"], []string{"text"})
	assert.Equal(t, b.Params["from_group"], true)
	assert.Equal(t, b.Params["sticker_id"], 1)
	assert.Equal(t, b.Params["guid"], "text")
}

func TestBoardDeleteCommentBulder(t *testing.T) {
	b := params.NewBoardDeleteCommentBulder()

	b.GroupID(1)
	b.TopicID(1)
	b.CommentID(1)

	assert.Equal(t, b.Params["group_id"], 1)
	assert.Equal(t, b.Params["topic_id"], 1)
	assert.Equal(t, b.Params["comment_id"], 1)
}

func TestBoardDeleteTopicBulder(t *testing.T) {
	b := params.NewBoardDeleteTopicBulder()

	b.GroupID(1)
	b.TopicID(1)

	assert.Equal(t, b.Params["group_id"], 1)
	assert.Equal(t, b.Params["topic_id"], 1)
}

func TestBoardEditCommentBulder(t *testing.T) {
	b := params.NewBoardEditCommentBulder()

	b.GroupID(1)
	b.TopicID(1)
	b.CommentID(1)
	b.Message("text")
	b.Attachments([]string{"text"})

	assert.Equal(t, b.Params["group_id"], 1)
	assert.Equal(t, b.Params["topic_id"], 1)
	assert.Equal(t, b.Params["comment_id"], 1)
	assert.Equal(t, b.Params["message"], "text")
	assert.Equal(t, b.Params["attachments"], []string{"text"})
}

func TestBoardEditTopicBulder(t *testing.T) {
	b := params.NewBoardEditTopicBulder()

	b.GroupID(1)
	b.TopicID(1)
	b.Title("text")

	assert.Equal(t, b.Params["group_id"], 1)
	assert.Equal(t, b.Params["topic_id"], 1)
	assert.Equal(t, b.Params["title"], "text")
}

func TestBoardFixTopicBulder(t *testing.T) {
	b := params.NewBoardFixTopicBulder()

	b.GroupID(1)
	b.TopicID(1)

	assert.Equal(t, b.Params["group_id"], 1)
	assert.Equal(t, b.Params["topic_id"], 1)
}

func TestBoardGetCommentsBulder(t *testing.T) {
	b := params.NewBoardGetCommentsBulder()

	b.GroupID(1)
	b.TopicID(1)
	b.NeedLikes(true)
	b.StartCommentID(1)
	b.Offset(1)
	b.Count(1)
	b.Extended(true)
	b.Sort("text")

	assert.Equal(t, b.Params["group_id"], 1)
	assert.Equal(t, b.Params["topic_id"], 1)
	assert.Equal(t, b.Params["need_likes"], true)
	assert.Equal(t, b.Params["start_comment_id"], 1)
	assert.Equal(t, b.Params["offset"], 1)
	assert.Equal(t, b.Params["count"], 1)
	assert.Equal(t, b.Params["extended"], true)
	assert.Equal(t, b.Params["sort"], "text")
}

func TestBoardGetTopicsBulder(t *testing.T) {
	b := params.NewBoardGetTopicsBulder()

	b.GroupID(1)
	b.TopicIDs([]int{1})
	b.Order(1)
	b.Offset(1)
	b.Count(1)
	b.Extended(true)
	b.Preview(1)
	b.PreviewLength(1)

	assert.Equal(t, b.Params["group_id"], 1)
	assert.Equal(t, b.Params["topic_ids"], []int{1})
	assert.Equal(t, b.Params["order"], 1)
	assert.Equal(t, b.Params["offset"], 1)
	assert.Equal(t, b.Params["count"], 1)
	assert.Equal(t, b.Params["extended"], true)
	assert.Equal(t, b.Params["preview"], 1)
	assert.Equal(t, b.Params["preview_length"], 1)
}

func TestBoardOpenTopicBulder(t *testing.T) {
	b := params.NewBoardOpenTopicBulder()

	b.GroupID(1)
	b.TopicID(1)

	assert.Equal(t, b.Params["group_id"], 1)
	assert.Equal(t, b.Params["topic_id"], 1)
}

func TestBoardRestoreCommentBulder(t *testing.T) {
	b := params.NewBoardRestoreCommentBulder()

	b.GroupID(1)
	b.TopicID(1)
	b.CommentID(1)

	assert.Equal(t, b.Params["group_id"], 1)
	assert.Equal(t, b.Params["topic_id"], 1)
	assert.Equal(t, b.Params["comment_id"], 1)
}

func TestBoardUnfixTopicBulder(t *testing.T) {
	b := params.NewBoardUnfixTopicBulder()

	b.GroupID(1)
	b.TopicID(1)

	assert.Equal(t, b.Params["group_id"], 1)
	assert.Equal(t, b.Params["topic_id"], 1)
}
