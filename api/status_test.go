package api_test

import (
	"testing"

	"github.com/SevereCloud/vksdk/v2/api"
)

func TestVK_StatusGet(t *testing.T) {
	t.Parallel()

	needUserToken(t)

	_, err := vkUser.StatusGet(nil)
	noError(t, err)
}

func TestVK_StatusSet(t *testing.T) {
	t.Parallel()

	needUserToken(t)

	_, err := vkUser.StatusSet(api.Params{
		"text": "Hello world",
	})
	noError(t, err)
}
