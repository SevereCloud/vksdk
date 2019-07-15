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

	t.Run("WallPost", func(t *testing.T) {
		_, gotVkErr := vkUser.WallPost(map[string]string{
			"message": "Test post 0",
		})
		if gotVkErr.Code != 0 {
			t.Errorf("VK.WallPost() gotVkErr = %v, want %v", gotVkErr, 0)
		}
	})
	post, vkErr := vkUser.WallPost(map[string]string{
		"message": "Test post",
	})
	if vkErr.Code != 0 {
		log.Fatal(vkErr)
	}

	t.Run("WallPin", func(t *testing.T) {
		resp, gotVkErr := vkUser.WallPin(map[string]string{
			"post_id": strconv.Itoa(post.PostID),
		})
		if resp != 1 {
			t.Errorf("VK.WallPin() resp = %v, want %v", resp, 1)
		}
		if gotVkErr.Code != 0 {
			t.Errorf("VK.WallPin() gotVkErr = %v, want %v", gotVkErr, 0)
		}
	})

	t.Run("WallUnpin", func(t *testing.T) {
		resp, gotVkErr := vkUser.WallUnpin(map[string]string{
			"post_id": strconv.Itoa(post.PostID),
		})
		if resp != 1 {
			t.Errorf("VK.WallUnpin() resp = %v, want %v", resp, 1)
		}
		if gotVkErr.Code != 0 {
			t.Errorf("VK.WallUnpin() gotVkErr = %v, want %v", gotVkErr, 0)
		}
	})

	t.Run("WallCloseComments", func(t *testing.T) {
		resp, gotVkErr := vkUser.WallCloseComments(map[string]string{
			"post_id":  strconv.Itoa(post.PostID),
			"owner_id": strconv.Itoa(vkUserID),
		})
		if resp != 1 {
			t.Errorf("VK.WallCloseComments() resp = %v, want %v", resp, 1)
		}
		if gotVkErr.Code != 0 {
			t.Errorf("VK.WallCloseComments() gotVkErr = %v, want %v", gotVkErr, 0)
		}
	})

	t.Run("WallOpenComments", func(t *testing.T) {
		resp, gotVkErr := vkUser.WallOpenComments(map[string]string{
			"post_id":  strconv.Itoa(post.PostID),
			"owner_id": strconv.Itoa(vkUserID),
		})
		if resp != 1 {
			t.Errorf("VK.WallOpenComments() resp = %v, want %v", resp, 1)
		}
		if gotVkErr.Code != 0 {
			t.Errorf("VK.WallOpenComments() gotVkErr = %v, want %v", gotVkErr, 0)
		}
	})

	t.Run("WallEdit", func(t *testing.T) {
		resp, gotVkErr := vkUser.WallEdit(map[string]string{
			"post_id": strconv.Itoa(post.PostID),
			"message": "Test post edited",
		})
		if resp != 1 {
			t.Errorf("VK.WallEdit() resp = %v, want %v", resp, 1)
		}
		if gotVkErr.Code != 0 {
			t.Errorf("VK.WallEdit() gotVkErr = %v, want %v", gotVkErr, 0)
		}
	})

	t.Run("WallDelete", func(t *testing.T) {
		resp, gotVkErr := vkUser.WallDelete(map[string]string{
			"post_id": strconv.Itoa(post.PostID),
		})
		if resp != 1 {
			t.Errorf("VK.WallDelete() resp = %v, want %v", resp, 1)
		}
		if gotVkErr.Code != 0 {
			t.Errorf("VK.WallDelete() gotVkErr = %v, want %v", gotVkErr, 0)
		}
	})

	t.Run("WallRestore", func(t *testing.T) {
		resp, gotVkErr := vkUser.WallRestore(map[string]string{
			"post_id": strconv.Itoa(post.PostID),
			"message": "Test post edited",
		})
		if resp != 1 {
			t.Errorf("VK.WallRestore() resp = %v, want %v", resp, 1)
		}
		if gotVkErr.Code != 0 {
			t.Errorf("VK.WallRestore() gotVkErr = %v, want %v", gotVkErr, 0)
		}
	})
}

func TestVK_WallGet(t *testing.T) {
	if vkService.AccessToken == "" {
		t.Skip("SERVICE_TOKEN empty")
	}

	params := map[string]string{
		"owner_id": "-86529522",
		"count":    "100",
	}

	t.Run("WallGet", func(t *testing.T) {
		_, gotVkErr := vkService.WallGet(params)
		if gotVkErr.Code != 0 {
			t.Errorf("VK.WallGet() gotVkErr = %v, want %v", gotVkErr, 0)
		}
	})
	t.Run("WallGetExtended", func(t *testing.T) {
		_, gotVkErr := vkService.WallGetExtended(params)
		if gotVkErr.Code != 0 {
			t.Errorf("VK.WallGetExtended() gotVkErr = %v, want %v", gotVkErr, 0)
		}
	})
}

func TestVK_WallGetByID(t *testing.T) {
	if vkService.AccessToken == "" {
		t.Skip("SERVICE_TOKEN empty")
	}

	params := map[string]string{
		"posts": "-1_340393",
	}

	t.Run("WallGetByID", func(t *testing.T) {
		_, gotVkErr := vkService.WallGetByID(params)
		if gotVkErr.Code != 0 {
			t.Errorf("VK.WallGetByID() gotVkErr = %v, want %v", gotVkErr, 0)
		}
	})
	t.Run("WallGetByIDExtended", func(t *testing.T) {
		_, gotVkErr := vkService.WallGetByIDExtended(params)
		if gotVkErr.Code != 0 {
			t.Errorf("VK.WallGetByIDExtended() gotVkErr = %v, want %v", gotVkErr, 0)
		}
	})
}

func TestVK_WallGetComment(t *testing.T) {
	if vkUser.AccessToken == "" {
		t.Skip("USER_TOKEN empty")
	}

	params := map[string]string{
		"owner_id":   "66559",
		"comment_id": "73674",
	}

	t.Run("WallGetComment", func(t *testing.T) {
		_, gotVkErr := vkUser.WallGetComment(params)
		if gotVkErr.Code != 0 {
			t.Errorf("VK.WallGetComment() gotVkErr = %v, want %v", gotVkErr, 0)
		}
	})
	t.Run("WallGetCommentExtended", func(t *testing.T) {
		_, gotVkErr := vkUser.WallGetCommentExtended(params)
		if gotVkErr.Code != 0 {
			t.Errorf("VK.WallGetCommentExtended() gotVkErr = %v, want %v", gotVkErr, 0)
		}
	})
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

	t.Run("WallGetComments", func(t *testing.T) {
		_, gotVkErr := vkService.WallGetComments(params)
		if gotVkErr.Code != 0 {
			t.Errorf("VK.WallGetComments() gotVkErr = %v, want %v", gotVkErr, 0)
		}
	})
	t.Run("WallGetCommentsExtended", func(t *testing.T) {
		_, gotVkErr := vkService.WallGetCommentsExtended(params)
		if gotVkErr.Code != 0 {
			t.Errorf("VK.WallGetCommentsExtended() gotVkErr = %v, want %v", gotVkErr, 0)
		}
	})
}

func TestVK_WallGetReposts(t *testing.T) {
	if vkService.AccessToken == "" {
		t.Skip("SERVICE_TOKEN empty")
	}

	t.Run("WallGetReposts", func(t *testing.T) {
		_, gotVkErr := vkService.WallGetReposts(map[string]string{
			"owner_id": "1",
			"post_id":  "2442097",
			"count":    "1000",
		})
		if gotVkErr.Code != 0 {
			t.Errorf("VK.WallGetReposts() gotVkErr = %v, want %v", gotVkErr, 0)
		}
	})
}

func TestVK_WallCreateComment(t *testing.T) {
	if vkUser.AccessToken == "" {
		t.Skip("USER_TOKEN empty")
	}

	t.Run("WallCreateComment", func(t *testing.T) {
		_, gotVkErr := vkUser.WallCreateComment(map[string]string{
			"owner_id": "540036751",
			"post_id":  "29",
			"message":  "test comment 0",
		})
		if gotVkErr.Code != 0 {
			t.Errorf("VK.WallCreateComment() gotVkErr = %v, want %v", gotVkErr, 0)
		}
	})
	comment, vkErr := vkUser.WallCreateComment(map[string]string{
		"owner_id": "540036751",
		"post_id":  "29",
		"message":  "Test comment",
	})
	if vkErr.Code != 0 {
		log.Fatal(vkErr)
	}

	t.Run("WallEditComment", func(t *testing.T) {
		resp, gotVkErr := vkUser.WallEditComment(map[string]string{
			"owner_id":   "540036751",
			"comment_id": strconv.Itoa(comment.CommentID),
			"message":    "Test comment edited",
		})
		if resp != 1 {
			t.Errorf("VK.WallEditComment() resp = %v, want %v", resp, 1)
		}
		if gotVkErr.Code != 0 {
			t.Errorf("VK.WallEditComment() gotVkErr = %v, want %v", gotVkErr, 0)
		}
	})

	t.Run("WallDeleteComment", func(t *testing.T) {
		resp, gotVkErr := vkUser.WallDeleteComment(map[string]string{
			"owner_id":   "540036751",
			"comment_id": strconv.Itoa(comment.CommentID),
		})
		if resp != 1 {
			t.Errorf("VK.WallDeleteComment() resp = %v, want %v", resp, 1)
		}
		if gotVkErr.Code != 0 {
			t.Errorf("VK.WallDeleteComment() gotVkErr = %v, want %v", gotVkErr, 0)
		}
	})

	t.Run("WallRestoreComment", func(t *testing.T) {
		resp, gotVkErr := vkUser.WallRestoreComment(map[string]string{
			"owner_id":   "540036751",
			"comment_id": strconv.Itoa(comment.CommentID),
		})
		if resp != 1 {
			t.Errorf("VK.WallRestoreComment() resp = %v, want %v", resp, 1)
		}
		if gotVkErr.Code != 0 {
			t.Errorf("VK.WallRestoreComment() gotVkErr = %v, want %v", gotVkErr, 0)
		}
	})
}

func TestVK_WallPostAdsStealth(t *testing.T) {
	if vkUser.AccessToken == "" || vkGroupID == 0 {
		t.Skip("USER_TOKEN empty or GroupID not found")
	}

	t.Run("Wall AdsStealth", func(t *testing.T) {
		post, gotVkErr := vkUser.WallPostAdsStealth(map[string]string{
			"owner_id": strconv.Itoa(-vkGroupID),
			"message":  "Test AdsStealth",
		})
		if gotVkErr.Code != 0 {
			t.Errorf("VK.WallPostAdsStealth() gotVkErr = %v, want %v", gotVkErr, 0)
		}

		_, gotVkErr = vkUser.WallEditAdsStealth(map[string]string{
			"owner_id": strconv.Itoa(-vkGroupID),
			"post_id":  strconv.Itoa(post.PostID),
			"message":  "Test AdsStealth edited",
		})
		if gotVkErr.Code != 0 {
			t.Errorf("VK.WallEditAdsStealth() gotVkErr = %v, want %v", gotVkErr, 0)
		}
	})
}

func TestVK_WallReportComment(t *testing.T) {
	if vkUser.AccessToken == "" {
		t.Skip("USER_TOKEN empty")
	}

	t.Run("WallReportComment", func(t *testing.T) {
		resp, gotVkErr := vkUser.WallReportComment(map[string]string{
			"owner_id":   "66748",
			"comment_id": "4136",
			// "reason":   "3",
		})
		if resp != 1 {
			t.Errorf("VK.WallReportComment() resp = %v, want %v", resp, 1)
		}
		if gotVkErr.Code != 0 {
			t.Errorf("VK.WallReportComment() gotVkErr = %v, want %v", gotVkErr, 0)
		}
	})
}

func TestVK_WallReportPost(t *testing.T) {
	if vkUser.AccessToken == "" {
		t.Skip("USER_TOKEN empty")
	}

	t.Run("WallReportPost", func(t *testing.T) {
		resp, gotVkErr := vkUser.WallReportPost(map[string]string{
			"owner_id": "66748",
			"post_id":  "3821",
			"reason":   "3",
		})
		if resp != 1 {
			t.Errorf("VK.WallReportPost() resp = %v, want %v", resp, 1)
		}
		if gotVkErr.Code != 0 {
			t.Errorf("VK.WallReportPost() gotVkErr = %v, want %v", gotVkErr, 0)
		}
	})
}

func TestVK_WallRepost(t *testing.T) {
	if vkUser.AccessToken == "" {
		t.Skip("USER_TOKEN empty")
	}

	t.Run("WallRepost", func(t *testing.T) {
		_, gotVkErr := vkUser.WallRepost(map[string]string{
			"object": "wall85635407_3133",
		})
		if gotVkErr.Code != 0 {
			t.Errorf("VK.WallRepost() gotVkErr = %v, want %v", gotVkErr, 0)
		}
	})
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

	t.Run("WallSearch", func(t *testing.T) {
		_, gotVkErr := vkService.WallSearch(params)
		if gotVkErr.Code != 0 {
			t.Errorf("VK.WallSearch() gotVkErr = %v, want %v", gotVkErr, 0)
		}
	})

	t.Run("WallSearchExtended", func(t *testing.T) {
		_, gotVkErr := vkService.WallSearchExtended(params)
		if gotVkErr.Code != 0 {
			t.Errorf("VK.WallSearchExtended() gotVkErr = %v, want %v", gotVkErr, 0)
		}
	})
}
