package params_test

import (
	"testing"

	"github.com/SevereCloud/vksdk/v2/api/params"
	"github.com/stretchr/testify/assert"
)

func TestDocsAddBuilder(t *testing.T) {
	t.Parallel()

	b := params.NewDocsAddBuilder()

	b.OwnerID(1)
	b.DocID(1)
	b.AccessKey("text")

	assert.Equal(t, b.Params["owner_id"], 1)
	assert.Equal(t, b.Params["doc_id"], 1)
	assert.Equal(t, b.Params["access_key"], "text")
}

func TestDocsDeleteBuilder(t *testing.T) {
	t.Parallel()

	b := params.NewDocsDeleteBuilder()

	b.OwnerID(1)
	b.DocID(1)

	assert.Equal(t, b.Params["owner_id"], 1)
	assert.Equal(t, b.Params["doc_id"], 1)
}

func TestDocsEditBuilder(t *testing.T) {
	t.Parallel()

	b := params.NewDocsEditBuilder()

	b.OwnerID(1)
	b.DocID(1)
	b.Title("text")
	b.Tags([]string{"text"})

	assert.Equal(t, b.Params["owner_id"], 1)
	assert.Equal(t, b.Params["doc_id"], 1)
	assert.Equal(t, b.Params["title"], "text")
	assert.Equal(t, b.Params["tags"], []string{"text"})
}

func TestDocsGetBuilder(t *testing.T) {
	t.Parallel()

	b := params.NewDocsGetBuilder()

	b.Count(1)
	b.Offset(1)
	b.Type(1)
	b.OwnerID(1)

	assert.Equal(t, b.Params["count"], 1)
	assert.Equal(t, b.Params["offset"], 1)
	assert.Equal(t, b.Params["type"], 1)
	assert.Equal(t, b.Params["owner_id"], 1)
}

func TestDocsGetByIDBuilder(t *testing.T) {
	t.Parallel()

	b := params.NewDocsGetByIDBuilder()

	b.Docs([]string{"text"})

	assert.Equal(t, b.Params["docs"], []string{"text"})
}

func TestDocsGetMessagesUploadServerBuilder(t *testing.T) {
	t.Parallel()

	b := params.NewDocsGetMessagesUploadServerBuilder()

	b.Type("text")
	b.PeerID(1)

	assert.Equal(t, b.Params["type"], "text")
	assert.Equal(t, b.Params["peer_id"], 1)
}

func TestDocsGetTypesBuilder(t *testing.T) {
	t.Parallel()

	b := params.NewDocsGetTypesBuilder()

	b.OwnerID(1)

	assert.Equal(t, b.Params["owner_id"], 1)
}

func TestDocsGetUploadServerBuilder(t *testing.T) {
	t.Parallel()

	b := params.NewDocsGetUploadServerBuilder()

	b.GroupID(1)

	assert.Equal(t, b.Params["group_id"], 1)
}

func TestDocsGetWallUploadServerBuilder(t *testing.T) {
	t.Parallel()

	b := params.NewDocsGetWallUploadServerBuilder()

	b.GroupID(1)

	assert.Equal(t, b.Params["group_id"], 1)
}

func TestDocsSaveBuilder(t *testing.T) {
	t.Parallel()

	b := params.NewDocsSaveBuilder()

	b.File("text")
	b.Title("text")
	b.Tags("text")

	assert.Equal(t, b.Params["file"], "text")
	assert.Equal(t, b.Params["title"], "text")
	assert.Equal(t, b.Params["tags"], "text")
}

func TestDocsSearchBuilder(t *testing.T) {
	t.Parallel()

	b := params.NewDocsSearchBuilder()

	b.Q("text")
	b.SearchOwn(true)
	b.Count(1)
	b.Offset(1)

	assert.Equal(t, b.Params["q"], "text")
	assert.Equal(t, b.Params["search_own"], true)
	assert.Equal(t, b.Params["count"], 1)
	assert.Equal(t, b.Params["offset"], 1)
}
