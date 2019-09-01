package api // import "github.com/SevereCloud/vksdk/5.92/api"

import (
	"bytes"
	"encoding/json"
	"io"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"strconv"

	"github.com/SevereCloud/vksdk/5.92/object"
)

func UploadFile(url string, file io.Reader, fieldname, filename string) (bodyContent []byte, err error) {
	body := new(bytes.Buffer)
	writer := multipart.NewWriter(body)
	part, err := writer.CreateFormFile(fieldname, filename)
	if err != nil {
		return
	}

	_, err = io.Copy(part, file)
	if err != nil {
		return
	}

	contentType := writer.FormDataContentType()
	writer.Close()

	resp, err := http.Post(url, contentType, body) // nolint: gosec
	if err != nil {
		return
	}
	defer resp.Body.Close()

	bodyContent, err = ioutil.ReadAll(resp.Body)
	return
}

func (vk *VK) uploadPhoto(params map[string]string, file io.Reader) (response PhotosSaveResponse, vkErr Error) {
	uploadServer, vkErr := vk.PhotosGetUploadServer(params)
	if vkErr.Code != 0 {
		return
	}

	bodyContent, err := UploadFile(uploadServer.UploadURL, file, "file1", "file1.jpeg")
	if err != nil {
		vkErr = NewError(-1, err.Error(), "", map[string]string{})
		return
	}

	var handler object.PhotosPhotoUploadResponse
	err = json.Unmarshal(bodyContent, &handler)
	if err != nil {
		vkErr = NewError(-1, err.Error(), "", map[string]string{})
		return
	}

	response, vkErr = vk.PhotosSave(map[string]string{
		"server":      strconv.Itoa(handler.Server),
		"photos_list": handler.PhotosList,
		"aid":         strconv.Itoa(handler.AID),
		"hash":        handler.Hash,
		"album_id":    params["album_id"],
		"group_id":    params["group_id"],
	})
	return
}

// UploadPhoto upload photo to album
func (vk *VK) UploadPhoto(albumID int, file io.Reader) (response PhotosSaveResponse, vkErr Error) {
	response, vkErr = vk.uploadPhoto(map[string]string{
		"album_id": strconv.Itoa(albumID),
	}, file)
	return
}

// UploadPhotoGroup upload photo to group album
func (vk *VK) UploadPhotoGroup(groupID, albumID int, file io.Reader) (response PhotosSaveResponse, vkErr Error) {
	response, vkErr = vk.uploadPhoto(map[string]string{
		"album_id": strconv.Itoa(albumID),
		"group_id": strconv.Itoa(groupID),
	}, file)
	return
}

// uploadWallPhoto upload photo for wall
func (vk *VK) uploadWallPhoto(params map[string]string, file io.Reader) (response PhotosSaveWallPhotoResponse, vkErr Error) {
	uploadServer, vkErr := vk.PhotosGetWallUploadServer(params)
	if vkErr.Code != 0 {
		return
	}

	bodyContent, err := UploadFile(uploadServer.UploadURL, file, "photo", "photo.jpeg")
	if err != nil {
		vkErr = NewError(-1, err.Error(), "", map[string]string{})
		return
	}

	var handler object.PhotosWallUploadResponse
	err = json.Unmarshal(bodyContent, &handler)
	if err != nil {
		vkErr = NewError(-1, err.Error(), "", map[string]string{})
		return
	}

	response, vkErr = vk.PhotosSaveWallPhoto(map[string]string{
		"server":   strconv.Itoa(handler.Server),
		"photo":    handler.Photo,
		"hash":     handler.Hash,
		"group_id": params["group_id"],
	})
	return
}

// UploadWallPhoto upload photo for wall
func (vk *VK) UploadWallPhoto(file io.Reader) (response PhotosSaveWallPhotoResponse, vkErr Error) {
	response, vkErr = vk.uploadWallPhoto(map[string]string{}, file)
	return
}

// UploadGroupWallPhoto upload photo for group wall
func (vk *VK) UploadGroupWallPhoto(groupID int, file io.Reader) (response PhotosSaveWallPhotoResponse, vkErr Error) {
	response, vkErr = vk.uploadWallPhoto(map[string]string{
		"group_id": strconv.Itoa(groupID),
	}, file)
	return
}
