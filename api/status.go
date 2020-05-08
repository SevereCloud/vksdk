package api // import "github.com/SevereCloud/vksdk/api"

import (
	"github.com/SevereCloud/vksdk/object"
)

// StatusGetResponse struct.
type StatusGetResponse struct {
	Audio object.AudioAudioFull `json:"audio"`
	Text  string                `json:"text"`
}

// StatusGet returns data required to show the status of a user or community.
func (vk *VK) StatusGet(params Params) (response StatusGetResponse, err error) {
	err = vk.RequestUnmarshal("status.get", params, &response)
	return
}

// StatusSet sets a new status for the current user.
func (vk *VK) StatusSet(params Params) (response int, err error) {
	err = vk.RequestUnmarshal("status.set", params, &response)
	return
}
