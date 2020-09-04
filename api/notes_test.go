package api_test

import (
	"testing"

	"github.com/SevereCloud/vksdk/api"
	"github.com/stretchr/testify/assert"
)

func TestVK_NotesAdd(t *testing.T) {
	t.Parallel()

	needUserToken(t)

	note, err := vkUser.NotesAdd(api.Params{
		"title": "Test note",
		"text":  "Text note",
	})
	noError(t, err)
	assert.NotEmpty(t, note)

	res, err := vkUser.NotesEdit(api.Params{
		"note_id": note,
		"title":   "Test note edited",
		"text":    "Text note edited",
	})
	noError(t, err)
	assert.NotEmpty(t, res)

	res, err = vkUser.NotesDelete(api.Params{
		"note_id": note,
	})
	noError(t, err)
	assert.NotEmpty(t, res)
}

func TestVK_NotesCreateComment(t *testing.T) {
	t.Parallel()

	needUserToken(t)

	comment, err := vkUser.NotesCreateComment(api.Params{
		"note_id":  11751647,
		"owner_id": 2314852,
		"message":  "Test note comment",
	})
	// NOTE: note comment deprecated
	noError(t, err)
	// assert.NotEmpty(t, comment)

	_, _ = vkUser.NotesEditComment(api.Params{
		"comment_id": comment,
		"owner_id":   2314852,
		"message":    "Test note comment edited",
	})
	// assert.Equal(t, errors.GetType(err), errors.Param)

	_, _ = vkUser.NotesDeleteComment(api.Params{
		"comment_id": comment,
		"owner_id":   2314852,
	})
	// assert.Equal(t, errors.GetType(err), errors.Param)

	// assert.Equal(t, errors.GetType(err), errors.Param)
	_, _ = vkUser.NotesRestoreComment(api.Params{
		"comment_id": comment,
		"owner_id":   2314852,
	})
}

func TestVK_NotesGet(t *testing.T) {
	t.Parallel()

	needUserToken(t)

	res, err := vkUser.NotesGet(api.Params{
		"user_id": 66748,
	})
	noError(t, err)
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
	t.Parallel()

	needUserToken(t)

	res, err := vkUser.NotesGetByID(api.Params{
		"note_id":   11582572,
		"owner_id":  66748,
		"need_wiki": true,
	})
	noError(t, err)
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
	t.Parallel()

	needUserToken(t)

	_, err := vkUser.NotesGetComments(api.Params{
		"note_id":  11582572,
		"owner_id": 66748,
	})
	noError(t, err)
}
