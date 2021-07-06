package oauth_test

import (
	"testing"

	"github.com/SevereCloud/vksdk/v2/api/oauth"
	"github.com/stretchr/testify/assert"
)

func BenchmarkCheckScope(b *testing.B) {
	for i := 0; i < b.N; i++ {
		oauth.CheckScope(
			1026,
			oauth.ScopeUserFriends,
			oauth.ScopeUserStatus,
			oauth.ScopeUserPhotos,
		)
	}
}

func TestCheckScope(t *testing.T) {
	f := func(expected bool, scope int, permissions ...int) {
		actual := oauth.CheckScope(scope, permissions...)

		assert.Equal(t, expected, actual)
	}

	f(
		true,
		1026,
		oauth.ScopeUserFriends, oauth.ScopeUserStatus,
	)
	f(
		false,
		1026,
		oauth.ScopeUserFriends, oauth.ScopeUserStatus, oauth.ScopeUserPhotos,
	)
}
