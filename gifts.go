package vksdk

import "encoding/json"

// GiftsGetResponse struct
type GiftsGetResponse struct {
	Count int         `json:"count,omitempty"`
	Items []giftsGift `json:"items,omitempty"`
}

// GiftsGet returns a list of user gifts.
func (vk *VK) GiftsGet(params map[string]string) (GiftsGetResponse, error) {
	var response GiftsGetResponse

	rawResponse, err := vk.Request("gifts.get", params)
	if err != nil {
		return response, err
	}

	err = json.Unmarshal(rawResponse, &response)
	if err != nil {
		panic(err)
	}

	return response, nil
}
