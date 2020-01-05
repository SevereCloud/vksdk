package api // import "github.com/SevereCloud/vksdk/api"

import (
	"github.com/SevereCloud/vksdk/object"
)

// OrdersCancelSubscription Allows to cancel subscription.
//
// https://vk.com/dev/orders.cancelSubscription
func (vk *VK) OrdersCancelSubscription(params Params) (response int, err error) {
	err = vk.RequestUnmarshal("orders.cancelSubscription", params, &response)
	return
}

// OrdersChangeStateResponse struct
type OrdersChangeStateResponse string // New state

// OrdersChangeState Changes order status.
//
// https://vk.com/dev/orders.changeState
func (vk *VK) OrdersChangeState(params Params) (response OrdersChangeStateResponse, err error) {
	err = vk.RequestUnmarshal("orders.changeState", params, &response)
	return
}

// OrdersGetResponse struct
type OrdersGetResponse []object.OrdersOrder

// OrdersGet Returns a list of orders.
//
// https://vk.com/dev/orders.get
func (vk *VK) OrdersGet(params Params) (response OrdersGetResponse, err error) {
	err = vk.RequestUnmarshal("orders.get", params, &response)
	return
}

// OrdersGetAmountResponse struct
type OrdersGetAmountResponse []object.OrdersAmount

// OrdersGetAmount Returns the cost of votes in the user's consent.
//
// https://vk.com/dev/orders.getAmount
func (vk *VK) OrdersGetAmount(params Params) (response OrdersGetAmountResponse, err error) {
	err = vk.RequestUnmarshal("orders.getAmount", params, &response)
	return
}

// OrdersGetByIDResponse struct
type OrdersGetByIDResponse []object.OrdersOrder

// OrdersGetByID Returns information about orders by their IDs.
//
// https://vk.com/dev/orders.getByID
func (vk *VK) OrdersGetByID(params Params) (response OrdersGetByIDResponse, err error) {
	err = vk.RequestUnmarshal("orders.getById", params, &response)
	return
}

// OrdersGetUserSubscriptionByIDResponse struct
type OrdersGetUserSubscriptionByIDResponse object.OrdersSubscription

// OrdersGetUserSubscriptionByID Allows to get subscription by its ID.
//
// https://vk.com/dev/orders.getUserSubscriptionById
func (vk *VK) OrdersGetUserSubscriptionByID(params Params) (response OrdersGetUserSubscriptionByIDResponse, err error) {
	err = vk.RequestUnmarshal("orders.getUserSubscriptionById", params, &response)
	return
}

// OrdersGetUserSubscriptionsResponse struct
type OrdersGetUserSubscriptionsResponse struct {
	Count int                         `json:"count"` // Total number
	Items []object.OrdersSubscription `json:"items"`
}

// OrdersGetUserSubscriptions Allows to get user's active subscriptions.
//
// https://vk.com/dev/orders.getUserSubscriptions
func (vk *VK) OrdersGetUserSubscriptions(params Params) (response OrdersGetUserSubscriptionsResponse, err error) {
	err = vk.RequestUnmarshal("orders.getUserSubscriptions", params, &response)
	return
}

// OrdersUpdateSubscription Allows to update subscription price.
//
// https://vk.com/dev/orders.updateSubscription
func (vk *VK) OrdersUpdateSubscription(params Params) (response int, err error) {
	err = vk.RequestUnmarshal("orders.updateSubscription", params, &response)
	return
}
