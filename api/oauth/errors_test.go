package oauth_test

import (
	"errors"
	"testing"

	"github.com/SevereCloud/vksdk/v2/api/oauth"
	"github.com/stretchr/testify/assert"
)

type otherError string

func (e otherError) Error() string {
	return string(e)
}

func TestErrorType(t *testing.T) {
	t.Parallel()

	err := oauth.ErrorType("aoa")
	assert.EqualError(t, err, "oauth: error with type aoa")
}

func TestErrorReason(t *testing.T) {
	t.Parallel()

	err := oauth.ErrorReason("aoa")
	assert.EqualError(t, err, "oauth: error with reason aoa")
}

func TestError_Error(t *testing.T) {
	t.Parallel()

	err := oauth.Error{
		Type:        oauth.ErrorType("aoa"),
		Description: "test message",
	}
	assert.EqualError(t, err, "oauth: test message")

	err = oauth.Error{
		Type: oauth.ErrorType("aoa"),
	}
	assert.EqualError(t, err, "oauth: error with type aoa")
}

func TestError_Is(t *testing.T) {
	t.Parallel()

	f := func(err *oauth.Error, target error, want bool) {
		t.Helper()

		assert.Equal(t, want, errors.Is(err, target))
	}

	f(&oauth.Error{Type: oauth.ErrorType("aoa")}, &oauth.Error{Type: oauth.ErrorType("aoa")}, true)
	f(&oauth.Error{Type: oauth.ErrAccessDenied}, oauth.ErrAccessDenied, true)
	f(&oauth.Error{Reason: oauth.ErrUserDenied}, oauth.ErrUserDenied, true)
	f(&oauth.Error{Type: oauth.ErrorType("aoa"), Description: "123"}, &oauth.Error{Type: oauth.ErrorType("aoa"), Description: "123"}, true)

	f(&oauth.Error{Type: oauth.ErrorType("aoa")}, &oauth.Error{Type: oauth.ErrorType("oao")}, false)
	f(&oauth.Error{Type: oauth.ErrorType("aoa")}, oauth.ErrorType("oao"), false)
	f(&oauth.Error{Reason: oauth.ErrorReason("aoa")}, oauth.ErrorReason("oao"), false)
	f(&oauth.Error{Type: oauth.ErrorType("aoa"), Description: "123"}, &oauth.Error{Type: oauth.ErrorType("aoa"), Description: "321"}, false)
	f(&oauth.Error{Type: oauth.ErrorType("aoa")}, otherError("test"), false)
}

func TestError_As(t *testing.T) {
	t.Parallel()

	var target *oauth.Error

	err := &oauth.Error{Type: oauth.ErrorType("aoa")}
	if !errors.As(err, &target) && target.Type == "aoa" {
		t.Error("As not working")
	}
}
