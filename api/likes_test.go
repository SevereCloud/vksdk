package api

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestVK_LikesAdd(t *testing.T) {
	needUserToken(t)

	res, err := vkUser.LikesAdd(Params{
		"type":     "post",
		"owner_id": 1,
		"item_id":  45546,
	})
	assert.NoError(t, err)
	assert.NotEmpty(t, res.Likes)
}

func TestVK_LikesDelete(t *testing.T) {
	needUserToken(t)

	res, err := vkUser.LikesDelete(Params{
		"type":     "post",
		"owner_id": 1,
		"item_id":  45546,
	})
	assert.NoError(t, err)
	assert.NotEmpty(t, res.Likes)
}

func TestVK_LikesGetList(t *testing.T) {
	needUserToken(t)

	params := Params{
		"type":     "post",
		"owner_id": 1,
		"item_id":  45546,
	}

	res, err := vkUser.LikesGetList(params)
	assert.NoError(t, err)
	assert.NotEmpty(t, res.Count)
	assert.NotEmpty(t, res.Items)

	_, err = vkUser.LikesGetListExtended(params)
	assert.NoError(t, err)
	assert.NotEmpty(t, res.Count)
	assert.NotEmpty(t, res.Items)
}

func TestVK_LikesIsLiked(t *testing.T) {
	needUserToken(t)

	_, err := vkUser.LikesIsLiked(Params{
		"type":     "post",
		"owner_id": 1,
		"item_id":  45546,
	})
	// assert.NotEmpty(t, res)
	assert.NoError(t, err)
}
