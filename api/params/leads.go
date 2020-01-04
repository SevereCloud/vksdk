package params // import "github.com/SevereCloud/vksdk/api/params"

import (
	"github.com/SevereCloud/vksdk/api"
)

// LeadsCheckUserBuilder builder
//
// Checks if the user can start the lead.
//
// https://vk.com/dev/leads.checkUser
type LeadsCheckUserBuilder struct {
	api.Params
}

// NewLeadsCheckUserBuilder func
func NewLeadsCheckUserBuilder() *LeadsCheckUserBuilder {
	return &LeadsCheckUserBuilder{api.Params{}}
}

// LeadID Lead ID.
func (b *LeadsCheckUserBuilder) LeadID(v int) {
	b.Params["lead_id"] = v
}

// TestResult Value to be return in 'result' field when test mode is used.
func (b *LeadsCheckUserBuilder) TestResult(v int) {
	b.Params["test_result"] = v
}

// TestMode parameter
func (b *LeadsCheckUserBuilder) TestMode(v bool) {
	b.Params["test_mode"] = v
}

// AutoStart parameter
func (b *LeadsCheckUserBuilder) AutoStart(v bool) {
	b.Params["auto_start"] = v
}

// Age User age.
func (b *LeadsCheckUserBuilder) Age(v int) {
	b.Params["age"] = v
}

// Country User country code.
func (b *LeadsCheckUserBuilder) Country(v string) {
	b.Params["country"] = v
}

// LeadsCompleteBuilder builder
//
// Completes the lead started by user.
//
// https://vk.com/dev/leads.complete
type LeadsCompleteBuilder struct {
	api.Params
}

// NewLeadsCompleteBuilder func
func NewLeadsCompleteBuilder() *LeadsCompleteBuilder {
	return &LeadsCompleteBuilder{api.Params{}}
}

// VkSID Session obtained as GET parameter when session started.
func (b *LeadsCompleteBuilder) VkSID(v string) {
	b.Params["vk_sid"] = v
}

// Secret Secret key from the lead testing interface.
func (b *LeadsCompleteBuilder) Secret(v string) {
	b.Params["secret"] = v
}

// Comment Comment text.
func (b *LeadsCompleteBuilder) Comment(v string) {
	b.Params["comment"] = v
}

// LeadsGetStatsBuilder builder
//
// Returns lead stats data.
//
// https://vk.com/dev/leads.getStats
type LeadsGetStatsBuilder struct {
	api.Params
}

// NewLeadsGetStatsBuilder func
func NewLeadsGetStatsBuilder() *LeadsGetStatsBuilder {
	return &LeadsGetStatsBuilder{api.Params{}}
}

// LeadID Lead ID.
func (b *LeadsGetStatsBuilder) LeadID(v int) {
	b.Params["lead_id"] = v
}

// Secret Secret key obtained from the lead testing interface.
func (b *LeadsGetStatsBuilder) Secret(v string) {
	b.Params["secret"] = v
}

// DateStart Day to start stats from (YYYY_MM_DD, e.g.2011-09-17).
func (b *LeadsGetStatsBuilder) DateStart(v string) {
	b.Params["date_start"] = v
}

// DateEnd Day to finish stats (YYYY_MM_DD, e.g.2011-09-17).
func (b *LeadsGetStatsBuilder) DateEnd(v string) {
	b.Params["date_end"] = v
}

// LeadsGetUsersBuilder builder
//
// Returns a list of last user actions for the offer.
//
// https://vk.com/dev/leads.getUsers
type LeadsGetUsersBuilder struct {
	api.Params
}

// NewLeadsGetUsersBuilder func
func NewLeadsGetUsersBuilder() *LeadsGetUsersBuilder {
	return &LeadsGetUsersBuilder{api.Params{}}
}

// OfferID Offer ID.
func (b *LeadsGetUsersBuilder) OfferID(v int) {
	b.Params["offer_id"] = v
}

// Secret Secret key obtained in the lead testing interface.
func (b *LeadsGetUsersBuilder) Secret(v string) {
	b.Params["secret"] = v
}

// Offset Offset needed to return a specific subset of results.
func (b *LeadsGetUsersBuilder) Offset(v int) {
	b.Params["offset"] = v
}

// Count Number of results to return.
func (b *LeadsGetUsersBuilder) Count(v int) {
	b.Params["count"] = v
}

// Status Action type. Possible values:
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
func (b *LeadsGetUsersBuilder) Status(v int) {
	b.Params["status"] = v
}

// Reverse Sort order. Possible values:
//
// * 1 — chronological;
//
// * 0 — reverse chronological.
func (b *LeadsGetUsersBuilder) Reverse(v bool) {
	b.Params["reverse"] = v
}

// LeadsMetricHitBuilder builder
//
// Counts the metric event.
//
// https://vk.com/dev/leads.metricHit
type LeadsMetricHitBuilder struct {
	api.Params
}

// NewLeadsMetricHitBuilder func
func NewLeadsMetricHitBuilder() *LeadsMetricHitBuilder {
	return &LeadsMetricHitBuilder{api.Params{}}
}

// Data Metric data obtained in the lead interface.
func (b *LeadsMetricHitBuilder) Data(v string) {
	b.Params["data"] = v
}

// LeadsStartBuilder builder
//
// Creates new session for the user passing the offer.
//
// https://vk.com/dev/leads.start
type LeadsStartBuilder struct {
	api.Params
}

// NewLeadsStartBuilder func
func NewLeadsStartBuilder() *LeadsStartBuilder {
	return &LeadsStartBuilder{api.Params{}}
}

// LeadID Lead ID.
func (b *LeadsStartBuilder) LeadID(v int) {
	b.Params["lead_id"] = v
}

// Secret Secret key from the lead testing interface.
func (b *LeadsStartBuilder) Secret(v string) {
	b.Params["secret"] = v
}

// UID parameter
func (b *LeadsStartBuilder) UID(v int) {
	b.Params["uid"] = v
}

// AID parameter
func (b *LeadsStartBuilder) AID(v int) {
	b.Params["aid"] = v
}

// TestMode parameter
func (b *LeadsStartBuilder) TestMode(v bool) {
	b.Params["test_mode"] = v
}

// Force parameter
func (b *LeadsStartBuilder) Force(v bool) {
	b.Params["force"] = v
}
