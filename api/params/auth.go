package params // import "github.com/SevereCloud/vksdk/api/params"

import (
	"github.com/SevereCloud/vksdk/api"
)

// AuthCheckPhoneBulder builder
//
// Checks a user's phone number for correctness.
//
// https://vk.com/dev/auth.checkPhone
type AuthCheckPhoneBulder struct {
	api.Params
}

// NewAuthCheckPhoneBulder func
func NewAuthCheckPhoneBulder() *AuthCheckPhoneBulder {
	return &AuthCheckPhoneBulder{api.Params{}}
}

// Phone Phone number.
func (b *AuthCheckPhoneBulder) Phone(v string) {
	b.Params["phone"] = v
}

// ClientID User ID.
func (b *AuthCheckPhoneBulder) ClientID(v int) {
	b.Params["client_id"] = v
}

// ClientSecret parameter
func (b *AuthCheckPhoneBulder) ClientSecret(v string) {
	b.Params["client_secret"] = v
}

// AuthByPhone parameter
func (b *AuthCheckPhoneBulder) AuthByPhone(v bool) {
	b.Params["auth_by_phone"] = v
}

// AuthRestoreBulder builder
//
// Allows to restore account access using a code received via SMS. " This method is only available for apps with [vk.com/dev/auth_direct|Direct authorization] access. "
//
// https://vk.com/dev/auth.restore
type AuthRestoreBulder struct {
	api.Params
}

// NewAuthRestoreBulder func
func NewAuthRestoreBulder() *AuthRestoreBulder {
	return &AuthRestoreBulder{api.Params{}}
}

// Phone User phone number.
func (b *AuthRestoreBulder) Phone(v string) {
	b.Params["phone"] = v
}

// LastName User last name.
func (b *AuthRestoreBulder) LastName(v string) {
	b.Params["last_name"] = v
}
