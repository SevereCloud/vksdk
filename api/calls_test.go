package api_test

import (
	"testing"

	"github.com/SevereCloud/vksdk/v2/api"
)

func TestVK_CallsStart(t *testing.T) {
	t.Parallel()

	needUserToken(t)

	startCall, err := vkUser.CallsStart(nil)
	noError(t, err)

	_, err = vkUser.CallsForceFinish(api.Params{
		"call_id": startCall.CallID,
	})
	noError(t, err)
}
