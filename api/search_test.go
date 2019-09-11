package api

import (
	"testing"
)

func TestVK_SearchGetHints(t *testing.T) {
	if vkUser.AccessToken == "" {
		t.Skip("USER_TOKEN empty")
	}

	_, err := vkUser.SearchGetHints(map[string]string{
		"q":             "VK API",
		"limit":         "20",
		"search_global": "1",
	})
	if err != nil {
		t.Errorf("VK.SearchGetHints() err = %v", err)
	}
}
