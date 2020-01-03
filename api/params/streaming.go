package params // import "github.com/SevereCloud/vksdk/api/params"

import (
	"github.com/SevereCloud/vksdk/api"
)

// StreamingSetSettingsBulder builder
//
// https://vk.com/dev/streaming.setSettings
type StreamingSetSettingsBulder struct {
	api.Params
}

// NewStreamingSetSettingsBulder func
func NewStreamingSetSettingsBulder() *StreamingSetSettingsBulder {
	return &StreamingSetSettingsBulder{api.Params{}}
}

// MonthlyTier parameter
func (b *StreamingSetSettingsBulder) MonthlyTier(v string) {
	b.Params["monthly_tier"] = v
}
