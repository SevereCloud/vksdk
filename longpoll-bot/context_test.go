package longpoll_test

import (
	"context"
	"testing"

	"github.com/Derad6709/vksdk/v2/internal"
	"github.com/Derad6709/vksdk/v2/longpoll-bot"
	"github.com/stretchr/testify/assert"
)

func TestTsFromContext(t *testing.T) {
	t.Parallel()

	const ts = 123
	ctx := context.WithValue(context.Background(), internal.LongPollTsKey, ts)
	assert.Equal(t, ts, longpoll.TsFromContext(ctx))
}
