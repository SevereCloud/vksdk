package api_test

import (
	"errors"
	"testing"

	"github.com/SevereCloud/vksdk/v2/api"
	"github.com/SevereCloud/vksdk/v2/streaming"
	"github.com/stretchr/testify/assert"
)

func TestErrorType(t *testing.T) {
	t.Parallel()

	err := api.ErrorType(1)
	assert.EqualError(t, err, "api: error with code 1")
}

func TestErrorSubtype(t *testing.T) {
	t.Parallel()

	err := api.ErrorSubtype(1)
	assert.EqualError(t, err, "api: error with subcode 1")
}

func TestError_Error(t *testing.T) {
	t.Parallel()

	err := api.Error{
		Code:    api.ErrorType(1),
		Message: "test message",
	}
	assert.EqualError(t, err, "api: test message")
}

func TestError_Is(t *testing.T) {
	t.Parallel()

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
	t.Parallel()

	var target *api.Error

	err := &api.Error{Code: api.ErrorType(1)}
	if !errors.As(err, &target) && target.Code == 1 {
		t.Error("As not working")
	}
}

func TestInvalidContentType(t *testing.T) {
	t.Parallel()

	err := api.InvalidContentType{}
	assert.EqualError(t, err, "api: invalid content-type")
}

func TestExecuteErrors(t *testing.T) {
	t.Parallel()

	err := api.ExecuteErrors{api.ExecuteError{}}
	assert.EqualError(t, err, "api: execute errors (1)")
}

func TestAdsError_Error(t *testing.T) {
	t.Parallel()

	err := api.AdsError{
		Code: 1,
		Desc: "test message",
	}
	assert.EqualError(t, err, "api: test message")
}

func TestAdsError_Is(t *testing.T) {
	t.Parallel()

	f := func(err *api.AdsError, target error, want bool) {
		assert.Equal(t, want, errors.Is(err, target))
	}

	f(&api.AdsError{Code: api.ErrorType(1)}, &api.AdsError{Code: api.ErrorType(1)}, true)
	f(&api.AdsError{Code: api.ErrAuth}, api.ErrAuth, true)
	f(&api.AdsError{Code: api.ErrorType(1), Desc: "123"}, &api.AdsError{Code: api.ErrorType(1), Desc: "123"}, true)

	f(&api.AdsError{Code: api.ErrorType(1)}, &api.AdsError{Code: api.ErrorType(2)}, false)
	f(&api.AdsError{Code: api.ErrorType(1)}, api.ErrorType(2), false)
	f(&api.AdsError{Code: api.ErrorType(1), Desc: "123"}, &api.AdsError{Code: api.ErrorType(1), Desc: "321"}, false)
	f(&api.AdsError{Code: api.ErrorType(1)}, &streaming.Error{}, false)
}

func TestAuthSilentTokenError_Error(t *testing.T) {
	t.Parallel()

	err := api.AuthSilentTokenError{
		Code:        1,
		Description: "test message",
	}
	assert.EqualError(t, err, "api: test message")
}

func TestAuthSilentTokenError_Is(t *testing.T) {
	t.Parallel()

	f := func(err *api.AuthSilentTokenError, target error, want bool) {
		assert.Equal(t, want, errors.Is(err, target))
	}

	f(&api.AuthSilentTokenError{Code: api.ErrorType(1)}, &api.AuthSilentTokenError{Code: api.ErrorType(1)}, true)
	f(&api.AuthSilentTokenError{Code: api.ErrAuth}, api.ErrAuth, true)
	f(&api.AuthSilentTokenError{Code: api.ErrorType(1), Description: "123"}, &api.AuthSilentTokenError{Code: api.ErrorType(1), Description: "123"}, true)

	f(&api.AuthSilentTokenError{Code: api.ErrorType(1)}, &api.AuthSilentTokenError{Code: api.ErrorType(2)}, false)
	f(&api.AuthSilentTokenError{Code: api.ErrorType(1)}, api.ErrorType(2), false)
	f(&api.AuthSilentTokenError{Code: api.ErrorType(1), Description: "123"}, &api.AuthSilentTokenError{Code: api.ErrorType(1), Description: "321"}, false)
	f(&api.AuthSilentTokenError{Code: api.ErrorType(1)}, &streaming.Error{}, false)
}
