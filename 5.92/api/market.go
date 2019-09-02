package api // import "github.com/SevereCloud/vksdk/5.92/api"

import (
	"github.com/SevereCloud/vksdk/5.92/object"
)

// MarketAddResponse struct
type MarketAddResponse struct {
	MarketItemID int `json:"market_item_id"` // Item ID
}

// MarketAdd adds a new item to the market.
//
// https://vk.com/dev/market.add
func (vk *VK) MarketAdd(params map[string]string) (response MarketAddResponse, vkErr Error) {
	vk.RequestUnmarshal("market.add", params, &response, &vkErr)
	return
}

// MarketAddAlbumResponse struct
type MarketAddAlbumResponse struct {
	MarketAlbumID int `json:"market_album_id"` // Album ID
}

// MarketAddAlbum creates new collection of items
//
// https://vk.com/dev/market.addAlbum
func (vk *VK) MarketAddAlbum(params map[string]string) (response MarketAddAlbumResponse, vkErr Error) {
	vk.RequestUnmarshal("market.addAlbum", params, &response, &vkErr)
	return
}

// MarketAddToAlbum adds an item to one or multiple collections.
//
// https://vk.com/dev/market.addToAlbum
func (vk *VK) MarketAddToAlbum(params map[string]string) (response int, vkErr Error) {
	vk.RequestUnmarshal("market.addToAlbum", params, &response, &vkErr)
	return
}

// MarketCreateComment creates a new comment for an item.
//
// https://vk.com/dev/market.createComment
func (vk *VK) MarketCreateComment(params map[string]string) (response int, vkErr Error) {
	vk.RequestUnmarshal("market.createComment", params, &response, &vkErr)
	return
}

// MarketDelete deletes an item.
//
// https://vk.com/dev/market.delete
func (vk *VK) MarketDelete(params map[string]string) (response int, vkErr Error) {
	vk.RequestUnmarshal("market.delete", params, &response, &vkErr)
	return
}

// MarketDeleteAlbum deletes a collection of items.
//
// https://vk.com/dev/market.deleteAlbum
func (vk *VK) MarketDeleteAlbum(params map[string]string) (response int, vkErr Error) {
	vk.RequestUnmarshal("market.deleteAlbum", params, &response, &vkErr)
	return
}

// MarketDeleteComment deletes an item's comment
//
// https://vk.com/dev/market.deleteComment
func (vk *VK) MarketDeleteComment(params map[string]string) (response int, vkErr Error) {
	vk.RequestUnmarshal("market.deleteComment", params, &response, &vkErr)
	return
}

// MarketEdit edits an item.
//
// https://vk.com/dev/market.edit
func (vk *VK) MarketEdit(params map[string]string) (response int, vkErr Error) {
	vk.RequestUnmarshal("market.edit", params, &response, &vkErr)
	return
}

// MarketEditAlbum edits a collection of items
//
// https://vk.com/dev/market.editAlbum
func (vk *VK) MarketEditAlbum(params map[string]string) (response int, vkErr Error) {
	vk.RequestUnmarshal("market.editAlbum", params, &response, &vkErr)
	return
}

// MarketEditComment changes item comment's text
//
// https://vk.com/dev/market.editComment
func (vk *VK) MarketEditComment(params map[string]string) (response int, vkErr Error) {
	vk.RequestUnmarshal("market.editComment", params, &response, &vkErr)
	return
}

// MarketGetResponse struct
type MarketGetResponse struct {
	Count int                       `json:"count"`
	Items []object.MarketMarketItem `json:"items"`
}

// MarketGet returns items list for a community.
//
// https://vk.com/dev/market.get
func (vk *VK) MarketGet(params map[string]string) (response MarketGetResponse, vkErr Error) {
	vk.RequestUnmarshal("market.get", params, &response, &vkErr)
	return
}

// MarketGetAlbumByIDResponse struct
type MarketGetAlbumByIDResponse struct {
	Count int                        `json:"count"`
	Items []object.MarketMarketAlbum `json:"items"`
}

// MarketGetAlbumByID returns items album's data
//
// https://vk.com/dev/market.getAlbumById
func (vk *VK) MarketGetAlbumByID(params map[string]string) (response MarketGetAlbumByIDResponse, vkErr Error) {
	vk.RequestUnmarshal("market.getAlbumById", params, &response, &vkErr)
	return
}

// MarketGetAlbumsResponse struct
type MarketGetAlbumsResponse struct {
	Count int                        `json:"count"`
	Items []object.MarketMarketAlbum `json:"items"`
}

// MarketGetAlbums returns community's collections list.
//
// https://vk.com/dev/market.getAlbums
func (vk *VK) MarketGetAlbums(params map[string]string) (response MarketGetAlbumsResponse, vkErr Error) {
	vk.RequestUnmarshal("market.getAlbums", params, &response, &vkErr)
	return
}

// MarketGetByIDResponse struct
type MarketGetByIDResponse struct {
	Count int                       `json:"count"`
	Items []object.MarketMarketItem `json:"items"`
}

// MarketGetByID returns information about market items by their iDs.
//
// https://vk.com/dev/market.getById
func (vk *VK) MarketGetByID(params map[string]string) (response MarketGetByIDResponse, vkErr Error) {
	vk.RequestUnmarshal("market.getById", params, &response, &vkErr)
	return
}

// MarketGetCategoriesResponse struct
type MarketGetCategoriesResponse struct {
	Count int                           `json:"count"`
	Items []object.MarketMarketCategory `json:"items"`
}

// MarketGetCategories returns a list of market categories.
//
// https://vk.com/dev/market.getCategories
func (vk *VK) MarketGetCategories(params map[string]string) (response MarketGetCategoriesResponse, vkErr Error) {
	vk.RequestUnmarshal("market.getCategories", params, &response, &vkErr)
	return
}

// MarketGetCommentsResponse struct
type MarketGetCommentsResponse struct {
	Count int                      `json:"count"`
	Items []object.WallWallComment `json:"items"`
}

// MarketGetComments returns comments list for an item.
//
// extended=0
//
// https://vk.com/dev/market.getComments
func (vk *VK) MarketGetComments(params map[string]string) (response MarketGetCommentsResponse, vkErr Error) {
	params["extended"] = "0"
	vk.RequestUnmarshal("market.getComments", params, &response, &vkErr)
	return
}

// MarketGetCommentsExtendedResponse struct
type MarketGetCommentsExtendedResponse struct {
	Count    int                      `json:"count"`
	Items    []object.WallWallComment `json:"items"`
	Profiles []object.UsersUser       `json:"profiles"`
	Groups   []object.GroupsGroup     `json:"groups"`
}

// MarketGetCommentsExtended returns comments list for an item.
//
// extended=1
//
// https://vk.com/dev/market.getComments
func (vk *VK) MarketGetCommentsExtended(params map[string]string) (response MarketGetCommentsExtendedResponse, vkErr Error) {
	params["extended"] = "1"
	vk.RequestUnmarshal("market.getComments", params, &response, &vkErr)
	return
}

// MarketRemoveFromAlbum removes an item from one or multiple collections.
//
// https://vk.com/dev/market.removeFromAlbum
func (vk *VK) MarketRemoveFromAlbum(params map[string]string) (response int, vkErr Error) {
	vk.RequestUnmarshal("market.removeFromAlbum", params, &response, &vkErr)
	return
}

// MarketReorderAlbums reorders the collections list.
//
// https://vk.com/dev/market.reorderAlbums
func (vk *VK) MarketReorderAlbums(params map[string]string) (response int, vkErr Error) {
	vk.RequestUnmarshal("market.reorderAlbums", params, &response, &vkErr)
	return
}

// MarketReorderItems changes item place in a collection.
//
// https://vk.com/dev/market.reorderItems
func (vk *VK) MarketReorderItems(params map[string]string) (response int, vkErr Error) {
	vk.RequestUnmarshal("market.reorderItems", params, &response, &vkErr)
	return
}

// MarketReport sends a complaint to the item.
//
// https://vk.com/dev/market.report
func (vk *VK) MarketReport(params map[string]string) (response int, vkErr Error) {
	vk.RequestUnmarshal("market.report", params, &response, &vkErr)
	return
}

// MarketReportComment sends a complaint to the item's comment.
//
// https://vk.com/dev/market.reportComment
func (vk *VK) MarketReportComment(params map[string]string) (response int, vkErr Error) {
	vk.RequestUnmarshal("market.reportComment", params, &response, &vkErr)
	return
}

// MarketRestore restores recently deleted item
//
// https://vk.com/dev/market.restore
func (vk *VK) MarketRestore(params map[string]string) (response int, vkErr Error) {
	vk.RequestUnmarshal("market.restore", params, &response, &vkErr)
	return
}

// MarketRestoreComment restores a recently deleted comment
//
// https://vk.com/dev/market.restoreComment
func (vk *VK) MarketRestoreComment(params map[string]string) (response int, vkErr Error) {
	vk.RequestUnmarshal("market.restoreComment", params, &response, &vkErr)
	return
}

// MarketSearchResponse struct
type MarketSearchResponse struct {
	Count int                       `json:"count"`
	Items []object.MarketMarketItem `json:"items"`
}

// MarketSearch searches market items in a community's catalog
//
// https://vk.com/dev/market.search
func (vk *VK) MarketSearch(params map[string]string) (response MarketSearchResponse, vkErr Error) {
	vk.RequestUnmarshal("market.search", params, &response, &vkErr)
	return
}
