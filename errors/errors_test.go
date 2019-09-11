package errors_test

import (
	"fmt"
	"testing"

	"github.com/SevereCloud/vksdk/errors"
	"github.com/SevereCloud/vksdk/object"
	"github.com/stretchr/testify/assert"
)

func TestContext(t *testing.T) {

	err := errors.Unknown.New("an_error")
	expectedContext := object.Error{Code: 1}
	errWithContext := errors.AddErrorContext(err, expectedContext)

	assert.Equal(t, errors.Unknown, errors.GetType(errWithContext))
	assert.Equal(t, expectedContext, errors.GetErrorContext(errWithContext))
	assert.Equal(t, err.Error(), errWithContext.Error())
}

func TestGetErrorContext(t *testing.T) {
	f := func(err error, wantContext object.Error) {
		t.Helper()

		context := errors.GetErrorContext(err)
		assert.Equal(t, context, wantContext)
	}

	f(fmt.Errorf("no type error"), object.Error{})
	f(errors.New(object.Error{Code: 1}), object.Error{Code: 1})
}

func TestGetType(t *testing.T) {
	f := func(err error, wantType errors.ErrorType) {
		t.Helper()

		errorType := errors.GetType(err)
		assert.Equal(t, errorType, wantType)
	}

	f(fmt.Errorf("no type error"), errors.NoType)
	f(errors.Unknown.New("Unknown type error"), errors.Unknown)
}
