package api // import "github.com/SevereCloud/vksdk/api"

// StoreAddStickersToFavorite add stickers to favorite.
//
// https://vk.com/dev/store.addStickersToFavorite
func (vk *VK) StoreAddStickersToFavorite(params Params) (response int, err error) {
	err = vk.RequestUnmarshal("store.addStickersToFavorite", &response, params)
	return
}
