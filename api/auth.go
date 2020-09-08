package api // import "github.com/SevereCloud/vksdk/v2/api"

// AuthCheckPhone checks a user's phone number for correctness.
//
// https://vk.com/dev/auth.checkPhone
func (vk *VK) AuthCheckPhone(params Params) (response int, err error) {
	err = vk.RequestUnmarshal("auth.checkPhone", &response, params)
	return
}

// AuthRestoreResponse struct.
type AuthRestoreResponse struct {
	Success int    `json:"success"`
	SID     string `json:"sid"`
}

// AuthRestore allows to restore account access using a code received via SMS.
//
// https://vk.com/dev/auth.restore
func (vk *VK) AuthRestore(params Params) (response AuthRestoreResponse, err error) {
	err = vk.RequestUnmarshal("auth.restore", &response, params)
	return
}
