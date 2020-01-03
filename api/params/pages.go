package params // import "github.com/SevereCloud/vksdk/api/params"

import (
	"github.com/SevereCloud/vksdk/api"
)

// PagesClearCacheBulder builder
//
// Allows to clear the cache of particular 'external' pages which may be attached to VK posts.
//
// https://vk.com/dev/pages.clearCache
type PagesClearCacheBulder struct {
	api.Params
}

// NewPagesClearCacheBulder func
func NewPagesClearCacheBulder() *PagesClearCacheBulder {
	return &PagesClearCacheBulder{api.Params{}}
}

// URL Address of the page where you need to refesh the cached version
func (b *PagesClearCacheBulder) URL(v string) {
	b.Params["url"] = v
}

// PagesGetBulder builder
//
// Returns information about a wiki page.
//
// https://vk.com/dev/pages.get
type PagesGetBulder struct {
	api.Params
}

// NewPagesGetBulder func
func NewPagesGetBulder() *PagesGetBulder {
	return &PagesGetBulder{api.Params{}}
}

// OwnerID Page owner ID.
func (b *PagesGetBulder) OwnerID(v int) {
	b.Params["owner_id"] = v
}

// PageID Wiki page ID.
func (b *PagesGetBulder) PageID(v int) {
	b.Params["page_id"] = v
}

// Global '1' — to return information about a global wiki page
func (b *PagesGetBulder) Global(v bool) {
	b.Params["global"] = v
}

// SitePreview '1' — resulting wiki page is a preview for the attached link
func (b *PagesGetBulder) SitePreview(v bool) {
	b.Params["site_preview"] = v
}

// Title Wiki page title.
func (b *PagesGetBulder) Title(v string) {
	b.Params["title"] = v
}

// NeedSource parameter
func (b *PagesGetBulder) NeedSource(v bool) {
	b.Params["need_source"] = v
}

// NeedHTML '1' — to return the page as HTML,
func (b *PagesGetBulder) NeedHTML(v bool) {
	b.Params["need_html"] = v
}

// PagesGetHistoryBulder builder
//
// Returns a list of all previous versions of a wiki page.
//
// https://vk.com/dev/pages.getHistory
type PagesGetHistoryBulder struct {
	api.Params
}

// NewPagesGetHistoryBulder func
func NewPagesGetHistoryBulder() *PagesGetHistoryBulder {
	return &PagesGetHistoryBulder{api.Params{}}
}

// PageID Wiki page ID.
func (b *PagesGetHistoryBulder) PageID(v int) {
	b.Params["page_id"] = v
}

// GroupID ID of the community that owns the wiki page.
func (b *PagesGetHistoryBulder) GroupID(v int) {
	b.Params["group_id"] = v
}

// UserID parameter
func (b *PagesGetHistoryBulder) UserID(v int) {
	b.Params["user_id"] = v
}

// PagesGetTitlesBulder builder
//
// Returns a list of wiki pages in a group.
//
// https://vk.com/dev/pages.getTitles
type PagesGetTitlesBulder struct {
	api.Params
}

// NewPagesGetTitlesBulder func
func NewPagesGetTitlesBulder() *PagesGetTitlesBulder {
	return &PagesGetTitlesBulder{api.Params{}}
}

// GroupID ID of the community that owns the wiki page.
func (b *PagesGetTitlesBulder) GroupID(v int) {
	b.Params["group_id"] = v
}

// PagesGetVersionBulder builder
//
// Returns the text of one of the previous versions of a wiki page.
//
// https://vk.com/dev/pages.getVersion
type PagesGetVersionBulder struct {
	api.Params
}

// NewPagesGetVersionBulder func
func NewPagesGetVersionBulder() *PagesGetVersionBulder {
	return &PagesGetVersionBulder{api.Params{}}
}

// VersionID parameter
func (b *PagesGetVersionBulder) VersionID(v int) {
	b.Params["version_id"] = v
}

// GroupID ID of the community that owns the wiki page.
func (b *PagesGetVersionBulder) GroupID(v int) {
	b.Params["group_id"] = v
}

// UserID parameter
func (b *PagesGetVersionBulder) UserID(v int) {
	b.Params["user_id"] = v
}

// NeedHTML '1' — to return the page as HTML
func (b *PagesGetVersionBulder) NeedHTML(v bool) {
	b.Params["need_html"] = v
}

// PagesParseWikiBulder builder
//
// Returns HTML representation of the wiki markup.
//
// https://vk.com/dev/pages.parseWiki
type PagesParseWikiBulder struct {
	api.Params
}

// NewPagesParseWikiBulder func
func NewPagesParseWikiBulder() *PagesParseWikiBulder {
	return &PagesParseWikiBulder{api.Params{}}
}

// Text Text of the wiki page.
func (b *PagesParseWikiBulder) Text(v string) {
	b.Params["text"] = v
}

// GroupID ID of the group in the context of which this markup is interpreted.
func (b *PagesParseWikiBulder) GroupID(v int) {
	b.Params["group_id"] = v
}

// PagesSaveBulder builder
//
// Saves the text of a wiki page.
//
// https://vk.com/dev/pages.save
type PagesSaveBulder struct {
	api.Params
}

// NewPagesSaveBulder func
func NewPagesSaveBulder() *PagesSaveBulder {
	return &PagesSaveBulder{api.Params{}}
}

// Text Text of the wiki page in wiki-format.
func (b *PagesSaveBulder) Text(v string) {
	b.Params["text"] = v
}

// PageID Wiki page ID. The 'title' parameter can be passed instead of 'pid'.
func (b *PagesSaveBulder) PageID(v int) {
	b.Params["page_id"] = v
}

// GroupID ID of the community that owns the wiki page.
func (b *PagesSaveBulder) GroupID(v int) {
	b.Params["group_id"] = v
}

// UserID User ID
func (b *PagesSaveBulder) UserID(v int) {
	b.Params["user_id"] = v
}

// Title Wiki page title.
func (b *PagesSaveBulder) Title(v string) {
	b.Params["title"] = v
}

// PagesSaveAccessBulder builder
//
// Saves modified read and edit access settings for a wiki page.
//
// https://vk.com/dev/pages.saveAccess
type PagesSaveAccessBulder struct {
	api.Params
}

// NewPagesSaveAccessBulder func
func NewPagesSaveAccessBulder() *PagesSaveAccessBulder {
	return &PagesSaveAccessBulder{api.Params{}}
}

// PageID Wiki page ID.
func (b *PagesSaveAccessBulder) PageID(v int) {
	b.Params["page_id"] = v
}

// GroupID ID of the community that owns the wiki page.
func (b *PagesSaveAccessBulder) GroupID(v int) {
	b.Params["group_id"] = v
}

// UserID parameter
func (b *PagesSaveAccessBulder) UserID(v int) {
	b.Params["user_id"] = v
}

// View Who can view the wiki page: '1' — only community members, '2' — all users can view the page, '0' — only community managers
func (b *PagesSaveAccessBulder) View(v int) {
	b.Params["view"] = v
}

// Edit Who can edit the wiki page: '1' — only community members, '2' — all users can edit the page, '0' — only community managers
func (b *PagesSaveAccessBulder) Edit(v int) {
	b.Params["edit"] = v
}
