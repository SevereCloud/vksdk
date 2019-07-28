package api

import (
	"testing"
)

func TestVK_GiftsGet(t *testing.T) {
	if vkUser.AccessToken == "" {
		t.Skip("USER_TOKEN empty")
	}

	_, gotVkErr := vkUser.GiftsGet(map[string]string{
		"user_id": "1",
		"count":   "20",
	})
	if gotVkErr.Code != 0 {
		t.Errorf("VK.GiftsGet() gotVkErr = %v, want %v", gotVkErr, 0)
	}
}

func TestVK_GiftsGetCatalog(t *testing.T) {
	if vkUser.AccessToken == "" {
		t.Skip("USER_TOKEN empty")
	}

	_, gotVkErr := vkUser.GiftsGetCatalog(map[string]string{})
	// NOTE: Access denied: method allowed only for official app
	if gotVkErr.Code != 0 && gotVkErr.Code != 15 {
		t.Errorf("VK.GiftsGetCatalog() gotVkErr = %v, want %v", gotVkErr, 0)
	}
}
