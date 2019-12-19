package api_test

import (
	"testing"

	"github.com/SevereCloud/vksdk/api"

	"github.com/stretchr/testify/assert"
)

func TestVK_PodcastsGet(t *testing.T) {
	// FIXME: 7 Permission to perform this action is denied
	t.Skip("7 Permission to perform this action is denied")
	needUserToken(t)

	_, err := vkUser.PodcastsGetCatalog(api.Params{})
	assert.NoError(t, err)

	_, err = vkUser.PodcastsGetCatalogExtended(api.Params{})
	assert.NoError(t, err)

	_, err = vkUser.PodcastsGetCategories(api.Params{})
	assert.NoError(t, err)

	_, err = vkUser.PodcastsGetEpisodes(api.Params{
		"owner_id": -37473931,
	})
	assert.NoError(t, err)

	_, err = vkUser.PodcastsGetFeed(api.Params{})
	assert.NoError(t, err)

	_, err = vkUser.PodcastsGetFeedExtended(api.Params{})
	assert.NoError(t, err)

	_, err = vkUser.PodcastsGetStartPage(api.Params{})
	assert.NoError(t, err)

	_, err = vkUser.PodcastsGetStartPageExtended(api.Params{})
	assert.NoError(t, err)
}

func TestVK_PodcastsMarkAsListened(t *testing.T) {
	needUserToken(t)

	_, err := vkUser.PodcastsMarkAsListened(api.Params{
		"owner_id":   -76982440,
		"episode_id": 456239890,
	})
	assert.NoError(t, err)
}

func TestVK_PodcastsSubscribe(t *testing.T) {
	needUserToken(t)

	_, err := vkUser.PodcastsSubscribe(api.Params{
		"owner_id": -37473931,
	})
	assert.NoError(t, err)

	_, err = vkUser.PodcastsUnsubscribe(api.Params{
		"owner_id": -37473931,
	})
	assert.NoError(t, err)
}
