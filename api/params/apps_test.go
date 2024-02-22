package params_test

import (
	"testing"

	"github.com/SevereCloud/vksdk/v2/api/params"
	"github.com/stretchr/testify/assert"
)

func TestAppsAddUsersToTestingGroupBuilder(t *testing.T) {
	t.Parallel()

	b := params.NewAppsAddUsersToTestingGroupBuilder()

	b.UserIDs([]int{1})
	b.GroupID(1)

	assert.Equal(t, []int{1}, b.Params["user_ids"])
	assert.Equal(t, 1, b.Params["group_id"])
}

func TestAppsGetBuilder(t *testing.T) {
	t.Parallel()

	b := params.NewAppsGetBuilder()

	b.AppID(1)
	b.AppIDs([]string{"text"})
	b.Platform("text")
	b.Extended(true)
	b.ReturnFriends(true)
	b.Fields([]string{"test"})
	b.NameCase("text")

	assert.Equal(t, 1, b.Params["app_id"])
	assert.Equal(t, []string{"text"}, b.Params["app_ids"])
	assert.Equal(t, "text", b.Params["platform"])
	assert.Equal(t, true, b.Params["extended"])
	assert.Equal(t, true, b.Params["return_friends"])
	assert.Equal(t, []string{"test"}, b.Params["fields"])
	assert.Equal(t, "text", b.Params["name_case"])
}

func TestAppsGetCatalogBuilder(t *testing.T) {
	t.Parallel()

	b := params.NewAppsGetCatalogBuilder()

	b.Sort("text")
	b.Offset(1)
	b.Count(1)
	b.Platform("text")
	b.Extended(true)
	b.ReturnFriends(true)
	b.Fields([]string{"test"})
	b.NameCase("text")
	b.Q("text")
	b.GenreID(1)
	b.Filter("text")

	assert.Equal(t, "text", b.Params["sort"])
	assert.Equal(t, 1, b.Params["offset"])
	assert.Equal(t, 1, b.Params["count"])
	assert.Equal(t, "text", b.Params["platform"])
	assert.Equal(t, true, b.Params["extended"])
	assert.Equal(t, true, b.Params["return_friends"])
	assert.Equal(t, []string{"test"}, b.Params["fields"])
	assert.Equal(t, "text", b.Params["name_case"])
	assert.Equal(t, "text", b.Params["q"])
	assert.Equal(t, 1, b.Params["genre_id"])
	assert.Equal(t, "text", b.Params["filter"])
}

func TestAppsGetFriendsListBuilder(t *testing.T) {
	t.Parallel()

	b := params.NewAppsGetFriendsListBuilder()

	b.Extended(true)
	b.Count(1)
	b.Offset(1)
	b.Type("text")
	b.Fields([]string{"test"})

	assert.Equal(t, true, b.Params["extended"])
	assert.Equal(t, 1, b.Params["count"])
	assert.Equal(t, 1, b.Params["offset"])
	assert.Equal(t, "text", b.Params["type"])
	assert.Equal(t, []string{"test"}, b.Params["fields"])
}

func TestAppsGetLeaderboardBuilder(t *testing.T) {
	t.Parallel()

	b := params.NewAppsGetLeaderboardBuilder()

	b.Type("text")
	b.Global(true)
	b.Extended(true)

	assert.Equal(t, "text", b.Params["type"])
	assert.Equal(t, true, b.Params["global"])
	assert.Equal(t, true, b.Params["extended"])
}

func TestAppsGetScopesBuilder(t *testing.T) {
	t.Parallel()

	b := params.NewAppsGetScopesBuilder()

	b.Type("text")

	assert.Equal(t, "text", b.Params["type"])
}

func TestAppsGetScoreBuilder(t *testing.T) {
	t.Parallel()

	b := params.NewAppsGetScoreBuilder()

	b.UserID(1)

	assert.Equal(t, 1, b.Params["user_id"])
}

func TestAppsGetTestingGroupsBuilder(t *testing.T) {
	t.Parallel()

	b := params.NewAppsGetTestingGroupsBuilder()

	b.GroupID(1)

	assert.Equal(t, 1, b.Params["group_id"])
}

func TestAppsRemoveUsersFromTestingGroupsBuilder(t *testing.T) {
	t.Parallel()

	b := params.NewAppsRemoveUsersFromTestingGroupsBuilder()

	b.UserIDs([]int{1})

	assert.Equal(t, []int{1}, b.Params["user_ids"])
}

func TestAppsRemoveTestingGroupBuilder(t *testing.T) {
	t.Parallel()

	b := params.NewAppsRemoveTestingGroupBuilder()

	b.GroupID(1)

	assert.Equal(t, 1, b.Params["group_id"])
}

func TestAppsSendRequestBuilder(t *testing.T) {
	t.Parallel()

	b := params.NewAppsSendRequestBuilder()

	b.UserID(1)
	b.Text("text")
	b.Type("text")
	b.Name("text")
	b.Key("text")
	b.Separate(true)

	assert.Equal(t, 1, b.Params["user_id"])
	assert.Equal(t, "text", b.Params["text"])
	assert.Equal(t, "text", b.Params["type"])
	assert.Equal(t, "text", b.Params["name"])
	assert.Equal(t, "text", b.Params["key"])
	assert.Equal(t, true, b.Params["separate"])
}

func TestAppsUpdateMetaForTestingGroupBuilder(t *testing.T) {
	t.Parallel()

	b := params.NewAppsUpdateMetaForTestingGroupBuilder()

	b.GroupID(1)
	b.Webview("text")
	b.Name("text")
	b.Platforms([]string{"text"})
	b.UserIDs([]int{1})

	assert.Equal(t, 1, b.Params["group_id"])
	assert.Equal(t, "text", b.Params["webview"])
	assert.Equal(t, "text", b.Params["name"])
	assert.Equal(t, []string{"text"}, b.Params["platforms"])
	assert.Equal(t, []int{1}, b.Params["user_ids"])
}
