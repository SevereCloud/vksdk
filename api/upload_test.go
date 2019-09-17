package api

import (
	"bytes"
	"io"
	"net/http"
	"strconv"
	"testing"

	"github.com/SevereCloud/vksdk/errors"
	"github.com/stretchr/testify/assert"
)

const photoURL = "https://sun9-17.userapi.com/c853620/v853620933/dedb8/_5CIRVR-UA8.jpg"

func TestVK_UploadFile(t *testing.T) {
	vk := Init("")

	f := func(url string, file io.Reader, fieldname, filename string, needErr bool) {
		t.Helper()
		_, err := vk.UploadFile(url, file, fieldname, filename)

		if (err != nil) && !needErr {
			t.Errorf("VK.UploadWallPhoto() err = %v, want %v", err, needErr)
		}
	}

	f("", new(bytes.Buffer), "", "", true)
}

func TestVK_UploadPhoto(t *testing.T) {
	needUserToken(t)

	album, err := vkUser.PhotosCreateAlbum(map[string]string{
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
	assert.NoError(t, err)
}

func TestVK_UploadPhotoGroup(t *testing.T) {
	needUserToken(t)
	needGroupToken(t)

	album, err := vkUser.PhotosCreateAlbum(map[string]string{
		"title":    "test",
		"group_id": strconv.Itoa(vkGroupID),
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
	assert.NoError(t, err)
}

func TestVK_UploadWallPhoto(t *testing.T) {
	needUserToken(t)

	response, err := http.Get(photoURL)
	if err != nil {
		t.Fatalf("%v", err)
	}
	defer response.Body.Close()

	_, err = vkUser.UploadWallPhoto(response.Body)
	assert.NoError(t, err)
}

func TestVK_UploadGroupWallPhoto(t *testing.T) {
	needUserToken(t)
	needGroupToken(t)

	response, err := http.Get(photoURL)
	if err != nil {
		t.Fatalf("%v", err)
	}
	defer response.Body.Close()

	_, err = vkUser.UploadGroupWallPhoto(vkGroupID, response.Body)
	assert.NoError(t, err)
}

func TestVK_UploadUserPhoto(t *testing.T) {
	needUserToken(t)

	response, err := http.Get(photoURL)
	if err != nil {
		t.Fatalf("%v", err)
	}
	defer response.Body.Close()

	_, err = vkUser.UploadUserPhoto(response.Body)
	assert.NoError(t, err)
}

func TestVK_UploadOwnerPhoto(t *testing.T) {
	needUserToken(t)
	needGroupToken(t)

	response, err := http.Get(photoURL)
	if err != nil {
		t.Fatalf("%v", err)
	}
	defer response.Body.Close()

	_, err = vkUser.UploadOwnerPhoto(-vkGroupID, "10,10,200", response.Body)
	assert.NoError(t, err)
}

func TestVK_UploadMessagesPhoto(t *testing.T) {
	needGroupToken(t)

	response, err := http.Get(photoURL)
	if err != nil {
		t.Fatalf("%v", err)
	}
	defer response.Body.Close()

	_, err = vkGroup.UploadMessagesPhoto(117253521, response.Body)
	assert.NoError(t, err)
}

func TestVK_UploadChatPhoto(t *testing.T) {
	needUserToken(t)

	response, err := http.Get(photoURL)
	if err != nil {
		t.Fatalf("%v", err)
	}
	defer response.Body.Close()

	_, err = vkUser.UploadChatPhoto(vkChatID, response.Body)
	assert.NoError(t, err)
}

func TestVK_UploadChatPhotoCrop(t *testing.T) {
	needUserToken(t)

	response, err := http.Get(photoURL)
	if err != nil {
		t.Fatalf("%v", err)
	}
	defer response.Body.Close()

	_, err = vkUser.UploadChatPhotoCrop(vkChatID, 0, 0, 200, response.Body)
	assert.NoError(t, err)
}

func TestVK_UploadMarketPhoto(t *testing.T) {
	needUserToken(t)
	needGroupToken(t)

	response, err := http.Get(photoURL)
	if err != nil {
		t.Fatalf("%v", err)
	}
	defer response.Body.Close()

	_, err = vkUser.UploadMarketPhoto(vkGroupID, false, response.Body)
	assert.NoError(t, err)
}

func TestVK_UploadMarketPhotoMain(t *testing.T) {
	needUserToken(t)
	needGroupToken(t)

	response, err := http.Get(photoURL)
	if err != nil {
		t.Fatalf("%v", err)
	}
	defer response.Body.Close()

	_, err = vkUser.UploadMarketPhoto(vkGroupID, true, response.Body)
	assert.NoError(t, err)
}

func TestVK_UploadMarketPhotoCrop(t *testing.T) {
	needUserToken(t)
	needGroupToken(t)

	response, err := http.Get(photoURL)
	if err != nil {
		t.Fatalf("%v", err)
	}
	defer response.Body.Close()

	_, err = vkUser.UploadMarketPhotoCrop(vkGroupID, 0, 0, 400, response.Body)
	assert.NoError(t, err)
}

func TestVK_UploadMarketAlbumPhoto(t *testing.T) {
	needUserToken(t)
	needGroupToken(t)

	response, err := http.Get(photoURL)
	if err != nil {
		t.Fatalf("%v", err)
	}
	defer response.Body.Close()

	_, err = vkUser.UploadMarketAlbumPhoto(vkGroupID, response.Body)
	assert.NoError(t, err)
}

func TestVK_UploadVideo_Error(t *testing.T) {
	needUserToken(t)

	response, err := http.Get(photoURL)
	if err != nil {
		t.Fatalf("%v", err)
	}
	defer response.Body.Close()

	_, err = vkUser.UploadVideo(map[string]string{}, response.Body)
	if errors.GetType(err) != errors.NoType {
		t.Errorf("VK.UploadVideo() err = %v, want %v", err, -1)
	}
}

func TestVK_uploadDoc_Error(t *testing.T) {
	needUserToken(t)

	_, err := vkUser.uploadDoc("", "", "", new(bytes.Buffer))
	if errors.GetType(err) != errors.NoType {
		t.Errorf("VK.uploadDoc() err = %v, want %v", err, -1)
	}
}

func TestVK_UploadDoc(t *testing.T) {
	needUserToken(t)

	response, err := http.Get(photoURL)
	if err != nil {
		t.Fatalf("%v", err)
	}
	defer response.Body.Close()

	_, err = vkUser.UploadDoc("test.jpeg", "test", response.Body)
	assert.NoError(t, err)
}

func TestVK_UploadGroupDoc(t *testing.T) {
	needUserToken(t)
	needGroupToken(t)

	response, err := http.Get(photoURL)
	if err != nil {
		t.Fatalf("%v", err)
	}
	defer response.Body.Close()

	_, err = vkUser.UploadGroupDoc(vkGroupID, "test.jpeg", "test", response.Body)
	assert.NoError(t, err)
}

func TestVK_UploadWallDoc(t *testing.T) {
	needUserToken(t)

	response, err := http.Get(photoURL)
	if err != nil {
		t.Fatalf("%v", err)
	}
	defer response.Body.Close()

	_, err = vkUser.UploadWallDoc("test.jpeg", "test", response.Body)
	assert.NoError(t, err)
}

func TestVK_UploadGroupWallDoc(t *testing.T) {
	needUserToken(t)
	needGroupToken(t)

	response, err := http.Get(photoURL)
	if err != nil {
		t.Fatalf("%v", err)
	}
	defer response.Body.Close()

	_, err = vkUser.UploadGroupWallDoc(vkGroupID, "test.jpeg", "test", response.Body)
	assert.NoError(t, err)
}

func TestVK_UploadMessagesDoc(t *testing.T) {
	needUserToken(t)

	response, err := http.Get(photoURL)
	if err != nil {
		t.Fatalf("%v", err)
	}
	defer response.Body.Close()

	_, err = vkUser.UploadMessagesDoc(117253521, "doc", "test.jpeg", "test", response.Body)
	assert.NoError(t, err)
}

func TestVK_UploadOwnerCoverPhoto(t *testing.T) {
	needUserToken(t)
	needGroupToken(t)

	response, err := http.Get(photoURL)
	if err != nil {
		t.Fatalf("%v", err)
	}
	defer response.Body.Close()

	_, err = vkGroup.UploadOwnerCoverPhoto(vkGroupID, 0, 0, 795, 200, response.Body)
	assert.NoError(t, err)
}

func TestVK_UploadStoriesPhoto_Error(t *testing.T) {
	needUserToken(t)

	_, err := vkUser.UploadStoriesPhoto(map[string]string{}, new(bytes.Buffer))
	if errors.GetType(err) != errors.NoType {
		t.Errorf("VK.UploadStoriesPhoto() err = %v, want %v", err, -1)
	}
}

func TestVK_UploadStoriesPhoto(t *testing.T) {
	needUserToken(t)

	response, err := http.Get(photoURL)
	if err != nil {
		t.Fatalf("%v", err)
	}
	defer response.Body.Close()

	_, err = vkUser.UploadStoriesPhoto(map[string]string{}, response.Body)
	assert.NoError(t, err)
}

func TestVK_UploadStoriesVideo_Error(t *testing.T) {
	needUserToken(t)

	response, err := http.Get(photoURL)
	if err != nil {
		t.Fatalf("%v", err)
	}
	defer response.Body.Close()

	_, err = vkUser.UploadStoriesVideo(map[string]string{}, response.Body)
	if errors.GetType(err) != errors.NoType {
		t.Errorf("VK.UploadStoriesVideo() err = %v, want %v", err, -1)
	}
}

func TestVK_UploadPollsPhoto(t *testing.T) {
	needUserToken(t)

	response, err := http.Get(photoURL)
	if err != nil {
		t.Fatalf("%v", err)
	}
	defer response.Body.Close()

	_, err = vkUser.UploadPollsPhoto(response.Body)
	assert.NoError(t, err)
}

func TestVK_UploadOwnerPollsPhoto(t *testing.T) {
	needUserToken(t)
	needGroupToken(t)

	response, err := http.Get(photoURL)
	if err != nil {
		t.Fatalf("%v", err)
	}
	defer response.Body.Close()

	_, err = vkUser.UploadOwnerPollsPhoto(vkGroupID, response.Body)
	assert.NoError(t, err)
}

func TestVK_UploadPrettyCardsPhoto(t *testing.T) {
	needUserToken(t)

	response, err := http.Get(photoURL)
	if err != nil {
		t.Fatalf("%v", err)
	}
	defer response.Body.Close()

	_, err = vkUser.UploadPrettyCardsPhoto(response.Body)
	assert.NoError(t, err)
}

func TestVK_UploadLeadFormsPhoto(t *testing.T) {
	needUserToken(t)

	response, err := http.Get(photoURL)
	if err != nil {
		t.Fatalf("%v", err)
	}
	defer response.Body.Close()

	_, err = vkUser.UploadLeadFormsPhoto(response.Body)
	assert.NoError(t, err)
}

func TestVK_UploadAppImage(t *testing.T) {
	t.Skip("Access denied: method available only for community applications")
	needServiceToken(t)

	response, err := http.Get(photoURL)
	if err != nil {
		t.Fatalf("%v", err)
	}
	defer response.Body.Close()

	_, err = vkService.UploadAppImage("160x160", response.Body)
	assert.NoError(t, err)
}

func TestVK_UploadGroupImage(t *testing.T) {
	t.Skip("Access rights required: app_widget.")
	needGroupToken(t)

	response, err := http.Get(photoURL)
	if err != nil {
		t.Fatalf("%v", err)
	}
	defer response.Body.Close()

	_, err = vkGroup.UploadGroupImage("160x160", response.Body)
	assert.NoError(t, err)
}

func TestVK_Upload_Error(t *testing.T) {
	vk := Init("")

	_, err := vk.uploadPhoto(map[string]string{}, new(bytes.Buffer))
	assert.Equal(t, errors.GetType(err), errors.Auth)

	_, err = vk.uploadWallPhoto(map[string]string{}, new(bytes.Buffer))
	assert.Equal(t, errors.GetType(err), errors.Auth)

	_, err = vk.uploadOwnerPhoto(map[string]string{}, "", new(bytes.Buffer))
	assert.Equal(t, errors.GetType(err), errors.Auth)

	_, err = vk.UploadMessagesPhoto(1, new(bytes.Buffer))
	assert.Equal(t, errors.GetType(err), errors.Auth)

	_, err = vk.uploadChatPhoto(map[string]string{}, new(bytes.Buffer))
	assert.Equal(t, errors.GetType(err), errors.Auth)

	_, err = vk.uploadMarketPhoto(map[string]string{}, new(bytes.Buffer))
	assert.Equal(t, errors.GetType(err), errors.Auth)

	_, err = vk.UploadMarketAlbumPhoto(1, new(bytes.Buffer))
	assert.Equal(t, errors.GetType(err), errors.Auth)

	_, err = vk.UploadVideo(map[string]string{}, new(bytes.Buffer))
	assert.Equal(t, errors.GetType(err), errors.Auth)

	_, err = vk.UploadDoc("", "", new(bytes.Buffer))
	assert.Equal(t, errors.GetType(err), errors.Auth)

	_, err = vk.UploadGroupDoc(1, "", "", new(bytes.Buffer))
	assert.Equal(t, errors.GetType(err), errors.Auth)

	_, err = vk.UploadWallDoc("", "", new(bytes.Buffer))
	assert.Equal(t, errors.GetType(err), errors.Auth)

	_, err = vk.UploadGroupWallDoc(1, "", "", new(bytes.Buffer))
	assert.Equal(t, errors.GetType(err), errors.Auth)

	_, err = vk.UploadMessagesDoc(1, "", "", "", new(bytes.Buffer))
	assert.Equal(t, errors.GetType(err), errors.Auth)

	_, err = vk.UploadOwnerCoverPhoto(1, 0, 0, 0, 0, new(bytes.Buffer))
	assert.Equal(t, errors.GetType(err), errors.Auth)

	_, err = vk.UploadStoriesPhoto(map[string]string{}, new(bytes.Buffer))
	assert.Equal(t, errors.GetType(err), errors.Auth)

	_, err = vk.UploadStoriesVideo(map[string]string{}, new(bytes.Buffer))
	assert.Equal(t, errors.GetType(err), errors.Auth)

	_, err = vk.uploadPollsPhoto(map[string]string{}, new(bytes.Buffer))
	assert.Equal(t, errors.GetType(err), errors.Auth)

	_, err = vk.UploadPrettyCardsPhoto(new(bytes.Buffer))
	assert.Equal(t, errors.GetType(err), errors.Auth)

	_, err = vk.UploadLeadFormsPhoto(new(bytes.Buffer))
	assert.Equal(t, errors.GetType(err), errors.Auth)

	_, err = vk.UploadAppImage("", new(bytes.Buffer))
	assert.Equal(t, errors.GetType(err), errors.Auth)

	_, err = vk.UploadGroupImage("", new(bytes.Buffer))
	assert.Equal(t, errors.GetType(err), errors.Auth)
}
