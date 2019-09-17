package api

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestVK_StatusGet(t *testing.T) {
	needUserToken(t)

	_, err := vkUser.StatusGet(map[string]string{})
	assert.NoError(t, err)
}

func TestVK_StatusSet(t *testing.T) {
	needUserToken(t)

	_, err := vkUser.StatusSet(map[string]string{
		"text": "Hello world",
	})
	assert.NoError(t, err)
}
