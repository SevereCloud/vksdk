package callback // import "github.com/SevereCloud/vksdk/v2/callback"

import (
	"context"

	"github.com/SevereCloud/vksdk/v2/internal"
)

// RetryCounterFromContext returns the X-Retry-Counter from context.
func RetryCounterFromContext(ctx context.Context) int {
	return ctx.Value(internal.CallbackRetryCounterKey).(int)
}
