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

func TestVK_StoreGetProducts(t *testing.T) {
	t.Parallel()

	needUserToken(t)

	products, err := vkUser.StoreGetProducts(api.Params{
		"type":    "stickers",
		"filters": "active",
	})
	noError(t, err)

	assert.NotEmpty(t, products.Count)

	if assert.NotEmpty(t, products.Items) {
		assert.NotEmpty(t, products.Items[0].ID)
		assert.NotEmpty(t, products.Items[0].Type)
		assert.NotEmpty(t, products.Items[0].Purchased)
		assert.NotEmpty(t, products.Items[0].Active)
		assert.NotEmpty(t, products.Items[0].PurchaseDate)
	}
}

func TestVK_StoreGetProductsExtended(t *testing.T) {
	t.Parallel()

	needUserToken(t)

	products, err := vkUser.StoreGetProductsExtended(api.Params{
		"type":    "stickers",
		"filters": "active",
	})
	noError(t, err)

	assert.NotEmpty(t, products.Count)

	if assert.NotEmpty(t, products.Items) {
		assert.NotEmpty(t, products.Items[0].ID)
		assert.NotEmpty(t, products.Items[0].Type)
		assert.NotEmpty(t, products.Items[0].Purchased)
		assert.NotEmpty(t, products.Items[0].Active)
		assert.NotEmpty(t, products.Items[0].PurchaseDate)
		assert.NotEmpty(t, products.Items[0].Title)
		assert.NotEmpty(t, products.Items[0].Stickers)
		assert.NotEmpty(t, products.Items[0].Icon)
		assert.NotEmpty(t, products.Items[0].Previews)
	}
}
