package longpoll_test

import (
	"context"
	"testing"

	"github.com/SevereCloud/vksdk/v3/internal"
	"github.com/SevereCloud/vksdk/v3/longpoll-bot"
	"github.com/stretchr/testify/assert"
)

func TestTsFromContext(t *testing.T) {
	t.Parallel()

	const ts = 123
	ctx := context.WithValue(context.Background(), internal.LongPollTsKey, ts)
	assert.Equal(t, ts, longpoll.TsFromContext(ctx))
}
