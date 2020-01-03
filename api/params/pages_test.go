package params_test

import (
	"testing"

	"github.com/SevereCloud/vksdk/api/params"
	"github.com/stretchr/testify/assert"
)

func TestPagesClearCacheBulder(t *testing.T) {
	b := params.NewPagesClearCacheBulder()

	b.URL("text")

	assert.Equal(t, b.Params["url"], "text")
}

func TestPagesGetBulder(t *testing.T) {
	b := params.NewPagesGetBulder()

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

func TestPagesGetHistoryBulder(t *testing.T) {
	b := params.NewPagesGetHistoryBulder()

	b.PageID(1)
	b.GroupID(1)
	b.UserID(1)

	assert.Equal(t, b.Params["page_id"], 1)
	assert.Equal(t, b.Params["group_id"], 1)
	assert.Equal(t, b.Params["user_id"], 1)
}

func TestPagesGetTitlesBulder(t *testing.T) {
	b := params.NewPagesGetTitlesBulder()

	b.GroupID(1)

	assert.Equal(t, b.Params["group_id"], 1)
}

func TestPagesGetVersionBulder(t *testing.T) {
	b := params.NewPagesGetVersionBulder()

	b.VersionID(1)
	b.GroupID(1)
	b.UserID(1)
	b.NeedHTML(true)

	assert.Equal(t, b.Params["version_id"], 1)
	assert.Equal(t, b.Params["group_id"], 1)
	assert.Equal(t, b.Params["user_id"], 1)
	assert.Equal(t, b.Params["need_html"], true)
}

func TestPagesParseWikiBulder(t *testing.T) {
	b := params.NewPagesParseWikiBulder()

	b.Text("text")
	b.GroupID(1)

	assert.Equal(t, b.Params["text"], "text")
	assert.Equal(t, b.Params["group_id"], 1)
}

func TestPagesSaveBulder(t *testing.T) {
	b := params.NewPagesSaveBulder()

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

func TestPagesSaveAccessBulder(t *testing.T) {
	b := params.NewPagesSaveAccessBulder()

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
