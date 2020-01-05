package params_test

import (
	"testing"

	"github.com/SevereCloud/vksdk/api/params"
	"github.com/stretchr/testify/assert"
)

func TestPrettyCardsCreateBuilder(t *testing.T) {
	b := params.NewPrettyCardsCreateBuilder()

	b.OwnerID(1)
	b.Photo("text")
	b.Title("text")
	b.Link("text")
	b.Price("text")
	b.PriceOld("text")
	b.Button("text")

	assert.Equal(t, b.Params["owner_id"], 1)
	assert.Equal(t, b.Params["photo"], "text")
	assert.Equal(t, b.Params["title"], "text")
	assert.Equal(t, b.Params["link"], "text")
	assert.Equal(t, b.Params["price"], "text")
	assert.Equal(t, b.Params["price_old"], "text")
	assert.Equal(t, b.Params["button"], "text")
}

func TestPrettyCardsDeleteBuilder(t *testing.T) {
	b := params.NewPrettyCardsDeleteBuilder()

	b.OwnerID(1)
	b.CardID(1)

	assert.Equal(t, b.Params["owner_id"], 1)
	assert.Equal(t, b.Params["card_id"], 1)
}

func TestPrettyCardsEditBuilder(t *testing.T) {
	b := params.NewPrettyCardsEditBuilder()

	b.OwnerID(1)
	b.CardID(1)
	b.Photo("text")
	b.Title("text")
	b.Link("text")
	b.Price("text")
	b.PriceOld("text")
	b.Button("text")

	assert.Equal(t, b.Params["owner_id"], 1)
	assert.Equal(t, b.Params["card_id"], 1)
	assert.Equal(t, b.Params["photo"], "text")
	assert.Equal(t, b.Params["title"], "text")
	assert.Equal(t, b.Params["link"], "text")
	assert.Equal(t, b.Params["price"], "text")
	assert.Equal(t, b.Params["price_old"], "text")
	assert.Equal(t, b.Params["button"], "text")
}

func TestPrettyCardsGetBuilder(t *testing.T) {
	b := params.NewPrettyCardsGetBuilder()

	b.OwnerID(1)
	b.Offset(1)
	b.Count(1)

	assert.Equal(t, b.Params["owner_id"], 1)
	assert.Equal(t, b.Params["offset"], 1)
	assert.Equal(t, b.Params["count"], 1)
}

func TestPrettyCardsGetByIDBuilder(t *testing.T) {
	b := params.NewPrettyCardsGetByIDBuilder()

	b.OwnerID(1)
	b.CardIDs([]int{1})

	assert.Equal(t, b.Params["owner_id"], 1)
	assert.Equal(t, b.Params["card_ids"], []int{1})
}
