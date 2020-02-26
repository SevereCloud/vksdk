package params // import "github.com/SevereCloud/vksdk/api/params"

import (
	"github.com/SevereCloud/vksdk/api"
)

// StoriesBanOwnerBuilder builder
//
// Allows to hide stories from chosen sources from current user's feed.
//
// https://vk.com/dev/stories.banOwner
type StoriesBanOwnerBuilder struct {
	api.Params
}

// NewStoriesBanOwnerBuilder func
func NewStoriesBanOwnerBuilder() *StoriesBanOwnerBuilder {
	return &StoriesBanOwnerBuilder{api.Params{}}
}

// OwnersIDs List of sources IDs
func (b *StoriesBanOwnerBuilder) OwnersIDs(v []int) *StoriesBanOwnerBuilder{
	b.Params["owners_ids"] = v
	return b
}

// StoriesDeleteBuilder builder
//
// Allows to delete story.
//
// https://vk.com/dev/stories.delete
type StoriesDeleteBuilder struct {
	api.Params
}

// NewStoriesDeleteBuilder func
func NewStoriesDeleteBuilder() *StoriesDeleteBuilder {
	return &StoriesDeleteBuilder{api.Params{}}
}

// OwnerID Story owner's ID.
// Current user id is used by default.
func (b *StoriesDeleteBuilder) OwnerID(v int) *StoriesDeleteBuilder{
	b.Params["owner_id"] = v
	return b
}

// StoryID Story ID.
func (b *StoriesDeleteBuilder) StoryID(v int) *StoriesDeleteBuilder{
	b.Params["story_id"] = v
	return b
}

// StoriesGetBuilder builder
//
// Returns stories available for current user.
//
// https://vk.com/dev/stories.get
type StoriesGetBuilder struct {
	api.Params
}

// NewStoriesGetBuilder func
func NewStoriesGetBuilder() *StoriesGetBuilder {
	return &StoriesGetBuilder{api.Params{}}
}

// OwnerID Owner ID.
func (b *StoriesGetBuilder) OwnerID(v int) *StoriesGetBuilder{
	b.Params["owner_id"] = v
	return b
}

// Extended parameter 
//
// * '1' — to return additional fields for users and communities.
// Default value is 0.
func (b *StoriesGetBuilder) Extended(v bool) *StoriesGetBuilder{
	b.Params["extended"] = v
	return b
}

// StoriesGetBannedBuilder builder
//
// Returns list of sources hidden from current user's feed.
//
// https://vk.com/dev/stories.getBanned
type StoriesGetBannedBuilder struct {
	api.Params
}

// NewStoriesGetBannedBuilder func
func NewStoriesGetBannedBuilder() *StoriesGetBannedBuilder {
	return &StoriesGetBannedBuilder{api.Params{}}
}

// Extended parameter 
//
// * '1' — to return additional fields for users and communities.
// Default value is 0.
func (b *StoriesGetBannedBuilder) Extended(v bool) *StoriesGetBannedBuilder{
	b.Params["extended"] = v
	return b
}

// Fields 
// Additional fields to return
func (b *StoriesGetBannedBuilder) Fields(v []string) *StoriesGetBannedBuilder{
	b.Params["fields"] = v
	return b
}

// StoriesGetByIDBuilder builder
//
// Returns story by its ID.
//
// https://vk.com/dev/stories.getById
type StoriesGetByIDBuilder struct {
	api.Params
}

// NewStoriesGetByIDBuilder func
func NewStoriesGetByIDBuilder() *StoriesGetByIDBuilder {
	return &StoriesGetByIDBuilder{api.Params{}}
}

// Stories Stories IDs separated by commas.
// Use format {owner_id}+'_'+{story_id}, for example, 12345_54331.
func (b *StoriesGetByIDBuilder) Stories(v []string) *StoriesGetByIDBuilder{
	b.Params["stories"] = v
	return b
}

// Extended parameter 
//
// * '1' — to return additional fields for users and communities.
// Default value is 0.
func (b *StoriesGetByIDBuilder) Extended(v bool) *StoriesGetByIDBuilder{
	b.Params["extended"] = v
	return b
}

// Fields 
// Additional fields to return
func (b *StoriesGetByIDBuilder) Fields(v []string) *StoriesGetByIDBuilder{
	b.Params["fields"] = v
	return b
}

// StoriesGetPhotoUploadServerBuilder builder
//
// Returns URL for uploading a story with photo.
//
// https://vk.com/dev/stories.getPhotoUploadServer
type StoriesGetPhotoUploadServerBuilder struct {
	api.Params
}

// NewStoriesGetPhotoUploadServerBuilder func
func NewStoriesGetPhotoUploadServerBuilder() *StoriesGetPhotoUploadServerBuilder {
	return &StoriesGetPhotoUploadServerBuilder{api.Params{}}
}

// AddToNews 1 — to add the story to friend's feed.
func (b *StoriesGetPhotoUploadServerBuilder) AddToNews(v bool) *StoriesGetPhotoUploadServerBuilder{
	b.Params["add_to_news"] = v
	return b
}

// UserIDs List of users IDs who can see the story.
func (b *StoriesGetPhotoUploadServerBuilder) UserIDs(v []int) *StoriesGetPhotoUploadServerBuilder{
	b.Params["user_ids"] = v
	return b
}

// ReplyToStory ID of the story to reply with the current.
func (b *StoriesGetPhotoUploadServerBuilder) ReplyToStory(v string) *StoriesGetPhotoUploadServerBuilder{
	b.Params["reply_to_story"] = v
	return b
}

// LinkText Link text (for community's stories only).
func (b *StoriesGetPhotoUploadServerBuilder) LinkText(v string) *StoriesGetPhotoUploadServerBuilder{
	b.Params["link_text"] = v
	return b
}

// LinkURL Link URL.
// Internal links on https://vk.com only.
func (b *StoriesGetPhotoUploadServerBuilder) LinkURL(v string) *StoriesGetPhotoUploadServerBuilder{
	b.Params["link_url"] = v
	return b
}

// GroupID ID of the community to upload the story (should be verified or with the fire icon).
func (b *StoriesGetPhotoUploadServerBuilder) GroupID(v int) *StoriesGetPhotoUploadServerBuilder{
	b.Params["group_id"] = v
	return b
}

// StoriesGetRepliesBuilder builder
//
// Returns replies to the story.
//
// https://vk.com/dev/stories.getReplies
type StoriesGetRepliesBuilder struct {
	api.Params
}

// NewStoriesGetRepliesBuilder func
func NewStoriesGetRepliesBuilder() *StoriesGetRepliesBuilder {
	return &StoriesGetRepliesBuilder{api.Params{}}
}

// OwnerID Story owner ID.
func (b *StoriesGetRepliesBuilder) OwnerID(v int) *StoriesGetRepliesBuilder{
	b.Params["owner_id"] = v
	return b
}

// StoryID Story ID.
func (b *StoriesGetRepliesBuilder) StoryID(v int) *StoriesGetRepliesBuilder{
	b.Params["story_id"] = v
	return b
}

// AccessKey Access key for the private object.
func (b *StoriesGetRepliesBuilder) AccessKey(v string) *StoriesGetRepliesBuilder{
	b.Params["access_key"] = v
	return b
}

// Extended parameter 
//
// * '1' — to return additional fields for users and communities.
// Default value is 0.
func (b *StoriesGetRepliesBuilder) Extended(v bool) *StoriesGetRepliesBuilder{
	b.Params["extended"] = v
	return b
}

// Fields 
// Additional fields to return
func (b *StoriesGetRepliesBuilder) Fields(v []string) *StoriesGetRepliesBuilder{
	b.Params["fields"] = v
	return b
}

// StoriesGetStatsBuilder builder
//
// Returns stories available for current user.
//
// https://vk.com/dev/stories.getStats
type StoriesGetStatsBuilder struct {
	api.Params
}

// NewStoriesGetStatsBuilder func
func NewStoriesGetStatsBuilder() *StoriesGetStatsBuilder {
	return &StoriesGetStatsBuilder{api.Params{}}
}

// OwnerID Story owner ID.
func (b *StoriesGetStatsBuilder) OwnerID(v int) *StoriesGetStatsBuilder{
	b.Params["owner_id"] = v
	return b
}

// StoryID Story ID.
func (b *StoriesGetStatsBuilder) StoryID(v int) *StoriesGetStatsBuilder{
	b.Params["story_id"] = v
	return b
}

// StoriesGetVideoUploadServerBuilder builder
//
// Allows to receive URL for uploading story with video.
//
// https://vk.com/dev/stories.getVideoUploadServer
type StoriesGetVideoUploadServerBuilder struct {
	api.Params
}

// NewStoriesGetVideoUploadServerBuilder func
func NewStoriesGetVideoUploadServerBuilder() *StoriesGetVideoUploadServerBuilder {
	return &StoriesGetVideoUploadServerBuilder{api.Params{}}
}

// AddToNews 1 — to add the story to friend's feed.
func (b *StoriesGetVideoUploadServerBuilder) AddToNews(v bool) *StoriesGetVideoUploadServerBuilder{
	b.Params["add_to_news"] = v
	return b
}

// UserIDs List of users IDs who can see the story.
func (b *StoriesGetVideoUploadServerBuilder) UserIDs(v []int) *StoriesGetVideoUploadServerBuilder{
	b.Params["user_ids"] = v
	return b
}

// ReplyToStory ID of the story to reply with the current.
func (b *StoriesGetVideoUploadServerBuilder) ReplyToStory(v string) *StoriesGetVideoUploadServerBuilder{
	b.Params["reply_to_story"] = v
	return b
}

// LinkText Link text (for community's stories only).
func (b *StoriesGetVideoUploadServerBuilder) LinkText(v string) *StoriesGetVideoUploadServerBuilder{
	b.Params["link_text"] = v
	return b
}

// LinkURL Link URL.
// Internal links on https://vk.com only.
func (b *StoriesGetVideoUploadServerBuilder) LinkURL(v string) *StoriesGetVideoUploadServerBuilder{
	b.Params["link_url"] = v
	return b
}

// GroupID ID of the community to upload the story (should be verified or with the fire icon).
func (b *StoriesGetVideoUploadServerBuilder) GroupID(v int) *StoriesGetVideoUploadServerBuilder{
	b.Params["group_id"] = v
	return b
}

// StoriesGetViewersBuilder builder
//
// Returns a list of story viewers.
//
// https://vk.com/dev/stories.getViewers
type StoriesGetViewersBuilder struct {
	api.Params
}

// NewStoriesGetViewersBuilder func
func NewStoriesGetViewersBuilder() *StoriesGetViewersBuilder {
	return &StoriesGetViewersBuilder{api.Params{}}
}

// OwnerID Story owner ID.
func (b *StoriesGetViewersBuilder) OwnerID(v int) *StoriesGetViewersBuilder{
	b.Params["owner_id"] = v
	return b
}

// StoryID Story ID.
func (b *StoriesGetViewersBuilder) StoryID(v int) *StoriesGetViewersBuilder{
	b.Params["story_id"] = v
	return b
}

// Count Maximum number of results.
func (b *StoriesGetViewersBuilder) Count(v int) *StoriesGetViewersBuilder{
	b.Params["count"] = v
	return b
}

// Offset Offset needed to return a specific subset of results.
func (b *StoriesGetViewersBuilder) Offset(v int) *StoriesGetViewersBuilder{
	b.Params["offset"] = v
	return b
}

// Extended parameter 
//
// * '1' — to return detailed information about photos
func (b *StoriesGetViewersBuilder) Extended(v bool) *StoriesGetViewersBuilder{
	b.Params["extended"] = v
	return b
}

// StoriesHideAllRepliesBuilder builder
//
// Hides all replies in the last 24 hours from the user to current user's stories.
//
// https://vk.com/dev/stories.hideAllReplies
type StoriesHideAllRepliesBuilder struct {
	api.Params
}

// NewStoriesHideAllRepliesBuilder func
func NewStoriesHideAllRepliesBuilder() *StoriesHideAllRepliesBuilder {
	return &StoriesHideAllRepliesBuilder{api.Params{}}
}

// OwnerID ID of the user whose replies should be hidden.
func (b *StoriesHideAllRepliesBuilder) OwnerID(v int) *StoriesHideAllRepliesBuilder{
	b.Params["owner_id"] = v
	return b
}

// GroupID parameter
func (b *StoriesHideAllRepliesBuilder) GroupID(v int) *StoriesHideAllRepliesBuilder{
	b.Params["group_id"] = v
	return b
}

// StoriesHideReplyBuilder builder
//
// Hides the reply to the current user's story.
//
// https://vk.com/dev/stories.hideReply
type StoriesHideReplyBuilder struct {
	api.Params
}

// NewStoriesHideReplyBuilder func
func NewStoriesHideReplyBuilder() *StoriesHideReplyBuilder {
	return &StoriesHideReplyBuilder{api.Params{}}
}

// OwnerID ID of the user whose replies should be hidden.
func (b *StoriesHideReplyBuilder) OwnerID(v int) *StoriesHideReplyBuilder{
	b.Params["owner_id"] = v
	return b
}

// StoryID Story ID.
func (b *StoriesHideReplyBuilder) StoryID(v int) *StoriesHideReplyBuilder{
	b.Params["story_id"] = v
	return b
}

// StoriesUnbanOwnerBuilder builder
//
// Allows to show stories from hidden sources in current user's feed.
//
// https://vk.com/dev/stories.unbanOwner
type StoriesUnbanOwnerBuilder struct {
	api.Params
}

// NewStoriesUnbanOwnerBuilder func
func NewStoriesUnbanOwnerBuilder() *StoriesUnbanOwnerBuilder {
	return &StoriesUnbanOwnerBuilder{api.Params{}}
}

// OwnersIDs List of hidden sources to show stories from.
func (b *StoriesUnbanOwnerBuilder) OwnersIDs(v []int) *StoriesUnbanOwnerBuilder{
	b.Params["owners_ids"] = v
	return b
}
