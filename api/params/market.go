package params // import "github.com/SevereCloud/vksdk/api/params"

import (
	"github.com/SevereCloud/vksdk/api"
)

// MarketAddBuilder builder
//
// Ads a new item to the market.
//
// https://vk.com/dev/market.add
type MarketAddBuilder struct {
	api.Params
}

// NewMarketAddBuilder func
func NewMarketAddBuilder() *MarketAddBuilder {
	return &MarketAddBuilder{api.Params{}}
}

// OwnerID ID of an item owner community.
func (b *MarketAddBuilder) OwnerID(v int) *MarketAddBuilder{
	b.Params["owner_id"] = v
	return b
}

// Name Item name.
func (b *MarketAddBuilder) Name(v string) *MarketAddBuilder{
	b.Params["name"] = v
	return b
}

// Description Item description.
func (b *MarketAddBuilder) Description(v string) *MarketAddBuilder{
	b.Params["description"] = v
	return b
}

// CategoryID Item category ID.
func (b *MarketAddBuilder) CategoryID(v int) *MarketAddBuilder{
	b.Params["category_id"] = v
	return b
}

// Price Item price.
func (b *MarketAddBuilder) Price(v float64) *MarketAddBuilder{
	b.Params["price"] = v
	return b
}

// OldPrice parameter
func (b *MarketAddBuilder) OldPrice(v float64) *MarketAddBuilder{
	b.Params["old_price"] = v
	return b
}

// Deleted Item status
//
// * '1' — deleted
//
// * '0' — not deleted.
func (b *MarketAddBuilder) Deleted(v bool) *MarketAddBuilder{
	b.Params["deleted"] = v
	return b
}

// MainPhotoID Cover photo ID.
func (b *MarketAddBuilder) MainPhotoID(v int) *MarketAddBuilder{
	b.Params["main_photo_id"] = v
	return b
}

// PhotoIDs IDs of additional photos.
func (b *MarketAddBuilder) PhotoIDs(v []int) *MarketAddBuilder{
	b.Params["photo_ids"] = v
	return b
}

// URL Url for button in market item.
func (b *MarketAddBuilder) URL(v string) *MarketAddBuilder{
	b.Params["url"] = v
	return b
}

// MarketAddAlbumBuilder builder
//
// Creates new collection of items
//
// https://vk.com/dev/market.addAlbum
type MarketAddAlbumBuilder struct {
	api.Params
}

// NewMarketAddAlbumBuilder func
func NewMarketAddAlbumBuilder() *MarketAddAlbumBuilder {
	return &MarketAddAlbumBuilder{api.Params{}}
}

// OwnerID ID of an item owner community.
func (b *MarketAddAlbumBuilder) OwnerID(v int) *MarketAddAlbumBuilder{
	b.Params["owner_id"] = v
	return b
}

// Title Collection title.
func (b *MarketAddAlbumBuilder) Title(v string) *MarketAddAlbumBuilder{
	b.Params["title"] = v
	return b
}

// PhotoID Cover photo ID.
func (b *MarketAddAlbumBuilder) PhotoID(v int) *MarketAddAlbumBuilder{
	b.Params["photo_id"] = v
	return b
}

// MainAlbum Set as main
//
// * '1' – set
//
// * '0' – no.
func (b *MarketAddAlbumBuilder) MainAlbum(v bool) *MarketAddAlbumBuilder{
	b.Params["main_album"] = v
	return b
}

// MarketAddToAlbumBuilder builder
//
// Adds an item to one or multiple collections.
//
// https://vk.com/dev/market.addToAlbum
type MarketAddToAlbumBuilder struct {
	api.Params
}

// NewMarketAddToAlbumBuilder func
func NewMarketAddToAlbumBuilder() *MarketAddToAlbumBuilder {
	return &MarketAddToAlbumBuilder{api.Params{}}
}

// OwnerID ID of an item owner community.
func (b *MarketAddToAlbumBuilder) OwnerID(v int) *MarketAddToAlbumBuilder{
	b.Params["owner_id"] = v
	return b
}

// ItemID Item ID.
func (b *MarketAddToAlbumBuilder) ItemID(v int) *MarketAddToAlbumBuilder{
	b.Params["item_id"] = v
	return b
}

// AlbumIDs Collections IDs to add item to.
func (b *MarketAddToAlbumBuilder) AlbumIDs(v []int) *MarketAddToAlbumBuilder{
	b.Params["album_ids"] = v
	return b
}

// MarketCreateCommentBuilder builder
//
// Creates a new comment for an item.
//
// https://vk.com/dev/market.createComment
type MarketCreateCommentBuilder struct {
	api.Params
}

// NewMarketCreateCommentBuilder func
func NewMarketCreateCommentBuilder() *MarketCreateCommentBuilder {
	return &MarketCreateCommentBuilder{api.Params{}}
}

// OwnerID ID of an item owner community.
func (b *MarketCreateCommentBuilder) OwnerID(v int) *MarketCreateCommentBuilder{
	b.Params["owner_id"] = v
	return b
}

// ItemID Item ID.
func (b *MarketCreateCommentBuilder) ItemID(v int) *MarketCreateCommentBuilder{
	b.Params["item_id"] = v
	return b
}

// Message Comment text (required if 'attachments' parameter is not specified)
func (b *MarketCreateCommentBuilder) Message(v string) *MarketCreateCommentBuilder{
	b.Params["message"] = v
	return b
}

// Attachments Comma-separated list of objects attached to a comment.
// The field is submitted the following way:
// '<owner_id>_<media_id>,<owner_id>_<media_id>'
//
// * '' - media attachment type:
//
// * 'photo' - photo
//
// * 'video' - video
//
// * 'audio' - audio
//
// * 'doc' - document
//
// * '<owner_id>' - media owner id
//
// * '<media_id>' - media attachment id
//
// For example: photo100172_166443618,photo66748_265827614,
func (b *MarketCreateCommentBuilder) Attachments(v []string) *MarketCreateCommentBuilder{
	b.Params["attachments"] = v
	return b
}

// FromGroup parameter 
//
// * '1' - comment will be published on behalf of a community
//
// * '0' - on behalf of a user (by default).
func (b *MarketCreateCommentBuilder) FromGroup(v bool) *MarketCreateCommentBuilder{
	b.Params["from_group"] = v
	return b
}

// ReplyToComment ID of a comment to reply with current comment to.
func (b *MarketCreateCommentBuilder) ReplyToComment(v int) *MarketCreateCommentBuilder{
	b.Params["reply_to_comment"] = v
	return b
}

// StickerID Sticker ID.
func (b *MarketCreateCommentBuilder) StickerID(v int) *MarketCreateCommentBuilder{
	b.Params["sticker_id"] = v
	return b
}

// GUID Random value to avoid resending one comment.
func (b *MarketCreateCommentBuilder) GUID(v string) *MarketCreateCommentBuilder{
	b.Params["guid"] = v
	return b
}

// MarketDeleteBuilder builder
//
// Deletes an item.
//
// https://vk.com/dev/market.delete
type MarketDeleteBuilder struct {
	api.Params
}

// NewMarketDeleteBuilder func
func NewMarketDeleteBuilder() *MarketDeleteBuilder {
	return &MarketDeleteBuilder{api.Params{}}
}

// OwnerID ID of an item owner community.
func (b *MarketDeleteBuilder) OwnerID(v int) *MarketDeleteBuilder{
	b.Params["owner_id"] = v
	return b
}

// ItemID Item ID.
func (b *MarketDeleteBuilder) ItemID(v int) *MarketDeleteBuilder{
	b.Params["item_id"] = v
	return b
}

// MarketDeleteAlbumBuilder builder
//
// Deletes a collection of items.
//
// https://vk.com/dev/market.deleteAlbum
type MarketDeleteAlbumBuilder struct {
	api.Params
}

// NewMarketDeleteAlbumBuilder func
func NewMarketDeleteAlbumBuilder() *MarketDeleteAlbumBuilder {
	return &MarketDeleteAlbumBuilder{api.Params{}}
}

// OwnerID ID of an collection owner community.
func (b *MarketDeleteAlbumBuilder) OwnerID(v int) *MarketDeleteAlbumBuilder{
	b.Params["owner_id"] = v
	return b
}

// AlbumID Collection ID.
func (b *MarketDeleteAlbumBuilder) AlbumID(v int) *MarketDeleteAlbumBuilder{
	b.Params["album_id"] = v
	return b
}

// MarketDeleteCommentBuilder builder
//
// Deletes an item's comment
//
// https://vk.com/dev/market.deleteComment
type MarketDeleteCommentBuilder struct {
	api.Params
}

// NewMarketDeleteCommentBuilder func
func NewMarketDeleteCommentBuilder() *MarketDeleteCommentBuilder {
	return &MarketDeleteCommentBuilder{api.Params{}}
}

// OwnerID identifier of an item owner community
// Note that community id in the 'owner_id' parameter should be negative number.
//
// For example 'owner_id'=-1 matches the [vk.com/apiclub|VK API] community 
func (b *MarketDeleteCommentBuilder) OwnerID(v int) *MarketDeleteCommentBuilder{
	b.Params["owner_id"] = v
	return b
}

// CommentID comment id
func (b *MarketDeleteCommentBuilder) CommentID(v int) *MarketDeleteCommentBuilder{
	b.Params["comment_id"] = v
	return b
}

// MarketEditBuilder builder
//
// Edits an item.
//
// https://vk.com/dev/market.edit
type MarketEditBuilder struct {
	api.Params
}

// NewMarketEditBuilder func
func NewMarketEditBuilder() *MarketEditBuilder {
	return &MarketEditBuilder{api.Params{}}
}

// OwnerID ID of an item owner community.
func (b *MarketEditBuilder) OwnerID(v int) *MarketEditBuilder{
	b.Params["owner_id"] = v
	return b
}

// ItemID Item ID.
func (b *MarketEditBuilder) ItemID(v int) *MarketEditBuilder{
	b.Params["item_id"] = v
	return b
}

// Name Item name.
func (b *MarketEditBuilder) Name(v string) *MarketEditBuilder{
	b.Params["name"] = v
	return b
}

// Description Item description.
func (b *MarketEditBuilder) Description(v string) *MarketEditBuilder{
	b.Params["description"] = v
	return b
}

// CategoryID Item category ID.
func (b *MarketEditBuilder) CategoryID(v int) *MarketEditBuilder{
	b.Params["category_id"] = v
	return b
}

// Price Item price.
func (b *MarketEditBuilder) Price(v float64) *MarketEditBuilder{
	b.Params["price"] = v
	return b
}

// Deleted Item status
//
// * '1' — deleted
//
// * '0' — not deleted.
func (b *MarketEditBuilder) Deleted(v bool) *MarketEditBuilder{
	b.Params["deleted"] = v
	return b
}

// MainPhotoID Cover photo ID.
func (b *MarketEditBuilder) MainPhotoID(v int) *MarketEditBuilder{
	b.Params["main_photo_id"] = v
	return b
}

// PhotoIDs IDs of additional photos.
func (b *MarketEditBuilder) PhotoIDs(v []int) *MarketEditBuilder{
	b.Params["photo_ids"] = v
	return b
}

// URL Url for button in market item.
func (b *MarketEditBuilder) URL(v string) *MarketEditBuilder{
	b.Params["url"] = v
	return b
}

// MarketEditAlbumBuilder builder
//
// Edits a collection of items
//
// https://vk.com/dev/market.editAlbum
type MarketEditAlbumBuilder struct {
	api.Params
}

// NewMarketEditAlbumBuilder func
func NewMarketEditAlbumBuilder() *MarketEditAlbumBuilder {
	return &MarketEditAlbumBuilder{api.Params{}}
}

// OwnerID ID of an collection owner community.
func (b *MarketEditAlbumBuilder) OwnerID(v int) *MarketEditAlbumBuilder{
	b.Params["owner_id"] = v
	return b
}

// AlbumID Collection ID.
func (b *MarketEditAlbumBuilder) AlbumID(v int) *MarketEditAlbumBuilder{
	b.Params["album_id"] = v
	return b
}

// Title Collection title.
func (b *MarketEditAlbumBuilder) Title(v string) *MarketEditAlbumBuilder{
	b.Params["title"] = v
	return b
}

// PhotoID Cover photo id
func (b *MarketEditAlbumBuilder) PhotoID(v int) *MarketEditAlbumBuilder{
	b.Params["photo_id"] = v
	return b
}

// MainAlbum Set as main
//
// * '1' – set
//
// * '0' – no.
func (b *MarketEditAlbumBuilder) MainAlbum(v bool) *MarketEditAlbumBuilder{
	b.Params["main_album"] = v
	return b
}

// MarketEditCommentBuilder builder
//
// Chages item comment's text
//
// https://vk.com/dev/market.editComment
type MarketEditCommentBuilder struct {
	api.Params
}

// NewMarketEditCommentBuilder func
func NewMarketEditCommentBuilder() *MarketEditCommentBuilder {
	return &MarketEditCommentBuilder{api.Params{}}
}

// OwnerID ID of an item owner community.
func (b *MarketEditCommentBuilder) OwnerID(v int) *MarketEditCommentBuilder{
	b.Params["owner_id"] = v
	return b
}

// CommentID Comment ID.
func (b *MarketEditCommentBuilder) CommentID(v int) *MarketEditCommentBuilder{
	b.Params["comment_id"] = v
	return b
}

// Message New comment text (required if 'attachments' are not specified) 2048 symbols maximum.
func (b *MarketEditCommentBuilder) Message(v string) *MarketEditCommentBuilder{
	b.Params["message"] = v
	return b
}

// Attachments Comma-separated list of objects attached to a comment.
// The field is submitted the following way:
// '<owner_id>_<media_id>,<owner_id>_<media_id>'
//
// * '' - media attachment type:
//
// * 'photo' - photo
//
// * 'video' - video
//
// * 'audio' - audio
//
// * 'doc' - document
//
// * '<owner_id>' - media owner id
//
// * '<media_id>' - media attachment id
//
// For example: photo100172_166443618,photo66748_265827614,
func (b *MarketEditCommentBuilder) Attachments(v []string) *MarketEditCommentBuilder{
	b.Params["attachments"] = v
	return b
}

// MarketGetBuilder builder
//
// Returns items list for a community.
//
// https://vk.com/dev/market.get
type MarketGetBuilder struct {
	api.Params
}

// NewMarketGetBuilder func
func NewMarketGetBuilder() *MarketGetBuilder {
	return &MarketGetBuilder{api.Params{}}
}

// OwnerID ID of an item owner community
// Note that community id in the 'owner_id' parameter should be negative number.
//
// For example 'owner_id'=-1 matches the [vk.com/apiclub|VK API] community 
func (b *MarketGetBuilder) OwnerID(v int) *MarketGetBuilder{
	b.Params["owner_id"] = v
	return b
}

// AlbumID parameter
func (b *MarketGetBuilder) AlbumID(v int) *MarketGetBuilder{
	b.Params["album_id"] = v
	return b
}

// Count Number of items to return.
func (b *MarketGetBuilder) Count(v int) *MarketGetBuilder{
	b.Params["count"] = v
	return b
}

// Offset Offset needed to return a specific subset of results.
func (b *MarketGetBuilder) Offset(v int) *MarketGetBuilder{
	b.Params["offset"] = v
	return b
}

// Extended parameter 
//
// * '1' – method will return additional fields: 'likes, can_comment, car_repost, photos'.
// These parameters are not returned by default.
func (b *MarketGetBuilder) Extended(v bool) *MarketGetBuilder{
	b.Params["extended"] = v
	return b
}

// MarketGetAlbumByIDBuilder builder
//
// Returns items album's data
//
// https://vk.com/dev/market.getAlbumById
type MarketGetAlbumByIDBuilder struct {
	api.Params
}

// NewMarketGetAlbumByIDBuilder func
func NewMarketGetAlbumByIDBuilder() *MarketGetAlbumByIDBuilder {
	return &MarketGetAlbumByIDBuilder{api.Params{}}
}

// OwnerID identifier of an album owner community
// Note that community id in the 'owner_id' parameter should be negative number.
//
// For example 'owner_id'=-1 matches the [vk.com/apiclub|VK API] community 
func (b *MarketGetAlbumByIDBuilder) OwnerID(v int) *MarketGetAlbumByIDBuilder{
	b.Params["owner_id"] = v
	return b
}

// AlbumIDs collections identifiers to obtain data from
func (b *MarketGetAlbumByIDBuilder) AlbumIDs(v []int) *MarketGetAlbumByIDBuilder{
	b.Params["album_ids"] = v
	return b
}

// MarketGetAlbumsBuilder builder
//
// Returns community's collections list.
//
// https://vk.com/dev/market.getAlbums
type MarketGetAlbumsBuilder struct {
	api.Params
}

// NewMarketGetAlbumsBuilder func
func NewMarketGetAlbumsBuilder() *MarketGetAlbumsBuilder {
	return &MarketGetAlbumsBuilder{api.Params{}}
}

// OwnerID ID of an items owner community.
func (b *MarketGetAlbumsBuilder) OwnerID(v int) *MarketGetAlbumsBuilder{
	b.Params["owner_id"] = v
	return b
}

// Offset Offset needed to return a specific subset of results.
func (b *MarketGetAlbumsBuilder) Offset(v int) *MarketGetAlbumsBuilder{
	b.Params["offset"] = v
	return b
}

// Count Number of items to return.
func (b *MarketGetAlbumsBuilder) Count(v int) *MarketGetAlbumsBuilder{
	b.Params["count"] = v
	return b
}

// MarketGetByIDBuilder builder
//
// Returns information about market items by their ids.
//
// https://vk.com/dev/market.getById
type MarketGetByIDBuilder struct {
	api.Params
}

// NewMarketGetByIDBuilder func
func NewMarketGetByIDBuilder() *MarketGetByIDBuilder {
	return &MarketGetByIDBuilder{api.Params{}}
}

// ItemIDs Comma-separated ids list: {user id}_{item id}.
// If an item belongs to a community -{community id} is used.
// 'Videos' value example:
// '-4363_136089719,13245770_137352259'
func (b *MarketGetByIDBuilder) ItemIDs(v []string) *MarketGetByIDBuilder{
	b.Params["item_ids"] = v
	return b
}

// Extended parameter 
//
// * '1' – to return additional fields: 'likes, can_comment, car_repost, photos'.
//
// By default: '0'.
func (b *MarketGetByIDBuilder) Extended(v bool) *MarketGetByIDBuilder{
	b.Params["extended"] = v
	return b
}

// MarketGetCategoriesBuilder builder
//
// Returns a list of market categories.
//
// https://vk.com/dev/market.getCategories
type MarketGetCategoriesBuilder struct {
	api.Params
}

// NewMarketGetCategoriesBuilder func
func NewMarketGetCategoriesBuilder() *MarketGetCategoriesBuilder {
	return &MarketGetCategoriesBuilder{api.Params{}}
}

// Count Number of results to return.
func (b *MarketGetCategoriesBuilder) Count(v int) *MarketGetCategoriesBuilder{
	b.Params["count"] = v
	return b
}

// Offset Offset needed to return a specific subset of results.
func (b *MarketGetCategoriesBuilder) Offset(v int) *MarketGetCategoriesBuilder{
	b.Params["offset"] = v
	return b
}

// MarketGetCommentsBuilder builder
//
// Returns comments list for an item.
//
// https://vk.com/dev/market.getComments
type MarketGetCommentsBuilder struct {
	api.Params
}

// NewMarketGetCommentsBuilder func
func NewMarketGetCommentsBuilder() *MarketGetCommentsBuilder {
	return &MarketGetCommentsBuilder{api.Params{}}
}

// OwnerID ID of an item owner community
func (b *MarketGetCommentsBuilder) OwnerID(v int) *MarketGetCommentsBuilder{
	b.Params["owner_id"] = v
	return b
}

// ItemID Item ID.
func (b *MarketGetCommentsBuilder) ItemID(v int) *MarketGetCommentsBuilder{
	b.Params["item_id"] = v
	return b
}

// NeedLikes parameter 
//
// * '1' — to return likes info.
func (b *MarketGetCommentsBuilder) NeedLikes(v bool) *MarketGetCommentsBuilder{
	b.Params["need_likes"] = v
	return b
}

// StartCommentID ID of a comment to start a list from (details below).
func (b *MarketGetCommentsBuilder) StartCommentID(v int) *MarketGetCommentsBuilder{
	b.Params["start_comment_id"] = v
	return b
}

// Offset parameter
func (b *MarketGetCommentsBuilder) Offset(v int) *MarketGetCommentsBuilder{
	b.Params["offset"] = v
	return b
}

// Count Number of results to return.
func (b *MarketGetCommentsBuilder) Count(v int) *MarketGetCommentsBuilder{
	b.Params["count"] = v
	return b
}

// Sort Sort order
//
// * 'asc' — from old to new
//
// * 'desc' — from new to old
func (b *MarketGetCommentsBuilder) Sort(v string) *MarketGetCommentsBuilder{
	b.Params["sort"] = v
	return b
}

// Extended parameter 
//
// * '1' — comments will be returned as numbered objects, in addition lists of 'profiles' and 'groups' objects will be returned.
func (b *MarketGetCommentsBuilder) Extended(v bool) *MarketGetCommentsBuilder{
	b.Params["extended"] = v
	return b
}

// Fields List of additional profile fields to return.
// See the [vk.com/dev/fields|details]
func (b *MarketGetCommentsBuilder) Fields(v []string) *MarketGetCommentsBuilder{
	b.Params["fields"] = v
	return b
}

// MarketRemoveFromAlbumBuilder builder
//
// Removes an item from one or multiple collections.
//
// https://vk.com/dev/market.removeFromAlbum
type MarketRemoveFromAlbumBuilder struct {
	api.Params
}

// NewMarketRemoveFromAlbumBuilder func
func NewMarketRemoveFromAlbumBuilder() *MarketRemoveFromAlbumBuilder {
	return &MarketRemoveFromAlbumBuilder{api.Params{}}
}

// OwnerID ID of an item owner community.
func (b *MarketRemoveFromAlbumBuilder) OwnerID(v int) *MarketRemoveFromAlbumBuilder{
	b.Params["owner_id"] = v
	return b
}

// ItemID Item ID.
func (b *MarketRemoveFromAlbumBuilder) ItemID(v int) *MarketRemoveFromAlbumBuilder{
	b.Params["item_id"] = v
	return b
}

// AlbumIDs Collections IDs to remove item from.
func (b *MarketRemoveFromAlbumBuilder) AlbumIDs(v []int) *MarketRemoveFromAlbumBuilder{
	b.Params["album_ids"] = v
	return b
}

// MarketReorderAlbumsBuilder builder
//
// Reorders the collections list.
//
// https://vk.com/dev/market.reorderAlbums
type MarketReorderAlbumsBuilder struct {
	api.Params
}

// NewMarketReorderAlbumsBuilder func
func NewMarketReorderAlbumsBuilder() *MarketReorderAlbumsBuilder {
	return &MarketReorderAlbumsBuilder{api.Params{}}
}

// OwnerID ID of an item owner community.
func (b *MarketReorderAlbumsBuilder) OwnerID(v int) *MarketReorderAlbumsBuilder{
	b.Params["owner_id"] = v
	return b
}

// AlbumID Collection ID.
func (b *MarketReorderAlbumsBuilder) AlbumID(v int) *MarketReorderAlbumsBuilder{
	b.Params["album_id"] = v
	return b
}

// Before ID of a collection to place current collection before it.
func (b *MarketReorderAlbumsBuilder) Before(v int) *MarketReorderAlbumsBuilder{
	b.Params["before"] = v
	return b
}

// After ID of a collection to place current collection after it.
func (b *MarketReorderAlbumsBuilder) After(v int) *MarketReorderAlbumsBuilder{
	b.Params["after"] = v
	return b
}

// MarketReorderItemsBuilder builder
//
// Changes item place in a collection.
//
// https://vk.com/dev/market.reorderItems
type MarketReorderItemsBuilder struct {
	api.Params
}

// NewMarketReorderItemsBuilder func
func NewMarketReorderItemsBuilder() *MarketReorderItemsBuilder {
	return &MarketReorderItemsBuilder{api.Params{}}
}

// OwnerID ID of an item owner community.
func (b *MarketReorderItemsBuilder) OwnerID(v int) *MarketReorderItemsBuilder{
	b.Params["owner_id"] = v
	return b
}

// AlbumID ID of a collection to reorder items in.
// Set 0 to reorder full items list.
func (b *MarketReorderItemsBuilder) AlbumID(v int) *MarketReorderItemsBuilder{
	b.Params["album_id"] = v
	return b
}

// ItemID Item ID.
func (b *MarketReorderItemsBuilder) ItemID(v int) *MarketReorderItemsBuilder{
	b.Params["item_id"] = v
	return b
}

// Before ID of an item to place current item before it.
func (b *MarketReorderItemsBuilder) Before(v int) *MarketReorderItemsBuilder{
	b.Params["before"] = v
	return b
}

// After ID of an item to place current item after it.
func (b *MarketReorderItemsBuilder) After(v int) *MarketReorderItemsBuilder{
	b.Params["after"] = v
	return b
}

// MarketReportBuilder builder
//
// Sends a complaint to the item.
//
// https://vk.com/dev/market.report
type MarketReportBuilder struct {
	api.Params
}

// NewMarketReportBuilder func
func NewMarketReportBuilder() *MarketReportBuilder {
	return &MarketReportBuilder{api.Params{}}
}

// OwnerID ID of an item owner community.
func (b *MarketReportBuilder) OwnerID(v int) *MarketReportBuilder{
	b.Params["owner_id"] = v
	return b
}

// ItemID Item ID.
func (b *MarketReportBuilder) ItemID(v int) *MarketReportBuilder{
	b.Params["item_id"] = v
	return b
}

// Reason Complaint reason. Possible values:
//
// * 0 — spam
//
// * 1 — child porn
//
// * 2 — extremism
//
// * 3 — violence
//
// * 4 — drugs propaganda
//
// * 5 — adult materials
//
// * 6 — insult.
func (b *MarketReportBuilder) Reason(v int) *MarketReportBuilder{
	b.Params["reason"] = v
	return b
}

// MarketReportCommentBuilder builder
//
// Sends a complaint to the item's comment.
//
// https://vk.com/dev/market.reportComment
type MarketReportCommentBuilder struct {
	api.Params
}

// NewMarketReportCommentBuilder func
func NewMarketReportCommentBuilder() *MarketReportCommentBuilder {
	return &MarketReportCommentBuilder{api.Params{}}
}

// OwnerID ID of an item owner community.
func (b *MarketReportCommentBuilder) OwnerID(v int) *MarketReportCommentBuilder{
	b.Params["owner_id"] = v
	return b
}

// CommentID Comment ID.
func (b *MarketReportCommentBuilder) CommentID(v int) *MarketReportCommentBuilder{
	b.Params["comment_id"] = v
	return b
}

// Reason Complaint reason. Possible values:
//
// * 0 — spam
//
// * 1 — child porn
//
// * 2 — extremism
//
// * 3 — violence
//
// * 4 — drugs propaganda
//
// * 5 — adult materials
//
// * 6 — insult.
func (b *MarketReportCommentBuilder) Reason(v int) *MarketReportCommentBuilder{
	b.Params["reason"] = v
	return b
}

// MarketRestoreBuilder builder
//
// Restores recently deleted item
//
// https://vk.com/dev/market.restore
type MarketRestoreBuilder struct {
	api.Params
}

// NewMarketRestoreBuilder func
func NewMarketRestoreBuilder() *MarketRestoreBuilder {
	return &MarketRestoreBuilder{api.Params{}}
}

// OwnerID ID of an item owner community.
func (b *MarketRestoreBuilder) OwnerID(v int) *MarketRestoreBuilder{
	b.Params["owner_id"] = v
	return b
}

// ItemID Deleted item ID.
func (b *MarketRestoreBuilder) ItemID(v int) *MarketRestoreBuilder{
	b.Params["item_id"] = v
	return b
}

// MarketRestoreCommentBuilder builder
//
// Restores a recently deleted comment
//
// https://vk.com/dev/market.restoreComment
type MarketRestoreCommentBuilder struct {
	api.Params
}

// NewMarketRestoreCommentBuilder func
func NewMarketRestoreCommentBuilder() *MarketRestoreCommentBuilder {
	return &MarketRestoreCommentBuilder{api.Params{}}
}

// OwnerID identifier of an item owner community
// Note that community id in the 'owner_id' parameter should be negative number.
//
// For example 'owner_id'=-1 matches the [vk.com/apiclub|VK API] community 
func (b *MarketRestoreCommentBuilder) OwnerID(v int) *MarketRestoreCommentBuilder{
	b.Params["owner_id"] = v
	return b
}

// CommentID deleted comment id
func (b *MarketRestoreCommentBuilder) CommentID(v int) *MarketRestoreCommentBuilder{
	b.Params["comment_id"] = v
	return b
}

// MarketSearchBuilder builder
//
// Searches market items in a community's catalog
//
// https://vk.com/dev/market.search
type MarketSearchBuilder struct {
	api.Params
}

// NewMarketSearchBuilder func
func NewMarketSearchBuilder() *MarketSearchBuilder {
	return &MarketSearchBuilder{api.Params{}}
}

// OwnerID ID of an items owner community.
func (b *MarketSearchBuilder) OwnerID(v int) *MarketSearchBuilder{
	b.Params["owner_id"] = v
	return b
}

// AlbumID parameter
func (b *MarketSearchBuilder) AlbumID(v int) *MarketSearchBuilder{
	b.Params["album_id"] = v
	return b
}

// Q Search query, for example pink slippers.
func (b *MarketSearchBuilder) Q(v string) *MarketSearchBuilder{
	b.Params["q"] = v
	return b
}

// PriceFrom Minimum item price value.
func (b *MarketSearchBuilder) PriceFrom(v int) *MarketSearchBuilder{
	b.Params["price_from"] = v
	return b
}

// PriceTo Maximum item price value.
func (b *MarketSearchBuilder) PriceTo(v int) *MarketSearchBuilder{
	b.Params["price_to"] = v
	return b
}

// Tags Comma-separated tag IDs list.
func (b *MarketSearchBuilder) Tags(v []int) *MarketSearchBuilder{
	b.Params["tags"] = v
	return b
}

// Sort parameter
func (b *MarketSearchBuilder) Sort(v int) *MarketSearchBuilder{
	b.Params["sort"] = v
	return b
}

// Rev parameter 
//
// * '0' — do not use reverse order
//
// * '1' — use reverse order
func (b *MarketSearchBuilder) Rev(v int) *MarketSearchBuilder{
	b.Params["rev"] = v
	return b
}

// Offset Offset needed to return a specific subset of results.
func (b *MarketSearchBuilder) Offset(v int) *MarketSearchBuilder{
	b.Params["offset"] = v
	return b
}

// Count Number of items to return.
func (b *MarketSearchBuilder) Count(v int) *MarketSearchBuilder{
	b.Params["count"] = v
	return b
}

// Extended parameter 
//
// * '1' – to return additional fields: 'likes, can_comment, car_repost, photos'.
//
// By default: '0'.
func (b *MarketSearchBuilder) Extended(v bool) *MarketSearchBuilder{
	b.Params["extended"] = v
	return b
}

// Status parameter
func (b *MarketSearchBuilder) Status(v int) *MarketSearchBuilder{
	b.Params["status"] = v
	return b
}
