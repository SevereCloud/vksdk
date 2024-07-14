package params_test

import (
	"testing"

	"github.com/SevereCloud/vksdk/v3/api/params"
	"github.com/stretchr/testify/assert"
)

func TestWidgetsGetCommentsBuilder(t *testing.T) {
	t.Parallel()

	b := params.NewWidgetsGetCommentsBuilder()

	b.WidgetAPIID(1)
	b.URL("text")
	b.PageID("text")
	b.Order("text")
	b.Fields([]string{"test"})
	b.Offset(1)
	b.Count(1)

	assert.Equal(t, 1, b.Params["widget_api_id"])
	assert.Equal(t, "text", b.Params["url"])
	assert.Equal(t, "text", b.Params["page_id"])
	assert.Equal(t, "text", b.Params["order"])
	assert.Equal(t, []string{"test"}, b.Params["fields"])
	assert.Equal(t, 1, b.Params["offset"])
	assert.Equal(t, 1, b.Params["count"])
}

func TestWidgetsGetPagesBuilder(t *testing.T) {
	t.Parallel()

	b := params.NewWidgetsGetPagesBuilder()

	b.WidgetAPIID(1)
	b.Order("text")
	b.Period("text")
	b.Offset(1)
	b.Count(1)

	assert.Equal(t, 1, b.Params["widget_api_id"])
	assert.Equal(t, "text", b.Params["order"])
	assert.Equal(t, "text", b.Params["period"])
	assert.Equal(t, 1, b.Params["offset"])
	assert.Equal(t, 1, b.Params["count"])
}
