package api

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestVK_NotificationsGet(t *testing.T) {
	needUserToken(t)

	res, err := vkUser.NotificationsGet(Params{
		"count": 30,
	})
	assert.NoError(t, err)
	// assert.NotEmpty(t, res.Count)
	// assert.NotEmpty(t, res.Items)
	assert.NotEmpty(t, res.LastViewed)
	assert.NotEmpty(t, res.TTL)
}

func TestVK_NotificationsMarkAsViewed(t *testing.T) {
	needUserToken(t)

	res, err := vkUser.NotificationsMarkAsViewed(Params{})
	assert.NoError(t, err)
	assert.NotEmpty(t, res)
}
