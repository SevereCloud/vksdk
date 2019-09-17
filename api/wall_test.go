package api

import (
	"log"
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestVK_WallPost(t *testing.T) {
	needUserToken(t)

	post, err := vkUser.WallPost(map[string]string{
		"message": "Test post",
	})
	if err != nil {
		log.Fatal(err)
	}

	_, err = vkUser.WallPin(map[string]string{
		"post_id": strconv.Itoa(post.PostID),
	})
	assert.NoError(t, err)

	_, err = vkUser.WallUnpin(map[string]string{
		"post_id": strconv.Itoa(post.PostID),
	})
	assert.NoError(t, err)

	_, err = vkUser.WallCloseComments(map[string]string{
		"post_id":  strconv.Itoa(post.PostID),
		"owner_id": strconv.Itoa(vkUserID),
	})
	assert.NoError(t, err)

	_, err = vkUser.WallOpenComments(map[string]string{
		"post_id":  strconv.Itoa(post.PostID),
		"owner_id": strconv.Itoa(vkUserID),
	})
	assert.NoError(t, err)

	_, err = vkUser.WallEdit(map[string]string{
		"post_id": strconv.Itoa(post.PostID),
		"message": "Test post edited",
	})
	assert.NoError(t, err)

	_, err = vkUser.WallDelete(map[string]string{
		"post_id": strconv.Itoa(post.PostID),
	})
	assert.NoError(t, err)

	_, err = vkUser.WallRestore(map[string]string{
		"post_id": strconv.Itoa(post.PostID),
		"message": "Test post edited",
	})
	assert.NoError(t, err)
}

func TestVK_WallGet(t *testing.T) {
	needServiceToken(t)

	params := map[string]string{
		"owner_id": "-86529522",
		"count":    "100",
	}

	_, err := vkService.WallGet(params)
	assert.NoError(t, err)

	_, err = vkService.WallGetExtended(params)
	assert.NoError(t, err)
}

func TestVK_WallGetByID(t *testing.T) {
	needServiceToken(t)

	params := map[string]string{
		"posts": "-1_340393",
	}

	_, err := vkService.WallGetByID(params)
	assert.NoError(t, err)

	_, err = vkService.WallGetByIDExtended(params)
	assert.NoError(t, err)
}

func TestVK_WallGetComment(t *testing.T) {
	// BUG(VK): https://github.com/SevereCloud/vksdk/issues/55
	t.Skip("BUG")
	needUserToken(t)

	params := map[string]string{
		"owner_id":   "66559",
		"comment_id": "73674",
	}

	_, err := vkUser.WallGetComment(params)
	assert.NoError(t, err)

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

	_, err := vkService.WallGetComments(params)
	assert.NoError(t, err)

	_, err = vkService.WallGetCommentsExtended(params)
	assert.NoError(t, err)
}

func TestVK_WallGetReposts(t *testing.T) {
	needServiceToken(t)

	_, err := vkService.WallGetReposts(map[string]string{
		"owner_id": "1",
		"post_id":  "2442097",
		"count":    "1000",
	})
	assert.NoError(t, err)
}

func TestVK_WallCreateComment(t *testing.T) {
	needUserToken(t)

	comment, err := vkUser.WallCreateComment(map[string]string{
		"owner_id": "117253521",
		"post_id":  "2342",
		"message":  "Test comment",
	})
	if err != nil {
		log.Fatal(err)
	}

	_, err = vkUser.WallEditComment(map[string]string{
		"owner_id":   "117253521",
		"comment_id": strconv.Itoa(comment.CommentID),
		"message":    "Test comment edited",
	})
	assert.NoError(t, err)

	_, err = vkUser.WallDeleteComment(map[string]string{
		"owner_id":   "117253521",
		"comment_id": strconv.Itoa(comment.CommentID),
	})
	assert.NoError(t, err)

	_, err = vkUser.WallRestoreComment(map[string]string{
		"owner_id":   "117253521",
		"comment_id": strconv.Itoa(comment.CommentID),
	})
	assert.NoError(t, err)
}

func TestVK_WallPostAdsStealth(t *testing.T) {
	needUserToken(t)
	needGroupToken(t)

	post, err := vkUser.WallPostAdsStealth(map[string]string{
		"owner_id": strconv.Itoa(-vkGroupID),
		"message":  "Test AdsStealth",
	})
	assert.NoError(t, err)

	_, err = vkUser.WallEditAdsStealth(map[string]string{
		"owner_id": strconv.Itoa(-vkGroupID),
		"post_id":  strconv.Itoa(post.PostID),
		"message":  "Test AdsStealth edited",
	})
	assert.NoError(t, err)
}

func TestVK_WallReportComment(t *testing.T) {
	needUserToken(t)

	_, err := vkUser.WallReportComment(map[string]string{
		"owner_id":   "66748",
		"comment_id": "4136",
		// "reason":   "3",
	})
	assert.NoError(t, err)
}

func TestVK_WallReportPost(t *testing.T) {
	needUserToken(t)

	_, err := vkUser.WallReportPost(map[string]string{
		"owner_id": "66748",
		"post_id":  "3821",
		"reason":   "3",
	})
	assert.NoError(t, err)
}

func TestVK_WallRepost(t *testing.T) {
	needUserToken(t)

	_, err := vkUser.WallRepost(map[string]string{
		"object": "wall85635407_3133",
	})
	assert.NoError(t, err)
}

func TestVK_WallSearch(t *testing.T) {
	needServiceToken(t)

	params := map[string]string{
		"owner_id": "6",
		"query":    "vk",
		"count":    "100",
	}

	_, err := vkService.WallSearch(params)
	assert.NoError(t, err)

	_, err = vkService.WallSearchExtended(params)
	assert.NoError(t, err)
}
