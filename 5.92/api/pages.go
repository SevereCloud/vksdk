package api // import "github.com/SevereCloud/vksdk/5.92/api"

import (
	"github.com/SevereCloud/vksdk/5.92/object"
)

// PagesClearCache allows to clear the cache of particular external pages which may be attached to VK posts.
// https://vk.com/dev/pages.clearCache
func (vk VK) PagesClearCache(params map[string]string) (vkErr Error) {
	_, vkErr = vk.Request("pages.clearCache", params)
	return
}

// PagesGetResponse struct
type PagesGetResponse object.PagesWikipageFull

// PagesGet returns information about a wiki page.
// https://vk.com/dev/pages.get
func (vk VK) PagesGet(params map[string]string) (response PagesGetResponse, vkErr Error) {
	vk.requestU("pages.get", params, &response, &vkErr)
	return
}

// PagesGetHistoryResponse struct
type PagesGetHistoryResponse []object.PagesWikipageVersion

// PagesGetHistory returns a list of all previous versions of a wiki page.
// https://vk.com/dev/pages.getHistory
func (vk VK) PagesGetHistory(params map[string]string) (response PagesGetHistoryResponse, vkErr Error) {
	vk.requestU("pages.getHistory", params, &response, &vkErr)
	return
}

// PagesGetTitlesResponse struct
type PagesGetTitlesResponse []object.PagesWikipageFull

// PagesGetTitles returns a list of wiki pages in a group.
// https://vk.com/dev/pages.getTitles
func (vk VK) PagesGetTitles(params map[string]string) (response PagesGetTitlesResponse, vkErr Error) {
	vk.requestU("pages.getTitles", params, &response, &vkErr)
	return
}

// PagesGetVersionResponse struct
type PagesGetVersionResponse []object.PagesWikipageFull

// PagesGetVersion returns the text of one of the previous versions of a wiki page.
// https://vk.com/dev/pages.getVersion
func (vk VK) PagesGetVersion(params map[string]string) (response PagesGetVersionResponse, vkErr Error) {
	vk.requestU("pages.getVersion", params, &response, &vkErr)
	return
}

// PagesParseWikiResponse struct
type PagesParseWikiResponse string

// PagesParseWiki returns HTML representation of the wiki markup.
// https://vk.com/dev/pages.parseWiki
func (vk VK) PagesParseWiki(params map[string]string) (response PagesParseWikiResponse, vkErr Error) {
	vk.requestU("pages.parseWiki", params, &response, &vkErr)
	return
}

// PagesSaveResponse struct
type PagesSaveResponse int

// PagesSave saves the text of a wiki page.
// https://vk.com/dev/pages.save
func (vk VK) PagesSave(params map[string]string) (response PagesSaveResponse, vkErr Error) {
	vk.requestU("pages.save", params, &response, &vkErr)
	return
}

// PagesSaveAccessResponse struct
type PagesSaveAccessResponse int

// PagesSaveAccess saves modified read and edit access settings for a wiki page.
// https://vk.com/dev/pages.saveAccess
func (vk VK) PagesSaveAccess(params map[string]string) (response PagesSaveAccessResponse, vkErr Error) {
	vk.requestU("pages.saveAccess", params, &response, &vkErr)
	return
}
