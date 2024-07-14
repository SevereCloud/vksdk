package api_test

import (
	"testing"

	"github.com/SevereCloud/vksdk/v3/api"

	"github.com/stretchr/testify/assert"
)

func TestVK_NotificationsGet(t *testing.T) {
	t.Parallel()

	needUserToken(t)

	res, err := vkUser.NotificationsGet(api.Params{
		"count": 30,
	})
	noError(t, err)
	// assert.NotEmpty(t, res.Count)
	// assert.NotEmpty(t, res.Items)
	// assert.NotEmpty(t, res.LastViewed)
	assert.NotEmpty(t, res.TTL)
}

func TestVK_NotificationsMarkAsViewed(t *testing.T) {
	t.Parallel()

	needUserToken(t)

	res, err := vkUser.NotificationsMarkAsViewed(nil)
	noError(t, err)
	assert.NotEmpty(t, res)
}
