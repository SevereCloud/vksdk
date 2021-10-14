package api // import "github.com/SevereCloud/vksdk/v2/api"

import (
	"github.com/SevereCloud/vksdk/v2/object"
)

// MarusiaGetPictureUploadLinkResponse struct.
type MarusiaGetPictureUploadLinkResponse struct {
	PictureUploadLink string `json:"picture_upload_link"` // Link
}

// MarusiaGetPictureUploadLink method.
//
// https://vk.com/dev/marusia_skill_docs10
func (vk *VK) MarusiaGetPictureUploadLink(params Params) (response MarusiaGetPictureUploadLinkResponse, err error) {
	err = vk.RequestUnmarshal("marusia.getPictureUploadLink", &response, params)
	return
}

// MarusiaSavePictureResponse struct.
type MarusiaSavePictureResponse struct {
	AppID   int `json:"app_id"`
	PhotoID int `json:"photo_id"`
}

// MarusiaSavePicture method.
//
// https://vk.com/dev/marusia_skill_docs10
func (vk *VK) MarusiaSavePicture(params Params) (response MarusiaSavePictureResponse, err error) {
	err = vk.RequestUnmarshal("marusia.savePicture", &response, params)
	return
}

// MarusiaGetPicturesResponse struct.
type MarusiaGetPicturesResponse struct {
	Count int                     `json:"count"`
	Items []object.MarusiaPicture `json:"items"`
}

// MarusiaGetPictures method.
//
// https://vk.com/dev/marusia_skill_docs10
func (vk *VK) MarusiaGetPictures(params Params) (response MarusiaGetPicturesResponse, err error) {
	err = vk.RequestUnmarshal("marusia.getPictures", &response, params)
	return
}

// MarusiaDeletePicture delete picture.
//
// https://vk.com/dev/marusia_skill_docs10
func (vk *VK) MarusiaDeletePicture(params Params) (response int, err error) {
	err = vk.RequestUnmarshal("marusia.deletePicture", &response, params)
	return
}
