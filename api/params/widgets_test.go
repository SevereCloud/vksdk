package params_test

import (
	"testing"

	"github.com/SevereCloud/vksdk/api/params"
	"github.com/stretchr/testify/assert"
)

func TestWidgetsGetCommentsBuilder(t *testing.T) {
	b := params.NewWidgetsGetCommentsBuilder()

	b.WidgetAPIID(1)
	b.URL("text")
	b.PageID("text")
	b.Order("text")
	b.Fields([]string{"test"})
	b.Offset(1)
	b.Count(1)

	assert.Equal(t, b.Params["widget_api_id"], 1)
	assert.Equal(t, b.Params["url"], "text")
	assert.Equal(t, b.Params["page_id"], "text")
	assert.Equal(t, b.Params["order"], "text")
	assert.Equal(t, b.Params["fields"], []string{"test"})
	assert.Equal(t, b.Params["offset"], 1)
	assert.Equal(t, b.Params["count"], 1)
}

func TestWidgetsGetPagesBuilder(t *testing.T) {
	b := params.NewWidgetsGetPagesBuilder()

	b.WidgetAPIID(1)
	b.Order("text")
	b.Period("text")
	b.Offset(1)
	b.Count(1)

	assert.Equal(t, b.Params["widget_api_id"], 1)
	assert.Equal(t, b.Params["order"], "text")
	assert.Equal(t, b.Params["period"], "text")
	assert.Equal(t, b.Params["offset"], 1)
	assert.Equal(t, b.Params["count"], 1)
}
