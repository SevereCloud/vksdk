package api

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestVK_AppsDeleteAppRequests(t *testing.T) {
	needUserToken(t)

	_, err := vkUser.AppsDeleteAppRequests(map[string]string{})
	assert.NoError(t, err)
}

func TestVK_AppsGet(t *testing.T) {
	needServiceToken(t)

	_, err := vkService.AppsGet(map[string]string{
		"app_id":   "4063926",
		"extended": "1",
	})
	assert.NoError(t, err)
}

func TestVK_AppsGetCatalog(t *testing.T) {
	needServiceToken(t)

	_, err := vkService.AppsGetCatalog(map[string]string{})
	assert.NoError(t, err)
}

func TestVK_AppsGetFriendsList(t *testing.T) {
	needUserToken(t)

	_, err := vkUser.AppsGetFriendsList(map[string]string{})
	assert.NoError(t, err)

	_, err = vkUser.AppsGetFriendsListExtended(map[string]string{})
	assert.NoError(t, err)
}

// TODO: TestVK_AppsGetLeaderboard

func TestVK_AppsGetScopes(t *testing.T) {
	needUserToken(t)

	_, err := vkUser.AppsGetScopes(map[string]string{})
	assert.NoError(t, err)
}

// TODO: TestVK_AppsGetScore
// TODO: TestVK_AppsSsendRequest
