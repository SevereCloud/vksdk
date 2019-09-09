package api

import (
	"testing"
)

func TestVK_NotificationsGet(t *testing.T) {
	if vkUser.AccessToken == "" {
		t.Skip("USER_TOKEN empty")
	}

	_, gotVkErr := vkUser.NotificationsGet(map[string]string{
		"count": "30",
	})
	if gotVkErr.Code != 0 {
		t.Errorf("VK.NotificationsGet() gotVkErr = %v, want %v", gotVkErr, 0)
	}
}

func TestVK_NotificationsMarkAsViewed(t *testing.T) {
	if vkUser.AccessToken == "" {
		t.Skip("USER_TOKEN empty")
	}

	_, gotVkErr := vkUser.NotificationsMarkAsViewed(map[string]string{})
	if gotVkErr.Code != 0 {
		t.Errorf("VK.NotificationsMarkAsViewed() gotVkErr = %v, want %v", gotVkErr, 0)
	}
}
