package params_test

import (
	"testing"

	"github.com/SevereCloud/vksdk/v2/api/params"
	"github.com/stretchr/testify/assert"
)

func TestUtilsCheckLinkBuilder(t *testing.T) {
	t.Parallel()

	b := params.NewUtilsCheckLinkBuilder()

	b.URL("text")

	assert.Equal(t, "text", b.Params["url"])
}

func TestUtilsDeleteFromLastShortenedBuilder(t *testing.T) {
	t.Parallel()

	b := params.NewUtilsDeleteFromLastShortenedBuilder()

	b.Key("text")

	assert.Equal(t, "text", b.Params["key"])
}

func TestUtilsGetLastShortenedLinksBuilder(t *testing.T) {
	t.Parallel()

	b := params.NewUtilsGetLastShortenedLinksBuilder()

	b.Count(1)
	b.Offset(1)

	assert.Equal(t, 1, b.Params["count"])
	assert.Equal(t, 1, b.Params["offset"])
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

	assert.Equal(t, "text", b.Params["key"])
	assert.Equal(t, "text", b.Params["source"])
	assert.Equal(t, "text", b.Params["access_key"])
	assert.Equal(t, "text", b.Params["interval"])
	assert.Equal(t, 1, b.Params["intervals_count"])
	assert.Equal(t, true, b.Params["extended"])
}

func TestUtilsGetShortLinkBuilder(t *testing.T) {
	t.Parallel()

	b := params.NewUtilsGetShortLinkBuilder()

	b.URL("text")
	b.Private(true)

	assert.Equal(t, "text", b.Params["url"])
	assert.Equal(t, true, b.Params["private"])
}

func TestUtilsResolveScreenNameBuilder(t *testing.T) {
	t.Parallel()

	b := params.NewUtilsResolveScreenNameBuilder()

	b.ScreenName("text")

	assert.Equal(t, "text", b.Params["screen_name"])
}
