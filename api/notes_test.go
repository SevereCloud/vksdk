package api

import (
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/SevereCloud/vksdk/errors"
)

func TestVK_NotesAdd(t *testing.T) {
	needUserToken(t)

	note, err := vkUser.NotesAdd(map[string]string{
		"title": "Test note",
		"text":  "Text note",
	})
	assert.NoError(t, err)

	_, err = vkUser.NotesEdit(map[string]string{
		"note_id": strconv.Itoa(note),
		"title":   "Test note edited",
		"text":    "Text note edited",
	})
	assert.NoError(t, err)

	_, err = vkUser.NotesDelete(map[string]string{
		"note_id": strconv.Itoa(note),
	})
	assert.NoError(t, err)
}

func TestVK_NotesCreateComment(t *testing.T) {
	needUserToken(t)

	comment, err := vkUser.NotesCreateComment(map[string]string{
		"note_id":  "11751647",
		"owner_id": "2314852",
		"message":  "Test note comment",
	})
	assert.NoError(t, err)

	_, err = vkUser.NotesEditComment(map[string]string{
		"comment_id": strconv.Itoa(comment),
		"owner_id":   "2314852",
		"message":    "Test note commentedited",
	})
	assert.Equal(t, errors.GetType(err), errors.Param)

	_, err = vkUser.NotesDeleteComment(map[string]string{
		"comment_id": strconv.Itoa(comment),
		"owner_id":   "2314852",
	})
	assert.Equal(t, errors.GetType(err), errors.Param)

	_, err = vkUser.NotesRestoreComment(map[string]string{
		"comment_id": strconv.Itoa(comment),
		"owner_id":   "2314852",
	})
	assert.Equal(t, errors.GetType(err), errors.Param)
}

func TestVK_NotesGet(t *testing.T) {
	needUserToken(t)

	_, err := vkUser.NotesGet(map[string]string{
		"user_id": "66748",
	})
	assert.NoError(t, err)
}

func TestVK_NotesGetByID(t *testing.T) {
	needUserToken(t)

	_, err := vkUser.NotesGetByID(map[string]string{
		"note_id":   "11582572",
		"owner_id":  "66748",
		"need_wiki": "1",
	})
	assert.NoError(t, err)
}

func TestVK_NotesGetComments(t *testing.T) {
	needUserToken(t)

	_, err := vkUser.NotesGetComments(map[string]string{
		"note_id":  "11582572",
		"owner_id": "66748",
	})
	assert.NoError(t, err)
}
