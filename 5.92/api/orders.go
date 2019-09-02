package api // import "github.com/SevereCloud/vksdk/5.92/api"

import (
	"github.com/SevereCloud/vksdk/5.92/object"
)

// OrdersCancelSubscription Allows to cancel subscription.
//
// https://vk.com/dev/orders.cancelSubscription
func (vk *VK) OrdersCancelSubscription(params map[string]string) (response int, vkErr Error) {
	vk.RequestUnmarshal("orders.cancelSubscription", params, &response, &vkErr)
	return
}

// OrdersChangeStateResponse struct
type OrdersChangeStateResponse string // New state

// OrdersChangeState Changes order status.
//
// https://vk.com/dev/orders.changeState
func (vk *VK) OrdersChangeState(params map[string]string) (response OrdersChangeStateResponse, vkErr Error) {
	vk.RequestUnmarshal("orders.changeState", params, &response, &vkErr)
	return
}

// OrdersGetResponse struct
type OrdersGetResponse []object.OrdersOrder

// OrdersGet Returns a list of orders.
//
// https://vk.com/dev/orders.get
func (vk *VK) OrdersGet(params map[string]string) (response OrdersGetResponse, vkErr Error) {
	vk.RequestUnmarshal("orders.get", params, &response, &vkErr)
	return
}

// OrdersGetAmountResponse struct
type OrdersGetAmountResponse []object.OrdersAmount

// OrdersGetAmount Returns the cost of votes in the user's consent.
//
// https://vk.com/dev/orders.getAmount
func (vk *VK) OrdersGetAmount(params map[string]string) (response OrdersGetAmountResponse, vkErr Error) {
	vk.RequestUnmarshal("orders.getAmount", params, &response, &vkErr)
	return
}

// OrdersGetByIDResponse struct
type OrdersGetByIDResponse []object.OrdersOrder

// OrdersGetByID Returns information about orders by their IDs.
//
// https://vk.com/dev/orders.getByID
func (vk *VK) OrdersGetByID(params map[string]string) (response OrdersGetByIDResponse, vkErr Error) {
	vk.RequestUnmarshal("orders.getById", params, &response, &vkErr)
	return
}

// OrdersGetUserSubscriptionByIDResponse struct
type OrdersGetUserSubscriptionByIDResponse object.OrdersSubscription

// OrdersGetUserSubscriptionByID Allows to get subscription by its ID.
//
// https://vk.com/dev/orders.getUserSubscriptionById
func (vk *VK) OrdersGetUserSubscriptionByID(params map[string]string) (response OrdersGetUserSubscriptionByIDResponse, vkErr Error) {
	vk.RequestUnmarshal("orders.getUserSubscriptionById", params, &response, &vkErr)
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
func (vk *VK) OrdersGetUserSubscriptions(params map[string]string) (response OrdersGetUserSubscriptionsResponse, vkErr Error) {
	vk.RequestUnmarshal("orders.getUserSubscriptions", params, &response, &vkErr)
	return
}

// OrdersUpdateSubscription Allows to update subscription price.
//
// https://vk.com/dev/orders.updateSubscription
func (vk *VK) OrdersUpdateSubscription(params map[string]string) (response int, vkErr Error) {
	vk.RequestUnmarshal("orders.updateSubscription", params, &response, &vkErr)
	return
}
