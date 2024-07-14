package object_test

import (
	"testing"

	"github.com/SevereCloud/vksdk/v3/object"

	"github.com/stretchr/testify/assert"
)

func TestClickableStickers(t *testing.T) {
	t.Parallel()

	const (
		height  = 1920
		width   = 1080
		hashtag = "#hashtage"
		mention = "[club1|@apiclub]"
	)

	cs := object.NewClickableStickers(width, height)
	assert.Equal(t, width, cs.OriginalWidth)
	assert.Equal(t, height, cs.OriginalHeight)

	cs.AddHashtag(hashtag, []object.StoriesClickablePoint{})

	if assert.NotEmpty(t, cs.ClickableStickers) {
		assert.Equal(t, object.ClickableStickerHashtag, cs.ClickableStickers[0].Type)
		assert.Equal(t, hashtag, cs.ClickableStickers[0].Hashtag)
	}

	cs.AddMention(mention, []object.StoriesClickablePoint{})

	if assert.NotEmpty(t, cs.ClickableStickers) {
		assert.Equal(t, object.ClickableStickerMention, cs.ClickableStickers[1].Type)
		assert.Equal(t, mention, cs.ClickableStickers[1].Mention)
	}

	json := cs.ToJSON()
	assert.NotEmpty(t, json)
}
