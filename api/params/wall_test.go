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

	assert.Equal(t, "test", b.Params["link"])
}

func TestWallCloseCommentsBuilder(t *testing.T) {
	t.Parallel()

	b := params.NewWallCloseCommentsBuilder()

	b.OwnerID(1)
	b.PostID(1)

	assert.Equal(t, 1, b.Params["owner_id"])
	assert.Equal(t, 1, b.Params["post_id"])
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

	assert.Equal(t, 1, b.Params["owner_id"])
	assert.Equal(t, 1, b.Params["post_id"])
	assert.Equal(t, 1, b.Params["from_group"])
	assert.Equal(t, "text", b.Params["message"])
	assert.Equal(t, 1, b.Params["reply_to_comment"])
	assert.Equal(t, []string{"text"}, b.Params["attachments"])
	assert.Equal(t, 1, b.Params["sticker_id"])
	assert.Equal(t, "text", b.Params["guid"])
}

func TestWallDeleteBuilder(t *testing.T) {
	t.Parallel()

	b := params.NewWallDeleteBuilder()

	b.OwnerID(1)
	b.PostID(1)

	assert.Equal(t, 1, b.Params["owner_id"])
	assert.Equal(t, 1, b.Params["post_id"])
}

func TestWallDeleteCommentBuilder(t *testing.T) {
	t.Parallel()

	b := params.NewWallDeleteCommentBuilder()

	b.OwnerID(1)
	b.CommentID(1)

	assert.Equal(t, 1, b.Params["owner_id"])
	assert.Equal(t, 1, b.Params["comment_id"])
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

	assert.Equal(t, 1, b.Params["owner_id"])
	assert.Equal(t, 1, b.Params["post_id"])
	assert.Equal(t, true, b.Params["friends_only"])
	assert.Equal(t, "text", b.Params["message"])
	assert.Equal(t, []string{"text"}, b.Params["attachments"])
	assert.Equal(t, "text", b.Params["services"])
	assert.Equal(t, true, b.Params["signed"])
	assert.Equal(t, 1, b.Params["publish_date"])
	assert.InEpsilon(t, 1.1, b.Params["lat"], 0.1)
	assert.InEpsilon(t, 1.1, b.Params["long"], 0.1)
	assert.Equal(t, 1, b.Params["place_id"])
	assert.Equal(t, true, b.Params["mark_as_ads"])
	assert.Equal(t, true, b.Params["close_comments"])
	assert.Equal(t, 1, b.Params["poster_bkg_id"])
	assert.Equal(t, 1, b.Params["poster_bkg_owner_id"])
	assert.Equal(t, "text", b.Params["poster_bkg_access_hash"])
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

	assert.Equal(t, 1, b.Params["owner_id"])
	assert.Equal(t, 1, b.Params["post_id"])
	assert.Equal(t, "text", b.Params["message"])
	assert.Equal(t, []string{"text"}, b.Params["attachments"])
	assert.Equal(t, true, b.Params["signed"])
	assert.InEpsilon(t, 1.1, b.Params["lat"], 0.1)
	assert.InEpsilon(t, 1.1, b.Params["long"], 0.1)
	assert.Equal(t, 1, b.Params["place_id"])
	assert.Equal(t, "text", b.Params["link_button"])
	assert.Equal(t, "text", b.Params["link_title"])
	assert.Equal(t, "text", b.Params["link_image"])
	assert.Equal(t, "text", b.Params["link_video"])
}

func TestWallEditCommentBuilder(t *testing.T) {
	t.Parallel()

	b := params.NewWallEditCommentBuilder()

	b.OwnerID(1)
	b.CommentID(1)
	b.Message("text")
	b.Attachments([]string{"text"})

	assert.Equal(t, 1, b.Params["owner_id"])
	assert.Equal(t, 1, b.Params["comment_id"])
	assert.Equal(t, "text", b.Params["message"])
	assert.Equal(t, []string{"text"}, b.Params["attachments"])
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

	assert.Equal(t, 1, b.Params["owner_id"])
	assert.Equal(t, "text", b.Params["domain"])
	assert.Equal(t, 1, b.Params["offset"])
	assert.Equal(t, 1, b.Params["count"])
	assert.Equal(t, "text", b.Params["filter"])
	assert.Equal(t, true, b.Params["extended"])
	assert.Equal(t, []string{"test"}, b.Params["fields"])
}

func TestWallGetByIDBuilder(t *testing.T) {
	t.Parallel()

	b := params.NewWallGetByIDBuilder()

	b.Posts([]string{"text"})
	b.Extended(true)
	b.CopyHistoryDepth(1)
	b.Fields([]string{"test"})

	assert.Equal(t, []string{"text"}, b.Params["posts"])
	assert.Equal(t, true, b.Params["extended"])
	assert.Equal(t, 1, b.Params["copy_history_depth"])
	assert.Equal(t, []string{"test"}, b.Params["fields"])
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

	assert.Equal(t, 1, b.Params["owner_id"])
	assert.Equal(t, 1, b.Params["post_id"])
	assert.Equal(t, true, b.Params["need_likes"])
	assert.Equal(t, 1, b.Params["start_comment_id"])
	assert.Equal(t, 1, b.Params["offset"])
	assert.Equal(t, 1, b.Params["count"])
	assert.Equal(t, "text", b.Params["sort"])
	assert.Equal(t, 1, b.Params["preview_length"])
	assert.Equal(t, true, b.Params["extended"])
	assert.Equal(t, []string{"test"}, b.Params["fields"])
	assert.Equal(t, 1, b.Params["comment_id"])
	assert.Equal(t, 1, b.Params["thread_items_count"])
}

func TestWallGetRepostsBuilder(t *testing.T) {
	t.Parallel()

	b := params.NewWallGetRepostsBuilder()

	b.OwnerID(1)
	b.PostID(1)
	b.Offset(1)
	b.Count(1)

	assert.Equal(t, 1, b.Params["owner_id"])
	assert.Equal(t, 1, b.Params["post_id"])
	assert.Equal(t, 1, b.Params["offset"])
	assert.Equal(t, 1, b.Params["count"])
}

func TestWallOpenCommentsBuilder(t *testing.T) {
	t.Parallel()

	b := params.NewWallOpenCommentsBuilder()

	b.OwnerID(1)
	b.PostID(1)

	assert.Equal(t, 1, b.Params["owner_id"])
	assert.Equal(t, 1, b.Params["post_id"])
}

func TestWallPinBuilder(t *testing.T) {
	t.Parallel()

	b := params.NewWallPinBuilder()

	b.OwnerID(1)
	b.PostID(1)

	assert.Equal(t, 1, b.Params["owner_id"])
	assert.Equal(t, 1, b.Params["post_id"])
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
	b.DonutPaidDuration(1)
	b.MuteNotifications(true)
	b.Copyright("test")
	b.TopicID(1)

	assert.Equal(t, 1, b.Params["owner_id"])
	assert.Equal(t, true, b.Params["friends_only"])
	assert.Equal(t, true, b.Params["from_group"])
	assert.Equal(t, "text", b.Params["message"])
	assert.Equal(t, []string{"text"}, b.Params["attachments"])
	assert.Equal(t, "text", b.Params["services"])
	assert.Equal(t, true, b.Params["signed"])
	assert.Equal(t, 1, b.Params["publish_date"])
	assert.InEpsilon(t, 1.1, b.Params["lat"], 0.1)
	assert.InEpsilon(t, 1.1, b.Params["long"], 0.1)
	assert.Equal(t, 1, b.Params["place_id"])
	assert.Equal(t, 1, b.Params["post_id"])
	assert.Equal(t, "text", b.Params["guid"])
	assert.Equal(t, true, b.Params["mark_as_ads"])
	assert.Equal(t, true, b.Params["close_comments"])
	assert.Equal(t, 1, b.Params["donut_paid_duration"])
	assert.Equal(t, true, b.Params["mute_notifications"])
	assert.Equal(t, "test", b.Params["copyright"])
	assert.Equal(t, 1, b.Params["topic_id"])
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

	assert.Equal(t, 1, b.Params["owner_id"])
	assert.Equal(t, "text", b.Params["message"])
	assert.Equal(t, []string{"text"}, b.Params["attachments"])
	assert.Equal(t, true, b.Params["signed"])
	assert.InEpsilon(t, 1.1, b.Params["lat"], 0.1)
	assert.InEpsilon(t, 1.1, b.Params["long"], 0.1)
	assert.Equal(t, 1, b.Params["place_id"])
	assert.Equal(t, "text", b.Params["guid"])
	assert.Equal(t, "text", b.Params["link_button"])
	assert.Equal(t, "text", b.Params["link_title"])
	assert.Equal(t, "text", b.Params["link_image"])
	assert.Equal(t, "text", b.Params["link_video"])
}

func TestWallReportCommentBuilder(t *testing.T) {
	t.Parallel()

	b := params.NewWallReportCommentBuilder()

	b.OwnerID(1)
	b.CommentID(1)
	b.Reason(1)

	assert.Equal(t, 1, b.Params["owner_id"])
	assert.Equal(t, 1, b.Params["comment_id"])
	assert.Equal(t, 1, b.Params["reason"])
}

func TestWallReportPostBuilder(t *testing.T) {
	t.Parallel()

	b := params.NewWallReportPostBuilder()

	b.OwnerID(1)
	b.PostID(1)
	b.Reason(1)

	assert.Equal(t, 1, b.Params["owner_id"])
	assert.Equal(t, 1, b.Params["post_id"])
	assert.Equal(t, 1, b.Params["reason"])
}

func TestWallRepostBuilder(t *testing.T) {
	t.Parallel()

	b := params.NewWallRepostBuilder()

	b.Object("text")
	b.Message("text")
	b.GroupID(1)
	b.MarkAsAds(true)
	b.MuteNotifications(true)

	assert.Equal(t, "text", b.Params["object"])
	assert.Equal(t, "text", b.Params["message"])
	assert.Equal(t, 1, b.Params["group_id"])
	assert.Equal(t, true, b.Params["mark_as_ads"])
	assert.Equal(t, true, b.Params["mute_notifications"])
}

func TestWallRestoreBuilder(t *testing.T) {
	t.Parallel()

	b := params.NewWallRestoreBuilder()

	b.OwnerID(1)
	b.PostID(1)

	assert.Equal(t, 1, b.Params["owner_id"])
	assert.Equal(t, 1, b.Params["post_id"])
}

func TestWallRestoreCommentBuilder(t *testing.T) {
	t.Parallel()

	b := params.NewWallRestoreCommentBuilder()

	b.OwnerID(1)
	b.CommentID(1)

	assert.Equal(t, 1, b.Params["owner_id"])
	assert.Equal(t, 1, b.Params["comment_id"])
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

	assert.Equal(t, 1, b.Params["owner_id"])
	assert.Equal(t, "text", b.Params["domain"])
	assert.Equal(t, "text", b.Params["query"])
	assert.Equal(t, true, b.Params["owners_only"])
	assert.Equal(t, 1, b.Params["count"])
	assert.Equal(t, 1, b.Params["offset"])
	assert.Equal(t, true, b.Params["extended"])
	assert.Equal(t, []string{"test"}, b.Params["fields"])
}

func TestWallUnpinBuilder(t *testing.T) {
	t.Parallel()

	b := params.NewWallUnpinBuilder()

	b.OwnerID(1)
	b.PostID(1)

	assert.Equal(t, 1, b.Params["owner_id"])
	assert.Equal(t, 1, b.Params["post_id"])
}
