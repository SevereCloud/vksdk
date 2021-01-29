package api_test

import (
	"bytes"
	"io"
	"net/http"
	"testing"

	"github.com/SevereCloud/vksdk/v2/api"
	"github.com/stretchr/testify/assert"
)

const (
	photoURL        = "https://sun9-17.userapi.com/c853620/v853620933/dedb8/_5CIRVR-UA8.jpg"
	photo480x480URL = "https://sun9-64.userapi.com/6MiMozkfYIEYi9p_GSdUUHNNV_H-D84aNLSabQ/59Rpw9oZq3M.jpg"
)

func TestVK_UploadFile(t *testing.T) {
	t.Parallel()

	vk := api.NewVK("")

	f := func(url string, file io.Reader, fieldname, filename string, needErr bool) {
		t.Helper()

		_, err := vk.UploadFile(url, file, fieldname, filename)
		if (err != nil) && !needErr {
			t.Errorf("VK.UploadFile() err = %v, want %v", err, needErr)
		}
	}

	f("", new(bytes.Buffer), "", "", true)
}

func TestVK_UploadPhoto(t *testing.T) {
	t.Parallel()

	needUserToken(t)

	album, err := vkUser.PhotosCreateAlbum(api.Params{
		"title": "test",
	})
	if err != nil {
		t.Fatalf("VK.PhotosCreateAlbum() err = %v", err)
	}

	response, err := http.Get(photoURL)
	if err != nil {
		t.Fatalf("%v", err)
	}
	defer response.Body.Close()

	_, err = vkUser.UploadPhoto(album.ID, response.Body)
	noError(t, err)
}

func TestVK_UploadPhotoGroup(t *testing.T) {
	t.Parallel()

	needUserToken(t)
	needGroupToken(t)

	album, err := vkUser.PhotosCreateAlbum(api.Params{
		"title":    "test",
		"group_id": vkGroupID,
	})
	if err != nil {
		t.Fatalf("VK.PhotosCreateAlbum() err = %v", err)
	}

	response, err := http.Get(photoURL)
	if err != nil {
		t.Fatalf("%v", err)
	}
	defer response.Body.Close()

	_, err = vkUser.UploadPhotoGroup(vkGroupID, album.ID, response.Body)
	noError(t, err)
}

func TestVK_UploadWallPhoto(t *testing.T) {
	t.Parallel()

	needUserToken(t)

	response, err := http.Get(photoURL)
	if err != nil {
		t.Fatalf("%v", err)
	}
	defer response.Body.Close()

	_, err = vkUser.UploadWallPhoto(response.Body)
	noError(t, err)
}

func TestVK_UploadGroupWallPhoto(t *testing.T) {
	t.Parallel()

	needUserToken(t)
	needGroupToken(t)

	response, err := http.Get(photoURL)
	if err != nil {
		t.Fatalf("%v", err)
	}
	defer response.Body.Close()

	_, err = vkUser.UploadGroupWallPhoto(vkGroupID, response.Body)
	noError(t, err)
}

func TestVK_UploadUserPhoto(t *testing.T) {
	t.Parallel()

	needUserToken(t)

	response, err := http.Get(photoURL)
	if err != nil {
		t.Fatalf("%v", err)
	}
	defer response.Body.Close()

	_, err = vkUser.UploadUserPhoto(response.Body)
	noError(t, err)
}

func TestVK_UploadOwnerPhoto(t *testing.T) {
	t.Parallel()

	needUserToken(t)
	needGroupToken(t)

	response, err := http.Get(photoURL)
	if err != nil {
		t.Fatalf("%v", err)
	}
	defer response.Body.Close()

	_, err = vkUser.UploadOwnerPhoto(-vkGroupID, "10,10,200", response.Body)
	noError(t, err)
}

func TestVK_UploadMessagesPhoto(t *testing.T) {
	t.Parallel()

	needGroupToken(t)

	response, err := http.Get(photoURL)
	if err != nil {
		t.Fatalf("%v", err)
	}
	defer response.Body.Close()

	_, err = vkGroup.UploadMessagesPhoto(117253521, response.Body)
	noError(t, err)
}

func TestVK_UploadChatPhoto(t *testing.T) {
	t.Parallel()

	needUserToken(t)

	response, err := http.Get(photoURL)
	if err != nil {
		t.Fatalf("%v", err)
	}
	defer response.Body.Close()

	_, err = vkUser.UploadChatPhoto(needChatID(t), response.Body)
	noError(t, err)
}

func TestVK_UploadChatPhotoCrop(t *testing.T) {
	t.Parallel()

	needUserToken(t)

	response, err := http.Get(photoURL)
	if err != nil {
		t.Fatalf("%v", err)
	}
	defer response.Body.Close()

	_, err = vkUser.UploadChatPhotoCrop(needChatID(t), 0, 0, 200, response.Body)
	noError(t, err)
}

func TestVK_UploadMarketPhoto(t *testing.T) {
	t.Parallel()

	needUserToken(t)
	needGroupToken(t)

	response, err := http.Get(photoURL)
	if err != nil {
		t.Fatalf("%v", err)
	}
	defer response.Body.Close()

	_, err = vkUser.UploadMarketPhoto(vkGroupID, false, response.Body)
	noError(t, err)
}

func TestVK_UploadMarketPhotoMain(t *testing.T) {
	t.Parallel()

	needUserToken(t)
	needGroupToken(t)

	response, err := http.Get(photoURL)
	if err != nil {
		t.Fatalf("%v", err)
	}
	defer response.Body.Close()

	_, err = vkUser.UploadMarketPhoto(vkGroupID, true, response.Body)
	noError(t, err)
}

func TestVK_UploadMarketPhotoCrop(t *testing.T) {
	t.Parallel()

	needUserToken(t)
	needGroupToken(t)

	response, err := http.Get(photoURL)
	if err != nil {
		t.Fatalf("%v", err)
	}
	defer response.Body.Close()

	_, err = vkUser.UploadMarketPhotoCrop(vkGroupID, 0, 0, 400, response.Body)
	noError(t, err)
}

func TestVK_UploadMarketAlbumPhoto(t *testing.T) {
	t.Parallel()

	needUserToken(t)
	needGroupToken(t)

	response, err := http.Get(photoURL)
	if err != nil {
		t.Fatalf("%v", err)
	}
	defer response.Body.Close()

	_, err = vkUser.UploadMarketAlbumPhoto(vkGroupID, response.Body)
	noError(t, err)
}

func TestVK_UploadVideo(t *testing.T) {
	t.Parallel()

	needUserToken(t)

	respVideoGet, err := vkUser.VideoGet(api.Params{
		"videos": "-169097025_456239034",
	})
	if !noError(t, err) {
		t.FailNow()
	}

	response, err := http.Get(respVideoGet.Items[0].Files.Mp4_1080)
	if err != nil {
		t.Fatalf("%v", err)
	}
	defer response.Body.Close()

	resp, err := vkUser.UploadVideo(api.Params{}, response.Body)
	noError(t, err)

	assert.NotEmpty(t, resp.OwnerID)
	assert.NotEmpty(t, resp.VideoID)

	_, err = vkUser.VideoDelete(api.Params{
		"owner_id": resp.OwnerID,
		"video_id": resp.VideoID,
	})
	noError(t, err)

	_, err = vkUser.VideoRestore(api.Params{
		"owner_id": resp.OwnerID,
		"video_id": resp.VideoID,
	})
	noError(t, err)
}

func TestVK_UploadVideo_Error(t *testing.T) {
	t.Parallel()

	needUserToken(t)

	response, err := http.Get(photoURL)
	if err != nil {
		t.Fatalf("%v", err)
	}
	defer response.Body.Close()

	_, err = vkUser.UploadVideo(api.Params{}, response.Body)
	if _, ok := err.(*api.UploadError); err != nil && !ok { // nolint:errorlint
		t.Errorf("VK.UploadVideo() err = %v, want %v", err, -1)
	}
}

func TestVK_uploadDoc_Error(t *testing.T) {
	t.Parallel()

	needUserToken(t)

	_, err := vkUser.UploadDoc("", "", new(bytes.Buffer))
	if _, ok := err.(*api.UploadError); err != nil && !ok { // nolint:errorlint
		t.Errorf("VK.UploadDoc() err = %v, want %v", err, -1)
	}
}

func TestVK_UploadDoc(t *testing.T) {
	t.Parallel()

	needUserToken(t)

	response, err := http.Get(photoURL)
	if err != nil {
		t.Fatalf("%v", err)
	}
	defer response.Body.Close()

	_, err = vkUser.UploadDoc("test.jpeg", "test", response.Body)
	noError(t, err)
}

func TestVK_UploadGroupDoc(t *testing.T) {
	t.Parallel()

	needUserToken(t)
	needGroupToken(t)

	response, err := http.Get(photoURL)
	if err != nil {
		t.Fatalf("%v", err)
	}
	defer response.Body.Close()

	_, err = vkUser.UploadGroupDoc(vkGroupID, "test.jpeg", "test", response.Body)
	noError(t, err)
}

func TestVK_UploadWallDoc(t *testing.T) {
	t.Parallel()

	needUserToken(t)

	response, err := http.Get(photoURL)
	if err != nil {
		t.Fatalf("%v", err)
	}
	defer response.Body.Close()

	_, err = vkUser.UploadWallDoc("test.jpeg", "test", response.Body)
	noError(t, err)
}

func TestVK_UploadGroupWallDoc(t *testing.T) {
	t.Parallel()

	needUserToken(t)
	needGroupToken(t)

	response, err := http.Get(photoURL)
	if err != nil {
		t.Fatalf("%v", err)
	}
	defer response.Body.Close()

	_, err = vkUser.UploadGroupWallDoc(vkGroupID, "doc", "test.jpeg", "test", response.Body)
	noError(t, err)
}

func TestVK_UploadMessagesDoc(t *testing.T) {
	t.Parallel()

	needUserToken(t)

	response, err := http.Get(photoURL)
	if err != nil {
		t.Fatalf("%v", err)
	}
	defer response.Body.Close()

	_, err = vkUser.UploadMessagesDoc(117253521, "doc", "test.jpeg", "test", response.Body)
	noError(t, err)
}

func TestVK_UploadOwnerCoverPhoto(t *testing.T) {
	t.Parallel()

	needUserToken(t)
	needGroupToken(t)

	response, err := http.Get(photoURL)
	if err != nil {
		t.Fatalf("%v", err)
	}
	defer response.Body.Close()

	_, err = vkGroup.UploadOwnerCoverPhoto(vkGroupID, 0, 0, 795, 200, response.Body)
	noError(t, err)
}

func TestVK_UploadStoriesPhoto_Error(t *testing.T) {
	t.Parallel()

	needUserToken(t)

	_, err := vkUser.UploadStoriesPhoto(api.Params{}, new(bytes.Buffer))
	if _, ok := err.(*api.UploadError); err != nil && !ok { // nolint:errorlint
		t.Errorf("VK.UploadStoriesPhoto() err = %v, want %v", err, -1)
	}
}

func TestVK_UploadStoriesVideo_Error(t *testing.T) {
	t.Parallel()

	needUserToken(t)

	response, err := http.Get(photoURL)
	if err != nil {
		t.Fatalf("%v", err)
	}
	defer response.Body.Close()

	_, err = vkUser.UploadStoriesVideo(api.Params{}, response.Body)
	if _, ok := err.(*api.UploadError); err != nil && !ok { // nolint:errorlint
		t.Errorf("VK.UploadStoriesVideo() err = %v, want %v", err, -1)
	}
}

func TestVK_UploadPollsPhoto(t *testing.T) {
	t.Parallel()

	needUserToken(t)

	response, err := http.Get(photoURL)
	if err != nil {
		t.Fatalf("%v", err)
	}
	defer response.Body.Close()

	_, err = vkUser.UploadPollsPhoto(response.Body)
	noError(t, err)
}

func TestVK_UploadOwnerPollsPhoto(t *testing.T) {
	t.Parallel()

	needUserToken(t)
	needGroupToken(t)

	response, err := http.Get(photoURL)
	if err != nil {
		t.Fatalf("%v", err)
	}
	defer response.Body.Close()

	_, err = vkUser.UploadOwnerPollsPhoto(vkGroupID, response.Body)
	noError(t, err)
}

func TestVK_UploadPrettyCardsPhoto(t *testing.T) {
	t.Parallel()

	needUserToken(t)

	response, err := http.Get(photoURL)
	if err != nil {
		t.Fatalf("%v", err)
	}
	defer response.Body.Close()

	_, err = vkUser.UploadPrettyCardsPhoto(response.Body)
	noError(t, err)
}

func TestVK_UploadLeadFormsPhoto(t *testing.T) {
	t.Parallel()

	needUserToken(t)

	response, err := http.Get(photoURL)
	if err != nil {
		t.Fatalf("%v", err)
	}
	defer response.Body.Close()

	_, err = vkUser.UploadLeadFormsPhoto(response.Body)
	noError(t, err)
}

func TestVK_UploadAppImage(t *testing.T) {
	t.Parallel()

	t.Skip("Access denied: method available only for community applications")
	needWidgetToken(t)

	response, err := http.Get(photo480x480URL)
	if err != nil {
		t.Fatalf("%v", err)
	}
	defer response.Body.Close()

	_, err = vkWidget.UploadAppImage("160x160", response.Body)
	noError(t, err)
}

func TestVK_UploadGroupImage(t *testing.T) {
	t.Parallel()

	needWidgetToken(t)

	response, err := http.Get(photo480x480URL)
	if err != nil {
		t.Fatalf("%v", err)
	}
	defer response.Body.Close()

	_, err = vkWidget.UploadGroupImage("160x160", response.Body)
	noError(t, err)
}

func TestVK_Upload_Error(t *testing.T) {
	t.Parallel()

	vk := api.NewVK("")

	_, _ = vk.UploadPhoto(0, new(bytes.Buffer))
	_, _ = vk.UploadWallPhoto(new(bytes.Buffer))
	_, _ = vk.UploadUserPhoto(new(bytes.Buffer))
	_, _ = vk.UploadMessagesPhoto(1, new(bytes.Buffer))
	_, _ = vk.UploadChatPhoto(1, new(bytes.Buffer))
	_, _ = vk.UploadMarketPhoto(1, false, new(bytes.Buffer))
	_, _ = vk.UploadMarketAlbumPhoto(1, new(bytes.Buffer))
	_, _ = vk.UploadVideo(api.Params{}, new(bytes.Buffer))
	_, _ = vk.UploadDoc("", "", new(bytes.Buffer))
	_, _ = vk.UploadGroupDoc(1, "", "", new(bytes.Buffer))
	_, _ = vk.UploadWallDoc("", "", new(bytes.Buffer))
	_, _ = vk.UploadGroupWallDoc(1, "", "", "", new(bytes.Buffer))
	_, _ = vk.UploadMessagesDoc(1, "", "", "", new(bytes.Buffer))
	_, _ = vk.UploadOwnerCoverPhoto(1, 0, 0, 0, 0, new(bytes.Buffer))
	_, _ = vk.UploadStoriesPhoto(api.Params{}, new(bytes.Buffer))
	_, _ = vk.UploadStoriesVideo(api.Params{}, new(bytes.Buffer))
	_, _ = vk.UploadPollsPhoto(new(bytes.Buffer))
	_, _ = vk.UploadPrettyCardsPhoto(new(bytes.Buffer))
	_, _ = vk.UploadLeadFormsPhoto(new(bytes.Buffer))
	_, _ = vk.UploadAppImage("", new(bytes.Buffer))
	_, _ = vk.UploadGroupImage("", new(bytes.Buffer))
}
