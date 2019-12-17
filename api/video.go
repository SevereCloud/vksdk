package api // import "github.com/SevereCloud/vksdk/api"

import (
	"github.com/SevereCloud/vksdk/object"
)

// VideoAdd adds a video to a user or community page.
//
// https://vk.com/dev/video.add
func (vk *VK) VideoAdd(params Params) (response int, err error) {
	err = vk.RequestUnmarshal("video.add", params, &response)
	return
}

// VideoAddAlbumResponse struct
type VideoAddAlbumResponse struct {
	AlbumID int `json:"album_id"`
}

// VideoAddAlbum creates an empty album for videos.
//
// https://vk.com/dev/video.addAlbum
func (vk *VK) VideoAddAlbum(params Params) (response VideoAddAlbumResponse, err error) {
	err = vk.RequestUnmarshal("video.addAlbum", params, &response)
	return
}

// VideoAddToAlbum allows you to add a video to the album.
//
// https://vk.com/dev/video.addToAlbum
func (vk *VK) VideoAddToAlbum(params Params) (response int, err error) {
	err = vk.RequestUnmarshal("video.addToAlbum", params, &response)
	return
}

// VideoCreateComment adds a new comment on a video.
//
// https://vk.com/dev/video.createComment
func (vk *VK) VideoCreateComment(params Params) (response int, err error) {
	err = vk.RequestUnmarshal("video.createComment", params, &response)
	return
}

// VideoDelete deletes a video from a user or community page.
//
// https://vk.com/dev/video.delete
func (vk *VK) VideoDelete(params Params) (response int, err error) {
	err = vk.RequestUnmarshal("video.delete", params, &response)
	return
}

// VideoDeleteAlbum deletes a video album.
//
// https://vk.com/dev/video.deleteAlbum
func (vk *VK) VideoDeleteAlbum(params Params) (response int, err error) {
	err = vk.RequestUnmarshal("video.deleteAlbum", params, &response)
	return
}

// VideoDeleteComment deletes a comment on a video.
//
// https://vk.com/dev/video.deleteComment
func (vk *VK) VideoDeleteComment(params Params) (response int, err error) {
	err = vk.RequestUnmarshal("video.deleteComment", params, &response)
	return
}

// VideoEdit edits information about a video on a user or community page.
//
// https://vk.com/dev/video.edit
func (vk *VK) VideoEdit(params Params) (response int, err error) {
	err = vk.RequestUnmarshal("video.edit", params, &response)
	return
}

// VideoEditAlbum edits the title of a video album.
//
// https://vk.com/dev/video.editAlbum
func (vk *VK) VideoEditAlbum(params Params) (response int, err error) {
	err = vk.RequestUnmarshal("video.editAlbum", params, &response)
	return
}

// VideoEditComment edits the text of a comment on a video.
//
// https://vk.com/dev/video.editComment
func (vk *VK) VideoEditComment(params Params) (response int, err error) {
	err = vk.RequestUnmarshal("video.editComment", params, &response)
	return
}

// VideoGetResponse struct
type VideoGetResponse struct {
	Count int                 `json:"count"`
	Items []object.VideoVideo `json:"items"`
}

// VideoGet returns detailed information about videos.
//
// extended=0
//
// https://vk.com/dev/video.get
func (vk *VK) VideoGet(params Params) (response VideoGetResponse, err error) {
	params["extended"] = "0"
	err = vk.RequestUnmarshal("video.get", params, &response)

	return
}

// VideoGetExtendedResponse struct
type VideoGetExtendedResponse struct {
	Count int                 `json:"count"`
	Items []object.VideoVideo `json:"items"`
	object.ExtendedResponse
}

// VideoGetExtended returns detailed information about videos.
//
// extended=1
//
// https://vk.com/dev/video.get
func (vk *VK) VideoGetExtended(params Params) (response VideoGetExtendedResponse, err error) {
	params["extended"] = true
	err = vk.RequestUnmarshal("video.get", params, &response)

	return
}

// VideoGetAlbumByIDResponse struct
type VideoGetAlbumByIDResponse object.VideoVideoAlbumFull

// VideoGetAlbumByID returns video album info
//
// https://vk.com/dev/video.getAlbumById
func (vk *VK) VideoGetAlbumByID(params Params) (response VideoGetAlbumByIDResponse, err error) {
	err = vk.RequestUnmarshal("video.getAlbumById", params, &response)
	return
}

// VideoGetAlbumsResponse struct
type VideoGetAlbumsResponse struct {
	Count int                      `json:"count"`
	Items []object.VideoVideoAlbum `json:"items"`
}

// VideoGetAlbums returns a list of video albums owned by a user or community.
//
// extended=0
//
// https://vk.com/dev/video.getAlbums
func (vk *VK) VideoGetAlbums(params Params) (response VideoGetAlbumsResponse, err error) {
	params["extended"] = "0"
	err = vk.RequestUnmarshal("video.getAlbums", params, &response)

	return
}

// VideoGetAlbumsExtendedResponse struct
type VideoGetAlbumsExtendedResponse struct {
	Count int                          `json:"count"`
	Items []object.VideoVideoAlbumFull `json:"items"`
}

// VideoGetAlbumsExtended returns a list of video albums owned by a user or community.
//
// extended=1
//
// https://vk.com/dev/video.getAlbums
func (vk *VK) VideoGetAlbumsExtended(params Params) (response VideoGetAlbumsExtendedResponse, err error) {
	params["extended"] = true
	err = vk.RequestUnmarshal("video.getAlbums", params, &response)

	return
}

// VideoGetAlbumsByVideoResponse struct
type VideoGetAlbumsByVideoResponse []int

// VideoGetAlbumsByVideo returns a list of albums in which the video is located.
//
// extended=0
//
// https://vk.com/dev/video.getAlbumsByVideo
func (vk *VK) VideoGetAlbumsByVideo(params Params) (response VideoGetAlbumsByVideoResponse, err error) {
	params["extended"] = "0"
	err = vk.RequestUnmarshal("video.getAlbumsByVideo", params, &response)

	return
}

// VideoGetAlbumsByVideoExtendedResponse struct
type VideoGetAlbumsByVideoExtendedResponse struct {
	Count int                          `json:"count"`
	Items []object.VideoVideoAlbumFull `json:"items"`
}

// VideoGetAlbumsByVideoExtended returns a list of albums in which the video is located.
//
// extended=1
//
// https://vk.com/dev/video.getAlbumsByVideo
func (vk *VK) VideoGetAlbumsByVideoExtended(params Params) (response VideoGetAlbumsByVideoExtendedResponse, err error) {
	params["extended"] = true
	err = vk.RequestUnmarshal("video.getAlbumsByVideo", params, &response)

	return
}

// VideoGetCommentsResponse struct
type VideoGetCommentsResponse struct {
	Count int                      `json:"count"`
	Items []object.WallWallComment `json:"items"`
}

// VideoGetComments returns a list of comments on a video.
//
// extended=0
//
// https://vk.com/dev/video.getComments
func (vk *VK) VideoGetComments(params Params) (response VideoGetCommentsResponse, err error) {
	params["extended"] = "0"
	err = vk.RequestUnmarshal("video.getComments", params, &response)

	return
}

// VideoGetCommentsExtendedResponse struct
type VideoGetCommentsExtendedResponse struct {
	Count int                      `json:"count"`
	Items []object.WallWallComment `json:"items"`
	object.ExtendedResponse
}

// VideoGetCommentsExtended returns a list of comments on a video.
//
// extended=1
//
// https://vk.com/dev/video.getComments
func (vk *VK) VideoGetCommentsExtended(params Params) (response VideoGetCommentsExtendedResponse, err error) {
	params["extended"] = true
	err = vk.RequestUnmarshal("video.getComments", params, &response)

	return
}

// VideoRemoveFromAlbum Allows you to remove the video from the album.
//
// https://vk.com/dev/video.removeFromAlbum
func (vk *VK) VideoRemoveFromAlbum(params Params) (response int, err error) {
	err = vk.RequestUnmarshal("video.removeFromAlbum", params, &response)
	return
}

// VideoReorderAlbums reorders the album in the list of user video albums.
//
// https://vk.com/dev/video.reorderAlbums
func (vk *VK) VideoReorderAlbums(params Params) (response int, err error) {
	err = vk.RequestUnmarshal("video.reorderAlbums", params, &response)
	return
}

// VideoReorderVideos reorders the video in the video album.
//
// https://vk.com/dev/video.reorderVideos
func (vk *VK) VideoReorderVideos(params Params) (response int, err error) {
	err = vk.RequestUnmarshal("video.reorderVideos", params, &response)
	return
}

// VideoReport reports (submits a complaint about) a video.
//
// https://vk.com/dev/video.report
func (vk *VK) VideoReport(params Params) (response int, err error) {
	err = vk.RequestUnmarshal("video.report", params, &response)
	return
}

// VideoReportComment reports (submits a complaint about) a comment on a video.
//
// https://vk.com/dev/video.reportComment
func (vk *VK) VideoReportComment(params Params) (response int, err error) {
	err = vk.RequestUnmarshal("video.reportComment", params, &response)
	return
}

// VideoRestore restores a previously deleted video.
//
// https://vk.com/dev/video.restore
func (vk *VK) VideoRestore(params Params) (response int, err error) {
	err = vk.RequestUnmarshal("video.restore", params, &response)
	return
}

// VideoRestoreComment restores a previously deleted comment on a video.
//
// https://vk.com/dev/video.restoreComment
func (vk *VK) VideoRestoreComment(params Params) (response int, err error) {
	err = vk.RequestUnmarshal("video.restoreComment", params, &response)
	return
}

// VideoSaveResponse struct
type VideoSaveResponse object.VideoSaveResult

// VideoSave returns a server address (required for upload) and video data.
//
// https://vk.com/dev/video.save
func (vk *VK) VideoSave(params Params) (response VideoSaveResponse, err error) {
	err = vk.RequestUnmarshal("video.save", params, &response)
	return
}

// VideoSearchResponse struct
type VideoSearchResponse struct {
	Count int                 `json:"count"`
	Items []object.VideoVideo `json:"items"`
}

// VideoSearch returns a list of videos under the set search criterion.
//
// extended=0
//
// https://vk.com/dev/video.search
func (vk *VK) VideoSearch(params Params) (response VideoSearchResponse, err error) {
	params["extended"] = "0"
	err = vk.RequestUnmarshal("video.search", params, &response)

	return
}

// VideoSearchExtendedResponse struct
type VideoSearchExtendedResponse struct {
	Count int                 `json:"count"`
	Items []object.VideoVideo `json:"items"`
	object.ExtendedResponse
}

// VideoSearchExtended returns a list of videos under the set search criterion.
//
// extended=1
//
// https://vk.com/dev/video.search
func (vk *VK) VideoSearchExtended(params Params) (response VideoSearchExtendedResponse, err error) {
	params["extended"] = true
	err = vk.RequestUnmarshal("video.search", params, &response)

	return
}
