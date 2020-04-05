package params_test

import (
	"testing"

	"github.com/SevereCloud/vksdk/api/params"
	"github.com/stretchr/testify/assert"
)

func TestBoardAddTopicBuilder(t *testing.T) {
	t.Parallel()

	b := params.NewBoardAddTopicBuilder()

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

func TestBoardCloseTopicBuilder(t *testing.T) {
	t.Parallel()

	b := params.NewBoardCloseTopicBuilder()

	b.GroupID(1)
	b.TopicID(1)

	assert.Equal(t, b.Params["group_id"], 1)
	assert.Equal(t, b.Params["topic_id"], 1)
}

func TestBoardCreateCommentBuilder(t *testing.T) {
	t.Parallel()

	b := params.NewBoardCreateCommentBuilder()

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

func TestBoardDeleteCommentBuilder(t *testing.T) {
	t.Parallel()

	b := params.NewBoardDeleteCommentBuilder()

	b.GroupID(1)
	b.TopicID(1)
	b.CommentID(1)

	assert.Equal(t, b.Params["group_id"], 1)
	assert.Equal(t, b.Params["topic_id"], 1)
	assert.Equal(t, b.Params["comment_id"], 1)
}

func TestBoardDeleteTopicBuilder(t *testing.T) {
	t.Parallel()

	b := params.NewBoardDeleteTopicBuilder()

	b.GroupID(1)
	b.TopicID(1)

	assert.Equal(t, b.Params["group_id"], 1)
	assert.Equal(t, b.Params["topic_id"], 1)
}

func TestBoardEditCommentBuilder(t *testing.T) {
	t.Parallel()

	b := params.NewBoardEditCommentBuilder()

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

func TestBoardEditTopicBuilder(t *testing.T) {
	t.Parallel()

	b := params.NewBoardEditTopicBuilder()

	b.GroupID(1)
	b.TopicID(1)
	b.Title("text")

	assert.Equal(t, b.Params["group_id"], 1)
	assert.Equal(t, b.Params["topic_id"], 1)
	assert.Equal(t, b.Params["title"], "text")
}

func TestBoardFixTopicBuilder(t *testing.T) {
	t.Parallel()

	b := params.NewBoardFixTopicBuilder()

	b.GroupID(1)
	b.TopicID(1)

	assert.Equal(t, b.Params["group_id"], 1)
	assert.Equal(t, b.Params["topic_id"], 1)
}

func TestBoardGetCommentsBuilder(t *testing.T) {
	t.Parallel()

	b := params.NewBoardGetCommentsBuilder()

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

func TestBoardGetTopicsBuilder(t *testing.T) {
	t.Parallel()

	b := params.NewBoardGetTopicsBuilder()

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

func TestBoardOpenTopicBuilder(t *testing.T) {
	t.Parallel()

	b := params.NewBoardOpenTopicBuilder()

	b.GroupID(1)
	b.TopicID(1)

	assert.Equal(t, b.Params["group_id"], 1)
	assert.Equal(t, b.Params["topic_id"], 1)
}

func TestBoardRestoreCommentBuilder(t *testing.T) {
	t.Parallel()

	b := params.NewBoardRestoreCommentBuilder()

	b.GroupID(1)
	b.TopicID(1)
	b.CommentID(1)

	assert.Equal(t, b.Params["group_id"], 1)
	assert.Equal(t, b.Params["topic_id"], 1)
	assert.Equal(t, b.Params["comment_id"], 1)
}

func TestBoardUnfixTopicBuilder(t *testing.T) {
	t.Parallel()

	b := params.NewBoardUnfixTopicBuilder()

	b.GroupID(1)
	b.TopicID(1)

	assert.Equal(t, b.Params["group_id"], 1)
	assert.Equal(t, b.Params["topic_id"], 1)
}
