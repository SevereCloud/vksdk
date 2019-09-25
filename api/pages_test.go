package api

import (
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestVK_PagesClearCache(t *testing.T) {
	needServiceToken(t)

	res, err := vkService.PagesClearCache(map[string]string{
		"url": "https://ya.ru",
	})
	assert.NoError(t, err)
	assert.NotEmpty(t, res)
}

func TestVK_PagesGet(t *testing.T) {
	needUserToken(t)

	res, err := vkUser.PagesGet(map[string]string{
		"owner_id":  "-87938575",
		"page_id":   "51298167",
		"need_html": "1",
	})
	assert.NoError(t, err)
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
	needUserToken(t)
	needGroupToken(t)

	page, err := vkUser.PagesSave(map[string]string{
		"group_id": strconv.Itoa(vkGroupID),
		"title":    "Test",
		"text":     "Test text",
	})
	assert.NoError(t, err)
	assert.NotEmpty(t, page)

	history, err := vkUser.PagesGetHistory(map[string]string{
		"group_id": strconv.Itoa(vkGroupID),
		"page_id":  strconv.Itoa(page),
	})
	assert.NoError(t, err)
	if assert.NotEmpty(t, history) {
		assert.NotEmpty(t, history[0].Date)
		assert.NotEmpty(t, history[0].EditorID)
		assert.NotEmpty(t, history[0].EditorName)
		assert.NotEmpty(t, history[0].ID)
		assert.NotEmpty(t, history[0].Length)
	}

	version, err := vkUser.PagesGetVersion(map[string]string{
		"group_id":   strconv.Itoa(vkGroupID),
		"version_id": strconv.Itoa(history[0].ID),
	})
	assert.NoError(t, err)
	assert.NotEmpty(t, version.ID)
	assert.NotEmpty(t, version.GroupID)
	assert.NotEmpty(t, version.Title)
	// assert.NotEmpty(t, version.WhoCanView)
	// assert.NotEmpty(t, version.WhoCanEdit)
	// assert.NotEmpty(t, version.Edited)
	// assert.NotEmpty(t, version.Created)
	// assert.NotEmpty(t, version.Views)

	res, err := vkUser.PagesSaveAccess(map[string]string{
		"group_id": strconv.Itoa(vkGroupID),
		"page_id":  strconv.Itoa(page),
		"view":     "0",
		"edit":     "0",
	})
	assert.NoError(t, err)
	assert.NotEmpty(t, res)
}

func TestVK_PagesGetTitles(t *testing.T) {
	needUserToken(t)

	res, err := vkUser.PagesGetTitles(map[string]string{
		"group_id": "87938575",
	})
	assert.NoError(t, err)
	assert.NotEmpty(t, res)
}

func TestVK_PagesParseWiki(t *testing.T) {
	needUserToken(t)

	res, err := vkUser.PagesParseWiki(map[string]string{
		"text":     `[[photo-37273781_295853750|nolink;| ]]`,
		"group_id": "37273781",
	})
	assert.NoError(t, err)
	assert.NotEmpty(t, res)
}
