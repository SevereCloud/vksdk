package api // import "github.com/SevereCloud/vksdk/api"

import (
	"github.com/SevereCloud/vksdk/object"
)

// PhotosConfirmTag confirms a tag on a photo.
//
// https://vk.com/dev/photos.confirmTag
func (vk *VK) PhotosConfirmTag(params map[string]string) (response int, err error) {
	err = vk.RequestUnmarshal("photos.confirmTag", params, &response)
	return
}

// PhotosCopy allows to copy a photo to the "Saved photos" album
//
// https://vk.com/dev/photos.copy
func (vk *VK) PhotosCopy(params map[string]string) (response int, err error) {
	err = vk.RequestUnmarshal("photos.copy", params, &response)
	return
}

// PhotosCreateAlbumResponse struct
type PhotosCreateAlbumResponse object.PhotosPhotoAlbumFull

// PhotosCreateAlbum creates an empty photo album.
//
// https://vk.com/dev/photos.createAlbum
func (vk *VK) PhotosCreateAlbum(params map[string]string) (response PhotosCreateAlbumResponse, err error) {
	err = vk.RequestUnmarshal("photos.createAlbum", params, &response)
	return
}

// PhotosCreateComment adds a new comment on the photo.
//
// https://vk.com/dev/photos.createComment
func (vk *VK) PhotosCreateComment(params map[string]string) (response int, err error) {
	err = vk.RequestUnmarshal("photos.createComment", params, &response)
	return
}

// PhotosDelete deletes a photo.
//
// https://vk.com/dev/photos.delete
func (vk *VK) PhotosDelete(params map[string]string) (response int, err error) {
	err = vk.RequestUnmarshal("photos.delete", params, &response)
	return
}

// PhotosDeleteAlbum deletes a photo album belonging to the current user.
//
// https://vk.com/dev/photos.deleteAlbum
func (vk *VK) PhotosDeleteAlbum(params map[string]string) (response int, err error) {
	err = vk.RequestUnmarshal("photos.deleteAlbum", params, &response)
	return
}

// PhotosDeleteComment deletes a comment on the photo.
//
// https://vk.com/dev/photos.deleteComment
func (vk *VK) PhotosDeleteComment(params map[string]string) (response int, err error) {
	err = vk.RequestUnmarshal("photos.deleteComment", params, &response)
	return
}

// PhotosEdit edits the caption of a photo.
//
// https://vk.com/dev/photos.edit
func (vk *VK) PhotosEdit(params map[string]string) (response int, err error) {
	err = vk.RequestUnmarshal("photos.edit", params, &response)
	return
}

// PhotosEditAlbum edits information about a photo album.
//
// https://vk.com/dev/photos.editAlbum
func (vk *VK) PhotosEditAlbum(params map[string]string) (response int, err error) {
	err = vk.RequestUnmarshal("photos.editAlbum", params, &response)
	return
}

// PhotosEditComment edits a comment on a photo.
//
// https://vk.com/dev/photos.editComment
func (vk *VK) PhotosEditComment(params map[string]string) (response int, err error) {
	err = vk.RequestUnmarshal("photos.editComment", params, &response)
	return
}

// PhotosGetResponse struct
type PhotosGetResponse struct {
	Count int                  `json:"count"` // Total number
	Items []object.PhotosPhoto `json:"items"`
}

// PhotosGet returns a list of a user's or community's photos.
//
// extended=0
//
// https://vk.com/dev/photos.get
func (vk *VK) PhotosGet(params map[string]string) (response PhotosGetResponse, err error) {
	params["extended"] = "0"
	err = vk.RequestUnmarshal("photos.get", params, &response)

	return
}

// PhotosGetExtendedResponse struct
type PhotosGetExtendedResponse struct {
	Count int                      `json:"count"` // Total number
	Items []object.PhotosPhotoFull `json:"items"`
}

// PhotosGetExtended returns a list of a user's or community's photos.
//
// extended=1
//
// https://vk.com/dev/photos.get
func (vk *VK) PhotosGetExtended(params map[string]string) (response PhotosGetExtendedResponse, err error) {
	params["extended"] = "1"
	err = vk.RequestUnmarshal("photos.get", params, &response)

	return
}

// PhotosGetAlbumsResponse struct
type PhotosGetAlbumsResponse struct {
	Count int                           `json:"count"` // Total number
	Items []object.PhotosPhotoAlbumFull `json:"items"`
}

// PhotosGetAlbums returns a list of a user's or community's photo albums.
//
// https://vk.com/dev/photos.getAlbums
func (vk *VK) PhotosGetAlbums(params map[string]string) (response PhotosGetAlbumsResponse, err error) {
	err = vk.RequestUnmarshal("photos.getAlbums", params, &response)
	return
}

// PhotosGetAlbumsCount returns the number of photo albums belonging to a user or community.
//
// https://vk.com/dev/photos.getAlbumsCount
func (vk *VK) PhotosGetAlbumsCount(params map[string]string) (response int, err error) {
	err = vk.RequestUnmarshal("photos.getAlbumsCount", params, &response)
	return
}

// PhotosGetAllResponse struct
type PhotosGetAllResponse struct {
	Count int                               `json:"count"` // Total number
	Items []object.PhotosPhotoXtrRealOffset `json:"items"`
	More  int                               `json:"more"` // Information whether next page is presented
}

// PhotosGetAll returns a list of photos belonging to a user or community, in reverse chronological order.
//
// extended=0
//
// https://vk.com/dev/photos.getAll
func (vk *VK) PhotosGetAll(params map[string]string) (response PhotosGetAllResponse, err error) {
	params["extended"] = "0"
	err = vk.RequestUnmarshal("photos.getAll", params, &response)

	return
}

// PhotosGetAllExtendedResponse struct
type PhotosGetAllExtendedResponse struct {
	Count int                                   `json:"count"` // Total number
	Items []object.PhotosPhotoFullXtrRealOffset `json:"items"`
	More  int                                   `json:"more"` // Information whether next page is presented
}

// PhotosGetAllExtended returns a list of photos belonging to a user or community, in reverse chronological order.
//
// extended=1
//
// https://vk.com/dev/photos.getAll
func (vk *VK) PhotosGetAllExtended(params map[string]string) (response PhotosGetAllExtendedResponse, err error) {
	params["extended"] = "1"
	err = vk.RequestUnmarshal("photos.getAll", params, &response)

	return
}

// PhotosGetAllCommentsResponse struct
type PhotosGetAllCommentsResponse struct {
	Count int                          `json:"count"` // Total number
	Items []object.PhotosCommentXtrPid `json:"items"`
}

// PhotosGetAllComments returns a list of comments on a specific
// photo album or all albums of the user sorted in reverse chronological order.
//
// https://vk.com/dev/photos.getAllComments
func (vk *VK) PhotosGetAllComments(params map[string]string) (response PhotosGetAllCommentsResponse, err error) {
	err = vk.RequestUnmarshal("photos.getAllComments", params, &response)
	return
}

// PhotosGetByIDResponse struct
type PhotosGetByIDResponse []object.PhotosPhoto

// PhotosGetByID returns information about photos by their IDs.
//
// extended=0
//
// https://vk.com/dev/photos.getById
func (vk *VK) PhotosGetByID(params map[string]string) (response PhotosGetByIDResponse, err error) {
	params["extended"] = "0"
	err = vk.RequestUnmarshal("photos.getById", params, &response)

	return
}

// PhotosGetByIDExtendedResponse struct
type PhotosGetByIDExtendedResponse []object.PhotosPhotoFull

// PhotosGetByIDExtended returns information about photos by their IDs.
//
// extended=1
//
// https://vk.com/dev/photos.getById
func (vk *VK) PhotosGetByIDExtended(params map[string]string) (response PhotosGetByIDExtendedResponse, err error) {
	params["extended"] = "1"
	err = vk.RequestUnmarshal("photos.getById", params, &response)

	return
}

// PhotosGetChatUploadServerResponse struct
type PhotosGetChatUploadServerResponse struct {
	UploadURL string `json:"upload_url"`
}

// PhotosGetChatUploadServer returns an upload link for chat cover pictures.
//
// https://vk.com/dev/photos.getChatUploadServer
func (vk *VK) PhotosGetChatUploadServer(params map[string]string) (response PhotosGetChatUploadServerResponse, err error) {
	err = vk.RequestUnmarshal("photos.getChatUploadServer", params, &response)
	return
}

// PhotosGetCommentsResponse struct
type PhotosGetCommentsResponse struct {
	Count      int                      `json:"count"`       // Total number
	RealOffset int                      `json:"real_offset"` // Real offset of the comments
	Items      []object.WallWallComment `json:"items"`
}

// PhotosGetComments returns a list of comments on a photo.
//
// extended=0
//
// https://vk.com/dev/photos.getComments
func (vk *VK) PhotosGetComments(params map[string]string) (response PhotosGetCommentsResponse, err error) {
	params["extended"] = "0"
	err = vk.RequestUnmarshal("photos.getComments", params, &response)

	return
}

// PhotosGetCommentsExtendedResponse struct
type PhotosGetCommentsExtendedResponse struct {
	Count      int                      `json:"count"`       // Total number
	RealOffset int                      `json:"real_offset"` // Real offset of the comments
	Items      []object.WallWallComment `json:"items"`
	Profiles   []object.UsersUser       `json:"profiles"`
	Groups     []object.GroupsGroup     `json:"groups"`
}

// PhotosGetCommentsExtended returns a list of comments on a photo.
//
// extended=1
//
// https://vk.com/dev/photos.getComments
func (vk *VK) PhotosGetCommentsExtended(params map[string]string) (response PhotosGetCommentsExtendedResponse, err error) {
	params["extended"] = "1"
	err = vk.RequestUnmarshal("photos.getComments", params, &response)

	return
}

// PhotosGetMarketAlbumUploadServerResponse struct
type PhotosGetMarketAlbumUploadServerResponse struct {
	UploadURL string `json:"upload_url"`
}

// PhotosGetMarketAlbumUploadServer returns the server address for market album photo upload.
//
// https://vk.com/dev/photos.getMarketAlbumUploadServer
func (vk *VK) PhotosGetMarketAlbumUploadServer(params map[string]string) (response PhotosGetMarketAlbumUploadServerResponse, err error) {
	err = vk.RequestUnmarshal("photos.getMarketAlbumUploadServer", params, &response)
	return
}

// PhotosGetMarketUploadServerResponse struct
type PhotosGetMarketUploadServerResponse struct {
	UploadURL string `json:"upload_url"`
}

// PhotosGetMarketUploadServer returns the server address for market photo upload.
//
// https://vk.com/dev/photos.getMarketUploadServer
func (vk *VK) PhotosGetMarketUploadServer(params map[string]string) (response PhotosGetMarketUploadServerResponse, err error) {
	err = vk.RequestUnmarshal("photos.getMarketUploadServer", params, &response)
	return
}

// PhotosGetMessagesUploadServerResponse struct
type PhotosGetMessagesUploadServerResponse struct {
	AlbumID   int    `json:"album_id"`
	UploadURL string `json:"upload_url"`
	UserID    int    `json:"user_id"`
}

// PhotosGetMessagesUploadServer returns the server address for photo upload onto a messages.
//
// https://vk.com/dev/photos.getMessagesUploadServer
func (vk *VK) PhotosGetMessagesUploadServer(params map[string]string) (response PhotosGetMessagesUploadServerResponse, err error) {
	err = vk.RequestUnmarshal("photos.getMessagesUploadServer", params, &response)
	return
}

// PhotosGetNewTagsResponse struct
type PhotosGetNewTagsResponse struct {
	Count int                            `json:"count"` // Total number
	Items []object.PhotosPhotoXtrTagInfo `json:"items"`
}

// PhotosGetNewTags returns a list of photos with tags that have not been viewed.
//
// https://vk.com/dev/photos.getNewTags
func (vk *VK) PhotosGetNewTags(params map[string]string) (response PhotosGetNewTagsResponse, err error) {
	err = vk.RequestUnmarshal("photos.getNewTags", params, &response)
	return
}

// PhotosGetOwnerCoverPhotoUploadServerResponse struct
type PhotosGetOwnerCoverPhotoUploadServerResponse struct {
	UploadURL string `json:"upload_url"`
}

// PhotosGetOwnerCoverPhotoUploadServer receives server address for uploading community cover.
//
// https://vk.com/dev/photos.getOwnerCoverPhotoUploadServer
func (vk *VK) PhotosGetOwnerCoverPhotoUploadServer(params map[string]string) (response PhotosGetOwnerCoverPhotoUploadServerResponse, err error) {
	err = vk.RequestUnmarshal("photos.getOwnerCoverPhotoUploadServer", params, &response)
	return
}

// PhotosGetOwnerPhotoUploadServerResponse struct
type PhotosGetOwnerPhotoUploadServerResponse struct {
	UploadURL string `json:"upload_url"`
}

// PhotosGetOwnerPhotoUploadServer returns an upload server address for a profile or community photo.
//
// https://vk.com/dev/photos.getOwnerPhotoUploadServer
func (vk *VK) PhotosGetOwnerPhotoUploadServer(params map[string]string) (response PhotosGetOwnerPhotoUploadServerResponse, err error) {
	err = vk.RequestUnmarshal("photos.getOwnerPhotoUploadServer", params, &response)
	return
}

// PhotosGetTagsResponse struct
type PhotosGetTagsResponse []object.PhotosPhotoTag

// PhotosGetTags returns a list of tags on a photo.
//
// https://vk.com/dev/photos.getTags
func (vk *VK) PhotosGetTags(params map[string]string) (response PhotosGetTagsResponse, err error) {
	err = vk.RequestUnmarshal("photos.getTags", params, &response)
	return
}

// PhotosGetUploadServerResponse struct
type PhotosGetUploadServerResponse object.PhotosPhotoUpload

// PhotosGetUploadServer returns the server address for photo upload.
//
// https://vk.com/dev/photos.getUploadServer
func (vk *VK) PhotosGetUploadServer(params map[string]string) (response PhotosGetUploadServerResponse, err error) {
	err = vk.RequestUnmarshal("photos.getUploadServer", params, &response)
	return
}

// PhotosGetUserPhotosResponse struct
type PhotosGetUserPhotosResponse struct {
	Count int                  `json:"count"` // Total number
	Items []object.PhotosPhoto `json:"items"`
}

// PhotosGetUserPhotos returns a list of photos in which a user is tagged.
//
// extended=0
//
// https://vk.com/dev/photos.getUserPhotos
func (vk *VK) PhotosGetUserPhotos(params map[string]string) (response PhotosGetUserPhotosResponse, err error) {
	params["extended"] = "0"
	err = vk.RequestUnmarshal("photos.getUserPhotos", params, &response)

	return
}

// PhotosGetUserPhotosExtendedResponse struct
type PhotosGetUserPhotosExtendedResponse struct {
	Count int                      `json:"count"` // Total number
	Items []object.PhotosPhotoFull `json:"items"`
}

// PhotosGetUserPhotosExtended returns a list of photos in which a user is tagged.
//
// extended=1
//
// https://vk.com/dev/photos.getUserPhotos
func (vk *VK) PhotosGetUserPhotosExtended(params map[string]string) (response PhotosGetUserPhotosExtendedResponse, err error) {
	params["extended"] = "1"
	err = vk.RequestUnmarshal("photos.getUserPhotos", params, &response)

	return
}

// PhotosGetWallUploadServerResponse struct
type PhotosGetWallUploadServerResponse object.PhotosPhotoUpload

// PhotosGetWallUploadServer returns the server address for photo upload onto a user's wall.
//
// https://vk.com/dev/photos.getWallUploadServer
func (vk *VK) PhotosGetWallUploadServer(params map[string]string) (response PhotosGetWallUploadServerResponse, err error) {
	err = vk.RequestUnmarshal("photos.getWallUploadServer", params, &response)
	return
}

// PhotosMakeCover makes a photo into an album cover.
//
// https://vk.com/dev/photos.makeCover
func (vk *VK) PhotosMakeCover(params map[string]string) (response int, err error) {
	err = vk.RequestUnmarshal("photos.makeCover", params, &response)
	return
}

// PhotosMoveMoves a photo from one album to another.
//
// https://vk.com/dev/photos.moveMoves
func (vk *VK) PhotosMoveMoves(params map[string]string) (response int, err error) {
	err = vk.RequestUnmarshal("photos.moveMoves", params, &response)
	return
}

// PhotosPutTag adds a tag on the photo.
//
// https://vk.com/dev/photos.putTag
func (vk *VK) PhotosPutTag(params map[string]string) (response int, err error) {
	err = vk.RequestUnmarshal("photos.putTag", params, &response)
	return
}

// PhotosRemoveTag removes a tag from a photo.
//
// https://vk.com/dev/photos.removeTag
func (vk *VK) PhotosRemoveTag(params map[string]string) (response int, err error) {
	err = vk.RequestUnmarshal("photos.removeTag", params, &response)
	return
}

// PhotosReorderAlbums reorders the album in the list of user albums.
//
// https://vk.com/dev/photos.reorderAlbums
func (vk *VK) PhotosReorderAlbums(params map[string]string) (response int, err error) {
	err = vk.RequestUnmarshal("photos.reorderAlbums", params, &response)
	return
}

// PhotosReorderPhotos reorders the photo in the list of photos of the user album.
//
// https://vk.com/dev/photos.reorderPhotos
func (vk *VK) PhotosReorderPhotos(params map[string]string) (response int, err error) {
	err = vk.RequestUnmarshal("photos.reorderPhotos", params, &response)
	return
}

// PhotosReport reports (submits a complaint about) a photo.
//
// https://vk.com/dev/photos.report
func (vk *VK) PhotosReport(params map[string]string) (response int, err error) {
	err = vk.RequestUnmarshal("photos.report", params, &response)
	return
}

// PhotosReportComment reports (submits a complaint about) a comment on a photo.
//
// https://vk.com/dev/photos.reportComment
func (vk *VK) PhotosReportComment(params map[string]string) (response int, err error) {
	err = vk.RequestUnmarshal("photos.reportComment", params, &response)
	return
}

// PhotosRestore restores a deleted photo.
//
// https://vk.com/dev/photos.restore
func (vk *VK) PhotosRestore(params map[string]string) (response int, err error) {
	err = vk.RequestUnmarshal("photos.restore", params, &response)
	return
}

// PhotosRestoreComment restores a deleted comment on a photo.
//
// https://vk.com/dev/photos.restoreComment
func (vk *VK) PhotosRestoreComment(params map[string]string) (response int, err error) {
	err = vk.RequestUnmarshal("photos.restoreComment", params, &response)
	return
}

// PhotosSaveResponse struct
type PhotosSaveResponse []object.PhotosPhoto

// PhotosSave saves photos after successful uploading.
//
// https://vk.com/dev/photos.save
func (vk *VK) PhotosSave(params map[string]string) (response PhotosSaveResponse, err error) {
	err = vk.RequestUnmarshal("photos.save", params, &response)
	return
}

// PhotosSaveMarketAlbumPhotoResponse struct
type PhotosSaveMarketAlbumPhotoResponse []object.PhotosPhoto

// PhotosSaveMarketAlbumPhoto photo Saves market album photos after successful uploading.
//
// https://vk.com/dev/photos.saveMarketAlbumPhoto
func (vk *VK) PhotosSaveMarketAlbumPhoto(params map[string]string) (response PhotosSaveMarketAlbumPhotoResponse, err error) {
	err = vk.RequestUnmarshal("photos.saveMarketAlbumPhoto", params, &response)
	return
}

// PhotosSaveMarketPhotoResponse struct
type PhotosSaveMarketPhotoResponse []object.PhotosPhoto

// PhotosSaveMarketPhoto saves market photos after successful uploading.
//
// https://vk.com/dev/photos.saveMarketPhoto
func (vk *VK) PhotosSaveMarketPhoto(params map[string]string) (response PhotosSaveMarketPhotoResponse, err error) {
	err = vk.RequestUnmarshal("photos.saveMarketPhoto", params, &response)
	return
}

// PhotosSaveMessagesPhotoResponse struct
type PhotosSaveMessagesPhotoResponse []object.PhotosPhoto

// PhotosSaveMessagesPhoto saves a photo after being successfully
//
// https://vk.com/dev/photos.saveMessagesPhoto
func (vk *VK) PhotosSaveMessagesPhoto(params map[string]string) (response PhotosSaveMessagesPhotoResponse, err error) {
	err = vk.RequestUnmarshal("photos.saveMessagesPhoto", params, &response)
	return
}

// PhotosSaveOwnerCoverPhotoResponse struct
type PhotosSaveOwnerCoverPhotoResponse struct {
	Images []object.PhotosImage `json:"images"`
}

// PhotosSaveOwnerCoverPhoto saves cover photo after successful uploading.
//
// https://vk.com/dev/photos.saveOwnerCoverPhoto
func (vk *VK) PhotosSaveOwnerCoverPhoto(params map[string]string) (response PhotosSaveOwnerCoverPhotoResponse, err error) {
	err = vk.RequestUnmarshal("photos.saveOwnerCoverPhoto", params, &response)
	return
}

// PhotosSaveOwnerPhotoResponse struct
type PhotosSaveOwnerPhotoResponse struct {
	PhotoHash     string `json:"photo_hash"`
	PhotoSrc      string `json:"photo_src"`
	PhotoSrcBig   string `json:"photo_src_big"`
	PhotoSrcSmall string `json:"photo_src_small"`
	Saved         int    `json:"saved"`
	PostID        int    `json:"post_id"`
}

// PhotosSaveOwnerPhoto saves a profile or community photo.
//
// https://vk.com/dev/photos.saveOwnerPhoto
func (vk *VK) PhotosSaveOwnerPhoto(params map[string]string) (response PhotosSaveOwnerPhotoResponse, err error) {
	err = vk.RequestUnmarshal("photos.saveOwnerPhoto", params, &response)
	return
}

// PhotosSaveWallPhotoResponse struct
type PhotosSaveWallPhotoResponse []object.PhotosPhoto

// PhotosSaveWallPhoto saves a photo to a user's or community's wall after being uploaded.
//
// https://vk.com/dev/photos.saveWallPhoto
func (vk *VK) PhotosSaveWallPhoto(params map[string]string) (response PhotosSaveWallPhotoResponse, err error) {
	err = vk.RequestUnmarshal("photos.saveWallPhoto", params, &response)
	return
}

// PhotosSearchResponse struct
type PhotosSearchResponse struct {
	Count int                      `json:"count"` // Total number
	Items []object.PhotosPhotoFull `json:"items"`
}

// PhotosSearch returns a list of photos.
//
// https://vk.com/dev/photos.search
func (vk *VK) PhotosSearch(params map[string]string) (response PhotosSearchResponse, err error) {
	err = vk.RequestUnmarshal("photos.search", params, &response)
	return
}
