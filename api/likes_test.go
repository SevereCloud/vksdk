package api

import (
	"testing"
)

func TestVK_LikesAdd(t *testing.T) {
	needUserToken(t)

	_, err := vkUser.LikesAdd(map[string]string{
		"type":     "post",
		"owner_id": "1",
		"item_id":  "45546",
	})
	if err != nil {
		t.Errorf("VK.LikesAdd() err = %v", err)
	}
}

func TestVK_LikesDelete(t *testing.T) {
	needUserToken(t)

	_, err := vkUser.LikesDelete(map[string]string{
		"type":     "post",
		"owner_id": "1",
		"item_id":  "45546",
	})
	if err != nil {
		t.Errorf("VK.LikesDelete() err = %v", err)
	}
}

func TestVK_LikesGetList(t *testing.T) {
	needUserToken(t)

	params := map[string]string{
		"type":     "post",
		"owner_id": "1",
		"item_id":  "45546",
	}

	_, err := vkUser.LikesGetList(params)
	if err != nil {
		t.Errorf("VK.LikesGetList() err = %v", err)
	}

	_, err = vkUser.LikesGetListExtended(params)
	if err != nil {
		t.Errorf("VK.LikesGetListExtended() err = %v", err)
	}
}

func TestVK_LikesIsLiked(t *testing.T) {
	needUserToken(t)

	_, err := vkUser.LikesIsLiked(map[string]string{
		"type":     "post",
		"owner_id": "1",
		"item_id":  "45546",
	})
	if err != nil {
		t.Errorf("VK.LikesIsLiked() err = %v", err)
	}
}
