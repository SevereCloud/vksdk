package params // import "github.com/SevereCloud/vksdk/api/params"

import (
	"github.com/SevereCloud/vksdk/api"
)

// NotesAddBuilder builder
//
// Creates a new note for the current user.
//
// https://vk.com/dev/notes.add
type NotesAddBuilder struct {
	api.Params
}

// NewNotesAddBuilder func
func NewNotesAddBuilder() *NotesAddBuilder {
	return &NotesAddBuilder{api.Params{}}
}

// Title Note title.
func (b *NotesAddBuilder) Title(v string) {
	b.Params["title"] = v
}

// Text Note text.
func (b *NotesAddBuilder) Text(v string) {
	b.Params["text"] = v
}

// PrivacyView parameter
func (b *NotesAddBuilder) PrivacyView(v []string) {
	b.Params["privacy_view"] = v
}

// PrivacyComment parameter
func (b *NotesAddBuilder) PrivacyComment(v []string) {
	b.Params["privacy_comment"] = v
}

// NotesCreateCommentBuilder builder
//
// Adds a new comment on a note.
//
// https://vk.com/dev/notes.createComment
type NotesCreateCommentBuilder struct {
	api.Params
}

// NewNotesCreateCommentBuilder func
func NewNotesCreateCommentBuilder() *NotesCreateCommentBuilder {
	return &NotesCreateCommentBuilder{api.Params{}}
}

// NoteID Note ID.
func (b *NotesCreateCommentBuilder) NoteID(v int) {
	b.Params["note_id"] = v
}

// OwnerID Note owner ID.
func (b *NotesCreateCommentBuilder) OwnerID(v int) {
	b.Params["owner_id"] = v
}

// ReplyTo ID of the user to whom the reply is addressed (if the comment is a reply to another comment).
func (b *NotesCreateCommentBuilder) ReplyTo(v int) {
	b.Params["reply_to"] = v
}

// Message Comment text.
func (b *NotesCreateCommentBuilder) Message(v string) {
	b.Params["message"] = v
}

// GUID parameter
func (b *NotesCreateCommentBuilder) GUID(v string) {
	b.Params["guid"] = v
}

// NotesDeleteBuilder builder
//
// Deletes a note of the current user.
//
// https://vk.com/dev/notes.delete
type NotesDeleteBuilder struct {
	api.Params
}

// NewNotesDeleteBuilder func
func NewNotesDeleteBuilder() *NotesDeleteBuilder {
	return &NotesDeleteBuilder{api.Params{}}
}

// NoteID Note ID.
func (b *NotesDeleteBuilder) NoteID(v int) {
	b.Params["note_id"] = v
}

// NotesDeleteCommentBuilder builder
//
// Deletes a comment on a note.
//
// https://vk.com/dev/notes.deleteComment
type NotesDeleteCommentBuilder struct {
	api.Params
}

// NewNotesDeleteCommentBuilder func
func NewNotesDeleteCommentBuilder() *NotesDeleteCommentBuilder {
	return &NotesDeleteCommentBuilder{api.Params{}}
}

// CommentID Comment ID.
func (b *NotesDeleteCommentBuilder) CommentID(v int) {
	b.Params["comment_id"] = v
}

// OwnerID Note owner ID.
func (b *NotesDeleteCommentBuilder) OwnerID(v int) {
	b.Params["owner_id"] = v
}

// NotesEditBuilder builder
//
// Edits a note of the current user.
//
// https://vk.com/dev/notes.edit
type NotesEditBuilder struct {
	api.Params
}

// NewNotesEditBuilder func
func NewNotesEditBuilder() *NotesEditBuilder {
	return &NotesEditBuilder{api.Params{}}
}

// NoteID Note ID.
func (b *NotesEditBuilder) NoteID(v int) {
	b.Params["note_id"] = v
}

// Title Note title.
func (b *NotesEditBuilder) Title(v string) {
	b.Params["title"] = v
}

// Text Note text.
func (b *NotesEditBuilder) Text(v string) {
	b.Params["text"] = v
}

// PrivacyView parameter
func (b *NotesEditBuilder) PrivacyView(v []string) {
	b.Params["privacy_view"] = v
}

// PrivacyComment parameter
func (b *NotesEditBuilder) PrivacyComment(v []string) {
	b.Params["privacy_comment"] = v
}

// NotesEditCommentBuilder builder
//
// Edits a comment on a note.
//
// https://vk.com/dev/notes.editComment
type NotesEditCommentBuilder struct {
	api.Params
}

// NewNotesEditCommentBuilder func
func NewNotesEditCommentBuilder() *NotesEditCommentBuilder {
	return &NotesEditCommentBuilder{api.Params{}}
}

// CommentID Comment ID.
func (b *NotesEditCommentBuilder) CommentID(v int) {
	b.Params["comment_id"] = v
}

// OwnerID Note owner ID.
func (b *NotesEditCommentBuilder) OwnerID(v int) {
	b.Params["owner_id"] = v
}

// Message New comment text.
func (b *NotesEditCommentBuilder) Message(v string) {
	b.Params["message"] = v
}

// NotesGetBuilder builder
//
// Returns a list of notes created by a user.
//
// https://vk.com/dev/notes.get
type NotesGetBuilder struct {
	api.Params
}

// NewNotesGetBuilder func
func NewNotesGetBuilder() *NotesGetBuilder {
	return &NotesGetBuilder{api.Params{}}
}

// NoteIDs Note IDs.
func (b *NotesGetBuilder) NoteIDs(v []int) {
	b.Params["note_ids"] = v
}

// UserID Note owner ID.
func (b *NotesGetBuilder) UserID(v int) {
	b.Params["user_id"] = v
}

// Offset parameter
func (b *NotesGetBuilder) Offset(v int) {
	b.Params["offset"] = v
}

// Count Number of notes to return.
func (b *NotesGetBuilder) Count(v int) {
	b.Params["count"] = v
}

// Sort parameter
func (b *NotesGetBuilder) Sort(v int) {
	b.Params["sort"] = v
}

// NotesGetByIDBuilder builder
//
// Returns a note by its ID.
//
// https://vk.com/dev/notes.getById
type NotesGetByIDBuilder struct {
	api.Params
}

// NewNotesGetByIDBuilder func
func NewNotesGetByIDBuilder() *NotesGetByIDBuilder {
	return &NotesGetByIDBuilder{api.Params{}}
}

// NoteID Note ID.
func (b *NotesGetByIDBuilder) NoteID(v int) {
	b.Params["note_id"] = v
}

// OwnerID Note owner ID.
func (b *NotesGetByIDBuilder) OwnerID(v int) {
	b.Params["owner_id"] = v
}

// NeedWiki parameter
func (b *NotesGetByIDBuilder) NeedWiki(v bool) {
	b.Params["need_wiki"] = v
}

// NotesGetCommentsBuilder builder
//
// Returns a list of comments on a note.
//
// https://vk.com/dev/notes.getComments
type NotesGetCommentsBuilder struct {
	api.Params
}

// NewNotesGetCommentsBuilder func
func NewNotesGetCommentsBuilder() *NotesGetCommentsBuilder {
	return &NotesGetCommentsBuilder{api.Params{}}
}

// NoteID Note ID.
func (b *NotesGetCommentsBuilder) NoteID(v int) {
	b.Params["note_id"] = v
}

// OwnerID Note owner ID.
func (b *NotesGetCommentsBuilder) OwnerID(v int) {
	b.Params["owner_id"] = v
}

// Sort parameter
func (b *NotesGetCommentsBuilder) Sort(v int) {
	b.Params["sort"] = v
}

// Offset parameter
func (b *NotesGetCommentsBuilder) Offset(v int) {
	b.Params["offset"] = v
}

// Count Number of comments to return.
func (b *NotesGetCommentsBuilder) Count(v int) {
	b.Params["count"] = v
}

// NotesRestoreCommentBuilder builder
//
// Restores a deleted comment on a note.
//
// https://vk.com/dev/notes.restoreComment
type NotesRestoreCommentBuilder struct {
	api.Params
}

// NewNotesRestoreCommentBuilder func
func NewNotesRestoreCommentBuilder() *NotesRestoreCommentBuilder {
	return &NotesRestoreCommentBuilder{api.Params{}}
}

// CommentID Comment ID.
func (b *NotesRestoreCommentBuilder) CommentID(v int) {
	b.Params["comment_id"] = v
}

// OwnerID Note owner ID.
func (b *NotesRestoreCommentBuilder) OwnerID(v int) {
	b.Params["owner_id"] = v
}
