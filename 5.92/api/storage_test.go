package api

import (
	"testing"
)

func TestVK_StorageSet(t *testing.T) {
	if vkUser.AccessToken == "" {
		t.Skip("USER_TOKEN empty")
	}

	t.Run("StorageSet", func(t *testing.T) {
		_, gotVkErr := vkUser.StorageSet(map[string]string{
			"key":   "test",
			"value": "Hello",
		})
		if gotVkErr.Code != 0 {
			t.Errorf("VK.StorageSet() gotVkErr = %v, want %v", gotVkErr, 0)
		}
	})
}

func TestVK_StorageGetKeys(t *testing.T) {
	if vkUser.AccessToken == "" {
		t.Skip("USER_TOKEN empty")
	}

	t.Run("StorageGetKeys", func(t *testing.T) {
		_, gotVkErr := vkUser.StorageGetKeys(map[string]string{})
		if gotVkErr.Code != 0 {
			t.Errorf("VK.StorageGetKeys() gotVkErr = %v, want %v", gotVkErr, 0)
		}
	})
}

func TestVK_StorageGet(t *testing.T) {
	if vkUser.AccessToken == "" {
		t.Skip("USER_TOKEN empty")
	}

	t.Run("StorageGet", func(t *testing.T) {
		_, gotVkErr := vkUser.StorageGet(map[string]string{
			"key": "test",
		})
		if gotVkErr.Code != 0 {
			t.Errorf("VK.StorageGet() gotVkErr = %v, want %v", gotVkErr, 0)
		}
	})
}
