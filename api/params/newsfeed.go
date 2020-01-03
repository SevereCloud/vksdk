package params // import "github.com/SevereCloud/vksdk/api/params"

import (
	"github.com/SevereCloud/vksdk/api"
)

// NewsfeedAddBanBulder builder
//
// Prevents news from specified users and communities from appearing in the current user's newsfeed.
//
// https://vk.com/dev/newsfeed.addBan
type NewsfeedAddBanBulder struct {
	api.Params
}

// NewNewsfeedAddBanBulder func
func NewNewsfeedAddBanBulder() *NewsfeedAddBanBulder {
	return &NewsfeedAddBanBulder{api.Params{}}
}

// UserIDs parameter
func (b *NewsfeedAddBanBulder) UserIDs(v []int) {
	b.Params["user_ids"] = v
}

// GroupIDs parameter
func (b *NewsfeedAddBanBulder) GroupIDs(v []int) {
	b.Params["group_ids"] = v
}

// NewsfeedDeleteBanBulder builder
//
// Allows news from previously banned users and communities to be shown in the current user's newsfeed.
//
// https://vk.com/dev/newsfeed.deleteBan
type NewsfeedDeleteBanBulder struct {
	api.Params
}

// NewNewsfeedDeleteBanBulder func
func NewNewsfeedDeleteBanBulder() *NewsfeedDeleteBanBulder {
	return &NewsfeedDeleteBanBulder{api.Params{}}
}

// UserIDs parameter
func (b *NewsfeedDeleteBanBulder) UserIDs(v []int) {
	b.Params["user_ids"] = v
}

// GroupIDs parameter
func (b *NewsfeedDeleteBanBulder) GroupIDs(v []int) {
	b.Params["group_ids"] = v
}

// NewsfeedDeleteListBulder builder
//
// https://vk.com/dev/newsfeed.deleteList
type NewsfeedDeleteListBulder struct {
	api.Params
}

// NewNewsfeedDeleteListBulder func
func NewNewsfeedDeleteListBulder() *NewsfeedDeleteListBulder {
	return &NewsfeedDeleteListBulder{api.Params{}}
}

// ListID parameter
func (b *NewsfeedDeleteListBulder) ListID(v int) {
	b.Params["list_id"] = v
}

// NewsfeedGetBulder builder
//
// Returns data required to show newsfeed for the current user.
//
// https://vk.com/dev/newsfeed.get
type NewsfeedGetBulder struct {
	api.Params
}

// NewNewsfeedGetBulder func
func NewNewsfeedGetBulder() *NewsfeedGetBulder {
	return &NewsfeedGetBulder{api.Params{}}
}

// Filters Filters to apply: 'post' — new wall posts, 'photo' — new photos, 'photo_tag' — new photo tags, 'wall_photo' — new wall photos, 'friend' — new friends, 'note' — new notes
func (b *NewsfeedGetBulder) Filters(v []string) {
	b.Params["filters"] = v
}

// ReturnBanned '1' — to return news items from banned sources
func (b *NewsfeedGetBulder) ReturnBanned(v bool) {
	b.Params["return_banned"] = v
}

// StartTime Earliest timestamp (in Unix time) of a news item to return. By default, 24 hours ago.
func (b *NewsfeedGetBulder) StartTime(v int) {
	b.Params["start_time"] = v
}

// EndTime Latest timestamp (in Unix time) of a news item to return. By default, the current time.
func (b *NewsfeedGetBulder) EndTime(v int) {
	b.Params["end_time"] = v
}

// MaxPhotos Maximum number of photos to return. By default, '5'.
func (b *NewsfeedGetBulder) MaxPhotos(v int) {
	b.Params["max_photos"] = v
}

// SourceIDs Sources to obtain news from, separated by commas. User IDs can be specified in formats '' or 'u' , where '' is the user's friend ID. Community IDs can be specified in formats '-' or 'g' , where '' is the community ID. If the parameter is not set, all of the user's friends and communities are returned, except for banned sources, which can be obtained with the [vk.com/dev/newsfeed.getBanned|newsfeed.getBanned] method.
func (b *NewsfeedGetBulder) SourceIDs(v string) {
	b.Params["source_ids"] = v
}

// StartFrom identifier required to get the next page of results. Value for this parameter is returned in 'next_from' field in a reply.
func (b *NewsfeedGetBulder) StartFrom(v string) {
	b.Params["start_from"] = v
}

// Count Number of news items to return (default 50, maximum 100). For auto feed, you can use the 'new_offset' parameter returned by this method.
func (b *NewsfeedGetBulder) Count(v int) {
	b.Params["count"] = v
}

// Fields Additional fields of [vk.com/dev/fields|profiles] and [vk.com/dev/fields_groups|communities] to return.
func (b *NewsfeedGetBulder) Fields(v []string) {
	b.Params["fields"] = v
}

// Section parameter
func (b *NewsfeedGetBulder) Section(v string) {
	b.Params["section"] = v
}

// NewsfeedGetBannedBulder builder
//
// Returns a list of users and communities banned from the current user's newsfeed.
//
// https://vk.com/dev/newsfeed.getBanned
type NewsfeedGetBannedBulder struct {
	api.Params
}

// NewNewsfeedGetBannedBulder func
func NewNewsfeedGetBannedBulder() *NewsfeedGetBannedBulder {
	return &NewsfeedGetBannedBulder{api.Params{}}
}

// Extended '1' — return extra information about users and communities
func (b *NewsfeedGetBannedBulder) Extended(v bool) {
	b.Params["extended"] = v
}

// Fields Profile fields to return.
func (b *NewsfeedGetBannedBulder) Fields(v []string) {
	b.Params["fields"] = v
}

// NameCase Case for declension of user name and surname: 'nom' — nominative (default), 'gen' — genitive , 'dat' — dative, 'acc' — accusative , 'ins' — instrumental , 'abl' — prepositional
func (b *NewsfeedGetBannedBulder) NameCase(v string) {
	b.Params["name_case"] = v
}

// NewsfeedGetCommentsBulder builder
//
// Returns a list of comments in the current user's newsfeed.
//
// https://vk.com/dev/newsfeed.getComments
type NewsfeedGetCommentsBulder struct {
	api.Params
}

// NewNewsfeedGetCommentsBulder func
func NewNewsfeedGetCommentsBulder() *NewsfeedGetCommentsBulder {
	return &NewsfeedGetCommentsBulder{api.Params{}}
}

// Count Number of comments to return. For auto feed, you can use the 'new_offset' parameter returned by this method.
func (b *NewsfeedGetCommentsBulder) Count(v int) {
	b.Params["count"] = v
}

// Filters Filters to apply: 'post' — new comments on wall posts, 'photo' — new comments on photos, 'video' — new comments on videos, 'topic' — new comments on discussions, 'note' — new comments on notes,
func (b *NewsfeedGetCommentsBulder) Filters(v []string) {
	b.Params["filters"] = v
}

// Reposts Object ID, comments on repost of which shall be returned, e.g. 'wall1_45486'. (If the parameter is set, the 'filters' parameter is optional.),
func (b *NewsfeedGetCommentsBulder) Reposts(v string) {
	b.Params["reposts"] = v
}

// StartTime Earliest timestamp (in Unix time) of a comment to return. By default, 24 hours ago.
func (b *NewsfeedGetCommentsBulder) StartTime(v int) {
	b.Params["start_time"] = v
}

// EndTime Latest timestamp (in Unix time) of a comment to return. By default, the current time.
func (b *NewsfeedGetCommentsBulder) EndTime(v int) {
	b.Params["end_time"] = v
}

// LastCommentsCount parameter
func (b *NewsfeedGetCommentsBulder) LastCommentsCount(v int) {
	b.Params["last_comments_count"] = v
}

// StartFrom Identificator needed to return the next page with results. Value for this parameter returns in 'next_from' field.
func (b *NewsfeedGetCommentsBulder) StartFrom(v string) {
	b.Params["start_from"] = v
}

// Fields Additional fields of [vk.com/dev/fields|profiles] and [vk.com/dev/fields_groups|communities] to return.
func (b *NewsfeedGetCommentsBulder) Fields(v []string) {
	b.Params["fields"] = v
}

// NewsfeedGetListsBulder builder
//
// Returns a list of newsfeeds followed by the current user.
//
// https://vk.com/dev/newsfeed.getLists
type NewsfeedGetListsBulder struct {
	api.Params
}

// NewNewsfeedGetListsBulder func
func NewNewsfeedGetListsBulder() *NewsfeedGetListsBulder {
	return &NewsfeedGetListsBulder{api.Params{}}
}

// ListIDs numeric list identifiers.
func (b *NewsfeedGetListsBulder) ListIDs(v []int) {
	b.Params["list_ids"] = v
}

// Extended Return additional list info
func (b *NewsfeedGetListsBulder) Extended(v bool) {
	b.Params["extended"] = v
}

// NewsfeedGetMentionsBulder builder
//
// Returns a list of posts on user walls in which the current user is mentioned.
//
// https://vk.com/dev/newsfeed.getMentions
type NewsfeedGetMentionsBulder struct {
	api.Params
}

// NewNewsfeedGetMentionsBulder func
func NewNewsfeedGetMentionsBulder() *NewsfeedGetMentionsBulder {
	return &NewsfeedGetMentionsBulder{api.Params{}}
}

// OwnerID Owner ID.
func (b *NewsfeedGetMentionsBulder) OwnerID(v int) {
	b.Params["owner_id"] = v
}

// StartTime Earliest timestamp (in Unix time) of a post to return. By default, 24 hours ago.
func (b *NewsfeedGetMentionsBulder) StartTime(v int) {
	b.Params["start_time"] = v
}

// EndTime Latest timestamp (in Unix time) of a post to return. By default, the current time.
func (b *NewsfeedGetMentionsBulder) EndTime(v int) {
	b.Params["end_time"] = v
}

// Offset Offset needed to return a specific subset of posts.
func (b *NewsfeedGetMentionsBulder) Offset(v int) {
	b.Params["offset"] = v
}

// Count Number of posts to return.
func (b *NewsfeedGetMentionsBulder) Count(v int) {
	b.Params["count"] = v
}

// NewsfeedGetRecommendedBulder builder
//
// , Returns a list of newsfeeds recommended to the current user.
//
// https://vk.com/dev/newsfeed.getRecommended
type NewsfeedGetRecommendedBulder struct {
	api.Params
}

// NewNewsfeedGetRecommendedBulder func
func NewNewsfeedGetRecommendedBulder() *NewsfeedGetRecommendedBulder {
	return &NewsfeedGetRecommendedBulder{api.Params{}}
}

// StartTime Earliest timestamp (in Unix time) of a news item to return. By default, 24 hours ago.
func (b *NewsfeedGetRecommendedBulder) StartTime(v int) {
	b.Params["start_time"] = v
}

// EndTime Latest timestamp (in Unix time) of a news item to return. By default, the current time.
func (b *NewsfeedGetRecommendedBulder) EndTime(v int) {
	b.Params["end_time"] = v
}

// MaxPhotos Maximum number of photos to return. By default, '5'.
func (b *NewsfeedGetRecommendedBulder) MaxPhotos(v int) {
	b.Params["max_photos"] = v
}

// StartFrom 'new_from' value obtained in previous call.
func (b *NewsfeedGetRecommendedBulder) StartFrom(v string) {
	b.Params["start_from"] = v
}

// Count Number of news items to return.
func (b *NewsfeedGetRecommendedBulder) Count(v int) {
	b.Params["count"] = v
}

// Fields Additional fields of [vk.com/dev/fields|profiles] and [vk.com/dev/fields_groups|communities] to return.
func (b *NewsfeedGetRecommendedBulder) Fields(v []string) {
	b.Params["fields"] = v
}

// NewsfeedGetSuggestedSourcesBulder builder
//
// Returns communities and users that current user is suggested to follow.
//
// https://vk.com/dev/newsfeed.getSuggestedSources
type NewsfeedGetSuggestedSourcesBulder struct {
	api.Params
}

// NewNewsfeedGetSuggestedSourcesBulder func
func NewNewsfeedGetSuggestedSourcesBulder() *NewsfeedGetSuggestedSourcesBulder {
	return &NewsfeedGetSuggestedSourcesBulder{api.Params{}}
}

// Offset offset required to choose a particular subset of communities or users.
func (b *NewsfeedGetSuggestedSourcesBulder) Offset(v int) {
	b.Params["offset"] = v
}

// Count amount of communities or users to return.
func (b *NewsfeedGetSuggestedSourcesBulder) Count(v int) {
	b.Params["count"] = v
}

// Shuffle shuffle the returned list or not.
func (b *NewsfeedGetSuggestedSourcesBulder) Shuffle(v bool) {
	b.Params["shuffle"] = v
}

// Fields list of extra fields to be returned. See available fields for [vk.com/dev/fields|users] and [vk.com/dev/fields_groups|communities].
func (b *NewsfeedGetSuggestedSourcesBulder) Fields(v []string) {
	b.Params["fields"] = v
}

// NewsfeedIgnoreItemBulder builder
//
// Hides an item from the newsfeed.
//
// https://vk.com/dev/newsfeed.ignoreItem
type NewsfeedIgnoreItemBulder struct {
	api.Params
}

// NewNewsfeedIgnoreItemBulder func
func NewNewsfeedIgnoreItemBulder() *NewsfeedIgnoreItemBulder {
	return &NewsfeedIgnoreItemBulder{api.Params{}}
}

// Type Item type. Possible values: *'wall' – post on the wall,, *'tag' – tag on a photo,, *'profilephoto' – profile photo,, *'video' – video,, *'audio' – audio.
func (b *NewsfeedIgnoreItemBulder) Type(v string) {
	b.Params["type"] = v
}

// OwnerID Item owner's identifier (user or community), "Note that community id must be negative. 'owner_id=1' – user , 'owner_id=-1' – community "
func (b *NewsfeedIgnoreItemBulder) OwnerID(v int) {
	b.Params["owner_id"] = v
}

// ItemID Item identifier
func (b *NewsfeedIgnoreItemBulder) ItemID(v int) {
	b.Params["item_id"] = v
}

// NewsfeedSaveListBulder builder
//
// Creates and edits user newsfeed lists
//
// https://vk.com/dev/newsfeed.saveList
type NewsfeedSaveListBulder struct {
	api.Params
}

// NewNewsfeedSaveListBulder func
func NewNewsfeedSaveListBulder() *NewsfeedSaveListBulder {
	return &NewsfeedSaveListBulder{api.Params{}}
}

// ListID numeric list identifier (if not sent, will be set automatically).
func (b *NewsfeedSaveListBulder) ListID(v int) {
	b.Params["list_id"] = v
}

// Title list name.
func (b *NewsfeedSaveListBulder) Title(v string) {
	b.Params["title"] = v
}

// SourceIDs users and communities identifiers to be added to the list. Community identifiers must be negative numbers.
func (b *NewsfeedSaveListBulder) SourceIDs(v []int) {
	b.Params["source_ids"] = v
}

// NoReposts reposts display on and off ('1' is for off).
func (b *NewsfeedSaveListBulder) NoReposts(v bool) {
	b.Params["no_reposts"] = v
}

// NewsfeedSearchBulder builder
//
// Returns search results by statuses.
//
// https://vk.com/dev/newsfeed.search
type NewsfeedSearchBulder struct {
	api.Params
}

// NewNewsfeedSearchBulder func
func NewNewsfeedSearchBulder() *NewsfeedSearchBulder {
	return &NewsfeedSearchBulder{api.Params{}}
}

// Q Search query string (e.g., 'New Year').
func (b *NewsfeedSearchBulder) Q(v string) {
	b.Params["q"] = v
}

// Extended '1' — to return additional information about the user or community that placed the post.
func (b *NewsfeedSearchBulder) Extended(v bool) {
	b.Params["extended"] = v
}

// Count Number of posts to return.
func (b *NewsfeedSearchBulder) Count(v int) {
	b.Params["count"] = v
}

// Latitude Geographical latitude point (in degrees, -90 to 90) within which to search.
func (b *NewsfeedSearchBulder) Latitude(v float64) {
	b.Params["latitude"] = v
}

// Longitude Geographical longitude point (in degrees, -180 to 180) within which to search.
func (b *NewsfeedSearchBulder) Longitude(v float64) {
	b.Params["longitude"] = v
}

// StartTime Earliest timestamp (in Unix time) of a news item to return. By default, 24 hours ago.
func (b *NewsfeedSearchBulder) StartTime(v int) {
	b.Params["start_time"] = v
}

// EndTime Latest timestamp (in Unix time) of a news item to return. By default, the current time.
func (b *NewsfeedSearchBulder) EndTime(v int) {
	b.Params["end_time"] = v
}

// StartFrom parameter
func (b *NewsfeedSearchBulder) StartFrom(v string) {
	b.Params["start_from"] = v
}

// Fields Additional fields of [vk.com/dev/fields|profiles] and [vk.com/dev/fields_groups|communities] to return.
func (b *NewsfeedSearchBulder) Fields(v []string) {
	b.Params["fields"] = v
}

// NewsfeedUnignoreItemBulder builder
//
// Returns a hidden item to the newsfeed.
//
// https://vk.com/dev/newsfeed.unignoreItem
type NewsfeedUnignoreItemBulder struct {
	api.Params
}

// NewNewsfeedUnignoreItemBulder func
func NewNewsfeedUnignoreItemBulder() *NewsfeedUnignoreItemBulder {
	return &NewsfeedUnignoreItemBulder{api.Params{}}
}

// Type Item type. Possible values: *'wall' – post on the wall,, *'tag' – tag on a photo,, *'profilephoto' – profile photo,, *'video' – video,, *'audio' – audio.
func (b *NewsfeedUnignoreItemBulder) Type(v string) {
	b.Params["type"] = v
}

// OwnerID Item owner's identifier (user or community), "Note that community id must be negative. 'owner_id=1' – user , 'owner_id=-1' – community "
func (b *NewsfeedUnignoreItemBulder) OwnerID(v int) {
	b.Params["owner_id"] = v
}

// ItemID Item identifier
func (b *NewsfeedUnignoreItemBulder) ItemID(v int) {
	b.Params["item_id"] = v
}

// NewsfeedUnsubscribeBulder builder
//
// Unsubscribes the current user from specified newsfeeds.
//
// https://vk.com/dev/newsfeed.unsubscribe
type NewsfeedUnsubscribeBulder struct {
	api.Params
}

// NewNewsfeedUnsubscribeBulder func
func NewNewsfeedUnsubscribeBulder() *NewsfeedUnsubscribeBulder {
	return &NewsfeedUnsubscribeBulder{api.Params{}}
}

// Type Type of object from which to unsubscribe: 'note' — note, 'photo' — photo, 'post' — post on user wall or community wall, 'topic' — topic, 'video' — video
func (b *NewsfeedUnsubscribeBulder) Type(v string) {
	b.Params["type"] = v
}

// OwnerID Object owner ID.
func (b *NewsfeedUnsubscribeBulder) OwnerID(v int) {
	b.Params["owner_id"] = v
}

// ItemID Object ID.
func (b *NewsfeedUnsubscribeBulder) ItemID(v int) {
	b.Params["item_id"] = v
}
