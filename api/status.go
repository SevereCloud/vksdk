package api // import "github.com/SevereCloud/vksdk/api"

import (
	"github.com/SevereCloud/vksdk/object"
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

// StatusSet sets a new status for the current user.
func (vk *VK) StatusSet(params map[string]string) (response int, vkErr Error) {
	vk.RequestUnmarshal("status.set", params, &response, &vkErr)
	return
}
