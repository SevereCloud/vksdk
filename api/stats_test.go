package api

import (
	"strconv"
	"testing"
)

func TestVK_StatsGet(t *testing.T) {
	t.Skip("See https://vk.com/bug136096")
	if vkUser.AccessToken == "" {
		t.Skip("USER_TOKEN empty")
	}

	_, gotVkErr := vkUser.StatsGet(map[string]string{
		"group_id":        strconv.Itoa(vkGroupID),
		"interval":        "day",
		"intervals_count": "10",
		"extended":        "1",
	})
	if gotVkErr.Code != 0 {
		t.Errorf("VK.StatsGet() gotVkErr = %v, want %v", gotVkErr, 0)
	}
}

func TestVK_StatsTrackVisitor(t *testing.T) {
	if vkUser.AccessToken == "" {
		t.Skip("USER_TOKEN empty")
	}

	_, gotVkErr := vkUser.StatsTrackVisitor(map[string]string{})
	if gotVkErr.Code != 0 {
		t.Errorf("VK.StatsTrackVisitor() gotVkErr = %v, want %v", gotVkErr, 0)
	}
}
