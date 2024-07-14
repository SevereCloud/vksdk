package params_test

import (
	"testing"

	"github.com/SevereCloud/vksdk/v3/api/params"
	"github.com/stretchr/testify/assert"
)

func TestStoriesBanOwnerBuilder(t *testing.T) {
	t.Parallel()

	b := params.NewStoriesBanOwnerBuilder()

	b.OwnersIDs([]int{1})

	assert.Equal(t, []int{1}, b.Params["owners_ids"])
}

func TestStoriesDeleteBuilder(t *testing.T) {
	t.Parallel()

	b := params.NewStoriesDeleteBuilder()

	b.OwnerID(1)
	b.StoryID(1)

	assert.Equal(t, 1, b.Params["owner_id"])
	assert.Equal(t, 1, b.Params["story_id"])
}

func TestStoriesGetBuilder(t *testing.T) {
	t.Parallel()

	b := params.NewStoriesGetBuilder()

	b.OwnerID(1)
	b.Extended(true)

	assert.Equal(t, 1, b.Params["owner_id"])
	assert.Equal(t, true, b.Params["extended"])
}

func TestStoriesGetBannedBuilder(t *testing.T) {
	t.Parallel()

	b := params.NewStoriesGetBannedBuilder()

	b.Extended(true)
	b.Fields([]string{"test"})

	assert.Equal(t, true, b.Params["extended"])
	assert.Equal(t, []string{"test"}, b.Params["fields"])
}

func TestStoriesGetByIDBuilder(t *testing.T) {
	t.Parallel()

	b := params.NewStoriesGetByIDBuilder()

	b.Stories([]string{"text"})
	b.Extended(true)
	b.Fields([]string{"test"})

	assert.Equal(t, []string{"text"}, b.Params["stories"])
	assert.Equal(t, true, b.Params["extended"])
	assert.Equal(t, []string{"test"}, b.Params["fields"])
}

func TestStoriesGetPhotoUploadServerBuilder(t *testing.T) {
	t.Parallel()

	b := params.NewStoriesGetPhotoUploadServerBuilder()

	b.AddToNews(true)
	b.UserIDs([]int{1})
	b.ReplyToStory("text")
	b.LinkText("text")
	b.LinkURL("text")
	b.GroupID(1)

	assert.Equal(t, true, b.Params["add_to_news"])
	assert.Equal(t, []int{1}, b.Params["user_ids"])
	assert.Equal(t, "text", b.Params["reply_to_story"])
	assert.Equal(t, "text", b.Params["link_text"])
	assert.Equal(t, "text", b.Params["link_url"])
	assert.Equal(t, 1, b.Params["group_id"])
}

func TestStoriesGetRepliesBuilder(t *testing.T) {
	t.Parallel()

	b := params.NewStoriesGetRepliesBuilder()

	b.OwnerID(1)
	b.StoryID(1)
	b.AccessKey("text")
	b.Extended(true)
	b.Fields([]string{"test"})

	assert.Equal(t, 1, b.Params["owner_id"])
	assert.Equal(t, 1, b.Params["story_id"])
	assert.Equal(t, "text", b.Params["access_key"])
	assert.Equal(t, true, b.Params["extended"])
	assert.Equal(t, []string{"test"}, b.Params["fields"])
}

func TestStoriesGetStatsBuilder(t *testing.T) {
	t.Parallel()

	b := params.NewStoriesGetStatsBuilder()

	b.OwnerID(1)
	b.StoryID(1)

	assert.Equal(t, 1, b.Params["owner_id"])
	assert.Equal(t, 1, b.Params["story_id"])
}

func TestStoriesGetVideoUploadServerBuilder(t *testing.T) {
	t.Parallel()

	b := params.NewStoriesGetVideoUploadServerBuilder()

	b.AddToNews(true)
	b.UserIDs([]int{1})
	b.ReplyToStory("text")
	b.LinkText("text")
	b.LinkURL("text")
	b.GroupID(1)

	assert.Equal(t, true, b.Params["add_to_news"])
	assert.Equal(t, []int{1}, b.Params["user_ids"])
	assert.Equal(t, "text", b.Params["reply_to_story"])
	assert.Equal(t, "text", b.Params["link_text"])
	assert.Equal(t, "text", b.Params["link_url"])
	assert.Equal(t, 1, b.Params["group_id"])
}

func TestStoriesGetViewersBuilder(t *testing.T) {
	t.Parallel()

	b := params.NewStoriesGetViewersBuilder()

	b.OwnerID(1)
	b.StoryID(1)
	b.Count(1)
	b.Offset(1)
	b.Extended(true)

	assert.Equal(t, 1, b.Params["owner_id"])
	assert.Equal(t, 1, b.Params["story_id"])
	assert.Equal(t, 1, b.Params["count"])
	assert.Equal(t, 1, b.Params["offset"])
	assert.Equal(t, true, b.Params["extended"])
}

func TestStoriesHideAllRepliesBuilder(t *testing.T) {
	t.Parallel()

	b := params.NewStoriesHideAllRepliesBuilder()

	b.OwnerID(1)
	b.GroupID(1)

	assert.Equal(t, 1, b.Params["owner_id"])
	assert.Equal(t, 1, b.Params["group_id"])
}

func TestStoriesHideReplyBuilder(t *testing.T) {
	t.Parallel()

	b := params.NewStoriesHideReplyBuilder()

	b.OwnerID(1)
	b.StoryID(1)

	assert.Equal(t, 1, b.Params["owner_id"])
	assert.Equal(t, 1, b.Params["story_id"])
}

func TestStoriesSaveBuilder(t *testing.T) {
	t.Parallel()

	b := params.NewStoriesSaveBuilder()

	b.UploadResults("test")

	assert.Equal(t, "test", b.Params["upload_results"])
}

func TestStoriesUnbanOwnerBuilder(t *testing.T) {
	t.Parallel()

	b := params.NewStoriesUnbanOwnerBuilder()

	b.OwnersIDs([]int{1})

	assert.Equal(t, []int{1}, b.Params["owners_ids"])
}

func TestStoriesSendInteractionBuilder(t *testing.T) {
	t.Parallel()

	b := params.NewStoriesSendInteractionBuilder()

	b.AccessKey("text")
	b.Message("text")
	b.IsBroadcast(true)
	b.IsAnonymous(true)
	b.UnseenMarker(true)

	assert.Equal(t, "text", b.Params["access_key"])
	assert.Equal(t, "text", b.Params["message"])
	assert.Equal(t, true, b.Params["is_broadcast"])
	assert.Equal(t, true, b.Params["is_anonymous"])
	assert.Equal(t, true, b.Params["unseen_marker"])
}
