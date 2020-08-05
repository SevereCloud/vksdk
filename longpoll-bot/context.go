package longpoll // import "github.com/SevereCloud/vksdk/longpoll-bot"

import (
	"context"

	"github.com/SevereCloud/vksdk/internal"
)

// TsFromContext returns the ts from context.
func TsFromContext(ctx context.Context) int {
	return ctx.Value(internal.LongpollTsKey).(int)
}
