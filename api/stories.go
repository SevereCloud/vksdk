package api // import "github.com/SevereCloud/vksdk/api"

import "github.com/SevereCloud/vksdk/object"

// StoriesBanOwner allows to hide stories from chosen sources from current user's feed.
//
// https://vk.com/dev/stories.banOwner
func (vk *VK) StoriesBanOwner(params map[string]string) (response int, err error) {
	err = vk.RequestUnmarshal("stories.banOwner", params, &response)
	return
}

// StoriesDelete allows to delete story.
//
// https://vk.com/dev/stories.delete
func (vk *VK) StoriesDelete(params map[string]string) (response int, err error) {
	err = vk.RequestUnmarshal("stories.delete", params, &response)
	return
}

// StoriesGetResponse struct
type StoriesGetResponse struct {
	Count int                   `json:"count"`
	Items []object.StoriesStory `json:"items"`
}

// StoriesGet returns stories available for current user.
//
// extended=0
//
// https://vk.com/dev/stories.get
func (vk *VK) StoriesGet(params map[string]string) (response StoriesGetResponse, err error) {
	params["extended"] = "0"
	err = vk.RequestUnmarshal("stories.get", params, &response)

	return
}

// StoriesGetExtendedResponse struct
type StoriesGetExtendedResponse struct {
	Count int                   `json:"count"`
	Items []object.StoriesStory `json:"items"`
	object.ExtendedResponse
}

// StoriesGetExtended returns stories available for current user.
//
// extended=1
//
// https://vk.com/dev/stories.get
func (vk *VK) StoriesGetExtended(params map[string]string) (response StoriesGetExtendedResponse, err error) {
	params["extended"] = "1"
	err = vk.RequestUnmarshal("stories.get", params, &response)

	return
}

// StoriesGetBannedResponse struct
type StoriesGetBannedResponse struct {
	Count int   `json:"count"`
	Items []int `json:"items"`
}

// StoriesGetBanned returns list of sources hidden from current user's feed.
//
// extended=0
//
// https://vk.com/dev/stories.getBanned
func (vk *VK) StoriesGetBanned(params map[string]string) (response StoriesGetBannedResponse, err error) {
	params["extended"] = "0"
	err = vk.RequestUnmarshal("stories.getBanned", params, &response)

	return
}

// StoriesGetBannedExtendedResponse struct
type StoriesGetBannedExtendedResponse struct {
	Count int   `json:"count"`
	Items []int `json:"items"`
	object.ExtendedResponse
}

// StoriesGetBannedExtended returns list of sources hidden from current user's feed.
//
// extended=1
//
// https://vk.com/dev/stories.getBanned
func (vk *VK) StoriesGetBannedExtended(params map[string]string) (response StoriesGetBannedExtendedResponse, err error) {
	params["extended"] = "1"
	err = vk.RequestUnmarshal("stories.getBanned", params, &response)

	return
}

// StoriesGetByIDResponse struct
type StoriesGetByIDResponse struct {
	Count int                   `json:"count"`
	Items []object.StoriesStory `json:"items"`
}

// StoriesGetByID returns story by its ID.
//
// extended=0
//
// https://vk.com/dev/stories.getById
func (vk *VK) StoriesGetByID(params map[string]string) (response StoriesGetByIDResponse, err error) {
	params["extended"] = "0"
	err = vk.RequestUnmarshal("stories.getById", params, &response)

	return
}

// StoriesGetByIDExtendedResponse struct
type StoriesGetByIDExtendedResponse struct {
	Count int                   `json:"count"`
	Items []object.StoriesStory `json:"items"`
	object.ExtendedResponse
}

// StoriesGetByIDExtended returns story by its ID.
//
// extended=1
//
// https://vk.com/dev/stories.getById
func (vk *VK) StoriesGetByIDExtended(params map[string]string) (response StoriesGetByIDExtendedResponse, err error) {
	params["extended"] = "1"
	err = vk.RequestUnmarshal("stories.getById", params, &response)

	return
}

// StoriesGetPhotoUploadServerResponse struct
type StoriesGetPhotoUploadServerResponse struct {
	UploadURL string `json:"upload_url"`
	PeerIDs   []int  `json:"peer_ids"`
	UserIDs   []int  `json:"user_ids"`
}

// StoriesGetPhotoUploadServer returns URL for uploading a story with photo.
//
// https://vk.com/dev/stories.getPhotoUploadServer
func (vk *VK) StoriesGetPhotoUploadServer(params map[string]string) (response StoriesGetPhotoUploadServerResponse, err error) {
	err = vk.RequestUnmarshal("stories.getPhotoUploadServer", params, &response)
	return
}

// StoriesGetRepliesResponse struct
type StoriesGetRepliesResponse struct {
	Count int                     `json:"count"`
	Items [][]object.StoriesStory `json:"items"`
}

// StoriesGetReplies returns replies to the story.
//
// extended=0
//
// https://vk.com/dev/stories.getReplies
func (vk *VK) StoriesGetReplies(params map[string]string) (response StoriesGetRepliesResponse, err error) {
	params["extended"] = "0"
	err = vk.RequestUnmarshal("stories.getReplies", params, &response)

	return
}

// StoriesGetRepliesExtendedResponse struct
type StoriesGetRepliesExtendedResponse struct {
	Count int                     `json:"count"`
	Items [][]object.StoriesStory `json:"items"`
	object.ExtendedResponse
}

// StoriesGetRepliesExtended returns replies to the story.
//
// extended=1
//
// https://vk.com/dev/stories.getReplies
func (vk *VK) StoriesGetRepliesExtended(params map[string]string) (response StoriesGetRepliesExtendedResponse, err error) {
	params["extended"] = "1"
	err = vk.RequestUnmarshal("stories.getReplies", params, &response)

	return
}

// StoriesGetStatsResponse struct
type StoriesGetStatsResponse object.StoriesStoryStats

// StoriesGetStats return statistics data for the story.
//
// https://vk.com/dev/stories.getStats
func (vk *VK) StoriesGetStats(params map[string]string) (response StoriesGetStatsResponse, err error) {
	err = vk.RequestUnmarshal("stories.getStats", params, &response)
	return
}

// StoriesGetVideoUploadServerResponse struct
type StoriesGetVideoUploadServerResponse struct {
	UploadURL string `json:"upload_url"`
	PeerIDs   []int  `json:"peer_ids"`
	UserIDs   []int  `json:"user_ids"`
}

// StoriesGetVideoUploadServer allows to receive URL for uploading story with video.
//
// https://vk.com/dev/stories.getVideoUploadServer
func (vk *VK) StoriesGetVideoUploadServer(params map[string]string) (response StoriesGetVideoUploadServerResponse, err error) {
	err = vk.RequestUnmarshal("stories.getVideoUploadServer", params, &response)
	return
}

// StoriesGetViewersResponse struct
type StoriesGetViewersResponse struct {
	Count int   `json:"count"`
	Items []int `json:"items"`
}

// StoriesGetViewers returns a list of story viewers.
//
// extended=0
//
// https://vk.com/dev/stories.getViewers
func (vk *VK) StoriesGetViewers(params map[string]string) (response StoriesGetViewersResponse, err error) {
	params["extended"] = "0"
	err = vk.RequestUnmarshal("stories.getViewers", params, &response)

	return
}

// StoriesGetViewersExtendedResponse struct
type StoriesGetViewersExtendedResponse struct {
	Count int                `json:"count"`
	Items []object.UsersUser `json:"items"`
}

// StoriesGetViewersExtended returns a list of story viewers.
//
// extended=1
//
// https://vk.com/dev/stories.getViewers
func (vk *VK) StoriesGetViewersExtended(params map[string]string) (response StoriesGetViewersExtendedResponse, err error) {
	params["extended"] = "1"
	err = vk.RequestUnmarshal("stories.getViewers", params, &response)

	return
}

// StoriesHideAllReplies hides all replies in the last 24 hours from the user to current user's stories.
//
// https://vk.com/dev/stories.hideAllReplies
func (vk *VK) StoriesHideAllReplies(params map[string]string) (response int, err error) {
	err = vk.RequestUnmarshal("stories.hideAllReplies", params, &response)
	return
}

// StoriesHideReply hides the reply to the current user's story.
//
// https://vk.com/dev/stories.hideReply
func (vk *VK) StoriesHideReply(params map[string]string) (response int, err error) {
	err = vk.RequestUnmarshal("stories.hideReply", params, &response)
	return
}

// StoriesSearchResponse struct
type StoriesSearchResponse struct {
	Count int                     `json:"count"`
	Items [][]object.StoriesStory `json:"items"`
}

// StoriesSearch returns search results for stories.
//
// extended=0
//
// https://vk.com/dev/stories.search
func (vk *VK) StoriesSearch(params map[string]string) (response StoriesSearchResponse, err error) {
	params["extended"] = "0"
	err = vk.RequestUnmarshal("stories.search", params, &response)

	return
}

// StoriesSearchExtendedResponse struct
type StoriesSearchExtendedResponse struct {
	Count int                     `json:"count"`
	Items [][]object.StoriesStory `json:"items"`
	object.ExtendedResponse
}

// StoriesSearchExtended returns search results for stories.
//
// extended=1
//
// https://vk.com/dev/stories.search
func (vk *VK) StoriesSearchExtended(params map[string]string) (response StoriesSearchExtendedResponse, err error) {
	params["extended"] = "1"
	err = vk.RequestUnmarshal("stories.search", params, &response)

	return
}

// StoriesUnbanOwner allows to show stories from hidden sources in current user's feed.
//
// https://vk.com/dev/stories.unbanOwner
func (vk *VK) StoriesUnbanOwner(params map[string]string) (response int, err error) {
	err = vk.RequestUnmarshal("stories.unbanOwner", params, &response)
	return
}
