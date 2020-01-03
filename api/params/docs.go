package params // import "github.com/SevereCloud/vksdk/api/params"

import (
	"github.com/SevereCloud/vksdk/api"
)

// DocsAddBulder builder
//
// Copies a document to a user's or community's document list.
//
// https://vk.com/dev/docs.add
type DocsAddBulder struct {
	api.Params
}

// NewDocsAddBulder func
func NewDocsAddBulder() *DocsAddBulder {
	return &DocsAddBulder{api.Params{}}
}

// OwnerID ID of the user or community that owns the document. Use a negative value to designate a community ID.
func (b *DocsAddBulder) OwnerID(v int) {
	b.Params["owner_id"] = v
}

// DocID Document ID.
func (b *DocsAddBulder) DocID(v int) {
	b.Params["doc_id"] = v
}

// AccessKey Access key. This parameter is required if 'access_key' was returned with the document's data.
func (b *DocsAddBulder) AccessKey(v string) {
	b.Params["access_key"] = v
}

// DocsDeleteBulder builder
//
// Deletes a user or community document.
//
// https://vk.com/dev/docs.delete
type DocsDeleteBulder struct {
	api.Params
}

// NewDocsDeleteBulder func
func NewDocsDeleteBulder() *DocsDeleteBulder {
	return &DocsDeleteBulder{api.Params{}}
}

// OwnerID ID of the user or community that owns the document. Use a negative value to designate a community ID.
func (b *DocsDeleteBulder) OwnerID(v int) {
	b.Params["owner_id"] = v
}

// DocID Document ID.
func (b *DocsDeleteBulder) DocID(v int) {
	b.Params["doc_id"] = v
}

// DocsEditBulder builder
//
// Edits a document.
//
// https://vk.com/dev/docs.edit
type DocsEditBulder struct {
	api.Params
}

// NewDocsEditBulder func
func NewDocsEditBulder() *DocsEditBulder {
	return &DocsEditBulder{api.Params{}}
}

// OwnerID User ID or community ID. Use a negative value to designate a community ID.
func (b *DocsEditBulder) OwnerID(v int) {
	b.Params["owner_id"] = v
}

// DocID Document ID.
func (b *DocsEditBulder) DocID(v int) {
	b.Params["doc_id"] = v
}

// Title Document title.
func (b *DocsEditBulder) Title(v string) {
	b.Params["title"] = v
}

// Tags Document tags.
func (b *DocsEditBulder) Tags(v []string) {
	b.Params["tags"] = v
}

// DocsGetBulder builder
//
// Returns detailed information about user or community documents.
//
// https://vk.com/dev/docs.get
type DocsGetBulder struct {
	api.Params
}

// NewDocsGetBulder func
func NewDocsGetBulder() *DocsGetBulder {
	return &DocsGetBulder{api.Params{}}
}

// Count Number of documents to return. By default, all documents.
func (b *DocsGetBulder) Count(v int) {
	b.Params["count"] = v
}

// Offset Offset needed to return a specific subset of documents.
func (b *DocsGetBulder) Offset(v int) {
	b.Params["offset"] = v
}

// Type parameter
func (b *DocsGetBulder) Type(v int) {
	b.Params["type"] = v
}

// OwnerID ID of the user or community that owns the documents. Use a negative value to designate a community ID.
func (b *DocsGetBulder) OwnerID(v int) {
	b.Params["owner_id"] = v
}

// DocsGetByIDBulder builder
//
// Returns information about documents by their IDs.
//
// https://vk.com/dev/docs.getById
type DocsGetByIDBulder struct {
	api.Params
}

// NewDocsGetByIDBulder func
func NewDocsGetByIDBulder() *DocsGetByIDBulder {
	return &DocsGetByIDBulder{api.Params{}}
}

// Docs Document IDs. Example: , "66748_91488,66748_91455",
func (b *DocsGetByIDBulder) Docs(v []string) {
	b.Params["docs"] = v
}

// DocsGetMessagesUploadServerBulder builder
//
// Returns the server address for document upload.
//
// https://vk.com/dev/docs.getMessagesUploadServer
type DocsGetMessagesUploadServerBulder struct {
	api.Params
}

// NewDocsGetMessagesUploadServerBulder func
func NewDocsGetMessagesUploadServerBulder() *DocsGetMessagesUploadServerBulder {
	return &DocsGetMessagesUploadServerBulder{api.Params{}}
}

// Type Document type.
func (b *DocsGetMessagesUploadServerBulder) Type(v string) {
	b.Params["type"] = v
}

// PeerID Destination ID. "For user: 'User ID', e.g. '12345'. For chat: '2000000000' + 'Chat ID', e.g. '2000000001'. For community: '- Community ID', e.g. '-12345'. "
func (b *DocsGetMessagesUploadServerBulder) PeerID(v int) {
	b.Params["peer_id"] = v
}

// DocsGetTypesBulder builder
//
// Returns documents types available for current user.
//
// https://vk.com/dev/docs.getTypes
type DocsGetTypesBulder struct {
	api.Params
}

// NewDocsGetTypesBulder func
func NewDocsGetTypesBulder() *DocsGetTypesBulder {
	return &DocsGetTypesBulder{api.Params{}}
}

// OwnerID ID of the user or community that owns the documents. Use a negative value to designate a community ID.
func (b *DocsGetTypesBulder) OwnerID(v int) {
	b.Params["owner_id"] = v
}

// DocsGetUploadServerBulder builder
//
// Returns the server address for document upload.
//
// https://vk.com/dev/docs.getUploadServer
type DocsGetUploadServerBulder struct {
	api.Params
}

// NewDocsGetUploadServerBulder func
func NewDocsGetUploadServerBulder() *DocsGetUploadServerBulder {
	return &DocsGetUploadServerBulder{api.Params{}}
}

// GroupID Community ID (if the document will be uploaded to the community).
func (b *DocsGetUploadServerBulder) GroupID(v int) {
	b.Params["group_id"] = v
}

// DocsGetWallUploadServerBulder builder
//
// Returns the server address for document upload onto a user's or community's wall.
//
// https://vk.com/dev/docs.getWallUploadServer
type DocsGetWallUploadServerBulder struct {
	api.Params
}

// NewDocsGetWallUploadServerBulder func
func NewDocsGetWallUploadServerBulder() *DocsGetWallUploadServerBulder {
	return &DocsGetWallUploadServerBulder{api.Params{}}
}

// GroupID Community ID (if the document will be uploaded to the community).
func (b *DocsGetWallUploadServerBulder) GroupID(v int) {
	b.Params["group_id"] = v
}

// DocsSaveBulder builder
//
// Saves a document after [vk.com/dev/upload_files_2|uploading it to a server].
//
// https://vk.com/dev/docs.save
type DocsSaveBulder struct {
	api.Params
}

// NewDocsSaveBulder func
func NewDocsSaveBulder() *DocsSaveBulder {
	return &DocsSaveBulder{api.Params{}}
}

// File This parameter is returned when the file is [vk.com/dev/upload_files_2|uploaded to the server].
func (b *DocsSaveBulder) File(v string) {
	b.Params["file"] = v
}

// Title Document title.
func (b *DocsSaveBulder) Title(v string) {
	b.Params["title"] = v
}

// Tags Document tags.
func (b *DocsSaveBulder) Tags(v string) {
	b.Params["tags"] = v
}

// DocsSearchBulder builder
//
// Returns a list of documents matching the search criteria.
//
// https://vk.com/dev/docs.search
type DocsSearchBulder struct {
	api.Params
}

// NewDocsSearchBulder func
func NewDocsSearchBulder() *DocsSearchBulder {
	return &DocsSearchBulder{api.Params{}}
}

// Q Search query string.
func (b *DocsSearchBulder) Q(v string) {
	b.Params["q"] = v
}

// SearchOwn parameter
func (b *DocsSearchBulder) SearchOwn(v bool) {
	b.Params["search_own"] = v
}

// Count Number of results to return.
func (b *DocsSearchBulder) Count(v int) {
	b.Params["count"] = v
}

// Offset Offset needed to return a specific subset of results.
func (b *DocsSearchBulder) Offset(v int) {
	b.Params["offset"] = v
}
