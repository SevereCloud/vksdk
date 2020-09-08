package params // import "github.com/SevereCloud/vksdk/v2/api/params"

import (
	"github.com/SevereCloud/vksdk/v2/api"
)

// FriendsAddBuilder builder.
//
// Approves or creates a friend request.
//
// https://vk.com/dev/friends.add
type FriendsAddBuilder struct {
	api.Params
}

// NewFriendsAddBuilder func.
func NewFriendsAddBuilder() *FriendsAddBuilder {
	return &FriendsAddBuilder{api.Params{}}
}

// UserID ID of the user whose friend request will be approved or to whom a friend request will be sent.
func (b *FriendsAddBuilder) UserID(v int) *FriendsAddBuilder {
	b.Params["user_id"] = v
	return b
}

// Text of the message (up to 500 characters) for the friend request, if any.
func (b *FriendsAddBuilder) Text(v string) *FriendsAddBuilder {
	b.Params["text"] = v
	return b
}

// Follow '1' to pass an incoming request to followers list.
func (b *FriendsAddBuilder) Follow(v bool) *FriendsAddBuilder {
	b.Params["follow"] = v
	return b
}

// FriendsAddListBuilder builder.
//
// Creates a new friend list for the current user.
//
// https://vk.com/dev/friends.addList
type FriendsAddListBuilder struct {
	api.Params
}

// NewFriendsAddListBuilder func.
func NewFriendsAddListBuilder() *FriendsAddListBuilder {
	return &FriendsAddListBuilder{api.Params{}}
}

// Name of the friend list.
func (b *FriendsAddListBuilder) Name(v string) *FriendsAddListBuilder {
	b.Params["name"] = v
	return b
}

// UserIDs IDs of users to be added to the friend list.
func (b *FriendsAddListBuilder) UserIDs(v []int) *FriendsAddListBuilder {
	b.Params["user_ids"] = v
	return b
}

// FriendsAreFriendsBuilder builder.
//
// Checks the current user's friendship status with other specified users.
//
// https://vk.com/dev/friends.areFriends
type FriendsAreFriendsBuilder struct {
	api.Params
}

// NewFriendsAreFriendsBuilder func.
func NewFriendsAreFriendsBuilder() *FriendsAreFriendsBuilder {
	return &FriendsAreFriendsBuilder{api.Params{}}
}

// UserIDs IDs of the users whose friendship status to check.
func (b *FriendsAreFriendsBuilder) UserIDs(v []int) *FriendsAreFriendsBuilder {
	b.Params["user_ids"] = v
	return b
}

// NeedSign parameter.
//
// * 1 — to return 'sign' field. 'sign' is md5("{id}_{user_id}_{friends_status}_{application_secret}"),
// where id is current user ID. This field allows to check that data has not been modified by the client.
//
// By default: '0'.
func (b *FriendsAreFriendsBuilder) NeedSign(v bool) *FriendsAreFriendsBuilder {
	b.Params["need_sign"] = v
	return b
}

// FriendsDeleteBuilder builder.
//
// Declines a friend request or deletes a user from the current user's friend list.
//
// https://vk.com/dev/friends.delete
type FriendsDeleteBuilder struct {
	api.Params
}

// NewFriendsDeleteBuilder func.
func NewFriendsDeleteBuilder() *FriendsDeleteBuilder {
	return &FriendsDeleteBuilder{api.Params{}}
}

// UserID ID of the user whose friend request is to be declined or who is to be deleted from the current user's
// friend list.
func (b *FriendsDeleteBuilder) UserID(v int) *FriendsDeleteBuilder {
	b.Params["user_id"] = v
	return b
}

// FriendsDeleteListBuilder builder.
//
// Deletes a friend list of the current user.
//
// https://vk.com/dev/friends.deleteList
type FriendsDeleteListBuilder struct {
	api.Params
}

// NewFriendsDeleteListBuilder func.
func NewFriendsDeleteListBuilder() *FriendsDeleteListBuilder {
	return &FriendsDeleteListBuilder{api.Params{}}
}

// ListID ID of the friend list to delete.
func (b *FriendsDeleteListBuilder) ListID(v int) *FriendsDeleteListBuilder {
	b.Params["list_id"] = v
	return b
}

// FriendsEditBuilder builder.
//
// Edits the friend lists of the selected user.
//
// https://vk.com/dev/friends.edit
type FriendsEditBuilder struct {
	api.Params
}

// NewFriendsEditBuilder func.
func NewFriendsEditBuilder() *FriendsEditBuilder {
	return &FriendsEditBuilder{api.Params{}}
}

// UserID ID of the user whose friend list is to be edited.
func (b *FriendsEditBuilder) UserID(v int) *FriendsEditBuilder {
	b.Params["user_id"] = v
	return b
}

// ListIDs IDs of the friend lists to which to add the user.
func (b *FriendsEditBuilder) ListIDs(v []int) *FriendsEditBuilder {
	b.Params["list_ids"] = v
	return b
}

// FriendsEditListBuilder builder.
//
// Edits a friend list of the current user.
//
// https://vk.com/dev/friends.editList
type FriendsEditListBuilder struct {
	api.Params
}

// NewFriendsEditListBuilder func.
func NewFriendsEditListBuilder() *FriendsEditListBuilder {
	return &FriendsEditListBuilder{api.Params{}}
}

// Name of the friend list.
func (b *FriendsEditListBuilder) Name(v string) *FriendsEditListBuilder {
	b.Params["name"] = v
	return b
}

// ListID friend list ID.
func (b *FriendsEditListBuilder) ListID(v int) *FriendsEditListBuilder {
	b.Params["list_id"] = v
	return b
}

// UserIDs IDs of users in the friend list.
func (b *FriendsEditListBuilder) UserIDs(v []int) *FriendsEditListBuilder {
	b.Params["user_ids"] = v
	return b
}

// AddUserIDs (Applies if 'user_ids' parameter is not set.), User IDs to add to the friend list.
func (b *FriendsEditListBuilder) AddUserIDs(v []int) *FriendsEditListBuilder {
	b.Params["add_user_ids"] = v
	return b
}

// DeleteUserIDs (Applies if 'user_ids' parameter is not set.), User IDs to delete from the friend list.
func (b *FriendsEditListBuilder) DeleteUserIDs(v []int) *FriendsEditListBuilder {
	b.Params["delete_user_ids"] = v
	return b
}

// FriendsGetBuilder builder.
//
// Returns a list of user IDs or detailed information about a user's friends.
//
// https://vk.com/dev/friends.get
type FriendsGetBuilder struct {
	api.Params
}

// NewFriendsGetBuilder func.
func NewFriendsGetBuilder() *FriendsGetBuilder {
	return &FriendsGetBuilder{api.Params{}}
}

// UserID parameter. By default, the current user ID.
func (b *FriendsGetBuilder) UserID(v int) *FriendsGetBuilder {
	b.Params["user_id"] = v
	return b
}

// Order sort order:
//
// * name — by name (enabled only if the 'fields' parameter is used);
//
// * hints — by rating, similar to how friends are sorted in My friends section.
// This parameter is available only for [vk.com/dev/standalone|desktop applications].
func (b *FriendsGetBuilder) Order(v string) *FriendsGetBuilder {
	b.Params["order"] = v
	return b
}

// ListID ID of the friend list returned by the [vk.com/dev/friends.getLists|friends.getLists] method to be used
// as the source. This parameter is taken into account only when the uid parameter is set to the current user ID.
// This parameter is available only for [vk.com/dev/standalone|desktop applications].
func (b *FriendsGetBuilder) ListID(v int) *FriendsGetBuilder {
	b.Params["list_id"] = v
	return b
}

// Count number of friends to return.
func (b *FriendsGetBuilder) Count(v int) *FriendsGetBuilder {
	b.Params["count"] = v
	return b
}

// Offset needed to return a specific subset of friends.
func (b *FriendsGetBuilder) Offset(v int) *FriendsGetBuilder {
	b.Params["offset"] = v
	return b
}

// Fields profile fields to return. Sample values: 'uid', 'first_name', 'last_name', 'nickname', 'sex',
// 'bdate' (birthdate), 'city', 'country', 'timezone', 'photo', 'photo_medium', 'photo_big', 'domain',
// 'has_mobile', 'rate', 'contacts', 'education'.
func (b *FriendsGetBuilder) Fields(v []string) *FriendsGetBuilder {
	b.Params["fields"] = v
	return b
}

// NameCase case for declension of user name and surname:
//
// * nom — nominative (default);
//
// * gen — genitive;
//
// * dat — dative;
//
// * acc — accusative;
//
// * ins — instrumental;
//
// * abl — prepositional.
func (b *FriendsGetBuilder) NameCase(v string) *FriendsGetBuilder {
	b.Params["name_case"] = v
	return b
}

// Ref parameter.
func (b *FriendsGetBuilder) Ref(v string) *FriendsGetBuilder {
	b.Params["ref"] = v
	return b
}

// FriendsGetByPhonesBuilder builder.
//
// Returns a list of the current user's friends whose phone numbers, validated or specified in a profile, are in a
// given list.
//
// https://vk.com/dev/friends.getByPhones
type FriendsGetByPhonesBuilder struct {
	api.Params
}

// NewFriendsGetByPhonesBuilder func.
func NewFriendsGetByPhonesBuilder() *FriendsGetByPhonesBuilder {
	return &FriendsGetByPhonesBuilder{api.Params{}}
}

// Phones list of phone numbers in MSISDN format (maximum 1000). Example: "+79219876543,+79111234567".
func (b *FriendsGetByPhonesBuilder) Phones(v []string) *FriendsGetByPhonesBuilder {
	b.Params["phones"] = v
	return b
}

// Fields profile fields to return. Sample values: 'nickname', 'screen_name', 'sex', 'bdate' (birthdate), 'city',
// 'country', 'timezone', 'photo', 'photo_medium', 'photo_big', 'has_mobile',
// 'rate', 'contacts', 'education', 'online, counters'.
func (b *FriendsGetByPhonesBuilder) Fields(v []string) *FriendsGetByPhonesBuilder {
	b.Params["fields"] = v
	return b
}

// FriendsGetListsBuilder builder.
//
// Returns a list of the user's friend lists.
//
// https://vk.com/dev/friends.getLists
type FriendsGetListsBuilder struct {
	api.Params
}

// NewFriendsGetListsBuilder func.
func NewFriendsGetListsBuilder() *FriendsGetListsBuilder {
	return &FriendsGetListsBuilder{api.Params{}}
}

// UserID parameter.
func (b *FriendsGetListsBuilder) UserID(v int) *FriendsGetListsBuilder {
	b.Params["user_id"] = v
	return b
}

// ReturnSystem parameter.
//
// * 1 — to return system friend lists. By default: '0'.
func (b *FriendsGetListsBuilder) ReturnSystem(v bool) *FriendsGetListsBuilder {
	b.Params["return_system"] = v
	return b
}

// FriendsGetMutualBuilder builder.
//
// Returns a list of user IDs of the mutual friends of two users.
//
// https://vk.com/dev/friends.getMutual
type FriendsGetMutualBuilder struct {
	api.Params
}

// NewFriendsGetMutualBuilder func.
func NewFriendsGetMutualBuilder() *FriendsGetMutualBuilder {
	return &FriendsGetMutualBuilder{api.Params{}}
}

// SourceUID ID of the user whose friends will be checked against the friends of the user specified in 'target_uid'.
func (b *FriendsGetMutualBuilder) SourceUID(v int) *FriendsGetMutualBuilder {
	b.Params["source_uid"] = v
	return b
}

// TargetUID ID of the user whose friends will be checked against the friends of the user specified in 'source_uid'.
func (b *FriendsGetMutualBuilder) TargetUID(v int) *FriendsGetMutualBuilder {
	b.Params["target_uid"] = v
	return b
}

// TargetUids IDs of the users whose friends will be checked against the friends of the user specified in 'source_uid'.
func (b *FriendsGetMutualBuilder) TargetUids(v []int) *FriendsGetMutualBuilder {
	b.Params["target_uids"] = v
	return b
}

// Order sort order: 'random' — random order.
func (b *FriendsGetMutualBuilder) Order(v string) *FriendsGetMutualBuilder {
	b.Params["order"] = v
	return b
}

// Count number of mutual friends to return.
func (b *FriendsGetMutualBuilder) Count(v int) *FriendsGetMutualBuilder {
	b.Params["count"] = v
	return b
}

// Offset needed to return a specific subset of mutual friends.
func (b *FriendsGetMutualBuilder) Offset(v int) *FriendsGetMutualBuilder {
	b.Params["offset"] = v
	return b
}

// FriendsGetOnlineBuilder builder.
//
// Returns a list of user IDs of a user's friends who are online.
//
// https://vk.com/dev/friends.getOnline
type FriendsGetOnlineBuilder struct {
	api.Params
}

// NewFriendsGetOnlineBuilder func.
func NewFriendsGetOnlineBuilder() *FriendsGetOnlineBuilder {
	return &FriendsGetOnlineBuilder{api.Params{}}
}

// UserID parameter.
func (b *FriendsGetOnlineBuilder) UserID(v int) *FriendsGetOnlineBuilder {
	b.Params["user_id"] = v
	return b
}

// ListID friend list ID. If this parameter is not set, information about all online friends is returned.
func (b *FriendsGetOnlineBuilder) ListID(v int) *FriendsGetOnlineBuilder {
	b.Params["list_id"] = v
	return b
}

// OnlineMobile parameter.
//
// * 1 — to return an additional 'online_mobile' field, '0' — (default).
func (b *FriendsGetOnlineBuilder) OnlineMobile(v bool) *FriendsGetOnlineBuilder {
	b.Params["online_mobile"] = v
	return b
}

// Order sort order: 'random' — random order.
func (b *FriendsGetOnlineBuilder) Order(v string) *FriendsGetOnlineBuilder {
	b.Params["order"] = v
	return b
}

// Count number of friends to return.
func (b *FriendsGetOnlineBuilder) Count(v int) *FriendsGetOnlineBuilder {
	b.Params["count"] = v
	return b
}

// Offset needed to return a specific subset of friends.
func (b *FriendsGetOnlineBuilder) Offset(v int) *FriendsGetOnlineBuilder {
	b.Params["offset"] = v
	return b
}

// FriendsGetRecentBuilder builder.
//
// Returns a list of user IDs of the current user's recently added friends.
//
// https://vk.com/dev/friends.getRecent
type FriendsGetRecentBuilder struct {
	api.Params
}

// NewFriendsGetRecentBuilder func.
func NewFriendsGetRecentBuilder() *FriendsGetRecentBuilder {
	return &FriendsGetRecentBuilder{api.Params{}}
}

// Count number of recently added friends to return.
func (b *FriendsGetRecentBuilder) Count(v int) *FriendsGetRecentBuilder {
	b.Params["count"] = v
	return b
}

// FriendsGetRequestsBuilder builder.
//
// Returns information about the current user's incoming and outgoing friend requests.
//
// https://vk.com/dev/friends.getRequests
type FriendsGetRequestsBuilder struct {
	api.Params
}

// NewFriendsGetRequestsBuilder func.
func NewFriendsGetRequestsBuilder() *FriendsGetRequestsBuilder {
	return &FriendsGetRequestsBuilder{api.Params{}}
}

// Offset needed to return a specific subset of friend requests.
func (b *FriendsGetRequestsBuilder) Offset(v int) *FriendsGetRequestsBuilder {
	b.Params["offset"] = v
	return b
}

// Count number of friend requests to return (default 100, maximum 1000).
func (b *FriendsGetRequestsBuilder) Count(v int) *FriendsGetRequestsBuilder {
	b.Params["count"] = v
	return b
}

// Extended parameter.
//
// * 1 — to return response messages from users who have sent a friend request or,
// if 'suggested' is set to '1', to return a list of suggested friends.
func (b *FriendsGetRequestsBuilder) Extended(v bool) *FriendsGetRequestsBuilder {
	b.Params["extended"] = v
	return b
}

// NeedMutual parameter.
//
// * 1 — to return a list of mutual friends (up to 20), if any.
func (b *FriendsGetRequestsBuilder) NeedMutual(v bool) *FriendsGetRequestsBuilder {
	b.Params["need_mutual"] = v
	return b
}

// Out parameter.
//
// * 1 — to return outgoing requests, '0' — to return incoming requests (default).
func (b *FriendsGetRequestsBuilder) Out(v bool) *FriendsGetRequestsBuilder {
	b.Params["out"] = v
	return b
}

// Sort order:
//
// * 1 — by number of mutual friends, '0' — by date.
func (b *FriendsGetRequestsBuilder) Sort(v int) *FriendsGetRequestsBuilder {
	b.Params["sort"] = v
	return b
}

// NeedViewed parameter.
func (b *FriendsGetRequestsBuilder) NeedViewed(v bool) *FriendsGetRequestsBuilder {
	b.Params["need_viewed"] = v
	return b
}

// Suggested parameter.
//
// * 1 — to return a list of suggested friends, '0' — to return friend requests (default).
func (b *FriendsGetRequestsBuilder) Suggested(v bool) *FriendsGetRequestsBuilder {
	b.Params["suggested"] = v
	return b
}

// Ref parameter.
func (b *FriendsGetRequestsBuilder) Ref(v string) *FriendsGetRequestsBuilder {
	b.Params["ref"] = v
	return b
}

// Fields parameter.
func (b *FriendsGetRequestsBuilder) Fields(v []string) *FriendsGetRequestsBuilder {
	b.Params["fields"] = v
	return b
}

// FriendsGetSuggestionsBuilder builder.
//
// Returns a list of profiles of users whom the current user may know.
//
// https://vk.com/dev/friends.getSuggestions
type FriendsGetSuggestionsBuilder struct {
	api.Params
}

// NewFriendsGetSuggestionsBuilder func.
func NewFriendsGetSuggestionsBuilder() *FriendsGetSuggestionsBuilder {
	return &FriendsGetSuggestionsBuilder{api.Params{}}
}

// Filter Types of potential friends to return:
//
// * mutual — users with many mutual friends;
//
// * contacts — users found with the [vk.com/dev/account.importContacts|account.importContacts] method;
//
// * mutual_contacts — users who imported the same contacts as the current user with the
// [vk.com/dev/account.importContacts|account.importContacts] method.
func (b *FriendsGetSuggestionsBuilder) Filter(v []string) *FriendsGetSuggestionsBuilder {
	b.Params["filter"] = v
	return b
}

// Count number of suggestions to return.
func (b *FriendsGetSuggestionsBuilder) Count(v int) *FriendsGetSuggestionsBuilder {
	b.Params["count"] = v
	return b
}

// Offset needed to return a specific subset of suggestions.
func (b *FriendsGetSuggestionsBuilder) Offset(v int) *FriendsGetSuggestionsBuilder {
	b.Params["offset"] = v
	return b
}

// Fields profile fields to return. Sample values: 'nickname', 'screen_name', 'sex', 'bdate' (birthdate), 'city',
// 'country', 'timezone', 'photo', 'photo_medium', 'photo_big', 'has_mobile',
// 'rate', 'contacts', 'education', 'online', 'counters'.
func (b *FriendsGetSuggestionsBuilder) Fields(v []string) *FriendsGetSuggestionsBuilder {
	b.Params["fields"] = v
	return b
}

// NameCase case for declension of user name and surname:
//
// * nom — nominative (default);
//
// * gen — genitive;
//
// * dat — dative;
//
// * acc — accusative;
//
// * ins — instrumental;
//
// * abl — prepositional.
func (b *FriendsGetSuggestionsBuilder) NameCase(v string) *FriendsGetSuggestionsBuilder {
	b.Params["name_case"] = v
	return b
}

// FriendsSearchBuilder builder.
//
// Returns a list of friends matching the search criteria.
//
// https://vk.com/dev/friends.search
type FriendsSearchBuilder struct {
	api.Params
}

// NewFriendsSearchBuilder func.
func NewFriendsSearchBuilder() *FriendsSearchBuilder {
	return &FriendsSearchBuilder{api.Params{}}
}

// UserID parameter.
func (b *FriendsSearchBuilder) UserID(v int) *FriendsSearchBuilder {
	b.Params["user_id"] = v
	return b
}

// Q search query string (e.g., 'Vasya Babich').
func (b *FriendsSearchBuilder) Q(v string) *FriendsSearchBuilder {
	b.Params["q"] = v
	return b
}

// Fields profile fields to return. Sample values: 'nickname', 'screen_name', 'sex', 'bdate' (birthdate), 'city',
// 'country', 'timezone', 'photo', 'photo_medium', 'photo_big', 'has_mobile',
// 'rate', 'contacts', 'education', 'online'.
func (b *FriendsSearchBuilder) Fields(v []string) *FriendsSearchBuilder {
	b.Params["fields"] = v
	return b
}

// NameCase case for declension of user name and surname:
//
// * nom — nominative (default),
//
// * gen — genitive;
//
// * dat — dative;
//
// * acc — accusative;
//
// * ins — instrumental;
//
// * abl — prepositional.
func (b *FriendsSearchBuilder) NameCase(v string) *FriendsSearchBuilder {
	b.Params["name_case"] = v
	return b
}

// Offset needed to return a specific subset of friends.
func (b *FriendsSearchBuilder) Offset(v int) *FriendsSearchBuilder {
	b.Params["offset"] = v
	return b
}

// Count number of friends to return.
func (b *FriendsSearchBuilder) Count(v int) *FriendsSearchBuilder {
	b.Params["count"] = v
	return b
}
