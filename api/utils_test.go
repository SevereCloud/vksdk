package api_test

import (
	"testing"

	"github.com/SevereCloud/vksdk/api"

	"github.com/stretchr/testify/assert"
)

func TestVK_UtilsCheckLink(t *testing.T) {
	t.Parallel()

	needUserToken(t)

	res, err := vkUser.UtilsCheckLink(api.Params{
		"url": "http://google.ru",
	})
	noError(t, err)
	assert.NotEmpty(t, res.Link)
	assert.NotEmpty(t, res.Status)
}

func TestVK_UtilsGetShortLink(t *testing.T) {
	t.Parallel()

	needUserToken(t)

	shortLink, err := vkUser.UtilsGetShortLink(api.Params{
		"url":     "http://google.ru",
		"private": 1,
	})
	noError(t, err)
	assert.NotEmpty(t, shortLink)

	res, err := vkUser.UtilsGetLastShortenedLinks(api.Params{})
	noError(t, err)
	assert.NotEmpty(t, res.Count)
	assert.NotEmpty(t, res.Items)

	del, err := vkUser.UtilsDeleteFromLastShortened(api.Params{
		"key": shortLink.Key,
	})
	noError(t, err)
	assert.NotEmpty(t, del)
}

func TestVK_UtilsGetLinkStats(t *testing.T) {
	t.Parallel()

	// BUG(VK): https://vk.com/bug202983
	needGroupToken(t)

	params := api.Params{
		"key":             "8TDuIz",
		"interval":        "month",
		"intervals_count": 12,
	}

	res, err := vkGroup.UtilsGetLinkStats(params)
	noError(t, err)
	assert.NotEmpty(t, res.Key)
	// assert.NotEmpty(t, res.Stats)

	resEx, err := vkGroup.UtilsGetLinkStatsExtended(params)
	noError(t, err)
	// assert.NotEmpty(t, resEx.Stats)
	assert.NotEmpty(t, resEx.Key)
}

func TestVK_UtilsGetServerTime(t *testing.T) {
	t.Parallel()

	needGroupToken(t)

	res, err := vkGroup.UtilsGetServerTime(api.Params{})
	noError(t, err)
	assert.NotEmpty(t, res)
}

func TestVK_UtilsResolveScreenName(t *testing.T) {
	t.Parallel()

	needGroupToken(t)

	f := func(name string) {
		t.Helper()

		_, err := vkGroup.UtilsResolveScreenName(api.Params{
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
