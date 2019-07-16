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
