package api // import "github.com/severecloud/vksdk/api"

import (
	"encoding/json"

	"github.com/severecloud/vksdk/object"
)

// GiftsGetResponse struct
type GiftsGetResponse struct {
	Count int                `json:"count"`
	Items []object.GiftsGift `json:"items"`
}

// GiftsGet returns a list of user gifts.
func (vk *VK) GiftsGet(params map[string]string) (response GiftsGetResponse, vkErr Error) {
	rawResponse, vkErr := vk.Request("gifts.get", params)
	if vkErr.Code != 0 {
		return
	}

	err := json.Unmarshal(rawResponse, &response)
	if err != nil {
		panic(err)
	}

	return
}
