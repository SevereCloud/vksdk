package params // import "github.com/SevereCloud/vksdk/api/params"

import (
	"github.com/SevereCloud/vksdk/api"
)

// AppsGetBulder builder
//
// Returns applications data.
//
// https://vk.com/dev/apps.get
type AppsGetBulder struct {
	api.Params
}

// NewAppsGetBulder func
func NewAppsGetBulder() *AppsGetBulder {
	return &AppsGetBulder{api.Params{}}
}

// AppID Application ID
func (b *AppsGetBulder) AppID(v int) {
	b.Params["app_id"] = v
}

// AppIDs List of application ID
func (b *AppsGetBulder) AppIDs(v []string) {
	b.Params["app_ids"] = v
}

// Platform platform. Possible values: *'ios' — iOS,, *'android' — Android,, *'winphone' — Windows Phone,, *'web' — приложения на vk.com. By default: 'web'.
func (b *AppsGetBulder) Platform(v string) {
	b.Params["platform"] = v
}

// Extended parameter
func (b *AppsGetBulder) Extended(v bool) {
	b.Params["extended"] = v
}

// ReturnFriends parameter
func (b *AppsGetBulder) ReturnFriends(v bool) {
	b.Params["return_friends"] = v
}

// Fields Profile fields to return. Sample values: 'nickname', 'screen_name', 'sex', 'bdate' (birthdate), 'city', 'country', 'timezone', 'photo', 'photo_medium', 'photo_big', 'has_mobile', 'contacts', 'education', 'online', 'counters', 'relation', 'last_seen', 'activity', 'can_write_private_message', 'can_see_all_posts', 'can_post', 'universities', (only if return_friends - 1)
func (b *AppsGetBulder) Fields(v []string) {
	b.Params["fields"] = v
}

// NameCase Case for declension of user name and surname: 'nom' — nominative (default),, 'gen' — genitive,, 'dat' — dative,, 'acc' — accusative,, 'ins' — instrumental,, 'abl' — prepositional. (only if 'return_friends' = '1')
func (b *AppsGetBulder) NameCase(v string) {
	b.Params["name_case"] = v
}

// AppsGetCatalogBulder builder
//
// Returns a list of applications (apps) available to users in the App Catalog.
//
// https://vk.com/dev/apps.getCatalog
type AppsGetCatalogBulder struct {
	api.Params
}

// NewAppsGetCatalogBulder func
func NewAppsGetCatalogBulder() *AppsGetCatalogBulder {
	return &AppsGetCatalogBulder{api.Params{}}
}

// Sort Sort order: 'popular_today' — popular for one day (default), 'visitors' — by visitors number , 'create_date' — by creation date, 'growth_rate' — by growth rate, 'popular_week' — popular for one week
func (b *AppsGetCatalogBulder) Sort(v string) {
	b.Params["sort"] = v
}

// Offset Offset required to return a specific subset of apps.
func (b *AppsGetCatalogBulder) Offset(v int) {
	b.Params["offset"] = v
}

// Count Number of apps to return.
func (b *AppsGetCatalogBulder) Count(v int) {
	b.Params["count"] = v
}

// Platform parameter
func (b *AppsGetCatalogBulder) Platform(v string) {
	b.Params["platform"] = v
}

// Extended '1' — to return additional fields 'screenshots', 'MAU', 'catalog_position', and 'international'. If set, 'count' must be less than or equal to '100'. '0' — not to return additional fields (default).
func (b *AppsGetCatalogBulder) Extended(v bool) {
	b.Params["extended"] = v
}

// ReturnFriends parameter
func (b *AppsGetCatalogBulder) ReturnFriends(v bool) {
	b.Params["return_friends"] = v
}

// Fields parameter
func (b *AppsGetCatalogBulder) Fields(v []string) {
	b.Params["fields"] = v
}

// NameCase parameter
func (b *AppsGetCatalogBulder) NameCase(v string) {
	b.Params["name_case"] = v
}

// Q Search query string.
func (b *AppsGetCatalogBulder) Q(v string) {
	b.Params["q"] = v
}

// GenreID parameter
func (b *AppsGetCatalogBulder) GenreID(v int) {
	b.Params["genre_id"] = v
}

// Filter 'installed' — to return list of installed apps (only for mobile platform).
func (b *AppsGetCatalogBulder) Filter(v string) {
	b.Params["filter"] = v
}

// AppsGetFriendsListBulder builder
//
// Creates friends list for requests and invites in current app.
//
// https://vk.com/dev/apps.getFriendsList
type AppsGetFriendsListBulder struct {
	api.Params
}

// NewAppsGetFriendsListBulder func
func NewAppsGetFriendsListBulder() *AppsGetFriendsListBulder {
	return &AppsGetFriendsListBulder{api.Params{}}
}

// Extended parameter
func (b *AppsGetFriendsListBulder) Extended(v bool) {
	b.Params["extended"] = v
}

// Count List size.
func (b *AppsGetFriendsListBulder) Count(v int) {
	b.Params["count"] = v
}

// Offset parameter
func (b *AppsGetFriendsListBulder) Offset(v int) {
	b.Params["offset"] = v
}

// Type List type. Possible values: * 'invite' — available for invites (don't play the game),, * 'request' — available for request (play the game). By default: 'invite'.
func (b *AppsGetFriendsListBulder) Type(v string) {
	b.Params["type"] = v
}

// Fields Additional profile fields, see [vk.com/dev/fields|description].
func (b *AppsGetFriendsListBulder) Fields(v []string) {
	b.Params["fields"] = v
}

// AppsGetLeaderboardBulder builder
//
// Returns players rating in the game.
//
// https://vk.com/dev/apps.getLeaderboard
type AppsGetLeaderboardBulder struct {
	api.Params
}

// NewAppsGetLeaderboardBulder func
func NewAppsGetLeaderboardBulder() *AppsGetLeaderboardBulder {
	return &AppsGetLeaderboardBulder{api.Params{}}
}

// Type Leaderboard type. Possible values: *'level' — by level,, *'points' — by mission points,, *'score' — by score ().
func (b *AppsGetLeaderboardBulder) Type(v string) {
	b.Params["type"] = v
}

// Global Rating type. Possible values: *'1' — global rating among all players,, *'0' — rating among user friends.
func (b *AppsGetLeaderboardBulder) Global(v bool) {
	b.Params["global"] = v
}

// Extended 1 — to return additional info about users
func (b *AppsGetLeaderboardBulder) Extended(v bool) {
	b.Params["extended"] = v
}

// AppsGetScopesBulder builder
//
// Returns scopes for auth
//
// https://vk.com/dev/apps.getScopes
type AppsGetScopesBulder struct {
	api.Params
}

// NewAppsGetScopesBulder func
func NewAppsGetScopesBulder() *AppsGetScopesBulder {
	return &AppsGetScopesBulder{api.Params{}}
}

// Type parameter
func (b *AppsGetScopesBulder) Type(v string) {
	b.Params["type"] = v
}

// AppsGetScoreBulder builder
//
// Returns user score in app
//
// https://vk.com/dev/apps.getScore
type AppsGetScoreBulder struct {
	api.Params
}

// NewAppsGetScoreBulder func
func NewAppsGetScoreBulder() *AppsGetScoreBulder {
	return &AppsGetScoreBulder{api.Params{}}
}

// UserID parameter
func (b *AppsGetScoreBulder) UserID(v int) {
	b.Params["user_id"] = v
}

// AppsSendRequestBulder builder
//
// Sends a request to another user in an app that uses VK authorization.
//
// https://vk.com/dev/apps.sendRequest
type AppsSendRequestBulder struct {
	api.Params
}

// NewAppsSendRequestBulder func
func NewAppsSendRequestBulder() *AppsSendRequestBulder {
	return &AppsSendRequestBulder{api.Params{}}
}

// UserID id of the user to send a request
func (b *AppsSendRequestBulder) UserID(v int) {
	b.Params["user_id"] = v
}

// Text request text
func (b *AppsSendRequestBulder) Text(v string) {
	b.Params["text"] = v
}

// Type request type. Values: 'invite' – if the request is sent to a user who does not have the app installed,, 'request' – if a user has already installed the app
func (b *AppsSendRequestBulder) Type(v string) {
	b.Params["type"] = v
}

// Name parameter
func (b *AppsSendRequestBulder) Name(v string) {
	b.Params["name"] = v
}

// Key special string key to be sent with the request
func (b *AppsSendRequestBulder) Key(v string) {
	b.Params["key"] = v
}

// Separate parameter
func (b *AppsSendRequestBulder) Separate(v bool) {
	b.Params["separate"] = v
}
