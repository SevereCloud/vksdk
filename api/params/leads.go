package params // import "github.com/SevereCloud/vksdk/api/params"

import (
	"github.com/SevereCloud/vksdk/api"
)

// LeadsCheckUserBulder builder
//
// Checks if the user can start the lead.
//
// https://vk.com/dev/leads.checkUser
type LeadsCheckUserBulder struct {
	api.Params
}

// NewLeadsCheckUserBulder func
func NewLeadsCheckUserBulder() *LeadsCheckUserBulder {
	return &LeadsCheckUserBulder{api.Params{}}
}

// LeadID Lead ID.
func (b *LeadsCheckUserBulder) LeadID(v int) {
	b.Params["lead_id"] = v
}

// TestResult Value to be return in 'result' field when test mode is used.
func (b *LeadsCheckUserBulder) TestResult(v int) {
	b.Params["test_result"] = v
}

// TestMode parameter
func (b *LeadsCheckUserBulder) TestMode(v bool) {
	b.Params["test_mode"] = v
}

// AutoStart parameter
func (b *LeadsCheckUserBulder) AutoStart(v bool) {
	b.Params["auto_start"] = v
}

// Age User age.
func (b *LeadsCheckUserBulder) Age(v int) {
	b.Params["age"] = v
}

// Country User country code.
func (b *LeadsCheckUserBulder) Country(v string) {
	b.Params["country"] = v
}

// LeadsCompleteBulder builder
//
// Completes the lead started by user.
//
// https://vk.com/dev/leads.complete
type LeadsCompleteBulder struct {
	api.Params
}

// NewLeadsCompleteBulder func
func NewLeadsCompleteBulder() *LeadsCompleteBulder {
	return &LeadsCompleteBulder{api.Params{}}
}

// VkSID Session obtained as GET parameter when session started.
func (b *LeadsCompleteBulder) VkSID(v string) {
	b.Params["vk_sid"] = v
}

// Secret Secret key from the lead testing interface.
func (b *LeadsCompleteBulder) Secret(v string) {
	b.Params["secret"] = v
}

// Comment Comment text.
func (b *LeadsCompleteBulder) Comment(v string) {
	b.Params["comment"] = v
}

// LeadsGetStatsBulder builder
//
// Returns lead stats data.
//
// https://vk.com/dev/leads.getStats
type LeadsGetStatsBulder struct {
	api.Params
}

// NewLeadsGetStatsBulder func
func NewLeadsGetStatsBulder() *LeadsGetStatsBulder {
	return &LeadsGetStatsBulder{api.Params{}}
}

// LeadID Lead ID.
func (b *LeadsGetStatsBulder) LeadID(v int) {
	b.Params["lead_id"] = v
}

// Secret Secret key obtained from the lead testing interface.
func (b *LeadsGetStatsBulder) Secret(v string) {
	b.Params["secret"] = v
}

// DateStart Day to start stats from (YYYY_MM_DD, e.g.2011-09-17).
func (b *LeadsGetStatsBulder) DateStart(v string) {
	b.Params["date_start"] = v
}

// DateEnd Day to finish stats (YYYY_MM_DD, e.g.2011-09-17).
func (b *LeadsGetStatsBulder) DateEnd(v string) {
	b.Params["date_end"] = v
}

// LeadsGetUsersBulder builder
//
// Returns a list of last user actions for the offer.
//
// https://vk.com/dev/leads.getUsers
type LeadsGetUsersBulder struct {
	api.Params
}

// NewLeadsGetUsersBulder func
func NewLeadsGetUsersBulder() *LeadsGetUsersBulder {
	return &LeadsGetUsersBulder{api.Params{}}
}

// OfferID Offer ID.
func (b *LeadsGetUsersBulder) OfferID(v int) {
	b.Params["offer_id"] = v
}

// Secret Secret key obtained in the lead testing interface.
func (b *LeadsGetUsersBulder) Secret(v string) {
	b.Params["secret"] = v
}

// Offset Offset needed to return a specific subset of results.
func (b *LeadsGetUsersBulder) Offset(v int) {
	b.Params["offset"] = v
}

// Count Number of results to return.
func (b *LeadsGetUsersBulder) Count(v int) {
	b.Params["count"] = v
}

// Status Action type. Possible values: *'0' — start,, *'1' — finish,, *'2' — blocking users,, *'3' — start in a test mode,, *'4' — finish in a test mode.
func (b *LeadsGetUsersBulder) Status(v int) {
	b.Params["status"] = v
}

// Reverse Sort order. Possible values: *'1' — chronological,, *'0' — reverse chronological.
func (b *LeadsGetUsersBulder) Reverse(v bool) {
	b.Params["reverse"] = v
}

// LeadsMetricHitBulder builder
//
// Counts the metric event.
//
// https://vk.com/dev/leads.metricHit
type LeadsMetricHitBulder struct {
	api.Params
}

// NewLeadsMetricHitBulder func
func NewLeadsMetricHitBulder() *LeadsMetricHitBulder {
	return &LeadsMetricHitBulder{api.Params{}}
}

// Data Metric data obtained in the lead interface.
func (b *LeadsMetricHitBulder) Data(v string) {
	b.Params["data"] = v
}

// LeadsStartBulder builder
//
// Creates new session for the user passing the offer.
//
// https://vk.com/dev/leads.start
type LeadsStartBulder struct {
	api.Params
}

// NewLeadsStartBulder func
func NewLeadsStartBulder() *LeadsStartBulder {
	return &LeadsStartBulder{api.Params{}}
}

// LeadID Lead ID.
func (b *LeadsStartBulder) LeadID(v int) {
	b.Params["lead_id"] = v
}

// Secret Secret key from the lead testing interface.
func (b *LeadsStartBulder) Secret(v string) {
	b.Params["secret"] = v
}

// UID parameter
func (b *LeadsStartBulder) UID(v int) {
	b.Params["uid"] = v
}

// AID parameter
func (b *LeadsStartBulder) AID(v int) {
	b.Params["aid"] = v
}

// TestMode parameter
func (b *LeadsStartBulder) TestMode(v bool) {
	b.Params["test_mode"] = v
}

// Force parameter
func (b *LeadsStartBulder) Force(v bool) {
	b.Params["force"] = v
}
