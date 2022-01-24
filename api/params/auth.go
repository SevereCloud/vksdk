package params // import "github.com/SevereCloud/vksdk/v2/api/params"

import (
	"github.com/SevereCloud/vksdk/v2/api"
)

// AuthCheckPhoneBuilder builder.
//
// Checks a user's phone number for correctness.
//
// https://vk.com/dev/auth.checkPhone
type AuthCheckPhoneBuilder struct {
	api.Params
}

// NewAuthCheckPhoneBuilder func.
func NewAuthCheckPhoneBuilder() *AuthCheckPhoneBuilder {
	return &AuthCheckPhoneBuilder{api.Params{}}
}

// Phone phone number.
func (b *AuthCheckPhoneBuilder) Phone(v string) *AuthCheckPhoneBuilder {
	b.Params["phone"] = v
	return b
}

// ClientID user ID.
func (b *AuthCheckPhoneBuilder) ClientID(v int) *AuthCheckPhoneBuilder {
	b.Params["client_id"] = v
	return b
}

// ClientSecret parameter.
func (b *AuthCheckPhoneBuilder) ClientSecret(v string) *AuthCheckPhoneBuilder {
	b.Params["client_secret"] = v
	return b
}

// AuthByPhone parameter.
func (b *AuthCheckPhoneBuilder) AuthByPhone(v bool) *AuthCheckPhoneBuilder {
	b.Params["auth_by_phone"] = v
	return b
}

// AuthRestoreBuilder builder.
//
// Allows to restore account access using a code received via SMS.
// This method is only available for apps with [vk.com/dev/auth_direct|Direct authorization] access.
//
// https://vk.com/dev/auth.restore
type AuthRestoreBuilder struct {
	api.Params
}

// NewAuthRestoreBuilder func.
func NewAuthRestoreBuilder() *AuthRestoreBuilder {
	return &AuthRestoreBuilder{api.Params{}}
}

// Phone user phone number.
func (b *AuthRestoreBuilder) Phone(v string) *AuthRestoreBuilder {
	b.Params["phone"] = v
	return b
}

// LastName user last name.
func (b *AuthRestoreBuilder) LastName(v string) *AuthRestoreBuilder {
	b.Params["last_name"] = v
	return b
}

// AuthGetProfileInfoBySilentTokenBuilder builder.
//
// https://platform.vk.com/?p=DocsDashboard&docs=tokens_silent-token
type AuthGetProfileInfoBySilentTokenBuilder struct {
	api.Params
}

// NewAuthGetProfileInfoBySilentTokenBuilder func.
func NewAuthGetProfileInfoBySilentTokenBuilder() *AuthGetProfileInfoBySilentTokenBuilder {
	return &AuthGetProfileInfoBySilentTokenBuilder{api.Params{
		"token": "",
		"uuid":  "",
		"event": "",
	}}
}

// Add token, uuid, event to params.
func (b *AuthGetProfileInfoBySilentTokenBuilder) Add(token, uuid, event string) *AuthGetProfileInfoBySilentTokenBuilder {
	separator := ","
	if b.Params["token"] == "" {
		separator = ""
	}

	b.Params["token"] = b.Params["token"].(string) + separator + token
	b.Params["uuid"] = b.Params["uuid"].(string) + separator + uuid
	b.Params["event"] = b.Params["event"].(string) + separator + event

	return b
}

// AuthExchangeSilentAuthTokenBuilder builder.
//
// https://platform.vk.com/?p=DocsDashboard&docs=tokens_silent-token
type AuthExchangeSilentAuthTokenBuilder struct {
	api.Params
}

// NewAuthExchangeSilentAuthTokenBuilder func.
func NewAuthExchangeSilentAuthTokenBuilder() *AuthExchangeSilentAuthTokenBuilder {
	return &AuthExchangeSilentAuthTokenBuilder{api.Params{}}
}

// Token parameter.
func (b *AuthExchangeSilentAuthTokenBuilder) Token(v string) *AuthExchangeSilentAuthTokenBuilder {
	b.Params["token"] = v
	return b
}

// UUID parameter.
func (b *AuthExchangeSilentAuthTokenBuilder) UUID(v string) *AuthExchangeSilentAuthTokenBuilder {
	b.Params["uuid"] = v
	return b
}
