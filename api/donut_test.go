package api_test

import (
	"errors"
	"testing"

	"github.com/SevereCloud/vksdk/v2/api"
)

func TestVK_DonutGetFriends(t *testing.T) {
	t.Parallel()

	needUserToken(t)

	_, err := vkUser.DonutGetFriends(api.Params{
		"owner_id": -173151748,
	})
	noError(t, err)
}

func TestVK_DonutGetSubscription(t *testing.T) {
	t.Parallel()

	needUserToken(t)

	_, err := vkUser.DonutGetSubscription(api.Params{
		"owner_id": -173151748,
	})
	if !errors.Is(err, api.ErrNotFound) {
		noError(t, err)
	}
}

func TestVK_DonutGetSubscriptions(t *testing.T) {
	t.Parallel()

	needUserToken(t)

	_, err := vkUser.DonutGetSubscriptions(nil)
	noError(t, err)
}

func TestVK_DonutIsDon(t *testing.T) {
	t.Parallel()

	needUserToken(t)

	_, err := vkUser.DonutIsDon(api.Params{
		"owner_id": -173151748,
	})
	noError(t, err)
}
