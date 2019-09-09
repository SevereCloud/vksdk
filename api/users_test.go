package api

import (
	"testing"
)

func TestVK_UsersGet(t *testing.T) {
	if vkUser.AccessToken == "" {
		t.Skip("USER_TOKEN empty")
	}

	_, gotVkErr := vkUser.UsersGet(map[string]string{
		"user_ids": "1",
		"fields":   "photo_id, verified, sex, bdate, city, country, home_town, has_photo, photo_50, photo_100, photo_200_orig, photo_200, photo_400_orig, photo_max, photo_max_orig, online, domain, has_mobile, contacts, site, education, universities, schools, status, last_seen, followers_count, common_count, occupation, nickname, relatives, relation, personal, connections, exports, activities, interests, music, movies, tv, books, games, about, quotes, can_post, can_see_all_posts, can_see_audio, can_write_private_message, can_send_friend_request, is_favorite, is_hidden_from_feed, timezone, screen_name, maiden_name, crop_photo, is_friend, friend_status, career, military, blacklisted, blacklisted_by_me, can_be_invited_group",
	})
	if gotVkErr.Code != 0 {
		t.Errorf("VK.UsersGet() gotVkErr = %v", gotVkErr)
	}
}

func TestVK_UsersGetFollowers(t *testing.T) {
	if vkUser.AccessToken == "" {
		t.Skip("USER_TOKEN empty")
	}

	_, gotVkErr := vkUser.UsersGetFollowers(map[string]string{
		"user_id": "1",
	})
	if gotVkErr.Code != 0 {
		t.Errorf("VK.UsersGetFollowers() gotVkErr = %v", gotVkErr)
	}

	_, gotVkErr = vkUser.UsersGetFollowersFields(map[string]string{
		"user_id": "1",
	})
	if gotVkErr.Code != 0 {
		t.Errorf("VK.UsersGetFollowersFields() gotVkErr = %v", gotVkErr)
	}
}

func TestVK_UsersGetSubscriptions(t *testing.T) {
	if vkUser.AccessToken == "" {
		t.Skip("USER_TOKEN empty")
	}

	_, gotVkErr := vkUser.UsersGetSubscriptions(map[string]string{
		"user_id": "117253521",
	})
	if gotVkErr.Code != 0 {
		t.Errorf("VK.UsersGetSubscriptions() gotVkErr = %v", gotVkErr)
	}
}

func TestVK_UsersIsAppUser(t *testing.T) {
	if vkUser.AccessToken == "" {
		t.Skip("USER_TOKEN empty")
	}

	_, gotVkErr := vkUser.UsersIsAppUser(map[string]string{
		"user_id": "1",
	})
	if gotVkErr.Code != 0 {
		t.Errorf("VK.UsersIsAppUser() gotVkErr = %v", gotVkErr)
	}
}

func TestVK_UsersReport(t *testing.T) {
	if vkUser.AccessToken == "" {
		t.Skip("USER_TOKEN empty")
	}

	_, gotVkErr := vkUser.UsersReport(map[string]string{
		"user_id": "1",
		"type":    "spam",
		"comment": "Тестовый репорт - github.com/SevereCloud/vksdk",
	})
	if gotVkErr.Code == 15 {
		t.Skip("Access denied")
	}
	if gotVkErr.Code != 0 {
		t.Errorf("VK.UsersReport() gotVkErr = %v", gotVkErr)
	}
}

func TestVK_UsersSearch(t *testing.T) {
	if vkUser.AccessToken == "" {
		t.Skip("USER_TOKEN empty")
	}

	_, gotVkErr := vkUser.UsersSearch(map[string]string{
		"q":      "Vasya Babich",
		"fields": "photo_id, verified, sex, bdate, city, country, home_town, has_photo, photo_50, photo_100, photo_200_orig, photo_200, photo_400_orig, photo_max, photo_max_orig, online, domain, has_mobile, contacts, site, education, universities, schools, status, last_seen, followers_count, common_count, occupation, nickname, relatives, relation, personal, connections, exports, activities, interests, music, movies, tv, books, games, about, quotes, can_post, can_see_all_posts, can_see_audio, can_write_private_message, can_send_friend_request, is_favorite, is_hidden_from_feed, timezone, screen_name, maiden_name, crop_photo, is_friend, friend_status, career, military, blacklisted, blacklisted_by_me, can_be_invited_group",
	})
	if gotVkErr.Code != 0 {
		t.Errorf("VK.UsersSearch() gotVkErr = %v", gotVkErr)
	}
}
