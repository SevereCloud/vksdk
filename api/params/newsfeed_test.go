package params_test

import (
	"testing"

	"github.com/SevereCloud/vksdk/api/params"
	"github.com/stretchr/testify/assert"
)

func TestNewsfeedAddBanBulder(t *testing.T) {
	b := params.NewNewsfeedAddBanBulder()

	b.UserIDs([]int{1})
	b.GroupIDs([]int{1})

	assert.Equal(t, b.Params["user_ids"], []int{1})
	assert.Equal(t, b.Params["group_ids"], []int{1})
}

func TestNewsfeedDeleteBanBulder(t *testing.T) {
	b := params.NewNewsfeedDeleteBanBulder()

	b.UserIDs([]int{1})
	b.GroupIDs([]int{1})

	assert.Equal(t, b.Params["user_ids"], []int{1})
	assert.Equal(t, b.Params["group_ids"], []int{1})
}

func TestNewsfeedDeleteListBulder(t *testing.T) {
	b := params.NewNewsfeedDeleteListBulder()

	b.ListID(1)

	assert.Equal(t, b.Params["list_id"], 1)
}

func TestNewsfeedGetBulder(t *testing.T) {
	b := params.NewNewsfeedGetBulder()

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

func TestNewsfeedGetBannedBulder(t *testing.T) {
	b := params.NewNewsfeedGetBannedBulder()

	b.Extended(true)
	b.Fields([]string{"test"})
	b.NameCase("text")

	assert.Equal(t, b.Params["extended"], true)
	assert.Equal(t, b.Params["fields"], []string{"test"})
	assert.Equal(t, b.Params["name_case"], "text")
}

func TestNewsfeedGetCommentsBulder(t *testing.T) {
	b := params.NewNewsfeedGetCommentsBulder()

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

func TestNewsfeedGetListsBulder(t *testing.T) {
	b := params.NewNewsfeedGetListsBulder()

	b.ListIDs([]int{1})
	b.Extended(true)

	assert.Equal(t, b.Params["list_ids"], []int{1})
	assert.Equal(t, b.Params["extended"], true)
}

func TestNewsfeedGetMentionsBulder(t *testing.T) {
	b := params.NewNewsfeedGetMentionsBulder()

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

func TestNewsfeedGetRecommendedBulder(t *testing.T) {
	b := params.NewNewsfeedGetRecommendedBulder()

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

func TestNewsfeedGetSuggestedSourcesBulder(t *testing.T) {
	b := params.NewNewsfeedGetSuggestedSourcesBulder()

	b.Offset(1)
	b.Count(1)
	b.Shuffle(true)
	b.Fields([]string{"test"})

	assert.Equal(t, b.Params["offset"], 1)
	assert.Equal(t, b.Params["count"], 1)
	assert.Equal(t, b.Params["shuffle"], true)
	assert.Equal(t, b.Params["fields"], []string{"test"})
}

func TestNewsfeedIgnoreItemBulder(t *testing.T) {
	b := params.NewNewsfeedIgnoreItemBulder()

	b.Type("text")
	b.OwnerID(1)
	b.ItemID(1)

	assert.Equal(t, b.Params["type"], "text")
	assert.Equal(t, b.Params["owner_id"], 1)
	assert.Equal(t, b.Params["item_id"], 1)
}

func TestNewsfeedSaveListBulder(t *testing.T) {
	b := params.NewNewsfeedSaveListBulder()

	b.ListID(1)
	b.Title("text")
	b.SourceIDs([]int{1})
	b.NoReposts(true)

	assert.Equal(t, b.Params["list_id"], 1)
	assert.Equal(t, b.Params["title"], "text")
	assert.Equal(t, b.Params["source_ids"], []int{1})
	assert.Equal(t, b.Params["no_reposts"], true)
}

func TestNewsfeedSearchBulder(t *testing.T) {
	b := params.NewNewsfeedSearchBulder()

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

func TestNewsfeedUnignoreItemBulder(t *testing.T) {
	b := params.NewNewsfeedUnignoreItemBulder()

	b.Type("text")
	b.OwnerID(1)
	b.ItemID(1)

	assert.Equal(t, b.Params["type"], "text")
	assert.Equal(t, b.Params["owner_id"], 1)
	assert.Equal(t, b.Params["item_id"], 1)
}

func TestNewsfeedUnsubscribeBulder(t *testing.T) {
	b := params.NewNewsfeedUnsubscribeBulder()

	b.Type("text")
	b.OwnerID(1)
	b.ItemID(1)

	assert.Equal(t, b.Params["type"], "text")
	assert.Equal(t, b.Params["owner_id"], 1)
	assert.Equal(t, b.Params["item_id"], 1)
}
