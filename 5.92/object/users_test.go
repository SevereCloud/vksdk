package object

import "testing"

func TestUsersUser_ToMention(t *testing.T) {
	type fields struct {
		ID        int
		FirstName string
		LastName  string
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			fields: fields{1, "Ivan", "Ivanov"},
			want:   "[id1|Ivan Ivanov]",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			user := UsersUser{
				ID:        tt.fields.ID,
				FirstName: tt.fields.FirstName,
				LastName:  tt.fields.LastName,
			}
			if got := user.ToMention(); got != tt.want {
				t.Errorf("UsersUser.ToMention() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUsersUserMin_ToMention(t *testing.T) {
	type fields struct {
		ID        int
		FirstName string
		LastName  string
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			fields: fields{1, "Ivan", "Ivanov"},
			want:   "[id1|Ivan Ivanov]",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			user := UsersUserMin{
				ID:        tt.fields.ID,
				FirstName: tt.fields.FirstName,
				LastName:  tt.fields.LastName,
			}
			if got := user.ToMention(); got != tt.want {
				t.Errorf("UsersUserMin.ToMention() = %v, want %v", got, tt.want)
			}
		})
	}
}
