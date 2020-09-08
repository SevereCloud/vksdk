package params_test

import (
	"testing"

	"github.com/SevereCloud/vksdk/v2/api/params"
	"github.com/stretchr/testify/assert"
)

func TestLikesAddBuilder(t *testing.T) {
	t.Parallel()

	b := params.NewLikesAddBuilder()

	b.Type("text")
	b.OwnerID(1)
	b.ItemID(1)
	b.ReactionID(1)
	b.AccessKey("text")

	assert.Equal(t, b.Params["type"], "text")
	assert.Equal(t, b.Params["owner_id"], 1)
	assert.Equal(t, b.Params["item_id"], 1)
	assert.Equal(t, b.Params["reaction_id"], 1)
	assert.Equal(t, b.Params["access_key"], "text")
}

func TestLikesDeleteBuilder(t *testing.T) {
	t.Parallel()

	b := params.NewLikesDeleteBuilder()

	b.Type("text")
	b.OwnerID(1)
	b.ItemID(1)

	assert.Equal(t, b.Params["type"], "text")
	assert.Equal(t, b.Params["owner_id"], 1)
	assert.Equal(t, b.Params["item_id"], 1)
}

func TestLikesGetListBuilder(t *testing.T) {
	t.Parallel()

	b := params.NewLikesGetListBuilder()

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

func TestLikesIsLikedBuilder(t *testing.T) {
	t.Parallel()

	b := params.NewLikesIsLikedBuilder()

	b.UserID(1)
	b.Type("text")
	b.OwnerID(1)
	b.ItemID(1)

	assert.Equal(t, b.Params["user_id"], 1)
	assert.Equal(t, b.Params["type"], "text")
	assert.Equal(t, b.Params["owner_id"], 1)
	assert.Equal(t, b.Params["item_id"], 1)
}
