package api

import (
	"bytes"
	"io"
	"net/http"
	"strconv"
	"testing"
)

const photoURL = "https://sun9-45.userapi.com/c638629/v638629852/2afba/o-dvykjSIB4.jpg"

func Test_UploadFile(t *testing.T) {
	f := func(url string, file io.Reader, fieldname, filename string, needErr bool) {
		t.Helper()
		_, err := UploadFile(url, file, fieldname, filename)

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
	if vkUser.AccessToken == "" {
		t.Skip("USER_TOKEN empty")
	}

	response, err := http.Get(photoURL)
	if err != nil {
		t.Fatalf("%v", err)
	}
	defer response.Body.Close()

	_, gotVkErr := vkUser.UploadMessagesPhoto(117253521, response.Body)
	if gotVkErr.Code != 0 {
		t.Errorf("VK.UploadMessagesPhoto() gotVkErr = %v, want %v", gotVkErr, 0)
	}
}
