package marusia_test

import (
	"testing"

	"github.com/SevereCloud/vksdk/v3/marusia"
	"github.com/stretchr/testify/assert"
)

func TestNewBigImage(t *testing.T) {
	t.Parallel()

	f := func(title, description string, imageID int, actual *marusia.Card) {
		t.Helper()

		card := marusia.NewBigImage(title, description, imageID)

		assert.Equal(t, card, actual)
	}

	f("title", "description", 1234,
		&marusia.Card{
			Type:        marusia.BigImage,
			Title:       "title",
			Description: "description",
			ImageID:     1234,
		},
	)
}

func TestNewItemsList(t *testing.T) {
	t.Parallel()

	f := func(items []marusia.CardItem, actual *marusia.Card) {
		t.Helper()

		card := marusia.NewItemsList(items...)

		assert.Equal(t, card, actual)
	}

	f([]marusia.CardItem{{1}, {2}},
		&marusia.Card{
			Type:  marusia.ItemsList,
			Items: []marusia.CardItem{{1}, {2}},
		},
	)
}

func TestNewMiniApp(t *testing.T) {
	t.Parallel()

	f := func(url string, actual *marusia.Card) {
		t.Helper()

		card := marusia.NewMiniApp(url)

		assert.Equal(t, card, actual)
	}

	f("url",
		&marusia.Card{
			Type: marusia.MiniApp,
			URL:  "url",
		},
	)
}

func TestNewLink(t *testing.T) {
	t.Parallel()

	f := func(url, title, text string, imageID int, actual *marusia.Card) {
		t.Helper()

		card := marusia.NewLink(url, title, text, imageID)

		assert.Equal(t, card, actual)
	}

	f("url", "title", "text", 1234,
		&marusia.Card{
			Type:    marusia.Link,
			URL:     "url",
			Title:   "title",
			Text:    "text",
			ImageID: 1234,
		},
	)
}

func TestNewImageList(t *testing.T) {
	t.Parallel()

	f := func(items []int, actual *marusia.Card) {
		t.Helper()

		card := marusia.NewImageList(items...)

		assert.Equal(t, card, actual)
	}

	f([]int{1, 2},
		&marusia.Card{
			Type:  marusia.ItemsList,
			Items: []marusia.CardItem{{1}, {2}},
		},
	)
}
