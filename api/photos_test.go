package api

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestVK_PhotosConfirmTag(t *testing.T) {
	needUserToken(t)

	_, err := vkUser.PhotosConfirmTag(Params{
		"photo_id": 1234,
		"tag_id":   5678,
	})
	assert.NoError(t, err)
}

func TestVK_PhotosCopy(t *testing.T) {
	needUserToken(t)

	_, err := vkUser.PhotosCopy(Params{
		"owner_id": 1,
		"photo_id": 306810815,
	})
	assert.NoError(t, err)
}

func TestVK_PhotosCreateAlbum(t *testing.T) {
	needUserToken(t)

	album, err := vkUser.PhotosCreateAlbum(Params{
		"title": "TestAlbum",
	})
	assert.NoError(t, err)

	_, err = vkUser.PhotosEditAlbum(Params{
		"album_id": album.ID,
		"title":    "TestAlbum edited",
	})
	assert.NoError(t, err)

	_, err = vkUser.PhotosDeleteAlbum(Params{
		"album_id": album.ID,
	})
	assert.NoError(t, err)
}

func TestVK_PhotosCreateComment(t *testing.T) {
	needUserToken(t)

	commentID, err := vkUser.PhotosCreateComment(Params{
		"owner_id": 233157978,
		"photo_id": 456250946,
		"message":  "Test photo comment",
	})
	assert.NoError(t, err)

	_, err = vkUser.PhotosEditComment(Params{
		"owner_id":   233157978,
		"comment_id": commentID,
		"message":    "Test photo comment edited",
	})
	assert.NoError(t, err)

	_, err = vkUser.PhotosDeleteComment(Params{
		"owner_id":   233157978,
		"comment_id": commentID,
	})
	assert.NoError(t, err)

	_, err = vkUser.PhotosRestoreComment(Params{
		"owner_id":   233157978,
		"comment_id": commentID,
	})
	assert.NoError(t, err)
}

// TODO: TestVK_PhotosDelete
// TODO: TestVK_PhotosEdit

func TestVK_PhotosGet(t *testing.T) {
	needUserToken(t)

	params := Params{
		"owner_id": -1,
		"album_id": "wall",
	}

	_, err := vkUser.PhotosGet(params)
	assert.NoError(t, err)

	_, err = vkUser.PhotosGetExtended(params)
	assert.NoError(t, err)
}

func TestVK_PhotosGetAlbums(t *testing.T) {
	needUserToken(t)

	_, err := vkUser.PhotosGetAlbums(Params{
		"owner_id":    185014513,
		"need_system": true,
		"need_covers": true,
		"photo_sizes": true,
	})
	assert.NoError(t, err)
}

func TestVK_PhotosGetAlbumsCount(t *testing.T) {
	needUserToken(t)

	_, err := vkUser.PhotosGetAlbumsCount(Params{
		"user_id": 185014513,
	})
	assert.NoError(t, err)
}

func TestVK_PhotosGetAll(t *testing.T) {
	needUserToken(t)

	params := Params{
		"owner_id":    1,
		"photo_sizes": true,
	}

	_, err := vkUser.PhotosGetAll(params)
	assert.NoError(t, err)

	_, err = vkUser.PhotosGetAllExtended(params)
	assert.NoError(t, err)
}

func TestVK_PhotosGetAllComments(t *testing.T) {
	needUserToken(t)

	_, err := vkUser.PhotosGetAllComments(Params{
		"owner_id": 66748,
	})
	assert.NoError(t, err)
}

func TestVK_PhotosGetByID(t *testing.T) {
	needUserToken(t)

	params := Params{
		"photos": "1_278184324,1_295770249",
	}

	_, err := vkUser.PhotosGetByID(params)
	assert.NoError(t, err)
	_, err = vkUser.PhotosGetByIDExtended(params)
	assert.NoError(t, err)
}

func TestVK_PhotosGetChatUploadServer(t *testing.T) {
	needUserToken(t)

	_, err := vkUser.PhotosGetChatUploadServer(Params{
		"chat_id": 1,
	})
	assert.NoError(t, err)
}

func TestVK_PhotosGetComments(t *testing.T) {
	needUserToken(t)

	params := Params{
		"owner_id":   1,
		"photo_id":   456264771,
		"need_likes": true,
	}

	_, err := vkUser.PhotosGetComments(params)
	assert.NoError(t, err)

	_, err = vkUser.PhotosGetCommentsExtended(params)
	assert.NoError(t, err)
}

func TestVK_PhotosGetMarketAlbumUploadServer(t *testing.T) {
	needUserToken(t)
	needGroupToken(t)

	_, err := vkUser.PhotosGetMarketAlbumUploadServer(Params{
		"group_id": vkGroupID,
	})
	assert.NoError(t, err)
}

func TestVK_PhotosGetMarketUploadServer(t *testing.T) {
	needUserToken(t)
	needGroupToken(t)

	_, err := vkUser.PhotosGetMarketUploadServer(Params{
		"group_id": vkGroupID,
	})
	assert.NoError(t, err)
}

func TestVK_PhotosGetMessagesUploadServer(t *testing.T) {
	needUserToken(t)

	_, err := vkUser.PhotosGetMessagesUploadServer(Params{
		"peer_id": 1,
	})
	assert.NoError(t, err)
}

func TestVK_PhotosGetNewTags(t *testing.T) {
	needUserToken(t)

	_, err := vkUser.PhotosGetNewTags(Params{})
	assert.NoError(t, err)
}

func TestVK_PhotosGetOwnerCoverPhotoUploadServer(t *testing.T) {
	needUserToken(t)
	needGroupToken(t)

	_, err := vkUser.PhotosGetOwnerCoverPhotoUploadServer(Params{
		"group_id": vkGroupID,
	})
	assert.NoError(t, err)
}

func TestVK_PhotosGetOwnerPhotoUploadServer(t *testing.T) {
	needUserToken(t)

	_, err := vkUser.PhotosGetOwnerPhotoUploadServer(Params{})
	assert.NoError(t, err)
}

func TestVK_PhotosGetTags(t *testing.T) {
	needUserToken(t)

	_, err := vkUser.PhotosGetTags(Params{
		"owner_id": 3420,
		"photo_id": 341642982,
	})
	assert.NoError(t, err)
}

func TestVK_PhotosGetUploadServer(t *testing.T) {
	needUserToken(t)

	_, err := vkUser.PhotosGetUploadServer(Params{
		"group_id": 25557243,
		"album_id": 156000927,
	})
	assert.NoError(t, err)
}

func TestVK_PhotosGetUserPhotos(t *testing.T) {
	needUserToken(t)

	params := Params{
		"user_id": 110794238,
	}

	_, err := vkUser.PhotosGetUserPhotos(params)
	assert.NoError(t, err)

	_, err = vkUser.PhotosGetUserPhotosExtended(params)
	assert.NoError(t, err)
}

func TestVK_PhotosGetWallUploadServer(t *testing.T) {
	needUserToken(t)

	_, err := vkUser.PhotosGetWallUploadServer(Params{
		"group_id": vkGroupID,
	})
	assert.NoError(t, err)
}

// TODO: TestVK_PhotosMakeCover
// TODO: TestVK_PhotosMove
// TODO: TestVK_PhotosPutTag
// TODO: TestVK_PhotosRemoveTag
// TODO: TestVK_PhotosReorderAlbums
// TODO: TestVK_PhotosReorderPhotos

func TestVK_PhotosReport(t *testing.T) {
	needUserToken(t)

	_, err := vkUser.PhotosReport(Params{
		"owner_id": 66748,
		"photo_id": 312071876,
	})
	assert.NoError(t, err)
}

func TestVK_PhotosReportComment(t *testing.T) {
	needUserToken(t)

	_, err := vkUser.PhotosReportComment(Params{
		"owner_id":   66748,
		"comment_id": 823,
	})
	assert.NoError(t, err)
}

func TestVK_PhotosSearch(t *testing.T) {
	needUserToken(t)

	_, err := vkUser.PhotosSearch(Params{
		"q": "Nature",
	})
	assert.NoError(t, err)
}
