// Package ssml implements Speech Synthesis Markup Language.
package ssml // import "github.com/Derad6709/vksdk/v2/marusia/ssml"

import (
	"bytes"
	"strconv"
	"time"
)

// Builder SSML.
type Builder struct {
	buf bytes.Buffer
}

// NewBuilder возвращает *Builder.
func NewBuilder() *Builder {
	return &Builder{}
}

// Say используется для произношения слов. Пропускает теги внутрь.
func (b *Builder) Say(s string) *Builder {
	_, _ = b.buf.WriteString(s)
	return b
}

// Paragraph используется для выделения абзацев. После абзаца ставится длинная интонационная пауза.
func (b *Builder) Paragraph(s string) *Builder {
	return b.Say(`<p>`).Say(s).Say(`</p>`)
}

// Sentence используется для выделения предложений. Предложение будет
// выделено интонационно, а в конце предложения будет пауза.
func (b *Builder) Sentence(s string) *Builder {
	return b.Say(`<s>`).Say(s).Say(`</s>`)
}

// Break используется для вставки пауз между произносимым текстом.
func (b *Builder) Break(duration time.Duration) *Builder {
	return b.Say(`<break time="`).Say(strconv.FormatInt(duration.Milliseconds(), 10)).Say(`ms"/>`)
}

// Speaker используется для вставки собственных аудиозаписей в текст.
func (b *Builder) Speaker(s string) *Builder {
	return b.Say(`<speaker audio_vk_id="`).Say(s).Say(`"/>`)
}

// SpeakerFromLibrary используется для вставки звуков из библиотеки в текст.
func (b *Builder) SpeakerFromLibrary(s string) *Builder {
	return b.Say(`<speaker audio="`).Say(s).Say(`"/>`)
}

// String возвращает SSML в виде строки.
func (b Builder) String() string {
	return "<speak>" + b.buf.String() + "</speak>"
}
