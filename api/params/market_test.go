package params_test

import (
	"testing"

	"github.com/SevereCloud/vksdk/api/params"
	"github.com/stretchr/testify/assert"
)

func TestMarketAddBulder(t *testing.T) {
	b := params.NewMarketAddBulder()

	b.OwnerID(1)
	b.Name("text")
	b.Description("text")
	b.CategoryID(1)
	b.Price(1.1)
	b.OldPrice(1.1)
	b.Deleted(true)
	b.MainPhotoID(1)
	b.PhotoIDs([]int{1})
	b.URL("text")

	assert.Equal(t, b.Params["owner_id"], 1)
	assert.Equal(t, b.Params["name"], "text")
	assert.Equal(t, b.Params["description"], "text")
	assert.Equal(t, b.Params["category_id"], 1)
	assert.Equal(t, b.Params["price"], 1.1)
	assert.Equal(t, b.Params["old_price"], 1.1)
	assert.Equal(t, b.Params["deleted"], true)
	assert.Equal(t, b.Params["main_photo_id"], 1)
	assert.Equal(t, b.Params["photo_ids"], []int{1})
	assert.Equal(t, b.Params["url"], "text")
}

func TestMarketAddAlbumBulder(t *testing.T) {
	b := params.NewMarketAddAlbumBulder()

	b.OwnerID(1)
	b.Title("text")
	b.PhotoID(1)
	b.MainAlbum(true)

	assert.Equal(t, b.Params["owner_id"], 1)
	assert.Equal(t, b.Params["title"], "text")
	assert.Equal(t, b.Params["photo_id"], 1)
	assert.Equal(t, b.Params["main_album"], true)
}

func TestMarketAddToAlbumBulder(t *testing.T) {
	b := params.NewMarketAddToAlbumBulder()

	b.OwnerID(1)
	b.ItemID(1)
	b.AlbumIDs([]int{1})

	assert.Equal(t, b.Params["owner_id"], 1)
	assert.Equal(t, b.Params["item_id"], 1)
	assert.Equal(t, b.Params["album_ids"], []int{1})
}

func TestMarketCreateCommentBulder(t *testing.T) {
	b := params.NewMarketCreateCommentBulder()

	b.OwnerID(1)
	b.ItemID(1)
	b.Message("text")
	b.Attachments([]string{"text"})
	b.FromGroup(true)
	b.ReplyToComment(1)
	b.StickerID(1)
	b.GUID("text")

	assert.Equal(t, b.Params["owner_id"], 1)
	assert.Equal(t, b.Params["item_id"], 1)
	assert.Equal(t, b.Params["message"], "text")
	assert.Equal(t, b.Params["attachments"], []string{"text"})
	assert.Equal(t, b.Params["from_group"], true)
	assert.Equal(t, b.Params["reply_to_comment"], 1)
	assert.Equal(t, b.Params["sticker_id"], 1)
	assert.Equal(t, b.Params["guid"], "text")
}

func TestMarketDeleteBulder(t *testing.T) {
	b := params.NewMarketDeleteBulder()

	b.OwnerID(1)
	b.ItemID(1)

	assert.Equal(t, b.Params["owner_id"], 1)
	assert.Equal(t, b.Params["item_id"], 1)
}

func TestMarketDeleteAlbumBulder(t *testing.T) {
	b := params.NewMarketDeleteAlbumBulder()

	b.OwnerID(1)
	b.AlbumID(1)

	assert.Equal(t, b.Params["owner_id"], 1)
	assert.Equal(t, b.Params["album_id"], 1)
}

func TestMarketDeleteCommentBulder(t *testing.T) {
	b := params.NewMarketDeleteCommentBulder()

	b.OwnerID(1)
	b.CommentID(1)

	assert.Equal(t, b.Params["owner_id"], 1)
	assert.Equal(t, b.Params["comment_id"], 1)
}

func TestMarketEditBulder(t *testing.T) {
	b := params.NewMarketEditBulder()

	b.OwnerID(1)
	b.ItemID(1)
	b.Name("text")
	b.Description("text")
	b.CategoryID(1)
	b.Price(1.1)
	b.Deleted(true)
	b.MainPhotoID(1)
	b.PhotoIDs([]int{1})
	b.URL("text")

	assert.Equal(t, b.Params["owner_id"], 1)
	assert.Equal(t, b.Params["item_id"], 1)
	assert.Equal(t, b.Params["name"], "text")
	assert.Equal(t, b.Params["description"], "text")
	assert.Equal(t, b.Params["category_id"], 1)
	assert.Equal(t, b.Params["price"], 1.1)
	assert.Equal(t, b.Params["deleted"], true)
	assert.Equal(t, b.Params["main_photo_id"], 1)
	assert.Equal(t, b.Params["photo_ids"], []int{1})
	assert.Equal(t, b.Params["url"], "text")
}

func TestMarketEditAlbumBulder(t *testing.T) {
	b := params.NewMarketEditAlbumBulder()

	b.OwnerID(1)
	b.AlbumID(1)
	b.Title("text")
	b.PhotoID(1)
	b.MainAlbum(true)

	assert.Equal(t, b.Params["owner_id"], 1)
	assert.Equal(t, b.Params["album_id"], 1)
	assert.Equal(t, b.Params["title"], "text")
	assert.Equal(t, b.Params["photo_id"], 1)
	assert.Equal(t, b.Params["main_album"], true)
}

func TestMarketEditCommentBulder(t *testing.T) {
	b := params.NewMarketEditCommentBulder()

	b.OwnerID(1)
	b.CommentID(1)
	b.Message("text")
	b.Attachments([]string{"text"})

	assert.Equal(t, b.Params["owner_id"], 1)
	assert.Equal(t, b.Params["comment_id"], 1)
	assert.Equal(t, b.Params["message"], "text")
	assert.Equal(t, b.Params["attachments"], []string{"text"})
}

func TestMarketGetBulder(t *testing.T) {
	b := params.NewMarketGetBulder()

	b.OwnerID(1)
	b.AlbumID(1)
	b.Count(1)
	b.Offset(1)
	b.Extended(true)

	assert.Equal(t, b.Params["owner_id"], 1)
	assert.Equal(t, b.Params["album_id"], 1)
	assert.Equal(t, b.Params["count"], 1)
	assert.Equal(t, b.Params["offset"], 1)
	assert.Equal(t, b.Params["extended"], true)
}

func TestMarketGetAlbumByIDBulder(t *testing.T) {
	b := params.NewMarketGetAlbumByIDBulder()

	b.OwnerID(1)
	b.AlbumIDs([]int{1})

	assert.Equal(t, b.Params["owner_id"], 1)
	assert.Equal(t, b.Params["album_ids"], []int{1})
}

func TestMarketGetAlbumsBulder(t *testing.T) {
	b := params.NewMarketGetAlbumsBulder()

	b.OwnerID(1)
	b.Offset(1)
	b.Count(1)

	assert.Equal(t, b.Params["owner_id"], 1)
	assert.Equal(t, b.Params["offset"], 1)
	assert.Equal(t, b.Params["count"], 1)
}

func TestMarketGetByIDBulder(t *testing.T) {
	b := params.NewMarketGetByIDBulder()

	b.ItemIDs([]string{"text"})
	b.Extended(true)

	assert.Equal(t, b.Params["item_ids"], []string{"text"})
	assert.Equal(t, b.Params["extended"], true)
}

func TestMarketGetCategoriesBulder(t *testing.T) {
	b := params.NewMarketGetCategoriesBulder()

	b.Count(1)
	b.Offset(1)

	assert.Equal(t, b.Params["count"], 1)
	assert.Equal(t, b.Params["offset"], 1)
}

func TestMarketGetCommentsBulder(t *testing.T) {
	b := params.NewMarketGetCommentsBulder()

	b.OwnerID(1)
	b.ItemID(1)
	b.NeedLikes(true)
	b.StartCommentID(1)
	b.Offset(1)
	b.Count(1)
	b.Sort("text")
	b.Extended(true)
	b.Fields([]string{"test"})

	assert.Equal(t, b.Params["owner_id"], 1)
	assert.Equal(t, b.Params["item_id"], 1)
	assert.Equal(t, b.Params["need_likes"], true)
	assert.Equal(t, b.Params["start_comment_id"], 1)
	assert.Equal(t, b.Params["offset"], 1)
	assert.Equal(t, b.Params["count"], 1)
	assert.Equal(t, b.Params["sort"], "text")
	assert.Equal(t, b.Params["extended"], true)
	assert.Equal(t, b.Params["fields"], []string{"test"})
}

func TestMarketRemoveFromAlbumBulder(t *testing.T) {
	b := params.NewMarketRemoveFromAlbumBulder()

	b.OwnerID(1)
	b.ItemID(1)
	b.AlbumIDs([]int{1})

	assert.Equal(t, b.Params["owner_id"], 1)
	assert.Equal(t, b.Params["item_id"], 1)
	assert.Equal(t, b.Params["album_ids"], []int{1})
}

func TestMarketReorderAlbumsBulder(t *testing.T) {
	b := params.NewMarketReorderAlbumsBulder()

	b.OwnerID(1)
	b.AlbumID(1)
	b.Before(1)
	b.After(1)

	assert.Equal(t, b.Params["owner_id"], 1)
	assert.Equal(t, b.Params["album_id"], 1)
	assert.Equal(t, b.Params["before"], 1)
	assert.Equal(t, b.Params["after"], 1)
}

func TestMarketReorderItemsBulder(t *testing.T) {
	b := params.NewMarketReorderItemsBulder()

	b.OwnerID(1)
	b.AlbumID(1)
	b.ItemID(1)
	b.Before(1)
	b.After(1)

	assert.Equal(t, b.Params["owner_id"], 1)
	assert.Equal(t, b.Params["album_id"], 1)
	assert.Equal(t, b.Params["item_id"], 1)
	assert.Equal(t, b.Params["before"], 1)
	assert.Equal(t, b.Params["after"], 1)
}

func TestMarketReportBulder(t *testing.T) {
	b := params.NewMarketReportBulder()

	b.OwnerID(1)
	b.ItemID(1)
	b.Reason(1)

	assert.Equal(t, b.Params["owner_id"], 1)
	assert.Equal(t, b.Params["item_id"], 1)
	assert.Equal(t, b.Params["reason"], 1)
}

func TestMarketReportCommentBulder(t *testing.T) {
	b := params.NewMarketReportCommentBulder()

	b.OwnerID(1)
	b.CommentID(1)
	b.Reason(1)

	assert.Equal(t, b.Params["owner_id"], 1)
	assert.Equal(t, b.Params["comment_id"], 1)
	assert.Equal(t, b.Params["reason"], 1)
}

func TestMarketRestoreBulder(t *testing.T) {
	b := params.NewMarketRestoreBulder()

	b.OwnerID(1)
	b.ItemID(1)

	assert.Equal(t, b.Params["owner_id"], 1)
	assert.Equal(t, b.Params["item_id"], 1)
}

func TestMarketRestoreCommentBulder(t *testing.T) {
	b := params.NewMarketRestoreCommentBulder()

	b.OwnerID(1)
	b.CommentID(1)

	assert.Equal(t, b.Params["owner_id"], 1)
	assert.Equal(t, b.Params["comment_id"], 1)
}

func TestMarketSearchBulder(t *testing.T) {
	b := params.NewMarketSearchBulder()

	b.OwnerID(1)
	b.AlbumID(1)
	b.Q("text")
	b.PriceFrom(1)
	b.PriceTo(1)
	b.Tags([]int{1})
	b.Sort(1)
	b.Rev(1)
	b.Offset(1)
	b.Count(1)
	b.Extended(true)
	b.Status(1)

	assert.Equal(t, b.Params["owner_id"], 1)
	assert.Equal(t, b.Params["album_id"], 1)
	assert.Equal(t, b.Params["q"], "text")
	assert.Equal(t, b.Params["price_from"], 1)
	assert.Equal(t, b.Params["price_to"], 1)
	assert.Equal(t, b.Params["tags"], []int{1})
	assert.Equal(t, b.Params["sort"], 1)
	assert.Equal(t, b.Params["rev"], 1)
	assert.Equal(t, b.Params["offset"], 1)
	assert.Equal(t, b.Params["count"], 1)
	assert.Equal(t, b.Params["extended"], true)
	assert.Equal(t, b.Params["status"], 1)
}
