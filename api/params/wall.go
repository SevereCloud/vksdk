package params // import "github.com/SevereCloud/vksdk/v2/api/params"

import (
	"github.com/SevereCloud/vksdk/v2/api"
)

// WallCheckCopyrightLinkBuilder builder.
//
// https://vk.com/dev/wall.checkCopyrightLink
type WallCheckCopyrightLinkBuilder struct {
	api.Params
}

// NewWallCheckCopyrightLinkBuilder func.
func NewWallCheckCopyrightLinkBuilder() *WallCheckCopyrightLinkBuilder {
	return &WallCheckCopyrightLinkBuilder{api.Params{}}
}

// Link parameter.
func (b *WallCheckCopyrightLinkBuilder) Link(v string) *WallCheckCopyrightLinkBuilder {
	b.Params["link"] = v
	return b
}

// WallCloseCommentsBuilder builder.
//
// https://vk.com/dev/wall.closeComments
type WallCloseCommentsBuilder struct {
	api.Params
}

// NewWallCloseCommentsBuilder func.
func NewWallCloseCommentsBuilder() *WallCloseCommentsBuilder {
	return &WallCloseCommentsBuilder{api.Params{}}
}

// OwnerID parameter.
func (b *WallCloseCommentsBuilder) OwnerID(v int) *WallCloseCommentsBuilder {
	b.Params["owner_id"] = v
	return b
}

// PostID parameter.
func (b *WallCloseCommentsBuilder) PostID(v int) *WallCloseCommentsBuilder {
	b.Params["post_id"] = v
	return b
}

// WallCreateCommentBuilder builder.
//
// Adds a comment to a post on a user wall or community wall.
//
// https://vk.com/dev/wall.createComment
type WallCreateCommentBuilder struct {
	api.Params
}

// NewWallCreateCommentBuilder func.
func NewWallCreateCommentBuilder() *WallCreateCommentBuilder {
	return &WallCreateCommentBuilder{api.Params{}}
}

// OwnerID user ID or community ID. Use a negative value to designate a community ID.
func (b *WallCreateCommentBuilder) OwnerID(v int) *WallCreateCommentBuilder {
	b.Params["owner_id"] = v
	return b
}

// PostID parameter.
func (b *WallCreateCommentBuilder) PostID(v int) *WallCreateCommentBuilder {
	b.Params["post_id"] = v
	return b
}

// FromGroup group ID.
func (b *WallCreateCommentBuilder) FromGroup(v int) *WallCreateCommentBuilder {
	b.Params["from_group"] = v
	return b
}

// Message (Required if 'attachments' is not set.) Text of the comment.
func (b *WallCreateCommentBuilder) Message(v string) *WallCreateCommentBuilder {
	b.Params["message"] = v
	return b
}

// ReplyToComment ID of comment to reply.
func (b *WallCreateCommentBuilder) ReplyToComment(v int) *WallCreateCommentBuilder {
	b.Params["reply_to_comment"] = v
	return b
}

// Attachments (Required if 'message' is not set.) List of media objects attached to the comment,
// in the following format:
// "<owner_id>_<media_id>,<owner_id>_<media_id>",
// ” — Type of media object: 'photo' — photo, 'video' — video, 'audio' — audio, 'doc' — document,
// '<owner_id>' — ID of the media owner. '<media_id>' — Media ID.
// For example: "photo100172_166443618,photo66748_265827614".
func (b *WallCreateCommentBuilder) Attachments(v interface{}) *WallCreateCommentBuilder {
	b.Params["attachments"] = v
	return b
}

// StickerID parameter.
func (b *WallCreateCommentBuilder) StickerID(v int) *WallCreateCommentBuilder {
	b.Params["sticker_id"] = v
	return b
}

// GUID unique identifier to avoid repeated comments.
func (b *WallCreateCommentBuilder) GUID(v string) *WallCreateCommentBuilder {
	b.Params["guid"] = v
	return b
}

// WallDeleteBuilder builder.
//
// Deletes a post from a user wall or community wall.
//
// https://vk.com/dev/wall.delete
type WallDeleteBuilder struct {
	api.Params
}

// NewWallDeleteBuilder func.
func NewWallDeleteBuilder() *WallDeleteBuilder {
	return &WallDeleteBuilder{api.Params{}}
}

// OwnerID user ID or community ID. Use a negative value to designate a community ID.
func (b *WallDeleteBuilder) OwnerID(v int) *WallDeleteBuilder {
	b.Params["owner_id"] = v
	return b
}

// PostID ID of the post to be deleted.
func (b *WallDeleteBuilder) PostID(v int) *WallDeleteBuilder {
	b.Params["post_id"] = v
	return b
}

// WallDeleteCommentBuilder builder.
//
// Deletes a comment on a post on a user wall or community wall.
//
// https://vk.com/dev/wall.deleteComment
type WallDeleteCommentBuilder struct {
	api.Params
}

// NewWallDeleteCommentBuilder func.
func NewWallDeleteCommentBuilder() *WallDeleteCommentBuilder {
	return &WallDeleteCommentBuilder{api.Params{}}
}

// OwnerID user ID or community ID. Use a negative value to designate a community ID.
func (b *WallDeleteCommentBuilder) OwnerID(v int) *WallDeleteCommentBuilder {
	b.Params["owner_id"] = v
	return b
}

// CommentID parameter.
func (b *WallDeleteCommentBuilder) CommentID(v int) *WallDeleteCommentBuilder {
	b.Params["comment_id"] = v
	return b
}

// WallEditBuilder builder.
//
// Edits a post on a user wall or community wall.
//
// https://vk.com/dev/wall.edit
type WallEditBuilder struct {
	api.Params
}

// NewWallEditBuilder func.
func NewWallEditBuilder() *WallEditBuilder {
	return &WallEditBuilder{api.Params{}}
}

// OwnerID user ID or community ID. Use a negative value to designate a community ID.
func (b *WallEditBuilder) OwnerID(v int) *WallEditBuilder {
	b.Params["owner_id"] = v
	return b
}

// PostID parameter.
func (b *WallEditBuilder) PostID(v int) *WallEditBuilder {
	b.Params["post_id"] = v
	return b
}

// FriendsOnly parameter.
func (b *WallEditBuilder) FriendsOnly(v bool) *WallEditBuilder {
	b.Params["friends_only"] = v
	return b
}

// Message (Required if 'attachments' is not set.) Text of the post.
func (b *WallEditBuilder) Message(v string) *WallEditBuilder {
	b.Params["message"] = v
	return b
}

// Attachments (Required if 'message' is not set.) List of objects attached to the post, in the following format:
// "<owner_id>_<media_id>,<owner_id>_<media_id>",
// ” — Type of media object: 'photo' — photo, 'video' — video, 'audio' — audio, 'doc' — document,
// '<owner_id>' — ID of the media owner. '<media_id>' — Media ID.
// May contain a link to an external page to include in the post.
// Example: "photo66748_265827614,http://habrahabr.ru",
//
// NOTE: If more than one link is being attached, an error is thrown.
func (b *WallEditBuilder) Attachments(v interface{}) *WallEditBuilder {
	b.Params["attachments"] = v
	return b
}

// Services parameter.
func (b *WallEditBuilder) Services(v string) *WallEditBuilder {
	b.Params["services"] = v
	return b
}

// Signed parameter.
func (b *WallEditBuilder) Signed(v bool) *WallEditBuilder {
	b.Params["signed"] = v
	return b
}

// PublishDate parameter.
func (b *WallEditBuilder) PublishDate(v int) *WallEditBuilder {
	b.Params["publish_date"] = v
	return b
}

// Lat parameter.
func (b *WallEditBuilder) Lat(v float64) *WallEditBuilder {
	b.Params["lat"] = v
	return b
}

// Long parameter.
func (b *WallEditBuilder) Long(v float64) *WallEditBuilder {
	b.Params["long"] = v
	return b
}

// PlaceID parameter.
func (b *WallEditBuilder) PlaceID(v int) *WallEditBuilder {
	b.Params["place_id"] = v
	return b
}

// MarkAsAds parameter.
func (b *WallEditBuilder) MarkAsAds(v bool) *WallEditBuilder {
	b.Params["mark_as_ads"] = v
	return b
}

// CloseComments parameter.
func (b *WallEditBuilder) CloseComments(v bool) *WallEditBuilder {
	b.Params["close_comments"] = v
	return b
}

// PosterBkgID parameter.
func (b *WallEditBuilder) PosterBkgID(v int) *WallEditBuilder {
	b.Params["poster_bkg_id"] = v
	return b
}

// PosterBkgOwnerID parameter.
func (b *WallEditBuilder) PosterBkgOwnerID(v int) *WallEditBuilder {
	b.Params["poster_bkg_owner_id"] = v
	return b
}

// PosterBkgAccessHash parameter.
func (b *WallEditBuilder) PosterBkgAccessHash(v string) *WallEditBuilder {
	b.Params["poster_bkg_access_hash"] = v
	return b
}

// WallEditAdsStealthBuilder builder.
//
// Allows to edit hidden post.
//
// https://vk.com/dev/wall.editAdsStealth
type WallEditAdsStealthBuilder struct {
	api.Params
}

// NewWallEditAdsStealthBuilder func.
func NewWallEditAdsStealthBuilder() *WallEditAdsStealthBuilder {
	return &WallEditAdsStealthBuilder{api.Params{}}
}

// OwnerID user ID or community ID. Use a negative value to designate a community ID.
func (b *WallEditAdsStealthBuilder) OwnerID(v int) *WallEditAdsStealthBuilder {
	b.Params["owner_id"] = v
	return b
}

// PostID parameter. Used for publishing of scheduled and suggested posts.
func (b *WallEditAdsStealthBuilder) PostID(v int) *WallEditAdsStealthBuilder {
	b.Params["post_id"] = v
	return b
}

// Message (Required if 'attachments' is not set.) Text of the post.
func (b *WallEditAdsStealthBuilder) Message(v string) *WallEditAdsStealthBuilder {
	b.Params["message"] = v
	return b
}

// Attachments (Required if 'message' is not set.) List of objects attached to the post, in the following format:
// "<owner_id>_<media_id>,<owner_id>_<media_id>",
// ” — Type of media object: 'photo' — photo, 'video' — video, 'audio' — audio, 'doc' — document,
// '<owner_id>' — ID of the media owner. '<media_id>' — Media ID.
// May contain a link to an external page to include in the post.
// Example: "photo66748_265827614,http://habrahabr.ru",
//
// NOTE: If more than one link is being attached, an error is thrown.
func (b *WallEditAdsStealthBuilder) Attachments(v interface{}) *WallEditAdsStealthBuilder {
	b.Params["attachments"] = v
	return b
}

// Signed only for posts in communities with 'from_group' set to '1':
//
// * 1 — post will be signed with the name of the posting user.
//
// * 0 — post will not be signed (default).
func (b *WallEditAdsStealthBuilder) Signed(v bool) *WallEditAdsStealthBuilder {
	b.Params["signed"] = v
	return b
}

// Lat geographical latitude of a check-in, in degrees (from -90 to 90).
func (b *WallEditAdsStealthBuilder) Lat(v float64) *WallEditAdsStealthBuilder {
	b.Params["lat"] = v
	return b
}

// Long geographical longitude of a check-in, in degrees (from -180 to 180).
func (b *WallEditAdsStealthBuilder) Long(v float64) *WallEditAdsStealthBuilder {
	b.Params["long"] = v
	return b
}

// PlaceID ID of the location where the user was tagged.
func (b *WallEditAdsStealthBuilder) PlaceID(v int) *WallEditAdsStealthBuilder {
	b.Params["place_id"] = v
	return b
}

// LinkButton link button ID.
func (b *WallEditAdsStealthBuilder) LinkButton(v string) *WallEditAdsStealthBuilder {
	b.Params["link_button"] = v
	return b
}

// LinkTitle link title.
func (b *WallEditAdsStealthBuilder) LinkTitle(v string) *WallEditAdsStealthBuilder {
	b.Params["link_title"] = v
	return b
}

// LinkImage link image url.
func (b *WallEditAdsStealthBuilder) LinkImage(v string) *WallEditAdsStealthBuilder {
	b.Params["link_image"] = v
	return b
}

// LinkVideo link video ID in format "<owner_id>_<media_id>".
func (b *WallEditAdsStealthBuilder) LinkVideo(v string) *WallEditAdsStealthBuilder {
	b.Params["link_video"] = v
	return b
}

// WallEditCommentBuilder builder.
//
// Edits a comment on a user wall or community wall.
//
// https://vk.com/dev/wall.editComment
type WallEditCommentBuilder struct {
	api.Params
}

// NewWallEditCommentBuilder func.
func NewWallEditCommentBuilder() *WallEditCommentBuilder {
	return &WallEditCommentBuilder{api.Params{}}
}

// OwnerID user ID or community ID. Use a negative value to designate a community ID.
func (b *WallEditCommentBuilder) OwnerID(v int) *WallEditCommentBuilder {
	b.Params["owner_id"] = v
	return b
}

// CommentID parameter.
func (b *WallEditCommentBuilder) CommentID(v int) *WallEditCommentBuilder {
	b.Params["comment_id"] = v
	return b
}

// Message new comment text.
func (b *WallEditCommentBuilder) Message(v string) *WallEditCommentBuilder {
	b.Params["message"] = v
	return b
}

// Attachments list of objects attached to the comment, in the following format:
// "<owner_id>_<media_id>,<owner_id>_<media_id>",
// ” — Type of media attachment: 'photo' — photo, 'video' — video, 'audio' — audio, 'doc' — document,
// 'page' — wiki-page, 'note' — note, 'poll' — poll, 'album' — photo album,
// '<owner_id>' — ID of the media application owner. '<media_id>' — Media application ID.
//
// Example: "photo100172_166443618,photo66748_265827614".
func (b *WallEditCommentBuilder) Attachments(v interface{}) *WallEditCommentBuilder {
	b.Params["attachments"] = v
	return b
}

// WallGetBuilder builder.
//
// Returns a list of posts on a user wall or community wall.
//
// https://vk.com/dev/wall.get
type WallGetBuilder struct {
	api.Params
}

// NewWallGetBuilder func.
func NewWallGetBuilder() *WallGetBuilder {
	return &WallGetBuilder{api.Params{}}
}

// OwnerID ID of the user or community that owns the wall. By default, current user ID.
// Use a negative value to designate a community ID.
func (b *WallGetBuilder) OwnerID(v int) *WallGetBuilder {
	b.Params["owner_id"] = v
	return b
}

// Domain user or community short address.
func (b *WallGetBuilder) Domain(v string) *WallGetBuilder {
	b.Params["domain"] = v
	return b
}

// Offset needed to return a specific subset of posts.
func (b *WallGetBuilder) Offset(v int) *WallGetBuilder {
	b.Params["offset"] = v
	return b
}

// Count number of posts to return (maximum 100).
func (b *WallGetBuilder) Count(v int) *WallGetBuilder {
	b.Params["count"] = v
	return b
}

// Filter to apply:
//
// * owner — posts by the wall owner;
//
// * others — posts by someone else;
//
// * all — posts by the wall owner and others (default);
//
// * postponed — timed posts (only available for calls with an 'access_token');
//
// * suggests — suggested posts on a community wall.
func (b *WallGetBuilder) Filter(v string) *WallGetBuilder {
	b.Params["filter"] = v
	return b
}

// Extended '1' — to return 'wall', 'profiles', and 'groups' fields, '0' — to return no additional fields (default).
func (b *WallGetBuilder) Extended(v bool) *WallGetBuilder {
	b.Params["extended"] = v
	return b
}

// Fields parameter.
func (b *WallGetBuilder) Fields(v []string) *WallGetBuilder {
	b.Params["fields"] = v
	return b
}

// WallGetByIDBuilder builder.
//
// Returns a list of posts from user or community walls by their IDs.
//
// https://vk.com/dev/wall.getById
type WallGetByIDBuilder struct {
	api.Params
}

// NewWallGetByIDBuilder func.
func NewWallGetByIDBuilder() *WallGetByIDBuilder {
	return &WallGetByIDBuilder{api.Params{}}
}

// Posts user or community IDs and post IDs, separated by underscores.
// Use a negative value to designate a community ID.
//
// Example: "93388_21539,93388_20904,2943_4276,-1_1".
func (b *WallGetByIDBuilder) Posts(v []string) *WallGetByIDBuilder {
	b.Params["posts"] = v
	return b
}

// Extended '1' — to return user and community objects needed to display posts, '0' — no additional fields
// are returned (default).
func (b *WallGetByIDBuilder) Extended(v bool) *WallGetByIDBuilder {
	b.Params["extended"] = v
	return b
}

// CopyHistoryDepth sets the number of parent elements to include in the array 'copy_history' that is returned
// if the post is a repost from another wall.
func (b *WallGetByIDBuilder) CopyHistoryDepth(v int) *WallGetByIDBuilder {
	b.Params["copy_history_depth"] = v
	return b
}

// Fields parameter.
func (b *WallGetByIDBuilder) Fields(v []string) *WallGetByIDBuilder {
	b.Params["fields"] = v
	return b
}

// WallGetCommentsBuilder builder.
//
// Returns a list of comments on a post on a user wall or community wall.
//
// https://vk.com/dev/wall.getComments
type WallGetCommentsBuilder struct {
	api.Params
}

// NewWallGetCommentsBuilder func.
func NewWallGetCommentsBuilder() *WallGetCommentsBuilder {
	return &WallGetCommentsBuilder{api.Params{}}
}

// OwnerID user ID or community ID. Use a negative value to designate a community ID.
func (b *WallGetCommentsBuilder) OwnerID(v int) *WallGetCommentsBuilder {
	b.Params["owner_id"] = v
	return b
}

// PostID parameter.
func (b *WallGetCommentsBuilder) PostID(v int) *WallGetCommentsBuilder {
	b.Params["post_id"] = v
	return b
}

// NeedLikes parameter.
//
// * '1' — to return the 'likes' field,
//
// * '0' — not to return the 'likes' field (default).
func (b *WallGetCommentsBuilder) NeedLikes(v bool) *WallGetCommentsBuilder {
	b.Params["need_likes"] = v
	return b
}

// StartCommentID parameter.
func (b *WallGetCommentsBuilder) StartCommentID(v int) *WallGetCommentsBuilder {
	b.Params["start_comment_id"] = v
	return b
}

// Offset needed to return a specific subset of comments.
func (b *WallGetCommentsBuilder) Offset(v int) *WallGetCommentsBuilder {
	b.Params["offset"] = v
	return b
}

// Count number of comments to return (maximum 100).
func (b *WallGetCommentsBuilder) Count(v int) *WallGetCommentsBuilder {
	b.Params["count"] = v
	return b
}

// Sort order: 'asc' — chronological, 'desc' — reverse chronological.
func (b *WallGetCommentsBuilder) Sort(v string) *WallGetCommentsBuilder {
	b.Params["sort"] = v
	return b
}

// PreviewLength number of characters at which to truncate comments when previewed. By default, '90'.
// Specify '0' if you do not want to truncate comments.
func (b *WallGetCommentsBuilder) PreviewLength(v int) *WallGetCommentsBuilder {
	b.Params["preview_length"] = v
	return b
}

// Extended parameter.
func (b *WallGetCommentsBuilder) Extended(v bool) *WallGetCommentsBuilder {
	b.Params["extended"] = v
	return b
}

// Fields parameter.
func (b *WallGetCommentsBuilder) Fields(v []string) *WallGetCommentsBuilder {
	b.Params["fields"] = v
	return b
}

// CommentID parameter.
func (b *WallGetCommentsBuilder) CommentID(v int) *WallGetCommentsBuilder {
	b.Params["comment_id"] = v
	return b
}

// ThreadItemsCount count items in threads.
func (b *WallGetCommentsBuilder) ThreadItemsCount(v int) *WallGetCommentsBuilder {
	b.Params["thread_items_count"] = v
	return b
}

// WallGetRepostsBuilder builder.
//
// Returns information about reposts of a post on user wall or community wall.
//
// https://vk.com/dev/wall.getReposts
type WallGetRepostsBuilder struct {
	api.Params
}

// NewWallGetRepostsBuilder func.
func NewWallGetRepostsBuilder() *WallGetRepostsBuilder {
	return &WallGetRepostsBuilder{api.Params{}}
}

// OwnerID user ID or community ID. By default, current user ID. Use a negative value to designate a community ID.
func (b *WallGetRepostsBuilder) OwnerID(v int) *WallGetRepostsBuilder {
	b.Params["owner_id"] = v
	return b
}

// PostID parameter.
func (b *WallGetRepostsBuilder) PostID(v int) *WallGetRepostsBuilder {
	b.Params["post_id"] = v
	return b
}

// Offset needed to return a specific subset of reposts.
func (b *WallGetRepostsBuilder) Offset(v int) *WallGetRepostsBuilder {
	b.Params["offset"] = v
	return b
}

// Count number of reposts to return.
func (b *WallGetRepostsBuilder) Count(v int) *WallGetRepostsBuilder {
	b.Params["count"] = v
	return b
}

// WallOpenCommentsBuilder builder.
//
// https://vk.com/dev/wall.openComments
type WallOpenCommentsBuilder struct {
	api.Params
}

// NewWallOpenCommentsBuilder func.
func NewWallOpenCommentsBuilder() *WallOpenCommentsBuilder {
	return &WallOpenCommentsBuilder{api.Params{}}
}

// OwnerID parameter.
func (b *WallOpenCommentsBuilder) OwnerID(v int) *WallOpenCommentsBuilder {
	b.Params["owner_id"] = v
	return b
}

// PostID parameter.
func (b *WallOpenCommentsBuilder) PostID(v int) *WallOpenCommentsBuilder {
	b.Params["post_id"] = v
	return b
}

// WallPinBuilder builder.
//
// Pins the post on wall.
//
// https://vk.com/dev/wall.pin
type WallPinBuilder struct {
	api.Params
}

// NewWallPinBuilder func.
func NewWallPinBuilder() *WallPinBuilder {
	return &WallPinBuilder{api.Params{}}
}

// OwnerID ID of the user or community that owns the wall. By default, current user ID.
// Use a negative value to designate a community ID.
func (b *WallPinBuilder) OwnerID(v int) *WallPinBuilder {
	b.Params["owner_id"] = v
	return b
}

// PostID parameter.
func (b *WallPinBuilder) PostID(v int) *WallPinBuilder {
	b.Params["post_id"] = v
	return b
}

// WallPostBuilder builder.
//
// Adds a new post on a user wall or community wall. Can also be used to publish suggested or scheduled posts.
//
// https://vk.com/dev/wall.post
type WallPostBuilder struct {
	api.Params
}

// NewWallPostBuilder func.
func NewWallPostBuilder() *WallPostBuilder {
	return &WallPostBuilder{api.Params{}}
}

// OwnerID user ID or community ID. Use a negative value to designate a community ID.
func (b *WallPostBuilder) OwnerID(v int) *WallPostBuilder {
	b.Params["owner_id"] = v
	return b
}

// FriendsOnly '1' — post will be available to friends only, '0' — post will be available to all users (default).
func (b *WallPostBuilder) FriendsOnly(v bool) *WallPostBuilder {
	b.Params["friends_only"] = v
	return b
}

// FromGroup for a community:
//
// * 1 — post will be published by the community.
//
// * 0 — post will be published by the user (default).
func (b *WallPostBuilder) FromGroup(v bool) *WallPostBuilder {
	b.Params["from_group"] = v
	return b
}

// Message (Required if 'attachments' is not set.) Text of the post.
func (b *WallPostBuilder) Message(v string) *WallPostBuilder {
	b.Params["message"] = v
	return b
}

// Attachments (Required if 'message' is not set.) List of objects attached to the post, in the following format:
// "<owner_id>_<media_id>,<owner_id>_<media_id>",
// ” — Type of media attachment: 'photo' — photo, 'video' — video, 'audio' — audio, 'doc' — document,
// 'page' — wiki-page, 'note' — note, 'poll' — poll, 'album' — photo album,
// '<owner_id>' — ID of the media application owner. '<media_id>' — Media application ID.
//
// Example: "photo100172_166443618,photo66748_265827614".
//
// May contain a link to an external page to include in the post.
//
// Example: "photo66748_265827614,http://habrahabr.ru".
//
// NOTE: If more than one link is being attached, an error will be thrown.
func (b *WallPostBuilder) Attachments(v interface{}) *WallPostBuilder {
	b.Params["attachments"] = v
	return b
}

// Services list of services or websites the update will be exported to, if the user has so requested.
// Sample values: 'twitter', 'facebook'.
func (b *WallPostBuilder) Services(v string) *WallPostBuilder {
	b.Params["services"] = v
	return b
}

// Signed only for posts in communities with 'from_group' set to '1':
//
// * 1 — post will be signed with the name of the posting user;
//
// * 0 — post will not be signed (default).
func (b *WallPostBuilder) Signed(v bool) *WallPostBuilder {
	b.Params["signed"] = v
	return b
}

// PublishDate in Unix time. If used, posting will be delayed until the set time.
func (b *WallPostBuilder) PublishDate(v int) *WallPostBuilder {
	b.Params["publish_date"] = v
	return b
}

// Lat geographical latitude of a check-in, in degrees (from -90 to 90).
func (b *WallPostBuilder) Lat(v float64) *WallPostBuilder {
	b.Params["lat"] = v
	return b
}

// Long geographical longitude of a check-in, in degrees (from -180 to 180).
func (b *WallPostBuilder) Long(v float64) *WallPostBuilder {
	b.Params["long"] = v
	return b
}

// PlaceID ID of the location where the user was tagged.
func (b *WallPostBuilder) PlaceID(v int) *WallPostBuilder {
	b.Params["place_id"] = v
	return b
}

// PostID parameter. Used for publishing of scheduled and suggested posts.
func (b *WallPostBuilder) PostID(v int) *WallPostBuilder {
	b.Params["post_id"] = v
	return b
}

// GUID parameter.
func (b *WallPostBuilder) GUID(v string) *WallPostBuilder {
	b.Params["guid"] = v
	return b
}

// MarkAsAds parameter.
func (b *WallPostBuilder) MarkAsAds(v bool) *WallPostBuilder {
	b.Params["mark_as_ads"] = v
	return b
}

// CloseComments parameter.
func (b *WallPostBuilder) CloseComments(v bool) *WallPostBuilder {
	b.Params["close_comments"] = v
	return b
}

// DonutPaidDuration period of time during which the post will be available for
// VK Donut paid subscribers.
//
// -1 - exclusively for donators;
//
// 86400 - for 1 day;
//
// 172800 - for 2 days;
//
// 172800 - for 3 days;
//
// 345600 - for 4 days;
//
// 432000 - for 5 days;
//
// 518400 - for 6 days;
//
// 604800 - for 7 days.
func (b *WallPostBuilder) DonutPaidDuration(v int) *WallPostBuilder {
	b.Params["donut_paid_duration"] = v
	return b
}

// MuteNotifications parameter.
func (b *WallPostBuilder) MuteNotifications(v bool) *WallPostBuilder {
	b.Params["mute_notifications"] = v
	return b
}

// Copyright parameter.
func (b *WallPostBuilder) Copyright(v string) *WallPostBuilder {
	b.Params["copyright"] = v
	return b
}

// TopicID parameter.
func (b *WallPostBuilder) TopicID(v int) *WallPostBuilder {
	b.Params["topic_id"] = v
	return b
}

// WallPostAdsStealthBuilder builder.
//
// Allows to create hidden post which will not be shown on the community's wall and can be used for creating an ad
// with type "Community post".
//
// https://vk.com/dev/wall.postAdsStealth
type WallPostAdsStealthBuilder struct {
	api.Params
}

// NewWallPostAdsStealthBuilder func.
func NewWallPostAdsStealthBuilder() *WallPostAdsStealthBuilder {
	return &WallPostAdsStealthBuilder{api.Params{}}
}

// OwnerID user ID or community ID. Use a negative value to designate a community ID.
func (b *WallPostAdsStealthBuilder) OwnerID(v int) *WallPostAdsStealthBuilder {
	b.Params["owner_id"] = v
	return b
}

// Message (Required if 'attachments' is not set.) Text of the post.
func (b *WallPostAdsStealthBuilder) Message(v string) *WallPostAdsStealthBuilder {
	b.Params["message"] = v
	return b
}

// Attachments (Required if 'message' is not set.) List of objects attached to the post, in the following format:
// "<owner_id>_<media_id>,<owner_id>_<media_id>",
// ” — Type of media attachment: 'photo' — photo, 'video' — video, 'audio' — audio, 'doc' — document,
// 'page' — wiki-page, 'note' — note, 'poll' — poll, 'album' — photo album,
// '<owner_id>' — ID of the media application owner. '<media_id>' — Media application ID.
//
// Example: "photo100172_166443618,photo66748_265827614".
//
// May contain a link to an external page to include in the post.
//
// Example: "photo66748_265827614,http://habrahabr.ru".
//
// NOTE: If more than one link is being attached, an error will be thrown.
func (b *WallPostAdsStealthBuilder) Attachments(v interface{}) *WallPostAdsStealthBuilder {
	b.Params["attachments"] = v
	return b
}

// Signed only for posts in communities with 'from_group' set to '1':
//
// * 1 — post will be signed with the name of the posting user;
//
// * 0 — post will not be signed (default).
func (b *WallPostAdsStealthBuilder) Signed(v bool) *WallPostAdsStealthBuilder {
	b.Params["signed"] = v
	return b
}

// Lat geographical latitude of a check-in, in degrees (from -90 to 90).
func (b *WallPostAdsStealthBuilder) Lat(v float64) *WallPostAdsStealthBuilder {
	b.Params["lat"] = v
	return b
}

// Long geographical longitude of a check-in, in degrees (from -180 to 180).
func (b *WallPostAdsStealthBuilder) Long(v float64) *WallPostAdsStealthBuilder {
	b.Params["long"] = v
	return b
}

// PlaceID ID of the location where the user was tagged.
func (b *WallPostAdsStealthBuilder) PlaceID(v int) *WallPostAdsStealthBuilder {
	b.Params["place_id"] = v
	return b
}

// GUID unique identifier to avoid duplication the same post.
func (b *WallPostAdsStealthBuilder) GUID(v string) *WallPostAdsStealthBuilder {
	b.Params["guid"] = v
	return b
}

// LinkButton link button ID.
func (b *WallPostAdsStealthBuilder) LinkButton(v string) *WallPostAdsStealthBuilder {
	b.Params["link_button"] = v
	return b
}

// LinkTitle link title.
func (b *WallPostAdsStealthBuilder) LinkTitle(v string) *WallPostAdsStealthBuilder {
	b.Params["link_title"] = v
	return b
}

// LinkImage link image url.
func (b *WallPostAdsStealthBuilder) LinkImage(v string) *WallPostAdsStealthBuilder {
	b.Params["link_image"] = v
	return b
}

// LinkVideo link video ID in format "<owner_id>_<media_id>".
func (b *WallPostAdsStealthBuilder) LinkVideo(v string) *WallPostAdsStealthBuilder {
	b.Params["link_video"] = v
	return b
}

// WallReportCommentBuilder builder.
//
// Reports (submits a complaint about) a comment on a post on a user wall or community wall.
//
// https://vk.com/dev/wall.reportComment
type WallReportCommentBuilder struct {
	api.Params
}

// NewWallReportCommentBuilder func.
func NewWallReportCommentBuilder() *WallReportCommentBuilder {
	return &WallReportCommentBuilder{api.Params{}}
}

// OwnerID ID of the user or community that owns the wall.
func (b *WallReportCommentBuilder) OwnerID(v int) *WallReportCommentBuilder {
	b.Params["owner_id"] = v
	return b
}

// CommentID parameter.
func (b *WallReportCommentBuilder) CommentID(v int) *WallReportCommentBuilder {
	b.Params["comment_id"] = v
	return b
}

// Reason for the complaint:
//
// * 0 – spam;
//
// * 1 – child pornography;
//
// * 2 – extremism;
//
// * 3 – violence;
//
// * 4 – drug propaganda;
//
// * 5 – adult material;
//
// * 6 – insult, abuse.
func (b *WallReportCommentBuilder) Reason(v int) *WallReportCommentBuilder {
	b.Params["reason"] = v
	return b
}

// WallReportPostBuilder builder.
//
// Reports (submits a complaint about) a post on a user wall or community wall.
//
// https://vk.com/dev/wall.reportPost
type WallReportPostBuilder struct {
	api.Params
}

// NewWallReportPostBuilder func.
func NewWallReportPostBuilder() *WallReportPostBuilder {
	return &WallReportPostBuilder{api.Params{}}
}

// OwnerID ID of the user or community that owns the wall.
func (b *WallReportPostBuilder) OwnerID(v int) *WallReportPostBuilder {
	b.Params["owner_id"] = v
	return b
}

// PostID parameter.
func (b *WallReportPostBuilder) PostID(v int) *WallReportPostBuilder {
	b.Params["post_id"] = v
	return b
}

// Reason for the complaint:
//
// * 0 – spam;
//
// * 1 – child pornography;
//
// * 2 – extremism;
//
// * 3 – violence;
//
// * 4 – drug propaganda;
//
// * 5 – adult material;
//
// * 6 – insult, abuse.
func (b *WallReportPostBuilder) Reason(v int) *WallReportPostBuilder {
	b.Params["reason"] = v
	return b
}

// WallRepostBuilder builder.
//
// Reposts (copies) an object to a user wall or community wall.
//
// https://vk.com/dev/wall.repost
type WallRepostBuilder struct {
	api.Params
}

// NewWallRepostBuilder func.
func NewWallRepostBuilder() *WallRepostBuilder {
	return &WallRepostBuilder{api.Params{}}
}

// Object ID of the object to be reposted on the wall. Example: "wall66748_3675".
func (b *WallRepostBuilder) Object(v string) *WallRepostBuilder {
	b.Params["object"] = v
	return b
}

// Message comment to be added along with the reposted object.
func (b *WallRepostBuilder) Message(v string) *WallRepostBuilder {
	b.Params["message"] = v
	return b
}

// GroupID Target community ID when reposting to a community.
func (b *WallRepostBuilder) GroupID(v int) *WallRepostBuilder {
	b.Params["group_id"] = v
	return b
}

// MarkAsAds parameter.
func (b *WallRepostBuilder) MarkAsAds(v bool) *WallRepostBuilder {
	b.Params["mark_as_ads"] = v
	return b
}

// MuteNotifications parameter.
func (b *WallRepostBuilder) MuteNotifications(v bool) *WallRepostBuilder {
	b.Params["mute_notifications"] = v
	return b
}

// WallRestoreBuilder builder.
//
// Restores a post deleted from a user wall or community wall.
//
// https://vk.com/dev/wall.restore
type WallRestoreBuilder struct {
	api.Params
}

// NewWallRestoreBuilder func.
func NewWallRestoreBuilder() *WallRestoreBuilder {
	return &WallRestoreBuilder{api.Params{}}
}

// OwnerID user ID or community ID from whose wall the post was deleted.
// Use a negative value to designate a community ID.
func (b *WallRestoreBuilder) OwnerID(v int) *WallRestoreBuilder {
	b.Params["owner_id"] = v
	return b
}

// PostID ID of the post to be restored.
func (b *WallRestoreBuilder) PostID(v int) *WallRestoreBuilder {
	b.Params["post_id"] = v
	return b
}

// WallRestoreCommentBuilder builder.
//
// Restores a comment deleted from a user wall or community wall.
//
// https://vk.com/dev/wall.restoreComment
type WallRestoreCommentBuilder struct {
	api.Params
}

// NewWallRestoreCommentBuilder func.
func NewWallRestoreCommentBuilder() *WallRestoreCommentBuilder {
	return &WallRestoreCommentBuilder{api.Params{}}
}

// OwnerID user ID or community ID. Use a negative value to designate a community ID.
func (b *WallRestoreCommentBuilder) OwnerID(v int) *WallRestoreCommentBuilder {
	b.Params["owner_id"] = v
	return b
}

// CommentID parameter.
func (b *WallRestoreCommentBuilder) CommentID(v int) *WallRestoreCommentBuilder {
	b.Params["comment_id"] = v
	return b
}

// WallSearchBuilder builder.
//
// Allows to search posts on user or community walls.
//
// https://vk.com/dev/wall.search
type WallSearchBuilder struct {
	api.Params
}

// NewWallSearchBuilder func.
func NewWallSearchBuilder() *WallSearchBuilder {
	return &WallSearchBuilder{api.Params{}}
}

// OwnerID user or community id. "Remember that for a community 'owner_id' must be negative.".
func (b *WallSearchBuilder) OwnerID(v int) *WallSearchBuilder {
	b.Params["owner_id"] = v
	return b
}

// Domain user or community screen name.
func (b *WallSearchBuilder) Domain(v string) *WallSearchBuilder {
	b.Params["domain"] = v
	return b
}

// Query search query string.
func (b *WallSearchBuilder) Query(v string) *WallSearchBuilder {
	b.Params["query"] = v
	return b
}

// OwnersOnly '1' – returns only page owner's posts.
func (b *WallSearchBuilder) OwnersOnly(v bool) *WallSearchBuilder {
	b.Params["owners_only"] = v
	return b
}

// Count of posts to return.
func (b *WallSearchBuilder) Count(v int) *WallSearchBuilder {
	b.Params["count"] = v
	return b
}

// Offset needed to return a specific subset of posts.
func (b *WallSearchBuilder) Offset(v int) *WallSearchBuilder {
	b.Params["offset"] = v
	return b
}

// Extended show extended post info.
func (b *WallSearchBuilder) Extended(v bool) *WallSearchBuilder {
	b.Params["extended"] = v
	return b
}

// Fields parameter.
func (b *WallSearchBuilder) Fields(v []string) *WallSearchBuilder {
	b.Params["fields"] = v
	return b
}

// WallUnpinBuilder builder.
//
// Unpins the post on wall.
//
// https://vk.com/dev/wall.unpin
type WallUnpinBuilder struct {
	api.Params
}

// NewWallUnpinBuilder func.
func NewWallUnpinBuilder() *WallUnpinBuilder {
	return &WallUnpinBuilder{api.Params{}}
}

// OwnerID ID of the user or community that owns the wall. By default, current user ID.
// Use a negative value to designate a community ID.
func (b *WallUnpinBuilder) OwnerID(v int) *WallUnpinBuilder {
	b.Params["owner_id"] = v
	return b
}

// PostID post ID.
func (b *WallUnpinBuilder) PostID(v int) *WallUnpinBuilder {
	b.Params["post_id"] = v
	return b
}
