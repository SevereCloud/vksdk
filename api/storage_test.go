package api_test

import (
	"testing"

	"github.com/SevereCloud/vksdk/api"

	"github.com/stretchr/testify/assert"
)

func TestVK_StorageSet(t *testing.T) {
	needUserToken(t)

	_, err := vkUser.StorageSet(api.Params{
		"key":   "test",
		"value": "Hello",
	})
	assert.NoError(t, err)
}

func TestVK_StorageGetKeys(t *testing.T) {
	needUserToken(t)

	_, err := vkUser.StorageGetKeys(api.Params{})
	assert.NoError(t, err)
}

func TestVK_StorageGet(t *testing.T) {
	needUserToken(t)

	_, err := vkUser.StorageGet(api.Params{
		"key": "test",
	})
	assert.NoError(t, err)
}
