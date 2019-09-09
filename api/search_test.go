package api

import (
	"testing"
)

func TestVK_SearchGetHints(t *testing.T) {
	if vkUser.AccessToken == "" {
		t.Skip("USER_TOKEN empty")
	}

	_, gotVkErr := vkUser.SearchGetHints(map[string]string{
		"q":             "VK API",
		"limit":         "20",
		"search_global": "1",
	})
	if gotVkErr.Code != 0 {
		t.Errorf("VK.SearchGetHints() gotVkErr = %v", gotVkErr)
	}
}
