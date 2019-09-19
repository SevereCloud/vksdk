package api

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestVK_AppsDeleteAppRequests(t *testing.T) {
	needUserToken(t)

	res, err := vkUser.AppsDeleteAppRequests(map[string]string{})
	assert.NoError(t, err)
	assert.Equal(t, res, 1)
}

func TestVK_AppsGet(t *testing.T) {
	needServiceToken(t)

	res, err := vkService.AppsGet(map[string]string{
		"app_id":   "4063926",
		"extended": "1",
	})
	assert.NoError(t, err)
	assert.NotEmpty(t, res.Count)
	if assert.NotEmpty(t, res.Items) {
		assert.NotEmpty(t, res.Items[0].Type)
		assert.NotEmpty(t, res.Items[0].Title)
		assert.NotEmpty(t, res.Items[0].Icon139)
		assert.NotEmpty(t, res.Items[0].Icon150)
		assert.NotEmpty(t, res.Items[0].Icon278)
		assert.NotEmpty(t, res.Items[0].AuthorGroup)
		assert.NotEmpty(t, res.Items[0].AuthorURL)
		assert.NotEmpty(t, res.Items[0].Banner1120)
		assert.NotEmpty(t, res.Items[0].Banner560)
		assert.NotEmpty(t, res.Items[0].Genre)
		assert.NotEmpty(t, res.Items[0].GenreID)
		// assert.NotEmpty(t, res.Items[0].IsInCatalog)
		// assert.NotEmpty(t, res.Items[0].Installed)
		// assert.NotEmpty(t, res.Items[0].LeaderboardType)
		assert.NotEmpty(t, res.Items[0].MembersCount)
		assert.NotEmpty(t, res.Items[0].PublishedDate)
		assert.NotEmpty(t, res.Items[0].Section)
	}
	// assert.NotEmpty(t, res.Profiles)
	// assert.NotEmpty(t, res.Groups)
}

func TestVK_AppsGetCatalog(t *testing.T) {
	needServiceToken(t)

	res, err := vkService.AppsGetCatalog(map[string]string{})
	assert.NoError(t, err)
	assert.NotEmpty(t, res.Count)
	assert.NotEmpty(t, res.Items)
}

func TestVK_AppsGetFriendsList(t *testing.T) {
	needUserToken(t)

	_, err := vkUser.AppsGetFriendsList(map[string]string{})
	assert.NoError(t, err)
	// assert.NotEmpty(t, res.Count)
	// assert.NotEmpty(t, res.Items)

	_, err = vkUser.AppsGetFriendsListExtended(map[string]string{})
	assert.NoError(t, err)
	// assert.NotEmpty(t, res2.Count)
	// assert.NotEmpty(t, res2.Items)
}

// TODO: TestVK_AppsGetLeaderboard

func TestVK_AppsGetScopes(t *testing.T) {
	needUserToken(t)

	res, err := vkUser.AppsGetScopes(map[string]string{})
	assert.NoError(t, err)
	assert.NotEmpty(t, res.Count)
	if assert.NotEmpty(t, res.Items) {
		assert.NotEmpty(t, res.Items[0].Name)
		assert.NotEmpty(t, res.Items[0].Title)
	}
}

// TODO: TestVK_AppsGetScore
// TODO: TestVK_AppsSsendRequest
