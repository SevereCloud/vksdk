package params // import "github.com/SevereCloud/vksdk/api/params"

import (
	"github.com/SevereCloud/vksdk/api"
)

// PhotosConfirmTagBulder builder
//
// Confirms a tag on a photo.
//
// https://vk.com/dev/photos.confirmTag
type PhotosConfirmTagBulder struct {
	api.Params
}

// NewPhotosConfirmTagBulder func
func NewPhotosConfirmTagBulder() *PhotosConfirmTagBulder {
	return &PhotosConfirmTagBulder{api.Params{}}
}

// OwnerID ID of the user or community that owns the photo.
func (b *PhotosConfirmTagBulder) OwnerID(v int) {
	b.Params["owner_id"] = v
}

// PhotoID Photo ID.
func (b *PhotosConfirmTagBulder) PhotoID(v string) {
	b.Params["photo_id"] = v
}

// TagID Tag ID.
func (b *PhotosConfirmTagBulder) TagID(v int) {
	b.Params["tag_id"] = v
}

// PhotosCopyBulder builder
//
// Allows to copy a photo to the "Saved photos" album
//
// https://vk.com/dev/photos.copy
type PhotosCopyBulder struct {
	api.Params
}

// NewPhotosCopyBulder func
func NewPhotosCopyBulder() *PhotosCopyBulder {
	return &PhotosCopyBulder{api.Params{}}
}

// OwnerID photo's owner ID
func (b *PhotosCopyBulder) OwnerID(v int) {
	b.Params["owner_id"] = v
}

// PhotoID photo ID
func (b *PhotosCopyBulder) PhotoID(v int) {
	b.Params["photo_id"] = v
}

// AccessKey for private photos
func (b *PhotosCopyBulder) AccessKey(v string) {
	b.Params["access_key"] = v
}

// PhotosCreateAlbumBulder builder
//
// Creates an empty photo album.
//
// https://vk.com/dev/photos.createAlbum
type PhotosCreateAlbumBulder struct {
	api.Params
}

// NewPhotosCreateAlbumBulder func
func NewPhotosCreateAlbumBulder() *PhotosCreateAlbumBulder {
	return &PhotosCreateAlbumBulder{api.Params{}}
}

// Title Album title.
func (b *PhotosCreateAlbumBulder) Title(v string) {
	b.Params["title"] = v
}

// GroupID ID of the community in which the album will be created.
func (b *PhotosCreateAlbumBulder) GroupID(v int) {
	b.Params["group_id"] = v
}

// Description Album description.
func (b *PhotosCreateAlbumBulder) Description(v string) {
	b.Params["description"] = v
}

// PrivacyView parameter
func (b *PhotosCreateAlbumBulder) PrivacyView(v []string) {
	b.Params["privacy_view"] = v
}

// PrivacyComment parameter
func (b *PhotosCreateAlbumBulder) PrivacyComment(v []string) {
	b.Params["privacy_comment"] = v
}

// UploadByAdminsOnly parameter
func (b *PhotosCreateAlbumBulder) UploadByAdminsOnly(v bool) {
	b.Params["upload_by_admins_only"] = v
}

// CommentsDisabled parameter
func (b *PhotosCreateAlbumBulder) CommentsDisabled(v bool) {
	b.Params["comments_disabled"] = v
}

// PhotosCreateCommentBulder builder
//
// Adds a new comment on the photo.
//
// https://vk.com/dev/photos.createComment
type PhotosCreateCommentBulder struct {
	api.Params
}

// NewPhotosCreateCommentBulder func
func NewPhotosCreateCommentBulder() *PhotosCreateCommentBulder {
	return &PhotosCreateCommentBulder{api.Params{}}
}

// OwnerID ID of the user or community that owns the photo.
func (b *PhotosCreateCommentBulder) OwnerID(v int) {
	b.Params["owner_id"] = v
}

// PhotoID Photo ID.
func (b *PhotosCreateCommentBulder) PhotoID(v int) {
	b.Params["photo_id"] = v
}

// Message Comment text.
func (b *PhotosCreateCommentBulder) Message(v string) {
	b.Params["message"] = v
}

// Attachments (Required if 'message' is not set.) List of objects attached to the post, in the following format: "<owner_id>_<media_id>,<owner_id>_<media_id>", '' — Type of media attachment: 'photo' — photo, 'video' — video, 'audio' — audio, 'doc' — document, '<owner_id>' — Media attachment owner ID. '<media_id>' — Media attachment ID. Example: "photo100172_166443618,photo66748_265827614"
func (b *PhotosCreateCommentBulder) Attachments(v []string) {
	b.Params["attachments"] = v
}

// FromGroup '1' — to post a comment from the community
func (b *PhotosCreateCommentBulder) FromGroup(v bool) {
	b.Params["from_group"] = v
}

// ReplyToComment parameter
func (b *PhotosCreateCommentBulder) ReplyToComment(v int) {
	b.Params["reply_to_comment"] = v
}

// StickerID parameter
func (b *PhotosCreateCommentBulder) StickerID(v int) {
	b.Params["sticker_id"] = v
}

// AccessKey parameter
func (b *PhotosCreateCommentBulder) AccessKey(v string) {
	b.Params["access_key"] = v
}

// GUID parameter
func (b *PhotosCreateCommentBulder) GUID(v string) {
	b.Params["guid"] = v
}

// PhotosDeleteBulder builder
//
// Deletes a photo.
//
// https://vk.com/dev/photos.delete
type PhotosDeleteBulder struct {
	api.Params
}

// NewPhotosDeleteBulder func
func NewPhotosDeleteBulder() *PhotosDeleteBulder {
	return &PhotosDeleteBulder{api.Params{}}
}

// OwnerID ID of the user or community that owns the photo.
func (b *PhotosDeleteBulder) OwnerID(v int) {
	b.Params["owner_id"] = v
}

// PhotoID Photo ID.
func (b *PhotosDeleteBulder) PhotoID(v int) {
	b.Params["photo_id"] = v
}

// PhotosDeleteAlbumBulder builder
//
// Deletes a photo album belonging to the current user.
//
// https://vk.com/dev/photos.deleteAlbum
type PhotosDeleteAlbumBulder struct {
	api.Params
}

// NewPhotosDeleteAlbumBulder func
func NewPhotosDeleteAlbumBulder() *PhotosDeleteAlbumBulder {
	return &PhotosDeleteAlbumBulder{api.Params{}}
}

// AlbumID Album ID.
func (b *PhotosDeleteAlbumBulder) AlbumID(v int) {
	b.Params["album_id"] = v
}

// GroupID ID of the community that owns the album.
func (b *PhotosDeleteAlbumBulder) GroupID(v int) {
	b.Params["group_id"] = v
}

// PhotosDeleteCommentBulder builder
//
// Deletes a comment on the photo.
//
// https://vk.com/dev/photos.deleteComment
type PhotosDeleteCommentBulder struct {
	api.Params
}

// NewPhotosDeleteCommentBulder func
func NewPhotosDeleteCommentBulder() *PhotosDeleteCommentBulder {
	return &PhotosDeleteCommentBulder{api.Params{}}
}

// OwnerID ID of the user or community that owns the photo.
func (b *PhotosDeleteCommentBulder) OwnerID(v int) {
	b.Params["owner_id"] = v
}

// CommentID Comment ID.
func (b *PhotosDeleteCommentBulder) CommentID(v int) {
	b.Params["comment_id"] = v
}

// PhotosEditBulder builder
//
// Edits the caption of a photo.
//
// https://vk.com/dev/photos.edit
type PhotosEditBulder struct {
	api.Params
}

// NewPhotosEditBulder func
func NewPhotosEditBulder() *PhotosEditBulder {
	return &PhotosEditBulder{api.Params{}}
}

// OwnerID ID of the user or community that owns the photo.
func (b *PhotosEditBulder) OwnerID(v int) {
	b.Params["owner_id"] = v
}

// PhotoID Photo ID.
func (b *PhotosEditBulder) PhotoID(v int) {
	b.Params["photo_id"] = v
}

// Caption New caption for the photo. If this parameter is not set, it is considered to be equal to an empty string.
func (b *PhotosEditBulder) Caption(v string) {
	b.Params["caption"] = v
}

// Latitude parameter
func (b *PhotosEditBulder) Latitude(v float64) {
	b.Params["latitude"] = v
}

// Longitude parameter
func (b *PhotosEditBulder) Longitude(v float64) {
	b.Params["longitude"] = v
}

// PlaceStr parameter
func (b *PhotosEditBulder) PlaceStr(v string) {
	b.Params["place_str"] = v
}

// FoursquareID parameter
func (b *PhotosEditBulder) FoursquareID(v string) {
	b.Params["foursquare_id"] = v
}

// DeletePlace parameter
func (b *PhotosEditBulder) DeletePlace(v bool) {
	b.Params["delete_place"] = v
}

// PhotosEditAlbumBulder builder
//
// Edits information about a photo album.
//
// https://vk.com/dev/photos.editAlbum
type PhotosEditAlbumBulder struct {
	api.Params
}

// NewPhotosEditAlbumBulder func
func NewPhotosEditAlbumBulder() *PhotosEditAlbumBulder {
	return &PhotosEditAlbumBulder{api.Params{}}
}

// AlbumID ID of the photo album to be edited.
func (b *PhotosEditAlbumBulder) AlbumID(v int) {
	b.Params["album_id"] = v
}

// Title New album title.
func (b *PhotosEditAlbumBulder) Title(v string) {
	b.Params["title"] = v
}

// Description New album description.
func (b *PhotosEditAlbumBulder) Description(v string) {
	b.Params["description"] = v
}

// OwnerID ID of the user or community that owns the album.
func (b *PhotosEditAlbumBulder) OwnerID(v int) {
	b.Params["owner_id"] = v
}

// PrivacyView parameter
func (b *PhotosEditAlbumBulder) PrivacyView(v []string) {
	b.Params["privacy_view"] = v
}

// PrivacyComment parameter
func (b *PhotosEditAlbumBulder) PrivacyComment(v []string) {
	b.Params["privacy_comment"] = v
}

// UploadByAdminsOnly parameter
func (b *PhotosEditAlbumBulder) UploadByAdminsOnly(v bool) {
	b.Params["upload_by_admins_only"] = v
}

// CommentsDisabled parameter
func (b *PhotosEditAlbumBulder) CommentsDisabled(v bool) {
	b.Params["comments_disabled"] = v
}

// PhotosEditCommentBulder builder
//
// Edits a comment on a photo.
//
// https://vk.com/dev/photos.editComment
type PhotosEditCommentBulder struct {
	api.Params
}

// NewPhotosEditCommentBulder func
func NewPhotosEditCommentBulder() *PhotosEditCommentBulder {
	return &PhotosEditCommentBulder{api.Params{}}
}

// OwnerID ID of the user or community that owns the photo.
func (b *PhotosEditCommentBulder) OwnerID(v int) {
	b.Params["owner_id"] = v
}

// CommentID Comment ID.
func (b *PhotosEditCommentBulder) CommentID(v int) {
	b.Params["comment_id"] = v
}

// Message New text of the comment.
func (b *PhotosEditCommentBulder) Message(v string) {
	b.Params["message"] = v
}

// Attachments (Required if 'message' is not set.) List of objects attached to the post, in the following format: "<owner_id>_<media_id>,<owner_id>_<media_id>", '' — Type of media attachment: 'photo' — photo, 'video' — video, 'audio' — audio, 'doc' — document, '<owner_id>' — Media attachment owner ID. '<media_id>' — Media attachment ID. Example: "photo100172_166443618,photo66748_265827614"
func (b *PhotosEditCommentBulder) Attachments(v []string) {
	b.Params["attachments"] = v
}

// PhotosGetBulder builder
//
// Returns a list of a user's or community's photos.
//
// https://vk.com/dev/photos.get
type PhotosGetBulder struct {
	api.Params
}

// NewPhotosGetBulder func
func NewPhotosGetBulder() *PhotosGetBulder {
	return &PhotosGetBulder{api.Params{}}
}

// OwnerID ID of the user or community that owns the photos. Use a negative value to designate a community ID.
func (b *PhotosGetBulder) OwnerID(v int) {
	b.Params["owner_id"] = v
}

// AlbumID Photo album ID. To return information about photos from service albums, use the following string values: 'profile, wall, saved'.
func (b *PhotosGetBulder) AlbumID(v string) {
	b.Params["album_id"] = v
}

// PhotoIDs Photo IDs.
func (b *PhotosGetBulder) PhotoIDs(v []string) {
	b.Params["photo_ids"] = v
}

// Rev Sort order: '1' — reverse chronological, '0' — chronological
func (b *PhotosGetBulder) Rev(v bool) {
	b.Params["rev"] = v
}

// Extended '1' — to return additional 'likes', 'comments', and 'tags' fields, '0' — (default)
func (b *PhotosGetBulder) Extended(v bool) {
	b.Params["extended"] = v
}

// FeedType Type of feed obtained in 'feed' field of the method.
func (b *PhotosGetBulder) FeedType(v string) {
	b.Params["feed_type"] = v
}

// Feed unixtime, that can be obtained with [vk.com/dev/newsfeed.get|newsfeed.get] method in date field to get all photos uploaded by the user on a specific day, or photos the user has been tagged on. Also, 'uid' parameter of the user the event happened with shall be specified.
func (b *PhotosGetBulder) Feed(v int) {
	b.Params["feed"] = v
}

// PhotoSizes '1' — to return photo sizes in a [vk.com/dev/photo_sizes|special format]
func (b *PhotosGetBulder) PhotoSizes(v bool) {
	b.Params["photo_sizes"] = v
}

// Offset parameter
func (b *PhotosGetBulder) Offset(v int) {
	b.Params["offset"] = v
}

// Count parameter
func (b *PhotosGetBulder) Count(v int) {
	b.Params["count"] = v
}

// PhotosGetAlbumsBulder builder
//
// Returns a list of a user's or community's photo albums.
//
// https://vk.com/dev/photos.getAlbums
type PhotosGetAlbumsBulder struct {
	api.Params
}

// NewPhotosGetAlbumsBulder func
func NewPhotosGetAlbumsBulder() *PhotosGetAlbumsBulder {
	return &PhotosGetAlbumsBulder{api.Params{}}
}

// OwnerID ID of the user or community that owns the albums.
func (b *PhotosGetAlbumsBulder) OwnerID(v int) {
	b.Params["owner_id"] = v
}

// AlbumIDs Album IDs.
func (b *PhotosGetAlbumsBulder) AlbumIDs(v []int) {
	b.Params["album_ids"] = v
}

// Offset Offset needed to return a specific subset of albums.
func (b *PhotosGetAlbumsBulder) Offset(v int) {
	b.Params["offset"] = v
}

// Count Number of albums to return.
func (b *PhotosGetAlbumsBulder) Count(v int) {
	b.Params["count"] = v
}

// NeedSystem '1' — to return system albums with negative IDs
func (b *PhotosGetAlbumsBulder) NeedSystem(v bool) {
	b.Params["need_system"] = v
}

// NeedCovers '1' — to return an additional 'thumb_src' field, '0' — (default)
func (b *PhotosGetAlbumsBulder) NeedCovers(v bool) {
	b.Params["need_covers"] = v
}

// PhotoSizes '1' — to return photo sizes in a
func (b *PhotosGetAlbumsBulder) PhotoSizes(v bool) {
	b.Params["photo_sizes"] = v
}

// PhotosGetAlbumsCountBulder builder
//
// Returns the number of photo albums belonging to a user or community.
//
// https://vk.com/dev/photos.getAlbumsCount
type PhotosGetAlbumsCountBulder struct {
	api.Params
}

// NewPhotosGetAlbumsCountBulder func
func NewPhotosGetAlbumsCountBulder() *PhotosGetAlbumsCountBulder {
	return &PhotosGetAlbumsCountBulder{api.Params{}}
}

// UserID User ID.
func (b *PhotosGetAlbumsCountBulder) UserID(v int) {
	b.Params["user_id"] = v
}

// GroupID Community ID.
func (b *PhotosGetAlbumsCountBulder) GroupID(v int) {
	b.Params["group_id"] = v
}

// PhotosGetAllBulder builder
//
// Returns a list of photos belonging to a user or community, in reverse chronological order.
//
// https://vk.com/dev/photos.getAll
type PhotosGetAllBulder struct {
	api.Params
}

// NewPhotosGetAllBulder func
func NewPhotosGetAllBulder() *PhotosGetAllBulder {
	return &PhotosGetAllBulder{api.Params{}}
}

// OwnerID ID of a user or community that owns the photos. Use a negative value to designate a community ID.
func (b *PhotosGetAllBulder) OwnerID(v int) {
	b.Params["owner_id"] = v
}

// Extended '1' — to return detailed information about photos
func (b *PhotosGetAllBulder) Extended(v bool) {
	b.Params["extended"] = v
}

// Offset Offset needed to return a specific subset of photos. By default, '0'.
func (b *PhotosGetAllBulder) Offset(v int) {
	b.Params["offset"] = v
}

// Count Number of photos to return.
func (b *PhotosGetAllBulder) Count(v int) {
	b.Params["count"] = v
}

// PhotoSizes '1' – to return image sizes in [vk.com/dev/photo_sizes|special format].
func (b *PhotosGetAllBulder) PhotoSizes(v bool) {
	b.Params["photo_sizes"] = v
}

// NoServiceAlbums '1' – to return photos only from standard albums, '0' – to return all photos including those in service albums, e.g., 'My wall photos' (default)
func (b *PhotosGetAllBulder) NoServiceAlbums(v bool) {
	b.Params["no_service_albums"] = v
}

// NeedHidden '1' – to show information about photos being hidden from the block above the wall.
func (b *PhotosGetAllBulder) NeedHidden(v bool) {
	b.Params["need_hidden"] = v
}

// SkipHidden '1' – not to return photos being hidden from the block above the wall. Works only with owner_id>0, no_service_albums is ignored.
func (b *PhotosGetAllBulder) SkipHidden(v bool) {
	b.Params["skip_hidden"] = v
}

// PhotosGetAllCommentsBulder builder
//
// Returns a list of comments on a specific photo album or all albums of the user sorted in reverse chronological order.
//
// https://vk.com/dev/photos.getAllComments
type PhotosGetAllCommentsBulder struct {
	api.Params
}

// NewPhotosGetAllCommentsBulder func
func NewPhotosGetAllCommentsBulder() *PhotosGetAllCommentsBulder {
	return &PhotosGetAllCommentsBulder{api.Params{}}
}

// OwnerID ID of the user or community that owns the album(s).
func (b *PhotosGetAllCommentsBulder) OwnerID(v int) {
	b.Params["owner_id"] = v
}

// AlbumID Album ID. If the parameter is not set, comments on all of the user's albums will be returned.
func (b *PhotosGetAllCommentsBulder) AlbumID(v int) {
	b.Params["album_id"] = v
}

// NeedLikes '1' — to return an additional 'likes' field, '0' — (default)
func (b *PhotosGetAllCommentsBulder) NeedLikes(v bool) {
	b.Params["need_likes"] = v
}

// Offset Offset needed to return a specific subset of comments. By default, '0'.
func (b *PhotosGetAllCommentsBulder) Offset(v int) {
	b.Params["offset"] = v
}

// Count Number of comments to return. By default, '20'. Maximum value, '100'.
func (b *PhotosGetAllCommentsBulder) Count(v int) {
	b.Params["count"] = v
}

// PhotosGetByIDBulder builder
//
// Returns information about photos by their IDs.
//
// https://vk.com/dev/photos.getById
type PhotosGetByIDBulder struct {
	api.Params
}

// NewPhotosGetByIDBulder func
func NewPhotosGetByIDBulder() *PhotosGetByIDBulder {
	return &PhotosGetByIDBulder{api.Params{}}
}

// Photos IDs separated with a comma, that are IDs of users who posted photos and IDs of photos themselves with an underscore character between such IDs. To get information about a photo in the group album, you shall specify group ID instead of user ID. Example: "1_129207899,6492_135055734, , -20629724_271945303"
func (b *PhotosGetByIDBulder) Photos(v []string) {
	b.Params["photos"] = v
}

// Extended '1' — to return additional fields, '0' — (default)
func (b *PhotosGetByIDBulder) Extended(v bool) {
	b.Params["extended"] = v
}

// PhotoSizes '1' — to return photo sizes in a
func (b *PhotosGetByIDBulder) PhotoSizes(v bool) {
	b.Params["photo_sizes"] = v
}

// PhotosGetChatUploadServerBulder builder
//
// Returns an upload link for chat cover pictures.
//
// https://vk.com/dev/photos.getChatUploadServer
type PhotosGetChatUploadServerBulder struct {
	api.Params
}

// NewPhotosGetChatUploadServerBulder func
func NewPhotosGetChatUploadServerBulder() *PhotosGetChatUploadServerBulder {
	return &PhotosGetChatUploadServerBulder{api.Params{}}
}

// ChatID ID of the chat for which you want to upload a cover photo.
func (b *PhotosGetChatUploadServerBulder) ChatID(v int) {
	b.Params["chat_id"] = v
}

// CropX parameter
func (b *PhotosGetChatUploadServerBulder) CropX(v int) {
	b.Params["crop_x"] = v
}

// CropY parameter
func (b *PhotosGetChatUploadServerBulder) CropY(v int) {
	b.Params["crop_y"] = v
}

// CropWidth Width (in pixels) of the photo after cropping.
func (b *PhotosGetChatUploadServerBulder) CropWidth(v int) {
	b.Params["crop_width"] = v
}

// PhotosGetCommentsBulder builder
//
// Returns a list of comments on a photo.
//
// https://vk.com/dev/photos.getComments
type PhotosGetCommentsBulder struct {
	api.Params
}

// NewPhotosGetCommentsBulder func
func NewPhotosGetCommentsBulder() *PhotosGetCommentsBulder {
	return &PhotosGetCommentsBulder{api.Params{}}
}

// OwnerID ID of the user or community that owns the photo.
func (b *PhotosGetCommentsBulder) OwnerID(v int) {
	b.Params["owner_id"] = v
}

// PhotoID Photo ID.
func (b *PhotosGetCommentsBulder) PhotoID(v int) {
	b.Params["photo_id"] = v
}

// NeedLikes '1' — to return an additional 'likes' field, '0' — (default)
func (b *PhotosGetCommentsBulder) NeedLikes(v bool) {
	b.Params["need_likes"] = v
}

// StartCommentID parameter
func (b *PhotosGetCommentsBulder) StartCommentID(v int) {
	b.Params["start_comment_id"] = v
}

// Offset Offset needed to return a specific subset of comments. By default, '0'.
func (b *PhotosGetCommentsBulder) Offset(v int) {
	b.Params["offset"] = v
}

// Count Number of comments to return.
func (b *PhotosGetCommentsBulder) Count(v int) {
	b.Params["count"] = v
}

// Sort Sort order: 'asc' — old first, 'desc' — new first
func (b *PhotosGetCommentsBulder) Sort(v string) {
	b.Params["sort"] = v
}

// AccessKey parameter
func (b *PhotosGetCommentsBulder) AccessKey(v string) {
	b.Params["access_key"] = v
}

// Extended parameter
func (b *PhotosGetCommentsBulder) Extended(v bool) {
	b.Params["extended"] = v
}

// Fields parameter
func (b *PhotosGetCommentsBulder) Fields(v []string) {
	b.Params["fields"] = v
}

// PhotosGetMarketAlbumUploadServerBulder builder
//
// Returns the server address for market album photo upload.
//
// https://vk.com/dev/photos.getMarketAlbumUploadServer
type PhotosGetMarketAlbumUploadServerBulder struct {
	api.Params
}

// NewPhotosGetMarketAlbumUploadServerBulder func
func NewPhotosGetMarketAlbumUploadServerBulder() *PhotosGetMarketAlbumUploadServerBulder {
	return &PhotosGetMarketAlbumUploadServerBulder{api.Params{}}
}

// GroupID Community ID.
func (b *PhotosGetMarketAlbumUploadServerBulder) GroupID(v int) {
	b.Params["group_id"] = v
}

// PhotosGetMarketUploadServerBulder builder
//
// Returns the server address for market photo upload.
//
// https://vk.com/dev/photos.getMarketUploadServer
type PhotosGetMarketUploadServerBulder struct {
	api.Params
}

// NewPhotosGetMarketUploadServerBulder func
func NewPhotosGetMarketUploadServerBulder() *PhotosGetMarketUploadServerBulder {
	return &PhotosGetMarketUploadServerBulder{api.Params{}}
}

// GroupID Community ID.
func (b *PhotosGetMarketUploadServerBulder) GroupID(v int) {
	b.Params["group_id"] = v
}

// MainPhoto '1' if you want to upload the main item photo.
func (b *PhotosGetMarketUploadServerBulder) MainPhoto(v bool) {
	b.Params["main_photo"] = v
}

// CropX X coordinate of the crop left upper corner.
func (b *PhotosGetMarketUploadServerBulder) CropX(v int) {
	b.Params["crop_x"] = v
}

// CropY Y coordinate of the crop left upper corner.
func (b *PhotosGetMarketUploadServerBulder) CropY(v int) {
	b.Params["crop_y"] = v
}

// CropWidth Width of the cropped photo in px.
func (b *PhotosGetMarketUploadServerBulder) CropWidth(v int) {
	b.Params["crop_width"] = v
}

// PhotosGetMessagesUploadServerBulder builder
//
// Returns the server address for photo upload in a private message for a user.
//
// https://vk.com/dev/photos.getMessagesUploadServer
type PhotosGetMessagesUploadServerBulder struct {
	api.Params
}

// NewPhotosGetMessagesUploadServerBulder func
func NewPhotosGetMessagesUploadServerBulder() *PhotosGetMessagesUploadServerBulder {
	return &PhotosGetMessagesUploadServerBulder{api.Params{}}
}

// PeerID Destination ID. "For user: 'User ID', e.g. '12345'. For chat: '2000000000' + 'Chat ID', e.g. '2000000001'. For community: '- Community ID', e.g. '-12345'. "
func (b *PhotosGetMessagesUploadServerBulder) PeerID(v int) {
	b.Params["peer_id"] = v
}

// PhotosGetNewTagsBulder builder
//
// Returns a list of photos with tags that have not been viewed.
//
// https://vk.com/dev/photos.getNewTags
type PhotosGetNewTagsBulder struct {
	api.Params
}

// NewPhotosGetNewTagsBulder func
func NewPhotosGetNewTagsBulder() *PhotosGetNewTagsBulder {
	return &PhotosGetNewTagsBulder{api.Params{}}
}

// Offset Offset needed to return a specific subset of photos.
func (b *PhotosGetNewTagsBulder) Offset(v int) {
	b.Params["offset"] = v
}

// Count Number of photos to return.
func (b *PhotosGetNewTagsBulder) Count(v int) {
	b.Params["count"] = v
}

// PhotosGetOwnerCoverPhotoUploadServerBulder builder
//
// Returns the server address for owner cover upload.
//
// https://vk.com/dev/photos.getOwnerCoverPhotoUploadServer
type PhotosGetOwnerCoverPhotoUploadServerBulder struct {
	api.Params
}

// NewPhotosGetOwnerCoverPhotoUploadServerBulder func
func NewPhotosGetOwnerCoverPhotoUploadServerBulder() *PhotosGetOwnerCoverPhotoUploadServerBulder {
	return &PhotosGetOwnerCoverPhotoUploadServerBulder{api.Params{}}
}

// GroupID ID of community that owns the album (if the photo will be uploaded to a community album).
func (b *PhotosGetOwnerCoverPhotoUploadServerBulder) GroupID(v int) {
	b.Params["group_id"] = v
}

// CropX X coordinate of the left-upper corner
func (b *PhotosGetOwnerCoverPhotoUploadServerBulder) CropX(v int) {
	b.Params["crop_x"] = v
}

// CropY Y coordinate of the left-upper corner
func (b *PhotosGetOwnerCoverPhotoUploadServerBulder) CropY(v int) {
	b.Params["crop_y"] = v
}

// CropX2 X coordinate of the right-bottom corner
func (b *PhotosGetOwnerCoverPhotoUploadServerBulder) CropX2(v int) {
	b.Params["crop_x2"] = v
}

// CropY2 Y coordinate of the right-bottom corner
func (b *PhotosGetOwnerCoverPhotoUploadServerBulder) CropY2(v int) {
	b.Params["crop_y2"] = v
}

// PhotosGetOwnerPhotoUploadServerBulder builder
//
// Returns an upload server address for a profile or community photo.
//
// https://vk.com/dev/photos.getOwnerPhotoUploadServer
type PhotosGetOwnerPhotoUploadServerBulder struct {
	api.Params
}

// NewPhotosGetOwnerPhotoUploadServerBulder func
func NewPhotosGetOwnerPhotoUploadServerBulder() *PhotosGetOwnerPhotoUploadServerBulder {
	return &PhotosGetOwnerPhotoUploadServerBulder{api.Params{}}
}

// OwnerID identifier of a community or current user. "Note that community id must be negative. 'owner_id=1' – user, 'owner_id=-1' – community, "
func (b *PhotosGetOwnerPhotoUploadServerBulder) OwnerID(v int) {
	b.Params["owner_id"] = v
}

// PhotosGetTagsBulder builder
//
// Returns a list of tags on a photo.
//
// https://vk.com/dev/photos.getTags
type PhotosGetTagsBulder struct {
	api.Params
}

// NewPhotosGetTagsBulder func
func NewPhotosGetTagsBulder() *PhotosGetTagsBulder {
	return &PhotosGetTagsBulder{api.Params{}}
}

// OwnerID ID of the user or community that owns the photo.
func (b *PhotosGetTagsBulder) OwnerID(v int) {
	b.Params["owner_id"] = v
}

// PhotoID Photo ID.
func (b *PhotosGetTagsBulder) PhotoID(v int) {
	b.Params["photo_id"] = v
}

// AccessKey parameter
func (b *PhotosGetTagsBulder) AccessKey(v string) {
	b.Params["access_key"] = v
}

// PhotosGetUploadServerBulder builder
//
// Returns the server address for photo upload.
//
// https://vk.com/dev/photos.getUploadServer
type PhotosGetUploadServerBulder struct {
	api.Params
}

// NewPhotosGetUploadServerBulder func
func NewPhotosGetUploadServerBulder() *PhotosGetUploadServerBulder {
	return &PhotosGetUploadServerBulder{api.Params{}}
}

// GroupID ID of community that owns the album (if the photo will be uploaded to a community album).
func (b *PhotosGetUploadServerBulder) GroupID(v int) {
	b.Params["group_id"] = v
}

// AlbumID parameter
func (b *PhotosGetUploadServerBulder) AlbumID(v int) {
	b.Params["album_id"] = v
}

// PhotosGetUserPhotosBulder builder
//
// Returns a list of photos in which a user is tagged.
//
// https://vk.com/dev/photos.getUserPhotos
type PhotosGetUserPhotosBulder struct {
	api.Params
}

// NewPhotosGetUserPhotosBulder func
func NewPhotosGetUserPhotosBulder() *PhotosGetUserPhotosBulder {
	return &PhotosGetUserPhotosBulder{api.Params{}}
}

// UserID User ID.
func (b *PhotosGetUserPhotosBulder) UserID(v int) {
	b.Params["user_id"] = v
}

// Offset Offset needed to return a specific subset of photos. By default, '0'.
func (b *PhotosGetUserPhotosBulder) Offset(v int) {
	b.Params["offset"] = v
}

// Count Number of photos to return. Maximum value is 1000.
func (b *PhotosGetUserPhotosBulder) Count(v int) {
	b.Params["count"] = v
}

// Extended '1' — to return an additional 'likes' field, '0' — (default)
func (b *PhotosGetUserPhotosBulder) Extended(v bool) {
	b.Params["extended"] = v
}

// Sort Sort order: '1' — by date the tag was added in ascending order, '0' — by date the tag was added in descending order
func (b *PhotosGetUserPhotosBulder) Sort(v string) {
	b.Params["sort"] = v
}

// PhotosGetWallUploadServerBulder builder
//
// Returns the server address for photo upload onto a user's wall.
//
// https://vk.com/dev/photos.getWallUploadServer
type PhotosGetWallUploadServerBulder struct {
	api.Params
}

// NewPhotosGetWallUploadServerBulder func
func NewPhotosGetWallUploadServerBulder() *PhotosGetWallUploadServerBulder {
	return &PhotosGetWallUploadServerBulder{api.Params{}}
}

// GroupID ID of community to whose wall the photo will be uploaded.
func (b *PhotosGetWallUploadServerBulder) GroupID(v int) {
	b.Params["group_id"] = v
}

// PhotosMakeCoverBulder builder
//
// Makes a photo into an album cover.
//
// https://vk.com/dev/photos.makeCover
type PhotosMakeCoverBulder struct {
	api.Params
}

// NewPhotosMakeCoverBulder func
func NewPhotosMakeCoverBulder() *PhotosMakeCoverBulder {
	return &PhotosMakeCoverBulder{api.Params{}}
}

// OwnerID ID of the user or community that owns the photo.
func (b *PhotosMakeCoverBulder) OwnerID(v int) {
	b.Params["owner_id"] = v
}

// PhotoID Photo ID.
func (b *PhotosMakeCoverBulder) PhotoID(v int) {
	b.Params["photo_id"] = v
}

// AlbumID Album ID.
func (b *PhotosMakeCoverBulder) AlbumID(v int) {
	b.Params["album_id"] = v
}

// PhotosMoveBulder builder
//
// Moves a photo from one album to another.
//
// https://vk.com/dev/photos.move
type PhotosMoveBulder struct {
	api.Params
}

// NewPhotosMoveBulder func
func NewPhotosMoveBulder() *PhotosMoveBulder {
	return &PhotosMoveBulder{api.Params{}}
}

// OwnerID ID of the user or community that owns the photo.
func (b *PhotosMoveBulder) OwnerID(v int) {
	b.Params["owner_id"] = v
}

// TargetAlbumID ID of the album to which the photo will be moved.
func (b *PhotosMoveBulder) TargetAlbumID(v int) {
	b.Params["target_album_id"] = v
}

// PhotoID Photo ID.
func (b *PhotosMoveBulder) PhotoID(v int) {
	b.Params["photo_id"] = v
}

// PhotosPutTagBulder builder
//
// Adds a tag on the photo.
//
// https://vk.com/dev/photos.putTag
type PhotosPutTagBulder struct {
	api.Params
}

// NewPhotosPutTagBulder func
func NewPhotosPutTagBulder() *PhotosPutTagBulder {
	return &PhotosPutTagBulder{api.Params{}}
}

// OwnerID ID of the user or community that owns the photo.
func (b *PhotosPutTagBulder) OwnerID(v int) {
	b.Params["owner_id"] = v
}

// PhotoID Photo ID.
func (b *PhotosPutTagBulder) PhotoID(v int) {
	b.Params["photo_id"] = v
}

// UserID ID of the user to be tagged.
func (b *PhotosPutTagBulder) UserID(v int) {
	b.Params["user_id"] = v
}

// X Upper left-corner coordinate of the tagged area (as a percentage of the photo's width).
func (b *PhotosPutTagBulder) X(v float64) {
	b.Params["x"] = v
}

// Y Upper left-corner coordinate of the tagged area (as a percentage of the photo's height).
func (b *PhotosPutTagBulder) Y(v float64) {
	b.Params["y"] = v
}

// X2 Lower right-corner coordinate of the tagged area (as a percentage of the photo's width).
func (b *PhotosPutTagBulder) X2(v float64) {
	b.Params["x2"] = v
}

// Y2 Lower right-corner coordinate of the tagged area (as a percentage of the photo's height).
func (b *PhotosPutTagBulder) Y2(v float64) {
	b.Params["y2"] = v
}

// PhotosRemoveTagBulder builder
//
// Removes a tag from a photo.
//
// https://vk.com/dev/photos.removeTag
type PhotosRemoveTagBulder struct {
	api.Params
}

// NewPhotosRemoveTagBulder func
func NewPhotosRemoveTagBulder() *PhotosRemoveTagBulder {
	return &PhotosRemoveTagBulder{api.Params{}}
}

// OwnerID ID of the user or community that owns the photo.
func (b *PhotosRemoveTagBulder) OwnerID(v int) {
	b.Params["owner_id"] = v
}

// PhotoID Photo ID.
func (b *PhotosRemoveTagBulder) PhotoID(v int) {
	b.Params["photo_id"] = v
}

// TagID Tag ID.
func (b *PhotosRemoveTagBulder) TagID(v int) {
	b.Params["tag_id"] = v
}

// PhotosReorderAlbumsBulder builder
//
// Reorders the album in the list of user albums.
//
// https://vk.com/dev/photos.reorderAlbums
type PhotosReorderAlbumsBulder struct {
	api.Params
}

// NewPhotosReorderAlbumsBulder func
func NewPhotosReorderAlbumsBulder() *PhotosReorderAlbumsBulder {
	return &PhotosReorderAlbumsBulder{api.Params{}}
}

// OwnerID ID of the user or community that owns the album.
func (b *PhotosReorderAlbumsBulder) OwnerID(v int) {
	b.Params["owner_id"] = v
}

// AlbumID Album ID.
func (b *PhotosReorderAlbumsBulder) AlbumID(v int) {
	b.Params["album_id"] = v
}

// Before ID of the album before which the album in question shall be placed.
func (b *PhotosReorderAlbumsBulder) Before(v int) {
	b.Params["before"] = v
}

// After ID of the album after which the album in question shall be placed.
func (b *PhotosReorderAlbumsBulder) After(v int) {
	b.Params["after"] = v
}

// PhotosReorderPhotosBulder builder
//
// Reorders the photo in the list of photos of the user album.
//
// https://vk.com/dev/photos.reorderPhotos
type PhotosReorderPhotosBulder struct {
	api.Params
}

// NewPhotosReorderPhotosBulder func
func NewPhotosReorderPhotosBulder() *PhotosReorderPhotosBulder {
	return &PhotosReorderPhotosBulder{api.Params{}}
}

// OwnerID ID of the user or community that owns the photo.
func (b *PhotosReorderPhotosBulder) OwnerID(v int) {
	b.Params["owner_id"] = v
}

// PhotoID Photo ID.
func (b *PhotosReorderPhotosBulder) PhotoID(v int) {
	b.Params["photo_id"] = v
}

// Before ID of the photo before which the photo in question shall be placed.
func (b *PhotosReorderPhotosBulder) Before(v int) {
	b.Params["before"] = v
}

// After ID of the photo after which the photo in question shall be placed.
func (b *PhotosReorderPhotosBulder) After(v int) {
	b.Params["after"] = v
}

// PhotosReportBulder builder
//
// Reports (submits a complaint about) a photo.
//
// https://vk.com/dev/photos.report
type PhotosReportBulder struct {
	api.Params
}

// NewPhotosReportBulder func
func NewPhotosReportBulder() *PhotosReportBulder {
	return &PhotosReportBulder{api.Params{}}
}

// OwnerID ID of the user or community that owns the photo.
func (b *PhotosReportBulder) OwnerID(v int) {
	b.Params["owner_id"] = v
}

// PhotoID Photo ID.
func (b *PhotosReportBulder) PhotoID(v int) {
	b.Params["photo_id"] = v
}

// Reason Reason for the complaint: '0' – spam, '1' – child pornography, '2' – extremism, '3' – violence, '4' – drug propaganda, '5' – adult material, '6' – insult, abuse
func (b *PhotosReportBulder) Reason(v int) {
	b.Params["reason"] = v
}

// PhotosReportCommentBulder builder
//
// Reports (submits a complaint about) a comment on a photo.
//
// https://vk.com/dev/photos.reportComment
type PhotosReportCommentBulder struct {
	api.Params
}

// NewPhotosReportCommentBulder func
func NewPhotosReportCommentBulder() *PhotosReportCommentBulder {
	return &PhotosReportCommentBulder{api.Params{}}
}

// OwnerID ID of the user or community that owns the photo.
func (b *PhotosReportCommentBulder) OwnerID(v int) {
	b.Params["owner_id"] = v
}

// CommentID ID of the comment being reported.
func (b *PhotosReportCommentBulder) CommentID(v int) {
	b.Params["comment_id"] = v
}

// Reason Reason for the complaint: '0' – spam, '1' – child pornography, '2' – extremism, '3' – violence, '4' – drug propaganda, '5' – adult material, '6' – insult, abuse
func (b *PhotosReportCommentBulder) Reason(v int) {
	b.Params["reason"] = v
}

// PhotosRestoreBulder builder
//
// Restores a deleted photo.
//
// https://vk.com/dev/photos.restore
type PhotosRestoreBulder struct {
	api.Params
}

// NewPhotosRestoreBulder func
func NewPhotosRestoreBulder() *PhotosRestoreBulder {
	return &PhotosRestoreBulder{api.Params{}}
}

// OwnerID ID of the user or community that owns the photo.
func (b *PhotosRestoreBulder) OwnerID(v int) {
	b.Params["owner_id"] = v
}

// PhotoID Photo ID.
func (b *PhotosRestoreBulder) PhotoID(v int) {
	b.Params["photo_id"] = v
}

// PhotosRestoreCommentBulder builder
//
// Restores a deleted comment on a photo.
//
// https://vk.com/dev/photos.restoreComment
type PhotosRestoreCommentBulder struct {
	api.Params
}

// NewPhotosRestoreCommentBulder func
func NewPhotosRestoreCommentBulder() *PhotosRestoreCommentBulder {
	return &PhotosRestoreCommentBulder{api.Params{}}
}

// OwnerID ID of the user or community that owns the photo.
func (b *PhotosRestoreCommentBulder) OwnerID(v int) {
	b.Params["owner_id"] = v
}

// CommentID ID of the deleted comment.
func (b *PhotosRestoreCommentBulder) CommentID(v int) {
	b.Params["comment_id"] = v
}

// PhotosSaveBulder builder
//
// Saves photos after successful uploading.
//
// https://vk.com/dev/photos.save
type PhotosSaveBulder struct {
	api.Params
}

// NewPhotosSaveBulder func
func NewPhotosSaveBulder() *PhotosSaveBulder {
	return &PhotosSaveBulder{api.Params{}}
}

// AlbumID ID of the album to save photos to.
func (b *PhotosSaveBulder) AlbumID(v int) {
	b.Params["album_id"] = v
}

// GroupID ID of the community to save photos to.
func (b *PhotosSaveBulder) GroupID(v int) {
	b.Params["group_id"] = v
}

// Server Parameter returned when photos are [vk.com/dev/upload_files|uploaded to server].
func (b *PhotosSaveBulder) Server(v int) {
	b.Params["server"] = v
}

// PhotosList Parameter returned when photos are [vk.com/dev/upload_files|uploaded to server].
func (b *PhotosSaveBulder) PhotosList(v string) {
	b.Params["photos_list"] = v
}

// Hash Parameter returned when photos are [vk.com/dev/upload_files|uploaded to server].
func (b *PhotosSaveBulder) Hash(v string) {
	b.Params["hash"] = v
}

// Latitude Geographical latitude, in degrees (from '-90' to '90').
func (b *PhotosSaveBulder) Latitude(v float64) {
	b.Params["latitude"] = v
}

// Longitude Geographical longitude, in degrees (from '-180' to '180').
func (b *PhotosSaveBulder) Longitude(v float64) {
	b.Params["longitude"] = v
}

// Caption Text describing the photo. 2048 digits max.
func (b *PhotosSaveBulder) Caption(v string) {
	b.Params["caption"] = v
}

// PhotosSaveMarketAlbumPhotoBulder builder
//
// Saves market album photos after successful uploading.
//
// https://vk.com/dev/photos.saveMarketAlbumPhoto
type PhotosSaveMarketAlbumPhotoBulder struct {
	api.Params
}

// NewPhotosSaveMarketAlbumPhotoBulder func
func NewPhotosSaveMarketAlbumPhotoBulder() *PhotosSaveMarketAlbumPhotoBulder {
	return &PhotosSaveMarketAlbumPhotoBulder{api.Params{}}
}

// GroupID Community ID.
func (b *PhotosSaveMarketAlbumPhotoBulder) GroupID(v int) {
	b.Params["group_id"] = v
}

// Photo Parameter returned when photos are [vk.com/dev/upload_files|uploaded to server].
func (b *PhotosSaveMarketAlbumPhotoBulder) Photo(v string) {
	b.Params["photo"] = v
}

// Server Parameter returned when photos are [vk.com/dev/upload_files|uploaded to server].
func (b *PhotosSaveMarketAlbumPhotoBulder) Server(v int) {
	b.Params["server"] = v
}

// Hash Parameter returned when photos are [vk.com/dev/upload_files|uploaded to server].
func (b *PhotosSaveMarketAlbumPhotoBulder) Hash(v string) {
	b.Params["hash"] = v
}

// PhotosSaveMarketPhotoBulder builder
//
// Saves market photos after successful uploading.
//
// https://vk.com/dev/photos.saveMarketPhoto
type PhotosSaveMarketPhotoBulder struct {
	api.Params
}

// NewPhotosSaveMarketPhotoBulder func
func NewPhotosSaveMarketPhotoBulder() *PhotosSaveMarketPhotoBulder {
	return &PhotosSaveMarketPhotoBulder{api.Params{}}
}

// GroupID Community ID.
func (b *PhotosSaveMarketPhotoBulder) GroupID(v int) {
	b.Params["group_id"] = v
}

// Photo Parameter returned when photos are [vk.com/dev/upload_files|uploaded to server].
func (b *PhotosSaveMarketPhotoBulder) Photo(v string) {
	b.Params["photo"] = v
}

// Server Parameter returned when photos are [vk.com/dev/upload_files|uploaded to server].
func (b *PhotosSaveMarketPhotoBulder) Server(v int) {
	b.Params["server"] = v
}

// Hash Parameter returned when photos are [vk.com/dev/upload_files|uploaded to server].
func (b *PhotosSaveMarketPhotoBulder) Hash(v string) {
	b.Params["hash"] = v
}

// CropData Parameter returned when photos are [vk.com/dev/upload_files|uploaded to server].
func (b *PhotosSaveMarketPhotoBulder) CropData(v string) {
	b.Params["crop_data"] = v
}

// CropHash Parameter returned when photos are [vk.com/dev/upload_files|uploaded to server].
func (b *PhotosSaveMarketPhotoBulder) CropHash(v string) {
	b.Params["crop_hash"] = v
}

// PhotosSaveMessagesPhotoBulder builder
//
// Saves a photo after being successfully uploaded. URL obtained with [vk.com/dev/photos.getMessagesUploadServer|photos.getMessagesUploadServer] method.
//
// https://vk.com/dev/photos.saveMessagesPhoto
type PhotosSaveMessagesPhotoBulder struct {
	api.Params
}

// NewPhotosSaveMessagesPhotoBulder func
func NewPhotosSaveMessagesPhotoBulder() *PhotosSaveMessagesPhotoBulder {
	return &PhotosSaveMessagesPhotoBulder{api.Params{}}
}

// Photo Parameter returned when the photo is [vk.com/dev/upload_files|uploaded to the server].
func (b *PhotosSaveMessagesPhotoBulder) Photo(v string) {
	b.Params["photo"] = v
}

// Server parameter
func (b *PhotosSaveMessagesPhotoBulder) Server(v int) {
	b.Params["server"] = v
}

// Hash parameter
func (b *PhotosSaveMessagesPhotoBulder) Hash(v string) {
	b.Params["hash"] = v
}

// PhotosSaveOwnerCoverPhotoBulder builder
//
// Saves cover photo after successful uploading.
//
// https://vk.com/dev/photos.saveOwnerCoverPhoto
type PhotosSaveOwnerCoverPhotoBulder struct {
	api.Params
}

// NewPhotosSaveOwnerCoverPhotoBulder func
func NewPhotosSaveOwnerCoverPhotoBulder() *PhotosSaveOwnerCoverPhotoBulder {
	return &PhotosSaveOwnerCoverPhotoBulder{api.Params{}}
}

// Hash Parameter returned when photos are [vk.com/dev/upload_files|uploaded to server].
func (b *PhotosSaveOwnerCoverPhotoBulder) Hash(v string) {
	b.Params["hash"] = v
}

// Photo Parameter returned when photos are [vk.com/dev/upload_files|uploaded to server].
func (b *PhotosSaveOwnerCoverPhotoBulder) Photo(v string) {
	b.Params["photo"] = v
}

// PhotosSaveOwnerPhotoBulder builder
//
// Saves a profile or community photo. Upload URL can be got with the [vk.com/dev/photos.getOwnerPhotoUploadServer|photos.getOwnerPhotoUploadServer] method.
//
// https://vk.com/dev/photos.saveOwnerPhoto
type PhotosSaveOwnerPhotoBulder struct {
	api.Params
}

// NewPhotosSaveOwnerPhotoBulder func
func NewPhotosSaveOwnerPhotoBulder() *PhotosSaveOwnerPhotoBulder {
	return &PhotosSaveOwnerPhotoBulder{api.Params{}}
}

// Server parameter returned after [vk.com/dev/upload_files|photo upload].
func (b *PhotosSaveOwnerPhotoBulder) Server(v string) {
	b.Params["server"] = v
}

// Hash parameter returned after [vk.com/dev/upload_files|photo upload].
func (b *PhotosSaveOwnerPhotoBulder) Hash(v string) {
	b.Params["hash"] = v
}

// Photo parameter returned after [vk.com/dev/upload_files|photo upload].
func (b *PhotosSaveOwnerPhotoBulder) Photo(v string) {
	b.Params["photo"] = v
}

// PhotosSaveWallPhotoBulder builder
//
// Saves a photo to a user's or community's wall after being uploaded.
//
// https://vk.com/dev/photos.saveWallPhoto
type PhotosSaveWallPhotoBulder struct {
	api.Params
}

// NewPhotosSaveWallPhotoBulder func
func NewPhotosSaveWallPhotoBulder() *PhotosSaveWallPhotoBulder {
	return &PhotosSaveWallPhotoBulder{api.Params{}}
}

// UserID ID of the user on whose wall the photo will be saved.
func (b *PhotosSaveWallPhotoBulder) UserID(v int) {
	b.Params["user_id"] = v
}

// GroupID ID of community on whose wall the photo will be saved.
func (b *PhotosSaveWallPhotoBulder) GroupID(v int) {
	b.Params["group_id"] = v
}

// Photo Parameter returned when the the photo is [vk.com/dev/upload_files|uploaded to the server].
func (b *PhotosSaveWallPhotoBulder) Photo(v string) {
	b.Params["photo"] = v
}

// Server parameter
func (b *PhotosSaveWallPhotoBulder) Server(v int) {
	b.Params["server"] = v
}

// Hash parameter
func (b *PhotosSaveWallPhotoBulder) Hash(v string) {
	b.Params["hash"] = v
}

// Latitude Geographical latitude, in degrees (from '-90' to '90').
func (b *PhotosSaveWallPhotoBulder) Latitude(v float64) {
	b.Params["latitude"] = v
}

// Longitude Geographical longitude, in degrees (from '-180' to '180').
func (b *PhotosSaveWallPhotoBulder) Longitude(v float64) {
	b.Params["longitude"] = v
}

// Caption Text describing the photo. 2048 digits max.
func (b *PhotosSaveWallPhotoBulder) Caption(v string) {
	b.Params["caption"] = v
}

// PhotosSearchBulder builder
//
// Returns a list of photos.
//
// https://vk.com/dev/photos.search
type PhotosSearchBulder struct {
	api.Params
}

// NewPhotosSearchBulder func
func NewPhotosSearchBulder() *PhotosSearchBulder {
	return &PhotosSearchBulder{api.Params{}}
}

// Q Search query string.
func (b *PhotosSearchBulder) Q(v string) {
	b.Params["q"] = v
}

// Lat Geographical latitude, in degrees (from '-90' to '90').
func (b *PhotosSearchBulder) Lat(v float64) {
	b.Params["lat"] = v
}

// Long Geographical longitude, in degrees (from '-180' to '180').
func (b *PhotosSearchBulder) Long(v float64) {
	b.Params["long"] = v
}

// StartTime parameter
func (b *PhotosSearchBulder) StartTime(v int) {
	b.Params["start_time"] = v
}

// EndTime parameter
func (b *PhotosSearchBulder) EndTime(v int) {
	b.Params["end_time"] = v
}

// Sort Sort order:
func (b *PhotosSearchBulder) Sort(v int) {
	b.Params["sort"] = v
}

// Offset Offset needed to return a specific subset of photos.
func (b *PhotosSearchBulder) Offset(v int) {
	b.Params["offset"] = v
}

// Count Number of photos to return.
func (b *PhotosSearchBulder) Count(v int) {
	b.Params["count"] = v
}

// Radius Radius of search in meters (works very approximately). Available values: '10', '100', '800', '6000', '50000'.
func (b *PhotosSearchBulder) Radius(v int) {
	b.Params["radius"] = v
}
