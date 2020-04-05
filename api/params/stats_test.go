package params_test

import (
	"testing"

	"github.com/SevereCloud/vksdk/api/params"
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

	assert.Equal(t, b.Params["group_id"], 1)
	assert.Equal(t, b.Params["app_id"], 1)
	assert.Equal(t, b.Params["timestamp_from"], 1)
	assert.Equal(t, b.Params["timestamp_to"], 1)
	assert.Equal(t, b.Params["interval"], "text")
	assert.Equal(t, b.Params["intervals_count"], 1)
	assert.Equal(t, b.Params["filters"], []string{"text"})
	assert.Equal(t, b.Params["stats_groups"], []string{"text"})
	assert.Equal(t, b.Params["extended"], true)
}

func TestStatsGetPostReachBuilder(t *testing.T) {
	t.Parallel()

	b := params.NewStatsGetPostReachBuilder()

	b.OwnerID("text")
	b.PostID(1)

	assert.Equal(t, b.Params["owner_id"], "text")
	assert.Equal(t, b.Params["post_id"], 1)
}

func TestStatsTrackVisitorBuilder(t *testing.T) {
	t.Parallel()

	b := params.NewStatsTrackVisitorBuilder()

	b.ID("text")

	assert.Equal(t, b.Params["id"], "text")
}
