package params // import "github.com/SevereCloud/vksdk/v2/api/params"

import (
	"github.com/SevereCloud/vksdk/v2/api"
)

// WidgetsGetCommentsBuilder builder.
//
// Gets a list of comments for the page added through the [vk.com/dev/Comments|Comments widget].
//
// https://vk.com/dev/widgets.getComments
type WidgetsGetCommentsBuilder struct {
	api.Params
}

// NewWidgetsGetCommentsBuilder func.
func NewWidgetsGetCommentsBuilder() *WidgetsGetCommentsBuilder {
	return &WidgetsGetCommentsBuilder{api.Params{}}
}

// WidgetAPIID parameter.
func (b *WidgetsGetCommentsBuilder) WidgetAPIID(v int) *WidgetsGetCommentsBuilder {
	b.Params["widget_api_id"] = v
	return b
}

// URL parameter.
func (b *WidgetsGetCommentsBuilder) URL(v string) *WidgetsGetCommentsBuilder {
	b.Params["url"] = v
	return b
}

// PageID parameter.
func (b *WidgetsGetCommentsBuilder) PageID(v string) *WidgetsGetCommentsBuilder {
	b.Params["page_id"] = v
	return b
}

// Order parameter.
func (b *WidgetsGetCommentsBuilder) Order(v string) *WidgetsGetCommentsBuilder {
	b.Params["order"] = v
	return b
}

// Fields parameter.
func (b *WidgetsGetCommentsBuilder) Fields(v []string) *WidgetsGetCommentsBuilder {
	b.Params["fields"] = v
	return b
}

// Offset parameter.
func (b *WidgetsGetCommentsBuilder) Offset(v int) *WidgetsGetCommentsBuilder {
	b.Params["offset"] = v
	return b
}

// Count parameter.
func (b *WidgetsGetCommentsBuilder) Count(v int) *WidgetsGetCommentsBuilder {
	b.Params["count"] = v
	return b
}

// WidgetsGetPagesBuilder builder.
//
// Gets a list of application/site pages where the [vk.com/dev/Comments|Comments widget]
// or [vk.com/dev/Like|Like widget] is installed.
//
// https://vk.com/dev/widgets.getPages
type WidgetsGetPagesBuilder struct {
	api.Params
}

// NewWidgetsGetPagesBuilder func.
func NewWidgetsGetPagesBuilder() *WidgetsGetPagesBuilder {
	return &WidgetsGetPagesBuilder{api.Params{}}
}

// WidgetAPIID parameter.
func (b *WidgetsGetPagesBuilder) WidgetAPIID(v int) *WidgetsGetPagesBuilder {
	b.Params["widget_api_id"] = v
	return b
}

// Order parameter.
func (b *WidgetsGetPagesBuilder) Order(v string) *WidgetsGetPagesBuilder {
	b.Params["order"] = v
	return b
}

// Period parameter.
func (b *WidgetsGetPagesBuilder) Period(v string) *WidgetsGetPagesBuilder {
	b.Params["period"] = v
	return b
}

// Offset parameter.
func (b *WidgetsGetPagesBuilder) Offset(v int) *WidgetsGetPagesBuilder {
	b.Params["offset"] = v
	return b
}

// Count parameter.
func (b *WidgetsGetPagesBuilder) Count(v int) *WidgetsGetPagesBuilder {
	b.Params["count"] = v
	return b
}
