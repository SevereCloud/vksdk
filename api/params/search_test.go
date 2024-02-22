package params_test

import (
	"testing"

	"github.com/SevereCloud/vksdk/v2/api/params"
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

	assert.Equal(t, "text", b.Params["q"])
	assert.Equal(t, 1, b.Params["offset"])
	assert.Equal(t, 1, b.Params["limit"])
	assert.Equal(t, []string{"text"}, b.Params["filters"])
	assert.Equal(t, []string{"text"}, b.Params["fields"])
	assert.Equal(t, true, b.Params["search_global"])
}
