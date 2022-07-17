package params // import "github.com/SevereCloud/vksdk/v2/api/params"

import (
	"github.com/SevereCloud/vksdk/v2/api"
)

// AppsAddUsersToTestingGroupBuilder builder.
//
// https://vk.com/dev/apps.addUsersToTestingGroup
type AppsAddUsersToTestingGroupBuilder struct {
	api.Params
}

// NewAppsAddUsersToTestingGroupBuilder func.
func NewAppsAddUsersToTestingGroupBuilder() *AppsAddUsersToTestingGroupBuilder {
	return &AppsAddUsersToTestingGroupBuilder{api.Params{}}
}

// UserIDs users ID.
func (b *AppsAddUsersToTestingGroupBuilder) UserIDs(v []int) *AppsAddUsersToTestingGroupBuilder {
	b.Params["user_ids"] = v
	return b
}

// GroupID group ID.
func (b *AppsAddUsersToTestingGroupBuilder) GroupID(v int) *AppsAddUsersToTestingGroupBuilder {
	b.Params["group_id"] = v
	return b
}

// AppsGetBuilder builder.
//
// Returns applications data.
//
// https://vk.com/dev/apps.get
type AppsGetBuilder struct {
	api.Params
}

// NewAppsGetBuilder func.
func NewAppsGetBuilder() *AppsGetBuilder {
	return &AppsGetBuilder{api.Params{}}
}

// AppID application ID.
func (b *AppsGetBuilder) AppID(v int) *AppsGetBuilder {
	b.Params["app_id"] = v
	return b
}

// AppIDs list of application ID.
func (b *AppsGetBuilder) AppIDs(v []string) *AppsGetBuilder {
	b.Params["app_ids"] = v
	return b
}

// Platform platform. Possible values:
//
// * ios — iOS;
//
// * android — Android;
//
// * winphone — Windows Phone;
//
// * web — приложения на vk.com.
//
// By default: web.
func (b *AppsGetBuilder) Platform(v string) *AppsGetBuilder {
	b.Params["platform"] = v
	return b
}

// Extended parameter.
func (b *AppsGetBuilder) Extended(v bool) *AppsGetBuilder {
	b.Params["extended"] = v
	return b
}

// ReturnFriends parameter.
func (b *AppsGetBuilder) ReturnFriends(v bool) *AppsGetBuilder {
	b.Params["return_friends"] = v
	return b
}

// Fields profile fields to return. Sample values:
//
// * nickname;
//
// * screen_name;
//
// * sex;
//
// * bdate (birthdate);
//
// * city;
//
// * country;
//
// * timezone;
//
// * photo;
//
// * photo_medium;
//
// * photo_big;
//
// * has_mobile;
//
// * contacts;
//
// * education;
//
// * online;
//
// * counters;
//
// * relation;
//
// * last_seen;
//
// * activity;
//
// * can_write_private_message;
//
// * can_see_all_posts;
//
// * can_post;
//
// * universities, (only if return_friends - 1).
func (b *AppsGetBuilder) Fields(v []string) *AppsGetBuilder {
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
// * abl — prepositional. (only if 'return_friends' = '1').
func (b *AppsGetBuilder) NameCase(v string) *AppsGetBuilder {
	b.Params["name_case"] = v
	return b
}

// AppsGetCatalogBuilder builder.
//
// Returns a list of applications (apps) available to users in the App Catalog.
//
// https://vk.com/dev/apps.getCatalog
type AppsGetCatalogBuilder struct {
	api.Params
}

// NewAppsGetCatalogBuilder func.
func NewAppsGetCatalogBuilder() *AppsGetCatalogBuilder {
	return &AppsGetCatalogBuilder{api.Params{}}
}

// Sort order:
//
// * popular_today — popular for one day (default)
//
// * visitors — by visitors number
//
// * create_date — by creation date
//
// * growth_rate — by growth rate
//
// * popular_week — popular for one week.
func (b *AppsGetCatalogBuilder) Sort(v string) *AppsGetCatalogBuilder {
	b.Params["sort"] = v
	return b
}

// Offset required to return a specific subset of apps.
func (b *AppsGetCatalogBuilder) Offset(v int) *AppsGetCatalogBuilder {
	b.Params["offset"] = v
	return b
}

// Count number of apps to return.
func (b *AppsGetCatalogBuilder) Count(v int) *AppsGetCatalogBuilder {
	b.Params["count"] = v
	return b
}

// Platform parameter.
func (b *AppsGetCatalogBuilder) Platform(v string) *AppsGetCatalogBuilder {
	b.Params["platform"] = v
	return b
}

// Extended parameter.
//
// * 1 — to return additional fields 'screenshots', 'MAU', 'catalog_position', and 'international'.
// If set, 'count' must be less than or equal to '100'.
//
// * 0 — not to return additional fields (default).
func (b *AppsGetCatalogBuilder) Extended(v bool) *AppsGetCatalogBuilder {
	b.Params["extended"] = v
	return b
}

// ReturnFriends parameter.
func (b *AppsGetCatalogBuilder) ReturnFriends(v bool) *AppsGetCatalogBuilder {
	b.Params["return_friends"] = v
	return b
}

// Fields parameter.
func (b *AppsGetCatalogBuilder) Fields(v []string) *AppsGetCatalogBuilder {
	b.Params["fields"] = v
	return b
}

// NameCase parameter.
func (b *AppsGetCatalogBuilder) NameCase(v string) *AppsGetCatalogBuilder {
	b.Params["name_case"] = v
	return b
}

// Q search query string.
func (b *AppsGetCatalogBuilder) Q(v string) *AppsGetCatalogBuilder {
	b.Params["q"] = v
	return b
}

// GenreID parameter.
func (b *AppsGetCatalogBuilder) GenreID(v int) *AppsGetCatalogBuilder {
	b.Params["genre_id"] = v
	return b
}

// Filter 'installed' — to return list of installed apps (only for mobile platform).
func (b *AppsGetCatalogBuilder) Filter(v string) *AppsGetCatalogBuilder {
	b.Params["filter"] = v
	return b
}

// AppsGetFriendsListBuilder builder.
//
// Creates friends list for requests and invites in current app.
//
// https://vk.com/dev/apps.getFriendsList
type AppsGetFriendsListBuilder struct {
	api.Params
}

// NewAppsGetFriendsListBuilder func.
func NewAppsGetFriendsListBuilder() *AppsGetFriendsListBuilder {
	return &AppsGetFriendsListBuilder{api.Params{}}
}

// Extended parameter.
func (b *AppsGetFriendsListBuilder) Extended(v bool) *AppsGetFriendsListBuilder {
	b.Params["extended"] = v
	return b
}

// Count - list size.
func (b *AppsGetFriendsListBuilder) Count(v int) *AppsGetFriendsListBuilder {
	b.Params["count"] = v
	return b
}

// Offset parameter.
func (b *AppsGetFriendsListBuilder) Offset(v int) *AppsGetFriendsListBuilder {
	b.Params["offset"] = v
	return b
}

// Type list. Possible values:
//
// * invite — available for invites (don't play the game);
//
// * request — available for request (play the game).
//
// By default: invite.
func (b *AppsGetFriendsListBuilder) Type(v string) *AppsGetFriendsListBuilder {
	b.Params["type"] = v
	return b
}

// Fields additional profile fields, see [vk.com/dev/fields|description].
func (b *AppsGetFriendsListBuilder) Fields(v []string) *AppsGetFriendsListBuilder {
	b.Params["fields"] = v
	return b
}

// AppsGetLeaderboardBuilder builder.
//
// Returns players rating in the game.
//
// https://vk.com/dev/apps.getLeaderboard
type AppsGetLeaderboardBuilder struct {
	api.Params
}

// NewAppsGetLeaderboardBuilder func.
func NewAppsGetLeaderboardBuilder() *AppsGetLeaderboardBuilder {
	return &AppsGetLeaderboardBuilder{api.Params{}}
}

// Type leaderboard. Possible values:
//
// * level — by level;
//
// * points — by mission points;
//
// * score — by score ().
func (b *AppsGetLeaderboardBuilder) Type(v string) *AppsGetLeaderboardBuilder {
	b.Params["type"] = v
	return b
}

// Global rating type. Possible values:
//
// * 1 — global rating among all players;
//
// * 0 — rating among user friends.
func (b *AppsGetLeaderboardBuilder) Global(v bool) *AppsGetLeaderboardBuilder {
	b.Params["global"] = v
	return b
}

// Extended 1 — to return additional info about users.
func (b *AppsGetLeaderboardBuilder) Extended(v bool) *AppsGetLeaderboardBuilder {
	b.Params["extended"] = v
	return b
}

// AppsGetScopesBuilder builder.
//
// Returns scopes for auth.
//
// https://vk.com/dev/apps.getScopes
type AppsGetScopesBuilder struct {
	api.Params
}

// NewAppsGetScopesBuilder func.
func NewAppsGetScopesBuilder() *AppsGetScopesBuilder {
	return &AppsGetScopesBuilder{api.Params{}}
}

// Type parameter.
func (b *AppsGetScopesBuilder) Type(v string) *AppsGetScopesBuilder {
	b.Params["type"] = v
	return b
}

// AppsGetScoreBuilder builder.
//
// Returns user score in app.
//
// https://vk.com/dev/apps.getScore
type AppsGetScoreBuilder struct {
	api.Params
}

// NewAppsGetScoreBuilder func.
func NewAppsGetScoreBuilder() *AppsGetScoreBuilder {
	return &AppsGetScoreBuilder{api.Params{}}
}

// UserID parameter.
func (b *AppsGetScoreBuilder) UserID(v int) *AppsGetScoreBuilder {
	b.Params["user_id"] = v
	return b
}

// AppsGetTestingGroupsBuilder builder.
//
// https://vk.com/dev/apps.getTestingGroups
type AppsGetTestingGroupsBuilder struct {
	api.Params
}

// NewAppsGetTestingGroupsBuilder func.
func NewAppsGetTestingGroupsBuilder() *AppsGetTestingGroupsBuilder {
	return &AppsGetTestingGroupsBuilder{api.Params{}}
}

// GroupID group ID.
func (b *AppsGetTestingGroupsBuilder) GroupID(v int) *AppsGetTestingGroupsBuilder {
	b.Params["group_id"] = v
	return b
}

// AppsRemoveUsersFromTestingGroupsBuilder builder.
//
// https://vk.com/dev/apps.removeUsersFromTestingGroups
type AppsRemoveUsersFromTestingGroupsBuilder struct {
	api.Params
}

// NewAppsRemoveUsersFromTestingGroupsBuilder func.
func NewAppsRemoveUsersFromTestingGroupsBuilder() *AppsRemoveUsersFromTestingGroupsBuilder {
	return &AppsRemoveUsersFromTestingGroupsBuilder{api.Params{}}
}

// UserIDs users ID.
func (b *AppsRemoveUsersFromTestingGroupsBuilder) UserIDs(v []int) *AppsRemoveUsersFromTestingGroupsBuilder {
	b.Params["user_ids"] = v
	return b
}

// AppsRemoveTestingGroupBuilder builder.
//
// https://vk.com/dev/apps.removeTestingGroup
type AppsRemoveTestingGroupBuilder struct {
	api.Params
}

// NewAppsRemoveTestingGroupBuilder func.
func NewAppsRemoveTestingGroupBuilder() *AppsRemoveTestingGroupBuilder {
	return &AppsRemoveTestingGroupBuilder{api.Params{}}
}

// GroupID group ID.
func (b *AppsRemoveTestingGroupBuilder) GroupID(v int) *AppsRemoveTestingGroupBuilder {
	b.Params["group_id"] = v
	return b
}

// AppsSendRequestBuilder builder.
//
// Sends a request to another user in an app that uses VK authorization.
//
// https://vk.com/dev/apps.sendRequest
type AppsSendRequestBuilder struct {
	api.Params
}

// NewAppsSendRequestBuilder func.
func NewAppsSendRequestBuilder() *AppsSendRequestBuilder {
	return &AppsSendRequestBuilder{api.Params{}}
}

// UserID id of the user to send a request.
func (b *AppsSendRequestBuilder) UserID(v int) *AppsSendRequestBuilder {
	b.Params["user_id"] = v
	return b
}

// Text request text.
func (b *AppsSendRequestBuilder) Text(v string) *AppsSendRequestBuilder {
	b.Params["text"] = v
	return b
}

// Type request type. Values:
//
// * invite – if the request is sent to a user who does not have the app installed
//
// * request – if a user has already installed the app.
func (b *AppsSendRequestBuilder) Type(v string) *AppsSendRequestBuilder {
	b.Params["type"] = v
	return b
}

// Name parameter.
func (b *AppsSendRequestBuilder) Name(v string) *AppsSendRequestBuilder {
	b.Params["name"] = v
	return b
}

// Key special string key to be sent with the request.
func (b *AppsSendRequestBuilder) Key(v string) *AppsSendRequestBuilder {
	b.Params["key"] = v
	return b
}

// Separate parameter.
func (b *AppsSendRequestBuilder) Separate(v bool) *AppsSendRequestBuilder {
	b.Params["separate"] = v
	return b
}

// AppsUpdateMetaForTestingGroupBuilder builder.
//
// https://vk.com/dev/apps.UpdateMetaForTestingGroup
type AppsUpdateMetaForTestingGroupBuilder struct {
	api.Params
}

// NewAppsUpdateMetaForTestingGroupBuilder func.
func NewAppsUpdateMetaForTestingGroupBuilder() *AppsUpdateMetaForTestingGroupBuilder {
	return &AppsUpdateMetaForTestingGroupBuilder{api.Params{}}
}

// GroupID parameter.
func (b *AppsUpdateMetaForTestingGroupBuilder) GroupID(v int) *AppsUpdateMetaForTestingGroupBuilder {
	b.Params["group_id"] = v
	return b
}

// Webview parameter.
func (b *AppsUpdateMetaForTestingGroupBuilder) Webview(v string) *AppsUpdateMetaForTestingGroupBuilder {
	b.Params["webview"] = v
	return b
}

// Name parameter.
func (b *AppsUpdateMetaForTestingGroupBuilder) Name(v string) *AppsUpdateMetaForTestingGroupBuilder {
	b.Params["name"] = v
	return b
}

// Platforms parameter.
func (b *AppsUpdateMetaForTestingGroupBuilder) Platforms(v []string) *AppsUpdateMetaForTestingGroupBuilder {
	b.Params["platforms"] = v
	return b
}

// UserIDs users ID.
func (b *AppsUpdateMetaForTestingGroupBuilder) UserIDs(v []int) *AppsUpdateMetaForTestingGroupBuilder {
	b.Params["user_ids"] = v
	return b
}
