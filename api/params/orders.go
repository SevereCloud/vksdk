package params // import "github.com/SevereCloud/vksdk/api/params"

import (
	"github.com/SevereCloud/vksdk/api"
)

// OrdersCancelSubscriptionBuilder builder
//
// https://vk.com/dev/orders.cancelSubscription
type OrdersCancelSubscriptionBuilder struct {
	api.Params
}

// NewOrdersCancelSubscriptionBuilder func
func NewOrdersCancelSubscriptionBuilder() *OrdersCancelSubscriptionBuilder {
	return &OrdersCancelSubscriptionBuilder{api.Params{}}
}

// UserID parameter
func (b *OrdersCancelSubscriptionBuilder) UserID(v int) {
	b.Params["user_id"] = v
}

// SubscriptionID parameter
func (b *OrdersCancelSubscriptionBuilder) SubscriptionID(v int) {
	b.Params["subscription_id"] = v
}

// PendingCancel parameter
func (b *OrdersCancelSubscriptionBuilder) PendingCancel(v bool) {
	b.Params["pending_cancel"] = v
}

// OrdersChangeStateBuilder builder
//
// Changes order status.
//
// https://vk.com/dev/orders.changeState
type OrdersChangeStateBuilder struct {
	api.Params
}

// NewOrdersChangeStateBuilder func
func NewOrdersChangeStateBuilder() *OrdersChangeStateBuilder {
	return &OrdersChangeStateBuilder{api.Params{}}
}

// OrderID order ID.
func (b *OrdersChangeStateBuilder) OrderID(v int) {
	b.Params["order_id"] = v
}

// Action action to be done with the order. Available actions: *cancel — to cancel unconfirmed order. *charge — to confirm unconfirmed order. Applies only if processing of [vk.com/dev/payments_status|order_change_state] notification failed. *refund — to cancel confirmed order.
func (b *OrdersChangeStateBuilder) Action(v string) {
	b.Params["action"] = v
}

// AppOrderID internal ID of the order in the application.
func (b *OrdersChangeStateBuilder) AppOrderID(v int) {
	b.Params["app_order_id"] = v
}

// TestMode if this parameter is set to 1, this method returns a list of test mode orders. By default — 0.
func (b *OrdersChangeStateBuilder) TestMode(v bool) {
	b.Params["test_mode"] = v
}

// OrdersGetBuilder builder
//
// Returns a list of orders.
//
// https://vk.com/dev/orders.get
type OrdersGetBuilder struct {
	api.Params
}

// NewOrdersGetBuilder func
func NewOrdersGetBuilder() *OrdersGetBuilder {
	return &OrdersGetBuilder{api.Params{}}
}

// Offset parameter
func (b *OrdersGetBuilder) Offset(v int) {
	b.Params["offset"] = v
}

// Count number of returned orders.
func (b *OrdersGetBuilder) Count(v int) {
	b.Params["count"] = v
}

// TestMode if this parameter is set to 1, this method returns a list of test mode orders. By default — 0.
func (b *OrdersGetBuilder) TestMode(v bool) {
	b.Params["test_mode"] = v
}

// OrdersGetAmountBuilder builder
//
// https://vk.com/dev/orders.getAmount
type OrdersGetAmountBuilder struct {
	api.Params
}

// NewOrdersGetAmountBuilder func
func NewOrdersGetAmountBuilder() *OrdersGetAmountBuilder {
	return &OrdersGetAmountBuilder{api.Params{}}
}

// UserID parameter
func (b *OrdersGetAmountBuilder) UserID(v int) {
	b.Params["user_id"] = v
}

// Votes parameter
func (b *OrdersGetAmountBuilder) Votes(v []string) {
	b.Params["votes"] = v
}

// OrdersGetByIDBuilder builder
//
// Returns information about orders by their IDs.
//
// https://vk.com/dev/orders.getById
type OrdersGetByIDBuilder struct {
	api.Params
}

// NewOrdersGetByIDBuilder func
func NewOrdersGetByIDBuilder() *OrdersGetByIDBuilder {
	return &OrdersGetByIDBuilder{api.Params{}}
}

// OrderID order ID.
func (b *OrdersGetByIDBuilder) OrderID(v int) {
	b.Params["order_id"] = v
}

// OrderIDs order IDs (when information about several orders is requested).
func (b *OrdersGetByIDBuilder) OrderIDs(v []int) {
	b.Params["order_ids"] = v
}

// TestMode if this parameter is set to 1, this method returns a list of test mode orders. By default — 0.
func (b *OrdersGetByIDBuilder) TestMode(v bool) {
	b.Params["test_mode"] = v
}

// OrdersGetUserSubscriptionByIDBuilder builder
//
// https://vk.com/dev/orders.getUserSubscriptionById
type OrdersGetUserSubscriptionByIDBuilder struct {
	api.Params
}

// NewOrdersGetUserSubscriptionByIDBuilder func
func NewOrdersGetUserSubscriptionByIDBuilder() *OrdersGetUserSubscriptionByIDBuilder {
	return &OrdersGetUserSubscriptionByIDBuilder{api.Params{}}
}

// UserID parameter
func (b *OrdersGetUserSubscriptionByIDBuilder) UserID(v int) {
	b.Params["user_id"] = v
}

// SubscriptionID parameter
func (b *OrdersGetUserSubscriptionByIDBuilder) SubscriptionID(v int) {
	b.Params["subscription_id"] = v
}

// OrdersGetUserSubscriptionsBuilder builder
//
// https://vk.com/dev/orders.getUserSubscriptions
type OrdersGetUserSubscriptionsBuilder struct {
	api.Params
}

// NewOrdersGetUserSubscriptionsBuilder func
func NewOrdersGetUserSubscriptionsBuilder() *OrdersGetUserSubscriptionsBuilder {
	return &OrdersGetUserSubscriptionsBuilder{api.Params{}}
}

// UserID parameter
func (b *OrdersGetUserSubscriptionsBuilder) UserID(v int) {
	b.Params["user_id"] = v
}

// OrdersUpdateSubscriptionBuilder builder
//
// https://vk.com/dev/orders.updateSubscription
type OrdersUpdateSubscriptionBuilder struct {
	api.Params
}

// NewOrdersUpdateSubscriptionBuilder func
func NewOrdersUpdateSubscriptionBuilder() *OrdersUpdateSubscriptionBuilder {
	return &OrdersUpdateSubscriptionBuilder{api.Params{}}
}

// UserID parameter
func (b *OrdersUpdateSubscriptionBuilder) UserID(v int) {
	b.Params["user_id"] = v
}

// SubscriptionID parameter
func (b *OrdersUpdateSubscriptionBuilder) SubscriptionID(v int) {
	b.Params["subscription_id"] = v
}

// Price parameter
func (b *OrdersUpdateSubscriptionBuilder) Price(v int) {
	b.Params["price"] = v
}
