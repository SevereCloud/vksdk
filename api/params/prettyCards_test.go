package params_test

import (
	"testing"

	"github.com/SevereCloud/vksdk/v2/api/params"
	"github.com/stretchr/testify/assert"
)

func TestPrettyCardsCreateBuilder(t *testing.T) {
	t.Parallel()

	b := params.NewPrettyCardsCreateBuilder()

	b.OwnerID(1)
	b.Photo("text")
	b.Title("text")
	b.Link("text")
	b.Price("text")
	b.PriceOld("text")
	b.Button("text")

	assert.Equal(t, 1, b.Params["owner_id"])
	assert.Equal(t, "text", b.Params["photo"])
	assert.Equal(t, "text", b.Params["title"])
	assert.Equal(t, "text", b.Params["link"])
	assert.Equal(t, "text", b.Params["price"])
	assert.Equal(t, "text", b.Params["price_old"])
	assert.Equal(t, "text", b.Params["button"])
}

func TestPrettyCardsDeleteBuilder(t *testing.T) {
	t.Parallel()

	b := params.NewPrettyCardsDeleteBuilder()

	b.OwnerID(1)
	b.CardID(1)

	assert.Equal(t, 1, b.Params["owner_id"])
	assert.Equal(t, 1, b.Params["card_id"])
}

func TestPrettyCardsEditBuilder(t *testing.T) {
	t.Parallel()

	b := params.NewPrettyCardsEditBuilder()

	b.OwnerID(1)
	b.CardID(1)
	b.Photo("text")
	b.Title("text")
	b.Link("text")
	b.Price("text")
	b.PriceOld("text")
	b.Button("text")

	assert.Equal(t, 1, b.Params["owner_id"])
	assert.Equal(t, 1, b.Params["card_id"])
	assert.Equal(t, "text", b.Params["photo"])
	assert.Equal(t, "text", b.Params["title"])
	assert.Equal(t, "text", b.Params["link"])
	assert.Equal(t, "text", b.Params["price"])
	assert.Equal(t, "text", b.Params["price_old"])
	assert.Equal(t, "text", b.Params["button"])
}

func TestPrettyCardsGetBuilder(t *testing.T) {
	t.Parallel()

	b := params.NewPrettyCardsGetBuilder()

	b.OwnerID(1)
	b.Offset(1)
	b.Count(1)

	assert.Equal(t, 1, b.Params["owner_id"])
	assert.Equal(t, 1, b.Params["offset"])
	assert.Equal(t, 1, b.Params["count"])
}

func TestPrettyCardsGetByIDBuilder(t *testing.T) {
	t.Parallel()

	b := params.NewPrettyCardsGetByIDBuilder()

	b.OwnerID(1)
	b.CardIDs([]int{1})

	assert.Equal(t, 1, b.Params["owner_id"])
	assert.Equal(t, []int{1}, b.Params["card_ids"])
}
