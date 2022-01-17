package object_test

import (
	"encoding/json"
	"testing"

	"github.com/SevereCloud/vksdk/v2/object"
	"github.com/stretchr/testify/assert"
	"github.com/vmihailenco/msgpack/v5"
	"github.com/vmihailenco/msgpack/v5/msgpcode"
)

func TestUsersUser_ToMention(t *testing.T) {
	t.Parallel()

	f := func(user object.UsersUser, want string) {
		if got := user.ToMention(); got != want {
			t.Errorf("UsersUser.ToMention() = %v, want %v", got, want)
		}
	}

	f(
		object.UsersUser{ID: 1, FirstName: "Ivan", LastName: "Ivanov"},
		"[id1|Ivan Ivanov]",
	)
}

func TestUsersUserMin_ToMention(t *testing.T) {
	t.Parallel()

	f := func(user object.UsersUserMin, want string) {
		if got := user.ToMention(); got != want {
			t.Errorf("UsersUserMin.ToMention() = %v, want %v", got, want)
		}
	}

	f(
		object.UsersUserMin{ID: 1, FirstName: "Ivan", LastName: "Ivanov"},
		"[id1|Ivan Ivanov]",
	)
}

func TestUsersPersonal_UnmarshalJSON(t *testing.T) {
	t.Parallel()

	f := func(data []byte, wantpersonal object.UsersPersonal) {
		t.Helper()

		var personal object.UsersPersonal

		err := json.Unmarshal(data, &personal)
		assert.NoError(t, err)
		assert.Equal(t, wantpersonal, personal)
	}

	f([]byte("[]"), object.UsersPersonal{})
	f([]byte(`{"alcohol":1}`), object.UsersPersonal{Alcohol: 1})

	var personal object.UsersPersonal
	err := json.Unmarshal([]byte("0"), &personal)
	assert.Error(t, err)
}

func TestUsersPersonal_DecodeMsgpack(t *testing.T) {
	t.Parallel()

	f := func(data []byte, wantpersonal object.UsersPersonal, wantErr string) {
		t.Helper()

		var personal object.UsersPersonal

		err := msgpack.Unmarshal(data, &personal)
		if err != nil || wantErr != "" {
			assert.EqualError(t, err, wantErr)
		}

		assert.Equal(t, wantpersonal, personal)
	}

	f([]byte{msgpcode.FixedArrayLow}, object.UsersPersonal{}, "")
	f([]byte{
		0x81, 0xA7, 0x61, 0x6C, 0x63, 0x6F, 0x68, 0x6F, 0x6C, 0x01,
	}, object.UsersPersonal{Alcohol: 1}, "")
	f([]byte("\xc2"), object.UsersPersonal{}, "msgpack: unexpected code=c2 decoding map length")
	f(nil, object.UsersPersonal{}, "EOF")
}
