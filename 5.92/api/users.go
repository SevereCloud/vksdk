package api // import "github.com/SevereCloud/vksdk/5.92/api"

import (
	"github.com/SevereCloud/vksdk/5.92/object"
)

// UsersGetResponse users.get response
type UsersGetResponse []object.UsersUser

// UsersGet returns detailed information on users
//
// https://vk.com/dev/users.get
func (vk *VK) UsersGet(params map[string]string) (response UsersGetResponse, vkErr Error) {
	vk.RequestUnmarshal("users.get", params, &response, &vkErr)
	return
}

// UsersGetFollowersResponse struct
type UsersGetFollowersResponse []object.UsersUser

// UsersGetFollowers returns a list of IDs of followers of the user in question, sorted by date added, most recent first.
//
// https://vk.com/dev/users.getFollowers
func (vk *VK) UsersGetFollowers(params map[string]string) (response UsersGetFollowersResponse, vkErr Error) {
	vk.RequestUnmarshal("users.getFollowers", params, &response, &vkErr)
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
// https://vk.com/dev/users.getSubscriptions
// BUG(SevereCloud): UsersGetSubscriptions bad response with extended=1
func (vk *VK) UsersGetSubscriptions(params map[string]string) (response UsersGetSubscriptionsResponse, vkErr Error) {
	params["extended"] = "0"
	vk.RequestUnmarshal("users.getSubscriptions", params, &response, &vkErr)
	return
}

// UsersIsAppUserResponse struct
type UsersIsAppUserResponse int

// UsersIsAppUser returns information whether a user installed the application.
//
// https://vk.com/dev/users.isAppUser
func (vk *VK) UsersIsAppUser(params map[string]string) (response UsersIsAppUserResponse, vkErr Error) {
	vk.RequestUnmarshal("users.isAppUser", params, &response, &vkErr)
	return
}

// UsersReport reports (submits a complain about) a user.
//
// https://vk.com/dev/users.report
func (vk *VK) UsersReport(params map[string]string) (vkErr Error) {
	_, vkErr = vk.Request("users.report", params)
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
func (vk *VK) UsersSearch(params map[string]string) (response UsersSearchResponse, vkErr Error) {
	vk.RequestUnmarshal("users.search", params, &response, &vkErr)
	return
}
