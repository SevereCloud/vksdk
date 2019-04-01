package api // import "github.com/severecloud/vksdk/5.92/api"

// AppWidgetsGetAppImageUploadServerResponse struct
type AppWidgetsGetAppImageUploadServerResponse struct{}

// TODO: appWidgets.getAppImageUploadServer returns a URL for uploading a photo to the app collection for community app widgets.
// https://vk.com/dev/appWidgets.getAppImageUploadServer

// AppWidgetsGetAppImagesResponse struct
type AppWidgetsGetAppImagesResponse struct{}

// TODO: appWidgets.getAppImages returns an app collection of images for community app widgets.
// https://vk.com/dev/appWidgets.getAppImages

// AppWidgetsGetGroupImageUploadServerResponse struct
type AppWidgetsGetGroupImageUploadServerResponse struct{}

// TODO: appWidgets.getGroupImageUploadServer returns a URL for uploading a photo to the community collection for community app widgets.
// https://vk.com/dev/appWidgets.getGroupImageUploadServer

// AppWidgetsGetGroupImagesResponse struct
type AppWidgetsGetGroupImagesResponse struct{}

// TODO: appWidgets.getGroupImages returns a community collection of images for community app widgets.
// https://vk.com/dev/appWidgets.getGroupImages

// AppWidgetsGetImagesByIDResponse struct
type AppWidgetsGetImagesByIDResponse struct{}

// TODO: appWidgets.getImagesById returns an image for community app widgets by its ID.
// https://vk.com/dev/appWidgets.getImagesById

// AppWidgetsSaveAppImageResponse struct
type AppWidgetsSaveAppImageResponse struct{}

// TODO: appWidgets.saveAppImage allows to save image into app collection for community app widgets.
// https://vk.com/dev/appWidgets.saveAppImage

// AppWidgetsSaveGroupImageResponse struct
type AppWidgetsSaveGroupImageResponse struct{}

// TODO: appWidgets.saveGroupImage allows to save image into community collection for community app widgets.
// https://vk.com/dev/appWidgets.saveGroupImage

// AppWidgetsUpdate allows to update community app widget.
// https://vk.com/dev/appWidgets.update
func (vk VK) AppWidgetsUpdate(params map[string]string) (vkErr Error) {
	_, vkErr = vk.Request("appWidgets.update", params)

	return
}
