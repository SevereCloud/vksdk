package params // import "github.com/SevereCloud/vksdk/api/params"

import (
	"github.com/SevereCloud/vksdk/api"
)

// FriendsAddBulder builder
//
// Approves or creates a friend request.
//
// https://vk.com/dev/friends.add
type FriendsAddBulder struct {
	api.Params
}

// NewFriendsAddBulder func
func NewFriendsAddBulder() *FriendsAddBulder {
	return &FriendsAddBulder{api.Params{}}
}

// UserID ID of the user whose friend request will be approved or to whom a friend request will be sent.
func (b *FriendsAddBulder) UserID(v int) {
	b.Params["user_id"] = v
}

// Text Text of the message (up to 500 characters) for the friend request, if any.
func (b *FriendsAddBulder) Text(v string) {
	b.Params["text"] = v
}

// Follow '1' to pass an incoming request to followers list.
func (b *FriendsAddBulder) Follow(v bool) {
	b.Params["follow"] = v
}

// FriendsAddListBulder builder
//
// Creates a new friend list for the current user.
//
// https://vk.com/dev/friends.addList
type FriendsAddListBulder struct {
	api.Params
}

// NewFriendsAddListBulder func
func NewFriendsAddListBulder() *FriendsAddListBulder {
	return &FriendsAddListBulder{api.Params{}}
}

// Name Name of the friend list.
func (b *FriendsAddListBulder) Name(v string) {
	b.Params["name"] = v
}

// UserIDs IDs of users to be added to the friend list.
func (b *FriendsAddListBulder) UserIDs(v []int) {
	b.Params["user_ids"] = v
}

// FriendsAreFriendsBulder builder
//
// Checks the current user's friendship status with other specified users.
//
// https://vk.com/dev/friends.areFriends
type FriendsAreFriendsBulder struct {
	api.Params
}

// NewFriendsAreFriendsBulder func
func NewFriendsAreFriendsBulder() *FriendsAreFriendsBulder {
	return &FriendsAreFriendsBulder{api.Params{}}
}

// UserIDs IDs of the users whose friendship status to check.
func (b *FriendsAreFriendsBulder) UserIDs(v []int) {
	b.Params["user_ids"] = v
}

// NeedSign '1' — to return 'sign' field. 'sign' is md5("{id}_{user_id}_{friends_status}_{application_secret}"), where id is current user ID. This field allows to check that data has not been modified by the client. By default: '0'.
func (b *FriendsAreFriendsBulder) NeedSign(v bool) {
	b.Params["need_sign"] = v
}

// FriendsDeleteBulder builder
//
// Declines a friend request or deletes a user from the current user's friend list.
//
// https://vk.com/dev/friends.delete
type FriendsDeleteBulder struct {
	api.Params
}

// NewFriendsDeleteBulder func
func NewFriendsDeleteBulder() *FriendsDeleteBulder {
	return &FriendsDeleteBulder{api.Params{}}
}

// UserID ID of the user whose friend request is to be declined or who is to be deleted from the current user's friend list.
func (b *FriendsDeleteBulder) UserID(v int) {
	b.Params["user_id"] = v
}

// FriendsDeleteListBulder builder
//
// Deletes a friend list of the current user.
//
// https://vk.com/dev/friends.deleteList
type FriendsDeleteListBulder struct {
	api.Params
}

// NewFriendsDeleteListBulder func
func NewFriendsDeleteListBulder() *FriendsDeleteListBulder {
	return &FriendsDeleteListBulder{api.Params{}}
}

// ListID ID of the friend list to delete.
func (b *FriendsDeleteListBulder) ListID(v int) {
	b.Params["list_id"] = v
}

// FriendsEditBulder builder
//
// Edits the friend lists of the selected user.
//
// https://vk.com/dev/friends.edit
type FriendsEditBulder struct {
	api.Params
}

// NewFriendsEditBulder func
func NewFriendsEditBulder() *FriendsEditBulder {
	return &FriendsEditBulder{api.Params{}}
}

// UserID ID of the user whose friend list is to be edited.
func (b *FriendsEditBulder) UserID(v int) {
	b.Params["user_id"] = v
}

// ListIDs IDs of the friend lists to which to add the user.
func (b *FriendsEditBulder) ListIDs(v []int) {
	b.Params["list_ids"] = v
}

// FriendsEditListBulder builder
//
// Edits a friend list of the current user.
//
// https://vk.com/dev/friends.editList
type FriendsEditListBulder struct {
	api.Params
}

// NewFriendsEditListBulder func
func NewFriendsEditListBulder() *FriendsEditListBulder {
	return &FriendsEditListBulder{api.Params{}}
}

// Name Name of the friend list.
func (b *FriendsEditListBulder) Name(v string) {
	b.Params["name"] = v
}

// ListID Friend list ID.
func (b *FriendsEditListBulder) ListID(v int) {
	b.Params["list_id"] = v
}

// UserIDs IDs of users in the friend list.
func (b *FriendsEditListBulder) UserIDs(v []int) {
	b.Params["user_ids"] = v
}

// AddUserIDs (Applies if 'user_ids' parameter is not set.), User IDs to add to the friend list.
func (b *FriendsEditListBulder) AddUserIDs(v []int) {
	b.Params["add_user_ids"] = v
}

// DeleteUserIDs (Applies if 'user_ids' parameter is not set.), User IDs to delete from the friend list.
func (b *FriendsEditListBulder) DeleteUserIDs(v []int) {
	b.Params["delete_user_ids"] = v
}

// FriendsGetBulder builder
//
// Returns a list of user IDs or detailed information about a user's friends.
//
// https://vk.com/dev/friends.get
type FriendsGetBulder struct {
	api.Params
}

// NewFriendsGetBulder func
func NewFriendsGetBulder() *FriendsGetBulder {
	return &FriendsGetBulder{api.Params{}}
}

// UserID User ID. By default, the current user ID.
func (b *FriendsGetBulder) UserID(v int) {
	b.Params["user_id"] = v
}

// Order Sort order: , 'name' — by name (enabled only if the 'fields' parameter is used), 'hints' — by rating, similar to how friends are sorted in My friends section, , This parameter is available only for [vk.com/dev/standalone|desktop applications].
func (b *FriendsGetBulder) Order(v string) {
	b.Params["order"] = v
}

// ListID ID of the friend list returned by the [vk.com/dev/friends.getLists|friends.getLists] method to be used as the source. This parameter is taken into account only when the uid parameter is set to the current user ID. This parameter is available only for [vk.com/dev/standalone|desktop applications].
func (b *FriendsGetBulder) ListID(v int) {
	b.Params["list_id"] = v
}

// Count Number of friends to return.
func (b *FriendsGetBulder) Count(v int) {
	b.Params["count"] = v
}

// Offset Offset needed to return a specific subset of friends.
func (b *FriendsGetBulder) Offset(v int) {
	b.Params["offset"] = v
}

// Fields Profile fields to return. Sample values: 'uid', 'first_name', 'last_name', 'nickname', 'sex', 'bdate' (birthdate), 'city', 'country', 'timezone', 'photo', 'photo_medium', 'photo_big', 'domain', 'has_mobile', 'rate', 'contacts', 'education'.
func (b *FriendsGetBulder) Fields(v []string) {
	b.Params["fields"] = v
}

// NameCase Case for declension of user name and surname: , 'nom' — nominative (default) , 'gen' — genitive , 'dat' — dative , 'acc' — accusative , 'ins' — instrumental , 'abl' — prepositional
func (b *FriendsGetBulder) NameCase(v string) {
	b.Params["name_case"] = v
}

// Ref parameter
func (b *FriendsGetBulder) Ref(v string) {
	b.Params["ref"] = v
}

// FriendsGetByPhonesBulder builder
//
// Returns a list of the current user's friends whose phone numbers, validated or specified in a profile, are in a given list.
//
// https://vk.com/dev/friends.getByPhones
type FriendsGetByPhonesBulder struct {
	api.Params
}

// NewFriendsGetByPhonesBulder func
func NewFriendsGetByPhonesBulder() *FriendsGetByPhonesBulder {
	return &FriendsGetByPhonesBulder{api.Params{}}
}

// Phones List of phone numbers in MSISDN format (maximum 1000). Example: "+79219876543,+79111234567"
func (b *FriendsGetByPhonesBulder) Phones(v []string) {
	b.Params["phones"] = v
}

// Fields Profile fields to return. Sample values: 'nickname', 'screen_name', 'sex', 'bdate' (birthdate), 'city', 'country', 'timezone', 'photo', 'photo_medium', 'photo_big', 'has_mobile', 'rate', 'contacts', 'education', 'online, counters'.
func (b *FriendsGetByPhonesBulder) Fields(v []string) {
	b.Params["fields"] = v
}

// FriendsGetListsBulder builder
//
// Returns a list of the user's friend lists.
//
// https://vk.com/dev/friends.getLists
type FriendsGetListsBulder struct {
	api.Params
}

// NewFriendsGetListsBulder func
func NewFriendsGetListsBulder() *FriendsGetListsBulder {
	return &FriendsGetListsBulder{api.Params{}}
}

// UserID User ID.
func (b *FriendsGetListsBulder) UserID(v int) {
	b.Params["user_id"] = v
}

// ReturnSystem '1' — to return system friend lists. By default: '0'.
func (b *FriendsGetListsBulder) ReturnSystem(v bool) {
	b.Params["return_system"] = v
}

// FriendsGetMutualBulder builder
//
// Returns a list of user IDs of the mutual friends of two users.
//
// https://vk.com/dev/friends.getMutual
type FriendsGetMutualBulder struct {
	api.Params
}

// NewFriendsGetMutualBulder func
func NewFriendsGetMutualBulder() *FriendsGetMutualBulder {
	return &FriendsGetMutualBulder{api.Params{}}
}

// SourceUID ID of the user whose friends will be checked against the friends of the user specified in 'target_uid'.
func (b *FriendsGetMutualBulder) SourceUID(v int) {
	b.Params["source_uid"] = v
}

// TargetUID ID of the user whose friends will be checked against the friends of the user specified in 'source_uid'.
func (b *FriendsGetMutualBulder) TargetUID(v int) {
	b.Params["target_uid"] = v
}

// TargetUids IDs of the users whose friends will be checked against the friends of the user specified in 'source_uid'.
func (b *FriendsGetMutualBulder) TargetUids(v []int) {
	b.Params["target_uids"] = v
}

// Order Sort order: 'random' — random order
func (b *FriendsGetMutualBulder) Order(v string) {
	b.Params["order"] = v
}

// Count Number of mutual friends to return.
func (b *FriendsGetMutualBulder) Count(v int) {
	b.Params["count"] = v
}

// Offset Offset needed to return a specific subset of mutual friends.
func (b *FriendsGetMutualBulder) Offset(v int) {
	b.Params["offset"] = v
}

// FriendsGetOnlineBulder builder
//
// Returns a list of user IDs of a user's friends who are online.
//
// https://vk.com/dev/friends.getOnline
type FriendsGetOnlineBulder struct {
	api.Params
}

// NewFriendsGetOnlineBulder func
func NewFriendsGetOnlineBulder() *FriendsGetOnlineBulder {
	return &FriendsGetOnlineBulder{api.Params{}}
}

// UserID User ID.
func (b *FriendsGetOnlineBulder) UserID(v int) {
	b.Params["user_id"] = v
}

// ListID Friend list ID. If this parameter is not set, information about all online friends is returned.
func (b *FriendsGetOnlineBulder) ListID(v int) {
	b.Params["list_id"] = v
}

// OnlineMobile '1' — to return an additional 'online_mobile' field, '0' — (default),
func (b *FriendsGetOnlineBulder) OnlineMobile(v bool) {
	b.Params["online_mobile"] = v
}

// Order Sort order: 'random' — random order
func (b *FriendsGetOnlineBulder) Order(v string) {
	b.Params["order"] = v
}

// Count Number of friends to return.
func (b *FriendsGetOnlineBulder) Count(v int) {
	b.Params["count"] = v
}

// Offset Offset needed to return a specific subset of friends.
func (b *FriendsGetOnlineBulder) Offset(v int) {
	b.Params["offset"] = v
}

// FriendsGetRecentBulder builder
//
// Returns a list of user IDs of the current user's recently added friends.
//
// https://vk.com/dev/friends.getRecent
type FriendsGetRecentBulder struct {
	api.Params
}

// NewFriendsGetRecentBulder func
func NewFriendsGetRecentBulder() *FriendsGetRecentBulder {
	return &FriendsGetRecentBulder{api.Params{}}
}

// Count Number of recently added friends to return.
func (b *FriendsGetRecentBulder) Count(v int) {
	b.Params["count"] = v
}

// FriendsGetRequestsBulder builder
//
// Returns information about the current user's incoming and outgoing friend requests.
//
// https://vk.com/dev/friends.getRequests
type FriendsGetRequestsBulder struct {
	api.Params
}

// NewFriendsGetRequestsBulder func
func NewFriendsGetRequestsBulder() *FriendsGetRequestsBulder {
	return &FriendsGetRequestsBulder{api.Params{}}
}

// Offset Offset needed to return a specific subset of friend requests.
func (b *FriendsGetRequestsBulder) Offset(v int) {
	b.Params["offset"] = v
}

// Count Number of friend requests to return (default 100, maximum 1000).
func (b *FriendsGetRequestsBulder) Count(v int) {
	b.Params["count"] = v
}

// Extended '1' — to return response messages from users who have sent a friend request or, if 'suggested' is set to '1', to return a list of suggested friends
func (b *FriendsGetRequestsBulder) Extended(v bool) {
	b.Params["extended"] = v
}

// NeedMutual '1' — to return a list of mutual friends (up to 20), if any
func (b *FriendsGetRequestsBulder) NeedMutual(v bool) {
	b.Params["need_mutual"] = v
}

// Out '1' — to return outgoing requests, '0' — to return incoming requests (default)
func (b *FriendsGetRequestsBulder) Out(v bool) {
	b.Params["out"] = v
}

// Sort Sort order: '1' — by number of mutual friends, '0' — by date
func (b *FriendsGetRequestsBulder) Sort(v int) {
	b.Params["sort"] = v
}

// NeedViewed parameter
func (b *FriendsGetRequestsBulder) NeedViewed(v bool) {
	b.Params["need_viewed"] = v
}

// Suggested '1' — to return a list of suggested friends, '0' — to return friend requests (default)
func (b *FriendsGetRequestsBulder) Suggested(v bool) {
	b.Params["suggested"] = v
}

// Ref parameter
func (b *FriendsGetRequestsBulder) Ref(v string) {
	b.Params["ref"] = v
}

// Fields parameter
func (b *FriendsGetRequestsBulder) Fields(v []string) {
	b.Params["fields"] = v
}

// FriendsGetSuggestionsBulder builder
//
// Returns a list of profiles of users whom the current user may know.
//
// https://vk.com/dev/friends.getSuggestions
type FriendsGetSuggestionsBulder struct {
	api.Params
}

// NewFriendsGetSuggestionsBulder func
func NewFriendsGetSuggestionsBulder() *FriendsGetSuggestionsBulder {
	return &FriendsGetSuggestionsBulder{api.Params{}}
}

// Filter Types of potential friends to return: 'mutual' — users with many mutual friends , 'contacts' — users found with the [vk.com/dev/account.importContacts|account.importContacts] method , 'mutual_contacts' — users who imported the same contacts as the current user with the [vk.com/dev/account.importContacts|account.importContacts] method
func (b *FriendsGetSuggestionsBulder) Filter(v []string) {
	b.Params["filter"] = v
}

// Count Number of suggestions to return.
func (b *FriendsGetSuggestionsBulder) Count(v int) {
	b.Params["count"] = v
}

// Offset Offset needed to return a specific subset of suggestions.
func (b *FriendsGetSuggestionsBulder) Offset(v int) {
	b.Params["offset"] = v
}

// Fields Profile fields to return. Sample values: 'nickname', 'screen_name', 'sex', 'bdate' (birthdate), 'city', 'country', 'timezone', 'photo', 'photo_medium', 'photo_big', 'has_mobile', 'rate', 'contacts', 'education', 'online', 'counters'.
func (b *FriendsGetSuggestionsBulder) Fields(v []string) {
	b.Params["fields"] = v
}

// NameCase Case for declension of user name and surname: , 'nom' — nominative (default) , 'gen' — genitive , 'dat' — dative , 'acc' — accusative , 'ins' — instrumental , 'abl' — prepositional
func (b *FriendsGetSuggestionsBulder) NameCase(v string) {
	b.Params["name_case"] = v
}

// FriendsSearchBulder builder
//
// Returns a list of friends matching the search criteria.
//
// https://vk.com/dev/friends.search
type FriendsSearchBulder struct {
	api.Params
}

// NewFriendsSearchBulder func
func NewFriendsSearchBulder() *FriendsSearchBulder {
	return &FriendsSearchBulder{api.Params{}}
}

// UserID User ID.
func (b *FriendsSearchBulder) UserID(v int) {
	b.Params["user_id"] = v
}

// Q Search query string (e.g., 'Vasya Babich').
func (b *FriendsSearchBulder) Q(v string) {
	b.Params["q"] = v
}

// Fields Profile fields to return. Sample values: 'nickname', 'screen_name', 'sex', 'bdate' (birthdate), 'city', 'country', 'timezone', 'photo', 'photo_medium', 'photo_big', 'has_mobile', 'rate', 'contacts', 'education', 'online',
func (b *FriendsSearchBulder) Fields(v []string) {
	b.Params["fields"] = v
}

// NameCase Case for declension of user name and surname: 'nom' — nominative (default), 'gen' — genitive , 'dat' — dative, 'acc' — accusative , 'ins' — instrumental , 'abl' — prepositional
func (b *FriendsSearchBulder) NameCase(v string) {
	b.Params["name_case"] = v
}

// Offset Offset needed to return a specific subset of friends.
func (b *FriendsSearchBulder) Offset(v int) {
	b.Params["offset"] = v
}

// Count Number of friends to return.
func (b *FriendsSearchBulder) Count(v int) {
	b.Params["count"] = v
}
