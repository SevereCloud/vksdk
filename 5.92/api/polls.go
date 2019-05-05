package api // import "github.com/SevereCloud/vksdk/5.92/api"

// PollsAddVoteResponse struct
type PollsAddVoteResponse struct{}

// TODO: polls.addVote adds the current user's vote to the selected answer in the poll.
// https://vk.com/dev/polls.addVote

// PollsCreateResponse struct
type PollsCreateResponse struct{}

// TODO: polls.create creates polls that can be attached to the users' or communities' posts.
// https://vk.com/dev/polls.create

// PollsDeleteVoteResponse struct
type PollsDeleteVoteResponse struct{}

// TODO: polls.deleteVote deletes the current user's vote from the selected answer in the poll.
// https://vk.com/dev/polls.deleteVote

// PollsEditResponse struct
type PollsEditResponse struct{}

// TODO: polls.edit edits created polls
// https://vk.com/dev/polls.edit

// PollsGetBackgroundsResponse struct
type PollsGetBackgroundsResponse struct{}

// TODO: polls.getBackgrounds return default backgrounds for polls.
// https://vk.com/dev/polls.getBackgrounds

// PollsGetByIDResponse struct
type PollsGetByIDResponse struct{}

// TODO: polls.getById returns detailed information about a poll by its ID.
// https://vk.com/dev/polls.getById

// PollsGetPhotoUploadServerResponse struct
type PollsGetPhotoUploadServerResponse struct{}

// TODO: polls.getPhotoUploadServer returns a URL for uploading a photo to a poll.
// https://vk.com/dev/polls.getPhotoUploadServer

// PollsGetVotersResponse struct
type PollsGetVotersResponse []struct {
	AnswerID int `json:"answer_id"`
	Users    struct {
		Count int   `json:"count"`
		Items []int `json:"items"`
	} `json:"users"`
}

// PollsGetVoters returns a list of IDs of users who selected specific answers in the poll.
// https://vk.com/dev/polls.getVoters
// FIXME: list if no filter
func (vk VK) PollsGetVoters(params map[string]string) (response PollsGetVotersResponse, vkErr Error) {
	vk.requestU("polls.getVoters", params, &response, &vkErr)
	return
}

// PollsSavePhotoResponse struct
type PollsSavePhotoResponse struct{}

// TODO: polls.savePhoto allows to save poll's uploaded photo.
// https://vk.com/dev/polls.savePhoto
