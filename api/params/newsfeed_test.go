package params_test

import (
	"testing"

	"github.com/SevereCloud/vksdk/api/params"
	"github.com/stretchr/testify/assert"
)

func TestNewsfeedAddBanBuilder(t *testing.T) {
	t.Parallel()

	b := params.NewNewsfeedAddBanBuilder()

	b.UserIDs([]int{1})
	b.GroupIDs([]int{1})

	assert.Equal(t, b.Params["user_ids"], []int{1})
	assert.Equal(t, b.Params["group_ids"], []int{1})
}

func TestNewsfeedDeleteBanBuilder(t *testing.T) {
	t.Parallel()

	b := params.NewNewsfeedDeleteBanBuilder()

	b.UserIDs([]int{1})
	b.GroupIDs([]int{1})

	assert.Equal(t, b.Params["user_ids"], []int{1})
	assert.Equal(t, b.Params["group_ids"], []int{1})
}

func TestNewsfeedDeleteListBuilder(t *testing.T) {
	t.Parallel()

	b := params.NewNewsfeedDeleteListBuilder()

	b.ListID(1)

	assert.Equal(t, b.Params["list_id"], 1)
}

func TestNewsfeedGetBuilder(t *testing.T) {
	t.Parallel()

	b := params.NewNewsfeedGetBuilder()

	b.Filters([]string{"test"})
	b.ReturnBanned(true)
	b.StartTime(1)
	b.EndTime(1)
	b.MaxPhotos(1)
	b.SourceIDs("text")
	b.StartFrom("text")
	b.Count(1)
	b.Fields([]string{"test"})
	b.Section("text")

	assert.Equal(t, b.Params["filters"], []string{"test"})
	assert.Equal(t, b.Params["return_banned"], true)
	assert.Equal(t, b.Params["start_time"], 1)
	assert.Equal(t, b.Params["end_time"], 1)
	assert.Equal(t, b.Params["max_photos"], 1)
	assert.Equal(t, b.Params["source_ids"], "text")
	assert.Equal(t, b.Params["start_from"], "text")
	assert.Equal(t, b.Params["count"], 1)
	assert.Equal(t, b.Params["fields"], []string{"test"})
	assert.Equal(t, b.Params["section"], "text")
}

func TestNewsfeedGetBannedBuilder(t *testing.T) {
	t.Parallel()

	b := params.NewNewsfeedGetBannedBuilder()

	b.Extended(true)
	b.Fields([]string{"test"})
	b.NameCase("text")

	assert.Equal(t, b.Params["extended"], true)
	assert.Equal(t, b.Params["fields"], []string{"test"})
	assert.Equal(t, b.Params["name_case"], "text")
}

func TestNewsfeedGetCommentsBuilder(t *testing.T) {
	t.Parallel()

	b := params.NewNewsfeedGetCommentsBuilder()

	b.Count(1)
	b.Filters([]string{"test"})
	b.Reposts("text")
	b.StartTime(1)
	b.EndTime(1)
	b.LastCommentsCount(1)
	b.StartFrom("text")
	b.Fields([]string{"test"})

	assert.Equal(t, b.Params["count"], 1)
	assert.Equal(t, b.Params["filters"], []string{"test"})
	assert.Equal(t, b.Params["reposts"], "text")
	assert.Equal(t, b.Params["start_time"], 1)
	assert.Equal(t, b.Params["end_time"], 1)
	assert.Equal(t, b.Params["last_comments_count"], 1)
	assert.Equal(t, b.Params["start_from"], "text")
	assert.Equal(t, b.Params["fields"], []string{"test"})
}

func TestNewsfeedGetListsBuilder(t *testing.T) {
	t.Parallel()

	b := params.NewNewsfeedGetListsBuilder()

	b.ListIDs([]int{1})
	b.Extended(true)

	assert.Equal(t, b.Params["list_ids"], []int{1})
	assert.Equal(t, b.Params["extended"], true)
}

func TestNewsfeedGetMentionsBuilder(t *testing.T) {
	t.Parallel()

	b := params.NewNewsfeedGetMentionsBuilder()

	b.OwnerID(1)
	b.StartTime(1)
	b.EndTime(1)
	b.Offset(1)
	b.Count(1)

	assert.Equal(t, b.Params["owner_id"], 1)
	assert.Equal(t, b.Params["start_time"], 1)
	assert.Equal(t, b.Params["end_time"], 1)
	assert.Equal(t, b.Params["offset"], 1)
	assert.Equal(t, b.Params["count"], 1)
}

func TestNewsfeedGetRecommendedBuilder(t *testing.T) {
	t.Parallel()

	b := params.NewNewsfeedGetRecommendedBuilder()

	b.StartTime(1)
	b.EndTime(1)
	b.MaxPhotos(1)
	b.StartFrom("text")
	b.Count(1)
	b.Fields([]string{"test"})

	assert.Equal(t, b.Params["start_time"], 1)
	assert.Equal(t, b.Params["end_time"], 1)
	assert.Equal(t, b.Params["max_photos"], 1)
	assert.Equal(t, b.Params["start_from"], "text")
	assert.Equal(t, b.Params["count"], 1)
	assert.Equal(t, b.Params["fields"], []string{"test"})
}

func TestNewsfeedGetSuggestedSourcesBuilder(t *testing.T) {
	t.Parallel()

	b := params.NewNewsfeedGetSuggestedSourcesBuilder()

	b.Offset(1)
	b.Count(1)
	b.Shuffle(true)
	b.Fields([]string{"test"})

	assert.Equal(t, b.Params["offset"], 1)
	assert.Equal(t, b.Params["count"], 1)
	assert.Equal(t, b.Params["shuffle"], true)
	assert.Equal(t, b.Params["fields"], []string{"test"})
}

func TestNewsfeedIgnoreItemBuilder(t *testing.T) {
	t.Parallel()

	b := params.NewNewsfeedIgnoreItemBuilder()

	b.Type("text")
	b.OwnerID(1)
	b.ItemID(1)

	assert.Equal(t, b.Params["type"], "text")
	assert.Equal(t, b.Params["owner_id"], 1)
	assert.Equal(t, b.Params["item_id"], 1)
}

func TestNewsfeedSaveListBuilder(t *testing.T) {
	t.Parallel()

	b := params.NewNewsfeedSaveListBuilder()

	b.ListID(1)
	b.Title("text")
	b.SourceIDs([]int{1})
	b.NoReposts(true)

	assert.Equal(t, b.Params["list_id"], 1)
	assert.Equal(t, b.Params["title"], "text")
	assert.Equal(t, b.Params["source_ids"], []int{1})
	assert.Equal(t, b.Params["no_reposts"], true)
}

func TestNewsfeedSearchBuilder(t *testing.T) {
	t.Parallel()

	b := params.NewNewsfeedSearchBuilder()

	b.Q("text")
	b.Extended(true)
	b.Count(1)
	b.Latitude(1.1)
	b.Longitude(1.1)
	b.StartTime(1)
	b.EndTime(1)
	b.StartFrom("text")
	b.Fields([]string{"test"})

	assert.Equal(t, b.Params["q"], "text")
	assert.Equal(t, b.Params["extended"], true)
	assert.Equal(t, b.Params["count"], 1)
	assert.Equal(t, b.Params["latitude"], 1.1)
	assert.Equal(t, b.Params["longitude"], 1.1)
	assert.Equal(t, b.Params["start_time"], 1)
	assert.Equal(t, b.Params["end_time"], 1)
	assert.Equal(t, b.Params["start_from"], "text")
	assert.Equal(t, b.Params["fields"], []string{"test"})
}

func TestNewsfeedUnignoreItemBuilder(t *testing.T) {
	t.Parallel()

	b := params.NewNewsfeedUnignoreItemBuilder()

	b.Type("text")
	b.OwnerID(1)
	b.ItemID(1)
	b.TrackCode("text")

	assert.Equal(t, b.Params["type"], "text")
	assert.Equal(t, b.Params["owner_id"], 1)
	assert.Equal(t, b.Params["item_id"], 1)
	assert.Equal(t, b.Params["track_code"], "text")
}

func TestNewsfeedUnsubscribeBuilder(t *testing.T) {
	t.Parallel()

	b := params.NewNewsfeedUnsubscribeBuilder()

	b.Type("text")
	b.OwnerID(1)
	b.ItemID(1)

	assert.Equal(t, b.Params["type"], "text")
	assert.Equal(t, b.Params["owner_id"], 1)
	assert.Equal(t, b.Params["item_id"], 1)
}
