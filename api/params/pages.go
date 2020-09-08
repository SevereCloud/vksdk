package params // import "github.com/SevereCloud/vksdk/v2/api/params"

import (
	"github.com/SevereCloud/vksdk/v2/api"
)

// PagesClearCacheBuilder builder.
//
// Allows to clear the cache of particular 'external' pages which may be attached to VK posts.
//
// https://vk.com/dev/pages.clearCache
type PagesClearCacheBuilder struct {
	api.Params
}

// NewPagesClearCacheBuilder func.
func NewPagesClearCacheBuilder() *PagesClearCacheBuilder {
	return &PagesClearCacheBuilder{api.Params{}}
}

// URL address of the page where you need to refresh the cached version.
func (b *PagesClearCacheBuilder) URL(v string) *PagesClearCacheBuilder {
	b.Params["url"] = v
	return b
}

// PagesGetBuilder builder.
//
// Returns information about a wiki page.
//
// https://vk.com/dev/pages.get
type PagesGetBuilder struct {
	api.Params
}

// NewPagesGetBuilder func.
func NewPagesGetBuilder() *PagesGetBuilder {
	return &PagesGetBuilder{api.Params{}}
}

// OwnerID page owner ID.
func (b *PagesGetBuilder) OwnerID(v int) *PagesGetBuilder {
	b.Params["owner_id"] = v
	return b
}

// PageID wiki page ID.
func (b *PagesGetBuilder) PageID(v int) *PagesGetBuilder {
	b.Params["page_id"] = v
	return b
}

// Global '1' — to return information about a global wiki page.
func (b *PagesGetBuilder) Global(v bool) *PagesGetBuilder {
	b.Params["global"] = v
	return b
}

// SitePreview '1' — resulting wiki page is a preview for the attached link.
func (b *PagesGetBuilder) SitePreview(v bool) *PagesGetBuilder {
	b.Params["site_preview"] = v
	return b
}

// Title wiki page title.
func (b *PagesGetBuilder) Title(v string) *PagesGetBuilder {
	b.Params["title"] = v
	return b
}

// NeedSource parameter.
func (b *PagesGetBuilder) NeedSource(v bool) *PagesGetBuilder {
	b.Params["need_source"] = v
	return b
}

// NeedHTML '1' — to return the page as HTML.
func (b *PagesGetBuilder) NeedHTML(v bool) *PagesGetBuilder {
	b.Params["need_html"] = v
	return b
}

// PagesGetHistoryBuilder builder.
//
// Returns a list of all previous versions of a wiki page.
//
// https://vk.com/dev/pages.getHistory
type PagesGetHistoryBuilder struct {
	api.Params
}

// NewPagesGetHistoryBuilder func.
func NewPagesGetHistoryBuilder() *PagesGetHistoryBuilder {
	return &PagesGetHistoryBuilder{api.Params{}}
}

// PageID wiki page ID.
func (b *PagesGetHistoryBuilder) PageID(v int) *PagesGetHistoryBuilder {
	b.Params["page_id"] = v
	return b
}

// GroupID ID of the community that owns the wiki page.
func (b *PagesGetHistoryBuilder) GroupID(v int) *PagesGetHistoryBuilder {
	b.Params["group_id"] = v
	return b
}

// UserID parameter.
func (b *PagesGetHistoryBuilder) UserID(v int) *PagesGetHistoryBuilder {
	b.Params["user_id"] = v
	return b
}

// PagesGetTitlesBuilder builder.
//
// Returns a list of wiki pages in a group.
//
// https://vk.com/dev/pages.getTitles
type PagesGetTitlesBuilder struct {
	api.Params
}

// NewPagesGetTitlesBuilder func.
func NewPagesGetTitlesBuilder() *PagesGetTitlesBuilder {
	return &PagesGetTitlesBuilder{api.Params{}}
}

// GroupID ID of the community that owns the wiki page.
func (b *PagesGetTitlesBuilder) GroupID(v int) *PagesGetTitlesBuilder {
	b.Params["group_id"] = v
	return b
}

// PagesGetVersionBuilder builder.
//
// Returns the text of one of the previous versions of a wiki page.
//
// https://vk.com/dev/pages.getVersion
type PagesGetVersionBuilder struct {
	api.Params
}

// NewPagesGetVersionBuilder func.
func NewPagesGetVersionBuilder() *PagesGetVersionBuilder {
	return &PagesGetVersionBuilder{api.Params{}}
}

// VersionID parameter.
func (b *PagesGetVersionBuilder) VersionID(v int) *PagesGetVersionBuilder {
	b.Params["version_id"] = v
	return b
}

// GroupID ID of the community that owns the wiki page.
func (b *PagesGetVersionBuilder) GroupID(v int) *PagesGetVersionBuilder {
	b.Params["group_id"] = v
	return b
}

// UserID parameter.
func (b *PagesGetVersionBuilder) UserID(v int) *PagesGetVersionBuilder {
	b.Params["user_id"] = v
	return b
}

// NeedHTML '1' — to return the page as HTML.
func (b *PagesGetVersionBuilder) NeedHTML(v bool) *PagesGetVersionBuilder {
	b.Params["need_html"] = v
	return b
}

// PagesParseWikiBuilder builder.
//
// Returns HTML representation of the wiki markup.
//
// https://vk.com/dev/pages.parseWiki
type PagesParseWikiBuilder struct {
	api.Params
}

// NewPagesParseWikiBuilder func.
func NewPagesParseWikiBuilder() *PagesParseWikiBuilder {
	return &PagesParseWikiBuilder{api.Params{}}
}

// Text Text of the wiki page.
func (b *PagesParseWikiBuilder) Text(v string) *PagesParseWikiBuilder {
	b.Params["text"] = v
	return b
}

// GroupID ID of the group in the context of which this markup is interpreted.
func (b *PagesParseWikiBuilder) GroupID(v int) *PagesParseWikiBuilder {
	b.Params["group_id"] = v
	return b
}

// PagesSaveBuilder builder.
//
// Saves the text of a wiki page.
//
// https://vk.com/dev/pages.save
type PagesSaveBuilder struct {
	api.Params
}

// NewPagesSaveBuilder func.
func NewPagesSaveBuilder() *PagesSaveBuilder {
	return &PagesSaveBuilder{api.Params{}}
}

// Text Text of the wiki page in wiki-format.
func (b *PagesSaveBuilder) Text(v string) *PagesSaveBuilder {
	b.Params["text"] = v
	return b
}

// PageID wiki page ID. The 'title' parameter can be passed instead of 'pid'.
func (b *PagesSaveBuilder) PageID(v int) *PagesSaveBuilder {
	b.Params["page_id"] = v
	return b
}

// GroupID ID of the community that owns the wiki page.
func (b *PagesSaveBuilder) GroupID(v int) *PagesSaveBuilder {
	b.Params["group_id"] = v
	return b
}

// UserID parameter.
func (b *PagesSaveBuilder) UserID(v int) *PagesSaveBuilder {
	b.Params["user_id"] = v
	return b
}

// Title wiki page title.
func (b *PagesSaveBuilder) Title(v string) *PagesSaveBuilder {
	b.Params["title"] = v
	return b
}

// PagesSaveAccessBuilder builder.
//
// Saves modified read and edit access settings for a wiki page.
//
// https://vk.com/dev/pages.saveAccess
type PagesSaveAccessBuilder struct {
	api.Params
}

// NewPagesSaveAccessBuilder func.
func NewPagesSaveAccessBuilder() *PagesSaveAccessBuilder {
	return &PagesSaveAccessBuilder{api.Params{}}
}

// PageID wiki page ID.
func (b *PagesSaveAccessBuilder) PageID(v int) *PagesSaveAccessBuilder {
	b.Params["page_id"] = v
	return b
}

// GroupID ID of the community that owns the wiki page.
func (b *PagesSaveAccessBuilder) GroupID(v int) *PagesSaveAccessBuilder {
	b.Params["group_id"] = v
	return b
}

// UserID parameter.
func (b *PagesSaveAccessBuilder) UserID(v int) *PagesSaveAccessBuilder {
	b.Params["user_id"] = v
	return b
}

// View who can view the wiki page:
//
// * 1 — only community members;
//
// * 2 — all users can view the page;
//
// * 0 — only community managers.
func (b *PagesSaveAccessBuilder) View(v int) *PagesSaveAccessBuilder {
	b.Params["view"] = v
	return b
}

// Edit who can edit the wiki page:
//
// * 1 — only community members;
//
// * 2 — all users can edit the page;
//
// * 0 — only community managers.
func (b *PagesSaveAccessBuilder) Edit(v int) *PagesSaveAccessBuilder {
	b.Params["edit"] = v
	return b
}
