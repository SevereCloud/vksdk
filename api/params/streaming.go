package params // import "github.com/SevereCloud/vksdk/v3/api/params"

import (
	"github.com/SevereCloud/vksdk/v3/api"
)

// StreamingSetSettingsBuilder builder.
//
// https://dev.vk.ru/method/streaming.setSettings
type StreamingSetSettingsBuilder struct {
	api.Params
}

// NewStreamingSetSettingsBuilder func.
func NewStreamingSetSettingsBuilder() *StreamingSetSettingsBuilder {
	return &StreamingSetSettingsBuilder{api.Params{}}
}

// MonthlyTier parameter.
func (b *StreamingSetSettingsBuilder) MonthlyTier(v string) *StreamingSetSettingsBuilder {
	b.Params["monthly_tier"] = v
	return b
}
