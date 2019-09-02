package api // import "github.com/SevereCloud/vksdk/5.92/api"

import (
	"bytes"
	"encoding/json"
	"io"
	"io/ioutil"
	"mime/multipart"
	"strconv"

	"github.com/SevereCloud/vksdk/5.92/object"
)

type uploadError struct {
	Error     string `json:"error"`
	ErrorCode int    `json:"error_code"`
}

// UploadFile uploading file
func (vk *VK) UploadFile(url string, file io.Reader, fieldname, filename string) (bodyContent []byte, err error) {
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

	resp, err := vk.Client.Post(url, contentType, body) // nolint: gosec
	if err != nil {
		return
	}
	defer resp.Body.Close()

	bodyContent, err = ioutil.ReadAll(resp.Body)
	return
}

// uploadPhoto uploading Photos into Album
//
// Supported formats: JPG, PNG, GIF.
//
// Limits: width+height not more than 14000 px, file size up to 50 Mb,
// aspect ratio of at least 1:20
func (vk *VK) uploadPhoto(params map[string]string, file io.Reader) (response PhotosSaveResponse, vkErr Error) {
	uploadServer, vkErr := vk.PhotosGetUploadServer(params)
	if vkErr.Code != 0 {
		return
	}

	bodyContent, err := vk.UploadFile(uploadServer.UploadURL, file, "file1", "file1.jpeg")
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

// UploadPhoto uploading Photos into User Album
//
// Supported formats: JPG, PNG, GIF.
//
// Limits: width+height not more than 14000 px, file size up to 50 Mb,
// aspect ratio of at least 1:20
func (vk *VK) UploadPhoto(albumID int, file io.Reader) (response PhotosSaveResponse, vkErr Error) {
	response, vkErr = vk.uploadPhoto(map[string]string{
		"album_id": strconv.Itoa(albumID),
	}, file)
	return
}

// UploadPhotoGroup uploading Photos into Group Album
//
// Supported formats: JPG, PNG, GIF.
//
// Limits: width+height not more than 14000 px, file size up to 50 Mb,
// aspect ratio of at least 1:20
func (vk *VK) UploadPhotoGroup(groupID, albumID int, file io.Reader) (response PhotosSaveResponse, vkErr Error) {
	response, vkErr = vk.uploadPhoto(map[string]string{
		"album_id": strconv.Itoa(albumID),
		"group_id": strconv.Itoa(groupID),
	}, file)
	return
}

// uploadWallPhoto uploading Photos on Wall
//
// Supported formats: JPG, PNG, GIF.
//
// Limits: width+height not more than 14000 px, file size up to 50 Mb,
// aspect ratio of at least 1:20
func (vk *VK) uploadWallPhoto(params map[string]string, file io.Reader) (response PhotosSaveWallPhotoResponse, vkErr Error) {
	uploadServer, vkErr := vk.PhotosGetWallUploadServer(params)
	if vkErr.Code != 0 {
		return
	}

	bodyContent, err := vk.UploadFile(uploadServer.UploadURL, file, "photo", "photo.jpeg")
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

// UploadWallPhoto uploading Photos on User Wall
//
// Supported formats: JPG, PNG, GIF.
//
// Limits: width+height not more than 14000 px, file size up to 50 Mb,
// aspect ratio of at least 1:20
func (vk *VK) UploadWallPhoto(file io.Reader) (response PhotosSaveWallPhotoResponse, vkErr Error) {
	response, vkErr = vk.uploadWallPhoto(map[string]string{}, file)
	return
}

// UploadGroupWallPhoto uploading Photos on Group Wall
//
// Supported formats: JPG, PNG, GIF.
//
// Limits: width+height not more than 14000 px, file size up to 50 Mb,
// aspect ratio of at least 1:20
func (vk *VK) UploadGroupWallPhoto(groupID int, file io.Reader) (response PhotosSaveWallPhotoResponse, vkErr Error) {
	response, vkErr = vk.uploadWallPhoto(map[string]string{
		"group_id": strconv.Itoa(groupID),
	}, file)
	return
}

// uploadOwnerPhoto uploading Photos into User Profile or Community
// To upload a photo to a community send its negative id in the owner_id parameter.
//
// Following parameters can be sent in addition:
// squareCrop in x,y,w (no quotes) format where x and y are the coordinates of
// the preview upper-right corner and w is square side length.
// That will create a square preview for a photo.
//
// Supported formats: JPG, PNG, GIF.
//
// Limits: size not less than 200x200px, aspect ratio from 0.25 to 3,
// width+height not more than 14000 px, file size up to 50 Mb.
func (vk *VK) uploadOwnerPhoto(params map[string]string, squareCrop string, file io.Reader) (response PhotosSaveOwnerPhotoResponse, vkErr Error) {
	uploadServer, vkErr := vk.PhotosGetOwnerPhotoUploadServer(params)
	if vkErr.Code != 0 {
		return
	}

	body := new(bytes.Buffer)
	writer := multipart.NewWriter(body)
	part, err := writer.CreateFormFile("photo", "photo.jpeg")
	if err != nil {
		vkErr = NewError(-1, err.Error(), "", map[string]string{})
		return
	}

	_, err = io.Copy(part, file)
	if err != nil {
		vkErr = NewError(-1, err.Error(), "", map[string]string{})
		return
	}

	contentType := writer.FormDataContentType()

	if squareCrop != "" {
		err = writer.WriteField("_square_crop", squareCrop)
		if err != nil {
			vkErr = NewError(-1, err.Error(), "", map[string]string{})
			return
		}
	}

	writer.Close()

	resp, err := vk.Client.Post(uploadServer.UploadURL, contentType, body) // nolint: gosec
	if err != nil {
		vkErr = NewError(-1, err.Error(), "", map[string]string{})
		return
	}
	defer resp.Body.Close()

	bodyContent, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		vkErr = NewError(-1, err.Error(), "", map[string]string{})
		return
	}

	var handler object.PhotosOwnerUploadResponse
	err = json.Unmarshal(bodyContent, &handler)
	if err != nil {
		vkErr = NewError(-1, err.Error(), "", map[string]string{})
		return
	}

	response, vkErr = vk.PhotosSaveOwnerPhoto(map[string]string{
		"server": strconv.Itoa(handler.Server),
		"photo":  handler.Photo,
		"hash":   handler.Hash,
	})
	return response, vkErr
}

// UploadUserPhoto uploading Photos into User Profile
//
// Supported formats: JPG, PNG, GIF.
//
// Limits: size not less than 200x200px, aspect ratio from 0.25 to 3,
// width+height not more than 14000 px, file size up to 50 Mb.
func (vk *VK) UploadUserPhoto(file io.Reader) (response PhotosSaveOwnerPhotoResponse, vkErr Error) {
	response, vkErr = vk.uploadOwnerPhoto(map[string]string{}, "", file)
	return
}

// UploadOwnerPhoto uploading Photos into User Profile or Community
// To upload a photo to a community send its negative id in the owner_id parameter.
//
// Following parameters can be sent in addition:
// squareCrop in x,y,w (no quotes) format where x and y are the coordinates of
// the preview upper-right corner and w is square side length.
// That will create a square preview for a photo.
//
// Supported formats: JPG, PNG, GIF.
//
// Limits: size not less than 200x200px, aspect ratio from 0.25 to 3,
// width+height not more than 14000 px, file size up to 50 Mb.
func (vk *VK) UploadOwnerPhoto(ownerID int, squareCrop string, file io.Reader) (response PhotosSaveOwnerPhotoResponse, vkErr Error) {
	response, vkErr = vk.uploadOwnerPhoto(map[string]string{
		"owner_id": strconv.Itoa(ownerID),
	}, squareCrop, file)
	return
}

// UploadMessagesPhoto uploading Photos into a Private Message
//
// Supported formats: JPG, PNG, GIF.
//
// Limits: width+height not more than 14000 px, file size up to 50 Mb,
// aspect ratio of at least 1:20
func (vk *VK) UploadMessagesPhoto(peerID int, file io.Reader) (response PhotosSaveMessagesPhotoResponse, vkErr Error) {
	uploadServer, vkErr := vk.PhotosGetMessagesUploadServer(map[string]string{
		"peer_id": strconv.Itoa(peerID),
	})
	if vkErr.Code != 0 {
		return
	}

	bodyContent, err := vk.UploadFile(uploadServer.UploadURL, file, "photo", "photo.jpeg")
	if err != nil {
		vkErr = NewError(-1, err.Error(), "", map[string]string{})
		return
	}

	var handler object.PhotosMessageUploadResponse
	err = json.Unmarshal(bodyContent, &handler)
	if err != nil {
		vkErr = NewError(-1, err.Error(), "", map[string]string{})
		return
	}

	response, vkErr = vk.PhotosSaveMessagesPhoto(map[string]string{
		"server": strconv.Itoa(handler.Server),
		"photo":  handler.Photo,
		"hash":   handler.Hash,
	})
	return
}

// uploadChatPhoto uploading a Main Photo to a Group Chat
//
// Supported formats: JPG, PNG, GIF.
//
// Limits: size not less than 200x200px,
// width+height not more than 14000 px, file size up to 50 Mb,
// aspect ratio of at least 1:20
func (vk *VK) uploadChatPhoto(params map[string]string, file io.Reader) (response MessagesSetChatPhotoResponse, vkErr Error) {
	uploadServer, vkErr := vk.PhotosGetChatUploadServer(params)
	if vkErr.Code != 0 {
		return
	}

	bodyContent, err := vk.UploadFile(uploadServer.UploadURL, file, "file", "photo.jpeg")
	if err != nil {
		vkErr = NewError(-1, err.Error(), "", map[string]string{})
		return
	}

	var handler object.PhotosChatUploadResponse
	err = json.Unmarshal(bodyContent, &handler)
	if err != nil {
		vkErr = NewError(-1, err.Error(), "", map[string]string{})
		return
	}

	response, vkErr = vk.MessagesSetChatPhoto(map[string]string{
		"file": handler.Response,
	})
	return
}

// UploadChatPhoto uploading a Main Photo to a Group Chat without crop
//
// Supported formats: JPG, PNG, GIF.
//
// Limits: size not less than 200x200px,
// width+height not more than 14000 px, file size up to 50 Mb,
// aspect ratio of at least 1:20
func (vk *VK) UploadChatPhoto(chatID int, file io.Reader) (response MessagesSetChatPhotoResponse, vkErr Error) {
	response, vkErr = vk.uploadChatPhoto(map[string]string{
		"chat_id": strconv.Itoa(chatID),
	}, file)
	return
}

// UploadChatPhotoCrop uploading a Main Photo to a Group Chat with crop
//
// Supported formats: JPG, PNG, GIF.
//
// Limits: size not less than 200x200px,
// width+height not more than 14000 px, file size up to 50 Mb,
// aspect ratio of at least 1:20
func (vk *VK) UploadChatPhotoCrop(chatID, cropX, cropY, cropWidth int, file io.Reader) (response MessagesSetChatPhotoResponse, vkErr Error) {
	response, vkErr = vk.uploadChatPhoto(map[string]string{
		"chat_id":    strconv.Itoa(chatID),
		"crop_x":     strconv.Itoa(cropX),
		"crop_y":     strconv.Itoa(cropY),
		"crop_width": strconv.Itoa(cropWidth),
	}, file)
	return
}

// uploadMarketPhoto uploading a Market Item Photo
//
// Supported formats: JPG, PNG, GIF.
//
// Limits: size not less than 400x400px,
// width+height not more than 14000 px, file size up to 50 Mb,
// aspect ratio of at least 1:20
func (vk *VK) uploadMarketPhoto(params map[string]string, file io.Reader) (response PhotosSaveMarketPhotoResponse, vkErr Error) {
	uploadServer, vkErr := vk.PhotosGetMarketUploadServer(params)
	if vkErr.Code != 0 {
		return
	}

	bodyContent, err := vk.UploadFile(uploadServer.UploadURL, file, "file", "photo.jpeg")
	if err != nil {
		vkErr = NewError(-1, err.Error(), "", map[string]string{})
		return
	}

	var handler object.PhotosMarketUploadResponse
	err = json.Unmarshal(bodyContent, &handler)
	if err != nil {
		vkErr = NewError(-1, err.Error(), "", map[string]string{})
		return
	}

	response, vkErr = vk.PhotosSaveMarketPhoto(map[string]string{
		"group_id":  params["group_id"],
		"server":    strconv.Itoa(handler.Server),
		"photo":     handler.Photo,
		"hash":      handler.Hash,
		"crop_data": handler.CropData,
		"crop_hash": handler.CropHash,
	})
	return
}

// UploadMarketPhoto uploading a Market Item Photo without crop
//
// Supported formats: JPG, PNG, GIF.
//
// Limits: size not less than 400x400px,
// width+height not more than 14000 px, file size up to 50 Mb,
// aspect ratio of at least 1:20
func (vk *VK) UploadMarketPhoto(groupID int, mainPhoto bool, file io.Reader) (response PhotosSaveMarketPhotoResponse, vkErr Error) {
	mainPhotoString := "0"
	if mainPhoto {
		mainPhotoString = "1"
	}
	response, vkErr = vk.uploadMarketPhoto(map[string]string{
		"group_id":   strconv.Itoa(groupID),
		"main_photo": mainPhotoString,
	}, file)
	return
}

// UploadMarketPhotoCrop uploading a Market Item Photo with crop
//
// Supported formats: JPG, PNG, GIF.
//
// Limits: size not less than 400x400px,
// width+height not more than 14000 px, file size up to 50 Mb,
// aspect ratio of at least 1:20
func (vk *VK) UploadMarketPhotoCrop(groupID, cropX, cropY, cropWidth int, file io.Reader) (response PhotosSaveMarketPhotoResponse, vkErr Error) {
	response, vkErr = vk.uploadMarketPhoto(map[string]string{
		"group_id":   strconv.Itoa(groupID),
		"main_photo": "1",
		"crop_x":     strconv.Itoa(cropX),
		"crop_y":     strconv.Itoa(cropY),
		"crop_width": strconv.Itoa(cropWidth),
	}, file)
	return
}

// uploadMarketAlbumPhoto uploading a Main Photo to a Group Chat
//
// Supported formats: JPG, PNG, GIF.
//
// Limits: size not less than 1280x720px,
// width+height not more than 14000 px, file size up to 50 Mb,
// aspect ratio of at least 1:20
func (vk *VK) UploadMarketAlbumPhoto(groupID int, file io.Reader) (response PhotosSaveMarketAlbumPhotoResponse, vkErr Error) {
	uploadServer, vkErr := vk.PhotosGetMarketAlbumUploadServer(map[string]string{
		"group_id": strconv.Itoa(groupID),
	})
	if vkErr.Code != 0 {
		return
	}

	bodyContent, err := vk.UploadFile(uploadServer.UploadURL, file, "file", "photo.jpeg")
	if err != nil {
		vkErr = NewError(-1, err.Error(), "", map[string]string{})
		return
	}

	var handler object.PhotosMarketAlbumUploadResponse
	err = json.Unmarshal(bodyContent, &handler)
	if err != nil {
		vkErr = NewError(-1, err.Error(), "", map[string]string{})
		return
	}

	response, vkErr = vk.PhotosSaveMarketAlbumPhoto(map[string]string{
		"group_id": strconv.Itoa(groupID),
		"server":   strconv.Itoa(handler.Server),
		"photo":    handler.Photo,
		"hash":     handler.Hash,
	})
	return
}

// UploadVideo uploading Video Files
//
// Supported formats: AVI, MP4, 3GP, MPEG, MOV, FLV, WMV.
func (vk *VK) UploadVideo(params map[string]string, file io.Reader) (response VideoSaveResponse, vkErr Error) {
	response, vkErr = vk.VideoSave(params)
	if vkErr.Code != 0 {
		return
	}

	bodyContent, err := vk.UploadFile(response.UploadURL, file, "video_file", "video")
	if err != nil {
		vkErr = NewError(-1, err.Error(), "", params)
		return
	}

	var videoUploadError uploadError
	err = json.Unmarshal(bodyContent, &videoUploadError)
	if err != nil {
		vkErr = NewError(-1, err.Error(), "", map[string]string{})
	}
	if videoUploadError.ErrorCode != 0 {
		vkErr = NewError(-1, videoUploadError.Error, "", map[string]string{})
	}
	return
}
