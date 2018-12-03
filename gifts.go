package vksdk

import "encoding/json"

// GiftsGetResponse struct
type GiftsGetResponse struct {
	Count int         `json:"count"`
	Items []giftsGift `json:"items"`
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
