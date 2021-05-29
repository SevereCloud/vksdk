package object_test

import (
	"encoding/json"
	"testing"

	"github.com/SevereCloud/vksdk/v2/object"
	"github.com/stretchr/testify/assert"
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

func TestUsersPersonal__UnmarshalJSON(t *testing.T) {
	t.Parallel()

	f := func(data []byte, wantpersonal object.UsersPersonal) {
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
