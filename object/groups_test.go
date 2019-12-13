package object

import "testing"

func TestGroupsGroup_ToMention(t *testing.T) {
	type fields struct {
		ID   int
		Name string
	}

	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			fields: fields{1, "Api club"},
			want:   "[club1|Api club]",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			group := GroupsGroup{
				ID:   tt.fields.ID,
				Name: tt.fields.Name,
			}
			if got := group.ToMention(); got != tt.want {
				t.Errorf("GroupsGroup.ToMention() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGroupsGroupXtrInvitedBy_ToMention(t *testing.T) {
	type fields struct {
		ID   int
		Name string
	}

	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			fields: fields{1, "Api club"},
			want:   "[club1|Api club]",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			group := GroupsGroupXtrInvitedBy{
				ID:   tt.fields.ID,
				Name: tt.fields.Name,
			}
			if got := group.ToMention(); got != tt.want {
				t.Errorf("GroupsGroupXtrInvitedBy.ToMention() = %v, want %v", got, tt.want)
			}
		})
	}
}
