package params_test

import (
	"testing"

	"github.com/SevereCloud/vksdk/api/params"
	"github.com/stretchr/testify/assert"
)

func TestStoriesBanOwnerBuilder(t *testing.T) {
	b := params.NewStoriesBanOwnerBuilder()

	b.OwnersIDs([]int{1})

	assert.Equal(t, b.Params["owners_ids"], []int{1})
}

func TestStoriesDeleteBuilder(t *testing.T) {
	b := params.NewStoriesDeleteBuilder()

	b.OwnerID(1)
	b.StoryID(1)

	assert.Equal(t, b.Params["owner_id"], 1)
	assert.Equal(t, b.Params["story_id"], 1)
}

func TestStoriesGetBuilder(t *testing.T) {
	b := params.NewStoriesGetBuilder()

	b.OwnerID(1)
	b.Extended(true)

	assert.Equal(t, b.Params["owner_id"], 1)
	assert.Equal(t, b.Params["extended"], true)
}

func TestStoriesGetBannedBuilder(t *testing.T) {
	b := params.NewStoriesGetBannedBuilder()

	b.Extended(true)
	b.Fields([]string{"test"})

	assert.Equal(t, b.Params["extended"], true)
	assert.Equal(t, b.Params["fields"], []string{"test"})
}

func TestStoriesGetByIDBuilder(t *testing.T) {
	b := params.NewStoriesGetByIDBuilder()

	b.Stories([]string{"text"})
	b.Extended(true)
	b.Fields([]string{"test"})

	assert.Equal(t, b.Params["stories"], []string{"text"})
	assert.Equal(t, b.Params["extended"], true)
	assert.Equal(t, b.Params["fields"], []string{"test"})
}

func TestStoriesGetPhotoUploadServerBuilder(t *testing.T) {
	b := params.NewStoriesGetPhotoUploadServerBuilder()

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

func TestStoriesGetRepliesBuilder(t *testing.T) {
	b := params.NewStoriesGetRepliesBuilder()

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

func TestStoriesGetStatsBuilder(t *testing.T) {
	b := params.NewStoriesGetStatsBuilder()

	b.OwnerID(1)
	b.StoryID(1)

	assert.Equal(t, b.Params["owner_id"], 1)
	assert.Equal(t, b.Params["story_id"], 1)
}

func TestStoriesGetVideoUploadServerBuilder(t *testing.T) {
	b := params.NewStoriesGetVideoUploadServerBuilder()

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

func TestStoriesGetViewersBuilder(t *testing.T) {
	b := params.NewStoriesGetViewersBuilder()

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

func TestStoriesHideAllRepliesBuilder(t *testing.T) {
	b := params.NewStoriesHideAllRepliesBuilder()

	b.OwnerID(1)
	b.GroupID(1)

	assert.Equal(t, b.Params["owner_id"], 1)
	assert.Equal(t, b.Params["group_id"], 1)
}

func TestStoriesHideReplyBuilder(t *testing.T) {
	b := params.NewStoriesHideReplyBuilder()

	b.OwnerID(1)
	b.StoryID(1)

	assert.Equal(t, b.Params["owner_id"], 1)
	assert.Equal(t, b.Params["story_id"], 1)
}

func TestStoriesUnbanOwnerBuilder(t *testing.T) {
	b := params.NewStoriesUnbanOwnerBuilder()

	b.OwnersIDs([]int{1})

	assert.Equal(t, b.Params["owners_ids"], []int{1})
}
