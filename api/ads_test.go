package api_test

import (
	"testing"

	"github.com/SevereCloud/vksdk/api"
)

func TestVK_AdsGetAccounts(t *testing.T) {
	t.Parallel()

	needUserToken(t)

	_, err := vkUser.AdsGetAccounts(api.Params{})
	noError(t, err)
}
