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

// Maps constant to string representation for debug.
var mapMessageFlagToName = map[MessageFlag]string{
	Unread:        "Unread",
	Outbox:        "Outbox",
	Important:     "Important",
	Chat:          "Chat",
	Friends:       "Friends",
	Spam:          "Spam",
	Deleted:       "Deleted",
	AudioListened: "AudioListened",
	Chat2:         "Chat2",
	CancelSpam:    "CancelSpam",
	Hidden:        "Hidden",
	DeletedAll:    "DeletedAll",
	ChatIn:        "ChatIn",
	NonameFlag:    "NonameFlag",
	ReplyMsg:      "ReplyMsg",
}

// String returns string representation of flag
func (b MessageFlag) String() string {
	result := ""

	for k, v := range mapMessageFlagToName {
		if b.Has(k) {
			result = result + v + ";"
		}
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

// Maps constant to string representation for debug.
var mapTypeIDToName = [...]string{
	ChatNameChange:          "ChatNameChange",
	ChatCoverChange:         "ChatCoverChange",
	ChatAdminAssigned:       "ChatAdminAssigned",
	ChatPermissionChange:    "ChatPermissionChange",
	ChatPinnedMessage:       "ChatPinnedMessage",
	ChatUserCome:            "ChatUserCome",
	ChatUserLeave:           "ChatUserLeave",
	ChatUserKicked:          "ChatUserKicked",
	ChatAdminDismissed:      "ChatAdminDismissed",
	ChatKeyboardStateChange: "ChatKeyboardStateChange",
}

// String returns string representation of flag
func (b TypeID) String() string {
	return mapTypeIDToName[b]
}

type MentionFlags int
