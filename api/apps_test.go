package api_test

import (
	"testing"

	"github.com/Derad6709/vksdk/v2/api"

	"github.com/stretchr/testify/assert"
)

func TestVK_AppsDeleteAppRequests(t *testing.T) {
	t.Parallel()

	needUserToken(t)

	res, err := vkUser.AppsDeleteAppRequests(nil)
	noError(t, err)
	assert.Equal(t, res, 1)
}

func TestVK_AppsGet(t *testing.T) {
	t.Parallel()

	needServiceToken(t)

	res, err := vkService.AppsGet(api.Params{
		"app_id":   4063926,
		"extended": true,
	})
	noError(t, err)
	assert.NotEmpty(t, res.Count)

	if assert.NotEmpty(t, res.Items) {
		assert.NotEmpty(t, res.Items[0].Type)
		assert.NotEmpty(t, res.Items[0].Title)
		assert.NotEmpty(t, res.Items[0].Icon139)
		assert.NotEmpty(t, res.Items[0].Icon150)
		assert.NotEmpty(t, res.Items[0].Icon278)
		assert.NotEmpty(t, res.Items[0].Banner1120)
		assert.NotEmpty(t, res.Items[0].Banner560)
		assert.NotEmpty(t, res.Items[0].Genre)
		assert.NotEmpty(t, res.Items[0].GenreID)
		// assert.NotEmpty(t, res.Items[0].IsInCatalog)
		// assert.NotEmpty(t, res.Items[0].Installed)
		// assert.NotEmpty(t, res.Items[0].LeaderboardType)
		assert.NotEmpty(t, res.Items[0].MembersCount)
		// assert.NotEmpty(t, res.Items[0].PublishedDate)
		assert.NotEmpty(t, res.Items[0].Section)
	}
}

func TestVK_AppsGetCatalog(t *testing.T) {
	t.Parallel()

	needServiceToken(t)

	res, err := vkService.AppsGetCatalog(nil)
	noError(t, err)
	assert.NotEmpty(t, res.Count)
	assert.NotEmpty(t, res.Items)
}

func TestVK_AppsGetFriendsList(t *testing.T) {
	t.Parallel()

	needUserToken(t)

	_, err := vkUser.AppsGetFriendsList(nil)
	noError(t, err)
	// assert.NotEmpty(t, res.Count)
	// assert.NotEmpty(t, res.Items)

	_, err = vkUser.AppsGetFriendsListExtended(nil)
	noError(t, err)
}

// TODO: TestVK_AppsGetLeaderboard

func TestVK_AppsGetScopes(t *testing.T) {
	t.Parallel()

	needUserToken(t)

	res, err := vkUser.AppsGetScopes(nil)
	noError(t, err)
	assert.NotEmpty(t, res.Count)

	if assert.NotEmpty(t, res.Items) {
		assert.NotEmpty(t, res.Items[0].Name)
		assert.NotEmpty(t, res.Items[0].Title)
	}
}

// TODO: TestVK_AppsGetScore
// TODO: TestVK_AppsSendRequest
