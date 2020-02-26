package wrapper // import "github.com/SevereCloud/vksdk/longpoll-user/v3"

// MessageFlag type
type MessageFlag int

// Has function
func (b MessageFlag) Has(flag MessageFlag) bool { return b&flag != 0 }

// Each message has a flag, which is a value received by summing up any of the following parameters
const (
	Unread       MessageFlag = 1 << iota // Message is unread
	Outbox                               // Message is outgoing
	Replied                              // Message was answered
	Important                            // Message is marked as important
	Chat                                 // Message sent via chat
	Friends                              // Message sent by a friend
	Spam                                 // Message marked as "Spam"
	Deleted                              // Message was deleted
	Fixed                                //	Message was user-checked for spam
	Media                                // Message has media content
	Hidden       MessageFlag = 1 << 16   // Greeting message from a community
	DeleteForAll             = 1 << 17   // Message was deleted for all
	NotDelivered             = 1 << 18   // Incoming message not delivered
)

// Maps constant to string representation for debug.
var mapMessageFlagToName = map[MessageFlag]string{
	Unread:       "Unread",
	Outbox:       "Outbox",
	Replied:      "Replied",
	Important:    "Important",
	Chat:         "Chat",
	Friends:      "Friends",
	Spam:         "Spam",
	Deleted:      "Deleted",
	Fixed:        "Fixed",
	Media:        "Media",
	Hidden:       "Hidden",
	DeleteForAll: "DeleteForAll",
	NotDelivered: "NotDelivered",
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

// DialogFlag type
type DialogFlag int

// Has func
func (b DialogFlag) Has(flag DialogFlag) bool { return b&flag != 0 }

// Each dialog has flags, which are values received by summing up any of the
// following parameters.
const (
	ImportantDialog  DialogFlag = 1 << iota // Important dialog
	UnansweredDialog                        // Dialog without a community reply
)

// Maps constant to string representation for debug.
var mapDialogFlagToName = map[DialogFlag]string{
	ImportantDialog:  "ImportantDialog",
	UnansweredDialog: "UnansweredDialog",
}

// String returns string representation of flag
func (b DialogFlag) String() string {
	result := ""

	for k, v := range mapDialogFlagToName {
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
	ChatNameChange     TypeID = iota + 1 // The name of the conversation has changed
	ChatCoverChange                      // The cover of the conversation has changed
	ChatAdminAssigned                    // New administrator has been appointed.
	ChatPinnedMessage                    // Message pin
	ChatUserCome                         // User has joined the conversation
	ChatUserLeave                        // User left the conversation.
	ChatUserKicked                       // User was excluded from the conversation.
	ChatAdminDismissed                   // Administrator rights have been removed from the user
)

// Maps constant to string representation for debug.
var mapTypeIDToName = [...]string{
	ChatNameChange:     "ChatNameChange",
	ChatCoverChange:    "ChatCoverChange",
	ChatAdminAssigned:  "ChatAdminAssigned",
	ChatPinnedMessage:  "ChatPinnedMessage",
	ChatUserCome:       "ChatUserCome",
	ChatUserLeave:      "ChatUserLeave",
	ChatUserKicked:     "ChatUserKicked",
	ChatAdminDismissed: "ChatAdminDismissed",
}

// String returns string representation of flag
func (b TypeID) String() string {
	return mapTypeIDToName[b]
}
