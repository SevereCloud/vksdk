package streaming_test

import (
	"fmt"
	"testing"

	"github.com/SevereCloud/vksdk/streaming"
	"github.com/stretchr/testify/assert"
)

func TestContext(t *testing.T) {
	err := streaming.BadKey.New("an_error")
	expectedContext := streaming.Error{ErrorCode: 1004}
	errWithContext := streaming.AddErrorContext(err, expectedContext)

	assert.Equal(t, streaming.BadKey, streaming.GetType(errWithContext))
	assert.Equal(t, expectedContext, streaming.GetErrorContext(errWithContext))
	assert.Equal(t, err.Error(), errWithContext.Error())
}

func TestGetErrorContext(t *testing.T) {
	f := func(err error, wantContext streaming.Error) {
		t.Helper()

		context := streaming.GetErrorContext(err)
		assert.Equal(t, context, wantContext)
	}

	f(fmt.Errorf("no type error"), streaming.Error{})
	f(streaming.NewError(streaming.Error{ErrorCode: 1004}), streaming.Error{ErrorCode: 1004})
}

func TestGetType(t *testing.T) {
	f := func(err error, wantType streaming.ErrorType) {
		t.Helper()

		errorType := streaming.GetType(err)
		assert.Equal(t, errorType, wantType)
	}

	f(fmt.Errorf("no type error"), streaming.NoType)
	f(streaming.BadKey.New("BadKey type error"), streaming.BadKey)
}
