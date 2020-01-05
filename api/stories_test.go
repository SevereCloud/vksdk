package api_test

import (
	"net/http"
	"testing"

	"github.com/SevereCloud/vksdk/api"
	"github.com/stretchr/testify/assert"
)

func TestVK_StoriesBanOwner(t *testing.T) {
	needUserToken(t)

	_, err := vkUser.StoriesBanOwner(api.Params{
		"owners_ids": []int{-1, 1},
	})
	noError(t, err)

	banned, err := vkUser.StoriesGetBanned(api.Params{})
	noError(t, err)
	assert.NotEmpty(t, banned.Count)
	assert.NotEmpty(t, banned.Items)

	bannedEx, err := vkUser.StoriesGetBannedExtended(api.Params{})
	noError(t, err)
	assert.NotEmpty(t, bannedEx.Count)
	assert.NotEmpty(t, bannedEx.Items)
	assert.NotEmpty(t, bannedEx.Profiles)
	assert.NotEmpty(t, bannedEx.Groups)

	_, err = vkUser.StoriesUnbanOwner(api.Params{
		"owners_ids": []int{-1, 1},
	})
	noError(t, err)
}

func TestVK_StoriesGet(t *testing.T) {
	needUserToken(t)

	_, err := vkUser.StoriesGet(api.Params{})
	noError(t, err)
}

func TestVK_StoriesGetExtended(t *testing.T) {
	needUserToken(t)

	_, err := vkUser.StoriesGetExtended(api.Params{})
	noError(t, err)
}

func TestVK_StoriesGetByID(t *testing.T) {
	needGroupToken(t)

	_, err := vkGroup.StoriesGetByID(api.Params{
		"stories": "93388_21539,93388_20904",
	})
	noError(t, err)
}

func TestVK_StoriesGetByIDExtended(t *testing.T) {
	needGroupToken(t)

	_, err := vkGroup.StoriesGetByIDExtended(api.Params{
		"stories": "93388_21539,93388_20904",
	})
	noError(t, err)
}

func TestVK_StoriesGetReplies(t *testing.T) {
	needGroupToken(t)

	response, err := http.Get(photoURL)
	if err != nil {
		t.Fatalf("%v", err)
	}
	defer response.Body.Close()

	res, err := vkGroup.UploadStoriesPhoto(api.Params{
		"add_to_news": true,
	}, response.Body)
	noError(t, err)
	assert.NotEmpty(t, res.Stories)
	// assert.NotEmpty(t, res.Sig)

	params := api.Params{
		"owner_id": res.Stories.OwnerID,
		"story_id": res.Stories.ID,
	}

	_, err = vkGroup.StoriesGetReplies(params)
	noError(t, err)

	_, err = vkGroup.StoriesGetRepliesExtended(params)
	noError(t, err)

	_, err = vkGroup.StoriesGetStats(params)
	noError(t, err)

	_, err = vkGroup.StoriesGetViewers(params)
	noError(t, err)

	_, err = vkGroup.StoriesGetViewersExtended(params)
	noError(t, err)

	_, err = vkGroup.StoriesDelete(params)
	noError(t, err)
}

func TestVK_StoriesHideAllReplies(t *testing.T) {
	needUserToken(t)

	_, err := vkUser.StoriesHideAllReplies(api.Params{
		"owner_id": 1,
	})
	noError(t, err)
}

func TestVK_StoriesHideReply(t *testing.T) {
	// TODO: Add test case
}

func TestVK_StoriesSearch(t *testing.T) {
	needUserToken(t)

	_, err := vkUser.StoriesSearch(api.Params{
		"q": "test",
	})
	noError(t, err)
}

func TestVK_StoriesSearchExtended(t *testing.T) {
	needUserToken(t)

	_, err := vkUser.StoriesSearchExtended(api.Params{
		"q": "test",
	})
	noError(t, err)
}
