package params_test

import (
	"testing"

	"github.com/SevereCloud/vksdk/api/params"
	"github.com/stretchr/testify/assert"
)

func TestPhotosConfirmTagBuilder(t *testing.T) {
	b := params.NewPhotosConfirmTagBuilder()

	b.OwnerID(1)
	b.PhotoID("text")
	b.TagID(1)

	assert.Equal(t, b.Params["owner_id"], 1)
	assert.Equal(t, b.Params["photo_id"], "text")
	assert.Equal(t, b.Params["tag_id"], 1)
}

func TestPhotosCopyBuilder(t *testing.T) {
	b := params.NewPhotosCopyBuilder()

	b.OwnerID(1)
	b.PhotoID(1)
	b.AccessKey("text")

	assert.Equal(t, b.Params["owner_id"], 1)
	assert.Equal(t, b.Params["photo_id"], 1)
	assert.Equal(t, b.Params["access_key"], "text")
}

func TestPhotosCreateAlbumBuilder(t *testing.T) {
	b := params.NewPhotosCreateAlbumBuilder()

	b.Title("text")
	b.GroupID(1)
	b.Description("text")
	b.PrivacyView([]string{"text"})
	b.PrivacyComment([]string{"text"})
	b.UploadByAdminsOnly(true)
	b.CommentsDisabled(true)

	assert.Equal(t, b.Params["title"], "text")
	assert.Equal(t, b.Params["group_id"], 1)
	assert.Equal(t, b.Params["description"], "text")
	assert.Equal(t, b.Params["privacy_view"], []string{"text"})
	assert.Equal(t, b.Params["privacy_comment"], []string{"text"})
	assert.Equal(t, b.Params["upload_by_admins_only"], true)
	assert.Equal(t, b.Params["comments_disabled"], true)
}

func TestPhotosCreateCommentBuilder(t *testing.T) {
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

	assert.Equal(t, b.Params["owner_id"], 1)
	assert.Equal(t, b.Params["photo_id"], 1)
	assert.Equal(t, b.Params["message"], "text")
	assert.Equal(t, b.Params["attachments"], []string{"text"})
	assert.Equal(t, b.Params["from_group"], true)
	assert.Equal(t, b.Params["reply_to_comment"], 1)
	assert.Equal(t, b.Params["sticker_id"], 1)
	assert.Equal(t, b.Params["access_key"], "text")
	assert.Equal(t, b.Params["guid"], "text")
}

func TestPhotosDeleteBuilder(t *testing.T) {
	b := params.NewPhotosDeleteBuilder()

	b.OwnerID(1)
	b.PhotoID(1)

	assert.Equal(t, b.Params["owner_id"], 1)
	assert.Equal(t, b.Params["photo_id"], 1)
}

func TestPhotosDeleteAlbumBuilder(t *testing.T) {
	b := params.NewPhotosDeleteAlbumBuilder()

	b.AlbumID(1)
	b.GroupID(1)

	assert.Equal(t, b.Params["album_id"], 1)
	assert.Equal(t, b.Params["group_id"], 1)
}

func TestPhotosDeleteCommentBuilder(t *testing.T) {
	b := params.NewPhotosDeleteCommentBuilder()

	b.OwnerID(1)
	b.CommentID(1)

	assert.Equal(t, b.Params["owner_id"], 1)
	assert.Equal(t, b.Params["comment_id"], 1)
}

func TestPhotosEditBuilder(t *testing.T) {
	b := params.NewPhotosEditBuilder()

	b.OwnerID(1)
	b.PhotoID(1)
	b.Caption("text")
	b.Latitude(1.1)
	b.Longitude(1.1)
	b.PlaceStr("text")
	b.FoursquareID("text")
	b.DeletePlace(true)

	assert.Equal(t, b.Params["owner_id"], 1)
	assert.Equal(t, b.Params["photo_id"], 1)
	assert.Equal(t, b.Params["caption"], "text")
	assert.Equal(t, b.Params["latitude"], 1.1)
	assert.Equal(t, b.Params["longitude"], 1.1)
	assert.Equal(t, b.Params["place_str"], "text")
	assert.Equal(t, b.Params["foursquare_id"], "text")
	assert.Equal(t, b.Params["delete_place"], true)
}

func TestPhotosEditAlbumBuilder(t *testing.T) {
	b := params.NewPhotosEditAlbumBuilder()

	b.AlbumID(1)
	b.Title("text")
	b.Description("text")
	b.OwnerID(1)
	b.PrivacyView([]string{"text"})
	b.PrivacyComment([]string{"text"})
	b.UploadByAdminsOnly(true)
	b.CommentsDisabled(true)

	assert.Equal(t, b.Params["album_id"], 1)
	assert.Equal(t, b.Params["title"], "text")
	assert.Equal(t, b.Params["description"], "text")
	assert.Equal(t, b.Params["owner_id"], 1)
	assert.Equal(t, b.Params["privacy_view"], []string{"text"})
	assert.Equal(t, b.Params["privacy_comment"], []string{"text"})
	assert.Equal(t, b.Params["upload_by_admins_only"], true)
	assert.Equal(t, b.Params["comments_disabled"], true)
}

func TestPhotosEditCommentBuilder(t *testing.T) {
	b := params.NewPhotosEditCommentBuilder()

	b.OwnerID(1)
	b.CommentID(1)
	b.Message("text")
	b.Attachments([]string{"text"})

	assert.Equal(t, b.Params["owner_id"], 1)
	assert.Equal(t, b.Params["comment_id"], 1)
	assert.Equal(t, b.Params["message"], "text")
	assert.Equal(t, b.Params["attachments"], []string{"text"})
}

func TestPhotosGetBuilder(t *testing.T) {
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

	assert.Equal(t, b.Params["owner_id"], 1)
	assert.Equal(t, b.Params["album_id"], "text")
	assert.Equal(t, b.Params["photo_ids"], []string{"text"})
	assert.Equal(t, b.Params["rev"], true)
	assert.Equal(t, b.Params["extended"], true)
	assert.Equal(t, b.Params["feed_type"], "text")
	assert.Equal(t, b.Params["feed"], 1)
	assert.Equal(t, b.Params["photo_sizes"], true)
	assert.Equal(t, b.Params["offset"], 1)
	assert.Equal(t, b.Params["count"], 1)
}

func TestPhotosGetAlbumsBuilder(t *testing.T) {
	b := params.NewPhotosGetAlbumsBuilder()

	b.OwnerID(1)
	b.AlbumIDs([]int{1})
	b.Offset(1)
	b.Count(1)
	b.NeedSystem(true)
	b.NeedCovers(true)
	b.PhotoSizes(true)

	assert.Equal(t, b.Params["owner_id"], 1)
	assert.Equal(t, b.Params["album_ids"], []int{1})
	assert.Equal(t, b.Params["offset"], 1)
	assert.Equal(t, b.Params["count"], 1)
	assert.Equal(t, b.Params["need_system"], true)
	assert.Equal(t, b.Params["need_covers"], true)
	assert.Equal(t, b.Params["photo_sizes"], true)
}

func TestPhotosGetAlbumsCountBuilder(t *testing.T) {
	b := params.NewPhotosGetAlbumsCountBuilder()

	b.UserID(1)
	b.GroupID(1)

	assert.Equal(t, b.Params["user_id"], 1)
	assert.Equal(t, b.Params["group_id"], 1)
}

func TestPhotosGetAllBuilder(t *testing.T) {
	b := params.NewPhotosGetAllBuilder()

	b.OwnerID(1)
	b.Extended(true)
	b.Offset(1)
	b.Count(1)
	b.PhotoSizes(true)
	b.NoServiceAlbums(true)
	b.NeedHidden(true)
	b.SkipHidden(true)

	assert.Equal(t, b.Params["owner_id"], 1)
	assert.Equal(t, b.Params["extended"], true)
	assert.Equal(t, b.Params["offset"], 1)
	assert.Equal(t, b.Params["count"], 1)
	assert.Equal(t, b.Params["photo_sizes"], true)
	assert.Equal(t, b.Params["no_service_albums"], true)
	assert.Equal(t, b.Params["need_hidden"], true)
	assert.Equal(t, b.Params["skip_hidden"], true)
}

func TestPhotosGetAllCommentsBuilder(t *testing.T) {
	b := params.NewPhotosGetAllCommentsBuilder()

	b.OwnerID(1)
	b.AlbumID(1)
	b.NeedLikes(true)
	b.Offset(1)
	b.Count(1)

	assert.Equal(t, b.Params["owner_id"], 1)
	assert.Equal(t, b.Params["album_id"], 1)
	assert.Equal(t, b.Params["need_likes"], true)
	assert.Equal(t, b.Params["offset"], 1)
	assert.Equal(t, b.Params["count"], 1)
}

func TestPhotosGetByIDBuilder(t *testing.T) {
	b := params.NewPhotosGetByIDBuilder()

	b.Photos([]string{"text"})
	b.Extended(true)
	b.PhotoSizes(true)

	assert.Equal(t, b.Params["photos"], []string{"text"})
	assert.Equal(t, b.Params["extended"], true)
	assert.Equal(t, b.Params["photo_sizes"], true)
}

func TestPhotosGetChatUploadServerBuilder(t *testing.T) {
	b := params.NewPhotosGetChatUploadServerBuilder()

	b.ChatID(1)
	b.CropX(1)
	b.CropY(1)
	b.CropWidth(1)

	assert.Equal(t, b.Params["chat_id"], 1)
	assert.Equal(t, b.Params["crop_x"], 1)
	assert.Equal(t, b.Params["crop_y"], 1)
	assert.Equal(t, b.Params["crop_width"], 1)
}

func TestPhotosGetCommentsBuilder(t *testing.T) {
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

	assert.Equal(t, b.Params["owner_id"], 1)
	assert.Equal(t, b.Params["photo_id"], 1)
	assert.Equal(t, b.Params["need_likes"], true)
	assert.Equal(t, b.Params["start_comment_id"], 1)
	assert.Equal(t, b.Params["offset"], 1)
	assert.Equal(t, b.Params["count"], 1)
	assert.Equal(t, b.Params["sort"], "text")
	assert.Equal(t, b.Params["access_key"], "text")
	assert.Equal(t, b.Params["extended"], true)
	assert.Equal(t, b.Params["fields"], []string{"test"})
}

func TestPhotosGetMarketAlbumUploadServerBuilder(t *testing.T) {
	b := params.NewPhotosGetMarketAlbumUploadServerBuilder()

	b.GroupID(1)

	assert.Equal(t, b.Params["group_id"], 1)
}

func TestPhotosGetMarketUploadServerBuilder(t *testing.T) {
	b := params.NewPhotosGetMarketUploadServerBuilder()

	b.GroupID(1)
	b.MainPhoto(true)
	b.CropX(1)
	b.CropY(1)
	b.CropWidth(1)

	assert.Equal(t, b.Params["group_id"], 1)
	assert.Equal(t, b.Params["main_photo"], true)
	assert.Equal(t, b.Params["crop_x"], 1)
	assert.Equal(t, b.Params["crop_y"], 1)
	assert.Equal(t, b.Params["crop_width"], 1)
}

func TestPhotosGetMessagesUploadServerBuilder(t *testing.T) {
	b := params.NewPhotosGetMessagesUploadServerBuilder()

	b.PeerID(1)

	assert.Equal(t, b.Params["peer_id"], 1)
}

func TestPhotosGetNewTagsBuilder(t *testing.T) {
	b := params.NewPhotosGetNewTagsBuilder()

	b.Offset(1)
	b.Count(1)

	assert.Equal(t, b.Params["offset"], 1)
	assert.Equal(t, b.Params["count"], 1)
}

func TestPhotosGetOwnerCoverPhotoUploadServerBuilder(t *testing.T) {
	b := params.NewPhotosGetOwnerCoverPhotoUploadServerBuilder()

	b.GroupID(1)
	b.CropX(1)
	b.CropY(1)
	b.CropX2(1)
	b.CropY2(1)

	assert.Equal(t, b.Params["group_id"], 1)
	assert.Equal(t, b.Params["crop_x"], 1)
	assert.Equal(t, b.Params["crop_y"], 1)
	assert.Equal(t, b.Params["crop_x2"], 1)
	assert.Equal(t, b.Params["crop_y2"], 1)
}

func TestPhotosGetOwnerPhotoUploadServerBuilder(t *testing.T) {
	b := params.NewPhotosGetOwnerPhotoUploadServerBuilder()

	b.OwnerID(1)

	assert.Equal(t, b.Params["owner_id"], 1)
}

func TestPhotosGetTagsBuilder(t *testing.T) {
	b := params.NewPhotosGetTagsBuilder()

	b.OwnerID(1)
	b.PhotoID(1)
	b.AccessKey("text")

	assert.Equal(t, b.Params["owner_id"], 1)
	assert.Equal(t, b.Params["photo_id"], 1)
	assert.Equal(t, b.Params["access_key"], "text")
}

func TestPhotosGetUploadServerBuilder(t *testing.T) {
	b := params.NewPhotosGetUploadServerBuilder()

	b.GroupID(1)
	b.AlbumID(1)

	assert.Equal(t, b.Params["group_id"], 1)
	assert.Equal(t, b.Params["album_id"], 1)
}

func TestPhotosGetUserPhotosBuilder(t *testing.T) {
	b := params.NewPhotosGetUserPhotosBuilder()

	b.UserID(1)
	b.Offset(1)
	b.Count(1)
	b.Extended(true)
	b.Sort("text")

	assert.Equal(t, b.Params["user_id"], 1)
	assert.Equal(t, b.Params["offset"], 1)
	assert.Equal(t, b.Params["count"], 1)
	assert.Equal(t, b.Params["extended"], true)
	assert.Equal(t, b.Params["sort"], "text")
}

func TestPhotosGetWallUploadServerBuilder(t *testing.T) {
	b := params.NewPhotosGetWallUploadServerBuilder()

	b.GroupID(1)

	assert.Equal(t, b.Params["group_id"], 1)
}

func TestPhotosMakeCoverBuilder(t *testing.T) {
	b := params.NewPhotosMakeCoverBuilder()

	b.OwnerID(1)
	b.PhotoID(1)
	b.AlbumID(1)

	assert.Equal(t, b.Params["owner_id"], 1)
	assert.Equal(t, b.Params["photo_id"], 1)
	assert.Equal(t, b.Params["album_id"], 1)
}

func TestPhotosMoveBuilder(t *testing.T) {
	b := params.NewPhotosMoveBuilder()

	b.OwnerID(1)
	b.TargetAlbumID(1)
	b.PhotoID(1)

	assert.Equal(t, b.Params["owner_id"], 1)
	assert.Equal(t, b.Params["target_album_id"], 1)
	assert.Equal(t, b.Params["photo_id"], 1)
}

func TestPhotosPutTagBuilder(t *testing.T) {
	b := params.NewPhotosPutTagBuilder()

	b.OwnerID(1)
	b.PhotoID(1)
	b.UserID(1)
	b.X(1.1)
	b.Y(1.1)
	b.X2(1.1)
	b.Y2(1.1)

	assert.Equal(t, b.Params["owner_id"], 1)
	assert.Equal(t, b.Params["photo_id"], 1)
	assert.Equal(t, b.Params["user_id"], 1)
	assert.Equal(t, b.Params["x"], 1.1)
	assert.Equal(t, b.Params["y"], 1.1)
	assert.Equal(t, b.Params["x2"], 1.1)
	assert.Equal(t, b.Params["y2"], 1.1)
}

func TestPhotosRemoveTagBuilder(t *testing.T) {
	b := params.NewPhotosRemoveTagBuilder()

	b.OwnerID(1)
	b.PhotoID(1)
	b.TagID(1)

	assert.Equal(t, b.Params["owner_id"], 1)
	assert.Equal(t, b.Params["photo_id"], 1)
	assert.Equal(t, b.Params["tag_id"], 1)
}

func TestPhotosReorderAlbumsBuilder(t *testing.T) {
	b := params.NewPhotosReorderAlbumsBuilder()

	b.OwnerID(1)
	b.AlbumID(1)
	b.Before(1)
	b.After(1)

	assert.Equal(t, b.Params["owner_id"], 1)
	assert.Equal(t, b.Params["album_id"], 1)
	assert.Equal(t, b.Params["before"], 1)
	assert.Equal(t, b.Params["after"], 1)
}

func TestPhotosReorderPhotosBuilder(t *testing.T) {
	b := params.NewPhotosReorderPhotosBuilder()

	b.OwnerID(1)
	b.PhotoID(1)
	b.Before(1)
	b.After(1)

	assert.Equal(t, b.Params["owner_id"], 1)
	assert.Equal(t, b.Params["photo_id"], 1)
	assert.Equal(t, b.Params["before"], 1)
	assert.Equal(t, b.Params["after"], 1)
}

func TestPhotosReportBuilder(t *testing.T) {
	b := params.NewPhotosReportBuilder()

	b.OwnerID(1)
	b.PhotoID(1)
	b.Reason(1)

	assert.Equal(t, b.Params["owner_id"], 1)
	assert.Equal(t, b.Params["photo_id"], 1)
	assert.Equal(t, b.Params["reason"], 1)
}

func TestPhotosReportCommentBuilder(t *testing.T) {
	b := params.NewPhotosReportCommentBuilder()

	b.OwnerID(1)
	b.CommentID(1)
	b.Reason(1)

	assert.Equal(t, b.Params["owner_id"], 1)
	assert.Equal(t, b.Params["comment_id"], 1)
	assert.Equal(t, b.Params["reason"], 1)
}

func TestPhotosRestoreBuilder(t *testing.T) {
	b := params.NewPhotosRestoreBuilder()

	b.OwnerID(1)
	b.PhotoID(1)

	assert.Equal(t, b.Params["owner_id"], 1)
	assert.Equal(t, b.Params["photo_id"], 1)
}

func TestPhotosRestoreCommentBuilder(t *testing.T) {
	b := params.NewPhotosRestoreCommentBuilder()

	b.OwnerID(1)
	b.CommentID(1)

	assert.Equal(t, b.Params["owner_id"], 1)
	assert.Equal(t, b.Params["comment_id"], 1)
}

func TestPhotosSaveBuilder(t *testing.T) {
	b := params.NewPhotosSaveBuilder()

	b.AlbumID(1)
	b.GroupID(1)
	b.Server(1)
	b.PhotosList("text")
	b.Hash("text")
	b.Latitude(1.1)
	b.Longitude(1.1)
	b.Caption("text")

	assert.Equal(t, b.Params["album_id"], 1)
	assert.Equal(t, b.Params["group_id"], 1)
	assert.Equal(t, b.Params["server"], 1)
	assert.Equal(t, b.Params["photos_list"], "text")
	assert.Equal(t, b.Params["hash"], "text")
	assert.Equal(t, b.Params["latitude"], 1.1)
	assert.Equal(t, b.Params["longitude"], 1.1)
	assert.Equal(t, b.Params["caption"], "text")
}

func TestPhotosSaveMarketAlbumPhotoBuilder(t *testing.T) {
	b := params.NewPhotosSaveMarketAlbumPhotoBuilder()

	b.GroupID(1)
	b.Photo("text")
	b.Server(1)
	b.Hash("text")

	assert.Equal(t, b.Params["group_id"], 1)
	assert.Equal(t, b.Params["photo"], "text")
	assert.Equal(t, b.Params["server"], 1)
	assert.Equal(t, b.Params["hash"], "text")
}

func TestPhotosSaveMarketPhotoBuilder(t *testing.T) {
	b := params.NewPhotosSaveMarketPhotoBuilder()

	b.GroupID(1)
	b.Photo("text")
	b.Server(1)
	b.Hash("text")
	b.CropData("text")
	b.CropHash("text")

	assert.Equal(t, b.Params["group_id"], 1)
	assert.Equal(t, b.Params["photo"], "text")
	assert.Equal(t, b.Params["server"], 1)
	assert.Equal(t, b.Params["hash"], "text")
	assert.Equal(t, b.Params["crop_data"], "text")
	assert.Equal(t, b.Params["crop_hash"], "text")
}

func TestPhotosSaveMessagesPhotoBuilder(t *testing.T) {
	b := params.NewPhotosSaveMessagesPhotoBuilder()

	b.Photo("text")
	b.Server(1)
	b.Hash("text")

	assert.Equal(t, b.Params["photo"], "text")
	assert.Equal(t, b.Params["server"], 1)
	assert.Equal(t, b.Params["hash"], "text")
}

func TestPhotosSaveOwnerCoverPhotoBuilder(t *testing.T) {
	b := params.NewPhotosSaveOwnerCoverPhotoBuilder()

	b.Hash("text")
	b.Photo("text")

	assert.Equal(t, b.Params["hash"], "text")
	assert.Equal(t, b.Params["photo"], "text")
}

func TestPhotosSaveOwnerPhotoBuilder(t *testing.T) {
	b := params.NewPhotosSaveOwnerPhotoBuilder()

	b.Server("text")
	b.Hash("text")
	b.Photo("text")

	assert.Equal(t, b.Params["server"], "text")
	assert.Equal(t, b.Params["hash"], "text")
	assert.Equal(t, b.Params["photo"], "text")
}

func TestPhotosSaveWallPhotoBuilder(t *testing.T) {
	b := params.NewPhotosSaveWallPhotoBuilder()

	b.UserID(1)
	b.GroupID(1)
	b.Photo("text")
	b.Server(1)
	b.Hash("text")
	b.Latitude(1.1)
	b.Longitude(1.1)
	b.Caption("text")

	assert.Equal(t, b.Params["user_id"], 1)
	assert.Equal(t, b.Params["group_id"], 1)
	assert.Equal(t, b.Params["photo"], "text")
	assert.Equal(t, b.Params["server"], 1)
	assert.Equal(t, b.Params["hash"], "text")
	assert.Equal(t, b.Params["latitude"], 1.1)
	assert.Equal(t, b.Params["longitude"], 1.1)
	assert.Equal(t, b.Params["caption"], "text")
}

func TestPhotosSearchBuilder(t *testing.T) {
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

	assert.Equal(t, b.Params["q"], "text")
	assert.Equal(t, b.Params["lat"], 1.1)
	assert.Equal(t, b.Params["long"], 1.1)
	assert.Equal(t, b.Params["start_time"], 1)
	assert.Equal(t, b.Params["end_time"], 1)
	assert.Equal(t, b.Params["sort"], 1)
	assert.Equal(t, b.Params["offset"], 1)
	assert.Equal(t, b.Params["count"], 1)
	assert.Equal(t, b.Params["radius"], 1)
}
