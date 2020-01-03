package params_test

import (
	"testing"

	"github.com/SevereCloud/vksdk/api/params"
	"github.com/stretchr/testify/assert"
)

func TestUtilsCheckLinkBulder(t *testing.T) {
	b := params.NewUtilsCheckLinkBulder()

	b.URL("text")

	assert.Equal(t, b.Params["url"], "text")
}

func TestUtilsDeleteFromLastShortenedBulder(t *testing.T) {
	b := params.NewUtilsDeleteFromLastShortenedBulder()

	b.Key("text")

	assert.Equal(t, b.Params["key"], "text")
}

func TestUtilsGetLastShortenedLinksBulder(t *testing.T) {
	b := params.NewUtilsGetLastShortenedLinksBulder()

	b.Count(1)
	b.Offset(1)

	assert.Equal(t, b.Params["count"], 1)
	assert.Equal(t, b.Params["offset"], 1)
}

func TestUtilsGetLinkStatsBulder(t *testing.T) {
	b := params.NewUtilsGetLinkStatsBulder()

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

func TestUtilsGetShortLinkBulder(t *testing.T) {
	b := params.NewUtilsGetShortLinkBulder()

	b.URL("text")
	b.Private(true)

	assert.Equal(t, b.Params["url"], "text")
	assert.Equal(t, b.Params["private"], true)
}

func TestUtilsResolveScreenNameBulder(t *testing.T) {
	b := params.NewUtilsResolveScreenNameBulder()

	b.ScreenName("text")

	assert.Equal(t, b.Params["screen_name"], "text")
}
