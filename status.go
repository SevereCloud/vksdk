package vksdk

import "encoding/json"

// StatusGetResponse struct
type StatusGetResponse struct {
	Audio audioAudioFull `json:"audio,omitempty"`
	Text  string         `json:"text,omitempty"`
}

// StatusGet returns data required to show the status of a user or community.
func (vk *VK) StatusGet(params map[string]string) (StatusGetResponse, error) {
	var response StatusGetResponse

	rawResponse, err := vk.Request("status.get", params)
	if err != nil {
		return response, err
	}

	err = json.Unmarshal(rawResponse, &response)
	if err != nil {
		panic(err)
	}

	return response, nil
}

// StatusSet sets a new status for the current user.
func (vk *VK) StatusSet(params map[string]string) error {
	_, err := vk.Request("status.set", params)
	if err != nil {
		return err
	}

	return nil
}
