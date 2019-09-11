package api

import (
	"log"
	"strconv"
	"testing"
)

func TestVK_WallPost(t *testing.T) {
	if vkUser.AccessToken == "" {
		t.Skip("USER_TOKEN empty")
	}

	post, err := vkUser.WallPost(map[string]string{
		"message": "Test post",
	})
	if err != nil {
		log.Fatal(err)
	}

	_, err = vkUser.WallPin(map[string]string{
		"post_id": strconv.Itoa(post.PostID),
	})
	if err != nil {
		t.Errorf("VK.WallPin() err = %v", err)
	}

	_, err = vkUser.WallUnpin(map[string]string{
		"post_id": strconv.Itoa(post.PostID),
	})
	if err != nil {
		t.Errorf("VK.WallUnpin() err = %v", err)
	}

	_, err = vkUser.WallCloseComments(map[string]string{
		"post_id":  strconv.Itoa(post.PostID),
		"owner_id": strconv.Itoa(vkUserID),
	})
	if err != nil {
		t.Errorf("VK.WallCloseComments() err = %v", err)
	}

	_, err = vkUser.WallOpenComments(map[string]string{
		"post_id":  strconv.Itoa(post.PostID),
		"owner_id": strconv.Itoa(vkUserID),
	})
	if err != nil {
		t.Errorf("VK.WallOpenComments() err = %v", err)
	}

	_, err = vkUser.WallEdit(map[string]string{
		"post_id": strconv.Itoa(post.PostID),
		"message": "Test post edited",
	})
	if err != nil {
		t.Errorf("VK.WallEdit() err = %v", err)
	}

	_, err = vkUser.WallDelete(map[string]string{
		"post_id": strconv.Itoa(post.PostID),
	})
	if err != nil {
		t.Errorf("VK.WallDelete() err = %v", err)
	}

	_, err = vkUser.WallRestore(map[string]string{
		"post_id": strconv.Itoa(post.PostID),
		"message": "Test post edited",
	})
	if err != nil {
		t.Errorf("VK.WallRestore() err = %v", err)
	}
}

func TestVK_WallGet(t *testing.T) {
	if vkService.AccessToken == "" {
		t.Skip("SERVICE_TOKEN empty")
	}

	params := map[string]string{
		"owner_id": "-86529522",
		"count":    "100",
	}

	_, err := vkService.WallGet(params)
	if err != nil {
		t.Errorf("VK.WallGet() err = %v", err)
	}

	_, err = vkService.WallGetExtended(params)
	if err != nil {
		t.Errorf("VK.WallGetExtended() err = %v", err)
	}
}

func TestVK_WallGetByID(t *testing.T) {
	if vkService.AccessToken == "" {
		t.Skip("SERVICE_TOKEN empty")
	}

	params := map[string]string{
		"posts": "-1_340393",
	}

	_, err := vkService.WallGetByID(params)
	if err != nil {
		t.Errorf("VK.WallGetByID() err = %v", err)
	}

	_, err = vkService.WallGetByIDExtended(params)
	if err != nil {
		t.Errorf("VK.WallGetByIDExtended() err = %v", err)
	}
}

func TestVK_WallGetComment(t *testing.T) {
	// BUG(VK): https://github.com/SevereCloud/vksdk/issues/55
	t.Skip("BUG")
	if vkUser.AccessToken == "" {
		t.Skip("USER_TOKEN empty")
	}

	params := map[string]string{
		"owner_id":   "66559",
		"comment_id": "73674",
	}

	_, err := vkUser.WallGetComment(params)
	if err != nil {
		t.Errorf("VK.WallGetComment() err = %v", err)
	}

	_, err = vkUser.WallGetCommentExtended(params)
	if err != nil {
		t.Errorf("VK.WallGetCommentExtended() err = %v", err)
	}
}

func TestVK_WallGetComments(t *testing.T) {
	if vkService.AccessToken == "" {
		t.Skip("SERVICE_TOKEN empty")
	}

	params := map[string]string{
		"owner_id":           "85635407",
		"post_id":            "3199",
		"need_likes":         "1",
		"count":              "100",
		"thread_items_count": "10",
	}

	_, err := vkService.WallGetComments(params)
	if err != nil {
		t.Errorf("VK.WallGetComments() err = %v", err)
	}

	_, err = vkService.WallGetCommentsExtended(params)
	if err != nil {
		t.Errorf("VK.WallGetCommentsExtended() err = %v", err)
	}
}

func TestVK_WallGetReposts(t *testing.T) {
	if vkService.AccessToken == "" {
		t.Skip("SERVICE_TOKEN empty")
	}

	_, err := vkService.WallGetReposts(map[string]string{
		"owner_id": "1",
		"post_id":  "2442097",
		"count":    "1000",
	})
	if err != nil {
		t.Errorf("VK.WallGetReposts() err = %v", err)
	}
}

func TestVK_WallCreateComment(t *testing.T) {
	if vkUser.AccessToken == "" {
		t.Skip("USER_TOKEN empty")
	}

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
	if err != nil {
		t.Errorf("VK.WallEditComment() err = %v", err)
	}

	_, err = vkUser.WallDeleteComment(map[string]string{
		"owner_id":   "117253521",
		"comment_id": strconv.Itoa(comment.CommentID),
	})
	if err != nil {
		t.Errorf("VK.WallDeleteComment() err = %v", err)
	}

	_, err = vkUser.WallRestoreComment(map[string]string{
		"owner_id":   "117253521",
		"comment_id": strconv.Itoa(comment.CommentID),
	})
	if err != nil {
		t.Errorf("VK.WallRestoreComment() err = %v", err)
	}
}

func TestVK_WallPostAdsStealth(t *testing.T) {
	if vkUser.AccessToken == "" || vkGroupID == 0 {
		t.Skip("USER_TOKEN empty or GroupID not found")
	}

	post, err := vkUser.WallPostAdsStealth(map[string]string{
		"owner_id": strconv.Itoa(-vkGroupID),
		"message":  "Test AdsStealth",
	})
	if err != nil {
		t.Errorf("VK.WallPostAdsStealth() err = %v", err)
	}

	_, err = vkUser.WallEditAdsStealth(map[string]string{
		"owner_id": strconv.Itoa(-vkGroupID),
		"post_id":  strconv.Itoa(post.PostID),
		"message":  "Test AdsStealth edited",
	})
	if err != nil {
		t.Errorf("VK.WallEditAdsStealth() err = %v", err)
	}
}

func TestVK_WallReportComment(t *testing.T) {
	if vkUser.AccessToken == "" {
		t.Skip("USER_TOKEN empty")
	}

	_, err := vkUser.WallReportComment(map[string]string{
		"owner_id":   "66748",
		"comment_id": "4136",
		// "reason":   "3",
	})
	if err != nil {
		t.Errorf("VK.WallReportComment() err = %v", err)
	}
}

func TestVK_WallReportPost(t *testing.T) {
	if vkUser.AccessToken == "" {
		t.Skip("USER_TOKEN empty")
	}

	_, err := vkUser.WallReportPost(map[string]string{
		"owner_id": "66748",
		"post_id":  "3821",
		"reason":   "3",
	})
	if err != nil {
		t.Errorf("VK.WallReportPost() err = %v", err)
	}
}

func TestVK_WallRepost(t *testing.T) {
	if vkUser.AccessToken == "" {
		t.Skip("USER_TOKEN empty")
	}

	_, err := vkUser.WallRepost(map[string]string{
		"object": "wall85635407_3133",
	})
	if err != nil {
		t.Errorf("VK.WallRepost() err = %v", err)
	}
}

func TestVK_WallSearch(t *testing.T) {
	if vkService.AccessToken == "" {
		t.Skip("SERVICE_TOKEN empty")
	}

	params := map[string]string{
		"owner_id": "6",
		"query":    "vk",
		"count":    "100",
	}

	_, err := vkService.WallSearch(params)
	if err != nil {
		t.Errorf("VK.WallSearch() err = %v", err)
	}

	_, err = vkService.WallSearchExtended(params)
	if err != nil {
		t.Errorf("VK.WallSearchExtended() err = %v", err)
	}
}
