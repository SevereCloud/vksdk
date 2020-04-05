package api_test

import (
	"testing"

	"github.com/SevereCloud/vksdk/api"
	"github.com/stretchr/testify/assert"
)

func TestVK_NewsfeedAddBan(t *testing.T) {
	t.Parallel()

	needUserToken(t)

	_, err := vkUser.NewsfeedAddBan(api.Params{
		"user_ids":  []int{1, 6492},
		"group_ids": []int{44278061, 1},
	})
	noError(t, err)

	banned, err := vkUser.NewsfeedGetBanned(api.Params{})
	noError(t, err)
	assert.NotEmpty(t, banned.Groups)
	assert.NotEmpty(t, banned.Members)

	bannedEx, err := vkUser.NewsfeedGetBannedExtended(api.Params{})
	noError(t, err)
	assert.NotEmpty(t, bannedEx.Groups)
	assert.NotEmpty(t, bannedEx.Profiles)

	_, err = vkUser.NewsfeedDeleteBan(api.Params{
		"user_ids":  []int{1, 6492},
		"group_ids": []int{44278061, 1},
	})
	noError(t, err)
}

func TestVK_NewsfeedGet(t *testing.T) {
	t.Parallel()

	needUserToken(t)

	res, err := vkUser.NewsfeedGet(api.Params{})
	noError(t, err)
	assert.NotEmpty(t, res.Items)
	assert.NotEmpty(t, res.Profiles)
	assert.NotEmpty(t, res.Groups)
	assert.NotEmpty(t, res.NextFrom)
}

func TestVK_NewsfeedGetComments(t *testing.T) {
	t.Parallel()

	needUserToken(t)

	res, err := vkUser.NewsfeedGetComments(api.Params{
		"last_comments_count": 0,
	})
	noError(t, err)
	// assert.NotEmpty(t, res.NextFrom)
	assert.NotEmpty(t, res.Items)
	assert.NotEmpty(t, res.Profiles)
	assert.NotEmpty(t, res.Groups)
}

func TestVK_NewsfeedGetMentions(t *testing.T) {
	t.Parallel()

	needUserToken(t)

	res, err := vkUser.NewsfeedGetMentions(api.Params{
		"owner_id": 66748,
	})
	noError(t, err)
	assert.NotEmpty(t, res.Count)
	assert.NotEmpty(t, res.Items)
}

func TestVK_NewsfeedGetRecommended(t *testing.T) {
	t.Parallel()

	needUserToken(t)

	res, err := vkUser.NewsfeedGetRecommended(api.Params{})
	noError(t, err)
	assert.NotEmpty(t, res.Items)
	assert.NotEmpty(t, res.Profiles)
	assert.NotEmpty(t, res.Groups)
	assert.NotEmpty(t, res.NextFrom)
}

func TestVK_NewsfeedGetSuggestedSources(t *testing.T) {
	// FIXME: NewsfeedGetSuggestedSources have bug
	// needUserToken(t)

	// res, err := vkUser.NewsfeedGetSuggestedSources(api.Params{})
	// noError(t, err)
	// assert.NotEmpty(t, res.Count)
	// assert.NotEmpty(t, res.Items)
}

func TestVK_NewsfeedIgnoreItem(t *testing.T) {
	t.Parallel()

	needUserToken(t)

	params := api.Params{
		"type":     "wall",
		"owner_id": 1,
		"item_id":  2442097,
	}
	_, err := vkUser.NewsfeedIgnoreItem(params)
	noError(t, err)

	_, err = vkUser.NewsfeedUnignoreItem(params)
	noError(t, err)
}

func TestVK_NewsfeedSaveList(t *testing.T) {
	t.Parallel()

	needUserToken(t)

	listID, err := vkUser.NewsfeedSaveList(api.Params{
		"title":      "Список",
		"source_ids": 1,
		"no_reposts": true,
	})
	noError(t, err)

	lists, err := vkUser.NewsfeedGetLists(api.Params{
		"list_ids": listID,
		"extended": true,
	})
	noError(t, err)
	assert.NotEmpty(t, lists.Count)

	if assert.NotEmpty(t, lists.Items) {
		assert.NotEmpty(t, lists.Items[0].ID)
		assert.NotEmpty(t, lists.Items[0].Title)
		assert.NotEmpty(t, lists.Items[0].SourceIDS)
	}

	_, err = vkUser.NewsfeedDeleteList(api.Params{
		"list_id": listID,
	})
	noError(t, err)
}

func TestVK_NewsfeedSearch(t *testing.T) {
	t.Parallel()

	needServiceToken(t)

	res, err := vkService.NewsfeedSearch(api.Params{
		"q": "VK API",
	})
	noError(t, err)
	assert.NotEmpty(t, res.Items)
	assert.NotEmpty(t, res.Count)
	assert.NotEmpty(t, res.TotalCount)
	assert.NotEmpty(t, res.NextFrom)
}

func TestVK_NewsfeedSearchExtended(t *testing.T) {
	t.Parallel()

	needServiceToken(t)

	res, err := vkService.NewsfeedSearchExtended(api.Params{
		"q": "VK API",
	})
	noError(t, err)
	assert.NotEmpty(t, res.Items)
	assert.NotEmpty(t, res.Count)
	assert.NotEmpty(t, res.TotalCount)
	assert.NotEmpty(t, res.NextFrom)
	assert.NotEmpty(t, res.Profiles)
	assert.NotEmpty(t, res.Groups)
}

func TestVK_NewsfeedUnsubscribe(t *testing.T) {
	t.Parallel()

	needUserToken(t)

	_, err := vkUser.NewsfeedIgnoreItem(api.Params{
		"type":     "wall",
		"owner_id": 1,
		"item_id":  2442097,
	})
	noError(t, err)
}
