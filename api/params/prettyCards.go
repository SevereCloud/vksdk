package params // import "github.com/SevereCloud/vksdk/api/params"

import (
	"github.com/SevereCloud/vksdk/api"
)

// PrettyCardsCreateBuilder builder
//
// https://vk.com/dev/prettyCards.create
type PrettyCardsCreateBuilder struct {
	api.Params
}

// NewPrettyCardsCreateBuilder func
func NewPrettyCardsCreateBuilder() *PrettyCardsCreateBuilder {
	return &PrettyCardsCreateBuilder{api.Params{}}
}

// OwnerID parameter
func (b *PrettyCardsCreateBuilder) OwnerID(v int) {
	b.Params["owner_id"] = v
}

// Photo parameter
func (b *PrettyCardsCreateBuilder) Photo(v string) {
	b.Params["photo"] = v
}

// Title parameter
func (b *PrettyCardsCreateBuilder) Title(v string) {
	b.Params["title"] = v
}

// Link parameter
func (b *PrettyCardsCreateBuilder) Link(v string) {
	b.Params["link"] = v
}

// Price parameter
func (b *PrettyCardsCreateBuilder) Price(v string) {
	b.Params["price"] = v
}

// PriceOld parameter
func (b *PrettyCardsCreateBuilder) PriceOld(v string) {
	b.Params["price_old"] = v
}

// Button parameter
func (b *PrettyCardsCreateBuilder) Button(v string) {
	b.Params["button"] = v
}

// PrettyCardsDeleteBuilder builder
//
// https://vk.com/dev/prettyCards.delete
type PrettyCardsDeleteBuilder struct {
	api.Params
}

// NewPrettyCardsDeleteBuilder func
func NewPrettyCardsDeleteBuilder() *PrettyCardsDeleteBuilder {
	return &PrettyCardsDeleteBuilder{api.Params{}}
}

// OwnerID parameter
func (b *PrettyCardsDeleteBuilder) OwnerID(v int) {
	b.Params["owner_id"] = v
}

// CardID parameter
func (b *PrettyCardsDeleteBuilder) CardID(v int) {
	b.Params["card_id"] = v
}

// PrettyCardsEditBuilder builder
//
// https://vk.com/dev/prettyCards.edit
type PrettyCardsEditBuilder struct {
	api.Params
}

// NewPrettyCardsEditBuilder func
func NewPrettyCardsEditBuilder() *PrettyCardsEditBuilder {
	return &PrettyCardsEditBuilder{api.Params{}}
}

// OwnerID parameter
func (b *PrettyCardsEditBuilder) OwnerID(v int) {
	b.Params["owner_id"] = v
}

// CardID parameter
func (b *PrettyCardsEditBuilder) CardID(v int) {
	b.Params["card_id"] = v
}

// Photo parameter
func (b *PrettyCardsEditBuilder) Photo(v string) {
	b.Params["photo"] = v
}

// Title parameter
func (b *PrettyCardsEditBuilder) Title(v string) {
	b.Params["title"] = v
}

// Link parameter
func (b *PrettyCardsEditBuilder) Link(v string) {
	b.Params["link"] = v
}

// Price parameter
func (b *PrettyCardsEditBuilder) Price(v string) {
	b.Params["price"] = v
}

// PriceOld parameter
func (b *PrettyCardsEditBuilder) PriceOld(v string) {
	b.Params["price_old"] = v
}

// Button parameter
func (b *PrettyCardsEditBuilder) Button(v string) {
	b.Params["button"] = v
}

// PrettyCardsGetBuilder builder
//
// https://vk.com/dev/prettyCards.get
type PrettyCardsGetBuilder struct {
	api.Params
}

// NewPrettyCardsGetBuilder func
func NewPrettyCardsGetBuilder() *PrettyCardsGetBuilder {
	return &PrettyCardsGetBuilder{api.Params{}}
}

// OwnerID parameter
func (b *PrettyCardsGetBuilder) OwnerID(v int) {
	b.Params["owner_id"] = v
}

// Offset parameter
func (b *PrettyCardsGetBuilder) Offset(v int) {
	b.Params["offset"] = v
}

// Count parameter
func (b *PrettyCardsGetBuilder) Count(v int) {
	b.Params["count"] = v
}

// PrettyCardsGetByIDBuilder builder
//
// https://vk.com/dev/prettyCards.getById
type PrettyCardsGetByIDBuilder struct {
	api.Params
}

// NewPrettyCardsGetByIDBuilder func
func NewPrettyCardsGetByIDBuilder() *PrettyCardsGetByIDBuilder {
	return &PrettyCardsGetByIDBuilder{api.Params{}}
}

// OwnerID parameter
func (b *PrettyCardsGetByIDBuilder) OwnerID(v int) {
	b.Params["owner_id"] = v
}

// CardIDs parameter
func (b *PrettyCardsGetByIDBuilder) CardIDs(v []int) {
	b.Params["card_ids"] = v
}
