package api_test

import (
	"testing"

	"github.com/SevereCloud/vksdk/api"

	"github.com/stretchr/testify/assert"
)

func TestVK_PagesClearCache(t *testing.T) {
	t.Parallel()

	needServiceToken(t)

	res, err := vkService.PagesClearCache(api.Params{
		"url": "https://ya.ru",
	})
	noError(t, err)
	assert.NotEmpty(t, res)
}

func TestVK_PagesGet(t *testing.T) {
	t.Parallel()

	needUserToken(t)

	res, err := vkUser.PagesGet(api.Params{
		"owner_id":  -87938575,
		"page_id":   51298167,
		"need_html": true,
	})
	noError(t, err)
	assert.NotEmpty(t, res.ID)
	assert.NotEmpty(t, res.GroupID)
	assert.NotEmpty(t, res.Title)
	assert.NotEmpty(t, res.WhoCanView)
	// assert.NotEmpty(t, res.WhoCanEdit)
	assert.NotEmpty(t, res.Edited)
	assert.NotEmpty(t, res.Created)
	assert.NotEmpty(t, res.Views)
	assert.NotEmpty(t, res.HTML)
	assert.NotEmpty(t, res.ViewURL)
}

func TestVK_PagesSave(t *testing.T) {
	t.Parallel()

	needUserToken(t)
	needGroupToken(t)

	page, err := vkUser.PagesSave(api.Params{
		"group_id": vkGroupID,
		"title":    "Test",
		"text":     "Test text",
	})
	noError(t, err)
	assert.NotEmpty(t, page)

	history, err := vkUser.PagesGetHistory(api.Params{
		"group_id": vkGroupID,
		"page_id":  page,
	})
	noError(t, err)

	if assert.NotEmpty(t, history) {
		assert.NotEmpty(t, history[0].Date)
		assert.NotEmpty(t, history[0].EditorID)
		assert.NotEmpty(t, history[0].EditorName)
		assert.NotEmpty(t, history[0].ID)
		assert.NotEmpty(t, history[0].Length)
	}

	version, err := vkUser.PagesGetVersion(api.Params{
		"group_id":   vkGroupID,
		"version_id": history[0].ID,
	})
	noError(t, err)
	assert.NotEmpty(t, version.ID)
	assert.NotEmpty(t, version.GroupID)
	assert.NotEmpty(t, version.Title)
	// assert.NotEmpty(t, version.WhoCanView)
	// assert.NotEmpty(t, version.WhoCanEdit)
	// assert.NotEmpty(t, version.Edited)
	// assert.NotEmpty(t, version.Created)
	// assert.NotEmpty(t, version.Views)

	res, err := vkUser.PagesSaveAccess(api.Params{
		"group_id": vkGroupID,
		"page_id":  page,
		"view":     false,
		"edit":     false,
	})
	noError(t, err)
	assert.NotEmpty(t, res)
}

func TestVK_PagesGetTitles(t *testing.T) {
	t.Parallel()

	needUserToken(t)

	res, err := vkUser.PagesGetTitles(api.Params{
		"group_id": 87938575,
	})
	noError(t, err)
	assert.NotEmpty(t, res)
}

func TestVK_PagesParseWiki(t *testing.T) {
	t.Parallel()

	needUserToken(t)

	res, err := vkUser.PagesParseWiki(api.Params{
		"text":     `[[photo-37273781_295853750|nolink;| ]]`,
		"group_id": 37273781,
	})
	noError(t, err)
	assert.NotEmpty(t, res)
}
