package params // import "github.com/SevereCloud/vksdk/api/params"

import (
	"github.com/SevereCloud/vksdk/api"
)

// WallCloseCommentsBulder builder
//
// https://vk.com/dev/wall.closeComments
type WallCloseCommentsBulder struct {
	api.Params
}

// NewWallCloseCommentsBulder func
func NewWallCloseCommentsBulder() *WallCloseCommentsBulder {
	return &WallCloseCommentsBulder{api.Params{}}
}

// OwnerID parameter
func (b *WallCloseCommentsBulder) OwnerID(v int) {
	b.Params["owner_id"] = v
}

// PostID parameter
func (b *WallCloseCommentsBulder) PostID(v int) {
	b.Params["post_id"] = v
}

// WallCreateCommentBulder builder
//
// Adds a comment to a post on a user wall or community wall.
//
// https://vk.com/dev/wall.createComment
type WallCreateCommentBulder struct {
	api.Params
}

// NewWallCreateCommentBulder func
func NewWallCreateCommentBulder() *WallCreateCommentBulder {
	return &WallCreateCommentBulder{api.Params{}}
}

// OwnerID User ID or community ID. Use a negative value to designate a community ID.
func (b *WallCreateCommentBulder) OwnerID(v int) {
	b.Params["owner_id"] = v
}

// PostID Post ID.
func (b *WallCreateCommentBulder) PostID(v int) {
	b.Params["post_id"] = v
}

// FromGroup Group ID.
func (b *WallCreateCommentBulder) FromGroup(v int) {
	b.Params["from_group"] = v
}

// Message (Required if 'attachments' is not set.) Text of the comment.
func (b *WallCreateCommentBulder) Message(v string) {
	b.Params["message"] = v
}

// ReplyToComment ID of comment to reply.
func (b *WallCreateCommentBulder) ReplyToComment(v int) {
	b.Params["reply_to_comment"] = v
}

// Attachments (Required if 'message' is not set.) List of media objects attached to the comment, in the following format: "<owner_id>_<media_id>,<owner_id>_<media_id>", '' — Type of media ojbect: 'photo' — photo, 'video' — video, 'audio' — audio, 'doc' — document, '<owner_id>' — ID of the media owner. '<media_id>' — Media ID. For example: "photo100172_166443618,photo66748_265827614"
func (b *WallCreateCommentBulder) Attachments(v []string) {
	b.Params["attachments"] = v
}

// StickerID Sticker ID.
func (b *WallCreateCommentBulder) StickerID(v int) {
	b.Params["sticker_id"] = v
}

// GUID Unique identifier to avoid repeated comments.
func (b *WallCreateCommentBulder) GUID(v string) {
	b.Params["guid"] = v
}

// WallDeleteBulder builder
//
// Deletes a post from a user wall or community wall.
//
// https://vk.com/dev/wall.delete
type WallDeleteBulder struct {
	api.Params
}

// NewWallDeleteBulder func
func NewWallDeleteBulder() *WallDeleteBulder {
	return &WallDeleteBulder{api.Params{}}
}

// OwnerID User ID or community ID. Use a negative value to designate a community ID.
func (b *WallDeleteBulder) OwnerID(v int) {
	b.Params["owner_id"] = v
}

// PostID ID of the post to be deleted.
func (b *WallDeleteBulder) PostID(v int) {
	b.Params["post_id"] = v
}

// WallDeleteCommentBulder builder
//
// Deletes a comment on a post on a user wall or community wall.
//
// https://vk.com/dev/wall.deleteComment
type WallDeleteCommentBulder struct {
	api.Params
}

// NewWallDeleteCommentBulder func
func NewWallDeleteCommentBulder() *WallDeleteCommentBulder {
	return &WallDeleteCommentBulder{api.Params{}}
}

// OwnerID User ID or community ID. Use a negative value to designate a community ID.
func (b *WallDeleteCommentBulder) OwnerID(v int) {
	b.Params["owner_id"] = v
}

// CommentID Comment ID.
func (b *WallDeleteCommentBulder) CommentID(v int) {
	b.Params["comment_id"] = v
}

// WallEditBulder builder
//
// Edits a post on a user wall or community wall.
//
// https://vk.com/dev/wall.edit
type WallEditBulder struct {
	api.Params
}

// NewWallEditBulder func
func NewWallEditBulder() *WallEditBulder {
	return &WallEditBulder{api.Params{}}
}

// OwnerID User ID or community ID. Use a negative value to designate a community ID.
func (b *WallEditBulder) OwnerID(v int) {
	b.Params["owner_id"] = v
}

// PostID parameter
func (b *WallEditBulder) PostID(v int) {
	b.Params["post_id"] = v
}

// FriendsOnly parameter
func (b *WallEditBulder) FriendsOnly(v bool) {
	b.Params["friends_only"] = v
}

// Message (Required if 'attachments' is not set.) Text of the post.
func (b *WallEditBulder) Message(v string) {
	b.Params["message"] = v
}

// Attachments (Required if 'message' is not set.) List of objects attached to the post, in the following format: "<owner_id>_<media_id>,<owner_id>_<media_id>", '' — Type of media attachment: 'photo' — photo, 'video' — video, 'audio' — audio, 'doc' — document, '<owner_id>' — ID of the media application owner. '<media_id>' — Media application ID. Example: "photo100172_166443618,photo66748_265827614", May contain a link to an external page to include in the post. Example: "photo66748_265827614,http://habrahabr.ru", "NOTE: If more than one link is being attached, an error is thrown."
func (b *WallEditBulder) Attachments(v []string) {
	b.Params["attachments"] = v
}

// Services parameter
func (b *WallEditBulder) Services(v string) {
	b.Params["services"] = v
}

// Signed parameter
func (b *WallEditBulder) Signed(v bool) {
	b.Params["signed"] = v
}

// PublishDate parameter
func (b *WallEditBulder) PublishDate(v int) {
	b.Params["publish_date"] = v
}

// Lat parameter
func (b *WallEditBulder) Lat(v float64) {
	b.Params["lat"] = v
}

// Long parameter
func (b *WallEditBulder) Long(v float64) {
	b.Params["long"] = v
}

// PlaceID parameter
func (b *WallEditBulder) PlaceID(v int) {
	b.Params["place_id"] = v
}

// MarkAsAds parameter
func (b *WallEditBulder) MarkAsAds(v bool) {
	b.Params["mark_as_ads"] = v
}

// CloseComments parameter
func (b *WallEditBulder) CloseComments(v bool) {
	b.Params["close_comments"] = v
}

// PosterBkgID parameter
func (b *WallEditBulder) PosterBkgID(v int) {
	b.Params["poster_bkg_id"] = v
}

// PosterBkgOwnerID parameter
func (b *WallEditBulder) PosterBkgOwnerID(v int) {
	b.Params["poster_bkg_owner_id"] = v
}

// PosterBkgAccessHash parameter
func (b *WallEditBulder) PosterBkgAccessHash(v string) {
	b.Params["poster_bkg_access_hash"] = v
}

// WallEditAdsStealthBulder builder
//
// Allows to edit hidden post.
//
// https://vk.com/dev/wall.editAdsStealth
type WallEditAdsStealthBulder struct {
	api.Params
}

// NewWallEditAdsStealthBulder func
func NewWallEditAdsStealthBulder() *WallEditAdsStealthBulder {
	return &WallEditAdsStealthBulder{api.Params{}}
}

// OwnerID User ID or community ID. Use a negative value to designate a community ID.
func (b *WallEditAdsStealthBulder) OwnerID(v int) {
	b.Params["owner_id"] = v
}

// PostID Post ID. Used for publishing of scheduled and suggested posts.
func (b *WallEditAdsStealthBulder) PostID(v int) {
	b.Params["post_id"] = v
}

// Message (Required if 'attachments' is not set.) Text of the post.
func (b *WallEditAdsStealthBulder) Message(v string) {
	b.Params["message"] = v
}

// Attachments (Required if 'message' is not set.) List of objects attached to the post, in the following format: "<owner_id>_<media_id>,<owner_id>_<media_id>", '' — Type of media attachment: 'photo' — photo, 'video' — video, 'audio' — audio, 'doc' — document, 'page' — wiki-page, 'note' — note, 'poll' — poll, 'album' — photo album, '<owner_id>' — ID of the media application owner. '<media_id>' — Media application ID. Example: "photo100172_166443618,photo66748_265827614", May contain a link to an external page to include in the post. Example: "photo66748_265827614,http://habrahabr.ru", "NOTE: If more than one link is being attached, an error will be thrown."
func (b *WallEditAdsStealthBulder) Attachments(v []string) {
	b.Params["attachments"] = v
}

// Signed Only for posts in communities with 'from_group' set to '1': '1' — post will be signed with the name of the posting user, '0' — post will not be signed (default)
func (b *WallEditAdsStealthBulder) Signed(v bool) {
	b.Params["signed"] = v
}

// Lat Geographical latitude of a check-in, in degrees (from -90 to 90).
func (b *WallEditAdsStealthBulder) Lat(v float64) {
	b.Params["lat"] = v
}

// Long Geographical longitude of a check-in, in degrees (from -180 to 180).
func (b *WallEditAdsStealthBulder) Long(v float64) {
	b.Params["long"] = v
}

// PlaceID ID of the location where the user was tagged.
func (b *WallEditAdsStealthBulder) PlaceID(v int) {
	b.Params["place_id"] = v
}

// LinkButton Link button ID
func (b *WallEditAdsStealthBulder) LinkButton(v string) {
	b.Params["link_button"] = v
}

// LinkTitle Link title
func (b *WallEditAdsStealthBulder) LinkTitle(v string) {
	b.Params["link_title"] = v
}

// LinkImage Link image url
func (b *WallEditAdsStealthBulder) LinkImage(v string) {
	b.Params["link_image"] = v
}

// LinkVideo Link video ID in format "<owner_id>_<media_id>"
func (b *WallEditAdsStealthBulder) LinkVideo(v string) {
	b.Params["link_video"] = v
}

// WallEditCommentBulder builder
//
// Edits a comment on a user wall or community wall.
//
// https://vk.com/dev/wall.editComment
type WallEditCommentBulder struct {
	api.Params
}

// NewWallEditCommentBulder func
func NewWallEditCommentBulder() *WallEditCommentBulder {
	return &WallEditCommentBulder{api.Params{}}
}

// OwnerID User ID or community ID. Use a negative value to designate a community ID.
func (b *WallEditCommentBulder) OwnerID(v int) {
	b.Params["owner_id"] = v
}

// CommentID Comment ID.
func (b *WallEditCommentBulder) CommentID(v int) {
	b.Params["comment_id"] = v
}

// Message New comment text.
func (b *WallEditCommentBulder) Message(v string) {
	b.Params["message"] = v
}

// Attachments List of objects attached to the comment, in the following format: , "<owner_id>_<media_id>,<owner_id>_<media_id>", '' — Type of media attachment: 'photo' — photo, 'video' — video, 'audio' — audio, 'doc' — document, '<owner_id>' — ID of the media attachment owner. '<media_id>' — Media attachment ID. For example: "photo100172_166443618,photo66748_265827614"
func (b *WallEditCommentBulder) Attachments(v []string) {
	b.Params["attachments"] = v
}

// WallGetBulder builder
//
// Returns a list of posts on a user wall or community wall.
//
// https://vk.com/dev/wall.get
type WallGetBulder struct {
	api.Params
}

// NewWallGetBulder func
func NewWallGetBulder() *WallGetBulder {
	return &WallGetBulder{api.Params{}}
}

// OwnerID ID of the user or community that owns the wall. By default, current user ID. Use a negative value to designate a community ID.
func (b *WallGetBulder) OwnerID(v int) {
	b.Params["owner_id"] = v
}

// Domain User or community short address.
func (b *WallGetBulder) Domain(v string) {
	b.Params["domain"] = v
}

// Offset Offset needed to return a specific subset of posts.
func (b *WallGetBulder) Offset(v int) {
	b.Params["offset"] = v
}

// Count Number of posts to return (maximum 100).
func (b *WallGetBulder) Count(v int) {
	b.Params["count"] = v
}

// Filter Filter to apply: 'owner' — posts by the wall owner, 'others' — posts by someone else, 'all' — posts by the wall owner and others (default), 'postponed' — timed posts (only available for calls with an 'access_token'), 'suggests' — suggested posts on a community wall
func (b *WallGetBulder) Filter(v string) {
	b.Params["filter"] = v
}

// Extended '1' — to return 'wall', 'profiles', and 'groups' fields, '0' — to return no additional fields (default)
func (b *WallGetBulder) Extended(v bool) {
	b.Params["extended"] = v
}

// Fields parameter
func (b *WallGetBulder) Fields(v []string) {
	b.Params["fields"] = v
}

// WallGetByIDBulder builder
//
// Returns a list of posts from user or community walls by their IDs.
//
// https://vk.com/dev/wall.getById
type WallGetByIDBulder struct {
	api.Params
}

// NewWallGetByIDBulder func
func NewWallGetByIDBulder() *WallGetByIDBulder {
	return &WallGetByIDBulder{api.Params{}}
}

// Posts User or community IDs and post IDs, separated by underscores. Use a negative value to designate a community ID. Example: "93388_21539,93388_20904,2943_4276,-1_1"
func (b *WallGetByIDBulder) Posts(v []string) {
	b.Params["posts"] = v
}

// Extended '1' — to return user and community objects needed to display posts, '0' — no additional fields are returned (default)
func (b *WallGetByIDBulder) Extended(v bool) {
	b.Params["extended"] = v
}

// CopyHistoryDepth Sets the number of parent elements to include in the array 'copy_history' that is returned if the post is a repost from another wall.
func (b *WallGetByIDBulder) CopyHistoryDepth(v int) {
	b.Params["copy_history_depth"] = v
}

// Fields parameter
func (b *WallGetByIDBulder) Fields(v []string) {
	b.Params["fields"] = v
}

// WallGetCommentsBulder builder
//
// Returns a list of comments on a post on a user wall or community wall.
//
// https://vk.com/dev/wall.getComments
type WallGetCommentsBulder struct {
	api.Params
}

// NewWallGetCommentsBulder func
func NewWallGetCommentsBulder() *WallGetCommentsBulder {
	return &WallGetCommentsBulder{api.Params{}}
}

// OwnerID User ID or community ID. Use a negative value to designate a community ID.
func (b *WallGetCommentsBulder) OwnerID(v int) {
	b.Params["owner_id"] = v
}

// PostID Post ID.
func (b *WallGetCommentsBulder) PostID(v int) {
	b.Params["post_id"] = v
}

// NeedLikes '1' — to return the 'likes' field, '0' — not to return the 'likes' field (default)
func (b *WallGetCommentsBulder) NeedLikes(v bool) {
	b.Params["need_likes"] = v
}

// StartCommentID parameter
func (b *WallGetCommentsBulder) StartCommentID(v int) {
	b.Params["start_comment_id"] = v
}

// Offset Offset needed to return a specific subset of comments.
func (b *WallGetCommentsBulder) Offset(v int) {
	b.Params["offset"] = v
}

// Count Number of comments to return (maximum 100).
func (b *WallGetCommentsBulder) Count(v int) {
	b.Params["count"] = v
}

// Sort Sort order: 'asc' — chronological, 'desc' — reverse chronological
func (b *WallGetCommentsBulder) Sort(v string) {
	b.Params["sort"] = v
}

// PreviewLength Number of characters at which to truncate comments when previewed. By default, '90'. Specify '0' if you do not want to truncate comments.
func (b *WallGetCommentsBulder) PreviewLength(v int) {
	b.Params["preview_length"] = v
}

// Extended parameter
func (b *WallGetCommentsBulder) Extended(v bool) {
	b.Params["extended"] = v
}

// Fields parameter
func (b *WallGetCommentsBulder) Fields(v []string) {
	b.Params["fields"] = v
}

// CommentID Comment ID.
func (b *WallGetCommentsBulder) CommentID(v int) {
	b.Params["comment_id"] = v
}

// ThreadItemsCount Count items in threads.
func (b *WallGetCommentsBulder) ThreadItemsCount(v int) {
	b.Params["thread_items_count"] = v
}

// WallGetRepostsBulder builder
//
// Returns information about reposts of a post on user wall or community wall.
//
// https://vk.com/dev/wall.getReposts
type WallGetRepostsBulder struct {
	api.Params
}

// NewWallGetRepostsBulder func
func NewWallGetRepostsBulder() *WallGetRepostsBulder {
	return &WallGetRepostsBulder{api.Params{}}
}

// OwnerID User ID or community ID. By default, current user ID. Use a negative value to designate a community ID.
func (b *WallGetRepostsBulder) OwnerID(v int) {
	b.Params["owner_id"] = v
}

// PostID Post ID.
func (b *WallGetRepostsBulder) PostID(v int) {
	b.Params["post_id"] = v
}

// Offset Offset needed to return a specific subset of reposts.
func (b *WallGetRepostsBulder) Offset(v int) {
	b.Params["offset"] = v
}

// Count Number of reposts to return.
func (b *WallGetRepostsBulder) Count(v int) {
	b.Params["count"] = v
}

// WallOpenCommentsBulder builder
//
// https://vk.com/dev/wall.openComments
type WallOpenCommentsBulder struct {
	api.Params
}

// NewWallOpenCommentsBulder func
func NewWallOpenCommentsBulder() *WallOpenCommentsBulder {
	return &WallOpenCommentsBulder{api.Params{}}
}

// OwnerID parameter
func (b *WallOpenCommentsBulder) OwnerID(v int) {
	b.Params["owner_id"] = v
}

// PostID parameter
func (b *WallOpenCommentsBulder) PostID(v int) {
	b.Params["post_id"] = v
}

// WallPinBulder builder
//
// Pins the post on wall.
//
// https://vk.com/dev/wall.pin
type WallPinBulder struct {
	api.Params
}

// NewWallPinBulder func
func NewWallPinBulder() *WallPinBulder {
	return &WallPinBulder{api.Params{}}
}

// OwnerID ID of the user or community that owns the wall. By default, current user ID. Use a negative value to designate a community ID.
func (b *WallPinBulder) OwnerID(v int) {
	b.Params["owner_id"] = v
}

// PostID Post ID.
func (b *WallPinBulder) PostID(v int) {
	b.Params["post_id"] = v
}

// WallPostBulder builder
//
// Adds a new post on a user wall or community wall. Can also be used to publish suggested or scheduled posts.
//
// https://vk.com/dev/wall.post
type WallPostBulder struct {
	api.Params
}

// NewWallPostBulder func
func NewWallPostBulder() *WallPostBulder {
	return &WallPostBulder{api.Params{}}
}

// OwnerID User ID or community ID. Use a negative value to designate a community ID.
func (b *WallPostBulder) OwnerID(v int) {
	b.Params["owner_id"] = v
}

// FriendsOnly '1' — post will be available to friends only, '0' — post will be available to all users (default)
func (b *WallPostBulder) FriendsOnly(v bool) {
	b.Params["friends_only"] = v
}

// FromGroup For a community: '1' — post will be published by the community, '0' — post will be published by the user (default)
func (b *WallPostBulder) FromGroup(v bool) {
	b.Params["from_group"] = v
}

// Message (Required if 'attachments' is not set.) Text of the post.
func (b *WallPostBulder) Message(v string) {
	b.Params["message"] = v
}

// Attachments (Required if 'message' is not set.) List of objects attached to the post, in the following format: "<owner_id>_<media_id>,<owner_id>_<media_id>", '' — Type of media attachment: 'photo' — photo, 'video' — video, 'audio' — audio, 'doc' — document, 'page' — wiki-page, 'note' — note, 'poll' — poll, 'album' — photo album, '<owner_id>' — ID of the media application owner. '<media_id>' — Media application ID. Example: "photo100172_166443618,photo66748_265827614", May contain a link to an external page to include in the post. Example: "photo66748_265827614,http://habrahabr.ru", "NOTE: If more than one link is being attached, an error will be thrown."
func (b *WallPostBulder) Attachments(v []string) {
	b.Params["attachments"] = v
}

// Services List of services or websites the update will be exported to, if the user has so requested. Sample values: 'twitter', 'facebook'.
func (b *WallPostBulder) Services(v string) {
	b.Params["services"] = v
}

// Signed Only for posts in communities with 'from_group' set to '1': '1' — post will be signed with the name of the posting user, '0' — post will not be signed (default)
func (b *WallPostBulder) Signed(v bool) {
	b.Params["signed"] = v
}

// PublishDate Publication date (in Unix time). If used, posting will be delayed until the set time.
func (b *WallPostBulder) PublishDate(v int) {
	b.Params["publish_date"] = v
}

// Lat Geographical latitude of a check-in, in degrees (from -90 to 90).
func (b *WallPostBulder) Lat(v float64) {
	b.Params["lat"] = v
}

// Long Geographical longitude of a check-in, in degrees (from -180 to 180).
func (b *WallPostBulder) Long(v float64) {
	b.Params["long"] = v
}

// PlaceID ID of the location where the user was tagged.
func (b *WallPostBulder) PlaceID(v int) {
	b.Params["place_id"] = v
}

// PostID Post ID. Used for publishing of scheduled and suggested posts.
func (b *WallPostBulder) PostID(v int) {
	b.Params["post_id"] = v
}

// GUID parameter
func (b *WallPostBulder) GUID(v string) {
	b.Params["guid"] = v
}

// MarkAsAds parameter
func (b *WallPostBulder) MarkAsAds(v bool) {
	b.Params["mark_as_ads"] = v
}

// CloseComments parameter
func (b *WallPostBulder) CloseComments(v bool) {
	b.Params["close_comments"] = v
}

// MuteNotifications parameter
func (b *WallPostBulder) MuteNotifications(v bool) {
	b.Params["mute_notifications"] = v
}

// WallPostAdsStealthBulder builder
//
// Allows to create hidden post which will not be shown on the community's wall and can be used for creating an ad with type "Community post".
//
// https://vk.com/dev/wall.postAdsStealth
type WallPostAdsStealthBulder struct {
	api.Params
}

// NewWallPostAdsStealthBulder func
func NewWallPostAdsStealthBulder() *WallPostAdsStealthBulder {
	return &WallPostAdsStealthBulder{api.Params{}}
}

// OwnerID User ID or community ID. Use a negative value to designate a community ID.
func (b *WallPostAdsStealthBulder) OwnerID(v int) {
	b.Params["owner_id"] = v
}

// Message (Required if 'attachments' is not set.) Text of the post.
func (b *WallPostAdsStealthBulder) Message(v string) {
	b.Params["message"] = v
}

// Attachments (Required if 'message' is not set.) List of objects attached to the post, in the following format: "<owner_id>_<media_id>,<owner_id>_<media_id>", '' — Type of media attachment: 'photo' — photo, 'video' — video, 'audio' — audio, 'doc' — document, 'page' — wiki-page, 'note' — note, 'poll' — poll, 'album' — photo album, '<owner_id>' — ID of the media application owner. '<media_id>' — Media application ID. Example: "photo100172_166443618,photo66748_265827614", May contain a link to an external page to include in the post. Example: "photo66748_265827614,http://habrahabr.ru", "NOTE: If more than one link is being attached, an error will be thrown."
func (b *WallPostAdsStealthBulder) Attachments(v []string) {
	b.Params["attachments"] = v
}

// Signed Only for posts in communities with 'from_group' set to '1': '1' — post will be signed with the name of the posting user, '0' — post will not be signed (default)
func (b *WallPostAdsStealthBulder) Signed(v bool) {
	b.Params["signed"] = v
}

// Lat Geographical latitude of a check-in, in degrees (from -90 to 90).
func (b *WallPostAdsStealthBulder) Lat(v float64) {
	b.Params["lat"] = v
}

// Long Geographical longitude of a check-in, in degrees (from -180 to 180).
func (b *WallPostAdsStealthBulder) Long(v float64) {
	b.Params["long"] = v
}

// PlaceID ID of the location where the user was tagged.
func (b *WallPostAdsStealthBulder) PlaceID(v int) {
	b.Params["place_id"] = v
}

// GUID Unique identifier to avoid duplication the same post.
func (b *WallPostAdsStealthBulder) GUID(v string) {
	b.Params["guid"] = v
}

// LinkButton Link button ID
func (b *WallPostAdsStealthBulder) LinkButton(v string) {
	b.Params["link_button"] = v
}

// LinkTitle Link title
func (b *WallPostAdsStealthBulder) LinkTitle(v string) {
	b.Params["link_title"] = v
}

// LinkImage Link image url
func (b *WallPostAdsStealthBulder) LinkImage(v string) {
	b.Params["link_image"] = v
}

// LinkVideo Link video ID in format "<owner_id>_<media_id>"
func (b *WallPostAdsStealthBulder) LinkVideo(v string) {
	b.Params["link_video"] = v
}

// WallReportCommentBulder builder
//
// Reports (submits a complaint about) a comment on a post on a user wall or community wall.
//
// https://vk.com/dev/wall.reportComment
type WallReportCommentBulder struct {
	api.Params
}

// NewWallReportCommentBulder func
func NewWallReportCommentBulder() *WallReportCommentBulder {
	return &WallReportCommentBulder{api.Params{}}
}

// OwnerID ID of the user or community that owns the wall.
func (b *WallReportCommentBulder) OwnerID(v int) {
	b.Params["owner_id"] = v
}

// CommentID Comment ID.
func (b *WallReportCommentBulder) CommentID(v int) {
	b.Params["comment_id"] = v
}

// Reason Reason for the complaint: '0' – spam, '1' – child pornography, '2' – extremism, '3' – violence, '4' – drug propaganda, '5' – adult material, '6' – insult, abuse
func (b *WallReportCommentBulder) Reason(v int) {
	b.Params["reason"] = v
}

// WallReportPostBulder builder
//
// Reports (submits a complaint about) a post on a user wall or community wall.
//
// https://vk.com/dev/wall.reportPost
type WallReportPostBulder struct {
	api.Params
}

// NewWallReportPostBulder func
func NewWallReportPostBulder() *WallReportPostBulder {
	return &WallReportPostBulder{api.Params{}}
}

// OwnerID ID of the user or community that owns the wall.
func (b *WallReportPostBulder) OwnerID(v int) {
	b.Params["owner_id"] = v
}

// PostID Post ID.
func (b *WallReportPostBulder) PostID(v int) {
	b.Params["post_id"] = v
}

// Reason Reason for the complaint: '0' – spam, '1' – child pornography, '2' – extremism, '3' – violence, '4' – drug propaganda, '5' – adult material, '6' – insult, abuse
func (b *WallReportPostBulder) Reason(v int) {
	b.Params["reason"] = v
}

// WallRepostBulder builder
//
// Reposts (copies) an object to a user wall or community wall.
//
// https://vk.com/dev/wall.repost
type WallRepostBulder struct {
	api.Params
}

// NewWallRepostBulder func
func NewWallRepostBulder() *WallRepostBulder {
	return &WallRepostBulder{api.Params{}}
}

// Object ID of the object to be reposted on the wall. Example: "wall66748_3675"
func (b *WallRepostBulder) Object(v string) {
	b.Params["object"] = v
}

// Message Comment to be added along with the reposted object.
func (b *WallRepostBulder) Message(v string) {
	b.Params["message"] = v
}

// GroupID Target community ID when reposting to a community.
func (b *WallRepostBulder) GroupID(v int) {
	b.Params["group_id"] = v
}

// MarkAsAds parameter
func (b *WallRepostBulder) MarkAsAds(v bool) {
	b.Params["mark_as_ads"] = v
}

// MuteNotifications parameter
func (b *WallRepostBulder) MuteNotifications(v bool) {
	b.Params["mute_notifications"] = v
}

// WallRestoreBulder builder
//
// Restores a post deleted from a user wall or community wall.
//
// https://vk.com/dev/wall.restore
type WallRestoreBulder struct {
	api.Params
}

// NewWallRestoreBulder func
func NewWallRestoreBulder() *WallRestoreBulder {
	return &WallRestoreBulder{api.Params{}}
}

// OwnerID User ID or community ID from whose wall the post was deleted. Use a negative value to designate a community ID.
func (b *WallRestoreBulder) OwnerID(v int) {
	b.Params["owner_id"] = v
}

// PostID ID of the post to be restored.
func (b *WallRestoreBulder) PostID(v int) {
	b.Params["post_id"] = v
}

// WallRestoreCommentBulder builder
//
// Restores a comment deleted from a user wall or community wall.
//
// https://vk.com/dev/wall.restoreComment
type WallRestoreCommentBulder struct {
	api.Params
}

// NewWallRestoreCommentBulder func
func NewWallRestoreCommentBulder() *WallRestoreCommentBulder {
	return &WallRestoreCommentBulder{api.Params{}}
}

// OwnerID User ID or community ID. Use a negative value to designate a community ID.
func (b *WallRestoreCommentBulder) OwnerID(v int) {
	b.Params["owner_id"] = v
}

// CommentID Comment ID.
func (b *WallRestoreCommentBulder) CommentID(v int) {
	b.Params["comment_id"] = v
}

// WallSearchBulder builder
//
// Allows to search posts on user or community walls.
//
// https://vk.com/dev/wall.search
type WallSearchBulder struct {
	api.Params
}

// NewWallSearchBulder func
func NewWallSearchBulder() *WallSearchBulder {
	return &WallSearchBulder{api.Params{}}
}

// OwnerID user or community id. "Remember that for a community 'owner_id' must be negative."
func (b *WallSearchBulder) OwnerID(v int) {
	b.Params["owner_id"] = v
}

// Domain user or community screen name.
func (b *WallSearchBulder) Domain(v string) {
	b.Params["domain"] = v
}

// Query search query string.
func (b *WallSearchBulder) Query(v string) {
	b.Params["query"] = v
}

// OwnersOnly '1' – returns only page owner's posts.
func (b *WallSearchBulder) OwnersOnly(v bool) {
	b.Params["owners_only"] = v
}

// Count count of posts to return.
func (b *WallSearchBulder) Count(v int) {
	b.Params["count"] = v
}

// Offset Offset needed to return a specific subset of posts.
func (b *WallSearchBulder) Offset(v int) {
	b.Params["offset"] = v
}

// Extended show extended post info.
func (b *WallSearchBulder) Extended(v bool) {
	b.Params["extended"] = v
}

// Fields parameter
func (b *WallSearchBulder) Fields(v []string) {
	b.Params["fields"] = v
}

// WallUnpinBulder builder
//
// Unpins the post on wall.
//
// https://vk.com/dev/wall.unpin
type WallUnpinBulder struct {
	api.Params
}

// NewWallUnpinBulder func
func NewWallUnpinBulder() *WallUnpinBulder {
	return &WallUnpinBulder{api.Params{}}
}

// OwnerID ID of the user or community that owns the wall. By default, current user ID. Use a negative value to designate a community ID.
func (b *WallUnpinBulder) OwnerID(v int) {
	b.Params["owner_id"] = v
}

// PostID Post ID.
func (b *WallUnpinBulder) PostID(v int) {
	b.Params["post_id"] = v
}
