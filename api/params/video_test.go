package params_test

import (
	"testing"

	"github.com/SevereCloud/vksdk/api/params"
	"github.com/stretchr/testify/assert"
)

func TestVideoAddBulder(t *testing.T) {
	b := params.NewVideoAddBulder()

	b.TargetID(1)
	b.VideoID(1)
	b.OwnerID(1)

	assert.Equal(t, b.Params["target_id"], 1)
	assert.Equal(t, b.Params["video_id"], 1)
	assert.Equal(t, b.Params["owner_id"], 1)
}

func TestVideoAddAlbumBulder(t *testing.T) {
	b := params.NewVideoAddAlbumBulder()

	b.GroupID(1)
	b.Title("text")
	b.Privacy([]string{"text"})

	assert.Equal(t, b.Params["group_id"], 1)
	assert.Equal(t, b.Params["title"], "text")
	assert.Equal(t, b.Params["privacy"], []string{"text"})
}

func TestVideoAddToAlbumBulder(t *testing.T) {
	b := params.NewVideoAddToAlbumBulder()

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

func TestVideoCreateCommentBulder(t *testing.T) {
	b := params.NewVideoCreateCommentBulder()

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

func TestVideoDeleteBulder(t *testing.T) {
	b := params.NewVideoDeleteBulder()

	b.VideoID(1)
	b.OwnerID(1)
	b.TargetID(1)

	assert.Equal(t, b.Params["video_id"], 1)
	assert.Equal(t, b.Params["owner_id"], 1)
	assert.Equal(t, b.Params["target_id"], 1)
}

func TestVideoDeleteAlbumBulder(t *testing.T) {
	b := params.NewVideoDeleteAlbumBulder()

	b.GroupID(1)
	b.AlbumID(1)

	assert.Equal(t, b.Params["group_id"], 1)
	assert.Equal(t, b.Params["album_id"], 1)
}

func TestVideoDeleteCommentBulder(t *testing.T) {
	b := params.NewVideoDeleteCommentBulder()

	b.OwnerID(1)
	b.CommentID(1)

	assert.Equal(t, b.Params["owner_id"], 1)
	assert.Equal(t, b.Params["comment_id"], 1)
}

func TestVideoEditBulder(t *testing.T) {
	b := params.NewVideoEditBulder()

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

func TestVideoEditAlbumBulder(t *testing.T) {
	b := params.NewVideoEditAlbumBulder()

	b.GroupID(1)
	b.AlbumID(1)
	b.Title("text")
	b.Privacy([]string{"text"})

	assert.Equal(t, b.Params["group_id"], 1)
	assert.Equal(t, b.Params["album_id"], 1)
	assert.Equal(t, b.Params["title"], "text")
	assert.Equal(t, b.Params["privacy"], []string{"text"})
}

func TestVideoEditCommentBulder(t *testing.T) {
	b := params.NewVideoEditCommentBulder()

	b.OwnerID(1)
	b.CommentID(1)
	b.Message("text")
	b.Attachments([]string{"text"})

	assert.Equal(t, b.Params["owner_id"], 1)
	assert.Equal(t, b.Params["comment_id"], 1)
	assert.Equal(t, b.Params["message"], "text")
	assert.Equal(t, b.Params["attachments"], []string{"text"})
}

func TestVideoGetBulder(t *testing.T) {
	b := params.NewVideoGetBulder()

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

func TestVideoGetAlbumByIDBulder(t *testing.T) {
	b := params.NewVideoGetAlbumByIDBulder()

	b.OwnerID(1)
	b.AlbumID(1)

	assert.Equal(t, b.Params["owner_id"], 1)
	assert.Equal(t, b.Params["album_id"], 1)
}

func TestVideoGetAlbumsBulder(t *testing.T) {
	b := params.NewVideoGetAlbumsBulder()

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

func TestVideoGetAlbumsByVideoBulder(t *testing.T) {
	b := params.NewVideoGetAlbumsByVideoBulder()

	b.TargetID(1)
	b.OwnerID(1)
	b.VideoID(1)
	b.Extended(true)

	assert.Equal(t, b.Params["target_id"], 1)
	assert.Equal(t, b.Params["owner_id"], 1)
	assert.Equal(t, b.Params["video_id"], 1)
	assert.Equal(t, b.Params["extended"], true)
}

func TestVideoGetCommentsBulder(t *testing.T) {
	b := params.NewVideoGetCommentsBulder()

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

func TestVideoRemoveFromAlbumBulder(t *testing.T) {
	b := params.NewVideoRemoveFromAlbumBulder()

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

func TestVideoReorderAlbumsBulder(t *testing.T) {
	b := params.NewVideoReorderAlbumsBulder()

	b.OwnerID(1)
	b.AlbumID(1)
	b.Before(1)
	b.After(1)

	assert.Equal(t, b.Params["owner_id"], 1)
	assert.Equal(t, b.Params["album_id"], 1)
	assert.Equal(t, b.Params["before"], 1)
	assert.Equal(t, b.Params["after"], 1)
}

func TestVideoReorderVideosBulder(t *testing.T) {
	b := params.NewVideoReorderVideosBulder()

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

func TestVideoReportBulder(t *testing.T) {
	b := params.NewVideoReportBulder()

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

func TestVideoReportCommentBulder(t *testing.T) {
	b := params.NewVideoReportCommentBulder()

	b.OwnerID(1)
	b.CommentID(1)
	b.Reason(1)

	assert.Equal(t, b.Params["owner_id"], 1)
	assert.Equal(t, b.Params["comment_id"], 1)
	assert.Equal(t, b.Params["reason"], 1)
}

func TestVideoRestoreBulder(t *testing.T) {
	b := params.NewVideoRestoreBulder()

	b.VideoID(1)
	b.OwnerID(1)

	assert.Equal(t, b.Params["video_id"], 1)
	assert.Equal(t, b.Params["owner_id"], 1)
}

func TestVideoRestoreCommentBulder(t *testing.T) {
	b := params.NewVideoRestoreCommentBulder()

	b.OwnerID(1)
	b.CommentID(1)

	assert.Equal(t, b.Params["owner_id"], 1)
	assert.Equal(t, b.Params["comment_id"], 1)
}

func TestVideoSaveBulder(t *testing.T) {
	b := params.NewVideoSaveBulder()

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

func TestVideoSearchBulder(t *testing.T) {
	b := params.NewVideoSearchBulder()

	b.Q("text")
	b.Sort(1)
	b.Hd(1)
	b.Adult(true)
	b.Filters([]string{"text"})
	b.SearchOwn(true)
	b.Offset(1)
	b.Longer(1)
	b.Shorter(1)
	b.Count(1)
	b.Extended(true)

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
}
