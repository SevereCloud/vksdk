package params_test

import (
	"testing"

	"github.com/SevereCloud/vksdk/api/params"
	"github.com/stretchr/testify/assert"
)

func TestWallCloseCommentsBulder(t *testing.T) {
	b := params.NewWallCloseCommentsBulder()

	b.OwnerID(1)
	b.PostID(1)

	assert.Equal(t, b.Params["owner_id"], 1)
	assert.Equal(t, b.Params["post_id"], 1)
}

func TestWallCreateCommentBulder(t *testing.T) {
	b := params.NewWallCreateCommentBulder()

	b.OwnerID(1)
	b.PostID(1)
	b.FromGroup(1)
	b.Message("text")
	b.ReplyToComment(1)
	b.Attachments([]string{"text"})
	b.StickerID(1)
	b.GUID("text")

	assert.Equal(t, b.Params["owner_id"], 1)
	assert.Equal(t, b.Params["post_id"], 1)
	assert.Equal(t, b.Params["from_group"], 1)
	assert.Equal(t, b.Params["message"], "text")
	assert.Equal(t, b.Params["reply_to_comment"], 1)
	assert.Equal(t, b.Params["attachments"], []string{"text"})
	assert.Equal(t, b.Params["sticker_id"], 1)
	assert.Equal(t, b.Params["guid"], "text")
}

func TestWallDeleteBulder(t *testing.T) {
	b := params.NewWallDeleteBulder()

	b.OwnerID(1)
	b.PostID(1)

	assert.Equal(t, b.Params["owner_id"], 1)
	assert.Equal(t, b.Params["post_id"], 1)
}

func TestWallDeleteCommentBulder(t *testing.T) {
	b := params.NewWallDeleteCommentBulder()

	b.OwnerID(1)
	b.CommentID(1)

	assert.Equal(t, b.Params["owner_id"], 1)
	assert.Equal(t, b.Params["comment_id"], 1)
}

func TestWallEditBulder(t *testing.T) {
	b := params.NewWallEditBulder()

	b.OwnerID(1)
	b.PostID(1)
	b.FriendsOnly(true)
	b.Message("text")
	b.Attachments([]string{"text"})
	b.Services("text")
	b.Signed(true)
	b.PublishDate(1)
	b.Lat(1.1)
	b.Long(1.1)
	b.PlaceID(1)
	b.MarkAsAds(true)
	b.CloseComments(true)
	b.PosterBkgID(1)
	b.PosterBkgOwnerID(1)
	b.PosterBkgAccessHash("text")

	assert.Equal(t, b.Params["owner_id"], 1)
	assert.Equal(t, b.Params["post_id"], 1)
	assert.Equal(t, b.Params["friends_only"], true)
	assert.Equal(t, b.Params["message"], "text")
	assert.Equal(t, b.Params["attachments"], []string{"text"})
	assert.Equal(t, b.Params["services"], "text")
	assert.Equal(t, b.Params["signed"], true)
	assert.Equal(t, b.Params["publish_date"], 1)
	assert.Equal(t, b.Params["lat"], 1.1)
	assert.Equal(t, b.Params["long"], 1.1)
	assert.Equal(t, b.Params["place_id"], 1)
	assert.Equal(t, b.Params["mark_as_ads"], true)
	assert.Equal(t, b.Params["close_comments"], true)
	assert.Equal(t, b.Params["poster_bkg_id"], 1)
	assert.Equal(t, b.Params["poster_bkg_owner_id"], 1)
	assert.Equal(t, b.Params["poster_bkg_access_hash"], "text")
}

func TestWallEditAdsStealthBulder(t *testing.T) {
	b := params.NewWallEditAdsStealthBulder()

	b.OwnerID(1)
	b.PostID(1)
	b.Message("text")
	b.Attachments([]string{"text"})
	b.Signed(true)
	b.Lat(1.1)
	b.Long(1.1)
	b.PlaceID(1)
	b.LinkButton("text")
	b.LinkTitle("text")
	b.LinkImage("text")
	b.LinkVideo("text")

	assert.Equal(t, b.Params["owner_id"], 1)
	assert.Equal(t, b.Params["post_id"], 1)
	assert.Equal(t, b.Params["message"], "text")
	assert.Equal(t, b.Params["attachments"], []string{"text"})
	assert.Equal(t, b.Params["signed"], true)
	assert.Equal(t, b.Params["lat"], 1.1)
	assert.Equal(t, b.Params["long"], 1.1)
	assert.Equal(t, b.Params["place_id"], 1)
	assert.Equal(t, b.Params["link_button"], "text")
	assert.Equal(t, b.Params["link_title"], "text")
	assert.Equal(t, b.Params["link_image"], "text")
	assert.Equal(t, b.Params["link_video"], "text")
}

func TestWallEditCommentBulder(t *testing.T) {
	b := params.NewWallEditCommentBulder()

	b.OwnerID(1)
	b.CommentID(1)
	b.Message("text")
	b.Attachments([]string{"text"})

	assert.Equal(t, b.Params["owner_id"], 1)
	assert.Equal(t, b.Params["comment_id"], 1)
	assert.Equal(t, b.Params["message"], "text")
	assert.Equal(t, b.Params["attachments"], []string{"text"})
}

func TestWallGetBulder(t *testing.T) {
	b := params.NewWallGetBulder()

	b.OwnerID(1)
	b.Domain("text")
	b.Offset(1)
	b.Count(1)
	b.Filter("text")
	b.Extended(true)
	b.Fields([]string{"test"})

	assert.Equal(t, b.Params["owner_id"], 1)
	assert.Equal(t, b.Params["domain"], "text")
	assert.Equal(t, b.Params["offset"], 1)
	assert.Equal(t, b.Params["count"], 1)
	assert.Equal(t, b.Params["filter"], "text")
	assert.Equal(t, b.Params["extended"], true)
	assert.Equal(t, b.Params["fields"], []string{"test"})
}

func TestWallGetByIDBulder(t *testing.T) {
	b := params.NewWallGetByIDBulder()

	b.Posts([]string{"text"})
	b.Extended(true)
	b.CopyHistoryDepth(1)
	b.Fields([]string{"test"})

	assert.Equal(t, b.Params["posts"], []string{"text"})
	assert.Equal(t, b.Params["extended"], true)
	assert.Equal(t, b.Params["copy_history_depth"], 1)
	assert.Equal(t, b.Params["fields"], []string{"test"})
}

func TestWallGetCommentsBulder(t *testing.T) {
	b := params.NewWallGetCommentsBulder()

	b.OwnerID(1)
	b.PostID(1)
	b.NeedLikes(true)
	b.StartCommentID(1)
	b.Offset(1)
	b.Count(1)
	b.Sort("text")
	b.PreviewLength(1)
	b.Extended(true)
	b.Fields([]string{"test"})
	b.CommentID(1)
	b.ThreadItemsCount(1)

	assert.Equal(t, b.Params["owner_id"], 1)
	assert.Equal(t, b.Params["post_id"], 1)
	assert.Equal(t, b.Params["need_likes"], true)
	assert.Equal(t, b.Params["start_comment_id"], 1)
	assert.Equal(t, b.Params["offset"], 1)
	assert.Equal(t, b.Params["count"], 1)
	assert.Equal(t, b.Params["sort"], "text")
	assert.Equal(t, b.Params["preview_length"], 1)
	assert.Equal(t, b.Params["extended"], true)
	assert.Equal(t, b.Params["fields"], []string{"test"})
	assert.Equal(t, b.Params["comment_id"], 1)
	assert.Equal(t, b.Params["thread_items_count"], 1)
}

func TestWallGetRepostsBulder(t *testing.T) {
	b := params.NewWallGetRepostsBulder()

	b.OwnerID(1)
	b.PostID(1)
	b.Offset(1)
	b.Count(1)

	assert.Equal(t, b.Params["owner_id"], 1)
	assert.Equal(t, b.Params["post_id"], 1)
	assert.Equal(t, b.Params["offset"], 1)
	assert.Equal(t, b.Params["count"], 1)
}

func TestWallOpenCommentsBulder(t *testing.T) {
	b := params.NewWallOpenCommentsBulder()

	b.OwnerID(1)
	b.PostID(1)

	assert.Equal(t, b.Params["owner_id"], 1)
	assert.Equal(t, b.Params["post_id"], 1)
}

func TestWallPinBulder(t *testing.T) {
	b := params.NewWallPinBulder()

	b.OwnerID(1)
	b.PostID(1)

	assert.Equal(t, b.Params["owner_id"], 1)
	assert.Equal(t, b.Params["post_id"], 1)
}

func TestWallPostBulder(t *testing.T) {
	b := params.NewWallPostBulder()

	b.OwnerID(1)
	b.FriendsOnly(true)
	b.FromGroup(true)
	b.Message("text")
	b.Attachments([]string{"text"})
	b.Services("text")
	b.Signed(true)
	b.PublishDate(1)
	b.Lat(1.1)
	b.Long(1.1)
	b.PlaceID(1)
	b.PostID(1)
	b.GUID("text")
	b.MarkAsAds(true)
	b.CloseComments(true)
	b.MuteNotifications(true)

	assert.Equal(t, b.Params["owner_id"], 1)
	assert.Equal(t, b.Params["friends_only"], true)
	assert.Equal(t, b.Params["from_group"], true)
	assert.Equal(t, b.Params["message"], "text")
	assert.Equal(t, b.Params["attachments"], []string{"text"})
	assert.Equal(t, b.Params["services"], "text")
	assert.Equal(t, b.Params["signed"], true)
	assert.Equal(t, b.Params["publish_date"], 1)
	assert.Equal(t, b.Params["lat"], 1.1)
	assert.Equal(t, b.Params["long"], 1.1)
	assert.Equal(t, b.Params["place_id"], 1)
	assert.Equal(t, b.Params["post_id"], 1)
	assert.Equal(t, b.Params["guid"], "text")
	assert.Equal(t, b.Params["mark_as_ads"], true)
	assert.Equal(t, b.Params["close_comments"], true)
	assert.Equal(t, b.Params["mute_notifications"], true)
}

func TestWallPostAdsStealthBulder(t *testing.T) {
	b := params.NewWallPostAdsStealthBulder()

	b.OwnerID(1)
	b.Message("text")
	b.Attachments([]string{"text"})
	b.Signed(true)
	b.Lat(1.1)
	b.Long(1.1)
	b.PlaceID(1)
	b.GUID("text")
	b.LinkButton("text")
	b.LinkTitle("text")
	b.LinkImage("text")
	b.LinkVideo("text")

	assert.Equal(t, b.Params["owner_id"], 1)
	assert.Equal(t, b.Params["message"], "text")
	assert.Equal(t, b.Params["attachments"], []string{"text"})
	assert.Equal(t, b.Params["signed"], true)
	assert.Equal(t, b.Params["lat"], 1.1)
	assert.Equal(t, b.Params["long"], 1.1)
	assert.Equal(t, b.Params["place_id"], 1)
	assert.Equal(t, b.Params["guid"], "text")
	assert.Equal(t, b.Params["link_button"], "text")
	assert.Equal(t, b.Params["link_title"], "text")
	assert.Equal(t, b.Params["link_image"], "text")
	assert.Equal(t, b.Params["link_video"], "text")
}

func TestWallReportCommentBulder(t *testing.T) {
	b := params.NewWallReportCommentBulder()

	b.OwnerID(1)
	b.CommentID(1)
	b.Reason(1)

	assert.Equal(t, b.Params["owner_id"], 1)
	assert.Equal(t, b.Params["comment_id"], 1)
	assert.Equal(t, b.Params["reason"], 1)
}

func TestWallReportPostBulder(t *testing.T) {
	b := params.NewWallReportPostBulder()

	b.OwnerID(1)
	b.PostID(1)
	b.Reason(1)

	assert.Equal(t, b.Params["owner_id"], 1)
	assert.Equal(t, b.Params["post_id"], 1)
	assert.Equal(t, b.Params["reason"], 1)
}

func TestWallRepostBulder(t *testing.T) {
	b := params.NewWallRepostBulder()

	b.Object("text")
	b.Message("text")
	b.GroupID(1)
	b.MarkAsAds(true)
	b.MuteNotifications(true)

	assert.Equal(t, b.Params["object"], "text")
	assert.Equal(t, b.Params["message"], "text")
	assert.Equal(t, b.Params["group_id"], 1)
	assert.Equal(t, b.Params["mark_as_ads"], true)
	assert.Equal(t, b.Params["mute_notifications"], true)
}

func TestWallRestoreBulder(t *testing.T) {
	b := params.NewWallRestoreBulder()

	b.OwnerID(1)
	b.PostID(1)

	assert.Equal(t, b.Params["owner_id"], 1)
	assert.Equal(t, b.Params["post_id"], 1)
}

func TestWallRestoreCommentBulder(t *testing.T) {
	b := params.NewWallRestoreCommentBulder()

	b.OwnerID(1)
	b.CommentID(1)

	assert.Equal(t, b.Params["owner_id"], 1)
	assert.Equal(t, b.Params["comment_id"], 1)
}

func TestWallSearchBulder(t *testing.T) {
	b := params.NewWallSearchBulder()

	b.OwnerID(1)
	b.Domain("text")
	b.Query("text")
	b.OwnersOnly(true)
	b.Count(1)
	b.Offset(1)
	b.Extended(true)
	b.Fields([]string{"test"})

	assert.Equal(t, b.Params["owner_id"], 1)
	assert.Equal(t, b.Params["domain"], "text")
	assert.Equal(t, b.Params["query"], "text")
	assert.Equal(t, b.Params["owners_only"], true)
	assert.Equal(t, b.Params["count"], 1)
	assert.Equal(t, b.Params["offset"], 1)
	assert.Equal(t, b.Params["extended"], true)
	assert.Equal(t, b.Params["fields"], []string{"test"})
}

func TestWallUnpinBulder(t *testing.T) {
	b := params.NewWallUnpinBulder()

	b.OwnerID(1)
	b.PostID(1)

	assert.Equal(t, b.Params["owner_id"], 1)
	assert.Equal(t, b.Params["post_id"], 1)
}
