package api // import "github.com/severecloud/vksdk/api"

import (
	"encoding/json"

	"github.com/severecloud/vksdk/object"
)

// StatusGetResponse struct
type StatusGetResponse struct {
	Audio object.AudioAudioFull `json:"audio"`
	Text  string                `json:"text"`
}

// StatusGet returns data required to show the status of a user or community.
func (vk VK) StatusGet(params map[string]string) (response StatusGetResponse, vkErr Error) {
	rawResponse, vkErr := vk.Request("status.get", params)
	if vkErr.Code != 0 {
		return
	}

	err := json.Unmarshal(rawResponse, &response)
	if err != nil {
		panic(err)
	}

	return
}

// StatusSet sets a new status for the current user.
func (vk VK) StatusSet(params map[string]string) (vkErr Error) {
	_, vkErr = vk.Request("status.set", params)
	return
}
