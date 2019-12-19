package api_test

import (
	"testing"

	"github.com/SevereCloud/vksdk/api"

	"github.com/stretchr/testify/assert"
)

func TestVK_SearchGetHints(t *testing.T) {
	needUserToken(t)

	_, err := vkUser.SearchGetHints(api.Params{
		"q":             "VK API",
		"limit":         20,
		"search_global": true,
	})
	assert.NoError(t, err)
}
