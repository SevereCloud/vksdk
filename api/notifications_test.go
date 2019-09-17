package api

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestVK_NotificationsGet(t *testing.T) {
	needUserToken(t)

	_, err := vkUser.NotificationsGet(map[string]string{
		"count": "30",
	})
	assert.NoError(t, err)
}

func TestVK_NotificationsMarkAsViewed(t *testing.T) {
	needUserToken(t)

	_, err := vkUser.NotificationsMarkAsViewed(map[string]string{})
	assert.NoError(t, err)
}
