package params // import "github.com/SevereCloud/vksdk/v2/api/params"

import (
	"github.com/SevereCloud/vksdk/v2/api"
)

// LeadsCheckUserBuilder builder.
//
// Checks if the user can start the lead.
//
// https://vk.com/dev/leads.checkUser
type LeadsCheckUserBuilder struct {
	api.Params
}

// NewLeadsCheckUserBuilder func.
func NewLeadsCheckUserBuilder() *LeadsCheckUserBuilder {
	return &LeadsCheckUserBuilder{api.Params{}}
}

// LeadID parameter.
func (b *LeadsCheckUserBuilder) LeadID(v int) *LeadsCheckUserBuilder {
	b.Params["lead_id"] = v
	return b
}

// TestResult value to be return in 'result' field when test mode is used.
func (b *LeadsCheckUserBuilder) TestResult(v int) *LeadsCheckUserBuilder {
	b.Params["test_result"] = v
	return b
}

// TestMode parameter.
func (b *LeadsCheckUserBuilder) TestMode(v bool) *LeadsCheckUserBuilder {
	b.Params["test_mode"] = v
	return b
}

// AutoStart parameter.
func (b *LeadsCheckUserBuilder) AutoStart(v bool) *LeadsCheckUserBuilder {
	b.Params["auto_start"] = v
	return b
}

// Age user age.
func (b *LeadsCheckUserBuilder) Age(v int) *LeadsCheckUserBuilder {
	b.Params["age"] = v
	return b
}

// Country user country code.
func (b *LeadsCheckUserBuilder) Country(v string) *LeadsCheckUserBuilder {
	b.Params["country"] = v
	return b
}

// LeadsCompleteBuilder builder.
//
// Completes the lead started by user.
//
// https://vk.com/dev/leads.complete
type LeadsCompleteBuilder struct {
	api.Params
}

// NewLeadsCompleteBuilder func.
func NewLeadsCompleteBuilder() *LeadsCompleteBuilder {
	return &LeadsCompleteBuilder{api.Params{}}
}

// VkSID session obtained as GET parameter when session started.
func (b *LeadsCompleteBuilder) VkSID(v string) *LeadsCompleteBuilder {
	b.Params["vk_sid"] = v
	return b
}

// Secret secret key from the lead testing interface.
func (b *LeadsCompleteBuilder) Secret(v string) *LeadsCompleteBuilder {
	b.Params["secret"] = v
	return b
}

// Comment text.
func (b *LeadsCompleteBuilder) Comment(v string) *LeadsCompleteBuilder {
	b.Params["comment"] = v
	return b
}

// LeadsGetStatsBuilder builder.
//
// Returns lead stats data.
//
// https://vk.com/dev/leads.getStats
type LeadsGetStatsBuilder struct {
	api.Params
}

// NewLeadsGetStatsBuilder func.
func NewLeadsGetStatsBuilder() *LeadsGetStatsBuilder {
	return &LeadsGetStatsBuilder{api.Params{}}
}

// LeadID lead ID.
func (b *LeadsGetStatsBuilder) LeadID(v int) *LeadsGetStatsBuilder {
	b.Params["lead_id"] = v
	return b
}

// Secret key obtained from the lead testing interface.
func (b *LeadsGetStatsBuilder) Secret(v string) *LeadsGetStatsBuilder {
	b.Params["secret"] = v
	return b
}

// DateStart day to start stats from (YYYY_MM_DD, e.g.2011-09-17).
func (b *LeadsGetStatsBuilder) DateStart(v string) *LeadsGetStatsBuilder {
	b.Params["date_start"] = v
	return b
}

// DateEnd day to finish stats (YYYY_MM_DD, e.g.2011-09-17).
func (b *LeadsGetStatsBuilder) DateEnd(v string) *LeadsGetStatsBuilder {
	b.Params["date_end"] = v
	return b
}

// LeadsGetUsersBuilder builder.
//
// Returns a list of last user actions for the offer.
//
// https://vk.com/dev/leads.getUsers
type LeadsGetUsersBuilder struct {
	api.Params
}

// NewLeadsGetUsersBuilder func.
func NewLeadsGetUsersBuilder() *LeadsGetUsersBuilder {
	return &LeadsGetUsersBuilder{api.Params{}}
}

// OfferID parameter.
func (b *LeadsGetUsersBuilder) OfferID(v int) *LeadsGetUsersBuilder {
	b.Params["offer_id"] = v
	return b
}

// Secret key obtained in the lead testing interface.
func (b *LeadsGetUsersBuilder) Secret(v string) *LeadsGetUsersBuilder {
	b.Params["secret"] = v
	return b
}

// Offset needed to return a specific subset of results.
func (b *LeadsGetUsersBuilder) Offset(v int) *LeadsGetUsersBuilder {
	b.Params["offset"] = v
	return b
}

// Count number of results to return.
func (b *LeadsGetUsersBuilder) Count(v int) *LeadsGetUsersBuilder {
	b.Params["count"] = v
	return b
}

// Status action type. Possible values:
//
// * 0 — start;
//
// * 1 — finish;
//
// * 2 — blocking users;
//
// * 3 — start in a test mode;
//
// * 4 — finish in a test mode.
func (b *LeadsGetUsersBuilder) Status(v int) *LeadsGetUsersBuilder {
	b.Params["status"] = v
	return b
}

// Reverse sort order. Possible values:
//
// * 1 — chronological;
//
// * 0 — reverse chronological.
func (b *LeadsGetUsersBuilder) Reverse(v bool) *LeadsGetUsersBuilder {
	b.Params["reverse"] = v
	return b
}

// LeadsMetricHitBuilder builder.
//
// Counts the metric event.
//
// https://vk.com/dev/leads.metricHit
type LeadsMetricHitBuilder struct {
	api.Params
}

// NewLeadsMetricHitBuilder func.
func NewLeadsMetricHitBuilder() *LeadsMetricHitBuilder {
	return &LeadsMetricHitBuilder{api.Params{}}
}

// Data metric data obtained in the lead interface.
func (b *LeadsMetricHitBuilder) Data(v string) *LeadsMetricHitBuilder {
	b.Params["data"] = v
	return b
}

// LeadsStartBuilder builder.
//
// Creates new session for the user passing the offer.
//
// https://vk.com/dev/leads.start
type LeadsStartBuilder struct {
	api.Params
}

// NewLeadsStartBuilder func.
func NewLeadsStartBuilder() *LeadsStartBuilder {
	return &LeadsStartBuilder{api.Params{}}
}

// LeadID lead ID.
func (b *LeadsStartBuilder) LeadID(v int) *LeadsStartBuilder {
	b.Params["lead_id"] = v
	return b
}

// Secret key from the lead testing interface.
func (b *LeadsStartBuilder) Secret(v string) *LeadsStartBuilder {
	b.Params["secret"] = v
	return b
}

// UID parameter.
func (b *LeadsStartBuilder) UID(v int) *LeadsStartBuilder {
	b.Params["uid"] = v
	return b
}

// AID parameter.
func (b *LeadsStartBuilder) AID(v int) *LeadsStartBuilder {
	b.Params["aid"] = v
	return b
}

// TestMode parameter.
func (b *LeadsStartBuilder) TestMode(v bool) *LeadsStartBuilder {
	b.Params["test_mode"] = v
	return b
}

// Force parameter.
func (b *LeadsStartBuilder) Force(v bool) *LeadsStartBuilder {
	b.Params["force"] = v
	return b
}
