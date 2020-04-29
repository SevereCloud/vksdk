package api_test

import (
	"net/http"
	"testing"

	"github.com/SevereCloud/vksdk/api"
	"github.com/stretchr/testify/assert"
)

func MarketAdd(t *testing.T) int {
	needUserToken(t)
	needGroupToken(t)

	response, err := http.Get(photoURL)
	if err != nil {
		t.Fatalf("%v", err)
	}
	defer response.Body.Close()

	res, err := vkUser.UploadMarketPhoto(vkGroupID, true, response.Body)
	noError(t, err)

	market, err := vkUser.MarketAdd(api.Params{
		"owner_id":      -vkGroupID,
		"name":          "Test",
		"description":   "Test description",
		"category_id":   1,
		"price":         1000,
		"main_photo_id": res[0].ID,
	})
	noError(t, err)
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
	noError(t, err)

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
	noError(t, err)
	assert.NotEmpty(t, album)

	_, err = vkUser.MarketEditAlbum(api.Params{
		"owner_id": -vkGroupID,
		"album_id": album.MarketAlbumID,
		"title":    "Test edit album",
	})
	noError(t, err)

	marketID := MarketAdd(t)
	marketBeforeID := MarketAdd(t)

	_, err = vkUser.MarketAddToAlbum(api.Params{
		"owner_id":  -vkGroupID,
		"album_ids": album.MarketAlbumID,
		"item_id":   marketID,
	})
	noError(t, err)

	_, err = vkUser.MarketAddToAlbum(api.Params{
		"owner_id":  -vkGroupID,
		"album_ids": album.MarketAlbumID,
		"item_id":   marketBeforeID,
	})
	noError(t, err)

	_, err = vkUser.MarketReorderItems(api.Params{
		"owner_id": -vkGroupID,
		"album_id": album.MarketAlbumID,
		"item_id":  marketID,
		"before":   marketBeforeID,
	})
	noError(t, err)

	_, err = vkUser.MarketRemoveFromAlbum(api.Params{
		"owner_id":  -vkGroupID,
		"album_ids": album.MarketAlbumID,
		"item_id":   marketBeforeID,
	})
	noError(t, err)

	albumAfter, err := vkUser.MarketAddAlbum(api.Params{
		"owner_id": -vkGroupID,
		"title":    "Test album2",
	})
	noError(t, err)

	_, err = vkUser.MarketReorderAlbums(api.Params{
		"owner_id": -vkGroupID,
		"album_id": album.MarketAlbumID,
		"after":    albumAfter.MarketAlbumID,
	})
	noError(t, err)

	_, err = vkUser.MarketDeleteAlbum(api.Params{
		"owner_id": -vkGroupID,
		"album_id": album.MarketAlbumID,
	})
	noError(t, err)
	_, err = vkUser.MarketDeleteAlbum(api.Params{
		"owner_id": -vkGroupID,
		"album_id": albumAfter.MarketAlbumID,
	})
	noError(t, err)
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
		assert.NotEmpty(t, res.Items[0].Date)
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
		assert.NotEmpty(t, res.Items[0].Reposts)
		assert.NotEmpty(t, res.Items[0].ViewsCount)
	}
}

func TestVK_MarketGetCategories(t *testing.T) {
	t.Parallel()

	needUserToken(t)

	res, err := vkUser.MarketGetCategories(api.Params{})
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
		assert.NotEmpty(t, res.Items[0].Date)
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
		assert.NotEmpty(t, res.Items[0].Reposts)
		assert.NotEmpty(t, res.Items[0].ViewsCount)
	}
}
