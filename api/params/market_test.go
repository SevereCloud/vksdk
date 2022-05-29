package params_test

import (
	"testing"

	"github.com/Derad6709/vksdk/v2/api/params"
	"github.com/stretchr/testify/assert"
)

func TestMarketAddBuilder(t *testing.T) {
	t.Parallel()

	b := params.NewMarketAddBuilder()

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

func TestMarketAddAlbumBuilder(t *testing.T) {
	t.Parallel()

	b := params.NewMarketAddAlbumBuilder()

	b.OwnerID(1)
	b.Title("text")
	b.PhotoID(1)
	b.MainAlbum(true)
	b.IsHidden(true)

	assert.Equal(t, b.Params["owner_id"], 1)
	assert.Equal(t, b.Params["title"], "text")
	assert.Equal(t, b.Params["photo_id"], 1)
	assert.Equal(t, b.Params["main_album"], true)
	assert.Equal(t, b.Params["is_hidden"], true)
}

func TestMarketAddToAlbumBuilder(t *testing.T) {
	t.Parallel()

	b := params.NewMarketAddToAlbumBuilder()

	b.OwnerID(1)
	b.ItemID(1)
	b.ItemIDs(1)
	b.AlbumIDs([]int{1})

	assert.Equal(t, b.Params["owner_id"], 1)
	assert.Equal(t, b.Params["item_id"], 1)
	assert.Equal(t, b.Params["item_ids"], 1)
	assert.Equal(t, b.Params["album_ids"], []int{1})
}

func TestMarketCreateCommentBuilder(t *testing.T) {
	t.Parallel()

	b := params.NewMarketCreateCommentBuilder()

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

func TestMarketDeleteBuilder(t *testing.T) {
	t.Parallel()

	b := params.NewMarketDeleteBuilder()

	b.OwnerID(1)
	b.ItemID(1)

	assert.Equal(t, b.Params["owner_id"], 1)
	assert.Equal(t, b.Params["item_id"], 1)
}

func TestMarketDeleteAlbumBuilder(t *testing.T) {
	t.Parallel()

	b := params.NewMarketDeleteAlbumBuilder()

	b.OwnerID(1)
	b.AlbumID(1)

	assert.Equal(t, b.Params["owner_id"], 1)
	assert.Equal(t, b.Params["album_id"], 1)
}

func TestMarketDeleteCommentBuilder(t *testing.T) {
	t.Parallel()

	b := params.NewMarketDeleteCommentBuilder()

	b.OwnerID(1)
	b.CommentID(1)

	assert.Equal(t, b.Params["owner_id"], 1)
	assert.Equal(t, b.Params["comment_id"], 1)
}

func TestMarketEditBuilder(t *testing.T) {
	t.Parallel()

	b := params.NewMarketEditBuilder()

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

func TestMarketEditOrderBuilder(t *testing.T) {
	t.Parallel()

	b := params.NewMarketEditOrderBuilder()

	b.UserID(1)
	b.OrderID(1)
	b.MerchantComment("text")
	b.Status(1)
	b.TrackNumber("text")
	b.PaymentStatus("text")
	b.DeliveryPrice(1)
	b.Width(1)
	b.Length(1)
	b.Height(1)
	b.Weight(1)
	b.CommentForUser("text")
	b.ReceiptLink("text")

	assert.Equal(t, b.Params["user_id"], 1)
	assert.Equal(t, b.Params["order_id"], 1)
	assert.Equal(t, b.Params["merchant_comment"], "text")
	assert.Equal(t, b.Params["status"], 1)
	assert.Equal(t, b.Params["track_number"], "text")
	assert.Equal(t, b.Params["payment_status"], "text")
	assert.Equal(t, b.Params["delivery_price"], 1)
	assert.Equal(t, b.Params["width"], 1)
	assert.Equal(t, b.Params["length"], 1)
	assert.Equal(t, b.Params["height"], 1)
	assert.Equal(t, b.Params["weight"], 1)
	assert.Equal(t, b.Params["comment_for_user"], "text")
	assert.Equal(t, b.Params["receipt_link"], "text")
}

func TestMarketEditAlbumBuilder(t *testing.T) {
	t.Parallel()

	b := params.NewMarketEditAlbumBuilder()

	b.OwnerID(1)
	b.AlbumID(1)
	b.Title("text")
	b.PhotoID(1)
	b.MainAlbum(true)
	b.IsHidden(true)

	assert.Equal(t, b.Params["owner_id"], 1)
	assert.Equal(t, b.Params["album_id"], 1)
	assert.Equal(t, b.Params["title"], "text")
	assert.Equal(t, b.Params["photo_id"], 1)
	assert.Equal(t, b.Params["main_album"], true)
	assert.Equal(t, b.Params["is_hidden"], true)
}

func TestMarketEditCommentBuilder(t *testing.T) {
	t.Parallel()

	b := params.NewMarketEditCommentBuilder()

	b.OwnerID(1)
	b.CommentID(1)
	b.Message("text")
	b.Attachments([]string{"text"})

	assert.Equal(t, b.Params["owner_id"], 1)
	assert.Equal(t, b.Params["comment_id"], 1)
	assert.Equal(t, b.Params["message"], "text")
	assert.Equal(t, b.Params["attachments"], []string{"text"})
}

func TestMarketGetBuilder(t *testing.T) {
	t.Parallel()

	b := params.NewMarketGetBuilder()

	b.OwnerID(1)
	b.AlbumID(1)
	b.Count(1)
	b.Offset(1)
	b.Extended(true)
	b.NeedVariants(true)
	b.WithDisabled(true)

	assert.Equal(t, b.Params["owner_id"], 1)
	assert.Equal(t, b.Params["album_id"], 1)
	assert.Equal(t, b.Params["count"], 1)
	assert.Equal(t, b.Params["offset"], 1)
	assert.Equal(t, b.Params["extended"], true)
	assert.Equal(t, b.Params["need_variants"], true)
	assert.Equal(t, b.Params["with_disabled"], true)
}

func TestMarketGetAlbumByIDBuilder(t *testing.T) {
	t.Parallel()

	b := params.NewMarketGetAlbumByIDBuilder()

	b.OwnerID(1)
	b.AlbumIDs([]int{1})
	b.NeedAllItemIDs(true)

	assert.Equal(t, b.Params["owner_id"], 1)
	assert.Equal(t, b.Params["album_ids"], []int{1})
	assert.Equal(t, b.Params["need_all_item_ids"], true)
}

func TestMarketGetAlbumsBuilder(t *testing.T) {
	t.Parallel()

	b := params.NewMarketGetAlbumsBuilder()

	b.OwnerID(1)
	b.Offset(1)
	b.Count(1)

	assert.Equal(t, b.Params["owner_id"], 1)
	assert.Equal(t, b.Params["offset"], 1)
	assert.Equal(t, b.Params["count"], 1)
}

func TestMarketGetByIDBuilder(t *testing.T) {
	t.Parallel()

	b := params.NewMarketGetByIDBuilder()

	b.ItemIDs([]string{"text"})
	b.Extended(true)

	assert.Equal(t, b.Params["item_ids"], []string{"text"})
	assert.Equal(t, b.Params["extended"], true)
}

func TestMarketGetCategoriesBuilder(t *testing.T) {
	t.Parallel()

	b := params.NewMarketGetCategoriesBuilder()

	b.Count(1)
	b.Offset(1)

	assert.Equal(t, b.Params["count"], 1)
	assert.Equal(t, b.Params["offset"], 1)
}

func TestMarketGetCommentsBuilder(t *testing.T) {
	t.Parallel()

	b := params.NewMarketGetCommentsBuilder()

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

func TestMarketGetGroupOrdersBuilder(t *testing.T) {
	t.Parallel()

	b := params.NewMarketGetGroupOrdersBuilder()

	b.GroupID(1)
	b.Offset(1)
	b.Count(1)

	assert.Equal(t, b.Params["group_id"], 1)
	assert.Equal(t, b.Params["offset"], 1)
	assert.Equal(t, b.Params["count"], 1)
}

func TestMarketGetOrderByIDBuilder(t *testing.T) {
	t.Parallel()

	b := params.NewMarketGetOrderByIDBuilder()

	b.UserID(1)
	b.OrderID(1)
	b.Extended(true)

	assert.Equal(t, b.Params["user_id"], 1)
	assert.Equal(t, b.Params["order_id"], 1)
	assert.Equal(t, b.Params["extended"], true)
}

func TestMarketGetOrderItemsBuilder(t *testing.T) {
	t.Parallel()

	b := params.NewMarketGetOrderItemsBuilder()

	b.OrderID(1)
	b.Offset(1)
	b.Count(1)

	assert.Equal(t, b.Params["order_id"], 1)
	assert.Equal(t, b.Params["offset"], 1)
	assert.Equal(t, b.Params["count"], 1)
}

func TestMarketRemoveFromAlbumBuilder(t *testing.T) {
	t.Parallel()

	b := params.NewMarketRemoveFromAlbumBuilder()

	b.OwnerID(1)
	b.ItemID(1)
	b.AlbumIDs([]int{1})

	assert.Equal(t, b.Params["owner_id"], 1)
	assert.Equal(t, b.Params["item_id"], 1)
	assert.Equal(t, b.Params["album_ids"], []int{1})
}

func TestMarketReorderAlbumsBuilder(t *testing.T) {
	t.Parallel()

	b := params.NewMarketReorderAlbumsBuilder()

	b.OwnerID(1)
	b.AlbumID(1)
	b.Before(1)
	b.After(1)

	assert.Equal(t, b.Params["owner_id"], 1)
	assert.Equal(t, b.Params["album_id"], 1)
	assert.Equal(t, b.Params["before"], 1)
	assert.Equal(t, b.Params["after"], 1)
}

func TestMarketReorderItemsBuilder(t *testing.T) {
	t.Parallel()

	b := params.NewMarketReorderItemsBuilder()

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

func TestMarketReportBuilder(t *testing.T) {
	t.Parallel()

	b := params.NewMarketReportBuilder()

	b.OwnerID(1)
	b.ItemID(1)
	b.Reason(1)

	assert.Equal(t, b.Params["owner_id"], 1)
	assert.Equal(t, b.Params["item_id"], 1)
	assert.Equal(t, b.Params["reason"], 1)
}

func TestMarketReportCommentBuilder(t *testing.T) {
	t.Parallel()

	b := params.NewMarketReportCommentBuilder()

	b.OwnerID(1)
	b.CommentID(1)
	b.Reason(1)

	assert.Equal(t, b.Params["owner_id"], 1)
	assert.Equal(t, b.Params["comment_id"], 1)
	assert.Equal(t, b.Params["reason"], 1)
}

func TestMarketRestoreBuilder(t *testing.T) {
	t.Parallel()

	b := params.NewMarketRestoreBuilder()

	b.OwnerID(1)
	b.ItemID(1)

	assert.Equal(t, b.Params["owner_id"], 1)
	assert.Equal(t, b.Params["item_id"], 1)
}

func TestMarketRestoreCommentBuilder(t *testing.T) {
	t.Parallel()

	b := params.NewMarketRestoreCommentBuilder()

	b.OwnerID(1)
	b.CommentID(1)

	assert.Equal(t, b.Params["owner_id"], 1)
	assert.Equal(t, b.Params["comment_id"], 1)
}

func TestMarketSearchBuilder(t *testing.T) {
	t.Parallel()

	b := params.NewMarketSearchBuilder()

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
	b.NeedVariants(true)

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
	assert.Equal(t, b.Params["need_variants"], true)
}

func TestMarketSearchItemsBuilder(t *testing.T) {
	t.Parallel()

	b := params.NewMarketSearchItemsBuilder()

	b.Q("text")
	b.Offset(1)
	b.Count(1)
	b.CategoryID(1)
	b.PriceFrom(1)
	b.PriceTo(1)
	b.SortBy(1)
	b.SortDirection(1)

	assert.Equal(t, b.Params["q"], "text")
	assert.Equal(t, b.Params["offset"], 1)
	assert.Equal(t, b.Params["count"], 1)
	assert.Equal(t, b.Params["category_id"], 1)
	assert.Equal(t, b.Params["price_from"], 1)
	assert.Equal(t, b.Params["price_to"], 1)
	assert.Equal(t, b.Params["sort_by"], 1)
	assert.Equal(t, b.Params["sort_direction"], 1)
}
