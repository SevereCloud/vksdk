package packer

import "testing"

func TestEscape(t *testing.T) {
	tests := []struct {
		name string
		args string
		want string
	}{
		{
			name: "test1",
			args: `сообщение", "peer_id":"1`,
			want: `сообщение\", \"peer_id\":\"1`,
		},
		{
			name: "test2",
			args: `сообщение\", "peer_id":"1`,
			want: `сообщение\", \"peer_id\":\"1`,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := escape(tt.args); got != tt.want {
				t.Errorf("escape() = %v, want %v", got, tt.want)
			}
		})
	}
}
