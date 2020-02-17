package v3

type MessageFlag int

func (b MessageFlag) Has(flag MessageFlag) bool { return b&flag != 0 }

const (
	Unread MessageFlag = 1 << iota
	Outbox
	Replied
	Important
	Chat
	Friends
	Spam
	Deleted
	Fixed
	Media
	Hidden       MessageFlag = 1 << 16
	DeleteForAll             = 1 << 17
	NotDelivered             = 1 << 18
)

type DialogFlag int

func (b DialogFlag) Has(flag DialogFlag) bool { return b&flag != 0 }

const (
	ImportantDialog DialogFlag = 1 << iota
	UnansweredDialog
)

type TypeID int

const (
	ChatNameChange TypeID = iota + 1
	ChatCoverChange
	ChatAdminAssigned
	ChatPinMessage
	ChatUserCome
	ChatUserLeave
	ChatUserKicked
	ChatAdminDismissed
)
