package longpoll_test

import (
	"testing"

	"github.com/SevereCloud/vksdk/longpoll-user"
	"github.com/stretchr/testify/assert"
)

func TestFailed_Error(t *testing.T) {
	err := longpoll.Failed{1}
	assert.EqualError(t, err, "longpoll: failed code 1")
}
