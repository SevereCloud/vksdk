package params_test

import (
	"testing"

	"github.com/SevereCloud/vksdk/v3/api/params"
	"github.com/stretchr/testify/assert"
)

func TestSecureAddAppEventBuilder(t *testing.T) {
	t.Parallel()

	b := params.NewSecureAddAppEventBuilder()

	b.UserID(1)
	b.ActivityID(1)
	b.Value(1)

	assert.Equal(t, 1, b.Params["user_id"])
	assert.Equal(t, 1, b.Params["activity_id"])
	assert.Equal(t, 1, b.Params["value"])
}

func TestSecureCheckTokenBuilder(t *testing.T) {
	t.Parallel()

	b := params.NewSecureCheckTokenBuilder()

	b.Token("text")
	b.IP("text")

	assert.Equal(t, "text", b.Params["token"])
	assert.Equal(t, "text", b.Params["ip"])
}

func TestSecureGetSMSHistoryBuilder(t *testing.T) {
	t.Parallel()

	b := params.NewSecureGetSMSHistoryBuilder()

	b.UserID(1)
	b.DateFrom(1)
	b.DateTo(1)
	b.Limit(1)

	assert.Equal(t, 1, b.Params["user_id"])
	assert.Equal(t, 1, b.Params["date_from"])
	assert.Equal(t, 1, b.Params["date_to"])
	assert.Equal(t, 1, b.Params["limit"])
}

func TestSecureGetTransactionsHistoryBuilder(t *testing.T) {
	t.Parallel()

	b := params.NewSecureGetTransactionsHistoryBuilder()

	b.Type(1)
	b.UIDFrom(1)
	b.UIDTo(1)
	b.DateFrom(1)
	b.DateTo(1)
	b.Limit(1)

	assert.Equal(t, 1, b.Params["type"])
	assert.Equal(t, 1, b.Params["uid_from"])
	assert.Equal(t, 1, b.Params["uid_to"])
	assert.Equal(t, 1, b.Params["date_from"])
	assert.Equal(t, 1, b.Params["date_to"])
	assert.Equal(t, 1, b.Params["limit"])
}

func TestSecureGetUserLevelBuilder(t *testing.T) {
	t.Parallel()

	b := params.NewSecureGetUserLevelBuilder()

	b.UserIDs([]int{1})

	assert.Equal(t, []int{1}, b.Params["user_ids"])
}

func TestSecureGiveEventStickerBuilder(t *testing.T) {
	t.Parallel()

	b := params.NewSecureGiveEventStickerBuilder()

	b.UserIDs([]int{1})
	b.AchievementID(1)

	assert.Equal(t, []int{1}, b.Params["user_ids"])
	assert.Equal(t, 1, b.Params["achievement_id"])
}

func TestSecureSendNotificationBuilder(t *testing.T) {
	t.Parallel()

	b := params.NewSecureSendNotificationBuilder()

	b.UserIDs([]int{1})
	b.UserID(1)
	b.Message("text")

	assert.Equal(t, []int{1}, b.Params["user_ids"])
	assert.Equal(t, 1, b.Params["user_id"])
	assert.Equal(t, "text", b.Params["message"])
}

func TestSecureSendSMSNotificationBuilder(t *testing.T) {
	t.Parallel()

	b := params.NewSecureSendSMSNotificationBuilder()

	b.UserID(1)
	b.Message("text")

	assert.Equal(t, 1, b.Params["user_id"])
	assert.Equal(t, "text", b.Params["message"])
}

func TestSecureSetCounterBuilder(t *testing.T) {
	t.Parallel()

	b := params.NewSecureSetCounterBuilder()

	b.Counters([]string{"text"})
	b.UserID(1)
	b.Counter(1)
	b.Increment(true)

	assert.Equal(t, []string{"text"}, b.Params["counters"])
	assert.Equal(t, 1, b.Params["user_id"])
	assert.Equal(t, 1, b.Params["counter"])
	assert.Equal(t, true, b.Params["increment"])
}
