package api

import (
	"testing"
)

func TestVK_LikesAdd(t *testing.T) {
	if vkUser.AccessToken == "" {
		t.Skip("USER_TOKEN empty")
	}

	_, gotVkErr := vkUser.LikesAdd(map[string]string{
		"type":     "post",
		"owner_id": "1",
		"item_id":  "45546",
	})
	if gotVkErr.Code != 0 {
		t.Errorf("VK.LikesAdd() gotVkErr = %v", gotVkErr)
	}
}

func TestVK_LikesDelete(t *testing.T) {
	if vkUser.AccessToken == "" {
		t.Skip("USER_TOKEN empty")
	}

	_, gotVkErr := vkUser.LikesDelete(map[string]string{
		"type":     "post",
		"owner_id": "1",
		"item_id":  "45546",
	})
	if gotVkErr.Code != 0 {
		t.Errorf("VK.LikesDelete() gotVkErr = %v", gotVkErr)
	}
}

func TestVK_LikesGetList(t *testing.T) {
	if vkUser.AccessToken == "" {
		t.Skip("USER_TOKEN empty")
	}

	params := map[string]string{
		"type":     "post",
		"owner_id": "1",
		"item_id":  "45546",
	}

	_, gotVkErr := vkUser.LikesGetList(params)
	if gotVkErr.Code != 0 {
		t.Errorf("VK.LikesGetList() gotVkErr = %v", gotVkErr)
	}

	_, gotVkErr = vkUser.LikesGetListExtended(params)
	if gotVkErr.Code != 0 {
		t.Errorf("VK.LikesGetListExtended() gotVkErr = %v", gotVkErr)
	}
}

func TestVK_LikesIsLiked(t *testing.T) {
	if vkUser.AccessToken == "" {
		t.Skip("USER_TOKEN empty")
	}

	_, gotVkErr := vkUser.LikesIsLiked(map[string]string{
		"type":     "post",
		"owner_id": "1",
		"item_id":  "45546",
	})
	if gotVkErr.Code != 0 {
		t.Errorf("VK.LikesIsLiked() gotVkErr = %v", gotVkErr)
	}
}
