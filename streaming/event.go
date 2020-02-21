package streaming

import (
	"github.com/SevereCloud/vksdk/object"
)

// Event object, if status code is 100
type Event struct {
	EventType EventType `json:"event_type"`
	EventID   struct {
		PostOwnerID  int `json:"post_owner_id"`  // (for EventType = Post, Comment, Share)
		PostID       int `json:"post_id"`        // (for EventType = Post, Comment, Share)
		CommentID    int `json:"comment_id"`     // (for EventType = Comment)
		SharedPostID int `json:"shared_post_id"` // (for EventType = Share)
		TopicOwnerID int `json:"topic_owner_id"` // (for EventType = TopicPost)
		TopicID      int `json:"topic_id"`       // (for EventType = TopicPost)
		TopicPostID  int `json:"topic_post_id"`  // (for EventType = TopicPost)
	} `json:"event_id"`
	EventURL               string                          `json:"event_url"` // Object URL
	Text                   string                          `json:"text"`
	Action                 Action                          `json:"action"` //  action with the object
	ActionTime             int                             `json:"action_time"`
	CreationTime           int                             `json:"creation_time"` // Time when the object has been create
	Attachments            []object.WallWallpostAttachment `json:"attachments"`
	Geo                    object.BaseGeo                  `json:"geo"`
	SharedPostText         string                          `json:"shared_post_text"`
	SharedPostCreationTime int                             `json:"shared_post_creation_time"`
	SignerID               int                             `json:"signer_id"`
	Tags                   []string                        `json:"tags"`   // List of tags which correspond to the event
	Author                 Author                          `json:"author"` // Information about author
}

// Action with the object
type Action string

// Possible values
const (
	New     Action = "new"     // new object has been created
	Update  Action = "update"  // object has been updated
	Delete  Action = "delete"  // object has been deleted
	Restore Action = "restore" // object has been restored
)

// EventType event type
type EventType string

// Possible values
const (
	Post      EventType = "post"
	Comment   EventType = "comment"
	Share     EventType = "share"
	TopicPost EventType = "topic_post"
)

// Author struct
type Author struct {
	ID        int    `json:"id"`         // author ID
	AuthorURL string `json:"author_url"` // URL of the author's page
	// ID of the author of the original post (for EventType = Share)
	SharedPostAuthorID int `json:"shared_post_author_id,omitempty"`
	// URL of the original post's author's page  (for EventType = Share)
	SharedPostAuthorURL string          `json:"shared_post_author_url,omitempty"`
	Platform            object.Platform `json:"platform"` // Platform used by author
}
