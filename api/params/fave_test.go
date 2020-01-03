package params_test

import (
	"testing"

	"github.com/SevereCloud/vksdk/api/params"
	"github.com/stretchr/testify/assert"
)

func TestFaveAddArticleBulder(t *testing.T) {
	b := params.NewFaveAddArticleBulder()

	b.URL("text")

	assert.Equal(t, b.Params["url"], "text")
}

func TestFaveAddLinkBulder(t *testing.T) {
	b := params.NewFaveAddLinkBulder()

	b.Link("text")

	assert.Equal(t, b.Params["link"], "text")
}

func TestFaveAddPageBulder(t *testing.T) {
	b := params.NewFaveAddPageBulder()

	b.UserID(1)
	b.GroupID(1)

	assert.Equal(t, b.Params["user_id"], 1)
	assert.Equal(t, b.Params["group_id"], 1)
}

func TestFaveAddPostBulder(t *testing.T) {
	b := params.NewFaveAddPostBulder()

	b.OwnerID(1)
	b.ID(1)
	b.AccessKey("text")

	assert.Equal(t, b.Params["owner_id"], 1)
	assert.Equal(t, b.Params["id"], 1)
	assert.Equal(t, b.Params["access_key"], "text")
}

func TestFaveAddProductBulder(t *testing.T) {
	b := params.NewFaveAddProductBulder()

	b.OwnerID(1)
	b.ID(1)
	b.AccessKey("text")

	assert.Equal(t, b.Params["owner_id"], 1)
	assert.Equal(t, b.Params["id"], 1)
	assert.Equal(t, b.Params["access_key"], "text")
}

func TestFaveAddTagBulder(t *testing.T) {
	b := params.NewFaveAddTagBulder()

	b.Name("text")

	assert.Equal(t, b.Params["name"], "text")
}

func TestFaveAddVideoBulder(t *testing.T) {
	b := params.NewFaveAddVideoBulder()

	b.OwnerID(1)
	b.ID(1)
	b.AccessKey("text")

	assert.Equal(t, b.Params["owner_id"], 1)
	assert.Equal(t, b.Params["id"], 1)
	assert.Equal(t, b.Params["access_key"], "text")
}

func TestFaveEditTagBulder(t *testing.T) {
	b := params.NewFaveEditTagBulder()

	b.ID(1)
	b.Name("text")

	assert.Equal(t, b.Params["id"], 1)
	assert.Equal(t, b.Params["name"], "text")
}

func TestFaveGetBulder(t *testing.T) {
	b := params.NewFaveGetBulder()

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

func TestFaveGetPagesBulder(t *testing.T) {
	b := params.NewFaveGetPagesBulder()

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

func TestFaveRemoveArticleBulder(t *testing.T) {
	b := params.NewFaveRemoveArticleBulder()

	b.OwnerID(1)
	b.ArticleID(1)

	assert.Equal(t, b.Params["owner_id"], 1)
	assert.Equal(t, b.Params["article_id"], 1)
}

func TestFaveRemoveLinkBulder(t *testing.T) {
	b := params.NewFaveRemoveLinkBulder()

	b.LinkID("text")
	b.Link("text")

	assert.Equal(t, b.Params["link_id"], "text")
	assert.Equal(t, b.Params["link"], "text")
}

func TestFaveRemovePageBulder(t *testing.T) {
	b := params.NewFaveRemovePageBulder()

	b.UserID(1)
	b.GroupID(1)

	assert.Equal(t, b.Params["user_id"], 1)
	assert.Equal(t, b.Params["group_id"], 1)
}

func TestFaveRemovePostBulder(t *testing.T) {
	b := params.NewFaveRemovePostBulder()

	b.OwnerID(1)
	b.ID(1)

	assert.Equal(t, b.Params["owner_id"], 1)
	assert.Equal(t, b.Params["id"], 1)
}

func TestFaveRemoveProductBulder(t *testing.T) {
	b := params.NewFaveRemoveProductBulder()

	b.OwnerID(1)
	b.ID(1)

	assert.Equal(t, b.Params["owner_id"], 1)
	assert.Equal(t, b.Params["id"], 1)
}

func TestFaveRemoveTagBulder(t *testing.T) {
	b := params.NewFaveRemoveTagBulder()

	b.ID(1)

	assert.Equal(t, b.Params["id"], 1)
}

func TestFaveReorderTagsBulder(t *testing.T) {
	b := params.NewFaveReorderTagsBulder()

	b.IDs([]int{1})

	assert.Equal(t, b.Params["ids"], []int{1})
}

func TestFaveSetPageTagsBulder(t *testing.T) {
	b := params.NewFaveSetPageTagsBulder()

	b.UserID(1)
	b.GroupID(1)
	b.TagIDs([]int{1})

	assert.Equal(t, b.Params["user_id"], 1)
	assert.Equal(t, b.Params["group_id"], 1)
	assert.Equal(t, b.Params["tag_ids"], []int{1})
}

func TestFaveSetTagsBulder(t *testing.T) {
	b := params.NewFaveSetTagsBulder()

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

func TestFaveTrackPageInteractionBulder(t *testing.T) {
	b := params.NewFaveTrackPageInteractionBulder()

	b.UserID(1)
	b.GroupID(1)

	assert.Equal(t, b.Params["user_id"], 1)
	assert.Equal(t, b.Params["group_id"], 1)
}
