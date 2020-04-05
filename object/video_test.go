package object_test

import (
	"testing"

	"github.com/SevereCloud/vksdk/object"
)

func TestVideoVideo_ToAttachment(t *testing.T) {
	t.Parallel()

	f := func(video object.VideoVideo, want string) {
		if got := video.ToAttachment(); got != want {
			t.Errorf("VideoVideo.ToAttachment() = %v, want %v", got, want)
		}
	}

	f(object.VideoVideo{ID: 10, OwnerID: 20}, "video20_10")
	f(object.VideoVideo{ID: 20, OwnerID: -10}, "video-10_20")
}

func TestVideoVideoFull_ToAttachment(t *testing.T) {
	t.Parallel()

	f := func(video object.VideoVideoFull, want string) {
		if got := video.ToAttachment(); got != want {
			t.Errorf("VideoVideoFull.ToAttachment() = %v, want %v", got, want)
		}
	}

	f(object.VideoVideoFull{ID: 10, OwnerID: 20}, "video20_10")
	f(object.VideoVideoFull{ID: 20, OwnerID: -10}, "video-10_20")
}
