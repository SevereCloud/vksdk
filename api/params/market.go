package params // import "github.com/SevereCloud/vksdk/api/params"

import (
	"github.com/SevereCloud/vksdk/api"
)

// MarketAddBulder builder
//
// Ads a new item to the market.
//
// https://vk.com/dev/market.add
type MarketAddBulder struct {
	api.Params
}

// NewMarketAddBulder func
func NewMarketAddBulder() *MarketAddBulder {
	return &MarketAddBulder{api.Params{}}
}

// OwnerID ID of an item owner community.
func (b *MarketAddBulder) OwnerID(v int) {
	b.Params["owner_id"] = v
}

// Name Item name.
func (b *MarketAddBulder) Name(v string) {
	b.Params["name"] = v
}

// Description Item description.
func (b *MarketAddBulder) Description(v string) {
	b.Params["description"] = v
}

// CategoryID Item category ID.
func (b *MarketAddBulder) CategoryID(v int) {
	b.Params["category_id"] = v
}

// Price Item price.
func (b *MarketAddBulder) Price(v float64) {
	b.Params["price"] = v
}

// OldPrice parameter
func (b *MarketAddBulder) OldPrice(v float64) {
	b.Params["old_price"] = v
}

// Deleted Item status ('1' — deleted, '0' — not deleted).
func (b *MarketAddBulder) Deleted(v bool) {
	b.Params["deleted"] = v
}

// MainPhotoID Cover photo ID.
func (b *MarketAddBulder) MainPhotoID(v int) {
	b.Params["main_photo_id"] = v
}

// PhotoIDs IDs of additional photos.
func (b *MarketAddBulder) PhotoIDs(v []int) {
	b.Params["photo_ids"] = v
}

// URL Url for button in market item.
func (b *MarketAddBulder) URL(v string) {
	b.Params["url"] = v
}

// MarketAddAlbumBulder builder
//
// Creates new collection of items
//
// https://vk.com/dev/market.addAlbum
type MarketAddAlbumBulder struct {
	api.Params
}

// NewMarketAddAlbumBulder func
func NewMarketAddAlbumBulder() *MarketAddAlbumBulder {
	return &MarketAddAlbumBulder{api.Params{}}
}

// OwnerID ID of an item owner community.
func (b *MarketAddAlbumBulder) OwnerID(v int) {
	b.Params["owner_id"] = v
}

// Title Collection title.
func (b *MarketAddAlbumBulder) Title(v string) {
	b.Params["title"] = v
}

// PhotoID Cover photo ID.
func (b *MarketAddAlbumBulder) PhotoID(v int) {
	b.Params["photo_id"] = v
}

// MainAlbum Set as main ('1' – set, '0' – no).
func (b *MarketAddAlbumBulder) MainAlbum(v bool) {
	b.Params["main_album"] = v
}

// MarketAddToAlbumBulder builder
//
// Adds an item to one or multiple collections.
//
// https://vk.com/dev/market.addToAlbum
type MarketAddToAlbumBulder struct {
	api.Params
}

// NewMarketAddToAlbumBulder func
func NewMarketAddToAlbumBulder() *MarketAddToAlbumBulder {
	return &MarketAddToAlbumBulder{api.Params{}}
}

// OwnerID ID of an item owner community.
func (b *MarketAddToAlbumBulder) OwnerID(v int) {
	b.Params["owner_id"] = v
}

// ItemID Item ID.
func (b *MarketAddToAlbumBulder) ItemID(v int) {
	b.Params["item_id"] = v
}

// AlbumIDs Collections IDs to add item to.
func (b *MarketAddToAlbumBulder) AlbumIDs(v []int) {
	b.Params["album_ids"] = v
}

// MarketCreateCommentBulder builder
//
// Creates a new comment for an item.
//
// https://vk.com/dev/market.createComment
type MarketCreateCommentBulder struct {
	api.Params
}

// NewMarketCreateCommentBulder func
func NewMarketCreateCommentBulder() *MarketCreateCommentBulder {
	return &MarketCreateCommentBulder{api.Params{}}
}

// OwnerID ID of an item owner community.
func (b *MarketCreateCommentBulder) OwnerID(v int) {
	b.Params["owner_id"] = v
}

// ItemID Item ID.
func (b *MarketCreateCommentBulder) ItemID(v int) {
	b.Params["item_id"] = v
}

// Message Comment text (required if 'attachments' parameter is not specified)
func (b *MarketCreateCommentBulder) Message(v string) {
	b.Params["message"] = v
}

// Attachments Comma-separated list of objects attached to a comment. The field is submitted the following way: , "'<owner_id>_<media_id>,<owner_id>_<media_id>'", , '' - media attachment type: "'photo' - photo, 'video' - video, 'audio' - audio, 'doc' - document", , '<owner_id>' - media owner id, '<media_id>' - media attachment id, , For example: "photo100172_166443618,photo66748_265827614",
func (b *MarketCreateCommentBulder) Attachments(v []string) {
	b.Params["attachments"] = v
}

// FromGroup '1' - comment will be published on behalf of a community, '0' - on behalf of a user (by default).
func (b *MarketCreateCommentBulder) FromGroup(v bool) {
	b.Params["from_group"] = v
}

// ReplyToComment ID of a comment to reply with current comment to.
func (b *MarketCreateCommentBulder) ReplyToComment(v int) {
	b.Params["reply_to_comment"] = v
}

// StickerID Sticker ID.
func (b *MarketCreateCommentBulder) StickerID(v int) {
	b.Params["sticker_id"] = v
}

// GUID Random value to avoid resending one comment.
func (b *MarketCreateCommentBulder) GUID(v string) {
	b.Params["guid"] = v
}

// MarketDeleteBulder builder
//
// Deletes an item.
//
// https://vk.com/dev/market.delete
type MarketDeleteBulder struct {
	api.Params
}

// NewMarketDeleteBulder func
func NewMarketDeleteBulder() *MarketDeleteBulder {
	return &MarketDeleteBulder{api.Params{}}
}

// OwnerID ID of an item owner community.
func (b *MarketDeleteBulder) OwnerID(v int) {
	b.Params["owner_id"] = v
}

// ItemID Item ID.
func (b *MarketDeleteBulder) ItemID(v int) {
	b.Params["item_id"] = v
}

// MarketDeleteAlbumBulder builder
//
// Deletes a collection of items.
//
// https://vk.com/dev/market.deleteAlbum
type MarketDeleteAlbumBulder struct {
	api.Params
}

// NewMarketDeleteAlbumBulder func
func NewMarketDeleteAlbumBulder() *MarketDeleteAlbumBulder {
	return &MarketDeleteAlbumBulder{api.Params{}}
}

// OwnerID ID of an collection owner community.
func (b *MarketDeleteAlbumBulder) OwnerID(v int) {
	b.Params["owner_id"] = v
}

// AlbumID Collection ID.
func (b *MarketDeleteAlbumBulder) AlbumID(v int) {
	b.Params["album_id"] = v
}

// MarketDeleteCommentBulder builder
//
// Deletes an item's comment
//
// https://vk.com/dev/market.deleteComment
type MarketDeleteCommentBulder struct {
	api.Params
}

// NewMarketDeleteCommentBulder func
func NewMarketDeleteCommentBulder() *MarketDeleteCommentBulder {
	return &MarketDeleteCommentBulder{api.Params{}}
}

// OwnerID identifier of an item owner community, "Note that community id in the 'owner_id' parameter should be negative number. For example 'owner_id'=-1 matches the [vk.com/apiclub|VK API] community "
func (b *MarketDeleteCommentBulder) OwnerID(v int) {
	b.Params["owner_id"] = v
}

// CommentID comment id
func (b *MarketDeleteCommentBulder) CommentID(v int) {
	b.Params["comment_id"] = v
}

// MarketEditBulder builder
//
// Edits an item.
//
// https://vk.com/dev/market.edit
type MarketEditBulder struct {
	api.Params
}

// NewMarketEditBulder func
func NewMarketEditBulder() *MarketEditBulder {
	return &MarketEditBulder{api.Params{}}
}

// OwnerID ID of an item owner community.
func (b *MarketEditBulder) OwnerID(v int) {
	b.Params["owner_id"] = v
}

// ItemID Item ID.
func (b *MarketEditBulder) ItemID(v int) {
	b.Params["item_id"] = v
}

// Name Item name.
func (b *MarketEditBulder) Name(v string) {
	b.Params["name"] = v
}

// Description Item description.
func (b *MarketEditBulder) Description(v string) {
	b.Params["description"] = v
}

// CategoryID Item category ID.
func (b *MarketEditBulder) CategoryID(v int) {
	b.Params["category_id"] = v
}

// Price Item price.
func (b *MarketEditBulder) Price(v float64) {
	b.Params["price"] = v
}

// Deleted Item status ('1' — deleted, '0' — not deleted).
func (b *MarketEditBulder) Deleted(v bool) {
	b.Params["deleted"] = v
}

// MainPhotoID Cover photo ID.
func (b *MarketEditBulder) MainPhotoID(v int) {
	b.Params["main_photo_id"] = v
}

// PhotoIDs IDs of additional photos.
func (b *MarketEditBulder) PhotoIDs(v []int) {
	b.Params["photo_ids"] = v
}

// URL Url for button in market item.
func (b *MarketEditBulder) URL(v string) {
	b.Params["url"] = v
}

// MarketEditAlbumBulder builder
//
// Edits a collection of items
//
// https://vk.com/dev/market.editAlbum
type MarketEditAlbumBulder struct {
	api.Params
}

// NewMarketEditAlbumBulder func
func NewMarketEditAlbumBulder() *MarketEditAlbumBulder {
	return &MarketEditAlbumBulder{api.Params{}}
}

// OwnerID ID of an collection owner community.
func (b *MarketEditAlbumBulder) OwnerID(v int) {
	b.Params["owner_id"] = v
}

// AlbumID Collection ID.
func (b *MarketEditAlbumBulder) AlbumID(v int) {
	b.Params["album_id"] = v
}

// Title Collection title.
func (b *MarketEditAlbumBulder) Title(v string) {
	b.Params["title"] = v
}

// PhotoID Cover photo id
func (b *MarketEditAlbumBulder) PhotoID(v int) {
	b.Params["photo_id"] = v
}

// MainAlbum Set as main ('1' – set, '0' – no).
func (b *MarketEditAlbumBulder) MainAlbum(v bool) {
	b.Params["main_album"] = v
}

// MarketEditCommentBulder builder
//
// Chages item comment's text
//
// https://vk.com/dev/market.editComment
type MarketEditCommentBulder struct {
	api.Params
}

// NewMarketEditCommentBulder func
func NewMarketEditCommentBulder() *MarketEditCommentBulder {
	return &MarketEditCommentBulder{api.Params{}}
}

// OwnerID ID of an item owner community.
func (b *MarketEditCommentBulder) OwnerID(v int) {
	b.Params["owner_id"] = v
}

// CommentID Comment ID.
func (b *MarketEditCommentBulder) CommentID(v int) {
	b.Params["comment_id"] = v
}

// Message New comment text (required if 'attachments' are not specified), , 2048 symbols maximum.
func (b *MarketEditCommentBulder) Message(v string) {
	b.Params["message"] = v
}

// Attachments Comma-separated list of objects attached to a comment. The field is submitted the following way: , "'<owner_id>_<media_id>,<owner_id>_<media_id>'", , '' - media attachment type: "'photo' - photo, 'video' - video, 'audio' - audio, 'doc' - document", , '<owner_id>' - media owner id, '<media_id>' - media attachment id, , For example: "photo100172_166443618,photo66748_265827614",
func (b *MarketEditCommentBulder) Attachments(v []string) {
	b.Params["attachments"] = v
}

// MarketGetBulder builder
//
// Returns items list for a community.
//
// https://vk.com/dev/market.get
type MarketGetBulder struct {
	api.Params
}

// NewMarketGetBulder func
func NewMarketGetBulder() *MarketGetBulder {
	return &MarketGetBulder{api.Params{}}
}

// OwnerID ID of an item owner community, "Note that community id in the 'owner_id' parameter should be negative number. For example 'owner_id'=-1 matches the [vk.com/apiclub|VK API] community "
func (b *MarketGetBulder) OwnerID(v int) {
	b.Params["owner_id"] = v
}

// AlbumID parameter
func (b *MarketGetBulder) AlbumID(v int) {
	b.Params["album_id"] = v
}

// Count Number of items to return.
func (b *MarketGetBulder) Count(v int) {
	b.Params["count"] = v
}

// Offset Offset needed to return a specific subset of results.
func (b *MarketGetBulder) Offset(v int) {
	b.Params["offset"] = v
}

// Extended '1' – method will return additional fields: 'likes, can_comment, car_repost, photos'. These parameters are not returned by default.
func (b *MarketGetBulder) Extended(v bool) {
	b.Params["extended"] = v
}

// MarketGetAlbumByIDBulder builder
//
// Returns items album's data
//
// https://vk.com/dev/market.getAlbumById
type MarketGetAlbumByIDBulder struct {
	api.Params
}

// NewMarketGetAlbumByIDBulder func
func NewMarketGetAlbumByIDBulder() *MarketGetAlbumByIDBulder {
	return &MarketGetAlbumByIDBulder{api.Params{}}
}

// OwnerID identifier of an album owner community, "Note that community id in the 'owner_id' parameter should be negative number. For example 'owner_id'=-1 matches the [vk.com/apiclub|VK API] community "
func (b *MarketGetAlbumByIDBulder) OwnerID(v int) {
	b.Params["owner_id"] = v
}

// AlbumIDs collections identifiers to obtain data from
func (b *MarketGetAlbumByIDBulder) AlbumIDs(v []int) {
	b.Params["album_ids"] = v
}

// MarketGetAlbumsBulder builder
//
// Returns community's collections list.
//
// https://vk.com/dev/market.getAlbums
type MarketGetAlbumsBulder struct {
	api.Params
}

// NewMarketGetAlbumsBulder func
func NewMarketGetAlbumsBulder() *MarketGetAlbumsBulder {
	return &MarketGetAlbumsBulder{api.Params{}}
}

// OwnerID ID of an items owner community.
func (b *MarketGetAlbumsBulder) OwnerID(v int) {
	b.Params["owner_id"] = v
}

// Offset Offset needed to return a specific subset of results.
func (b *MarketGetAlbumsBulder) Offset(v int) {
	b.Params["offset"] = v
}

// Count Number of items to return.
func (b *MarketGetAlbumsBulder) Count(v int) {
	b.Params["count"] = v
}

// MarketGetByIDBulder builder
//
// Returns information about market items by their ids.
//
// https://vk.com/dev/market.getById
type MarketGetByIDBulder struct {
	api.Params
}

// NewMarketGetByIDBulder func
func NewMarketGetByIDBulder() *MarketGetByIDBulder {
	return &MarketGetByIDBulder{api.Params{}}
}

// ItemIDs Comma-separated ids list: {user id}_{item id}. If an item belongs to a community -{community id} is used. " 'Videos' value example: , '-4363_136089719,13245770_137352259'"
func (b *MarketGetByIDBulder) ItemIDs(v []string) {
	b.Params["item_ids"] = v
}

// Extended '1' – to return additional fields: 'likes, can_comment, car_repost, photos'. By default: '0'.
func (b *MarketGetByIDBulder) Extended(v bool) {
	b.Params["extended"] = v
}

// MarketGetCategoriesBulder builder
//
// Returns a list of market categories.
//
// https://vk.com/dev/market.getCategories
type MarketGetCategoriesBulder struct {
	api.Params
}

// NewMarketGetCategoriesBulder func
func NewMarketGetCategoriesBulder() *MarketGetCategoriesBulder {
	return &MarketGetCategoriesBulder{api.Params{}}
}

// Count Number of results to return.
func (b *MarketGetCategoriesBulder) Count(v int) {
	b.Params["count"] = v
}

// Offset Offset needed to return a specific subset of results.
func (b *MarketGetCategoriesBulder) Offset(v int) {
	b.Params["offset"] = v
}

// MarketGetCommentsBulder builder
//
// Returns comments list for an item.
//
// https://vk.com/dev/market.getComments
type MarketGetCommentsBulder struct {
	api.Params
}

// NewMarketGetCommentsBulder func
func NewMarketGetCommentsBulder() *MarketGetCommentsBulder {
	return &MarketGetCommentsBulder{api.Params{}}
}

// OwnerID ID of an item owner community
func (b *MarketGetCommentsBulder) OwnerID(v int) {
	b.Params["owner_id"] = v
}

// ItemID Item ID.
func (b *MarketGetCommentsBulder) ItemID(v int) {
	b.Params["item_id"] = v
}

// NeedLikes '1' — to return likes info.
func (b *MarketGetCommentsBulder) NeedLikes(v bool) {
	b.Params["need_likes"] = v
}

// StartCommentID ID of a comment to start a list from (details below).
func (b *MarketGetCommentsBulder) StartCommentID(v int) {
	b.Params["start_comment_id"] = v
}

// Offset parameter
func (b *MarketGetCommentsBulder) Offset(v int) {
	b.Params["offset"] = v
}

// Count Number of results to return.
func (b *MarketGetCommentsBulder) Count(v int) {
	b.Params["count"] = v
}

// Sort Sort order ('asc' — from old to new, 'desc' — from new to old)
func (b *MarketGetCommentsBulder) Sort(v string) {
	b.Params["sort"] = v
}

// Extended '1' — comments will be returned as numbered objects, in addition lists of 'profiles' and 'groups' objects will be returned.
func (b *MarketGetCommentsBulder) Extended(v bool) {
	b.Params["extended"] = v
}

// Fields List of additional profile fields to return. See the [vk.com/dev/fields|details]
func (b *MarketGetCommentsBulder) Fields(v []string) {
	b.Params["fields"] = v
}

// MarketRemoveFromAlbumBulder builder
//
// Removes an item from one or multiple collections.
//
// https://vk.com/dev/market.removeFromAlbum
type MarketRemoveFromAlbumBulder struct {
	api.Params
}

// NewMarketRemoveFromAlbumBulder func
func NewMarketRemoveFromAlbumBulder() *MarketRemoveFromAlbumBulder {
	return &MarketRemoveFromAlbumBulder{api.Params{}}
}

// OwnerID ID of an item owner community.
func (b *MarketRemoveFromAlbumBulder) OwnerID(v int) {
	b.Params["owner_id"] = v
}

// ItemID Item ID.
func (b *MarketRemoveFromAlbumBulder) ItemID(v int) {
	b.Params["item_id"] = v
}

// AlbumIDs Collections IDs to remove item from.
func (b *MarketRemoveFromAlbumBulder) AlbumIDs(v []int) {
	b.Params["album_ids"] = v
}

// MarketReorderAlbumsBulder builder
//
// Reorders the collections list.
//
// https://vk.com/dev/market.reorderAlbums
type MarketReorderAlbumsBulder struct {
	api.Params
}

// NewMarketReorderAlbumsBulder func
func NewMarketReorderAlbumsBulder() *MarketReorderAlbumsBulder {
	return &MarketReorderAlbumsBulder{api.Params{}}
}

// OwnerID ID of an item owner community.
func (b *MarketReorderAlbumsBulder) OwnerID(v int) {
	b.Params["owner_id"] = v
}

// AlbumID Collection ID.
func (b *MarketReorderAlbumsBulder) AlbumID(v int) {
	b.Params["album_id"] = v
}

// Before ID of a collection to place current collection before it.
func (b *MarketReorderAlbumsBulder) Before(v int) {
	b.Params["before"] = v
}

// After ID of a collection to place current collection after it.
func (b *MarketReorderAlbumsBulder) After(v int) {
	b.Params["after"] = v
}

// MarketReorderItemsBulder builder
//
// Changes item place in a collection.
//
// https://vk.com/dev/market.reorderItems
type MarketReorderItemsBulder struct {
	api.Params
}

// NewMarketReorderItemsBulder func
func NewMarketReorderItemsBulder() *MarketReorderItemsBulder {
	return &MarketReorderItemsBulder{api.Params{}}
}

// OwnerID ID of an item owner community.
func (b *MarketReorderItemsBulder) OwnerID(v int) {
	b.Params["owner_id"] = v
}

// AlbumID ID of a collection to reorder items in. Set 0 to reorder full items list.
func (b *MarketReorderItemsBulder) AlbumID(v int) {
	b.Params["album_id"] = v
}

// ItemID Item ID.
func (b *MarketReorderItemsBulder) ItemID(v int) {
	b.Params["item_id"] = v
}

// Before ID of an item to place current item before it.
func (b *MarketReorderItemsBulder) Before(v int) {
	b.Params["before"] = v
}

// After ID of an item to place current item after it.
func (b *MarketReorderItemsBulder) After(v int) {
	b.Params["after"] = v
}

// MarketReportBulder builder
//
// Sends a complaint to the item.
//
// https://vk.com/dev/market.report
type MarketReportBulder struct {
	api.Params
}

// NewMarketReportBulder func
func NewMarketReportBulder() *MarketReportBulder {
	return &MarketReportBulder{api.Params{}}
}

// OwnerID ID of an item owner community.
func (b *MarketReportBulder) OwnerID(v int) {
	b.Params["owner_id"] = v
}

// ItemID Item ID.
func (b *MarketReportBulder) ItemID(v int) {
	b.Params["item_id"] = v
}

// Reason Complaint reason. Possible values: *'0' — spam,, *'1' — child porn,, *'2' — extremism,, *'3' — violence,, *'4' — drugs propaganda,, *'5' — adult materials,, *'6' — insult.
func (b *MarketReportBulder) Reason(v int) {
	b.Params["reason"] = v
}

// MarketReportCommentBulder builder
//
// Sends a complaint to the item's comment.
//
// https://vk.com/dev/market.reportComment
type MarketReportCommentBulder struct {
	api.Params
}

// NewMarketReportCommentBulder func
func NewMarketReportCommentBulder() *MarketReportCommentBulder {
	return &MarketReportCommentBulder{api.Params{}}
}

// OwnerID ID of an item owner community.
func (b *MarketReportCommentBulder) OwnerID(v int) {
	b.Params["owner_id"] = v
}

// CommentID Comment ID.
func (b *MarketReportCommentBulder) CommentID(v int) {
	b.Params["comment_id"] = v
}

// Reason Complaint reason. Possible values: *'0' — spam,, *'1' — child porn,, *'2' — extremism,, *'3' — violence,, *'4' — drugs propaganda,, *'5' — adult materials,, *'6' — insult.
func (b *MarketReportCommentBulder) Reason(v int) {
	b.Params["reason"] = v
}

// MarketRestoreBulder builder
//
// Restores recently deleted item
//
// https://vk.com/dev/market.restore
type MarketRestoreBulder struct {
	api.Params
}

// NewMarketRestoreBulder func
func NewMarketRestoreBulder() *MarketRestoreBulder {
	return &MarketRestoreBulder{api.Params{}}
}

// OwnerID ID of an item owner community.
func (b *MarketRestoreBulder) OwnerID(v int) {
	b.Params["owner_id"] = v
}

// ItemID Deleted item ID.
func (b *MarketRestoreBulder) ItemID(v int) {
	b.Params["item_id"] = v
}

// MarketRestoreCommentBulder builder
//
// Restores a recently deleted comment
//
// https://vk.com/dev/market.restoreComment
type MarketRestoreCommentBulder struct {
	api.Params
}

// NewMarketRestoreCommentBulder func
func NewMarketRestoreCommentBulder() *MarketRestoreCommentBulder {
	return &MarketRestoreCommentBulder{api.Params{}}
}

// OwnerID identifier of an item owner community, "Note that community id in the 'owner_id' parameter should be negative number. For example 'owner_id'=-1 matches the [vk.com/apiclub|VK API] community "
func (b *MarketRestoreCommentBulder) OwnerID(v int) {
	b.Params["owner_id"] = v
}

// CommentID deleted comment id
func (b *MarketRestoreCommentBulder) CommentID(v int) {
	b.Params["comment_id"] = v
}

// MarketSearchBulder builder
//
// Searches market items in a community's catalog
//
// https://vk.com/dev/market.search
type MarketSearchBulder struct {
	api.Params
}

// NewMarketSearchBulder func
func NewMarketSearchBulder() *MarketSearchBulder {
	return &MarketSearchBulder{api.Params{}}
}

// OwnerID ID of an items owner community.
func (b *MarketSearchBulder) OwnerID(v int) {
	b.Params["owner_id"] = v
}

// AlbumID parameter
func (b *MarketSearchBulder) AlbumID(v int) {
	b.Params["album_id"] = v
}

// Q Search query, for example "pink slippers".
func (b *MarketSearchBulder) Q(v string) {
	b.Params["q"] = v
}

// PriceFrom Minimum item price value.
func (b *MarketSearchBulder) PriceFrom(v int) {
	b.Params["price_from"] = v
}

// PriceTo Maximum item price value.
func (b *MarketSearchBulder) PriceTo(v int) {
	b.Params["price_to"] = v
}

// Tags Comma-separated tag IDs list.
func (b *MarketSearchBulder) Tags(v []int) {
	b.Params["tags"] = v
}

// Sort parameter
func (b *MarketSearchBulder) Sort(v int) {
	b.Params["sort"] = v
}

// Rev '0' — do not use reverse order, '1' — use reverse order
func (b *MarketSearchBulder) Rev(v int) {
	b.Params["rev"] = v
}

// Offset Offset needed to return a specific subset of results.
func (b *MarketSearchBulder) Offset(v int) {
	b.Params["offset"] = v
}

// Count Number of items to return.
func (b *MarketSearchBulder) Count(v int) {
	b.Params["count"] = v
}

// Extended '1' – to return additional fields: 'likes, can_comment, car_repost, photos'. By default: '0'.
func (b *MarketSearchBulder) Extended(v bool) {
	b.Params["extended"] = v
}

// Status parameter
func (b *MarketSearchBulder) Status(v int) {
	b.Params["status"] = v
}
