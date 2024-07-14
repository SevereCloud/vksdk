package params_test

import (
	"testing"

	"github.com/SevereCloud/vksdk/v3/api/params"
	"github.com/stretchr/testify/assert"
)

func TestFaveAddArticleBuilder(t *testing.T) {
	t.Parallel()

	b := params.NewFaveAddArticleBuilder()

	b.URL("text")

	assert.Equal(t, "text", b.Params["url"])
}

func TestFaveAddLinkBuilder(t *testing.T) {
	t.Parallel()

	b := params.NewFaveAddLinkBuilder()

	b.Link("text")

	assert.Equal(t, "text", b.Params["link"])
}

func TestFaveAddPageBuilder(t *testing.T) {
	t.Parallel()

	b := params.NewFaveAddPageBuilder()

	b.UserID(1)
	b.GroupID(1)

	assert.Equal(t, 1, b.Params["user_id"])
	assert.Equal(t, 1, b.Params["group_id"])
}

func TestFaveAddPostBuilder(t *testing.T) {
	t.Parallel()

	b := params.NewFaveAddPostBuilder()

	b.OwnerID(1)
	b.ID(1)
	b.AccessKey("text")

	assert.Equal(t, 1, b.Params["owner_id"])
	assert.Equal(t, 1, b.Params["id"])
	assert.Equal(t, "text", b.Params["access_key"])
}

func TestFaveAddProductBuilder(t *testing.T) {
	t.Parallel()

	b := params.NewFaveAddProductBuilder()

	b.OwnerID(1)
	b.ID(1)
	b.AccessKey("text")

	assert.Equal(t, 1, b.Params["owner_id"])
	assert.Equal(t, 1, b.Params["id"])
	assert.Equal(t, "text", b.Params["access_key"])
}

func TestFaveAddTagBuilder(t *testing.T) {
	t.Parallel()

	b := params.NewFaveAddTagBuilder()

	b.Name("text")

	assert.Equal(t, "text", b.Params["name"])
}

func TestFaveAddVideoBuilder(t *testing.T) {
	t.Parallel()

	b := params.NewFaveAddVideoBuilder()

	b.OwnerID(1)
	b.ID(1)
	b.AccessKey("text")

	assert.Equal(t, 1, b.Params["owner_id"])
	assert.Equal(t, 1, b.Params["id"])
	assert.Equal(t, "text", b.Params["access_key"])
}

func TestFaveEditTagBuilder(t *testing.T) {
	t.Parallel()

	b := params.NewFaveEditTagBuilder()

	b.ID(1)
	b.Name("text")

	assert.Equal(t, 1, b.Params["id"])
	assert.Equal(t, "text", b.Params["name"])
}

func TestFaveGetBuilder(t *testing.T) {
	t.Parallel()

	b := params.NewFaveGetBuilder()

	b.Extended(true)
	b.ItemType("text")
	b.TagID(1)
	b.Offset(1)
	b.Count(1)
	b.Fields("text")
	b.IsFromSnackbar(true)

	assert.Equal(t, true, b.Params["extended"])
	assert.Equal(t, "text", b.Params["item_type"])
	assert.Equal(t, 1, b.Params["tag_id"])
	assert.Equal(t, 1, b.Params["offset"])
	assert.Equal(t, 1, b.Params["count"])
	assert.Equal(t, "text", b.Params["fields"])
	assert.Equal(t, true, b.Params["is_from_snackbar"])
}

func TestFaveGetPagesBuilder(t *testing.T) {
	t.Parallel()

	b := params.NewFaveGetPagesBuilder()

	b.Offset(1)
	b.Count(1)
	b.Type("text")
	b.Fields([]string{"test"})
	b.TagID(1)

	assert.Equal(t, 1, b.Params["offset"])
	assert.Equal(t, 1, b.Params["count"])
	assert.Equal(t, "text", b.Params["type"])
	assert.Equal(t, []string{"test"}, b.Params["fields"])
	assert.Equal(t, 1, b.Params["tag_id"])
}

func TestFaveRemoveArticleBuilder(t *testing.T) {
	t.Parallel()

	b := params.NewFaveRemoveArticleBuilder()

	b.OwnerID(1)
	b.ArticleID(1)

	assert.Equal(t, 1, b.Params["owner_id"])
	assert.Equal(t, 1, b.Params["article_id"])
}

func TestFaveRemoveLinkBuilder(t *testing.T) {
	t.Parallel()

	b := params.NewFaveRemoveLinkBuilder()

	b.LinkID("text")
	b.Link("text")

	assert.Equal(t, "text", b.Params["link_id"])
	assert.Equal(t, "text", b.Params["link"])
}

func TestFaveRemovePageBuilder(t *testing.T) {
	t.Parallel()

	b := params.NewFaveRemovePageBuilder()

	b.UserID(1)
	b.GroupID(1)

	assert.Equal(t, 1, b.Params["user_id"])
	assert.Equal(t, 1, b.Params["group_id"])
}

func TestFaveRemovePostBuilder(t *testing.T) {
	t.Parallel()

	b := params.NewFaveRemovePostBuilder()

	b.OwnerID(1)
	b.ID(1)

	assert.Equal(t, 1, b.Params["owner_id"])
	assert.Equal(t, 1, b.Params["id"])
}

func TestFaveRemoveProductBuilder(t *testing.T) {
	t.Parallel()

	b := params.NewFaveRemoveProductBuilder()

	b.OwnerID(1)
	b.ID(1)

	assert.Equal(t, 1, b.Params["owner_id"])
	assert.Equal(t, 1, b.Params["id"])
}

func TestFaveRemoveTagBuilder(t *testing.T) {
	t.Parallel()

	b := params.NewFaveRemoveTagBuilder()

	b.ID(1)

	assert.Equal(t, 1, b.Params["id"])
}

func TestFaveReorderTagsBuilder(t *testing.T) {
	t.Parallel()

	b := params.NewFaveReorderTagsBuilder()

	b.IDs([]int{1})

	assert.Equal(t, []int{1}, b.Params["ids"])
}

func TestFaveSetPageTagsBuilder(t *testing.T) {
	t.Parallel()

	b := params.NewFaveSetPageTagsBuilder()

	b.UserID(1)
	b.GroupID(1)
	b.TagIDs([]int{1})

	assert.Equal(t, 1, b.Params["user_id"])
	assert.Equal(t, 1, b.Params["group_id"])
	assert.Equal(t, []int{1}, b.Params["tag_ids"])
}

func TestFaveSetTagsBuilder(t *testing.T) {
	t.Parallel()

	b := params.NewFaveSetTagsBuilder()

	b.ItemType("text")
	b.ItemOwnerID(1)
	b.ItemID(1)
	b.TagIDs([]int{1})
	b.LinkID("text")
	b.LinkURL("text")

	assert.Equal(t, "text", b.Params["item_type"])
	assert.Equal(t, 1, b.Params["item_owner_id"])
	assert.Equal(t, 1, b.Params["item_id"])
	assert.Equal(t, []int{1}, b.Params["tag_ids"])
	assert.Equal(t, "text", b.Params["link_id"])
	assert.Equal(t, "text", b.Params["link_url"])
}

func TestFaveTrackPageInteractionBuilder(t *testing.T) {
	t.Parallel()

	b := params.NewFaveTrackPageInteractionBuilder()

	b.UserID(1)
	b.GroupID(1)

	assert.Equal(t, 1, b.Params["user_id"])
	assert.Equal(t, 1, b.Params["group_id"])
}
