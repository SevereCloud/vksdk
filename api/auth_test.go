package api

import (
	"os"
	"testing"
)

func TestVK_AuthCheckPhone(t *testing.T) {
	clientSecret := os.Getenv("CLIENT_SECRET")
	clientID := os.Getenv("CLIENT_ID")
	if vkService.AccessToken == "" || clientSecret == "" || clientID == "" {
		t.Skip("need params")
	}

	_, gotVkErr := vkUser.AuthCheckPhone(map[string]string{
		"phone":         "+79523071234",
		"client_id":     clientID,
		"client_secret": clientSecret,
	})
	if gotVkErr.Code != 0 {
		t.Errorf("VK.AuthCheckPhone() gotVkErr = %v, want %v", gotVkErr, 0)
	}
}
