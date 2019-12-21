package api_test

import (
	"net/http"
	"testing"

	"github.com/SevereCloud/vksdk/api"
	"github.com/SevereCloud/vksdk/object"
)

func TestVK_PhotosConfirmTag(t *testing.T) {
	needUserToken(t)

	_, err := vkUser.PhotosConfirmTag(api.Params{
		"photo_id": 1234,
		"tag_id":   5678,
	})
	noError(t, err)
}

func TestVK_PhotosCopy(t *testing.T) {
	needUserToken(t)

	_, err := vkUser.PhotosCopy(api.Params{
		"owner_id": 1,
		"photo_id": 306810815,
	})
	noError(t, err)
}

func TestVK_PhotosCreateAlbum(t *testing.T) {
	needUserToken(t)

	album, err := vkUser.PhotosCreateAlbum(api.Params{
		"title": "TestAlbum",
	})
	noError(t, err)

	afterAlbum, err := vkUser.PhotosCreateAlbum(api.Params{
		"title": "TestAlbum",
	})
	noError(t, err)

	photo := LoadPhoto(t, album.ID)
	afterPhoto := LoadPhoto(t, album.ID)

	_, err = vkUser.PhotosMakeCover(api.Params{
		"photo_id": photo.ID,
		"album_id": album.ID,
	})
	noError(t, err)

	_, err = vkUser.PhotosReorderPhotos(api.Params{
		"photo_id": photo.ID,
		"after":    afterPhoto.ID,
	})
	noError(t, err)

	_, err = vkUser.PhotosMove(api.Params{
		"photo_id":        photo.ID,
		"target_album_id": afterAlbum.ID,
	})
	noError(t, err)

	_, err = vkUser.PhotosReorderAlbums(api.Params{
		"album_id": album.ID,
		"after":    afterAlbum.ID,
	})
	noError(t, err)

	_, err = vkUser.PhotosEditAlbum(api.Params{
		"album_id": album.ID,
		"title":    "TestAlbum edited",
	})
	noError(t, err)

	_, err = vkUser.PhotosDeleteAlbum(api.Params{
		"album_id": album.ID,
	})
	noError(t, err)

	_, err = vkUser.PhotosDeleteAlbum(api.Params{
		"album_id": afterAlbum.ID,
	})
	noError(t, err)
}

func TestVK_PhotosCreateComment(t *testing.T) {
	needUserToken(t)

	commentID, err := vkUser.PhotosCreateComment(api.Params{
		"owner_id": 233157978,
		"photo_id": 456250946,
		"message":  "Test photo comment",
	})
	noError(t, err)

	_, err = vkUser.PhotosEditComment(api.Params{
		"owner_id":   233157978,
		"comment_id": commentID,
		"message":    "Test photo comment edited",
	})
	noError(t, err)

	_, err = vkUser.PhotosDeleteComment(api.Params{
		"owner_id":   233157978,
		"comment_id": commentID,
	})
	noError(t, err)

	_, err = vkUser.PhotosRestoreComment(api.Params{
		"owner_id":   233157978,
		"comment_id": commentID,
	})
	noError(t, err)
}

func LoadPhoto(t *testing.T, albumID int) object.PhotosPhoto {
	response, err := http.Get(photoURL)
	if err != nil {
		t.Fatalf("%v", err)
	}
	defer response.Body.Close()

	photoSave, err := vkUser.UploadPhoto(albumID, response.Body)
	noError(t, err)

	photo := photoSave[0]

	tagID, err := vkUser.PhotosPutTag(api.Params{
		"photo_id": photo.ID,
		"user_id":  vkUserID,
		"x":        10,
		"y":        10,
		"x2":       80,
		"y2":       80,
	})
	noError(t, err)

	_, err = vkUser.PhotosRemoveTag(api.Params{
		"photo_id": photo.ID,
		"tag_id":   tagID,
	})
	noError(t, err)

	_, err = vkUser.PhotosEdit(api.Params{
		"photo_id": photo.ID,
		"caption":  "test",
	})
	noError(t, err)

	_, err = vkUser.PhotosDelete(api.Params{
		"photo_id": photo.ID,
	})
	noError(t, err)

	_, err = vkUser.PhotosRestore(api.Params{
		"photo_id": photo.ID,
	})
	noError(t, err)

	return photo
}

func TestVK_PhotosGet(t *testing.T) {
	needUserToken(t)

	params := api.Params{
		"owner_id": -1,
		"album_id": "wall",
	}

	_, err := vkUser.PhotosGet(params)
	noError(t, err)

	_, err = vkUser.PhotosGetExtended(params)
	noError(t, err)
}

func TestVK_PhotosGetAlbums(t *testing.T) {
	needUserToken(t)

	_, err := vkUser.PhotosGetAlbums(api.Params{
		"owner_id":    185014513,
		"need_system": true,
		"need_covers": true,
		"photo_sizes": true,
	})
	noError(t, err)
}

func TestVK_PhotosGetAlbumsCount(t *testing.T) {
	needUserToken(t)

	_, err := vkUser.PhotosGetAlbumsCount(api.Params{
		"user_id": 185014513,
	})
	noError(t, err)
}

func TestVK_PhotosGetAll(t *testing.T) {
	needUserToken(t)

	params := api.Params{
		"owner_id":    1,
		"photo_sizes": true,
	}

	_, err := vkUser.PhotosGetAll(params)
	noError(t, err)

	_, err = vkUser.PhotosGetAllExtended(params)
	noError(t, err)
}

func TestVK_PhotosGetAllComments(t *testing.T) {
	needUserToken(t)

	_, err := vkUser.PhotosGetAllComments(api.Params{
		"owner_id": 66748,
	})
	noError(t, err)
}

func TestVK_PhotosGetByID(t *testing.T) {
	needUserToken(t)

	params := api.Params{
		"photos": "1_278184324,1_295770249",
	}

	_, err := vkUser.PhotosGetByID(params)
	noError(t, err)
	_, err = vkUser.PhotosGetByIDExtended(params)
	noError(t, err)
}

func TestVK_PhotosGetChatUploadServer(t *testing.T) {
	needUserToken(t)

	_, err := vkUser.PhotosGetChatUploadServer(api.Params{
		"chat_id": 1,
	})
	noError(t, err)
}

func TestVK_PhotosGetComments(t *testing.T) {
	needUserToken(t)

	params := api.Params{
		"owner_id":   1,
		"photo_id":   456264771,
		"need_likes": true,
	}

	_, err := vkUser.PhotosGetComments(params)
	noError(t, err)

	_, err = vkUser.PhotosGetCommentsExtended(params)
	noError(t, err)
}

func TestVK_PhotosGetMarketAlbumUploadServer(t *testing.T) {
	needUserToken(t)
	needGroupToken(t)

	_, err := vkUser.PhotosGetMarketAlbumUploadServer(api.Params{
		"group_id": vkGroupID,
	})
	noError(t, err)
}

func TestVK_PhotosGetMarketUploadServer(t *testing.T) {
	needUserToken(t)
	needGroupToken(t)

	_, err := vkUser.PhotosGetMarketUploadServer(api.Params{
		"group_id": vkGroupID,
	})
	noError(t, err)
}

func TestVK_PhotosGetMessagesUploadServer(t *testing.T) {
	needUserToken(t)

	_, err := vkUser.PhotosGetMessagesUploadServer(api.Params{
		"peer_id": 1,
	})
	noError(t, err)
}

func TestVK_PhotosGetNewTags(t *testing.T) {
	needUserToken(t)

	_, err := vkUser.PhotosGetNewTags(api.Params{})
	noError(t, err)
}

func TestVK_PhotosGetOwnerCoverPhotoUploadServer(t *testing.T) {
	needUserToken(t)
	needGroupToken(t)

	_, err := vkUser.PhotosGetOwnerCoverPhotoUploadServer(api.Params{
		"group_id": vkGroupID,
	})
	noError(t, err)
}

func TestVK_PhotosGetOwnerPhotoUploadServer(t *testing.T) {
	needUserToken(t)

	_, err := vkUser.PhotosGetOwnerPhotoUploadServer(api.Params{})
	noError(t, err)
}

func TestVK_PhotosGetTags(t *testing.T) {
	needUserToken(t)

	_, err := vkUser.PhotosGetTags(api.Params{
		"owner_id": 3420,
		"photo_id": 341642982,
	})
	noError(t, err)
}

func TestVK_PhotosGetUploadServer(t *testing.T) {
	needUserToken(t)

	_, err := vkUser.PhotosGetUploadServer(api.Params{
		"group_id": 25557243,
		"album_id": 156000927,
	})
	noError(t, err)
}

func TestVK_PhotosGetUserPhotos(t *testing.T) {
	needUserToken(t)

	params := api.Params{
		"user_id": 110794238,
	}

	_, err := vkUser.PhotosGetUserPhotos(params)
	noError(t, err)

	_, err = vkUser.PhotosGetUserPhotosExtended(params)
	noError(t, err)
}

func TestVK_PhotosGetWallUploadServer(t *testing.T) {
	needUserToken(t)

	_, err := vkUser.PhotosGetWallUploadServer(api.Params{
		"group_id": vkGroupID,
	})
	noError(t, err)
}

func TestVK_PhotosReport(t *testing.T) {
	needUserToken(t)

	_, err := vkUser.PhotosReport(api.Params{
		"owner_id": 66748,
		"photo_id": 312071876,
	})
	noError(t, err)
}

func TestVK_PhotosReportComment(t *testing.T) {
	needUserToken(t)

	_, err := vkUser.PhotosReportComment(api.Params{
		"owner_id":   66748,
		"comment_id": 823,
	})
	noError(t, err)
}

func TestVK_PhotosSearch(t *testing.T) {
	needUserToken(t)

	_, err := vkUser.PhotosSearch(api.Params{
		"q": "Nature",
	})
	noError(t, err)
}
