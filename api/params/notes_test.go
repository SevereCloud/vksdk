package params_test

import (
	"testing"

	"github.com/SevereCloud/vksdk/api/params"
	"github.com/stretchr/testify/assert"
)

func TestNotesAddBulder(t *testing.T) {
	b := params.NewNotesAddBulder()

	b.Title("text")
	b.Text("text")
	b.PrivacyView([]string{"text"})
	b.PrivacyComment([]string{"text"})

	assert.Equal(t, b.Params["title"], "text")
	assert.Equal(t, b.Params["text"], "text")
	assert.Equal(t, b.Params["privacy_view"], []string{"text"})
	assert.Equal(t, b.Params["privacy_comment"], []string{"text"})
}

func TestNotesCreateCommentBulder(t *testing.T) {
	b := params.NewNotesCreateCommentBulder()

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

func TestNotesDeleteBulder(t *testing.T) {
	b := params.NewNotesDeleteBulder()

	b.NoteID(1)

	assert.Equal(t, b.Params["note_id"], 1)
}

func TestNotesDeleteCommentBulder(t *testing.T) {
	b := params.NewNotesDeleteCommentBulder()

	b.CommentID(1)
	b.OwnerID(1)

	assert.Equal(t, b.Params["comment_id"], 1)
	assert.Equal(t, b.Params["owner_id"], 1)
}

func TestNotesEditBulder(t *testing.T) {
	b := params.NewNotesEditBulder()

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

func TestNotesEditCommentBulder(t *testing.T) {
	b := params.NewNotesEditCommentBulder()

	b.CommentID(1)
	b.OwnerID(1)
	b.Message("text")

	assert.Equal(t, b.Params["comment_id"], 1)
	assert.Equal(t, b.Params["owner_id"], 1)
	assert.Equal(t, b.Params["message"], "text")
}

func TestNotesGetBulder(t *testing.T) {
	b := params.NewNotesGetBulder()

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

func TestNotesGetByIDBulder(t *testing.T) {
	b := params.NewNotesGetByIDBulder()

	b.NoteID(1)
	b.OwnerID(1)
	b.NeedWiki(true)

	assert.Equal(t, b.Params["note_id"], 1)
	assert.Equal(t, b.Params["owner_id"], 1)
	assert.Equal(t, b.Params["need_wiki"], true)
}

func TestNotesGetCommentsBulder(t *testing.T) {
	b := params.NewNotesGetCommentsBulder()

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

func TestNotesRestoreCommentBulder(t *testing.T) {
	b := params.NewNotesRestoreCommentBulder()

	b.CommentID(1)
	b.OwnerID(1)

	assert.Equal(t, b.Params["comment_id"], 1)
	assert.Equal(t, b.Params["owner_id"], 1)
}
