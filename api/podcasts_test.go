package api

import (
	"testing"
)

func TestVK_PodcastsGet(t *testing.T) {
	// FIXME: 7 Permission to perform this action is denied
	t.Skip("7 Permission to perform this action is denied")
	needUserToken(t)

	_, err := vkUser.PodcastsGetCatalog(map[string]string{})
	if err != nil {
		t.Errorf("VK.PodcastsGetCatalog() err = %v", err)
	}

	_, err = vkUser.PodcastsGetCatalogExtended(map[string]string{})
	if err != nil {
		t.Errorf("VK.PodcastsGetCatalogExtended() err = %v", err)
	}

	_, err = vkUser.PodcastsGetCategories(map[string]string{})
	if err != nil {
		t.Errorf("VK.PodcastsGetCategories() err = %v", err)
	}

	_, err = vkUser.PodcastsGetEpisodes(map[string]string{
		"owner_id": "-37473931",
	})
	if err != nil {
		t.Errorf("VK.PodcastsGetEpisodes() err = %v", err)
	}

	_, err = vkUser.PodcastsGetFeed(map[string]string{})
	if err != nil {
		t.Errorf("VK.PodcastsGetFeed() err = %v", err)
	}

	_, err = vkUser.PodcastsGetFeedExtended(map[string]string{})
	if err != nil {
		t.Errorf("VK.PodcastsGetFeedExtended() err = %v", err)
	}

	_, err = vkUser.PodcastsGetStartPage(map[string]string{})
	if err != nil {
		t.Errorf("VK.PodcastsGetStartPage() err = %v", err)
	}

	_, err = vkUser.PodcastsGetStartPageExtended(map[string]string{})
	if err != nil {
		t.Errorf("VK.PodcastsGetStartPageExtended() err = %v", err)
	}
}

func TestVK_PodcastsMarkAsListened(t *testing.T) {
	needUserToken(t)

	_, err := vkUser.PodcastsMarkAsListened(map[string]string{
		"owner_id":   "-37473931",
		"episode_id": "456239025",
	})
	if err != nil {
		t.Errorf("VK.PodcastsMarkAsListened() err = %v", err)
	}
}

func TestVK_PodcastsSubscribe(t *testing.T) {
	needUserToken(t)

	_, err := vkUser.PodcastsSubscribe(map[string]string{
		"owner_id": "-37473931",
	})
	if err != nil {
		t.Errorf("VK.PodcastsSubscribe() err = %v", err)
	}

	_, err = vkUser.PodcastsUnsubscribe(map[string]string{
		"owner_id": "-37473931",
	})
	if err != nil {
		t.Errorf("VK.PodcastsUnsubscribe() err = %v", err)
	}
}
