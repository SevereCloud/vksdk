package params_test

import (
	"testing"

	"github.com/SevereCloud/vksdk/v2/api/params"
	"github.com/stretchr/testify/assert"
)

func TestNotesAddBuilder(t *testing.T) {
	t.Parallel()

	b := params.NewNotesAddBuilder()

	b.Title("text")
	b.Text("text")
	b.PrivacyView([]string{"text"})
	b.PrivacyComment([]string{"text"})

	assert.Equal(t, "text", b.Params["title"])
	assert.Equal(t, "text", b.Params["text"])
	assert.Equal(t, []string{"text"}, b.Params["privacy_view"])
	assert.Equal(t, []string{"text"}, b.Params["privacy_comment"])
}

func TestNotesCreateCommentBuilder(t *testing.T) {
	t.Parallel()

	b := params.NewNotesCreateCommentBuilder()

	b.NoteID(1)
	b.OwnerID(1)
	b.ReplyTo(1)
	b.Message("text")
	b.GUID("text")

	assert.Equal(t, 1, b.Params["note_id"])
	assert.Equal(t, 1, b.Params["owner_id"])
	assert.Equal(t, 1, b.Params["reply_to"])
	assert.Equal(t, "text", b.Params["message"])
	assert.Equal(t, "text", b.Params["guid"])
}

func TestNotesDeleteBuilder(t *testing.T) {
	t.Parallel()

	b := params.NewNotesDeleteBuilder()

	b.NoteID(1)

	assert.Equal(t, 1, b.Params["note_id"])
}

func TestNotesDeleteCommentBuilder(t *testing.T) {
	t.Parallel()

	b := params.NewNotesDeleteCommentBuilder()

	b.CommentID(1)
	b.OwnerID(1)

	assert.Equal(t, 1, b.Params["comment_id"])
	assert.Equal(t, 1, b.Params["owner_id"])
}

func TestNotesEditBuilder(t *testing.T) {
	t.Parallel()

	b := params.NewNotesEditBuilder()

	b.NoteID(1)
	b.Title("text")
	b.Text("text")
	b.PrivacyView([]string{"text"})
	b.PrivacyComment([]string{"text"})

	assert.Equal(t, 1, b.Params["note_id"])
	assert.Equal(t, "text", b.Params["title"])
	assert.Equal(t, "text", b.Params["text"])
	assert.Equal(t, []string{"text"}, b.Params["privacy_view"])
	assert.Equal(t, []string{"text"}, b.Params["privacy_comment"])
}

func TestNotesEditCommentBuilder(t *testing.T) {
	t.Parallel()

	b := params.NewNotesEditCommentBuilder()

	b.CommentID(1)
	b.OwnerID(1)
	b.Message("text")

	assert.Equal(t, 1, b.Params["comment_id"])
	assert.Equal(t, 1, b.Params["owner_id"])
	assert.Equal(t, "text", b.Params["message"])
}

func TestNotesGetBuilder(t *testing.T) {
	t.Parallel()

	b := params.NewNotesGetBuilder()

	b.NoteIDs([]int{1})
	b.UserID(1)
	b.Offset(1)
	b.Count(1)
	b.Sort(1)

	assert.Equal(t, []int{1}, b.Params["note_ids"])
	assert.Equal(t, 1, b.Params["user_id"])
	assert.Equal(t, 1, b.Params["offset"])
	assert.Equal(t, 1, b.Params["count"])
	assert.Equal(t, 1, b.Params["sort"])
}

func TestNotesGetByIDBuilder(t *testing.T) {
	t.Parallel()

	b := params.NewNotesGetByIDBuilder()

	b.NoteID(1)
	b.OwnerID(1)
	b.NeedWiki(true)

	assert.Equal(t, 1, b.Params["note_id"])
	assert.Equal(t, 1, b.Params["owner_id"])
	assert.Equal(t, true, b.Params["need_wiki"])
}

func TestNotesGetCommentsBuilder(t *testing.T) {
	t.Parallel()

	b := params.NewNotesGetCommentsBuilder()

	b.NoteID(1)
	b.OwnerID(1)
	b.Sort(1)
	b.Offset(1)
	b.Count(1)

	assert.Equal(t, 1, b.Params["note_id"])
	assert.Equal(t, 1, b.Params["owner_id"])
	assert.Equal(t, 1, b.Params["sort"])
	assert.Equal(t, 1, b.Params["offset"])
	assert.Equal(t, 1, b.Params["count"])
}

func TestNotesRestoreCommentBuilder(t *testing.T) {
	t.Parallel()

	b := params.NewNotesRestoreCommentBuilder()

	b.CommentID(1)
	b.OwnerID(1)

	assert.Equal(t, 1, b.Params["comment_id"])
	assert.Equal(t, 1, b.Params["owner_id"])
}
