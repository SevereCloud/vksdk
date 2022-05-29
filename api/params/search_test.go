package params_test

import (
	"testing"

	"github.com/Derad6709/vksdk/v2/api/params"
	"github.com/stretchr/testify/assert"
)

func TestSearchGetHintsBuilder(t *testing.T) {
	t.Parallel()

	b := params.NewSearchGetHintsBuilder()

	b.Q("text")
	b.Offset(1)
	b.Limit(1)
	b.Filters([]string{"text"})
	b.Fields([]string{"text"})
	b.SearchGlobal(true)

	assert.Equal(t, b.Params["q"], "text")
	assert.Equal(t, b.Params["offset"], 1)
	assert.Equal(t, b.Params["limit"], 1)
	assert.Equal(t, b.Params["filters"], []string{"text"})
	assert.Equal(t, b.Params["fields"], []string{"text"})
	assert.Equal(t, b.Params["search_global"], true)
}
