package params // import "github.com/SevereCloud/vksdk/v2/api/params"

import (
	"github.com/SevereCloud/vksdk/v2/api"
)

// AppWidgetsUpdateBuilder builder.
//
// Allows to update community app widget.
//
// https://vk.com/dev/appWidgets.update
type AppWidgetsUpdateBuilder struct {
	api.Params
}

// NewAppWidgetsUpdateBuilder func.
func NewAppWidgetsUpdateBuilder() *AppWidgetsUpdateBuilder {
	return &AppWidgetsUpdateBuilder{api.Params{}}
}

// Code parameter.
func (b *AppWidgetsUpdateBuilder) Code(v string) *AppWidgetsUpdateBuilder {
	b.Params["code"] = v
	return b
}

// Type parameter.
func (b *AppWidgetsUpdateBuilder) Type(v string) *AppWidgetsUpdateBuilder {
	b.Params["type"] = v
	return b
}
