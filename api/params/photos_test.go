package params_test

import (
	"testing"

	"github.com/SevereCloud/vksdk/v3/api/params"
	"github.com/stretchr/testify/assert"
)

func TestPhotosConfirmTagBuilder(t *testing.T) {
	t.Parallel()

	b := params.NewPhotosConfirmTagBuilder()

	b.OwnerID(1)
	b.PhotoID("text")
	b.TagID(1)

	assert.Equal(t, 1, b.Params["owner_id"])
	assert.Equal(t, "text", b.Params["photo_id"])
	assert.Equal(t, 1, b.Params["tag_id"])
}

func TestPhotosCopyBuilder(t *testing.T) {
	t.Parallel()

	b := params.NewPhotosCopyBuilder()

	b.OwnerID(1)
	b.PhotoID(1)
	b.AccessKey("text")

	assert.Equal(t, 1, b.Params["owner_id"])
	assert.Equal(t, 1, b.Params["photo_id"])
	assert.Equal(t, "text", b.Params["access_key"])
}

func TestPhotosCreateAlbumBuilder(t *testing.T) {
	t.Parallel()

	b := params.NewPhotosCreateAlbumBuilder()

	b.Title("text")
	b.GroupID(1)
	b.Description("text")
	b.PrivacyView([]string{"text"})
	b.PrivacyComment([]string{"text"})
	b.UploadByAdminsOnly(true)
	b.CommentsDisabled(true)

	assert.Equal(t, "text", b.Params["title"])
	assert.Equal(t, 1, b.Params["group_id"])
	assert.Equal(t, "text", b.Params["description"])
	assert.Equal(t, []string{"text"}, b.Params["privacy_view"])
	assert.Equal(t, []string{"text"}, b.Params["privacy_comment"])
	assert.Equal(t, true, b.Params["upload_by_admins_only"])
	assert.Equal(t, true, b.Params["comments_disabled"])
}

func TestPhotosCreateCommentBuilder(t *testing.T) {
	t.Parallel()

	b := params.NewPhotosCreateCommentBuilder()

	b.OwnerID(1)
	b.PhotoID(1)
	b.Message("text")
	b.Attachments([]string{"text"})
	b.FromGroup(true)
	b.ReplyToComment(1)
	b.StickerID(1)
	b.AccessKey("text")
	b.GUID("text")

	assert.Equal(t, 1, b.Params["owner_id"])
	assert.Equal(t, 1, b.Params["photo_id"])
	assert.Equal(t, "text", b.Params["message"])
	assert.Equal(t, []string{"text"}, b.Params["attachments"])
	assert.Equal(t, true, b.Params["from_group"])
	assert.Equal(t, 1, b.Params["reply_to_comment"])
	assert.Equal(t, 1, b.Params["sticker_id"])
	assert.Equal(t, "text", b.Params["access_key"])
	assert.Equal(t, "text", b.Params["guid"])
}

func TestPhotosDeleteBuilder(t *testing.T) {
	t.Parallel()

	b := params.NewPhotosDeleteBuilder()

	b.OwnerID(1)
	b.PhotoID(1)

	assert.Equal(t, 1, b.Params["owner_id"])
	assert.Equal(t, 1, b.Params["photo_id"])
}

func TestPhotosDeleteAlbumBuilder(t *testing.T) {
	t.Parallel()

	b := params.NewPhotosDeleteAlbumBuilder()

	b.AlbumID(1)
	b.GroupID(1)

	assert.Equal(t, 1, b.Params["album_id"])
	assert.Equal(t, 1, b.Params["group_id"])
}

func TestPhotosDeleteCommentBuilder(t *testing.T) {
	t.Parallel()

	b := params.NewPhotosDeleteCommentBuilder()

	b.OwnerID(1)
	b.CommentID(1)

	assert.Equal(t, 1, b.Params["owner_id"])
	assert.Equal(t, 1, b.Params["comment_id"])
}

func TestPhotosEditBuilder(t *testing.T) {
	t.Parallel()

	b := params.NewPhotosEditBuilder()

	b.OwnerID(1)
	b.PhotoID(1)
	b.Caption("text")
	b.Latitude(1.1)
	b.Longitude(1.1)
	b.PlaceStr("text")
	b.FoursquareID("text")
	b.DeletePlace(true)

	assert.Equal(t, 1, b.Params["owner_id"])
	assert.Equal(t, 1, b.Params["photo_id"])
	assert.Equal(t, "text", b.Params["caption"])
	assert.InEpsilon(t, 1.1, b.Params["latitude"], 0.1)
	assert.InEpsilon(t, 1.1, b.Params["longitude"], 0.1)
	assert.Equal(t, "text", b.Params["place_str"])
	assert.Equal(t, "text", b.Params["foursquare_id"])
	assert.Equal(t, true, b.Params["delete_place"])
}

func TestPhotosEditAlbumBuilder(t *testing.T) {
	t.Parallel()

	b := params.NewPhotosEditAlbumBuilder()

	b.AlbumID(1)
	b.Title("text")
	b.Description("text")
	b.OwnerID(1)
	b.PrivacyView([]string{"text"})
	b.PrivacyComment([]string{"text"})
	b.UploadByAdminsOnly(true)
	b.CommentsDisabled(true)

	assert.Equal(t, 1, b.Params["album_id"])
	assert.Equal(t, "text", b.Params["title"])
	assert.Equal(t, "text", b.Params["description"])
	assert.Equal(t, 1, b.Params["owner_id"])
	assert.Equal(t, []string{"text"}, b.Params["privacy_view"])
	assert.Equal(t, []string{"text"}, b.Params["privacy_comment"])
	assert.Equal(t, true, b.Params["upload_by_admins_only"])
	assert.Equal(t, true, b.Params["comments_disabled"])
}

func TestPhotosEditCommentBuilder(t *testing.T) {
	t.Parallel()

	b := params.NewPhotosEditCommentBuilder()

	b.OwnerID(1)
	b.CommentID(1)
	b.Message("text")
	b.Attachments([]string{"text"})

	assert.Equal(t, 1, b.Params["owner_id"])
	assert.Equal(t, 1, b.Params["comment_id"])
	assert.Equal(t, "text", b.Params["message"])
	assert.Equal(t, []string{"text"}, b.Params["attachments"])
}

func TestPhotosGetBuilder(t *testing.T) {
	t.Parallel()

	b := params.NewPhotosGetBuilder()

	b.OwnerID(1)
	b.AlbumID("text")
	b.PhotoIDs([]string{"text"})
	b.Rev(true)
	b.Extended(true)
	b.FeedType("text")
	b.Feed(1)
	b.PhotoSizes(true)
	b.Offset(1)
	b.Count(1)

	assert.Equal(t, 1, b.Params["owner_id"])
	assert.Equal(t, "text", b.Params["album_id"])
	assert.Equal(t, []string{"text"}, b.Params["photo_ids"])
	assert.Equal(t, true, b.Params["rev"])
	assert.Equal(t, true, b.Params["extended"])
	assert.Equal(t, "text", b.Params["feed_type"])
	assert.Equal(t, 1, b.Params["feed"])
	assert.Equal(t, true, b.Params["photo_sizes"])
	assert.Equal(t, 1, b.Params["offset"])
	assert.Equal(t, 1, b.Params["count"])
}

func TestPhotosGetAlbumsBuilder(t *testing.T) {
	t.Parallel()

	b := params.NewPhotosGetAlbumsBuilder()

	b.OwnerID(1)
	b.AlbumIDs([]int{1})
	b.Offset(1)
	b.Count(1)
	b.NeedSystem(true)
	b.NeedCovers(true)
	b.PhotoSizes(true)

	assert.Equal(t, 1, b.Params["owner_id"])
	assert.Equal(t, []int{1}, b.Params["album_ids"])
	assert.Equal(t, 1, b.Params["offset"])
	assert.Equal(t, 1, b.Params["count"])
	assert.Equal(t, true, b.Params["need_system"])
	assert.Equal(t, true, b.Params["need_covers"])
	assert.Equal(t, true, b.Params["photo_sizes"])
}

func TestPhotosGetAlbumsCountBuilder(t *testing.T) {
	t.Parallel()

	b := params.NewPhotosGetAlbumsCountBuilder()

	b.UserID(1)
	b.GroupID(1)

	assert.Equal(t, 1, b.Params["user_id"])
	assert.Equal(t, 1, b.Params["group_id"])
}

func TestPhotosGetAllBuilder(t *testing.T) {
	t.Parallel()

	b := params.NewPhotosGetAllBuilder()

	b.OwnerID(1)
	b.Extended(true)
	b.Offset(1)
	b.Count(1)
	b.PhotoSizes(true)
	b.NoServiceAlbums(true)
	b.NeedHidden(true)
	b.SkipHidden(true)

	assert.Equal(t, 1, b.Params["owner_id"])
	assert.Equal(t, true, b.Params["extended"])
	assert.Equal(t, 1, b.Params["offset"])
	assert.Equal(t, 1, b.Params["count"])
	assert.Equal(t, true, b.Params["photo_sizes"])
	assert.Equal(t, true, b.Params["no_service_albums"])
	assert.Equal(t, true, b.Params["need_hidden"])
	assert.Equal(t, true, b.Params["skip_hidden"])
}

func TestPhotosGetAllCommentsBuilder(t *testing.T) {
	t.Parallel()

	b := params.NewPhotosGetAllCommentsBuilder()

	b.OwnerID(1)
	b.AlbumID(1)
	b.NeedLikes(true)
	b.Offset(1)
	b.Count(1)

	assert.Equal(t, 1, b.Params["owner_id"])
	assert.Equal(t, 1, b.Params["album_id"])
	assert.Equal(t, true, b.Params["need_likes"])
	assert.Equal(t, 1, b.Params["offset"])
	assert.Equal(t, 1, b.Params["count"])
}

func TestPhotosGetByIDBuilder(t *testing.T) {
	t.Parallel()

	b := params.NewPhotosGetByIDBuilder()

	b.Photos([]string{"text"})
	b.Extended(true)
	b.PhotoSizes(true)

	assert.Equal(t, []string{"text"}, b.Params["photos"])
	assert.Equal(t, true, b.Params["extended"])
	assert.Equal(t, true, b.Params["photo_sizes"])
}

func TestPhotosGetChatUploadServerBuilder(t *testing.T) {
	t.Parallel()

	b := params.NewPhotosGetChatUploadServerBuilder()

	b.ChatID(1)
	b.CropX(1)
	b.CropY(1)
	b.CropWidth(1)

	assert.Equal(t, 1, b.Params["chat_id"])
	assert.Equal(t, 1, b.Params["crop_x"])
	assert.Equal(t, 1, b.Params["crop_y"])
	assert.Equal(t, 1, b.Params["crop_width"])
}

func TestPhotosGetCommentsBuilder(t *testing.T) {
	t.Parallel()

	b := params.NewPhotosGetCommentsBuilder()

	b.OwnerID(1)
	b.PhotoID(1)
	b.NeedLikes(true)
	b.StartCommentID(1)
	b.Offset(1)
	b.Count(1)
	b.Sort("text")
	b.AccessKey("text")
	b.Extended(true)
	b.Fields([]string{"test"})

	assert.Equal(t, 1, b.Params["owner_id"])
	assert.Equal(t, 1, b.Params["photo_id"])
	assert.Equal(t, true, b.Params["need_likes"])
	assert.Equal(t, 1, b.Params["start_comment_id"])
	assert.Equal(t, 1, b.Params["offset"])
	assert.Equal(t, 1, b.Params["count"])
	assert.Equal(t, "text", b.Params["sort"])
	assert.Equal(t, "text", b.Params["access_key"])
	assert.Equal(t, true, b.Params["extended"])
	assert.Equal(t, []string{"test"}, b.Params["fields"])
}

func TestPhotosGetMarketAlbumUploadServerBuilder(t *testing.T) {
	t.Parallel()

	b := params.NewPhotosGetMarketAlbumUploadServerBuilder()

	b.GroupID(1)

	assert.Equal(t, 1, b.Params["group_id"])
}

func TestPhotosGetMarketUploadServerBuilder(t *testing.T) {
	t.Parallel()

	b := params.NewPhotosGetMarketUploadServerBuilder()

	b.GroupID(1)
	b.MainPhoto(true)
	b.CropX(1)
	b.CropY(1)
	b.CropWidth(1)

	assert.Equal(t, 1, b.Params["group_id"])
	assert.Equal(t, true, b.Params["main_photo"])
	assert.Equal(t, 1, b.Params["crop_x"])
	assert.Equal(t, 1, b.Params["crop_y"])
	assert.Equal(t, 1, b.Params["crop_width"])
}

func TestPhotosGetMessagesUploadServerBuilder(t *testing.T) {
	t.Parallel()

	b := params.NewPhotosGetMessagesUploadServerBuilder()

	b.PeerID(1)

	assert.Equal(t, 1, b.Params["peer_id"])
}

func TestPhotosGetNewTagsBuilder(t *testing.T) {
	t.Parallel()

	b := params.NewPhotosGetNewTagsBuilder()

	b.Offset(1)
	b.Count(1)

	assert.Equal(t, 1, b.Params["offset"])
	assert.Equal(t, 1, b.Params["count"])
}

func TestPhotosGetOwnerCoverPhotoUploadServerBuilder(t *testing.T) {
	t.Parallel()

	b := params.NewPhotosGetOwnerCoverPhotoUploadServerBuilder()

	b.GroupID(1)
	b.CropX(1)
	b.CropY(1)
	b.CropX2(1)
	b.CropY2(1)

	assert.Equal(t, 1, b.Params["group_id"])
	assert.Equal(t, 1, b.Params["crop_x"])
	assert.Equal(t, 1, b.Params["crop_y"])
	assert.Equal(t, 1, b.Params["crop_x2"])
	assert.Equal(t, 1, b.Params["crop_y2"])
}

func TestPhotosGetOwnerPhotoUploadServerBuilder(t *testing.T) {
	t.Parallel()

	b := params.NewPhotosGetOwnerPhotoUploadServerBuilder()

	b.OwnerID(1)

	assert.Equal(t, 1, b.Params["owner_id"])
}

func TestPhotosGetTagsBuilder(t *testing.T) {
	t.Parallel()

	b := params.NewPhotosGetTagsBuilder()

	b.OwnerID(1)
	b.PhotoID(1)
	b.AccessKey("text")

	assert.Equal(t, 1, b.Params["owner_id"])
	assert.Equal(t, 1, b.Params["photo_id"])
	assert.Equal(t, "text", b.Params["access_key"])
}

func TestPhotosGetUploadServerBuilder(t *testing.T) {
	t.Parallel()

	b := params.NewPhotosGetUploadServerBuilder()

	b.GroupID(1)
	b.AlbumID(1)

	assert.Equal(t, 1, b.Params["group_id"])
	assert.Equal(t, 1, b.Params["album_id"])
}

func TestPhotosGetUserPhotosBuilder(t *testing.T) {
	t.Parallel()

	b := params.NewPhotosGetUserPhotosBuilder()

	b.UserID(1)
	b.Offset(1)
	b.Count(1)
	b.Extended(true)
	b.Sort("text")

	assert.Equal(t, 1, b.Params["user_id"])
	assert.Equal(t, 1, b.Params["offset"])
	assert.Equal(t, 1, b.Params["count"])
	assert.Equal(t, true, b.Params["extended"])
	assert.Equal(t, "text", b.Params["sort"])
}

func TestPhotosGetWallUploadServerBuilder(t *testing.T) {
	t.Parallel()

	b := params.NewPhotosGetWallUploadServerBuilder()

	b.GroupID(1)

	assert.Equal(t, 1, b.Params["group_id"])
}

func TestPhotosMakeCoverBuilder(t *testing.T) {
	t.Parallel()

	b := params.NewPhotosMakeCoverBuilder()

	b.OwnerID(1)
	b.PhotoID(1)
	b.AlbumID(1)

	assert.Equal(t, 1, b.Params["owner_id"])
	assert.Equal(t, 1, b.Params["photo_id"])
	assert.Equal(t, 1, b.Params["album_id"])
}

func TestPhotosMoveBuilder(t *testing.T) {
	t.Parallel()

	b := params.NewPhotosMoveBuilder()

	b.OwnerID(1)
	b.TargetAlbumID(1)
	b.PhotoID(1)

	assert.Equal(t, 1, b.Params["owner_id"])
	assert.Equal(t, 1, b.Params["target_album_id"])
	assert.Equal(t, 1, b.Params["photo_id"])
}

func TestPhotosPutTagBuilder(t *testing.T) {
	t.Parallel()

	b := params.NewPhotosPutTagBuilder()

	b.OwnerID(1)
	b.PhotoID(1)
	b.UserID(1)
	b.X(1.1)
	b.Y(1.1)
	b.X2(1.1)
	b.Y2(1.1)

	assert.Equal(t, 1, b.Params["owner_id"])
	assert.Equal(t, 1, b.Params["photo_id"])
	assert.Equal(t, 1, b.Params["user_id"])
	assert.InEpsilon(t, 1.1, b.Params["x"], 0.1)
	assert.InEpsilon(t, 1.1, b.Params["y"], 0.1)
	assert.InEpsilon(t, 1.1, b.Params["x2"], 0.1)
	assert.InEpsilon(t, 1.1, b.Params["y2"], 0.1)
}

func TestPhotosRemoveTagBuilder(t *testing.T) {
	t.Parallel()

	b := params.NewPhotosRemoveTagBuilder()

	b.OwnerID(1)
	b.PhotoID(1)
	b.TagID(1)

	assert.Equal(t, 1, b.Params["owner_id"])
	assert.Equal(t, 1, b.Params["photo_id"])
	assert.Equal(t, 1, b.Params["tag_id"])
}

func TestPhotosReorderAlbumsBuilder(t *testing.T) {
	t.Parallel()

	b := params.NewPhotosReorderAlbumsBuilder()

	b.OwnerID(1)
	b.AlbumID(1)
	b.Before(1)
	b.After(1)

	assert.Equal(t, 1, b.Params["owner_id"])
	assert.Equal(t, 1, b.Params["album_id"])
	assert.Equal(t, 1, b.Params["before"])
	assert.Equal(t, 1, b.Params["after"])
}

func TestPhotosReorderPhotosBuilder(t *testing.T) {
	t.Parallel()

	b := params.NewPhotosReorderPhotosBuilder()

	b.OwnerID(1)
	b.PhotoID(1)
	b.Before(1)
	b.After(1)

	assert.Equal(t, 1, b.Params["owner_id"])
	assert.Equal(t, 1, b.Params["photo_id"])
	assert.Equal(t, 1, b.Params["before"])
	assert.Equal(t, 1, b.Params["after"])
}

func TestPhotosReportBuilder(t *testing.T) {
	t.Parallel()

	b := params.NewPhotosReportBuilder()

	b.OwnerID(1)
	b.PhotoID(1)
	b.Reason(1)

	assert.Equal(t, 1, b.Params["owner_id"])
	assert.Equal(t, 1, b.Params["photo_id"])
	assert.Equal(t, 1, b.Params["reason"])
}

func TestPhotosReportCommentBuilder(t *testing.T) {
	t.Parallel()

	b := params.NewPhotosReportCommentBuilder()

	b.OwnerID(1)
	b.CommentID(1)
	b.Reason(1)

	assert.Equal(t, 1, b.Params["owner_id"])
	assert.Equal(t, 1, b.Params["comment_id"])
	assert.Equal(t, 1, b.Params["reason"])
}

func TestPhotosRestoreBuilder(t *testing.T) {
	t.Parallel()

	b := params.NewPhotosRestoreBuilder()

	b.OwnerID(1)
	b.PhotoID(1)

	assert.Equal(t, 1, b.Params["owner_id"])
	assert.Equal(t, 1, b.Params["photo_id"])
}

func TestPhotosRestoreCommentBuilder(t *testing.T) {
	t.Parallel()

	b := params.NewPhotosRestoreCommentBuilder()

	b.OwnerID(1)
	b.CommentID(1)

	assert.Equal(t, 1, b.Params["owner_id"])
	assert.Equal(t, 1, b.Params["comment_id"])
}

func TestPhotosSaveBuilder(t *testing.T) {
	t.Parallel()

	b := params.NewPhotosSaveBuilder()

	b.AlbumID(1)
	b.GroupID(1)
	b.Server(1)
	b.PhotosList("text")
	b.Hash("text")
	b.Latitude(1.1)
	b.Longitude(1.1)
	b.Caption("text")

	assert.Equal(t, 1, b.Params["album_id"])
	assert.Equal(t, 1, b.Params["group_id"])
	assert.Equal(t, 1, b.Params["server"])
	assert.Equal(t, "text", b.Params["photos_list"])
	assert.Equal(t, "text", b.Params["hash"])
	assert.InEpsilon(t, 1.1, b.Params["latitude"], 0.1)
	assert.InEpsilon(t, 1.1, b.Params["longitude"], 0.1)
	assert.Equal(t, "text", b.Params["caption"])
}

func TestPhotosSaveMarketAlbumPhotoBuilder(t *testing.T) {
	t.Parallel()

	b := params.NewPhotosSaveMarketAlbumPhotoBuilder()

	b.GroupID(1)
	b.Photo("text")
	b.Server(1)
	b.Hash("text")

	assert.Equal(t, 1, b.Params["group_id"])
	assert.Equal(t, "text", b.Params["photo"])
	assert.Equal(t, 1, b.Params["server"])
	assert.Equal(t, "text", b.Params["hash"])
}

func TestPhotosSaveMarketPhotoBuilder(t *testing.T) {
	t.Parallel()

	b := params.NewPhotosSaveMarketPhotoBuilder()

	b.GroupID(1)
	b.Photo("text")
	b.Server(1)
	b.Hash("text")
	b.CropData("text")
	b.CropHash("text")

	assert.Equal(t, 1, b.Params["group_id"])
	assert.Equal(t, "text", b.Params["photo"])
	assert.Equal(t, 1, b.Params["server"])
	assert.Equal(t, "text", b.Params["hash"])
	assert.Equal(t, "text", b.Params["crop_data"])
	assert.Equal(t, "text", b.Params["crop_hash"])
}

func TestPhotosSaveMessagesPhotoBuilder(t *testing.T) {
	t.Parallel()

	b := params.NewPhotosSaveMessagesPhotoBuilder()

	b.Photo("text")
	b.Server(1)
	b.Hash("text")

	assert.Equal(t, "text", b.Params["photo"])
	assert.Equal(t, 1, b.Params["server"])
	assert.Equal(t, "text", b.Params["hash"])
}

func TestPhotosSaveOwnerCoverPhotoBuilder(t *testing.T) {
	t.Parallel()

	b := params.NewPhotosSaveOwnerCoverPhotoBuilder()

	b.Hash("text")
	b.Photo("text")

	assert.Equal(t, "text", b.Params["hash"])
	assert.Equal(t, "text", b.Params["photo"])
}

func TestPhotosSaveOwnerPhotoBuilder(t *testing.T) {
	t.Parallel()

	b := params.NewPhotosSaveOwnerPhotoBuilder()

	b.Server("text")
	b.Hash("text")
	b.Photo("text")

	assert.Equal(t, "text", b.Params["server"])
	assert.Equal(t, "text", b.Params["hash"])
	assert.Equal(t, "text", b.Params["photo"])
}

func TestPhotosSaveWallPhotoBuilder(t *testing.T) {
	t.Parallel()

	b := params.NewPhotosSaveWallPhotoBuilder()

	b.UserID(1)
	b.GroupID(1)
	b.Photo("text")
	b.Server(1)
	b.Hash("text")
	b.Latitude(1.1)
	b.Longitude(1.1)
	b.Caption("text")

	assert.Equal(t, 1, b.Params["user_id"])
	assert.Equal(t, 1, b.Params["group_id"])
	assert.Equal(t, "text", b.Params["photo"])
	assert.Equal(t, 1, b.Params["server"])
	assert.Equal(t, "text", b.Params["hash"])
	assert.InEpsilon(t, 1.1, b.Params["latitude"], 0.1)
	assert.InEpsilon(t, 1.1, b.Params["longitude"], 0.1)
	assert.Equal(t, "text", b.Params["caption"])
}

func TestPhotosSearchBuilder(t *testing.T) {
	t.Parallel()

	b := params.NewPhotosSearchBuilder()

	b.Q("text")
	b.Lat(1.1)
	b.Long(1.1)
	b.StartTime(1)
	b.EndTime(1)
	b.Sort(1)
	b.Offset(1)
	b.Count(1)
	b.Radius(1)

	assert.Equal(t, "text", b.Params["q"])
	assert.InEpsilon(t, 1.1, b.Params["lat"], 0.1)
	assert.InEpsilon(t, 1.1, b.Params["long"], 0.1)
	assert.Equal(t, 1, b.Params["start_time"])
	assert.Equal(t, 1, b.Params["end_time"])
	assert.Equal(t, 1, b.Params["sort"])
	assert.Equal(t, 1, b.Params["offset"])
	assert.Equal(t, 1, b.Params["count"])
	assert.Equal(t, 1, b.Params["radius"])
}
