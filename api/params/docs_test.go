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

	assert.Equal(t, 1, b.Params["owner_id"])
	assert.Equal(t, 1, b.Params["doc_id"])
	assert.Equal(t, "text", b.Params["access_key"])
}

func TestDocsDeleteBuilder(t *testing.T) {
	t.Parallel()

	b := params.NewDocsDeleteBuilder()

	b.OwnerID(1)
	b.DocID(1)

	assert.Equal(t, 1, b.Params["owner_id"])
	assert.Equal(t, 1, b.Params["doc_id"])
}

func TestDocsEditBuilder(t *testing.T) {
	t.Parallel()

	b := params.NewDocsEditBuilder()

	b.OwnerID(1)
	b.DocID(1)
	b.Title("text")
	b.Tags([]string{"text"})

	assert.Equal(t, 1, b.Params["owner_id"])
	assert.Equal(t, 1, b.Params["doc_id"])
	assert.Equal(t, "text", b.Params["title"])
	assert.Equal(t, []string{"text"}, b.Params["tags"])
}

func TestDocsGetBuilder(t *testing.T) {
	t.Parallel()

	b := params.NewDocsGetBuilder()

	b.Count(1)
	b.Offset(1)
	b.Type(1)
	b.OwnerID(1)

	assert.Equal(t, 1, b.Params["count"])
	assert.Equal(t, 1, b.Params["offset"])
	assert.Equal(t, 1, b.Params["type"])
	assert.Equal(t, 1, b.Params["owner_id"])
}

func TestDocsGetByIDBuilder(t *testing.T) {
	t.Parallel()

	b := params.NewDocsGetByIDBuilder()

	b.Docs([]string{"text"})

	assert.Equal(t, []string{"text"}, b.Params["docs"])
}

func TestDocsGetMessagesUploadServerBuilder(t *testing.T) {
	t.Parallel()

	b := params.NewDocsGetMessagesUploadServerBuilder()

	b.Type("text")
	b.PeerID(1)

	assert.Equal(t, "text", b.Params["type"])
	assert.Equal(t, 1, b.Params["peer_id"])
}

func TestDocsGetTypesBuilder(t *testing.T) {
	t.Parallel()

	b := params.NewDocsGetTypesBuilder()

	b.OwnerID(1)

	assert.Equal(t, 1, b.Params["owner_id"])
}

func TestDocsGetUploadServerBuilder(t *testing.T) {
	t.Parallel()

	b := params.NewDocsGetUploadServerBuilder()

	b.GroupID(1)

	assert.Equal(t, 1, b.Params["group_id"])
}

func TestDocsGetWallUploadServerBuilder(t *testing.T) {
	t.Parallel()

	b := params.NewDocsGetWallUploadServerBuilder()

	b.GroupID(1)

	assert.Equal(t, 1, b.Params["group_id"])
}

func TestDocsSaveBuilder(t *testing.T) {
	t.Parallel()

	b := params.NewDocsSaveBuilder()

	b.File("text")
	b.Title("text")
	b.Tags("text")

	assert.Equal(t, "text", b.Params["file"])
	assert.Equal(t, "text", b.Params["title"])
	assert.Equal(t, "text", b.Params["tags"])
}

func TestDocsSearchBuilder(t *testing.T) {
	t.Parallel()

	b := params.NewDocsSearchBuilder()

	b.Q("text")
	b.SearchOwn(true)
	b.Count(1)
	b.Offset(1)

	assert.Equal(t, "text", b.Params["q"])
	assert.Equal(t, true, b.Params["search_own"])
	assert.Equal(t, 1, b.Params["count"])
	assert.Equal(t, 1, b.Params["offset"])
}
