package api

import (
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestVK_PagesClearCache(t *testing.T) {
	needServiceToken(t)

	_, err := vkService.PagesClearCache(map[string]string{
		"url": "https://ya.ru",
	})
	assert.NoError(t, err)
}

func TestVK_PagesGet(t *testing.T) {
	needUserToken(t)

	_, err := vkUser.PagesGet(map[string]string{
		"owner_id":  "-87938575",
		"page_id":   "51298167",
		"need_html": "1",
	})
	assert.NoError(t, err)
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

	history, err := vkUser.PagesGetHistory(map[string]string{
		"group_id": strconv.Itoa(vkGroupID),
		"page_id":  strconv.Itoa(page),
	})
	assert.NoError(t, err)

	_, err = vkUser.PagesGetVersion(map[string]string{
		"group_id":   strconv.Itoa(vkGroupID),
		"version_id": strconv.Itoa(history[0].ID),
	})
	assert.NoError(t, err)

	_, err = vkUser.PagesSaveAccess(map[string]string{
		"group_id": strconv.Itoa(vkGroupID),
		"page_id":  strconv.Itoa(page),
		"view":     "0",
		"edit":     "0",
	})
	assert.NoError(t, err)
}

func TestVK_PagesGetTitles(t *testing.T) {
	needUserToken(t)

	_, err := vkUser.PagesGetTitles(map[string]string{
		"group_id": "87938575",
	})
	assert.NoError(t, err)
}

func TestVK_PagesParseWiki(t *testing.T) {
	needUserToken(t)

	_, err := vkUser.PagesParseWiki(map[string]string{
		"text":     `[[photo-37273781_295853750|nolink;| ]]`,
		"group_id": "37273781",
	})
	assert.NoError(t, err)
}
