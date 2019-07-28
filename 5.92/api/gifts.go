package api // import "github.com/SevereCloud/vksdk/5.92/api"

import "github.com/SevereCloud/vksdk/5.92/object"

// GiftsGetResponse struct
type GiftsGetResponse struct {
	Count int                `json:"count"`
	Items []object.GiftsGift `json:"items"`
}

// GiftsGet returns a list of user gifts.
//
// https://vk.com/dev/gifts.get
func (vk *VK) GiftsGet(params map[string]string) (response GiftsGetResponse, vkErr Error) {
	vk.RequestUnmarshal("gifts.get", params, &response, &vkErr)
	return
}

// GiftsGetCatalogResponse struct
type GiftsGetCatalogResponse []struct {
	Name  string             `json:"name"`
	Title string             `json:"title"`
	Items []object.GiftsGift `json:"items"`
}

// GiftsGetCatalog returns catalog.
//
// https://vk.com/dev/gifts.get
func (vk *VK) GiftsGetCatalog(params map[string]string) (response GiftsGetCatalogResponse, vkErr Error) {
	vk.RequestUnmarshal("gifts.getCatalog", params, &response, &vkErr)
	return
}
