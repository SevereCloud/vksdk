package object

import "testing"

func TestVideoVideo_ToAttachment(t *testing.T) {
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
			name:   "video20_10",
			fields: fields{10, 20},
			want:   "video20_10",
		},
		{
			name:   "video-10_20",
			fields: fields{20, -10},
			want:   "video-10_20",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			video := VideoVideo{
				ID:      tt.fields.ID,
				OwnerID: tt.fields.OwnerID,
			}
			if got := video.ToAttachment(); got != tt.want {
				t.Errorf("VideoVideo.ToAttachment() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestVideoVideoFull_ToAttachment(t *testing.T) {
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
			name:   "video20_10",
			fields: fields{10, 20},
			want:   "video20_10",
		},
		{
			name:   "video-10_20",
			fields: fields{20, -10},
			want:   "video-10_20",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			video := VideoVideoFull{
				ID:      tt.fields.ID,
				OwnerID: tt.fields.OwnerID,
			}
			if got := video.ToAttachment(); got != tt.want {
				t.Errorf("VideoVideoFull.ToAttachment() = %v, want %v", got, tt.want)
			}
		})
	}
}
