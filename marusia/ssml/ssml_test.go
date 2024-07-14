package ssml_test

import (
	"fmt"
	"time"

	"github.com/SevereCloud/vksdk/v3/marusia/ssml"
)

func ExampleNewBuilder() {
	b := ssml.NewBuilder()
	b.Say("Привет")
	b.Break(500 * time.Millisecond).Say("Я Маруся")

	fmt.Print(b)
	// Output: <speak>Привет<break time="500ms"/>Я Маруся</speak>
}

func ExampleBuilder_Say() {
	b := ssml.NewBuilder()
	b.Say("Привет")

	fmt.Print(b)
	// Output: <speak>Привет</speak>
}

func ExampleBuilder_Paragraph() {
	b := ssml.NewBuilder()
	b.Paragraph("Привет")

	fmt.Print(b)
	// Output: <speak><p>Привет</p></speak>
}

func ExampleBuilder_Sentence() {
	b := ssml.NewBuilder()
	b.Sentence("Привет")

	fmt.Print(b)
	// Output: <speak><s>Привет</s></speak>
}

func ExampleBuilder_Break() {
	b := ssml.NewBuilder()
	b.Say("Привет").Break(500 * time.Millisecond).Say("Я Маруся")

	fmt.Print(b)
	// Output: <speak>Привет<break time="500ms"/>Я Маруся</speak>
}

func ExampleBuilder_Speaker() {
	b := ssml.NewBuilder()
	b.Say("Угадайте, чей это голос?").Speaker("-2000000002_123456789")

	fmt.Print(b)
	// Output: <speak>Угадайте, чей это голос?<speaker audio_vk_id="-2000000002_123456789"/></speak>
}

func ExampleBuilder_SpeakerFromLibrary() {
	b := ssml.NewBuilder()
	b.Say("Так мычит корова").SpeakerFromLibrary("marusia-sounds/animals-cow-1")

	fmt.Print(b)
	// Output: <speak>Так мычит корова<speaker audio="marusia-sounds/animals-cow-1"/></speak>
}
