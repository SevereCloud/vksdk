package api

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/SevereCloud/vksdk/errors"
)

func TestVK_NotesAdd(t *testing.T) {
	needUserToken(t)

	note, err := vkUser.NotesAdd(Params{
		"title": "Test note",
		"text":  "Text note",
	})
	assert.NoError(t, err)
	assert.NotEmpty(t, note)

	res, err := vkUser.NotesEdit(Params{
		"note_id": note,
		"title":   "Test note edited",
		"text":    "Text note edited",
	})
	assert.NoError(t, err)
	assert.NotEmpty(t, res)

	res, err = vkUser.NotesDelete(Params{
		"note_id": note,
	})
	assert.NoError(t, err)
	assert.NotEmpty(t, res)
}

func TestVK_NotesCreateComment(t *testing.T) {
	needUserToken(t)

	comment, err := vkUser.NotesCreateComment(Params{
		"note_id":  11751647,
		"owner_id": 2314852,
		"message":  "Test note comment",
	})
	// NOTE: note comment deprecated
	assert.NoError(t, err)
	// assert.NotEmpty(t, comment)

	_, err = vkUser.NotesEditComment(Params{
		"comment_id": comment,
		"owner_id":   2314852,
		"message":    "Test note commentedited",
	})
	assert.Equal(t, errors.GetType(err), errors.Param)

	_, err = vkUser.NotesDeleteComment(Params{
		"comment_id": comment,
		"owner_id":   2314852,
	})
	assert.Equal(t, errors.GetType(err), errors.Param)

	// assert.Equal(t, errors.GetType(err), errors.Param)
	_, _ = vkUser.NotesRestoreComment(Params{
		"comment_id": comment,
		"owner_id":   2314852,
	})
}

func TestVK_NotesGet(t *testing.T) {
	needUserToken(t)

	res, err := vkUser.NotesGet(Params{
		"user_id": 66748,
	})
	assert.NoError(t, err)
	assert.NotEmpty(t, res.Count)

	if assert.NotEmpty(t, res.Items) {
		assert.NotEmpty(t, res.Items[0].ID)
		assert.NotEmpty(t, res.Items[0].OwnerID)
		// assert.NotEmpty(t, res.Items[0].Comments)
		assert.NotEmpty(t, res.Items[0].Date)
		assert.NotEmpty(t, res.Items[0].Title)
		assert.NotEmpty(t, res.Items[0].Text)
		assert.NotEmpty(t, res.Items[0].ViewURL)
	}
}

func TestVK_NotesGetByID(t *testing.T) {
	needUserToken(t)

	res, err := vkUser.NotesGetByID(Params{
		"note_id":   11582572,
		"owner_id":  66748,
		"need_wiki": true,
	})
	assert.NoError(t, err)
	assert.NotEmpty(t, res.ID)
	assert.NotEmpty(t, res.OwnerID)
	// assert.NotEmpty(t, res.Items[0].Comments)
	assert.NotEmpty(t, res.Date)
	assert.NotEmpty(t, res.Title)
	assert.NotEmpty(t, res.Text)
	assert.NotEmpty(t, res.CanComment)
	assert.NotEmpty(t, res.ViewURL)
}

func TestVK_NotesGetComments(t *testing.T) {
	needUserToken(t)

	_, err := vkUser.NotesGetComments(Params{
		"note_id":  11582572,
		"owner_id": 66748,
	})
	assert.NoError(t, err)
}
