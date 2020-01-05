package params_test

import (
	"testing"

	"github.com/SevereCloud/vksdk/api/params"
	"github.com/stretchr/testify/assert"
)

func TestNotesAddBuilder(t *testing.T) {
	b := params.NewNotesAddBuilder()

	b.Title("text")
	b.Text("text")
	b.PrivacyView([]string{"text"})
	b.PrivacyComment([]string{"text"})

	assert.Equal(t, b.Params["title"], "text")
	assert.Equal(t, b.Params["text"], "text")
	assert.Equal(t, b.Params["privacy_view"], []string{"text"})
	assert.Equal(t, b.Params["privacy_comment"], []string{"text"})
}

func TestNotesCreateCommentBuilder(t *testing.T) {
	b := params.NewNotesCreateCommentBuilder()

	b.NoteID(1)
	b.OwnerID(1)
	b.ReplyTo(1)
	b.Message("text")
	b.GUID("text")

	assert.Equal(t, b.Params["note_id"], 1)
	assert.Equal(t, b.Params["owner_id"], 1)
	assert.Equal(t, b.Params["reply_to"], 1)
	assert.Equal(t, b.Params["message"], "text")
	assert.Equal(t, b.Params["guid"], "text")
}

func TestNotesDeleteBuilder(t *testing.T) {
	b := params.NewNotesDeleteBuilder()

	b.NoteID(1)

	assert.Equal(t, b.Params["note_id"], 1)
}

func TestNotesDeleteCommentBuilder(t *testing.T) {
	b := params.NewNotesDeleteCommentBuilder()

	b.CommentID(1)
	b.OwnerID(1)

	assert.Equal(t, b.Params["comment_id"], 1)
	assert.Equal(t, b.Params["owner_id"], 1)
}

func TestNotesEditBuilder(t *testing.T) {
	b := params.NewNotesEditBuilder()

	b.NoteID(1)
	b.Title("text")
	b.Text("text")
	b.PrivacyView([]string{"text"})
	b.PrivacyComment([]string{"text"})

	assert.Equal(t, b.Params["note_id"], 1)
	assert.Equal(t, b.Params["title"], "text")
	assert.Equal(t, b.Params["text"], "text")
	assert.Equal(t, b.Params["privacy_view"], []string{"text"})
	assert.Equal(t, b.Params["privacy_comment"], []string{"text"})
}

func TestNotesEditCommentBuilder(t *testing.T) {
	b := params.NewNotesEditCommentBuilder()

	b.CommentID(1)
	b.OwnerID(1)
	b.Message("text")

	assert.Equal(t, b.Params["comment_id"], 1)
	assert.Equal(t, b.Params["owner_id"], 1)
	assert.Equal(t, b.Params["message"], "text")
}

func TestNotesGetBuilder(t *testing.T) {
	b := params.NewNotesGetBuilder()

	b.NoteIDs([]int{1})
	b.UserID(1)
	b.Offset(1)
	b.Count(1)
	b.Sort(1)

	assert.Equal(t, b.Params["note_ids"], []int{1})
	assert.Equal(t, b.Params["user_id"], 1)
	assert.Equal(t, b.Params["offset"], 1)
	assert.Equal(t, b.Params["count"], 1)
	assert.Equal(t, b.Params["sort"], 1)
}

func TestNotesGetByIDBuilder(t *testing.T) {
	b := params.NewNotesGetByIDBuilder()

	b.NoteID(1)
	b.OwnerID(1)
	b.NeedWiki(true)

	assert.Equal(t, b.Params["note_id"], 1)
	assert.Equal(t, b.Params["owner_id"], 1)
	assert.Equal(t, b.Params["need_wiki"], true)
}

func TestNotesGetCommentsBuilder(t *testing.T) {
	b := params.NewNotesGetCommentsBuilder()

	b.NoteID(1)
	b.OwnerID(1)
	b.Sort(1)
	b.Offset(1)
	b.Count(1)

	assert.Equal(t, b.Params["note_id"], 1)
	assert.Equal(t, b.Params["owner_id"], 1)
	assert.Equal(t, b.Params["sort"], 1)
	assert.Equal(t, b.Params["offset"], 1)
	assert.Equal(t, b.Params["count"], 1)
}

func TestNotesRestoreCommentBuilder(t *testing.T) {
	b := params.NewNotesRestoreCommentBuilder()

	b.CommentID(1)
	b.OwnerID(1)

	assert.Equal(t, b.Params["comment_id"], 1)
	assert.Equal(t, b.Params["owner_id"], 1)
}
