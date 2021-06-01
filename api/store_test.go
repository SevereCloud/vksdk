package api_test

import (
	"testing"

	"github.com/SevereCloud/vksdk/v2/api"

	"github.com/stretchr/testify/assert"
)

func TestVK_StoreGetFavoriteStickers(t *testing.T) {
	t.Parallel()

	needUserToken(t)

	_, err := vkUser.StoreAddStickersToFavorite(api.Params{
		"sticker_ids": 9008,
	})
	noError(t, err)

	res, err := vkUser.StoreGetFavoriteStickers(api.Params{})
	noError(t, err)
	assert.NotEmpty(t, res.Count)
	assert.NotEmpty(t, res.Items)

	_, err = vkUser.StoreRemoveStickersFromFavorite(api.Params{
		"sticker_ids": 9008,
	})
	noError(t, err)
}
