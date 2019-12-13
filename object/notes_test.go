package object

import "testing"

func TestNotesNote_ToAttachment(t *testing.T) {
	type fields struct {
		ID      int
		OwnerID int
	}

	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name:   "note20_10",
			fields: fields{10, 20},
			want:   "note20_10",
		},
		{
			name:   "note-10_20",
			fields: fields{20, -10},
			want:   "note-10_20",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			note := NotesNote{
				ID:      tt.fields.ID,
				OwnerID: tt.fields.OwnerID,
			}
			if got := note.ToAttachment(); got != tt.want {
				t.Errorf("NotesNote.ToAttachment() = %v, want %v", got, tt.want)
			}
		})
	}
}
