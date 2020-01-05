package object_test

import (
	"testing"

	"github.com/SevereCloud/vksdk/object"
)

func TestNotesNote_ToAttachment(t *testing.T) {
	f := func(note object.NotesNote, want string) {
		if got := note.ToAttachment(); got != want {
			t.Errorf("NotesNote.ToAttachment() = %v, want %v", got, want)
		}
	}

	f(object.NotesNote{ID: 10, OwnerID: 20}, "note20_10")
	f(object.NotesNote{ID: 20, OwnerID: -10}, "note-10_20")
}
