package params // import "github.com/SevereCloud/vksdk/api/params"

import (
	"github.com/SevereCloud/vksdk/api"
)

// WallCloseCommentsBuilder builder
//
// https://vk.com/dev/wall.closeComments
type WallCloseCommentsBuilder struct {
	api.Params
}

// NewWallCloseCommentsBuilder func
func NewWallCloseCommentsBuilder() *WallCloseCommentsBuilder {
	return &WallCloseCommentsBuilder{api.Params{}}
}

// OwnerID parameter
func (b *WallCloseCommentsBuilder) OwnerID(v int) {
	b.Params["owner_id"] = v
}

// PostID parameter
func (b *WallCloseCommentsBuilder) PostID(v int) {
	b.Params["post_id"] = v
}

// WallCreateCommentBuilder builder
//
// Adds a comment to a post on a user wall or community wall.
//
// https://vk.com/dev/wall.createComment
type WallCreateCommentBuilder struct {
	api.Params
}

// NewWallCreateCommentBuilder func
func NewWallCreateCommentBuilder() *WallCreateCommentBuilder {
	return &WallCreateCommentBuilder{api.Params{}}
}

// OwnerID User ID or community ID. Use a negative value to designate a community ID.
func (b *WallCreateCommentBuilder) OwnerID(v int) {
	b.Params["owner_id"] = v
}

// PostID Post ID.
func (b *WallCreateCommentBuilder) PostID(v int) {
	b.Params["post_id"] = v
}

// FromGroup Group ID.
func (b *WallCreateCommentBuilder) FromGroup(v int) {
	b.Params["from_group"] = v
}

// Message (Required if 'attachments' is not set.) Text of the comment.
func (b *WallCreateCommentBuilder) Message(v string) {
	b.Params["message"] = v
}

// ReplyToComment ID of comment to reply.
func (b *WallCreateCommentBuilder) ReplyToComment(v int) {
	b.Params["reply_to_comment"] = v
}

// Attachments (Required if 'message' is not set.) List of media objects attached to the comment, in the following format: "<owner_id>_<media_id>,<owner_id>_<media_id>", '' — Type of media ojbect: 'photo' — photo, 'video' — video, 'audio' — audio, 'doc' — document, '<owner_id>' — ID of the media owner. '<media_id>' — Media ID. For example: "photo100172_166443618,photo66748_265827614"
func (b *WallCreateCommentBuilder) Attachments(v []string) {
	b.Params["attachments"] = v
}

// StickerID Sticker ID.
func (b *WallCreateCommentBuilder) StickerID(v int) {
	b.Params["sticker_id"] = v
}

// GUID Unique identifier to avoid repeated comments.
func (b *WallCreateCommentBuilder) GUID(v string) {
	b.Params["guid"] = v
}

// WallDeleteBuilder builder
//
// Deletes a post from a user wall or community wall.
//
// https://vk.com/dev/wall.delete
type WallDeleteBuilder struct {
	api.Params
}

// NewWallDeleteBuilder func
func NewWallDeleteBuilder() *WallDeleteBuilder {
	return &WallDeleteBuilder{api.Params{}}
}

// OwnerID User ID or community ID. Use a negative value to designate a community ID.
func (b *WallDeleteBuilder) OwnerID(v int) {
	b.Params["owner_id"] = v
}

// PostID ID of the post to be deleted.
func (b *WallDeleteBuilder) PostID(v int) {
	b.Params["post_id"] = v
}

// WallDeleteCommentBuilder builder
//
// Deletes a comment on a post on a user wall or community wall.
//
// https://vk.com/dev/wall.deleteComment
type WallDeleteCommentBuilder struct {
	api.Params
}

// NewWallDeleteCommentBuilder func
func NewWallDeleteCommentBuilder() *WallDeleteCommentBuilder {
	return &WallDeleteCommentBuilder{api.Params{}}
}

// OwnerID User ID or community ID. Use a negative value to designate a community ID.
func (b *WallDeleteCommentBuilder) OwnerID(v int) {
	b.Params["owner_id"] = v
}

// CommentID Comment ID.
func (b *WallDeleteCommentBuilder) CommentID(v int) {
	b.Params["comment_id"] = v
}

// WallEditBuilder builder
//
// Edits a post on a user wall or community wall.
//
// https://vk.com/dev/wall.edit
type WallEditBuilder struct {
	api.Params
}

// NewWallEditBuilder func
func NewWallEditBuilder() *WallEditBuilder {
	return &WallEditBuilder{api.Params{}}
}

// OwnerID User ID or community ID. Use a negative value to designate a community ID.
func (b *WallEditBuilder) OwnerID(v int) {
	b.Params["owner_id"] = v
}

// PostID parameter
func (b *WallEditBuilder) PostID(v int) {
	b.Params["post_id"] = v
}

// FriendsOnly parameter
func (b *WallEditBuilder) FriendsOnly(v bool) {
	b.Params["friends_only"] = v
}

// Message (Required if 'attachments' is not set.) Text of the post.
func (b *WallEditBuilder) Message(v string) {
	b.Params["message"] = v
}

// Attachments (Required if 'message' is not set.) List of objects attached to the post, in the following format: "<owner_id>_<media_id>,<owner_id>_<media_id>", '' — Type of media attachment: 'photo' — photo, 'video' — video, 'audio' — audio, 'doc' — document, '<owner_id>' — ID of the media application owner. '<media_id>' — Media application ID. Example: "photo100172_166443618,photo66748_265827614", May contain a link to an external page to include in the post. Example: "photo66748_265827614,http://habrahabr.ru", "NOTE: If more than one link is being attached, an error is thrown."
func (b *WallEditBuilder) Attachments(v []string) {
	b.Params["attachments"] = v
}

// Services parameter
func (b *WallEditBuilder) Services(v string) {
	b.Params["services"] = v
}

// Signed parameter
func (b *WallEditBuilder) Signed(v bool) {
	b.Params["signed"] = v
}

// PublishDate parameter
func (b *WallEditBuilder) PublishDate(v int) {
	b.Params["publish_date"] = v
}

// Lat parameter
func (b *WallEditBuilder) Lat(v float64) {
	b.Params["lat"] = v
}

// Long parameter
func (b *WallEditBuilder) Long(v float64) {
	b.Params["long"] = v
}

// PlaceID parameter
func (b *WallEditBuilder) PlaceID(v int) {
	b.Params["place_id"] = v
}

// MarkAsAds parameter
func (b *WallEditBuilder) MarkAsAds(v bool) {
	b.Params["mark_as_ads"] = v
}

// CloseComments parameter
func (b *WallEditBuilder) CloseComments(v bool) {
	b.Params["close_comments"] = v
}

// PosterBkgID parameter
func (b *WallEditBuilder) PosterBkgID(v int) {
	b.Params["poster_bkg_id"] = v
}

// PosterBkgOwnerID parameter
func (b *WallEditBuilder) PosterBkgOwnerID(v int) {
	b.Params["poster_bkg_owner_id"] = v
}

// PosterBkgAccessHash parameter
func (b *WallEditBuilder) PosterBkgAccessHash(v string) {
	b.Params["poster_bkg_access_hash"] = v
}

// WallEditAdsStealthBuilder builder
//
// Allows to edit hidden post.
//
// https://vk.com/dev/wall.editAdsStealth
type WallEditAdsStealthBuilder struct {
	api.Params
}

// NewWallEditAdsStealthBuilder func
func NewWallEditAdsStealthBuilder() *WallEditAdsStealthBuilder {
	return &WallEditAdsStealthBuilder{api.Params{}}
}

// OwnerID User ID or community ID. Use a negative value to designate a community ID.
func (b *WallEditAdsStealthBuilder) OwnerID(v int) {
	b.Params["owner_id"] = v
}

// PostID Post ID. Used for publishing of scheduled and suggested posts.
func (b *WallEditAdsStealthBuilder) PostID(v int) {
	b.Params["post_id"] = v
}

// Message (Required if 'attachments' is not set.) Text of the post.
func (b *WallEditAdsStealthBuilder) Message(v string) {
	b.Params["message"] = v
}

// Attachments (Required if 'message' is not set.) List of objects attached to the post, in the following format: "<owner_id>_<media_id>,<owner_id>_<media_id>", '' — Type of media attachment: 'photo' — photo, 'video' — video, 'audio' — audio, 'doc' — document, 'page' — wiki-page, 'note' — note, 'poll' — poll, 'album' — photo album, '<owner_id>' — ID of the media application owner. '<media_id>' — Media application ID. Example: "photo100172_166443618,photo66748_265827614", May contain a link to an external page to include in the post. Example: "photo66748_265827614,http://habrahabr.ru", "NOTE: If more than one link is being attached, an error will be thrown."
func (b *WallEditAdsStealthBuilder) Attachments(v []string) {
	b.Params["attachments"] = v
}

// Signed Only for posts in communities with 'from_group' set to '1': '1' — post will be signed with the name of the posting user, '0' — post will not be signed (default)
func (b *WallEditAdsStealthBuilder) Signed(v bool) {
	b.Params["signed"] = v
}

// Lat Geographical latitude of a check-in, in degrees (from -90 to 90).
func (b *WallEditAdsStealthBuilder) Lat(v float64) {
	b.Params["lat"] = v
}

// Long Geographical longitude of a check-in, in degrees (from -180 to 180).
func (b *WallEditAdsStealthBuilder) Long(v float64) {
	b.Params["long"] = v
}

// PlaceID ID of the location where the user was tagged.
func (b *WallEditAdsStealthBuilder) PlaceID(v int) {
	b.Params["place_id"] = v
}

// LinkButton Link button ID
func (b *WallEditAdsStealthBuilder) LinkButton(v string) {
	b.Params["link_button"] = v
}

// LinkTitle Link title
func (b *WallEditAdsStealthBuilder) LinkTitle(v string) {
	b.Params["link_title"] = v
}

// LinkImage Link image url
func (b *WallEditAdsStealthBuilder) LinkImage(v string) {
	b.Params["link_image"] = v
}

// LinkVideo Link video ID in format "<owner_id>_<media_id>"
func (b *WallEditAdsStealthBuilder) LinkVideo(v string) {
	b.Params["link_video"] = v
}

// WallEditCommentBuilder builder
//
// Edits a comment on a user wall or community wall.
//
// https://vk.com/dev/wall.editComment
type WallEditCommentBuilder struct {
	api.Params
}

// NewWallEditCommentBuilder func
func NewWallEditCommentBuilder() *WallEditCommentBuilder {
	return &WallEditCommentBuilder{api.Params{}}
}

// OwnerID User ID or community ID. Use a negative value to designate a community ID.
func (b *WallEditCommentBuilder) OwnerID(v int) {
	b.Params["owner_id"] = v
}

// CommentID Comment ID.
func (b *WallEditCommentBuilder) CommentID(v int) {
	b.Params["comment_id"] = v
}

// Message New comment text.
func (b *WallEditCommentBuilder) Message(v string) {
	b.Params["message"] = v
}

// Attachments List of objects attached to the comment, in the following format: , "<owner_id>_<media_id>,<owner_id>_<media_id>", '' — Type of media attachment: 'photo' — photo, 'video' — video, 'audio' — audio, 'doc' — document, '<owner_id>' — ID of the media attachment owner. '<media_id>' — Media attachment ID. For example: "photo100172_166443618,photo66748_265827614"
func (b *WallEditCommentBuilder) Attachments(v []string) {
	b.Params["attachments"] = v
}

// WallGetBuilder builder
//
// Returns a list of posts on a user wall or community wall.
//
// https://vk.com/dev/wall.get
type WallGetBuilder struct {
	api.Params
}

// NewWallGetBuilder func
func NewWallGetBuilder() *WallGetBuilder {
	return &WallGetBuilder{api.Params{}}
}

// OwnerID ID of the user or community that owns the wall. By default, current user ID. Use a negative value to designate a community ID.
func (b *WallGetBuilder) OwnerID(v int) {
	b.Params["owner_id"] = v
}

// Domain User or community short address.
func (b *WallGetBuilder) Domain(v string) {
	b.Params["domain"] = v
}

// Offset Offset needed to return a specific subset of posts.
func (b *WallGetBuilder) Offset(v int) {
	b.Params["offset"] = v
}

// Count Number of posts to return (maximum 100).
func (b *WallGetBuilder) Count(v int) {
	b.Params["count"] = v
}

// Filter Filter to apply: 'owner' — posts by the wall owner, 'others' — posts by someone else, 'all' — posts by the wall owner and others (default), 'postponed' — timed posts (only available for calls with an 'access_token'), 'suggests' — suggested posts on a community wall
func (b *WallGetBuilder) Filter(v string) {
	b.Params["filter"] = v
}

// Extended '1' — to return 'wall', 'profiles', and 'groups' fields, '0' — to return no additional fields (default)
func (b *WallGetBuilder) Extended(v bool) {
	b.Params["extended"] = v
}

// Fields parameter
func (b *WallGetBuilder) Fields(v []string) {
	b.Params["fields"] = v
}

// WallGetByIDBuilder builder
//
// Returns a list of posts from user or community walls by their IDs.
//
// https://vk.com/dev/wall.getById
type WallGetByIDBuilder struct {
	api.Params
}

// NewWallGetByIDBuilder func
func NewWallGetByIDBuilder() *WallGetByIDBuilder {
	return &WallGetByIDBuilder{api.Params{}}
}

// Posts User or community IDs and post IDs, separated by underscores. Use a negative value to designate a community ID. Example: "93388_21539,93388_20904,2943_4276,-1_1"
func (b *WallGetByIDBuilder) Posts(v []string) {
	b.Params["posts"] = v
}

// Extended '1' — to return user and community objects needed to display posts, '0' — no additional fields are returned (default)
func (b *WallGetByIDBuilder) Extended(v bool) {
	b.Params["extended"] = v
}

// CopyHistoryDepth Sets the number of parent elements to include in the array 'copy_history' that is returned if the post is a repost from another wall.
func (b *WallGetByIDBuilder) CopyHistoryDepth(v int) {
	b.Params["copy_history_depth"] = v
}

// Fields parameter
func (b *WallGetByIDBuilder) Fields(v []string) {
	b.Params["fields"] = v
}

// WallGetCommentsBuilder builder
//
// Returns a list of comments on a post on a user wall or community wall.
//
// https://vk.com/dev/wall.getComments
type WallGetCommentsBuilder struct {
	api.Params
}

// NewWallGetCommentsBuilder func
func NewWallGetCommentsBuilder() *WallGetCommentsBuilder {
	return &WallGetCommentsBuilder{api.Params{}}
}

// OwnerID User ID or community ID. Use a negative value to designate a community ID.
func (b *WallGetCommentsBuilder) OwnerID(v int) {
	b.Params["owner_id"] = v
}

// PostID Post ID.
func (b *WallGetCommentsBuilder) PostID(v int) {
	b.Params["post_id"] = v
}

// NeedLikes '1' — to return the 'likes' field, '0' — not to return the 'likes' field (default)
func (b *WallGetCommentsBuilder) NeedLikes(v bool) {
	b.Params["need_likes"] = v
}

// StartCommentID parameter
func (b *WallGetCommentsBuilder) StartCommentID(v int) {
	b.Params["start_comment_id"] = v
}

// Offset Offset needed to return a specific subset of comments.
func (b *WallGetCommentsBuilder) Offset(v int) {
	b.Params["offset"] = v
}

// Count Number of comments to return (maximum 100).
func (b *WallGetCommentsBuilder) Count(v int) {
	b.Params["count"] = v
}

// Sort Sort order: 'asc' — chronological, 'desc' — reverse chronological
func (b *WallGetCommentsBuilder) Sort(v string) {
	b.Params["sort"] = v
}

// PreviewLength Number of characters at which to truncate comments when previewed. By default, '90'. Specify '0' if you do not want to truncate comments.
func (b *WallGetCommentsBuilder) PreviewLength(v int) {
	b.Params["preview_length"] = v
}

// Extended parameter
func (b *WallGetCommentsBuilder) Extended(v bool) {
	b.Params["extended"] = v
}

// Fields parameter
func (b *WallGetCommentsBuilder) Fields(v []string) {
	b.Params["fields"] = v
}

// CommentID Comment ID.
func (b *WallGetCommentsBuilder) CommentID(v int) {
	b.Params["comment_id"] = v
}

// ThreadItemsCount Count items in threads.
func (b *WallGetCommentsBuilder) ThreadItemsCount(v int) {
	b.Params["thread_items_count"] = v
}

// WallGetRepostsBuilder builder
//
// Returns information about reposts of a post on user wall or community wall.
//
// https://vk.com/dev/wall.getReposts
type WallGetRepostsBuilder struct {
	api.Params
}

// NewWallGetRepostsBuilder func
func NewWallGetRepostsBuilder() *WallGetRepostsBuilder {
	return &WallGetRepostsBuilder{api.Params{}}
}

// OwnerID User ID or community ID. By default, current user ID. Use a negative value to designate a community ID.
func (b *WallGetRepostsBuilder) OwnerID(v int) {
	b.Params["owner_id"] = v
}

// PostID Post ID.
func (b *WallGetRepostsBuilder) PostID(v int) {
	b.Params["post_id"] = v
}

// Offset Offset needed to return a specific subset of reposts.
func (b *WallGetRepostsBuilder) Offset(v int) {
	b.Params["offset"] = v
}

// Count Number of reposts to return.
func (b *WallGetRepostsBuilder) Count(v int) {
	b.Params["count"] = v
}

// WallOpenCommentsBuilder builder
//
// https://vk.com/dev/wall.openComments
type WallOpenCommentsBuilder struct {
	api.Params
}

// NewWallOpenCommentsBuilder func
func NewWallOpenCommentsBuilder() *WallOpenCommentsBuilder {
	return &WallOpenCommentsBuilder{api.Params{}}
}

// OwnerID parameter
func (b *WallOpenCommentsBuilder) OwnerID(v int) {
	b.Params["owner_id"] = v
}

// PostID parameter
func (b *WallOpenCommentsBuilder) PostID(v int) {
	b.Params["post_id"] = v
}

// WallPinBuilder builder
//
// Pins the post on wall.
//
// https://vk.com/dev/wall.pin
type WallPinBuilder struct {
	api.Params
}

// NewWallPinBuilder func
func NewWallPinBuilder() *WallPinBuilder {
	return &WallPinBuilder{api.Params{}}
}

// OwnerID ID of the user or community that owns the wall. By default, current user ID. Use a negative value to designate a community ID.
func (b *WallPinBuilder) OwnerID(v int) {
	b.Params["owner_id"] = v
}

// PostID Post ID.
func (b *WallPinBuilder) PostID(v int) {
	b.Params["post_id"] = v
}

// WallPostBuilder builder
//
// Adds a new post on a user wall or community wall. Can also be used to publish suggested or scheduled posts.
//
// https://vk.com/dev/wall.post
type WallPostBuilder struct {
	api.Params
}

// NewWallPostBuilder func
func NewWallPostBuilder() *WallPostBuilder {
	return &WallPostBuilder{api.Params{}}
}

// OwnerID User ID or community ID. Use a negative value to designate a community ID.
func (b *WallPostBuilder) OwnerID(v int) {
	b.Params["owner_id"] = v
}

// FriendsOnly '1' — post will be available to friends only, '0' — post will be available to all users (default)
func (b *WallPostBuilder) FriendsOnly(v bool) {
	b.Params["friends_only"] = v
}

// FromGroup For a community: '1' — post will be published by the community, '0' — post will be published by the user (default)
func (b *WallPostBuilder) FromGroup(v bool) {
	b.Params["from_group"] = v
}

// Message (Required if 'attachments' is not set.) Text of the post.
func (b *WallPostBuilder) Message(v string) {
	b.Params["message"] = v
}

// Attachments (Required if 'message' is not set.) List of objects attached to the post, in the following format: "<owner_id>_<media_id>,<owner_id>_<media_id>", '' — Type of media attachment: 'photo' — photo, 'video' — video, 'audio' — audio, 'doc' — document, 'page' — wiki-page, 'note' — note, 'poll' — poll, 'album' — photo album, '<owner_id>' — ID of the media application owner. '<media_id>' — Media application ID. Example: "photo100172_166443618,photo66748_265827614", May contain a link to an external page to include in the post. Example: "photo66748_265827614,http://habrahabr.ru", "NOTE: If more than one link is being attached, an error will be thrown."
func (b *WallPostBuilder) Attachments(v []string) {
	b.Params["attachments"] = v
}

// Services List of services or websites the update will be exported to, if the user has so requested. Sample values: 'twitter', 'facebook'.
func (b *WallPostBuilder) Services(v string) {
	b.Params["services"] = v
}

// Signed Only for posts in communities with 'from_group' set to '1': '1' — post will be signed with the name of the posting user, '0' — post will not be signed (default)
func (b *WallPostBuilder) Signed(v bool) {
	b.Params["signed"] = v
}

// PublishDate Publication date (in Unix time). If used, posting will be delayed until the set time.
func (b *WallPostBuilder) PublishDate(v int) {
	b.Params["publish_date"] = v
}

// Lat Geographical latitude of a check-in, in degrees (from -90 to 90).
func (b *WallPostBuilder) Lat(v float64) {
	b.Params["lat"] = v
}

// Long Geographical longitude of a check-in, in degrees (from -180 to 180).
func (b *WallPostBuilder) Long(v float64) {
	b.Params["long"] = v
}

// PlaceID ID of the location where the user was tagged.
func (b *WallPostBuilder) PlaceID(v int) {
	b.Params["place_id"] = v
}

// PostID Post ID. Used for publishing of scheduled and suggested posts.
func (b *WallPostBuilder) PostID(v int) {
	b.Params["post_id"] = v
}

// GUID parameter
func (b *WallPostBuilder) GUID(v string) {
	b.Params["guid"] = v
}

// MarkAsAds parameter
func (b *WallPostBuilder) MarkAsAds(v bool) {
	b.Params["mark_as_ads"] = v
}

// CloseComments parameter
func (b *WallPostBuilder) CloseComments(v bool) {
	b.Params["close_comments"] = v
}

// MuteNotifications parameter
func (b *WallPostBuilder) MuteNotifications(v bool) {
	b.Params["mute_notifications"] = v
}

// WallPostAdsStealthBuilder builder
//
// Allows to create hidden post which will not be shown on the community's wall and can be used for creating an ad with type "Community post".
//
// https://vk.com/dev/wall.postAdsStealth
type WallPostAdsStealthBuilder struct {
	api.Params
}

// NewWallPostAdsStealthBuilder func
func NewWallPostAdsStealthBuilder() *WallPostAdsStealthBuilder {
	return &WallPostAdsStealthBuilder{api.Params{}}
}

// OwnerID User ID or community ID. Use a negative value to designate a community ID.
func (b *WallPostAdsStealthBuilder) OwnerID(v int) {
	b.Params["owner_id"] = v
}

// Message (Required if 'attachments' is not set.) Text of the post.
func (b *WallPostAdsStealthBuilder) Message(v string) {
	b.Params["message"] = v
}

// Attachments (Required if 'message' is not set.) List of objects attached to the post, in the following format: "<owner_id>_<media_id>,<owner_id>_<media_id>", '' — Type of media attachment: 'photo' — photo, 'video' — video, 'audio' — audio, 'doc' — document, 'page' — wiki-page, 'note' — note, 'poll' — poll, 'album' — photo album, '<owner_id>' — ID of the media application owner. '<media_id>' — Media application ID. Example: "photo100172_166443618,photo66748_265827614", May contain a link to an external page to include in the post. Example: "photo66748_265827614,http://habrahabr.ru", "NOTE: If more than one link is being attached, an error will be thrown."
func (b *WallPostAdsStealthBuilder) Attachments(v []string) {
	b.Params["attachments"] = v
}

// Signed Only for posts in communities with 'from_group' set to '1': '1' — post will be signed with the name of the posting user, '0' — post will not be signed (default)
func (b *WallPostAdsStealthBuilder) Signed(v bool) {
	b.Params["signed"] = v
}

// Lat Geographical latitude of a check-in, in degrees (from -90 to 90).
func (b *WallPostAdsStealthBuilder) Lat(v float64) {
	b.Params["lat"] = v
}

// Long Geographical longitude of a check-in, in degrees (from -180 to 180).
func (b *WallPostAdsStealthBuilder) Long(v float64) {
	b.Params["long"] = v
}

// PlaceID ID of the location where the user was tagged.
func (b *WallPostAdsStealthBuilder) PlaceID(v int) {
	b.Params["place_id"] = v
}

// GUID Unique identifier to avoid duplication the same post.
func (b *WallPostAdsStealthBuilder) GUID(v string) {
	b.Params["guid"] = v
}

// LinkButton Link button ID
func (b *WallPostAdsStealthBuilder) LinkButton(v string) {
	b.Params["link_button"] = v
}

// LinkTitle Link title
func (b *WallPostAdsStealthBuilder) LinkTitle(v string) {
	b.Params["link_title"] = v
}

// LinkImage Link image url
func (b *WallPostAdsStealthBuilder) LinkImage(v string) {
	b.Params["link_image"] = v
}

// LinkVideo Link video ID in format "<owner_id>_<media_id>"
func (b *WallPostAdsStealthBuilder) LinkVideo(v string) {
	b.Params["link_video"] = v
}

// WallReportCommentBuilder builder
//
// Reports (submits a complaint about) a comment on a post on a user wall or community wall.
//
// https://vk.com/dev/wall.reportComment
type WallReportCommentBuilder struct {
	api.Params
}

// NewWallReportCommentBuilder func
func NewWallReportCommentBuilder() *WallReportCommentBuilder {
	return &WallReportCommentBuilder{api.Params{}}
}

// OwnerID ID of the user or community that owns the wall.
func (b *WallReportCommentBuilder) OwnerID(v int) {
	b.Params["owner_id"] = v
}

// CommentID Comment ID.
func (b *WallReportCommentBuilder) CommentID(v int) {
	b.Params["comment_id"] = v
}

// Reason Reason for the complaint: '0' – spam, '1' – child pornography, '2' – extremism, '3' – violence, '4' – drug propaganda, '5' – adult material, '6' – insult, abuse
func (b *WallReportCommentBuilder) Reason(v int) {
	b.Params["reason"] = v
}

// WallReportPostBuilder builder
//
// Reports (submits a complaint about) a post on a user wall or community wall.
//
// https://vk.com/dev/wall.reportPost
type WallReportPostBuilder struct {
	api.Params
}

// NewWallReportPostBuilder func
func NewWallReportPostBuilder() *WallReportPostBuilder {
	return &WallReportPostBuilder{api.Params{}}
}

// OwnerID ID of the user or community that owns the wall.
func (b *WallReportPostBuilder) OwnerID(v int) {
	b.Params["owner_id"] = v
}

// PostID Post ID.
func (b *WallReportPostBuilder) PostID(v int) {
	b.Params["post_id"] = v
}

// Reason Reason for the complaint: '0' – spam, '1' – child pornography, '2' – extremism, '3' – violence, '4' – drug propaganda, '5' – adult material, '6' – insult, abuse
func (b *WallReportPostBuilder) Reason(v int) {
	b.Params["reason"] = v
}

// WallRepostBuilder builder
//
// Reposts (copies) an object to a user wall or community wall.
//
// https://vk.com/dev/wall.repost
type WallRepostBuilder struct {
	api.Params
}

// NewWallRepostBuilder func
func NewWallRepostBuilder() *WallRepostBuilder {
	return &WallRepostBuilder{api.Params{}}
}

// Object ID of the object to be reposted on the wall. Example: "wall66748_3675"
func (b *WallRepostBuilder) Object(v string) {
	b.Params["object"] = v
}

// Message Comment to be added along with the reposted object.
func (b *WallRepostBuilder) Message(v string) {
	b.Params["message"] = v
}

// GroupID Target community ID when reposting to a community.
func (b *WallRepostBuilder) GroupID(v int) {
	b.Params["group_id"] = v
}

// MarkAsAds parameter
func (b *WallRepostBuilder) MarkAsAds(v bool) {
	b.Params["mark_as_ads"] = v
}

// MuteNotifications parameter
func (b *WallRepostBuilder) MuteNotifications(v bool) {
	b.Params["mute_notifications"] = v
}

// WallRestoreBuilder builder
//
// Restores a post deleted from a user wall or community wall.
//
// https://vk.com/dev/wall.restore
type WallRestoreBuilder struct {
	api.Params
}

// NewWallRestoreBuilder func
func NewWallRestoreBuilder() *WallRestoreBuilder {
	return &WallRestoreBuilder{api.Params{}}
}

// OwnerID User ID or community ID from whose wall the post was deleted. Use a negative value to designate a community ID.
func (b *WallRestoreBuilder) OwnerID(v int) {
	b.Params["owner_id"] = v
}

// PostID ID of the post to be restored.
func (b *WallRestoreBuilder) PostID(v int) {
	b.Params["post_id"] = v
}

// WallRestoreCommentBuilder builder
//
// Restores a comment deleted from a user wall or community wall.
//
// https://vk.com/dev/wall.restoreComment
type WallRestoreCommentBuilder struct {
	api.Params
}

// NewWallRestoreCommentBuilder func
func NewWallRestoreCommentBuilder() *WallRestoreCommentBuilder {
	return &WallRestoreCommentBuilder{api.Params{}}
}

// OwnerID User ID or community ID. Use a negative value to designate a community ID.
func (b *WallRestoreCommentBuilder) OwnerID(v int) {
	b.Params["owner_id"] = v
}

// CommentID Comment ID.
func (b *WallRestoreCommentBuilder) CommentID(v int) {
	b.Params["comment_id"] = v
}

// WallSearchBuilder builder
//
// Allows to search posts on user or community walls.
//
// https://vk.com/dev/wall.search
type WallSearchBuilder struct {
	api.Params
}

// NewWallSearchBuilder func
func NewWallSearchBuilder() *WallSearchBuilder {
	return &WallSearchBuilder{api.Params{}}
}

// OwnerID user or community id. "Remember that for a community 'owner_id' must be negative."
func (b *WallSearchBuilder) OwnerID(v int) {
	b.Params["owner_id"] = v
}

// Domain user or community screen name.
func (b *WallSearchBuilder) Domain(v string) {
	b.Params["domain"] = v
}

// Query search query string.
func (b *WallSearchBuilder) Query(v string) {
	b.Params["query"] = v
}

// OwnersOnly '1' – returns only page owner's posts.
func (b *WallSearchBuilder) OwnersOnly(v bool) {
	b.Params["owners_only"] = v
}

// Count count of posts to return.
func (b *WallSearchBuilder) Count(v int) {
	b.Params["count"] = v
}

// Offset Offset needed to return a specific subset of posts.
func (b *WallSearchBuilder) Offset(v int) {
	b.Params["offset"] = v
}

// Extended show extended post info.
func (b *WallSearchBuilder) Extended(v bool) {
	b.Params["extended"] = v
}

// Fields parameter
func (b *WallSearchBuilder) Fields(v []string) {
	b.Params["fields"] = v
}

// WallUnpinBuilder builder
//
// Unpins the post on wall.
//
// https://vk.com/dev/wall.unpin
type WallUnpinBuilder struct {
	api.Params
}

// NewWallUnpinBuilder func
func NewWallUnpinBuilder() *WallUnpinBuilder {
	return &WallUnpinBuilder{api.Params{}}
}

// OwnerID ID of the user or community that owns the wall. By default, current user ID. Use a negative value to designate a community ID.
func (b *WallUnpinBuilder) OwnerID(v int) {
	b.Params["owner_id"] = v
}

// PostID Post ID.
func (b *WallUnpinBuilder) PostID(v int) {
	b.Params["post_id"] = v
}
