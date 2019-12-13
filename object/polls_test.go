package object

import "testing"

func TestPollsPoll_ToAttachment(t *testing.T) {
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
			name:   "poll20_10",
			fields: fields{10, 20},
			want:   "poll20_10",
		},
		{
			name:   "poll-10_20",
			fields: fields{20, -10},
			want:   "poll-10_20",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			poll := PollsPoll{
				ID:      tt.fields.ID,
				OwnerID: tt.fields.OwnerID,
			}
			if got := poll.ToAttachment(); got != tt.want {
				t.Errorf("PollsPoll.ToAttachment() = %v, want %v", got, tt.want)
			}
		})
	}
}
