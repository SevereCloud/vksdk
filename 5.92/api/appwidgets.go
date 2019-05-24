package api // import "github.com/SevereCloud/vksdk/5.92/api"

// AppWidgetsGetAppImageUploadServerResponse struct
type AppWidgetsGetAppImageUploadServerResponse struct {
	UploadURL string `json:"upload_url"`
}

// AppWidgetsGetAppImageUploadServer returns a URL for uploading a
// photo to the app collection for community app widgets.
//
// https://vk.com/dev/appWidgets.getAppImageUploadServer
func (vk *VK) AppWidgetsGetAppImageUploadServer(params map[string]string) (response AppWidgetsGetAppImageUploadServerResponse, vkErr Error) {
	vk.RequestUnmarshal("appWidgets.getAppImageUploadServer", params, &response, &vkErr)
	return
}

// AppWidgetsGetAppImagesResponse struct
type AppWidgetsGetAppImagesResponse struct {
	Count int `json:"count"`
	Items []struct {
		ID     string `json:"id"`
		Type   string `json:"type"`
		Images struct {
			URL    string `json:"url"`
			Width  int    `json:"width"`
			Height int    `json:"height"`
		} `json:"images"`
	} `json:"items"`
}

// AppWidgetsGetAppImages returns an app collection of images for community app widgets.
//
// https://vk.com/dev/appWidgets.getAppImages
func (vk *VK) AppWidgetsGetAppImages(params map[string]string) (response AppWidgetsGetAppImagesResponse, vkErr Error) {
	vk.RequestUnmarshal("appWidgets.getAppImages", params, &response, &vkErr)
	return
}

// AppWidgetsGetGroupImageUploadServerResponse struct
type AppWidgetsGetGroupImageUploadServerResponse struct {
	UploadURL string `json:"upload_url"`
}

// AppWidgetsGetGroupImageUploadServer returns a URL for uploading
// a photo to the community collection for community app widgets.
//
// https://vk.com/dev/appWidgets.getGroupImageUploadServer
func (vk *VK) AppWidgetsGetGroupImageUploadServer(params map[string]string) (response AppWidgetsGetGroupImageUploadServerResponse, vkErr Error) {
	vk.RequestUnmarshal("appWidgets.getGroupImageUploadServer", params, &response, &vkErr)
	return
}

// AppWidgetsGetGroupImagesResponse struct
type AppWidgetsGetGroupImagesResponse struct {
	Count int `json:"count"`
	Items []struct {
		ID     string `json:"id"`
		Type   string `json:"type"`
		Images struct {
			URL    string `json:"url"`
			Width  int    `json:"width"`
			Height int    `json:"height"`
		} `json:"images"`
	} `json:"items"`
}

// AppWidgetsGetGroupImages returns a community collection of images for community app widgets.
//
// https://vk.com/dev/appWidgets.getGroupImages
func (vk *VK) AppWidgetsGetGroupImages(params map[string]string) (response AppWidgetsGetGroupImagesResponse, vkErr Error) {
	vk.RequestUnmarshal("appWidgets.getGroupImages", params, &response, &vkErr)
	return
}

// AppWidgetsGetImagesByIDResponse struct
type AppWidgetsGetImagesByIDResponse struct {
	ID     string `json:"id"`
	Type   string `json:"type"`
	Images struct {
		URL    string `json:"url"`
		Width  int    `json:"width"`
		Height int    `json:"height"`
	} `json:"images"`
}

// AppWidgetsGetImagesByID returns an image for community app widgets by its ID.
//
// https://vk.com/dev/appWidgets.getImagesById
func (vk *VK) AppWidgetsGetImagesByID(params map[string]string) (response AppWidgetsGetImagesByIDResponse, vkErr Error) {
	vk.RequestUnmarshal("appWidgets.getImagesById", params, &response, &vkErr)
	return
}

// AppWidgetsSaveAppImageResponse struct
type AppWidgetsSaveAppImageResponse struct {
	ID     string `json:"id"`
	Type   string `json:"type"`
	Images struct {
		URL    string `json:"url"`
		Width  int    `json:"width"`
		Height int    `json:"height"`
	} `json:"images"`
}

// AppWidgetsSaveAppImage allows to save image into app collection for community app widgets.
//
// https://vk.com/dev/appWidgets.saveAppImage
func (vk *VK) AppWidgetsSaveAppImage(params map[string]string) (response AppWidgetsSaveAppImageResponse, vkErr Error) {
	vk.RequestUnmarshal("appWidgets.saveAppImage", params, &response, &vkErr)
	return
}

// AppWidgetsSaveGroupImageResponse struct
type AppWidgetsSaveGroupImageResponse struct {
	ID     string `json:"id"`
	Type   string `json:"type"`
	Images struct {
		URL    string `json:"url"`
		Width  int    `json:"width"`
		Height int    `json:"height"`
	} `json:"images"`
}

// AppWidgetsSaveGroupImage allows to save image into community collection for community app widgets.
//
// https://vk.com/dev/appWidgets.saveGroupImage
func (vk *VK) AppWidgetsSaveGroupImage(params map[string]string) (response AppWidgetsSaveGroupImageResponse, vkErr Error) {
	vk.RequestUnmarshal("appWidgets.saveGroupImage", params, &response, &vkErr)
	return
}

// AppWidgetsUpdateResponse struct
type AppWidgetsUpdateResponse int

// AppWidgetsUpdate allows to update community app widget.
//
// https://vk.com/dev/appWidgets.update
func (vk *VK) AppWidgetsUpdate(params map[string]string) (response AppWidgetsUpdateResponse, vkErr Error) {
	vk.RequestUnmarshal("appWidgets.update", params, &response, &vkErr)

	return
}
