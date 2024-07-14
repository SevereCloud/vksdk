package api_test

import (
	"testing"

	"github.com/SevereCloud/vksdk/v3/api"

	"github.com/stretchr/testify/assert"
)

func TestVK_WallCheckCopyrightLink(t *testing.T) {
	t.Parallel()

	needUserToken(t)

	get, err := vkUser.WallCheckCopyrightLink(api.Params{
		"link": "https://vk.com/wall1_2442097",
	})
	noError(t, err)
	assert.Equal(t, 1, get)

	_, err = vkUser.WallCheckCopyrightLink(api.Params{
		"link": "https://vk.com/severecloud",
	})
	assert.ErrorIs(t, err, api.ErrWallCheckLinkCantDetermineSource)
}

func TestVK_WallPost(t *testing.T) {
	t.Parallel()

	needUserToken(t)

	post, err := vkUser.WallPost(api.Params{
		"message": "Test post",
	})
	if !noError(t, err) {
		t.Fatal(err)
	}

	assert.NotEmpty(t, post)

	res, err := vkUser.WallPin(api.Params{
		"post_id": post.PostID,
	})
	noError(t, err)
	assert.NotEmpty(t, res)

	res, err = vkUser.WallUnpin(api.Params{
		"post_id": post.PostID,
	})
	noError(t, err)
	assert.NotEmpty(t, res)

	res, err = vkUser.WallCloseComments(api.Params{
		"post_id":  post.PostID,
		"owner_id": vkUserID,
	})
	noError(t, err)
	assert.NotEmpty(t, res)

	res, err = vkUser.WallOpenComments(api.Params{
		"post_id":  post.PostID,
		"owner_id": vkUserID,
	})
	noError(t, err)
	assert.NotEmpty(t, res)

	edit, err := vkUser.WallEdit(api.Params{
		"post_id": post.PostID,
		"message": "Test post edited",
	})
	noError(t, err)
	assert.NotEmpty(t, edit.PostID)

	res, err = vkUser.WallDelete(api.Params{
		"post_id": post.PostID,
	})
	noError(t, err)
	assert.NotEmpty(t, res)

	res, err = vkUser.WallRestore(api.Params{
		"post_id": post.PostID,
		"message": "Test post edited",
	})
	noError(t, err)
	assert.NotEmpty(t, res)
}

func TestVK_WallGet(t *testing.T) {
	t.Parallel()

	needServiceToken(t)

	params := api.Params{
		"owner_id": -22822305,
		"count":    10,
	}

	get, err := vkService.WallGet(params)
	noError(t, err)
	assert.NotEmpty(t, get.Count)

	if assert.NotEmpty(t, get.Items) {
		for _, item := range get.Items {
			assert.NotEmpty(t, item.ID)
			assert.NotEmpty(t, item.FromID)
			assert.NotEmpty(t, item.OwnerID)
			assert.NotEmpty(t, item.Date)
			// assert.NotEmpty(t, item.MarkedAsAds)
			assert.NotEmpty(t, item.PostType)
			// assert.NotEmpty(t, item.Text)
			assert.NotEmpty(t, item.PostSource)
			assert.NotEmpty(t, item.Comments)
			assert.NotEmpty(t, item.Likes)
			assert.NotEmpty(t, item.Reposts)
			assert.NotEmpty(t, item.Views)
		}
	}

	_, err = vkService.WallGetExtended(params)
	noError(t, err)
}

func TestVK_WallGetByID(t *testing.T) {
	t.Parallel()

	needServiceToken(t)

	params := api.Params{
		"posts": "-1_340393",
	}

	res, err := vkService.WallGetByID(params)
	noError(t, err)
	assert.NotEmpty(t, res.Items[0].ID)
	assert.NotEmpty(t, res.Items[0].FromID)
	assert.NotEmpty(t, res.Items[0].OwnerID)
	assert.NotEmpty(t, res.Items[0].Date)
	// assert.NotEmpty(t, item.MarkedAsAds)
	assert.NotEmpty(t, res.Items[0].PostType)
	// assert.NotEmpty(t, res.Items[0].Text)
	assert.NotEmpty(t, res.Items[0].PostSource)
	assert.NotEmpty(t, res.Items[0].Comments)
	assert.NotEmpty(t, res.Items[0].Likes)
	assert.NotEmpty(t, res.Items[0].Reposts)
	// assert.NotEmpty(t, res.Items[0].Views)

	_, err = vkService.WallGetByIDExtended(params)
	noError(t, err)
}

// See https://github.com/SevereCloud/vksdk/issues/147
func TestVK_WallGetByID_issues147(t *testing.T) {
	t.Parallel()

	needServiceToken(t)

	res, err := vkService.WallGetByID(api.Params{
		"posts": "-195263939_4",
	})
	noError(t, err)
	assert.NotEmpty(t, res)

	res, err = vkService.WallGetByID(api.Params{
		"posts": "-169097025_377",
	})
	noError(t, err)
	assert.NotEmpty(t, res.Items[0].Attachments[0].Market.ID)
}

func TestVK_WallGetComment(t *testing.T) {
	t.Parallel()

	needUserToken(t)

	params := api.Params{
		"owner_id":   66559,
		"comment_id": 73674,
	}

	res, err := vkUser.WallGetComment(params)
	noError(t, err)

	if assert.NotEmpty(t, res.Items) {
		assert.NotEmpty(t, res.Items[0].ID)
		assert.NotEmpty(t, res.Items[0].FromID)
		assert.NotEmpty(t, res.Items[0].PostID)
		assert.NotEmpty(t, res.Items[0].OwnerID)
		// assert.NotEmpty(t, res.Items[0].ParentsStack)
		assert.NotEmpty(t, res.Items[0].Date)
		assert.NotEmpty(t, res.Items[0].Text)
		assert.NotEmpty(t, res.Items[0].Likes)
		assert.NotEmpty(t, res.Items[0].ReplyToUser)
		assert.NotEmpty(t, res.Items[0].ReplyToUser)
		assert.NotEmpty(t, res.Items[0].ReplyToComment)
		assert.NotEmpty(t, res.Items[0].Thread)
	}

	_, err = vkUser.WallGetCommentExtended(params)
	noError(t, err)
}

func TestVK_WallGetComments(t *testing.T) {
	t.Parallel()

	needServiceToken(t)

	params := api.Params{
		"owner_id":           85635407,
		"post_id":            3199,
		"need_likes":         1,
		"count":              100,
		"thread_items_count": 10,
	}

	res, err := vkService.WallGetComments(params)
	noError(t, err)
	assert.NotEmpty(t, res.Count)
	assert.NotEmpty(t, res.Items)

	resEx, err := vkService.WallGetCommentsExtended(params)
	noError(t, err)
	assert.NotEmpty(t, resEx.Count)
	assert.NotEmpty(t, resEx.Items)
}

func TestVK_WallGetReposts(t *testing.T) {
	t.Parallel()

	needServiceToken(t)

	_, err := vkService.WallGetReposts(api.Params{
		"owner_id": 1,
		"post_id":  2442097,
		"count":    1000,
	})
	// assert.NotEmpty(t, res.Items)
	noError(t, err)
}

func TestVK_WallCreateComment(t *testing.T) {
	t.Parallel()

	needUserToken(t)

	comment, err := vkUser.WallCreateComment(api.Params{
		"owner_id": 117253521,
		"post_id":  2342,
		"message":  "Test comment",
	})
	if !noError(t, err) {
		t.Fatal(err)
	}

	assert.NotEmpty(t, comment.CommentID)

	res, err := vkUser.WallEditComment(api.Params{
		"owner_id":   117253521,
		"comment_id": comment.CommentID,
		"message":    "Test comment edited",
	})
	noError(t, err)
	assert.NotEmpty(t, res)

	res, err = vkUser.WallDeleteComment(api.Params{
		"owner_id":   117253521,
		"comment_id": comment.CommentID,
	})
	noError(t, err)
	assert.NotEmpty(t, res)

	res, err = vkUser.WallRestoreComment(api.Params{
		"owner_id":   117253521,
		"comment_id": comment.CommentID,
	})
	noError(t, err)
	assert.NotEmpty(t, res)
}

func TestVK_WallPostAdsStealth(t *testing.T) {
	t.Parallel()

	needUserToken(t)
	needGroupToken(t)

	post, err := vkUser.WallPostAdsStealth(api.Params{
		"owner_id": -vkGroupID,
		"message":  "Test AdsStealth",
	})
	noError(t, err)
	assert.NotEmpty(t, post.PostID)

	res, err := vkUser.WallEditAdsStealth(api.Params{
		"owner_id": -vkGroupID,
		"post_id":  post.PostID,
		"message":  "Test AdsStealth edited",
	})
	noError(t, err)
	assert.NotEmpty(t, res)
}

func TestVK_WallReportComment(t *testing.T) {
	t.Parallel()

	needUserToken(t)

	res, err := vkUser.WallReportComment(api.Params{
		"owner_id":   66748,
		"comment_id": 4136,
		"reason":     3,
	})
	noError(t, err)
	assert.NotEmpty(t, res)
}

func TestVK_WallReportPost(t *testing.T) {
	t.Parallel()

	needUserToken(t)

	res, err := vkUser.WallReportPost(api.Params{
		"owner_id": 66748,
		"post_id":  3821,
		"reason":   3,
	})
	noError(t, err)
	assert.NotEmpty(t, res)
}

func TestVK_WallRepost(t *testing.T) {
	t.Parallel()

	needUserToken(t)

	res, err := vkUser.WallRepost(api.Params{
		"object": "wall85635407_3133",
	})
	noError(t, err)
	assert.NotEmpty(t, res.Success)
	assert.NotEmpty(t, res.RepostsCount)
	assert.NotEmpty(t, res.PostID)
	assert.NotEmpty(t, res.LikesCount)
}

func TestVK_WallSearch(t *testing.T) {
	t.Parallel()

	needServiceToken(t)

	params := api.Params{
		"owner_id": 6,
		"query":    "vk",
		"count":    100,
	}

	res, err := vkService.WallSearch(params)
	noError(t, err)
	assert.NotEmpty(t, res.Count)
	assert.NotEmpty(t, res.Items)

	resEx, err := vkService.WallSearchExtended(params)
	noError(t, err)
	assert.NotEmpty(t, resEx.Count)
	assert.NotEmpty(t, resEx.Items)
}
