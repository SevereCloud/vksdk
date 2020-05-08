package streaming // import "github.com/SevereCloud/vksdk/streaming"

import (
	"fmt"

	"github.com/pkg/errors"
)

// ErrorType is the type of an error.
type ErrorType int

// Error codes. See https://vk.com/dev/streaming_api_docs_2
const (
	NoType ErrorType = 0 // NoType error
	// Settings for updating the connection to WebSocket are incorrectly transmitted
	UpgradeWebsocket             ErrorType = 1000
	UnsupportedHTTP              ErrorType = 1001 // Unsupported HTTP method
	ContentType                  ErrorType = 1002 // “Content-type” key is absent or invalid
	MissingKey                   ErrorType = 1003 // "key" parameter is absent
	BadKey                       ErrorType = 1004 // "key" parameter is invalid
	BadStreamID                  ErrorType = 1005 // "stream_id" parameter is invalid
	ConnectionAlreadyEstablished ErrorType = 1006 // Such connection has been established already
	InvalidJSON                  ErrorType = 2000 // Invalid JSON
	TagAlreadyExist              ErrorType = 2001 // The rule this such tag has been added already
	TagNotExist                  ErrorType = 2002 // The this such tag is absent
	InvalidRule                  ErrorType = 2003 // Invalid rule
	TooManyFilters               ErrorType = 2004 // Too many filters in one rule
	UnbalancedQuotes             ErrorType = 2005 // Unbalanced quotes
	TooManyRules                 ErrorType = 2006 // Too many rules in one stream
	MinusKeywordsOnly            ErrorType = 2008 // You can't use minus keywords only
)

// Error description.
type Error struct {
	Message   string `json:"message"`    // error message;
	ErrorCode int    `json:"error_code"` // error code
}

type customError struct {
	errorType     ErrorType
	originalError error
	context       Error
}

// New creates a new customError.
func (errorType ErrorType) New(msg string) error {
	return customError{errorType: errorType, originalError: errors.New(msg)}
}

// Newf creates a new customError with formatted message.
func (errorType ErrorType) Newf(msg string, args ...interface{}) error {
	return customError{errorType: errorType, originalError: fmt.Errorf(msg, args...)}
}

// Wrap creates a new wrapped error.
func (errorType ErrorType) Wrap(err error, msg string) error {
	return errorType.Wrapf(err, msg)
}

// Wrapf creates a new wrapped error with formatted message.
func (errorType ErrorType) Wrapf(err error, msg string, args ...interface{}) error {
	return customError{errorType: errorType, originalError: errors.Wrapf(err, msg, args...)}
}

// Error returns the mssage of a customError.
func (error customError) Error() string {
	return error.originalError.Error()
}

// NewError creates a no type error.
func NewError(vkErr Error) error {
	if vkErr.ErrorCode == 0 {
		return nil
	}

	return customError{
		errorType:     ErrorType(vkErr.ErrorCode),
		originalError: errors.New(vkErr.Message),
		context:       vkErr,
	}
}

// Cause gives the original error.
func Cause(err error) error {
	return errors.Cause(err)
}

// AddErrorContext adds a context to an error.
func AddErrorContext(err error, context Error) error {
	if customErr, ok := err.(customError); ok {
		return customError{
			errorType:     customErr.errorType,
			originalError: customErr.originalError,
			context:       context,
		}
	}

	return customError{errorType: NoType, originalError: err, context: context}
}

// GetErrorContext returns the error context.
func GetErrorContext(err error) Error {
	if customErr, ok := err.(customError); ok {
		return customErr.context
	}

	return Error{}
}

// GetType returns the error type.
func GetType(err error) ErrorType {
	if customErr, ok := err.(customError); ok {
		return customErr.errorType
	}

	return NoType
}
