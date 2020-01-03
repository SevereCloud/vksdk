package params_test

import (
	"testing"

	"github.com/SevereCloud/vksdk/api/params"
	"github.com/stretchr/testify/assert"
)

func TestDocsAddBulder(t *testing.T) {
	b := params.NewDocsAddBulder()

	b.OwnerID(1)
	b.DocID(1)
	b.AccessKey("text")

	assert.Equal(t, b.Params["owner_id"], 1)
	assert.Equal(t, b.Params["doc_id"], 1)
	assert.Equal(t, b.Params["access_key"], "text")
}

func TestDocsDeleteBulder(t *testing.T) {
	b := params.NewDocsDeleteBulder()

	b.OwnerID(1)
	b.DocID(1)

	assert.Equal(t, b.Params["owner_id"], 1)
	assert.Equal(t, b.Params["doc_id"], 1)
}

func TestDocsEditBulder(t *testing.T) {
	b := params.NewDocsEditBulder()

	b.OwnerID(1)
	b.DocID(1)
	b.Title("text")
	b.Tags([]string{"text"})

	assert.Equal(t, b.Params["owner_id"], 1)
	assert.Equal(t, b.Params["doc_id"], 1)
	assert.Equal(t, b.Params["title"], "text")
	assert.Equal(t, b.Params["tags"], []string{"text"})
}

func TestDocsGetBulder(t *testing.T) {
	b := params.NewDocsGetBulder()

	b.Count(1)
	b.Offset(1)
	b.Type(1)
	b.OwnerID(1)

	assert.Equal(t, b.Params["count"], 1)
	assert.Equal(t, b.Params["offset"], 1)
	assert.Equal(t, b.Params["type"], 1)
	assert.Equal(t, b.Params["owner_id"], 1)
}

func TestDocsGetByIDBulder(t *testing.T) {
	b := params.NewDocsGetByIDBulder()

	b.Docs([]string{"text"})

	assert.Equal(t, b.Params["docs"], []string{"text"})
}

func TestDocsGetMessagesUploadServerBulder(t *testing.T) {
	b := params.NewDocsGetMessagesUploadServerBulder()

	b.Type("text")
	b.PeerID(1)

	assert.Equal(t, b.Params["type"], "text")
	assert.Equal(t, b.Params["peer_id"], 1)
}

func TestDocsGetTypesBulder(t *testing.T) {
	b := params.NewDocsGetTypesBulder()

	b.OwnerID(1)

	assert.Equal(t, b.Params["owner_id"], 1)
}

func TestDocsGetUploadServerBulder(t *testing.T) {
	b := params.NewDocsGetUploadServerBulder()

	b.GroupID(1)

	assert.Equal(t, b.Params["group_id"], 1)
}

func TestDocsGetWallUploadServerBulder(t *testing.T) {
	b := params.NewDocsGetWallUploadServerBulder()

	b.GroupID(1)

	assert.Equal(t, b.Params["group_id"], 1)
}

func TestDocsSaveBulder(t *testing.T) {
	b := params.NewDocsSaveBulder()

	b.File("text")
	b.Title("text")
	b.Tags("text")

	assert.Equal(t, b.Params["file"], "text")
	assert.Equal(t, b.Params["title"], "text")
	assert.Equal(t, b.Params["tags"], "text")
}

func TestDocsSearchBulder(t *testing.T) {
	b := params.NewDocsSearchBulder()

	b.Q("text")
	b.SearchOwn(true)
	b.Count(1)
	b.Offset(1)

	assert.Equal(t, b.Params["q"], "text")
	assert.Equal(t, b.Params["search_own"], true)
	assert.Equal(t, b.Params["count"], 1)
	assert.Equal(t, b.Params["offset"], 1)
}
