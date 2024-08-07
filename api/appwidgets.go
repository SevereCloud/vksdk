package api // import "github.com/SevereCloud/vksdk/v3/api"

import (
	"github.com/SevereCloud/vksdk/v3/object"
)

// AppWidgetsGetAppImageUploadServerResponse struct.
type AppWidgetsGetAppImageUploadServerResponse struct {
	UploadURL string `json:"upload_url"`
}

// AppWidgetsGetAppImageUploadServer returns a URL for uploading a
// photo to the app collection for community app widgets.
//
// https://dev.vk.com/method/appWidgets.getAppImageUploadServer
func (vk *VK) AppWidgetsGetAppImageUploadServer(params Params) (
	response AppWidgetsGetAppImageUploadServerResponse,
	err error,
) {
	err = vk.RequestUnmarshal("appWidgets.getAppImageUploadServer", &response, params)
	return
}

// AppWidgetsGetAppImagesResponse struct.
type AppWidgetsGetAppImagesResponse struct {
	Count int                      `json:"count"`
	Items []object.AppWidgetsImage `json:"items"`
}

// AppWidgetsGetAppImages returns an app collection of images for community app widgets.
//
// https://dev.vk.com/method/appWidgets.getAppImages
func (vk *VK) AppWidgetsGetAppImages(params Params) (response AppWidgetsGetAppImagesResponse, err error) {
	err = vk.RequestUnmarshal("appWidgets.getAppImages", &response, params)
	return
}

// AppWidgetsGetGroupImageUploadServerResponse struct.
type AppWidgetsGetGroupImageUploadServerResponse struct {
	UploadURL string `json:"upload_url"`
}

// AppWidgetsGetGroupImageUploadServer returns a URL for uploading
// a photo to the community collection for community app widgets.
//
// https://dev.vk.com/method/appWidgets.getGroupImageUploadServer
func (vk *VK) AppWidgetsGetGroupImageUploadServer(params Params) (
	response AppWidgetsGetGroupImageUploadServerResponse,
	err error,
) {
	err = vk.RequestUnmarshal("appWidgets.getGroupImageUploadServer", &response, params)
	return
}

// AppWidgetsGetGroupImagesResponse struct.
type AppWidgetsGetGroupImagesResponse struct {
	Count int                      `json:"count"`
	Items []object.AppWidgetsImage `json:"items"`
}

// AppWidgetsGetGroupImages returns a community collection of images for community app widgets.
//
// https://dev.vk.com/method/appWidgets.getGroupImages
func (vk *VK) AppWidgetsGetGroupImages(params Params) (response AppWidgetsGetGroupImagesResponse, err error) {
	err = vk.RequestUnmarshal("appWidgets.getGroupImages", &response, params)
	return
}

// AppWidgetsGetImagesByID returns an image for community app widgets by its ID.
//
// https://dev.vk.com/method/appWidgets.getImagesById
func (vk *VK) AppWidgetsGetImagesByID(params Params) (response object.AppWidgetsImage, err error) {
	err = vk.RequestUnmarshal("appWidgets.getImagesById", &response, params)
	return
}

// AppWidgetsSaveAppImage allows to save image into app collection for community app widgets.
//
// https://dev.vk.com/method/appWidgets.saveAppImage
func (vk *VK) AppWidgetsSaveAppImage(params Params) (response object.AppWidgetsImage, err error) {
	err = vk.RequestUnmarshal("appWidgets.saveAppImage", &response, params)
	return
}

// AppWidgetsSaveGroupImage allows to save image into community collection for community app widgets.
//
// https://dev.vk.com/method/appWidgets.saveGroupImage
func (vk *VK) AppWidgetsSaveGroupImage(params Params) (response object.AppWidgetsImage, err error) {
	err = vk.RequestUnmarshal("appWidgets.saveGroupImage", &response, params)
	return
}

// AppWidgetsUpdate allows to update community app widget.
//
// https://dev.vk.com/method/appWidgets.update
func (vk *VK) AppWidgetsUpdate(params Params) (response int, err error) {
	err = vk.RequestUnmarshal("appWidgets.update", &response, params)

	return
}
