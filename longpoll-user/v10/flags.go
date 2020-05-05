package wrapper // import "github.com/SevereCloud/vksdk/longpoll-user/v10"

// MessageFlag type
type MessageFlag int

// Has function
func (b MessageFlag) Has(flag MessageFlag) bool { return b&flag != 0 }

// Each message has a flag, which is a value received by summing up any of the following parameters
const (
	Unread        MessageFlag = 1 << 0  // Message is unread
	Outbox        MessageFlag = 1 << 1  // Message is outgoing
	Important     MessageFlag = 1 << 3  // Message is marked as important
	Chat          MessageFlag = 1 << 4  // Message sent via vk.com
	Friends       MessageFlag = 1 << 5  // Message sent by a friend
	Spam          MessageFlag = 1 << 6  // Message marked as "Spam"
	Deleted       MessageFlag = 1 << 7  // Message was deleted
	AudioListened MessageFlag = 1 << 12 // Voice messages was listened
	Chat2         MessageFlag = 1 << 13 // Message sent via mobile client
	CancelSpam    MessageFlag = 1 << 15 // Message unmarked as "Spam"
	Hidden        MessageFlag = 1 << 16 // Greeting message from a community
	DeletedAll    MessageFlag = 1 << 17 // Message was deleted for all
	ChatIn        MessageFlag = 1 << 19 // Message is ingoing
	NonameFlag    MessageFlag = 1 << 20 // Unknown flag. TODO specify flag.
	ReplyMsg      MessageFlag = 1 << 21 // Reply message
)

// String returns string representation of flag
func (b MessageFlag) String() string {
	result := ""

	switch {
	case b.Has(Unread):
		result += "Unread;"
		fallthrough
	case b.Has(Outbox):
		result += "Outbox;"
		fallthrough
	case b.Has(Important):
		result += "Important;"
		fallthrough
	case b.Has(Chat):
		result += "Chat;"
		fallthrough
	case b.Has(Friends):
		result += "Friends;"
		fallthrough
	case b.Has(Spam):
		result += "Spam;"
		fallthrough
	case b.Has(Deleted):
		result += "Deleted;"
		fallthrough
	case b.Has(AudioListened):
		result += "AudioListened;"
		fallthrough
	case b.Has(Chat2):
		result += "Chat2;"
		fallthrough
	case b.Has(CancelSpam):
		result += "CancelSpam;"
		fallthrough
	case b.Has(Hidden):
		result += "Hidden;"
		fallthrough
	case b.Has(DeletedAll):
		result += "DeletedAll;"
		fallthrough
	case b.Has(ChatIn):
		result += "ChatIn;"
		fallthrough
	case b.Has(NonameFlag):
		result += "NonameFlag;"
		fallthrough
	case b.Has(ReplyMsg):
		result += "ReplyMsg;"
	}

	return result
}

// TypeID chat change type identifier
type TypeID int

// TypeID const
const (
	ChatNameChange          TypeID = iota + 1 // The name of the conversation has changed
	ChatCoverChange                           // The cover of the conversation has changed
	ChatAdminAssigned                         // New administrator has been appointed.
	ChatPermissionChange                      // The permission of the conversation has changed.
	ChatPinnedMessage                         // Message pin
	ChatUserCome                              // User has joined the conversation
	ChatUserLeave                             // User left the conversation.
	ChatUserKicked                            // User was excluded from the conversation.
	ChatAdminDismissed                        // Administrator rights have been removed from the user
	ChatKeyboardStateChange                   // Keyboard shown/hidden
)

// String returns string representation of flag
func (b TypeID) String() string {
	switch b {
	case ChatNameChange:
		return "ChatNameChange"
	case ChatCoverChange:
		return "ChatCoverChange"
	case ChatAdminAssigned:
		return "ChatAdminAssigned"
	case ChatPermissionChange:
		return "ChatPermissionChange"
	case ChatPinnedMessage:
		return "ChatPinnedMessage"
	case ChatUserCome:
		return "ChatUserCome"
	case ChatUserLeave:
		return "ChatUserLeave"
	case ChatUserKicked:
		return "ChatUserKicked"
	case ChatAdminDismissed:
		return "ChatAdminDismissed"
	case ChatKeyboardStateChange:
		return "ChatKeyboardStateChange"
	}

	return ""
}

// MentionFlags type
type MentionFlags int

// Has function
func (b MentionFlags) Has(flag MentionFlags) bool { return b&flag != 0 }
