package api

import (
	"strconv"
	"testing"
)

func TestVK_PagesClearCache(t *testing.T) {
	if vkService.AccessToken == "" {
		t.Skip("USER_TOKEN empty")
	}

	_, gotVkErr := vkService.PagesClearCache(map[string]string{
		"url": "https://ya.ru",
	})
	if gotVkErr.Code != 0 {
		t.Errorf("VK.PagesClearCache() gotVkErr = %v, want %v", gotVkErr, 0)
	}
}

func TestVK_PagesGet(t *testing.T) {
	if vkUser.AccessToken == "" {
		t.Skip("USER_TOKEN empty")
	}

	_, gotVkErr := vkUser.PagesGet(map[string]string{
		"owner_id":  "-87938575",
		"page_id":   "51298167",
		"need_html": "1",
	})
	if gotVkErr.Code != 0 {
		t.Errorf("VK.PagesGet() gotVkErr = %v, want %v", gotVkErr, 0)
	}
}

func TestVK_PagesSave(t *testing.T) {
	if vkUser.AccessToken == "" || vkGroupID == 0 {
		t.Skip("USER_TOKEN empty or vkGroupID=0")
	}

	page, gotVkErr := vkUser.PagesSave(map[string]string{
		"group_id": strconv.Itoa(vkGroupID),
		"title":    "Test",
		"text":     "Test text",
	})
	if gotVkErr.Code != 0 {
		t.Errorf("VK.PagesSave() gotVkErr = %v, want %v", gotVkErr, 0)
	}

	history, gotVkErr := vkUser.PagesGetHistory(map[string]string{
		"group_id": strconv.Itoa(vkGroupID),
		"page_id":  strconv.Itoa(page),
	})
	if gotVkErr.Code != 0 {
		t.Errorf("VK.PagesGetHistory() gotVkErr = %v, want %v", gotVkErr, 0)
	}

	_, gotVkErr = vkUser.PagesGetVersion(map[string]string{
		"group_id":   strconv.Itoa(vkGroupID),
		"version_id": strconv.Itoa(history[0].ID),
	})
	if gotVkErr.Code != 0 {
		t.Errorf("VK.PagesGetVersion() gotVkErr = %v, want %v", gotVkErr, 0)
	}

	_, gotVkErr = vkUser.PagesSaveAccess(map[string]string{
		"group_id": strconv.Itoa(vkGroupID),
		"page_id":  strconv.Itoa(page),
		"view":     "0",
		"edit":     "0",
	})
	if gotVkErr.Code != 0 {
		t.Errorf("VK.PagesSaveAccess() gotVkErr = %v, want %v", gotVkErr, 0)
	}
}

func TestVK_PagesGetTitles(t *testing.T) {
	if vkUser.AccessToken == "" {
		t.Skip("USER_TOKEN empty")
	}

	_, gotVkErr := vkUser.PagesGetTitles(map[string]string{
		"group_id": "87938575",
	})
	if gotVkErr.Code != 0 {
		t.Errorf("VK.PagesGetTitles() gotVkErr = %v, want %v", gotVkErr, 0)
	}
}

func TestVK_PagesParseWiki(t *testing.T) {
	if vkUser.AccessToken == "" {
		t.Skip("USER_TOKEN empty")
	}

	_, gotVkErr := vkUser.PagesParseWiki(map[string]string{
		"text":     `[[photo-37273781_295853750|nolink;| ]]`,
		"group_id": "37273781",
	})
	if gotVkErr.Code != 0 {
		t.Errorf("VK.PagesParseWiki() gotVkErr = %v, want %v", gotVkErr, 0)
	}
}
