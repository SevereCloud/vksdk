package params_test

import (
	"testing"

	"github.com/SevereCloud/vksdk/api/params"
	"github.com/stretchr/testify/assert"
)

func TestOrdersCancelSubscriptionBulder(t *testing.T) {
	b := params.NewOrdersCancelSubscriptionBulder()

	b.UserID(1)
	b.SubscriptionID(1)
	b.PendingCancel(true)

	assert.Equal(t, b.Params["user_id"], 1)
	assert.Equal(t, b.Params["subscription_id"], 1)
	assert.Equal(t, b.Params["pending_cancel"], true)
}

func TestOrdersChangeStateBulder(t *testing.T) {
	b := params.NewOrdersChangeStateBulder()

	b.OrderID(1)
	b.Action("text")
	b.AppOrderID(1)
	b.TestMode(true)

	assert.Equal(t, b.Params["order_id"], 1)
	assert.Equal(t, b.Params["action"], "text")
	assert.Equal(t, b.Params["app_order_id"], 1)
	assert.Equal(t, b.Params["test_mode"], true)
}

func TestOrdersGetBulder(t *testing.T) {
	b := params.NewOrdersGetBulder()

	b.Offset(1)
	b.Count(1)
	b.TestMode(true)

	assert.Equal(t, b.Params["offset"], 1)
	assert.Equal(t, b.Params["count"], 1)
	assert.Equal(t, b.Params["test_mode"], true)
}

func TestOrdersGetAmountBulder(t *testing.T) {
	b := params.NewOrdersGetAmountBulder()

	b.UserID(1)
	b.Votes([]string{"text"})

	assert.Equal(t, b.Params["user_id"], 1)
	assert.Equal(t, b.Params["votes"], []string{"text"})
}

func TestOrdersGetByIDBulder(t *testing.T) {
	b := params.NewOrdersGetByIDBulder()

	b.OrderID(1)
	b.OrderIDs([]int{1})
	b.TestMode(true)

	assert.Equal(t, b.Params["order_id"], 1)
	assert.Equal(t, b.Params["order_ids"], []int{1})
	assert.Equal(t, b.Params["test_mode"], true)
}

func TestOrdersGetUserSubscriptionByIDBulder(t *testing.T) {
	b := params.NewOrdersGetUserSubscriptionByIDBulder()

	b.UserID(1)
	b.SubscriptionID(1)

	assert.Equal(t, b.Params["user_id"], 1)
	assert.Equal(t, b.Params["subscription_id"], 1)
}

func TestOrdersGetUserSubscriptionsBulder(t *testing.T) {
	b := params.NewOrdersGetUserSubscriptionsBulder()

	b.UserID(1)

	assert.Equal(t, b.Params["user_id"], 1)
}

func TestOrdersUpdateSubscriptionBulder(t *testing.T) {
	b := params.NewOrdersUpdateSubscriptionBulder()

	b.UserID(1)
	b.SubscriptionID(1)
	b.Price(1)

	assert.Equal(t, b.Params["user_id"], 1)
	assert.Equal(t, b.Params["subscription_id"], 1)
	assert.Equal(t, b.Params["price"], 1)
}
