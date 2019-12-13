package object

import "testing"

func TestAudioAudioFull_ToAttachment(t *testing.T) {
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
			name:   "audio20_10",
			fields: fields{10, 20},
			want:   "audio20_10",
		},
		{
			name:   "audio-10_20",
			fields: fields{20, -10},
			want:   "audio-10_20",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			audio := AudioAudioFull{
				ID:      tt.fields.ID,
				OwnerID: tt.fields.OwnerID,
			}
			if got := audio.ToAttachment(); got != tt.want {
				t.Errorf("AudioAudioFull.ToAttachment() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAudioAudio_ToAttachment(t *testing.T) {
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
			name:   "audio20_10",
			fields: fields{10, 20},
			want:   "audio20_10",
		},
		{
			name:   "audio-10_20",
			fields: fields{20, -10},
			want:   "audio-10_20",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			audio := AudioAudio{
				ID:      tt.fields.ID,
				OwnerID: tt.fields.OwnerID,
			}
			if got := audio.ToAttachment(); got != tt.want {
				t.Errorf("AudioAudio.ToAttachment() = %v, want %v", got, tt.want)
			}
		})
	}
}
