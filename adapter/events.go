package adapter

import (
	"github.com/SevereCloud/vksdk/object"
)

type ChatCreate struct {
	Channel string // The id of created channel.
	Text    string // Chat title

	// original response (objects.MessageNewObject)
	Data interface{}
}

type ChatTitleUpdate struct {
	Channel string // The id of created channel.
	NewText string // New chat title

	// original response (objects.MessageNewObject)
	Data interface{}
}

type ChatPhotoUpdate struct {
	Channel  string                            // The id of created channel.
	NewPhoto object.MessagesMessageActionPhoto // The object with new cover photo urls
	Removed  bool                              // true if cover photo has changed, otherwise false

	// original response (objects.MessageNewObject)
	Data interface{}
}

type ChatPinUpdate struct {
	Channel   string // The id of created channel.
	UserID    string // A string identifying the user who changed pin
	MessageID string // A string identifying the changed pin message
	Unpinned  bool   // true if message has unpinned, otherwise false

	// original response (objects.MessageNewObject)
	Data interface{}
}

type UserEnteredChat struct {
	Channel string // The channel over which the message was received.
	UserID  string // A string identifying the new user in chat
	ByLink  bool

	// original response (objects.MessageNewObject)
	Data interface{}
}

type UserLeavedChat struct {
	Channel string // The channel over which the message was received.
	UserID  string // A string identifying the leaved user in chat

	// original response (objects.MessageNewObject)
	Data interface{}
}
