package params // import "github.com/SevereCloud/vksdk/api/params"

import (
	"github.com/SevereCloud/vksdk/api"
)

// AuthCheckPhoneBuilder builder
//
// Checks a user's phone number for correctness.
//
// https://vk.com/dev/auth.checkPhone
type AuthCheckPhoneBuilder struct {
	api.Params
}

// NewAuthCheckPhoneBuilder func
func NewAuthCheckPhoneBuilder() *AuthCheckPhoneBuilder {
	return &AuthCheckPhoneBuilder{api.Params{}}
}

// Phone Phone number.
func (b *AuthCheckPhoneBuilder) Phone(v string) {
	b.Params["phone"] = v
}

// ClientID User ID.
func (b *AuthCheckPhoneBuilder) ClientID(v int) {
	b.Params["client_id"] = v
}

// ClientSecret parameter
func (b *AuthCheckPhoneBuilder) ClientSecret(v string) {
	b.Params["client_secret"] = v
}

// AuthByPhone parameter
func (b *AuthCheckPhoneBuilder) AuthByPhone(v bool) {
	b.Params["auth_by_phone"] = v
}

// AuthRestoreBuilder builder
//
// Allows to restore account access using a code received via SMS. " This method is only available for apps with [vk.com/dev/auth_direct|Direct authorization] access. "
//
// https://vk.com/dev/auth.restore
type AuthRestoreBuilder struct {
	api.Params
}

// NewAuthRestoreBuilder func
func NewAuthRestoreBuilder() *AuthRestoreBuilder {
	return &AuthRestoreBuilder{api.Params{}}
}

// Phone User phone number.
func (b *AuthRestoreBuilder) Phone(v string) {
	b.Params["phone"] = v
}

// LastName User last name.
func (b *AuthRestoreBuilder) LastName(v string) {
	b.Params["last_name"] = v
}
