package params // import "github.com/SevereCloud/vksdk/api/params"

import (
	"github.com/SevereCloud/vksdk/api"
)

// AppWidgetsUpdateBulder builder
//
// Allows to update community app widget
//
// https://vk.com/dev/appWidgets.update
type AppWidgetsUpdateBulder struct {
	api.Params
}

// NewAppWidgetsUpdateBulder func
func NewAppWidgetsUpdateBulder() *AppWidgetsUpdateBulder {
	return &AppWidgetsUpdateBulder{api.Params{}}
}

// Code parameter
func (b *AppWidgetsUpdateBulder) Code(v string) {
	b.Params["code"] = v
}

// Type parameter
func (b *AppWidgetsUpdateBulder) Type(v string) {
	b.Params["type"] = v
}
