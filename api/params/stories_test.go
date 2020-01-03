package params_test

import (
	"testing"

	"github.com/SevereCloud/vksdk/api/params"
	"github.com/stretchr/testify/assert"
)

func TestStoriesBanOwnerBulder(t *testing.T) {
	b := params.NewStoriesBanOwnerBulder()

	b.OwnersIDs([]int{1})

	assert.Equal(t, b.Params["owners_ids"], []int{1})
}

func TestStoriesDeleteBulder(t *testing.T) {
	b := params.NewStoriesDeleteBulder()

	b.OwnerID(1)
	b.StoryID(1)

	assert.Equal(t, b.Params["owner_id"], 1)
	assert.Equal(t, b.Params["story_id"], 1)
}

func TestStoriesGetBulder(t *testing.T) {
	b := params.NewStoriesGetBulder()

	b.OwnerID(1)
	b.Extended(true)

	assert.Equal(t, b.Params["owner_id"], 1)
	assert.Equal(t, b.Params["extended"], true)
}

func TestStoriesGetBannedBulder(t *testing.T) {
	b := params.NewStoriesGetBannedBulder()

	b.Extended(true)
	b.Fields([]string{"test"})

	assert.Equal(t, b.Params["extended"], true)
	assert.Equal(t, b.Params["fields"], []string{"test"})
}

func TestStoriesGetByIDBulder(t *testing.T) {
	b := params.NewStoriesGetByIDBulder()

	b.Stories([]string{"text"})
	b.Extended(true)
	b.Fields([]string{"test"})

	assert.Equal(t, b.Params["stories"], []string{"text"})
	assert.Equal(t, b.Params["extended"], true)
	assert.Equal(t, b.Params["fields"], []string{"test"})
}

func TestStoriesGetPhotoUploadServerBulder(t *testing.T) {
	b := params.NewStoriesGetPhotoUploadServerBulder()

	b.AddToNews(true)
	b.UserIDs([]int{1})
	b.ReplyToStory("text")
	b.LinkText("text")
	b.LinkURL("text")
	b.GroupID(1)

	assert.Equal(t, b.Params["add_to_news"], true)
	assert.Equal(t, b.Params["user_ids"], []int{1})
	assert.Equal(t, b.Params["reply_to_story"], "text")
	assert.Equal(t, b.Params["link_text"], "text")
	assert.Equal(t, b.Params["link_url"], "text")
	assert.Equal(t, b.Params["group_id"], 1)
}

func TestStoriesGetRepliesBulder(t *testing.T) {
	b := params.NewStoriesGetRepliesBulder()

	b.OwnerID(1)
	b.StoryID(1)
	b.AccessKey("text")
	b.Extended(true)
	b.Fields([]string{"test"})

	assert.Equal(t, b.Params["owner_id"], 1)
	assert.Equal(t, b.Params["story_id"], 1)
	assert.Equal(t, b.Params["access_key"], "text")
	assert.Equal(t, b.Params["extended"], true)
	assert.Equal(t, b.Params["fields"], []string{"test"})
}

func TestStoriesGetStatsBulder(t *testing.T) {
	b := params.NewStoriesGetStatsBulder()

	b.OwnerID(1)
	b.StoryID(1)

	assert.Equal(t, b.Params["owner_id"], 1)
	assert.Equal(t, b.Params["story_id"], 1)
}

func TestStoriesGetVideoUploadServerBulder(t *testing.T) {
	b := params.NewStoriesGetVideoUploadServerBulder()

	b.AddToNews(true)
	b.UserIDs([]int{1})
	b.ReplyToStory("text")
	b.LinkText("text")
	b.LinkURL("text")
	b.GroupID(1)

	assert.Equal(t, b.Params["add_to_news"], true)
	assert.Equal(t, b.Params["user_ids"], []int{1})
	assert.Equal(t, b.Params["reply_to_story"], "text")
	assert.Equal(t, b.Params["link_text"], "text")
	assert.Equal(t, b.Params["link_url"], "text")
	assert.Equal(t, b.Params["group_id"], 1)
}

func TestStoriesGetViewersBulder(t *testing.T) {
	b := params.NewStoriesGetViewersBulder()

	b.OwnerID(1)
	b.StoryID(1)
	b.Count(1)
	b.Offset(1)
	b.Extended(true)

	assert.Equal(t, b.Params["owner_id"], 1)
	assert.Equal(t, b.Params["story_id"], 1)
	assert.Equal(t, b.Params["count"], 1)
	assert.Equal(t, b.Params["offset"], 1)
	assert.Equal(t, b.Params["extended"], true)
}

func TestStoriesHideAllRepliesBulder(t *testing.T) {
	b := params.NewStoriesHideAllRepliesBulder()

	b.OwnerID(1)
	b.GroupID(1)

	assert.Equal(t, b.Params["owner_id"], 1)
	assert.Equal(t, b.Params["group_id"], 1)
}

func TestStoriesHideReplyBulder(t *testing.T) {
	b := params.NewStoriesHideReplyBulder()

	b.OwnerID(1)
	b.StoryID(1)

	assert.Equal(t, b.Params["owner_id"], 1)
	assert.Equal(t, b.Params["story_id"], 1)
}

func TestStoriesUnbanOwnerBulder(t *testing.T) {
	b := params.NewStoriesUnbanOwnerBulder()

	b.OwnersIDs([]int{1})

	assert.Equal(t, b.Params["owners_ids"], []int{1})
}
