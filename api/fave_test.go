package api

import (
	"log"
	"strconv"
	"testing"
	"time"
)

const sleepTime = 300 * time.Millisecond

func TestVK_FaveArticle(t *testing.T) {
	if vkUser.AccessToken == "" {
		t.Skip("USER_TOKEN empty")
	}

	time.Sleep(sleepTime)
	_, err := vkUser.FaveAddArticle(map[string]string{
		"url": "https://vk.com/@vkappsdev-vk-apps-kak-popast-v-katalog",
	})
	if err != nil {
		t.Errorf("VK.FaveAddArticle() err = %v", err)
	}

	res, err := vkUser.FaveGet(map[string]string{"item_type": "article"})
	if err != nil {
		log.Fatal(err)
	}

	time.Sleep(sleepTime)
	_, err = vkUser.FaveRemoveArticle(map[string]string{
		"owner_id":   strconv.Itoa(res.Items[0].Article.OwnerID),
		"article_id": strconv.Itoa(res.Items[0].Article.ID),
	})
	if err != nil {
		t.Errorf("VK.FaveAddArticle() err = %v", err)
	}
}

func TestVK_FaveLink(t *testing.T) {
	if vkUser.AccessToken == "" {
		t.Skip("USER_TOKEN empty")
	}

	time.Sleep(sleepTime)
	_, err := vkUser.FaveAddLink(map[string]string{
		"link": "https://google.ru",
	})
	if err != nil {
		t.Errorf("VK.FaveAddLink() err = %v", err)
	}

	res, err := vkUser.FaveGet(map[string]string{"item_type": "link"})
	if err != nil {
		log.Fatal(err)
	}

	time.Sleep(sleepTime)
	_, err = vkUser.FaveRemoveLink(map[string]string{
		"link_id": res.Items[0].Link.ID,
	})
	if err != nil {
		t.Errorf("VK.FaveRemoveLink() err = %v", err)
	}
}

func TestVK_FavePage(t *testing.T) {
	if vkUser.AccessToken == "" {
		t.Skip("USER_TOKEN empty")
	}

	time.Sleep(sleepTime)
	_, err := vkUser.FaveAddPage(map[string]string{
		"user_id": "1",
	})
	if err != nil {
		t.Errorf("VK.FaveAddPage() err = %v", err)
	}

	time.Sleep(sleepTime)
	_, err = vkUser.FaveTrackPageInteraction(map[string]string{
		"user_id": "1",
	})
	if err != nil {
		t.Errorf("VK.FaveAddPage() err = %v", err)
	}
	time.Sleep(sleepTime)
	_, err = vkUser.FaveRemovePage(map[string]string{
		"user_id": "1",
	})
	if err != nil {
		t.Errorf("VK.FaveRemovePage() err = %v", err)
	}
}

func TestVK_FavePost(t *testing.T) {
	if vkUser.AccessToken == "" {
		t.Skip("USER_TOKEN empty")
	}

	time.Sleep(sleepTime)
	_, err := vkUser.FaveAddPost(map[string]string{
		"owner_id": "-28551727",
		"id":       "5713",
	})
	if err != nil {
		t.Errorf("VK.FaveAddPost() err = %v", err)
	}

	time.Sleep(sleepTime)
	_, err = vkUser.FaveRemovePost(map[string]string{
		"owner_id": "-28551727",
		"id":       "5713",
	})
	if err != nil {
		t.Errorf("VK.FaveRemovePost() err = %v", err)
	}
}

func TestVK_FaveProduct(t *testing.T) {
	if vkUser.AccessToken == "" {
		t.Skip("USER_TOKEN empty")
	}

	time.Sleep(sleepTime)
	_, err := vkUser.FaveAddProduct(map[string]string{
		"owner_id": "-169097025",
		"id":       "3398864",
	})
	if err != nil {
		t.Errorf("VK.FaveAddProduct() err = %v", err)
	}

	time.Sleep(sleepTime)
	_, err = vkUser.FaveRemoveProduct(map[string]string{
		"owner_id": "-169097025",
		"id":       "3398864",
	})
	if err != nil {
		t.Errorf("VK.FaveRemoveProduct() err = %v", err)
	}
}

func TestVK_FaveTag(t *testing.T) {
	if vkUser.AccessToken == "" {
		t.Skip("USER_TOKEN empty")
	}

	time.Sleep(sleepTime)
	_, err := vkUser.FaveAddTag(map[string]string{
		"name": "Test0",
	})
	if err != nil {
		log.Fatal(err)
	}

	time.Sleep(sleepTime)
	tagResponse, err := vkUser.FaveAddTag(map[string]string{
		"name": "Test1",
	})
	if err != nil {
		log.Fatal(err)
	}

	time.Sleep(sleepTime)
	_, err = vkUser.FaveAddLink(map[string]string{
		"link": "https://google.ru",
	})
	if err != nil {
		t.Errorf("VK.FaveAddLink() err = %v", err)
	}

	time.Sleep(sleepTime)
	_, err = vkUser.FaveSetTags(map[string]string{
		"link_url": "https://google.ru",
		"tag_ids":  strconv.Itoa(tagResponse.ID),
	})
	if err != nil {
		t.Errorf("VK.FaveSetTags() err = %v", err)
	}

	time.Sleep(sleepTime)
	_, err = vkUser.FaveAddPage(map[string]string{
		"group_id": "1",
	})
	if err != nil {
		t.Errorf("VK.FaveAddPage() err = %v", err)
	}

	time.Sleep(sleepTime)
	_, err = vkUser.FaveSetPageTags(map[string]string{
		"group_id": "1",
		"tag_ids":  strconv.Itoa(tagResponse.ID),
	})
	if err != nil {
		t.Errorf("VK.FaveSetPageTags() err = %v", err)
	}

	time.Sleep(sleepTime)
	_, err = vkUser.FaveEditTag(map[string]string{
		"id":   strconv.Itoa(tagResponse.ID),
		"name": "Test2",
	})
	if err != nil {
		t.Errorf("VK.FaveEditTag() err = %v", err)
	}

	tags, err := vkUser.FaveGetTags(map[string]string{})
	if err != nil {
		log.Fatal(err)
	}

	// FaveReorderTags need all tags
	var ids string
	for i, tag := range tags.Items {
		ids += strconv.Itoa(tag.ID)
		if i != tags.Count-1 {
			ids += ","
		}
	}

	time.Sleep(sleepTime)
	_, err = vkUser.FaveReorderTags(map[string]string{
		"ids": ids,
	})
	if err != nil {
		t.Errorf("VK.FaveReorderTags() err = %v", err)
	}

	time.Sleep(sleepTime)
	_, err = vkUser.FaveRemoveTag(map[string]string{
		"id": strconv.Itoa(tags.Items[tags.Count-1].ID),
	})
	if err != nil {
		t.Errorf("VK.FaveRemoveTag() err = %v", err)
	}

	_, _ = vkUser.FaveRemoveTag(map[string]string{
		"id": strconv.Itoa(tags.Items[tags.Count-2].ID),
	})
}
func TestVK_FaveVideo(t *testing.T) {
	if vkUser.AccessToken == "" {
		t.Skip("USER_TOKEN empty")
	}

	time.Sleep(sleepTime)
	_, err := vkUser.FaveAddVideo(map[string]string{
		"owner_id": "-84585194",
		"id":       "456239018",
	})
	if err != nil {
		t.Errorf("VK.FaveAddVideo() err = %v", err)
	}

	time.Sleep(sleepTime)
	_, err = vkUser.FaveRemoveVideo(map[string]string{
		"owner_id": "-84585194",
		"id":       "456239018",
	})
	if err != nil {
		t.Errorf("VK.FaveRemoveVideo() err = %v", err)
	}
}

func TestVK_FaveGet(t *testing.T) {
	if vkUser.AccessToken == "" {
		t.Skip("USER_TOKEN empty")
	}

	time.Sleep(sleepTime)
	_, err := vkUser.FaveGet(map[string]string{})
	if err != nil {
		t.Errorf("VK.FaveGet() err = %v", err)
	}

	time.Sleep(sleepTime)
	_, err = vkUser.FaveGetExtended(map[string]string{})
	if err != nil {
		t.Errorf("VK.FaveGetExtended() err = %v", err)
	}
}

func TestVK_FaveGetPages(t *testing.T) {
	if vkUser.AccessToken == "" {
		t.Skip("USER_TOKEN empty")
	}

	time.Sleep(sleepTime)
	_, err := vkUser.FaveGetPages(map[string]string{})
	if err != nil {
		t.Errorf("VK.FaveGetPages() err = %v", err)
	}
}

func TestVK_FaveMarkSeen(t *testing.T) {
	if vkUser.AccessToken == "" {
		t.Skip("USER_TOKEN empty")
	}

	time.Sleep(sleepTime)
	_, err := vkUser.FaveMarkSeen(map[string]string{})
	if err != nil {
		t.Errorf("VK.FaveMarkSeen() err = %v", err)
	}
}
