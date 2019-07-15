package api

import (
	"testing"
)

func TestVK_LikesAdd(t *testing.T) {
	if vkUser.AccessToken == "" {
		t.Skip("USER_TOKEN empty")
	}

	t.Run("LikesAdd", func(t *testing.T) {
		_, gotVkErr := vkUser.LikesAdd(map[string]string{
			"type":     "post",
			"owner_id": "1",
			"item_id":  "45546",
		})
		if gotVkErr.Code != 0 {
			t.Errorf("VK.LikesAdd() gotVkErr = %v, want %v", gotVkErr, 0)
		}
	})
}

func TestVK_LikesDelete(t *testing.T) {
	if vkUser.AccessToken == "" {
		t.Skip("USER_TOKEN empty")
	}

	t.Run("LikesDelete", func(t *testing.T) {
		_, gotVkErr := vkUser.LikesDelete(map[string]string{
			"type":     "post",
			"owner_id": "1",
			"item_id":  "45546",
		})
		if gotVkErr.Code != 0 {
			t.Errorf("VK.LikesDelete() gotVkErr = %v, want %v", gotVkErr, 0)
		}
	})
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

	t.Run("LikesGetList", func(t *testing.T) {
		_, gotVkErr := vkUser.LikesGetList(params)
		if gotVkErr.Code != 0 {
			t.Errorf("VK.LikesGetList() gotVkErr = %v, want %v", gotVkErr, 0)
		}
	})

	t.Run("LikesGetListExtended", func(t *testing.T) {
		_, gotVkErr := vkUser.LikesGetListExtended(params)
		if gotVkErr.Code != 0 {
			t.Errorf("VK.LikesGetListExtended() gotVkErr = %v, want %v", gotVkErr, 0)
		}
	})
}

func TestVK_LikesIsLiked(t *testing.T) {
	if vkUser.AccessToken == "" {
		t.Skip("USER_TOKEN empty")
	}

	t.Run("LikesIsLiked", func(t *testing.T) {
		_, gotVkErr := vkUser.LikesIsLiked(map[string]string{
			"type":     "post",
			"owner_id": "1",
			"item_id":  "45546",
		})
		if gotVkErr.Code != 0 {
			t.Errorf("VK.LikesIsLiked() gotVkErr = %v, want %v", gotVkErr, 0)
		}
	})
}
