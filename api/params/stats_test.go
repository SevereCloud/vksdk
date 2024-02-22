package params_test

import (
	"testing"

	"github.com/SevereCloud/vksdk/v2/api/params"
	"github.com/stretchr/testify/assert"
)

func TestStatsGetBuilder(t *testing.T) {
	t.Parallel()

	b := params.NewStatsGetBuilder()

	b.GroupID(1)
	b.AppID(1)
	b.TimestampFrom(1)
	b.TimestampTo(1)
	b.Interval("text")
	b.IntervalsCount(1)
	b.Filters([]string{"text"})
	b.StatsGroups([]string{"text"})
	b.Extended(true)

	assert.Equal(t, 1, b.Params["group_id"])
	assert.Equal(t, 1, b.Params["app_id"])
	assert.Equal(t, 1, b.Params["timestamp_from"])
	assert.Equal(t, 1, b.Params["timestamp_to"])
	assert.Equal(t, "text", b.Params["interval"])
	assert.Equal(t, 1, b.Params["intervals_count"])
	assert.Equal(t, []string{"text"}, b.Params["filters"])
	assert.Equal(t, []string{"text"}, b.Params["stats_groups"])
	assert.Equal(t, true, b.Params["extended"])
}

func TestStatsGetPostReachBuilder(t *testing.T) {
	t.Parallel()

	b := params.NewStatsGetPostReachBuilder()

	b.OwnerID("text")
	b.PostID(1)

	assert.Equal(t, "text", b.Params["owner_id"])
	assert.Equal(t, 1, b.Params["post_id"])
}

func TestStatsTrackVisitorBuilder(t *testing.T) {
	t.Parallel()

	b := params.NewStatsTrackVisitorBuilder()

	b.ID("text")

	assert.Equal(t, "text", b.Params["id"])
}
