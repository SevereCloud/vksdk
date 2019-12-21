package api_test

import (
	"testing"

	"github.com/SevereCloud/vksdk/api"
)

func TestVK_StorageSet(t *testing.T) {
	needUserToken(t)

	_, err := vkUser.StorageSet(api.Params{
		"key":   "test",
		"value": "Hello",
	})
	noError(t, err)
}

func TestVK_StorageGetKeys(t *testing.T) {
	needUserToken(t)

	_, err := vkUser.StorageGetKeys(api.Params{})
	noError(t, err)
}

func TestVK_StorageGet(t *testing.T) {
	needUserToken(t)

	_, err := vkUser.StorageGet(api.Params{
		"key": "test",
	})
	noError(t, err)
}
