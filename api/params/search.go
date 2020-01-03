package params // import "github.com/SevereCloud/vksdk/api/params"

import (
	"github.com/SevereCloud/vksdk/api"
)

// SearchGetHintsBulder builder
//
// Allows the programmer to do a quick search for any substring.
//
// https://vk.com/dev/search.getHints
type SearchGetHintsBulder struct {
	api.Params
}

// NewSearchGetHintsBulder func
func NewSearchGetHintsBulder() *SearchGetHintsBulder {
	return &SearchGetHintsBulder{api.Params{}}
}

// Q Search query string.
func (b *SearchGetHintsBulder) Q(v string) {
	b.Params["q"] = v
}

// Offset Offset for querying specific result subset
func (b *SearchGetHintsBulder) Offset(v int) {
	b.Params["offset"] = v
}

// Limit Maximum number of results to return.
func (b *SearchGetHintsBulder) Limit(v int) {
	b.Params["limit"] = v
}

// Filters parameter
func (b *SearchGetHintsBulder) Filters(v []string) {
	b.Params["filters"] = v
}

// Fields parameter
func (b *SearchGetHintsBulder) Fields(v []string) {
	b.Params["fields"] = v
}

// SearchGlobal parameter
func (b *SearchGetHintsBulder) SearchGlobal(v bool) {
	b.Params["search_global"] = v
}
