package api_test

import (
	"testing"

	"github.com/Derad6709/vksdk/v2/api"
)

func TestVK_StatsGet(t *testing.T) {
	t.Parallel()

	needUserToken(t)
	needGroupToken(t)

	_, err := vkUser.StatsGet(api.Params{
		"group_id":        vkGroupID,
		"interval":        "day",
		"intervals_count": 10,
		"extended":        true,
	})
	noError(t, err)
}

func TestVK_StatsTrackVisitor(t *testing.T) {
	t.Parallel()

	needUserToken(t)

	_, err := vkUser.StatsTrackVisitor(nil)
	noError(t, err)
}
