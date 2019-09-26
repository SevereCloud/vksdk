package api

import (
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestVK_WallPost(t *testing.T) {
	needUserToken(t)

	post, err := vkUser.WallPost(map[string]string{
		"message": "Test post",
	})
	if !assert.NoError(t, err) {
		t.Fatal(err)
	}
	assert.NotEmpty(t, post)

	res, err := vkUser.WallPin(map[string]string{
		"post_id": strconv.Itoa(post.PostID),
	})
	assert.NoError(t, err)
	assert.NotEmpty(t, res)

	res, err = vkUser.WallUnpin(map[string]string{
		"post_id": strconv.Itoa(post.PostID),
	})
	assert.NoError(t, err)
	assert.NotEmpty(t, res)

	res, err = vkUser.WallCloseComments(map[string]string{
		"post_id":  strconv.Itoa(post.PostID),
		"owner_id": strconv.Itoa(vkUserID),
	})
	assert.NoError(t, err)
	assert.NotEmpty(t, res)

	res, err = vkUser.WallOpenComments(map[string]string{
		"post_id":  strconv.Itoa(post.PostID),
		"owner_id": strconv.Itoa(vkUserID),
	})
	assert.NoError(t, err)
	assert.NotEmpty(t, res)

	edit, err := vkUser.WallEdit(map[string]string{
		"post_id": strconv.Itoa(post.PostID),
		"message": "Test post edited",
	})
	assert.NoError(t, err)
	assert.NotEmpty(t, edit.PostID)

	res, err = vkUser.WallDelete(map[string]string{
		"post_id": strconv.Itoa(post.PostID),
	})
	assert.NoError(t, err)
	assert.NotEmpty(t, res)

	res, err = vkUser.WallRestore(map[string]string{
		"post_id": strconv.Itoa(post.PostID),
		"message": "Test post edited",
	})
	assert.NoError(t, err)
	assert.NotEmpty(t, res)
}

func TestVK_WallGet(t *testing.T) {
	needServiceToken(t)

	params := map[string]string{
		"owner_id": "-86529522",
		"count":    "100",
	}

	get, err := vkService.WallGet(params)
	assert.NoError(t, err)
	assert.NotEmpty(t, get.Count)
	if assert.NotEmpty(t, get.Items) {
		for _, item := range get.Items {
			assert.NotEmpty(t, item.ID)
			assert.NotEmpty(t, item.FromID)
			assert.NotEmpty(t, item.OwnerID)
			assert.NotEmpty(t, item.Date)
			// assert.NotEmpty(t, item.MarkedAsAds)
			assert.NotEmpty(t, item.PostType)
			assert.NotEmpty(t, item.Text)
			assert.NotEmpty(t, item.PostSource)
			assert.NotEmpty(t, item.Comments)
			assert.NotEmpty(t, item.Likes)
			assert.NotEmpty(t, item.Reposts)
			assert.NotEmpty(t, item.Views)
		}
	}

	_, err = vkService.WallGetExtended(params)
	assert.NoError(t, err)
}

func TestVK_WallGetByID(t *testing.T) {
	needServiceToken(t)

	params := map[string]string{
		"posts": "-1_340393",
	}

	res, err := vkService.WallGetByID(params)
	assert.NoError(t, err)
	assert.NotEmpty(t, res[0].ID)
	assert.NotEmpty(t, res[0].FromID)
	assert.NotEmpty(t, res[0].OwnerID)
	assert.NotEmpty(t, res[0].Date)
	// assert.NotEmpty(t, item.MarkedAsAds)
	assert.NotEmpty(t, res[0].PostType)
	// assert.NotEmpty(t, res[0].Text)
	assert.NotEmpty(t, res[0].PostSource)
	assert.NotEmpty(t, res[0].Comments)
	assert.NotEmpty(t, res[0].Likes)
	assert.NotEmpty(t, res[0].Reposts)
	// assert.NotEmpty(t, res[0].Views)

	_, err = vkService.WallGetByIDExtended(params)
	assert.NoError(t, err)
}

func TestVK_WallGetComment(t *testing.T) {
	needUserToken(t)

	params := map[string]string{
		"owner_id":   "66559",
		"comment_id": "73674",
	}

	res, err := vkUser.WallGetComment(params)
	assert.NoError(t, err)
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
	assert.NoError(t, err)
}

func TestVK_WallGetComments(t *testing.T) {
	needServiceToken(t)

	params := map[string]string{
		"owner_id":           "85635407",
		"post_id":            "3199",
		"need_likes":         "1",
		"count":              "100",
		"thread_items_count": "10",
	}

	res, err := vkService.WallGetComments(params)
	assert.NoError(t, err)
	assert.NotEmpty(t, res.Count)
	assert.NotEmpty(t, res.Items)

	resEx, err := vkService.WallGetCommentsExtended(params)
	assert.NoError(t, err)
	assert.NotEmpty(t, resEx.Count)
	assert.NotEmpty(t, resEx.Items)
}

func TestVK_WallGetReposts(t *testing.T) {
	needServiceToken(t)

	_, err := vkService.WallGetReposts(map[string]string{
		"owner_id": "1",
		"post_id":  "2442097",
		"count":    "1000",
	})
	assert.NoError(t, err)
	// assert.NotEmpty(t, res.Items)
}

func TestVK_WallCreateComment(t *testing.T) {
	needUserToken(t)

	comment, err := vkUser.WallCreateComment(map[string]string{
		"owner_id": "117253521",
		"post_id":  "2342",
		"message":  "Test comment",
	})
	if !assert.NoError(t, err) {
		t.Fatal(err)
	}
	assert.NotEmpty(t, comment.CommentID)

	res, err := vkUser.WallEditComment(map[string]string{
		"owner_id":   "117253521",
		"comment_id": strconv.Itoa(comment.CommentID),
		"message":    "Test comment edited",
	})
	assert.NoError(t, err)
	assert.NotEmpty(t, res)

	res, err = vkUser.WallDeleteComment(map[string]string{
		"owner_id":   "117253521",
		"comment_id": strconv.Itoa(comment.CommentID),
	})
	assert.NoError(t, err)
	assert.NotEmpty(t, res)

	res, err = vkUser.WallRestoreComment(map[string]string{
		"owner_id":   "117253521",
		"comment_id": strconv.Itoa(comment.CommentID),
	})
	assert.NoError(t, err)
	assert.NotEmpty(t, res)
}

func TestVK_WallPostAdsStealth(t *testing.T) {
	needUserToken(t)
	needGroupToken(t)

	post, err := vkUser.WallPostAdsStealth(map[string]string{
		"owner_id": strconv.Itoa(-vkGroupID),
		"message":  "Test AdsStealth",
	})
	assert.NoError(t, err)
	assert.NotEmpty(t, post.PostID)

	res, err := vkUser.WallEditAdsStealth(map[string]string{
		"owner_id": strconv.Itoa(-vkGroupID),
		"post_id":  strconv.Itoa(post.PostID),
		"message":  "Test AdsStealth edited",
	})
	assert.NoError(t, err)
	assert.NotEmpty(t, res)
}

func TestVK_WallReportComment(t *testing.T) {
	needUserToken(t)

	res, err := vkUser.WallReportComment(map[string]string{
		"owner_id":   "66748",
		"comment_id": "4136",
		// "reason":   "3",
	})
	assert.NoError(t, err)
	assert.NotEmpty(t, res)
}

func TestVK_WallReportPost(t *testing.T) {
	needUserToken(t)

	res, err := vkUser.WallReportPost(map[string]string{
		"owner_id": "66748",
		"post_id":  "3821",
		"reason":   "3",
	})
	assert.NoError(t, err)
	assert.NotEmpty(t, res)
}

func TestVK_WallRepost(t *testing.T) {
	needUserToken(t)

	res, err := vkUser.WallRepost(map[string]string{
		"object": "wall85635407_3133",
	})
	assert.NoError(t, err)
	assert.NotEmpty(t, res.Success)
	assert.NotEmpty(t, res.RepostsCount)
	assert.NotEmpty(t, res.PostID)
	assert.NotEmpty(t, res.LikesCount)
}

func TestVK_WallSearch(t *testing.T) {
	needServiceToken(t)

	params := map[string]string{
		"owner_id": "6",
		"query":    "vk",
		"count":    "100",
	}

	res, err := vkService.WallSearch(params)
	assert.NoError(t, err)
	assert.NotEmpty(t, res.Count)
	assert.NotEmpty(t, res.Items)

	resEx, err := vkService.WallSearchExtended(params)
	assert.NoError(t, err)
	assert.NotEmpty(t, resEx.Count)
	assert.NotEmpty(t, resEx.Items)
}
