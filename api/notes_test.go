package api

import (
	"strconv"
	"testing"

	"github.com/SevereCloud/vksdk/errors"
)

func TestVK_NotesAdd(t *testing.T) {
	needUserToken(t)

	note, err := vkUser.NotesAdd(map[string]string{
		"title": "Test note",
		"text":  "Text note",
	})
	if err != nil {
		t.Errorf("VK.NotesAdd() err = %v", err)
	}

	_, err = vkUser.NotesEdit(map[string]string{
		"note_id": strconv.Itoa(note),
		"title":   "Test note edited",
		"text":    "Text note edited",
	})
	if err != nil {
		t.Errorf("VK.NotesEdit() err = %v", err)
	}

	_, err = vkUser.NotesDelete(map[string]string{
		"note_id": strconv.Itoa(note),
	})
	if err != nil {
		t.Errorf("VK.NotesDelete() err = %v", err)
	}
}

func TestVK_NotesCreateComment(t *testing.T) {
	needUserToken(t)

	comment, err := vkUser.NotesCreateComment(map[string]string{
		"note_id":  "11751647",
		"owner_id": "2314852",
		"message":  "Test note comment",
	})
	if err != nil {
		t.Errorf("VK.NotesCreateComment() err = %v", err)
	}

	_, err = vkUser.NotesEditComment(map[string]string{
		"comment_id": strconv.Itoa(comment),
		"owner_id":   "2314852",
		"message":    "Test note commentedited",
	})
	if errors.GetType(err) != errors.Param {
		t.Errorf("VK.NotesEditComment() err = %v", err)
	}

	_, err = vkUser.NotesDeleteComment(map[string]string{
		"comment_id": strconv.Itoa(comment),
		"owner_id":   "2314852",
	})
	if errors.GetType(err) != errors.Param {
		t.Errorf("VK.NotesDeleteComment() err = %v", err)
	}

	_, err = vkUser.NotesRestoreComment(map[string]string{
		"comment_id": strconv.Itoa(comment),
		"owner_id":   "2314852",
	})
	if errors.GetType(err) != errors.Param {
		t.Errorf("VK.NotesRestoreComment() err = %v", err)
	}
}

func TestVK_NotesGet(t *testing.T) {
	needUserToken(t)

	_, err := vkUser.NotesGet(map[string]string{
		"user_id": "66748",
	})
	if err != nil {
		t.Errorf("VK.NotesGet() err = %v", err)
	}
}

func TestVK_NotesGetByID(t *testing.T) {
	needUserToken(t)

	_, err := vkUser.NotesGetByID(map[string]string{
		"note_id":   "11582572",
		"owner_id":  "66748",
		"need_wiki": "1",
	})
	if err != nil {
		t.Errorf("VK.NotesGetByID() err = %v", err)
	}
}

func TestVK_NotesGetComments(t *testing.T) {
	needUserToken(t)

	_, err := vkUser.NotesGetComments(map[string]string{
		"note_id":  "11582572",
		"owner_id": "66748",
	})
	if err != nil {
		t.Errorf("VK.NotesGetComments() err = %v", err)
	}
}
