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
