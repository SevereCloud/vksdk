package object_test

import (
	"testing"

	"github.com/SevereCloud/vksdk/object"
)

func TestGroupsGroup_ToMention(t *testing.T) {
	f := func(group object.GroupsGroup, want string) {
		if got := group.ToMention(); got != want {
			t.Errorf("GroupsGroup.ToMention() = %v, want %v", got, want)
		}
	}

	f(
		object.GroupsGroup{ID: 1, Name: "Api club"},
		"[club1|Api club]",
	)
}

func TestGroupsGroupXtrInvitedBy_ToMention(t *testing.T) {
	f := func(group object.GroupsGroupXtrInvitedBy, want string) {
		if got := group.ToMention(); got != want {
			t.Errorf("GroupsGroupXtrInvitedBy.ToMention() = %v, want %v", got, want)
		}
	}

	f(
		object.GroupsGroupXtrInvitedBy{ID: 1, Name: "Api club"},
		"[club1|Api club]",
	)
}
