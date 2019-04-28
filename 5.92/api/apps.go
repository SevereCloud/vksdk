package api // import "github.com/severecloud/vksdk/5.92/api"

import (
	"encoding/json"

	"github.com/severecloud/vksdk/5.92/object"
)

// AppsDeleteAppRequests deletes all request notifications from the current app.
// https://vk.com/dev/apps.deleteAppRequests
func (vk VK) AppsDeleteAppRequests() (vkErr Error) {
	_, vkErr = vk.Request("apps.deleteAppRequests", map[string]string{})

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
// https://vk.com/dev/apps.get
func (vk VK) AppsGet(params map[string]string) (response AppsGetResponse, vkErr Error) {
	rawResponse, vkErr := vk.Request("apps.get", params)
	if vkErr.Code != 0 {
		return
	}

	err := json.Unmarshal(rawResponse, &response)
	if err != nil {
		panic(err)
	}

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
// https://vk.com/dev/apps.getCatalog
func (vk VK) AppsGetCatalog(params map[string]string) (response AppsGetCatalogResponse, vkErr Error) {
	rawResponse, vkErr := vk.Request("apps.getCatalog", params)
	if vkErr.Code != 0 {
		return
	}

	err := json.Unmarshal(rawResponse, &response)
	if err != nil {
		panic(err)
	}

	return
}

// AppsGetFriendsListResponse struct
type AppsGetFriendsListResponse struct {
	Count int                `json:"count"`
	Items []object.UsersUser `json:"profiles"`
}

// AppsGetFriendsList creates friends list for requests and invites in current app.
// https://vk.com/dev/apps.getFriendsList
func (vk VK) AppsGetFriendsList(params map[string]string) (response AppsGetFriendsListResponse, vkErr Error) {
	rawResponse, vkErr := vk.Request("apps.getFriendsList", params)
	if vkErr.Code != 0 {
		return
	}

	err := json.Unmarshal(rawResponse, &response)
	if err != nil {
		panic(err)
	}

	return
}

// AppsGetLeaderboardResponse struct
type AppsGetLeaderboardResponse struct {
	Count int `json:"count"`
	Items []struct {
		Score  int `json:"score"`
		UserID int `json:"user_id"`
	} `json:"items"`
}

// AppsGetLeaderboard returns players rating in the game.
// https://vk.com/dev/apps.getLeaderboard
func (vk VK) AppsGetLeaderboard(params map[string]string) (response AppsGetLeaderboardResponse, vkErr Error) {
	params["extended"] = "0"
	rawResponse, vkErr := vk.Request("apps.getLeaderboard", params)
	if vkErr.Code != 0 {
		return
	}

	err := json.Unmarshal(rawResponse, &response)
	if err != nil {
		panic(err)
	}

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
// https://vk.com/dev/apps.getLeaderboard
func (vk VK) AppsGetLeaderboardExtended(params map[string]string) (response AppsGetLeaderboardExtendedResponse, vkErr Error) {
	params["extended"] = "1"
	rawResponse, vkErr := vk.Request("apps.getLeaderboard", params)
	if vkErr.Code != 0 {
		return
	}

	err := json.Unmarshal(rawResponse, &response)
	if err != nil {
		panic(err)
	}

	return
}

// AppsGetScopesResponse struct
type AppsGetScopesResponse struct {
	Count int `json:"count"`
	Items []struct {
		Name  string `json:"name"`
		Title string `json:"title"`
	} `json:"items"`
}

// AppsGetScopes x
// https://vk.com/dev/apps.getScopes
func (vk VK) AppsGetScopes(params map[string]string) (response AppsGetScopesResponse, vkErr Error) {
	rawResponse, vkErr := vk.Request("apps.getScopes", params)
	if vkErr.Code != 0 {
		return
	}

	err := json.Unmarshal(rawResponse, &response)
	if err != nil {
		panic(err)
	}

	return
}

// AppsGetScoreResponse struct
// NOTE: vk wtf!?
type AppsGetScoreResponse string

// AppsGetScore returns user score in app.
// https://vk.com/dev/apps.getScore
func (vk VK) AppsGetScore(params map[string]string) (response AppsGetScoreResponse, vkErr Error) {
	rawResponse, vkErr := vk.Request("apps.getScore", params)
	if vkErr.Code != 0 {
		return
	}

	err := json.Unmarshal(rawResponse, &response)
	if err != nil {
		panic(err)
	}

	return
}

// AppsSendRequestResponse struct
type AppsSendRequestResponse int

// AppsSendRequest sends a request to another user in an app that uses VK authorization.
// https://vk.com/dev/apps.sendRequest
func (vk VK) AppsSendRequest(params map[string]string) (response AppsSendRequestResponse, vkErr Error) {
	rawResponse, vkErr := vk.Request("apps.sendRequest", params)
	if vkErr.Code != 0 {
		return
	}

	err := json.Unmarshal(rawResponse, &response)
	if err != nil {
		panic(err)
	}

	return
}
