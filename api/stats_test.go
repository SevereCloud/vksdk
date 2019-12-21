package api_test

import (
	"testing"

	"github.com/SevereCloud/vksdk/api"
)

func TestVK_StatsGet(t *testing.T) {
	t.Skip("See https://vk.com/bug136096")
	needUserToken(t)

	_, err := vkUser.StatsGet(api.Params{
		"group_id":        vkGroupID,
		"interval":        "day",
		"intervals_count": 10,
		"extended":        true,
	})
	noError(t, err)
}

func TestVK_StatsTrackVisitor(t *testing.T) {
	needUserToken(t)

	_, err := vkUser.StatsTrackVisitor(api.Params{})
	noError(t, err)
}
