package api // import "github.com/SevereCloud/vksdk/api"

import (
	"github.com/SevereCloud/vksdk/object"
)

// DocsAdd copies a document to a user's or community's document list.
//
// https://vk.com/dev/docs.add
func (vk *VK) DocsAdd(params Params) (response int, err error) {
	err = vk.RequestUnmarshal("docs.add", params, &response)
	return
}

// DocsDelete deletes a user or community document.
//
// https://vk.com/dev/docs.delete
func (vk *VK) DocsDelete(params Params) (response int, err error) {
	err = vk.RequestUnmarshal("docs.delete", params, &response)
	return
}

// DocsEdit edits a document.
//
// https://vk.com/dev/docs.edit
func (vk *VK) DocsEdit(params Params) (response int, err error) {
	err = vk.RequestUnmarshal("docs.edit", params, &response)
	return
}

// DocsGetResponse struct.
type DocsGetResponse struct {
	Count int              `json:"count"`
	Items []object.DocsDoc `json:"items"`
}

// DocsGet returns detailed information about user or community documents.
//
// https://vk.com/dev/docs.get
func (vk *VK) DocsGet(params Params) (response DocsGetResponse, err error) {
	err = vk.RequestUnmarshal("docs.get", params, &response)
	return
}

// DocsGetByIDResponse struct.
type DocsGetByIDResponse []object.DocsDoc

// DocsGetByID returns information about documents by their IDs.
//
// https://vk.com/dev/docs.getById
func (vk *VK) DocsGetByID(params Params) (response DocsGetByIDResponse, err error) {
	err = vk.RequestUnmarshal("docs.getById", params, &response)
	return
}

// DocsGetMessagesUploadServerResponse struct.
type DocsGetMessagesUploadServerResponse struct {
	UploadURL string `json:"upload_url"`
}

// DocsGetMessagesUploadServer returns the server address for document upload.
//
// https://vk.com/dev/docs.getMessagesUploadServer
func (vk *VK) DocsGetMessagesUploadServer(params Params) (response DocsGetMessagesUploadServerResponse, err error) {
	err = vk.RequestUnmarshal("docs.getMessagesUploadServer", params, &response)
	return
}

// DocsGetTypesResponse struct.
type DocsGetTypesResponse struct {
	Count int                   `json:"count"`
	Items []object.DocsDocTypes `json:"items"`
}

// DocsGetTypes returns documents types available for current user.
//
// https://vk.com/dev/docs.getTypes
func (vk *VK) DocsGetTypes(params Params) (response DocsGetTypesResponse, err error) {
	err = vk.RequestUnmarshal("docs.getTypes", params, &response)
	return
}

// DocsGetUploadServerResponse struct.
type DocsGetUploadServerResponse struct {
	UploadURL string `json:"upload_url"`
}

// DocsGetUploadServer returns the server address for document upload.
//
// https://vk.com/dev/docs.getUploadServer
func (vk *VK) DocsGetUploadServer(params Params) (response DocsGetUploadServerResponse, err error) {
	err = vk.RequestUnmarshal("docs.getUploadServer", params, &response)
	return
}

// DocsGetWallUploadServerResponse struct.
type DocsGetWallUploadServerResponse struct {
	UploadURL string `json:"upload_url"`
}

// DocsGetWallUploadServer returns the server address for document upload onto a user's or community's wall.
//
// https://vk.com/dev/docs.getWallUploadServer
func (vk *VK) DocsGetWallUploadServer(params Params) (response DocsGetWallUploadServerResponse, err error) {
	err = vk.RequestUnmarshal("docs.getWallUploadServer", params, &response)
	return
}

// DocsSaveResponse struct.
type DocsSaveResponse struct {
	Type         string                      `json:"string"`
	AudioMessage object.MessagesAudioMessage `json:"audio_message"`
	Doc          object.DocsDoc              `json:"doc"`
	Graffiti     object.MessagesGraffiti     `json:"graffiti"`
}

// DocsSave saves a document after uploading it to a server.
//
// https://vk.com/dev/docs.save
func (vk *VK) DocsSave(params Params) (response DocsSaveResponse, err error) {
	err = vk.RequestUnmarshal("docs.save", params, &response)
	return
}

// DocsSearchResponse struct.
type DocsSearchResponse struct {
	Count int              `json:"count"`
	Items []object.DocsDoc `json:"items"`
}

// DocsSearch returns a list of documents matching the search criteria.
//
// https://vk.com/dev/docs.search
func (vk *VK) DocsSearch(params Params) (response DocsSearchResponse, err error) {
	err = vk.RequestUnmarshal("docs.search", params, &response)
	return
}
