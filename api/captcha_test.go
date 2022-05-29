package api_test

import (
	"errors"
	"testing"

	"github.com/Derad6709/vksdk/v2/api"
)

func TestVK_CaptchaForce(t *testing.T) {
	t.Parallel()

	needUserToken(t)

	_, err := vkUser.CaptchaForce(nil)

	if !errors.Is(err, api.ErrCaptcha) {
		t.Errorf("VK.CaptchaForce() err=%v, want 14", err)
	}
}
