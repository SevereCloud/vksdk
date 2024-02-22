package params_test

import (
	"testing"

	"github.com/SevereCloud/vksdk/v2/api/params"
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

	assert.Equal(t, 1, b.Params["group_id"])
	assert.Equal(t, "text", b.Params["title"])
	assert.Equal(t, "text", b.Params["text"])
	assert.Equal(t, true, b.Params["from_group"])
	assert.Equal(t, []string{"text"}, b.Params["attachments"])
}

func TestBoardCloseTopicBuilder(t *testing.T) {
	t.Parallel()

	b := params.NewBoardCloseTopicBuilder()

	b.GroupID(1)
	b.TopicID(1)

	assert.Equal(t, 1, b.Params["group_id"])
	assert.Equal(t, 1, b.Params["topic_id"])
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

	assert.Equal(t, 1, b.Params["group_id"])
	assert.Equal(t, 1, b.Params["topic_id"])
	assert.Equal(t, "text", b.Params["message"])
	assert.Equal(t, []string{"text"}, b.Params["attachments"])
	assert.Equal(t, true, b.Params["from_group"])
	assert.Equal(t, 1, b.Params["sticker_id"])
	assert.Equal(t, "text", b.Params["guid"])
}

func TestBoardDeleteCommentBuilder(t *testing.T) {
	t.Parallel()

	b := params.NewBoardDeleteCommentBuilder()

	b.GroupID(1)
	b.TopicID(1)
	b.CommentID(1)

	assert.Equal(t, 1, b.Params["group_id"])
	assert.Equal(t, 1, b.Params["topic_id"])
	assert.Equal(t, 1, b.Params["comment_id"])
}

func TestBoardDeleteTopicBuilder(t *testing.T) {
	t.Parallel()

	b := params.NewBoardDeleteTopicBuilder()

	b.GroupID(1)
	b.TopicID(1)

	assert.Equal(t, 1, b.Params["group_id"])
	assert.Equal(t, 1, b.Params["topic_id"])
}

func TestBoardEditCommentBuilder(t *testing.T) {
	t.Parallel()

	b := params.NewBoardEditCommentBuilder()

	b.GroupID(1)
	b.TopicID(1)
	b.CommentID(1)
	b.Message("text")
	b.Attachments([]string{"text"})

	assert.Equal(t, 1, b.Params["group_id"])
	assert.Equal(t, 1, b.Params["topic_id"])
	assert.Equal(t, 1, b.Params["comment_id"])
	assert.Equal(t, "text", b.Params["message"])
	assert.Equal(t, []string{"text"}, b.Params["attachments"])
}

func TestBoardEditTopicBuilder(t *testing.T) {
	t.Parallel()

	b := params.NewBoardEditTopicBuilder()

	b.GroupID(1)
	b.TopicID(1)
	b.Title("text")

	assert.Equal(t, 1, b.Params["group_id"])
	assert.Equal(t, 1, b.Params["topic_id"])
	assert.Equal(t, "text", b.Params["title"])
}

func TestBoardFixTopicBuilder(t *testing.T) {
	t.Parallel()

	b := params.NewBoardFixTopicBuilder()

	b.GroupID(1)
	b.TopicID(1)

	assert.Equal(t, 1, b.Params["group_id"])
	assert.Equal(t, 1, b.Params["topic_id"])
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

	assert.Equal(t, 1, b.Params["group_id"])
	assert.Equal(t, 1, b.Params["topic_id"])
	assert.Equal(t, true, b.Params["need_likes"])
	assert.Equal(t, 1, b.Params["start_comment_id"])
	assert.Equal(t, 1, b.Params["offset"])
	assert.Equal(t, 1, b.Params["count"])
	assert.Equal(t, true, b.Params["extended"])
	assert.Equal(t, "text", b.Params["sort"])
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

	assert.Equal(t, 1, b.Params["group_id"])
	assert.Equal(t, []int{1}, b.Params["topic_ids"])
	assert.Equal(t, 1, b.Params["order"])
	assert.Equal(t, 1, b.Params["offset"])
	assert.Equal(t, 1, b.Params["count"])
	assert.Equal(t, true, b.Params["extended"])
	assert.Equal(t, 1, b.Params["preview"])
	assert.Equal(t, 1, b.Params["preview_length"])
}

func TestBoardOpenTopicBuilder(t *testing.T) {
	t.Parallel()

	b := params.NewBoardOpenTopicBuilder()

	b.GroupID(1)
	b.TopicID(1)

	assert.Equal(t, 1, b.Params["group_id"])
	assert.Equal(t, 1, b.Params["topic_id"])
}

func TestBoardRestoreCommentBuilder(t *testing.T) {
	t.Parallel()

	b := params.NewBoardRestoreCommentBuilder()

	b.GroupID(1)
	b.TopicID(1)
	b.CommentID(1)

	assert.Equal(t, 1, b.Params["group_id"])
	assert.Equal(t, 1, b.Params["topic_id"])
	assert.Equal(t, 1, b.Params["comment_id"])
}

func TestBoardUnfixTopicBuilder(t *testing.T) {
	t.Parallel()

	b := params.NewBoardUnfixTopicBuilder()

	b.GroupID(1)
	b.TopicID(1)

	assert.Equal(t, 1, b.Params["group_id"])
	assert.Equal(t, 1, b.Params["topic_id"])
}
