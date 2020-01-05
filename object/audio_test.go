package object_test

import (
	"testing"

	"github.com/SevereCloud/vksdk/object"
)

func TestAudioAudioFull_ToAttachment(t *testing.T) {
	f := func(audio object.AudioAudioFull, want string) {
		if got := audio.ToAttachment(); got != want {
			t.Errorf("AudioAudioFull.ToAttachment() = %v, want %v", got, want)
		}
	}

	f(object.AudioAudioFull{ID: 10, OwnerID: 20}, "audio20_10")
	f(object.AudioAudioFull{ID: 20, OwnerID: -10}, "audio-10_20")
}

func TestAudioAudio_ToAttachment(t *testing.T) {
	f := func(audio object.AudioAudio, want string) {
		if got := audio.ToAttachment(); got != want {
			t.Errorf("AudioAudio.ToAttachment() = %v, want %v", got, want)
		}
	}

	f(object.AudioAudio{ID: 10, OwnerID: 20}, "audio20_10")
	f(object.AudioAudio{ID: 20, OwnerID: -10}, "audio-10_20")
}
