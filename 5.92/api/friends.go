package api // import "github.com/severecloud/vksdk/5.92/api"

import (
	"encoding/json"

	"github.com/severecloud/vksdk/5.92/object"
)

// FriendsAddResponse struct
type FriendsAddResponse struct{}

// TODO: friends.add approves or creates a friend request.
// https://vk.com/dev/friends.add

// FriendsAddListResponse struct
type FriendsAddListResponse struct{}

// TODO: friends.addList creates a new friend list for the current user.
// https://vk.com/dev/friends.addList

// FriendsAreFriendsResponse struct
type FriendsAreFriendsResponse struct{}

// TODO: friends.areFriends checks the current user's friendship status with other specified users.
// https://vk.com/dev/friends.areFriends

// FriendsDeleteResponse struct
type FriendsDeleteResponse struct{}

// TODO: friends.delete declines a friend request or deletes a user from the current user's friend list.
// https://vk.com/dev/friends.delete

// FriendsDeleteAllRequestsResponse struct
type FriendsDeleteAllRequestsResponse struct{}

// TODO: friends.deleteAllRequests marks all incoming friend requests as viewed.
// https://vk.com/dev/friends.deleteAllRequests

// FriendsDeleteListResponse struct
type FriendsDeleteListResponse struct{}

// TODO: friends.deleteList deletes a friend list of the current user.
// https://vk.com/dev/friends.deleteList

// FriendsEditResponse struct
type FriendsEditResponse struct{}

// TODO: friends.edit edits the friend lists of the selected user.
// https://vk.com/dev/friends.edit

// FriendsEditListResponse struct
type FriendsEditListResponse struct{}

// TODO: friends.editList edits a friend list of the current user.
// https://vk.com/dev/friends.editList

// FriendsGetResponse struct
type FriendsGetResponse struct {
	Count int                `json:"count"`
	Items []object.UsersUser `json:"items"`
}

// FriendsGet returns a list of user IDs or detailed information about a user's friends
// TODO: testing
// https://vk.com/dev/friends.get
func (vk VK) FriendsGet(params map[string]string) (response FriendsGetResponse, vkErr Error) {
	rawResponse, vkErr := vk.Request("friends.get", params)
	if vkErr.Code != 0 {
		return
	}

	err := json.Unmarshal(rawResponse, &response)
	if err != nil {
		panic(err)
	}

	return
}

// FriendsGetAppUsersResponse struct
type FriendsGetAppUsersResponse struct{}

// TODO: friends.getAppUsers returns a list of IDs of the current user's friends who installed the application.
// https://vk.com/dev/friends.getAppUsers

// FriendsGetByPhonesResponse struct
type FriendsGetByPhonesResponse struct{}

// TODO: friends.getByPhones returns a list of the current user's friends whose phone numbers, validated or specified in a profile, are in a given list.
// https://vk.com/dev/friends.getByPhones

// FriendsGetListsResponse struct
type FriendsGetListsResponse struct{}

// TODO: friends.getLists returns a list of the user's friend lists.
// https://vk.com/dev/friends.getLists

// FriendsGetMutualResponse struct
type FriendsGetMutualResponse struct{}

// TODO: friends.getMutual returns a list of user IDs of the mutual friends of two users.
// https://vk.com/dev/friends.getMutual

// FriendsGetOnlineResponse struct
type FriendsGetOnlineResponse struct{}

// TODO: friends.getOnline returns a list of user IDs of a user's friends who are online.
// https://vk.com/dev/friends.getOnline

// FriendsGetRecentResponse struct
type FriendsGetRecentResponse struct{}

// TODO: friends.getRecent returns a list of user IDs of the current user's recently added friends.
// https://vk.com/dev/friends.getRecent

// FriendsGetRequestsResponse struct
type FriendsGetRequestsResponse struct{}

// TODO: friends.getRequests returns information about the current user's incoming and outgoing friend requests.
// https://vk.com/dev/friends.getRequests

// FriendsGetSuggestionsResponse struct
type FriendsGetSuggestionsResponse struct{}

// TODO: friends.getSuggestions returns a list of profiles of users whom the current user may know.
// https://vk.com/dev/friends.getSuggestions

// FriendsSearchResponse struct
type FriendsSearchResponse struct{}

// TODO: friends.search returns a list of friends matching the search criteria.
// https://vk.com/dev/friends.search
