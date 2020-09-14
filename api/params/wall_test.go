package params_test

import (
	"testing"

	"github.com/SevereCloud/vksdk/v2/api/params"
	"github.com/stretchr/testify/assert"
)

func TestWallCheckCopyrightLinkBuilder(t *testing.T) {
	t.Parallel()

	b := params.NewWallCheckCopyrightLinkBuilder()

	b.Link("test")

	assert.Equal(t, b.Params["link"], "test")
}

func TestWallCloseCommentsBuilder(t *testing.T) {
	t.Parallel()

	b := params.NewWallCloseCommentsBuilder()

	b.OwnerID(1)
	b.PostID(1)

	assert.Equal(t, b.Params["owner_id"], 1)
	assert.Equal(t, b.Params["post_id"], 1)
}

func TestWallCreateCommentBuilder(t *testing.T) {
	t.Parallel()

	b := params.NewWallCreateCommentBuilder()

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

func TestWallDeleteBuilder(t *testing.T) {
	t.Parallel()

	b := params.NewWallDeleteBuilder()

	b.OwnerID(1)
	b.PostID(1)

	assert.Equal(t, b.Params["owner_id"], 1)
	assert.Equal(t, b.Params["post_id"], 1)
}

func TestWallDeleteCommentBuilder(t *testing.T) {
	t.Parallel()

	b := params.NewWallDeleteCommentBuilder()

	b.OwnerID(1)
	b.CommentID(1)

	assert.Equal(t, b.Params["owner_id"], 1)
	assert.Equal(t, b.Params["comment_id"], 1)
}

func TestWallEditBuilder(t *testing.T) {
	t.Parallel()

	b := params.NewWallEditBuilder()

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

func TestWallEditAdsStealthBuilder(t *testing.T) {
	t.Parallel()

	b := params.NewWallEditAdsStealthBuilder()

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

func TestWallEditCommentBuilder(t *testing.T) {
	t.Parallel()

	b := params.NewWallEditCommentBuilder()

	b.OwnerID(1)
	b.CommentID(1)
	b.Message("text")
	b.Attachments([]string{"text"})

	assert.Equal(t, b.Params["owner_id"], 1)
	assert.Equal(t, b.Params["comment_id"], 1)
	assert.Equal(t, b.Params["message"], "text")
	assert.Equal(t, b.Params["attachments"], []string{"text"})
}

func TestWallGetBuilder(t *testing.T) {
	t.Parallel()

	b := params.NewWallGetBuilder()

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

func TestWallGetByIDBuilder(t *testing.T) {
	t.Parallel()

	b := params.NewWallGetByIDBuilder()

	b.Posts([]string{"text"})
	b.Extended(true)
	b.CopyHistoryDepth(1)
	b.Fields([]string{"test"})

	assert.Equal(t, b.Params["posts"], []string{"text"})
	assert.Equal(t, b.Params["extended"], true)
	assert.Equal(t, b.Params["copy_history_depth"], 1)
	assert.Equal(t, b.Params["fields"], []string{"test"})
}

func TestWallGetCommentsBuilder(t *testing.T) {
	t.Parallel()

	b := params.NewWallGetCommentsBuilder()

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

func TestWallGetRepostsBuilder(t *testing.T) {
	t.Parallel()

	b := params.NewWallGetRepostsBuilder()

	b.OwnerID(1)
	b.PostID(1)
	b.Offset(1)
	b.Count(1)

	assert.Equal(t, b.Params["owner_id"], 1)
	assert.Equal(t, b.Params["post_id"], 1)
	assert.Equal(t, b.Params["offset"], 1)
	assert.Equal(t, b.Params["count"], 1)
}

func TestWallOpenCommentsBuilder(t *testing.T) {
	t.Parallel()

	b := params.NewWallOpenCommentsBuilder()

	b.OwnerID(1)
	b.PostID(1)

	assert.Equal(t, b.Params["owner_id"], 1)
	assert.Equal(t, b.Params["post_id"], 1)
}

func TestWallPinBuilder(t *testing.T) {
	t.Parallel()

	b := params.NewWallPinBuilder()

	b.OwnerID(1)
	b.PostID(1)

	assert.Equal(t, b.Params["owner_id"], 1)
	assert.Equal(t, b.Params["post_id"], 1)
}

func TestWallPostBuilder(t *testing.T) {
	t.Parallel()

	b := params.NewWallPostBuilder()

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

func TestWallPostAdsStealthBuilder(t *testing.T) {
	t.Parallel()

	b := params.NewWallPostAdsStealthBuilder()

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

func TestWallReportCommentBuilder(t *testing.T) {
	t.Parallel()

	b := params.NewWallReportCommentBuilder()

	b.OwnerID(1)
	b.CommentID(1)
	b.Reason(1)

	assert.Equal(t, b.Params["owner_id"], 1)
	assert.Equal(t, b.Params["comment_id"], 1)
	assert.Equal(t, b.Params["reason"], 1)
}

func TestWallReportPostBuilder(t *testing.T) {
	t.Parallel()

	b := params.NewWallReportPostBuilder()

	b.OwnerID(1)
	b.PostID(1)
	b.Reason(1)

	assert.Equal(t, b.Params["owner_id"], 1)
	assert.Equal(t, b.Params["post_id"], 1)
	assert.Equal(t, b.Params["reason"], 1)
}

func TestWallRepostBuilder(t *testing.T) {
	t.Parallel()

	b := params.NewWallRepostBuilder()

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

func TestWallRestoreBuilder(t *testing.T) {
	t.Parallel()

	b := params.NewWallRestoreBuilder()

	b.OwnerID(1)
	b.PostID(1)

	assert.Equal(t, b.Params["owner_id"], 1)
	assert.Equal(t, b.Params["post_id"], 1)
}

func TestWallRestoreCommentBuilder(t *testing.T) {
	t.Parallel()

	b := params.NewWallRestoreCommentBuilder()

	b.OwnerID(1)
	b.CommentID(1)

	assert.Equal(t, b.Params["owner_id"], 1)
	assert.Equal(t, b.Params["comment_id"], 1)
}

func TestWallSearchBuilder(t *testing.T) {
	t.Parallel()

	b := params.NewWallSearchBuilder()

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

func TestWallUnpinBuilder(t *testing.T) {
	t.Parallel()

	b := params.NewWallUnpinBuilder()

	b.OwnerID(1)
	b.PostID(1)

	assert.Equal(t, b.Params["owner_id"], 1)
	assert.Equal(t, b.Params["post_id"], 1)
}
