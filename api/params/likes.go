package params // import "github.com/SevereCloud/vksdk/v2/api/params"

import (
	"github.com/SevereCloud/vksdk/v2/api"
)

// LikesAddBuilder builder.
//
// Adds the specified object to the 'Likes' list of the current user.
//
// https://vk.com/dev/likes.add
type LikesAddBuilder struct {
	api.Params
}

// NewLikesAddBuilder func.
func NewLikesAddBuilder() *LikesAddBuilder {
	return &LikesAddBuilder{api.Params{}}
}

// Type object:
//
// * post — post on user or community wall;
//
// * comment — comment on a wall post;
//
// * photo — photo;
//
// * audio — audio;
//
// * video — video;
//
// * note — note;
//
// * photo_comment — comment on the photo;
//
// * video_comment — comment on the video;
//
// * topic_comment — comment in the discussion;
//
// * sitepage — page of the site where the [vk.com/dev/Like|Like widget] is installed.
func (b *LikesAddBuilder) Type(v string) *LikesAddBuilder {
	b.Params["type"] = v
	return b
}

// OwnerID ID of the user or community that owns the object.
func (b *LikesAddBuilder) OwnerID(v int) *LikesAddBuilder {
	b.Params["owner_id"] = v
	return b
}

// ItemID parameter.
func (b *LikesAddBuilder) ItemID(v int) *LikesAddBuilder {
	b.Params["item_id"] = v
	return b
}

// ReactionID parameter.
func (b *LikesAddBuilder) ReactionID(v int) *LikesAddBuilder {
	b.Params["reaction_id"] = v
	return b
}

// AccessKey access key required for an object owned by a private entity.
func (b *LikesAddBuilder) AccessKey(v string) *LikesAddBuilder {
	b.Params["access_key"] = v
	return b
}

// LikesDeleteBuilder builder.
//
// Deletes the specified object from the 'Likes' list of the current user.
//
// https://vk.com/dev/likes.delete
type LikesDeleteBuilder struct {
	api.Params
}

// NewLikesDeleteBuilder func.
func NewLikesDeleteBuilder() *LikesDeleteBuilder {
	return &LikesDeleteBuilder{api.Params{}}
}

// Type object:
//
// * post — post on user or community wall;
//
// * comment — comment on a wall post;
//
// * photo — photo;
//
// * audio — audio;
//
// * video — video;
//
// * note — note;
//
// * photo_comment — comment on the photo;
//
// * video_comment — comment on the video;
//
// * topic_comment — comment in the discussion;
//
// * sitepage — page of the site where the [vk.com/dev/Like|Like widget] is installed.
func (b *LikesDeleteBuilder) Type(v string) *LikesDeleteBuilder {
	b.Params["type"] = v
	return b
}

// OwnerID ID of the user or community that owns the object.
func (b *LikesDeleteBuilder) OwnerID(v int) *LikesDeleteBuilder {
	b.Params["owner_id"] = v
	return b
}

// ItemID parameter.
func (b *LikesDeleteBuilder) ItemID(v int) *LikesDeleteBuilder {
	b.Params["item_id"] = v
	return b
}

// LikesGetListBuilder builder.
//
// Returns a list of IDs of users who added the specified object to their 'Likes' list.
//
// https://vk.com/dev/likes.getList
type LikesGetListBuilder struct {
	api.Params
}

// NewLikesGetListBuilder func.
func NewLikesGetListBuilder() *LikesGetListBuilder {
	return &LikesGetListBuilder{api.Params{}}
}

// Type object:
//
// * post — post on user or community wall;
//
// * comment — comment on a wall post;
//
// * photo — photo;
//
// * audio — audio;
//
// * video — video;
//
// * note — note;
//
// * photo_comment — comment on the photo;
//
// * video_comment — comment on the video;
//
// * topic_comment — comment in the discussion;
//
// * sitepage — page of the site where the [vk.com/dev/Like|Like widget] is installed.
func (b *LikesGetListBuilder) Type(v string) *LikesGetListBuilder {
	b.Params["type"] = v
	return b
}

// OwnerID ID of the user, community, or application that owns the object.
// If the 'type' parameter is set as 'sitepage', the application ID is passed as 'owner_id'.
// Use negative value for a community id. If the 'type' parameter is not set, the 'owner_id' is assumed to be
// either the current user or the same application ID as if the 'type' parameter was set to 'sitepage'.
func (b *LikesGetListBuilder) OwnerID(v int) *LikesGetListBuilder {
	b.Params["owner_id"] = v
	return b
}

// ItemID parameter. If 'type' is set as 'sitepage', 'item_id' can include the 'page_id' parameter value used during
// initialization of the [vk.com/dev/Like|Like widget].
func (b *LikesGetListBuilder) ItemID(v int) *LikesGetListBuilder {
	b.Params["item_id"] = v
	return b
}

// PageURL URL of the page where the [vk.com/dev/Like|Like widget] is installed. Used instead of the 'item_id'
// parameter.
func (b *LikesGetListBuilder) PageURL(v string) *LikesGetListBuilder {
	b.Params["page_url"] = v
	return b
}

// Filter filters to apply:
//
// * likes — returns information about all users who liked the object (default);
//
// * copies — returns information only about users who told their friends about the object.
func (b *LikesGetListBuilder) Filter(v string) *LikesGetListBuilder {
	b.Params["filter"] = v
	return b
}

// FriendsOnly specifies which users are returned:
//
// * 1 — to return only the current user's friends;
//
// * 0 — to return all users (default).
func (b *LikesGetListBuilder) FriendsOnly(v int) *LikesGetListBuilder {
	b.Params["friends_only"] = v
	return b
}

// Extended specifies whether extended information will be returned.
//
// * 1 — to return extended information about users and communities from the 'Likes' list;
//
// * 0 — to return no additional information (default).
func (b *LikesGetListBuilder) Extended(v bool) *LikesGetListBuilder {
	b.Params["extended"] = v
	return b
}

// Offset needed to select a specific subset of users.
func (b *LikesGetListBuilder) Offset(v int) *LikesGetListBuilder {
	b.Params["offset"] = v
	return b
}

// Count number of user IDs to return (maximum '1000'). Default is '100' if 'friends_only' is set to '0',
// otherwise, the default is '10' if 'friends_only' is set to '1'.
func (b *LikesGetListBuilder) Count(v int) *LikesGetListBuilder {
	b.Params["count"] = v
	return b
}

// SkipOwn parameter.
func (b *LikesGetListBuilder) SkipOwn(v bool) *LikesGetListBuilder {
	b.Params["skip_own"] = v
	return b
}

// LikesIsLikedBuilder builder.
//
// Checks for the object in the 'Likes' list of the specified user.
//
// https://vk.com/dev/likes.isLiked
type LikesIsLikedBuilder struct {
	api.Params
}

// NewLikesIsLikedBuilder func.
func NewLikesIsLikedBuilder() *LikesIsLikedBuilder {
	return &LikesIsLikedBuilder{api.Params{}}
}

// UserID parameter.
func (b *LikesIsLikedBuilder) UserID(v int) *LikesIsLikedBuilder {
	b.Params["user_id"] = v
	return b
}

// Type object:
//
// * post — post on user or community wall;
//
// * comment — comment on a wall post;
//
// * photo — photo;
//
// * audio — audio;
//
// * video — video;
//
// * note — note;
//
// * photo_comment — comment on the photo;
//
// * video_comment — comment on the video;
//
// * topic_comment — comment in the discussion.
func (b *LikesIsLikedBuilder) Type(v string) *LikesIsLikedBuilder {
	b.Params["type"] = v
	return b
}

// OwnerID ID of the user or community that owns the object.
func (b *LikesIsLikedBuilder) OwnerID(v int) *LikesIsLikedBuilder {
	b.Params["owner_id"] = v
	return b
}

// ItemID parameter.
func (b *LikesIsLikedBuilder) ItemID(v int) *LikesIsLikedBuilder {
	b.Params["item_id"] = v
	return b
}
