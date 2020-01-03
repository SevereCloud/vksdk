package params // import "github.com/SevereCloud/vksdk/api/params"

import (
	"github.com/SevereCloud/vksdk/api"
)

// StoriesBanOwnerBulder builder
//
// Allows to hide stories from chosen sources from current user's feed.
//
// https://vk.com/dev/stories.banOwner
type StoriesBanOwnerBulder struct {
	api.Params
}

// NewStoriesBanOwnerBulder func
func NewStoriesBanOwnerBulder() *StoriesBanOwnerBulder {
	return &StoriesBanOwnerBulder{api.Params{}}
}

// OwnersIDs List of sources IDs
func (b *StoriesBanOwnerBulder) OwnersIDs(v []int) {
	b.Params["owners_ids"] = v
}

// StoriesDeleteBulder builder
//
// Allows to delete story.
//
// https://vk.com/dev/stories.delete
type StoriesDeleteBulder struct {
	api.Params
}

// NewStoriesDeleteBulder func
func NewStoriesDeleteBulder() *StoriesDeleteBulder {
	return &StoriesDeleteBulder{api.Params{}}
}

// OwnerID Story owner's ID. Current user id is used by default.
func (b *StoriesDeleteBulder) OwnerID(v int) {
	b.Params["owner_id"] = v
}

// StoryID Story ID.
func (b *StoriesDeleteBulder) StoryID(v int) {
	b.Params["story_id"] = v
}

// StoriesGetBulder builder
//
// Returns stories available for current user.
//
// https://vk.com/dev/stories.get
type StoriesGetBulder struct {
	api.Params
}

// NewStoriesGetBulder func
func NewStoriesGetBulder() *StoriesGetBulder {
	return &StoriesGetBulder{api.Params{}}
}

// OwnerID Owner ID.
func (b *StoriesGetBulder) OwnerID(v int) {
	b.Params["owner_id"] = v
}

// Extended '1' — to return additional fields for users and communities. Default value is 0.
func (b *StoriesGetBulder) Extended(v bool) {
	b.Params["extended"] = v
}

// StoriesGetBannedBulder builder
//
// Returns list of sources hidden from current user's feed.
//
// https://vk.com/dev/stories.getBanned
type StoriesGetBannedBulder struct {
	api.Params
}

// NewStoriesGetBannedBulder func
func NewStoriesGetBannedBulder() *StoriesGetBannedBulder {
	return &StoriesGetBannedBulder{api.Params{}}
}

// Extended '1' — to return additional fields for users and communities. Default value is 0.
func (b *StoriesGetBannedBulder) Extended(v bool) {
	b.Params["extended"] = v
}

// Fields Additional fields to return
func (b *StoriesGetBannedBulder) Fields(v []string) {
	b.Params["fields"] = v
}

// StoriesGetByIDBulder builder
//
// Returns story by its ID.
//
// https://vk.com/dev/stories.getById
type StoriesGetByIDBulder struct {
	api.Params
}

// NewStoriesGetByIDBulder func
func NewStoriesGetByIDBulder() *StoriesGetByIDBulder {
	return &StoriesGetByIDBulder{api.Params{}}
}

// Stories Stories IDs separated by commas. Use format {owner_id}+'_'+{story_id}, for example, 12345_54331.
func (b *StoriesGetByIDBulder) Stories(v []string) {
	b.Params["stories"] = v
}

// Extended '1' — to return additional fields for users and communities. Default value is 0.
func (b *StoriesGetByIDBulder) Extended(v bool) {
	b.Params["extended"] = v
}

// Fields Additional fields to return
func (b *StoriesGetByIDBulder) Fields(v []string) {
	b.Params["fields"] = v
}

// StoriesGetPhotoUploadServerBulder builder
//
// Returns URL for uploading a story with photo.
//
// https://vk.com/dev/stories.getPhotoUploadServer
type StoriesGetPhotoUploadServerBulder struct {
	api.Params
}

// NewStoriesGetPhotoUploadServerBulder func
func NewStoriesGetPhotoUploadServerBulder() *StoriesGetPhotoUploadServerBulder {
	return &StoriesGetPhotoUploadServerBulder{api.Params{}}
}

// AddToNews 1 — to add the story to friend's feed.
func (b *StoriesGetPhotoUploadServerBulder) AddToNews(v bool) {
	b.Params["add_to_news"] = v
}

// UserIDs List of users IDs who can see the story.
func (b *StoriesGetPhotoUploadServerBulder) UserIDs(v []int) {
	b.Params["user_ids"] = v
}

// ReplyToStory ID of the story to reply with the current.
func (b *StoriesGetPhotoUploadServerBulder) ReplyToStory(v string) {
	b.Params["reply_to_story"] = v
}

// LinkText Link text (for community's stories only).
func (b *StoriesGetPhotoUploadServerBulder) LinkText(v string) {
	b.Params["link_text"] = v
}

// LinkURL Link URL. Internal links on https://vk.com only.
func (b *StoriesGetPhotoUploadServerBulder) LinkURL(v string) {
	b.Params["link_url"] = v
}

// GroupID ID of the community to upload the story (should be verified or with the "fire" icon).
func (b *StoriesGetPhotoUploadServerBulder) GroupID(v int) {
	b.Params["group_id"] = v
}

// StoriesGetRepliesBulder builder
//
// Returns replies to the story.
//
// https://vk.com/dev/stories.getReplies
type StoriesGetRepliesBulder struct {
	api.Params
}

// NewStoriesGetRepliesBulder func
func NewStoriesGetRepliesBulder() *StoriesGetRepliesBulder {
	return &StoriesGetRepliesBulder{api.Params{}}
}

// OwnerID Story owner ID.
func (b *StoriesGetRepliesBulder) OwnerID(v int) {
	b.Params["owner_id"] = v
}

// StoryID Story ID.
func (b *StoriesGetRepliesBulder) StoryID(v int) {
	b.Params["story_id"] = v
}

// AccessKey Access key for the private object.
func (b *StoriesGetRepliesBulder) AccessKey(v string) {
	b.Params["access_key"] = v
}

// Extended '1' — to return additional fields for users and communities. Default value is 0.
func (b *StoriesGetRepliesBulder) Extended(v bool) {
	b.Params["extended"] = v
}

// Fields Additional fields to return
func (b *StoriesGetRepliesBulder) Fields(v []string) {
	b.Params["fields"] = v
}

// StoriesGetStatsBulder builder
//
// Returns stories available for current user.
//
// https://vk.com/dev/stories.getStats
type StoriesGetStatsBulder struct {
	api.Params
}

// NewStoriesGetStatsBulder func
func NewStoriesGetStatsBulder() *StoriesGetStatsBulder {
	return &StoriesGetStatsBulder{api.Params{}}
}

// OwnerID Story owner ID.
func (b *StoriesGetStatsBulder) OwnerID(v int) {
	b.Params["owner_id"] = v
}

// StoryID Story ID.
func (b *StoriesGetStatsBulder) StoryID(v int) {
	b.Params["story_id"] = v
}

// StoriesGetVideoUploadServerBulder builder
//
// Allows to receive URL for uploading story with video.
//
// https://vk.com/dev/stories.getVideoUploadServer
type StoriesGetVideoUploadServerBulder struct {
	api.Params
}

// NewStoriesGetVideoUploadServerBulder func
func NewStoriesGetVideoUploadServerBulder() *StoriesGetVideoUploadServerBulder {
	return &StoriesGetVideoUploadServerBulder{api.Params{}}
}

// AddToNews 1 — to add the story to friend's feed.
func (b *StoriesGetVideoUploadServerBulder) AddToNews(v bool) {
	b.Params["add_to_news"] = v
}

// UserIDs List of users IDs who can see the story.
func (b *StoriesGetVideoUploadServerBulder) UserIDs(v []int) {
	b.Params["user_ids"] = v
}

// ReplyToStory ID of the story to reply with the current.
func (b *StoriesGetVideoUploadServerBulder) ReplyToStory(v string) {
	b.Params["reply_to_story"] = v
}

// LinkText Link text (for community's stories only).
func (b *StoriesGetVideoUploadServerBulder) LinkText(v string) {
	b.Params["link_text"] = v
}

// LinkURL Link URL. Internal links on https://vk.com only.
func (b *StoriesGetVideoUploadServerBulder) LinkURL(v string) {
	b.Params["link_url"] = v
}

// GroupID ID of the community to upload the story (should be verified or with the "fire" icon).
func (b *StoriesGetVideoUploadServerBulder) GroupID(v int) {
	b.Params["group_id"] = v
}

// StoriesGetViewersBulder builder
//
// Returns a list of story viewers.
//
// https://vk.com/dev/stories.getViewers
type StoriesGetViewersBulder struct {
	api.Params
}

// NewStoriesGetViewersBulder func
func NewStoriesGetViewersBulder() *StoriesGetViewersBulder {
	return &StoriesGetViewersBulder{api.Params{}}
}

// OwnerID Story owner ID.
func (b *StoriesGetViewersBulder) OwnerID(v int) {
	b.Params["owner_id"] = v
}

// StoryID Story ID.
func (b *StoriesGetViewersBulder) StoryID(v int) {
	b.Params["story_id"] = v
}

// Count Maximum number of results.
func (b *StoriesGetViewersBulder) Count(v int) {
	b.Params["count"] = v
}

// Offset Offset needed to return a specific subset of results.
func (b *StoriesGetViewersBulder) Offset(v int) {
	b.Params["offset"] = v
}

// Extended '1' — to return detailed information about photos
func (b *StoriesGetViewersBulder) Extended(v bool) {
	b.Params["extended"] = v
}

// StoriesHideAllRepliesBulder builder
//
// Hides all replies in the last 24 hours from the user to current user's stories.
//
// https://vk.com/dev/stories.hideAllReplies
type StoriesHideAllRepliesBulder struct {
	api.Params
}

// NewStoriesHideAllRepliesBulder func
func NewStoriesHideAllRepliesBulder() *StoriesHideAllRepliesBulder {
	return &StoriesHideAllRepliesBulder{api.Params{}}
}

// OwnerID ID of the user whose replies should be hidden.
func (b *StoriesHideAllRepliesBulder) OwnerID(v int) {
	b.Params["owner_id"] = v
}

// GroupID parameter
func (b *StoriesHideAllRepliesBulder) GroupID(v int) {
	b.Params["group_id"] = v
}

// StoriesHideReplyBulder builder
//
// Hides the reply to the current user's story.
//
// https://vk.com/dev/stories.hideReply
type StoriesHideReplyBulder struct {
	api.Params
}

// NewStoriesHideReplyBulder func
func NewStoriesHideReplyBulder() *StoriesHideReplyBulder {
	return &StoriesHideReplyBulder{api.Params{}}
}

// OwnerID ID of the user whose replies should be hidden.
func (b *StoriesHideReplyBulder) OwnerID(v int) {
	b.Params["owner_id"] = v
}

// StoryID Story ID.
func (b *StoriesHideReplyBulder) StoryID(v int) {
	b.Params["story_id"] = v
}

// StoriesUnbanOwnerBulder builder
//
// Allows to show stories from hidden sources in current user's feed.
//
// https://vk.com/dev/stories.unbanOwner
type StoriesUnbanOwnerBulder struct {
	api.Params
}

// NewStoriesUnbanOwnerBulder func
func NewStoriesUnbanOwnerBulder() *StoriesUnbanOwnerBulder {
	return &StoriesUnbanOwnerBulder{api.Params{}}
}

// OwnersIDs List of hidden sources to show stories from.
func (b *StoriesUnbanOwnerBulder) OwnersIDs(v []int) {
	b.Params["owners_ids"] = v
}
