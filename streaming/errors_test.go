package streaming_test

import (
	"errors"
	"testing"

	"github.com/SevereCloud/vksdk/api"
	"github.com/SevereCloud/vksdk/streaming"
	"github.com/stretchr/testify/assert"
)

func TestErrorType(t *testing.T) {
	err := streaming.ErrorType(1)
	assert.EqualError(t, err, "streaming: error with code 1")
}

func TestError_Error(t *testing.T) {
	err := streaming.Error{
		Code:    streaming.ErrorType(1),
		Message: "test message",
	}
	assert.EqualError(t, err, "streaming: test message")
}

func TestError_Is(t *testing.T) {
	f := func(err *streaming.Error, target error, want bool) {
		assert.Equal(t, want, errors.Is(err, target))
	}

	e1 := streaming.ErrorType(1)
	e2 := streaming.ErrorType(2)

	f(&streaming.Error{Code: streaming.ErrorType(1)}, &streaming.Error{Code: streaming.ErrorType(1)}, true)
	f(&streaming.Error{Code: streaming.ErrorType(1)}, &e1, true)
	f(&streaming.Error{Code: streaming.ErrorType(1), Message: "123"}, &streaming.Error{Code: streaming.ErrorType(1), Message: "123"}, true)

	f(&streaming.Error{Code: streaming.ErrorType(1)}, &streaming.Error{Code: streaming.ErrorType(2)}, false)
	f(&streaming.Error{Code: streaming.ErrorType(1)}, &e2, false)
	f(&streaming.Error{Code: streaming.ErrorType(1), Message: "123"}, &streaming.Error{Code: streaming.ErrorType(1), Message: "321"}, false)
	f(&streaming.Error{Code: streaming.ErrorType(1)}, &api.Error{}, false)
}

func TestError_As(t *testing.T) {
	f := func(err *streaming.Error, target interface{}, want bool) {
		assert.Equal(t, want, errors.As(err, target))
	}

	e1 := streaming.ErrorType(1)
	e2 := streaming.ErrorType(2)

	f(&streaming.Error{Code: streaming.ErrorType(1)}, &streaming.Error{Code: streaming.ErrorType(1)}, true)
	f(&streaming.Error{Code: streaming.ErrorType(1)}, &e1, true)
	f(&streaming.Error{Code: streaming.ErrorType(1), Message: "123"}, &streaming.Error{Code: streaming.ErrorType(1), Message: "123"}, true)
	f(&streaming.Error{Code: streaming.ErrorType(1), Message: "123"}, &streaming.Error{Code: streaming.ErrorType(1), Message: "321"}, true)

	f(&streaming.Error{Code: streaming.ErrorType(1)}, &streaming.Error{Code: streaming.ErrorType(2)}, false)
	f(&streaming.Error{Code: streaming.ErrorType(1)}, &e2, false)
	f(&streaming.Error{Code: streaming.ErrorType(1)}, &api.Error{}, false)
}
