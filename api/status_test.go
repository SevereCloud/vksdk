package api

import (
	"testing"
)

func TestVK_StatusGet(t *testing.T) {
	needUserToken(t)

	_, err := vkUser.StatusGet(map[string]string{})
	if err != nil {
		t.Errorf("VK.StatusGet() err = %v", err)
	}
}

func TestVK_StatusSet(t *testing.T) {
	needUserToken(t)

	_, err := vkUser.StatusSet(map[string]string{
		"text": "Hello world",
	})
	if err != nil {
		t.Errorf("VK.StatusSet() err = %v", err)
	}
}
