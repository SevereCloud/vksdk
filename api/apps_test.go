package api

import (
	"testing"
)

func TestVK_AppsDeleteAppRequests(t *testing.T) {
	if vkUser.AccessToken == "" {
		t.Skip("USER_TOKEN empty")
	}

	_, gotVkErr := vkUser.AppsDeleteAppRequests(map[string]string{})
	if gotVkErr.Code != 0 {
		t.Errorf("VK.AppsDeleteAppRequests() gotVkErr = %v", gotVkErr)
	}
}

func TestVK_AppsGet(t *testing.T) {
	if vkService.AccessToken == "" {
		t.Skip("USER_TOKEN empty")
	}

	_, gotVkErr := vkService.AppsGet(map[string]string{
		"app_id":   "4063926",
		"extended": "1",
	})
	if gotVkErr.Code != 0 {
		t.Errorf("VK.AppsGet() gotVkErr = %v", gotVkErr)
	}
}

func TestVK_AppsGetCatalog(t *testing.T) {
	if vkService.AccessToken == "" {
		t.Skip("SERVICE_TOKEN empty")
	}

	_, gotVkErr := vkService.AppsGetCatalog(map[string]string{})
	if gotVkErr.Code != 0 {
		t.Errorf("VK.AppsGetCatalog() gotVkErr = %v", gotVkErr)
	}
}

func TestVK_AppsGetFriendsList(t *testing.T) {
	if vkUser.AccessToken == "" {
		t.Skip("USER_TOKEN empty")
	}

	_, gotVkErr := vkUser.AppsGetFriendsList(map[string]string{})
	if gotVkErr.Code != 0 {
		t.Errorf("VK.AppsGetFriendsList() gotVkErr = %v", gotVkErr)
	}

	_, gotVkErr = vkUser.AppsGetFriendsListExtended(map[string]string{})
	if gotVkErr.Code != 0 {
		t.Errorf("VK.AppsGetFriendsList() gotVkErr = %v", gotVkErr)
	}
}

// TODO: TestVK_AppsGetLeaderboard

func TestVK_AppsGetScopes(t *testing.T) {
	if vkUser.AccessToken == "" {
		t.Skip("USER_TOKEN empty")
	}

	_, gotVkErr := vkUser.AppsGetScopes(map[string]string{})
	if gotVkErr.Code != 0 {
		t.Errorf("VK.AppsGetScopes() gotVkErr = %v", gotVkErr)
	}
}

// TODO: TestVK_AppsGetScore
// TODO: TestVK_AppsSsendRequest
