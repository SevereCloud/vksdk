package params // import "github.com/SevereCloud/vksdk/api/params"

import (
	"github.com/SevereCloud/vksdk/api"
)

// NotesAddBulder builder
//
// Creates a new note for the current user.
//
// https://vk.com/dev/notes.add
type NotesAddBulder struct {
	api.Params
}

// NewNotesAddBulder func
func NewNotesAddBulder() *NotesAddBulder {
	return &NotesAddBulder{api.Params{}}
}

// Title Note title.
func (b *NotesAddBulder) Title(v string) {
	b.Params["title"] = v
}

// Text Note text.
func (b *NotesAddBulder) Text(v string) {
	b.Params["text"] = v
}

// PrivacyView parameter
func (b *NotesAddBulder) PrivacyView(v []string) {
	b.Params["privacy_view"] = v
}

// PrivacyComment parameter
func (b *NotesAddBulder) PrivacyComment(v []string) {
	b.Params["privacy_comment"] = v
}

// NotesCreateCommentBulder builder
//
// Adds a new comment on a note.
//
// https://vk.com/dev/notes.createComment
type NotesCreateCommentBulder struct {
	api.Params
}

// NewNotesCreateCommentBulder func
func NewNotesCreateCommentBulder() *NotesCreateCommentBulder {
	return &NotesCreateCommentBulder{api.Params{}}
}

// NoteID Note ID.
func (b *NotesCreateCommentBulder) NoteID(v int) {
	b.Params["note_id"] = v
}

// OwnerID Note owner ID.
func (b *NotesCreateCommentBulder) OwnerID(v int) {
	b.Params["owner_id"] = v
}

// ReplyTo ID of the user to whom the reply is addressed (if the comment is a reply to another comment).
func (b *NotesCreateCommentBulder) ReplyTo(v int) {
	b.Params["reply_to"] = v
}

// Message Comment text.
func (b *NotesCreateCommentBulder) Message(v string) {
	b.Params["message"] = v
}

// GUID parameter
func (b *NotesCreateCommentBulder) GUID(v string) {
	b.Params["guid"] = v
}

// NotesDeleteBulder builder
//
// Deletes a note of the current user.
//
// https://vk.com/dev/notes.delete
type NotesDeleteBulder struct {
	api.Params
}

// NewNotesDeleteBulder func
func NewNotesDeleteBulder() *NotesDeleteBulder {
	return &NotesDeleteBulder{api.Params{}}
}

// NoteID Note ID.
func (b *NotesDeleteBulder) NoteID(v int) {
	b.Params["note_id"] = v
}

// NotesDeleteCommentBulder builder
//
// Deletes a comment on a note.
//
// https://vk.com/dev/notes.deleteComment
type NotesDeleteCommentBulder struct {
	api.Params
}

// NewNotesDeleteCommentBulder func
func NewNotesDeleteCommentBulder() *NotesDeleteCommentBulder {
	return &NotesDeleteCommentBulder{api.Params{}}
}

// CommentID Comment ID.
func (b *NotesDeleteCommentBulder) CommentID(v int) {
	b.Params["comment_id"] = v
}

// OwnerID Note owner ID.
func (b *NotesDeleteCommentBulder) OwnerID(v int) {
	b.Params["owner_id"] = v
}

// NotesEditBulder builder
//
// Edits a note of the current user.
//
// https://vk.com/dev/notes.edit
type NotesEditBulder struct {
	api.Params
}

// NewNotesEditBulder func
func NewNotesEditBulder() *NotesEditBulder {
	return &NotesEditBulder{api.Params{}}
}

// NoteID Note ID.
func (b *NotesEditBulder) NoteID(v int) {
	b.Params["note_id"] = v
}

// Title Note title.
func (b *NotesEditBulder) Title(v string) {
	b.Params["title"] = v
}

// Text Note text.
func (b *NotesEditBulder) Text(v string) {
	b.Params["text"] = v
}

// PrivacyView parameter
func (b *NotesEditBulder) PrivacyView(v []string) {
	b.Params["privacy_view"] = v
}

// PrivacyComment parameter
func (b *NotesEditBulder) PrivacyComment(v []string) {
	b.Params["privacy_comment"] = v
}

// NotesEditCommentBulder builder
//
// Edits a comment on a note.
//
// https://vk.com/dev/notes.editComment
type NotesEditCommentBulder struct {
	api.Params
}

// NewNotesEditCommentBulder func
func NewNotesEditCommentBulder() *NotesEditCommentBulder {
	return &NotesEditCommentBulder{api.Params{}}
}

// CommentID Comment ID.
func (b *NotesEditCommentBulder) CommentID(v int) {
	b.Params["comment_id"] = v
}

// OwnerID Note owner ID.
func (b *NotesEditCommentBulder) OwnerID(v int) {
	b.Params["owner_id"] = v
}

// Message New comment text.
func (b *NotesEditCommentBulder) Message(v string) {
	b.Params["message"] = v
}

// NotesGetBulder builder
//
// Returns a list of notes created by a user.
//
// https://vk.com/dev/notes.get
type NotesGetBulder struct {
	api.Params
}

// NewNotesGetBulder func
func NewNotesGetBulder() *NotesGetBulder {
	return &NotesGetBulder{api.Params{}}
}

// NoteIDs Note IDs.
func (b *NotesGetBulder) NoteIDs(v []int) {
	b.Params["note_ids"] = v
}

// UserID Note owner ID.
func (b *NotesGetBulder) UserID(v int) {
	b.Params["user_id"] = v
}

// Offset parameter
func (b *NotesGetBulder) Offset(v int) {
	b.Params["offset"] = v
}

// Count Number of notes to return.
func (b *NotesGetBulder) Count(v int) {
	b.Params["count"] = v
}

// Sort parameter
func (b *NotesGetBulder) Sort(v int) {
	b.Params["sort"] = v
}

// NotesGetByIDBulder builder
//
// Returns a note by its ID.
//
// https://vk.com/dev/notes.getById
type NotesGetByIDBulder struct {
	api.Params
}

// NewNotesGetByIDBulder func
func NewNotesGetByIDBulder() *NotesGetByIDBulder {
	return &NotesGetByIDBulder{api.Params{}}
}

// NoteID Note ID.
func (b *NotesGetByIDBulder) NoteID(v int) {
	b.Params["note_id"] = v
}

// OwnerID Note owner ID.
func (b *NotesGetByIDBulder) OwnerID(v int) {
	b.Params["owner_id"] = v
}

// NeedWiki parameter
func (b *NotesGetByIDBulder) NeedWiki(v bool) {
	b.Params["need_wiki"] = v
}

// NotesGetCommentsBulder builder
//
// Returns a list of comments on a note.
//
// https://vk.com/dev/notes.getComments
type NotesGetCommentsBulder struct {
	api.Params
}

// NewNotesGetCommentsBulder func
func NewNotesGetCommentsBulder() *NotesGetCommentsBulder {
	return &NotesGetCommentsBulder{api.Params{}}
}

// NoteID Note ID.
func (b *NotesGetCommentsBulder) NoteID(v int) {
	b.Params["note_id"] = v
}

// OwnerID Note owner ID.
func (b *NotesGetCommentsBulder) OwnerID(v int) {
	b.Params["owner_id"] = v
}

// Sort parameter
func (b *NotesGetCommentsBulder) Sort(v int) {
	b.Params["sort"] = v
}

// Offset parameter
func (b *NotesGetCommentsBulder) Offset(v int) {
	b.Params["offset"] = v
}

// Count Number of comments to return.
func (b *NotesGetCommentsBulder) Count(v int) {
	b.Params["count"] = v
}

// NotesRestoreCommentBulder builder
//
// Restores a deleted comment on a note.
//
// https://vk.com/dev/notes.restoreComment
type NotesRestoreCommentBulder struct {
	api.Params
}

// NewNotesRestoreCommentBulder func
func NewNotesRestoreCommentBulder() *NotesRestoreCommentBulder {
	return &NotesRestoreCommentBulder{api.Params{}}
}

// CommentID Comment ID.
func (b *NotesRestoreCommentBulder) CommentID(v int) {
	b.Params["comment_id"] = v
}

// OwnerID Note owner ID.
func (b *NotesRestoreCommentBulder) OwnerID(v int) {
	b.Params["owner_id"] = v
}
