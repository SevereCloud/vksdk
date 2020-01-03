package params_test

import (
	"testing"

	"github.com/SevereCloud/vksdk/api/params"
	"github.com/stretchr/testify/assert"
)

func TestNotificationsGetBuilder(t *testing.T) {
	b := params.NewNotificationsGetBuilder()

	b.Count(1)
	b.StartFrom("text")
	b.Filters([]string{"text"})
	b.StartTime(1)
	b.EndTime(1)

	assert.Equal(t, b.Params["count"], 1)
	assert.Equal(t, b.Params["start_from"], "text")
	assert.Equal(t, b.Params["filters"], []string{"text"})
	assert.Equal(t, b.Params["start_time"], 1)
	assert.Equal(t, b.Params["end_time"], 1)
}

func TestNotificationsSendMessageBuilder(t *testing.T) {
	b := params.NewNotificationsSendMessageBuilder()

	b.UserIDs([]int{1})
	b.Message("text")
	b.Fragment("text")
	b.GroupID(1)

	assert.Equal(t, b.Params["user_ids"], []int{1})
	assert.Equal(t, b.Params["message"], "text")
	assert.Equal(t, b.Params["fragment"], "text")
	assert.Equal(t, b.Params["group_id"], 1)
}
