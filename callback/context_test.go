package callback_test

import (
	"context"
	"testing"

	"github.com/SevereCloud/vksdk/v2/callback"
	"github.com/SevereCloud/vksdk/v2/internal"
	"github.com/stretchr/testify/assert"
)

func TestRetryCounterFromContext(t *testing.T) {
	const retryCounter = 123
	ctx := context.WithValue(
		context.Background(),
		internal.CallbackRetryCounterKey,
		retryCounter,
	)
	assert.Equal(t, retryCounter, callback.RetryCounterFromContext(ctx))
}
