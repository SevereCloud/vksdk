package api_test

import (
	"testing"

	"github.com/SevereCloud/vksdk/api"
	"github.com/stretchr/testify/assert"
)

func TestVK_StorageSet(t *testing.T) {
	t.Parallel()

	needUserToken(t)

	_, err := vkUser.StorageSet(api.Params{
		"key":   "test",
		"value": "Hello",
	})
	noError(t, err)
}

func TestVK_StorageGetKeys(t *testing.T) {
	t.Parallel()

	needUserToken(t)

	_, err := vkUser.StorageGetKeys(nil)
	noError(t, err)
}

func TestVK_StorageGet(t *testing.T) {
	t.Parallel()

	needUserToken(t)

	r, err := vkUser.StorageGet(api.Params{
		"key": "test",
	})
	noError(t, err)

	if assert.NotEmpty(t, r) {
		assert.NotEmpty(t, r[0].Key)
	}
}

func TestStorageGetResponse_ToMap(t *testing.T) {
	t.Parallel()

	f := func(s api.StorageGetResponse, wantMap map[string]string) {
		t.Helper()
		assert.Equal(t, s.ToMap(), wantMap)
	}

	f(
		api.StorageGetResponse{
			{
				Key:   "key",
				Value: "value",
			},
		},
		map[string]string{
			"key": "value",
		},
	)
	f(nil, map[string]string{})
}
