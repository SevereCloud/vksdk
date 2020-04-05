package object_test

import (
	"testing"

	"github.com/SevereCloud/vksdk/object"
	"github.com/stretchr/testify/assert"
)

func TestGroupsGroup_ToMention(t *testing.T) {
	t.Parallel()

	f := func(group object.GroupsGroup, want string) {
		t.Helper()

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
	t.Parallel()

	f := func(group object.GroupsGroupXtrInvitedBy, want string) {
		t.Helper()

		if got := group.ToMention(); got != want {
			t.Errorf("GroupsGroupXtrInvitedBy.ToMention() = %v, want %v", got, want)
		}
	}

	f(
		object.GroupsGroupXtrInvitedBy{ID: 1, Name: "Api club"},
		"[club1|Api club]",
	)
}

func TestGroupsSectionsList_UnmarshalJSON(t *testing.T) {
	t.Parallel()

	f := func(data []byte, want object.GroupsSectionsList, wanErr bool) {
		t.Helper()

		var got object.GroupsSectionsList

		err := got.UnmarshalJSON(data)

		if wanErr {
			assert.Error(t, err)
		} else {
			assert.NoError(t, err)
			assert.Equal(t, want, got)
		}
	}

	f(
		[]byte(`[123, "test"]`),
		object.GroupsSectionsList{123, "test"},
		false,
	)

	// Errors:
	f(
		[]byte(`123`),
		object.GroupsSectionsList{},
		true,
	)
	f(
		[]byte(`[123, "123", "123"]`),
		object.GroupsSectionsList{},
		true,
	)
	f(
		[]byte(`["123", "123"]`),
		object.GroupsSectionsList{},
		true,
	)
	f(
		[]byte(`[123, 123]`),
		object.GroupsSectionsList{},
		true,
	)
}
