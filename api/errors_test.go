package api_test

import (
	"errors"
	"testing"

	"github.com/SevereCloud/vksdk/api"
	"github.com/SevereCloud/vksdk/streaming"
	"github.com/stretchr/testify/assert"
)

func TestErrorType(t *testing.T) {
	err := api.ErrorType(1)
	assert.EqualError(t, err, "api: error with code 1")
}

func TestError_Error(t *testing.T) {
	err := api.Error{
		Code:    api.ErrorType(1),
		Message: "test message",
	}
	assert.EqualError(t, err, "api: test message")
}

func TestError_Is(t *testing.T) {
	f := func(err *api.Error, target error, want bool) {
		assert.Equal(t, want, errors.Is(err, target))
	}

	f(&api.Error{Code: api.ErrorType(1)}, &api.Error{Code: api.ErrorType(1)}, true)
	f(&api.Error{Code: api.ErrAuth}, api.ErrAuth, true)
	f(&api.Error{Code: api.ErrorType(1), Message: "123"}, &api.Error{Code: api.ErrorType(1), Message: "123"}, true)

	f(&api.Error{Code: api.ErrorType(1)}, &api.Error{Code: api.ErrorType(2)}, false)
	f(&api.Error{Code: api.ErrorType(1)}, api.ErrorType(2), false)
	f(&api.Error{Code: api.ErrorType(1), Message: "123"}, &api.Error{Code: api.ErrorType(1), Message: "321"}, false)
	f(&api.Error{Code: api.ErrorType(1)}, &streaming.Error{}, false)
}

func TestError_As(t *testing.T) {
	var target *api.Error

	err := &api.Error{Code: api.ErrorType(1)}
	if !errors.As(err, &target) && target.Code == 1 {
		t.Error("As not working")
	}
}

func TestInvalidContentType(t *testing.T) {
	err := api.InvalidContentType{}
	assert.EqualError(t, err, "api: invalid content-type")
}

func TestExecuteErrors(t *testing.T) {
	err := api.ExecuteErrors{api.ExecuteError{}}
	assert.EqualError(t, err, "api: execute errors (1)")
}
