package errors_test

import (
	"testing"

	"github.com/SevereCloud/vksdk/object"
	"github.com/SevereCloud/vksdk/object/errors"
	"github.com/stretchr/testify/assert"
)

func TestContext(t *testing.T) {

	err := errors.Unknown.New("an_error")
	expectedContext := object.Error{}
	errWithContext := errors.AddErrorContext(err, expectedContext)

	assert.Equal(t, errors.Unknown, errors.GetType(errWithContext))
	assert.Equal(t, expectedContext, errors.GetErrorContext(errWithContext))
	assert.Equal(t, err.Error(), errWithContext.Error())
}

func TestContextInNoTypeError(t *testing.T) {
	err := errors.New("a custom error")

	expectedContext := object.Error{}
	errWithContext := errors.AddErrorContext(err, expectedContext)

	assert.Equal(t, errors.NoType, errors.GetType(errWithContext))
	assert.Equal(t, expectedContext, errors.GetErrorContext(errWithContext))
	assert.Equal(t, err.Error(), errWithContext.Error())
}

func TestWrapf(t *testing.T) {
	err := errors.New("an_error")
	wrappedError := errors.Unknown.Wrapf(err, "error %s", "1")

	assert.Equal(t, errors.Unknown, errors.GetType(wrappedError))
	assert.EqualError(t, wrappedError, "error 1: an_error")
}

func TestWrapfInNoTypeError(t *testing.T) {
	err := errors.Newf("an_error %s", "2")
	wrappedError := errors.Wrapf(err, "error %s", "1")

	assert.Equal(t, errors.NoType, errors.GetType(wrappedError))
	assert.EqualError(t, wrappedError, "error 1: an_error 2")
}
