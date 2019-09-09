package api

import (
	"testing"
)

func TestVK_StorageSet(t *testing.T) {
	if vkUser.AccessToken == "" {
		t.Skip("USER_TOKEN empty")
	}

	_, gotVkErr := vkUser.StorageSet(map[string]string{
		"key":   "test",
		"value": "Hello",
	})
	if gotVkErr.Code != 0 {
		t.Errorf("VK.StorageSet() gotVkErr = %v", gotVkErr)
	}
}

func TestVK_StorageGetKeys(t *testing.T) {
	if vkUser.AccessToken == "" {
		t.Skip("USER_TOKEN empty")
	}

	_, gotVkErr := vkUser.StorageGetKeys(map[string]string{})
	if gotVkErr.Code != 0 {
		t.Errorf("VK.StorageGetKeys() gotVkErr = %v", gotVkErr)
	}
}

func TestVK_StorageGet(t *testing.T) {
	if vkUser.AccessToken == "" {
		t.Skip("USER_TOKEN empty")
	}

	_, gotVkErr := vkUser.StorageGet(map[string]string{
		"key": "test",
	})
	if gotVkErr.Code != 0 {
		t.Errorf("VK.StorageGet() gotVkErr = %v", gotVkErr)
	}
}
