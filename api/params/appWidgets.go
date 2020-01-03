package params // import "github.com/SevereCloud/vksdk/api/params"

import (
	"github.com/SevereCloud/vksdk/api"
)

// AppWidgetsUpdateBuilder builder
//
// Allows to update community app widget
//
// https://vk.com/dev/appWidgets.update
type AppWidgetsUpdateBuilder struct {
	api.Params
}

// NewAppWidgetsUpdateBuilder func
func NewAppWidgetsUpdateBuilder() *AppWidgetsUpdateBuilder {
	return &AppWidgetsUpdateBuilder{api.Params{}}
}

// Code parameter
func (b *AppWidgetsUpdateBuilder) Code(v string) {
	b.Params["code"] = v
}

// Type parameter
func (b *AppWidgetsUpdateBuilder) Type(v string) {
	b.Params["type"] = v
}
