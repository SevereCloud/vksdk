package api

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestVK_StatsGet(t *testing.T) {
	t.Skip("See https://vk.com/bug136096")
	needUserToken(t)

	_, err := vkUser.StatsGet(Params{
		"group_id":        vkGroupID,
		"interval":        "day",
		"intervals_count": 10,
		"extended":        true,
	})
	assert.NoError(t, err)
}

func TestVK_StatsTrackVisitor(t *testing.T) {
	needUserToken(t)

	_, err := vkUser.StatsTrackVisitor(Params{})
	assert.NoError(t, err)
}
