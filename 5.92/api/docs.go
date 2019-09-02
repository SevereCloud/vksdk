package api // import "github.com/SevereCloud/vksdk/5.92/api"

import (
	"github.com/SevereCloud/vksdk/5.92/object"
)

// DocsAdd copies a document to a user's or community's document list.
//
// https://vk.com/dev/docs.add
func (vk *VK) DocsAdd(params map[string]string) (response int, vkErr Error) {
	vk.RequestUnmarshal("docs.add", params, &response, &vkErr)
	return
}

// DocsDelete deletes a user or community document.
//
// https://vk.com/dev/docs.delete
func (vk *VK) DocsDelete(params map[string]string) (response int, vkErr Error) {
	vk.RequestUnmarshal("docs.delete", params, &response, &vkErr)
	return
}

// DocsEdit edits a document.
//
// https://vk.com/dev/docs.edit
func (vk *VK) DocsEdit(params map[string]string) (response int, vkErr Error) {
	vk.RequestUnmarshal("docs.edit", params, &response, &vkErr)
	return
}

// DocsGetResponse struct
type DocsGetResponse struct {
	Count int              `json:"count"`
	Items []object.DocsDoc `json:"items"`
}

// DocsGet returns detailed information about user or community documents.
//
// https://vk.com/dev/docs.get
func (vk *VK) DocsGet(params map[string]string) (response DocsGetResponse, vkErr Error) {
	vk.RequestUnmarshal("docs.get", params, &response, &vkErr)
	return
}

// DocsGetByIDResponse struct
type DocsGetByIDResponse []object.DocsDoc

// DocsGetByID returns information about documents by their IDs.
//
// https://vk.com/dev/docs.getById
func (vk *VK) DocsGetByID(params map[string]string) (response DocsGetByIDResponse, vkErr Error) {
	vk.RequestUnmarshal("docs.getById", params, &response, &vkErr)
	return
}

// DocsGetMessagesUploadServerResponse struct
type DocsGetMessagesUploadServerResponse struct {
	UploadURL string `json:"upload_url"`
}

// DocsGetMessagesUploadServer returns the server address for document upload.
//
// https://vk.com/dev/docs.getMessagesUploadServer
func (vk *VK) DocsGetMessagesUploadServer(params map[string]string) (response DocsGetMessagesUploadServerResponse, vkErr Error) {
	vk.RequestUnmarshal("docs.getMessagesUploadServer", params, &response, &vkErr)
	return
}

// DocsGetTypesResponse struct
type DocsGetTypesResponse struct {
	Count int                   `json:"count"`
	Items []object.DocsDocTypes `json:"items"`
}

// DocsGetTypes returns documents types available for current user.
//
// https://vk.com/dev/docs.getTypes
func (vk *VK) DocsGetTypes(params map[string]string) (response DocsGetTypesResponse, vkErr Error) {
	vk.RequestUnmarshal("docs.getTypes", params, &response, &vkErr)
	return
}

// DocsGetUploadServerResponse struct
type DocsGetUploadServerResponse struct {
	UploadURL string `json:"upload_url"`
}

// DocsGetUploadServer returns the server address for document upload.
//
// https://vk.com/dev/docs.getUploadServer
func (vk *VK) DocsGetUploadServer(params map[string]string) (response DocsGetUploadServerResponse, vkErr Error) {
	vk.RequestUnmarshal("docs.getUploadServer", params, &response, &vkErr)
	return
}

// DocsGetWallUploadServerResponse struct
type DocsGetWallUploadServerResponse struct {
	UploadURL string `json:"upload_url"`
}

// DocsGetWallUploadServer returns the server address for document upload onto a user's or community's wall.
//
// https://vk.com/dev/docs.getWallUploadServer
func (vk *VK) DocsGetWallUploadServer(params map[string]string) (response DocsGetWallUploadServerResponse, vkErr Error) {
	vk.RequestUnmarshal("docs.getWallUploadServer", params, &response, &vkErr)
	return
}

// DocsSaveResponse struct
type DocsSaveResponse struct {
	Type         string                      `json:"string"`
	AudioMessage object.MessagesAudioMessage `json:"audio_message"`
	Doc          object.DocsDoc              `json:"doc"`
	Graffiti     object.MessagesGraffiti     `json:"graffiti"`
}

// DocsSave saves a document after uploading it to a server.
//
// https://vk.com/dev/docs.save
func (vk *VK) DocsSave(params map[string]string) (response DocsSaveResponse, vkErr Error) {
	vk.RequestUnmarshal("docs.save", params, &response, &vkErr)
	return
}

// DocsSearchResponse struct
type DocsSearchResponse struct {
	Count int              `json:"count"`
	Items []object.DocsDoc `json:"items"`
}

// DocsSearch returns a list of documents matching the search criteria.
//
// https://vk.com/dev/docs.search
func (vk *VK) DocsSearch(params map[string]string) (response DocsSearchResponse, vkErr Error) {
	vk.RequestUnmarshal("docs.search", params, &response, &vkErr)
	return
}
