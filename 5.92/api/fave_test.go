package api

import (
	"log"
	"strconv"
	"testing"
	"time"
)

func TestVK_FaveArticle(t *testing.T) {
	if vkUser.AccessToken == "" {
		t.Skip("USER_TOKEN empty")
	}

	time.Sleep(300 * time.Millisecond)
	t.Run("FaveAddArticle", func(t *testing.T) {
		gotResponse, gotVkErr := vkUser.FaveAddArticle(map[string]string{
			"url": "https://vk.com/@vkappsdev-vk-apps-kak-popast-v-katalog",
		})
		if gotResponse != 1 {
			t.Errorf("VK.FaveAddArticle() gotResponse = %v, want %v", gotResponse, 1)
		}
		if gotVkErr.Code != 0 {
			t.Errorf("VK.FaveAddArticle() gotVkErr = %v, want %v", gotVkErr, 0)
		}
	})

	res, vkErr := vkUser.FaveGet(map[string]string{"item_type": "article"})
	if vkErr.Code != 0 {
		log.Fatal(vkErr)
	}

	time.Sleep(300 * time.Millisecond)
	t.Run("FaveRemoveArticle", func(t *testing.T) {
		gotResponse, gotVkErr := vkUser.FaveRemoveArticle(map[string]string{
			"owner_id":   strconv.Itoa(res.Items[0].Article.OwnerID),
			"article_id": strconv.Itoa(res.Items[0].Article.ID),
		})
		if gotResponse != 1 {
			t.Errorf("VK.FaveAddArticle() gotResponse = %v, want %v", gotResponse, 1)
		}
		if gotVkErr.Code != 0 {
			t.Errorf("VK.FaveAddArticle() gotVkErr = %v, want %v", gotVkErr, 0)
		}
	})
}

func TestVK_FaveLink(t *testing.T) {
	if vkUser.AccessToken == "" {
		t.Skip("USER_TOKEN empty")
	}

	time.Sleep(300 * time.Millisecond)
	t.Run("FaveAddLink", func(t *testing.T) {
		gotResponse, gotVkErr := vkUser.FaveAddLink(map[string]string{
			"link": "https://ya.ru",
		})
		if gotResponse != 1 {
			t.Errorf("VK.FaveAddLink() gotResponse = %v, want %v", gotResponse, 1)
		}
		if gotVkErr.Code != 0 {
			t.Errorf("VK.FaveAddLink() gotVkErr = %v, want %v", gotVkErr, 0)
		}
	})

	res, vkErr := vkUser.FaveGet(map[string]string{"item_type": "link"})
	if vkErr.Code != 0 {
		log.Fatal(vkErr)
	}

	time.Sleep(300 * time.Millisecond)
	t.Run("FaveRemoveLink", func(t *testing.T) {
		gotResponse, gotVkErr := vkUser.FaveRemoveLink(map[string]string{
			"link_id": res.Items[0].Link.ID,
		})
		if gotResponse != 1 {
			t.Errorf("VK.FaveRemoveLink() gotResponse = %v, want %v", gotResponse, 1)
		}
		if gotVkErr.Code != 0 {
			t.Errorf("VK.FaveRemoveLink() gotVkErr = %v, want %v", gotVkErr, 0)
		}
	})
}

func TestVK_FavePage(t *testing.T) {
	if vkUser.AccessToken == "" {
		t.Skip("USER_TOKEN empty")
	}

	time.Sleep(300 * time.Millisecond)
	t.Run("FaveAddPage", func(t *testing.T) {
		gotResponse, gotVkErr := vkUser.FaveAddPage(map[string]string{
			"user_id": "1",
		})
		if gotResponse != 1 {
			t.Errorf("VK.FaveAddPage() gotResponse = %v, want %v", gotResponse, 1)
		}
		if gotVkErr.Code != 0 {
			t.Errorf("VK.FaveAddPage() gotVkErr = %v, want %v", gotVkErr, 0)
		}
	})

	time.Sleep(300 * time.Millisecond)
	t.Run("FaveTrackPageInteraction", func(t *testing.T) {
		gotResponse, gotVkErr := vkUser.FaveTrackPageInteraction(map[string]string{
			"user_id": "1",
		})
		if gotResponse != 1 {
			t.Errorf("VK.FaveAddPage() gotResponse = %v, want %v", gotResponse, 1)
		}
		if gotVkErr.Code != 0 {
			t.Errorf("VK.FaveAddPage() gotVkErr = %v, want %v", gotVkErr, 0)
		}
	})

	time.Sleep(300 * time.Millisecond)
	t.Run("FaveRemovePage", func(t *testing.T) {
		gotResponse, gotVkErr := vkUser.FaveRemovePage(map[string]string{
			"user_id": "1",
		})
		if gotResponse != 1 {
			t.Errorf("VK.FaveRemovePage() gotResponse = %v, want %v", gotResponse, 1)
		}
		if gotVkErr.Code != 0 {
			t.Errorf("VK.FaveRemovePage() gotVkErr = %v, want %v", gotVkErr, 0)
		}
	})
}

func TestVK_FavePost(t *testing.T) {
	if vkUser.AccessToken == "" {
		t.Skip("USER_TOKEN empty")
	}

	time.Sleep(300 * time.Millisecond)
	t.Run("FaveAddPost", func(t *testing.T) {
		gotResponse, gotVkErr := vkUser.FaveAddPost(map[string]string{
			"owner_id": "-28551727",
			"id":       "5713",
		})
		if gotResponse != 1 {
			t.Errorf("VK.FaveAddPost() gotResponse = %v, want %v", gotResponse, 1)
		}
		if gotVkErr.Code != 0 {
			t.Errorf("VK.FaveAddPost() gotVkErr = %v, want %v", gotVkErr, 0)
		}
	})

	time.Sleep(300 * time.Millisecond)
	t.Run("FaveRemovePost", func(t *testing.T) {
		gotResponse, gotVkErr := vkUser.FaveRemovePost(map[string]string{
			"owner_id": "-28551727",
			"id":       "5713",
		})
		if gotResponse != 1 {
			t.Errorf("VK.FaveRemovePost() gotResponse = %v, want %v", gotResponse, 1)
		}
		if gotVkErr.Code != 0 {
			t.Errorf("VK.FaveRemovePost() gotVkErr = %v, want %v", gotVkErr, 0)
		}
	})
}

func TestVK_FaveProduct(t *testing.T) {
	if vkUser.AccessToken == "" {
		t.Skip("USER_TOKEN empty")
	}

	time.Sleep(300 * time.Millisecond)
	t.Run("FaveAddProduct", func(t *testing.T) {
		gotResponse, gotVkErr := vkUser.FaveAddProduct(map[string]string{
			"owner_id": "-169097025",
			"id":       "3398864",
		})
		if gotResponse != 1 {
			t.Errorf("VK.FaveAddProduct() gotResponse = %v, want %v", gotResponse, 1)
		}
		if gotVkErr.Code != 0 {
			t.Errorf("VK.FaveAddProduct() gotVkErr = %v, want %v", gotVkErr, 0)
		}
	})

	time.Sleep(300 * time.Millisecond)
	t.Run("FaveRemoveProduct", func(t *testing.T) {
		gotResponse, gotVkErr := vkUser.FaveRemoveProduct(map[string]string{
			"owner_id": "-169097025",
			"id":       "3398864",
		})
		if gotResponse != 1 {
			t.Errorf("VK.FaveRemoveProduct() gotResponse = %v, want %v", gotResponse, 1)
		}
		if gotVkErr.Code != 0 {
			t.Errorf("VK.FaveRemoveProduct() gotVkErr = %v, want %v", gotVkErr, 0)
		}
	})
}

func TestVK_FaveTag(t *testing.T) {
	if vkUser.AccessToken == "" {
		t.Skip("USER_TOKEN empty")
	}

	time.Sleep(300 * time.Millisecond)
	t.Run("FaveAddTag", func(t *testing.T) {
		_, gotVkErr := vkUser.FaveAddTag(map[string]string{
			"name": "Test",
		})
		if gotVkErr.Code != 0 {
			t.Errorf("VK.FaveAddTag() gotVkErr = %v, want %v", gotVkErr, 0)
		}
	})
	tagResponse, vkErr := vkUser.FaveAddTag(map[string]string{
		"name": "Test1",
	})
	if vkErr.Code != 0 {
		log.Fatal(vkErr)
	}

	time.Sleep(300 * time.Millisecond)
	t.Run("FaveSetTags", func(t *testing.T) {
		_, gotVkErr := vkUser.FaveSetTags(map[string]string{
			"link_url": "https://google.ru",
			"tag_ids":  strconv.Itoa(tagResponse.ID),
		})
		if gotVkErr.Code != 0 {
			t.Errorf("VK.FaveSetTags() gotVkErr = %v, want %v", gotVkErr, 0)
		}
	})

	time.Sleep(300 * time.Millisecond)
	t.Run("FaveSetPageTags", func(t *testing.T) {
		_, gotVkErr := vkUser.FaveSetPageTags(map[string]string{
			"group_id": "1",
			"tag_ids":  strconv.Itoa(tagResponse.ID),
		})
		if gotVkErr.Code != 0 {
			t.Errorf("VK.FaveSetPageTags() gotVkErr = %v, want %v", gotVkErr, 0)
		}
	})

	time.Sleep(300 * time.Millisecond)
	t.Run("FaveEditTag", func(t *testing.T) {
		gotResponse, gotVkErr := vkUser.FaveEditTag(map[string]string{
			"id":   strconv.Itoa(tagResponse.ID),
			"name": "Test2",
		})
		if gotResponse != 1 {
			t.Errorf("VK.FaveEditTag() gotResponse = %v, want %v", gotResponse, 1)
		}
		if gotVkErr.Code != 0 {
			t.Errorf("VK.FaveEditTag() gotVkErr = %v, want %v", gotVkErr, 0)
		}
	})
	tags, vkErr := vkUser.FaveGetTags()
	if vkErr.Code != 0 {
		log.Fatal(vkErr)
	}

	// FaveReorderTags need all tags
	var ids string
	for i, tag := range tags.Items {
		ids += strconv.Itoa(tag.ID)
		if i != tags.Count-1 {
			ids += ","
		}
	}

	time.Sleep(300 * time.Millisecond)
	t.Run("FaveReorderTags", func(t *testing.T) {
		gotResponse, gotVkErr := vkUser.FaveReorderTags(map[string]string{
			"ids": ids,
		})
		if gotResponse != 1 {
			t.Errorf("VK.FaveReorderTags() gotResponse = %v, want %v", gotResponse, 1)
		}
		if gotVkErr.Code != 0 {
			t.Errorf("VK.FaveReorderTags() gotVkErr = %v, want %v", gotVkErr, 0)
		}
	})

	time.Sleep(300 * time.Millisecond)
	t.Run("FaveRemoveTag", func(t *testing.T) {
		gotResponse, gotVkErr := vkUser.FaveRemoveTag(map[string]string{
			"id": strconv.Itoa(tags.Items[tags.Count-1].ID),
		})
		if gotResponse != 1 {
			t.Errorf("VK.FaveRemoveTag() gotResponse = %v, want %v", gotResponse, 1)
		}
		if gotVkErr.Code != 0 {
			t.Errorf("VK.FaveRemoveTag() gotVkErr = %v, want %v", gotVkErr, 0)
		}
	})

	_, _ = vkUser.FaveRemoveTag(map[string]string{
		"id": strconv.Itoa(tags.Items[tags.Count-2].ID),
	})
}

func TestVK_FaveVideo(t *testing.T) {
	if vkUser.AccessToken == "" {
		t.Skip("USER_TOKEN empty")
	}

	time.Sleep(300 * time.Millisecond)
	t.Run("FaveAddVideo", func(t *testing.T) {
		gotResponse, gotVkErr := vkUser.FaveAddVideo(map[string]string{
			"owner_id": "-84585194",
			"id":       "456239018",
		})
		if gotResponse != 1 {
			t.Errorf("VK.FaveAddVideo() gotResponse = %v, want %v", gotResponse, 1)
		}
		if gotVkErr.Code != 0 {
			t.Errorf("VK.FaveAddVideo() gotVkErr = %v, want %v", gotVkErr, 0)
		}
	})

	time.Sleep(300 * time.Millisecond)
	t.Run("FaveRemoveVideo", func(t *testing.T) {
		gotResponse, gotVkErr := vkUser.FaveRemoveVideo(map[string]string{
			"owner_id": "-84585194",
			"id":       "456239018",
		})
		if gotResponse != 1 {
			t.Errorf("VK.FaveRemoveVideo() gotResponse = %v, want %v", gotResponse, 1)
		}
		if gotVkErr.Code != 0 {
			t.Errorf("VK.FaveRemoveVideo() gotVkErr = %v, want %v", gotVkErr, 0)
		}
	})
}

func TestVK_FaveGet(t *testing.T) {
	if vkUser.AccessToken == "" {
		t.Skip("USER_TOKEN empty")
	}

	time.Sleep(300 * time.Millisecond)
	t.Run("FaveGet", func(t *testing.T) {
		_, gotVkErr := vkUser.FaveGet(map[string]string{})
		if gotVkErr.Code != 0 {
			t.Errorf("VK.FaveGet() gotVkErr = %v, want %v", gotVkErr, 0)
		}
	})
	t.Run("FaveGetExtended", func(t *testing.T) {
		_, gotVkErr := vkUser.FaveGetExtended(map[string]string{})
		if gotVkErr.Code != 0 {
			t.Errorf("VK.FaveGetExtended() gotVkErr = %v, want %v", gotVkErr, 0)
		}
	})
}

func TestVK_FaveGetPages(t *testing.T) {
	if vkUser.AccessToken == "" {
		t.Skip("USER_TOKEN empty")
	}

	time.Sleep(300 * time.Millisecond)
	t.Run("FaveGetPages", func(t *testing.T) {
		_, gotVkErr := vkUser.FaveGetPages(map[string]string{})
		if gotVkErr.Code != 0 {
			t.Errorf("VK.FaveGetPages() gotVkErr = %v, want %v", gotVkErr, 0)
		}
	})
}

func TestVK_FaveMarkSeen(t *testing.T) {
	if vkUser.AccessToken == "" {
		t.Skip("USER_TOKEN empty")
	}

	time.Sleep(300 * time.Millisecond)
	t.Run("FaveMarkSeen", func(t *testing.T) {
		_, gotVkErr := vkUser.FaveMarkSeen()
		if gotVkErr.Code != 0 {
			t.Errorf("VK.FaveMarkSeen() gotVkErr = %v, want %v", gotVkErr, 0)
		}
	})
}
