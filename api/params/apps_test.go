package params_test

import (
	"testing"

	"github.com/SevereCloud/vksdk/api/params"
	"github.com/stretchr/testify/assert"
)

func TestAppsGetBulder(t *testing.T) {
	b := params.NewAppsGetBulder()

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

func TestAppsGetCatalogBulder(t *testing.T) {
	b := params.NewAppsGetCatalogBulder()

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

func TestAppsGetFriendsListBulder(t *testing.T) {
	b := params.NewAppsGetFriendsListBulder()

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

func TestAppsGetLeaderboardBulder(t *testing.T) {
	b := params.NewAppsGetLeaderboardBulder()

	b.Type("text")
	b.Global(true)
	b.Extended(true)

	assert.Equal(t, b.Params["type"], "text")
	assert.Equal(t, b.Params["global"], true)
	assert.Equal(t, b.Params["extended"], true)
}

func TestAppsGetScopesBulder(t *testing.T) {
	b := params.NewAppsGetScopesBulder()

	b.Type("text")

	assert.Equal(t, b.Params["type"], "text")
}

func TestAppsGetScoreBulder(t *testing.T) {
	b := params.NewAppsGetScoreBulder()

	b.UserID(1)

	assert.Equal(t, b.Params["user_id"], 1)
}

func TestAppsSendRequestBulder(t *testing.T) {
	b := params.NewAppsSendRequestBulder()

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
