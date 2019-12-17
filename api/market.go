package api // import "github.com/SevereCloud/vksdk/api"

import (
	"github.com/SevereCloud/vksdk/object"
)

// MarketAddResponse struct
type MarketAddResponse struct {
	MarketItemID int `json:"market_item_id"` // Item ID
}

// MarketAdd adds a new item to the market.
//
// https://vk.com/dev/market.add
func (vk *VK) MarketAdd(params Params) (response MarketAddResponse, err error) {
	err = vk.RequestUnmarshal("market.add", params, &response)
	return
}

// MarketAddAlbumResponse struct
type MarketAddAlbumResponse struct {
	MarketAlbumID int `json:"market_album_id"` // Album ID
}

// MarketAddAlbum creates new collection of items
//
// https://vk.com/dev/market.addAlbum
func (vk *VK) MarketAddAlbum(params Params) (response MarketAddAlbumResponse, err error) {
	err = vk.RequestUnmarshal("market.addAlbum", params, &response)
	return
}

// MarketAddToAlbum adds an item to one or multiple collections.
//
// https://vk.com/dev/market.addToAlbum
func (vk *VK) MarketAddToAlbum(params Params) (response int, err error) {
	err = vk.RequestUnmarshal("market.addToAlbum", params, &response)
	return
}

// MarketCreateComment creates a new comment for an item.
//
// https://vk.com/dev/market.createComment
func (vk *VK) MarketCreateComment(params Params) (response int, err error) {
	err = vk.RequestUnmarshal("market.createComment", params, &response)
	return
}

// MarketDelete deletes an item.
//
// https://vk.com/dev/market.delete
func (vk *VK) MarketDelete(params Params) (response int, err error) {
	err = vk.RequestUnmarshal("market.delete", params, &response)
	return
}

// MarketDeleteAlbum deletes a collection of items.
//
// https://vk.com/dev/market.deleteAlbum
func (vk *VK) MarketDeleteAlbum(params Params) (response int, err error) {
	err = vk.RequestUnmarshal("market.deleteAlbum", params, &response)
	return
}

// MarketDeleteComment deletes an item's comment
//
// https://vk.com/dev/market.deleteComment
func (vk *VK) MarketDeleteComment(params Params) (response int, err error) {
	err = vk.RequestUnmarshal("market.deleteComment", params, &response)
	return
}

// MarketEdit edits an item.
//
// https://vk.com/dev/market.edit
func (vk *VK) MarketEdit(params Params) (response int, err error) {
	err = vk.RequestUnmarshal("market.edit", params, &response)
	return
}

// MarketEditAlbum edits a collection of items
//
// https://vk.com/dev/market.editAlbum
func (vk *VK) MarketEditAlbum(params Params) (response int, err error) {
	err = vk.RequestUnmarshal("market.editAlbum", params, &response)
	return
}

// MarketEditComment changes item comment's text
//
// https://vk.com/dev/market.editComment
func (vk *VK) MarketEditComment(params Params) (response int, err error) {
	err = vk.RequestUnmarshal("market.editComment", params, &response)
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
func (vk *VK) MarketGet(params Params) (response MarketGetResponse, err error) {
	err = vk.RequestUnmarshal("market.get", params, &response)
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
func (vk *VK) MarketGetAlbumByID(params Params) (response MarketGetAlbumByIDResponse, err error) {
	err = vk.RequestUnmarshal("market.getAlbumById", params, &response)
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
func (vk *VK) MarketGetAlbums(params Params) (response MarketGetAlbumsResponse, err error) {
	err = vk.RequestUnmarshal("market.getAlbums", params, &response)
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
func (vk *VK) MarketGetByID(params Params) (response MarketGetByIDResponse, err error) {
	err = vk.RequestUnmarshal("market.getById", params, &response)
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
func (vk *VK) MarketGetCategories(params Params) (response MarketGetCategoriesResponse, err error) {
	err = vk.RequestUnmarshal("market.getCategories", params, &response)
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
func (vk *VK) MarketGetComments(params Params) (response MarketGetCommentsResponse, err error) {
	params["extended"] = "0"
	err = vk.RequestUnmarshal("market.getComments", params, &response)

	return
}

// MarketGetCommentsExtendedResponse struct
type MarketGetCommentsExtendedResponse struct {
	Count int                      `json:"count"`
	Items []object.WallWallComment `json:"items"`
	object.ExtendedResponse
}

// MarketGetCommentsExtended returns comments list for an item.
//
// extended=1
//
// https://vk.com/dev/market.getComments
func (vk *VK) MarketGetCommentsExtended(params Params) (response MarketGetCommentsExtendedResponse, err error) {
	params["extended"] = true
	err = vk.RequestUnmarshal("market.getComments", params, &response)

	return
}

// MarketRemoveFromAlbum removes an item from one or multiple collections.
//
// https://vk.com/dev/market.removeFromAlbum
func (vk *VK) MarketRemoveFromAlbum(params Params) (response int, err error) {
	err = vk.RequestUnmarshal("market.removeFromAlbum", params, &response)
	return
}

// MarketReorderAlbums reorders the collections list.
//
// https://vk.com/dev/market.reorderAlbums
func (vk *VK) MarketReorderAlbums(params Params) (response int, err error) {
	err = vk.RequestUnmarshal("market.reorderAlbums", params, &response)
	return
}

// MarketReorderItems changes item place in a collection.
//
// https://vk.com/dev/market.reorderItems
func (vk *VK) MarketReorderItems(params Params) (response int, err error) {
	err = vk.RequestUnmarshal("market.reorderItems", params, &response)
	return
}

// MarketReport sends a complaint to the item.
//
// https://vk.com/dev/market.report
func (vk *VK) MarketReport(params Params) (response int, err error) {
	err = vk.RequestUnmarshal("market.report", params, &response)
	return
}

// MarketReportComment sends a complaint to the item's comment.
//
// https://vk.com/dev/market.reportComment
func (vk *VK) MarketReportComment(params Params) (response int, err error) {
	err = vk.RequestUnmarshal("market.reportComment", params, &response)
	return
}

// MarketRestore restores recently deleted item
//
// https://vk.com/dev/market.restore
func (vk *VK) MarketRestore(params Params) (response int, err error) {
	err = vk.RequestUnmarshal("market.restore", params, &response)
	return
}

// MarketRestoreComment restores a recently deleted comment
//
// https://vk.com/dev/market.restoreComment
func (vk *VK) MarketRestoreComment(params Params) (response int, err error) {
	err = vk.RequestUnmarshal("market.restoreComment", params, &response)
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
func (vk *VK) MarketSearch(params Params) (response MarketSearchResponse, err error) {
	err = vk.RequestUnmarshal("market.search", params, &response)
	return
}
