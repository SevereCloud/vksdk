package marusia_test

import (
	"fmt"
	"testing"

	"github.com/SevereCloud/vksdk/v3/marusia"
	"github.com/stretchr/testify/assert"
)

type my struct {
	Name string `json:"name"`
}

func TestInterfaces_IsScreen(t *testing.T) {
	t.Parallel()

	f := func(i marusia.Interfaces, actual bool) {
		t.Helper()

		assert.Equal(t, i.IsScreen(), actual)
	}

	f(marusia.Interfaces{}, false)
	f(marusia.Interfaces{
		Screen: nil,
	}, false)
	f(marusia.Interfaces{
		Screen: &marusia.Screen{},
	}, true)
}

func TestResponse_AddURL(t *testing.T) {
	t.Parallel()

	f := func(title string, link string, actual marusia.Response) {
		t.Helper()

		resp := marusia.Response{}
		resp.AddURL(title, link)

		assert.Equal(t, resp, actual)
	}

	f("title", "https://vk.com",
		marusia.Response{
			Buttons: []marusia.Button{
				{
					Title: "title",
					URL:   "https://vk.com",
				},
			},
		},
	)
}

func TestResponse_AddButton(t *testing.T) {
	t.Parallel()

	f := func(title string, payload interface{}, actual marusia.Response) {
		t.Helper()

		resp := marusia.Response{}
		resp.AddButton(title, payload)

		assert.Equal(t, resp, actual)
	}

	f(
		"title",
		nil,
		marusia.Response{
			Buttons: []marusia.Button{
				{
					Title: "title",
				},
			},
		},
	)
	f(
		"title",
		my{"test"},
		marusia.Response{
			Buttons: []marusia.Button{
				{
					Title:   "title",
					Payload: my{"test"},
				},
			},
		},
	)
}

func TestSpeakerAudioVKID(t *testing.T) {
	t.Parallel()

	assert.Equal(
		t,
		"<speaker audio_vk_id=\"-2000000002_123456789\">",
		marusia.SpeakerAudioVKID("-2000000002_123456789"),
	)
}

func TestSpeakerAudio(t *testing.T) {
	t.Parallel()

	assert.Equal(
		t,
		"<speaker audio=\"marusia-sounds/game-win-1\">",
		marusia.SpeakerAudio("marusia-sounds/game-win-1"),
	)
}

func ExampleSpeakerAudioVKID() {
	tts := "Угадайте, чей это голос? " + marusia.SpeakerAudioVKID("-2000000002_123456789")
	fmt.Println(tts)

	// Output:
	// Угадайте, чей это голос? <speaker audio_vk_id="-2000000002_123456789">
}

func ExampleSpeakerAudio() {
	tts := fmt.Sprintf(
		"Поздравляю! %s Вы правильно ответили на все мои вопросы!",
		marusia.SpeakerAudio("marusia-sounds/game-win-1"),
	)
	fmt.Println(tts)

	// Output:
	// Поздравляю! <speaker audio="marusia-sounds/game-win-1"> Вы правильно ответили на все мои вопросы!
}
