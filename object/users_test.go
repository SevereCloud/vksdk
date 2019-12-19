package object_test

import (
	"testing"

	"github.com/SevereCloud/vksdk/object"
)

func TestUsersUser_ToMention(t *testing.T) {
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
