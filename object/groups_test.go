package object_test

import (
	"encoding/json"
	"testing"

	"github.com/SevereCloud/vksdk/v3/object"
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
	f([]byte("\xc2"), object.GroupsCountersGroup{}, "object.GroupsCountersGroup: msgpack: unexpected code=c2 decoding map length")
	f(nil, object.GroupsCountersGroup{}, "object.GroupsCountersGroup: EOF")
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

func TestGroupsLongPollServer_GetURL(t *testing.T) {
	t.Parallel()

	f := func(s object.GroupsLongPollServer, wait int, want string) {
		t.Helper()
		assert.Equal(t, want, s.GetURL(wait))
	}

	f(object.GroupsLongPollServer{
		Key:    "abc",
		Server: "https://vk.ru",
		Ts:     "123",
	}, 25, "https://vk.ru?act=a_check&key=abc&ts=123&wait=25")
	f(object.GroupsLongPollServer{
		Key:    "abc",
		Server: "https://vk.ru",
		Ts:     "123",
	}, 10, "https://vk.ru?act=a_check&key=abc&ts=123&wait=10")
}
