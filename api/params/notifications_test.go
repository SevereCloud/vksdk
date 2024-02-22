package params_test

import (
	"testing"

	"github.com/SevereCloud/vksdk/v2/api/params"
	"github.com/stretchr/testify/assert"
)

func TestNotificationsGetBuilder(t *testing.T) {
	t.Parallel()

	b := params.NewNotificationsGetBuilder()

	b.Count(1)
	b.StartFrom("text")
	b.Filters([]string{"text"})
	b.StartTime(1)
	b.EndTime(1)

	assert.Equal(t, 1, b.Params["count"])
	assert.Equal(t, "text", b.Params["start_from"])
	assert.Equal(t, []string{"text"}, b.Params["filters"])
	assert.Equal(t, 1, b.Params["start_time"])
	assert.Equal(t, 1, b.Params["end_time"])
}

func TestNotificationsSendMessageBuilder(t *testing.T) {
	t.Parallel()

	b := params.NewNotificationsSendMessageBuilder()

	b.UserIDs([]int{1})
	b.Message("text")
	b.Fragment("text")
	b.GroupID(1)
	b.RandomID(1)

	assert.Equal(t, []int{1}, b.Params["user_ids"])
	assert.Equal(t, "text", b.Params["message"])
	assert.Equal(t, "text", b.Params["fragment"])
	assert.Equal(t, 1, b.Params["group_id"])
	assert.Equal(t, int64(1), b.Params["random_id"])
}
