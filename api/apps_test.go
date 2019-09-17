package api

import (
	"testing"
)

func TestVK_AppsDeleteAppRequests(t *testing.T) {
	needUserToken(t)

	_, err := vkUser.AppsDeleteAppRequests(map[string]string{})
	if err != nil {
		t.Errorf("VK.AppsDeleteAppRequests() err = %v", err)
	}
}

func TestVK_AppsGet(t *testing.T) {
	needServiceToken(t)

	_, err := vkService.AppsGet(map[string]string{
		"app_id":   "4063926",
		"extended": "1",
	})
	if err != nil {
		t.Errorf("VK.AppsGet() err = %v", err)
	}
}

func TestVK_AppsGetCatalog(t *testing.T) {
	needServiceToken(t)

	_, err := vkService.AppsGetCatalog(map[string]string{})
	if err != nil {
		t.Errorf("VK.AppsGetCatalog() err = %v", err)
	}
}

func TestVK_AppsGetFriendsList(t *testing.T) {
	needUserToken(t)

	_, err := vkUser.AppsGetFriendsList(map[string]string{})
	if err != nil {
		t.Errorf("VK.AppsGetFriendsList() err = %v", err)
	}

	_, err = vkUser.AppsGetFriendsListExtended(map[string]string{})
	if err != nil {
		t.Errorf("VK.AppsGetFriendsList() err = %v", err)
	}
}

// TODO: TestVK_AppsGetLeaderboard

func TestVK_AppsGetScopes(t *testing.T) {
	needUserToken(t)

	_, err := vkUser.AppsGetScopes(map[string]string{})
	if err != nil {
		t.Errorf("VK.AppsGetScopes() err = %v", err)
	}
}

// TODO: TestVK_AppsGetScore
// TODO: TestVK_AppsSsendRequest
