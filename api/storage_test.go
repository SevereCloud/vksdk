package api

import (
	"testing"
)

func TestVK_StorageSet(t *testing.T) {
	needUserToken(t)

	_, err := vkUser.StorageSet(map[string]string{
		"key":   "test",
		"value": "Hello",
	})
	if err != nil {
		t.Errorf("VK.StorageSet() err = %v", err)
	}
}

func TestVK_StorageGetKeys(t *testing.T) {
	needUserToken(t)

	_, err := vkUser.StorageGetKeys(map[string]string{})
	if err != nil {
		t.Errorf("VK.StorageGetKeys() err = %v", err)
	}
}

func TestVK_StorageGet(t *testing.T) {
	needUserToken(t)

	_, err := vkUser.StorageGet(map[string]string{
		"key": "test",
	})
	if err != nil {
		t.Errorf("VK.StorageGet() err = %v", err)
	}
}
