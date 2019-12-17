package api

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestVK_PodcastsGet(t *testing.T) {
	// FIXME: 7 Permission to perform this action is denied
	t.Skip("7 Permission to perform this action is denied")
	needUserToken(t)

	_, err := vkUser.PodcastsGetCatalog(Params{})
	assert.NoError(t, err)

	_, err = vkUser.PodcastsGetCatalogExtended(Params{})
	assert.NoError(t, err)

	_, err = vkUser.PodcastsGetCategories(Params{})
	assert.NoError(t, err)

	_, err = vkUser.PodcastsGetEpisodes(Params{
		"owner_id": -37473931,
	})
	assert.NoError(t, err)

	_, err = vkUser.PodcastsGetFeed(Params{})
	assert.NoError(t, err)

	_, err = vkUser.PodcastsGetFeedExtended(Params{})
	assert.NoError(t, err)

	_, err = vkUser.PodcastsGetStartPage(Params{})
	assert.NoError(t, err)

	_, err = vkUser.PodcastsGetStartPageExtended(Params{})
	assert.NoError(t, err)
}

func TestVK_PodcastsMarkAsListened(t *testing.T) {
	needUserToken(t)

	_, err := vkUser.PodcastsMarkAsListened(Params{
		"owner_id":   -76982440,
		"episode_id": 456239890,
	})
	assert.NoError(t, err)
}

func TestVK_PodcastsSubscribe(t *testing.T) {
	needUserToken(t)

	_, err := vkUser.PodcastsSubscribe(Params{
		"owner_id": -37473931,
	})
	assert.NoError(t, err)

	_, err = vkUser.PodcastsUnsubscribe(Params{
		"owner_id": -37473931,
	})
	assert.NoError(t, err)
}
