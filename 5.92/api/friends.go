package api // import "github.com/SevereCloud/vksdk/5.92/api"

import (
	"github.com/SevereCloud/vksdk/5.92/object"
)

// FriendsAddResponse struct
type FriendsAddResponse int

// FriendsAdd approves or creates a friend request.
// https://vk.com/dev/friends.add
func (vk *VK) FriendsAdd(params map[string]string) (response FriendsAddResponse, vkErr Error) {
	vk.requestU("friends.add", params, &response, &vkErr)
	return
}

// FriendsAddListResponse struct
type FriendsAddListResponse struct {
	ListID int `json:"list_id"`
}

// FriendsAddList creates a new friend list for the current user.
// https://vk.com/dev/friends.addList
func (vk *VK) FriendsAddList(params map[string]string) (response FriendsAddListResponse, vkErr Error) {
	vk.requestU("friends.addList", params, &response, &vkErr)
	return
}

// FriendsAreFriendsResponse struct
type FriendsAreFriendsResponse []object.FriendsFriendStatus

// FriendsAreFriends checks the current user's friendship status with other specified users.
// https://vk.com/dev/friends.areFriends
func (vk *VK) FriendsAreFriends(params map[string]string) (response FriendsAreFriendsResponse, vkErr Error) {
	vk.requestU("friends.areFriends", params, &response, &vkErr)
	return
}

// FriendsDeleteResponse struct
type FriendsDeleteResponse struct {
	Success           int `json:"success"`
	FriendDeleted     int `json:"friend_deleted"`
	OutRequestDeleted int `json:"out_request_deleted"`
	InRequestDeleted  int `json:"in_request_deleted"`
	SuggestionDeleted int `json:"suggestion_deleted"`
}

// FriendsDelete declines a friend request or deletes a user from the current user's friend list.
// https://vk.com/dev/friends.delete
func (vk *VK) FriendsDelete(params map[string]string) (response FriendsDeleteResponse, vkErr Error) {
	vk.requestU("friends.delete", params, &response, &vkErr)
	return
}

// FriendsDeleteAllRequests marks all incoming friend requests as viewed.
// https://vk.com/dev/friends.deleteAllRequests
func (vk *VK) FriendsDeleteAllRequests() (vkErr Error) {
	_, vkErr = vk.Request("friends.deleteAllRequests", map[string]string{})
	return
}

// FriendsDeleteList deletes a friend list of the current user.
// https://vk.com/dev/friends.deleteList
func (vk *VK) FriendsDeleteList(params map[string]string) (vkErr Error) {
	_, vkErr = vk.Request("friends.deleteList", params)
	return
}

// FriendsEdit edits the friend lists of the selected user.
// https://vk.com/dev/friends.edit
func (vk *VK) FriendsEdit(params map[string]string) (vkErr Error) {
	_, vkErr = vk.Request("friends.edit", params)
	return
}

// FriendsEditList edits a friend list of the current user.
// https://vk.com/dev/friends.editList
func (vk *VK) FriendsEditList(params map[string]string) (vkErr Error) {
	_, vkErr = vk.Request("friends.editList", params)
	return
}

// FriendsGetResponse struct
type FriendsGetResponse struct {
	Count int                `json:"count"`
	Items []object.UsersUser `json:"items"`
}

// FriendsGet returns a list of user IDs or detailed information about a user's friends
// https://vk.com/dev/friends.get
func (vk *VK) FriendsGet(params map[string]string) (response FriendsGetResponse, vkErr Error) {
	vk.requestU("friends.get", params, &response, &vkErr)
	return
}

// FriendsGetAppUsersResponse struct
type FriendsGetAppUsersResponse []int

// FriendsGetAppUsers returns a list of IDs of the current user's friends who installed the application.
// https://vk.com/dev/friends.getAppUsers
func (vk *VK) FriendsGetAppUsers(params map[string]string) (response FriendsGetAppUsersResponse, vkErr Error) {
	vk.requestU("riends.getAppUsers", params, &response, &vkErr)
	return
}

// FriendsGetByPhonesResponse struct
type FriendsGetByPhonesResponse []object.UsersUser

// FriendsGetByPhones returns a list of the current user's friends whose phone numbers, validated or specified in a profile, are in a given list.
// https://vk.com/dev/friends.getByPhones
func (vk *VK) FriendsGetByPhones(params map[string]string) (response FriendsGetByPhonesResponse, vkErr Error) {
	vk.requestU("friends.getByPhones", params, &response, &vkErr)
	return
}

// FriendsGetListsResponse struct
type FriendsGetListsResponse struct {
	Count int                         `json:"count"`
	Items []object.FriendsFriendsList `json:"items"`
}

// FriendsGetLists returns a list of the user's friend lists.
// https://vk.com/dev/friends.getLists
func (vk *VK) FriendsGetLists(params map[string]string) (response FriendsGetListsResponse, vkErr Error) {
	vk.requestU("friends.getLists", params, &response, &vkErr)
	return
}

// FriendsGetMutualResponse struct
type FriendsGetMutualResponse []int

// FriendsGetMutual returns a list of user IDs of the mutual friends of two users.
// https://vk.com/dev/friends.getMutual
func (vk *VK) FriendsGetMutual(params map[string]string) (response FriendsGetMutualResponse, vkErr Error) {
	vk.requestU("friends.getMutual", params, &response, &vkErr)
	return
}

// FriendsGetOnlineResponse struct
type FriendsGetOnlineResponse struct {
	Online       []int `json:"online"`
	OnlineMobile []int `json:"online_mobile"`
}

// FriendsGetOnline returns a list of user IDs of a user's friends who are online.
// https://vk.com/dev/friends.getOnline
func (vk *VK) FriendsGetOnline(params map[string]string) (response FriendsGetOnlineResponse, vkErr Error) {
	vk.requestU("friends.getOnline", params, &response, &vkErr)
	return
}

// FriendsGetRecentResponse struct
type FriendsGetRecentResponse []int

// FriendsGetRecent returns a list of user IDs of the current user's recently added friends.
// https://vk.com/dev/friends.getRecent
func (vk *VK) FriendsGetRecent(params map[string]string) (response FriendsGetRecentResponse, vkErr Error) {
	vk.requestU("friends.getRecent", params, &response, &vkErr)
	return
}

// FriendsGetRequestsResponse struct
type FriendsGetRequestsResponse struct {
	Count int                      `json:"count"`
	Items []object.FriendsRequests `json:"items"`
}

// FriendsGetRequests returns information about the current user's incoming and outgoing friend requests.
// https://vk.com/dev/friends.getRequests
func (vk *VK) FriendsGetRequests(params map[string]string) (response FriendsGetRequestsResponse, vkErr Error) {
	vk.requestU("friends.getRequests", params, &response, &vkErr)
	return
}

// FriendsGetSuggestionsResponse struct
type FriendsGetSuggestionsResponse struct {
	Count int                `json:"count"`
	Items []object.UsersUser `json:"items"`
}

// FriendsGetSuggestions returns a list of profiles of users whom the current user may know.
// https://vk.com/dev/friends.getSuggestions
func (vk *VK) FriendsGetSuggestions(params map[string]string) (response FriendsGetSuggestionsResponse, vkErr Error) {
	vk.requestU("friends.getSuggestions", params, &response, &vkErr)
	return
}

// FriendsSearchResponse struct
type FriendsSearchResponse struct {
	Count int                `json:"count"`
	Items []object.UsersUser `json:"items"`
}

// FriendsSearch returns a list of friends matching the search criteria.
// https://vk.com/dev/friends.search
func (vk *VK) FriendsSearch(params map[string]string) (response FriendsSearchResponse, vkErr Error) {
	vk.requestU("friends.search", params, &response, &vkErr)
	return
}
