package api // import "github.com/SevereCloud/vksdk/5.92/api"

import (
	"encoding/json"

	"github.com/SevereCloud/vksdk/5.92/object"
)

// DocsAddResponse struct
type DocsAddResponse int

// DocsAdd copies a document to a user's or community's document list.
// https://vk.com/dev/docs.add
func (vk VK) DocsAdd(params map[string]string) (response DocsAddResponse, vkErr Error) {
	rawResponse, vkErr := vk.Request("docs.add", params)
	if vkErr.Code != 0 {
		return
	}

	err := json.Unmarshal(rawResponse, &response)
	if err != nil {
		panic(err)
	}

	return
}

// DocsDelete deletes a user or community document.
// https://vk.com/dev/docs.delete
func (vk VK) DocsDelete(params map[string]string) (vkErr Error) {
	_, vkErr = vk.Request("docs.delete", params)

	return
}

// DocsEdit edits a document.
// https://vk.com/dev/docs.edit
func (vk VK) DocsEdit(params map[string]string) (vkErr Error) {
	_, vkErr = vk.Request("docs.edit", params)

	return
}

// DocsGetResponse struct
type DocsGetResponse struct {
	Count int              `json:"count"`
	Items []object.DocsDoc `json:"items"`
}

// DocsGet returns detailed information about user or community documents.
// https://vk.com/dev/docs.get
func (vk VK) DocsGet(params map[string]string) (response DocsGetResponse, vkErr Error) {
	rawResponse, vkErr := vk.Request("docs.get", params)
	if vkErr.Code != 0 {
		return
	}

	err := json.Unmarshal(rawResponse, &response)
	if err != nil {
		panic(err)
	}

	return
}

// DocsGetByIDResponse struct
type DocsGetByIDResponse []object.DocsDoc

// DocsGetByID returns information about documents by their IDs.
// https://vk.com/dev/docs.getById
func (vk VK) DocsGetByID(params map[string]string) (response DocsGetByIDResponse, vkErr Error) {
	rawResponse, vkErr := vk.Request("docs.getById", params)
	if vkErr.Code != 0 {
		return
	}

	err := json.Unmarshal(rawResponse, &response)
	if err != nil {
		panic(err)
	}

	return
}

// DocsGetMessagesUploadServerResponse struct
type DocsGetMessagesUploadServerResponse struct {
	UploadURL string `json:"upload_url"`
}

// DocsGetMessagesUploadServer returns the server address for document upload.
// https://vk.com/dev/docs.getMessagesUploadServer
func (vk VK) DocsGetMessagesUploadServer(params map[string]string) (response DocsGetMessagesUploadServerResponse, vkErr Error) {
	rawResponse, vkErr := vk.Request("docs.getMessagesUploadServer", params)
	if vkErr.Code != 0 {
		return
	}

	err := json.Unmarshal(rawResponse, &response)
	if err != nil {
		panic(err)
	}

	return
}

// DocsGetTypesResponse struct
type DocsGetTypesResponse struct {
	Count int                   `json:"count"`
	Items []object.DocsDocTypes `json:"items"`
}

// DocsGetTypes returns documents types available for current user.
// https://vk.com/dev/docs.getTypes
func (vk VK) DocsGetTypes(params map[string]string) (response DocsGetTypesResponse, vkErr Error) {
	rawResponse, vkErr := vk.Request("docs.getTypes", params)
	if vkErr.Code != 0 {
		return
	}

	err := json.Unmarshal(rawResponse, &response)
	if err != nil {
		panic(err)
	}

	return
}

// DocsGetUploadServerResponse struct
type DocsGetUploadServerResponse struct {
	UploadURL string `json:"upload_url"`
}

// DocsGetUploadServer returns the server address for document upload.
// https://vk.com/dev/docs.getUploadServer
func (vk VK) DocsGetUploadServer(params map[string]string) (response DocsGetUploadServerResponse, vkErr Error) {
	rawResponse, vkErr := vk.Request("docs.getUploadServer", params)
	if vkErr.Code != 0 {
		return
	}

	err := json.Unmarshal(rawResponse, &response)
	if err != nil {
		panic(err)
	}

	return
}

// DocsGetWallUploadServerResponse struct
type DocsGetWallUploadServerResponse struct {
	UploadURL string `json:"upload_url"`
}

// DocsGetWallUploadServer returns the server address for document upload onto a user's or community's wall.
// https://vk.com/dev/docs.getWallUploadServer
func (vk VK) DocsGetWallUploadServer(params map[string]string) (response DocsGetWallUploadServerResponse, vkErr Error) {
	rawResponse, vkErr := vk.Request("docs.getWallUploadServer", params)
	if vkErr.Code != 0 {
		return
	}

	err := json.Unmarshal(rawResponse, &response)
	if err != nil {
		panic(err)
	}

	return
}

// DocsSaveResponse struct
type DocsSaveResponse []object.DocsDoc

// DocsSave saves a document after uploading it to a server.
// https://vk.com/dev/docs.save
func (vk VK) DocsSave(params map[string]string) (response DocsSaveResponse, vkErr Error) {
	rawResponse, vkErr := vk.Request("docs.save", params)
	if vkErr.Code != 0 {
		return
	}

	err := json.Unmarshal(rawResponse, &response)
	if err != nil {
		panic(err)
	}

	return
}

// DocsSearchResponse struct
type DocsSearchResponse struct {
	Count int              `json:"count"`
	Items []object.DocsDoc `json:"items"`
}

// DocsSearch returns a list of documents matching the search criteria.
// https://vk.com/dev/docs.search
func (vk VK) DocsSearch(params map[string]string) (response DocsSearchResponse, vkErr Error) {
	rawResponse, vkErr := vk.Request("docs.search", params)
	if vkErr.Code != 0 {
		return
	}

	err := json.Unmarshal(rawResponse, &response)
	if err != nil {
		panic(err)
	}

	return
}
