package params_test

import (
	"testing"

	"github.com/SevereCloud/vksdk/v2/api/params"
	"github.com/stretchr/testify/assert"
)

func TestVideoAddBuilder(t *testing.T) {
	t.Parallel()

	b := params.NewVideoAddBuilder()

	b.TargetID(1)
	b.VideoID(1)
	b.OwnerID(1)

	assert.Equal(t, 1, b.Params["target_id"])
	assert.Equal(t, 1, b.Params["video_id"])
	assert.Equal(t, 1, b.Params["owner_id"])
}

func TestVideoAddAlbumBuilder(t *testing.T) {
	t.Parallel()

	b := params.NewVideoAddAlbumBuilder()

	b.GroupID(1)
	b.Title("text")
	b.Privacy([]string{"text"})

	assert.Equal(t, 1, b.Params["group_id"])
	assert.Equal(t, "text", b.Params["title"])
	assert.Equal(t, []string{"text"}, b.Params["privacy"])
}

func TestVideoAddToAlbumBuilder(t *testing.T) {
	t.Parallel()

	b := params.NewVideoAddToAlbumBuilder()

	b.TargetID(1)
	b.AlbumID(1)
	b.AlbumIDs([]int{1})
	b.OwnerID(1)
	b.VideoID(1)

	assert.Equal(t, 1, b.Params["target_id"])
	assert.Equal(t, 1, b.Params["album_id"])
	assert.Equal(t, []int{1}, b.Params["album_ids"])
	assert.Equal(t, 1, b.Params["owner_id"])
	assert.Equal(t, 1, b.Params["video_id"])
}

func TestVideoCreateCommentBuilder(t *testing.T) {
	t.Parallel()

	b := params.NewVideoCreateCommentBuilder()

	b.OwnerID(1)
	b.VideoID(1)
	b.Message("text")
	b.Attachments([]string{"text"})
	b.FromGroup(true)
	b.ReplyToComment(1)
	b.StickerID(1)
	b.GUID("text")

	assert.Equal(t, 1, b.Params["owner_id"])
	assert.Equal(t, 1, b.Params["video_id"])
	assert.Equal(t, "text", b.Params["message"])
	assert.Equal(t, []string{"text"}, b.Params["attachments"])
	assert.Equal(t, true, b.Params["from_group"])
	assert.Equal(t, 1, b.Params["reply_to_comment"])
	assert.Equal(t, 1, b.Params["sticker_id"])
	assert.Equal(t, "text", b.Params["guid"])
}

func TestVideoDeleteBuilder(t *testing.T) {
	t.Parallel()

	b := params.NewVideoDeleteBuilder()

	b.VideoID(1)
	b.OwnerID(1)
	b.TargetID(1)

	assert.Equal(t, 1, b.Params["video_id"])
	assert.Equal(t, 1, b.Params["owner_id"])
	assert.Equal(t, 1, b.Params["target_id"])
}

func TestVideoDeleteAlbumBuilder(t *testing.T) {
	t.Parallel()

	b := params.NewVideoDeleteAlbumBuilder()

	b.GroupID(1)
	b.AlbumID(1)

	assert.Equal(t, 1, b.Params["group_id"])
	assert.Equal(t, 1, b.Params["album_id"])
}

func TestVideoDeleteCommentBuilder(t *testing.T) {
	t.Parallel()

	b := params.NewVideoDeleteCommentBuilder()

	b.OwnerID(1)
	b.CommentID(1)

	assert.Equal(t, 1, b.Params["owner_id"])
	assert.Equal(t, 1, b.Params["comment_id"])
}

func TestVideoEditBuilder(t *testing.T) {
	t.Parallel()

	b := params.NewVideoEditBuilder()

	b.OwnerID(1)
	b.VideoID(1)
	b.Name("text")
	b.Desc("text")
	b.PrivacyView([]string{"text"})
	b.PrivacyComment([]string{"text"})
	b.NoComments(true)
	b.Repeat(true)

	assert.Equal(t, 1, b.Params["owner_id"])
	assert.Equal(t, 1, b.Params["video_id"])
	assert.Equal(t, "text", b.Params["name"])
	assert.Equal(t, "text", b.Params["desc"])
	assert.Equal(t, []string{"text"}, b.Params["privacy_view"])
	assert.Equal(t, []string{"text"}, b.Params["privacy_comment"])
	assert.Equal(t, true, b.Params["no_comments"])
	assert.Equal(t, true, b.Params["repeat"])
}

func TestVideoEditAlbumBuilder(t *testing.T) {
	t.Parallel()

	b := params.NewVideoEditAlbumBuilder()

	b.GroupID(1)
	b.AlbumID(1)
	b.Title("text")
	b.Privacy([]string{"text"})

	assert.Equal(t, 1, b.Params["group_id"])
	assert.Equal(t, 1, b.Params["album_id"])
	assert.Equal(t, "text", b.Params["title"])
	assert.Equal(t, []string{"text"}, b.Params["privacy"])
}

func TestVideoEditCommentBuilder(t *testing.T) {
	t.Parallel()

	b := params.NewVideoEditCommentBuilder()

	b.OwnerID(1)
	b.CommentID(1)
	b.Message("text")
	b.Attachments([]string{"text"})

	assert.Equal(t, 1, b.Params["owner_id"])
	assert.Equal(t, 1, b.Params["comment_id"])
	assert.Equal(t, "text", b.Params["message"])
	assert.Equal(t, []string{"text"}, b.Params["attachments"])
}

func TestVideoGetBuilder(t *testing.T) {
	t.Parallel()

	b := params.NewVideoGetBuilder()

	b.OwnerID(1)
	b.Videos([]string{"text"})
	b.AlbumID(1)
	b.Count(1)
	b.Offset(1)
	b.Extended(true)

	assert.Equal(t, 1, b.Params["owner_id"])
	assert.Equal(t, []string{"text"}, b.Params["videos"])
	assert.Equal(t, 1, b.Params["album_id"])
	assert.Equal(t, 1, b.Params["count"])
	assert.Equal(t, 1, b.Params["offset"])
	assert.Equal(t, true, b.Params["extended"])
}

func TestVideoGetAlbumByIDBuilder(t *testing.T) {
	t.Parallel()

	b := params.NewVideoGetAlbumByIDBuilder()

	b.OwnerID(1)
	b.AlbumID(1)

	assert.Equal(t, 1, b.Params["owner_id"])
	assert.Equal(t, 1, b.Params["album_id"])
}

func TestVideoGetAlbumsBuilder(t *testing.T) {
	t.Parallel()

	b := params.NewVideoGetAlbumsBuilder()

	b.OwnerID(1)
	b.Offset(1)
	b.Count(1)
	b.Extended(true)
	b.NeedSystem(true)

	assert.Equal(t, 1, b.Params["owner_id"])
	assert.Equal(t, 1, b.Params["offset"])
	assert.Equal(t, 1, b.Params["count"])
	assert.Equal(t, true, b.Params["extended"])
	assert.Equal(t, true, b.Params["need_system"])
}

func TestVideoGetAlbumsByVideoBuilder(t *testing.T) {
	t.Parallel()

	b := params.NewVideoGetAlbumsByVideoBuilder()

	b.TargetID(1)
	b.OwnerID(1)
	b.VideoID(1)
	b.Extended(true)

	assert.Equal(t, 1, b.Params["target_id"])
	assert.Equal(t, 1, b.Params["owner_id"])
	assert.Equal(t, 1, b.Params["video_id"])
	assert.Equal(t, true, b.Params["extended"])
}

func TestVideoGetCommentsBuilder(t *testing.T) {
	t.Parallel()

	b := params.NewVideoGetCommentsBuilder()

	b.OwnerID(1)
	b.VideoID(1)
	b.NeedLikes(true)
	b.StartCommentID(1)
	b.Offset(1)
	b.Count(1)
	b.Sort("text")
	b.Extended(true)
	b.Fields([]string{"text"})

	assert.Equal(t, 1, b.Params["owner_id"])
	assert.Equal(t, 1, b.Params["video_id"])
	assert.Equal(t, true, b.Params["need_likes"])
	assert.Equal(t, 1, b.Params["start_comment_id"])
	assert.Equal(t, 1, b.Params["offset"])
	assert.Equal(t, 1, b.Params["count"])
	assert.Equal(t, "text", b.Params["sort"])
	assert.Equal(t, true, b.Params["extended"])
	assert.Equal(t, []string{"text"}, b.Params["fields"])
}

func TestVideoRemoveFromAlbumBuilder(t *testing.T) {
	t.Parallel()

	b := params.NewVideoRemoveFromAlbumBuilder()

	b.TargetID(1)
	b.AlbumID(1)
	b.AlbumIDs([]int{1})
	b.OwnerID(1)
	b.VideoID(1)

	assert.Equal(t, 1, b.Params["target_id"])
	assert.Equal(t, 1, b.Params["album_id"])
	assert.Equal(t, []int{1}, b.Params["album_ids"])
	assert.Equal(t, 1, b.Params["owner_id"])
	assert.Equal(t, 1, b.Params["video_id"])
}

func TestVideoReorderAlbumsBuilder(t *testing.T) {
	t.Parallel()

	b := params.NewVideoReorderAlbumsBuilder()

	b.OwnerID(1)
	b.AlbumID(1)
	b.Before(1)
	b.After(1)

	assert.Equal(t, 1, b.Params["owner_id"])
	assert.Equal(t, 1, b.Params["album_id"])
	assert.Equal(t, 1, b.Params["before"])
	assert.Equal(t, 1, b.Params["after"])
}

func TestVideoReorderVideosBuilder(t *testing.T) {
	t.Parallel()

	b := params.NewVideoReorderVideosBuilder()

	b.TargetID(1)
	b.AlbumID(1)
	b.OwnerID(1)
	b.VideoID(1)
	b.BeforeOwnerID(1)
	b.BeforeVideoID(1)
	b.AfterOwnerID(1)
	b.AfterVideoID(1)

	assert.Equal(t, 1, b.Params["target_id"])
	assert.Equal(t, 1, b.Params["album_id"])
	assert.Equal(t, 1, b.Params["owner_id"])
	assert.Equal(t, 1, b.Params["video_id"])
	assert.Equal(t, 1, b.Params["before_owner_id"])
	assert.Equal(t, 1, b.Params["before_video_id"])
	assert.Equal(t, 1, b.Params["after_owner_id"])
	assert.Equal(t, 1, b.Params["after_video_id"])
}

func TestVideoReportBuilder(t *testing.T) {
	t.Parallel()

	b := params.NewVideoReportBuilder()

	b.OwnerID(1)
	b.VideoID(1)
	b.Reason(1)
	b.Comment("text")
	b.SearchQuery("text")

	assert.Equal(t, 1, b.Params["owner_id"])
	assert.Equal(t, 1, b.Params["video_id"])
	assert.Equal(t, 1, b.Params["reason"])
	assert.Equal(t, "text", b.Params["comment"])
	assert.Equal(t, "text", b.Params["search_query"])
}

func TestVideoReportCommentBuilder(t *testing.T) {
	t.Parallel()

	b := params.NewVideoReportCommentBuilder()

	b.OwnerID(1)
	b.CommentID(1)
	b.Reason(1)

	assert.Equal(t, 1, b.Params["owner_id"])
	assert.Equal(t, 1, b.Params["comment_id"])
	assert.Equal(t, 1, b.Params["reason"])
}

func TestVideoRestoreBuilder(t *testing.T) {
	t.Parallel()

	b := params.NewVideoRestoreBuilder()

	b.VideoID(1)
	b.OwnerID(1)

	assert.Equal(t, 1, b.Params["video_id"])
	assert.Equal(t, 1, b.Params["owner_id"])
}

func TestVideoRestoreCommentBuilder(t *testing.T) {
	t.Parallel()

	b := params.NewVideoRestoreCommentBuilder()

	b.OwnerID(1)
	b.CommentID(1)

	assert.Equal(t, 1, b.Params["owner_id"])
	assert.Equal(t, 1, b.Params["comment_id"])
}

func TestVideoSaveBuilder(t *testing.T) {
	t.Parallel()

	b := params.NewVideoSaveBuilder()

	b.Name("text")
	b.Description("text")
	b.IsUnitedVideoUpload(true)
	b.IsPrivate(true)
	b.Wallpost(true)
	b.Link("text")
	b.GroupID(1)
	b.AlbumID(1)
	b.PrivacyView([]string{"text"})
	b.PrivacyComment([]string{"text"})
	b.NoComments(true)
	b.Repeat(true)
	b.Compression(true)

	assert.Equal(t, "text", b.Params["name"])
	assert.Equal(t, "text", b.Params["description"])
	assert.Equal(t, true, b.Params["is_united_video_upload"])
	assert.Equal(t, true, b.Params["is_private"])
	assert.Equal(t, true, b.Params["wallpost"])
	assert.Equal(t, "text", b.Params["link"])
	assert.Equal(t, 1, b.Params["group_id"])
	assert.Equal(t, 1, b.Params["album_id"])
	assert.Equal(t, []string{"text"}, b.Params["privacy_view"])
	assert.Equal(t, []string{"text"}, b.Params["privacy_comment"])
	assert.Equal(t, true, b.Params["no_comments"])
	assert.Equal(t, true, b.Params["repeat"])
	assert.Equal(t, true, b.Params["compression"])
}

func TestVideoSearchBuilder(t *testing.T) {
	t.Parallel()

	b := params.NewVideoSearchBuilder()

	b.Q("text")
	b.Sort(1)
	b.HD(1)
	b.Adult(true)
	b.Filters([]string{"text"})
	b.SearchOwn(true)
	b.Offset(1)
	b.Longer(1)
	b.Shorter(1)
	b.Count(1)
	b.Extended(true)
	b.Legal(true)
	b.LegalOwner(1)

	assert.Equal(t, "text", b.Params["q"])
	assert.Equal(t, 1, b.Params["sort"])
	assert.Equal(t, 1, b.Params["hd"])
	assert.Equal(t, true, b.Params["adult"])
	assert.Equal(t, []string{"text"}, b.Params["filters"])
	assert.Equal(t, true, b.Params["search_own"])
	assert.Equal(t, 1, b.Params["offset"])
	assert.Equal(t, 1, b.Params["longer"])
	assert.Equal(t, 1, b.Params["shorter"])
	assert.Equal(t, 1, b.Params["count"])
	assert.Equal(t, true, b.Params["extended"])
	assert.Equal(t, true, b.Params["legal"])
	assert.Equal(t, 1, b.Params["legal_owner"])
}

func TestVideoStartStreamingBuilder(t *testing.T) {
	t.Parallel()

	b := params.NewVideoStartStreamingBuilder()

	b.VideoID(1)
	b.Name("text")
	b.Description("text")
	b.Wallpost(true)
	b.GroupID(1)
	b.PrivacyView([]string{"text"})
	b.PrivacyComment([]string{"text"})
	b.NoComments(true)
	b.CategoryID(1)
	b.Publish(true)

	assert.Equal(t, 1, b.Params["video_id"])
	assert.Equal(t, "text", b.Params["name"])
	assert.Equal(t, "text", b.Params["description"])
	assert.Equal(t, true, b.Params["wallpost"])
	assert.Equal(t, 1, b.Params["group_id"])
	assert.Equal(t, []string{"text"}, b.Params["privacy_view"])
	assert.Equal(t, []string{"text"}, b.Params["privacy_comment"])
	assert.Equal(t, true, b.Params["no_comments"])
	assert.Equal(t, 1, b.Params["category_id"])
	assert.Equal(t, true, b.Params["publish"])
}

func TestVideoStopStreamingBuilder(t *testing.T) {
	t.Parallel()

	b := params.NewVideoStopStreamingBuilder()

	b.GroupID(1)
	b.OwnerID(1)
	b.VideoID(1)

	assert.Equal(t, 1, b.Params["group_id"])
	assert.Equal(t, 1, b.Params["owner_id"])
	assert.Equal(t, 1, b.Params["video_id"])
}
