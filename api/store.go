package api // import "github.com/SevereCloud/vksdk/v3/api"

import (
	"github.com/SevereCloud/vksdk/v3/object"
)

// StoreAddStickersToFavorite add stickers to favorite.
//
// https://dev.vk.com/method/store.addStickersToFavorite
func (vk *VK) StoreAddStickersToFavorite(params Params) (response int, err error) {
	err = vk.RequestUnmarshal("store.addStickersToFavorite", &response, params)
	return
}

// StoreGetFavoriteStickersResponse struct.
type StoreGetFavoriteStickersResponse struct {
	Count int                  `json:"count"`
	Items []object.BaseSticker `json:"items"`
}

// StoreGetFavoriteStickers return favorite stickers.
//
// https://dev.vk.com/method/store.getFavoriteStickers
func (vk *VK) StoreGetFavoriteStickers(params Params) (response StoreGetFavoriteStickersResponse, err error) {
	err = vk.RequestUnmarshal("store.getFavoriteStickers", &response, params)
	return
}

// StoreRemoveStickersFromFavorite remove stickers from favorite.
//
// https://dev.vk.com/method/store.removeStickersFromFavorite
func (vk *VK) StoreRemoveStickersFromFavorite(params Params) (response int, err error) {
	err = vk.RequestUnmarshal("store.removeStickersFromFavorite", &response, params)
	return
}

// StoreActivateProduct method.
//
// https://vk.com/dev/store.activateProduct
func (vk *VK) StoreActivateProduct(params Params) (response int, err error) {
	err = vk.RequestUnmarshal("store.activateProduct", &response, params)
	return
}

// StoreDeactivateProduct method.
//
// https://vk.com/dev/store.deactivateProduct
func (vk *VK) StoreDeactivateProduct(params Params) (response int, err error) {
	err = vk.RequestUnmarshal("store.deactivateProduct", &response, params)
	return
}

// StoreGetProductsResponse struct.
type StoreGetProductsResponse struct {
	Count int                   `json:"count"`
	Items []object.StoreProduct `json:"items"`
}

// StoreGetProducts method.
//
// extended=0
//
// https://vk.com/dev/store.getProducts
func (vk *VK) StoreGetProducts(params Params) (response StoreGetProductsResponse, err error) {
	params["extended"] = false
	err = vk.RequestUnmarshal("store.getProducts", &response, params)

	return
}

// StoreGetProductsExtendedResponse struct.
type StoreGetProductsExtendedResponse struct {
	Count int                           `json:"count"`
	Items []object.StoreProductExtended `json:"items"`
}

// StoreGetProductsExtended method.
//
// extended=1
//
// https://vk.com/dev/store.getProducts
func (vk *VK) StoreGetProductsExtended(params Params) (response StoreGetProductsExtendedResponse, err error) {
	params["extended"] = true
	err = vk.RequestUnmarshal("store.getProducts", &response, params)

	return
}

// StoreGetStickersResponse struct.
type StoreGetStickersResponse []object.BaseSticker

// StoreGetStickers method.
//
// https://vk.com/dev/store.getStickers
func (vk *VK) StoreGetStickers(params Params) (response StoreGetStickersResponse, err error) {
	err = vk.RequestUnmarshal("store.getStickers", &response, params)
	return
}

// StoreGetStickersKeywordsResponse struct.
type StoreGetStickersKeywordsResponse struct {
	Count      int                              `json:"count"`
	Dictionary []object.StoreStickersDictionary `json:"dictionary"`
}

// StoreGetStickersKeywords method.
//
// https://vk.com/dev/store.getStickersKeywords
func (vk *VK) StoreGetStickersKeywords(params Params) (response StoreGetStickersKeywordsResponse, err error) {
	err = vk.RequestUnmarshal("store.getStickersKeywords", &response, params)
	return
}

// StoreMarkAsViewed method.
//
// https://vk.com/dev/store.markAsViewed
func (vk *VK) StoreMarkAsViewed(params Params) (response int, err error) {
	err = vk.RequestUnmarshal("store.markAsViewed", &response, params)
	return
}

// StoreReorderProducts method.
//
// https://vk.com/dev/store.reorderProducts
func (vk *VK) StoreReorderProducts(params Params) (response int, err error) {
	err = vk.RequestUnmarshal("store.reorderProducts", &response, params)
	return
}
