package api // import "github.com/SevereCloud/vksdk/api"

import "github.com/SevereCloud/vksdk/object"

// PollsAddVote adds the current user's vote to the selected answer in the poll.
//
// https://vk.com/dev/polls.addVote
func (vk *VK) PollsAddVote(params map[string]string) (response int, err error) {
	err = vk.RequestUnmarshal("polls.addVote", params, &response)
	return
}

// PollsCreateResponse struct
type PollsCreateResponse object.PollsPoll

// PollsCreate creates polls that can be attached to the users' or communities' posts.
//
// https://vk.com/dev/polls.create
func (vk *VK) PollsCreate(params map[string]string) (response PollsCreateResponse, err error) {
	err = vk.RequestUnmarshal("polls.create", params, &response)
	return
}

// PollsDeleteVote deletes the current user's vote from the selected answer in the poll.
//
// https://vk.com/dev/polls.deleteVote
func (vk *VK) PollsDeleteVote(params map[string]string) (response int, err error) {
	err = vk.RequestUnmarshal("polls.deleteVote", params, &response)
	return
}

// PollsEdit edits created polls
//
// https://vk.com/dev/polls.edit
func (vk *VK) PollsEdit(params map[string]string) (response int, err error) {
	err = vk.RequestUnmarshal("polls.edit", params, &response)
	return
}

// PollsGetBackgroundsResponse struct
type PollsGetBackgroundsResponse []object.PollsBackground

// PollsGetBackgrounds return default backgrounds for polls.
//
// https://vk.com/dev/polls.getBackgrounds
func (vk *VK) PollsGetBackgrounds(params map[string]string) (response PollsGetBackgroundsResponse, err error) {
	err = vk.RequestUnmarshal("polls.getBackgrounds", params, &response)
	return
}

// PollsGetByIDResponse struct
type PollsGetByIDResponse object.PollsPoll

// PollsGetByID returns detailed information about a poll by its ID.
//
// https://vk.com/dev/polls.getById
func (vk *VK) PollsGetByID(params map[string]string) (response PollsGetByIDResponse, err error) {
	err = vk.RequestUnmarshal("polls.getById", params, &response)
	return
}

// PollsGetPhotoUploadServerResponse struct
type PollsGetPhotoUploadServerResponse struct {
	UploadURL string `json:"upload_url"`
}

// PollsGetPhotoUploadServer returns a URL for uploading a photo to a poll.
//
// https://vk.com/dev/polls.getPhotoUploadServer
func (vk *VK) PollsGetPhotoUploadServer(params map[string]string) (response PollsGetPhotoUploadServerResponse, err error) {
	err = vk.RequestUnmarshal("polls.getPhotoUploadServer", params, &response)
	return
}

// PollsGetVotersResponse struct
type PollsGetVotersResponse []object.PollsVoters

// PollsGetVoters returns a list of IDs of users who selected specific answers in the poll.
//
// https://vk.com/dev/polls.getVoters
func (vk *VK) PollsGetVoters(params map[string]string) (response PollsGetVotersResponse, err error) {
	err = vk.RequestUnmarshal("polls.getVoters", params, &response)
	return
}

// PollsGetVotersFieldsResponse struct
type PollsGetVotersFieldsResponse []object.PollsVotersFields

// PollsGetVotersFields returns a list of IDs of users who selected specific answers in the poll.
//
// https://vk.com/dev/polls.getVoters
func (vk *VK) PollsGetVotersFields(params map[string]string) (response PollsGetVotersFieldsResponse, err error) {
	err = vk.RequestUnmarshal("polls.getVoters", params, &response)
	return
}

// PollsSavePhotoResponse struct
type PollsSavePhotoResponse object.PollsPhoto

// PollsSavePhoto allows to save poll's uploaded photo.
//
// https://vk.com/dev/polls.savePhoto
func (vk *VK) PollsSavePhoto(params map[string]string) (response PollsSavePhotoResponse, err error) {
	err = vk.RequestUnmarshal("polls.savePhoto", params, &response)
	return
}
