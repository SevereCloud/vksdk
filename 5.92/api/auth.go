package api // import "github.com/SevereCloud/vksdk/5.92/api"

import "encoding/json"

// AuthCheckPhone checks a user's phone number for correctness.
// https://vk.com/dev/auth.checkPhone
func (vk VK) AuthCheckPhone(params map[string]string) (vkErr Error) {
	_, vkErr = vk.Request("auth.checkPhone", params)
	return
}

// AuthRetoreResponse struct
type AuthRetoreResponse struct {
	Success int    `json:"success"`
	SID     string `json:"sid"`
}

// AuthRetore allows to restore account access using a code received via SMS.
// https://vk.com/dev/auth.restore
func (vk VK) AuthRetore(params map[string]string) (response AuthRetoreResponse, vkErr Error) {
	rawResponse, vkErr := vk.Request("auth.restore", params)
	if vkErr.Code != 0 {
		return
	}

	err := json.Unmarshal(rawResponse, &response)
	if err != nil {
		panic(err)
	}

	return
}
