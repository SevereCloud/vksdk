package api

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestVK_SearchGetHints(t *testing.T) {
	needUserToken(t)

	_, err := vkUser.SearchGetHints(Params{
		"q":             "VK API",
		"limit":         20,
		"search_global": true,
	})
	assert.NoError(t, err)
}
