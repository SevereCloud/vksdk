package streaming // import "github.com/SevereCloud/vksdk/v3/streaming"

import (
	"errors"
	"fmt"
)

// ErrorType is the type of an error.
type ErrorType int

// Error returns the message of a ErrorType.
func (e ErrorType) Error() string {
	return fmt.Sprintf("streaming: error with code %d", e)
}

// Error codes. See https://dev.vk.ru/ru/api/streaming/getting-started_2
const (
	ErrNoType ErrorType = 0 // NoType error
	// Settings for updating the connection to WebSocket are incorrectly
	// transmitted.
	ErrUpgradeWebsocket             ErrorType = 1000
	ErrUnsupportedHTTP              ErrorType = 1001 // Unsupported HTTP method
	ErrContentType                  ErrorType = 1002 // “Content-type” key is absent or invalid
	ErrMissingKey                   ErrorType = 1003 // "key" parameter is absent
	ErrBadKey                       ErrorType = 1004 // "key" parameter is invalid
	ErrBadStreamID                  ErrorType = 1005 // "stream_id" parameter is invalid
	ErrConnectionAlreadyEstablished ErrorType = 1006 // Such connection has been established already
	ErrInvalidJSON                  ErrorType = 2000 // Invalid JSON
	ErrTagAlreadyExist              ErrorType = 2001 // The rule this such tag has been added already
	ErrTagNotExist                  ErrorType = 2002 // The this such tag is absent
	ErrInvalidRule                  ErrorType = 2003 // Invalid rule
	ErrTooManyFilters               ErrorType = 2004 // Too many filters in one rule
	ErrUnbalancedQuotes             ErrorType = 2005 // Unbalanced quotes
	ErrTooManyRules                 ErrorType = 2006 // Too many rules in one stream
	ErrMinusKeywordsOnly            ErrorType = 2008 // You can't use minus keywords only
)

// Error description.
type Error struct {
	Code    ErrorType `json:"error_code"` // error code
	Message string    `json:"message"`    // error message
}

// Error returns the message of a Error.
func (e Error) Error() string {
	return "streaming: " + e.Message
}

// Is unwraps its first argument sequentially looking for an error that matches
// the second.
func (e Error) Is(target error) bool {
	var tError *Error
	if errors.As(target, &tError) {
		return e.Code == tError.Code && e.Message == tError.Message
	}

	var tErrorType ErrorType
	if errors.As(target, &tErrorType) {
		return e.Code == tErrorType
	}

	return false
}
