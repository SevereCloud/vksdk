package params // import "github.com/SevereCloud/vksdk/api/params"

import (
	"github.com/SevereCloud/vksdk/api"
)

// AppsGetBuilder builder
//
// Returns applications data.
//
// https://vk.com/dev/apps.get
type AppsGetBuilder struct {
	api.Params
}

// NewAppsGetBuilder func
func NewAppsGetBuilder() *AppsGetBuilder {
	return &AppsGetBuilder{api.Params{}}
}

// AppID Application ID
func (b *AppsGetBuilder) AppID(v int) {
	b.Params["app_id"] = v
}

// AppIDs List of application ID
func (b *AppsGetBuilder) AppIDs(v []string) {
	b.Params["app_ids"] = v
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
func (b *AppsGetBuilder) Platform(v string) {
	b.Params["platform"] = v
}

// Extended parameter
func (b *AppsGetBuilder) Extended(v bool) {
	b.Params["extended"] = v
}

// ReturnFriends parameter
func (b *AppsGetBuilder) ReturnFriends(v bool) {
	b.Params["return_friends"] = v
}

// Fields Profile fields to return. Sample values:
//
// * nickname
//
// * screen_name
//
// * sex
//
// * bdate (birthdate)
//
// * city
//
// * country
//
// * timezone
//
// * photo
//
// * photo_medium
//
// * photo_big
//
// * has_mobile
//
// * contacts
//
// * education
//
// * online
//
// * counters
//
// * relation
//
// * last_seen
//
// * activity
//
// * can_write_private_message
//
// * can_see_all_posts
//
// * can_post
//
// * universities, (only if return_friends - 1)
func (b *AppsGetBuilder) Fields(v []string) {
	b.Params["fields"] = v
}

// NameCase Case for declension of user name and surname:
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
// * abl — prepositional. (only if 'return_friends' = '1')
func (b *AppsGetBuilder) NameCase(v string) {
	b.Params["name_case"] = v
}

// AppsGetCatalogBuilder builder
//
// Returns a list of applications (apps) available to users in the App Catalog.
//
// https://vk.com/dev/apps.getCatalog
type AppsGetCatalogBuilder struct {
	api.Params
}

// NewAppsGetCatalogBuilder func
func NewAppsGetCatalogBuilder() *AppsGetCatalogBuilder {
	return &AppsGetCatalogBuilder{api.Params{}}
}

// Sort Sort order:
//
// * popular_today — popular for one day (default)
//
// * visitors — by visitors number
//
// * create_date — by creation date
//
// * growth_rate — by growth rate
//
// * popular_week — popular for one week
func (b *AppsGetCatalogBuilder) Sort(v string) {
	b.Params["sort"] = v
}

// Offset Offset required to return a specific subset of apps.
func (b *AppsGetCatalogBuilder) Offset(v int) {
	b.Params["offset"] = v
}

// Count Number of apps to return.
func (b *AppsGetCatalogBuilder) Count(v int) {
	b.Params["count"] = v
}

// Platform parameter
func (b *AppsGetCatalogBuilder) Platform(v string) {
	b.Params["platform"] = v
}

// Extended parameter
//
// * 1 — to return additional fields 'screenshots', 'MAU', 'catalog_position', and 'international'.
// If set, 'count' must be less than or equal to '100'.
//
// * 0 — not to return additional fields (default).
func (b *AppsGetCatalogBuilder) Extended(v bool) {
	b.Params["extended"] = v
}

// ReturnFriends parameter
func (b *AppsGetCatalogBuilder) ReturnFriends(v bool) {
	b.Params["return_friends"] = v
}

// Fields parameter
func (b *AppsGetCatalogBuilder) Fields(v []string) {
	b.Params["fields"] = v
}

// NameCase parameter
func (b *AppsGetCatalogBuilder) NameCase(v string) {
	b.Params["name_case"] = v
}

// Q Search query string.
func (b *AppsGetCatalogBuilder) Q(v string) {
	b.Params["q"] = v
}

// GenreID parameter
func (b *AppsGetCatalogBuilder) GenreID(v int) {
	b.Params["genre_id"] = v
}

// Filter 'installed' — to return list of installed apps (only for mobile platform).
func (b *AppsGetCatalogBuilder) Filter(v string) {
	b.Params["filter"] = v
}

// AppsGetFriendsListBuilder builder
//
// Creates friends list for requests and invites in current app.
//
// https://vk.com/dev/apps.getFriendsList
type AppsGetFriendsListBuilder struct {
	api.Params
}

// NewAppsGetFriendsListBuilder func
func NewAppsGetFriendsListBuilder() *AppsGetFriendsListBuilder {
	return &AppsGetFriendsListBuilder{api.Params{}}
}

// Extended parameter
func (b *AppsGetFriendsListBuilder) Extended(v bool) {
	b.Params["extended"] = v
}

// Count List size.
func (b *AppsGetFriendsListBuilder) Count(v int) {
	b.Params["count"] = v
}

// Offset parameter
func (b *AppsGetFriendsListBuilder) Offset(v int) {
	b.Params["offset"] = v
}

// Type List type. Possible values:
//
// * invite — available for invites (don't play the game);
//
// * request — available for request (play the game).
//
// By default: invite.
func (b *AppsGetFriendsListBuilder) Type(v string) {
	b.Params["type"] = v
}

// Fields Additional profile fields, see [vk.com/dev/fields|description].
func (b *AppsGetFriendsListBuilder) Fields(v []string) {
	b.Params["fields"] = v
}

// AppsGetLeaderboardBuilder builder
//
// Returns players rating in the game.
//
// https://vk.com/dev/apps.getLeaderboard
type AppsGetLeaderboardBuilder struct {
	api.Params
}

// NewAppsGetLeaderboardBuilder func
func NewAppsGetLeaderboardBuilder() *AppsGetLeaderboardBuilder {
	return &AppsGetLeaderboardBuilder{api.Params{}}
}

// Type Leaderboard type. Possible values:
//
// * level — by level;
//
// * points — by mission points;
//
// * score — by score ().
func (b *AppsGetLeaderboardBuilder) Type(v string) {
	b.Params["type"] = v
}

// Global Rating type. Possible values:
//
// * 1 — global rating among all players;
//
// * 0 — rating among user friends.
func (b *AppsGetLeaderboardBuilder) Global(v bool) {
	b.Params["global"] = v
}

// Extended 1 — to return additional info about users
func (b *AppsGetLeaderboardBuilder) Extended(v bool) {
	b.Params["extended"] = v
}

// AppsGetScopesBuilder builder
//
// Returns scopes for auth
//
// https://vk.com/dev/apps.getScopes
type AppsGetScopesBuilder struct {
	api.Params
}

// NewAppsGetScopesBuilder func
func NewAppsGetScopesBuilder() *AppsGetScopesBuilder {
	return &AppsGetScopesBuilder{api.Params{}}
}

// Type parameter
func (b *AppsGetScopesBuilder) Type(v string) {
	b.Params["type"] = v
}

// AppsGetScoreBuilder builder
//
// Returns user score in app
//
// https://vk.com/dev/apps.getScore
type AppsGetScoreBuilder struct {
	api.Params
}

// NewAppsGetScoreBuilder func
func NewAppsGetScoreBuilder() *AppsGetScoreBuilder {
	return &AppsGetScoreBuilder{api.Params{}}
}

// UserID parameter
func (b *AppsGetScoreBuilder) UserID(v int) {
	b.Params["user_id"] = v
}

// AppsSendRequestBuilder builder
//
// Sends a request to another user in an app that uses VK authorization.
//
// https://vk.com/dev/apps.sendRequest
type AppsSendRequestBuilder struct {
	api.Params
}

// NewAppsSendRequestBuilder func
func NewAppsSendRequestBuilder() *AppsSendRequestBuilder {
	return &AppsSendRequestBuilder{api.Params{}}
}

// UserID id of the user to send a request
func (b *AppsSendRequestBuilder) UserID(v int) {
	b.Params["user_id"] = v
}

// Text request text
func (b *AppsSendRequestBuilder) Text(v string) {
	b.Params["text"] = v
}

// Type request type. Values:
//
// * invite – if the request is sent to a user who does not have the app installed
//
// * request – if a user has already installed the app
func (b *AppsSendRequestBuilder) Type(v string) {
	b.Params["type"] = v
}

// Name parameter
func (b *AppsSendRequestBuilder) Name(v string) {
	b.Params["name"] = v
}

// Key special string key to be sent with the request
func (b *AppsSendRequestBuilder) Key(v string) {
	b.Params["key"] = v
}

// Separate parameter
func (b *AppsSendRequestBuilder) Separate(v bool) {
	b.Params["separate"] = v
}
