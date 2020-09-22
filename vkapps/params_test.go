package vkapps_test

import (
	"net/url"
	"testing"

	"github.com/SevereCloud/vksdk/v2/vkapps"
	"github.com/stretchr/testify/assert"
)

func TestParams(t *testing.T) {
	t.Parallel()

	link := "https://severecloud.github.io/vk-limits/?vk_access_token_settings=notify,menu&vk_app_id=7573302&vk_are_notifications_enabled=0&vk_is_app_user=1&vk_is_favorite=0&vk_language=ru&vk_platform=desktop_web&vk_ref=left_nav&vk_user_id=117253521&sign=sign"
	u, _ := url.Parse(link)

	params, err := vkapps.NewParams(u)
	if !assert.NoError(t, err) {
		t.FailNow()
	}

	assert.Equal(t, "notify,menu", params.VkAccessTokenSettings)
	assert.Equal(t, 7573302, params.VkAppID)
	assert.Equal(t, false, params.VkAreNotificationsEnabled)
	assert.Equal(t, true, params.VkIsAppUser)
	assert.Equal(t, false, params.VkIsFavorite)
	assert.Equal(t, "ru", params.VkLanguage)
	assert.Equal(t, vkapps.DesktopWeb, params.VkPlatform)
	assert.Equal(t, vkapps.LeftNav, params.VkRef)
	assert.Equal(t, 117253521, params.VkUserID)
	assert.Equal(t, "sign", params.Sign)
}

func TestReferral_Catalog(t *testing.T) {
	t.Parallel()

	f := func(expected string, ref vkapps.Referral) {
		assert.Equal(t, expected, ref.Catalog())
	}

	f("", vkapps.LeftNav)
	f("test", vkapps.Referral("catalog_test"))
}

func TestReferral_Story(t *testing.T) {
	t.Parallel()

	f := func(expectedUserID int, expectedParams string, ref vkapps.Referral) {
		userID, params := ref.Story()
		assert.Equal(t, expectedUserID, userID)
		assert.Equal(t, expectedParams, params)
	}

	f(0, "", vkapps.LeftNav)
	f(0, "", vkapps.Referral("storya_test"))
	f(1, "1_test", vkapps.Referral("story1_1_test"))
}
