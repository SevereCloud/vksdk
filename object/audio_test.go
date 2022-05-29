package object_test

import (
	"testing"

	"github.com/Derad6709/vksdk/v2/object"
)

func TestAudioAudio_ToAttachment(t *testing.T) {
	t.Parallel()

	f := func(audio object.AudioAudio, want string) {
		if got := audio.ToAttachment(); got != want {
			t.Errorf("AudioAudio.ToAttachment() = %v, want %v", got, want)
		}
	}

	f(object.AudioAudio{ID: 10, OwnerID: 20}, "audio20_10")
	f(object.AudioAudio{ID: 20, OwnerID: -10}, "audio-10_20")
}
