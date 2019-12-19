package object_test

import (
	"testing"

	"github.com/SevereCloud/vksdk/object"
)

func TestPollsPoll_ToAttachment(t *testing.T) {
	f := func(poll object.PollsPoll, want string) {
		if got := poll.ToAttachment(); got != want {
			t.Errorf("PollsPoll.ToAttachment() = %v, want %v", got, want)
		}
	}

	f(object.PollsPoll{ID: 10, OwnerID: 20}, "poll20_10")
	f(object.PollsPoll{ID: 20, OwnerID: -10}, "poll-10_20")
}
