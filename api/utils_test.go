package api

import (
	"testing"
)

func TestVK_UtilsCheckLink(t *testing.T) {
	if vkUser.AccessToken == "" {
		t.Skip("USER_TOKEN empty")
	}

	_, gotVkErr := vkUser.UtilsCheckLink(map[string]string{
		"url": "http://google.ru",
	})
	if gotVkErr.Code != 0 {
		t.Errorf("VK.UtilsCheckLink() gotVkErr = %v", gotVkErr)
	}
}

func TestVK_UtilsGetShortLink(t *testing.T) {
	if vkUser.AccessToken == "" {
		t.Skip("USER_TOKEN empty")
	}

	shortLink, gotVkErr := vkUser.UtilsGetShortLink(map[string]string{
		"url":     "http://google.ru",
		"private": "1",
	})
	if gotVkErr.Code != 0 {
		t.Errorf("VK.UtilsGetShortLink() gotVkErr = %v", gotVkErr)
	}

	_, gotVkErr = vkUser.UtilsGetLastShortenedLinks(map[string]string{})
	if gotVkErr.Code != 0 {
		t.Errorf("VK.UtilsGetLastShortenedLinks() gotVkErr = %v", gotVkErr)
	}

	_, gotVkErr = vkUser.UtilsDeleteFromLastShortened(map[string]string{
		"key": shortLink.Key,
	})
	if gotVkErr.Code != 0 {
		t.Errorf("VK.UtilsDeleteFromLastShortened() gotVkErr = %v", gotVkErr)
	}
}

func TestVK_UtilsGetLinkStats(t *testing.T) {
	if vkGroup.AccessToken == "" {
		t.Skip("GROUP_TOKEN empty")
	}

	params := map[string]string{
		"key":             "8TDuIz",
		"interval":        "month",
		"intervals_count": "12",
	}

	_, gotVkErr := vkGroup.UtilsGetLinkStats(params)
	if gotVkErr.Code != 0 {
		t.Errorf("VK.UtilsGetLinkStats() gotVkErr = %v", gotVkErr)
	}

	_, gotVkErr = vkGroup.UtilsGetLinkStatsExtended(params)
	if gotVkErr.Code != 0 {
		t.Errorf("VK.UtilsGetLinkStats() gotVkErr = %v", gotVkErr)
	}
}

func TestVK_UtilsGetServerTime(t *testing.T) {
	if vkGroup.AccessToken == "" {
		t.Skip("GROUP_TOKEN empty")
	}

	_, gotVkErr := vkGroup.UtilsGetServerTime(map[string]string{})
	if gotVkErr.Code != 0 {
		t.Errorf("VK.UtilsGetServerTime() gotVkErr = %v", gotVkErr)
	}
}

func TestVK_UtilsResolveScreenName(t *testing.T) {
	if vkGroup.AccessToken == "" {
		t.Skip("GROUP_TOKEN empty")
	}

	f := func(name string) {
		t.Helper()
		_, gotVkErr := vkGroup.UtilsResolveScreenName(map[string]string{
			"screen_name": name,
		})
		if gotVkErr.Code != 0 {
			t.Errorf("VK.UtilsResolveScreenName() gotVkErr = %v", gotVkErr)
		}
	}

	f("durov")
	f("api")
	f("app6991405")
	f("z")
}
