package api

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestVK_StorageSet(t *testing.T) {
	needUserToken(t)

	_, err := vkUser.StorageSet(Params{
		"key":   "test",
		"value": "Hello",
	})
	assert.NoError(t, err)
}

func TestVK_StorageGetKeys(t *testing.T) {
	needUserToken(t)

	_, err := vkUser.StorageGetKeys(Params{})
	assert.NoError(t, err)
}

func TestVK_StorageGet(t *testing.T) {
	needUserToken(t)

	_, err := vkUser.StorageGet(Params{
		"key": "test",
	})
	assert.NoError(t, err)
}
