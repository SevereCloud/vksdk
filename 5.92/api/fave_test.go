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
	_, gotVkErr := vkUser.FaveAddArticle(map[string]string{
		"url": "https://vk.com/@vkappsdev-vk-apps-kak-popast-v-katalog",
	})
	if gotVkErr.Code != 0 {
		t.Errorf("VK.FaveAddArticle() gotVkErr = %v, want %v", gotVkErr, 0)
	}

	res, vkErr := vkUser.FaveGet(map[string]string{"item_type": "article"})
	if vkErr.Code != 0 {
		log.Fatal(vkErr)
	}

	time.Sleep(300 * time.Millisecond)
	_, gotVkErr = vkUser.FaveRemoveArticle(map[string]string{
		"owner_id":   strconv.Itoa(res.Items[0].Article.OwnerID),
		"article_id": strconv.Itoa(res.Items[0].Article.ID),
	})
	if gotVkErr.Code != 0 {
		t.Errorf("VK.FaveAddArticle() gotVkErr = %v, want %v", gotVkErr, 0)
	}
}

func TestVK_FaveLink(t *testing.T) {
	if vkUser.AccessToken == "" {
		t.Skip("USER_TOKEN empty")
	}

	time.Sleep(300 * time.Millisecond)
	_, gotVkErr := vkUser.FaveAddLink(map[string]string{
		"link": "https://ya.ru",
	})
	if gotVkErr.Code != 0 {
		t.Errorf("VK.FaveAddLink() gotVkErr = %v, want %v", gotVkErr, 0)
	}

	res, vkErr := vkUser.FaveGet(map[string]string{"item_type": "link"})
	if vkErr.Code != 0 {
		log.Fatal(vkErr)
	}

	time.Sleep(300 * time.Millisecond)
	_, gotVkErr = vkUser.FaveRemoveLink(map[string]string{
		"link_id": res.Items[0].Link.ID,
	})
	if gotVkErr.Code != 0 {
		t.Errorf("VK.FaveRemoveLink() gotVkErr = %v, want %v", gotVkErr, 0)
	}
}

func TestVK_FavePage(t *testing.T) {
	if vkUser.AccessToken == "" {
		t.Skip("USER_TOKEN empty")
	}

	time.Sleep(300 * time.Millisecond)
	_, gotVkErr := vkUser.FaveAddPage(map[string]string{
		"user_id": "1",
	})
	if gotVkErr.Code != 0 {
		t.Errorf("VK.FaveAddPage() gotVkErr = %v, want %v", gotVkErr, 0)
	}

	time.Sleep(300 * time.Millisecond)
	_, gotVkErr = vkUser.FaveTrackPageInteraction(map[string]string{
		"user_id": "1",
	})
	if gotVkErr.Code != 0 {
		t.Errorf("VK.FaveAddPage() gotVkErr = %v, want %v", gotVkErr, 0)
	}
	time.Sleep(300 * time.Millisecond)
	_, gotVkErr = vkUser.FaveRemovePage(map[string]string{
		"user_id": "1",
	})
	if gotVkErr.Code != 0 {
		t.Errorf("VK.FaveRemovePage() gotVkErr = %v, want %v", gotVkErr, 0)
	}
}

func TestVK_FavePost(t *testing.T) {
	if vkUser.AccessToken == "" {
		t.Skip("USER_TOKEN empty")
	}

	time.Sleep(300 * time.Millisecond)
	_, gotVkErr := vkUser.FaveAddPost(map[string]string{
		"owner_id": "-28551727",
		"id":       "5713",
	})
	if gotVkErr.Code != 0 {
		t.Errorf("VK.FaveAddPost() gotVkErr = %v, want %v", gotVkErr, 0)
	}

	time.Sleep(300 * time.Millisecond)
	_, gotVkErr = vkUser.FaveRemovePost(map[string]string{
		"owner_id": "-28551727",
		"id":       "5713",
	})
	if gotVkErr.Code != 0 {
		t.Errorf("VK.FaveRemovePost() gotVkErr = %v, want %v", gotVkErr, 0)
	}
}

func TestVK_FaveProduct(t *testing.T) {
	if vkUser.AccessToken == "" {
		t.Skip("USER_TOKEN empty")
	}

	time.Sleep(300 * time.Millisecond)
	_, gotVkErr := vkUser.FaveAddProduct(map[string]string{
		"owner_id": "-169097025",
		"id":       "3398864",
	})
	if gotVkErr.Code != 0 {
		t.Errorf("VK.FaveAddProduct() gotVkErr = %v, want %v", gotVkErr, 0)
	}

	time.Sleep(300 * time.Millisecond)
	_, gotVkErr = vkUser.FaveRemoveProduct(map[string]string{
		"owner_id": "-169097025",
		"id":       "3398864",
	})
	if gotVkErr.Code != 0 {
		t.Errorf("VK.FaveRemoveProduct() gotVkErr = %v, want %v", gotVkErr, 0)
	}
}

func TestVK_FaveTag(t *testing.T) {
	if vkUser.AccessToken == "" {
		t.Skip("USER_TOKEN empty")
	}

	time.Sleep(300 * time.Millisecond)
	tagResponse, vkErr := vkUser.FaveAddTag(map[string]string{
		"name": "Test1",
	})
	if vkErr.Code != 0 {
		log.Fatal(vkErr)
	}

	time.Sleep(300 * time.Millisecond)
	_, gotVkErr := vkUser.FaveSetTags(map[string]string{
		"link_url": "https://google.ru",
		"tag_ids":  strconv.Itoa(tagResponse.ID),
	})
	if gotVkErr.Code != 0 {
		t.Errorf("VK.FaveSetTags() gotVkErr = %v, want %v", gotVkErr, 0)
	}

	time.Sleep(300 * time.Millisecond)
	_, gotVkErr = vkUser.FaveSetPageTags(map[string]string{
		"group_id": "1",
		"tag_ids":  strconv.Itoa(tagResponse.ID),
	})
	if gotVkErr.Code != 0 {
		t.Errorf("VK.FaveSetPageTags() gotVkErr = %v, want %v", gotVkErr, 0)
	}

	time.Sleep(300 * time.Millisecond)
	_, gotVkErr = vkUser.FaveEditTag(map[string]string{
		"id":   strconv.Itoa(tagResponse.ID),
		"name": "Test2",
	})
	if gotVkErr.Code != 0 {
		t.Errorf("VK.FaveEditTag() gotVkErr = %v, want %v", gotVkErr, 0)
	}

	tags, vkErr := vkUser.FaveGetTags(map[string]string{})
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
	_, gotVkErr = vkUser.FaveReorderTags(map[string]string{
		"ids": ids,
	})
	if gotVkErr.Code != 0 {
		t.Errorf("VK.FaveReorderTags() gotVkErr = %v, want %v", gotVkErr, 0)
	}

	time.Sleep(300 * time.Millisecond)
	_, gotVkErr = vkUser.FaveRemoveTag(map[string]string{
		"id": strconv.Itoa(tags.Items[tags.Count-1].ID),
	})
	if gotVkErr.Code != 0 {
		t.Errorf("VK.FaveRemoveTag() gotVkErr = %v, want %v", gotVkErr, 0)
	}

	_, _ = vkUser.FaveRemoveTag(map[string]string{
		"id": strconv.Itoa(tags.Items[tags.Count-2].ID),
	})
}
func TestVK_FaveVideo(t *testing.T) {
	if vkUser.AccessToken == "" {
		t.Skip("USER_TOKEN empty")
	}

	time.Sleep(300 * time.Millisecond)
	_, gotVkErr := vkUser.FaveAddVideo(map[string]string{
		"owner_id": "-84585194",
		"id":       "456239018",
	})
	if gotVkErr.Code != 0 {
		t.Errorf("VK.FaveAddVideo() gotVkErr = %v, want %v", gotVkErr, 0)
	}

	time.Sleep(300 * time.Millisecond)
	_, gotVkErr = vkUser.FaveRemoveVideo(map[string]string{
		"owner_id": "-84585194",
		"id":       "456239018",
	})
	if gotVkErr.Code != 0 {
		t.Errorf("VK.FaveRemoveVideo() gotVkErr = %v, want %v", gotVkErr, 0)
	}
}

func TestVK_FaveGet(t *testing.T) {
	if vkUser.AccessToken == "" {
		t.Skip("USER_TOKEN empty")
	}

	time.Sleep(300 * time.Millisecond)
	_, gotVkErr := vkUser.FaveGet(map[string]string{})
	if gotVkErr.Code != 0 {
		t.Errorf("VK.FaveGet() gotVkErr = %v, want %v", gotVkErr, 0)
	}

	_, gotVkErr = vkUser.FaveGetExtended(map[string]string{})
	if gotVkErr.Code != 0 {
		t.Errorf("VK.FaveGetExtended() gotVkErr = %v, want %v", gotVkErr, 0)
	}
}

func TestVK_FaveGetPages(t *testing.T) {
	if vkUser.AccessToken == "" {
		t.Skip("USER_TOKEN empty")
	}

	time.Sleep(300 * time.Millisecond)
	_, gotVkErr := vkUser.FaveGetPages(map[string]string{})
	if gotVkErr.Code != 0 {
		t.Errorf("VK.FaveGetPages() gotVkErr = %v, want %v", gotVkErr, 0)
	}
}

func TestVK_FaveMarkSeen(t *testing.T) {
	if vkUser.AccessToken == "" {
		t.Skip("USER_TOKEN empty")
	}

	time.Sleep(300 * time.Millisecond)
	_, gotVkErr := vkUser.FaveMarkSeen(map[string]string{})
	if gotVkErr.Code != 0 {
		t.Errorf("VK.FaveMarkSeen() gotVkErr = %v, want %v", gotVkErr, 0)
	}
}
