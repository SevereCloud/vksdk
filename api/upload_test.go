package api

import (
	"bytes"
	"io"
	"net/http"
	"strconv"
	"testing"

	"github.com/SevereCloud/vksdk/errors"
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
	if vkUser.AccessToken == "" {
		t.Skip("USER_TOKEN empty")
	}

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
	if err != nil {
		t.Errorf("VK.UploadPhoto() err = %v", err)
	}
}

func TestVK_UploadPhotoGroup(t *testing.T) {
	if vkUser.AccessToken == "" && vkGroupID == 0 {
		t.Skip("USER_TOKEN or GROUP_ID empty")
	}

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
	if err != nil {
		t.Errorf("VK.UploadPhotoGroup() err = %v", err)
	}
}

func TestVK_UploadWallPhoto(t *testing.T) {
	if vkUser.AccessToken == "" {
		t.Skip("USER_TOKEN empty")
	}

	response, err := http.Get(photoURL)
	if err != nil {
		t.Fatalf("%v", err)
	}
	defer response.Body.Close()

	_, err = vkUser.UploadWallPhoto(response.Body)
	if err != nil {
		t.Errorf("VK.UploadWallPhoto() err = %v", err)
	}
}

func TestVK_UploadGroupWallPhoto(t *testing.T) {
	if vkUser.AccessToken == "" && vkGroupID == 0 {
		t.Skip("USER_TOKEN or GROUP_ID empty")
	}

	response, err := http.Get(photoURL)
	if err != nil {
		t.Fatalf("%v", err)
	}
	defer response.Body.Close()

	_, err = vkUser.UploadGroupWallPhoto(vkGroupID, response.Body)
	if err != nil {
		t.Errorf("VK.UploadGroupWallPhoto() err = %v", err)
	}
}

func TestVK_UploadUserPhoto(t *testing.T) {
	if vkUser.AccessToken == "" {
		t.Skip("USER_TOKEN empty")
	}

	response, err := http.Get(photoURL)
	if err != nil {
		t.Fatalf("%v", err)
	}
	defer response.Body.Close()

	_, err = vkUser.UploadUserPhoto(response.Body)
	if err != nil {
		t.Errorf("VK.UploadUserPhoto() err = %v", err)
	}
}

func TestVK_UploadOwnerPhoto(t *testing.T) {
	if vkUser.AccessToken == "" && vkGroupID == 0 {
		t.Skip("USER_TOKEN or GROUP_ID empty")
	}

	response, err := http.Get(photoURL)
	if err != nil {
		t.Fatalf("%v", err)
	}
	defer response.Body.Close()

	_, err = vkUser.UploadOwnerPhoto(-vkGroupID, "10,10,200", response.Body)
	if err != nil {
		t.Errorf("VK.UploadOwnerPhoto() err = %v", err)
	}
}

func TestVK_UploadMessagesPhoto(t *testing.T) {
	if vkGroup.AccessToken == "" {
		t.Skip("GROUP_TOKEN empty")
	}

	response, err := http.Get(photoURL)
	if err != nil {
		t.Fatalf("%v", err)
	}
	defer response.Body.Close()

	_, err = vkGroup.UploadMessagesPhoto(117253521, response.Body)
	if err != nil {
		t.Errorf("VK.UploadMessagesPhoto() err = %v", err)
	}
}

func TestVK_UploadChatPhoto(t *testing.T) {
	if vkUser.AccessToken == "" {
		t.Skip("USER_TOKEN empty")
	}

	response, err := http.Get(photoURL)
	if err != nil {
		t.Fatalf("%v", err)
	}
	defer response.Body.Close()

	_, err = vkUser.UploadChatPhoto(vkChatID, response.Body)
	if err != nil {
		t.Errorf("VK.UploadChatPhoto() err = %v", err)
	}
}

func TestVK_UploadChatPhotoCrop(t *testing.T) {
	if vkUser.AccessToken == "" {
		t.Skip("USER_TOKEN empty")
	}

	response, err := http.Get(photoURL)
	if err != nil {
		t.Fatalf("%v", err)
	}
	defer response.Body.Close()

	_, err = vkUser.UploadChatPhotoCrop(vkChatID, 0, 0, 200, response.Body)
	if err != nil {
		t.Errorf("VK.UploadChatPhotoCrop() err = %v", err)
	}
}

func TestVK_UploadMarketPhoto(t *testing.T) {
	if vkUser.AccessToken == "" && vkGroupID == 0 {
		t.Skip("USER_TOKEN or GROUP_ID empty")
	}

	response, err := http.Get(photoURL)
	if err != nil {
		t.Fatalf("%v", err)
	}
	defer response.Body.Close()

	_, err = vkUser.UploadMarketPhoto(vkGroupID, false, response.Body)
	if err != nil {
		t.Errorf("VK.UploadMarketPhoto() err = %v", err)
	}
}

func TestVK_UploadMarketPhotoMain(t *testing.T) {
	if vkUser.AccessToken == "" && vkGroupID == 0 {
		t.Skip("USER_TOKEN or GROUP_ID empty")
	}

	response, err := http.Get(photoURL)
	if err != nil {
		t.Fatalf("%v", err)
	}
	defer response.Body.Close()

	_, err = vkUser.UploadMarketPhoto(vkGroupID, true, response.Body)
	if err != nil {
		t.Errorf("VK.UploadMarketPhoto() err = %v", err)
	}
}

func TestVK_UploadMarketPhotoCrop(t *testing.T) {
	if vkUser.AccessToken == "" && vkGroupID == 0 {
		t.Skip("USER_TOKEN or GROUP_ID empty")
	}

	response, err := http.Get(photoURL)
	if err != nil {
		t.Fatalf("%v", err)
	}
	defer response.Body.Close()

	_, err = vkUser.UploadMarketPhotoCrop(vkGroupID, 0, 0, 400, response.Body)
	if err != nil {
		t.Errorf("VK.UploadMarketPhotoCrop() err = %v", err)
	}
}

func TestVK_UploadMarketAlbumPhoto(t *testing.T) {
	if vkUser.AccessToken == "" && vkGroupID == 0 {
		t.Skip("USER_TOKEN or GROUP_ID empty")
	}

	response, err := http.Get(photoURL)
	if err != nil {
		t.Fatalf("%v", err)
	}
	defer response.Body.Close()

	_, err = vkUser.UploadMarketAlbumPhoto(vkGroupID, response.Body)
	if err != nil {
		t.Errorf("VK.UploadMarketAlbumPhoto() err = %v", err)
	}
}

func TestVK_UploadVideo_Error(t *testing.T) {
	if vkUser.AccessToken == "" {
		t.Skip("USER_TOKEN empty")
	}

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
	if vkUser.AccessToken == "" {
		t.Skip("USER_TOKEN empty")
	}

	_, err := vkUser.uploadDoc("", "", "", new(bytes.Buffer))
	if errors.GetType(err) != errors.NoType {
		t.Errorf("VK.uploadDoc() err = %v, want %v", err, -1)
	}
}

func TestVK_UploadDoc(t *testing.T) {
	if vkUser.AccessToken == "" {
		t.Skip("USER_TOKEN empty")
	}

	response, err := http.Get(photoURL)
	if err != nil {
		t.Fatalf("%v", err)
	}
	defer response.Body.Close()

	_, err = vkUser.UploadDoc("test.jpeg", "test", response.Body)
	if err != nil {
		t.Errorf("VK.UploadDoc() err = %v", err)
	}
}

func TestVK_UploadGroupDoc(t *testing.T) {
	if vkUser.AccessToken == "" && vkGroupID == 0 {
		t.Skip("USER_TOKEN or GROUP_ID empty")
	}

	response, err := http.Get(photoURL)
	if err != nil {
		t.Fatalf("%v", err)
	}
	defer response.Body.Close()

	_, err = vkUser.UploadGroupDoc(vkGroupID, "test.jpeg", "test", response.Body)
	if err != nil {
		t.Errorf("VK.UploadGroupDoc() err = %v", err)
	}
}

func TestVK_UploadWallDoc(t *testing.T) {
	if vkUser.AccessToken == "" {
		t.Skip("USER_TOKEN empty")
	}

	response, err := http.Get(photoURL)
	if err != nil {
		t.Fatalf("%v", err)
	}
	defer response.Body.Close()

	_, err = vkUser.UploadWallDoc("test.jpeg", "test", response.Body)
	if err != nil {
		t.Errorf("VK.UploadWallDoc() err = %v", err)
	}
}

func TestVK_UploadGroupWallDoc(t *testing.T) {
	if vkUser.AccessToken == "" && vkGroupID == 0 {
		t.Skip("USER_TOKEN or GROUP_ID empty")
	}

	response, err := http.Get(photoURL)
	if err != nil {
		t.Fatalf("%v", err)
	}
	defer response.Body.Close()

	_, err = vkUser.UploadGroupWallDoc(vkGroupID, "test.jpeg", "test", response.Body)
	if err != nil {
		t.Errorf("VK.UploadGroupWallDoc() err = %v", err)
	}
}

func TestVK_UploadMessagesDoc(t *testing.T) {
	if vkUser.AccessToken == "" {
		t.Skip("USER_TOKEN empty")
	}

	response, err := http.Get(photoURL)
	if err != nil {
		t.Fatalf("%v", err)
	}
	defer response.Body.Close()

	_, err = vkUser.UploadMessagesDoc(117253521, "doc", "test.jpeg", "test", response.Body)
	if err != nil {
		t.Errorf("VK.UploadMessagesDoc() err = %v", err)
	}
}

func TestVK_UploadOwnerCoverPhoto(t *testing.T) {
	if vkGroup.AccessToken == "" && vkGroupID == 0 {
		t.Skip("GROUP_TOKEN or GROUP_ID empty")
	}

	response, err := http.Get(photoURL)
	if err != nil {
		t.Fatalf("%v", err)
	}
	defer response.Body.Close()

	_, err = vkGroup.UploadOwnerCoverPhoto(vkGroupID, 0, 0, 795, 200, response.Body)
	if err != nil {
		t.Errorf("VK.UploadOwnerCoverPhoto() err = %v", err)
	}
}

func TestVK_UploadStoriesPhoto_Error(t *testing.T) {
	if vkUser.AccessToken == "" {
		t.Skip("USER_TOKEN empty")
	}

	_, err := vkUser.UploadStoriesPhoto(map[string]string{}, new(bytes.Buffer))
	if errors.GetType(err) != errors.NoType {
		t.Errorf("VK.UploadStoriesPhoto() err = %v, want %v", err, -1)
	}
}

func TestVK_UploadStoriesPhoto(t *testing.T) {
	if vkUser.AccessToken == "" {
		t.Skip("USER_TOKEN empty")
	}

	response, err := http.Get(photoURL)
	if err != nil {
		t.Fatalf("%v", err)
	}
	defer response.Body.Close()

	_, err = vkUser.UploadStoriesPhoto(map[string]string{}, response.Body)
	if err != nil {
		t.Errorf("VK.UploadStoriesPhoto() err = %v", err)
	}
}

func TestVK_UploadStoriesVideo_Error(t *testing.T) {
	if vkUser.AccessToken == "" {
		t.Skip("USER_TOKEN empty")
	}

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

func TestVK_Upload_Error(t *testing.T) {
	vk := Init("")

	_, err := vk.uploadPhoto(map[string]string{}, new(bytes.Buffer))
	if errors.GetType(err) != errors.Auth {
		t.Errorf("VK.uploadPhoto() err = %v, want %v", err, 5)
	}
	_, err = vk.uploadWallPhoto(map[string]string{}, new(bytes.Buffer))
	if errors.GetType(err) != errors.Auth {
		t.Errorf("VK.uploadWallPhoto() err = %v, want %v", err, 5)
	}

	_, err = vk.uploadOwnerPhoto(map[string]string{}, "", new(bytes.Buffer))
	if errors.GetType(err) != errors.Auth {
		t.Errorf("VK.uploadOwnerPhoto() err = %v, want %v", err, 5)
	}

	_, err = vk.UploadMessagesPhoto(1, new(bytes.Buffer))
	if errors.GetType(err) != errors.Auth {
		t.Errorf("VK.uploadOwnerPhoto() err = %v, want %v", err, 5)
	}

	_, err = vk.uploadChatPhoto(map[string]string{}, new(bytes.Buffer))
	if errors.GetType(err) != errors.Auth {
		t.Errorf("VK.uploadChatPhoto() err = %v, want %v", err, 5)
	}

	_, err = vk.uploadMarketPhoto(map[string]string{}, new(bytes.Buffer))
	if errors.GetType(err) != errors.Auth {
		t.Errorf("VK.uploadMarketPhoto() err = %v, want %v", err, 5)
	}

	_, err = vk.UploadMarketAlbumPhoto(1, new(bytes.Buffer))
	if errors.GetType(err) != errors.Auth {
		t.Errorf("VK.UploadMarketAlbumPhoto() err = %v, want %v", err, 5)
	}

	_, err = vk.UploadVideo(map[string]string{}, new(bytes.Buffer))
	if errors.GetType(err) != errors.Auth {
		t.Errorf("VK.UploadVideo() err = %v, want %v", err, 5)
	}

	_, err = vk.UploadDoc("", "", new(bytes.Buffer))
	if errors.GetType(err) != errors.Auth {
		t.Errorf("VK.UploadDoc() err = %v, want %v", err, 5)
	}

	_, err = vk.UploadGroupDoc(1, "", "", new(bytes.Buffer))
	if errors.GetType(err) != errors.Auth {
		t.Errorf("VK.UploadGroupDoc() err = %v, want %v", err, 5)
	}

	_, err = vk.UploadWallDoc("", "", new(bytes.Buffer))
	if errors.GetType(err) != errors.Auth {
		t.Errorf("VK.UploadWallDoc() err = %v, want %v", err, 5)
	}

	_, err = vk.UploadGroupWallDoc(1, "", "", new(bytes.Buffer))
	if errors.GetType(err) != errors.Auth {
		t.Errorf("VK.UploadGroupWallDoc() err = %v, want %v", err, 5)
	}

	_, err = vk.UploadMessagesDoc(1, "", "", "", new(bytes.Buffer))
	if errors.GetType(err) != errors.Auth {
		t.Errorf("VK.UploadMessagesDoc() err = %v, want %v", err, 5)
	}

	_, err = vk.UploadOwnerCoverPhoto(1, 0, 0, 0, 0, new(bytes.Buffer))
	if errors.GetType(err) != errors.Auth {
		t.Errorf("VK.UploadOwnerCoverPhoto() err = %v, want %v", err, 5)
	}

	_, err = vk.UploadStoriesPhoto(map[string]string{}, new(bytes.Buffer))
	if errors.GetType(err) != errors.Auth {
		t.Errorf("VK.UploadStoriesPhoto() err = %v, want %v", err, 5)
	}

	_, err = vk.UploadStoriesVideo(map[string]string{}, new(bytes.Buffer))
	if errors.GetType(err) != errors.Auth {
		t.Errorf("VK.UploadStoriesVideo() err = %v, want %v", err, 5)
	}
}
