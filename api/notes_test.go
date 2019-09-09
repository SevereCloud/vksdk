package api

import (
	"strconv"
	"testing"

	"github.com/SevereCloud/vksdk/object"
)

func TestVK_NotesAdd(t *testing.T) {
	if vkUser.AccessToken == "" {
		t.Skip("USER_TOKEN empty")
	}

	note, gotVkErr := vkUser.NotesAdd(map[string]string{
		"title": "Test note",
		"text":  "Text note",
	})
	if gotVkErr.Code != 0 {
		t.Errorf("VK.NotesAdd() gotVkErr = %v, want %v", gotVkErr, 0)
	}

	_, gotVkErr = vkUser.NotesEdit(map[string]string{
		"note_id": strconv.Itoa(note),
		"title":   "Test note edited",
		"text":    "Text note edited",
	})
	if gotVkErr.Code != 0 {
		t.Errorf("VK.NotesEdit() gotVkErr = %v, want %v", gotVkErr, 0)
	}

	_, gotVkErr = vkUser.NotesDelete(map[string]string{
		"note_id": strconv.Itoa(note),
	})
	if gotVkErr.Code != 0 {
		t.Errorf("VK.NotesDelete() gotVkErr = %v, want %v", gotVkErr, 0)
	}
}

func TestVK_NotesCreateComment(t *testing.T) {
	if vkUser.AccessToken == "" {
		t.Skip("USER_TOKEN empty")
	}

	comment, gotVkErr := vkUser.NotesCreateComment(map[string]string{
		"note_id":  "11751647",
		"owner_id": "2314852",
		"message":  "Test note comment",
	})
	if gotVkErr.Code != 0 {
		t.Errorf("VK.NotesCreateComment() gotVkErr = %v, want %v", gotVkErr, 0)
	}

	_, gotVkErr = vkUser.NotesEditComment(map[string]string{
		"comment_id": strconv.Itoa(comment),
		"owner_id":   "2314852",
		"message":    "Test note commentedited",
	})
	if gotVkErr.Code != object.ErrorParam {
		t.Errorf("VK.NotesEditComment() gotVkErr = %v, want %v", gotVkErr, 0)
	}

	_, gotVkErr = vkUser.NotesDeleteComment(map[string]string{
		"comment_id": strconv.Itoa(comment),
		"owner_id":   "2314852",
	})
	if gotVkErr.Code != object.ErrorParam {
		t.Errorf("VK.NotesDeleteComment() gotVkErr = %v, want %v", gotVkErr, 0)
	}

	_, gotVkErr = vkUser.NotesRestoreComment(map[string]string{
		"comment_id": strconv.Itoa(comment),
		"owner_id":   "2314852",
	})
	if gotVkErr.Code != object.ErrorParam {
		t.Errorf("VK.NotesRestoreComment() gotVkErr = %v, want %v", gotVkErr, 0)
	}
}

func TestVK_NotesGet(t *testing.T) {
	if vkUser.AccessToken == "" {
		t.Skip("USER_TOKEN empty")
	}

	_, gotVkErr := vkUser.NotesGet(map[string]string{
		"user_id": "66748",
	})
	if gotVkErr.Code != 0 {
		t.Errorf("VK.NotesGet() gotVkErr = %v, want %v", gotVkErr, 0)
	}
}

func TestVK_NotesGetByID(t *testing.T) {
	if vkUser.AccessToken == "" {
		t.Skip("USER_TOKEN empty")
	}

	_, gotVkErr := vkUser.NotesGetByID(map[string]string{
		"note_id":   "11582572",
		"owner_id":  "66748",
		"need_wiki": "1",
	})
	if gotVkErr.Code != 0 {
		t.Errorf("VK.NotesGetByID() gotVkErr = %v, want %v", gotVkErr, 0)
	}
}

func TestVK_NotesGetComments(t *testing.T) {
	if vkUser.AccessToken == "" {
		t.Skip("USER_TOKEN empty")
	}

	_, gotVkErr := vkUser.NotesGetComments(map[string]string{
		"note_id":  "11582572",
		"owner_id": "66748",
	})
	if gotVkErr.Code != 0 {
		t.Errorf("VK.NotesGetComments() gotVkErr = %v, want %v", gotVkErr, 0)
	}
}
