package params_test

import (
	"testing"

	"github.com/SevereCloud/vksdk/v3/api/params"
	"github.com/stretchr/testify/assert"
)

func TestPagesClearCacheBuilder(t *testing.T) {
	t.Parallel()

	b := params.NewPagesClearCacheBuilder()

	b.URL("text")

	assert.Equal(t, "text", b.Params["url"])
}

func TestPagesGetBuilder(t *testing.T) {
	t.Parallel()

	b := params.NewPagesGetBuilder()

	b.OwnerID(1)
	b.PageID(1)
	b.Global(true)
	b.SitePreview(true)
	b.Title("text")
	b.NeedSource(true)
	b.NeedHTML(true)

	assert.Equal(t, 1, b.Params["owner_id"])
	assert.Equal(t, 1, b.Params["page_id"])
	assert.Equal(t, true, b.Params["global"])
	assert.Equal(t, true, b.Params["site_preview"])
	assert.Equal(t, "text", b.Params["title"])
	assert.Equal(t, true, b.Params["need_source"])
	assert.Equal(t, true, b.Params["need_html"])
}

func TestPagesGetHistoryBuilder(t *testing.T) {
	t.Parallel()

	b := params.NewPagesGetHistoryBuilder()

	b.PageID(1)
	b.GroupID(1)
	b.UserID(1)

	assert.Equal(t, 1, b.Params["page_id"])
	assert.Equal(t, 1, b.Params["group_id"])
	assert.Equal(t, 1, b.Params["user_id"])
}

func TestPagesGetTitlesBuilder(t *testing.T) {
	t.Parallel()

	b := params.NewPagesGetTitlesBuilder()

	b.GroupID(1)

	assert.Equal(t, 1, b.Params["group_id"])
}

func TestPagesGetVersionBuilder(t *testing.T) {
	t.Parallel()

	b := params.NewPagesGetVersionBuilder()

	b.VersionID(1)
	b.GroupID(1)
	b.UserID(1)
	b.NeedHTML(true)

	assert.Equal(t, 1, b.Params["version_id"])
	assert.Equal(t, 1, b.Params["group_id"])
	assert.Equal(t, 1, b.Params["user_id"])
	assert.Equal(t, true, b.Params["need_html"])
}

func TestPagesParseWikiBuilder(t *testing.T) {
	t.Parallel()

	b := params.NewPagesParseWikiBuilder()

	b.Text("text")
	b.GroupID(1)

	assert.Equal(t, "text", b.Params["text"])
	assert.Equal(t, 1, b.Params["group_id"])
}

func TestPagesSaveBuilder(t *testing.T) {
	t.Parallel()

	b := params.NewPagesSaveBuilder()

	b.Text("text")
	b.PageID(1)
	b.GroupID(1)
	b.UserID(1)
	b.Title("text")

	assert.Equal(t, "text", b.Params["text"])
	assert.Equal(t, 1, b.Params["page_id"])
	assert.Equal(t, 1, b.Params["group_id"])
	assert.Equal(t, 1, b.Params["user_id"])
	assert.Equal(t, "text", b.Params["title"])
}

func TestPagesSaveAccessBuilder(t *testing.T) {
	t.Parallel()

	b := params.NewPagesSaveAccessBuilder()

	b.PageID(1)
	b.GroupID(1)
	b.UserID(1)
	b.View(1)
	b.Edit(1)

	assert.Equal(t, 1, b.Params["page_id"])
	assert.Equal(t, 1, b.Params["group_id"])
	assert.Equal(t, 1, b.Params["user_id"])
	assert.Equal(t, 1, b.Params["view"])
	assert.Equal(t, 1, b.Params["edit"])
}
