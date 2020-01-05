package params // import "github.com/SevereCloud/vksdk/api/params"

import (
	"github.com/SevereCloud/vksdk/api"
)

// WidgetsGetCommentsBuilder builder
//
// Gets a list of comments for the page added through the [vk.com/dev/Comments|Comments widget].
//
// https://vk.com/dev/widgets.getComments
type WidgetsGetCommentsBuilder struct {
	api.Params
}

// NewWidgetsGetCommentsBuilder func
func NewWidgetsGetCommentsBuilder() *WidgetsGetCommentsBuilder {
	return &WidgetsGetCommentsBuilder{api.Params{}}
}

// WidgetAPIID parameter
func (b *WidgetsGetCommentsBuilder) WidgetAPIID(v int) {
	b.Params["widget_api_id"] = v
}

// URL parameter
func (b *WidgetsGetCommentsBuilder) URL(v string) {
	b.Params["url"] = v
}

// PageID parameter
func (b *WidgetsGetCommentsBuilder) PageID(v string) {
	b.Params["page_id"] = v
}

// Order parameter
func (b *WidgetsGetCommentsBuilder) Order(v string) {
	b.Params["order"] = v
}

// Fields parameter
func (b *WidgetsGetCommentsBuilder) Fields(v []string) {
	b.Params["fields"] = v
}

// Offset parameter
func (b *WidgetsGetCommentsBuilder) Offset(v int) {
	b.Params["offset"] = v
}

// Count parameter
func (b *WidgetsGetCommentsBuilder) Count(v int) {
	b.Params["count"] = v
}

// WidgetsGetPagesBuilder builder
//
// Gets a list of application/site pages where the [vk.com/dev/Comments|Comments widget]
// or [vk.com/dev/Like|Like widget] is installed.
//
// https://vk.com/dev/widgets.getPages
type WidgetsGetPagesBuilder struct {
	api.Params
}

// NewWidgetsGetPagesBuilder func
func NewWidgetsGetPagesBuilder() *WidgetsGetPagesBuilder {
	return &WidgetsGetPagesBuilder{api.Params{}}
}

// WidgetAPIID parameter
func (b *WidgetsGetPagesBuilder) WidgetAPIID(v int) {
	b.Params["widget_api_id"] = v
}

// Order parameter
func (b *WidgetsGetPagesBuilder) Order(v string) {
	b.Params["order"] = v
}

// Period parameter
func (b *WidgetsGetPagesBuilder) Period(v string) {
	b.Params["period"] = v
}

// Offset parameter
func (b *WidgetsGetPagesBuilder) Offset(v int) {
	b.Params["offset"] = v
}

// Count parameter
func (b *WidgetsGetPagesBuilder) Count(v int) {
	b.Params["count"] = v
}
