package vksdk

import "encoding/json"

// FriendsAddResponse struct
type FriendsAddResponse struct{}

// TODO friends.add Approves or creates a friend request.

// FriendsAddListResponse struct
type FriendsAddListResponse struct{}

// TODO friends.addList Creates a new friend list for the current user.

// FriendsAreFriendsResponse struct
type FriendsAreFriendsResponse struct{}

// TODO friends.areFriends Checks the current user's friendship status with other specified users.

// FriendsDeleteResponse struct
type FriendsDeleteResponse struct{}

// TODO friends.delete Declines a friend request or deletes a user from the current user's friend list.

// FriendsDeleteAllRequestsResponse struct
type FriendsDeleteAllRequestsResponse struct{}

// TODO friends.deleteAllRequests Marks all incoming friend requests as viewed.

// FriendsDeleteListResponse struct
type FriendsDeleteListResponse struct{}

// TODO friends.deleteList Deletes a friend list of the current user.

// FriendsEditResponse struct
type FriendsEditResponse struct{}

// TODO friends.edit Edits the friend lists of the selected user.

// FriendsEditListResponse struct
type FriendsEditListResponse struct{}

// TODO friends.editList Edits a friend list of the current user.

// FriendsGetResponse struct
type FriendsGetResponse struct {
	Count int         `json:"count"`
	Items []usersUser `json:"items"`
}

// FriendsGet returns a list of user IDs or detailed information about a user's friends
// TODO testing
func (vk *VK) FriendsGet(params map[string]string) (response FriendsGetResponse, err error) {
	rawResponse, err := vk.Request("friends.get", params)
	if err != nil {
		return
	}

	err = json.Unmarshal(rawResponse, &response)
	if err != nil {
		panic(err)
	}

	return
}

// FriendsGetAppUsersResponse struct
type FriendsGetAppUsersResponse struct{}

// TODO friends.getAppUsers Returns a list of IDs of the current user's friends who installed the application.

// FriendsGetByPhonesResponse struct
type FriendsGetByPhonesResponse struct{}

// TODO friends.getByPhones Returns a list of the current user's friends whose phone numbers, validated or specified in a profile, are in a given list.

// FriendsGetListsResponse struct
type FriendsGetListsResponse struct{}

// TODO friends.getLists Returns a list of the user's friend lists.

// FriendsGetMutualResponse struct
type FriendsGetMutualResponse struct{}

// TODO friends.getMutual Returns a list of user IDs of the mutual friends of two users.

// FriendsGetOnlineResponse struct
type FriendsGetOnlineResponse struct{}

// TODO friends.getOnline Returns a list of user IDs of a user's friends who are online.

// FriendsGetRecentResponse struct
type FriendsGetRecentResponse struct{}

// TODO friends.getRecent Returns a list of user IDs of the current user's recently added friends.

// FriendsGetRequestsResponse struct
type FriendsGetRequestsResponse struct{}

// TODO friends.getRequests Returns information about the current user's incoming and outgoing friend requests.

// FriendsGetSuggestionsResponse struct
type FriendsGetSuggestionsResponse struct{}

// TODO friends.getSuggestions Returns a list of profiles of users whom the current user may know.

// FriendsSearchResponse struct
type FriendsSearchResponse struct{}

// TODO friends.search Returns a list of friends matching the search criteria.
