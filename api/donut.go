package api

import "github.com/SevereCloud/vksdk/v3/object"

// DonutGetFriendsResponse struct.
type DonutGetFriendsResponse struct {
	Count int                `json:"count"`
	Items []object.UsersUser `json:"items"`
}

// DonutGetFriends method.
//
// https://dev.vk.com/method/donut.getFriends
func (vk *VK) DonutGetFriends(params Params) (response DonutGetFriendsResponse, err error) {
	err = vk.RequestUnmarshal("donut.getFriends", &response, params)
	return
}

// DonutGetSubscription method.
//
// https://dev.vk.com/method/donut.getSubscription
func (vk *VK) DonutGetSubscription(params Params) (response object.DonutDonatorSubscriptionInfo, err error) {
	err = vk.RequestUnmarshal("donut.getSubscription", &response, params)
	return
}

// DonutGetSubscriptionsResponse struct.
type DonutGetSubscriptionsResponse struct {
	Subscriptions []object.DonutDonatorSubscriptionInfo `json:"subscriptions"`
	Count         int                                   `json:"count"`
	Profiles      []object.UsersUser                    `json:"profiles"`
	Groups        []object.GroupsGroup                  `json:"groups"`
}

// DonutGetSubscriptions method.
//
// https://dev.vk.com/method/donut.getSubscriptions
func (vk *VK) DonutGetSubscriptions(params Params) (response DonutGetSubscriptionsResponse, err error) {
	err = vk.RequestUnmarshal("donut.getSubscriptions", &response, params)
	return
}

// DonutIsDon method.
//
// https://dev.vk.com/method/donut.isDon
func (vk *VK) DonutIsDon(params Params) (response int, err error) {
	err = vk.RequestUnmarshal("donut.isDon", &response, params)
	return
}
