package api

import (
	"os"
	"testing"
)

func TestVK_AuthCheckPhone(t *testing.T) {
	needServiceToken(t)
	clientSecret := os.Getenv("CLIENT_SECRET")
	clientID := os.Getenv("CLIENT_ID")
	if clientSecret == "" || clientID == "" {
		t.Skip("need params")
	}

	_, err := vkUser.AuthCheckPhone(map[string]string{
		"phone":         "+79523071234",
		"client_id":     clientID,
		"client_secret": clientSecret,
	})
	if err != nil {
		t.Errorf("VK.AuthCheckPhone() err = %v", err)
	}
}
