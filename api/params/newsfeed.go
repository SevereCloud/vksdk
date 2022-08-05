package params // import "github.com/SevereCloud/vksdk/v2/api/params"

import (
	"github.com/SevereCloud/vksdk/v2/api"
)

// NewsfeedAddBanBuilder builder.
//
// Prevents news from specified users and communities from appearing in the current user's newsfeed.
//
// https://vk.com/dev/newsfeed.addBan
type NewsfeedAddBanBuilder struct {
	api.Params
}

// NewNewsfeedAddBanBuilder func.
func NewNewsfeedAddBanBuilder() *NewsfeedAddBanBuilder {
	return &NewsfeedAddBanBuilder{api.Params{}}
}

// UserIDs parameter.
func (b *NewsfeedAddBanBuilder) UserIDs(v []int) *NewsfeedAddBanBuilder {
	b.Params["user_ids"] = v
	return b
}

// GroupIDs parameter.
func (b *NewsfeedAddBanBuilder) GroupIDs(v []int) *NewsfeedAddBanBuilder {
	b.Params["group_ids"] = v
	return b
}

// NewsfeedDeleteBanBuilder builder.
//
// Allows news from previously banned users and communities to be shown in the current user's newsfeed.
//
// https://vk.com/dev/newsfeed.deleteBan
type NewsfeedDeleteBanBuilder struct {
	api.Params
}

// NewNewsfeedDeleteBanBuilder func.
func NewNewsfeedDeleteBanBuilder() *NewsfeedDeleteBanBuilder {
	return &NewsfeedDeleteBanBuilder{api.Params{}}
}

// UserIDs parameter.
func (b *NewsfeedDeleteBanBuilder) UserIDs(v []int) *NewsfeedDeleteBanBuilder {
	b.Params["user_ids"] = v
	return b
}

// GroupIDs parameter.
func (b *NewsfeedDeleteBanBuilder) GroupIDs(v []int) *NewsfeedDeleteBanBuilder {
	b.Params["group_ids"] = v
	return b
}

// NewsfeedDeleteListBuilder builder.
//
// https://vk.com/dev/newsfeed.deleteList
type NewsfeedDeleteListBuilder struct {
	api.Params
}

// NewNewsfeedDeleteListBuilder func.
func NewNewsfeedDeleteListBuilder() *NewsfeedDeleteListBuilder {
	return &NewsfeedDeleteListBuilder{api.Params{}}
}

// ListID parameter.
func (b *NewsfeedDeleteListBuilder) ListID(v int) *NewsfeedDeleteListBuilder {
	b.Params["list_id"] = v
	return b
}

// NewsfeedGetBuilder builder.
//
// Returns data required to show newsfeed for the current user.
//
// https://vk.com/dev/newsfeed.get
type NewsfeedGetBuilder struct {
	api.Params
}

// NewNewsfeedGetBuilder func.
func NewNewsfeedGetBuilder() *NewsfeedGetBuilder {
	return &NewsfeedGetBuilder{api.Params{}}
}

// Filters to apply:
//
// * post — new wall posts;
//
// * photo — new photos;
//
// * photo_tag — new photo tags;
//
// * wall_photo — new wall photos;
//
// * friend — new friends;
//
// * note — new notes.
func (b *NewsfeedGetBuilder) Filters(v []string) *NewsfeedGetBuilder {
	b.Params["filters"] = v
	return b
}

// ReturnBanned '1' — to return news items from banned sources.
func (b *NewsfeedGetBuilder) ReturnBanned(v bool) *NewsfeedGetBuilder {
	b.Params["return_banned"] = v
	return b
}

// StartTime earliest timestamp (in Unix time) of a news item to return. By default, 24 hours ago.
func (b *NewsfeedGetBuilder) StartTime(v int) *NewsfeedGetBuilder {
	b.Params["start_time"] = v
	return b
}

// EndTime latest timestamp (in Unix time) of a news item to return. By default, the current time.
func (b *NewsfeedGetBuilder) EndTime(v int) *NewsfeedGetBuilder {
	b.Params["end_time"] = v
	return b
}

// MaxPhotos maximum number of photos to return. By default, '5'.
func (b *NewsfeedGetBuilder) MaxPhotos(v int) *NewsfeedGetBuilder {
	b.Params["max_photos"] = v
	return b
}

// SourceIDs sources to obtain news from, separated by commas. User IDs can be specified in formats ” or 'u' ,
// where ” is the user's friend ID.
// Community IDs can be specified in formats '-' or 'g' , where ” is the community ID.
//
// If the parameter is not set, all of the user's friends and communities are returned, except for banned sources,
// which can be obtained with the [vk.com/dev/newsfeed.getBanned|newsfeed.getBanned] method.
func (b *NewsfeedGetBuilder) SourceIDs(v string) *NewsfeedGetBuilder {
	b.Params["source_ids"] = v
	return b
}

// StartFrom identifier required to get the next page of results. Value for this parameter is returned in 'next_from'
// field in a reply.
func (b *NewsfeedGetBuilder) StartFrom(v string) *NewsfeedGetBuilder {
	b.Params["start_from"] = v
	return b
}

// Count number of news items to return (default 50, maximum 100). For auto feed, you can use the 'new_offset'
// parameter returned by this method.
func (b *NewsfeedGetBuilder) Count(v int) *NewsfeedGetBuilder {
	b.Params["count"] = v
	return b
}

// Fields additional fields of [vk.com/dev/fields|profiles] and [vk.com/dev/fields_groups|communities] to return.
func (b *NewsfeedGetBuilder) Fields(v []string) *NewsfeedGetBuilder {
	b.Params["fields"] = v
	return b
}

// Section parameter.
func (b *NewsfeedGetBuilder) Section(v string) *NewsfeedGetBuilder {
	b.Params["section"] = v
	return b
}

// NewsfeedGetBannedBuilder builder.
//
// Returns a list of users and communities banned from the current user's newsfeed.
//
// https://vk.com/dev/newsfeed.getBanned
type NewsfeedGetBannedBuilder struct {
	api.Params
}

// NewNewsfeedGetBannedBuilder func.
func NewNewsfeedGetBannedBuilder() *NewsfeedGetBannedBuilder {
	return &NewsfeedGetBannedBuilder{api.Params{}}
}

// Extended '1' — return extra information about users and communities.
func (b *NewsfeedGetBannedBuilder) Extended(v bool) *NewsfeedGetBannedBuilder {
	b.Params["extended"] = v
	return b
}

// Fields profile fields to return.
func (b *NewsfeedGetBannedBuilder) Fields(v []string) *NewsfeedGetBannedBuilder {
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
func (b *NewsfeedGetBannedBuilder) NameCase(v string) *NewsfeedGetBannedBuilder {
	b.Params["name_case"] = v
	return b
}

// NewsfeedGetCommentsBuilder builder.
//
// Returns a list of comments in the current user's newsfeed.
//
// https://vk.com/dev/newsfeed.getComments
type NewsfeedGetCommentsBuilder struct {
	api.Params
}

// NewNewsfeedGetCommentsBuilder func.
func NewNewsfeedGetCommentsBuilder() *NewsfeedGetCommentsBuilder {
	return &NewsfeedGetCommentsBuilder{api.Params{}}
}

// Count number of comments to return. For auto feed, you can use the 'new_offset' parameter returned by this method.
func (b *NewsfeedGetCommentsBuilder) Count(v int) *NewsfeedGetCommentsBuilder {
	b.Params["count"] = v
	return b
}

// Filters to apply:
//
// * post — new comments on wall posts;
//
// * photo — new comments on photos;
//
// * video — new comments on videos;
//
// * topic — new comments on discussions;
//
// * note — new comments on notes.
func (b *NewsfeedGetCommentsBuilder) Filters(v []string) *NewsfeedGetCommentsBuilder {
	b.Params["filters"] = v
	return b
}

// Reposts object ID, comments on repost of which shall be returned, e.g. 'wall1_45486'.
// (If the parameter is set, the 'filters' parameter is optional.).
func (b *NewsfeedGetCommentsBuilder) Reposts(v string) *NewsfeedGetCommentsBuilder {
	b.Params["reposts"] = v
	return b
}

// StartTime earliest timestamp (in Unix time) of a comment to return. By default, 24 hours ago.
func (b *NewsfeedGetCommentsBuilder) StartTime(v int) *NewsfeedGetCommentsBuilder {
	b.Params["start_time"] = v
	return b
}

// EndTime latest timestamp (in Unix time) of a comment to return. By default, the current time.
func (b *NewsfeedGetCommentsBuilder) EndTime(v int) *NewsfeedGetCommentsBuilder {
	b.Params["end_time"] = v
	return b
}

// LastCommentsCount parameter.
func (b *NewsfeedGetCommentsBuilder) LastCommentsCount(v int) *NewsfeedGetCommentsBuilder {
	b.Params["last_comments_count"] = v
	return b
}

// StartFrom identificator needed to return the next page with results.
// Value for this parameter returns in 'next_from' field.
func (b *NewsfeedGetCommentsBuilder) StartFrom(v string) *NewsfeedGetCommentsBuilder {
	b.Params["start_from"] = v
	return b
}

// Fields additional fields of [vk.com/dev/fields|profiles] and [vk.com/dev/fields_groups|communities] to return.
func (b *NewsfeedGetCommentsBuilder) Fields(v []string) *NewsfeedGetCommentsBuilder {
	b.Params["fields"] = v
	return b
}

// NewsfeedGetListsBuilder builder.
//
// Returns a list of newsfeeds followed by the current user.
//
// https://vk.com/dev/newsfeed.getLists
type NewsfeedGetListsBuilder struct {
	api.Params
}

// NewNewsfeedGetListsBuilder func.
func NewNewsfeedGetListsBuilder() *NewsfeedGetListsBuilder {
	return &NewsfeedGetListsBuilder{api.Params{}}
}

// ListIDs numeric list identifiers.
func (b *NewsfeedGetListsBuilder) ListIDs(v []int) *NewsfeedGetListsBuilder {
	b.Params["list_ids"] = v
	return b
}

// Extended return additional list info.
func (b *NewsfeedGetListsBuilder) Extended(v bool) *NewsfeedGetListsBuilder {
	b.Params["extended"] = v
	return b
}

// NewsfeedGetMentionsBuilder builder.
//
// Returns a list of posts on user walls in which the current user is mentioned.
//
// https://vk.com/dev/newsfeed.getMentions
type NewsfeedGetMentionsBuilder struct {
	api.Params
}

// NewNewsfeedGetMentionsBuilder func.
func NewNewsfeedGetMentionsBuilder() *NewsfeedGetMentionsBuilder {
	return &NewsfeedGetMentionsBuilder{api.Params{}}
}

// OwnerID parameter.
func (b *NewsfeedGetMentionsBuilder) OwnerID(v int) *NewsfeedGetMentionsBuilder {
	b.Params["owner_id"] = v
	return b
}

// StartTime earliest timestamp (in Unix time) of a post to return. By default, 24 hours ago.
func (b *NewsfeedGetMentionsBuilder) StartTime(v int) *NewsfeedGetMentionsBuilder {
	b.Params["start_time"] = v
	return b
}

// EndTime latest timestamp (in Unix time) of a post to return. By default, the current time.
func (b *NewsfeedGetMentionsBuilder) EndTime(v int) *NewsfeedGetMentionsBuilder {
	b.Params["end_time"] = v
	return b
}

// Offset needed to return a specific subset of posts.
func (b *NewsfeedGetMentionsBuilder) Offset(v int) *NewsfeedGetMentionsBuilder {
	b.Params["offset"] = v
	return b
}

// Count number of posts to return.
func (b *NewsfeedGetMentionsBuilder) Count(v int) *NewsfeedGetMentionsBuilder {
	b.Params["count"] = v
	return b
}

// NewsfeedGetRecommendedBuilder builder.
//
// Returns a list of newsfeeds recommended to the current user.
//
// https://vk.com/dev/newsfeed.getRecommended
type NewsfeedGetRecommendedBuilder struct {
	api.Params
}

// NewNewsfeedGetRecommendedBuilder func.
func NewNewsfeedGetRecommendedBuilder() *NewsfeedGetRecommendedBuilder {
	return &NewsfeedGetRecommendedBuilder{api.Params{}}
}

// StartTime earliest timestamp (in Unix time) of a news item to return. By default, 24 hours ago.
func (b *NewsfeedGetRecommendedBuilder) StartTime(v int) *NewsfeedGetRecommendedBuilder {
	b.Params["start_time"] = v
	return b
}

// EndTime latest timestamp (in Unix time) of a news item to return. By default, the current time.
func (b *NewsfeedGetRecommendedBuilder) EndTime(v int) *NewsfeedGetRecommendedBuilder {
	b.Params["end_time"] = v
	return b
}

// MaxPhotos maximum number of photos to return. By default, '5'.
func (b *NewsfeedGetRecommendedBuilder) MaxPhotos(v int) *NewsfeedGetRecommendedBuilder {
	b.Params["max_photos"] = v
	return b
}

// StartFrom 'new_from' value obtained in previous call.
func (b *NewsfeedGetRecommendedBuilder) StartFrom(v string) *NewsfeedGetRecommendedBuilder {
	b.Params["start_from"] = v
	return b
}

// Count number of news items to return.
func (b *NewsfeedGetRecommendedBuilder) Count(v int) *NewsfeedGetRecommendedBuilder {
	b.Params["count"] = v
	return b
}

// Fields additional fields of [vk.com/dev/fields|profiles] and [vk.com/dev/fields_groups|communities] to return.
func (b *NewsfeedGetRecommendedBuilder) Fields(v []string) *NewsfeedGetRecommendedBuilder {
	b.Params["fields"] = v
	return b
}

// NewsfeedGetSuggestedSourcesBuilder builder.
//
// Returns communities and users that current user is suggested to follow.
//
// https://vk.com/dev/newsfeed.getSuggestedSources
type NewsfeedGetSuggestedSourcesBuilder struct {
	api.Params
}

// NewNewsfeedGetSuggestedSourcesBuilder func.
func NewNewsfeedGetSuggestedSourcesBuilder() *NewsfeedGetSuggestedSourcesBuilder {
	return &NewsfeedGetSuggestedSourcesBuilder{api.Params{}}
}

// Offset offset required to choose a particular subset of communities or users.
func (b *NewsfeedGetSuggestedSourcesBuilder) Offset(v int) *NewsfeedGetSuggestedSourcesBuilder {
	b.Params["offset"] = v
	return b
}

// Count amount of communities or users to return.
func (b *NewsfeedGetSuggestedSourcesBuilder) Count(v int) *NewsfeedGetSuggestedSourcesBuilder {
	b.Params["count"] = v
	return b
}

// Shuffle shuffle the returned list or not.
func (b *NewsfeedGetSuggestedSourcesBuilder) Shuffle(v bool) *NewsfeedGetSuggestedSourcesBuilder {
	b.Params["shuffle"] = v
	return b
}

// Fields list of extra fields to be returned. See available fields for [vk.com/dev/fields|users]
// and [vk.com/dev/fields_groups|communities].
func (b *NewsfeedGetSuggestedSourcesBuilder) Fields(v []string) *NewsfeedGetSuggestedSourcesBuilder {
	b.Params["fields"] = v
	return b
}

// NewsfeedIgnoreItemBuilder builder.
//
// Hides an item from the newsfeed.
//
// https://vk.com/dev/newsfeed.ignoreItem
type NewsfeedIgnoreItemBuilder struct {
	api.Params
}

// NewNewsfeedIgnoreItemBuilder func.
func NewNewsfeedIgnoreItemBuilder() *NewsfeedIgnoreItemBuilder {
	return &NewsfeedIgnoreItemBuilder{api.Params{}}
}

// Type item. Possible values:
//
// * wall – post on the wall;
//
// * tag – tag on a photo;
//
// * profilephoto – profile photo;
//
// * video – video;
//
// * audio – audio.
func (b *NewsfeedIgnoreItemBuilder) Type(v string) *NewsfeedIgnoreItemBuilder {
	b.Params["type"] = v
	return b
}

// OwnerID item owner's identifier (user or community),
// Note that community id must be negative. 'owner_id=1' – user , 'owner_id=-1' – community.
func (b *NewsfeedIgnoreItemBuilder) OwnerID(v int) *NewsfeedIgnoreItemBuilder {
	b.Params["owner_id"] = v
	return b
}

// ItemID item identifier.
func (b *NewsfeedIgnoreItemBuilder) ItemID(v int) *NewsfeedIgnoreItemBuilder {
	b.Params["item_id"] = v
	return b
}

// NewsfeedSaveListBuilder builder.
//
// Creates and edits user newsfeed lists.
//
// https://vk.com/dev/newsfeed.saveList
type NewsfeedSaveListBuilder struct {
	api.Params
}

// NewNewsfeedSaveListBuilder func.
func NewNewsfeedSaveListBuilder() *NewsfeedSaveListBuilder {
	return &NewsfeedSaveListBuilder{api.Params{}}
}

// ListID numeric list identifier (if not sent, will be set automatically).
func (b *NewsfeedSaveListBuilder) ListID(v int) *NewsfeedSaveListBuilder {
	b.Params["list_id"] = v
	return b
}

// Title list name.
func (b *NewsfeedSaveListBuilder) Title(v string) *NewsfeedSaveListBuilder {
	b.Params["title"] = v
	return b
}

// SourceIDs users and communities identifiers to be added to the list. Community identifiers must be negative numbers.
func (b *NewsfeedSaveListBuilder) SourceIDs(v []int) *NewsfeedSaveListBuilder {
	b.Params["source_ids"] = v
	return b
}

// NoReposts reposts display on and off ('1' is for off).
func (b *NewsfeedSaveListBuilder) NoReposts(v bool) *NewsfeedSaveListBuilder {
	b.Params["no_reposts"] = v
	return b
}

// NewsfeedSearchBuilder builder.
//
// Returns search results by statuses.
//
// https://vk.com/dev/newsfeed.search
type NewsfeedSearchBuilder struct {
	api.Params
}

// NewNewsfeedSearchBuilder func.
func NewNewsfeedSearchBuilder() *NewsfeedSearchBuilder {
	return &NewsfeedSearchBuilder{api.Params{}}
}

// Q search query string (e.g., 'New Year').
func (b *NewsfeedSearchBuilder) Q(v string) *NewsfeedSearchBuilder {
	b.Params["q"] = v
	return b
}

// Extended '1' — to return additional information about the user or community that placed the post.
func (b *NewsfeedSearchBuilder) Extended(v bool) *NewsfeedSearchBuilder {
	b.Params["extended"] = v
	return b
}

// Count number of posts to return.
func (b *NewsfeedSearchBuilder) Count(v int) *NewsfeedSearchBuilder {
	b.Params["count"] = v
	return b
}

// Latitude geographical latitude point (in degrees, -90 to 90) within which to search.
func (b *NewsfeedSearchBuilder) Latitude(v float64) *NewsfeedSearchBuilder {
	b.Params["latitude"] = v
	return b
}

// Longitude geographical longitude point (in degrees, -180 to 180) within which to search.
func (b *NewsfeedSearchBuilder) Longitude(v float64) *NewsfeedSearchBuilder {
	b.Params["longitude"] = v
	return b
}

// StartTime earliest timestamp (in Unix time) of a news item to return. By default, 24 hours ago.
func (b *NewsfeedSearchBuilder) StartTime(v int) *NewsfeedSearchBuilder {
	b.Params["start_time"] = v
	return b
}

// EndTime latest timestamp (in Unix time) of a news item to return. By default, the current time.
func (b *NewsfeedSearchBuilder) EndTime(v int) *NewsfeedSearchBuilder {
	b.Params["end_time"] = v
	return b
}

// StartFrom parameter.
func (b *NewsfeedSearchBuilder) StartFrom(v string) *NewsfeedSearchBuilder {
	b.Params["start_from"] = v
	return b
}

// Fields additional fields of [vk.com/dev/fields|profiles] and [vk.com/dev/fields_groups|communities] to return.
func (b *NewsfeedSearchBuilder) Fields(v []string) *NewsfeedSearchBuilder {
	b.Params["fields"] = v
	return b
}

// NewsfeedUnignoreItemBuilder builder.
//
// Returns a hidden item to the newsfeed.
//
// https://vk.com/dev/newsfeed.unignoreItem
type NewsfeedUnignoreItemBuilder struct {
	api.Params
}

// NewNewsfeedUnignoreItemBuilder func.
func NewNewsfeedUnignoreItemBuilder() *NewsfeedUnignoreItemBuilder {
	return &NewsfeedUnignoreItemBuilder{api.Params{}}
}

// Type item. Possible values:
//
// * wall – post on the wall;
//
// * tag – tag on a photo;
//
// * profilephoto – profile photo;
//
// * video – video;
//
// * audio – audio.
func (b *NewsfeedUnignoreItemBuilder) Type(v string) *NewsfeedUnignoreItemBuilder {
	b.Params["type"] = v
	return b
}

// OwnerID item owner's identifier (user or community),
// Note that community id must be negative. 'owner_id=1' – user , 'owner_id=-1' – community.
func (b *NewsfeedUnignoreItemBuilder) OwnerID(v int) *NewsfeedUnignoreItemBuilder {
	b.Params["owner_id"] = v
	return b
}

// ItemID parameter.
func (b *NewsfeedUnignoreItemBuilder) ItemID(v int) *NewsfeedUnignoreItemBuilder {
	b.Params["item_id"] = v
	return b
}

// TrackCode parameter.
func (b *NewsfeedUnignoreItemBuilder) TrackCode(v string) *NewsfeedUnignoreItemBuilder {
	b.Params["track_code"] = v
	return b
}

// NewsfeedUnsubscribeBuilder builder.
//
// Unsubscribes the current user from specified newsfeeds.
//
// https://vk.com/dev/newsfeed.unsubscribe
type NewsfeedUnsubscribeBuilder struct {
	api.Params
}

// NewNewsfeedUnsubscribeBuilder func.
func NewNewsfeedUnsubscribeBuilder() *NewsfeedUnsubscribeBuilder {
	return &NewsfeedUnsubscribeBuilder{api.Params{}}
}

// Type Type of object from which to unsubscribe:
//
// * note — note;
//
// * photo — photo;
//
// * post — post on user wall or community wall;
//
// * topic — topic;
//
// * video — video.
func (b *NewsfeedUnsubscribeBuilder) Type(v string) *NewsfeedUnsubscribeBuilder {
	b.Params["type"] = v
	return b
}

// OwnerID object owner ID.
func (b *NewsfeedUnsubscribeBuilder) OwnerID(v int) *NewsfeedUnsubscribeBuilder {
	b.Params["owner_id"] = v
	return b
}

// ItemID object ID.
func (b *NewsfeedUnsubscribeBuilder) ItemID(v int) *NewsfeedUnsubscribeBuilder {
	b.Params["item_id"] = v
	return b
}
