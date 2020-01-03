package params_test

import (
	"testing"

	"github.com/SevereCloud/vksdk/api/params"
	"github.com/stretchr/testify/assert"
)

func TestLikesAddBulder(t *testing.T) {
	b := params.NewLikesAddBulder()

	b.Type("text")
	b.OwnerID(1)
	b.ItemID(1)
	b.AccessKey("text")

	assert.Equal(t, b.Params["type"], "text")
	assert.Equal(t, b.Params["owner_id"], 1)
	assert.Equal(t, b.Params["item_id"], 1)
	assert.Equal(t, b.Params["access_key"], "text")
}

func TestLikesDeleteBulder(t *testing.T) {
	b := params.NewLikesDeleteBulder()

	b.Type("text")
	b.OwnerID(1)
	b.ItemID(1)

	assert.Equal(t, b.Params["type"], "text")
	assert.Equal(t, b.Params["owner_id"], 1)
	assert.Equal(t, b.Params["item_id"], 1)
}

func TestLikesGetListBulder(t *testing.T) {
	b := params.NewLikesGetListBulder()

	b.Type("text")
	b.OwnerID(1)
	b.ItemID(1)
	b.PageURL("text")
	b.Filter("text")
	b.FriendsOnly(1)
	b.Extended(true)
	b.Offset(1)
	b.Count(1)
	b.SkipOwn(true)

	assert.Equal(t, b.Params["type"], "text")
	assert.Equal(t, b.Params["owner_id"], 1)
	assert.Equal(t, b.Params["item_id"], 1)
	assert.Equal(t, b.Params["page_url"], "text")
	assert.Equal(t, b.Params["filter"], "text")
	assert.Equal(t, b.Params["friends_only"], 1)
	assert.Equal(t, b.Params["extended"], true)
	assert.Equal(t, b.Params["offset"], 1)
	assert.Equal(t, b.Params["count"], 1)
	assert.Equal(t, b.Params["skip_own"], true)
}

func TestLikesIsLikedBulder(t *testing.T) {
	b := params.NewLikesIsLikedBulder()

	b.UserID(1)
	b.Type("text")
	b.OwnerID(1)
	b.ItemID(1)

	assert.Equal(t, b.Params["user_id"], 1)
	assert.Equal(t, b.Params["type"], "text")
	assert.Equal(t, b.Params["owner_id"], 1)
	assert.Equal(t, b.Params["item_id"], 1)
}
