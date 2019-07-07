package api // import "github.com/SevereCloud/vksdk/5.92/api"

import "github.com/SevereCloud/vksdk/5.92/object"

// PollsAddVoteResponse struct
type PollsAddVoteResponse int

// PollsAddVote adds the current user's vote to the selected answer in the poll.
//
// https://vk.com/dev/polls.addVote
func (vk *VK) PollsAddVote(params map[string]string) (response PollsAddVoteResponse, vkErr Error) {
	vk.RequestUnmarshal("polls.addVote", params, &response, &vkErr)
	return
}

// PollsCreateResponse struct
type PollsCreateResponse object.PollsPoll

// PollsCreate creates polls that can be attached to the users' or communities' posts.
//
// https://vk.com/dev/polls.create
func (vk *VK) PollsCreate(params map[string]string) (response PollsCreateResponse, vkErr Error) {
	vk.RequestUnmarshal("polls.create", params, &response, &vkErr)
	return
}

// PollsDeleteVoteResponse struct
type PollsDeleteVoteResponse int

// PollsDeleteVote deletes the current user's vote from the selected answer in the poll.
//
// https://vk.com/dev/polls.deleteVote
func (vk *VK) PollsDeleteVote(params map[string]string) (response PollsDeleteVoteResponse, vkErr Error) {
	vk.RequestUnmarshal("polls.deleteVote", params, &response, &vkErr)
	return
}

// PollsEditResponse struct
type PollsEditResponse int

// PollsEdit edits created polls
//
// https://vk.com/dev/polls.edit
func (vk *VK) PollsEdit(params map[string]string) (response PollsEditResponse, vkErr Error) {
	vk.RequestUnmarshal("polls.edit", params, &response, &vkErr)
	return
}

// PollsGetBackgroundsResponse struct
type PollsGetBackgroundsResponse []object.PollsBackground

// PollsGetBackgrounds return default backgrounds for polls.
//
// https://vk.com/dev/polls.getBackgrounds
func (vk *VK) PollsGetBackgrounds(params map[string]string) (response PollsGetBackgroundsResponse, vkErr Error) {
	vk.RequestUnmarshal("polls.getBackgrounds", params, &response, &vkErr)
	return
}

// PollsGetByIDResponse struct
type PollsGetByIDResponse object.PollsPoll

// PollsGetByID returns detailed information about a poll by its ID.
//
// https://vk.com/dev/polls.getById
func (vk *VK) PollsGetByID(params map[string]string) (response PollsGetByIDResponse, vkErr Error) {
	vk.RequestUnmarshal("polls.getById", params, &response, &vkErr)
	return
}

// PollsGetPhotoUploadServerResponse struct
type PollsGetPhotoUploadServerResponse struct {
	UploadURL string `json:"upload_url"`
}

// PollsGetPhotoUploadServer returns a URL for uploading a photo to a poll.
//
// https://vk.com/dev/polls.getPhotoUploadServer
func (vk *VK) PollsGetPhotoUploadServer(params map[string]string) (response PollsGetPhotoUploadServerResponse, vkErr Error) {
	vk.RequestUnmarshal("polls.getPhotoUploadServer", params, &response, &vkErr)
	return
}

// PollsGetVotersResponse struct
type PollsGetVotersResponse []object.PollsVoters

// PollsGetVoters returns a list of IDs of users who selected specific answers in the poll.
//
// https://vk.com/dev/polls.getVoters
func (vk *VK) PollsGetVoters(params map[string]string) (response PollsGetVotersResponse, vkErr Error) {
	vk.RequestUnmarshal("polls.getVoters", params, &response, &vkErr)
	return
}

// PollsGetVotersFieldsResponse struct
type PollsGetVotersFieldsResponse []object.PollsVotersFields

// PollsGetVotersFields returns a list of IDs of users who selected specific answers in the poll.
//
// https://vk.com/dev/polls.getVoters
func (vk *VK) PollsGetVotersFields(params map[string]string) (response PollsGetVotersFieldsResponse, vkErr Error) {
	vk.RequestUnmarshal("polls.getVoters", params, &response, &vkErr)
	return
}

// PollsSavePhotoResponse struct
type PollsSavePhotoResponse object.PollsPhoto

// PollsSavePhoto allows to save poll's uploaded photo.
//
// https://vk.com/dev/polls.savePhoto
func (vk *VK) PollsSavePhoto(params map[string]string) (response PollsSavePhotoResponse, vkErr Error) {
	vk.RequestUnmarshal("polls.savePhoto", params, &response, &vkErr)
	return
}
