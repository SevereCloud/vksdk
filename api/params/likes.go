package params // import "github.com/SevereCloud/vksdk/api/params"

import (
	"github.com/SevereCloud/vksdk/api"
)

// LikesAddBulder builder
//
// Adds the specified object to the 'Likes' list of the current user.
//
// https://vk.com/dev/likes.add
type LikesAddBulder struct {
	api.Params
}

// NewLikesAddBulder func
func NewLikesAddBulder() *LikesAddBulder {
	return &LikesAddBulder{api.Params{}}
}

// Type Object type: 'post' — post on user or community wall, 'comment' — comment on a wall post, 'photo' — photo, 'audio' — audio, 'video' — video, 'note' — note, 'photo_comment' — comment on the photo, 'video_comment' — comment on the video, 'topic_comment' — comment in the discussion, 'sitepage' — page of the site where the [vk.com/dev/Like|Like widget] is installed
func (b *LikesAddBulder) Type(v string) {
	b.Params["type"] = v
}

// OwnerID ID of the user or community that owns the object.
func (b *LikesAddBulder) OwnerID(v int) {
	b.Params["owner_id"] = v
}

// ItemID Object ID.
func (b *LikesAddBulder) ItemID(v int) {
	b.Params["item_id"] = v
}

// AccessKey Access key required for an object owned by a private entity.
func (b *LikesAddBulder) AccessKey(v string) {
	b.Params["access_key"] = v
}

// LikesDeleteBulder builder
//
// Deletes the specified object from the 'Likes' list of the current user.
//
// https://vk.com/dev/likes.delete
type LikesDeleteBulder struct {
	api.Params
}

// NewLikesDeleteBulder func
func NewLikesDeleteBulder() *LikesDeleteBulder {
	return &LikesDeleteBulder{api.Params{}}
}

// Type Object type: 'post' — post on user or community wall, 'comment' — comment on a wall post, 'photo' — photo, 'audio' — audio, 'video' — video, 'note' — note, 'photo_comment' — comment on the photo, 'video_comment' — comment on the video, 'topic_comment' — comment in the discussion, 'sitepage' — page of the site where the [vk.com/dev/Like|Like widget] is installed
func (b *LikesDeleteBulder) Type(v string) {
	b.Params["type"] = v
}

// OwnerID ID of the user or community that owns the object.
func (b *LikesDeleteBulder) OwnerID(v int) {
	b.Params["owner_id"] = v
}

// ItemID Object ID.
func (b *LikesDeleteBulder) ItemID(v int) {
	b.Params["item_id"] = v
}

// LikesGetListBulder builder
//
// Returns a list of IDs of users who added the specified object to their 'Likes' list.
//
// https://vk.com/dev/likes.getList
type LikesGetListBulder struct {
	api.Params
}

// NewLikesGetListBulder func
func NewLikesGetListBulder() *LikesGetListBulder {
	return &LikesGetListBulder{api.Params{}}
}

// Type , Object type: 'post' — post on user or community wall, 'comment' — comment on a wall post, 'photo' — photo, 'audio' — audio, 'video' — video, 'note' — note, 'photo_comment' — comment on the photo, 'video_comment' — comment on the video, 'topic_comment' — comment in the discussion, 'sitepage' — page of the site where the [vk.com/dev/Like|Like widget] is installed
func (b *LikesGetListBulder) Type(v string) {
	b.Params["type"] = v
}

// OwnerID ID of the user, community, or application that owns the object. If the 'type' parameter is set as 'sitepage', the application ID is passed as 'owner_id'. Use negative value for a community id. If the 'type' parameter is not set, the 'owner_id' is assumed to be either the current user or the same application ID as if the 'type' parameter was set to 'sitepage'.
func (b *LikesGetListBulder) OwnerID(v int) {
	b.Params["owner_id"] = v
}

// ItemID Object ID. If 'type' is set as 'sitepage', 'item_id' can include the 'page_id' parameter value used during initialization of the [vk.com/dev/Like|Like widget].
func (b *LikesGetListBulder) ItemID(v int) {
	b.Params["item_id"] = v
}

// PageURL URL of the page where the [vk.com/dev/Like|Like widget] is installed. Used instead of the 'item_id' parameter.
func (b *LikesGetListBulder) PageURL(v string) {
	b.Params["page_url"] = v
}

// Filter Filters to apply: 'likes' — returns information about all users who liked the object (default), 'copies' — returns information only about users who told their friends about the object
func (b *LikesGetListBulder) Filter(v string) {
	b.Params["filter"] = v
}

// FriendsOnly Specifies which users are returned: '1' — to return only the current user's friends, '0' — to return all users (default)
func (b *LikesGetListBulder) FriendsOnly(v int) {
	b.Params["friends_only"] = v
}

// Extended Specifies whether extended information will be returned. '1' — to return extended information about users and communities from the 'Likes' list, '0' — to return no additional information (default)
func (b *LikesGetListBulder) Extended(v bool) {
	b.Params["extended"] = v
}

// Offset Offset needed to select a specific subset of users.
func (b *LikesGetListBulder) Offset(v int) {
	b.Params["offset"] = v
}

// Count Number of user IDs to return (maximum '1000'). Default is '100' if 'friends_only' is set to '0', otherwise, the default is '10' if 'friends_only' is set to '1'.
func (b *LikesGetListBulder) Count(v int) {
	b.Params["count"] = v
}

// SkipOwn parameter
func (b *LikesGetListBulder) SkipOwn(v bool) {
	b.Params["skip_own"] = v
}

// LikesIsLikedBulder builder
//
// Checks for the object in the 'Likes' list of the specified user.
//
// https://vk.com/dev/likes.isLiked
type LikesIsLikedBulder struct {
	api.Params
}

// NewLikesIsLikedBulder func
func NewLikesIsLikedBulder() *LikesIsLikedBulder {
	return &LikesIsLikedBulder{api.Params{}}
}

// UserID User ID.
func (b *LikesIsLikedBulder) UserID(v int) {
	b.Params["user_id"] = v
}

// Type Object type: 'post' — post on user or community wall, 'comment' — comment on a wall post, 'photo' — photo, 'audio' — audio, 'video' — video, 'note' — note, 'photo_comment' — comment on the photo, 'video_comment' — comment on the video, 'topic_comment' — comment in the discussion
func (b *LikesIsLikedBulder) Type(v string) {
	b.Params["type"] = v
}

// OwnerID ID of the user or community that owns the object.
func (b *LikesIsLikedBulder) OwnerID(v int) {
	b.Params["owner_id"] = v
}

// ItemID Object ID.
func (b *LikesIsLikedBulder) ItemID(v int) {
	b.Params["item_id"] = v
}
