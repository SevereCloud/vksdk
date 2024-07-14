package api_test

import (
	"testing"

	"github.com/SevereCloud/vksdk/v3/api"

	"github.com/stretchr/testify/assert"
)

func TestVK_LikesAdd(t *testing.T) {
	t.Parallel()

	needUserToken(t)

	res, err := vkUser.LikesAdd(api.Params{
		"type":     "post",
		"owner_id": 1,
		"item_id":  45546,
	})
	noError(t, err)
	assert.NotEmpty(t, res.Likes)

	resDel, err := vkUser.LikesDelete(api.Params{
		"type":     "post",
		"owner_id": 1,
		"item_id":  45546,
	})
	noError(t, err)
	assert.NotEmpty(t, resDel.Likes)
}

func TestVK_LikesGetList(t *testing.T) {
	t.Parallel()

	needUserToken(t)

	params := api.Params{
		"type":     "post",
		"owner_id": 1,
		"item_id":  45546,
	}

	res, err := vkUser.LikesGetList(params)
	noError(t, err)
	assert.NotEmpty(t, res.Count)
	assert.NotEmpty(t, res.Items)

	_, err = vkUser.LikesGetListExtended(params)
	noError(t, err)
	assert.NotEmpty(t, res.Count)
	assert.NotEmpty(t, res.Items)
}

func TestVK_LikesIsLiked(t *testing.T) {
	t.Parallel()

	needUserToken(t)

	_, err := vkUser.LikesIsLiked(api.Params{
		"type":     "post",
		"owner_id": 1,
		"item_id":  45546,
	})
	// assert.NotEmpty(t, res)
	noError(t, err)
}
