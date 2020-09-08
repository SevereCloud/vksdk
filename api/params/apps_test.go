package params_test

import (
	"testing"

	"github.com/SevereCloud/vksdk/v2/api/params"
	"github.com/stretchr/testify/assert"
)

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

	assert.Equal(t, b.Params["app_id"], 1)
	assert.Equal(t, b.Params["app_ids"], []string{"text"})
	assert.Equal(t, b.Params["platform"], "text")
	assert.Equal(t, b.Params["extended"], true)
	assert.Equal(t, b.Params["return_friends"], true)
	assert.Equal(t, b.Params["fields"], []string{"test"})
	assert.Equal(t, b.Params["name_case"], "text")
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

	assert.Equal(t, b.Params["sort"], "text")
	assert.Equal(t, b.Params["offset"], 1)
	assert.Equal(t, b.Params["count"], 1)
	assert.Equal(t, b.Params["platform"], "text")
	assert.Equal(t, b.Params["extended"], true)
	assert.Equal(t, b.Params["return_friends"], true)
	assert.Equal(t, b.Params["fields"], []string{"test"})
	assert.Equal(t, b.Params["name_case"], "text")
	assert.Equal(t, b.Params["q"], "text")
	assert.Equal(t, b.Params["genre_id"], 1)
	assert.Equal(t, b.Params["filter"], "text")
}

func TestAppsGetFriendsListBuilder(t *testing.T) {
	t.Parallel()

	b := params.NewAppsGetFriendsListBuilder()

	b.Extended(true)
	b.Count(1)
	b.Offset(1)
	b.Type("text")
	b.Fields([]string{"test"})

	assert.Equal(t, b.Params["extended"], true)
	assert.Equal(t, b.Params["count"], 1)
	assert.Equal(t, b.Params["offset"], 1)
	assert.Equal(t, b.Params["type"], "text")
	assert.Equal(t, b.Params["fields"], []string{"test"})
}

func TestAppsGetLeaderboardBuilder(t *testing.T) {
	t.Parallel()

	b := params.NewAppsGetLeaderboardBuilder()

	b.Type("text")
	b.Global(true)
	b.Extended(true)

	assert.Equal(t, b.Params["type"], "text")
	assert.Equal(t, b.Params["global"], true)
	assert.Equal(t, b.Params["extended"], true)
}

func TestAppsGetScopesBuilder(t *testing.T) {
	t.Parallel()

	b := params.NewAppsGetScopesBuilder()

	b.Type("text")

	assert.Equal(t, b.Params["type"], "text")
}

func TestAppsGetScoreBuilder(t *testing.T) {
	t.Parallel()

	b := params.NewAppsGetScoreBuilder()

	b.UserID(1)

	assert.Equal(t, b.Params["user_id"], 1)
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

	assert.Equal(t, b.Params["user_id"], 1)
	assert.Equal(t, b.Params["text"], "text")
	assert.Equal(t, b.Params["type"], "text")
	assert.Equal(t, b.Params["name"], "text")
	assert.Equal(t, b.Params["key"], "text")
	assert.Equal(t, b.Params["separate"], true)
}
