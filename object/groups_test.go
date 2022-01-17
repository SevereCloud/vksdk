package object_test

import (
	"encoding/json"
	"testing"

	"github.com/SevereCloud/vksdk/v2/object"
	"github.com/stretchr/testify/assert"
	"github.com/vmihailenco/msgpack/v5"
	"github.com/vmihailenco/msgpack/v5/msgpcode"
)

func TestGroupsCountersGroup_UnmarshalJSON(t *testing.T) {
	t.Parallel()

	f := func(data []byte, want object.GroupsCountersGroup) {
		t.Helper()

		var counters object.GroupsCountersGroup

		err := json.Unmarshal(data, &counters)
		assert.NoError(t, err)
		assert.Equal(t, want, counters)
	}

	f([]byte("[]"), object.GroupsCountersGroup{})
	f([]byte(`{"docs":1}`), object.GroupsCountersGroup{Docs: 1})

	var counters object.GroupsCountersGroup
	err := json.Unmarshal([]byte("0"), &counters)
	assert.Error(t, err)
}

func TestGroupsCountersGroup_DecodeMsgpack(t *testing.T) {
	t.Parallel()

	f := func(data []byte, want object.GroupsCountersGroup, wantErr string) {
		t.Helper()

		var counters object.GroupsCountersGroup

		err := msgpack.Unmarshal(data, &counters)
		if err != nil || wantErr != "" {
			assert.EqualError(t, err, wantErr)
		}

		assert.Equal(t, want, counters)
	}

	f([]byte{msgpcode.FixedArrayLow}, object.GroupsCountersGroup{}, "")
	f([]byte{0x81, 0xA4, 'd', 'o', 'c', 's', 0x01}, object.GroupsCountersGroup{Docs: 1}, "")
	f([]byte("\xc2"), object.GroupsCountersGroup{}, "msgpack: unexpected code=c2 decoding map length")
	f(nil, object.GroupsCountersGroup{}, "EOF")
}

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
		object.GroupsSectionsList{ID: 123, Name: "test"},
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

func TestGroupsSectionsList_DecodeMsgpack(t *testing.T) {
	t.Parallel()

	f := func(data []byte, want object.GroupsSectionsList, wanErr bool) {
		t.Helper()

		var got object.GroupsSectionsList

		err := msgpack.Unmarshal(data, &got)

		if wanErr {
			assert.Error(t, err)
		} else {
			assert.NoError(t, err)
			assert.Equal(t, want, got)
		}
	}

	f(
		[]byte{0x92, 0x7B, 0xA4, 0x74, 0x65, 0x73, 0x74},
		object.GroupsSectionsList{ID: 123, Name: "test"},
		false,
	)

	// Errors:
	f(
		[]byte{0x7B},
		object.GroupsSectionsList{},
		true,
	)
	f(
		[]byte{0x93, 0x7B, 0xA3, 0x31, 0x32, 0x33, 0xA3, 0x31, 0x32, 0x33},
		object.GroupsSectionsList{},
		true,
	)
	f(
		[]byte{0x92, 0xA3, 0x31, 0x32, 0x33, 0xA3, 0x31, 0x32, 0x33},
		object.GroupsSectionsList{},
		true,
	)
	f(
		[]byte{0x92, 0x7B, 0x7B},
		object.GroupsSectionsList{},
		true,
	)
	f(
		nil,
		object.GroupsSectionsList{},
		true,
	)
}
