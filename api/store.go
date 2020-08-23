package api // import "github.com/SevereCloud/vksdk/api"

import (
	"github.com/SevereCloud/vksdk/v2/object"
)

// StoreAddStickersToFavorite add stickers to favorite.
//
// https://vk.com/dev/store.addStickersToFavorite
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
// https://vk.com/dev/store.getFavoriteStickers
func (vk *VK) StoreGetFavoriteStickers(params Params) (response StoreGetFavoriteStickersResponse, err error) {
	err = vk.RequestUnmarshal("store.getFavoriteStickers", &response, params)
	return
}

// StoreRemoveStickersFromFavorite remove stickers from favorite.
//
// https://vk.com/dev/store.removeStickersFromFavorite
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
