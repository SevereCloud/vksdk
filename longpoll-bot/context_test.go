package longpoll_test

import (
	"context"
	"testing"

	"github.com/SevereCloud/vksdk/v2/internal"
	"github.com/SevereCloud/vksdk/v2/longpoll-bot"
	"github.com/stretchr/testify/assert"
)

func TestTsFromContext(t *testing.T) {
	const ts = 123
	ctx := context.WithValue(context.Background(), internal.LongPollTsKey, ts)
	assert.Equal(t, ts, longpoll.TsFromContext(ctx))
}
