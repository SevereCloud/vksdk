package params // import "github.com/SevereCloud/vksdk/api/params"

import (
	"github.com/SevereCloud/vksdk/api"
)

// StatsGetBulder builder
//
// Returns statistics of a community or an application.
//
// https://vk.com/dev/stats.get
type StatsGetBulder struct {
	api.Params
}

// NewStatsGetBulder func
func NewStatsGetBulder() *StatsGetBulder {
	return &StatsGetBulder{api.Params{}}
}

// GroupID Community ID.
func (b *StatsGetBulder) GroupID(v int) {
	b.Params["group_id"] = v
}

// AppID Application ID.
func (b *StatsGetBulder) AppID(v int) {
	b.Params["app_id"] = v
}

// TimestampFrom parameter
func (b *StatsGetBulder) TimestampFrom(v int) {
	b.Params["timestamp_from"] = v
}

// TimestampTo parameter
func (b *StatsGetBulder) TimestampTo(v int) {
	b.Params["timestamp_to"] = v
}

// Interval parameter
func (b *StatsGetBulder) Interval(v string) {
	b.Params["interval"] = v
}

// IntervalsCount parameter
func (b *StatsGetBulder) IntervalsCount(v int) {
	b.Params["intervals_count"] = v
}

// Filters parameter
func (b *StatsGetBulder) Filters(v []string) {
	b.Params["filters"] = v
}

// StatsGroups parameter
func (b *StatsGetBulder) StatsGroups(v []string) {
	b.Params["stats_groups"] = v
}

// Extended parameter
func (b *StatsGetBulder) Extended(v bool) {
	b.Params["extended"] = v
}

// StatsGetPostReachBulder builder
//
// Returns stats for a wall post.
//
// https://vk.com/dev/stats.getPostReach
type StatsGetPostReachBulder struct {
	api.Params
}

// NewStatsGetPostReachBulder func
func NewStatsGetPostReachBulder() *StatsGetPostReachBulder {
	return &StatsGetPostReachBulder{api.Params{}}
}

// OwnerID post owner community id. Specify with "-" sign.
func (b *StatsGetPostReachBulder) OwnerID(v string) {
	b.Params["owner_id"] = v
}

// PostID wall post id. Note that stats are available only for '300' last (newest) posts on a community wall.
func (b *StatsGetPostReachBulder) PostID(v int) {
	b.Params["post_id"] = v
}

// StatsTrackVisitorBulder builder
//
// https://vk.com/dev/stats.trackVisitor
type StatsTrackVisitorBulder struct {
	api.Params
}

// NewStatsTrackVisitorBulder func
func NewStatsTrackVisitorBulder() *StatsTrackVisitorBulder {
	return &StatsTrackVisitorBulder{api.Params{}}
}

// ID parameter
func (b *StatsTrackVisitorBulder) ID(v string) {
	b.Params["id"] = v
}
