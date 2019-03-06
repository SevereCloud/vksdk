package api // import "github.com/severecloud/vksdk/api"

import (
	"encoding/json"

	"github.com/severecloud/vksdk/object"
)

// PhotosConfirmTagResponse struct
type PhotosConfirmTagResponse struct{}

// TODO: photos.confirmTag confirms a tag on a photo.
// https://vk.com/dev/photos.confirmTag

// PhotosCopyResponse struct
type PhotosCopyResponse struct{}

// TODO: photos.copy allows to copy a photo to the "Saved photos" album
// https://vk.com/dev/photos.copy

// PhotosCreateAlbumResponse struct
type PhotosCreateAlbumResponse struct{}

// TODO: photos.createAlbum creates an empty photo album.
// https://vk.com/dev/photos.createAlbum

// PhotosCreateCommentResponse struct
type PhotosCreateCommentResponse struct{}

// TODO: photos.createComment adds a new comment on the photo.
// https://vk.com/dev/photos.createComment

// PhotosDeleteResponse struct
type PhotosDeleteResponse struct{}

// TODO: photos.delete deletes a photo.
// https://vk.com/dev/photos.delete

// PhotosDeleteAlbumResponse struct
type PhotosDeleteAlbumResponse struct{}

// TODO: photos.deleteAlbum deletes a photo album belonging to the current user.
// https://vk.com/dev/photos.deleteAlbum

// PhotosDeleteCommentResponse struct
type PhotosDeleteCommentResponse struct{}

// TODO: photos.deleteComment deletes a comment on the photo.
// https://vk.com/dev/photos.deleteComment

// PhotosEditResponse struct
type PhotosEditResponse struct{}

// TODO: photos.edit edits the caption of a photo.
// https://vk.com/dev/photos.edit

// PhotosEditAlbumResponse struct
type PhotosEditAlbumResponse struct{}

// TODO: photos.editAlbum edits information about a photo album.
// https://vk.com/dev/photos.editAlbum

// PhotosEditCommentResponse struct
type PhotosEditCommentResponse struct{}

// TODO: photos.editComment edits a comment on a photo.
// https://vk.com/dev/photos.editComment

// PhotosGetResponse struct
type PhotosGetResponse struct{}

// TODO: photos.get returns a list of a user's or community's photos.
// https://vk.com/dev/photos.get

// PhotosGetAlbumsResponse struct
type PhotosGetAlbumsResponse struct{}

// TODO: photos.getAlbums returns a list of a user's or community's photo albums.
// https://vk.com/dev/photos.getAlbums

// PhotosGetAlbumsCountResponse struct
type PhotosGetAlbumsCountResponse struct{}

// TODO: photos.getAlbumsCount returns the number of photo albums belonging to a user or community.
// https://vk.com/dev/photos.getAlbumsCount

// PhotosGetAllResponse struct
type PhotosGetAllResponse struct{}

// TODO: photos.getAll returns a list of photos belonging to a user or community, in reverse chronological order.
// https://vk.com/dev/photos.getAll

// PhotosGetAllCommentsResponse struct
type PhotosGetAllCommentsResponse struct{}

// TODO: photos.getAllComments returns a list of comments on a specific photo album or all albums of the user sorted in reverse chronological order.
// https://vk.com/dev/photos.getAllComments

// PhotosGetByIDResponse struct
type PhotosGetByIDResponse struct{}

// TODO: photos.getById returns information about photos by their IDs.
// https://vk.com/dev/photos.getById

// PhotosGetChatUploadServerResponse struct
type PhotosGetChatUploadServerResponse struct{}

// TODO: photos.getChatUploadServer returns an upload link for chat cover pictures.
// https://vk.com/dev/photos.getChatUploadServer

// PhotosGetCommentsResponse struct
type PhotosGetCommentsResponse struct{}

// TODO: photos.getComments returns a list of comments on a photo.
// https://vk.com/dev/photos.getComments

// PhotosGetMarketAlbumUploadServerResponse struct
type PhotosGetMarketAlbumUploadServerResponse struct{}

// TODO: photos.getMarketAlbumUploadServer returns the server address for market album photo upload.
// https://vk.com/dev/photos.getMarketAlbumUploadServer

// PhotosGetMarketUploadServerResponse struct
type PhotosGetMarketUploadServerResponse struct{}

// TODO: photos.getMarketUploadServer returns the server address for market photo upload.
// https://vk.com/dev/photos.getMarketUploadServer

// PhotosGetMessagesUploadServerResponse struct
type PhotosGetMessagesUploadServerResponse struct{}

// TODO: photos.getMessagesUploadServer returns the server address for photo upload in a private message for a user.
// https://vk.com/dev/photos.getMessagesUploadServer

// PhotosGetNewTagsResponse struct
type PhotosGetNewTagsResponse struct{}

// TODO: photos.getNewTags returns a list of photos with tags that have not been viewed.
// https://vk.com/dev/photos.getNewTags

// PhotosGetOwnerCoverPhotoUploadServerResponse struct
type PhotosGetOwnerCoverPhotoUploadServerResponse struct{}

// TODO: photos.getOwnerCoverPhotoUploadServer receives server address for uploading community cover.
// https://vk.com/dev/photos.getOwnerCoverPhotoUploadServer

// PhotosGetOwnerPhotoUploadServerResponse struct
type PhotosGetOwnerPhotoUploadServerResponse struct{}

// TODO: photos.getOwnerPhotoUploadServer returns an upload server address for a profile or community photo.
// https://vk.com/dev/photos.getOwnerPhotoUploadServer

// PhotosGetTagsResponse struct
type PhotosGetTagsResponse struct{}

// TODO: photos.getTags returns a list of tags on a photo.
// https://vk.com/dev/photos.getTags

// PhotosGetUploadServerResponse struct
type PhotosGetUploadServerResponse struct{}

// TODO: photos.getUploadServer returns the server address for photo upload.
// https://vk.com/dev/photos.getUploadServer

// PhotosGetUserPhotosResponse struct
type PhotosGetUserPhotosResponse struct{}

// TODO: photos.getUserPhotos returns a list of photos in which a user is tagged.
// https://vk.com/dev/photos.getUserPhotos

// PhotosGetWallUploadServerResponse struct
type PhotosGetWallUploadServerResponse struct {
	AlbumID   int    `json:"album_id"`
	UploadURL string `json:"upload_url"`
	UserID    int    `json:"user_id"`
}

// PhotosGetWallUploadServer returns the server address for photo upload onto a user's wall.
// https://vk.com/dev/photos.getWallUploadServer
func (vk VK) PhotosGetWallUploadServer(params map[string]string) (response PhotosGetWallUploadServerResponse, vkErr Error) {
	rawResponse, vkErr := vk.Request("photos.getWallUploadServer", params)
	if vkErr.Code != 0 {
		return
	}

	err := json.Unmarshal(rawResponse, &response)
	if err != nil {
		panic(err)
	}

	return
}

// PhotosMakeCoverResponse struct
type PhotosMakeCoverResponse struct{}

// TODO: photos.makeCover makes a photo into an album cover.
// https://vk.com/dev/photos.makeCover

// PhotosMoveMovesResponse struct
type PhotosMoveMovesResponse struct{}

// TODO: photos.moveMoves a photo from one album to another.
// https://vk.com/dev/photos.moveMoves

// PhotosPutTagResponse struct
type PhotosPutTagResponse struct{}

// TODO: photos.putTag adds a tag on the photo.
// https://vk.com/dev/photos.putTag

// PhotosRemoveTagResponse struct
type PhotosRemoveTagResponse struct{}

// TODO: photos.removeTag removes a tag from a photo.
// https://vk.com/dev/photos.removeTag

// PhotosReorderAlbumsResponse struct
type PhotosReorderAlbumsResponse struct{}

// TODO: photos.reorderAlbums reorders the album in the list of user albums.
// https://vk.com/dev/photos.reorderAlbums

// PhotosReorderPhotosResponse struct
type PhotosReorderPhotosResponse struct{}

// TODO: photos.reorderPhotos reorders the photo in the list of photos of the user album.
// https://vk.com/dev/photos.reorderPhotos

// PhotosReportResponse struct
type PhotosReportResponse struct{}

// TODO: photos.report reports (submits a complaint about) a photo.
// https://vk.com/dev/photos.report

// PhotosReportCommentResponse struct
type PhotosReportCommentResponse struct{}

// TODO: photos.reportComment reports (submits a complaint about) a comment on a photo.
// https://vk.com/dev/photos.reportComment

// PhotosRestoreResponse struct
type PhotosRestoreResponse struct{}

// TODO: photos.restore restores a deleted photo.
// https://vk.com/dev/photos.restore

// PhotosRestoreCommentResponse struct
type PhotosRestoreCommentResponse struct{}

// TODO: photos.restoreComment restores a deleted comment on a photo.
// https://vk.com/dev/photos.restoreComment

// PhotosSaveResponse struct
type PhotosSaveResponse struct{}

// TODO: photos.save saves photos after successful uploading.
// https://vk.com/dev/photos.save

// PhotosSaveMarketAlbumResponse struct
type PhotosSaveMarketAlbumResponse struct{}

// TODO: photos.saveMarketAlbum photo Saves market album photos after successful uploading.
// https://vk.com/dev/photos.saveMarketAlbum

// PhotosSaveMarketPhotoResponse struct
type PhotosSaveMarketPhotoResponse struct{}

// TODO: photos.saveMarketPhoto saves market photos after successful uploading.
// https://vk.com/dev/photos.saveMarketPhoto

// PhotosSaveMessagesPhotoResponse struct
type PhotosSaveMessagesPhotoResponse struct{}

// TODO: photos.saveMessagesPhoto saves a photo after being successfully uploaded.URL obtained with photos.getMessagesUploadServer method.
// https://vk.com/dev/photos.saveMessagesPhoto

// PhotosSaveOwnerCoverPhotoResponse struct
type PhotosSaveOwnerCoverPhotoResponse struct{}

// TODO: photos.saveOwnerCoverPhoto saves cover photo after successful uploading.
// https://vk.com/dev/photos.saveOwnerCoverPhoto

// PhotosSaveOwnerPhotoResponse struct
type PhotosSaveOwnerPhotoResponse struct{}

// TODO: photos.saveOwnerPhoto saves a profile or community photo.
// https://vk.com/dev/photos.saveOwnerPhoto

// PhotosSaveWallPhotoResponse struct
type PhotosSaveWallPhotoResponse []object.PhotosPhoto

// PhotosSaveWallPhoto saves a photo to a user's or community's wall after being uploaded.
// https://vk.com/dev/photos.saveWallPhoto
func (vk VK) PhotosSaveWallPhoto(params map[string]string) (response PhotosSaveWallPhotoResponse, vkErr Error) {
	rawResponse, vkErr := vk.Request("photos.saveWallPhoto", params)
	if vkErr.Code != 0 {
		return
	}

	err := json.Unmarshal(rawResponse, &response)
	if err != nil {
		panic(err)
	}

	return
}

// PhotosSearchResponse struct
type PhotosSearchResponse struct{}

// TODO: photos.search returns a list of photos.
// https://vk.com/dev/photos.search
