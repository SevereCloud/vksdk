package params_test

import (
	"testing"

	"github.com/Derad6709/vksdk/v2/api/params"
	"github.com/stretchr/testify/assert"
)

func TestUtilsCheckLinkBuilder(t *testing.T) {
	t.Parallel()

	b := params.NewUtilsCheckLinkBuilder()

	b.URL("text")

	assert.Equal(t, b.Params["url"], "text")
}

func TestUtilsDeleteFromLastShortenedBuilder(t *testing.T) {
	t.Parallel()

	b := params.NewUtilsDeleteFromLastShortenedBuilder()

	b.Key("text")

	assert.Equal(t, b.Params["key"], "text")
}

func TestUtilsGetLastShortenedLinksBuilder(t *testing.T) {
	t.Parallel()

	b := params.NewUtilsGetLastShortenedLinksBuilder()

	b.Count(1)
	b.Offset(1)

	assert.Equal(t, b.Params["count"], 1)
	assert.Equal(t, b.Params["offset"], 1)
}

func TestUtilsGetLinkStatsBuilder(t *testing.T) {
	t.Parallel()

	b := params.NewUtilsGetLinkStatsBuilder()

	b.Key("text")
	b.Source("text")
	b.AccessKey("text")
	b.Interval("text")
	b.IntervalsCount(1)
	b.Extended(true)

	assert.Equal(t, b.Params["key"], "text")
	assert.Equal(t, b.Params["source"], "text")
	assert.Equal(t, b.Params["access_key"], "text")
	assert.Equal(t, b.Params["interval"], "text")
	assert.Equal(t, b.Params["intervals_count"], 1)
	assert.Equal(t, b.Params["extended"], true)
}

func TestUtilsGetShortLinkBuilder(t *testing.T) {
	t.Parallel()

	b := params.NewUtilsGetShortLinkBuilder()

	b.URL("text")
	b.Private(true)

	assert.Equal(t, b.Params["url"], "text")
	assert.Equal(t, b.Params["private"], true)
}

func TestUtilsResolveScreenNameBuilder(t *testing.T) {
	t.Parallel()

	b := params.NewUtilsResolveScreenNameBuilder()

	b.ScreenName("text")

	assert.Equal(t, b.Params["screen_name"], "text")
}
