package wrapper

type MessageFlag int

func (b MessageFlag) Has(flag MessageFlag) bool { return b&flag != 0 }

const (
	Unread       MessageFlag = 1
	Outbox       MessageFlag = 2
	Replied      MessageFlag = 4
	Important    MessageFlag = 8
	Chat         MessageFlag = 16
	Friends      MessageFlag = 32
	Spam         MessageFlag = 64
	Deleted      MessageFlag = 128
	Fixed        MessageFlag = 256
	Media        MessageFlag = 512
	Hidden       MessageFlag = 65536
	DeleteForAll MessageFlag = 131072
	NotDelivered MessageFlag = 262144
)

type DialogFlag int

func (b DialogFlag) Has(flag DialogFlag) bool { return b&flag != 0 }

const (
	ImportantDialog  DialogFlag = 1
	UnansweredDialog DialogFlag = 2
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
