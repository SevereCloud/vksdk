package params_test

import (
	"testing"

	"github.com/SevereCloud/vksdk/v2/api/params"
	"github.com/stretchr/testify/assert"
)

func TestFaveAddArticleBuilder(t *testing.T) {
	t.Parallel()

	b := params.NewFaveAddArticleBuilder()

	b.URL("text")

	assert.Equal(t, b.Params["url"], "text")
}

func TestFaveAddLinkBuilder(t *testing.T) {
	t.Parallel()

	b := params.NewFaveAddLinkBuilder()

	b.Link("text")

	assert.Equal(t, b.Params["link"], "text")
}

func TestFaveAddPageBuilder(t *testing.T) {
	t.Parallel()

	b := params.NewFaveAddPageBuilder()

	b.UserID(1)
	b.GroupID(1)

	assert.Equal(t, b.Params["user_id"], 1)
	assert.Equal(t, b.Params["group_id"], 1)
}

func TestFaveAddPostBuilder(t *testing.T) {
	t.Parallel()

	b := params.NewFaveAddPostBuilder()

	b.OwnerID(1)
	b.ID(1)
	b.AccessKey("text")

	assert.Equal(t, b.Params["owner_id"], 1)
	assert.Equal(t, b.Params["id"], 1)
	assert.Equal(t, b.Params["access_key"], "text")
}

func TestFaveAddProductBuilder(t *testing.T) {
	t.Parallel()

	b := params.NewFaveAddProductBuilder()

	b.OwnerID(1)
	b.ID(1)
	b.AccessKey("text")

	assert.Equal(t, b.Params["owner_id"], 1)
	assert.Equal(t, b.Params["id"], 1)
	assert.Equal(t, b.Params["access_key"], "text")
}

func TestFaveAddTagBuilder(t *testing.T) {
	t.Parallel()

	b := params.NewFaveAddTagBuilder()

	b.Name("text")

	assert.Equal(t, b.Params["name"], "text")
}

func TestFaveAddVideoBuilder(t *testing.T) {
	t.Parallel()

	b := params.NewFaveAddVideoBuilder()

	b.OwnerID(1)
	b.ID(1)
	b.AccessKey("text")

	assert.Equal(t, b.Params["owner_id"], 1)
	assert.Equal(t, b.Params["id"], 1)
	assert.Equal(t, b.Params["access_key"], "text")
}

func TestFaveEditTagBuilder(t *testing.T) {
	t.Parallel()

	b := params.NewFaveEditTagBuilder()

	b.ID(1)
	b.Name("text")

	assert.Equal(t, b.Params["id"], 1)
	assert.Equal(t, b.Params["name"], "text")
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

	assert.Equal(t, b.Params["extended"], true)
	assert.Equal(t, b.Params["item_type"], "text")
	assert.Equal(t, b.Params["tag_id"], 1)
	assert.Equal(t, b.Params["offset"], 1)
	assert.Equal(t, b.Params["count"], 1)
	assert.Equal(t, b.Params["fields"], "text")
	assert.Equal(t, b.Params["is_from_snackbar"], true)
}

func TestFaveGetPagesBuilder(t *testing.T) {
	t.Parallel()

	b := params.NewFaveGetPagesBuilder()

	b.Offset(1)
	b.Count(1)
	b.Type("text")
	b.Fields([]string{"test"})
	b.TagID(1)

	assert.Equal(t, b.Params["offset"], 1)
	assert.Equal(t, b.Params["count"], 1)
	assert.Equal(t, b.Params["type"], "text")
	assert.Equal(t, b.Params["fields"], []string{"test"})
	assert.Equal(t, b.Params["tag_id"], 1)
}

func TestFaveRemoveArticleBuilder(t *testing.T) {
	t.Parallel()

	b := params.NewFaveRemoveArticleBuilder()

	b.OwnerID(1)
	b.ArticleID(1)

	assert.Equal(t, b.Params["owner_id"], 1)
	assert.Equal(t, b.Params["article_id"], 1)
}

func TestFaveRemoveLinkBuilder(t *testing.T) {
	t.Parallel()

	b := params.NewFaveRemoveLinkBuilder()

	b.LinkID("text")
	b.Link("text")

	assert.Equal(t, b.Params["link_id"], "text")
	assert.Equal(t, b.Params["link"], "text")
}

func TestFaveRemovePageBuilder(t *testing.T) {
	t.Parallel()

	b := params.NewFaveRemovePageBuilder()

	b.UserID(1)
	b.GroupID(1)

	assert.Equal(t, b.Params["user_id"], 1)
	assert.Equal(t, b.Params["group_id"], 1)
}

func TestFaveRemovePostBuilder(t *testing.T) {
	t.Parallel()

	b := params.NewFaveRemovePostBuilder()

	b.OwnerID(1)
	b.ID(1)

	assert.Equal(t, b.Params["owner_id"], 1)
	assert.Equal(t, b.Params["id"], 1)
}

func TestFaveRemoveProductBuilder(t *testing.T) {
	t.Parallel()

	b := params.NewFaveRemoveProductBuilder()

	b.OwnerID(1)
	b.ID(1)

	assert.Equal(t, b.Params["owner_id"], 1)
	assert.Equal(t, b.Params["id"], 1)
}

func TestFaveRemoveTagBuilder(t *testing.T) {
	t.Parallel()

	b := params.NewFaveRemoveTagBuilder()

	b.ID(1)

	assert.Equal(t, b.Params["id"], 1)
}

func TestFaveReorderTagsBuilder(t *testing.T) {
	t.Parallel()

	b := params.NewFaveReorderTagsBuilder()

	b.IDs([]int{1})

	assert.Equal(t, b.Params["ids"], []int{1})
}

func TestFaveSetPageTagsBuilder(t *testing.T) {
	t.Parallel()

	b := params.NewFaveSetPageTagsBuilder()

	b.UserID(1)
	b.GroupID(1)
	b.TagIDs([]int{1})

	assert.Equal(t, b.Params["user_id"], 1)
	assert.Equal(t, b.Params["group_id"], 1)
	assert.Equal(t, b.Params["tag_ids"], []int{1})
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

	assert.Equal(t, b.Params["item_type"], "text")
	assert.Equal(t, b.Params["item_owner_id"], 1)
	assert.Equal(t, b.Params["item_id"], 1)
	assert.Equal(t, b.Params["tag_ids"], []int{1})
	assert.Equal(t, b.Params["link_id"], "text")
	assert.Equal(t, b.Params["link_url"], "text")
}

func TestFaveTrackPageInteractionBuilder(t *testing.T) {
	t.Parallel()

	b := params.NewFaveTrackPageInteractionBuilder()

	b.UserID(1)
	b.GroupID(1)

	assert.Equal(t, b.Params["user_id"], 1)
	assert.Equal(t, b.Params["group_id"], 1)
}
