package api // import "github.com/SevereCloud/vksdk/api"

import (
	"github.com/SevereCloud/vksdk/object"
)

// AdsDeleteTargetGroup deletes target group.
//
// https://vk.com/dev/ads.deleteTargetGroup
func (vk *VK) AdsDeleteTargetGroup(params Params) (response int, err error) {
	err = vk.RequestUnmarshal("ads.deleteTargetGroup", params, &response)
	return
}

// AdsDeleteTargetPixel deletes target pixel.
//
// https://vk.com/dev/ads.deleteTargetPixel
func (vk *VK) AdsDeleteTargetPixel(params Params) (response int, err error) {
	err = vk.RequestUnmarshal("ads.deleteTargetPixel", params, &response)
	return
}

// AdsGetAccountsResponse struct
type AdsGetAccountsResponse []object.AdsAccount

// AdsGetAccounts returns a list of advertising accounts.
//
// https://vk.com/dev/ads.getAccounts
func (vk *VK) AdsGetAccounts(params Params) (response AdsGetAccountsResponse, err error) {
	err = vk.RequestUnmarshal("ads.getAccounts", params, &response)
	return
}

// AdsRemoveTargetContacts accepts the request to exclude the advertiser's
// contacts from the retargeting audience.
//
// The maximum allowed number of contacts to be excluded by a single
// request is 1000.
//
// Contacts are excluded within a few hours of the request.
//
// https://vk.com/dev/ads.removeTargetContacts
func (vk *VK) AdsRemoveTargetContacts(params Params) (response int, err error) {
	err = vk.RequestUnmarshal("ads.removeTargetContacts", params, &response)
	return
}

// AdsUpdateTargetGroup edits target group.
//
// https://vk.com/dev/ads.updateTargetGroup
func (vk *VK) AdsUpdateTargetGroup(params Params) (response int, err error) {
	err = vk.RequestUnmarshal("ads.updateTargetGroup", params, &response)
	return
}

// AdsUpdateTargetPixel edits target pixel.
//
// https://vk.com/dev/ads.updateTargetPixel
func (vk *VK) AdsUpdateTargetPixel(params Params) (response int, err error) {
	err = vk.RequestUnmarshal("ads. updateTargetPixel", params, &response)
	return
}
