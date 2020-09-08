package params // import "github.com/SevereCloud/vksdk/v2/api/params"

import (
	"github.com/SevereCloud/vksdk/v2/api"
)

// DocsAddBuilder builder.
//
// Copies a document to a user's or community's document list.
//
// https://vk.com/dev/docs.add
type DocsAddBuilder struct {
	api.Params
}

// NewDocsAddBuilder func.
func NewDocsAddBuilder() *DocsAddBuilder {
	return &DocsAddBuilder{api.Params{}}
}

// OwnerID ID of the user or community that owns the document. Use a negative value to designate a community ID.
func (b *DocsAddBuilder) OwnerID(v int) *DocsAddBuilder {
	b.Params["owner_id"] = v
	return b
}

// DocID document ID.
func (b *DocsAddBuilder) DocID(v int) *DocsAddBuilder {
	b.Params["doc_id"] = v
	return b
}

// AccessKey access key. This parameter is required if 'access_key' was returned with the document's data.
func (b *DocsAddBuilder) AccessKey(v string) *DocsAddBuilder {
	b.Params["access_key"] = v
	return b
}

// DocsDeleteBuilder builder.
//
// Deletes a user or community document.
//
// https://vk.com/dev/docs.delete
type DocsDeleteBuilder struct {
	api.Params
}

// NewDocsDeleteBuilder func.
func NewDocsDeleteBuilder() *DocsDeleteBuilder {
	return &DocsDeleteBuilder{api.Params{}}
}

// OwnerID ID of the user or community that owns the document. Use a negative value to designate a community ID.
func (b *DocsDeleteBuilder) OwnerID(v int) *DocsDeleteBuilder {
	b.Params["owner_id"] = v
	return b
}

// DocID document ID.
func (b *DocsDeleteBuilder) DocID(v int) *DocsDeleteBuilder {
	b.Params["doc_id"] = v
	return b
}

// DocsEditBuilder builder.
//
// Edits a document.
//
// https://vk.com/dev/docs.edit
type DocsEditBuilder struct {
	api.Params
}

// NewDocsEditBuilder func.
func NewDocsEditBuilder() *DocsEditBuilder {
	return &DocsEditBuilder{api.Params{}}
}

// OwnerID user ID or community ID. Use a negative value to designate a community ID.
func (b *DocsEditBuilder) OwnerID(v int) *DocsEditBuilder {
	b.Params["owner_id"] = v
	return b
}

// DocID document ID.
func (b *DocsEditBuilder) DocID(v int) *DocsEditBuilder {
	b.Params["doc_id"] = v
	return b
}

// Title document.
func (b *DocsEditBuilder) Title(v string) *DocsEditBuilder {
	b.Params["title"] = v
	return b
}

// Tags document.
func (b *DocsEditBuilder) Tags(v []string) *DocsEditBuilder {
	b.Params["tags"] = v
	return b
}

// DocsGetBuilder builder.
//
// Returns detailed information about user or community documents.
//
// https://vk.com/dev/docs.get
type DocsGetBuilder struct {
	api.Params
}

// NewDocsGetBuilder func.
func NewDocsGetBuilder() *DocsGetBuilder {
	return &DocsGetBuilder{api.Params{}}
}

// Count number of documents to return. By default, all documents.
func (b *DocsGetBuilder) Count(v int) *DocsGetBuilder {
	b.Params["count"] = v
	return b
}

// Offset needed to return a specific subset of documents.
func (b *DocsGetBuilder) Offset(v int) *DocsGetBuilder {
	b.Params["offset"] = v
	return b
}

// Type parameter.
func (b *DocsGetBuilder) Type(v int) *DocsGetBuilder {
	b.Params["type"] = v
	return b
}

// OwnerID ID of the user or community that owns the documents. Use a negative value to designate a community ID.
func (b *DocsGetBuilder) OwnerID(v int) *DocsGetBuilder {
	b.Params["owner_id"] = v
	return b
}

// DocsGetByIDBuilder builder.
//
// Returns information about documents by their IDs.
//
// https://vk.com/dev/docs.getById
type DocsGetByIDBuilder struct {
	api.Params
}

// NewDocsGetByIDBuilder func.
func NewDocsGetByIDBuilder() *DocsGetByIDBuilder {
	return &DocsGetByIDBuilder{api.Params{}}
}

// Docs IDs. Example: , "66748_91488,66748_91455".
func (b *DocsGetByIDBuilder) Docs(v []string) *DocsGetByIDBuilder {
	b.Params["docs"] = v
	return b
}

// DocsGetMessagesUploadServerBuilder builder.
//
// Returns the server address for document upload.
//
// https://vk.com/dev/docs.getMessagesUploadServer
type DocsGetMessagesUploadServerBuilder struct {
	api.Params
}

// NewDocsGetMessagesUploadServerBuilder func.
func NewDocsGetMessagesUploadServerBuilder() *DocsGetMessagesUploadServerBuilder {
	return &DocsGetMessagesUploadServerBuilder{api.Params{}}
}

// Type document.
func (b *DocsGetMessagesUploadServerBuilder) Type(v string) *DocsGetMessagesUploadServerBuilder {
	b.Params["type"] = v
	return b
}

// PeerID destination ID. For user: 'User ID', e.g. '12345'.
// For chat: '2000000000' + 'Chat ID', e.g. '2000000001'.
// For community: '- community ID', e.g. '-12345'.
func (b *DocsGetMessagesUploadServerBuilder) PeerID(v int) *DocsGetMessagesUploadServerBuilder {
	b.Params["peer_id"] = v
	return b
}

// DocsGetTypesBuilder builder.
//
// Returns documents types available for current user.
//
// https://vk.com/dev/docs.getTypes
type DocsGetTypesBuilder struct {
	api.Params
}

// NewDocsGetTypesBuilder func.
func NewDocsGetTypesBuilder() *DocsGetTypesBuilder {
	return &DocsGetTypesBuilder{api.Params{}}
}

// OwnerID ID of the user or community that owns the documents. Use a negative value to designate a community ID.
func (b *DocsGetTypesBuilder) OwnerID(v int) *DocsGetTypesBuilder {
	b.Params["owner_id"] = v
	return b
}

// DocsGetUploadServerBuilder builder.
//
// Returns the server address for document upload.
//
// https://vk.com/dev/docs.getUploadServer
type DocsGetUploadServerBuilder struct {
	api.Params
}

// NewDocsGetUploadServerBuilder func.
func NewDocsGetUploadServerBuilder() *DocsGetUploadServerBuilder {
	return &DocsGetUploadServerBuilder{api.Params{}}
}

// GroupID community ID (if the document will be uploaded to the community).
func (b *DocsGetUploadServerBuilder) GroupID(v int) *DocsGetUploadServerBuilder {
	b.Params["group_id"] = v
	return b
}

// DocsGetWallUploadServerBuilder builder.
//
// Returns the server address for document upload onto a user's or community's wall.
//
// https://vk.com/dev/docs.getWallUploadServer
type DocsGetWallUploadServerBuilder struct {
	api.Params
}

// NewDocsGetWallUploadServerBuilder func.
func NewDocsGetWallUploadServerBuilder() *DocsGetWallUploadServerBuilder {
	return &DocsGetWallUploadServerBuilder{api.Params{}}
}

// GroupID community ID (if the document will be uploaded to the community).
func (b *DocsGetWallUploadServerBuilder) GroupID(v int) *DocsGetWallUploadServerBuilder {
	b.Params["group_id"] = v
	return b
}

// DocsSaveBuilder builder.
//
// Saves a document after [vk.com/dev/upload_files_2|uploading it to a server].
//
// https://vk.com/dev/docs.save
type DocsSaveBuilder struct {
	api.Params
}

// NewDocsSaveBuilder func.
func NewDocsSaveBuilder() *DocsSaveBuilder {
	return &DocsSaveBuilder{api.Params{}}
}

// File This parameter is returned when the file is [vk.com/dev/upload_files_2|uploaded to the server].
func (b *DocsSaveBuilder) File(v string) *DocsSaveBuilder {
	b.Params["file"] = v
	return b
}

// Title document.
func (b *DocsSaveBuilder) Title(v string) *DocsSaveBuilder {
	b.Params["title"] = v
	return b
}

// Tags document.
func (b *DocsSaveBuilder) Tags(v string) *DocsSaveBuilder {
	b.Params["tags"] = v
	return b
}

// DocsSearchBuilder builder.
//
// Returns a list of documents matching the search criteria.
//
// https://vk.com/dev/docs.search
type DocsSearchBuilder struct {
	api.Params
}

// NewDocsSearchBuilder func.
func NewDocsSearchBuilder() *DocsSearchBuilder {
	return &DocsSearchBuilder{api.Params{}}
}

// Q search query string.
func (b *DocsSearchBuilder) Q(v string) *DocsSearchBuilder {
	b.Params["q"] = v
	return b
}

// SearchOwn parameter.
func (b *DocsSearchBuilder) SearchOwn(v bool) *DocsSearchBuilder {
	b.Params["search_own"] = v
	return b
}

// Count number of results to return.
func (b *DocsSearchBuilder) Count(v int) *DocsSearchBuilder {
	b.Params["count"] = v
	return b
}

// Offset needed to return a specific subset of results.
func (b *DocsSearchBuilder) Offset(v int) *DocsSearchBuilder {
	b.Params["offset"] = v
	return b
}
