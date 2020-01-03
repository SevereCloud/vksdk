package params_test

import (
	"testing"

	"github.com/SevereCloud/vksdk/api/params"
	"github.com/stretchr/testify/assert"
)

func TestPhotosConfirmTagBulder(t *testing.T) {
	b := params.NewPhotosConfirmTagBulder()

	b.OwnerID(1)
	b.PhotoID("text")
	b.TagID(1)

	assert.Equal(t, b.Params["owner_id"], 1)
	assert.Equal(t, b.Params["photo_id"], "text")
	assert.Equal(t, b.Params["tag_id"], 1)
}

func TestPhotosCopyBulder(t *testing.T) {
	b := params.NewPhotosCopyBulder()

	b.OwnerID(1)
	b.PhotoID(1)
	b.AccessKey("text")

	assert.Equal(t, b.Params["owner_id"], 1)
	assert.Equal(t, b.Params["photo_id"], 1)
	assert.Equal(t, b.Params["access_key"], "text")
}

func TestPhotosCreateAlbumBulder(t *testing.T) {
	b := params.NewPhotosCreateAlbumBulder()

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

func TestPhotosCreateCommentBulder(t *testing.T) {
	b := params.NewPhotosCreateCommentBulder()

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

func TestPhotosDeleteBulder(t *testing.T) {
	b := params.NewPhotosDeleteBulder()

	b.OwnerID(1)
	b.PhotoID(1)

	assert.Equal(t, b.Params["owner_id"], 1)
	assert.Equal(t, b.Params["photo_id"], 1)
}

func TestPhotosDeleteAlbumBulder(t *testing.T) {
	b := params.NewPhotosDeleteAlbumBulder()

	b.AlbumID(1)
	b.GroupID(1)

	assert.Equal(t, b.Params["album_id"], 1)
	assert.Equal(t, b.Params["group_id"], 1)
}

func TestPhotosDeleteCommentBulder(t *testing.T) {
	b := params.NewPhotosDeleteCommentBulder()

	b.OwnerID(1)
	b.CommentID(1)

	assert.Equal(t, b.Params["owner_id"], 1)
	assert.Equal(t, b.Params["comment_id"], 1)
}

func TestPhotosEditBulder(t *testing.T) {
	b := params.NewPhotosEditBulder()

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

func TestPhotosEditAlbumBulder(t *testing.T) {
	b := params.NewPhotosEditAlbumBulder()

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

func TestPhotosEditCommentBulder(t *testing.T) {
	b := params.NewPhotosEditCommentBulder()

	b.OwnerID(1)
	b.CommentID(1)
	b.Message("text")
	b.Attachments([]string{"text"})

	assert.Equal(t, b.Params["owner_id"], 1)
	assert.Equal(t, b.Params["comment_id"], 1)
	assert.Equal(t, b.Params["message"], "text")
	assert.Equal(t, b.Params["attachments"], []string{"text"})
}

func TestPhotosGetBulder(t *testing.T) {
	b := params.NewPhotosGetBulder()

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

func TestPhotosGetAlbumsBulder(t *testing.T) {
	b := params.NewPhotosGetAlbumsBulder()

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

func TestPhotosGetAlbumsCountBulder(t *testing.T) {
	b := params.NewPhotosGetAlbumsCountBulder()

	b.UserID(1)
	b.GroupID(1)

	assert.Equal(t, b.Params["user_id"], 1)
	assert.Equal(t, b.Params["group_id"], 1)
}

func TestPhotosGetAllBulder(t *testing.T) {
	b := params.NewPhotosGetAllBulder()

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

func TestPhotosGetAllCommentsBulder(t *testing.T) {
	b := params.NewPhotosGetAllCommentsBulder()

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

func TestPhotosGetByIDBulder(t *testing.T) {
	b := params.NewPhotosGetByIDBulder()

	b.Photos([]string{"text"})
	b.Extended(true)
	b.PhotoSizes(true)

	assert.Equal(t, b.Params["photos"], []string{"text"})
	assert.Equal(t, b.Params["extended"], true)
	assert.Equal(t, b.Params["photo_sizes"], true)
}

func TestPhotosGetChatUploadServerBulder(t *testing.T) {
	b := params.NewPhotosGetChatUploadServerBulder()

	b.ChatID(1)
	b.CropX(1)
	b.CropY(1)
	b.CropWidth(1)

	assert.Equal(t, b.Params["chat_id"], 1)
	assert.Equal(t, b.Params["crop_x"], 1)
	assert.Equal(t, b.Params["crop_y"], 1)
	assert.Equal(t, b.Params["crop_width"], 1)
}

func TestPhotosGetCommentsBulder(t *testing.T) {
	b := params.NewPhotosGetCommentsBulder()

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

func TestPhotosGetMarketAlbumUploadServerBulder(t *testing.T) {
	b := params.NewPhotosGetMarketAlbumUploadServerBulder()

	b.GroupID(1)

	assert.Equal(t, b.Params["group_id"], 1)
}

func TestPhotosGetMarketUploadServerBulder(t *testing.T) {
	b := params.NewPhotosGetMarketUploadServerBulder()

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

func TestPhotosGetMessagesUploadServerBulder(t *testing.T) {
	b := params.NewPhotosGetMessagesUploadServerBulder()

	b.PeerID(1)

	assert.Equal(t, b.Params["peer_id"], 1)
}

func TestPhotosGetNewTagsBulder(t *testing.T) {
	b := params.NewPhotosGetNewTagsBulder()

	b.Offset(1)
	b.Count(1)

	assert.Equal(t, b.Params["offset"], 1)
	assert.Equal(t, b.Params["count"], 1)
}

func TestPhotosGetOwnerCoverPhotoUploadServerBulder(t *testing.T) {
	b := params.NewPhotosGetOwnerCoverPhotoUploadServerBulder()

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

func TestPhotosGetOwnerPhotoUploadServerBulder(t *testing.T) {
	b := params.NewPhotosGetOwnerPhotoUploadServerBulder()

	b.OwnerID(1)

	assert.Equal(t, b.Params["owner_id"], 1)
}

func TestPhotosGetTagsBulder(t *testing.T) {
	b := params.NewPhotosGetTagsBulder()

	b.OwnerID(1)
	b.PhotoID(1)
	b.AccessKey("text")

	assert.Equal(t, b.Params["owner_id"], 1)
	assert.Equal(t, b.Params["photo_id"], 1)
	assert.Equal(t, b.Params["access_key"], "text")
}

func TestPhotosGetUploadServerBulder(t *testing.T) {
	b := params.NewPhotosGetUploadServerBulder()

	b.GroupID(1)
	b.AlbumID(1)

	assert.Equal(t, b.Params["group_id"], 1)
	assert.Equal(t, b.Params["album_id"], 1)
}

func TestPhotosGetUserPhotosBulder(t *testing.T) {
	b := params.NewPhotosGetUserPhotosBulder()

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

func TestPhotosGetWallUploadServerBulder(t *testing.T) {
	b := params.NewPhotosGetWallUploadServerBulder()

	b.GroupID(1)

	assert.Equal(t, b.Params["group_id"], 1)
}

func TestPhotosMakeCoverBulder(t *testing.T) {
	b := params.NewPhotosMakeCoverBulder()

	b.OwnerID(1)
	b.PhotoID(1)
	b.AlbumID(1)

	assert.Equal(t, b.Params["owner_id"], 1)
	assert.Equal(t, b.Params["photo_id"], 1)
	assert.Equal(t, b.Params["album_id"], 1)
}

func TestPhotosMoveBulder(t *testing.T) {
	b := params.NewPhotosMoveBulder()

	b.OwnerID(1)
	b.TargetAlbumID(1)
	b.PhotoID(1)

	assert.Equal(t, b.Params["owner_id"], 1)
	assert.Equal(t, b.Params["target_album_id"], 1)
	assert.Equal(t, b.Params["photo_id"], 1)
}

func TestPhotosPutTagBulder(t *testing.T) {
	b := params.NewPhotosPutTagBulder()

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

func TestPhotosRemoveTagBulder(t *testing.T) {
	b := params.NewPhotosRemoveTagBulder()

	b.OwnerID(1)
	b.PhotoID(1)
	b.TagID(1)

	assert.Equal(t, b.Params["owner_id"], 1)
	assert.Equal(t, b.Params["photo_id"], 1)
	assert.Equal(t, b.Params["tag_id"], 1)
}

func TestPhotosReorderAlbumsBulder(t *testing.T) {
	b := params.NewPhotosReorderAlbumsBulder()

	b.OwnerID(1)
	b.AlbumID(1)
	b.Before(1)
	b.After(1)

	assert.Equal(t, b.Params["owner_id"], 1)
	assert.Equal(t, b.Params["album_id"], 1)
	assert.Equal(t, b.Params["before"], 1)
	assert.Equal(t, b.Params["after"], 1)
}

func TestPhotosReorderPhotosBulder(t *testing.T) {
	b := params.NewPhotosReorderPhotosBulder()

	b.OwnerID(1)
	b.PhotoID(1)
	b.Before(1)
	b.After(1)

	assert.Equal(t, b.Params["owner_id"], 1)
	assert.Equal(t, b.Params["photo_id"], 1)
	assert.Equal(t, b.Params["before"], 1)
	assert.Equal(t, b.Params["after"], 1)
}

func TestPhotosReportBulder(t *testing.T) {
	b := params.NewPhotosReportBulder()

	b.OwnerID(1)
	b.PhotoID(1)
	b.Reason(1)

	assert.Equal(t, b.Params["owner_id"], 1)
	assert.Equal(t, b.Params["photo_id"], 1)
	assert.Equal(t, b.Params["reason"], 1)
}

func TestPhotosReportCommentBulder(t *testing.T) {
	b := params.NewPhotosReportCommentBulder()

	b.OwnerID(1)
	b.CommentID(1)
	b.Reason(1)

	assert.Equal(t, b.Params["owner_id"], 1)
	assert.Equal(t, b.Params["comment_id"], 1)
	assert.Equal(t, b.Params["reason"], 1)
}

func TestPhotosRestoreBulder(t *testing.T) {
	b := params.NewPhotosRestoreBulder()

	b.OwnerID(1)
	b.PhotoID(1)

	assert.Equal(t, b.Params["owner_id"], 1)
	assert.Equal(t, b.Params["photo_id"], 1)
}

func TestPhotosRestoreCommentBulder(t *testing.T) {
	b := params.NewPhotosRestoreCommentBulder()

	b.OwnerID(1)
	b.CommentID(1)

	assert.Equal(t, b.Params["owner_id"], 1)
	assert.Equal(t, b.Params["comment_id"], 1)
}

func TestPhotosSaveBulder(t *testing.T) {
	b := params.NewPhotosSaveBulder()

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

func TestPhotosSaveMarketAlbumPhotoBulder(t *testing.T) {
	b := params.NewPhotosSaveMarketAlbumPhotoBulder()

	b.GroupID(1)
	b.Photo("text")
	b.Server(1)
	b.Hash("text")

	assert.Equal(t, b.Params["group_id"], 1)
	assert.Equal(t, b.Params["photo"], "text")
	assert.Equal(t, b.Params["server"], 1)
	assert.Equal(t, b.Params["hash"], "text")
}

func TestPhotosSaveMarketPhotoBulder(t *testing.T) {
	b := params.NewPhotosSaveMarketPhotoBulder()

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

func TestPhotosSaveMessagesPhotoBulder(t *testing.T) {
	b := params.NewPhotosSaveMessagesPhotoBulder()

	b.Photo("text")
	b.Server(1)
	b.Hash("text")

	assert.Equal(t, b.Params["photo"], "text")
	assert.Equal(t, b.Params["server"], 1)
	assert.Equal(t, b.Params["hash"], "text")
}

func TestPhotosSaveOwnerCoverPhotoBulder(t *testing.T) {
	b := params.NewPhotosSaveOwnerCoverPhotoBulder()

	b.Hash("text")
	b.Photo("text")

	assert.Equal(t, b.Params["hash"], "text")
	assert.Equal(t, b.Params["photo"], "text")
}

func TestPhotosSaveOwnerPhotoBulder(t *testing.T) {
	b := params.NewPhotosSaveOwnerPhotoBulder()

	b.Server("text")
	b.Hash("text")
	b.Photo("text")

	assert.Equal(t, b.Params["server"], "text")
	assert.Equal(t, b.Params["hash"], "text")
	assert.Equal(t, b.Params["photo"], "text")
}

func TestPhotosSaveWallPhotoBulder(t *testing.T) {
	b := params.NewPhotosSaveWallPhotoBulder()

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

func TestPhotosSearchBulder(t *testing.T) {
	b := params.NewPhotosSearchBulder()

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
