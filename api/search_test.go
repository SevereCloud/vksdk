package api

import (
	"testing"
)

func TestVK_SearchGetHints(t *testing.T) {
	needUserToken(t)

	_, err := vkUser.SearchGetHints(map[string]string{
		"q":             "VK API",
		"limit":         "20",
		"search_global": "1",
	})
	if err != nil {
		t.Errorf("VK.SearchGetHints() err = %v", err)
	}
}
