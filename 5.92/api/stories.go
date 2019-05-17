package api // import "github.com/SevereCloud/vksdk/5.92/api"

import "github.com/SevereCloud/vksdk/5.92/object"

// StoriesBanOwner allows to hide stories from chosen sources from current user's feed.
//
// https://vk.com/dev/stories.banOwner
func (vk *VK) StoriesBanOwner(params map[string]string) (vkErr Error) {
	_, vkErr = vk.Request("stories.banOwner", params)
	return
}

// StoriesDelete allows to delete story.
//
// https://vk.com/dev/stories.delete
func (vk *VK) StoriesDelete(params map[string]string) (vkErr Error) {
	_, vkErr = vk.Request("stories.delete", params)
	return
}

// StoriesGetResponse struct
type StoriesGetResponse struct {
	Count int                   `json:"count"`
	Items []object.StoriesStory `json:"items"`
}

// StoriesGet returns stories available for current user.
//
// https://vk.com/dev/stories.get
func (vk *VK) StoriesGet(params map[string]string) (response StoriesGetResponse, vkErr Error) {
	params["extended"] = "0"
	vk.RequestUnmarshal("stories.get", params, &response, &vkErr)
	return
}

// StoriesGetExtendedResponse struct
type StoriesGetExtendedResponse struct {
	Count    int                   `json:"count"`
	Items    []object.StoriesStory `json:"items"`
	Profiles []object.UsersUser    `json:"profiles"`
	Groups   []object.GroupsGroup  `json:"groups"`
}

// StoriesGetExtended returns stories available for current user.
//
// https://vk.com/dev/stories.get
func (vk *VK) StoriesGetExtended(params map[string]string) (response StoriesGetExtendedResponse, vkErr Error) {
	params["extended"] = "1"
	vk.RequestUnmarshal("stories.get", params, &response, &vkErr)
	return
}

// StoriesGetBannedResponse struct
type StoriesGetBannedResponse struct {
	Count int   `json:"count"`
	Items []int `json:"items"`
}

// StoriesGetBanned returns list of sources hidden from current user's feed.
//
// https://vk.com/dev/stories.getBanned
func (vk *VK) StoriesGetBanned(params map[string]string) (response StoriesGetBannedResponse, vkErr Error) {
	params["extended"] = "0"
	vk.RequestUnmarshal("stories.getBanned", params, &response, &vkErr)
	return
}

// StoriesGetBannedExtendedResponse struct
type StoriesGetBannedExtendedResponse struct {
	Count    int                  `json:"count"`
	Items    []int                `json:"items"`
	Profiles []object.UsersUser   `json:"profiles"`
	Groups   []object.GroupsGroup `json:"groups"`
}

// StoriesGetBannedExtended returns list of sources hidden from current user's feed.
//
// https://vk.com/dev/stories.getBanned
func (vk *VK) StoriesGetBannedExtended(params map[string]string) (response StoriesGetBannedExtendedResponse, vkErr Error) {
	params["extended"] = "1"
	vk.RequestUnmarshal("stories.getBanned", params, &response, &vkErr)
	return
}

// StoriesGetByIDResponse struct
type StoriesGetByIDResponse struct {
	Count int                   `json:"count"`
	Items []object.StoriesStory `json:"items"`
}

// StoriesGetByID returns story by its ID.
//
// https://vk.com/dev/stories.getById
func (vk *VK) StoriesGetByID(params map[string]string) (response StoriesGetByIDResponse, vkErr Error) {
	params["extended"] = "0"
	vk.RequestUnmarshal("stories.getById", params, &response, &vkErr)
	return
}

// StoriesGetByIDExtendedResponse struct
type StoriesGetByIDExtendedResponse struct {
	Count    int                   `json:"count"`
	Items    []object.StoriesStory `json:"items"`
	Profiles []object.UsersUser    `json:"profiles"`
	Groups   []object.GroupsGroup  `json:"groups"`
}

// StoriesGetByIDExtended returns story by its ID.
//
// https://vk.com/dev/stories.getById
func (vk *VK) StoriesGetByIDExtended(params map[string]string) (response StoriesGetByIDExtendedResponse, vkErr Error) {
	params["extended"] = "1"
	vk.RequestUnmarshal("stories.getById", params, &response, &vkErr)
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
func (vk *VK) StoriesGetPhotoUploadServer(params map[string]string) (response StoriesGetPhotoUploadServerResponse, vkErr Error) {
	vk.RequestUnmarshal("stories.getPhotoUploadServer", params, &response, &vkErr)
	return
}

// StoriesGetRepliesResponse struct
type StoriesGetRepliesResponse struct {
	Count int                     `json:"count"`
	Items [][]object.StoriesStory `json:"items"`
}

// StoriesGetReplies returns replies to the story.
//
// https://vk.com/dev/stories.getReplies
func (vk *VK) StoriesGetReplies(params map[string]string) (response StoriesGetRepliesResponse, vkErr Error) {
	params["extended"] = "0"
	vk.RequestUnmarshal("stories.getReplies", params, &response, &vkErr)
	return
}

// StoriesGetRepliesExtendedResponse struct
type StoriesGetRepliesExtendedResponse struct {
	Count    int                     `json:"count"`
	Items    [][]object.StoriesStory `json:"items"`
	Profiles []object.UsersUser      `json:"profiles"`
	Groups   []object.GroupsGroup    `json:"groups"`
}

// StoriesGetRepliesExtended returns replies to the story.
//
// https://vk.com/dev/stories.getReplies
func (vk *VK) StoriesGetRepliesExtended(params map[string]string) (response StoriesGetRepliesExtendedResponse, vkErr Error) {
	params["extended"] = "1"
	vk.RequestUnmarshal("stories.getReplies", params, &response, &vkErr)
	return
}

// StoriesGetStatsResponse struct
type StoriesGetStatsResponse object.StoriesStoryStats

// StoriesGetStats return statistics data for the story.
//
// https://vk.com/dev/stories.getStats
func (vk *VK) StoriesGetStats(params map[string]string) (response StoriesGetStatsResponse, vkErr Error) {
	vk.RequestUnmarshal("stories.getStats", params, &response, &vkErr)
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
func (vk *VK) StoriesGetVideoUploadServer(params map[string]string) (response StoriesGetVideoUploadServerResponse, vkErr Error) {
	vk.RequestUnmarshal("stories.getVideoUploadServer", params, &response, &vkErr)
	return
}

// StoriesGetViewersResponse struct
type StoriesGetViewersResponse struct {
	Count int   `json:"count"`
	Items []int `json:"items"`
}

// StoriesGetViewers returns a list of story viewers.
//
// https://vk.com/dev/stories.getViewers
func (vk *VK) StoriesGetViewers(params map[string]string) (response StoriesGetViewersResponse, vkErr Error) {
	params["extended"] = "0"
	vk.RequestUnmarshal("stories.getViewers", params, &response, &vkErr)
	return
}

// StoriesGetViewersExtendedResponse struct
type StoriesGetViewersExtendedResponse struct {
	Count int                `json:"count"`
	Items []object.UsersUser `json:"items"`
}

// StoriesGetViewersExtended returns a list of story viewers.
//
// https://vk.com/dev/stories.getViewers
func (vk *VK) StoriesGetViewersExtended(params map[string]string) (response StoriesGetViewersExtendedResponse, vkErr Error) {
	params["extended"] = "1"
	vk.RequestUnmarshal("stories.getViewers", params, &response, &vkErr)
	return
}

// StoriesHideAllReplies hides all replies in the last 24 hours from the user to current user's stories.
//
// https://vk.com/dev/stories.hideAllReplies
func (vk *VK) StoriesHideAllReplies(params map[string]string) (vkErr Error) {
	_, vkErr = vk.Request("stories.hideAllReplies", params)
	return
}

// StoriesHideReply hides the reply to the current user's story.
//
// https://vk.com/dev/stories.hideReply
func (vk *VK) StoriesHideReply(params map[string]string) (vkErr Error) {
	_, vkErr = vk.Request("stories.hideReply", params)
	return
}

// StoriesUnbanOwner allows to show stories from hidden sources in current user's feed.
//
// https://vk.com/dev/stories.unbanOwner
func (vk *VK) StoriesUnbanOwner(params map[string]string) (vkErr Error) {
	_, vkErr = vk.Request("stories.unbanOwner", params)
	return
}
