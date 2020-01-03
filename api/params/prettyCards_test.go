package params_test

import (
	"testing"

	"github.com/SevereCloud/vksdk/api/params"
	"github.com/stretchr/testify/assert"
)

func TestPrettyCardsCreateBulder(t *testing.T) {
	b := params.NewPrettyCardsCreateBulder()

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

func TestPrettyCardsDeleteBulder(t *testing.T) {
	b := params.NewPrettyCardsDeleteBulder()

	b.OwnerID(1)
	b.CardID(1)

	assert.Equal(t, b.Params["owner_id"], 1)
	assert.Equal(t, b.Params["card_id"], 1)
}

func TestPrettyCardsEditBulder(t *testing.T) {
	b := params.NewPrettyCardsEditBulder()

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

func TestPrettyCardsGetBulder(t *testing.T) {
	b := params.NewPrettyCardsGetBulder()

	b.OwnerID(1)
	b.Offset(1)
	b.Count(1)

	assert.Equal(t, b.Params["owner_id"], 1)
	assert.Equal(t, b.Params["offset"], 1)
	assert.Equal(t, b.Params["count"], 1)
}

func TestPrettyCardsGetByIDBulder(t *testing.T) {
	b := params.NewPrettyCardsGetByIDBulder()

	b.OwnerID(1)
	b.CardIDs([]int{1})

	assert.Equal(t, b.Params["owner_id"], 1)
	assert.Equal(t, b.Params["card_ids"], []int{1})
}
