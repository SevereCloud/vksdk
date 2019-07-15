package api

import (
	"testing"
)

func TestVK_StatusGet(t *testing.T) {
	if vkUser.AccessToken == "" {
		t.Skip("USER_TOKEN empty")
	}

	t.Run("StatusGet", func(t *testing.T) {
		_, gotVkErr := vkUser.StatusGet(map[string]string{})
		if gotVkErr.Code != 0 {
			t.Errorf("VK.StatusGet() gotVkErr = %v, want %v", gotVkErr, 0)
		}
	})
}

func TestVK_StatusSet(t *testing.T) {
	if vkUser.AccessToken == "" {
		t.Skip("USER_TOKEN empty")
	}

	t.Run("StatusSet", func(t *testing.T) {
		gotResponse, gotVkErr := vkUser.StatusSet(map[string]string{
			"text": "Hello world",
		})
		if gotResponse != 1 {
			t.Errorf("VK.StatusSet() gotResponse = %v, want %v", gotResponse, 1)
		}
		if gotVkErr.Code != 0 {
			t.Errorf("VK.StatusSet() gotVkErr = %v, want %v", gotVkErr, 0)
		}
	})
}
