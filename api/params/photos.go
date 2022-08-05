package params // import "github.com/SevereCloud/vksdk/v2/api/params"

import (
	"github.com/SevereCloud/vksdk/v2/api"
)

// PhotosConfirmTagBuilder builder.
//
// Confirms a tag on a photo.
//
// https://vk.com/dev/photos.confirmTag
type PhotosConfirmTagBuilder struct {
	api.Params
}

// NewPhotosConfirmTagBuilder func.
func NewPhotosConfirmTagBuilder() *PhotosConfirmTagBuilder {
	return &PhotosConfirmTagBuilder{api.Params{}}
}

// OwnerID ID of the user or community that owns the photo.
func (b *PhotosConfirmTagBuilder) OwnerID(v int) *PhotosConfirmTagBuilder {
	b.Params["owner_id"] = v
	return b
}

// PhotoID parameter.
func (b *PhotosConfirmTagBuilder) PhotoID(v string) *PhotosConfirmTagBuilder {
	b.Params["photo_id"] = v
	return b
}

// TagID parameter.
func (b *PhotosConfirmTagBuilder) TagID(v int) *PhotosConfirmTagBuilder {
	b.Params["tag_id"] = v
	return b
}

// PhotosCopyBuilder builder.
//
// Allows to copy a photo to the "Saved photos" album.
//
// https://vk.com/dev/photos.copy
type PhotosCopyBuilder struct {
	api.Params
}

// NewPhotosCopyBuilder func.
func NewPhotosCopyBuilder() *PhotosCopyBuilder {
	return &PhotosCopyBuilder{api.Params{}}
}

// OwnerID photo's owner ID.
func (b *PhotosCopyBuilder) OwnerID(v int) *PhotosCopyBuilder {
	b.Params["owner_id"] = v
	return b
}

// PhotoID photo ID.
func (b *PhotosCopyBuilder) PhotoID(v int) *PhotosCopyBuilder {
	b.Params["photo_id"] = v
	return b
}

// AccessKey for private photos.
func (b *PhotosCopyBuilder) AccessKey(v string) *PhotosCopyBuilder {
	b.Params["access_key"] = v
	return b
}

// PhotosCreateAlbumBuilder builder.
//
// Creates an empty photo album.
//
// https://vk.com/dev/photos.createAlbum
type PhotosCreateAlbumBuilder struct {
	api.Params
}

// NewPhotosCreateAlbumBuilder func.
func NewPhotosCreateAlbumBuilder() *PhotosCreateAlbumBuilder {
	return &PhotosCreateAlbumBuilder{api.Params{}}
}

// Title album title.
func (b *PhotosCreateAlbumBuilder) Title(v string) *PhotosCreateAlbumBuilder {
	b.Params["title"] = v
	return b
}

// GroupID ID of the community in which the album will be created.
func (b *PhotosCreateAlbumBuilder) GroupID(v int) *PhotosCreateAlbumBuilder {
	b.Params["group_id"] = v
	return b
}

// Description album description.
func (b *PhotosCreateAlbumBuilder) Description(v string) *PhotosCreateAlbumBuilder {
	b.Params["description"] = v
	return b
}

// PrivacyView parameter.
func (b *PhotosCreateAlbumBuilder) PrivacyView(v []string) *PhotosCreateAlbumBuilder {
	b.Params["privacy_view"] = v
	return b
}

// PrivacyComment parameter.
func (b *PhotosCreateAlbumBuilder) PrivacyComment(v []string) *PhotosCreateAlbumBuilder {
	b.Params["privacy_comment"] = v
	return b
}

// UploadByAdminsOnly parameter.
func (b *PhotosCreateAlbumBuilder) UploadByAdminsOnly(v bool) *PhotosCreateAlbumBuilder {
	b.Params["upload_by_admins_only"] = v
	return b
}

// CommentsDisabled parameter.
func (b *PhotosCreateAlbumBuilder) CommentsDisabled(v bool) *PhotosCreateAlbumBuilder {
	b.Params["comments_disabled"] = v
	return b
}

// PhotosCreateCommentBuilder builder.
//
// Adds a new comment on the photo.
//
// https://vk.com/dev/photos.createComment
type PhotosCreateCommentBuilder struct {
	api.Params
}

// NewPhotosCreateCommentBuilder func.
func NewPhotosCreateCommentBuilder() *PhotosCreateCommentBuilder {
	return &PhotosCreateCommentBuilder{api.Params{}}
}

// OwnerID ID of the user or community that owns the photo.
func (b *PhotosCreateCommentBuilder) OwnerID(v int) *PhotosCreateCommentBuilder {
	b.Params["owner_id"] = v
	return b
}

// PhotoID parameter.
func (b *PhotosCreateCommentBuilder) PhotoID(v int) *PhotosCreateCommentBuilder {
	b.Params["photo_id"] = v
	return b
}

// Message comment text.
func (b *PhotosCreateCommentBuilder) Message(v string) *PhotosCreateCommentBuilder {
	b.Params["message"] = v
	return b
}

// Attachments (Required if 'message' is not set.) List of objects attached to the post, in the following format:
// "<owner_id>_<media_id>,<owner_id>_<media_id>",
// ” — Type of media attachment: 'photo' — photo, 'video' — video, 'audio' — audio, 'doc' — document,
// '<owner_id>' — Media attachment owner ID. '<media_id>' — Media attachment ID.
// Example: "photo100172_166443618,photo66748_265827614".
func (b *PhotosCreateCommentBuilder) Attachments(v interface{}) *PhotosCreateCommentBuilder {
	b.Params["attachments"] = v
	return b
}

// FromGroup '1' — to post a comment from the community.
func (b *PhotosCreateCommentBuilder) FromGroup(v bool) *PhotosCreateCommentBuilder {
	b.Params["from_group"] = v
	return b
}

// ReplyToComment parameter.
func (b *PhotosCreateCommentBuilder) ReplyToComment(v int) *PhotosCreateCommentBuilder {
	b.Params["reply_to_comment"] = v
	return b
}

// StickerID parameter.
func (b *PhotosCreateCommentBuilder) StickerID(v int) *PhotosCreateCommentBuilder {
	b.Params["sticker_id"] = v
	return b
}

// AccessKey parameter.
func (b *PhotosCreateCommentBuilder) AccessKey(v string) *PhotosCreateCommentBuilder {
	b.Params["access_key"] = v
	return b
}

// GUID parameter.
func (b *PhotosCreateCommentBuilder) GUID(v string) *PhotosCreateCommentBuilder {
	b.Params["guid"] = v
	return b
}

// PhotosDeleteBuilder builder.
//
// Deletes a photo.
//
// https://vk.com/dev/photos.delete
type PhotosDeleteBuilder struct {
	api.Params
}

// NewPhotosDeleteBuilder func.
func NewPhotosDeleteBuilder() *PhotosDeleteBuilder {
	return &PhotosDeleteBuilder{api.Params{}}
}

// OwnerID ID of the user or community that owns the photo.
func (b *PhotosDeleteBuilder) OwnerID(v int) *PhotosDeleteBuilder {
	b.Params["owner_id"] = v
	return b
}

// PhotoID parameter.
func (b *PhotosDeleteBuilder) PhotoID(v int) *PhotosDeleteBuilder {
	b.Params["photo_id"] = v
	return b
}

// PhotosDeleteAlbumBuilder builder.
//
// Deletes a photo album belonging to the current user.
//
// https://vk.com/dev/photos.deleteAlbum
type PhotosDeleteAlbumBuilder struct {
	api.Params
}

// NewPhotosDeleteAlbumBuilder func.
func NewPhotosDeleteAlbumBuilder() *PhotosDeleteAlbumBuilder {
	return &PhotosDeleteAlbumBuilder{api.Params{}}
}

// AlbumID parameter.
func (b *PhotosDeleteAlbumBuilder) AlbumID(v int) *PhotosDeleteAlbumBuilder {
	b.Params["album_id"] = v
	return b
}

// GroupID ID of the community that owns the album.
func (b *PhotosDeleteAlbumBuilder) GroupID(v int) *PhotosDeleteAlbumBuilder {
	b.Params["group_id"] = v
	return b
}

// PhotosDeleteCommentBuilder builder.
//
// Deletes a comment on the photo.
//
// https://vk.com/dev/photos.deleteComment
type PhotosDeleteCommentBuilder struct {
	api.Params
}

// NewPhotosDeleteCommentBuilder func.
func NewPhotosDeleteCommentBuilder() *PhotosDeleteCommentBuilder {
	return &PhotosDeleteCommentBuilder{api.Params{}}
}

// OwnerID ID of the user or community that owns the photo.
func (b *PhotosDeleteCommentBuilder) OwnerID(v int) *PhotosDeleteCommentBuilder {
	b.Params["owner_id"] = v
	return b
}

// CommentID parameter.
func (b *PhotosDeleteCommentBuilder) CommentID(v int) *PhotosDeleteCommentBuilder {
	b.Params["comment_id"] = v
	return b
}

// PhotosEditBuilder builder.
//
// Edits the caption of a photo.
//
// https://vk.com/dev/photos.edit
type PhotosEditBuilder struct {
	api.Params
}

// NewPhotosEditBuilder func.
func NewPhotosEditBuilder() *PhotosEditBuilder {
	return &PhotosEditBuilder{api.Params{}}
}

// OwnerID ID of the user or community that owns the photo.
func (b *PhotosEditBuilder) OwnerID(v int) *PhotosEditBuilder {
	b.Params["owner_id"] = v
	return b
}

// PhotoID parameter.
func (b *PhotosEditBuilder) PhotoID(v int) *PhotosEditBuilder {
	b.Params["photo_id"] = v
	return b
}

// Caption new caption for the photo. If this parameter is not set, it is considered to be equal to an empty string.
func (b *PhotosEditBuilder) Caption(v string) *PhotosEditBuilder {
	b.Params["caption"] = v
	return b
}

// Latitude parameter.
func (b *PhotosEditBuilder) Latitude(v float64) *PhotosEditBuilder {
	b.Params["latitude"] = v
	return b
}

// Longitude parameter.
func (b *PhotosEditBuilder) Longitude(v float64) *PhotosEditBuilder {
	b.Params["longitude"] = v
	return b
}

// PlaceStr parameter.
func (b *PhotosEditBuilder) PlaceStr(v string) *PhotosEditBuilder {
	b.Params["place_str"] = v
	return b
}

// FoursquareID parameter.
func (b *PhotosEditBuilder) FoursquareID(v string) *PhotosEditBuilder {
	b.Params["foursquare_id"] = v
	return b
}

// DeletePlace parameter.
func (b *PhotosEditBuilder) DeletePlace(v bool) *PhotosEditBuilder {
	b.Params["delete_place"] = v
	return b
}

// PhotosEditAlbumBuilder builder.
//
// Edits information about a photo album.
//
// https://vk.com/dev/photos.editAlbum
type PhotosEditAlbumBuilder struct {
	api.Params
}

// NewPhotosEditAlbumBuilder func.
func NewPhotosEditAlbumBuilder() *PhotosEditAlbumBuilder {
	return &PhotosEditAlbumBuilder{api.Params{}}
}

// AlbumID ID of the photo album to be edited.
func (b *PhotosEditAlbumBuilder) AlbumID(v int) *PhotosEditAlbumBuilder {
	b.Params["album_id"] = v
	return b
}

// Title new album title.
func (b *PhotosEditAlbumBuilder) Title(v string) *PhotosEditAlbumBuilder {
	b.Params["title"] = v
	return b
}

// Description new album description.
func (b *PhotosEditAlbumBuilder) Description(v string) *PhotosEditAlbumBuilder {
	b.Params["description"] = v
	return b
}

// OwnerID ID of the user or community that owns the album.
func (b *PhotosEditAlbumBuilder) OwnerID(v int) *PhotosEditAlbumBuilder {
	b.Params["owner_id"] = v
	return b
}

// PrivacyView parameter.
func (b *PhotosEditAlbumBuilder) PrivacyView(v []string) *PhotosEditAlbumBuilder {
	b.Params["privacy_view"] = v
	return b
}

// PrivacyComment parameter.
func (b *PhotosEditAlbumBuilder) PrivacyComment(v []string) *PhotosEditAlbumBuilder {
	b.Params["privacy_comment"] = v
	return b
}

// UploadByAdminsOnly parameter.
func (b *PhotosEditAlbumBuilder) UploadByAdminsOnly(v bool) *PhotosEditAlbumBuilder {
	b.Params["upload_by_admins_only"] = v
	return b
}

// CommentsDisabled parameter.
func (b *PhotosEditAlbumBuilder) CommentsDisabled(v bool) *PhotosEditAlbumBuilder {
	b.Params["comments_disabled"] = v
	return b
}

// PhotosEditCommentBuilder builder.
//
// Edits a comment on a photo.
//
// https://vk.com/dev/photos.editComment
type PhotosEditCommentBuilder struct {
	api.Params
}

// NewPhotosEditCommentBuilder func.
func NewPhotosEditCommentBuilder() *PhotosEditCommentBuilder {
	return &PhotosEditCommentBuilder{api.Params{}}
}

// OwnerID ID of the user or community that owns the photo.
func (b *PhotosEditCommentBuilder) OwnerID(v int) *PhotosEditCommentBuilder {
	b.Params["owner_id"] = v
	return b
}

// CommentID parameter.
func (b *PhotosEditCommentBuilder) CommentID(v int) *PhotosEditCommentBuilder {
	b.Params["comment_id"] = v
	return b
}

// Message new text of the comment.
func (b *PhotosEditCommentBuilder) Message(v string) *PhotosEditCommentBuilder {
	b.Params["message"] = v
	return b
}

// Attachments (Required if 'message' is not set.)
// List of objects attached to the post, in the following format:
// "<owner_id>_<media_id>,<owner_id>_<media_id>",
// ” — Type of media attachment: 'photo' — photo, 'video' — video, 'audio' — audio, 'doc' — document,
// '<owner_id>' — Media attachment owner ID. '<media_id>' — Media attachment ID.
// Example: "photo100172_166443618,photo66748_265827614".
func (b *PhotosEditCommentBuilder) Attachments(v interface{}) *PhotosEditCommentBuilder {
	b.Params["attachments"] = v
	return b
}

// PhotosGetBuilder builder.
//
// Returns a list of a user's or community's photos.
//
// https://vk.com/dev/photos.get
type PhotosGetBuilder struct {
	api.Params
}

// NewPhotosGetBuilder func.
func NewPhotosGetBuilder() *PhotosGetBuilder {
	return &PhotosGetBuilder{api.Params{}}
}

// OwnerID ID of the user or community that owns the photos. Use a negative value to designate a community ID.
func (b *PhotosGetBuilder) OwnerID(v int) *PhotosGetBuilder {
	b.Params["owner_id"] = v
	return b
}

// AlbumID photo album ID. To return information about photos from service albums, use the following string values:
// 'profile, wall, saved'.
func (b *PhotosGetBuilder) AlbumID(v string) *PhotosGetBuilder {
	b.Params["album_id"] = v
	return b
}

// PhotoIDs parameter.
func (b *PhotosGetBuilder) PhotoIDs(v []string) *PhotosGetBuilder {
	b.Params["photo_ids"] = v
	return b
}

// Rev sort order: '1' — reverse chronological, '0' — chronological.
func (b *PhotosGetBuilder) Rev(v bool) *PhotosGetBuilder {
	b.Params["rev"] = v
	return b
}

// Extended '1' — to return additional 'likes', 'comments', and 'tags' fields, '0' — (default).
func (b *PhotosGetBuilder) Extended(v bool) *PhotosGetBuilder {
	b.Params["extended"] = v
	return b
}

// FeedType Type of feed obtained in 'feed' field of the method.
func (b *PhotosGetBuilder) FeedType(v string) *PhotosGetBuilder {
	b.Params["feed_type"] = v
	return b
}

// Feed unixtime, that can be obtained with [vk.com/dev/newsfeed.get|newsfeed.get] method in date field to get all
// photos uploaded by the user on a specific day, or photos the user has been tagged on. Also, 'uid' parameter of
// the user the event happened with shall be specified.
func (b *PhotosGetBuilder) Feed(v int) *PhotosGetBuilder {
	b.Params["feed"] = v
	return b
}

// PhotoSizes '1' — to return photo sizes in a [vk.com/dev/photo_sizes|special format].
func (b *PhotosGetBuilder) PhotoSizes(v bool) *PhotosGetBuilder {
	b.Params["photo_sizes"] = v
	return b
}

// Offset parameter.
func (b *PhotosGetBuilder) Offset(v int) *PhotosGetBuilder {
	b.Params["offset"] = v
	return b
}

// Count parameter.
func (b *PhotosGetBuilder) Count(v int) *PhotosGetBuilder {
	b.Params["count"] = v
	return b
}

// PhotosGetAlbumsBuilder builder.
//
// Returns a list of a user's or community's photo albums.
//
// https://vk.com/dev/photos.getAlbums
type PhotosGetAlbumsBuilder struct {
	api.Params
}

// NewPhotosGetAlbumsBuilder func.
func NewPhotosGetAlbumsBuilder() *PhotosGetAlbumsBuilder {
	return &PhotosGetAlbumsBuilder{api.Params{}}
}

// OwnerID ID of the user or community that owns the albums.
func (b *PhotosGetAlbumsBuilder) OwnerID(v int) *PhotosGetAlbumsBuilder {
	b.Params["owner_id"] = v
	return b
}

// AlbumIDs album IDs.
func (b *PhotosGetAlbumsBuilder) AlbumIDs(v []int) *PhotosGetAlbumsBuilder {
	b.Params["album_ids"] = v
	return b
}

// Offset needed to return a specific subset of albums.
func (b *PhotosGetAlbumsBuilder) Offset(v int) *PhotosGetAlbumsBuilder {
	b.Params["offset"] = v
	return b
}

// Count number of albums to return.
func (b *PhotosGetAlbumsBuilder) Count(v int) *PhotosGetAlbumsBuilder {
	b.Params["count"] = v
	return b
}

// NeedSystem '1' — to return system albums with negative IDs.
func (b *PhotosGetAlbumsBuilder) NeedSystem(v bool) *PhotosGetAlbumsBuilder {
	b.Params["need_system"] = v
	return b
}

// NeedCovers '1' — to return an additional 'thumb_src' field, '0' — (default).
func (b *PhotosGetAlbumsBuilder) NeedCovers(v bool) *PhotosGetAlbumsBuilder {
	b.Params["need_covers"] = v
	return b
}

// PhotoSizes '1' — to return photo sizes in a.
func (b *PhotosGetAlbumsBuilder) PhotoSizes(v bool) *PhotosGetAlbumsBuilder {
	b.Params["photo_sizes"] = v
	return b
}

// PhotosGetAlbumsCountBuilder builder.
//
// Returns the number of photo albums belonging to a user or community.
//
// https://vk.com/dev/photos.getAlbumsCount
type PhotosGetAlbumsCountBuilder struct {
	api.Params
}

// NewPhotosGetAlbumsCountBuilder func.
func NewPhotosGetAlbumsCountBuilder() *PhotosGetAlbumsCountBuilder {
	return &PhotosGetAlbumsCountBuilder{api.Params{}}
}

// UserID parameter.
func (b *PhotosGetAlbumsCountBuilder) UserID(v int) *PhotosGetAlbumsCountBuilder {
	b.Params["user_id"] = v
	return b
}

// GroupID community ID.
func (b *PhotosGetAlbumsCountBuilder) GroupID(v int) *PhotosGetAlbumsCountBuilder {
	b.Params["group_id"] = v
	return b
}

// PhotosGetAllBuilder builder.
//
// Returns a list of photos belonging to a user or community, in reverse chronological order.
//
// https://vk.com/dev/photos.getAll
type PhotosGetAllBuilder struct {
	api.Params
}

// NewPhotosGetAllBuilder func.
func NewPhotosGetAllBuilder() *PhotosGetAllBuilder {
	return &PhotosGetAllBuilder{api.Params{}}
}

// OwnerID ID of a user or community that owns the photos. Use a negative value to designate a community ID.
func (b *PhotosGetAllBuilder) OwnerID(v int) *PhotosGetAllBuilder {
	b.Params["owner_id"] = v
	return b
}

// Extended '1' — to return detailed information about photos.
func (b *PhotosGetAllBuilder) Extended(v bool) *PhotosGetAllBuilder {
	b.Params["extended"] = v
	return b
}

// Offset needed to return a specific subset of photos. By default, '0'.
func (b *PhotosGetAllBuilder) Offset(v int) *PhotosGetAllBuilder {
	b.Params["offset"] = v
	return b
}

// Count number of photos to return.
func (b *PhotosGetAllBuilder) Count(v int) *PhotosGetAllBuilder {
	b.Params["count"] = v
	return b
}

// PhotoSizes '1' – to return image sizes in [vk.com/dev/photo_sizes|special format].
func (b *PhotosGetAllBuilder) PhotoSizes(v bool) *PhotosGetAllBuilder {
	b.Params["photo_sizes"] = v
	return b
}

// NoServiceAlbums '1' – to return photos only from standard albums, '0' – to return all photos including those in
// service albums, e.g., 'My wall photos' (default).
func (b *PhotosGetAllBuilder) NoServiceAlbums(v bool) *PhotosGetAllBuilder {
	b.Params["no_service_albums"] = v
	return b
}

// NeedHidden '1' – to show information about photos being hidden from the block above the wall.
func (b *PhotosGetAllBuilder) NeedHidden(v bool) *PhotosGetAllBuilder {
	b.Params["need_hidden"] = v
	return b
}

// SkipHidden '1' – not to return photos being hidden from the block above the wall. Works only with owner_id>0,
// no_service_albums is ignored.
func (b *PhotosGetAllBuilder) SkipHidden(v bool) *PhotosGetAllBuilder {
	b.Params["skip_hidden"] = v
	return b
}

// PhotosGetAllCommentsBuilder builder.
//
// Returns a list of comments on a specific photo album or all albums of the user sorted in reverse chronological order.
//
// https://vk.com/dev/photos.getAllComments
type PhotosGetAllCommentsBuilder struct {
	api.Params
}

// NewPhotosGetAllCommentsBuilder func.
func NewPhotosGetAllCommentsBuilder() *PhotosGetAllCommentsBuilder {
	return &PhotosGetAllCommentsBuilder{api.Params{}}
}

// OwnerID ID of the user or community that owns the album(s).
func (b *PhotosGetAllCommentsBuilder) OwnerID(v int) *PhotosGetAllCommentsBuilder {
	b.Params["owner_id"] = v
	return b
}

// AlbumID parameter. If the parameter is not set, comments on all of the user's albums will be returned.
func (b *PhotosGetAllCommentsBuilder) AlbumID(v int) *PhotosGetAllCommentsBuilder {
	b.Params["album_id"] = v
	return b
}

// NeedLikes '1' — to return an additional 'likes' field, '0' — (default).
func (b *PhotosGetAllCommentsBuilder) NeedLikes(v bool) *PhotosGetAllCommentsBuilder {
	b.Params["need_likes"] = v
	return b
}

// Offset needed to return a specific subset of comments. By default, '0'.
func (b *PhotosGetAllCommentsBuilder) Offset(v int) *PhotosGetAllCommentsBuilder {
	b.Params["offset"] = v
	return b
}

// Count number of comments to return. By default, '20'. Maximum value, '100'.
func (b *PhotosGetAllCommentsBuilder) Count(v int) *PhotosGetAllCommentsBuilder {
	b.Params["count"] = v
	return b
}

// PhotosGetByIDBuilder builder.
//
// Returns information about photos by their IDs.
//
// https://vk.com/dev/photos.getById
type PhotosGetByIDBuilder struct {
	api.Params
}

// NewPhotosGetByIDBuilder func.
func NewPhotosGetByIDBuilder() *PhotosGetByIDBuilder {
	return &PhotosGetByIDBuilder{api.Params{}}
}

// Photos IDs separated with a comma, that are IDs of users who posted photos and IDs of photos themselves with an
// underscore character between such IDs. To get information about a photo in the group album, you shall specify
// group ID instead of user ID. Example: "1_129207899,6492_135055734, , -20629724_271945303".
func (b *PhotosGetByIDBuilder) Photos(v []string) *PhotosGetByIDBuilder {
	b.Params["photos"] = v
	return b
}

// Extended '1' — to return additional fields, '0' — (default).
func (b *PhotosGetByIDBuilder) Extended(v bool) *PhotosGetByIDBuilder {
	b.Params["extended"] = v
	return b
}

// PhotoSizes '1' — to return photo sizes in a.
func (b *PhotosGetByIDBuilder) PhotoSizes(v bool) *PhotosGetByIDBuilder {
	b.Params["photo_sizes"] = v
	return b
}

// PhotosGetChatUploadServerBuilder builder.
//
// Returns an upload link for chat cover pictures.
//
// https://vk.com/dev/photos.getChatUploadServer
type PhotosGetChatUploadServerBuilder struct {
	api.Params
}

// NewPhotosGetChatUploadServerBuilder func.
func NewPhotosGetChatUploadServerBuilder() *PhotosGetChatUploadServerBuilder {
	return &PhotosGetChatUploadServerBuilder{api.Params{}}
}

// ChatID ID of the chat for which you want to upload a cover photo.
func (b *PhotosGetChatUploadServerBuilder) ChatID(v int) *PhotosGetChatUploadServerBuilder {
	b.Params["chat_id"] = v
	return b
}

// CropX parameter.
func (b *PhotosGetChatUploadServerBuilder) CropX(v int) *PhotosGetChatUploadServerBuilder {
	b.Params["crop_x"] = v
	return b
}

// CropY parameter.
func (b *PhotosGetChatUploadServerBuilder) CropY(v int) *PhotosGetChatUploadServerBuilder {
	b.Params["crop_y"] = v
	return b
}

// CropWidth width (in pixels) of the photo after cropping.
func (b *PhotosGetChatUploadServerBuilder) CropWidth(v int) *PhotosGetChatUploadServerBuilder {
	b.Params["crop_width"] = v
	return b
}

// PhotosGetCommentsBuilder builder.
//
// Returns a list of comments on a photo.
//
// https://vk.com/dev/photos.getComments
type PhotosGetCommentsBuilder struct {
	api.Params
}

// NewPhotosGetCommentsBuilder func.
func NewPhotosGetCommentsBuilder() *PhotosGetCommentsBuilder {
	return &PhotosGetCommentsBuilder{api.Params{}}
}

// OwnerID ID of the user or community that owns the photo.
func (b *PhotosGetCommentsBuilder) OwnerID(v int) *PhotosGetCommentsBuilder {
	b.Params["owner_id"] = v
	return b
}

// PhotoID parameter.
func (b *PhotosGetCommentsBuilder) PhotoID(v int) *PhotosGetCommentsBuilder {
	b.Params["photo_id"] = v
	return b
}

// NeedLikes '1' — to return an additional 'likes' field, '0' — (default).
func (b *PhotosGetCommentsBuilder) NeedLikes(v bool) *PhotosGetCommentsBuilder {
	b.Params["need_likes"] = v
	return b
}

// StartCommentID parameter.
func (b *PhotosGetCommentsBuilder) StartCommentID(v int) *PhotosGetCommentsBuilder {
	b.Params["start_comment_id"] = v
	return b
}

// Offset needed to return a specific subset of comments. By default, '0'.
func (b *PhotosGetCommentsBuilder) Offset(v int) *PhotosGetCommentsBuilder {
	b.Params["offset"] = v
	return b
}

// Count number of comments to return.
func (b *PhotosGetCommentsBuilder) Count(v int) *PhotosGetCommentsBuilder {
	b.Params["count"] = v
	return b
}

// Sort order: 'asc' — old first, 'desc' — new first.
func (b *PhotosGetCommentsBuilder) Sort(v string) *PhotosGetCommentsBuilder {
	b.Params["sort"] = v
	return b
}

// AccessKey parameter.
func (b *PhotosGetCommentsBuilder) AccessKey(v string) *PhotosGetCommentsBuilder {
	b.Params["access_key"] = v
	return b
}

// Extended parameter.
func (b *PhotosGetCommentsBuilder) Extended(v bool) *PhotosGetCommentsBuilder {
	b.Params["extended"] = v
	return b
}

// Fields parameter.
func (b *PhotosGetCommentsBuilder) Fields(v []string) *PhotosGetCommentsBuilder {
	b.Params["fields"] = v
	return b
}

// PhotosGetMarketAlbumUploadServerBuilder builder.
//
// Returns the server address for market album photo upload.
//
// https://vk.com/dev/photos.getMarketAlbumUploadServer
type PhotosGetMarketAlbumUploadServerBuilder struct {
	api.Params
}

// NewPhotosGetMarketAlbumUploadServerBuilder func.
func NewPhotosGetMarketAlbumUploadServerBuilder() *PhotosGetMarketAlbumUploadServerBuilder {
	return &PhotosGetMarketAlbumUploadServerBuilder{api.Params{}}
}

// GroupID community ID.
func (b *PhotosGetMarketAlbumUploadServerBuilder) GroupID(v int) *PhotosGetMarketAlbumUploadServerBuilder {
	b.Params["group_id"] = v
	return b
}

// PhotosGetMarketUploadServerBuilder builder.
//
// Returns the server address for market photo upload.
//
// https://vk.com/dev/photos.getMarketUploadServer
type PhotosGetMarketUploadServerBuilder struct {
	api.Params
}

// NewPhotosGetMarketUploadServerBuilder func.
func NewPhotosGetMarketUploadServerBuilder() *PhotosGetMarketUploadServerBuilder {
	return &PhotosGetMarketUploadServerBuilder{api.Params{}}
}

// GroupID community ID.
func (b *PhotosGetMarketUploadServerBuilder) GroupID(v int) *PhotosGetMarketUploadServerBuilder {
	b.Params["group_id"] = v
	return b
}

// MainPhoto '1' if you want to upload the main item photo.
func (b *PhotosGetMarketUploadServerBuilder) MainPhoto(v bool) *PhotosGetMarketUploadServerBuilder {
	b.Params["main_photo"] = v
	return b
}

// CropX X coordinate of the crop left upper corner.
func (b *PhotosGetMarketUploadServerBuilder) CropX(v int) *PhotosGetMarketUploadServerBuilder {
	b.Params["crop_x"] = v
	return b
}

// CropY Y coordinate of the crop left upper corner.
func (b *PhotosGetMarketUploadServerBuilder) CropY(v int) *PhotosGetMarketUploadServerBuilder {
	b.Params["crop_y"] = v
	return b
}

// CropWidth width of the cropped photo in px.
func (b *PhotosGetMarketUploadServerBuilder) CropWidth(v int) *PhotosGetMarketUploadServerBuilder {
	b.Params["crop_width"] = v
	return b
}

// PhotosGetMessagesUploadServerBuilder builder.
//
// Returns the server address for photo upload in a private message for a user.
//
// https://vk.com/dev/photos.getMessagesUploadServer
type PhotosGetMessagesUploadServerBuilder struct {
	api.Params
}

// NewPhotosGetMessagesUploadServerBuilder func.
func NewPhotosGetMessagesUploadServerBuilder() *PhotosGetMessagesUploadServerBuilder {
	return &PhotosGetMessagesUploadServerBuilder{api.Params{}}
}

// PeerID destination ID. For user: 'User ID', e.g. '12345'. For chat: '2000000000' + 'Chat ID', e.g. '2000000001'.
// For community: '- Community ID', e.g. '-12345'.
func (b *PhotosGetMessagesUploadServerBuilder) PeerID(v int) *PhotosGetMessagesUploadServerBuilder {
	b.Params["peer_id"] = v
	return b
}

// PhotosGetNewTagsBuilder builder.
//
// Returns a list of photos with tags that have not been viewed.
//
// https://vk.com/dev/photos.getNewTags
type PhotosGetNewTagsBuilder struct {
	api.Params
}

// NewPhotosGetNewTagsBuilder func.
func NewPhotosGetNewTagsBuilder() *PhotosGetNewTagsBuilder {
	return &PhotosGetNewTagsBuilder{api.Params{}}
}

// Offset needed to return a specific subset of photos.
func (b *PhotosGetNewTagsBuilder) Offset(v int) *PhotosGetNewTagsBuilder {
	b.Params["offset"] = v
	return b
}

// Count number of photos to return.
func (b *PhotosGetNewTagsBuilder) Count(v int) *PhotosGetNewTagsBuilder {
	b.Params["count"] = v
	return b
}

// PhotosGetOwnerCoverPhotoUploadServerBuilder builder.
//
// Returns the server address for owner cover upload.
//
// https://vk.com/dev/photos.getOwnerCoverPhotoUploadServer
type PhotosGetOwnerCoverPhotoUploadServerBuilder struct {
	api.Params
}

// NewPhotosGetOwnerCoverPhotoUploadServerBuilder func.
func NewPhotosGetOwnerCoverPhotoUploadServerBuilder() *PhotosGetOwnerCoverPhotoUploadServerBuilder {
	return &PhotosGetOwnerCoverPhotoUploadServerBuilder{api.Params{}}
}

// GroupID ID of community that owns the album (if the photo will be uploaded to a community album).
func (b *PhotosGetOwnerCoverPhotoUploadServerBuilder) GroupID(v int) *PhotosGetOwnerCoverPhotoUploadServerBuilder {
	b.Params["group_id"] = v
	return b
}

// CropX X coordinate of the left-upper corner.
func (b *PhotosGetOwnerCoverPhotoUploadServerBuilder) CropX(v int) *PhotosGetOwnerCoverPhotoUploadServerBuilder {
	b.Params["crop_x"] = v
	return b
}

// CropY Y coordinate of the left-upper corner.
func (b *PhotosGetOwnerCoverPhotoUploadServerBuilder) CropY(v int) *PhotosGetOwnerCoverPhotoUploadServerBuilder {
	b.Params["crop_y"] = v
	return b
}

// CropX2 X coordinate of the right-bottom corner.
func (b *PhotosGetOwnerCoverPhotoUploadServerBuilder) CropX2(v int) *PhotosGetOwnerCoverPhotoUploadServerBuilder {
	b.Params["crop_x2"] = v
	return b
}

// CropY2 Y coordinate of the right-bottom corner.
func (b *PhotosGetOwnerCoverPhotoUploadServerBuilder) CropY2(v int) *PhotosGetOwnerCoverPhotoUploadServerBuilder {
	b.Params["crop_y2"] = v
	return b
}

// PhotosGetOwnerPhotoUploadServerBuilder builder.
//
// Returns an upload server address for a profile or community photo.
//
// https://vk.com/dev/photos.getOwnerPhotoUploadServer
type PhotosGetOwnerPhotoUploadServerBuilder struct {
	api.Params
}

// NewPhotosGetOwnerPhotoUploadServerBuilder func.
func NewPhotosGetOwnerPhotoUploadServerBuilder() *PhotosGetOwnerPhotoUploadServerBuilder {
	return &PhotosGetOwnerPhotoUploadServerBuilder{api.Params{}}
}

// OwnerID identifier of a community or current user. "Note that community id must be negative. 'owner_id=1' – user,
// 'owner_id=-1' – community, ".
func (b *PhotosGetOwnerPhotoUploadServerBuilder) OwnerID(v int) *PhotosGetOwnerPhotoUploadServerBuilder {
	b.Params["owner_id"] = v
	return b
}

// PhotosGetTagsBuilder builder.
//
// Returns a list of tags on a photo.
//
// https://vk.com/dev/photos.getTags
type PhotosGetTagsBuilder struct {
	api.Params
}

// NewPhotosGetTagsBuilder func.
func NewPhotosGetTagsBuilder() *PhotosGetTagsBuilder {
	return &PhotosGetTagsBuilder{api.Params{}}
}

// OwnerID ID of the user or community that owns the photo.
func (b *PhotosGetTagsBuilder) OwnerID(v int) *PhotosGetTagsBuilder {
	b.Params["owner_id"] = v
	return b
}

// PhotoID parameter.
func (b *PhotosGetTagsBuilder) PhotoID(v int) *PhotosGetTagsBuilder {
	b.Params["photo_id"] = v
	return b
}

// AccessKey parameter.
func (b *PhotosGetTagsBuilder) AccessKey(v string) *PhotosGetTagsBuilder {
	b.Params["access_key"] = v
	return b
}

// PhotosGetUploadServerBuilder builder.
//
// Returns the server address for photo upload.
//
// https://vk.com/dev/photos.getUploadServer
type PhotosGetUploadServerBuilder struct {
	api.Params
}

// NewPhotosGetUploadServerBuilder func.
func NewPhotosGetUploadServerBuilder() *PhotosGetUploadServerBuilder {
	return &PhotosGetUploadServerBuilder{api.Params{}}
}

// GroupID ID of community that owns the album (if the photo will be uploaded to a community album).
func (b *PhotosGetUploadServerBuilder) GroupID(v int) *PhotosGetUploadServerBuilder {
	b.Params["group_id"] = v
	return b
}

// AlbumID parameter.
func (b *PhotosGetUploadServerBuilder) AlbumID(v int) *PhotosGetUploadServerBuilder {
	b.Params["album_id"] = v
	return b
}

// PhotosGetUserPhotosBuilder builder.
//
// Returns a list of photos in which a user is tagged.
//
// https://vk.com/dev/photos.getUserPhotos
type PhotosGetUserPhotosBuilder struct {
	api.Params
}

// NewPhotosGetUserPhotosBuilder func.
func NewPhotosGetUserPhotosBuilder() *PhotosGetUserPhotosBuilder {
	return &PhotosGetUserPhotosBuilder{api.Params{}}
}

// UserID parameter.
func (b *PhotosGetUserPhotosBuilder) UserID(v int) *PhotosGetUserPhotosBuilder {
	b.Params["user_id"] = v
	return b
}

// Offset needed to return a specific subset of photos. By default, '0'.
func (b *PhotosGetUserPhotosBuilder) Offset(v int) *PhotosGetUserPhotosBuilder {
	b.Params["offset"] = v
	return b
}

// Count number of photos to return. Maximum value is 1000.
func (b *PhotosGetUserPhotosBuilder) Count(v int) *PhotosGetUserPhotosBuilder {
	b.Params["count"] = v
	return b
}

// Extended '1' — to return an additional 'likes' field, '0' — (default).
func (b *PhotosGetUserPhotosBuilder) Extended(v bool) *PhotosGetUserPhotosBuilder {
	b.Params["extended"] = v
	return b
}

// Sort order: '1' — by date the tag was added in ascending order, '0' — by date the tag was added in descending
// order.
func (b *PhotosGetUserPhotosBuilder) Sort(v string) *PhotosGetUserPhotosBuilder {
	b.Params["sort"] = v
	return b
}

// PhotosGetWallUploadServerBuilder builder.
//
// Returns the server address for photo upload onto a user's wall.
//
// https://vk.com/dev/photos.getWallUploadServer
type PhotosGetWallUploadServerBuilder struct {
	api.Params
}

// NewPhotosGetWallUploadServerBuilder func.
func NewPhotosGetWallUploadServerBuilder() *PhotosGetWallUploadServerBuilder {
	return &PhotosGetWallUploadServerBuilder{api.Params{}}
}

// GroupID ID of community to whose wall the photo will be uploaded.
func (b *PhotosGetWallUploadServerBuilder) GroupID(v int) *PhotosGetWallUploadServerBuilder {
	b.Params["group_id"] = v
	return b
}

// PhotosMakeCoverBuilder builder.
//
// Makes a photo into an album cover.
//
// https://vk.com/dev/photos.makeCover
type PhotosMakeCoverBuilder struct {
	api.Params
}

// NewPhotosMakeCoverBuilder func.
func NewPhotosMakeCoverBuilder() *PhotosMakeCoverBuilder {
	return &PhotosMakeCoverBuilder{api.Params{}}
}

// OwnerID ID of the user or community that owns the photo.
func (b *PhotosMakeCoverBuilder) OwnerID(v int) *PhotosMakeCoverBuilder {
	b.Params["owner_id"] = v
	return b
}

// PhotoID parameter.
func (b *PhotosMakeCoverBuilder) PhotoID(v int) *PhotosMakeCoverBuilder {
	b.Params["photo_id"] = v
	return b
}

// AlbumID parameter.
func (b *PhotosMakeCoverBuilder) AlbumID(v int) *PhotosMakeCoverBuilder {
	b.Params["album_id"] = v
	return b
}

// PhotosMoveBuilder builder.
//
// Moves a photo from one album to another.
//
// https://vk.com/dev/photos.move
type PhotosMoveBuilder struct {
	api.Params
}

// NewPhotosMoveBuilder func.
func NewPhotosMoveBuilder() *PhotosMoveBuilder {
	return &PhotosMoveBuilder{api.Params{}}
}

// OwnerID ID of the user or community that owns the photo.
func (b *PhotosMoveBuilder) OwnerID(v int) *PhotosMoveBuilder {
	b.Params["owner_id"] = v
	return b
}

// TargetAlbumID ID of the album to which the photo will be moved.
func (b *PhotosMoveBuilder) TargetAlbumID(v int) *PhotosMoveBuilder {
	b.Params["target_album_id"] = v
	return b
}

// PhotoID parameter.
func (b *PhotosMoveBuilder) PhotoID(v int) *PhotosMoveBuilder {
	b.Params["photo_id"] = v
	return b
}

// PhotosPutTagBuilder builder.
//
// Adds a tag on the photo.
//
// https://vk.com/dev/photos.putTag
type PhotosPutTagBuilder struct {
	api.Params
}

// NewPhotosPutTagBuilder func.
func NewPhotosPutTagBuilder() *PhotosPutTagBuilder {
	return &PhotosPutTagBuilder{api.Params{}}
}

// OwnerID ID of the user or community that owns the photo.
func (b *PhotosPutTagBuilder) OwnerID(v int) *PhotosPutTagBuilder {
	b.Params["owner_id"] = v
	return b
}

// PhotoID parameter.
func (b *PhotosPutTagBuilder) PhotoID(v int) *PhotosPutTagBuilder {
	b.Params["photo_id"] = v
	return b
}

// UserID ID of the user to be tagged.
func (b *PhotosPutTagBuilder) UserID(v int) *PhotosPutTagBuilder {
	b.Params["user_id"] = v
	return b
}

// X upper left-corner coordinate of the tagged area (as a percentage of the photo's width).
func (b *PhotosPutTagBuilder) X(v float64) *PhotosPutTagBuilder {
	b.Params["x"] = v
	return b
}

// Y upper left-corner coordinate of the tagged area (as a percentage of the photo's height).
func (b *PhotosPutTagBuilder) Y(v float64) *PhotosPutTagBuilder {
	b.Params["y"] = v
	return b
}

// X2 Lower right-corner coordinate of the tagged area (as a percentage of the photo's width).
func (b *PhotosPutTagBuilder) X2(v float64) *PhotosPutTagBuilder {
	b.Params["x2"] = v
	return b
}

// Y2 Lower right-corner coordinate of the tagged area (as a percentage of the photo's height).
func (b *PhotosPutTagBuilder) Y2(v float64) *PhotosPutTagBuilder {
	b.Params["y2"] = v
	return b
}

// PhotosRemoveTagBuilder builder.
//
// Removes a tag from a photo.
//
// https://vk.com/dev/photos.removeTag
type PhotosRemoveTagBuilder struct {
	api.Params
}

// NewPhotosRemoveTagBuilder func.
func NewPhotosRemoveTagBuilder() *PhotosRemoveTagBuilder {
	return &PhotosRemoveTagBuilder{api.Params{}}
}

// OwnerID ID of the user or community that owns the photo.
func (b *PhotosRemoveTagBuilder) OwnerID(v int) *PhotosRemoveTagBuilder {
	b.Params["owner_id"] = v
	return b
}

// PhotoID parameter.
func (b *PhotosRemoveTagBuilder) PhotoID(v int) *PhotosRemoveTagBuilder {
	b.Params["photo_id"] = v
	return b
}

// TagID parameter.
func (b *PhotosRemoveTagBuilder) TagID(v int) *PhotosRemoveTagBuilder {
	b.Params["tag_id"] = v
	return b
}

// PhotosReorderAlbumsBuilder builder.
//
// Reorders the album in the list of user albums.
//
// https://vk.com/dev/photos.reorderAlbums
type PhotosReorderAlbumsBuilder struct {
	api.Params
}

// NewPhotosReorderAlbumsBuilder func.
func NewPhotosReorderAlbumsBuilder() *PhotosReorderAlbumsBuilder {
	return &PhotosReorderAlbumsBuilder{api.Params{}}
}

// OwnerID ID of the user or community that owns the album.
func (b *PhotosReorderAlbumsBuilder) OwnerID(v int) *PhotosReorderAlbumsBuilder {
	b.Params["owner_id"] = v
	return b
}

// AlbumID parameter.
func (b *PhotosReorderAlbumsBuilder) AlbumID(v int) *PhotosReorderAlbumsBuilder {
	b.Params["album_id"] = v
	return b
}

// Before ID of the album before which the album in question shall be placed.
func (b *PhotosReorderAlbumsBuilder) Before(v int) *PhotosReorderAlbumsBuilder {
	b.Params["before"] = v
	return b
}

// After ID of the album after which the album in question shall be placed.
func (b *PhotosReorderAlbumsBuilder) After(v int) *PhotosReorderAlbumsBuilder {
	b.Params["after"] = v
	return b
}

// PhotosReorderPhotosBuilder builder.
//
// Reorders the photo in the list of photos of the user album.
//
// https://vk.com/dev/photos.reorderPhotos
type PhotosReorderPhotosBuilder struct {
	api.Params
}

// NewPhotosReorderPhotosBuilder func.
func NewPhotosReorderPhotosBuilder() *PhotosReorderPhotosBuilder {
	return &PhotosReorderPhotosBuilder{api.Params{}}
}

// OwnerID ID of the user or community that owns the photo.
func (b *PhotosReorderPhotosBuilder) OwnerID(v int) *PhotosReorderPhotosBuilder {
	b.Params["owner_id"] = v
	return b
}

// PhotoID parameter.
func (b *PhotosReorderPhotosBuilder) PhotoID(v int) *PhotosReorderPhotosBuilder {
	b.Params["photo_id"] = v
	return b
}

// Before ID of the photo before which the photo in question shall be placed.
func (b *PhotosReorderPhotosBuilder) Before(v int) *PhotosReorderPhotosBuilder {
	b.Params["before"] = v
	return b
}

// After ID of the photo after which the photo in question shall be placed.
func (b *PhotosReorderPhotosBuilder) After(v int) *PhotosReorderPhotosBuilder {
	b.Params["after"] = v
	return b
}

// PhotosReportBuilder builder.
//
// Reports (submits a complaint about) a photo.
//
// https://vk.com/dev/photos.report
type PhotosReportBuilder struct {
	api.Params
}

// NewPhotosReportBuilder func.
func NewPhotosReportBuilder() *PhotosReportBuilder {
	return &PhotosReportBuilder{api.Params{}}
}

// OwnerID ID of the user or community that owns the photo.
func (b *PhotosReportBuilder) OwnerID(v int) *PhotosReportBuilder {
	b.Params["owner_id"] = v
	return b
}

// PhotoID parameter.
func (b *PhotosReportBuilder) PhotoID(v int) *PhotosReportBuilder {
	b.Params["photo_id"] = v
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
func (b *PhotosReportBuilder) Reason(v int) *PhotosReportBuilder {
	b.Params["reason"] = v
	return b
}

// PhotosReportCommentBuilder builder.
//
// Reports (submits a complaint about) a comment on a photo.
//
// https://vk.com/dev/photos.reportComment
type PhotosReportCommentBuilder struct {
	api.Params
}

// NewPhotosReportCommentBuilder func.
func NewPhotosReportCommentBuilder() *PhotosReportCommentBuilder {
	return &PhotosReportCommentBuilder{api.Params{}}
}

// OwnerID ID of the user or community that owns the photo.
func (b *PhotosReportCommentBuilder) OwnerID(v int) *PhotosReportCommentBuilder {
	b.Params["owner_id"] = v
	return b
}

// CommentID ID of the comment being reported.
func (b *PhotosReportCommentBuilder) CommentID(v int) *PhotosReportCommentBuilder {
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
func (b *PhotosReportCommentBuilder) Reason(v int) *PhotosReportCommentBuilder {
	b.Params["reason"] = v
	return b
}

// PhotosRestoreBuilder builder.
//
// Restores a deleted photo.
//
// https://vk.com/dev/photos.restore
type PhotosRestoreBuilder struct {
	api.Params
}

// NewPhotosRestoreBuilder func.
func NewPhotosRestoreBuilder() *PhotosRestoreBuilder {
	return &PhotosRestoreBuilder{api.Params{}}
}

// OwnerID ID of the user or community that owns the photo.
func (b *PhotosRestoreBuilder) OwnerID(v int) *PhotosRestoreBuilder {
	b.Params["owner_id"] = v
	return b
}

// PhotoID parameter.
func (b *PhotosRestoreBuilder) PhotoID(v int) *PhotosRestoreBuilder {
	b.Params["photo_id"] = v
	return b
}

// PhotosRestoreCommentBuilder builder.
//
// Restores a deleted comment on a photo.
//
// https://vk.com/dev/photos.restoreComment
type PhotosRestoreCommentBuilder struct {
	api.Params
}

// NewPhotosRestoreCommentBuilder func.
func NewPhotosRestoreCommentBuilder() *PhotosRestoreCommentBuilder {
	return &PhotosRestoreCommentBuilder{api.Params{}}
}

// OwnerID ID of the user or community that owns the photo.
func (b *PhotosRestoreCommentBuilder) OwnerID(v int) *PhotosRestoreCommentBuilder {
	b.Params["owner_id"] = v
	return b
}

// CommentID ID of the deleted comment.
func (b *PhotosRestoreCommentBuilder) CommentID(v int) *PhotosRestoreCommentBuilder {
	b.Params["comment_id"] = v
	return b
}

// PhotosSaveBuilder builder.
//
// Saves photos after successful uploading.
//
// https://vk.com/dev/photos.save
type PhotosSaveBuilder struct {
	api.Params
}

// NewPhotosSaveBuilder func.
func NewPhotosSaveBuilder() *PhotosSaveBuilder {
	return &PhotosSaveBuilder{api.Params{}}
}

// AlbumID ID of the album to save photos to.
func (b *PhotosSaveBuilder) AlbumID(v int) *PhotosSaveBuilder {
	b.Params["album_id"] = v
	return b
}

// GroupID ID of the community to save photos to.
func (b *PhotosSaveBuilder) GroupID(v int) *PhotosSaveBuilder {
	b.Params["group_id"] = v
	return b
}

// Server parameter returned when photos are [vk.com/dev/upload_files|uploaded to server].
func (b *PhotosSaveBuilder) Server(v int) *PhotosSaveBuilder {
	b.Params["server"] = v
	return b
}

// PhotosList parameter returned when photos are [vk.com/dev/upload_files|uploaded to server].
func (b *PhotosSaveBuilder) PhotosList(v string) *PhotosSaveBuilder {
	b.Params["photos_list"] = v
	return b
}

// Hash parameter returned when photos are [vk.com/dev/upload_files|uploaded to server].
func (b *PhotosSaveBuilder) Hash(v string) *PhotosSaveBuilder {
	b.Params["hash"] = v
	return b
}

// Latitude geographical latitude, in degrees (from '-90' to '90').
func (b *PhotosSaveBuilder) Latitude(v float64) *PhotosSaveBuilder {
	b.Params["latitude"] = v
	return b
}

// Longitude geographical longitude, in degrees (from '-180' to '180').
func (b *PhotosSaveBuilder) Longitude(v float64) *PhotosSaveBuilder {
	b.Params["longitude"] = v
	return b
}

// Caption Text describing the photo. 2048 digits max.
func (b *PhotosSaveBuilder) Caption(v string) *PhotosSaveBuilder {
	b.Params["caption"] = v
	return b
}

// PhotosSaveMarketAlbumPhotoBuilder builder.
//
// Saves market album photos after successful uploading.
//
// https://vk.com/dev/photos.saveMarketAlbumPhoto
type PhotosSaveMarketAlbumPhotoBuilder struct {
	api.Params
}

// NewPhotosSaveMarketAlbumPhotoBuilder func.
func NewPhotosSaveMarketAlbumPhotoBuilder() *PhotosSaveMarketAlbumPhotoBuilder {
	return &PhotosSaveMarketAlbumPhotoBuilder{api.Params{}}
}

// GroupID community ID.
func (b *PhotosSaveMarketAlbumPhotoBuilder) GroupID(v int) *PhotosSaveMarketAlbumPhotoBuilder {
	b.Params["group_id"] = v
	return b
}

// Photo parameter returned when photos are [vk.com/dev/upload_files|uploaded to server].
func (b *PhotosSaveMarketAlbumPhotoBuilder) Photo(v string) *PhotosSaveMarketAlbumPhotoBuilder {
	b.Params["photo"] = v
	return b
}

// Server parameter returned when photos are [vk.com/dev/upload_files|uploaded to server].
func (b *PhotosSaveMarketAlbumPhotoBuilder) Server(v int) *PhotosSaveMarketAlbumPhotoBuilder {
	b.Params["server"] = v
	return b
}

// Hash parameter returned when photos are [vk.com/dev/upload_files|uploaded to server].
func (b *PhotosSaveMarketAlbumPhotoBuilder) Hash(v string) *PhotosSaveMarketAlbumPhotoBuilder {
	b.Params["hash"] = v
	return b
}

// PhotosSaveMarketPhotoBuilder builder.
//
// Saves market photos after successful uploading.
//
// https://vk.com/dev/photos.saveMarketPhoto
type PhotosSaveMarketPhotoBuilder struct {
	api.Params
}

// NewPhotosSaveMarketPhotoBuilder func.
func NewPhotosSaveMarketPhotoBuilder() *PhotosSaveMarketPhotoBuilder {
	return &PhotosSaveMarketPhotoBuilder{api.Params{}}
}

// GroupID community ID.
func (b *PhotosSaveMarketPhotoBuilder) GroupID(v int) *PhotosSaveMarketPhotoBuilder {
	b.Params["group_id"] = v
	return b
}

// Photo parameter returned when photos are [vk.com/dev/upload_files|uploaded to server].
func (b *PhotosSaveMarketPhotoBuilder) Photo(v string) *PhotosSaveMarketPhotoBuilder {
	b.Params["photo"] = v
	return b
}

// Server parameter returned when photos are [vk.com/dev/upload_files|uploaded to server].
func (b *PhotosSaveMarketPhotoBuilder) Server(v int) *PhotosSaveMarketPhotoBuilder {
	b.Params["server"] = v
	return b
}

// Hash parameter returned when photos are [vk.com/dev/upload_files|uploaded to server].
func (b *PhotosSaveMarketPhotoBuilder) Hash(v string) *PhotosSaveMarketPhotoBuilder {
	b.Params["hash"] = v
	return b
}

// CropData parameter returned when photos are [vk.com/dev/upload_files|uploaded to server].
func (b *PhotosSaveMarketPhotoBuilder) CropData(v string) *PhotosSaveMarketPhotoBuilder {
	b.Params["crop_data"] = v
	return b
}

// CropHash parameter returned when photos are [vk.com/dev/upload_files|uploaded to server].
func (b *PhotosSaveMarketPhotoBuilder) CropHash(v string) *PhotosSaveMarketPhotoBuilder {
	b.Params["crop_hash"] = v
	return b
}

// PhotosSaveMessagesPhotoBuilder builder.
//
// Saves a photo after being successfully uploaded. URL obtained with
// [vk.com/dev/photos.getMessagesUploadServer|photos.getMessagesUploadServer] method.
//
// https://vk.com/dev/photos.saveMessagesPhoto
type PhotosSaveMessagesPhotoBuilder struct {
	api.Params
}

// NewPhotosSaveMessagesPhotoBuilder func.
func NewPhotosSaveMessagesPhotoBuilder() *PhotosSaveMessagesPhotoBuilder {
	return &PhotosSaveMessagesPhotoBuilder{api.Params{}}
}

// Photo parameter returned when the photo is [vk.com/dev/upload_files|uploaded to the server].
func (b *PhotosSaveMessagesPhotoBuilder) Photo(v string) *PhotosSaveMessagesPhotoBuilder {
	b.Params["photo"] = v
	return b
}

// Server parameter.
func (b *PhotosSaveMessagesPhotoBuilder) Server(v int) *PhotosSaveMessagesPhotoBuilder {
	b.Params["server"] = v
	return b
}

// Hash parameter.
func (b *PhotosSaveMessagesPhotoBuilder) Hash(v string) *PhotosSaveMessagesPhotoBuilder {
	b.Params["hash"] = v
	return b
}

// PhotosSaveOwnerCoverPhotoBuilder builder.
//
// Saves cover photo after successful uploading.
//
// https://vk.com/dev/photos.saveOwnerCoverPhoto
type PhotosSaveOwnerCoverPhotoBuilder struct {
	api.Params
}

// NewPhotosSaveOwnerCoverPhotoBuilder func.
func NewPhotosSaveOwnerCoverPhotoBuilder() *PhotosSaveOwnerCoverPhotoBuilder {
	return &PhotosSaveOwnerCoverPhotoBuilder{api.Params{}}
}

// Hash parameter returned when photos are [vk.com/dev/upload_files|uploaded to server].
func (b *PhotosSaveOwnerCoverPhotoBuilder) Hash(v string) *PhotosSaveOwnerCoverPhotoBuilder {
	b.Params["hash"] = v
	return b
}

// Photo parameter returned when photos are [vk.com/dev/upload_files|uploaded to server].
func (b *PhotosSaveOwnerCoverPhotoBuilder) Photo(v string) *PhotosSaveOwnerCoverPhotoBuilder {
	b.Params["photo"] = v
	return b
}

// PhotosSaveOwnerPhotoBuilder builder.
//
// Saves a profile or community photo. Upload URL can be got with the
// [vk.com/dev/photos.getOwnerPhotoUploadServer|photos.getOwnerPhotoUploadServer] method.
//
// https://vk.com/dev/photos.saveOwnerPhoto
type PhotosSaveOwnerPhotoBuilder struct {
	api.Params
}

// NewPhotosSaveOwnerPhotoBuilder func.
func NewPhotosSaveOwnerPhotoBuilder() *PhotosSaveOwnerPhotoBuilder {
	return &PhotosSaveOwnerPhotoBuilder{api.Params{}}
}

// Server parameter returned after [vk.com/dev/upload_files|photo upload].
func (b *PhotosSaveOwnerPhotoBuilder) Server(v string) *PhotosSaveOwnerPhotoBuilder {
	b.Params["server"] = v
	return b
}

// Hash parameter returned after [vk.com/dev/upload_files|photo upload].
func (b *PhotosSaveOwnerPhotoBuilder) Hash(v string) *PhotosSaveOwnerPhotoBuilder {
	b.Params["hash"] = v
	return b
}

// Photo parameter returned after [vk.com/dev/upload_files|photo upload].
func (b *PhotosSaveOwnerPhotoBuilder) Photo(v string) *PhotosSaveOwnerPhotoBuilder {
	b.Params["photo"] = v
	return b
}

// PhotosSaveWallPhotoBuilder builder.
//
// Saves a photo to a user's or community's wall after being uploaded.
//
// https://vk.com/dev/photos.saveWallPhoto
type PhotosSaveWallPhotoBuilder struct {
	api.Params
}

// NewPhotosSaveWallPhotoBuilder func.
func NewPhotosSaveWallPhotoBuilder() *PhotosSaveWallPhotoBuilder {
	return &PhotosSaveWallPhotoBuilder{api.Params{}}
}

// UserID ID of the user on whose wall the photo will be saved.
func (b *PhotosSaveWallPhotoBuilder) UserID(v int) *PhotosSaveWallPhotoBuilder {
	b.Params["user_id"] = v
	return b
}

// GroupID ID of community on whose wall the photo will be saved.
func (b *PhotosSaveWallPhotoBuilder) GroupID(v int) *PhotosSaveWallPhotoBuilder {
	b.Params["group_id"] = v
	return b
}

// Photo parameter returned when the the photo is [vk.com/dev/upload_files|uploaded to the server].
func (b *PhotosSaveWallPhotoBuilder) Photo(v string) *PhotosSaveWallPhotoBuilder {
	b.Params["photo"] = v
	return b
}

// Server parameter.
func (b *PhotosSaveWallPhotoBuilder) Server(v int) *PhotosSaveWallPhotoBuilder {
	b.Params["server"] = v
	return b
}

// Hash parameter.
func (b *PhotosSaveWallPhotoBuilder) Hash(v string) *PhotosSaveWallPhotoBuilder {
	b.Params["hash"] = v
	return b
}

// Latitude geographical latitude, in degrees (from '-90' to '90').
func (b *PhotosSaveWallPhotoBuilder) Latitude(v float64) *PhotosSaveWallPhotoBuilder {
	b.Params["latitude"] = v
	return b
}

// Longitude geographical longitude, in degrees (from '-180' to '180').
func (b *PhotosSaveWallPhotoBuilder) Longitude(v float64) *PhotosSaveWallPhotoBuilder {
	b.Params["longitude"] = v
	return b
}

// Caption Text describing the photo. 2048 digits max.
func (b *PhotosSaveWallPhotoBuilder) Caption(v string) *PhotosSaveWallPhotoBuilder {
	b.Params["caption"] = v
	return b
}

// PhotosSearchBuilder builder.
//
// Returns a list of photos.
//
// https://vk.com/dev/photos.search
type PhotosSearchBuilder struct {
	api.Params
}

// NewPhotosSearchBuilder func.
func NewPhotosSearchBuilder() *PhotosSearchBuilder {
	return &PhotosSearchBuilder{api.Params{}}
}

// Q search query string.
func (b *PhotosSearchBuilder) Q(v string) *PhotosSearchBuilder {
	b.Params["q"] = v
	return b
}

// Lat geographical latitude, in degrees (from '-90' to '90').
func (b *PhotosSearchBuilder) Lat(v float64) *PhotosSearchBuilder {
	b.Params["lat"] = v
	return b
}

// Long geographical longitude, in degrees (from '-180' to '180').
func (b *PhotosSearchBuilder) Long(v float64) *PhotosSearchBuilder {
	b.Params["long"] = v
	return b
}

// StartTime parameter.
func (b *PhotosSearchBuilder) StartTime(v int) *PhotosSearchBuilder {
	b.Params["start_time"] = v
	return b
}

// EndTime parameter.
func (b *PhotosSearchBuilder) EndTime(v int) *PhotosSearchBuilder {
	b.Params["end_time"] = v
	return b
}

// Sort order.
func (b *PhotosSearchBuilder) Sort(v int) *PhotosSearchBuilder {
	b.Params["sort"] = v
	return b
}

// Offset needed to return a specific subset of photos.
func (b *PhotosSearchBuilder) Offset(v int) *PhotosSearchBuilder {
	b.Params["offset"] = v
	return b
}

// Count number of photos to return.
func (b *PhotosSearchBuilder) Count(v int) *PhotosSearchBuilder {
	b.Params["count"] = v
	return b
}

// Radius of search in meters (works very approximately). Available values: '10', '100', '800', '6000', '50000'.
func (b *PhotosSearchBuilder) Radius(v int) *PhotosSearchBuilder {
	b.Params["radius"] = v
	return b
}
