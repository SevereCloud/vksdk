package callback // import "github.com/SevereCloud/vksdk/v3/callback"

import (
	"context"
	"time"

	"github.com/SevereCloud/vksdk/v3/internal"
)

// RetryCounterFromContext returns the X-Retry-Counter from context.
func RetryCounterFromContext(ctx context.Context) int {
	return ctx.Value(internal.CallbackRetryCounterKey).(int)
}

// RetryAfter send the "Retry-After" header field to indicate how long the
// VK Callback ought to wait before making a repeated request.
// Not work with goroutine mode!
//
// Possible HTTP status codes:
//
//	http.StatusGone
//	http.StatusTooManyRequests
//	http.StatusServiceUnavailable
//
// The resend time range must be less than 3 hours. The actual time of
// forwarding an event notification may be longer than the specified time.
func RetryAfter(ctx context.Context, code int, date time.Time) {
	ctx.Value(internal.CallbackRetryAfterKey).(func(int, time.Time))(
		code,
		date,
	)
}

// Remove VK Callback server.
func Remove(ctx context.Context) {
	ctx.Value(internal.CallbackRemove).(func())()
}
