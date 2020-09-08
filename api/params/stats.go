package params // import "github.com/SevereCloud/vksdk/v2/api/params"

import (
	"github.com/SevereCloud/vksdk/v2/api"
)

// StatsGetBuilder builder.
//
// Returns statistics of a community or an application.
//
// https://vk.com/dev/stats.get
type StatsGetBuilder struct {
	api.Params
}

// NewStatsGetBuilder func.
func NewStatsGetBuilder() *StatsGetBuilder {
	return &StatsGetBuilder{api.Params{}}
}

// GroupID community ID.
func (b *StatsGetBuilder) GroupID(v int) *StatsGetBuilder {
	b.Params["group_id"] = v
	return b
}

// AppID application ID.
func (b *StatsGetBuilder) AppID(v int) *StatsGetBuilder {
	b.Params["app_id"] = v
	return b
}

// TimestampFrom parameter.
func (b *StatsGetBuilder) TimestampFrom(v int) *StatsGetBuilder {
	b.Params["timestamp_from"] = v
	return b
}

// TimestampTo parameter.
func (b *StatsGetBuilder) TimestampTo(v int) *StatsGetBuilder {
	b.Params["timestamp_to"] = v
	return b
}

// Interval parameter.
func (b *StatsGetBuilder) Interval(v string) *StatsGetBuilder {
	b.Params["interval"] = v
	return b
}

// IntervalsCount parameter.
func (b *StatsGetBuilder) IntervalsCount(v int) *StatsGetBuilder {
	b.Params["intervals_count"] = v
	return b
}

// Filters parameter.
func (b *StatsGetBuilder) Filters(v []string) *StatsGetBuilder {
	b.Params["filters"] = v
	return b
}

// StatsGroups parameter.
func (b *StatsGetBuilder) StatsGroups(v []string) *StatsGetBuilder {
	b.Params["stats_groups"] = v
	return b
}

// Extended parameter.
func (b *StatsGetBuilder) Extended(v bool) *StatsGetBuilder {
	b.Params["extended"] = v
	return b
}

// StatsGetPostReachBuilder builder.
//
// Returns stats for a wall post.
//
// https://vk.com/dev/stats.getPostReach
type StatsGetPostReachBuilder struct {
	api.Params
}

// NewStatsGetPostReachBuilder func.
func NewStatsGetPostReachBuilder() *StatsGetPostReachBuilder {
	return &StatsGetPostReachBuilder{api.Params{}}
}

// OwnerID post owner community id. Specify with "-" sign.
func (b *StatsGetPostReachBuilder) OwnerID(v string) *StatsGetPostReachBuilder {
	b.Params["owner_id"] = v
	return b
}

// PostID wall post id. Note that stats are available only for '300' last (newest) posts on a community wall.
func (b *StatsGetPostReachBuilder) PostID(v int) *StatsGetPostReachBuilder {
	b.Params["post_id"] = v
	return b
}

// StatsTrackVisitorBuilder builder.
//
// https://vk.com/dev/stats.trackVisitor
type StatsTrackVisitorBuilder struct {
	api.Params
}

// NewStatsTrackVisitorBuilder func.
func NewStatsTrackVisitorBuilder() *StatsTrackVisitorBuilder {
	return &StatsTrackVisitorBuilder{api.Params{}}
}

// ID parameter.
func (b *StatsTrackVisitorBuilder) ID(v string) *StatsTrackVisitorBuilder {
	b.Params["id"] = v
	return b
}
