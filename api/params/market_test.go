package params_test

import (
	"testing"

	"github.com/SevereCloud/vksdk/v2/api/params"
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

	assert.Equal(t, 1, b.Params["owner_id"])
	assert.Equal(t, "text", b.Params["name"])
	assert.Equal(t, "text", b.Params["description"])
	assert.Equal(t, 1, b.Params["category_id"])
	assert.InEpsilon(t, 1.1, b.Params["price"], 0.1)
	assert.InEpsilon(t, 1.1, b.Params["old_price"], 0.1)
	assert.Equal(t, true, b.Params["deleted"])
	assert.Equal(t, 1, b.Params["main_photo_id"])
	assert.Equal(t, []int{1}, b.Params["photo_ids"])
	assert.Equal(t, "text", b.Params["url"])
}

func TestMarketAddAlbumBuilder(t *testing.T) {
	t.Parallel()

	b := params.NewMarketAddAlbumBuilder()

	b.OwnerID(1)
	b.Title("text")
	b.PhotoID(1)
	b.MainAlbum(true)
	b.IsHidden(true)

	assert.Equal(t, 1, b.Params["owner_id"])
	assert.Equal(t, "text", b.Params["title"])
	assert.Equal(t, 1, b.Params["photo_id"])
	assert.Equal(t, true, b.Params["main_album"])
	assert.Equal(t, true, b.Params["is_hidden"])
}

func TestMarketAddToAlbumBuilder(t *testing.T) {
	t.Parallel()

	b := params.NewMarketAddToAlbumBuilder()

	b.OwnerID(1)
	b.ItemID(1)
	b.ItemIDs(1)
	b.AlbumIDs([]int{1})

	assert.Equal(t, 1, b.Params["owner_id"])
	assert.Equal(t, 1, b.Params["item_id"])
	assert.Equal(t, 1, b.Params["item_ids"])
	assert.Equal(t, []int{1}, b.Params["album_ids"])
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

	assert.Equal(t, 1, b.Params["owner_id"])
	assert.Equal(t, 1, b.Params["item_id"])
	assert.Equal(t, "text", b.Params["message"])
	assert.Equal(t, []string{"text"}, b.Params["attachments"])
	assert.Equal(t, true, b.Params["from_group"])
	assert.Equal(t, 1, b.Params["reply_to_comment"])
	assert.Equal(t, 1, b.Params["sticker_id"])
	assert.Equal(t, "text", b.Params["guid"])
}

func TestMarketDeleteBuilder(t *testing.T) {
	t.Parallel()

	b := params.NewMarketDeleteBuilder()

	b.OwnerID(1)
	b.ItemID(1)

	assert.Equal(t, 1, b.Params["owner_id"])
	assert.Equal(t, 1, b.Params["item_id"])
}

func TestMarketDeleteAlbumBuilder(t *testing.T) {
	t.Parallel()

	b := params.NewMarketDeleteAlbumBuilder()

	b.OwnerID(1)
	b.AlbumID(1)

	assert.Equal(t, 1, b.Params["owner_id"])
	assert.Equal(t, 1, b.Params["album_id"])
}

func TestMarketDeleteCommentBuilder(t *testing.T) {
	t.Parallel()

	b := params.NewMarketDeleteCommentBuilder()

	b.OwnerID(1)
	b.CommentID(1)

	assert.Equal(t, 1, b.Params["owner_id"])
	assert.Equal(t, 1, b.Params["comment_id"])
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

	assert.Equal(t, 1, b.Params["owner_id"])
	assert.Equal(t, 1, b.Params["item_id"])
	assert.Equal(t, "text", b.Params["name"])
	assert.Equal(t, "text", b.Params["description"])
	assert.Equal(t, 1, b.Params["category_id"])
	assert.InEpsilon(t, 1.1, b.Params["price"], 0.1)
	assert.Equal(t, true, b.Params["deleted"])
	assert.Equal(t, 1, b.Params["main_photo_id"])
	assert.Equal(t, []int{1}, b.Params["photo_ids"])
	assert.Equal(t, "text", b.Params["url"])
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

	assert.Equal(t, 1, b.Params["user_id"])
	assert.Equal(t, 1, b.Params["order_id"])
	assert.Equal(t, "text", b.Params["merchant_comment"])
	assert.Equal(t, 1, b.Params["status"])
	assert.Equal(t, "text", b.Params["track_number"])
	assert.Equal(t, "text", b.Params["payment_status"])
	assert.Equal(t, 1, b.Params["delivery_price"])
	assert.Equal(t, 1, b.Params["width"])
	assert.Equal(t, 1, b.Params["length"])
	assert.Equal(t, 1, b.Params["height"])
	assert.Equal(t, 1, b.Params["weight"])
	assert.Equal(t, "text", b.Params["comment_for_user"])
	assert.Equal(t, "text", b.Params["receipt_link"])
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

	assert.Equal(t, 1, b.Params["owner_id"])
	assert.Equal(t, 1, b.Params["album_id"])
	assert.Equal(t, "text", b.Params["title"])
	assert.Equal(t, 1, b.Params["photo_id"])
	assert.Equal(t, true, b.Params["main_album"])
	assert.Equal(t, true, b.Params["is_hidden"])
}

func TestMarketEditCommentBuilder(t *testing.T) {
	t.Parallel()

	b := params.NewMarketEditCommentBuilder()

	b.OwnerID(1)
	b.CommentID(1)
	b.Message("text")
	b.Attachments([]string{"text"})

	assert.Equal(t, 1, b.Params["owner_id"])
	assert.Equal(t, 1, b.Params["comment_id"])
	assert.Equal(t, "text", b.Params["message"])
	assert.Equal(t, []string{"text"}, b.Params["attachments"])
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

	assert.Equal(t, 1, b.Params["owner_id"])
	assert.Equal(t, 1, b.Params["album_id"])
	assert.Equal(t, 1, b.Params["count"])
	assert.Equal(t, 1, b.Params["offset"])
	assert.Equal(t, true, b.Params["extended"])
	assert.Equal(t, true, b.Params["need_variants"])
	assert.Equal(t, true, b.Params["with_disabled"])
}

func TestMarketGetAlbumByIDBuilder(t *testing.T) {
	t.Parallel()

	b := params.NewMarketGetAlbumByIDBuilder()

	b.OwnerID(1)
	b.AlbumIDs([]int{1})
	b.NeedAllItemIDs(true)

	assert.Equal(t, 1, b.Params["owner_id"])
	assert.Equal(t, []int{1}, b.Params["album_ids"])
	assert.Equal(t, true, b.Params["need_all_item_ids"])
}

func TestMarketGetAlbumsBuilder(t *testing.T) {
	t.Parallel()

	b := params.NewMarketGetAlbumsBuilder()

	b.OwnerID(1)
	b.Offset(1)
	b.Count(1)

	assert.Equal(t, 1, b.Params["owner_id"])
	assert.Equal(t, 1, b.Params["offset"])
	assert.Equal(t, 1, b.Params["count"])
}

func TestMarketGetByIDBuilder(t *testing.T) {
	t.Parallel()

	b := params.NewMarketGetByIDBuilder()

	b.ItemIDs([]string{"text"})
	b.Extended(true)

	assert.Equal(t, []string{"text"}, b.Params["item_ids"])
	assert.Equal(t, true, b.Params["extended"])
}

func TestMarketGetCategoriesBuilder(t *testing.T) {
	t.Parallel()

	b := params.NewMarketGetCategoriesBuilder()

	b.Count(1)
	b.Offset(1)

	assert.Equal(t, 1, b.Params["count"])
	assert.Equal(t, 1, b.Params["offset"])
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

	assert.Equal(t, 1, b.Params["owner_id"])
	assert.Equal(t, 1, b.Params["item_id"])
	assert.Equal(t, true, b.Params["need_likes"])
	assert.Equal(t, 1, b.Params["start_comment_id"])
	assert.Equal(t, 1, b.Params["offset"])
	assert.Equal(t, 1, b.Params["count"])
	assert.Equal(t, "text", b.Params["sort"])
	assert.Equal(t, true, b.Params["extended"])
	assert.Equal(t, []string{"test"}, b.Params["fields"])
}

func TestMarketGetGroupOrdersBuilder(t *testing.T) {
	t.Parallel()

	b := params.NewMarketGetGroupOrdersBuilder()

	b.GroupID(1)
	b.Offset(1)
	b.Count(1)

	assert.Equal(t, 1, b.Params["group_id"])
	assert.Equal(t, 1, b.Params["offset"])
	assert.Equal(t, 1, b.Params["count"])
}

func TestMarketGetOrderByIDBuilder(t *testing.T) {
	t.Parallel()

	b := params.NewMarketGetOrderByIDBuilder()

	b.UserID(1)
	b.OrderID(1)
	b.Extended(true)

	assert.Equal(t, 1, b.Params["user_id"])
	assert.Equal(t, 1, b.Params["order_id"])
	assert.Equal(t, true, b.Params["extended"])
}

func TestMarketGetOrderItemsBuilder(t *testing.T) {
	t.Parallel()

	b := params.NewMarketGetOrderItemsBuilder()

	b.OrderID(1)
	b.Offset(1)
	b.Count(1)

	assert.Equal(t, 1, b.Params["order_id"])
	assert.Equal(t, 1, b.Params["offset"])
	assert.Equal(t, 1, b.Params["count"])
}

func TestMarketRemoveFromAlbumBuilder(t *testing.T) {
	t.Parallel()

	b := params.NewMarketRemoveFromAlbumBuilder()

	b.OwnerID(1)
	b.ItemID(1)
	b.AlbumIDs([]int{1})

	assert.Equal(t, 1, b.Params["owner_id"])
	assert.Equal(t, 1, b.Params["item_id"])
	assert.Equal(t, []int{1}, b.Params["album_ids"])
}

func TestMarketReorderAlbumsBuilder(t *testing.T) {
	t.Parallel()

	b := params.NewMarketReorderAlbumsBuilder()

	b.OwnerID(1)
	b.AlbumID(1)
	b.Before(1)
	b.After(1)

	assert.Equal(t, 1, b.Params["owner_id"])
	assert.Equal(t, 1, b.Params["album_id"])
	assert.Equal(t, 1, b.Params["before"])
	assert.Equal(t, 1, b.Params["after"])
}

func TestMarketReorderItemsBuilder(t *testing.T) {
	t.Parallel()

	b := params.NewMarketReorderItemsBuilder()

	b.OwnerID(1)
	b.AlbumID(1)
	b.ItemID(1)
	b.Before(1)
	b.After(1)

	assert.Equal(t, 1, b.Params["owner_id"])
	assert.Equal(t, 1, b.Params["album_id"])
	assert.Equal(t, 1, b.Params["item_id"])
	assert.Equal(t, 1, b.Params["before"])
	assert.Equal(t, 1, b.Params["after"])
}

func TestMarketReportBuilder(t *testing.T) {
	t.Parallel()

	b := params.NewMarketReportBuilder()

	b.OwnerID(1)
	b.ItemID(1)
	b.Reason(1)

	assert.Equal(t, 1, b.Params["owner_id"])
	assert.Equal(t, 1, b.Params["item_id"])
	assert.Equal(t, 1, b.Params["reason"])
}

func TestMarketReportCommentBuilder(t *testing.T) {
	t.Parallel()

	b := params.NewMarketReportCommentBuilder()

	b.OwnerID(1)
	b.CommentID(1)
	b.Reason(1)

	assert.Equal(t, 1, b.Params["owner_id"])
	assert.Equal(t, 1, b.Params["comment_id"])
	assert.Equal(t, 1, b.Params["reason"])
}

func TestMarketRestoreBuilder(t *testing.T) {
	t.Parallel()

	b := params.NewMarketRestoreBuilder()

	b.OwnerID(1)
	b.ItemID(1)

	assert.Equal(t, 1, b.Params["owner_id"])
	assert.Equal(t, 1, b.Params["item_id"])
}

func TestMarketRestoreCommentBuilder(t *testing.T) {
	t.Parallel()

	b := params.NewMarketRestoreCommentBuilder()

	b.OwnerID(1)
	b.CommentID(1)

	assert.Equal(t, 1, b.Params["owner_id"])
	assert.Equal(t, 1, b.Params["comment_id"])
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

	assert.Equal(t, 1, b.Params["owner_id"])
	assert.Equal(t, 1, b.Params["album_id"])
	assert.Equal(t, "text", b.Params["q"])
	assert.Equal(t, 1, b.Params["price_from"])
	assert.Equal(t, 1, b.Params["price_to"])
	assert.Equal(t, []int{1}, b.Params["tags"])
	assert.Equal(t, 1, b.Params["sort"])
	assert.Equal(t, 1, b.Params["rev"])
	assert.Equal(t, 1, b.Params["offset"])
	assert.Equal(t, 1, b.Params["count"])
	assert.Equal(t, true, b.Params["extended"])
	assert.Equal(t, 1, b.Params["status"])
	assert.Equal(t, true, b.Params["need_variants"])
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

	assert.Equal(t, "text", b.Params["q"])
	assert.Equal(t, 1, b.Params["offset"])
	assert.Equal(t, 1, b.Params["count"])
	assert.Equal(t, 1, b.Params["category_id"])
	assert.Equal(t, 1, b.Params["price_from"])
	assert.Equal(t, 1, b.Params["price_to"])
	assert.Equal(t, 1, b.Params["sort_by"])
	assert.Equal(t, 1, b.Params["sort_direction"])
}
