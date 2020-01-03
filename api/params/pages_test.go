package params_test

import (
	"testing"

	"github.com/SevereCloud/vksdk/api/params"
	"github.com/stretchr/testify/assert"
)

func TestPagesClearCacheBuilder(t *testing.T) {
	b := params.NewPagesClearCacheBuilder()

	b.URL("text")

	assert.Equal(t, b.Params["url"], "text")
}

func TestPagesGetBuilder(t *testing.T) {
	b := params.NewPagesGetBuilder()

	b.OwnerID(1)
	b.PageID(1)
	b.Global(true)
	b.SitePreview(true)
	b.Title("text")
	b.NeedSource(true)
	b.NeedHTML(true)

	assert.Equal(t, b.Params["owner_id"], 1)
	assert.Equal(t, b.Params["page_id"], 1)
	assert.Equal(t, b.Params["global"], true)
	assert.Equal(t, b.Params["site_preview"], true)
	assert.Equal(t, b.Params["title"], "text")
	assert.Equal(t, b.Params["need_source"], true)
	assert.Equal(t, b.Params["need_html"], true)
}

func TestPagesGetHistoryBuilder(t *testing.T) {
	b := params.NewPagesGetHistoryBuilder()

	b.PageID(1)
	b.GroupID(1)
	b.UserID(1)

	assert.Equal(t, b.Params["page_id"], 1)
	assert.Equal(t, b.Params["group_id"], 1)
	assert.Equal(t, b.Params["user_id"], 1)
}

func TestPagesGetTitlesBuilder(t *testing.T) {
	b := params.NewPagesGetTitlesBuilder()

	b.GroupID(1)

	assert.Equal(t, b.Params["group_id"], 1)
}

func TestPagesGetVersionBuilder(t *testing.T) {
	b := params.NewPagesGetVersionBuilder()

	b.VersionID(1)
	b.GroupID(1)
	b.UserID(1)
	b.NeedHTML(true)

	assert.Equal(t, b.Params["version_id"], 1)
	assert.Equal(t, b.Params["group_id"], 1)
	assert.Equal(t, b.Params["user_id"], 1)
	assert.Equal(t, b.Params["need_html"], true)
}

func TestPagesParseWikiBuilder(t *testing.T) {
	b := params.NewPagesParseWikiBuilder()

	b.Text("text")
	b.GroupID(1)

	assert.Equal(t, b.Params["text"], "text")
	assert.Equal(t, b.Params["group_id"], 1)
}

func TestPagesSaveBuilder(t *testing.T) {
	b := params.NewPagesSaveBuilder()

	b.Text("text")
	b.PageID(1)
	b.GroupID(1)
	b.UserID(1)
	b.Title("text")

	assert.Equal(t, b.Params["text"], "text")
	assert.Equal(t, b.Params["page_id"], 1)
	assert.Equal(t, b.Params["group_id"], 1)
	assert.Equal(t, b.Params["user_id"], 1)
	assert.Equal(t, b.Params["title"], "text")
}

func TestPagesSaveAccessBuilder(t *testing.T) {
	b := params.NewPagesSaveAccessBuilder()

	b.PageID(1)
	b.GroupID(1)
	b.UserID(1)
	b.View(1)
	b.Edit(1)

	assert.Equal(t, b.Params["page_id"], 1)
	assert.Equal(t, b.Params["group_id"], 1)
	assert.Equal(t, b.Params["user_id"], 1)
	assert.Equal(t, b.Params["view"], 1)
	assert.Equal(t, b.Params["edit"], 1)
}
