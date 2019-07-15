package api

import (
	"testing"
)

// BUG(VK): https://vk.com/bug136096
// func TestVK_StatsGet(t *testing.T) {
// 	if vkUser.AccessToken == "" {
// 		t.Skip("USER_TOKEN empty")
// 	}

// 	t.Run("StatsGet", func(t *testing.T) {
// 		_, gotVkErr := vkUser.StatsGet(map[string]string{
// 			"group_id":        strconv.Itoa(vkGroupID),
// 			"interval":        "day",
// 			"intervals_count": "10",
// 			"extended":        "1",
// 		})
// 		if gotVkErr.Code != 0 {
// 			t.Errorf("VK.StatsGet() gotVkErr = %v, want %v", gotVkErr, 0)
// 		}
// 	})
// }

func TestVK_StatsTrackVisitor(t *testing.T) {
	if vkUser.AccessToken == "" {
		t.Skip("USER_TOKEN empty")
	}

	t.Run("StatsTrackVisitor", func(t *testing.T) {
		_, gotVkErr := vkUser.StatsTrackVisitor(map[string]string{})
		if gotVkErr.Code != 0 {
			t.Errorf("VK.StatsTrackVisitor() gotVkErr = %v, want %v", gotVkErr, 0)
		}
	})
}
