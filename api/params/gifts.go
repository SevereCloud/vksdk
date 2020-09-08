package params // import "github.com/SevereCloud/vksdk/v2/api/params"

import (
	"github.com/SevereCloud/vksdk/v2/api"
)

// GiftsGetBuilder builder.
//
// Returns a list of user gifts.
//
// https://vk.com/dev/gifts.get
type GiftsGetBuilder struct {
	api.Params
}

// NewGiftsGetBuilder func.
func NewGiftsGetBuilder() *GiftsGetBuilder {
	return &GiftsGetBuilder{api.Params{}}
}

// UserID parameter.
func (b *GiftsGetBuilder) UserID(v int) *GiftsGetBuilder {
	b.Params["user_id"] = v
	return b
}

// Count number of gifts to return.
func (b *GiftsGetBuilder) Count(v int) *GiftsGetBuilder {
	b.Params["count"] = v
	return b
}

// Offset needed to return a specific subset of results.
func (b *GiftsGetBuilder) Offset(v int) *GiftsGetBuilder {
	b.Params["offset"] = v
	return b
}
