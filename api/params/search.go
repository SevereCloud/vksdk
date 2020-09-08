package params // import "github.com/SevereCloud/vksdk/v2/api/params"

import (
	"github.com/SevereCloud/vksdk/v2/api"
)

// SearchGetHintsBuilder builder.
//
// Allows the programmer to do a quick search for any substring.
//
// https://vk.com/dev/search.getHints
type SearchGetHintsBuilder struct {
	api.Params
}

// NewSearchGetHintsBuilder func.
func NewSearchGetHintsBuilder() *SearchGetHintsBuilder {
	return &SearchGetHintsBuilder{api.Params{}}
}

// Q search query string.
func (b *SearchGetHintsBuilder) Q(v string) *SearchGetHintsBuilder {
	b.Params["q"] = v
	return b
}

// Offset for querying specific result subset.
func (b *SearchGetHintsBuilder) Offset(v int) *SearchGetHintsBuilder {
	b.Params["offset"] = v
	return b
}

// Limit maximum number of results to return.
func (b *SearchGetHintsBuilder) Limit(v int) *SearchGetHintsBuilder {
	b.Params["limit"] = v
	return b
}

// Filters parameter.
func (b *SearchGetHintsBuilder) Filters(v []string) *SearchGetHintsBuilder {
	b.Params["filters"] = v
	return b
}

// Fields parameter.
func (b *SearchGetHintsBuilder) Fields(v []string) *SearchGetHintsBuilder {
	b.Params["fields"] = v
	return b
}

// SearchGlobal parameter.
func (b *SearchGetHintsBuilder) SearchGlobal(v bool) *SearchGetHintsBuilder {
	b.Params["search_global"] = v
	return b
}
