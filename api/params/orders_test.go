package params_test

import (
	"testing"

	"github.com/SevereCloud/vksdk/v2/api/params"
	"github.com/stretchr/testify/assert"
)

func TestOrdersCancelSubscriptionBuilder(t *testing.T) {
	t.Parallel()

	b := params.NewOrdersCancelSubscriptionBuilder()

	b.UserID(1)
	b.SubscriptionID(1)
	b.PendingCancel(true)

	assert.Equal(t, b.Params["user_id"], 1)
	assert.Equal(t, b.Params["subscription_id"], 1)
	assert.Equal(t, b.Params["pending_cancel"], true)
}

func TestOrdersChangeStateBuilder(t *testing.T) {
	t.Parallel()

	b := params.NewOrdersChangeStateBuilder()

	b.OrderID(1)
	b.Action("text")
	b.AppOrderID(1)
	b.TestMode(true)

	assert.Equal(t, b.Params["order_id"], 1)
	assert.Equal(t, b.Params["action"], "text")
	assert.Equal(t, b.Params["app_order_id"], 1)
	assert.Equal(t, b.Params["test_mode"], true)
}

func TestOrdersGetBuilder(t *testing.T) {
	t.Parallel()

	b := params.NewOrdersGetBuilder()

	b.Offset(1)
	b.Count(1)
	b.TestMode(true)

	assert.Equal(t, b.Params["offset"], 1)
	assert.Equal(t, b.Params["count"], 1)
	assert.Equal(t, b.Params["test_mode"], true)
}

func TestOrdersGetAmountBuilder(t *testing.T) {
	t.Parallel()

	b := params.NewOrdersGetAmountBuilder()

	b.UserID(1)
	b.Votes([]string{"text"})

	assert.Equal(t, b.Params["user_id"], 1)
	assert.Equal(t, b.Params["votes"], []string{"text"})
}

func TestOrdersGetByIDBuilder(t *testing.T) {
	t.Parallel()

	b := params.NewOrdersGetByIDBuilder()

	b.OrderID(1)
	b.OrderIDs([]int{1})
	b.TestMode(true)

	assert.Equal(t, b.Params["order_id"], 1)
	assert.Equal(t, b.Params["order_ids"], []int{1})
	assert.Equal(t, b.Params["test_mode"], true)
}

func TestOrdersGetUserSubscriptionByIDBuilder(t *testing.T) {
	t.Parallel()

	b := params.NewOrdersGetUserSubscriptionByIDBuilder()

	b.UserID(1)
	b.SubscriptionID(1)

	assert.Equal(t, b.Params["user_id"], 1)
	assert.Equal(t, b.Params["subscription_id"], 1)
}

func TestOrdersGetUserSubscriptionsBuilder(t *testing.T) {
	t.Parallel()

	b := params.NewOrdersGetUserSubscriptionsBuilder()

	b.UserID(1)

	assert.Equal(t, b.Params["user_id"], 1)
}

func TestOrdersUpdateSubscriptionBuilder(t *testing.T) {
	t.Parallel()

	b := params.NewOrdersUpdateSubscriptionBuilder()

	b.UserID(1)
	b.SubscriptionID(1)
	b.Price(1)

	assert.Equal(t, b.Params["user_id"], 1)
	assert.Equal(t, b.Params["subscription_id"], 1)
	assert.Equal(t, b.Params["price"], 1)
}
