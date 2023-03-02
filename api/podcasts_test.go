package api_test

import (
	"testing"

	"github.com/SevereCloud/vksdk/v2/api"
)

func TestVK_PodcastsGet(t *testing.T) {
	t.Parallel()

	// FIXME: 7 Permission to perform this action is denied
	t.Skip("7 Permission to perform this action is denied")
	needUserToken(t)

	_, err := vkUser.PodcastsGetCatalog(nil)
	noError(t, err)

	_, err = vkUser.PodcastsGetCatalogExtended(nil)
	noError(t, err)

	_, err = vkUser.PodcastsGetCategories(nil)
	noError(t, err)

	_, err = vkUser.PodcastsGetEpisodes(api.Params{
		"owner_id": -37473931,
	})
	noError(t, err)

	_, err = vkUser.PodcastsGetFeed(nil)
	noError(t, err)

	_, err = vkUser.PodcastsGetFeedExtended(nil)
	noError(t, err)

	_, err = vkUser.PodcastsGetStartPage(nil)
	noError(t, err)

	_, err = vkUser.PodcastsGetStartPageExtended(nil)
	noError(t, err)
}

func TestVK_PodcastsMarkAsListened(t *testing.T) {
	t.Parallel()

	t.Skip("2301 Podcasts is not available for this owner")
	needUserToken(t)

	_, err := vkUser.PodcastsMarkAsListened(api.Params{
		"owner_id":   -76982440,
		"episode_id": 456239890,
	})
	noError(t, err)
}

func TestVK_PodcastsSubscribe(t *testing.T) {
	t.Parallel()

	t.Skip("104 Not found params")
	needUserToken(t)

	_, err := vkUser.PodcastsSubscribe(api.Params{
		"owner_id": -37473931,
	})
	noError(t, err)

	_, err = vkUser.PodcastsUnsubscribe(api.Params{
		"owner_id": -37473931,
	})
	noError(t, err)
}
