package params // import "github.com/SevereCloud/vksdk/v2/api/params"

import (
	"github.com/SevereCloud/vksdk/v2/api"
)

// PrettyCardsCreateBuilder builder.
//
// https://vk.com/dev/prettyCards.create
type PrettyCardsCreateBuilder struct {
	api.Params
}

// NewPrettyCardsCreateBuilder func.
func NewPrettyCardsCreateBuilder() *PrettyCardsCreateBuilder {
	return &PrettyCardsCreateBuilder{api.Params{}}
}

// OwnerID parameter.
func (b *PrettyCardsCreateBuilder) OwnerID(v int) *PrettyCardsCreateBuilder {
	b.Params["owner_id"] = v
	return b
}

// Photo parameter.
func (b *PrettyCardsCreateBuilder) Photo(v string) *PrettyCardsCreateBuilder {
	b.Params["photo"] = v
	return b
}

// Title parameter.
func (b *PrettyCardsCreateBuilder) Title(v string) *PrettyCardsCreateBuilder {
	b.Params["title"] = v
	return b
}

// Link parameter.
func (b *PrettyCardsCreateBuilder) Link(v string) *PrettyCardsCreateBuilder {
	b.Params["link"] = v
	return b
}

// Price parameter.
func (b *PrettyCardsCreateBuilder) Price(v string) *PrettyCardsCreateBuilder {
	b.Params["price"] = v
	return b
}

// PriceOld parameter.
func (b *PrettyCardsCreateBuilder) PriceOld(v string) *PrettyCardsCreateBuilder {
	b.Params["price_old"] = v
	return b
}

// Button parameter.
func (b *PrettyCardsCreateBuilder) Button(v string) *PrettyCardsCreateBuilder {
	b.Params["button"] = v
	return b
}

// PrettyCardsDeleteBuilder builder.
//
// https://vk.com/dev/prettyCards.delete
type PrettyCardsDeleteBuilder struct {
	api.Params
}

// NewPrettyCardsDeleteBuilder func.
func NewPrettyCardsDeleteBuilder() *PrettyCardsDeleteBuilder {
	return &PrettyCardsDeleteBuilder{api.Params{}}
}

// OwnerID parameter.
func (b *PrettyCardsDeleteBuilder) OwnerID(v int) *PrettyCardsDeleteBuilder {
	b.Params["owner_id"] = v
	return b
}

// CardID parameter.
func (b *PrettyCardsDeleteBuilder) CardID(v int) *PrettyCardsDeleteBuilder {
	b.Params["card_id"] = v
	return b
}

// PrettyCardsEditBuilder builder.
//
// https://vk.com/dev/prettyCards.edit
type PrettyCardsEditBuilder struct {
	api.Params
}

// NewPrettyCardsEditBuilder func.
func NewPrettyCardsEditBuilder() *PrettyCardsEditBuilder {
	return &PrettyCardsEditBuilder{api.Params{}}
}

// OwnerID parameter.
func (b *PrettyCardsEditBuilder) OwnerID(v int) *PrettyCardsEditBuilder {
	b.Params["owner_id"] = v
	return b
}

// CardID parameter.
func (b *PrettyCardsEditBuilder) CardID(v int) *PrettyCardsEditBuilder {
	b.Params["card_id"] = v
	return b
}

// Photo parameter.
func (b *PrettyCardsEditBuilder) Photo(v string) *PrettyCardsEditBuilder {
	b.Params["photo"] = v
	return b
}

// Title parameter.
func (b *PrettyCardsEditBuilder) Title(v string) *PrettyCardsEditBuilder {
	b.Params["title"] = v
	return b
}

// Link parameter.
func (b *PrettyCardsEditBuilder) Link(v string) *PrettyCardsEditBuilder {
	b.Params["link"] = v
	return b
}

// Price parameter.
func (b *PrettyCardsEditBuilder) Price(v string) *PrettyCardsEditBuilder {
	b.Params["price"] = v
	return b
}

// PriceOld parameter.
func (b *PrettyCardsEditBuilder) PriceOld(v string) *PrettyCardsEditBuilder {
	b.Params["price_old"] = v
	return b
}

// Button parameter.
func (b *PrettyCardsEditBuilder) Button(v string) *PrettyCardsEditBuilder {
	b.Params["button"] = v
	return b
}

// PrettyCardsGetBuilder builder.
//
// https://vk.com/dev/prettyCards.get
type PrettyCardsGetBuilder struct {
	api.Params
}

// NewPrettyCardsGetBuilder func.
func NewPrettyCardsGetBuilder() *PrettyCardsGetBuilder {
	return &PrettyCardsGetBuilder{api.Params{}}
}

// OwnerID parameter.
func (b *PrettyCardsGetBuilder) OwnerID(v int) *PrettyCardsGetBuilder {
	b.Params["owner_id"] = v
	return b
}

// Offset parameter.
func (b *PrettyCardsGetBuilder) Offset(v int) *PrettyCardsGetBuilder {
	b.Params["offset"] = v
	return b
}

// Count parameter.
func (b *PrettyCardsGetBuilder) Count(v int) *PrettyCardsGetBuilder {
	b.Params["count"] = v
	return b
}

// PrettyCardsGetByIDBuilder builder.
//
// https://vk.com/dev/prettyCards.getById
type PrettyCardsGetByIDBuilder struct {
	api.Params
}

// NewPrettyCardsGetByIDBuilder func.
func NewPrettyCardsGetByIDBuilder() *PrettyCardsGetByIDBuilder {
	return &PrettyCardsGetByIDBuilder{api.Params{}}
}

// OwnerID parameter.
func (b *PrettyCardsGetByIDBuilder) OwnerID(v int) *PrettyCardsGetByIDBuilder {
	b.Params["owner_id"] = v
	return b
}

// CardIDs parameter.
func (b *PrettyCardsGetByIDBuilder) CardIDs(v []int) *PrettyCardsGetByIDBuilder {
	b.Params["card_ids"] = v
	return b
}
