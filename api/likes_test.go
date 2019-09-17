package api

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestVK_LikesAdd(t *testing.T) {
	needUserToken(t)

	_, err := vkUser.LikesAdd(map[string]string{
		"type":     "post",
		"owner_id": "1",
		"item_id":  "45546",
	})
	assert.NoError(t, err)
}

func TestVK_LikesDelete(t *testing.T) {
	needUserToken(t)

	_, err := vkUser.LikesDelete(map[string]string{
		"type":     "post",
		"owner_id": "1",
		"item_id":  "45546",
	})
	assert.NoError(t, err)
}

func TestVK_LikesGetList(t *testing.T) {
	needUserToken(t)

	params := map[string]string{
		"type":     "post",
		"owner_id": "1",
		"item_id":  "45546",
	}

	_, err := vkUser.LikesGetList(params)
	assert.NoError(t, err)

	_, err = vkUser.LikesGetListExtended(params)
	assert.NoError(t, err)
}

func TestVK_LikesIsLiked(t *testing.T) {
	needUserToken(t)

	_, err := vkUser.LikesIsLiked(map[string]string{
		"type":     "post",
		"owner_id": "1",
		"item_id":  "45546",
	})
	assert.NoError(t, err)
}
