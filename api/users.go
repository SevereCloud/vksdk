package api // import "github.com/severecloud/vksdk/api"

import (
	"encoding/json"

	"github.com/severecloud/vksdk/object"
)

// UsersGetResponse users.get response
type UsersGetResponse []object.UsersUser

// UsersGet returns detailed information on users
// https://vk.com/dev/users.get
func (vk VK) UsersGet(params map[string]string) (response UsersGetResponse, vkErr Error) {
	rawResponse, vkErr := vk.Request("users.get", params)
	if vkErr.Code != 0 {
		return
	}
	err := json.Unmarshal(rawResponse, &response)
	if err != nil {
		panic(err)
	}

	return
}

// UsersGetFollowersResponse struct
type UsersGetFollowersResponse struct{}

// TODO: users.getFollowers returns a list of IDs of followers of the user in question, sorted by date added, most recent first.
// https://vk.com/dev/users.getFollowers

// UsersGetSubscriptionsResponse struct
type UsersGetSubscriptionsResponse struct{}

// TODO: users.getSubscriptions returns a list of IDs of users and public pages followed by the user.
// https://vk.com/dev/users.getSubscriptions

// UsersIsAppUserResponse struct
type UsersIsAppUserResponse struct{}

// TODO: users.isAppUser returns information whether a user installed the application.
// https://vk.com/dev/users.isAppUser

// UsersReportResponse struct
type UsersReportResponse struct{}

// TODO: users.report reports (submits a complain about) a user.
// https://vk.com/dev/users.report

// UsersSearchResponse struct
type UsersSearchResponse struct{}

// TODO: users.search returns a list of users matching the search criteria.
// https://vk.com/dev/users.search
