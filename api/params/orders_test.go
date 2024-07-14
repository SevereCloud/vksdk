package params_test

import (
	"testing"

	"github.com/SevereCloud/vksdk/v3/api/params"
	"github.com/stretchr/testify/assert"
)

func TestOrdersCancelSubscriptionBuilder(t *testing.T) {
	t.Parallel()

	b := params.NewOrdersCancelSubscriptionBuilder()

	b.UserID(1)
	b.SubscriptionID(1)
	b.PendingCancel(true)

	assert.Equal(t, 1, b.Params["user_id"])
	assert.Equal(t, 1, b.Params["subscription_id"])
	assert.Equal(t, true, b.Params["pending_cancel"])
}

func TestOrdersChangeStateBuilder(t *testing.T) {
	t.Parallel()

	b := params.NewOrdersChangeStateBuilder()

	b.OrderID(1)
	b.Action("text")
	b.AppOrderID(1)
	b.TestMode(true)

	assert.Equal(t, 1, b.Params["order_id"])
	assert.Equal(t, "text", b.Params["action"])
	assert.Equal(t, 1, b.Params["app_order_id"])
	assert.Equal(t, true, b.Params["test_mode"])
}

func TestOrdersGetBuilder(t *testing.T) {
	t.Parallel()

	b := params.NewOrdersGetBuilder()

	b.Offset(1)
	b.Count(1)
	b.TestMode(true)

	assert.Equal(t, 1, b.Params["offset"])
	assert.Equal(t, 1, b.Params["count"])
	assert.Equal(t, true, b.Params["test_mode"])
}

func TestOrdersGetAmountBuilder(t *testing.T) {
	t.Parallel()

	b := params.NewOrdersGetAmountBuilder()

	b.UserID(1)
	b.Votes([]string{"text"})

	assert.Equal(t, 1, b.Params["user_id"])
	assert.Equal(t, []string{"text"}, b.Params["votes"])
}

func TestOrdersGetByIDBuilder(t *testing.T) {
	t.Parallel()

	b := params.NewOrdersGetByIDBuilder()

	b.OrderID(1)
	b.OrderIDs([]int{1})
	b.TestMode(true)

	assert.Equal(t, 1, b.Params["order_id"])
	assert.Equal(t, []int{1}, b.Params["order_ids"])
	assert.Equal(t, true, b.Params["test_mode"])
}

func TestOrdersGetUserSubscriptionByIDBuilder(t *testing.T) {
	t.Parallel()

	b := params.NewOrdersGetUserSubscriptionByIDBuilder()

	b.UserID(1)
	b.SubscriptionID(1)

	assert.Equal(t, 1, b.Params["user_id"])
	assert.Equal(t, 1, b.Params["subscription_id"])
}

func TestOrdersGetUserSubscriptionsBuilder(t *testing.T) {
	t.Parallel()

	b := params.NewOrdersGetUserSubscriptionsBuilder()

	b.UserID(1)

	assert.Equal(t, 1, b.Params["user_id"])
}

func TestOrdersUpdateSubscriptionBuilder(t *testing.T) {
	t.Parallel()

	b := params.NewOrdersUpdateSubscriptionBuilder()

	b.UserID(1)
	b.SubscriptionID(1)
	b.Price(1)

	assert.Equal(t, 1, b.Params["user_id"])
	assert.Equal(t, 1, b.Params["subscription_id"])
	assert.Equal(t, 1, b.Params["price"])
}
