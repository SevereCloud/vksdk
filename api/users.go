package api // import "github.com/SevereCloud/vksdk/api"

import (
	"github.com/SevereCloud/vksdk/object"
)

// UsersGetResponse users.get response
type UsersGetResponse []object.UsersUser

// UsersGet returns detailed information on users
//
// https://vk.com/dev/users.get
func (vk *VK) UsersGet(params Params) (response UsersGetResponse, err error) {
	err = vk.RequestUnmarshal("users.get", params, &response)
	return
}

// UsersGetFollowersResponse struct
type UsersGetFollowersResponse struct {
	Count int   `json:"count"`
	Items []int `json:"items"`
}

// UsersGetFollowers returns a list of IDs of followers of the user in question, sorted by date added, most recent first.
//
// fields=""
//
// https://vk.com/dev/users.getFollowers
func (vk *VK) UsersGetFollowers(params Params) (response UsersGetFollowersResponse, err error) {
	params["fields"] = ""
	err = vk.RequestUnmarshal("users.getFollowers", params, &response)

	return
}

// UsersGetFollowersFieldsResponse struct
type UsersGetFollowersFieldsResponse struct {
	Count int                `json:"count"`
	Items []object.UsersUser `json:"items"`
}

// UsersGetFollowersFields returns a list of IDs of followers of the user in question, sorted by date added, most recent first.
//
// fields not empty
//
// https://vk.com/dev/users.getFollowers
func (vk *VK) UsersGetFollowersFields(params Params) (response UsersGetFollowersFieldsResponse, err error) {
	if v, prs := params["fields"]; v == "" || !prs {
		params["fields"] = "id"
	}

	err = vk.RequestUnmarshal("users.getFollowers", params, &response)

	return
}

// UsersGetSubscriptionsResponse struct
type UsersGetSubscriptionsResponse struct {
	Users struct {
		Count int   `json:"count"`
		Items []int `json:"items"`
	} `json:"users"`
	Groups struct {
		Count int   `json:"count"`
		Items []int `json:"items"`
	} `json:"groups"`
}

// UsersGetSubscriptions returns a list of IDs of users and public pages followed by the user.
//
// extended=0
//
// https://vk.com/dev/users.getSubscriptions
// BUG(SevereCloud): UsersGetSubscriptions bad response with extended=1
func (vk *VK) UsersGetSubscriptions(params Params) (response UsersGetSubscriptionsResponse, err error) {
	params["extended"] = "0"
	err = vk.RequestUnmarshal("users.getSubscriptions", params, &response)

	return
}

// UsersReport reports (submits a complain about) a user.
//
// https://vk.com/dev/users.report
func (vk *VK) UsersReport(params Params) (response int, err error) {
	err = vk.RequestUnmarshal("users.report", params, &response)
	return
}

// UsersSearchResponse struct
type UsersSearchResponse struct {
	Count int                `json:"count"`
	Items []object.UsersUser `json:"items"`
}

// UsersSearch returns a list of users matching the search criteria.
//
// https://vk.com/dev/users.search
func (vk *VK) UsersSearch(params Params) (response UsersSearchResponse, err error) {
	err = vk.RequestUnmarshal("users.search", params, &response)
	return
}
