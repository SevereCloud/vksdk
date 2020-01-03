package params // import "github.com/SevereCloud/vksdk/api/params"

import (
	"github.com/SevereCloud/vksdk/api"
)

// UtilsCheckLinkBuilder builder
//
// Checks whether a link is blocked in VK.
//
// https://vk.com/dev/utils.checkLink
type UtilsCheckLinkBuilder struct {
	api.Params
}

// NewUtilsCheckLinkBuilder func
func NewUtilsCheckLinkBuilder() *UtilsCheckLinkBuilder {
	return &UtilsCheckLinkBuilder{api.Params{}}
}

// URL Link to check (e.g., 'http://google.com').
func (b *UtilsCheckLinkBuilder) URL(v string) {
	b.Params["url"] = v
}

// UtilsDeleteFromLastShortenedBuilder builder
//
// Deletes shortened link from user's list.
//
// https://vk.com/dev/utils.deleteFromLastShortened
type UtilsDeleteFromLastShortenedBuilder struct {
	api.Params
}

// NewUtilsDeleteFromLastShortenedBuilder func
func NewUtilsDeleteFromLastShortenedBuilder() *UtilsDeleteFromLastShortenedBuilder {
	return &UtilsDeleteFromLastShortenedBuilder{api.Params{}}
}

// Key Link key (characters after vk.cc/).
func (b *UtilsDeleteFromLastShortenedBuilder) Key(v string) {
	b.Params["key"] = v
}

// UtilsGetLastShortenedLinksBuilder builder
//
// Returns a list of user's shortened links.
//
// https://vk.com/dev/utils.getLastShortenedLinks
type UtilsGetLastShortenedLinksBuilder struct {
	api.Params
}

// NewUtilsGetLastShortenedLinksBuilder func
func NewUtilsGetLastShortenedLinksBuilder() *UtilsGetLastShortenedLinksBuilder {
	return &UtilsGetLastShortenedLinksBuilder{api.Params{}}
}

// Count Number of links to return.
func (b *UtilsGetLastShortenedLinksBuilder) Count(v int) {
	b.Params["count"] = v
}

// Offset Offset needed to return a specific subset of links.
func (b *UtilsGetLastShortenedLinksBuilder) Offset(v int) {
	b.Params["offset"] = v
}

// UtilsGetLinkStatsBuilder builder
//
// Returns stats data for shortened link.
//
// https://vk.com/dev/utils.getLinkStats
type UtilsGetLinkStatsBuilder struct {
	api.Params
}

// NewUtilsGetLinkStatsBuilder func
func NewUtilsGetLinkStatsBuilder() *UtilsGetLinkStatsBuilder {
	return &UtilsGetLinkStatsBuilder{api.Params{}}
}

// Key Link key (characters after vk.cc/).
func (b *UtilsGetLinkStatsBuilder) Key(v string) {
	b.Params["key"] = v
}

// Source Source of scope
func (b *UtilsGetLinkStatsBuilder) Source(v string) {
	b.Params["source"] = v
}

// AccessKey Access key for private link stats.
func (b *UtilsGetLinkStatsBuilder) AccessKey(v string) {
	b.Params["access_key"] = v
}

// Interval Interval.
func (b *UtilsGetLinkStatsBuilder) Interval(v string) {
	b.Params["interval"] = v
}

// IntervalsCount Number of intervals to return.
func (b *UtilsGetLinkStatsBuilder) IntervalsCount(v int) {
	b.Params["intervals_count"] = v
}

// Extended 1 — to return extended stats data (sex, age, geo). 0 — to return views number only.
func (b *UtilsGetLinkStatsBuilder) Extended(v bool) {
	b.Params["extended"] = v
}

// UtilsGetShortLinkBuilder builder
//
// Allows to receive a link shortened via vk.cc.
//
// https://vk.com/dev/utils.getShortLink
type UtilsGetShortLinkBuilder struct {
	api.Params
}

// NewUtilsGetShortLinkBuilder func
func NewUtilsGetShortLinkBuilder() *UtilsGetShortLinkBuilder {
	return &UtilsGetShortLinkBuilder{api.Params{}}
}

// URL URL to be shortened.
func (b *UtilsGetShortLinkBuilder) URL(v string) {
	b.Params["url"] = v
}

// Private 1 — private stats, 0 — public stats.
func (b *UtilsGetShortLinkBuilder) Private(v bool) {
	b.Params["private"] = v
}

// UtilsResolveScreenNameBuilder builder
//
// Detects a type of object (e.g., user, community, application) and its ID by screen name.
//
// https://vk.com/dev/utils.resolveScreenName
type UtilsResolveScreenNameBuilder struct {
	api.Params
}

// NewUtilsResolveScreenNameBuilder func
func NewUtilsResolveScreenNameBuilder() *UtilsResolveScreenNameBuilder {
	return &UtilsResolveScreenNameBuilder{api.Params{}}
}

// ScreenName Screen name of the user, community (e.g., 'apiclub,' 'andrew', or 'rules_of_war'), or application.
func (b *UtilsResolveScreenNameBuilder) ScreenName(v string) {
	b.Params["screen_name"] = v
}
