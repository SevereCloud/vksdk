package api

import (
	"testing"

	"github.com/SevereCloud/vksdk/errors"
)

func TestVK_GiftsGet(t *testing.T) {
	needUserToken(t)

	_, err := vkUser.GiftsGet(map[string]string{
		"user_id": "1",
		"count":   "20",
	})
	if err != nil {
		t.Errorf("VK.GiftsGet() err = %v", err)
	}
}

func TestVK_GiftsGetCatalog(t *testing.T) {
	needUserToken(t)

	_, err := vkUser.GiftsGetCatalog(map[string]string{})
	// NOTE: Access denied: method allowed only for official app
	if err != nil && errors.GetType(err) != errors.Access {
		t.Errorf("VK.GiftsGetCatalog() err = %v", err)
	}
}
