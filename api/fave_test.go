package api_test

import (
	"log"
	"testing"
	"time"

	"github.com/SevereCloud/vksdk/v2/api"

	"github.com/SevereCloud/vksdk/v2/object"
	"github.com/stretchr/testify/assert"
)

func testFave(t *testing.T, f object.FaveItem) {
	t.Helper()

	assert.NotEmpty(t, f.AddedDate, "f.AddedDate")
	// assert.NotEmpty(t, f.Seen, "f.Seen")
	assert.NotEmpty(t, f.Type, "f.Type")

	switch f.Type {
	case "link":
		if assert.NotEmpty(t, f.Link, "f.Link") {
			assert.NotEmpty(t, f.Link.URL, "f.Link.URL")
			assert.NotEmpty(t, f.Link.Title, "f.Link.Title")
			assert.NotEmpty(t, f.Link.Caption, "f.Link.Caption")
			// assert.NotEmpty(t, f.Link.Description, "f.Link.Description")
			// assert.NotEmpty(t, f.Link.Photo, "f.Link.Photo")
			assert.NotEmpty(t, f.Link.ID, "f.Link.ID")
		}
	case "post":
		if assert.NotEmpty(t, f.Post, "f.Post") {
			assert.NotEmpty(t, f.Post.ID, "f.Post.ID")
			assert.NotEmpty(t, f.Post.FromID, "f.Post.FromID")
			assert.NotEmpty(t, f.Post.OwnerID, "f.Post.OwnerID")
			assert.NotEmpty(t, f.Post.Date, "f.Post.Date")
			assert.NotEmpty(t, f.Post.PostType, "f.Post.PostType")
			assert.NotEmpty(t, f.Post.Text, "f.Post.Text")
			// assert.NotEmpty(t, f.Post.SignerID, "f.Post.SignerID")
			// assert.NotEmpty(t,f.Post.MarkedAsAds)
			assert.NotEmpty(t, f.Post.PostSource, "f.Post.PostSource")
			assert.NotEmpty(t, f.Post.Comments, "f.Post.Comments")
		}
	case "video":
		if assert.NotEmpty(t, f.Video, "f.Video") {
			assert.NotEmpty(t, f.Video.ID, "f.Video.ID")
			assert.NotEmpty(t, f.Video.OwnerID, "f.Video.OwnerID")
			assert.NotEmpty(t, f.Video.Title, "f.Video.Title")
			assert.NotEmpty(t, f.Video.Duration, "f.Video.Duration")
			// assert.NotEmpty(t, f.Video.Description, "f.Video.Description")
			assert.NotEmpty(t, f.Video.Date, "f.Video.Date")
			// assert.NotEmpty(t, f.Video.Comments, "f.Video.Comments")
			assert.NotEmpty(t, f.Video.Views, "f.Video.Views")
			assert.NotEmpty(t, f.Video.Width, "f.Video.Width")
			assert.NotEmpty(t, f.Video.Height, "f.Video.Height")
			assert.NotEmpty(t, f.Video.Image, "f.Video.Image")
			// assert.NotEmpty(t,f.Video.IsFavorite)
			assert.NotEmpty(t, f.Video.FirstFrame, "f.Video.FirstFrame")
			assert.NotEmpty(t, f.Video.Player, "f.Video.Player")
			assert.NotEmpty(t, f.Video.CanAdd, "f.Video.CanAdd")
			// assert.NotEmpty(t, f.Video.Added, "f.Video.Added")
			// assert.NotEmpty(t, f.Video.CanComment, "f.Video.CanComment")
			assert.NotEmpty(t, f.Video.CanRepost, "f.Video.CanRepost")
			assert.NotEmpty(t, f.Video.CanLike, "f.Video.CanLike")
			assert.NotEmpty(t, f.Video.CanAddToFaves, "f.Video.CanAddToFaves")
			assert.NotEmpty(t, f.Video.Type, "f.Video.Type")
			assert.NotEmpty(t, f.Video.Likes, "f.Video.Likes")
			assert.NotEmpty(t, f.Video.Reposts, "f.Video.Reposts")
		}
	case "product":
		if assert.NotEmpty(t, f.Product, "f.Product") {
			// assert.NotEmpty(t,f.Product.Availability)
			if assert.NotEmpty(t, f.Product.Category, "f.Product.Category") {
				assert.NotEmpty(t, f.Product.Category.ID, "f.Product.Category.ID")
				assert.NotEmpty(t, f.Product.Category.Name, "f.Product.Category.Name")

				if assert.NotEmpty(t, f.Product.Category.Section, "f.Product.Category.Section") {
					// assert.NotEmpty(t, f.Product.Category.Section.ID, "f.Product.Category.Section.ID")
					assert.NotEmpty(t, f.Product.Category.Section.Name, "f.Product.Category.Section.Name")
				}
			}

			// assert.NotEmpty(t, f.Product.Date, "f.Product.Date")
			assert.NotEmpty(t, f.Product.Description, "f.Product.Description")
			// assert.NotEmpty(t,f.Product.ExternalID)
			assert.NotEmpty(t, f.Product.ID, "f.Product.ID")
			// assert.NotEmpty(t,f.Product.IsFavorite)
			assert.NotEmpty(t, f.Product.OwnerID, "f.Product.OwnerID")

			if assert.NotEmpty(t, f.Product.Price, "f.Product.Price") {
				assert.NotEmpty(t, f.Product.Price.Amount, "f.Product.Price.Amount")

				if assert.NotEmpty(t, f.Product.Price.Currency, "f.Product.Price.Currency") {
					assert.NotEmpty(t, f.Product.Price.Currency.ID, "f.Product.Price.Currency.ID")
					assert.NotEmpty(t, f.Product.Price.Currency.Name, "f.Product.Price.Currency.Name")
				}

				assert.NotEmpty(t, f.Product.Price.Text, "f.Product.Price.Text")
			}

			assert.NotEmpty(t, f.Product.ThumbPhoto, "f.Product.ThumbPhoto")
			assert.NotEmpty(t, f.Product.Title, "f.Product.Title")
			// assert.NotEmpty(t,f.Product.AlbumsIDs)
			// assert.NotEmpty(t, f.Product.CanComment, "f.Product.CanComment")
			// assert.NotEmpty(t, f.Product.CanRepost, "f.Product.CanRepost")
			// assert.NotEmpty(t, f.Product.Likes, "f.Product.Likes")
			// assert.NotEmpty(t, f.Product.Reposts, "f.Product.Reposts")
			// assert.NotEmpty(t, f.Product.ViewsCount, "f.Product.ViewsCount")
			assert.NotEmpty(t, f.Product.Photos, "f.Product.Photos")
		}
	case "article":
		if assert.NotEmpty(t, f.Article, "f.Article") {
			assert.NotEmpty(t, f.Article.ID, "f.Article.ID")
			assert.NotEmpty(t, f.Article.OwnerID, "f.Article.OwnerID")
			assert.NotEmpty(t, f.Article.OwnerName, "f.Article.OwnerName")
			assert.NotEmpty(t, f.Article.OwnerPhoto, "f.Article.OwnerPhoto")
			assert.NotEmpty(t, f.Article.State, "f.Article.State")
			// assert.NotEmpty(t, f.Article.CanReport, "f.Article.CanReport")
			assert.NotEmpty(t, f.Article.Title, "f.Article.Title")
			assert.NotEmpty(t, f.Article.Subtitle, "f.Article.Subtitle")
			assert.NotEmpty(t, f.Article.Views, "f.Article.Views")
			assert.NotEmpty(t, f.Article.Shares, "f.Article.Shares")
			// assert.NotEmpty(t, f.Article.IsFavorite, "f.Article.IsFavorite")
			assert.NotEmpty(t, f.Article.URL, "f.Article.URL")
			assert.NotEmpty(t, f.Article.ViewURL, "f.Article.ViewURL")
			assert.NotEmpty(t, f.Article.AccessKey, "f.Article.AccessKey")
			assert.NotEmpty(t, f.Article.PublishedDate, "f.Article.PublishedDate")
			assert.NotEmpty(t, f.Article.Photo, "f.Article.Photo")
		}
	}
}

func TestVK_Fave(t *testing.T) {
	t.Parallel()

	needUserToken(t)

	fave, err := vkUser.FaveGet(nil)
	if !noError(t, err) {
		log.Fatal(err)
	}

	assert.NotEmpty(t, fave.Count)

	if assert.NotEmpty(t, fave.Items) {
		for _, f := range fave.Items {
			testFave(t, f)
		}
	}
}

func TestVK_FaveArticle(t *testing.T) {
	t.Parallel()

	needUserToken(t)

	time.Sleep(sleepTime)

	res, err := vkUser.FaveAddArticle(api.Params{
		"url": "https://vk.com/@vkappsdev-vk-apps-kak-popast-v-katalog",
	})
	noError(t, err)
	assert.NotEmpty(t, res)

	fave, err := vkUser.FaveGet(api.Params{"item_type": "article"})
	if !noError(t, err) {
		log.Fatal(err)
	}

	assert.NotEmpty(t, fave.Count)

	if assert.NotEmpty(t, fave.Items) {
		for _, f := range fave.Items {
			testFave(t, f)
		}
	}

	time.Sleep(sleepTime)

	res, err = vkUser.FaveRemoveArticle(api.Params{
		"owner_id":   fave.Items[0].Article.OwnerID,
		"article_id": fave.Items[0].Article.ID,
	})
	noError(t, err)
	assert.NotEmpty(t, res)
}

func TestVK_FaveLink(t *testing.T) {
	t.Parallel()

	needUserToken(t)

	time.Sleep(sleepTime)

	res, err := vkUser.FaveAddLink(api.Params{
		"link": "https://google.ru",
	})
	noError(t, err)
	assert.NotEmpty(t, res)

	fave, err := vkUser.FaveGet(api.Params{"item_type": "link"})
	if !noError(t, err) {
		log.Fatal(err)
	}

	assert.NotEmpty(t, fave.Count)

	if assert.NotEmpty(t, fave.Items) {
		for _, f := range fave.Items {
			testFave(t, f)
		}
	}

	time.Sleep(sleepTime)

	res, err = vkUser.FaveRemoveLink(api.Params{
		"link_id": fave.Items[0].Link.ID,
	})
	noError(t, err)
	assert.NotEmpty(t, res)
}

func TestVK_FavePage(t *testing.T) {
	t.Parallel()

	needUserToken(t)

	time.Sleep(sleepTime)

	res, err := vkUser.FaveAddPage(api.Params{
		"user_id": 1,
	})
	noError(t, err)
	assert.NotEmpty(t, res)

	time.Sleep(sleepTime)

	res, err = vkUser.FaveTrackPageInteraction(api.Params{
		"user_id": 1,
	})
	noError(t, err)
	assert.NotEmpty(t, res)

	time.Sleep(sleepTime)

	res, err = vkUser.FaveRemovePage(api.Params{
		"user_id": 1,
	})
	noError(t, err)
	assert.NotEmpty(t, res)
}

func TestVK_FavePost(t *testing.T) {
	t.Parallel()

	needUserToken(t)

	time.Sleep(sleepTime)

	res, err := vkUser.FaveAddPost(api.Params{
		"owner_id": -28551727,
		"id":       5713,
	})
	noError(t, err)
	assert.NotEmpty(t, res)

	time.Sleep(sleepTime)

	res, err = vkUser.FaveRemovePost(api.Params{
		"owner_id": -28551727,
		"id":       5713,
	})
	noError(t, err)
	assert.NotEmpty(t, res)
}

func TestVK_FaveProduct(t *testing.T) {
	t.Parallel()

	needUserToken(t)

	time.Sleep(sleepTime)

	res, err := vkUser.FaveAddProduct(api.Params{
		"owner_id": -169097025,
		"id":       3398864,
	})
	noError(t, err)
	assert.NotEmpty(t, res)

	time.Sleep(sleepTime)

	res, err = vkUser.FaveRemoveProduct(api.Params{
		"owner_id": -169097025,
		"id":       3398864,
	})
	noError(t, err)
	assert.NotEmpty(t, res)
}

func TestVK_FaveTag(t *testing.T) {
	t.Parallel()

	needUserToken(t)

	time.Sleep(sleepTime)

	tag, err := vkUser.FaveAddTag(api.Params{
		"name": "Test0",
	})
	noError(t, err)

	if assert.NotEmpty(t, tag) {
		assert.NotEmpty(t, tag.ID)
		assert.NotEmpty(t, tag.Name)
	}

	time.Sleep(sleepTime)

	tagResponse, err := vkUser.FaveAddTag(api.Params{
		"name": "Test1",
	})
	noError(t, err)
	assert.NotEmpty(t, tagResponse)

	time.Sleep(sleepTime)

	res, err := vkUser.FaveAddLink(api.Params{
		"link": "https://google.ru",
	})
	noError(t, err)
	assert.NotEmpty(t, res)

	time.Sleep(sleepTime)

	res, err = vkUser.FaveSetTags(api.Params{
		"link_url": "https://google.ru",
		"tag_ids":  tagResponse.ID,
	})
	noError(t, err)
	assert.NotEmpty(t, res)

	time.Sleep(sleepTime)

	res, err = vkUser.FaveAddPage(api.Params{
		"group_id": 1,
	})
	noError(t, err)
	assert.NotEmpty(t, res)

	time.Sleep(sleepTime)

	res, err = vkUser.FaveSetPageTags(api.Params{
		"group_id": 1,
		"tag_ids":  tagResponse.ID,
	})
	noError(t, err)
	assert.NotEmpty(t, res)

	time.Sleep(sleepTime)

	res, err = vkUser.FaveEditTag(api.Params{
		"id":   tagResponse.ID,
		"name": "Test2",
	})
	noError(t, err)
	assert.NotEmpty(t, res)

	tags, err := vkUser.FaveGetTags(nil)
	noError(t, err)
	assert.NotEmpty(t, tags)

	// FaveReorderTags need all tags
	ids := make([]int, 0)
	for _, item := range tags.Items {
		ids = append(ids, item.ID)
	}

	time.Sleep(sleepTime)

	res, err = vkUser.FaveReorderTags(api.Params{
		"ids": ids,
	})
	noError(t, err)
	assert.NotEmpty(t, res)

	time.Sleep(sleepTime)

	res, err = vkUser.FaveRemoveTag(api.Params{
		"id": tags.Items[tags.Count-1].ID,
	})
	noError(t, err)
	assert.NotEmpty(t, res)

	_, _ = vkUser.FaveRemoveTag(api.Params{
		"id": tags.Items[tags.Count-2].ID,
	})
}

func TestVK_FaveVideo(t *testing.T) {
	t.Parallel()

	needUserToken(t)

	time.Sleep(sleepTime)

	res, err := vkUser.FaveAddVideo(api.Params{
		"owner_id": -84585194,
		"id":       456239018,
	})
	noError(t, err)
	assert.NotEmpty(t, res)

	time.Sleep(sleepTime)

	res, err = vkUser.FaveRemoveVideo(api.Params{
		"owner_id": -84585194,
		"id":       456239018,
	})
	noError(t, err)
	assert.NotEmpty(t, res)
}

func TestVK_FaveGet(t *testing.T) {
	t.Parallel()

	needUserToken(t)

	time.Sleep(sleepTime)

	res, err := vkUser.FaveGet(nil)
	noError(t, err)
	assert.NotEmpty(t, res)

	time.Sleep(sleepTime)

	_, err = vkUser.FaveGetExtended(nil)
	noError(t, err)
	assert.NotEmpty(t, res)
}

func TestVK_FaveGetPages(t *testing.T) {
	t.Parallel()

	needUserToken(t)

	time.Sleep(sleepTime)

	res, err := vkUser.FaveGetPages(nil)
	noError(t, err)
	assert.NotEmpty(t, res.Count)

	if assert.NotEmpty(t, res.Items) {
		for _, page := range res.Items {
			assert.NotEmpty(t, page.Description)
			assert.NotEmpty(t, page.Type)
			// assert.NotEmpty(t, page.Tags)
			assert.NotEmpty(t, page.UpdatedDate)

			switch page.Type {
			case "user":
				assert.NotEmpty(t, page.User.ID)
				assert.NotEmpty(t, page.User.FirstName)
				assert.NotEmpty(t, page.User.LastName)
				// assert.NotEmpty(t, page.User.IsClosed)
				// assert.NotEmpty(t, page.User.CanAccessClosed)
			case "group":
				assert.NotEmpty(t, page.Group.ID)
				assert.NotEmpty(t, page.Group.Name)
				assert.NotEmpty(t, page.Group.ScreenName)
				// assert.NotEmpty(t, page.Group.IsClosed)
				assert.NotEmpty(t, page.Group.Type)
				assert.NotEmpty(t, page.Group.Photo50)
				assert.NotEmpty(t, page.Group.Photo100)
				assert.NotEmpty(t, page.Group.Photo200)
			}

			assert.NotEmpty(t, page.Description)
		}
	}
}

func TestVK_FaveMarkSeen(t *testing.T) {
	t.Parallel()

	needUserToken(t)

	time.Sleep(sleepTime)

	res, err := vkUser.FaveMarkSeen(nil)
	noError(t, err)
	assert.NotEmpty(t, res)
}
