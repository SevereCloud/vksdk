package api

import (
	"strconv"
	"testing"
)

func TestVK_PhotosConfirmTag(t *testing.T) {
	if vkUser.AccessToken == "" {
		t.Skip("USER_TOKEN empty")
	}

	_, gotVkErr := vkUser.PhotosConfirmTag(map[string]string{
		"photo_id": "1234",
		"tag_id":   "5678",
	})
	if gotVkErr.Code != 0 {
		t.Errorf("VK.PhotosConfirmTag() gotVkErr = %v, want %v", gotVkErr, 0)
	}
}

func TestVK_PhotosCopy(t *testing.T) {
	if vkUser.AccessToken == "" {
		t.Skip("USER_TOKEN empty")
	}

	_, gotVkErr := vkUser.PhotosCopy(map[string]string{
		"owner_id": "1",
		"photo_id": "306810815",
	})
	if gotVkErr.Code != 0 {
		t.Errorf("VK.PhotosCopy() gotVkErr = %v, want %v", gotVkErr, 0)
	}
}

func TestVK_PhotosCreateAlbum(t *testing.T) {
	if vkUser.AccessToken == "" {
		t.Skip("USER_TOKEN empty")
	}

	album, gotVkErr := vkUser.PhotosCreateAlbum(map[string]string{
		"title": "TestAlbum",
	})
	if gotVkErr.Code != 0 {
		t.Errorf("VK.PhotosCreateAlbum() gotVkErr = %v, want %v", gotVkErr, 0)
	}

	_, gotVkErr = vkUser.PhotosEditAlbum(map[string]string{
		"album_id": strconv.Itoa(album.ID),
		"title":    "TestAlbum edited",
	})
	if gotVkErr.Code != 0 {
		t.Errorf("VK.PhotosEditAlbum() gotVkErr = %v, want %v", gotVkErr, 0)
	}

	_, gotVkErr = vkUser.PhotosDeleteAlbum(map[string]string{
		"album_id": strconv.Itoa(album.ID),
	})
	if gotVkErr.Code != 0 {
		t.Errorf("VK.PhotosDeleteAlbum() gotVkErr = %v, want %v", gotVkErr, 0)
	}
}

func TestVK_PhotosCreateComment(t *testing.T) {
	if vkUser.AccessToken == "" {
		t.Skip("USER_TOKEN empty")
	}

	comment, gotVkErr := vkUser.PhotosCreateComment(map[string]string{
		"owner_id": "540036751",
		"photo_id": "457239020",
		"message":  "Test photo comment",
	})
	if gotVkErr.Code != 0 {
		t.Errorf("VK.PhotosCreateComment() gotVkErr = %v, want %v", gotVkErr, 0)
	}
	commentID := int(comment)

	_, gotVkErr = vkUser.PhotosEditComment(map[string]string{
		"owner_id":   "540036751",
		"comment_id": strconv.Itoa(commentID),
		"message":    "Test photo comment edited",
	})
	if gotVkErr.Code != 0 {
		t.Errorf("VK.PhotosEditComment() gotVkErr = %v, want %v", gotVkErr, 0)
	}

	_, gotVkErr = vkUser.PhotosDeleteComment(map[string]string{
		"owner_id":   "540036751",
		"comment_id": strconv.Itoa(commentID),
	})
	if gotVkErr.Code != 0 {
		t.Errorf("VK.PhotosDeleteComment() gotVkErr = %v, want %v", gotVkErr, 0)
	}

	_, gotVkErr = vkUser.PhotosRestoreComment(map[string]string{
		"owner_id":   "540036751",
		"comment_id": strconv.Itoa(commentID),
	})
	if gotVkErr.Code != 0 {
		t.Errorf("VK.PhotosRestoreComment() gotVkErr = %v, want %v", gotVkErr, 0)
	}
}

// TODO: TestVK_PhotosDelete
// TODO: TestVK_PhotosEdit

func TestVK_PhotosGet(t *testing.T) {
	if vkUser.AccessToken == "" {
		t.Skip("USER_TOKEN empty")
	}

	params := map[string]string{
		"owner_id": "-1",
		"album_id": "wall",
	}

	_, gotVkErr := vkUser.PhotosGet(params)
	if gotVkErr.Code != 0 {
		t.Errorf("VK.PhotosGet() gotVkErr = %v, want %v", gotVkErr, 0)
	}

	_, gotVkErr = vkUser.PhotosGetExtended(params)
	if gotVkErr.Code != 0 {
		t.Errorf("VK.PhotosGetExtended() gotVkErr = %v, want %v", gotVkErr, 0)
	}
}

func TestVK_PhotosGetAlbums(t *testing.T) {
	if vkUser.AccessToken == "" {
		t.Skip("USER_TOKEN empty")
	}

	_, gotVkErr := vkUser.PhotosGetAlbums(map[string]string{
		"owner_id":    "185014513",
		"need_system": "1",
		"need_covers": "1",
		"photo_sizes": "1",
	})
	if gotVkErr.Code != 0 {
		t.Errorf("VK.PhotosGetAlbums() gotVkErr = %v, want %v", gotVkErr, 0)
	}
}

func TestVK_PhotosGetAlbumsCount(t *testing.T) {
	if vkUser.AccessToken == "" {
		t.Skip("USER_TOKEN empty")
	}

	_, gotVkErr := vkUser.PhotosGetAlbumsCount(map[string]string{
		"user_id": "185014513",
	})
	if gotVkErr.Code != 0 {
		t.Errorf("VK.PhotosGetAlbumsCount() gotVkErr = %v, want %v", gotVkErr, 0)
	}
}

func TestVK_PhotosGetAll(t *testing.T) {
	if vkUser.AccessToken == "" {
		t.Skip("USER_TOKEN empty")
	}

	params := map[string]string{
		"owner_id":    "1",
		"photo_sizes": "1",
	}

	_, gotVkErr := vkUser.PhotosGetAll(params)
	if gotVkErr.Code != 0 {
		t.Errorf("VK.PhotosGetAll() gotVkErr = %v, want %v", gotVkErr, 0)
	}

	_, gotVkErr = vkUser.PhotosGetAllExtended(params)
	if gotVkErr.Code != 0 {
		t.Errorf("VK.PhotosGetAllExtended() gotVkErr = %v, want %v", gotVkErr, 0)
	}
}

func TestVK_PhotosGetAllComments(t *testing.T) {
	if vkUser.AccessToken == "" {
		t.Skip("USER_TOKEN empty")
	}

	_, gotVkErr := vkUser.PhotosGetAllComments(map[string]string{
		"owner_id": "66748",
	})
	if gotVkErr.Code != 0 {
		t.Errorf("VK.PhotosGetAllComments() gotVkErr = %v, want %v", gotVkErr, 0)
	}
}

func TestVK_PhotosGetByID(t *testing.T) {
	if vkUser.AccessToken == "" {
		t.Skip("USER_TOKEN empty")
	}

	params := map[string]string{
		"photos": "1_278184324,1_295770249",
	}

	_, gotVkErr := vkUser.PhotosGetByID(params)
	if gotVkErr.Code != 0 {
		t.Errorf("VK.PhotosGetByID() gotVkErr = %v, want %v", gotVkErr, 0)
	}
	_, gotVkErr = vkUser.PhotosGetByIDExtended(params)
	if gotVkErr.Code != 0 {
		t.Errorf("VK.PhotosGetByIDExtended() gotVkErr = %v, want %v", gotVkErr, 0)
	}
}

func TestVK_PhotosGetChatUploadServer(t *testing.T) {
	if vkUser.AccessToken == "" {
		t.Skip("USER_TOKEN empty")
	}

	_, gotVkErr := vkUser.PhotosGetChatUploadServer(map[string]string{
		"chat_id": "1",
	})
	if gotVkErr.Code != 0 {
		t.Errorf("VK.PhotosGetChatUploadServer() gotVkErr = %v, want %v", gotVkErr, 0)
	}
}

func TestVK_PhotosGetComments(t *testing.T) {
	if vkUser.AccessToken == "" {
		t.Skip("USER_TOKEN empty")
	}

	params := map[string]string{
		"owner_id":   "1",
		"photo_id":   "456264771",
		"need_likes": "1",
	}

	_, gotVkErr := vkUser.PhotosGetComments(params)
	if gotVkErr.Code != 0 {
		t.Errorf("VK.PhotosGetComments() gotVkErr = %v, want %v", gotVkErr, 0)
	}

	_, gotVkErr = vkUser.PhotosGetCommentsExtended(params)
	if gotVkErr.Code != 0 {
		t.Errorf("VK.PhotosGetCommentsExtended() gotVkErr = %v, want %v", gotVkErr, 0)
	}
}

func TestVK_PhotosGetMarketAlbumUploadServer(t *testing.T) {
	if vkUser.AccessToken == "" || vkGroupID == 0 {
		t.Skip("USER_TOKEN empty")
	}

	_, gotVkErr := vkUser.PhotosGetMarketAlbumUploadServer(map[string]string{
		"group_id": strconv.Itoa(vkGroupID),
	})
	if gotVkErr.Code != 0 {
		t.Errorf("VK.PhotosGetMarketAlbumUploadServer() gotVkErr = %v, want %v", gotVkErr, 0)
	}
}

func TestVK_PhotosGetMarketUploadServer(t *testing.T) {
	if vkUser.AccessToken == "" || vkGroupID == 0 {
		t.Skip("USER_TOKEN empty")
	}

	_, gotVkErr := vkUser.PhotosGetMarketUploadServer(map[string]string{
		"group_id": strconv.Itoa(vkGroupID),
	})
	if gotVkErr.Code != 0 {
		t.Errorf("VK.PhotosGetMarketUploadServer() gotVkErr = %v, want %v", gotVkErr, 0)
	}
}

func TestVK_PhotosGetMessagesUploadServer(t *testing.T) {
	if vkUser.AccessToken == "" {
		t.Skip("USER_TOKEN empty")
	}

	_, gotVkErr := vkUser.PhotosGetMessagesUploadServer(map[string]string{
		"peer_id": "1",
	})
	if gotVkErr.Code != 0 {
		t.Errorf("VK.PhotosGetMessagesUploadServer() gotVkErr = %v, want %v", gotVkErr, 0)
	}
}

func TestVK_PhotosGetNewTags(t *testing.T) {
	if vkUser.AccessToken == "" {
		t.Skip("USER_TOKEN empty")
	}

	_, gotVkErr := vkUser.PhotosGetNewTags(map[string]string{})
	if gotVkErr.Code != 0 {
		t.Errorf("VK.PhotosGetNewTags() gotVkErr = %v, want %v", gotVkErr, 0)
	}
}

func TestVK_PhotosGetOwnerCoverPhotoUploadServer(t *testing.T) {
	if vkUser.AccessToken == "" || vkGroupID == 0 {
		t.Skip("USER_TOKEN empty")
	}

	_, gotVkErr := vkUser.PhotosGetOwnerCoverPhotoUploadServer(map[string]string{
		"group_id": strconv.Itoa(vkGroupID),
	})
	if gotVkErr.Code != 0 {
		t.Errorf("VK.PhotosGetOwnerCoverPhotoUploadServer() gotVkErr = %v, want %v", gotVkErr, 0)
	}
}

func TestVK_PhotosGetOwnerPhotoUploadServer(t *testing.T) {
	if vkUser.AccessToken == "" {
		t.Skip("USER_TOKEN empty")
	}

	_, gotVkErr := vkUser.PhotosGetOwnerPhotoUploadServer(map[string]string{})
	if gotVkErr.Code != 0 {
		t.Errorf("VK.PhotosGetOwnerPhotoUploadServer() gotVkErr = %v, want %v", gotVkErr, 0)
	}
}

func TestVK_PhotosGetTags(t *testing.T) {
	if vkUser.AccessToken == "" {
		t.Skip("USER_TOKEN empty")
	}

	_, gotVkErr := vkUser.PhotosGetTags(map[string]string{
		"owner_id": "3420",
		"photo_id": "341642982",
	})
	if gotVkErr.Code != 0 {
		t.Errorf("VK.PhotosGetTags() gotVkErr = %v, want %v", gotVkErr, 0)
	}
}

func TestVK_PhotosGetUploadServer(t *testing.T) {
	if vkUser.AccessToken == "" {
		t.Skip("USER_TOKEN empty")
	}

	_, gotVkErr := vkUser.PhotosGetUploadServer(map[string]string{
		"group_id": "25557243",
		"album_id": "156000927",
	})
	if gotVkErr.Code != 0 {
		t.Errorf("VK.PhotosGetUploadServer() gotVkErr = %v, want %v", gotVkErr, 0)
	}
}

func TestVK_PhotosGetUserPhotos(t *testing.T) {
	if vkUser.AccessToken == "" {
		t.Skip("USER_TOKEN empty")
	}

	params := map[string]string{
		"user_id": "110794238",
	}

	_, gotVkErr := vkUser.PhotosGetUserPhotos(params)
	if gotVkErr.Code != 0 {
		t.Errorf("VK.PhotosGetUserPhotos() gotVkErr = %v, want %v", gotVkErr, 0)
	}

	_, gotVkErr = vkUser.PhotosGetUserPhotosExtended(params)
	if gotVkErr.Code != 0 {
		t.Errorf("VK.PhotosGetUserPhotosExtended() gotVkErr = %v, want %v", gotVkErr, 0)
	}
}

func TestVK_PhotosGetWallUploadServer(t *testing.T) {
	if vkUser.AccessToken == "" {
		t.Skip("USER_TOKEN empty")
	}

	_, gotVkErr := vkUser.PhotosGetWallUploadServer(map[string]string{
		"group_id": strconv.Itoa(vkGroupID),
	})
	if gotVkErr.Code != 0 {
		t.Errorf("VK.PhotosGetWallUploadServer() gotVkErr = %v, want %v", gotVkErr, 0)
	}
}

// TODO: TestVK_PhotosMakeCover
// TODO: TestVK_PhotosMove
// TODO: TestVK_PhotosPutTag
// TODO: TestVK_PhotosRemoveTag
// TODO: TestVK_PhotosReorderAlbums
// TODO: TestVK_PhotosReorderPhotos

func TestVK_PhotosReport(t *testing.T) {
	if vkUser.AccessToken == "" {
		t.Skip("USER_TOKEN empty")
	}

	_, gotVkErr := vkUser.PhotosReport(map[string]string{
		"owner_id": "66748",
		"photo_id": "312071876",
	})
	if gotVkErr.Code != 0 {
		t.Errorf("VK.PhotosReport() gotVkErr = %v, want %v", gotVkErr, 0)
	}
}

func TestVK_PhotosReportComment(t *testing.T) {
	if vkUser.AccessToken == "" {
		t.Skip("USER_TOKEN empty")
	}

	_, gotVkErr := vkUser.PhotosReportComment(map[string]string{
		"owner_id":   "66748",
		"comment_id": "823",
	})
	if gotVkErr.Code != 0 {
		t.Errorf("VK.PhotosReportComment() gotVkErr = %v, want %v", gotVkErr, 0)
	}
}

// TODO: TestVK_PhotosSave
// TODO: TestVK_PhotosSaveMarketAlbum
// TODO: TestVK_PhotosSaveMarketPhoto
// TODO: TestVK_PhotosSaveMessagesPhoto
// TODO: TestVK_PhotosSaveOwnerCoverPhoto
// TODO: TestVK_PhotosSaveOwnerPhoto
// TODO: TestVK_PhotosSaveWallPhoto

func TestVK_PhotosSearch(t *testing.T) {
	if vkUser.AccessToken == "" {
		t.Skip("USER_TOKEN empty")
	}

	_, gotVkErr := vkUser.PhotosSearch(map[string]string{
		"q": "Nature",
	})
	if gotVkErr.Code != 0 {
		t.Errorf("VK.PhotosSearch() gotVkErr = %v, want %v", gotVkErr, 0)
	}
}
