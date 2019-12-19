package api_test

import (
	"testing"

	"github.com/SevereCloud/vksdk/api"

	"github.com/stretchr/testify/assert"
)

func TestVK_StatusGet(t *testing.T) {
	needUserToken(t)

	_, err := vkUser.StatusGet(api.Params{})
	assert.NoError(t, err)
}

func TestVK_StatusSet(t *testing.T) {
	needUserToken(t)

	_, err := vkUser.StatusSet(api.Params{
		"text": "Hello world",
	})
	assert.NoError(t, err)
}
