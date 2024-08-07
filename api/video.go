package api // import "github.com/SevereCloud/vksdk/v3/api"

import (
	"github.com/SevereCloud/vksdk/v3/object"
)

// VideoAdd adds a video to a user or community page.
//
// https://dev.vk.com/method/video.add
func (vk *VK) VideoAdd(params Params) (response int, err error) {
	err = vk.RequestUnmarshal("video.add", &response, params)
	return
}

// VideoAddAlbumResponse struct.
type VideoAddAlbumResponse struct {
	AlbumID int `json:"album_id"`
}

// VideoAddAlbum creates an empty album for videos.
//
// https://dev.vk.com/method/video.addAlbum
func (vk *VK) VideoAddAlbum(params Params) (response VideoAddAlbumResponse, err error) {
	err = vk.RequestUnmarshal("video.addAlbum", &response, params)
	return
}

// VideoAddToAlbum allows you to add a video to the album.
//
// https://dev.vk.com/method/video.addToAlbum
func (vk *VK) VideoAddToAlbum(params Params) (response int, err error) {
	err = vk.RequestUnmarshal("video.addToAlbum", &response, params)
	return
}

// VideoCreateComment adds a new comment on a video.
//
// https://dev.vk.com/method/video.createComment
func (vk *VK) VideoCreateComment(params Params) (response int, err error) {
	err = vk.RequestUnmarshal("video.createComment", &response, params)
	return
}

// VideoDelete deletes a video from a user or community page.
//
// https://dev.vk.com/method/video.delete
func (vk *VK) VideoDelete(params Params) (response int, err error) {
	err = vk.RequestUnmarshal("video.delete", &response, params)
	return
}

// VideoDeleteAlbum deletes a video album.
//
// https://dev.vk.com/method/video.deleteAlbum
func (vk *VK) VideoDeleteAlbum(params Params) (response int, err error) {
	err = vk.RequestUnmarshal("video.deleteAlbum", &response, params)
	return
}

// VideoDeleteComment deletes a comment on a video.
//
// https://dev.vk.com/method/video.deleteComment
func (vk *VK) VideoDeleteComment(params Params) (response int, err error) {
	err = vk.RequestUnmarshal("video.deleteComment", &response, params)
	return
}

// VideoEditResponse struct.
type VideoEditResponse struct {
	Success   object.BaseBoolInt `json:"success"`
	AccessKey string             `json:"access_key"`
}

// VideoEdit edits information about a video on a user or community page.
//
// https://dev.vk.com/method/video.edit
func (vk *VK) VideoEdit(params Params) (response VideoEditResponse, err error) {
	err = vk.RequestUnmarshal("video.edit", &response, params)
	return
}

// VideoEditAlbum edits the title of a video album.
//
// https://dev.vk.com/method/video.editAlbum
func (vk *VK) VideoEditAlbum(params Params) (response int, err error) {
	err = vk.RequestUnmarshal("video.editAlbum", &response, params)
	return
}

// VideoEditComment edits the text of a comment on a video.
//
// https://dev.vk.com/method/video.editComment
func (vk *VK) VideoEditComment(params Params) (response int, err error) {
	err = vk.RequestUnmarshal("video.editComment", &response, params)
	return
}

// VideoGetResponse struct.
type VideoGetResponse struct {
	Count int                 `json:"count"`
	Items []object.VideoVideo `json:"items"`
}

// VideoGet returns detailed information about videos.
//
//	extended=0
//
// https://dev.vk.com/method/video.get
func (vk *VK) VideoGet(params Params) (response VideoGetResponse, err error) {
	err = vk.RequestUnmarshal("video.get", &response, params, Params{"extended": false})

	return
}

// VideoGetExtendedResponse struct.
type VideoGetExtendedResponse struct {
	Count int                 `json:"count"`
	Items []object.VideoVideo `json:"items"`
	object.ExtendedResponse
}

// VideoGetExtended returns detailed information about videos.
//
//	extended=1
//
// https://dev.vk.com/method/video.get
func (vk *VK) VideoGetExtended(params Params) (response VideoGetExtendedResponse, err error) {
	err = vk.RequestUnmarshal("video.get", &response, params, Params{"extended": true})

	return
}

// VideoGetAlbumByIDResponse struct.
type VideoGetAlbumByIDResponse object.VideoVideoAlbumFull

// VideoGetAlbumByID returns video album info.
//
// https://dev.vk.com/method/video.getAlbumById
func (vk *VK) VideoGetAlbumByID(params Params) (response VideoGetAlbumByIDResponse, err error) {
	err = vk.RequestUnmarshal("video.getAlbumById", &response, params)
	return
}

// VideoGetAlbumsResponse struct.
type VideoGetAlbumsResponse struct {
	Count int                      `json:"count"`
	Items []object.VideoVideoAlbum `json:"items"`
}

// VideoGetAlbums returns a list of video albums owned by a user or community.
//
//	extended=0
//
// https://dev.vk.com/method/video.getAlbums
func (vk *VK) VideoGetAlbums(params Params) (response VideoGetAlbumsResponse, err error) {
	err = vk.RequestUnmarshal("video.getAlbums", &response, params, Params{"extended": false})

	return
}

// VideoGetAlbumsExtendedResponse struct.
type VideoGetAlbumsExtendedResponse struct {
	Count int                          `json:"count"`
	Items []object.VideoVideoAlbumFull `json:"items"`
}

// VideoGetAlbumsExtended returns a list of video albums owned by a user or community.
//
//	extended=1
//
// https://dev.vk.com/method/video.getAlbums
func (vk *VK) VideoGetAlbumsExtended(params Params) (response VideoGetAlbumsExtendedResponse, err error) {
	err = vk.RequestUnmarshal("video.getAlbums", &response, params, Params{"extended": true})

	return
}

// VideoGetAlbumsByVideoResponse struct.
type VideoGetAlbumsByVideoResponse []int

// VideoGetAlbumsByVideo returns a list of albums in which the video is located.
//
//	extended=0
//
// https://dev.vk.com/method/video.getAlbumsByVideo
func (vk *VK) VideoGetAlbumsByVideo(params Params) (response VideoGetAlbumsByVideoResponse, err error) {
	err = vk.RequestUnmarshal("video.getAlbumsByVideo", &response, params, Params{"extended": false})

	return
}

// VideoGetAlbumsByVideoExtendedResponse struct.
type VideoGetAlbumsByVideoExtendedResponse struct {
	Count int                          `json:"count"`
	Items []object.VideoVideoAlbumFull `json:"items"`
}

// VideoGetAlbumsByVideoExtended returns a list of albums in which the video is located.
//
//	extended=1
//
// https://dev.vk.com/method/video.getAlbumsByVideo
func (vk *VK) VideoGetAlbumsByVideoExtended(params Params) (response VideoGetAlbumsByVideoExtendedResponse, err error) {
	err = vk.RequestUnmarshal("video.getAlbumsByVideo", &response, params, Params{"extended": true})

	return
}

// VideoGetCommentsResponse struct.
type VideoGetCommentsResponse struct {
	Count int                      `json:"count"`
	Items []object.WallWallComment `json:"items"`
}

// VideoGetComments returns a list of comments on a video.
//
//	extended=0
//
// https://dev.vk.com/method/video.getComments
func (vk *VK) VideoGetComments(params Params) (response VideoGetCommentsResponse, err error) {
	err = vk.RequestUnmarshal("video.getComments", &response, params, Params{"extended": false})

	return
}

// VideoGetCommentsExtendedResponse struct.
type VideoGetCommentsExtendedResponse struct {
	Count int                      `json:"count"`
	Items []object.WallWallComment `json:"items"`
	object.ExtendedResponse
}

// VideoGetCommentsExtended returns a list of comments on a video.
//
//	extended=1
//
// https://dev.vk.com/method/video.getComments
func (vk *VK) VideoGetCommentsExtended(params Params) (response VideoGetCommentsExtendedResponse, err error) {
	err = vk.RequestUnmarshal("video.getComments", &response, params, Params{"extended": true})

	return
}

// VideoLiveGetCategoriesResponse struct.
type VideoLiveGetCategoriesResponse []object.VideoLiveCategory

// VideoLiveGetCategories method.
//
// https://dev.vk.com/method/video.liveGetCategories
func (vk *VK) VideoLiveGetCategories(params Params) (response VideoLiveGetCategoriesResponse, err error) {
	err = vk.RequestUnmarshal("video.liveGetCategories", &response, params)
	return
}

// VideoRemoveFromAlbum allows you to remove the video from the album.
//
// https://dev.vk.com/method/video.removeFromAlbum
func (vk *VK) VideoRemoveFromAlbum(params Params) (response int, err error) {
	err = vk.RequestUnmarshal("video.removeFromAlbum", &response, params)
	return
}

// VideoReorderAlbums reorders the album in the list of user video albums.
//
// https://dev.vk.com/method/video.reorderAlbums
func (vk *VK) VideoReorderAlbums(params Params) (response int, err error) {
	err = vk.RequestUnmarshal("video.reorderAlbums", &response, params)
	return
}

// VideoReorderVideos reorders the video in the video album.
//
// https://dev.vk.com/method/video.reorderVideos
func (vk *VK) VideoReorderVideos(params Params) (response int, err error) {
	err = vk.RequestUnmarshal("video.reorderVideos", &response, params)
	return
}

// VideoReport reports (submits a complaint about) a video.
//
// https://dev.vk.com/method/video.report
func (vk *VK) VideoReport(params Params) (response int, err error) {
	err = vk.RequestUnmarshal("video.report", &response, params)
	return
}

// VideoReportComment reports (submits a complaint about) a comment on a video.
//
// https://dev.vk.com/method/video.reportComment
func (vk *VK) VideoReportComment(params Params) (response int, err error) {
	err = vk.RequestUnmarshal("video.reportComment", &response, params)
	return
}

// VideoRestore restores a previously deleted video.
//
// https://dev.vk.com/method/video.restore
func (vk *VK) VideoRestore(params Params) (response int, err error) {
	err = vk.RequestUnmarshal("video.restore", &response, params)
	return
}

// VideoRestoreComment restores a previously deleted comment on a video.
//
// https://dev.vk.com/method/video.restoreComment
func (vk *VK) VideoRestoreComment(params Params) (response int, err error) {
	err = vk.RequestUnmarshal("video.restoreComment", &response, params)
	return
}

// VideoSaveResponse struct.
type VideoSaveResponse object.VideoSaveResult

// VideoSave returns a server address (required for upload) and video data.
//
// https://dev.vk.com/method/video.save
func (vk *VK) VideoSave(params Params) (response VideoSaveResponse, err error) {
	err = vk.RequestUnmarshal("video.save", &response, params)
	return
}

// VideoSearchResponse struct.
type VideoSearchResponse struct {
	Count int                 `json:"count"`
	Items []object.VideoVideo `json:"items"`
}

// VideoSearch returns a list of videos under the set search criterion.
//
//	extended=0
//
// https://dev.vk.com/method/video.search
func (vk *VK) VideoSearch(params Params) (response VideoSearchResponse, err error) {
	err = vk.RequestUnmarshal("video.search", &response, params, Params{"extended": false})

	return
}

// VideoSearchExtendedResponse struct.
type VideoSearchExtendedResponse struct {
	Count int                 `json:"count"`
	Items []object.VideoVideo `json:"items"`
	object.ExtendedResponse
}

// VideoSearchExtended returns a list of videos under the set search criterion.
//
//	extended=1
//
// https://dev.vk.com/method/video.search
func (vk *VK) VideoSearchExtended(params Params) (response VideoSearchExtendedResponse, err error) {
	err = vk.RequestUnmarshal("video.search", &response, params, Params{"extended": true})

	return
}

// VideoStartStreamingResponse struct.
type VideoStartStreamingResponse object.VideoLive

// VideoStartStreaming method.
//
// https://dev.vk.com/method/video.startStreaming
func (vk *VK) VideoStartStreaming(params Params) (response VideoStartStreamingResponse, err error) {
	err = vk.RequestUnmarshal("video.startStreaming", &response, params)
	return
}

// VideoStopStreamingResponse struct.
type VideoStopStreamingResponse struct {
	UniqueViewers int `json:"unique_viewers"`
}

// VideoStopStreaming method.
//
// https://dev.vk.com/method/video.stopStreaming
func (vk *VK) VideoStopStreaming(params Params) (response VideoStopStreamingResponse, err error) {
	err = vk.RequestUnmarshal("video.stopStreaming", &response, params)
	return
}
