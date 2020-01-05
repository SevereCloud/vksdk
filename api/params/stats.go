package params // import "github.com/SevereCloud/vksdk/api/params"

import (
	"github.com/SevereCloud/vksdk/api"
)

// StatsGetBuilder builder
//
// Returns statistics of a community or an application.
//
// https://vk.com/dev/stats.get
type StatsGetBuilder struct {
	api.Params
}

// NewStatsGetBuilder func
func NewStatsGetBuilder() *StatsGetBuilder {
	return &StatsGetBuilder{api.Params{}}
}

// GroupID Community ID.
func (b *StatsGetBuilder) GroupID(v int) {
	b.Params["group_id"] = v
}

// AppID Application ID.
func (b *StatsGetBuilder) AppID(v int) {
	b.Params["app_id"] = v
}

// TimestampFrom parameter
func (b *StatsGetBuilder) TimestampFrom(v int) {
	b.Params["timestamp_from"] = v
}

// TimestampTo parameter
func (b *StatsGetBuilder) TimestampTo(v int) {
	b.Params["timestamp_to"] = v
}

// Interval parameter
func (b *StatsGetBuilder) Interval(v string) {
	b.Params["interval"] = v
}

// IntervalsCount parameter
func (b *StatsGetBuilder) IntervalsCount(v int) {
	b.Params["intervals_count"] = v
}

// Filters parameter
func (b *StatsGetBuilder) Filters(v []string) {
	b.Params["filters"] = v
}

// StatsGroups parameter
func (b *StatsGetBuilder) StatsGroups(v []string) {
	b.Params["stats_groups"] = v
}

// Extended parameter
func (b *StatsGetBuilder) Extended(v bool) {
	b.Params["extended"] = v
}

// StatsGetPostReachBuilder builder
//
// Returns stats for a wall post.
//
// https://vk.com/dev/stats.getPostReach
type StatsGetPostReachBuilder struct {
	api.Params
}

// NewStatsGetPostReachBuilder func
func NewStatsGetPostReachBuilder() *StatsGetPostReachBuilder {
	return &StatsGetPostReachBuilder{api.Params{}}
}

// OwnerID post owner community id. Specify with "-" sign.
func (b *StatsGetPostReachBuilder) OwnerID(v string) {
	b.Params["owner_id"] = v
}

// PostID wall post id. Note that stats are available only for '300' last (newest) posts on a community wall.
func (b *StatsGetPostReachBuilder) PostID(v int) {
	b.Params["post_id"] = v
}

// StatsTrackVisitorBuilder builder
//
// https://vk.com/dev/stats.trackVisitor
type StatsTrackVisitorBuilder struct {
	api.Params
}

// NewStatsTrackVisitorBuilder func
func NewStatsTrackVisitorBuilder() *StatsTrackVisitorBuilder {
	return &StatsTrackVisitorBuilder{api.Params{}}
}

// ID parameter
func (b *StatsTrackVisitorBuilder) ID(v string) {
	b.Params["id"] = v
}
