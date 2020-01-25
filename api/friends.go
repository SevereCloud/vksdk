package api // import "github.com/SevereCloud/vksdk/api"

import (
	"github.com/SevereCloud/vksdk/object"
)

// FriendsAdd approves or creates a friend request.
//
// https://vk.com/dev/friends.add
func (vk *VK) FriendsAdd(params Params) (response int, err error) {
	err = vk.RequestUnmarshal("friends.add", params, &response)
	return
}

// FriendsAddListResponse struct
type FriendsAddListResponse struct {
	ListID int `json:"list_id"`
}

// FriendsAddList creates a new friend list for the current user.
//
// https://vk.com/dev/friends.addList
func (vk *VK) FriendsAddList(params Params) (response FriendsAddListResponse, err error) {
	err = vk.RequestUnmarshal("friends.addList", params, &response)
	return
}

// FriendsAreFriendsResponse struct
type FriendsAreFriendsResponse []object.FriendsFriendStatus

// FriendsAreFriends checks the current user's friendship status with other specified users.
//
// https://vk.com/dev/friends.areFriends
func (vk *VK) FriendsAreFriends(params Params) (response FriendsAreFriendsResponse, err error) {
	err = vk.RequestUnmarshal("friends.areFriends", params, &response)
	return
}

// FriendsDeleteResponse struct
type FriendsDeleteResponse struct {
	Success           object.BaseBoolInt `json:"success"`
	FriendDeleted     object.BaseBoolInt `json:"friend_deleted"`
	OutRequestDeleted object.BaseBoolInt `json:"out_request_deleted"`
	InRequestDeleted  object.BaseBoolInt `json:"in_request_deleted"`
	SuggestionDeleted object.BaseBoolInt `json:"suggestion_deleted"`
}

// FriendsDelete declines a friend request or deletes a user from the current user's friend list.
//
// https://vk.com/dev/friends.delete
func (vk *VK) FriendsDelete(params Params) (response FriendsDeleteResponse, err error) {
	err = vk.RequestUnmarshal("friends.delete", params, &response)
	return
}

// FriendsDeleteAllRequests marks all incoming friend requests as viewed.
//
// https://vk.com/dev/friends.deleteAllRequests
func (vk *VK) FriendsDeleteAllRequests(params Params) (response int, err error) {
	err = vk.RequestUnmarshal("friends.deleteAllRequests", params, &response)
	return
}

// FriendsDeleteList deletes a friend list of the current user.
//
// https://vk.com/dev/friends.deleteList
func (vk *VK) FriendsDeleteList(params Params) (response int, err error) {
	err = vk.RequestUnmarshal("friends.deleteList", params, &response)
	return
}

// FriendsEdit edits the friend lists of the selected user.
//
// https://vk.com/dev/friends.edit
func (vk *VK) FriendsEdit(params Params) (response int, err error) {
	err = vk.RequestUnmarshal("friends.edit", params, &response)
	return
}

// FriendsEditList edits a friend list of the current user.
//
// https://vk.com/dev/friends.editList
func (vk *VK) FriendsEditList(params Params) (response int, err error) {
	err = vk.RequestUnmarshal("friends.editList", params, &response)
	return
}

// FriendsGetResponse struct
type FriendsGetResponse struct {
	Count int   `json:"count"`
	Items []int `json:"items"`
}

// FriendsGet returns a list of user IDs or detailed information about a user's friends
//
// https://vk.com/dev/friends.get
func (vk *VK) FriendsGet(params Params) (response FriendsGetResponse, err error) {
	err = vk.RequestUnmarshal("friends.get", params, &response)
	return
}

// FriendsGetFieldsResponse struct
type FriendsGetFieldsResponse struct {
	Count int                          `json:"count"`
	Items []object.FriendsUserXtrLists `json:"items"`
}

// FriendsGetFields returns a list of user IDs or detailed information about a user's friends
//
// https://vk.com/dev/friends.get
func (vk *VK) FriendsGetFields(params Params) (response FriendsGetFieldsResponse, err error) {
	if v, prs := params["fields"]; v == "" || !prs {
		params["fields"] = "id"
	}

	err = vk.RequestUnmarshal("friends.get", params, &response)

	return
}

// FriendsGetAppUsersResponse struct
type FriendsGetAppUsersResponse []int

// FriendsGetAppUsers returns a list of IDs of the current user's friends who installed the application.
//
// https://vk.com/dev/friends.getAppUsers
func (vk *VK) FriendsGetAppUsers(params Params) (response FriendsGetAppUsersResponse, err error) {
	err = vk.RequestUnmarshal("friends.getAppUsers", params, &response)
	return
}

// FriendsGetByPhonesResponse struct
type FriendsGetByPhonesResponse []object.FriendsUserXtrPhone

// FriendsGetByPhones returns a list of the current user's friends
// whose phone numbers, validated or specified in a profile, are in a given list.
//
// https://vk.com/dev/friends.getByPhones
func (vk *VK) FriendsGetByPhones(params Params) (response FriendsGetByPhonesResponse, err error) {
	err = vk.RequestUnmarshal("friends.getByPhones", params, &response)
	return
}

// FriendsGetListsResponse struct
type FriendsGetListsResponse struct {
	Count int                         `json:"count"`
	Items []object.FriendsFriendsList `json:"items"`
}

// FriendsGetLists returns a list of the user's friend lists.
//
// https://vk.com/dev/friends.getLists
func (vk *VK) FriendsGetLists(params Params) (response FriendsGetListsResponse, err error) {
	err = vk.RequestUnmarshal("friends.getLists", params, &response)
	return
}

// FriendsGetMutualResponse struct
type FriendsGetMutualResponse []int

// FriendsGetMutual returns a list of user IDs of the mutual friends of two users.
//
// https://vk.com/dev/friends.getMutual
func (vk *VK) FriendsGetMutual(params Params) (response FriendsGetMutualResponse, err error) {
	err = vk.RequestUnmarshal("friends.getMutual", params, &response)
	return
}

// FriendsGetOnline returns a list of user IDs of a user's friends who are online.
//
// online_mobile=0
//
// https://vk.com/dev/friends.getOnline
func (vk *VK) FriendsGetOnline(params Params) (response []int, err error) {
	params["online_mobile"] = false
	err = vk.RequestUnmarshal("friends.getOnline", params, &response)

	return
}

// FriendsGetOnlineOnlineMobileResponse struct
type FriendsGetOnlineOnlineMobileResponse struct {
	Online       []int `json:"online"`
	OnlineMobile []int `json:"online_mobile"`
}

// FriendsGetOnlineOnlineMobile returns a list of user IDs of a user's friends who are online.
//
// online_mobile=1
//
// https://vk.com/dev/friends.getOnline
func (vk *VK) FriendsGetOnlineOnlineMobile(params Params) (response FriendsGetOnlineOnlineMobileResponse, err error) {
	params["online_mobile"] = true
	err = vk.RequestUnmarshal("friends.getOnline", params, &response)

	return
}

// FriendsGetRecentResponse struct
type FriendsGetRecentResponse []int

// FriendsGetRecent returns a list of user IDs of the current user's recently added friends.
//
// https://vk.com/dev/friends.getRecent
func (vk *VK) FriendsGetRecent(params Params) (response FriendsGetRecentResponse, err error) {
	err = vk.RequestUnmarshal("friends.getRecent", params, &response)
	return
}

// FriendsGetRequestsResponse struct
type FriendsGetRequestsResponse struct {
	Count int   `json:"count"` // Total requests number
	Items []int `json:"items"`
}

// FriendsGetRequests returns information about the current user's incoming and outgoing friend requests.
//
// https://vk.com/dev/friends.getRequests
func (vk *VK) FriendsGetRequests(params Params) (response FriendsGetRequestsResponse, err error) {
	params["need_mutual"] = false
	params["extended"] = false
	err = vk.RequestUnmarshal("friends.getRequests", params, &response)

	return
}

// FriendsGetRequestsNeedMutualResponse struct
type FriendsGetRequestsNeedMutualResponse struct {
	Count int                      `json:"count"` // Total requests number
	Items []object.FriendsRequests `json:"items"`
}

// FriendsGetRequestsNeedMutual returns information about the current user's incoming and outgoing friend requests.
//
// https://vk.com/dev/friends.getRequests
func (vk *VK) FriendsGetRequestsNeedMutual(params Params) (response FriendsGetRequestsNeedMutualResponse, err error) {
	params["need_mutual"] = true
	params["extended"] = false
	err = vk.RequestUnmarshal("friends.getRequests", params, &response)

	return
}

// FriendsGetRequestsExtendedResponse struct
type FriendsGetRequestsExtendedResponse struct {
	Count int                                `json:"count"`
	Items []object.FriendsRequestsXtrMessage `json:"items"`
}

// FriendsGetRequestsExtended returns information about the current user's incoming and outgoing friend requests.
//
// https://vk.com/dev/friends.getRequests
func (vk *VK) FriendsGetRequestsExtended(params Params) (response FriendsGetRequestsExtendedResponse, err error) {
	params["need_mutual"] = false
	params["extended"] = true
	err = vk.RequestUnmarshal("friends.getRequests", params, &response)

	return
}

// FriendsGetSuggestionsResponse struct
type FriendsGetSuggestionsResponse struct {
	Count int                `json:"count"`
	Items []object.UsersUser `json:"items"`
}

// FriendsGetSuggestions returns a list of profiles of users whom the current user may know.
//
// https://vk.com/dev/friends.getSuggestions
func (vk *VK) FriendsGetSuggestions(params Params) (response FriendsGetSuggestionsResponse, err error) {
	err = vk.RequestUnmarshal("friends.getSuggestions", params, &response)
	return
}

// FriendsSearchResponse struct
type FriendsSearchResponse struct {
	Count int                `json:"count"`
	Items []object.UsersUser `json:"items"`
}

// FriendsSearch returns a list of friends matching the search criteria.
//
// https://vk.com/dev/friends.search
func (vk *VK) FriendsSearch(params Params) (response FriendsSearchResponse, err error) {
	err = vk.RequestUnmarshal("friends.search", params, &response)
	return
}
