package params_test

import (
	"testing"

	"github.com/SevereCloud/vksdk/v2/api/params"
	"github.com/stretchr/testify/assert"
)

func TestNewsfeedAddBanBuilder(t *testing.T) {
	t.Parallel()

	b := params.NewNewsfeedAddBanBuilder()

	b.UserIDs([]int{1})
	b.GroupIDs([]int{1})

	assert.Equal(t, []int{1}, b.Params["user_ids"])
	assert.Equal(t, []int{1}, b.Params["group_ids"])
}

func TestNewsfeedDeleteBanBuilder(t *testing.T) {
	t.Parallel()

	b := params.NewNewsfeedDeleteBanBuilder()

	b.UserIDs([]int{1})
	b.GroupIDs([]int{1})

	assert.Equal(t, []int{1}, b.Params["user_ids"])
	assert.Equal(t, []int{1}, b.Params["group_ids"])
}

func TestNewsfeedDeleteListBuilder(t *testing.T) {
	t.Parallel()

	b := params.NewNewsfeedDeleteListBuilder()

	b.ListID(1)

	assert.Equal(t, 1, b.Params["list_id"])
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

	assert.Equal(t, []string{"test"}, b.Params["filters"])
	assert.Equal(t, true, b.Params["return_banned"])
	assert.Equal(t, 1, b.Params["start_time"])
	assert.Equal(t, 1, b.Params["end_time"])
	assert.Equal(t, 1, b.Params["max_photos"])
	assert.Equal(t, "text", b.Params["source_ids"])
	assert.Equal(t, "text", b.Params["start_from"])
	assert.Equal(t, 1, b.Params["count"])
	assert.Equal(t, []string{"test"}, b.Params["fields"])
	assert.Equal(t, "text", b.Params["section"])
}

func TestNewsfeedGetBannedBuilder(t *testing.T) {
	t.Parallel()

	b := params.NewNewsfeedGetBannedBuilder()

	b.Extended(true)
	b.Fields([]string{"test"})
	b.NameCase("text")

	assert.Equal(t, true, b.Params["extended"])
	assert.Equal(t, []string{"test"}, b.Params["fields"])
	assert.Equal(t, "text", b.Params["name_case"])
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

	assert.Equal(t, 1, b.Params["count"])
	assert.Equal(t, []string{"test"}, b.Params["filters"])
	assert.Equal(t, "text", b.Params["reposts"])
	assert.Equal(t, 1, b.Params["start_time"])
	assert.Equal(t, 1, b.Params["end_time"])
	assert.Equal(t, 1, b.Params["last_comments_count"])
	assert.Equal(t, "text", b.Params["start_from"])
	assert.Equal(t, []string{"test"}, b.Params["fields"])
}

func TestNewsfeedGetListsBuilder(t *testing.T) {
	t.Parallel()

	b := params.NewNewsfeedGetListsBuilder()

	b.ListIDs([]int{1})
	b.Extended(true)

	assert.Equal(t, []int{1}, b.Params["list_ids"])
	assert.Equal(t, true, b.Params["extended"])
}

func TestNewsfeedGetMentionsBuilder(t *testing.T) {
	t.Parallel()

	b := params.NewNewsfeedGetMentionsBuilder()

	b.OwnerID(1)
	b.StartTime(1)
	b.EndTime(1)
	b.Offset(1)
	b.Count(1)

	assert.Equal(t, 1, b.Params["owner_id"])
	assert.Equal(t, 1, b.Params["start_time"])
	assert.Equal(t, 1, b.Params["end_time"])
	assert.Equal(t, 1, b.Params["offset"])
	assert.Equal(t, 1, b.Params["count"])
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

	assert.Equal(t, 1, b.Params["start_time"])
	assert.Equal(t, 1, b.Params["end_time"])
	assert.Equal(t, 1, b.Params["max_photos"])
	assert.Equal(t, "text", b.Params["start_from"])
	assert.Equal(t, 1, b.Params["count"])
	assert.Equal(t, []string{"test"}, b.Params["fields"])
}

func TestNewsfeedGetSuggestedSourcesBuilder(t *testing.T) {
	t.Parallel()

	b := params.NewNewsfeedGetSuggestedSourcesBuilder()

	b.Offset(1)
	b.Count(1)
	b.Shuffle(true)
	b.Fields([]string{"test"})

	assert.Equal(t, 1, b.Params["offset"])
	assert.Equal(t, 1, b.Params["count"])
	assert.Equal(t, true, b.Params["shuffle"])
	assert.Equal(t, []string{"test"}, b.Params["fields"])
}

func TestNewsfeedIgnoreItemBuilder(t *testing.T) {
	t.Parallel()

	b := params.NewNewsfeedIgnoreItemBuilder()

	b.Type("text")
	b.OwnerID(1)
	b.ItemID(1)

	assert.Equal(t, "text", b.Params["type"])
	assert.Equal(t, 1, b.Params["owner_id"])
	assert.Equal(t, 1, b.Params["item_id"])
}

func TestNewsfeedSaveListBuilder(t *testing.T) {
	t.Parallel()

	b := params.NewNewsfeedSaveListBuilder()

	b.ListID(1)
	b.Title("text")
	b.SourceIDs([]int{1})
	b.NoReposts(true)

	assert.Equal(t, 1, b.Params["list_id"])
	assert.Equal(t, "text", b.Params["title"])
	assert.Equal(t, []int{1}, b.Params["source_ids"])
	assert.Equal(t, true, b.Params["no_reposts"])
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

	assert.Equal(t, "text", b.Params["q"])
	assert.Equal(t, true, b.Params["extended"])
	assert.Equal(t, 1, b.Params["count"])
	assert.InEpsilon(t, 1.1, b.Params["latitude"], 0.1)
	assert.InEpsilon(t, 1.1, b.Params["longitude"], 0.1)
	assert.Equal(t, 1, b.Params["start_time"])
	assert.Equal(t, 1, b.Params["end_time"])
	assert.Equal(t, "text", b.Params["start_from"])
	assert.Equal(t, []string{"test"}, b.Params["fields"])
}

func TestNewsfeedUnignoreItemBuilder(t *testing.T) {
	t.Parallel()

	b := params.NewNewsfeedUnignoreItemBuilder()

	b.Type("text")
	b.OwnerID(1)
	b.ItemID(1)
	b.TrackCode("text")

	assert.Equal(t, "text", b.Params["type"])
	assert.Equal(t, 1, b.Params["owner_id"])
	assert.Equal(t, 1, b.Params["item_id"])
	assert.Equal(t, "text", b.Params["track_code"])
}

func TestNewsfeedUnsubscribeBuilder(t *testing.T) {
	t.Parallel()

	b := params.NewNewsfeedUnsubscribeBuilder()

	b.Type("text")
	b.OwnerID(1)
	b.ItemID(1)

	assert.Equal(t, "text", b.Params["type"])
	assert.Equal(t, 1, b.Params["owner_id"])
	assert.Equal(t, 1, b.Params["item_id"])
}
