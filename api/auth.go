package api // import "github.com/SevereCloud/vksdk/api"

// AuthCheckPhone checks a user's phone number for correctness.
//
// https://vk.com/dev/auth.checkPhone
func (vk *VK) AuthCheckPhone(params Params) (response int, err error) {
	err = vk.RequestUnmarshal("auth.checkPhone", params, &response)
	return
}

// AuthRetoreResponse struct
type AuthRetoreResponse struct {
	Success int    `json:"success"`
	SID     string `json:"sid"`
}

// AuthRetore allows to restore account access using a code received via SMS.
//
// https://vk.com/dev/auth.restore
func (vk *VK) AuthRetore(params Params) (response AuthRetoreResponse, err error) {
	err = vk.RequestUnmarshal("auth.restore", params, &response)
	return
}
