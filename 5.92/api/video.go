package api // import "github.com/SevereCloud/vksdk/5.92/api"

import "github.com/SevereCloud/vksdk/5.92/object"

// VideoAddResponse struct
type VideoAddResponse int

// VideoAdd adds a video to a user or community page.
//
// https://vk.com/dev/video.add
func (vk *VK) VideoAdd(params map[string]string) (response VideoAddResponse, vkErr Error) {
	vk.RequestUnmarshal("video.add", params, &response, &vkErr)
	return
}

// VideoAddAlbumResponse struct
type VideoAddAlbumResponse struct {
	AlbumID int `json:"album_id"`
}

// VideoAddAlbum creates an empty album for videos.
//
// https://vk.com/dev/video.addAlbum
func (vk *VK) VideoAddAlbum(params map[string]string) (response VideoAddAlbumResponse, vkErr Error) {
	vk.RequestUnmarshal("video.addAlbum", params, &response, &vkErr)
	return
}

// VideoAddToAlbumResponse struct
type VideoAddToAlbumResponse int

// VideoAddToAlbum allows you to add a video to the album.
//
// https://vk.com/dev/video.addToAlbum
func (vk *VK) VideoAddToAlbum(params map[string]string) (response VideoAddToAlbumResponse, vkErr Error) {
	vk.RequestUnmarshal("video.addToAlbum", params, &response, &vkErr)
	return
}

// VideoCreateCommentResponse struct
type VideoCreateCommentResponse int

// VideoCreateComment adds a new comment on a video.
//
// https://vk.com/dev/video.createComment
func (vk *VK) VideoCreateComment(params map[string]string) (response VideoCreateCommentResponse, vkErr Error) {
	vk.RequestUnmarshal("video.createComment", params, &response, &vkErr)
	return
}

// VideoDeleteResponse struct
type VideoDeleteResponse int

// VideoDelete deletes a video from a user or community page.
//
// https://vk.com/dev/video.delete
func (vk *VK) VideoDelete(params map[string]string) (response VideoDeleteResponse, vkErr Error) {
	vk.RequestUnmarshal("video.delete", params, &response, &vkErr)
	return
}

// VideoDeleteAlbumResponse struct
type VideoDeleteAlbumResponse int

// VideoDeleteAlbum deletes a video album.
//
// https://vk.com/dev/video.deleteAlbum
func (vk *VK) VideoDeleteAlbum(params map[string]string) (response VideoDeleteAlbumResponse, vkErr Error) {
	vk.RequestUnmarshal("video.deleteAlbum", params, &response, &vkErr)
	return
}

// VideoDeleteCommentResponse struct
type VideoDeleteCommentResponse int

// VideoDeleteComment deletes a comment on a video.
//
// https://vk.com/dev/video.deleteComment
func (vk *VK) VideoDeleteComment(params map[string]string) (response VideoDeleteCommentResponse, vkErr Error) {
	vk.RequestUnmarshal("video.deleteComment", params, &response, &vkErr)
	return
}

// VideoEditResponse struct
type VideoEditResponse int

// VideoEdit edits information about a video on a user or community page.
//
// https://vk.com/dev/video.edit
func (vk *VK) VideoEdit(params map[string]string) (response VideoEditResponse, vkErr Error) {
	vk.RequestUnmarshal("video.edit", params, &response, &vkErr)
	return
}

// VideoEditAlbumResponse struct
type VideoEditAlbumResponse int

// VideoEditAlbum edits the title of a video album.
//
// https://vk.com/dev/video.editAlbum
func (vk *VK) VideoEditAlbum(params map[string]string) (response VideoEditAlbumResponse, vkErr Error) {
	vk.RequestUnmarshal("video.editAlbum", params, &response, &vkErr)
	return
}

// VideoEditCommentResponse struct
type VideoEditCommentResponse int

// VideoEditComment edits the text of a comment on a video.
//
// https://vk.com/dev/video.editComment
func (vk *VK) VideoEditComment(params map[string]string) (response VideoEditCommentResponse, vkErr Error) {
	vk.RequestUnmarshal("video.editComment", params, &response, &vkErr)
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
func (vk *VK) VideoGet(params map[string]string) (response VideoGetResponse, vkErr Error) {
	params["extended"] = "0"
	vk.RequestUnmarshal("video.get", params, &response, &vkErr)
	return
}

// VideoGetExtendedResponse struct
type VideoGetExtendedResponse struct {
	Count    int                  `json:"count"`
	Items    []object.VideoVideo  `json:"items"`
	Profiles []object.UsersUser   `json:"profiles"`
	Groups   []object.GroupsGroup `json:"groups"`
}

// VideoGetExtended returns detailed information about videos.
//
// extended=1
//
// https://vk.com/dev/video.get
func (vk *VK) VideoGetExtended(params map[string]string) (response VideoGetExtendedResponse, vkErr Error) {
	params["extended"] = "1"
	vk.RequestUnmarshal("video.get", params, &response, &vkErr)
	return
}

// VideoGetAlbumByIDResponse struct
type VideoGetAlbumByIDResponse object.VideoVideoAlbumFull

// VideoGetAlbumByID returns video album info
//
// https://vk.com/dev/video.getAlbumById
func (vk *VK) VideoGetAlbumByID(params map[string]string) (response VideoGetAlbumByIDResponse, vkErr Error) {
	vk.RequestUnmarshal("video.getAlbumById", params, &response, &vkErr)
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
func (vk *VK) VideoGetAlbums(params map[string]string) (response VideoGetAlbumsResponse, vkErr Error) {
	params["extended"] = "0"
	vk.RequestUnmarshal("video.getAlbums", params, &response, &vkErr)
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
func (vk *VK) VideoGetAlbumsExtended(params map[string]string) (response VideoGetAlbumsExtendedResponse, vkErr Error) {
	params["extended"] = "1"
	vk.RequestUnmarshal("video.getAlbums", params, &response, &vkErr)
	return
}

// VideoGetAlbumsByVideoResponse struct
type VideoGetAlbumsByVideoResponse []int

// VideoGetAlbumsByVideo returns a list of albums in which the video is located.
//
// extended=0
//
// https://vk.com/dev/video.getAlbumsByVideo
func (vk *VK) VideoGetAlbumsByVideo(params map[string]string) (response VideoGetAlbumsByVideoResponse, vkErr Error) {
	params["extended"] = "0"
	vk.RequestUnmarshal("video.getAlbumsByVideo", params, &response, &vkErr)
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
func (vk *VK) VideoGetAlbumsByVideoExtended(params map[string]string) (response VideoGetAlbumsByVideoExtendedResponse, vkErr Error) {
	params["extended"] = "1"
	vk.RequestUnmarshal("video.getAlbumsByVideo", params, &response, &vkErr)
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
func (vk *VK) VideoGetComments(params map[string]string) (response VideoGetCommentsResponse, vkErr Error) {
	params["extended"] = "0"
	vk.RequestUnmarshal("video.getComments", params, &response, &vkErr)
	return
}

// VideoGetCommentsExtendedResponse struct
type VideoGetCommentsExtendedResponse struct {
	Count    int                      `json:"count"`
	Items    []object.WallWallComment `json:"items"`
	Profiles []object.UsersUser       `json:"profiles"`
	Groups   []object.GroupsGroup     `json:"groups"`
}

// VideoGetCommentsExtended returns a list of comments on a video.
//
// extended=1
//
// https://vk.com/dev/video.getComments
func (vk *VK) VideoGetCommentsExtended(params map[string]string) (response VideoGetCommentsExtendedResponse, vkErr Error) {
	params["extended"] = "1"
	vk.RequestUnmarshal("video.getComments", params, &response, &vkErr)
	return
}

// VideoRemoveFromAlbumResponse struct
type VideoRemoveFromAlbumResponse int

// VideoRemoveFromAlbum Allows you to remove the video from the album.
//
// https://vk.com/dev/video.removeFromAlbum
func (vk *VK) VideoRemoveFromAlbum(params map[string]string) (response VideoRemoveFromAlbumResponse, vkErr Error) {
	vk.RequestUnmarshal("video.removeFromAlbum", params, &response, &vkErr)
	return
}

// VideoReorderAlbumsResponse struct
type VideoReorderAlbumsResponse int

// VideoReorderAlbums reorders the album in the list of user video albums.
//
// https://vk.com/dev/video.reorderAlbums
func (vk *VK) VideoReorderAlbums(params map[string]string) (response VideoReorderAlbumsResponse, vkErr Error) {
	vk.RequestUnmarshal("video.reorderAlbums", params, &response, &vkErr)
	return
}

// VideoReorderVideosResponse struct
type VideoReorderVideosResponse int

// VideoReorderVideos reorders the video in the video album.
//
// https://vk.com/dev/video.reorderVideos
func (vk *VK) VideoReorderVideos(params map[string]string) (response VideoReorderVideosResponse, vkErr Error) {
	vk.RequestUnmarshal("video.reorderVideos", params, &response, &vkErr)
	return
}

// VideoReportResponse struct
type VideoReportResponse int

// VideoReport reports (submits a complaint about) a video.
//
// https://vk.com/dev/video.report
func (vk *VK) VideoReport(params map[string]string) (response VideoReportResponse, vkErr Error) {
	vk.RequestUnmarshal("video.report", params, &response, &vkErr)
	return
}

// VideoReportCommentResponse struct
type VideoReportCommentResponse int

// VideoReportComment reports (submits a complaint about) a comment on a video.
//
// https://vk.com/dev/video.reportComment
func (vk *VK) VideoReportComment(params map[string]string) (response VideoReportCommentResponse, vkErr Error) {
	vk.RequestUnmarshal("video.reportComment", params, &response, &vkErr)
	return
}

// VideoRestoreResponse struct
type VideoRestoreResponse int

// VideoRestore restores a previously deleted video.
//
// https://vk.com/dev/video.restore
func (vk *VK) VideoRestore(params map[string]string) (response VideoRestoreResponse, vkErr Error) {
	vk.RequestUnmarshal("video.restore", params, &response, &vkErr)
	return
}

// VideoRestoreCommentResponse struct
type VideoRestoreCommentResponse int

// VideoRestoreComment restores a previously deleted comment on a video.
//
// https://vk.com/dev/video.restoreComment
func (vk *VK) VideoRestoreComment(params map[string]string) (response VideoRestoreCommentResponse, vkErr Error) {
	vk.RequestUnmarshal("video.restoreComment", params, &response, &vkErr)
	return
}

// VideoSaveResponse struct
type VideoSaveResponse object.VideoSaveResult

// VideoSave returns a server address (required for upload) and video data.
//
// https://vk.com/dev/video.save
func (vk *VK) VideoSave(params map[string]string) (response VideoSaveResponse, vkErr Error) {
	vk.RequestUnmarshal("video.save", params, &response, &vkErr)
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
func (vk *VK) VideoSearch(params map[string]string) (response VideoSearchResponse, vkErr Error) {
	params["extended"] = "0"
	vk.RequestUnmarshal("video.search", params, &response, &vkErr)
	return
}

// VideoSearchExtendedResponse struct
type VideoSearchExtendedResponse struct {
	Count    int                  `json:"count"`
	Items    []object.VideoVideo  `json:"items"`
	Profiles []object.UsersUser   `json:"profiles"`
	Groups   []object.GroupsGroup `json:"groups"`
}

// VideoSearchExtended returns a list of videos under the set search criterion.
//
// extended=1
//
// https://vk.com/dev/video.search
func (vk *VK) VideoSearchExtended(params map[string]string) (response VideoSearchExtendedResponse, vkErr Error) {
	params["extended"] = "1"
	vk.RequestUnmarshal("video.search", params, &response, &vkErr)
	return
}
