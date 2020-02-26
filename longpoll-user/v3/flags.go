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
	case b.Has(Replied):
		result += "Replied;"
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
	case b.Has(Fixed):
		result += "Fixed;"
		fallthrough
	case b.Has(Media):
		result += "Media;"
		fallthrough
	case b.Has(Hidden):
		result += "Hidden;"
		fallthrough
	case b.Has(DeleteForAll):
		result += "DeleteForAll;"
		fallthrough
	case b.Has(NotDelivered):
		result += "NotDelivered;"
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

// String returns string representation of flag
func (b DialogFlag) String() string {
	result := ""

	switch {
	case b.Has(ImportantDialog):
		result += "ImportantDialog;"
		fallthrough
	case b.Has(UnansweredDialog):
		result += "UnansweredDialog;"
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

// String returns string representation of flag
func (b TypeID) String() string {
	switch b {
	case ChatNameChange:
		return "ChatNameChange"
	case ChatCoverChange:
		return "ChatCoverChange"
	case ChatAdminAssigned:
		return "ChatAdminAssigned"
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
	}

	return ""
}
