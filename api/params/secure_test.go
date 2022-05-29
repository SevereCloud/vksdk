package params_test

import (
	"testing"

	"github.com/Derad6709/vksdk/v2/api/params"
	"github.com/stretchr/testify/assert"
)

func TestSecureAddAppEventBuilder(t *testing.T) {
	t.Parallel()

	b := params.NewSecureAddAppEventBuilder()

	b.UserID(1)
	b.ActivityID(1)
	b.Value(1)

	assert.Equal(t, b.Params["user_id"], 1)
	assert.Equal(t, b.Params["activity_id"], 1)
	assert.Equal(t, b.Params["value"], 1)
}

func TestSecureCheckTokenBuilder(t *testing.T) {
	t.Parallel()

	b := params.NewSecureCheckTokenBuilder()

	b.Token("text")
	b.IP("text")

	assert.Equal(t, b.Params["token"], "text")
	assert.Equal(t, b.Params["ip"], "text")
}

func TestSecureGetSMSHistoryBuilder(t *testing.T) {
	t.Parallel()

	b := params.NewSecureGetSMSHistoryBuilder()

	b.UserID(1)
	b.DateFrom(1)
	b.DateTo(1)
	b.Limit(1)

	assert.Equal(t, b.Params["user_id"], 1)
	assert.Equal(t, b.Params["date_from"], 1)
	assert.Equal(t, b.Params["date_to"], 1)
	assert.Equal(t, b.Params["limit"], 1)
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

	assert.Equal(t, b.Params["type"], 1)
	assert.Equal(t, b.Params["uid_from"], 1)
	assert.Equal(t, b.Params["uid_to"], 1)
	assert.Equal(t, b.Params["date_from"], 1)
	assert.Equal(t, b.Params["date_to"], 1)
	assert.Equal(t, b.Params["limit"], 1)
}

func TestSecureGetUserLevelBuilder(t *testing.T) {
	t.Parallel()

	b := params.NewSecureGetUserLevelBuilder()

	b.UserIDs([]int{1})

	assert.Equal(t, b.Params["user_ids"], []int{1})
}

func TestSecureGiveEventStickerBuilder(t *testing.T) {
	t.Parallel()

	b := params.NewSecureGiveEventStickerBuilder()

	b.UserIDs([]int{1})
	b.AchievementID(1)

	assert.Equal(t, b.Params["user_ids"], []int{1})
	assert.Equal(t, b.Params["achievement_id"], 1)
}

func TestSecureSendNotificationBuilder(t *testing.T) {
	t.Parallel()

	b := params.NewSecureSendNotificationBuilder()

	b.UserIDs([]int{1})
	b.UserID(1)
	b.Message("text")

	assert.Equal(t, b.Params["user_ids"], []int{1})
	assert.Equal(t, b.Params["user_id"], 1)
	assert.Equal(t, b.Params["message"], "text")
}

func TestSecureSendSMSNotificationBuilder(t *testing.T) {
	t.Parallel()

	b := params.NewSecureSendSMSNotificationBuilder()

	b.UserID(1)
	b.Message("text")

	assert.Equal(t, b.Params["user_id"], 1)
	assert.Equal(t, b.Params["message"], "text")
}

func TestSecureSetCounterBuilder(t *testing.T) {
	t.Parallel()

	b := params.NewSecureSetCounterBuilder()

	b.Counters([]string{"text"})
	b.UserID(1)
	b.Counter(1)
	b.Increment(true)

	assert.Equal(t, b.Params["counters"], []string{"text"})
	assert.Equal(t, b.Params["user_id"], 1)
	assert.Equal(t, b.Params["counter"], 1)
	assert.Equal(t, b.Params["increment"], true)
}
