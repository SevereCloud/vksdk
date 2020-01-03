package params // import "github.com/SevereCloud/vksdk/api/params"

import (
	"github.com/SevereCloud/vksdk/api"
)

// PhotosConfirmTagBuilder builder
//
// Confirms a tag on a photo.
//
// https://vk.com/dev/photos.confirmTag
type PhotosConfirmTagBuilder struct {
	api.Params
}

// NewPhotosConfirmTagBuilder func
func NewPhotosConfirmTagBuilder() *PhotosConfirmTagBuilder {
	return &PhotosConfirmTagBuilder{api.Params{}}
}

// OwnerID ID of the user or community that owns the photo.
func (b *PhotosConfirmTagBuilder) OwnerID(v int) {
	b.Params["owner_id"] = v
}

// PhotoID Photo ID.
func (b *PhotosConfirmTagBuilder) PhotoID(v string) {
	b.Params["photo_id"] = v
}

// TagID Tag ID.
func (b *PhotosConfirmTagBuilder) TagID(v int) {
	b.Params["tag_id"] = v
}

// PhotosCopyBuilder builder
//
// Allows to copy a photo to the "Saved photos" album
//
// https://vk.com/dev/photos.copy
type PhotosCopyBuilder struct {
	api.Params
}

// NewPhotosCopyBuilder func
func NewPhotosCopyBuilder() *PhotosCopyBuilder {
	return &PhotosCopyBuilder{api.Params{}}
}

// OwnerID photo's owner ID
func (b *PhotosCopyBuilder) OwnerID(v int) {
	b.Params["owner_id"] = v
}

// PhotoID photo ID
func (b *PhotosCopyBuilder) PhotoID(v int) {
	b.Params["photo_id"] = v
}

// AccessKey for private photos
func (b *PhotosCopyBuilder) AccessKey(v string) {
	b.Params["access_key"] = v
}

// PhotosCreateAlbumBuilder builder
//
// Creates an empty photo album.
//
// https://vk.com/dev/photos.createAlbum
type PhotosCreateAlbumBuilder struct {
	api.Params
}

// NewPhotosCreateAlbumBuilder func
func NewPhotosCreateAlbumBuilder() *PhotosCreateAlbumBuilder {
	return &PhotosCreateAlbumBuilder{api.Params{}}
}

// Title Album title.
func (b *PhotosCreateAlbumBuilder) Title(v string) {
	b.Params["title"] = v
}

// GroupID ID of the community in which the album will be created.
func (b *PhotosCreateAlbumBuilder) GroupID(v int) {
	b.Params["group_id"] = v
}

// Description Album description.
func (b *PhotosCreateAlbumBuilder) Description(v string) {
	b.Params["description"] = v
}

// PrivacyView parameter
func (b *PhotosCreateAlbumBuilder) PrivacyView(v []string) {
	b.Params["privacy_view"] = v
}

// PrivacyComment parameter
func (b *PhotosCreateAlbumBuilder) PrivacyComment(v []string) {
	b.Params["privacy_comment"] = v
}

// UploadByAdminsOnly parameter
func (b *PhotosCreateAlbumBuilder) UploadByAdminsOnly(v bool) {
	b.Params["upload_by_admins_only"] = v
}

// CommentsDisabled parameter
func (b *PhotosCreateAlbumBuilder) CommentsDisabled(v bool) {
	b.Params["comments_disabled"] = v
}

// PhotosCreateCommentBuilder builder
//
// Adds a new comment on the photo.
//
// https://vk.com/dev/photos.createComment
type PhotosCreateCommentBuilder struct {
	api.Params
}

// NewPhotosCreateCommentBuilder func
func NewPhotosCreateCommentBuilder() *PhotosCreateCommentBuilder {
	return &PhotosCreateCommentBuilder{api.Params{}}
}

// OwnerID ID of the user or community that owns the photo.
func (b *PhotosCreateCommentBuilder) OwnerID(v int) {
	b.Params["owner_id"] = v
}

// PhotoID Photo ID.
func (b *PhotosCreateCommentBuilder) PhotoID(v int) {
	b.Params["photo_id"] = v
}

// Message Comment text.
func (b *PhotosCreateCommentBuilder) Message(v string) {
	b.Params["message"] = v
}

// Attachments (Required if 'message' is not set.) List of objects attached to the post, in the following format: "<owner_id>_<media_id>,<owner_id>_<media_id>", '' — Type of media attachment: 'photo' — photo, 'video' — video, 'audio' — audio, 'doc' — document, '<owner_id>' — Media attachment owner ID. '<media_id>' — Media attachment ID. Example: "photo100172_166443618,photo66748_265827614"
func (b *PhotosCreateCommentBuilder) Attachments(v []string) {
	b.Params["attachments"] = v
}

// FromGroup '1' — to post a comment from the community
func (b *PhotosCreateCommentBuilder) FromGroup(v bool) {
	b.Params["from_group"] = v
}

// ReplyToComment parameter
func (b *PhotosCreateCommentBuilder) ReplyToComment(v int) {
	b.Params["reply_to_comment"] = v
}

// StickerID parameter
func (b *PhotosCreateCommentBuilder) StickerID(v int) {
	b.Params["sticker_id"] = v
}

// AccessKey parameter
func (b *PhotosCreateCommentBuilder) AccessKey(v string) {
	b.Params["access_key"] = v
}

// GUID parameter
func (b *PhotosCreateCommentBuilder) GUID(v string) {
	b.Params["guid"] = v
}

// PhotosDeleteBuilder builder
//
// Deletes a photo.
//
// https://vk.com/dev/photos.delete
type PhotosDeleteBuilder struct {
	api.Params
}

// NewPhotosDeleteBuilder func
func NewPhotosDeleteBuilder() *PhotosDeleteBuilder {
	return &PhotosDeleteBuilder{api.Params{}}
}

// OwnerID ID of the user or community that owns the photo.
func (b *PhotosDeleteBuilder) OwnerID(v int) {
	b.Params["owner_id"] = v
}

// PhotoID Photo ID.
func (b *PhotosDeleteBuilder) PhotoID(v int) {
	b.Params["photo_id"] = v
}

// PhotosDeleteAlbumBuilder builder
//
// Deletes a photo album belonging to the current user.
//
// https://vk.com/dev/photos.deleteAlbum
type PhotosDeleteAlbumBuilder struct {
	api.Params
}

// NewPhotosDeleteAlbumBuilder func
func NewPhotosDeleteAlbumBuilder() *PhotosDeleteAlbumBuilder {
	return &PhotosDeleteAlbumBuilder{api.Params{}}
}

// AlbumID Album ID.
func (b *PhotosDeleteAlbumBuilder) AlbumID(v int) {
	b.Params["album_id"] = v
}

// GroupID ID of the community that owns the album.
func (b *PhotosDeleteAlbumBuilder) GroupID(v int) {
	b.Params["group_id"] = v
}

// PhotosDeleteCommentBuilder builder
//
// Deletes a comment on the photo.
//
// https://vk.com/dev/photos.deleteComment
type PhotosDeleteCommentBuilder struct {
	api.Params
}

// NewPhotosDeleteCommentBuilder func
func NewPhotosDeleteCommentBuilder() *PhotosDeleteCommentBuilder {
	return &PhotosDeleteCommentBuilder{api.Params{}}
}

// OwnerID ID of the user or community that owns the photo.
func (b *PhotosDeleteCommentBuilder) OwnerID(v int) {
	b.Params["owner_id"] = v
}

// CommentID Comment ID.
func (b *PhotosDeleteCommentBuilder) CommentID(v int) {
	b.Params["comment_id"] = v
}

// PhotosEditBuilder builder
//
// Edits the caption of a photo.
//
// https://vk.com/dev/photos.edit
type PhotosEditBuilder struct {
	api.Params
}

// NewPhotosEditBuilder func
func NewPhotosEditBuilder() *PhotosEditBuilder {
	return &PhotosEditBuilder{api.Params{}}
}

// OwnerID ID of the user or community that owns the photo.
func (b *PhotosEditBuilder) OwnerID(v int) {
	b.Params["owner_id"] = v
}

// PhotoID Photo ID.
func (b *PhotosEditBuilder) PhotoID(v int) {
	b.Params["photo_id"] = v
}

// Caption New caption for the photo. If this parameter is not set, it is considered to be equal to an empty string.
func (b *PhotosEditBuilder) Caption(v string) {
	b.Params["caption"] = v
}

// Latitude parameter
func (b *PhotosEditBuilder) Latitude(v float64) {
	b.Params["latitude"] = v
}

// Longitude parameter
func (b *PhotosEditBuilder) Longitude(v float64) {
	b.Params["longitude"] = v
}

// PlaceStr parameter
func (b *PhotosEditBuilder) PlaceStr(v string) {
	b.Params["place_str"] = v
}

// FoursquareID parameter
func (b *PhotosEditBuilder) FoursquareID(v string) {
	b.Params["foursquare_id"] = v
}

// DeletePlace parameter
func (b *PhotosEditBuilder) DeletePlace(v bool) {
	b.Params["delete_place"] = v
}

// PhotosEditAlbumBuilder builder
//
// Edits information about a photo album.
//
// https://vk.com/dev/photos.editAlbum
type PhotosEditAlbumBuilder struct {
	api.Params
}

// NewPhotosEditAlbumBuilder func
func NewPhotosEditAlbumBuilder() *PhotosEditAlbumBuilder {
	return &PhotosEditAlbumBuilder{api.Params{}}
}

// AlbumID ID of the photo album to be edited.
func (b *PhotosEditAlbumBuilder) AlbumID(v int) {
	b.Params["album_id"] = v
}

// Title New album title.
func (b *PhotosEditAlbumBuilder) Title(v string) {
	b.Params["title"] = v
}

// Description New album description.
func (b *PhotosEditAlbumBuilder) Description(v string) {
	b.Params["description"] = v
}

// OwnerID ID of the user or community that owns the album.
func (b *PhotosEditAlbumBuilder) OwnerID(v int) {
	b.Params["owner_id"] = v
}

// PrivacyView parameter
func (b *PhotosEditAlbumBuilder) PrivacyView(v []string) {
	b.Params["privacy_view"] = v
}

// PrivacyComment parameter
func (b *PhotosEditAlbumBuilder) PrivacyComment(v []string) {
	b.Params["privacy_comment"] = v
}

// UploadByAdminsOnly parameter
func (b *PhotosEditAlbumBuilder) UploadByAdminsOnly(v bool) {
	b.Params["upload_by_admins_only"] = v
}

// CommentsDisabled parameter
func (b *PhotosEditAlbumBuilder) CommentsDisabled(v bool) {
	b.Params["comments_disabled"] = v
}

// PhotosEditCommentBuilder builder
//
// Edits a comment on a photo.
//
// https://vk.com/dev/photos.editComment
type PhotosEditCommentBuilder struct {
	api.Params
}

// NewPhotosEditCommentBuilder func
func NewPhotosEditCommentBuilder() *PhotosEditCommentBuilder {
	return &PhotosEditCommentBuilder{api.Params{}}
}

// OwnerID ID of the user or community that owns the photo.
func (b *PhotosEditCommentBuilder) OwnerID(v int) {
	b.Params["owner_id"] = v
}

// CommentID Comment ID.
func (b *PhotosEditCommentBuilder) CommentID(v int) {
	b.Params["comment_id"] = v
}

// Message New text of the comment.
func (b *PhotosEditCommentBuilder) Message(v string) {
	b.Params["message"] = v
}

// Attachments (Required if 'message' is not set.) List of objects attached to the post, in the following format: "<owner_id>_<media_id>,<owner_id>_<media_id>", '' — Type of media attachment: 'photo' — photo, 'video' — video, 'audio' — audio, 'doc' — document, '<owner_id>' — Media attachment owner ID. '<media_id>' — Media attachment ID. Example: "photo100172_166443618,photo66748_265827614"
func (b *PhotosEditCommentBuilder) Attachments(v []string) {
	b.Params["attachments"] = v
}

// PhotosGetBuilder builder
//
// Returns a list of a user's or community's photos.
//
// https://vk.com/dev/photos.get
type PhotosGetBuilder struct {
	api.Params
}

// NewPhotosGetBuilder func
func NewPhotosGetBuilder() *PhotosGetBuilder {
	return &PhotosGetBuilder{api.Params{}}
}

// OwnerID ID of the user or community that owns the photos. Use a negative value to designate a community ID.
func (b *PhotosGetBuilder) OwnerID(v int) {
	b.Params["owner_id"] = v
}

// AlbumID Photo album ID. To return information about photos from service albums, use the following string values: 'profile, wall, saved'.
func (b *PhotosGetBuilder) AlbumID(v string) {
	b.Params["album_id"] = v
}

// PhotoIDs Photo IDs.
func (b *PhotosGetBuilder) PhotoIDs(v []string) {
	b.Params["photo_ids"] = v
}

// Rev Sort order: '1' — reverse chronological, '0' — chronological
func (b *PhotosGetBuilder) Rev(v bool) {
	b.Params["rev"] = v
}

// Extended '1' — to return additional 'likes', 'comments', and 'tags' fields, '0' — (default)
func (b *PhotosGetBuilder) Extended(v bool) {
	b.Params["extended"] = v
}

// FeedType Type of feed obtained in 'feed' field of the method.
func (b *PhotosGetBuilder) FeedType(v string) {
	b.Params["feed_type"] = v
}

// Feed unixtime, that can be obtained with [vk.com/dev/newsfeed.get|newsfeed.get] method in date field to get all photos uploaded by the user on a specific day, or photos the user has been tagged on. Also, 'uid' parameter of the user the event happened with shall be specified.
func (b *PhotosGetBuilder) Feed(v int) {
	b.Params["feed"] = v
}

// PhotoSizes '1' — to return photo sizes in a [vk.com/dev/photo_sizes|special format]
func (b *PhotosGetBuilder) PhotoSizes(v bool) {
	b.Params["photo_sizes"] = v
}

// Offset parameter
func (b *PhotosGetBuilder) Offset(v int) {
	b.Params["offset"] = v
}

// Count parameter
func (b *PhotosGetBuilder) Count(v int) {
	b.Params["count"] = v
}

// PhotosGetAlbumsBuilder builder
//
// Returns a list of a user's or community's photo albums.
//
// https://vk.com/dev/photos.getAlbums
type PhotosGetAlbumsBuilder struct {
	api.Params
}

// NewPhotosGetAlbumsBuilder func
func NewPhotosGetAlbumsBuilder() *PhotosGetAlbumsBuilder {
	return &PhotosGetAlbumsBuilder{api.Params{}}
}

// OwnerID ID of the user or community that owns the albums.
func (b *PhotosGetAlbumsBuilder) OwnerID(v int) {
	b.Params["owner_id"] = v
}

// AlbumIDs Album IDs.
func (b *PhotosGetAlbumsBuilder) AlbumIDs(v []int) {
	b.Params["album_ids"] = v
}

// Offset Offset needed to return a specific subset of albums.
func (b *PhotosGetAlbumsBuilder) Offset(v int) {
	b.Params["offset"] = v
}

// Count Number of albums to return.
func (b *PhotosGetAlbumsBuilder) Count(v int) {
	b.Params["count"] = v
}

// NeedSystem '1' — to return system albums with negative IDs
func (b *PhotosGetAlbumsBuilder) NeedSystem(v bool) {
	b.Params["need_system"] = v
}

// NeedCovers '1' — to return an additional 'thumb_src' field, '0' — (default)
func (b *PhotosGetAlbumsBuilder) NeedCovers(v bool) {
	b.Params["need_covers"] = v
}

// PhotoSizes '1' — to return photo sizes in a
func (b *PhotosGetAlbumsBuilder) PhotoSizes(v bool) {
	b.Params["photo_sizes"] = v
}

// PhotosGetAlbumsCountBuilder builder
//
// Returns the number of photo albums belonging to a user or community.
//
// https://vk.com/dev/photos.getAlbumsCount
type PhotosGetAlbumsCountBuilder struct {
	api.Params
}

// NewPhotosGetAlbumsCountBuilder func
func NewPhotosGetAlbumsCountBuilder() *PhotosGetAlbumsCountBuilder {
	return &PhotosGetAlbumsCountBuilder{api.Params{}}
}

// UserID User ID.
func (b *PhotosGetAlbumsCountBuilder) UserID(v int) {
	b.Params["user_id"] = v
}

// GroupID Community ID.
func (b *PhotosGetAlbumsCountBuilder) GroupID(v int) {
	b.Params["group_id"] = v
}

// PhotosGetAllBuilder builder
//
// Returns a list of photos belonging to a user or community, in reverse chronological order.
//
// https://vk.com/dev/photos.getAll
type PhotosGetAllBuilder struct {
	api.Params
}

// NewPhotosGetAllBuilder func
func NewPhotosGetAllBuilder() *PhotosGetAllBuilder {
	return &PhotosGetAllBuilder{api.Params{}}
}

// OwnerID ID of a user or community that owns the photos. Use a negative value to designate a community ID.
func (b *PhotosGetAllBuilder) OwnerID(v int) {
	b.Params["owner_id"] = v
}

// Extended '1' — to return detailed information about photos
func (b *PhotosGetAllBuilder) Extended(v bool) {
	b.Params["extended"] = v
}

// Offset Offset needed to return a specific subset of photos. By default, '0'.
func (b *PhotosGetAllBuilder) Offset(v int) {
	b.Params["offset"] = v
}

// Count Number of photos to return.
func (b *PhotosGetAllBuilder) Count(v int) {
	b.Params["count"] = v
}

// PhotoSizes '1' – to return image sizes in [vk.com/dev/photo_sizes|special format].
func (b *PhotosGetAllBuilder) PhotoSizes(v bool) {
	b.Params["photo_sizes"] = v
}

// NoServiceAlbums '1' – to return photos only from standard albums, '0' – to return all photos including those in service albums, e.g., 'My wall photos' (default)
func (b *PhotosGetAllBuilder) NoServiceAlbums(v bool) {
	b.Params["no_service_albums"] = v
}

// NeedHidden '1' – to show information about photos being hidden from the block above the wall.
func (b *PhotosGetAllBuilder) NeedHidden(v bool) {
	b.Params["need_hidden"] = v
}

// SkipHidden '1' – not to return photos being hidden from the block above the wall. Works only with owner_id>0, no_service_albums is ignored.
func (b *PhotosGetAllBuilder) SkipHidden(v bool) {
	b.Params["skip_hidden"] = v
}

// PhotosGetAllCommentsBuilder builder
//
// Returns a list of comments on a specific photo album or all albums of the user sorted in reverse chronological order.
//
// https://vk.com/dev/photos.getAllComments
type PhotosGetAllCommentsBuilder struct {
	api.Params
}

// NewPhotosGetAllCommentsBuilder func
func NewPhotosGetAllCommentsBuilder() *PhotosGetAllCommentsBuilder {
	return &PhotosGetAllCommentsBuilder{api.Params{}}
}

// OwnerID ID of the user or community that owns the album(s).
func (b *PhotosGetAllCommentsBuilder) OwnerID(v int) {
	b.Params["owner_id"] = v
}

// AlbumID Album ID. If the parameter is not set, comments on all of the user's albums will be returned.
func (b *PhotosGetAllCommentsBuilder) AlbumID(v int) {
	b.Params["album_id"] = v
}

// NeedLikes '1' — to return an additional 'likes' field, '0' — (default)
func (b *PhotosGetAllCommentsBuilder) NeedLikes(v bool) {
	b.Params["need_likes"] = v
}

// Offset Offset needed to return a specific subset of comments. By default, '0'.
func (b *PhotosGetAllCommentsBuilder) Offset(v int) {
	b.Params["offset"] = v
}

// Count Number of comments to return. By default, '20'. Maximum value, '100'.
func (b *PhotosGetAllCommentsBuilder) Count(v int) {
	b.Params["count"] = v
}

// PhotosGetByIDBuilder builder
//
// Returns information about photos by their IDs.
//
// https://vk.com/dev/photos.getById
type PhotosGetByIDBuilder struct {
	api.Params
}

// NewPhotosGetByIDBuilder func
func NewPhotosGetByIDBuilder() *PhotosGetByIDBuilder {
	return &PhotosGetByIDBuilder{api.Params{}}
}

// Photos IDs separated with a comma, that are IDs of users who posted photos and IDs of photos themselves with an underscore character between such IDs. To get information about a photo in the group album, you shall specify group ID instead of user ID. Example: "1_129207899,6492_135055734, , -20629724_271945303"
func (b *PhotosGetByIDBuilder) Photos(v []string) {
	b.Params["photos"] = v
}

// Extended '1' — to return additional fields, '0' — (default)
func (b *PhotosGetByIDBuilder) Extended(v bool) {
	b.Params["extended"] = v
}

// PhotoSizes '1' — to return photo sizes in a
func (b *PhotosGetByIDBuilder) PhotoSizes(v bool) {
	b.Params["photo_sizes"] = v
}

// PhotosGetChatUploadServerBuilder builder
//
// Returns an upload link for chat cover pictures.
//
// https://vk.com/dev/photos.getChatUploadServer
type PhotosGetChatUploadServerBuilder struct {
	api.Params
}

// NewPhotosGetChatUploadServerBuilder func
func NewPhotosGetChatUploadServerBuilder() *PhotosGetChatUploadServerBuilder {
	return &PhotosGetChatUploadServerBuilder{api.Params{}}
}

// ChatID ID of the chat for which you want to upload a cover photo.
func (b *PhotosGetChatUploadServerBuilder) ChatID(v int) {
	b.Params["chat_id"] = v
}

// CropX parameter
func (b *PhotosGetChatUploadServerBuilder) CropX(v int) {
	b.Params["crop_x"] = v
}

// CropY parameter
func (b *PhotosGetChatUploadServerBuilder) CropY(v int) {
	b.Params["crop_y"] = v
}

// CropWidth Width (in pixels) of the photo after cropping.
func (b *PhotosGetChatUploadServerBuilder) CropWidth(v int) {
	b.Params["crop_width"] = v
}

// PhotosGetCommentsBuilder builder
//
// Returns a list of comments on a photo.
//
// https://vk.com/dev/photos.getComments
type PhotosGetCommentsBuilder struct {
	api.Params
}

// NewPhotosGetCommentsBuilder func
func NewPhotosGetCommentsBuilder() *PhotosGetCommentsBuilder {
	return &PhotosGetCommentsBuilder{api.Params{}}
}

// OwnerID ID of the user or community that owns the photo.
func (b *PhotosGetCommentsBuilder) OwnerID(v int) {
	b.Params["owner_id"] = v
}

// PhotoID Photo ID.
func (b *PhotosGetCommentsBuilder) PhotoID(v int) {
	b.Params["photo_id"] = v
}

// NeedLikes '1' — to return an additional 'likes' field, '0' — (default)
func (b *PhotosGetCommentsBuilder) NeedLikes(v bool) {
	b.Params["need_likes"] = v
}

// StartCommentID parameter
func (b *PhotosGetCommentsBuilder) StartCommentID(v int) {
	b.Params["start_comment_id"] = v
}

// Offset Offset needed to return a specific subset of comments. By default, '0'.
func (b *PhotosGetCommentsBuilder) Offset(v int) {
	b.Params["offset"] = v
}

// Count Number of comments to return.
func (b *PhotosGetCommentsBuilder) Count(v int) {
	b.Params["count"] = v
}

// Sort Sort order: 'asc' — old first, 'desc' — new first
func (b *PhotosGetCommentsBuilder) Sort(v string) {
	b.Params["sort"] = v
}

// AccessKey parameter
func (b *PhotosGetCommentsBuilder) AccessKey(v string) {
	b.Params["access_key"] = v
}

// Extended parameter
func (b *PhotosGetCommentsBuilder) Extended(v bool) {
	b.Params["extended"] = v
}

// Fields parameter
func (b *PhotosGetCommentsBuilder) Fields(v []string) {
	b.Params["fields"] = v
}

// PhotosGetMarketAlbumUploadServerBuilder builder
//
// Returns the server address for market album photo upload.
//
// https://vk.com/dev/photos.getMarketAlbumUploadServer
type PhotosGetMarketAlbumUploadServerBuilder struct {
	api.Params
}

// NewPhotosGetMarketAlbumUploadServerBuilder func
func NewPhotosGetMarketAlbumUploadServerBuilder() *PhotosGetMarketAlbumUploadServerBuilder {
	return &PhotosGetMarketAlbumUploadServerBuilder{api.Params{}}
}

// GroupID Community ID.
func (b *PhotosGetMarketAlbumUploadServerBuilder) GroupID(v int) {
	b.Params["group_id"] = v
}

// PhotosGetMarketUploadServerBuilder builder
//
// Returns the server address for market photo upload.
//
// https://vk.com/dev/photos.getMarketUploadServer
type PhotosGetMarketUploadServerBuilder struct {
	api.Params
}

// NewPhotosGetMarketUploadServerBuilder func
func NewPhotosGetMarketUploadServerBuilder() *PhotosGetMarketUploadServerBuilder {
	return &PhotosGetMarketUploadServerBuilder{api.Params{}}
}

// GroupID Community ID.
func (b *PhotosGetMarketUploadServerBuilder) GroupID(v int) {
	b.Params["group_id"] = v
}

// MainPhoto '1' if you want to upload the main item photo.
func (b *PhotosGetMarketUploadServerBuilder) MainPhoto(v bool) {
	b.Params["main_photo"] = v
}

// CropX X coordinate of the crop left upper corner.
func (b *PhotosGetMarketUploadServerBuilder) CropX(v int) {
	b.Params["crop_x"] = v
}

// CropY Y coordinate of the crop left upper corner.
func (b *PhotosGetMarketUploadServerBuilder) CropY(v int) {
	b.Params["crop_y"] = v
}

// CropWidth Width of the cropped photo in px.
func (b *PhotosGetMarketUploadServerBuilder) CropWidth(v int) {
	b.Params["crop_width"] = v
}

// PhotosGetMessagesUploadServerBuilder builder
//
// Returns the server address for photo upload in a private message for a user.
//
// https://vk.com/dev/photos.getMessagesUploadServer
type PhotosGetMessagesUploadServerBuilder struct {
	api.Params
}

// NewPhotosGetMessagesUploadServerBuilder func
func NewPhotosGetMessagesUploadServerBuilder() *PhotosGetMessagesUploadServerBuilder {
	return &PhotosGetMessagesUploadServerBuilder{api.Params{}}
}

// PeerID Destination ID. "For user: 'User ID', e.g. '12345'. For chat: '2000000000' + 'Chat ID', e.g. '2000000001'. For community: '- Community ID', e.g. '-12345'. "
func (b *PhotosGetMessagesUploadServerBuilder) PeerID(v int) {
	b.Params["peer_id"] = v
}

// PhotosGetNewTagsBuilder builder
//
// Returns a list of photos with tags that have not been viewed.
//
// https://vk.com/dev/photos.getNewTags
type PhotosGetNewTagsBuilder struct {
	api.Params
}

// NewPhotosGetNewTagsBuilder func
func NewPhotosGetNewTagsBuilder() *PhotosGetNewTagsBuilder {
	return &PhotosGetNewTagsBuilder{api.Params{}}
}

// Offset Offset needed to return a specific subset of photos.
func (b *PhotosGetNewTagsBuilder) Offset(v int) {
	b.Params["offset"] = v
}

// Count Number of photos to return.
func (b *PhotosGetNewTagsBuilder) Count(v int) {
	b.Params["count"] = v
}

// PhotosGetOwnerCoverPhotoUploadServerBuilder builder
//
// Returns the server address for owner cover upload.
//
// https://vk.com/dev/photos.getOwnerCoverPhotoUploadServer
type PhotosGetOwnerCoverPhotoUploadServerBuilder struct {
	api.Params
}

// NewPhotosGetOwnerCoverPhotoUploadServerBuilder func
func NewPhotosGetOwnerCoverPhotoUploadServerBuilder() *PhotosGetOwnerCoverPhotoUploadServerBuilder {
	return &PhotosGetOwnerCoverPhotoUploadServerBuilder{api.Params{}}
}

// GroupID ID of community that owns the album (if the photo will be uploaded to a community album).
func (b *PhotosGetOwnerCoverPhotoUploadServerBuilder) GroupID(v int) {
	b.Params["group_id"] = v
}

// CropX X coordinate of the left-upper corner
func (b *PhotosGetOwnerCoverPhotoUploadServerBuilder) CropX(v int) {
	b.Params["crop_x"] = v
}

// CropY Y coordinate of the left-upper corner
func (b *PhotosGetOwnerCoverPhotoUploadServerBuilder) CropY(v int) {
	b.Params["crop_y"] = v
}

// CropX2 X coordinate of the right-bottom corner
func (b *PhotosGetOwnerCoverPhotoUploadServerBuilder) CropX2(v int) {
	b.Params["crop_x2"] = v
}

// CropY2 Y coordinate of the right-bottom corner
func (b *PhotosGetOwnerCoverPhotoUploadServerBuilder) CropY2(v int) {
	b.Params["crop_y2"] = v
}

// PhotosGetOwnerPhotoUploadServerBuilder builder
//
// Returns an upload server address for a profile or community photo.
//
// https://vk.com/dev/photos.getOwnerPhotoUploadServer
type PhotosGetOwnerPhotoUploadServerBuilder struct {
	api.Params
}

// NewPhotosGetOwnerPhotoUploadServerBuilder func
func NewPhotosGetOwnerPhotoUploadServerBuilder() *PhotosGetOwnerPhotoUploadServerBuilder {
	return &PhotosGetOwnerPhotoUploadServerBuilder{api.Params{}}
}

// OwnerID identifier of a community or current user. "Note that community id must be negative. 'owner_id=1' – user, 'owner_id=-1' – community, "
func (b *PhotosGetOwnerPhotoUploadServerBuilder) OwnerID(v int) {
	b.Params["owner_id"] = v
}

// PhotosGetTagsBuilder builder
//
// Returns a list of tags on a photo.
//
// https://vk.com/dev/photos.getTags
type PhotosGetTagsBuilder struct {
	api.Params
}

// NewPhotosGetTagsBuilder func
func NewPhotosGetTagsBuilder() *PhotosGetTagsBuilder {
	return &PhotosGetTagsBuilder{api.Params{}}
}

// OwnerID ID of the user or community that owns the photo.
func (b *PhotosGetTagsBuilder) OwnerID(v int) {
	b.Params["owner_id"] = v
}

// PhotoID Photo ID.
func (b *PhotosGetTagsBuilder) PhotoID(v int) {
	b.Params["photo_id"] = v
}

// AccessKey parameter
func (b *PhotosGetTagsBuilder) AccessKey(v string) {
	b.Params["access_key"] = v
}

// PhotosGetUploadServerBuilder builder
//
// Returns the server address for photo upload.
//
// https://vk.com/dev/photos.getUploadServer
type PhotosGetUploadServerBuilder struct {
	api.Params
}

// NewPhotosGetUploadServerBuilder func
func NewPhotosGetUploadServerBuilder() *PhotosGetUploadServerBuilder {
	return &PhotosGetUploadServerBuilder{api.Params{}}
}

// GroupID ID of community that owns the album (if the photo will be uploaded to a community album).
func (b *PhotosGetUploadServerBuilder) GroupID(v int) {
	b.Params["group_id"] = v
}

// AlbumID parameter
func (b *PhotosGetUploadServerBuilder) AlbumID(v int) {
	b.Params["album_id"] = v
}

// PhotosGetUserPhotosBuilder builder
//
// Returns a list of photos in which a user is tagged.
//
// https://vk.com/dev/photos.getUserPhotos
type PhotosGetUserPhotosBuilder struct {
	api.Params
}

// NewPhotosGetUserPhotosBuilder func
func NewPhotosGetUserPhotosBuilder() *PhotosGetUserPhotosBuilder {
	return &PhotosGetUserPhotosBuilder{api.Params{}}
}

// UserID User ID.
func (b *PhotosGetUserPhotosBuilder) UserID(v int) {
	b.Params["user_id"] = v
}

// Offset Offset needed to return a specific subset of photos. By default, '0'.
func (b *PhotosGetUserPhotosBuilder) Offset(v int) {
	b.Params["offset"] = v
}

// Count Number of photos to return. Maximum value is 1000.
func (b *PhotosGetUserPhotosBuilder) Count(v int) {
	b.Params["count"] = v
}

// Extended '1' — to return an additional 'likes' field, '0' — (default)
func (b *PhotosGetUserPhotosBuilder) Extended(v bool) {
	b.Params["extended"] = v
}

// Sort Sort order: '1' — by date the tag was added in ascending order, '0' — by date the tag was added in descending order
func (b *PhotosGetUserPhotosBuilder) Sort(v string) {
	b.Params["sort"] = v
}

// PhotosGetWallUploadServerBuilder builder
//
// Returns the server address for photo upload onto a user's wall.
//
// https://vk.com/dev/photos.getWallUploadServer
type PhotosGetWallUploadServerBuilder struct {
	api.Params
}

// NewPhotosGetWallUploadServerBuilder func
func NewPhotosGetWallUploadServerBuilder() *PhotosGetWallUploadServerBuilder {
	return &PhotosGetWallUploadServerBuilder{api.Params{}}
}

// GroupID ID of community to whose wall the photo will be uploaded.
func (b *PhotosGetWallUploadServerBuilder) GroupID(v int) {
	b.Params["group_id"] = v
}

// PhotosMakeCoverBuilder builder
//
// Makes a photo into an album cover.
//
// https://vk.com/dev/photos.makeCover
type PhotosMakeCoverBuilder struct {
	api.Params
}

// NewPhotosMakeCoverBuilder func
func NewPhotosMakeCoverBuilder() *PhotosMakeCoverBuilder {
	return &PhotosMakeCoverBuilder{api.Params{}}
}

// OwnerID ID of the user or community that owns the photo.
func (b *PhotosMakeCoverBuilder) OwnerID(v int) {
	b.Params["owner_id"] = v
}

// PhotoID Photo ID.
func (b *PhotosMakeCoverBuilder) PhotoID(v int) {
	b.Params["photo_id"] = v
}

// AlbumID Album ID.
func (b *PhotosMakeCoverBuilder) AlbumID(v int) {
	b.Params["album_id"] = v
}

// PhotosMoveBuilder builder
//
// Moves a photo from one album to another.
//
// https://vk.com/dev/photos.move
type PhotosMoveBuilder struct {
	api.Params
}

// NewPhotosMoveBuilder func
func NewPhotosMoveBuilder() *PhotosMoveBuilder {
	return &PhotosMoveBuilder{api.Params{}}
}

// OwnerID ID of the user or community that owns the photo.
func (b *PhotosMoveBuilder) OwnerID(v int) {
	b.Params["owner_id"] = v
}

// TargetAlbumID ID of the album to which the photo will be moved.
func (b *PhotosMoveBuilder) TargetAlbumID(v int) {
	b.Params["target_album_id"] = v
}

// PhotoID Photo ID.
func (b *PhotosMoveBuilder) PhotoID(v int) {
	b.Params["photo_id"] = v
}

// PhotosPutTagBuilder builder
//
// Adds a tag on the photo.
//
// https://vk.com/dev/photos.putTag
type PhotosPutTagBuilder struct {
	api.Params
}

// NewPhotosPutTagBuilder func
func NewPhotosPutTagBuilder() *PhotosPutTagBuilder {
	return &PhotosPutTagBuilder{api.Params{}}
}

// OwnerID ID of the user or community that owns the photo.
func (b *PhotosPutTagBuilder) OwnerID(v int) {
	b.Params["owner_id"] = v
}

// PhotoID Photo ID.
func (b *PhotosPutTagBuilder) PhotoID(v int) {
	b.Params["photo_id"] = v
}

// UserID ID of the user to be tagged.
func (b *PhotosPutTagBuilder) UserID(v int) {
	b.Params["user_id"] = v
}

// X Upper left-corner coordinate of the tagged area (as a percentage of the photo's width).
func (b *PhotosPutTagBuilder) X(v float64) {
	b.Params["x"] = v
}

// Y Upper left-corner coordinate of the tagged area (as a percentage of the photo's height).
func (b *PhotosPutTagBuilder) Y(v float64) {
	b.Params["y"] = v
}

// X2 Lower right-corner coordinate of the tagged area (as a percentage of the photo's width).
func (b *PhotosPutTagBuilder) X2(v float64) {
	b.Params["x2"] = v
}

// Y2 Lower right-corner coordinate of the tagged area (as a percentage of the photo's height).
func (b *PhotosPutTagBuilder) Y2(v float64) {
	b.Params["y2"] = v
}

// PhotosRemoveTagBuilder builder
//
// Removes a tag from a photo.
//
// https://vk.com/dev/photos.removeTag
type PhotosRemoveTagBuilder struct {
	api.Params
}

// NewPhotosRemoveTagBuilder func
func NewPhotosRemoveTagBuilder() *PhotosRemoveTagBuilder {
	return &PhotosRemoveTagBuilder{api.Params{}}
}

// OwnerID ID of the user or community that owns the photo.
func (b *PhotosRemoveTagBuilder) OwnerID(v int) {
	b.Params["owner_id"] = v
}

// PhotoID Photo ID.
func (b *PhotosRemoveTagBuilder) PhotoID(v int) {
	b.Params["photo_id"] = v
}

// TagID Tag ID.
func (b *PhotosRemoveTagBuilder) TagID(v int) {
	b.Params["tag_id"] = v
}

// PhotosReorderAlbumsBuilder builder
//
// Reorders the album in the list of user albums.
//
// https://vk.com/dev/photos.reorderAlbums
type PhotosReorderAlbumsBuilder struct {
	api.Params
}

// NewPhotosReorderAlbumsBuilder func
func NewPhotosReorderAlbumsBuilder() *PhotosReorderAlbumsBuilder {
	return &PhotosReorderAlbumsBuilder{api.Params{}}
}

// OwnerID ID of the user or community that owns the album.
func (b *PhotosReorderAlbumsBuilder) OwnerID(v int) {
	b.Params["owner_id"] = v
}

// AlbumID Album ID.
func (b *PhotosReorderAlbumsBuilder) AlbumID(v int) {
	b.Params["album_id"] = v
}

// Before ID of the album before which the album in question shall be placed.
func (b *PhotosReorderAlbumsBuilder) Before(v int) {
	b.Params["before"] = v
}

// After ID of the album after which the album in question shall be placed.
func (b *PhotosReorderAlbumsBuilder) After(v int) {
	b.Params["after"] = v
}

// PhotosReorderPhotosBuilder builder
//
// Reorders the photo in the list of photos of the user album.
//
// https://vk.com/dev/photos.reorderPhotos
type PhotosReorderPhotosBuilder struct {
	api.Params
}

// NewPhotosReorderPhotosBuilder func
func NewPhotosReorderPhotosBuilder() *PhotosReorderPhotosBuilder {
	return &PhotosReorderPhotosBuilder{api.Params{}}
}

// OwnerID ID of the user or community that owns the photo.
func (b *PhotosReorderPhotosBuilder) OwnerID(v int) {
	b.Params["owner_id"] = v
}

// PhotoID Photo ID.
func (b *PhotosReorderPhotosBuilder) PhotoID(v int) {
	b.Params["photo_id"] = v
}

// Before ID of the photo before which the photo in question shall be placed.
func (b *PhotosReorderPhotosBuilder) Before(v int) {
	b.Params["before"] = v
}

// After ID of the photo after which the photo in question shall be placed.
func (b *PhotosReorderPhotosBuilder) After(v int) {
	b.Params["after"] = v
}

// PhotosReportBuilder builder
//
// Reports (submits a complaint about) a photo.
//
// https://vk.com/dev/photos.report
type PhotosReportBuilder struct {
	api.Params
}

// NewPhotosReportBuilder func
func NewPhotosReportBuilder() *PhotosReportBuilder {
	return &PhotosReportBuilder{api.Params{}}
}

// OwnerID ID of the user or community that owns the photo.
func (b *PhotosReportBuilder) OwnerID(v int) {
	b.Params["owner_id"] = v
}

// PhotoID Photo ID.
func (b *PhotosReportBuilder) PhotoID(v int) {
	b.Params["photo_id"] = v
}

// Reason Reason for the complaint: '0' – spam, '1' – child pornography, '2' – extremism, '3' – violence, '4' – drug propaganda, '5' – adult material, '6' – insult, abuse
func (b *PhotosReportBuilder) Reason(v int) {
	b.Params["reason"] = v
}

// PhotosReportCommentBuilder builder
//
// Reports (submits a complaint about) a comment on a photo.
//
// https://vk.com/dev/photos.reportComment
type PhotosReportCommentBuilder struct {
	api.Params
}

// NewPhotosReportCommentBuilder func
func NewPhotosReportCommentBuilder() *PhotosReportCommentBuilder {
	return &PhotosReportCommentBuilder{api.Params{}}
}

// OwnerID ID of the user or community that owns the photo.
func (b *PhotosReportCommentBuilder) OwnerID(v int) {
	b.Params["owner_id"] = v
}

// CommentID ID of the comment being reported.
func (b *PhotosReportCommentBuilder) CommentID(v int) {
	b.Params["comment_id"] = v
}

// Reason Reason for the complaint: '0' – spam, '1' – child pornography, '2' – extremism, '3' – violence, '4' – drug propaganda, '5' – adult material, '6' – insult, abuse
func (b *PhotosReportCommentBuilder) Reason(v int) {
	b.Params["reason"] = v
}

// PhotosRestoreBuilder builder
//
// Restores a deleted photo.
//
// https://vk.com/dev/photos.restore
type PhotosRestoreBuilder struct {
	api.Params
}

// NewPhotosRestoreBuilder func
func NewPhotosRestoreBuilder() *PhotosRestoreBuilder {
	return &PhotosRestoreBuilder{api.Params{}}
}

// OwnerID ID of the user or community that owns the photo.
func (b *PhotosRestoreBuilder) OwnerID(v int) {
	b.Params["owner_id"] = v
}

// PhotoID Photo ID.
func (b *PhotosRestoreBuilder) PhotoID(v int) {
	b.Params["photo_id"] = v
}

// PhotosRestoreCommentBuilder builder
//
// Restores a deleted comment on a photo.
//
// https://vk.com/dev/photos.restoreComment
type PhotosRestoreCommentBuilder struct {
	api.Params
}

// NewPhotosRestoreCommentBuilder func
func NewPhotosRestoreCommentBuilder() *PhotosRestoreCommentBuilder {
	return &PhotosRestoreCommentBuilder{api.Params{}}
}

// OwnerID ID of the user or community that owns the photo.
func (b *PhotosRestoreCommentBuilder) OwnerID(v int) {
	b.Params["owner_id"] = v
}

// CommentID ID of the deleted comment.
func (b *PhotosRestoreCommentBuilder) CommentID(v int) {
	b.Params["comment_id"] = v
}

// PhotosSaveBuilder builder
//
// Saves photos after successful uploading.
//
// https://vk.com/dev/photos.save
type PhotosSaveBuilder struct {
	api.Params
}

// NewPhotosSaveBuilder func
func NewPhotosSaveBuilder() *PhotosSaveBuilder {
	return &PhotosSaveBuilder{api.Params{}}
}

// AlbumID ID of the album to save photos to.
func (b *PhotosSaveBuilder) AlbumID(v int) {
	b.Params["album_id"] = v
}

// GroupID ID of the community to save photos to.
func (b *PhotosSaveBuilder) GroupID(v int) {
	b.Params["group_id"] = v
}

// Server Parameter returned when photos are [vk.com/dev/upload_files|uploaded to server].
func (b *PhotosSaveBuilder) Server(v int) {
	b.Params["server"] = v
}

// PhotosList Parameter returned when photos are [vk.com/dev/upload_files|uploaded to server].
func (b *PhotosSaveBuilder) PhotosList(v string) {
	b.Params["photos_list"] = v
}

// Hash Parameter returned when photos are [vk.com/dev/upload_files|uploaded to server].
func (b *PhotosSaveBuilder) Hash(v string) {
	b.Params["hash"] = v
}

// Latitude Geographical latitude, in degrees (from '-90' to '90').
func (b *PhotosSaveBuilder) Latitude(v float64) {
	b.Params["latitude"] = v
}

// Longitude Geographical longitude, in degrees (from '-180' to '180').
func (b *PhotosSaveBuilder) Longitude(v float64) {
	b.Params["longitude"] = v
}

// Caption Text describing the photo. 2048 digits max.
func (b *PhotosSaveBuilder) Caption(v string) {
	b.Params["caption"] = v
}

// PhotosSaveMarketAlbumPhotoBuilder builder
//
// Saves market album photos after successful uploading.
//
// https://vk.com/dev/photos.saveMarketAlbumPhoto
type PhotosSaveMarketAlbumPhotoBuilder struct {
	api.Params
}

// NewPhotosSaveMarketAlbumPhotoBuilder func
func NewPhotosSaveMarketAlbumPhotoBuilder() *PhotosSaveMarketAlbumPhotoBuilder {
	return &PhotosSaveMarketAlbumPhotoBuilder{api.Params{}}
}

// GroupID Community ID.
func (b *PhotosSaveMarketAlbumPhotoBuilder) GroupID(v int) {
	b.Params["group_id"] = v
}

// Photo Parameter returned when photos are [vk.com/dev/upload_files|uploaded to server].
func (b *PhotosSaveMarketAlbumPhotoBuilder) Photo(v string) {
	b.Params["photo"] = v
}

// Server Parameter returned when photos are [vk.com/dev/upload_files|uploaded to server].
func (b *PhotosSaveMarketAlbumPhotoBuilder) Server(v int) {
	b.Params["server"] = v
}

// Hash Parameter returned when photos are [vk.com/dev/upload_files|uploaded to server].
func (b *PhotosSaveMarketAlbumPhotoBuilder) Hash(v string) {
	b.Params["hash"] = v
}

// PhotosSaveMarketPhotoBuilder builder
//
// Saves market photos after successful uploading.
//
// https://vk.com/dev/photos.saveMarketPhoto
type PhotosSaveMarketPhotoBuilder struct {
	api.Params
}

// NewPhotosSaveMarketPhotoBuilder func
func NewPhotosSaveMarketPhotoBuilder() *PhotosSaveMarketPhotoBuilder {
	return &PhotosSaveMarketPhotoBuilder{api.Params{}}
}

// GroupID Community ID.
func (b *PhotosSaveMarketPhotoBuilder) GroupID(v int) {
	b.Params["group_id"] = v
}

// Photo Parameter returned when photos are [vk.com/dev/upload_files|uploaded to server].
func (b *PhotosSaveMarketPhotoBuilder) Photo(v string) {
	b.Params["photo"] = v
}

// Server Parameter returned when photos are [vk.com/dev/upload_files|uploaded to server].
func (b *PhotosSaveMarketPhotoBuilder) Server(v int) {
	b.Params["server"] = v
}

// Hash Parameter returned when photos are [vk.com/dev/upload_files|uploaded to server].
func (b *PhotosSaveMarketPhotoBuilder) Hash(v string) {
	b.Params["hash"] = v
}

// CropData Parameter returned when photos are [vk.com/dev/upload_files|uploaded to server].
func (b *PhotosSaveMarketPhotoBuilder) CropData(v string) {
	b.Params["crop_data"] = v
}

// CropHash Parameter returned when photos are [vk.com/dev/upload_files|uploaded to server].
func (b *PhotosSaveMarketPhotoBuilder) CropHash(v string) {
	b.Params["crop_hash"] = v
}

// PhotosSaveMessagesPhotoBuilder builder
//
// Saves a photo after being successfully uploaded. URL obtained with [vk.com/dev/photos.getMessagesUploadServer|photos.getMessagesUploadServer] method.
//
// https://vk.com/dev/photos.saveMessagesPhoto
type PhotosSaveMessagesPhotoBuilder struct {
	api.Params
}

// NewPhotosSaveMessagesPhotoBuilder func
func NewPhotosSaveMessagesPhotoBuilder() *PhotosSaveMessagesPhotoBuilder {
	return &PhotosSaveMessagesPhotoBuilder{api.Params{}}
}

// Photo Parameter returned when the photo is [vk.com/dev/upload_files|uploaded to the server].
func (b *PhotosSaveMessagesPhotoBuilder) Photo(v string) {
	b.Params["photo"] = v
}

// Server parameter
func (b *PhotosSaveMessagesPhotoBuilder) Server(v int) {
	b.Params["server"] = v
}

// Hash parameter
func (b *PhotosSaveMessagesPhotoBuilder) Hash(v string) {
	b.Params["hash"] = v
}

// PhotosSaveOwnerCoverPhotoBuilder builder
//
// Saves cover photo after successful uploading.
//
// https://vk.com/dev/photos.saveOwnerCoverPhoto
type PhotosSaveOwnerCoverPhotoBuilder struct {
	api.Params
}

// NewPhotosSaveOwnerCoverPhotoBuilder func
func NewPhotosSaveOwnerCoverPhotoBuilder() *PhotosSaveOwnerCoverPhotoBuilder {
	return &PhotosSaveOwnerCoverPhotoBuilder{api.Params{}}
}

// Hash Parameter returned when photos are [vk.com/dev/upload_files|uploaded to server].
func (b *PhotosSaveOwnerCoverPhotoBuilder) Hash(v string) {
	b.Params["hash"] = v
}

// Photo Parameter returned when photos are [vk.com/dev/upload_files|uploaded to server].
func (b *PhotosSaveOwnerCoverPhotoBuilder) Photo(v string) {
	b.Params["photo"] = v
}

// PhotosSaveOwnerPhotoBuilder builder
//
// Saves a profile or community photo. Upload URL can be got with the [vk.com/dev/photos.getOwnerPhotoUploadServer|photos.getOwnerPhotoUploadServer] method.
//
// https://vk.com/dev/photos.saveOwnerPhoto
type PhotosSaveOwnerPhotoBuilder struct {
	api.Params
}

// NewPhotosSaveOwnerPhotoBuilder func
func NewPhotosSaveOwnerPhotoBuilder() *PhotosSaveOwnerPhotoBuilder {
	return &PhotosSaveOwnerPhotoBuilder{api.Params{}}
}

// Server parameter returned after [vk.com/dev/upload_files|photo upload].
func (b *PhotosSaveOwnerPhotoBuilder) Server(v string) {
	b.Params["server"] = v
}

// Hash parameter returned after [vk.com/dev/upload_files|photo upload].
func (b *PhotosSaveOwnerPhotoBuilder) Hash(v string) {
	b.Params["hash"] = v
}

// Photo parameter returned after [vk.com/dev/upload_files|photo upload].
func (b *PhotosSaveOwnerPhotoBuilder) Photo(v string) {
	b.Params["photo"] = v
}

// PhotosSaveWallPhotoBuilder builder
//
// Saves a photo to a user's or community's wall after being uploaded.
//
// https://vk.com/dev/photos.saveWallPhoto
type PhotosSaveWallPhotoBuilder struct {
	api.Params
}

// NewPhotosSaveWallPhotoBuilder func
func NewPhotosSaveWallPhotoBuilder() *PhotosSaveWallPhotoBuilder {
	return &PhotosSaveWallPhotoBuilder{api.Params{}}
}

// UserID ID of the user on whose wall the photo will be saved.
func (b *PhotosSaveWallPhotoBuilder) UserID(v int) {
	b.Params["user_id"] = v
}

// GroupID ID of community on whose wall the photo will be saved.
func (b *PhotosSaveWallPhotoBuilder) GroupID(v int) {
	b.Params["group_id"] = v
}

// Photo Parameter returned when the the photo is [vk.com/dev/upload_files|uploaded to the server].
func (b *PhotosSaveWallPhotoBuilder) Photo(v string) {
	b.Params["photo"] = v
}

// Server parameter
func (b *PhotosSaveWallPhotoBuilder) Server(v int) {
	b.Params["server"] = v
}

// Hash parameter
func (b *PhotosSaveWallPhotoBuilder) Hash(v string) {
	b.Params["hash"] = v
}

// Latitude Geographical latitude, in degrees (from '-90' to '90').
func (b *PhotosSaveWallPhotoBuilder) Latitude(v float64) {
	b.Params["latitude"] = v
}

// Longitude Geographical longitude, in degrees (from '-180' to '180').
func (b *PhotosSaveWallPhotoBuilder) Longitude(v float64) {
	b.Params["longitude"] = v
}

// Caption Text describing the photo. 2048 digits max.
func (b *PhotosSaveWallPhotoBuilder) Caption(v string) {
	b.Params["caption"] = v
}

// PhotosSearchBuilder builder
//
// Returns a list of photos.
//
// https://vk.com/dev/photos.search
type PhotosSearchBuilder struct {
	api.Params
}

// NewPhotosSearchBuilder func
func NewPhotosSearchBuilder() *PhotosSearchBuilder {
	return &PhotosSearchBuilder{api.Params{}}
}

// Q Search query string.
func (b *PhotosSearchBuilder) Q(v string) {
	b.Params["q"] = v
}

// Lat Geographical latitude, in degrees (from '-90' to '90').
func (b *PhotosSearchBuilder) Lat(v float64) {
	b.Params["lat"] = v
}

// Long Geographical longitude, in degrees (from '-180' to '180').
func (b *PhotosSearchBuilder) Long(v float64) {
	b.Params["long"] = v
}

// StartTime parameter
func (b *PhotosSearchBuilder) StartTime(v int) {
	b.Params["start_time"] = v
}

// EndTime parameter
func (b *PhotosSearchBuilder) EndTime(v int) {
	b.Params["end_time"] = v
}

// Sort Sort order:
func (b *PhotosSearchBuilder) Sort(v int) {
	b.Params["sort"] = v
}

// Offset Offset needed to return a specific subset of photos.
func (b *PhotosSearchBuilder) Offset(v int) {
	b.Params["offset"] = v
}

// Count Number of photos to return.
func (b *PhotosSearchBuilder) Count(v int) {
	b.Params["count"] = v
}

// Radius Radius of search in meters (works very approximately). Available values: '10', '100', '800', '6000', '50000'.
func (b *PhotosSearchBuilder) Radius(v int) {
	b.Params["radius"] = v
}
