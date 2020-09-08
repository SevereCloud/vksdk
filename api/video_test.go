package api_test

import (
	"testing"
	"time"

	"github.com/SevereCloud/vksdk/v2/api"

	"github.com/stretchr/testify/assert"
)

func TestVK_VideoAddDelete(t *testing.T) {
	t.Parallel()

	needUserToken(t)

	res, err := vkUser.VideoAdd(api.Params{
		"target_id": vkUserID,
		"owner_id":  -139533130,
		"video_id":  456239332,
	})
	noError(t, err)
	assert.NotEmpty(t, res)

	res, err = vkUser.VideoDelete(api.Params{
		"target_id": vkUserID,
		"owner_id":  -139533130,
		"video_id":  456239332,
	})
	noError(t, err)
	assert.NotEmpty(t, res)
}

func TestVK_VideoAddAlbum(t *testing.T) {
	t.Parallel()

	needUserToken(t)

	album, err := vkUser.VideoAddAlbum(api.Params{
		"title": "Test",
	})
	noError(t, err)
	assert.NotEmpty(t, album.AlbumID)

	afterAlbum, err := vkUser.VideoAddAlbum(api.Params{
		"title": "Test2",
	})
	noError(t, err)

	time.Sleep(time.Second)

	_, err = vkUser.VideoAddToAlbum(api.Params{
		"album_id": album.AlbumID,
		"owner_id": -139533130,
		"video_id": 456239332,
	})
	noError(t, err)

	_, err = vkUser.VideoAddToAlbum(api.Params{
		"album_id": album.AlbumID,
		"owner_id": 2314852,
		"video_id": 165705596,
	})
	noError(t, err)

	_, err = vkUser.VideoReorderVideos(api.Params{
		"album_id":       album.AlbumID,
		"owner_id":       2314852,
		"video_id":       165705596,
		"after_owner_id": -139533130,
		"after_video_id": 456239332,
	})
	noError(t, err)

	_, err = vkUser.VideoReorderAlbums(api.Params{
		"album_id": album.AlbumID,
		"after":    afterAlbum.AlbumID,
	})
	noError(t, err)

	_, err = vkUser.VideoRemoveFromAlbum(api.Params{
		"album_id": album.AlbumID,
		"owner_id": 2314852,
		"video_id": 165705596,
	})
	noError(t, err)

	_, err = vkUser.VideoEditAlbum(api.Params{
		"album_id": album.AlbumID,
		"title":    "Test edit",
	})
	noError(t, err)

	_, err = vkUser.VideoDeleteAlbum(api.Params{
		"album_id": album.AlbumID,
	})
	noError(t, err)
}

func TestVK_VideoCreateComment(t *testing.T) {
	t.Parallel()

	needUserToken(t)

	commentID, err := vkUser.VideoCreateComment(api.Params{
		"owner_id": -169097025,
		"video_id": 456239034,
		"message":  "Test video.createComment",
	})
	noError(t, err)
	assert.NotEmpty(t, commentID)

	_, err = vkUser.VideoEditComment(api.Params{
		"owner_id":   -169097025,
		"comment_id": commentID,
		"message":    "Test video.editComment",
	})
	noError(t, err)

	_, err = vkUser.VideoDeleteComment(api.Params{
		"owner_id":   -169097025,
		"comment_id": commentID,
	})
	noError(t, err)

	_, err = vkUser.VideoRestoreComment(api.Params{
		"owner_id":   -169097025,
		"comment_id": commentID,
	})
	noError(t, err)
}

func TestVK_VideoEdit(t *testing.T) {
	// TODO: write test
}

func TestVK_VideoGet(t *testing.T) {
	t.Parallel()

	needUserToken(t)

	res, err := vkUser.VideoGet(api.Params{
		"owner_id": -76477496,
	})
	noError(t, err)

	assert.NotEmpty(t, res.Count)
	assert.NotEmpty(t, res.Items)
}

func TestVK_VideoGetExtended(t *testing.T) {
	t.Parallel()

	needUserToken(t)

	res, err := vkUser.VideoGetExtended(api.Params{
		// "owner_id": 92093600,
	})
	noError(t, err)

	assert.NotEmpty(t, res.Count)
	assert.NotEmpty(t, res.Items)
}

func TestVK_VideoGetAlbumByID(t *testing.T) {
	t.Parallel()

	needUserToken(t)

	res, err := vkUser.VideoGetAlbumByID(api.Params{
		"owner_id": 185014513,
		"album_id": 1,
	})
	noError(t, err)

	assert.NotEmpty(t, res.Count)
	assert.NotEmpty(t, res.OwnerID)
	assert.NotEmpty(t, res.Title)
	assert.NotEmpty(t, res.UpdatedTime)
	// assert.NotEmpty(t, res.Image)
	assert.NotEmpty(t, res.ID)
}

func TestVK_VideoGetAlbums(t *testing.T) {
	t.Parallel()

	needUserToken(t)

	res, err := vkUser.VideoGetAlbums(api.Params{
		"owner_id": 66748,
	})
	noError(t, err)

	assert.NotEmpty(t, res.Count)
	assert.NotEmpty(t, res.Items)
}

func TestVK_VideoGetAlbumsExtended(t *testing.T) {
	t.Parallel()

	needUserToken(t)

	res, err := vkUser.VideoGetAlbumsExtended(api.Params{
		"owner_id": 66748,
	})
	noError(t, err)

	assert.NotEmpty(t, res.Count)
	assert.NotEmpty(t, res.Items)
}

func TestVK_VideoGetAlbumsByVideo(t *testing.T) {
	t.Parallel()

	needUserToken(t)

	_, err := vkUser.VideoGetAlbumsByVideo(api.Params{
		"owner_id": 2314852,
		"video_id": 165705596,
	})
	noError(t, err)
}

func TestVK_VideoGetAlbumsByVideoExtended(t *testing.T) {
	t.Parallel()

	needUserToken(t)

	_, err := vkUser.VideoGetAlbumsByVideoExtended(api.Params{
		"owner_id": 2314852,
		"video_id": 165705596,
	})
	noError(t, err)
}

func TestVK_VideoGetComments(t *testing.T) {
	t.Parallel()

	needUserToken(t)

	res, err := vkUser.VideoGetComments(api.Params{
		"owner_id": -169097025,
		"video_id": 456239034,
	})
	noError(t, err)

	assert.NotEmpty(t, res.Count)
	assert.NotEmpty(t, res.Items)
}

func TestVK_VideoGetCommentsExtended(t *testing.T) {
	t.Parallel()

	needUserToken(t)

	res, err := vkUser.VideoGetCommentsExtended(api.Params{
		"owner_id": -169097025,
		"video_id": 456239034,
	})
	noError(t, err)

	assert.NotEmpty(t, res.Count)
	assert.NotEmpty(t, res.Items)
	assert.NotEmpty(t, res.Profiles)
}

func TestVK_VideoReport(t *testing.T) {
	t.Parallel()

	needUserToken(t)

	_, err := vkUser.VideoReport(api.Params{
		"owner_id": -45172787,
		"video_id": 163704979,
	})
	noError(t, err)
}

func TestVK_VideoReportComment(t *testing.T) {
	t.Parallel()

	needUserToken(t)

	_, err := vkUser.VideoReportComment(api.Params{
		"owner_id":   1,
		"comment_id": 17142,
	})
	noError(t, err)
}

func TestVK_VideoSearch(t *testing.T) {
	t.Parallel()

	needUserToken(t)

	res, err := vkUser.VideoSearch(api.Params{
		"q": "Nature",
	})
	noError(t, err)

	assert.NotEmpty(t, res.Count)
	assert.NotEmpty(t, res.Items)
}

func TestVK_VideoSearchExtended(t *testing.T) {
	t.Parallel()

	needUserToken(t)

	res, err := vkUser.VideoSearchExtended(api.Params{
		"q": "Nature",
	})
	noError(t, err)

	assert.NotEmpty(t, res.Count)
	assert.NotEmpty(t, res.Items)
}
