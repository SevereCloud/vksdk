// Package oauth ...
package oauth // import "github.com/SevereCloud/vksdk/v3/api/oauth"

import "errors"

// ValidationType ...
type ValidationType string

// Possible values.
const (
	ValidationSMS ValidationType = "2fa_sms"
	ValidationApp ValidationType = "2fa_app"
)

// ErrorType for oauth.
type ErrorType string

// Error types.
//
// See https://tools.ietf.org/html/rfc6749#section-4.2.2.1
const (
	ErrInvalidRequest          ErrorType = "invalid_request"
	ErrUnauthorizedClient      ErrorType = "unauthorized_client"
	ErrUnsupportedResponseType ErrorType = "unsupported_response_type"
	ErrInvalidScope            ErrorType = "invalid_scope"
	ErrServerError             ErrorType = "server_error"
	ErrTemporarilyUnavailable  ErrorType = "temporarily_unavailable"
	ErrAccessDenied            ErrorType = "access_denied"

	ErrInvalidGrant ErrorType = "invalid_grant"

	ErrNeedValidation ErrorType = "need_validation"
	ErrNeedCaptcha    ErrorType = "need_captcha"
)

// Error returns the message of a Error.
func (e ErrorType) Error() string {
	return "oauth: error with type " + string(e)
}

// ErrorReason for oauth.
type ErrorReason string

// Error returns the message of a Error.
func (e ErrorReason) Error() string {
	return "oauth: error with reason " + string(e)
}

// ErrorReason types.
const (
	ErrUserDenied ErrorReason = "user_denied"
)

// Error for oauth.
type Error struct {
	Type        ErrorType   `json:"error"`
	Reason      ErrorReason `json:"error_reason,omitempty"`
	Description string      `json:"error_description,omitempty"`

	// For auth direct
	CaptchaSID     string         `json:"captcha_sid,omitempty"`
	CaptchaImg     string         `json:"captcha_img,omitempty"`
	RedirectURI    string         `json:"redirect_uri,omitempty"`
	ValidationType ValidationType `json:"validation_type,omitempty"`
	PhoneMask      string         `json:"phone_mask,omitempty"`
}

// Error returns the message of a Error.
func (e Error) Error() string {
	if e.Description != "" {
		return "oauth: " + e.Description
	}

	return e.Type.Error()
}

// Is unwraps its first argument sequentially looking for an error that matches
// the second.
func (e Error) Is(target error) bool {
	var tError *Error
	if errors.As(target, &tError) {
		return e.Type == tError.Type && e.Description == tError.Description
	}

	var tErrorType ErrorType
	if errors.As(target, &tErrorType) {
		return e.Type == tErrorType
	}

	var tErrorReason ErrorReason
	if errors.As(target, &tErrorReason) {
		return e.Reason == tErrorReason
	}

	return false
}
