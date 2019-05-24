package api // import "github.com/SevereCloud/vksdk/5.92/api"

import (
	"github.com/SevereCloud/vksdk/5.92/object"
)

// StatusGetResponse struct
type StatusGetResponse struct {
	Audio object.AudioAudioFull `json:"audio"`
	Text  string                `json:"text"`
}

// StatusGet returns data required to show the status of a user or community.
func (vk *VK) StatusGet(params map[string]string) (response StatusGetResponse, vkErr Error) {
	vk.RequestUnmarshal("status.get", params, &response, &vkErr)
	return
}

// StatusSetResponse struct
type StatusSetResponse int

// StatusSet sets a new status for the current user.
func (vk *VK) StatusSet(params map[string]string) (response StatusSetResponse, vkErr Error) {
	vk.RequestUnmarshal("status.set", params, &response, &vkErr)
	return
}
