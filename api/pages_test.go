package api

import (
	"strconv"
	"testing"
)

func TestVK_PagesClearCache(t *testing.T) {
	needServiceToken(t)

	_, err := vkService.PagesClearCache(map[string]string{
		"url": "https://ya.ru",
	})
	if err != nil {
		t.Errorf("VK.PagesClearCache() err = %v", err)
	}
}

func TestVK_PagesGet(t *testing.T) {
	needUserToken(t)

	_, err := vkUser.PagesGet(map[string]string{
		"owner_id":  "-87938575",
		"page_id":   "51298167",
		"need_html": "1",
	})
	if err != nil {
		t.Errorf("VK.PagesGet() err = %v", err)
	}
}

func TestVK_PagesSave(t *testing.T) {
	needUserToken(t)
	needGroupToken(t)

	page, err := vkUser.PagesSave(map[string]string{
		"group_id": strconv.Itoa(vkGroupID),
		"title":    "Test",
		"text":     "Test text",
	})
	if err != nil {
		t.Errorf("VK.PagesSave() err = %v", err)
	}

	history, err := vkUser.PagesGetHistory(map[string]string{
		"group_id": strconv.Itoa(vkGroupID),
		"page_id":  strconv.Itoa(page),
	})
	if err != nil {
		t.Errorf("VK.PagesGetHistory() err = %v", err)
	}

	_, err = vkUser.PagesGetVersion(map[string]string{
		"group_id":   strconv.Itoa(vkGroupID),
		"version_id": strconv.Itoa(history[0].ID),
	})
	if err != nil {
		t.Errorf("VK.PagesGetVersion() err = %v", err)
	}

	_, err = vkUser.PagesSaveAccess(map[string]string{
		"group_id": strconv.Itoa(vkGroupID),
		"page_id":  strconv.Itoa(page),
		"view":     "0",
		"edit":     "0",
	})
	if err != nil {
		t.Errorf("VK.PagesSaveAccess() err = %v", err)
	}
}

func TestVK_PagesGetTitles(t *testing.T) {
	needUserToken(t)

	_, err := vkUser.PagesGetTitles(map[string]string{
		"group_id": "87938575",
	})
	if err != nil {
		t.Errorf("VK.PagesGetTitles() err = %v", err)
	}
}

func TestVK_PagesParseWiki(t *testing.T) {
	needUserToken(t)

	_, err := vkUser.PagesParseWiki(map[string]string{
		"text":     `[[photo-37273781_295853750|nolink;| ]]`,
		"group_id": "37273781",
	})
	if err != nil {
		t.Errorf("VK.PagesParseWiki() err = %v", err)
	}
}
