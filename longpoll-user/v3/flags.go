package wrapper // import "github.com/SevereCloud/vksdk/longpoll-user/v3"

// MessageFlag type.
type MessageFlag int

// Has function.
func (b MessageFlag) Has(flag MessageFlag) bool { return b&flag != 0 }

// Each message has a flag, which is a value received by summing up any of the following parameters.
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

// DialogFlag type.
type DialogFlag int

// Has func.
func (b DialogFlag) Has(flag DialogFlag) bool { return b&flag != 0 }

// Each dialog has flags, which are values received by summing up any of the
// following parameters.
const (
	ImportantDialog  DialogFlag = 1 << iota // Important dialog
	UnansweredDialog                        // Dialog without a community reply
)

// TypeID chat change type identifier.
type TypeID int

// TypeID const.
const (
	ChatNameChange     TypeID = iota + 1 // The name of the conversation has changed
	ChatCoverChange                      // The cover of the conversation has changed
	ChatAdminAssigned                    // New administrator has been appointed.
	ChatPinMessage                       // Message pin
	ChatUserCome                         // User has joined the conversation
	ChatUserLeave                        // User left the conversation.
	ChatUserKicked                       // User was excluded from the conversation.
	ChatAdminDismissed                   // Administrator rights have been removed from the user
)
