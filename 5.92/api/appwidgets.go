package api // import "github.com/severecloud/vksdk/5.92/api"

import "encoding/json"

// AppWidgetsGetAppImageUploadServerResponse struct
type AppWidgetsGetAppImageUploadServerResponse struct {
	UploadURL string `json:"upload_url"`
}

// AppWidgetsGetAppImageUploadServer returns a URL for uploading a photo to the app collection for community app widgets.
// https://vk.com/dev/appWidgets.getAppImageUploadServer
func (vk VK) AppWidgetsGetAppImageUploadServer(params map[string]string) (response AppWidgetsGetAppImageUploadServerResponse, vkErr Error) {
	rawResponse, vkErr := vk.Request("appWidgets.getAppImageUploadServer", params)
	if vkErr.Code != 0 {
		return
	}

	err := json.Unmarshal(rawResponse, &response)
	if err != nil {
		panic(err)
	}

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
// https://vk.com/dev/appWidgets.getAppImages
func (vk VK) AppWidgetsGetAppImages(params map[string]string) (response AppWidgetsGetAppImagesResponse, vkErr Error) {
	rawResponse, vkErr := vk.Request("appWidgets.getAppImages", params)
	if vkErr.Code != 0 {
		return
	}

	err := json.Unmarshal(rawResponse, &response)
	if err != nil {
		panic(err)
	}

	return
}

// AppWidgetsGetGroupImageUploadServerResponse struct
type AppWidgetsGetGroupImageUploadServerResponse struct {
	UploadURL string `json:"upload_url"`
}

// AppWidgetsGetGroupImageUploadServer returns a URL for uploading a photo to the community collection for community app widgets.
// https://vk.com/dev/appWidgets.getGroupImageUploadServer
func (vk VK) AppWidgetsGetGroupImageUploadServer(params map[string]string) (response AppWidgetsGetGroupImageUploadServerResponse, vkErr Error) {
	rawResponse, vkErr := vk.Request("appWidgets.getGroupImageUploadServer", params)
	if vkErr.Code != 0 {
		return
	}

	err := json.Unmarshal(rawResponse, &response)
	if err != nil {
		panic(err)
	}

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
// https://vk.com/dev/appWidgets.getGroupImages
func (vk VK) AppWidgetsGetGroupImages(params map[string]string) (response AppWidgetsGetGroupImagesResponse, vkErr Error) {
	rawResponse, vkErr := vk.Request("appWidgets.getGroupImages", params)
	if vkErr.Code != 0 {
		return
	}

	err := json.Unmarshal(rawResponse, &response)
	if err != nil {
		panic(err)
	}

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
// https://vk.com/dev/appWidgets.getImagesById
func (vk VK) AppWidgetsGetImagesByID(params map[string]string) (response AppWidgetsGetImagesByIDResponse, vkErr Error) {
	rawResponse, vkErr := vk.Request("appWidgets.getImagesById", params)
	if vkErr.Code != 0 {
		return
	}

	err := json.Unmarshal(rawResponse, &response)
	if err != nil {
		panic(err)
	}

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
// https://vk.com/dev/appWidgets.saveAppImage
func (vk VK) AppWidgetsSaveAppImage(params map[string]string) (response AppWidgetsSaveAppImageResponse, vkErr Error) {
	rawResponse, vkErr := vk.Request("appWidgets.saveAppImage", params)
	if vkErr.Code != 0 {
		return
	}

	err := json.Unmarshal(rawResponse, &response)
	if err != nil {
		panic(err)
	}

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
// https://vk.com/dev/appWidgets.saveGroupImage
func (vk VK) AppWidgetsSaveGroupImage(params map[string]string) (response AppWidgetsSaveGroupImageResponse, vkErr Error) {
	rawResponse, vkErr := vk.Request("appWidgets.saveGroupImage", params)
	if vkErr.Code != 0 {
		return
	}

	err := json.Unmarshal(rawResponse, &response)
	if err != nil {
		panic(err)
	}

	return
}

// AppWidgetsUpdate allows to update community app widget.
// https://vk.com/dev/appWidgets.update
func (vk VK) AppWidgetsUpdate(params map[string]string) (vkErr Error) {
	_, vkErr = vk.Request("appWidgets.update", params)

	return
}
