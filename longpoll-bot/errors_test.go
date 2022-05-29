package longpoll_test

import (
	"testing"

	"github.com/Derad6709/vksdk/v2/longpoll-bot"
	"github.com/stretchr/testify/assert"
)

func TestFailed_Error(t *testing.T) {
	t.Parallel()

	err := longpoll.Failed{1}
	assert.EqualError(t, err, "longpoll: failed code 1")
}
