package params // import "github.com/SevereCloud/vksdk/api/params"

import (
	"github.com/SevereCloud/vksdk/api"
)

// OrdersCancelSubscriptionBulder builder
//
// https://vk.com/dev/orders.cancelSubscription
type OrdersCancelSubscriptionBulder struct {
	api.Params
}

// NewOrdersCancelSubscriptionBulder func
func NewOrdersCancelSubscriptionBulder() *OrdersCancelSubscriptionBulder {
	return &OrdersCancelSubscriptionBulder{api.Params{}}
}

// UserID parameter
func (b *OrdersCancelSubscriptionBulder) UserID(v int) {
	b.Params["user_id"] = v
}

// SubscriptionID parameter
func (b *OrdersCancelSubscriptionBulder) SubscriptionID(v int) {
	b.Params["subscription_id"] = v
}

// PendingCancel parameter
func (b *OrdersCancelSubscriptionBulder) PendingCancel(v bool) {
	b.Params["pending_cancel"] = v
}

// OrdersChangeStateBulder builder
//
// Changes order status.
//
// https://vk.com/dev/orders.changeState
type OrdersChangeStateBulder struct {
	api.Params
}

// NewOrdersChangeStateBulder func
func NewOrdersChangeStateBulder() *OrdersChangeStateBulder {
	return &OrdersChangeStateBulder{api.Params{}}
}

// OrderID order ID.
func (b *OrdersChangeStateBulder) OrderID(v int) {
	b.Params["order_id"] = v
}

// Action action to be done with the order. Available actions: *cancel — to cancel unconfirmed order. *charge — to confirm unconfirmed order. Applies only if processing of [vk.com/dev/payments_status|order_change_state] notification failed. *refund — to cancel confirmed order.
func (b *OrdersChangeStateBulder) Action(v string) {
	b.Params["action"] = v
}

// AppOrderID internal ID of the order in the application.
func (b *OrdersChangeStateBulder) AppOrderID(v int) {
	b.Params["app_order_id"] = v
}

// TestMode if this parameter is set to 1, this method returns a list of test mode orders. By default — 0.
func (b *OrdersChangeStateBulder) TestMode(v bool) {
	b.Params["test_mode"] = v
}

// OrdersGetBulder builder
//
// Returns a list of orders.
//
// https://vk.com/dev/orders.get
type OrdersGetBulder struct {
	api.Params
}

// NewOrdersGetBulder func
func NewOrdersGetBulder() *OrdersGetBulder {
	return &OrdersGetBulder{api.Params{}}
}

// Offset parameter
func (b *OrdersGetBulder) Offset(v int) {
	b.Params["offset"] = v
}

// Count number of returned orders.
func (b *OrdersGetBulder) Count(v int) {
	b.Params["count"] = v
}

// TestMode if this parameter is set to 1, this method returns a list of test mode orders. By default — 0.
func (b *OrdersGetBulder) TestMode(v bool) {
	b.Params["test_mode"] = v
}

// OrdersGetAmountBulder builder
//
// https://vk.com/dev/orders.getAmount
type OrdersGetAmountBulder struct {
	api.Params
}

// NewOrdersGetAmountBulder func
func NewOrdersGetAmountBulder() *OrdersGetAmountBulder {
	return &OrdersGetAmountBulder{api.Params{}}
}

// UserID parameter
func (b *OrdersGetAmountBulder) UserID(v int) {
	b.Params["user_id"] = v
}

// Votes parameter
func (b *OrdersGetAmountBulder) Votes(v []string) {
	b.Params["votes"] = v
}

// OrdersGetByIDBulder builder
//
// Returns information about orders by their IDs.
//
// https://vk.com/dev/orders.getById
type OrdersGetByIDBulder struct {
	api.Params
}

// NewOrdersGetByIDBulder func
func NewOrdersGetByIDBulder() *OrdersGetByIDBulder {
	return &OrdersGetByIDBulder{api.Params{}}
}

// OrderID order ID.
func (b *OrdersGetByIDBulder) OrderID(v int) {
	b.Params["order_id"] = v
}

// OrderIDs order IDs (when information about several orders is requested).
func (b *OrdersGetByIDBulder) OrderIDs(v []int) {
	b.Params["order_ids"] = v
}

// TestMode if this parameter is set to 1, this method returns a list of test mode orders. By default — 0.
func (b *OrdersGetByIDBulder) TestMode(v bool) {
	b.Params["test_mode"] = v
}

// OrdersGetUserSubscriptionByIDBulder builder
//
// https://vk.com/dev/orders.getUserSubscriptionById
type OrdersGetUserSubscriptionByIDBulder struct {
	api.Params
}

// NewOrdersGetUserSubscriptionByIDBulder func
func NewOrdersGetUserSubscriptionByIDBulder() *OrdersGetUserSubscriptionByIDBulder {
	return &OrdersGetUserSubscriptionByIDBulder{api.Params{}}
}

// UserID parameter
func (b *OrdersGetUserSubscriptionByIDBulder) UserID(v int) {
	b.Params["user_id"] = v
}

// SubscriptionID parameter
func (b *OrdersGetUserSubscriptionByIDBulder) SubscriptionID(v int) {
	b.Params["subscription_id"] = v
}

// OrdersGetUserSubscriptionsBulder builder
//
// https://vk.com/dev/orders.getUserSubscriptions
type OrdersGetUserSubscriptionsBulder struct {
	api.Params
}

// NewOrdersGetUserSubscriptionsBulder func
func NewOrdersGetUserSubscriptionsBulder() *OrdersGetUserSubscriptionsBulder {
	return &OrdersGetUserSubscriptionsBulder{api.Params{}}
}

// UserID parameter
func (b *OrdersGetUserSubscriptionsBulder) UserID(v int) {
	b.Params["user_id"] = v
}

// OrdersUpdateSubscriptionBulder builder
//
// https://vk.com/dev/orders.updateSubscription
type OrdersUpdateSubscriptionBulder struct {
	api.Params
}

// NewOrdersUpdateSubscriptionBulder func
func NewOrdersUpdateSubscriptionBulder() *OrdersUpdateSubscriptionBulder {
	return &OrdersUpdateSubscriptionBulder{api.Params{}}
}

// UserID parameter
func (b *OrdersUpdateSubscriptionBulder) UserID(v int) {
	b.Params["user_id"] = v
}

// SubscriptionID parameter
func (b *OrdersUpdateSubscriptionBulder) SubscriptionID(v int) {
	b.Params["subscription_id"] = v
}

// Price parameter
func (b *OrdersUpdateSubscriptionBulder) Price(v int) {
	b.Params["price"] = v
}
