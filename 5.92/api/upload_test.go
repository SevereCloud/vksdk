package api

import (
	"bytes"
	"io"
	"net/http"
	"strconv"
	"testing"
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

	album, gotVkErr := vkUser.PhotosCreateAlbum(map[string]string{
		"title": "test",
	})
	if gotVkErr.Code != 0 {
		t.Fatalf("VK.PhotosCreateAlbum() gotVkErr = %v, want %v", gotVkErr, 0)
	}

	response, err := http.Get(photoURL)
	if err != nil {
		t.Fatalf("%v", err)
	}
	defer response.Body.Close()

	_, gotVkErr = vkUser.UploadPhoto(album.ID, response.Body)
	if gotVkErr.Code != 0 {
		t.Errorf("VK.UploadPhoto() gotVkErr = %v, want %v", gotVkErr, 0)
	}
}

func TestVK_UploadPhotoGroup(t *testing.T) {
	if vkUser.AccessToken == "" && vkGroupID == 0 {
		t.Skip("USER_TOKEN or GROUP_ID empty")
	}

	album, gotVkErr := vkUser.PhotosCreateAlbum(map[string]string{
		"title":    "test",
		"group_id": strconv.Itoa(vkGroupID),
	})
	if gotVkErr.Code != 0 {
		t.Fatalf("VK.PhotosCreateAlbum() gotVkErr = %v, want %v", gotVkErr, 0)
	}

	response, err := http.Get(photoURL)
	if err != nil {
		t.Fatalf("%v", err)
	}
	defer response.Body.Close()

	_, gotVkErr = vkUser.UploadPhotoGroup(vkGroupID, album.ID, response.Body)
	if gotVkErr.Code != 0 {
		t.Errorf("VK.UploadPhotoGroup() gotVkErr = %v, want %v", gotVkErr, 0)
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

	_, gotVkErr := vkUser.UploadWallPhoto(response.Body)
	if gotVkErr.Code != 0 {
		t.Errorf("VK.UploadWallPhoto() gotVkErr = %v, want %v", gotVkErr, 0)
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

	_, gotVkErr := vkUser.UploadGroupWallPhoto(vkGroupID, response.Body)
	if gotVkErr.Code != 0 {
		t.Errorf("VK.UploadGroupWallPhoto() gotVkErr = %v, want %v", gotVkErr, 0)
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

	_, gotVkErr := vkUser.UploadUserPhoto(response.Body)
	if gotVkErr.Code != 0 {
		t.Errorf("VK.UploadUserPhoto() gotVkErr = %v, want %v", gotVkErr, 0)
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

	_, gotVkErr := vkUser.UploadOwnerPhoto(-vkGroupID, "10,10,200", response.Body)
	if gotVkErr.Code != 0 {
		t.Errorf("VK.UploadOwnerPhoto() gotVkErr = %v, want %v", gotVkErr, 0)
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

	_, gotVkErr := vkGroup.UploadMessagesPhoto(117253521, response.Body)
	if gotVkErr.Code != 0 {
		t.Errorf("VK.UploadMessagesPhoto() gotVkErr = %v, want %v", gotVkErr, 0)
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

	_, gotVkErr := vkUser.UploadChatPhoto(vkChatID, response.Body)
	if gotVkErr.Code != 0 {
		t.Errorf("VK.UploadChatPhoto() gotVkErr = %v, want %v", gotVkErr, 0)
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

	_, gotVkErr := vkUser.UploadChatPhotoCrop(vkChatID, 0, 0, 200, response.Body)
	if gotVkErr.Code != 0 {
		t.Errorf("VK.UploadChatPhotoCrop() gotVkErr = %v, want %v", gotVkErr, 0)
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

	_, gotVkErr := vkUser.UploadMarketPhoto(vkGroupID, false, response.Body)
	if gotVkErr.Code != 0 {
		t.Errorf("VK.UploadMarketPhoto() gotVkErr = %v, want %v", gotVkErr, 0)
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

	_, gotVkErr := vkUser.UploadMarketPhoto(vkGroupID, true, response.Body)
	if gotVkErr.Code != 0 {
		t.Errorf("VK.UploadMarketPhoto() gotVkErr = %v, want %v", gotVkErr, 0)
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

	_, gotVkErr := vkUser.UploadMarketPhotoCrop(vkGroupID, 0, 0, 400, response.Body)
	if gotVkErr.Code != 0 {
		t.Errorf("VK.UploadMarketPhotoCrop() gotVkErr = %v, want %v", gotVkErr, 0)
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

	_, gotVkErr := vkUser.UploadMarketAlbumPhoto(vkGroupID, response.Body)
	if gotVkErr.Code != 0 {
		t.Errorf("VK.UploadMarketAlbumPhoto() gotVkErr = %v, want %v", gotVkErr, 0)
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

	_, gotVkErr := vkUser.UploadVideo(map[string]string{}, response.Body)
	if gotVkErr.Code != -1 {
		t.Errorf("VK.UploadVideo() gotVkErr = %v, want %v", gotVkErr, -1)
	}
}

func TestVK_uploadDoc_Error(t *testing.T) {
	if vkUser.AccessToken == "" {
		t.Skip("USER_TOKEN empty")
	}

	_, gotVkErr := vkUser.uploadDoc("", "", "", new(bytes.Buffer))
	if gotVkErr.Code != -1 {
		t.Errorf("VK.uploadDoc() gotVkErr = %v, want %v", gotVkErr, -1)
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

	_, gotVkErr := vkUser.UploadDoc("test.jpeg", "test", response.Body)
	if gotVkErr.Code != 0 {
		t.Errorf("VK.UploadDoc() gotVkErr = %v, want %v", gotVkErr, 0)
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

	_, gotVkErr := vkUser.UploadGroupDoc(vkGroupID, "test.jpeg", "test", response.Body)
	if gotVkErr.Code != 0 {
		t.Errorf("VK.UploadGroupDoc() gotVkErr = %v, want %v", gotVkErr, 0)
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

	_, gotVkErr := vkUser.UploadWallDoc("test.jpeg", "test", response.Body)
	if gotVkErr.Code != 0 {
		t.Errorf("VK.UploadWallDoc() gotVkErr = %v, want %v", gotVkErr, 0)
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

	_, gotVkErr := vkUser.UploadGroupWallDoc(vkGroupID, "test.jpeg", "test", response.Body)
	if gotVkErr.Code != 0 {
		t.Errorf("VK.UploadGroupWallDoc() gotVkErr = %v, want %v", gotVkErr, 0)
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

	_, gotVkErr := vkUser.UploadMessagesDoc(117253521, "doc", "test.jpeg", "test", response.Body)
	if gotVkErr.Code != 0 {
		t.Errorf("VK.UploadMessagesDoc() gotVkErr = %v, want %v", gotVkErr, 0)
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

	_, gotVkErr := vkGroup.UploadOwnerCoverPhoto(vkGroupID, 0, 0, 795, 200, response.Body)
	if gotVkErr.Code != 0 {
		t.Errorf("VK.UploadOwnerCoverPhoto() gotVkErr = %v, want %v", gotVkErr, 0)
	}
}

func TestVK_UploadStoriesPhoto_Error(t *testing.T) {
	if vkUser.AccessToken == "" {
		t.Skip("USER_TOKEN empty")
	}

	_, gotVkErr := vkUser.UploadStoriesPhoto(map[string]string{}, new(bytes.Buffer))
	if gotVkErr.Code != -1 {
		t.Errorf("VK.UploadStoriesPhoto() gotVkErr = %v, want %v", gotVkErr, -1)
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

	_, gotVkErr := vkUser.UploadStoriesPhoto(map[string]string{}, response.Body)
	if gotVkErr.Code != 0 {
		t.Errorf("VK.UploadStoriesPhoto() gotVkErr = %v, want %v", gotVkErr, 0)
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

	_, gotVkErr := vkUser.UploadStoriesVideo(map[string]string{}, response.Body)
	if gotVkErr.Code != -1 {
		t.Errorf("VK.UploadStoriesVideo() gotVkErr = %v, want %v", gotVkErr, -1)
	}
}

func TestVK_Upload_Error(t *testing.T) {
	vk := Init("")

	_, gotVkErr := vk.uploadPhoto(map[string]string{}, new(bytes.Buffer))
	if gotVkErr.Code != 5 {
		t.Errorf("VK.uploadPhoto() gotVkErr = %v, want %v", gotVkErr, 5)
	}
	_, gotVkErr = vk.uploadWallPhoto(map[string]string{}, new(bytes.Buffer))
	if gotVkErr.Code != 5 {
		t.Errorf("VK.uploadWallPhoto() gotVkErr = %v, want %v", gotVkErr, 5)
	}

	_, gotVkErr = vk.uploadOwnerPhoto(map[string]string{}, "", new(bytes.Buffer))
	if gotVkErr.Code != 5 {
		t.Errorf("VK.uploadOwnerPhoto() gotVkErr = %v, want %v", gotVkErr, 5)
	}

	_, gotVkErr = vk.UploadMessagesPhoto(1, new(bytes.Buffer))
	if gotVkErr.Code != 5 {
		t.Errorf("VK.uploadOwnerPhoto() gotVkErr = %v, want %v", gotVkErr, 5)
	}

	_, gotVkErr = vk.uploadChatPhoto(map[string]string{}, new(bytes.Buffer))
	if gotVkErr.Code != 5 {
		t.Errorf("VK.uploadChatPhoto() gotVkErr = %v, want %v", gotVkErr, 5)
	}

	_, gotVkErr = vk.uploadMarketPhoto(map[string]string{}, new(bytes.Buffer))
	if gotVkErr.Code != 5 {
		t.Errorf("VK.uploadMarketPhoto() gotVkErr = %v, want %v", gotVkErr, 5)
	}

	_, gotVkErr = vk.UploadMarketAlbumPhoto(1, new(bytes.Buffer))
	if gotVkErr.Code != 5 {
		t.Errorf("VK.UploadMarketAlbumPhoto() gotVkErr = %v, want %v", gotVkErr, 5)
	}

	_, gotVkErr = vk.UploadVideo(map[string]string{}, new(bytes.Buffer))
	if gotVkErr.Code != 5 {
		t.Errorf("VK.UploadVideo() gotVkErr = %v, want %v", gotVkErr, 5)
	}

	_, gotVkErr = vk.uploadDoc("", "", "", new(bytes.Buffer))
	if gotVkErr.Code != 5 {
		t.Errorf("VK.uploadDoc() gotVkErr = %v, want %v", gotVkErr, 5)
	}

	_, gotVkErr = vk.UploadDoc("", "", new(bytes.Buffer))
	if gotVkErr.Code != 5 {
		t.Errorf("VK.UploadDoc() gotVkErr = %v, want %v", gotVkErr, 5)
	}

	_, gotVkErr = vk.UploadGroupDoc(1, "", "", new(bytes.Buffer))
	if gotVkErr.Code != 5 {
		t.Errorf("VK.UploadGroupDoc() gotVkErr = %v, want %v", gotVkErr, 5)
	}

	_, gotVkErr = vk.UploadWallDoc("", "", new(bytes.Buffer))
	if gotVkErr.Code != 5 {
		t.Errorf("VK.UploadWallDoc() gotVkErr = %v, want %v", gotVkErr, 5)
	}

	_, gotVkErr = vk.UploadGroupWallDoc(1, "", "", new(bytes.Buffer))
	if gotVkErr.Code != 5 {
		t.Errorf("VK.UploadGroupWallDoc() gotVkErr = %v, want %v", gotVkErr, 5)
	}

	_, gotVkErr = vk.UploadMessagesDoc(1, "", "", "", new(bytes.Buffer))
	if gotVkErr.Code != 5 {
		t.Errorf("VK.UploadMessagesDoc() gotVkErr = %v, want %v", gotVkErr, 5)
	}

	_, gotVkErr = vk.UploadOwnerCoverPhoto(1, 0, 0, 0, 0, new(bytes.Buffer))
	if gotVkErr.Code != 5 {
		t.Errorf("VK.UploadOwnerCoverPhoto() gotVkErr = %v, want %v", gotVkErr, 5)
	}

	_, gotVkErr = vk.UploadStoriesPhoto(map[string]string{}, new(bytes.Buffer))
	if gotVkErr.Code != 5 {
		t.Errorf("VK.UploadStoriesPhoto() gotVkErr = %v, want %v", gotVkErr, 5)
	}

	_, gotVkErr = vk.UploadStoriesVideo(map[string]string{}, new(bytes.Buffer))
	if gotVkErr.Code != 5 {
		t.Errorf("VK.UploadStoriesVideo() gotVkErr = %v, want %v", gotVkErr, 5)
	}
}
