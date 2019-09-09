package api

import (
	"testing"
)

func TestVK_PodcastsGet(t *testing.T) {
	// FIXME: 7 Permission to perform this action is denied
	t.Skip("7 Permission to perform this action is denied")
	if vkUser.AccessToken == "" {
		t.Skip("USER_TOKEN empty")
	}

	_, gotVkErr := vkUser.PodcastsGetCatalog(map[string]string{})
	if gotVkErr.Code != 0 {
		t.Errorf("VK.PodcastsGetCatalog() gotVkErr = %v, want %v", gotVkErr, 0)
	}

	_, gotVkErr = vkUser.PodcastsGetCatalogExtended(map[string]string{})
	if gotVkErr.Code != 0 {
		t.Errorf("VK.PodcastsGetCatalogExtended() gotVkErr = %v, want %v", gotVkErr, 0)
	}

	_, gotVkErr = vkUser.PodcastsGetCategories(map[string]string{})
	if gotVkErr.Code != 0 {
		t.Errorf("VK.PodcastsGetCategories() gotVkErr = %v, want %v", gotVkErr, 0)
	}

	_, gotVkErr = vkUser.PodcastsGetEpisodes(map[string]string{
		"owner_id": "-37473931",
	})
	if gotVkErr.Code != 0 {
		t.Errorf("VK.PodcastsGetEpisodes() gotVkErr = %v, want %v", gotVkErr, 0)
	}

	_, gotVkErr = vkUser.PodcastsGetFeed(map[string]string{})
	if gotVkErr.Code != 0 {
		t.Errorf("VK.PodcastsGetFeed() gotVkErr = %v, want %v", gotVkErr, 0)
	}

	_, gotVkErr = vkUser.PodcastsGetFeedExtended(map[string]string{})
	if gotVkErr.Code != 0 {
		t.Errorf("VK.PodcastsGetFeedExtended() gotVkErr = %v, want %v", gotVkErr, 0)
	}

	_, gotVkErr = vkUser.PodcastsGetStartPage(map[string]string{})
	if gotVkErr.Code != 0 {
		t.Errorf("VK.PodcastsGetStartPage() gotVkErr = %v, want %v", gotVkErr, 0)
	}

	_, gotVkErr = vkUser.PodcastsGetStartPageExtended(map[string]string{})
	if gotVkErr.Code != 0 {
		t.Errorf("VK.PodcastsGetStartPageExtended() gotVkErr = %v, want %v", gotVkErr, 0)
	}
}

func TestVK_PodcastsMarkAsListened(t *testing.T) {
	if vkUser.AccessToken == "" {
		t.Skip("USER_TOKEN empty")
	}

	_, gotVkErr := vkUser.PodcastsMarkAsListened(map[string]string{
		"owner_id":   "-37473931",
		"episode_id": "456239025",
	})
	if gotVkErr.Code != 0 {
		t.Errorf("VK.PodcastsMarkAsListened() gotVkErr = %v, want %v", gotVkErr, 0)
	}
}

func TestVK_PodcastsSubscribe(t *testing.T) {
	if vkUser.AccessToken == "" {
		t.Skip("USER_TOKEN empty")
	}

	_, gotVkErr := vkUser.PodcastsSubscribe(map[string]string{
		"owner_id": "-37473931",
	})
	if gotVkErr.Code != 0 {
		t.Errorf("VK.PodcastsSubscribe() gotVkErr = %v, want %v", gotVkErr, 0)
	}

	_, gotVkErr = vkUser.PodcastsUnsubscribe(map[string]string{
		"owner_id": "-37473931",
	})
	if gotVkErr.Code != 0 {
		t.Errorf("VK.PodcastsUnsubscribe() gotVkErr = %v, want %v", gotVkErr, 0)
	}
}
