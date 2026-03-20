package longpoll // import "github.com/SevereCloud/vksdk/v3/longpoll-bot"

import (
	"context"

	"github.com/SevereCloud/vksdk/v3/internal"
)

// TsFromContext returns the ts from context.
func TsFromContext(ctx context.Context) string {
	if ts, ok := ctx.Value(internal.LongPollTsKey).(string); ok {
		return ts
	}

	return ""
}
