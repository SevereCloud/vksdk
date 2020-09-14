package api

import "github.com/SevereCloud/vksdk/object"

// StoreActivateProduct method.
//
// https://vk.com/dev/store.activateProduct
func (vk *VK) StoreActivateProduct(params Params) (response int, err error) {
	err = vk.RequestUnmarshal("store.activateProduct", params, &response)
	return
}

// StoreAddStickersToFavorite method.
//
// https://vk.com/dev/store.addStickersToFavorite
func (vk *VK) StoreAddStickersToFavorite(params Params) (response int, err error) {
	err = vk.RequestUnmarshal("store.addStickersToFavorite", params, &response)
	return
}

// StoreDeactivateProduct method.
//
// https://vk.com/dev/store.deactivateProduct
func (vk *VK) StoreDeactivateProduct(params Params) (response int, err error) {
	err = vk.RequestUnmarshal("store.deactivateProduct", params, &response)
	return
}

// StoreGetFavoriteStickersResponse struct.
type StoreGetFavoriteStickersResponse struct {
	Count int                  `json:"count"`
	Items []object.BaseSticker `json:"items"`
}

// StoreGetFavoriteStickers method.
//
// https://vk.com/dev/store.getFavoriteStickers
func (vk *VK) StoreGetFavoriteStickers(params Params) (response StoreGetFavoriteStickersResponse, err error) {
	err = vk.RequestUnmarshal("store.getFavoriteStickers", params, &response)
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
	err = vk.RequestUnmarshal("store.getProducts", params, &response)

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
	err = vk.RequestUnmarshal("store.getProducts", params, &response)

	return
}

// StoreGetStickersResponse struct.
type StoreGetStickersResponse []object.BaseSticker

// StoreGetStickers method.
//
// https://vk.com/dev/store.getStickers
func (vk *VK) StoreGetStickers(params Params) (response StoreGetStickersResponse, err error) {
	err = vk.RequestUnmarshal("store.getStickers", params, &response)
	return
}

// StoreMarkAsViewed method.
//
// https://vk.com/dev/store.markAsViewed
func (vk *VK) StoreMarkAsViewed(params Params) (response int, err error) {
	err = vk.RequestUnmarshal("store.markAsViewed", params, &response)
	return
}

// StoreRemoveStickersFromFavorite method.
//
// https://vk.com/dev/store.removeStickersFromFavorite
func (vk *VK) StoreRemoveStickersFromFavorite(params Params) (response int, err error) {
	err = vk.RequestUnmarshal("store.removeStickersFromFavorite", params, &response)
	return
}

// StoreReorderProducts method.
//
// https://vk.com/dev/store.reorderProducts
func (vk *VK) StoreReorderProducts(params Params) (response int, err error) {
	err = vk.RequestUnmarshal("store.reorderProducts", params, &response)
	return
}
