package api // import "github.com/SevereCloud/vksdk/v3/api"

import "github.com/SevereCloud/vksdk/v3/object"

// GiftsGetResponse struct.
type GiftsGetResponse struct {
	Count int                `json:"count"`
	Items []object.GiftsGift `json:"items"`
}

// GiftsGet returns a list of user gifts.
//
// https://dev.vk.com/method/gifts.get
func (vk *VK) GiftsGet(params Params) (response GiftsGetResponse, err error) {
	err = vk.RequestUnmarshal("gifts.get", &response, params)
	return
}

// GiftsGetCatalogResponse struct.
type GiftsGetCatalogResponse []struct {
	Name  string             `json:"name"`
	Title string             `json:"title"`
	Items []object.GiftsGift `json:"items"`
}

// GiftsGetCatalog returns catalog.
//
// https://dev.vk.com/method/gifts.get
func (vk *VK) GiftsGetCatalog(params Params) (response GiftsGetCatalogResponse, err error) {
	err = vk.RequestUnmarshal("gifts.getCatalog", &response, params)
	return
}
