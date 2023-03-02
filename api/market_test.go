package api_test

import (
	"errors"
	"net/http"
	"testing"
	"time"

	"github.com/SevereCloud/vksdk/v2/api"
	"github.com/stretchr/testify/assert"
)

func MarketAdd(t *testing.T) int {
	t.Helper()

	needUserToken(t)
	needGroupToken(t)

	response, err := http.Get(photoURL)
	if err != nil {
		t.Fatalf("%v", err)
	}
	defer response.Body.Close()

	res, err := vkUser.UploadMarketPhoto(vkGroupID, true, response.Body)
	noErrorOrFail(t, err)

	market, err := vkUser.MarketAdd(api.Params{
		"owner_id":      -vkGroupID,
		"name":          "Test",
		"description":   "Test description",
		"category_id":   1,
		"price":         1000,
		"main_photo_id": res[0].ID,
	})
	noErrorOrFail(t, err)
	assert.NotEmpty(t, market.MarketItemID)

	_, err = vkUser.MarketEdit(api.Params{
		"owner_id": -vkGroupID,
		"item_id":  market.MarketItemID,
		"name":     "Test name",
	})
	noError(t, err)

	_, err = vkUser.MarketDelete(api.Params{
		"owner_id": -vkGroupID,
		"item_id":  market.MarketItemID,
	})
	noError(t, err)

	_, err = vkUser.MarketRestore(api.Params{
		"owner_id": -vkGroupID,
		"item_id":  market.MarketItemID,
	})
	noErrorOrFail(t, err)

	return market.MarketItemID
}

func TestVK_MarketAddAlbum(t *testing.T) {
	t.Parallel()

	needUserToken(t)
	needGroupToken(t)

	album, err := vkUser.MarketAddAlbum(api.Params{
		"owner_id": -vkGroupID,
		"title":    "Test album",
	})
	noErrorOrFail(t, err)
	assert.NotEmpty(t, album)

	time.Sleep(sleepTime)

	_, err = vkUser.MarketEditAlbum(api.Params{
		"owner_id": -vkGroupID,
		"album_id": album.MarketAlbumID,
		"title":    "Test edit album",
	})
	noErrorOrFail(t, err)

	time.Sleep(sleepTime)

	marketID := MarketAdd(t)

	time.Sleep(sleepTime)

	marketBeforeID := MarketAdd(t)

	time.Sleep(sleepTime)

	_, err = vkUser.MarketAddToAlbum(api.Params{
		"owner_id":  -vkGroupID,
		"album_ids": album.MarketAlbumID,
		"item_id":   marketID,
	})
	noErrorOrFail(t, err)

	time.Sleep(sleepTime)

	_, err = vkUser.MarketAddToAlbum(api.Params{
		"owner_id":  -vkGroupID,
		"album_ids": album.MarketAlbumID,
		"item_id":   marketBeforeID,
	})
	noErrorOrFail(t, err)

	time.Sleep(2 * sleepTime)

	_, err = vkUser.MarketReorderItems(api.Params{
		"owner_id": -vkGroupID,
		"album_id": album.MarketAlbumID,
		"item_id":  marketID,
		"before":   marketBeforeID,
	})
	if err != nil {
		if errors.Is(err, api.ErrMarketNotEnabled) {
			t.Skip("VK bad error: " + err.Error())
		} else {
			noErrorOrFail(t, err)
		}
	}

	time.Sleep(sleepTime)

	_, err = vkUser.MarketRemoveFromAlbum(api.Params{
		"owner_id":  -vkGroupID,
		"album_ids": album.MarketAlbumID,
		"item_id":   marketBeforeID,
	})
	noErrorOrFail(t, err)

	time.Sleep(sleepTime)

	albumAfter, err := vkUser.MarketAddAlbum(api.Params{
		"owner_id": -vkGroupID,
		"title":    "Test album2",
	})
	noErrorOrFail(t, err)

	time.Sleep(sleepTime)

	_, err = vkUser.MarketReorderAlbums(api.Params{
		"owner_id": -vkGroupID,
		"album_id": album.MarketAlbumID,
		"after":    albumAfter.MarketAlbumID,
	})
	noErrorOrFail(t, err)

	time.Sleep(sleepTime)

	_, err = vkUser.MarketDeleteAlbum(api.Params{
		"owner_id": -vkGroupID,
		"album_id": album.MarketAlbumID,
	})
	noErrorOrFail(t, err)

	time.Sleep(sleepTime)

	_, err = vkUser.MarketDeleteAlbum(api.Params{
		"owner_id": -vkGroupID,
		"album_id": albumAfter.MarketAlbumID,
	})
	noErrorOrFail(t, err)
}

func TestVK_MarketEditOrder(t *testing.T) {
	t.Parallel()
	// TODO: TestVK_MarketEditOrder
	t.Skip("TestVK_MarketEditOrder")
}

func TestVK_MarketCreateComment(t *testing.T) {
	t.Parallel()

	needUserToken(t)

	commentID, err := vkUser.MarketCreateComment(api.Params{
		"owner_id": -124527492,
		"item_id":  250407,
		"message":  "Comment text",
	})
	noError(t, err)
	assert.NotEmpty(t, commentID)

	_, err = vkUser.MarketEditComment(api.Params{
		"owner_id":   -124527492,
		"comment_id": commentID,
		"message":    "Edited text",
	})
	noError(t, err)

	_, err = vkUser.MarketDeleteComment(api.Params{
		"owner_id":   -124527492,
		"comment_id": commentID,
	})
	noError(t, err)

	_, err = vkUser.MarketRestoreComment(api.Params{
		"owner_id":   -124527492,
		"comment_id": commentID,
	})
	noError(t, err)
}

func TestVK_MarketGet(t *testing.T) {
	t.Parallel()

	needUserToken(t)

	res, err := vkUser.MarketGet(api.Params{
		"owner_id": -124527492,
		"extended": true,
	})
	noError(t, err)
	assert.NotEmpty(t, res.Count)
	assert.NotEmpty(t, res.Items)
}

func TestVK_MarketGetAlbumByID(t *testing.T) {
	t.Parallel()

	needUserToken(t)

	res, err := vkUser.MarketGetAlbumByID(api.Params{
		"owner_id":  -124527492,
		"album_ids": 2,
	})
	noError(t, err)
	assert.NotEmpty(t, res.Count)

	if assert.NotEmpty(t, res.Items) {
		assert.NotEmpty(t, res.Items[0].ID)
		assert.NotEmpty(t, res.Items[0].OwnerID)
		assert.NotEmpty(t, res.Items[0].Title)
		assert.NotEmpty(t, res.Items[0].Count)
		assert.NotEmpty(t, res.Items[0].UpdatedTime)
	}
}

func TestVK_MarketGetAlbums(t *testing.T) {
	t.Parallel()

	needUserToken(t)

	res, err := vkUser.MarketGetAlbums(api.Params{
		"owner_id": -124527492,
	})
	noError(t, err)
	assert.NotEmpty(t, res.Count)

	if assert.NotEmpty(t, res.Items) {
		assert.NotEmpty(t, res.Items[0].ID)
		assert.NotEmpty(t, res.Items[0].OwnerID)
		assert.NotEmpty(t, res.Items[0].Title)
		assert.NotEmpty(t, res.Items[0].Count)
		assert.NotEmpty(t, res.Items[0].UpdatedTime)
	}
}

func TestVK_MarketGetByID(t *testing.T) {
	t.Parallel()

	needUserToken(t)

	res, err := vkUser.MarketGetByID(api.Params{
		"item_ids": "-124527492_250407",
		"extended": true,
	})
	noError(t, err)
	assert.NotEmpty(t, res.Count)

	if assert.NotEmpty(t, res.Items) {
		// assert.NotEmpty(t, res.Items[0].Availability)
		assert.NotEmpty(t, res.Items[0].Category.ID)
		assert.NotEmpty(t, res.Items[0].Category.Name)
		assert.NotEmpty(t, res.Items[0].Category.Section.ID)
		assert.NotEmpty(t, res.Items[0].Category.Section.Name)
		// assert.NotEmpty(t, res.Items[0].Date)
		assert.NotEmpty(t, res.Items[0].Description)
		// assert.NotEmpty(t, res.Items[0].ExternalID)
		assert.NotEmpty(t, res.Items[0].ID)
		assert.NotEmpty(t, res.Items[0].OwnerID)
		assert.NotEmpty(t, res.Items[0].Price.Amount)
		assert.NotEmpty(t, res.Items[0].Price.Currency.Name)
		assert.NotEmpty(t, res.Items[0].Price.Currency.ID)
		assert.NotEmpty(t, res.Items[0].Price.Text)
		assert.NotEmpty(t, res.Items[0].ThumbPhoto)
		assert.NotEmpty(t, res.Items[0].Title)
		assert.NotEmpty(t, res.Items[0].AlbumsIDs)
		assert.NotEmpty(t, res.Items[0].Photos)
		assert.NotEmpty(t, res.Items[0].CanComment)
		assert.NotEmpty(t, res.Items[0].CanRepost)
		assert.NotEmpty(t, res.Items[0].Likes)
		// assert.NotEmpty(t, res.Items[0].ViewsCount)
		assert.NotEmpty(t, res.Items[0].Reposts)
	}
}

func TestVK_MarketGetCategories(t *testing.T) {
	t.Parallel()

	needUserToken(t)

	res, err := vkUser.MarketGetCategories(nil)
	noError(t, err)
	assert.NotEmpty(t, res.Count)

	if assert.NotEmpty(t, res.Items) {
		assert.NotEmpty(t, res.Items[0].ID)
		assert.NotEmpty(t, res.Items[0].Name)
		assert.NotEmpty(t, res.Items[0].Section.Name)
	}
}

func TestVK_MarketGetComments(t *testing.T) {
	t.Parallel()

	needUserToken(t)

	res, err := vkUser.MarketGetComments(api.Params{
		"owner_id": -124527492,
		"item_id":  250407,
	})
	noError(t, err)
	assert.NotEmpty(t, res.Count)
	assert.NotEmpty(t, res.Items)
}

func TestVK_MarketGetCommentsExtended(t *testing.T) {
	t.Parallel()

	needUserToken(t)

	res, err := vkUser.MarketGetCommentsExtended(api.Params{
		"owner_id": -124527492,
		"item_id":  250407,
	})
	noError(t, err)
	assert.NotEmpty(t, res.Count)
	assert.NotEmpty(t, res.Items)
}

func TestVK_MarketGetGroupOrders(t *testing.T) {
	t.Parallel()
	// TODO: TestVK_MarketGetGroupOrders
	t.Skip("TestVK_MarketGetGroupOrders")
}

func TestVK_MarketGetOrderByID(t *testing.T) {
	t.Parallel()
	// TODO: TestVK_MarketGetOrderByID
	t.Skip("TestVK_MarketGetOrderByID")
}

func TestVK_MarketGetOrderItems(t *testing.T) {
	t.Parallel()
	// TODO: TestVK_MarketGetOrderItems
	t.Skip("TestVK_MarketGetOrderItems")
}

func TestVK_MarketReport(t *testing.T) {
	t.Parallel()

	needUserToken(t)

	_, err := vkUser.MarketReport(api.Params{
		"owner_id": -124527492,
		"item_id":  250407,
		"reason":   0,
	})
	noError(t, err)
}

func TestVK_MarketReportComment(t *testing.T) {
	t.Parallel()

	needUserToken(t)

	_, err := vkUser.MarketReportComment(api.Params{
		"owner_id":   -124527492,
		"comment_id": 54,
		"reason":     6,
	})
	noError(t, err)
}

func TestVK_MarketSearch(t *testing.T) {
	t.Parallel()

	t.Skip("TODO: check vk bug")
	needUserToken(t)

	res, err := vkUser.MarketSearch(api.Params{
		"q":        "spotty",
		"owner_id": -124527492,
		"extended": true,
	})
	noError(t, err)
	assert.NotEmpty(t, res.Count)

	if assert.NotEmpty(t, res.Items) {
		assert.NotEmpty(t, res.Items[0].Category.ID)
		assert.NotEmpty(t, res.Items[0].Category.Name)
		assert.NotEmpty(t, res.Items[0].Category.Section.ID)
		assert.NotEmpty(t, res.Items[0].Category.Section.Name)
		// assert.NotEmpty(t, res.Items[0].Date)
		assert.NotEmpty(t, res.Items[0].Description)
		// assert.NotEmpty(t, res.Items[0].ExternalID)
		assert.NotEmpty(t, res.Items[0].ID)
		assert.NotEmpty(t, res.Items[0].OwnerID)
		assert.NotEmpty(t, res.Items[0].Price.Amount)
		assert.NotEmpty(t, res.Items[0].Price.Currency.Name)
		assert.NotEmpty(t, res.Items[0].Price.Currency.ID)
		assert.NotEmpty(t, res.Items[0].Price.Text)
		assert.NotEmpty(t, res.Items[0].ThumbPhoto)
		assert.NotEmpty(t, res.Items[0].Title)
		assert.NotEmpty(t, res.Items[0].AlbumsIDs)
		assert.NotEmpty(t, res.Items[0].Photos)
		assert.NotEmpty(t, res.Items[0].Likes)
		// assert.NotEmpty(t, res.Items[0].ViewsCount)
		assert.NotEmpty(t, res.Items[0].Reposts)
	}
}

func TestVK_MarketSearchItems(t *testing.T) {
	t.Parallel()

	needUserToken(t)

	res, err := vkUser.MarketSearchItems(api.Params{
		"q": "test",
	})

	noError(t, err)
	assert.NotEmpty(t, res.Count)
	assert.NotEmpty(t, res.ViewType)
}
