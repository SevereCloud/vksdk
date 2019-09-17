package api

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestVK_SearchGetHints(t *testing.T) {
	needUserToken(t)

	_, err := vkUser.SearchGetHints(map[string]string{
		"q":             "VK API",
		"limit":         "20",
		"search_global": "1",
	})
	assert.NoError(t, err)
}
