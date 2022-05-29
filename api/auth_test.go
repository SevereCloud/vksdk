package api_test

import (
	"os"
	"testing"

	"github.com/Derad6709/vksdk/v2/api"
)

func TestVK_AuthCheckPhone(t *testing.T) {
	t.Parallel()

	needServiceToken(t)

	clientSecret := os.Getenv("CLIENT_SECRET")
	clientID := os.Getenv("CLIENT_ID")

	if clientSecret == "" || clientID == "" {
		t.Skip("need params")
	}

	_, _ = vkUser.AuthCheckPhone(api.Params{
		"phone":         "+79523071234",
		"client_id":     clientID,
		"client_secret": clientSecret,
	})
}
