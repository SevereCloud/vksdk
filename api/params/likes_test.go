package params_test

import (
	"testing"

	"github.com/SevereCloud/vksdk/v3/api/params"
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

	assert.Equal(t, "text", b.Params["type"])
	assert.Equal(t, 1, b.Params["owner_id"])
	assert.Equal(t, 1, b.Params["item_id"])
	assert.Equal(t, 1, b.Params["reaction_id"])
	assert.Equal(t, "text", b.Params["access_key"])
}

func TestLikesDeleteBuilder(t *testing.T) {
	t.Parallel()

	b := params.NewLikesDeleteBuilder()

	b.Type("text")
	b.OwnerID(1)
	b.ItemID(1)

	assert.Equal(t, "text", b.Params["type"])
	assert.Equal(t, 1, b.Params["owner_id"])
	assert.Equal(t, 1, b.Params["item_id"])
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

	assert.Equal(t, "text", b.Params["type"])
	assert.Equal(t, 1, b.Params["owner_id"])
	assert.Equal(t, 1, b.Params["item_id"])
	assert.Equal(t, "text", b.Params["page_url"])
	assert.Equal(t, "text", b.Params["filter"])
	assert.Equal(t, 1, b.Params["friends_only"])
	assert.Equal(t, true, b.Params["extended"])
	assert.Equal(t, 1, b.Params["offset"])
	assert.Equal(t, 1, b.Params["count"])
	assert.Equal(t, true, b.Params["skip_own"])
}

func TestLikesIsLikedBuilder(t *testing.T) {
	t.Parallel()

	b := params.NewLikesIsLikedBuilder()

	b.UserID(1)
	b.Type("text")
	b.OwnerID(1)
	b.ItemID(1)

	assert.Equal(t, 1, b.Params["user_id"])
	assert.Equal(t, "text", b.Params["type"])
	assert.Equal(t, 1, b.Params["owner_id"])
	assert.Equal(t, 1, b.Params["item_id"])
}
