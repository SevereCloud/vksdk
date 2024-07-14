package api_test

import (
	"errors"
	"testing"

	"github.com/SevereCloud/vksdk/v3/api"
)

func TestVK_CaptchaForce(t *testing.T) {
	t.Parallel()

	needUserToken(t)

	_, err := vkUser.CaptchaForce(nil)

	if !errors.Is(err, api.ErrCaptcha) {
		t.Errorf("VK.CaptchaForce() err=%v, want 14", err)
	}
}
