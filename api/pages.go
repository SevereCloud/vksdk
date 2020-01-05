package api // import "github.com/SevereCloud/vksdk/api"

import (
	"github.com/SevereCloud/vksdk/object"
)

// PagesClearCache allows to clear the cache of particular external pages which may be attached to VK posts.
//
// https://vk.com/dev/pages.clearCache
func (vk *VK) PagesClearCache(params Params) (response int, err error) {
	err = vk.RequestUnmarshal("pages.clearCache", params, &response)
	return
}

// PagesGetResponse struct
type PagesGetResponse object.PagesWikipageFull

// PagesGet returns information about a wiki page.
//
// https://vk.com/dev/pages.get
func (vk *VK) PagesGet(params Params) (response PagesGetResponse, err error) {
	err = vk.RequestUnmarshal("pages.get", params, &response)
	return
}

// PagesGetHistoryResponse struct
type PagesGetHistoryResponse []object.PagesWikipageHistory

// PagesGetHistory returns a list of all previous versions of a wiki page.
//
// https://vk.com/dev/pages.getHistory
func (vk *VK) PagesGetHistory(params Params) (response PagesGetHistoryResponse, err error) {
	err = vk.RequestUnmarshal("pages.getHistory", params, &response)
	return
}

// PagesGetTitlesResponse struct
type PagesGetTitlesResponse []object.PagesWikipageFull

// PagesGetTitles returns a list of wiki pages in a group.
//
// https://vk.com/dev/pages.getTitles
func (vk *VK) PagesGetTitles(params Params) (response PagesGetTitlesResponse, err error) {
	err = vk.RequestUnmarshal("pages.getTitles", params, &response)
	return
}

// PagesGetVersionResponse struct
type PagesGetVersionResponse object.PagesWikipageFull

// PagesGetVersion returns the text of one of the previous versions of a wiki page.
//
// https://vk.com/dev/pages.getVersion
func (vk *VK) PagesGetVersion(params Params) (response PagesGetVersionResponse, err error) {
	err = vk.RequestUnmarshal("pages.getVersion", params, &response)
	return
}

// PagesParseWiki returns HTML representation of the wiki markup.
//
// https://vk.com/dev/pages.parseWiki
func (vk *VK) PagesParseWiki(params Params) (response string, err error) {
	err = vk.RequestUnmarshal("pages.parseWiki", params, &response)
	return
}

// PagesSave saves the text of a wiki page.
//
// https://vk.com/dev/pages.save
func (vk *VK) PagesSave(params Params) (response int, err error) {
	err = vk.RequestUnmarshal("pages.save", params, &response)
	return
}

// PagesSaveAccess saves modified read and edit access settings for a wiki page.
//
// https://vk.com/dev/pages.saveAccess
func (vk *VK) PagesSaveAccess(params Params) (response int, err error) {
	err = vk.RequestUnmarshal("pages.saveAccess", params, &response)
	return
}
