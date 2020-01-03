package params // import "github.com/SevereCloud/vksdk/api/params"

import (
	"github.com/SevereCloud/vksdk/api"
)

// WidgetsGetCommentsBulder builder
//
// Gets a list of comments for the page added through the [vk.com/dev/Comments|Comments widget].
//
// https://vk.com/dev/widgets.getComments
type WidgetsGetCommentsBulder struct {
	api.Params
}

// NewWidgetsGetCommentsBulder func
func NewWidgetsGetCommentsBulder() *WidgetsGetCommentsBulder {
	return &WidgetsGetCommentsBulder{api.Params{}}
}

// WidgetAPIID parameter
func (b *WidgetsGetCommentsBulder) WidgetAPIID(v int) {
	b.Params["widget_api_id"] = v
}

// URL parameter
func (b *WidgetsGetCommentsBulder) URL(v string) {
	b.Params["url"] = v
}

// PageID parameter
func (b *WidgetsGetCommentsBulder) PageID(v string) {
	b.Params["page_id"] = v
}

// Order parameter
func (b *WidgetsGetCommentsBulder) Order(v string) {
	b.Params["order"] = v
}

// Fields parameter
func (b *WidgetsGetCommentsBulder) Fields(v []string) {
	b.Params["fields"] = v
}

// Offset parameter
func (b *WidgetsGetCommentsBulder) Offset(v int) {
	b.Params["offset"] = v
}

// Count parameter
func (b *WidgetsGetCommentsBulder) Count(v int) {
	b.Params["count"] = v
}

// WidgetsGetPagesBulder builder
//
// Gets a list of application/site pages where the [vk.com/dev/Comments|Comments widget] or [vk.com/dev/Like|Like widget] is installed.
//
// https://vk.com/dev/widgets.getPages
type WidgetsGetPagesBulder struct {
	api.Params
}

// NewWidgetsGetPagesBulder func
func NewWidgetsGetPagesBulder() *WidgetsGetPagesBulder {
	return &WidgetsGetPagesBulder{api.Params{}}
}

// WidgetAPIID parameter
func (b *WidgetsGetPagesBulder) WidgetAPIID(v int) {
	b.Params["widget_api_id"] = v
}

// Order parameter
func (b *WidgetsGetPagesBulder) Order(v string) {
	b.Params["order"] = v
}

// Period parameter
func (b *WidgetsGetPagesBulder) Period(v string) {
	b.Params["period"] = v
}

// Offset parameter
func (b *WidgetsGetPagesBulder) Offset(v int) {
	b.Params["offset"] = v
}

// Count parameter
func (b *WidgetsGetPagesBulder) Count(v int) {
	b.Params["count"] = v
}
