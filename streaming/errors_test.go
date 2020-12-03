package streaming_test

import (
	"errors"
	"testing"

	"github.com/SevereCloud/vksdk/v2/api"
	"github.com/SevereCloud/vksdk/v2/streaming"
	"github.com/stretchr/testify/assert"
)

func TestErrorType(t *testing.T) {
	t.Parallel()

	err := streaming.ErrorType(1)
	assert.EqualError(t, err, "streaming: error with code 1")
}

func TestError_Error(t *testing.T) {
	t.Parallel()

	err := streaming.Error{
		Code:    streaming.ErrorType(1),
		Message: "test message",
	}
	assert.EqualError(t, err, "streaming: test message")
}

func TestError_Is(t *testing.T) {
	t.Parallel()

	f := func(err *streaming.Error, target error, want bool) {
		assert.Equal(t, want, errors.Is(err, target))
	}

	f(&streaming.Error{Code: streaming.ErrorType(1)}, &streaming.Error{Code: streaming.ErrorType(1)}, true)
	f(&streaming.Error{Code: streaming.ErrorType(1)}, streaming.ErrorType(1), true)
	f(&streaming.Error{Code: streaming.ErrorType(1), Message: "123"}, &streaming.Error{Code: streaming.ErrorType(1), Message: "123"}, true)

	f(&streaming.Error{Code: streaming.ErrorType(1)}, &streaming.Error{Code: streaming.ErrorType(2)}, false)
	f(&streaming.Error{Code: streaming.ErrorType(1)}, streaming.ErrorType(2), false)
	f(&streaming.Error{Code: streaming.ErrorType(1), Message: "123"}, &streaming.Error{Code: streaming.ErrorType(1), Message: "321"}, false)
	f(&streaming.Error{Code: streaming.ErrorType(1)}, &api.Error{}, false)
}

func TestError_As(t *testing.T) {
	t.Parallel()

	var target *streaming.Error

	err := &streaming.Error{Code: streaming.ErrorType(1)}
	if !errors.As(err, &target) && target.Code == 1 {
		t.Error("As not working")
	}
}
