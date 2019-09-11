package api // import "github.com/SevereCloud/vksdk/api"

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"mime/multipart"
	"strconv"

	"github.com/SevereCloud/vksdk/object"
)

type uploadError struct {
	Error         string `json:"error"`
	ErrorCode     int    `json:"error_code"`
	ErrorDescr    string `json:"error_descr"`
	ErrorIsLogged bool   `json:"error_is_logged"`
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
func (vk *VK) uploadPhoto(params map[string]string, file io.Reader) (response PhotosSaveResponse, err error) {
	uploadServer, err := vk.PhotosGetUploadServer(params)
	if err != nil {
		return
	}

	bodyContent, err := vk.UploadFile(uploadServer.UploadURL, file, "file1", "file1.jpeg")
	if err != nil {
		return
	}

	var handler object.PhotosPhotoUploadResponse
	err = json.Unmarshal(bodyContent, &handler)
	if err != nil {
		return
	}

	response, err = vk.PhotosSave(map[string]string{
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
func (vk *VK) UploadPhoto(albumID int, file io.Reader) (response PhotosSaveResponse, err error) {
	response, err = vk.uploadPhoto(map[string]string{
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
func (vk *VK) UploadPhotoGroup(groupID, albumID int, file io.Reader) (response PhotosSaveResponse, err error) {
	response, err = vk.uploadPhoto(map[string]string{
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
func (vk *VK) uploadWallPhoto(params map[string]string, file io.Reader) (response PhotosSaveWallPhotoResponse, err error) {
	uploadServer, err := vk.PhotosGetWallUploadServer(params)
	if err != nil {
		return
	}

	bodyContent, err := vk.UploadFile(uploadServer.UploadURL, file, "photo", "photo.jpeg")
	if err != nil {
		return
	}

	var handler object.PhotosWallUploadResponse
	err = json.Unmarshal(bodyContent, &handler)
	if err != nil {
		return
	}

	response, err = vk.PhotosSaveWallPhoto(map[string]string{
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
func (vk *VK) UploadWallPhoto(file io.Reader) (response PhotosSaveWallPhotoResponse, err error) {
	response, err = vk.uploadWallPhoto(map[string]string{}, file)
	return
}

// UploadGroupWallPhoto uploading Photos on Group Wall
//
// Supported formats: JPG, PNG, GIF.
//
// Limits: width+height not more than 14000 px, file size up to 50 Mb,
// aspect ratio of at least 1:20
func (vk *VK) UploadGroupWallPhoto(groupID int, file io.Reader) (response PhotosSaveWallPhotoResponse, err error) {
	response, err = vk.uploadWallPhoto(map[string]string{
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
func (vk *VK) uploadOwnerPhoto(params map[string]string, squareCrop string, file io.Reader) (response PhotosSaveOwnerPhotoResponse, err error) {
	uploadServer, err := vk.PhotosGetOwnerPhotoUploadServer(params)
	if err != nil {
		return
	}

	body := new(bytes.Buffer)
	writer := multipart.NewWriter(body)
	part, err := writer.CreateFormFile("photo", "photo.jpeg")
	if err != nil {
		return
	}

	_, err = io.Copy(part, file)
	if err != nil {
		return
	}

	contentType := writer.FormDataContentType()

	if squareCrop != "" {
		err = writer.WriteField("_square_crop", squareCrop)
		if err != nil {
			return
		}
	}

	writer.Close()

	resp, err := vk.Client.Post(uploadServer.UploadURL, contentType, body) // nolint: gosec
	if err != nil {
		return
	}
	defer resp.Body.Close()

	bodyContent, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return
	}

	var handler object.PhotosOwnerUploadResponse
	err = json.Unmarshal(bodyContent, &handler)
	if err != nil {
		return
	}

	response, err = vk.PhotosSaveOwnerPhoto(map[string]string{
		"server": strconv.Itoa(handler.Server),
		"photo":  handler.Photo,
		"hash":   handler.Hash,
	})
	return response, err
}

// UploadUserPhoto uploading Photos into User Profile
//
// Supported formats: JPG, PNG, GIF.
//
// Limits: size not less than 200x200px, aspect ratio from 0.25 to 3,
// width+height not more than 14000 px, file size up to 50 Mb.
func (vk *VK) UploadUserPhoto(file io.Reader) (response PhotosSaveOwnerPhotoResponse, err error) {
	response, err = vk.uploadOwnerPhoto(map[string]string{}, "", file)
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
func (vk *VK) UploadOwnerPhoto(ownerID int, squareCrop string, file io.Reader) (response PhotosSaveOwnerPhotoResponse, err error) {
	response, err = vk.uploadOwnerPhoto(map[string]string{
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
func (vk *VK) UploadMessagesPhoto(peerID int, file io.Reader) (response PhotosSaveMessagesPhotoResponse, err error) {
	uploadServer, err := vk.PhotosGetMessagesUploadServer(map[string]string{
		"peer_id": strconv.Itoa(peerID),
	})
	if err != nil {
		return
	}

	bodyContent, err := vk.UploadFile(uploadServer.UploadURL, file, "photo", "photo.jpeg")
	if err != nil {
		return
	}

	var handler object.PhotosMessageUploadResponse
	err = json.Unmarshal(bodyContent, &handler)
	if err != nil {
		return
	}

	response, err = vk.PhotosSaveMessagesPhoto(map[string]string{
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
func (vk *VK) uploadChatPhoto(params map[string]string, file io.Reader) (response MessagesSetChatPhotoResponse, err error) {
	uploadServer, err := vk.PhotosGetChatUploadServer(params)
	if err != nil {
		return
	}

	bodyContent, err := vk.UploadFile(uploadServer.UploadURL, file, "file", "photo.jpeg")
	if err != nil {
		return
	}

	var handler object.PhotosChatUploadResponse
	err = json.Unmarshal(bodyContent, &handler)
	if err != nil {
		return
	}

	response, err = vk.MessagesSetChatPhoto(map[string]string{
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
func (vk *VK) UploadChatPhoto(chatID int, file io.Reader) (response MessagesSetChatPhotoResponse, err error) {
	response, err = vk.uploadChatPhoto(map[string]string{
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
func (vk *VK) UploadChatPhotoCrop(chatID, cropX, cropY, cropWidth int, file io.Reader) (response MessagesSetChatPhotoResponse, err error) {
	response, err = vk.uploadChatPhoto(map[string]string{
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
func (vk *VK) uploadMarketPhoto(params map[string]string, file io.Reader) (response PhotosSaveMarketPhotoResponse, err error) {
	uploadServer, err := vk.PhotosGetMarketUploadServer(params)
	if err != nil {
		return
	}

	bodyContent, err := vk.UploadFile(uploadServer.UploadURL, file, "file", "photo.jpeg")
	if err != nil {
		return
	}

	var handler object.PhotosMarketUploadResponse
	err = json.Unmarshal(bodyContent, &handler)
	if err != nil {
		return
	}

	response, err = vk.PhotosSaveMarketPhoto(map[string]string{
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
func (vk *VK) UploadMarketPhoto(groupID int, mainPhoto bool, file io.Reader) (response PhotosSaveMarketPhotoResponse, err error) {
	mainPhotoString := "0"
	if mainPhoto {
		mainPhotoString = "1"
	}
	response, err = vk.uploadMarketPhoto(map[string]string{
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
func (vk *VK) UploadMarketPhotoCrop(groupID, cropX, cropY, cropWidth int, file io.Reader) (response PhotosSaveMarketPhotoResponse, err error) {
	response, err = vk.uploadMarketPhoto(map[string]string{
		"group_id":   strconv.Itoa(groupID),
		"main_photo": "1",
		"crop_x":     strconv.Itoa(cropX),
		"crop_y":     strconv.Itoa(cropY),
		"crop_width": strconv.Itoa(cropWidth),
	}, file)
	return
}

// UploadMarketAlbumPhoto uploading a Main Photo to a Group Chat
//
// Supported formats: JPG, PNG, GIF.
//
// Limits: size not less than 1280x720px,
// width+height not more than 14000 px, file size up to 50 Mb,
// aspect ratio of at least 1:20
func (vk *VK) UploadMarketAlbumPhoto(groupID int, file io.Reader) (response PhotosSaveMarketAlbumPhotoResponse, err error) {
	uploadServer, err := vk.PhotosGetMarketAlbumUploadServer(map[string]string{
		"group_id": strconv.Itoa(groupID),
	})
	if err != nil {
		return
	}

	bodyContent, err := vk.UploadFile(uploadServer.UploadURL, file, "file", "photo.jpeg")
	if err != nil {
		return
	}

	var handler object.PhotosMarketAlbumUploadResponse
	err = json.Unmarshal(bodyContent, &handler)
	if err != nil {
		return
	}

	response, err = vk.PhotosSaveMarketAlbumPhoto(map[string]string{
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
func (vk *VK) UploadVideo(params map[string]string, file io.Reader) (response VideoSaveResponse, err error) {
	response, err = vk.VideoSave(params)
	if err != nil {
		return
	}

	bodyContent, err := vk.UploadFile(response.UploadURL, file, "video_file", "video.mp4")
	if err != nil {
		return
	}

	var videoUploadError uploadError
	err = json.Unmarshal(bodyContent, &videoUploadError)
	if err != nil {
		return
	}
	if videoUploadError.ErrorCode != 0 {
		err = fmt.Errorf(videoUploadError.Error)
	}
	return
}

// uploadDoc uploading Documents
//
// Supported formats: any formats excepting mp3 and executable files.
//
// Limits: file size up to 200 MB.
func (vk *VK) uploadDoc(url, title, tags string, file io.Reader) (response DocsSaveResponse, err error) {
	bodyContent, err := vk.UploadFile(url, file, "file", title)
	if err != nil {
		return
	}

	var docUploadError uploadError
	err = json.Unmarshal(bodyContent, &docUploadError)
	if err != nil {
		return
	}
	if docUploadError.Error != "" {
		err = fmt.Errorf(docUploadError.Error)
		return
	}

	var handler object.DocsDocUploadResponse
	err = json.Unmarshal(bodyContent, &handler)
	if err != nil {
		return
	}

	response, err = vk.DocsSave(map[string]string{
		"file":  handler.File,
		"title": title,
		"tags":  tags,
	})
	return response, err
}

// UploadDoc uploading Documents
//
// Supported formats: any formats excepting mp3 and executable files.
//
// Limits: file size up to 200 MB.
func (vk *VK) UploadDoc(title, tags string, file io.Reader) (response DocsSaveResponse, err error) {
	uploadServer, err := vk.DocsGetUploadServer(map[string]string{})
	if err != nil {
		return
	}
	response, err = vk.uploadDoc(uploadServer.UploadURL, title, tags, file)
	return
}

// UploadGroupDoc uploading Documents into Community
//
// Supported formats: any formats excepting mp3 and executable files.
//
// Limits: file size up to 200 MB.
func (vk *VK) UploadGroupDoc(groupID int, title, tags string, file io.Reader) (response DocsSaveResponse, err error) {
	uploadServer, err := vk.DocsGetUploadServer(map[string]string{
		"group_id": strconv.Itoa(groupID),
	})
	if err != nil {
		return
	}
	response, err = vk.uploadDoc(uploadServer.UploadURL, title, tags, file)
	return
}

// UploadWallDoc uploading Documents on Wall
//
// Supported formats: any formats excepting mp3 and executable files.
//
// Limits: file size up to 200 MB.
func (vk *VK) UploadWallDoc(title, tags string, file io.Reader) (response DocsSaveResponse, err error) {
	uploadServer, err := vk.DocsGetWallUploadServer(map[string]string{})
	if err != nil {
		return
	}
	response, err = vk.uploadDoc(uploadServer.UploadURL, title, tags, file)
	return
}

// UploadGroupWallDoc uploading Documents on Group Wall
//
// Supported formats: any formats excepting mp3 and executable files.
//
// Limits: file size up to 200 MB.
func (vk *VK) UploadGroupWallDoc(groupID int, title, tags string, file io.Reader) (response DocsSaveResponse, err error) {
	uploadServer, err := vk.DocsGetWallUploadServer(map[string]string{
		"group_id": strconv.Itoa(groupID),
	})
	if err != nil {
		return
	}
	response, err = vk.uploadDoc(uploadServer.UploadURL, title, tags, file)
	return
}

// UploadMessagesDoc uploading Documents into a Private Message
//
// Supported formats: any formats excepting mp3 and executable files.
//
// Limits: file size up to 200 MB.
func (vk *VK) UploadMessagesDoc(peerID int, typeDoc, title, tags string, file io.Reader) (response DocsSaveResponse, err error) {
	uploadServer, err := vk.DocsGetMessagesUploadServer(map[string]string{
		"peer_id": strconv.Itoa(peerID),
		"type":    typeDoc,
	})
	if err != nil {
		return
	}
	response, err = vk.uploadDoc(uploadServer.UploadURL, title, tags, file)
	return
}

// UploadOwneCoverrPhoto uploading a Main Photo to a Group Chat
//
// Supported formats: JPG, PNG, GIF.
//
// Limits: minimum photo size 795x200px, width+height not more than 14000px,
// file size up to 50 MB. Recommended size: 1590x400px.
func (vk *VK) UploadOwnerCoverPhoto(groupID, cropX, cropY, cropX2, cropY2 int, file io.Reader) (response PhotosSaveOwnerCoverPhotoResponse, err error) {
	uploadServer, err := vk.PhotosGetOwnerCoverPhotoUploadServer(map[string]string{
		"group_id": strconv.Itoa(groupID),
		"crop_x":   strconv.Itoa(cropX),
		"crop_y":   strconv.Itoa(cropY),
		"crop_x2":  strconv.Itoa(cropX2),
		"crop_y2":  strconv.Itoa(cropY2),
	})
	if err != nil {
		return
	}

	bodyContent, err := vk.UploadFile(uploadServer.UploadURL, file, "photo", "photo.jpeg")
	if err != nil {
		return
	}

	var handler object.PhotosOwnerUploadResponse
	err = json.Unmarshal(bodyContent, &handler)
	if err != nil {
		return
	}

	response, err = vk.PhotosSaveOwnerCoverPhoto(map[string]string{
		"photo": handler.Photo,
		"hash":  handler.Hash,
	})
	return
}

type UploadStories struct {
	Stories object.StoriesStory `json:"stories"`
	Sig     string              `json:"_sig"`
}

type rawUploadStoriesPhoto struct {
	UploadStories
	Error struct {
		ErrorCode int    `json:"error_code"`
		Type      string `json:"type"`
	} `json:"error"`
}

type rawUploadStoriesVideo struct {
	UploadStories
	uploadError
}

// UploadStoriesPhoto uploading Story
//
// Supported formats: JPG, PNG, GIF.
// Limits: sum of with and height no more than 14000px, file size no more than 10 MB. Video format: h264 video, aac audio, maximum 720х1280, 30fps.
func (vk *VK) UploadStoriesPhoto(params map[string]string, file io.Reader) (response UploadStories, err error) {
	uploadServer, err := vk.StoriesGetPhotoUploadServer(params)
	if err != nil {
		return
	}

	bodyContent, err := vk.UploadFile(uploadServer.UploadURL, file, "file", "file.jpeg")
	if err != nil {
		return
	}

	var handler rawUploadStoriesPhoto
	err = json.Unmarshal(bodyContent, &handler)
	if err != nil {
		return
	}
	if handler.Error.ErrorCode != 0 {
		err = fmt.Errorf(handler.Error.Type)
	} else {
		response.Sig = handler.Sig
		response.Stories = handler.Stories
	}

	return
}

// UploadStoriesVideo uploading Story
//
// Video format: h264 video, aac audio, maximum 720х1280, 30fps.
func (vk *VK) UploadStoriesVideo(params map[string]string, file io.Reader) (response UploadStories, err error) {
	uploadServer, err := vk.StoriesGetVideoUploadServer(params)
	if err != nil {
		return
	}

	bodyContent, err := vk.UploadFile(uploadServer.UploadURL, file, "video_file", "video.mp4")
	if err != nil {
		return
	}

	var handler rawUploadStoriesVideo
	err = json.Unmarshal(bodyContent, &handler)
	if err != nil {
		return
	}
	if handler.ErrorCode != 0 {
		// TODO: new type error
		err = fmt.Errorf(handler.Error)
	} else {
		response.Sig = handler.Sig
		response.Stories = handler.Stories
	}

	return
}
