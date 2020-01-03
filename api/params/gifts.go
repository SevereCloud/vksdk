package params // import "github.com/SevereCloud/vksdk/api/params"

import (
	"github.com/SevereCloud/vksdk/api"
)

// GiftsGetBulder builder
//
// Returns a list of user gifts.
//
// https://vk.com/dev/gifts.get
type GiftsGetBulder struct {
	api.Params
}

// NewGiftsGetBulder func
func NewGiftsGetBulder() *GiftsGetBulder {
	return &GiftsGetBulder{api.Params{}}
}

// UserID User ID.
func (b *GiftsGetBulder) UserID(v int) {
	b.Params["user_id"] = v
}

// Count Number of gifts to return.
func (b *GiftsGetBulder) Count(v int) {
	b.Params["count"] = v
}

// Offset Offset needed to return a specific subset of results.
func (b *GiftsGetBulder) Offset(v int) {
	b.Params["offset"] = v
}
