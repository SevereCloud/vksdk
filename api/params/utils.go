package params // import "github.com/SevereCloud/vksdk/v3/api/params"

import (
	"github.com/SevereCloud/vksdk/v3/api"
)

// UtilsCheckLinkBuilder builder.
//
// Checks whether a link is blocked in VK.
//
// https://dev.vk.ru/method/utils.checkLink
type UtilsCheckLinkBuilder struct {
	api.Params
}

// NewUtilsCheckLinkBuilder func.
func NewUtilsCheckLinkBuilder() *UtilsCheckLinkBuilder {
	return &UtilsCheckLinkBuilder{api.Params{}}
}

// URL link to check (e.g., 'http://google.com').
func (b *UtilsCheckLinkBuilder) URL(v string) *UtilsCheckLinkBuilder {
	b.Params["url"] = v
	return b
}

// UtilsDeleteFromLastShortenedBuilder builder.
//
// Deletes shortened link from user's list.
//
// https://dev.vk.ru/method/utils.deleteFromLastShortened
type UtilsDeleteFromLastShortenedBuilder struct {
	api.Params
}

// NewUtilsDeleteFromLastShortenedBuilder func.
func NewUtilsDeleteFromLastShortenedBuilder() *UtilsDeleteFromLastShortenedBuilder {
	return &UtilsDeleteFromLastShortenedBuilder{api.Params{}}
}

// Key link (characters after vk.cc/).
func (b *UtilsDeleteFromLastShortenedBuilder) Key(v string) *UtilsDeleteFromLastShortenedBuilder {
	b.Params["key"] = v
	return b
}

// UtilsGetLastShortenedLinksBuilder builder.
//
// Returns a list of user's shortened links.
//
// https://dev.vk.ru/method/utils.getLastShortenedLinks
type UtilsGetLastShortenedLinksBuilder struct {
	api.Params
}

// NewUtilsGetLastShortenedLinksBuilder func.
func NewUtilsGetLastShortenedLinksBuilder() *UtilsGetLastShortenedLinksBuilder {
	return &UtilsGetLastShortenedLinksBuilder{api.Params{}}
}

// Count number of links to return.
func (b *UtilsGetLastShortenedLinksBuilder) Count(v int) *UtilsGetLastShortenedLinksBuilder {
	b.Params["count"] = v
	return b
}

// Offset needed to return a specific subset of links.
func (b *UtilsGetLastShortenedLinksBuilder) Offset(v int) *UtilsGetLastShortenedLinksBuilder {
	b.Params["offset"] = v
	return b
}

// UtilsGetLinkStatsBuilder builder.
//
// Returns stats data for shortened link.
//
// https://dev.vk.ru/method/utils.getLinkStats
type UtilsGetLinkStatsBuilder struct {
	api.Params
}

// NewUtilsGetLinkStatsBuilder func.
func NewUtilsGetLinkStatsBuilder() *UtilsGetLinkStatsBuilder {
	return &UtilsGetLinkStatsBuilder{api.Params{}}
}

// Key link (characters after vk.cc/).
func (b *UtilsGetLinkStatsBuilder) Key(v string) *UtilsGetLinkStatsBuilder {
	b.Params["key"] = v
	return b
}

// Source of scope.
func (b *UtilsGetLinkStatsBuilder) Source(v string) *UtilsGetLinkStatsBuilder {
	b.Params["source"] = v
	return b
}

// AccessKey access key for private link stats.
func (b *UtilsGetLinkStatsBuilder) AccessKey(v string) *UtilsGetLinkStatsBuilder {
	b.Params["access_key"] = v
	return b
}

// Interval parameter.
func (b *UtilsGetLinkStatsBuilder) Interval(v string) *UtilsGetLinkStatsBuilder {
	b.Params["interval"] = v
	return b
}

// IntervalsCount number of intervals to return.
func (b *UtilsGetLinkStatsBuilder) IntervalsCount(v int) *UtilsGetLinkStatsBuilder {
	b.Params["intervals_count"] = v
	return b
}

// Extended 1 — to return extended stats data (sex, age, geo). 0 — to return views number only.
func (b *UtilsGetLinkStatsBuilder) Extended(v bool) *UtilsGetLinkStatsBuilder {
	b.Params["extended"] = v
	return b
}

// UtilsGetShortLinkBuilder builder.
//
// Allows to receive a link shortened via vk.cc.
//
// https://dev.vk.ru/method/utils.getShortLink
type UtilsGetShortLinkBuilder struct {
	api.Params
}

// NewUtilsGetShortLinkBuilder func.
func NewUtilsGetShortLinkBuilder() *UtilsGetShortLinkBuilder {
	return &UtilsGetShortLinkBuilder{api.Params{}}
}

// URL URL to be shortened.
func (b *UtilsGetShortLinkBuilder) URL(v string) *UtilsGetShortLinkBuilder {
	b.Params["url"] = v
	return b
}

// Private 1 — private stats, 0 — public stats.
func (b *UtilsGetShortLinkBuilder) Private(v bool) *UtilsGetShortLinkBuilder {
	b.Params["private"] = v
	return b
}

// UtilsResolveScreenNameBuilder builder.
//
// Detects a type of object (e.g., user, community, application) and its ID by screen name.
//
// https://dev.vk.ru/method/utils.resolveScreenName
type UtilsResolveScreenNameBuilder struct {
	api.Params
}

// NewUtilsResolveScreenNameBuilder func.
func NewUtilsResolveScreenNameBuilder() *UtilsResolveScreenNameBuilder {
	return &UtilsResolveScreenNameBuilder{api.Params{}}
}

// ScreenName of the user, community (e.g., 'apiclub,' 'andrew', or 'rules_of_war'), or application.
func (b *UtilsResolveScreenNameBuilder) ScreenName(v string) *UtilsResolveScreenNameBuilder {
	b.Params["screen_name"] = v
	return b
}
