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

	assert.Equal(t, b.Params["target_id"], 1)
	assert.Equal(t, b.Params["video_id"], 1)
	assert.Equal(t, b.Params["owner_id"], 1)
}

func TestVideoAddAlbumBuilder(t *testing.T) {
	t.Parallel()

	b := params.NewVideoAddAlbumBuilder()

	b.GroupID(1)
	b.Title("text")
	b.Privacy([]string{"text"})

	assert.Equal(t, b.Params["group_id"], 1)
	assert.Equal(t, b.Params["title"], "text")
	assert.Equal(t, b.Params["privacy"], []string{"text"})
}

func TestVideoAddToAlbumBuilder(t *testing.T) {
	t.Parallel()

	b := params.NewVideoAddToAlbumBuilder()

	b.TargetID(1)
	b.AlbumID(1)
	b.AlbumIDs([]int{1})
	b.OwnerID(1)
	b.VideoID(1)

	assert.Equal(t, b.Params["target_id"], 1)
	assert.Equal(t, b.Params["album_id"], 1)
	assert.Equal(t, b.Params["album_ids"], []int{1})
	assert.Equal(t, b.Params["owner_id"], 1)
	assert.Equal(t, b.Params["video_id"], 1)
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

	assert.Equal(t, b.Params["owner_id"], 1)
	assert.Equal(t, b.Params["video_id"], 1)
	assert.Equal(t, b.Params["message"], "text")
	assert.Equal(t, b.Params["attachments"], []string{"text"})
	assert.Equal(t, b.Params["from_group"], true)
	assert.Equal(t, b.Params["reply_to_comment"], 1)
	assert.Equal(t, b.Params["sticker_id"], 1)
	assert.Equal(t, b.Params["guid"], "text")
}

func TestVideoDeleteBuilder(t *testing.T) {
	t.Parallel()

	b := params.NewVideoDeleteBuilder()

	b.VideoID(1)
	b.OwnerID(1)
	b.TargetID(1)

	assert.Equal(t, b.Params["video_id"], 1)
	assert.Equal(t, b.Params["owner_id"], 1)
	assert.Equal(t, b.Params["target_id"], 1)
}

func TestVideoDeleteAlbumBuilder(t *testing.T) {
	t.Parallel()

	b := params.NewVideoDeleteAlbumBuilder()

	b.GroupID(1)
	b.AlbumID(1)

	assert.Equal(t, b.Params["group_id"], 1)
	assert.Equal(t, b.Params["album_id"], 1)
}

func TestVideoDeleteCommentBuilder(t *testing.T) {
	t.Parallel()

	b := params.NewVideoDeleteCommentBuilder()

	b.OwnerID(1)
	b.CommentID(1)

	assert.Equal(t, b.Params["owner_id"], 1)
	assert.Equal(t, b.Params["comment_id"], 1)
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

	assert.Equal(t, b.Params["owner_id"], 1)
	assert.Equal(t, b.Params["video_id"], 1)
	assert.Equal(t, b.Params["name"], "text")
	assert.Equal(t, b.Params["desc"], "text")
	assert.Equal(t, b.Params["privacy_view"], []string{"text"})
	assert.Equal(t, b.Params["privacy_comment"], []string{"text"})
	assert.Equal(t, b.Params["no_comments"], true)
	assert.Equal(t, b.Params["repeat"], true)
}

func TestVideoEditAlbumBuilder(t *testing.T) {
	t.Parallel()

	b := params.NewVideoEditAlbumBuilder()

	b.GroupID(1)
	b.AlbumID(1)
	b.Title("text")
	b.Privacy([]string{"text"})

	assert.Equal(t, b.Params["group_id"], 1)
	assert.Equal(t, b.Params["album_id"], 1)
	assert.Equal(t, b.Params["title"], "text")
	assert.Equal(t, b.Params["privacy"], []string{"text"})
}

func TestVideoEditCommentBuilder(t *testing.T) {
	t.Parallel()

	b := params.NewVideoEditCommentBuilder()

	b.OwnerID(1)
	b.CommentID(1)
	b.Message("text")
	b.Attachments([]string{"text"})

	assert.Equal(t, b.Params["owner_id"], 1)
	assert.Equal(t, b.Params["comment_id"], 1)
	assert.Equal(t, b.Params["message"], "text")
	assert.Equal(t, b.Params["attachments"], []string{"text"})
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

	assert.Equal(t, b.Params["owner_id"], 1)
	assert.Equal(t, b.Params["videos"], []string{"text"})
	assert.Equal(t, b.Params["album_id"], 1)
	assert.Equal(t, b.Params["count"], 1)
	assert.Equal(t, b.Params["offset"], 1)
	assert.Equal(t, b.Params["extended"], true)
}

func TestVideoGetAlbumByIDBuilder(t *testing.T) {
	t.Parallel()

	b := params.NewVideoGetAlbumByIDBuilder()

	b.OwnerID(1)
	b.AlbumID(1)

	assert.Equal(t, b.Params["owner_id"], 1)
	assert.Equal(t, b.Params["album_id"], 1)
}

func TestVideoGetAlbumsBuilder(t *testing.T) {
	t.Parallel()

	b := params.NewVideoGetAlbumsBuilder()

	b.OwnerID(1)
	b.Offset(1)
	b.Count(1)
	b.Extended(true)
	b.NeedSystem(true)

	assert.Equal(t, b.Params["owner_id"], 1)
	assert.Equal(t, b.Params["offset"], 1)
	assert.Equal(t, b.Params["count"], 1)
	assert.Equal(t, b.Params["extended"], true)
	assert.Equal(t, b.Params["need_system"], true)
}

func TestVideoGetAlbumsByVideoBuilder(t *testing.T) {
	t.Parallel()

	b := params.NewVideoGetAlbumsByVideoBuilder()

	b.TargetID(1)
	b.OwnerID(1)
	b.VideoID(1)
	b.Extended(true)

	assert.Equal(t, b.Params["target_id"], 1)
	assert.Equal(t, b.Params["owner_id"], 1)
	assert.Equal(t, b.Params["video_id"], 1)
	assert.Equal(t, b.Params["extended"], true)
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

	assert.Equal(t, b.Params["owner_id"], 1)
	assert.Equal(t, b.Params["video_id"], 1)
	assert.Equal(t, b.Params["need_likes"], true)
	assert.Equal(t, b.Params["start_comment_id"], 1)
	assert.Equal(t, b.Params["offset"], 1)
	assert.Equal(t, b.Params["count"], 1)
	assert.Equal(t, b.Params["sort"], "text")
	assert.Equal(t, b.Params["extended"], true)
	assert.Equal(t, b.Params["fields"], []string{"text"})
}

func TestVideoRemoveFromAlbumBuilder(t *testing.T) {
	t.Parallel()

	b := params.NewVideoRemoveFromAlbumBuilder()

	b.TargetID(1)
	b.AlbumID(1)
	b.AlbumIDs([]int{1})
	b.OwnerID(1)
	b.VideoID(1)

	assert.Equal(t, b.Params["target_id"], 1)
	assert.Equal(t, b.Params["album_id"], 1)
	assert.Equal(t, b.Params["album_ids"], []int{1})
	assert.Equal(t, b.Params["owner_id"], 1)
	assert.Equal(t, b.Params["video_id"], 1)
}

func TestVideoReorderAlbumsBuilder(t *testing.T) {
	t.Parallel()

	b := params.NewVideoReorderAlbumsBuilder()

	b.OwnerID(1)
	b.AlbumID(1)
	b.Before(1)
	b.After(1)

	assert.Equal(t, b.Params["owner_id"], 1)
	assert.Equal(t, b.Params["album_id"], 1)
	assert.Equal(t, b.Params["before"], 1)
	assert.Equal(t, b.Params["after"], 1)
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

	assert.Equal(t, b.Params["target_id"], 1)
	assert.Equal(t, b.Params["album_id"], 1)
	assert.Equal(t, b.Params["owner_id"], 1)
	assert.Equal(t, b.Params["video_id"], 1)
	assert.Equal(t, b.Params["before_owner_id"], 1)
	assert.Equal(t, b.Params["before_video_id"], 1)
	assert.Equal(t, b.Params["after_owner_id"], 1)
	assert.Equal(t, b.Params["after_video_id"], 1)
}

func TestVideoReportBuilder(t *testing.T) {
	t.Parallel()

	b := params.NewVideoReportBuilder()

	b.OwnerID(1)
	b.VideoID(1)
	b.Reason(1)
	b.Comment("text")
	b.SearchQuery("text")

	assert.Equal(t, b.Params["owner_id"], 1)
	assert.Equal(t, b.Params["video_id"], 1)
	assert.Equal(t, b.Params["reason"], 1)
	assert.Equal(t, b.Params["comment"], "text")
	assert.Equal(t, b.Params["search_query"], "text")
}

func TestVideoReportCommentBuilder(t *testing.T) {
	t.Parallel()

	b := params.NewVideoReportCommentBuilder()

	b.OwnerID(1)
	b.CommentID(1)
	b.Reason(1)

	assert.Equal(t, b.Params["owner_id"], 1)
	assert.Equal(t, b.Params["comment_id"], 1)
	assert.Equal(t, b.Params["reason"], 1)
}

func TestVideoRestoreBuilder(t *testing.T) {
	t.Parallel()

	b := params.NewVideoRestoreBuilder()

	b.VideoID(1)
	b.OwnerID(1)

	assert.Equal(t, b.Params["video_id"], 1)
	assert.Equal(t, b.Params["owner_id"], 1)
}

func TestVideoRestoreCommentBuilder(t *testing.T) {
	t.Parallel()

	b := params.NewVideoRestoreCommentBuilder()

	b.OwnerID(1)
	b.CommentID(1)

	assert.Equal(t, b.Params["owner_id"], 1)
	assert.Equal(t, b.Params["comment_id"], 1)
}

func TestVideoSaveBuilder(t *testing.T) {
	t.Parallel()

	b := params.NewVideoSaveBuilder()

	b.Name("text")
	b.Description("text")
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

	assert.Equal(t, b.Params["name"], "text")
	assert.Equal(t, b.Params["description"], "text")
	assert.Equal(t, b.Params["is_private"], true)
	assert.Equal(t, b.Params["wallpost"], true)
	assert.Equal(t, b.Params["link"], "text")
	assert.Equal(t, b.Params["group_id"], 1)
	assert.Equal(t, b.Params["album_id"], 1)
	assert.Equal(t, b.Params["privacy_view"], []string{"text"})
	assert.Equal(t, b.Params["privacy_comment"], []string{"text"})
	assert.Equal(t, b.Params["no_comments"], true)
	assert.Equal(t, b.Params["repeat"], true)
	assert.Equal(t, b.Params["compression"], true)
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

	assert.Equal(t, b.Params["q"], "text")
	assert.Equal(t, b.Params["sort"], 1)
	assert.Equal(t, b.Params["hd"], 1)
	assert.Equal(t, b.Params["adult"], true)
	assert.Equal(t, b.Params["filters"], []string{"text"})
	assert.Equal(t, b.Params["search_own"], true)
	assert.Equal(t, b.Params["offset"], 1)
	assert.Equal(t, b.Params["longer"], 1)
	assert.Equal(t, b.Params["shorter"], 1)
	assert.Equal(t, b.Params["count"], 1)
	assert.Equal(t, b.Params["extended"], true)
	assert.Equal(t, b.Params["legal"], true)
	assert.Equal(t, b.Params["legal_owner"], 1)
}
