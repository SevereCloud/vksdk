package params // import "github.com/SevereCloud/vksdk/v2/api/params"

import (
	"github.com/SevereCloud/vksdk/v2/api"
)

// NotesAddBuilder builder.
//
// Creates a new note for the current user.
//
// https://vk.com/dev/notes.add
type NotesAddBuilder struct {
	api.Params
}

// NewNotesAddBuilder func.
func NewNotesAddBuilder() *NotesAddBuilder {
	return &NotesAddBuilder{api.Params{}}
}

// Title note title.
func (b *NotesAddBuilder) Title(v string) *NotesAddBuilder {
	b.Params["title"] = v
	return b
}

// Text note text.
func (b *NotesAddBuilder) Text(v string) *NotesAddBuilder {
	b.Params["text"] = v
	return b
}

// PrivacyView parameter.
func (b *NotesAddBuilder) PrivacyView(v []string) *NotesAddBuilder {
	b.Params["privacy_view"] = v
	return b
}

// PrivacyComment parameter.
func (b *NotesAddBuilder) PrivacyComment(v []string) *NotesAddBuilder {
	b.Params["privacy_comment"] = v
	return b
}

// NotesCreateCommentBuilder builder.
//
// Adds a new comment on a note.
//
// https://vk.com/dev/notes.createComment
type NotesCreateCommentBuilder struct {
	api.Params
}

// NewNotesCreateCommentBuilder func.
func NewNotesCreateCommentBuilder() *NotesCreateCommentBuilder {
	return &NotesCreateCommentBuilder{api.Params{}}
}

// NoteID parameter.
func (b *NotesCreateCommentBuilder) NoteID(v int) *NotesCreateCommentBuilder {
	b.Params["note_id"] = v
	return b
}

// OwnerID note owner ID.
func (b *NotesCreateCommentBuilder) OwnerID(v int) *NotesCreateCommentBuilder {
	b.Params["owner_id"] = v
	return b
}

// ReplyTo ID of the user to whom the reply is addressed (if the comment is a reply to another comment).
func (b *NotesCreateCommentBuilder) ReplyTo(v int) *NotesCreateCommentBuilder {
	b.Params["reply_to"] = v
	return b
}

// Message comment text.
func (b *NotesCreateCommentBuilder) Message(v string) *NotesCreateCommentBuilder {
	b.Params["message"] = v
	return b
}

// GUID parameter.
func (b *NotesCreateCommentBuilder) GUID(v string) *NotesCreateCommentBuilder {
	b.Params["guid"] = v
	return b
}

// NotesDeleteBuilder builder.
//
// Deletes a note of the current user.
//
// https://vk.com/dev/notes.delete
type NotesDeleteBuilder struct {
	api.Params
}

// NewNotesDeleteBuilder func.
func NewNotesDeleteBuilder() *NotesDeleteBuilder {
	return &NotesDeleteBuilder{api.Params{}}
}

// NoteID parameter.
func (b *NotesDeleteBuilder) NoteID(v int) *NotesDeleteBuilder {
	b.Params["note_id"] = v
	return b
}

// NotesDeleteCommentBuilder builder.
//
// Deletes a comment on a note.
//
// https://vk.com/dev/notes.deleteComment
type NotesDeleteCommentBuilder struct {
	api.Params
}

// NewNotesDeleteCommentBuilder func.
func NewNotesDeleteCommentBuilder() *NotesDeleteCommentBuilder {
	return &NotesDeleteCommentBuilder{api.Params{}}
}

// CommentID parameter.
func (b *NotesDeleteCommentBuilder) CommentID(v int) *NotesDeleteCommentBuilder {
	b.Params["comment_id"] = v
	return b
}

// OwnerID note owner ID.
func (b *NotesDeleteCommentBuilder) OwnerID(v int) *NotesDeleteCommentBuilder {
	b.Params["owner_id"] = v
	return b
}

// NotesEditBuilder builder.
//
// Edits a note of the current user.
//
// https://vk.com/dev/notes.edit
type NotesEditBuilder struct {
	api.Params
}

// NewNotesEditBuilder func.
func NewNotesEditBuilder() *NotesEditBuilder {
	return &NotesEditBuilder{api.Params{}}
}

// NoteID parameter.
func (b *NotesEditBuilder) NoteID(v int) *NotesEditBuilder {
	b.Params["note_id"] = v
	return b
}

// Title note title.
func (b *NotesEditBuilder) Title(v string) *NotesEditBuilder {
	b.Params["title"] = v
	return b
}

// Text note text.
func (b *NotesEditBuilder) Text(v string) *NotesEditBuilder {
	b.Params["text"] = v
	return b
}

// PrivacyView parameter.
func (b *NotesEditBuilder) PrivacyView(v []string) *NotesEditBuilder {
	b.Params["privacy_view"] = v
	return b
}

// PrivacyComment parameter.
func (b *NotesEditBuilder) PrivacyComment(v []string) *NotesEditBuilder {
	b.Params["privacy_comment"] = v
	return b
}

// NotesEditCommentBuilder builder.
//
// Edits a comment on a note.
//
// https://vk.com/dev/notes.editComment
type NotesEditCommentBuilder struct {
	api.Params
}

// NewNotesEditCommentBuilder func.
func NewNotesEditCommentBuilder() *NotesEditCommentBuilder {
	return &NotesEditCommentBuilder{api.Params{}}
}

// CommentID parameter.
func (b *NotesEditCommentBuilder) CommentID(v int) *NotesEditCommentBuilder {
	b.Params["comment_id"] = v
	return b
}

// OwnerID note owner ID.
func (b *NotesEditCommentBuilder) OwnerID(v int) *NotesEditCommentBuilder {
	b.Params["owner_id"] = v
	return b
}

// Message new comment text.
func (b *NotesEditCommentBuilder) Message(v string) *NotesEditCommentBuilder {
	b.Params["message"] = v
	return b
}

// NotesGetBuilder builder.
//
// Returns a list of notes created by a user.
//
// https://vk.com/dev/notes.get
type NotesGetBuilder struct {
	api.Params
}

// NewNotesGetBuilder func.
func NewNotesGetBuilder() *NotesGetBuilder {
	return &NotesGetBuilder{api.Params{}}
}

// NoteIDs note IDs.
func (b *NotesGetBuilder) NoteIDs(v []int) *NotesGetBuilder {
	b.Params["note_ids"] = v
	return b
}

// UserID note owner ID.
func (b *NotesGetBuilder) UserID(v int) *NotesGetBuilder {
	b.Params["user_id"] = v
	return b
}

// Offset parameter.
func (b *NotesGetBuilder) Offset(v int) *NotesGetBuilder {
	b.Params["offset"] = v
	return b
}

// Count number of notes to return.
func (b *NotesGetBuilder) Count(v int) *NotesGetBuilder {
	b.Params["count"] = v
	return b
}

// Sort parameter.
func (b *NotesGetBuilder) Sort(v int) *NotesGetBuilder {
	b.Params["sort"] = v
	return b
}

// NotesGetByIDBuilder builder.
//
// Returns a note by its ID.
//
// https://vk.com/dev/notes.getById
type NotesGetByIDBuilder struct {
	api.Params
}

// NewNotesGetByIDBuilder func.
func NewNotesGetByIDBuilder() *NotesGetByIDBuilder {
	return &NotesGetByIDBuilder{api.Params{}}
}

// NoteID parameter.
func (b *NotesGetByIDBuilder) NoteID(v int) *NotesGetByIDBuilder {
	b.Params["note_id"] = v
	return b
}

// OwnerID note owner ID.
func (b *NotesGetByIDBuilder) OwnerID(v int) *NotesGetByIDBuilder {
	b.Params["owner_id"] = v
	return b
}

// NeedWiki parameter.
func (b *NotesGetByIDBuilder) NeedWiki(v bool) *NotesGetByIDBuilder {
	b.Params["need_wiki"] = v
	return b
}

// NotesGetCommentsBuilder builder.
//
// Returns a list of comments on a note.
//
// https://vk.com/dev/notes.getComments
type NotesGetCommentsBuilder struct {
	api.Params
}

// NewNotesGetCommentsBuilder func.
func NewNotesGetCommentsBuilder() *NotesGetCommentsBuilder {
	return &NotesGetCommentsBuilder{api.Params{}}
}

// NoteID parameter.
func (b *NotesGetCommentsBuilder) NoteID(v int) *NotesGetCommentsBuilder {
	b.Params["note_id"] = v
	return b
}

// OwnerID note owner ID.
func (b *NotesGetCommentsBuilder) OwnerID(v int) *NotesGetCommentsBuilder {
	b.Params["owner_id"] = v
	return b
}

// Sort parameter.
func (b *NotesGetCommentsBuilder) Sort(v int) *NotesGetCommentsBuilder {
	b.Params["sort"] = v
	return b
}

// Offset parameter.
func (b *NotesGetCommentsBuilder) Offset(v int) *NotesGetCommentsBuilder {
	b.Params["offset"] = v
	return b
}

// Count number of comments to return.
func (b *NotesGetCommentsBuilder) Count(v int) *NotesGetCommentsBuilder {
	b.Params["count"] = v
	return b
}

// NotesRestoreCommentBuilder builder.
//
// Restores a deleted comment on a note.
//
// https://vk.com/dev/notes.restoreComment
type NotesRestoreCommentBuilder struct {
	api.Params
}

// NewNotesRestoreCommentBuilder func.
func NewNotesRestoreCommentBuilder() *NotesRestoreCommentBuilder {
	return &NotesRestoreCommentBuilder{api.Params{}}
}

// CommentID parameter.
func (b *NotesRestoreCommentBuilder) CommentID(v int) *NotesRestoreCommentBuilder {
	b.Params["comment_id"] = v
	return b
}

// OwnerID note owner ID.
func (b *NotesRestoreCommentBuilder) OwnerID(v int) *NotesRestoreCommentBuilder {
	b.Params["owner_id"] = v
	return b
}
