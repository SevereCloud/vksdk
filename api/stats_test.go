package api

import (
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestVK_StatsGet(t *testing.T) {
	t.Skip("See https://vk.com/bug136096")
	needUserToken(t)

	_, err := vkUser.StatsGet(map[string]string{
		"group_id":        strconv.Itoa(vkGroupID),
		"interval":        "day",
		"intervals_count": "10",
		"extended":        "1",
	})
	assert.NoError(t, err)
}

func TestVK_StatsTrackVisitor(t *testing.T) {
	needUserToken(t)

	_, err := vkUser.StatsTrackVisitor(map[string]string{})
	assert.NoError(t, err)
}
