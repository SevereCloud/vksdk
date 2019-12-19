package api_test

import (
	"testing"

	"github.com/SevereCloud/vksdk/api"

	"github.com/stretchr/testify/assert"
)

func TestVK_VideoAddDeleteRestore(t *testing.T) {
	needUserToken(t)

	res, err := vkUser.VideoAdd(api.Params{
		"target_id": vkUserID,
		"owner_id":  -139533130,
		"video_id":  456239332,
	})
	assert.NoError(t, err)
	assert.NotEmpty(t, res)

	res, err = vkUser.VideoDelete(api.Params{
		"target_id": vkUserID,
		"owner_id":  -139533130,
		"video_id":  456239332,
	})
	assert.NoError(t, err)
	assert.NotEmpty(t, res)

	res, err = vkUser.VideoRestore(api.Params{
		"target_id": vkUserID,
		"owner_id":  -139533130,
		"video_id":  456239332,
	})
	assert.NoError(t, err)
	assert.NotEmpty(t, res)

	_, _ = vkUser.VideoDelete(api.Params{
		"target_id": vkUserID,
		"owner_id":  -139533130,
		"video_id":  456239332,
	})
}

func TestVK_VideoAddAlbum(t *testing.T) {
	album, err := vkUser.VideoAddAlbum(api.Params{
		"title": "Test",
	})
	assert.NoError(t, err)
	assert.NotEmpty(t, album.AlbumID)

	res, err := vkUser.VideoEditAlbum(api.Params{
		"album_id": album.AlbumID,
		"title":    "Test edit",
	})
	assert.NoError(t, err)
	assert.NotEmpty(t, res)

	res, err = vkUser.VideoDeleteAlbum(api.Params{
		"album_id": album.AlbumID,
	})
	assert.NoError(t, err)
	assert.NotEmpty(t, res)
}

func TestVK_VideoAddToAlbum(t *testing.T) {
	// TODO: write test
}

func TestVK_VideoCreateComment(t *testing.T) {
	// TODO: write test
}

func TestVK_VideoDeleteComment(t *testing.T) {
	// TODO: write test
}

func TestVK_VideoEdit(t *testing.T) {
	// TODO: write test
}

func TestVK_VideoEditComment(t *testing.T) {
	// TODO: write test
}

func TestVK_VideoGet(t *testing.T) {
	// TODO: write test
}

func TestVK_VideoGetExtended(t *testing.T) {
	// TODO: write test
}

func TestVK_VideoGetAlbumByID(t *testing.T) {
	// TODO: write test
}

func TestVK_VideoGetAlbums(t *testing.T) {
	// TODO: write test
}

func TestVK_VideoGetAlbumsExtended(t *testing.T) {
	// TODO: write test
}

func TestVK_VideoGetAlbumsByVideo(t *testing.T) {
	// TODO: write test
}

func TestVK_VideoGetAlbumsByVideoExtended(t *testing.T) {
	// TODO: write test
}

func TestVK_VideoGetComments(t *testing.T) {
	// TODO: write test
}

func TestVK_VideoGetCommentsExtended(t *testing.T) {
	// TODO: write test
}

func TestVK_VideoRemoveFromAlbum(t *testing.T) {
	// TODO: write test
}

func TestVK_VideoReorderAlbums(t *testing.T) {
	// TODO: write test
}

func TestVK_VideoReorderVideos(t *testing.T) {
	// TODO: write test
}

func TestVK_VideoReport(t *testing.T) {
	// TODO: write test
}

func TestVK_VideoReportComment(t *testing.T) {
	// TODO: write test
}

func TestVK_VideoRestoreComment(t *testing.T) {
	// TODO: write test
}

func TestVK_VideoSave(t *testing.T) {
	// TODO: write test
}

func TestVK_VideoSearch(t *testing.T) {
	// TODO: write test
}

func TestVK_VideoSearchExtended(t *testing.T) {
	// TODO: write test
}
