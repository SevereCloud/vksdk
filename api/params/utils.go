package params // import "github.com/SevereCloud/vksdk/api/params"

import (
	"github.com/SevereCloud/vksdk/api"
)

// UtilsCheckLinkBulder builder
//
// Checks whether a link is blocked in VK.
//
// https://vk.com/dev/utils.checkLink
type UtilsCheckLinkBulder struct {
	api.Params
}

// NewUtilsCheckLinkBulder func
func NewUtilsCheckLinkBulder() *UtilsCheckLinkBulder {
	return &UtilsCheckLinkBulder{api.Params{}}
}

// URL Link to check (e.g., 'http://google.com').
func (b *UtilsCheckLinkBulder) URL(v string) {
	b.Params["url"] = v
}

// UtilsDeleteFromLastShortenedBulder builder
//
// Deletes shortened link from user's list.
//
// https://vk.com/dev/utils.deleteFromLastShortened
type UtilsDeleteFromLastShortenedBulder struct {
	api.Params
}

// NewUtilsDeleteFromLastShortenedBulder func
func NewUtilsDeleteFromLastShortenedBulder() *UtilsDeleteFromLastShortenedBulder {
	return &UtilsDeleteFromLastShortenedBulder{api.Params{}}
}

// Key Link key (characters after vk.cc/).
func (b *UtilsDeleteFromLastShortenedBulder) Key(v string) {
	b.Params["key"] = v
}

// UtilsGetLastShortenedLinksBulder builder
//
// Returns a list of user's shortened links.
//
// https://vk.com/dev/utils.getLastShortenedLinks
type UtilsGetLastShortenedLinksBulder struct {
	api.Params
}

// NewUtilsGetLastShortenedLinksBulder func
func NewUtilsGetLastShortenedLinksBulder() *UtilsGetLastShortenedLinksBulder {
	return &UtilsGetLastShortenedLinksBulder{api.Params{}}
}

// Count Number of links to return.
func (b *UtilsGetLastShortenedLinksBulder) Count(v int) {
	b.Params["count"] = v
}

// Offset Offset needed to return a specific subset of links.
func (b *UtilsGetLastShortenedLinksBulder) Offset(v int) {
	b.Params["offset"] = v
}

// UtilsGetLinkStatsBulder builder
//
// Returns stats data for shortened link.
//
// https://vk.com/dev/utils.getLinkStats
type UtilsGetLinkStatsBulder struct {
	api.Params
}

// NewUtilsGetLinkStatsBulder func
func NewUtilsGetLinkStatsBulder() *UtilsGetLinkStatsBulder {
	return &UtilsGetLinkStatsBulder{api.Params{}}
}

// Key Link key (characters after vk.cc/).
func (b *UtilsGetLinkStatsBulder) Key(v string) {
	b.Params["key"] = v
}

// Source Source of scope
func (b *UtilsGetLinkStatsBulder) Source(v string) {
	b.Params["source"] = v
}

// AccessKey Access key for private link stats.
func (b *UtilsGetLinkStatsBulder) AccessKey(v string) {
	b.Params["access_key"] = v
}

// Interval Interval.
func (b *UtilsGetLinkStatsBulder) Interval(v string) {
	b.Params["interval"] = v
}

// IntervalsCount Number of intervals to return.
func (b *UtilsGetLinkStatsBulder) IntervalsCount(v int) {
	b.Params["intervals_count"] = v
}

// Extended 1 — to return extended stats data (sex, age, geo). 0 — to return views number only.
func (b *UtilsGetLinkStatsBulder) Extended(v bool) {
	b.Params["extended"] = v
}

// UtilsGetShortLinkBulder builder
//
// Allows to receive a link shortened via vk.cc.
//
// https://vk.com/dev/utils.getShortLink
type UtilsGetShortLinkBulder struct {
	api.Params
}

// NewUtilsGetShortLinkBulder func
func NewUtilsGetShortLinkBulder() *UtilsGetShortLinkBulder {
	return &UtilsGetShortLinkBulder{api.Params{}}
}

// URL URL to be shortened.
func (b *UtilsGetShortLinkBulder) URL(v string) {
	b.Params["url"] = v
}

// Private 1 — private stats, 0 — public stats.
func (b *UtilsGetShortLinkBulder) Private(v bool) {
	b.Params["private"] = v
}

// UtilsResolveScreenNameBulder builder
//
// Detects a type of object (e.g., user, community, application) and its ID by screen name.
//
// https://vk.com/dev/utils.resolveScreenName
type UtilsResolveScreenNameBulder struct {
	api.Params
}

// NewUtilsResolveScreenNameBulder func
func NewUtilsResolveScreenNameBulder() *UtilsResolveScreenNameBulder {
	return &UtilsResolveScreenNameBulder{api.Params{}}
}

// ScreenName Screen name of the user, community (e.g., 'apiclub,' 'andrew', or 'rules_of_war'), or application.
func (b *UtilsResolveScreenNameBulder) ScreenName(v string) {
	b.Params["screen_name"] = v
}
