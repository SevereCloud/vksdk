package api // import "github.com/SevereCloud/vksdk/v2/api"

import (
	"github.com/SevereCloud/vksdk/v2/object"
)

// AuthCheckPhone checks a user's phone number for correctness.
//
// https://vk.com/dev/auth.checkPhone
//
// Deprecated: This method is deprecated and may be disabled soon, please avoid
// using it.
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

// AuthGetProfileInfoBySilentTokenResponse struct.
type AuthGetProfileInfoBySilentTokenResponse struct {
	Success []object.AuthSilentTokenProfile `json:"success"`
	Errors  []AuthSilentTokenError          `json:"errors"`
}

// AuthGetProfileInfoBySilentToken method.
//
// https://platform.vk.com/?p=DocsDashboard&docs=tokens_silent-token
func (vk *VK) AuthGetProfileInfoBySilentToken(params Params) (response AuthGetProfileInfoBySilentTokenResponse, err error) {
	err = vk.RequestUnmarshal("auth.getProfileInfoBySilentToken", &response, params)
	return
}
