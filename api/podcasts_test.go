package api

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestVK_PodcastsGet(t *testing.T) {
	// FIXME: 7 Permission to perform this action is denied
	t.Skip("7 Permission to perform this action is denied")
	needUserToken(t)

	_, err := vkUser.PodcastsGetCatalog(map[string]string{})
	assert.NoError(t, err)

	_, err = vkUser.PodcastsGetCatalogExtended(map[string]string{})
	assert.NoError(t, err)

	_, err = vkUser.PodcastsGetCategories(map[string]string{})
	assert.NoError(t, err)

	_, err = vkUser.PodcastsGetEpisodes(map[string]string{
		"owner_id": "-37473931",
	})
	assert.NoError(t, err)

	_, err = vkUser.PodcastsGetFeed(map[string]string{})
	assert.NoError(t, err)

	_, err = vkUser.PodcastsGetFeedExtended(map[string]string{})
	assert.NoError(t, err)

	_, err = vkUser.PodcastsGetStartPage(map[string]string{})
	assert.NoError(t, err)

	_, err = vkUser.PodcastsGetStartPageExtended(map[string]string{})
	assert.NoError(t, err)
}

func TestVK_PodcastsMarkAsListened(t *testing.T) {
	needUserToken(t)

	_, err := vkUser.PodcastsMarkAsListened(map[string]string{
		"owner_id":   "-37473931",
		"episode_id": "456239025",
	})
	assert.NoError(t, err)
}

func TestVK_PodcastsSubscribe(t *testing.T) {
	needUserToken(t)

	_, err := vkUser.PodcastsSubscribe(map[string]string{
		"owner_id": "-37473931",
	})
	assert.NoError(t, err)

	_, err = vkUser.PodcastsUnsubscribe(map[string]string{
		"owner_id": "-37473931",
	})
	assert.NoError(t, err)
}
