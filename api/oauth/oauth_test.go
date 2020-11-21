package oauth_test

import (
	"testing"

	"github.com/SevereCloud/vksdk/v2/api/oauth"
	"github.com/stretchr/testify/assert"
)

func TestCheckScope(t *testing.T) {
	t.Parallel()

	f := func(expected bool, scope int, permissions ...int) {
		t.Helper()

		actual := oauth.CheckScope(scope, permissions...)

		assert.Equal(t, expected, actual)
	}

	f(
		true,
		1026,
		oauth.ScopeUserFriends, oauth.ScopeUserStatus,
	)
	f(
		true,
		1030,
		oauth.ScopeUserFriends, oauth.ScopeUserStatus,
	)
	f(
		false,
		1026,
		oauth.ScopeUserFriends, oauth.ScopeUserStatus, oauth.ScopeUserPhotos,
	)
}
