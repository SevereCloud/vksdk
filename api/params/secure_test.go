package params_test

import (
	"testing"

	"github.com/SevereCloud/vksdk/api/params"
	"github.com/stretchr/testify/assert"
)

func TestSecureAddAppEventBulder(t *testing.T) {
	b := params.NewSecureAddAppEventBulder()

	b.UserID(1)
	b.ActivityID(1)
	b.Value(1)

	assert.Equal(t, b.Params["user_id"], 1)
	assert.Equal(t, b.Params["activity_id"], 1)
	assert.Equal(t, b.Params["value"], 1)
}

func TestSecureCheckTokenBulder(t *testing.T) {
	b := params.NewSecureCheckTokenBulder()

	b.Token("text")
	b.IP("text")

	assert.Equal(t, b.Params["token"], "text")
	assert.Equal(t, b.Params["ip"], "text")
}

func TestSecureGetSMSHistoryBulder(t *testing.T) {
	b := params.NewSecureGetSMSHistoryBulder()

	b.UserID(1)
	b.DateFrom(1)
	b.DateTo(1)
	b.Limit(1)

	assert.Equal(t, b.Params["user_id"], 1)
	assert.Equal(t, b.Params["date_from"], 1)
	assert.Equal(t, b.Params["date_to"], 1)
	assert.Equal(t, b.Params["limit"], 1)
}

func TestSecureGetTransactionsHistoryBulder(t *testing.T) {
	b := params.NewSecureGetTransactionsHistoryBulder()

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

func TestSecureGetUserLevelBulder(t *testing.T) {
	b := params.NewSecureGetUserLevelBulder()

	b.UserIDs([]int{1})

	assert.Equal(t, b.Params["user_ids"], []int{1})
}

func TestSecureGiveEventStickerBulder(t *testing.T) {
	b := params.NewSecureGiveEventStickerBulder()

	b.UserIDs([]int{1})
	b.AchievementID(1)

	assert.Equal(t, b.Params["user_ids"], []int{1})
	assert.Equal(t, b.Params["achievement_id"], 1)
}

func TestSecureSendNotificationBulder(t *testing.T) {
	b := params.NewSecureSendNotificationBulder()

	b.UserIDs([]int{1})
	b.UserID(1)
	b.Message("text")

	assert.Equal(t, b.Params["user_ids"], []int{1})
	assert.Equal(t, b.Params["user_id"], 1)
	assert.Equal(t, b.Params["message"], "text")
}

func TestSecureSendSMSNotificationBulder(t *testing.T) {
	b := params.NewSecureSendSMSNotificationBulder()

	b.UserID(1)
	b.Message("text")

	assert.Equal(t, b.Params["user_id"], 1)
	assert.Equal(t, b.Params["message"], "text")
}

func TestSecureSetCounterBulder(t *testing.T) {
	b := params.NewSecureSetCounterBulder()

	b.Counters([]string{"text"})
	b.UserID(1)
	b.Counter(1)
	b.Increment(true)

	assert.Equal(t, b.Params["counters"], []string{"text"})
	assert.Equal(t, b.Params["user_id"], 1)
	assert.Equal(t, b.Params["counter"], 1)
	assert.Equal(t, b.Params["increment"], true)
}
