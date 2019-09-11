package api // import "github.com/SevereCloud/vksdk/api"

import (
	"github.com/SevereCloud/vksdk/object"
)

// AppsDeleteAppRequests deletes all request notifications from the current app.
//
// https://vk.com/dev/apps.deleteAppRequests
func (vk *VK) AppsDeleteAppRequests(params map[string]string) (response int, err error) {
	err = vk.RequestUnmarshal("apps.deleteAppRequests", params, &response)
	return
}

// AppsGetResponse struct
type AppsGetResponse struct {
	Count    int                  `json:"count"`
	Items    []object.AppsApp     `json:"items"`
	Profiles []object.UsersUser   `json:"profiles"`
	Groups   []object.GroupsGroup `json:"groups"`
}

// AppsGet returns applications data.
//
// https://vk.com/dev/apps.get
func (vk *VK) AppsGet(params map[string]string) (response AppsGetResponse, err error) {
	err = vk.RequestUnmarshal("apps.get", params, &response)
	return
}

// AppsGetCatalogResponse struct
type AppsGetCatalogResponse struct {
	Count    int                  `json:"count"`
	Items    []object.AppsApp     `json:"items"`
	Profiles []object.UsersUser   `json:"profiles"`
	Groups   []object.GroupsGroup `json:"groups"`
}

// AppsGetCatalog returns a list of applications (apps) available to users in the App Catalog.
//
// https://vk.com/dev/apps.getCatalog
func (vk *VK) AppsGetCatalog(params map[string]string) (response AppsGetCatalogResponse, err error) {
	err = vk.RequestUnmarshal("apps.getCatalog", params, &response)
	return
}

// AppsGetFriendsListResponse struct
type AppsGetFriendsListResponse struct {
	Count int   `json:"count"`
	Items []int `json:"profiles"`
}

// AppsGetFriendsList creates friends list for requests and invites in current app.
//
// extended=0
//
// https://vk.com/dev/apps.getFriendsList
func (vk *VK) AppsGetFriendsList(params map[string]string) (response AppsGetFriendsListResponse, err error) {
	params["extended"] = "0"
	err = vk.RequestUnmarshal("apps.getFriendsList", params, &response)
	return
}

// AppsGetFriendsListExtendedResponse struct
type AppsGetFriendsListExtendedResponse struct {
	Count int                `json:"count"`
	Items []object.UsersUser `json:"profiles"`
}

// AppsGetFriendsListExtended creates friends list for requests and invites in current app.
//
// extended=1
//
// https://vk.com/dev/apps.getFriendsList
func (vk *VK) AppsGetFriendsListExtended(params map[string]string) (response AppsGetFriendsListExtendedResponse, err error) {
	params["extended"] = "1"
	err = vk.RequestUnmarshal("apps.getFriendsList", params, &response)
	return
}

// AppsGetLeaderboardResponse struct
type AppsGetLeaderboardResponse struct {
	Count int                      `json:"count"`
	Items []object.AppsLeaderboard `json:"items"`
}

// AppsGetLeaderboard returns players rating in the game.
//
// extended=0
//
// https://vk.com/dev/apps.getLeaderboard
func (vk *VK) AppsGetLeaderboard(params map[string]string) (response AppsGetLeaderboardResponse, err error) {
	params["extended"] = "0"
	err = vk.RequestUnmarshal("apps.getLeaderboard", params, &response)
	return
}

// AppsGetLeaderboardExtendedResponse struct
type AppsGetLeaderboardExtendedResponse struct {
	Count int `json:"count"`
	Items []struct {
		Score  int `json:"score"`
		UserID int `json:"user_id"`
	} `json:"items"`
	Profiles []object.UsersUser `json:"profiles"`
}

// AppsGetLeaderboardExtended returns players rating in the game.
//
// extended=1
//
// https://vk.com/dev/apps.getLeaderboard
func (vk *VK) AppsGetLeaderboardExtended(params map[string]string) (response AppsGetLeaderboardExtendedResponse, err error) {
	params["extended"] = "1"
	err = vk.RequestUnmarshal("apps.getLeaderboard", params, &response)
	return
}

// AppsGetScopesResponse struct
type AppsGetScopesResponse struct {
	Count int                `json:"count"`
	Items []object.AppsScope `json:"items"`
}

// AppsGetScopes x
//
// https://vk.com/dev/apps.getScopes
func (vk *VK) AppsGetScopes(params map[string]string) (response AppsGetScopesResponse, err error) {
	err = vk.RequestUnmarshal("apps.getScopes", params, &response)
	return
}

// AppsGetScore returns user score in app.
//
// NOTE: vk wtf!?
//
// https://vk.com/dev/apps.getScore
func (vk *VK) AppsGetScore(params map[string]string) (response string, err error) {
	err = vk.RequestUnmarshal("apps.getScore", params, &response)
	return
}

// AppsSendRequest sends a request to another user in an app that uses VK authorization.
//
// https://vk.com/dev/apps.sendRequest
func (vk *VK) AppsSendRequest(params map[string]string) (response int, err error) {
	err = vk.RequestUnmarshal("apps.sendRequest", params, &response)
	return
}
