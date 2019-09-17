package api

import (
	"testing"
)

func TestVK_UtilsCheckLink(t *testing.T) {
	needUserToken(t)

	_, err := vkUser.UtilsCheckLink(map[string]string{
		"url": "http://google.ru",
	})
	if err != nil {
		t.Errorf("VK.UtilsCheckLink() err = %v", err)
	}
}

func TestVK_UtilsGetShortLink(t *testing.T) {
	needUserToken(t)

	shortLink, err := vkUser.UtilsGetShortLink(map[string]string{
		"url":     "http://google.ru",
		"private": "1",
	})
	if err != nil {
		t.Errorf("VK.UtilsGetShortLink() err = %v", err)
	}

	_, err = vkUser.UtilsGetLastShortenedLinks(map[string]string{})
	if err != nil {
		t.Errorf("VK.UtilsGetLastShortenedLinks() err = %v", err)
	}

	_, err = vkUser.UtilsDeleteFromLastShortened(map[string]string{
		"key": shortLink.Key,
	})
	if err != nil {
		t.Errorf("VK.UtilsDeleteFromLastShortened() err = %v", err)
	}
}

func TestVK_UtilsGetLinkStats(t *testing.T) {
	needGroupToken(t)

	params := map[string]string{
		"key":             "8TDuIz",
		"interval":        "month",
		"intervals_count": "12",
	}

	_, err := vkGroup.UtilsGetLinkStats(params)
	if err != nil {
		t.Errorf("VK.UtilsGetLinkStats() err = %v", err)
	}

	_, err = vkGroup.UtilsGetLinkStatsExtended(params)
	if err != nil {
		t.Errorf("VK.UtilsGetLinkStats() err = %v", err)
	}
}

func TestVK_UtilsGetServerTime(t *testing.T) {
	needGroupToken(t)

	_, err := vkGroup.UtilsGetServerTime(map[string]string{})
	if err != nil {
		t.Errorf("VK.UtilsGetServerTime() err = %v", err)
	}
}

func TestVK_UtilsResolveScreenName(t *testing.T) {
	needGroupToken(t)

	f := func(name string) {
		t.Helper()
		_, err := vkGroup.UtilsResolveScreenName(map[string]string{
			"screen_name": name,
		})
		if err != nil {
			t.Errorf("VK.UtilsResolveScreenName() err = %v", err)
		}
	}

	f("durov")
	f("api")
	f("app6991405")
	f("z")
}
