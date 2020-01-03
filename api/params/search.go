package params // import "github.com/SevereCloud/vksdk/api/params"

import (
	"github.com/SevereCloud/vksdk/api"
)

// SearchGetHintsBuilder builder
//
// Allows the programmer to do a quick search for any substring.
//
// https://vk.com/dev/search.getHints
type SearchGetHintsBuilder struct {
	api.Params
}

// NewSearchGetHintsBuilder func
func NewSearchGetHintsBuilder() *SearchGetHintsBuilder {
	return &SearchGetHintsBuilder{api.Params{}}
}

// Q Search query string.
func (b *SearchGetHintsBuilder) Q(v string) {
	b.Params["q"] = v
}

// Offset Offset for querying specific result subset
func (b *SearchGetHintsBuilder) Offset(v int) {
	b.Params["offset"] = v
}

// Limit Maximum number of results to return.
func (b *SearchGetHintsBuilder) Limit(v int) {
	b.Params["limit"] = v
}

// Filters parameter
func (b *SearchGetHintsBuilder) Filters(v []string) {
	b.Params["filters"] = v
}

// Fields parameter
func (b *SearchGetHintsBuilder) Fields(v []string) {
	b.Params["fields"] = v
}

// SearchGlobal parameter
func (b *SearchGetHintsBuilder) SearchGlobal(v bool) {
	b.Params["search_global"] = v
}
