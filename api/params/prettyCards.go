package params // import "github.com/SevereCloud/vksdk/api/params"

import (
	"github.com/SevereCloud/vksdk/api"
)

// PrettyCardsCreateBulder builder
//
// https://vk.com/dev/prettyCards.create
type PrettyCardsCreateBulder struct {
	api.Params
}

// NewPrettyCardsCreateBulder func
func NewPrettyCardsCreateBulder() *PrettyCardsCreateBulder {
	return &PrettyCardsCreateBulder{api.Params{}}
}

// OwnerID parameter
func (b *PrettyCardsCreateBulder) OwnerID(v int) {
	b.Params["owner_id"] = v
}

// Photo parameter
func (b *PrettyCardsCreateBulder) Photo(v string) {
	b.Params["photo"] = v
}

// Title parameter
func (b *PrettyCardsCreateBulder) Title(v string) {
	b.Params["title"] = v
}

// Link parameter
func (b *PrettyCardsCreateBulder) Link(v string) {
	b.Params["link"] = v
}

// Price parameter
func (b *PrettyCardsCreateBulder) Price(v string) {
	b.Params["price"] = v
}

// PriceOld parameter
func (b *PrettyCardsCreateBulder) PriceOld(v string) {
	b.Params["price_old"] = v
}

// Button parameter
func (b *PrettyCardsCreateBulder) Button(v string) {
	b.Params["button"] = v
}

// PrettyCardsDeleteBulder builder
//
// https://vk.com/dev/prettyCards.delete
type PrettyCardsDeleteBulder struct {
	api.Params
}

// NewPrettyCardsDeleteBulder func
func NewPrettyCardsDeleteBulder() *PrettyCardsDeleteBulder {
	return &PrettyCardsDeleteBulder{api.Params{}}
}

// OwnerID parameter
func (b *PrettyCardsDeleteBulder) OwnerID(v int) {
	b.Params["owner_id"] = v
}

// CardID parameter
func (b *PrettyCardsDeleteBulder) CardID(v int) {
	b.Params["card_id"] = v
}

// PrettyCardsEditBulder builder
//
// https://vk.com/dev/prettyCards.edit
type PrettyCardsEditBulder struct {
	api.Params
}

// NewPrettyCardsEditBulder func
func NewPrettyCardsEditBulder() *PrettyCardsEditBulder {
	return &PrettyCardsEditBulder{api.Params{}}
}

// OwnerID parameter
func (b *PrettyCardsEditBulder) OwnerID(v int) {
	b.Params["owner_id"] = v
}

// CardID parameter
func (b *PrettyCardsEditBulder) CardID(v int) {
	b.Params["card_id"] = v
}

// Photo parameter
func (b *PrettyCardsEditBulder) Photo(v string) {
	b.Params["photo"] = v
}

// Title parameter
func (b *PrettyCardsEditBulder) Title(v string) {
	b.Params["title"] = v
}

// Link parameter
func (b *PrettyCardsEditBulder) Link(v string) {
	b.Params["link"] = v
}

// Price parameter
func (b *PrettyCardsEditBulder) Price(v string) {
	b.Params["price"] = v
}

// PriceOld parameter
func (b *PrettyCardsEditBulder) PriceOld(v string) {
	b.Params["price_old"] = v
}

// Button parameter
func (b *PrettyCardsEditBulder) Button(v string) {
	b.Params["button"] = v
}

// PrettyCardsGetBulder builder
//
// https://vk.com/dev/prettyCards.get
type PrettyCardsGetBulder struct {
	api.Params
}

// NewPrettyCardsGetBulder func
func NewPrettyCardsGetBulder() *PrettyCardsGetBulder {
	return &PrettyCardsGetBulder{api.Params{}}
}

// OwnerID parameter
func (b *PrettyCardsGetBulder) OwnerID(v int) {
	b.Params["owner_id"] = v
}

// Offset parameter
func (b *PrettyCardsGetBulder) Offset(v int) {
	b.Params["offset"] = v
}

// Count parameter
func (b *PrettyCardsGetBulder) Count(v int) {
	b.Params["count"] = v
}

// PrettyCardsGetByIDBulder builder
//
// https://vk.com/dev/prettyCards.getById
type PrettyCardsGetByIDBulder struct {
	api.Params
}

// NewPrettyCardsGetByIDBulder func
func NewPrettyCardsGetByIDBulder() *PrettyCardsGetByIDBulder {
	return &PrettyCardsGetByIDBulder{api.Params{}}
}

// OwnerID parameter
func (b *PrettyCardsGetByIDBulder) OwnerID(v int) {
	b.Params["owner_id"] = v
}

// CardIDs parameter
func (b *PrettyCardsGetByIDBulder) CardIDs(v []int) {
	b.Params["card_ids"] = v
}
