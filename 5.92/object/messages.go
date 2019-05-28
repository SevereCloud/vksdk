package object // import "github.com/SevereCloud/vksdk/5.92/object"

// MessagesMessage struct
type MessagesMessage struct {
	AdminAuthorID         int                         `json:"admin_author_id"` // Only for messages from community. Contains user ID of community admin, who sent this message.
	Action                messagesMessageAction       `json:"action"`
	Attachments           []messagesMessageAttachment `json:"attachments"`
	ConversationMessageID int                         `json:"conversation_message_id"` // Unique auto-incremented number for all messages with this peer
	Date                  int                         `json:"date"`                    // Date when the message has been sent in Unixtime
	Deleted               int                         `json:"deleted"`                 // Is it an deleted message
	FromID                int                         `json:"from_id"`                 // Message author's ID
	FwdMessages           []MessagesMessage           `json:"fwd_messages"`            // Forwarded messages
	ReplyMessage          *MessagesMessage            `json:"reply_message"`
	Geo                   baseGeo                     `json:"geo"`
	ID                    int                         `json:"id"`        // Message ID
	Important             bool                        `json:"important"` // Is it an important message
	IsHidden              bool                        `json:"is_hidden"`
	Keyboard              MessagesKeyboard            `json:"keyboard"`
	Out                   int                         `json:"out"` // Information whether the message is outcoming
	Payload               string                      `json:"payload"`
	PeerID                int                         `json:"peer_id"`   // Peer ID
	RandomID              int                         `json:"random_id"` // ID used for sending messages. It returned only for outgoing messages
	Ref                   string                      `json:"ref"`
	RefSource             string                      `json:"ref_source"`
	Text                  string                      `json:"text"`          // Message text
	UpdateTime            int                         `json:"update_time"`   // Date when the message has been updated in Unixtime
	MembersСount          int                         `json:"members_count"` // Members number
}

// MessagesKeyboard struct
type MessagesKeyboard struct {
	AuthorID int                        `json:"author_id,omitempty"` // Community or bot, which set this keyboard
	Buttons  [][]MessagesKeyboardButton `json:"buttons"`
	OneTime  bool                       `json:"one_time"` // Should this keyboard disappear on first use
}

// AddRow add row in MessagesKeyboard
func (keyboard *MessagesKeyboard) AddRow() {
	if len(keyboard.Buttons) == 0 {
		keyboard.Buttons = make([][]MessagesKeyboardButton, 1)
	} else {
		row := make([]MessagesKeyboardButton, 0)
		keyboard.Buttons = append(keyboard.Buttons, row)
	}
}

// AddButton add button in last row
func (keyboard *MessagesKeyboard) AddButton(label string, payload string, color string) {
	button := MessagesKeyboardButton{
		Action: MessagesKeyboardButtonAction{
			Label:   label,
			Payload: payload,
			Type:    "text",
		},
		Color: color,
	}

	lastRow := len(keyboard.Buttons) - 1
	keyboard.Buttons[lastRow] = append(keyboard.Buttons[lastRow], button)
}

// MessagesKeyboardButton struct
type MessagesKeyboardButton struct {
	Action MessagesKeyboardButtonAction `json:"action"`
	Color  string                       `json:"color"` // Button color
}

// MessagesKeyboardButtonAction struct
type MessagesKeyboardButtonAction struct {
	AppID   int    `json:"app_id,omitempty"`   // Fragment value in app link like vk.com/app{app_id}_-654321#hash
	Hash    string `json:"hash,omitempty"`     // Fragment value in app link like vk.com/app123456_-654321#{hash}
	Label   string `json:"label,omitempty"`    // Label for button
	OwnerID int    `json:"owner_id,omitempty"` // Fragment value in app link like vk.com/app123456_{owner_id}#hash
	Payload string `json:"payload"`            // Additional data sent along with message for developer convenience
	Type    string `json:"type"`               // Button type
}

// MessagesChat struct
type MessagesChat struct {
	AdminID      int                       `json:"admin_id"`  // Chat creator ID
	ID           int                       `json:"id"`        // Chat ID
	Kicked       int                       `json:"kicked"`    // Shows that user has been kicked from the chat
	Left         int                       `json:"left"`      // Shows that user has been left the chat
	Photo100     string                    `json:"photo_100"` // URL of the preview image with 100 px in width
	Photo200     string                    `json:"photo_200"` // URL of the preview image with 200 px in width
	Photo50      string                    `json:"photo_50"`  // URL of the preview image with 50 px in width
	PushSettings messagesChatPushSettings  `json:"push_settings"`
	Title        string                    `json:"title"` // Chat title
	Type         string                    `json:"type"`  // Chat type
	Users        []int                     `json:"users"`
	MembersCount int                       `json:"members_count"`
	Members      []int                     `json:"members"`
	Photo        messagesChatSettingsPhoto `json:"photo"`
	Joined       bool                      `json:"joined"`
}

// MessagesChatFull struct
type MessagesChatFull struct {
	AdminID      int                        `json:"admin_id"`  // Chat creator ID
	ID           int                        `json:"id"`        // Chat ID
	Kicked       int                        `json:"kicked"`    // Shows that user has been kicked from the chat
	Left         int                        `json:"left"`      // Shows that user has been left the chat
	Photo100     string                     `json:"photo_100"` // URL of the preview image with 100 px in width
	Photo200     string                     `json:"photo_200"` // URL of the preview image with 200 px in width
	Photo50      string                     `json:"photo_50"`  // URL of the preview image with 50 px in width
	PushSettings messagesChatPushSettings   `json:"push_settings"`
	Title        string                     `json:"title"` // Chat title
	Type         string                     `json:"type"`  // Chat type
	Users        []messagesUserXtrInvitedBy `json:"users"`
}

type messagesChatPushSettings struct {
	DisabledUntil int `json:"disabled_until"` // Time until that notifications are disabled
	Sound         int `json:"sound"`          // Information whether the sound is on
}

type messagesChatSettingsPhoto struct {
	Photo100 string `json:"photo_100"`
	Photo200 string `json:"photo_200"`
	Photo50  string `json:"photo_50"`
}

// MessagesConversation struct
type MessagesConversation struct {
	CanWrite        messagesConversationCanWrite     `json:"can_write"`
	ChatSettings    messagesConversationChatSettings `json:"chat_settings"`
	InRead          int                              `json:"in_read"`         // Last message user have read
	LastMessageID   int                              `json:"last_message_id"` // ID of the last message in conversation
	Mentions        []int                            `json:"mentions"`        // IDs of messages with mentions
	MessageRequest  string                           `json:"message_request"`
	OutRead         int                              `json:"out_read"` // Last outcoming message have been read by the opponent
	Peer            messagesConversationPeer         `json:"peer"`
	PushSettings    messagesConversationPushSettings `json:"push_settings"`
	Important       bool                             `json:"important"`
	Unanswered      bool                             `json:"unanswered"`
	UnreadCount     int                              `json:"unread_count"` // Unread messages number
	CurrentKeyboard MessagesKeyboard                 `json:"current_keyboard"`
}

type messagesConversationCanWrite struct {
	Allowed bool `json:"allowed"`
	Reason  int  `json:"reason"`
}

type messagesConversationChatSettings struct {
	MembersCount  int                       `json:"members_count"`
	Photo         messagesChatSettingsPhoto `json:"photo"`
	PinnedMessage messagesPinnedMessage     `json:"pinned_message"`
	State         string                    `json:"state"`
	Title         string                    `json:"title"`
	ActiveIds     []int                     `json:"active_ids"`
	ACL           struct {
		CanInvite           bool `json:"can_invite"`
		CanChangeInfo       bool `json:"can_change_info"`
		CanChangePin        bool `json:"can_change_pin"`
		CanPromoteUsers     bool `json:"can_promote_users"`
		CanSeeInviteLink    bool `json:"can_see_invite_link"`
		CanChangeInviteLink bool `json:"can_change_invite_link"`
	} `json:"acl"`
	IsGroupChannel bool `json:"is_group_channel"`
	OwnerID        int  `json:"owner_id"`
}

type messagesConversationPeer struct {
	ID      int    `json:"id"`
	LocalID int    `json:"local_id"`
	Type    string `json:"type"`
}

type messagesConversationPushSettings struct {
	DisabledUntil   int  `json:"disabled_until"`
	DisabledForever bool `json:"disabled_forever"`
	NoSound         bool `json:"no_sound"`
}

// MessagesConversationWithMessage struct
type MessagesConversationWithMessage struct {
	Conversation MessagesConversation `json:"conversation"`
	LastMessage  MessagesMessage      `json:"last_message"`
}

// MessagesDialog struct
type MessagesDialog struct {
	Important  int             `json:"important"`
	InRead     int             `json:"in_read"`
	Message    MessagesMessage `json:"message"`
	OutRead    int             `json:"out_read"`
	Unanswered int             `json:"unanswered"`
	Unread     int             `json:"unread"`
}

// MessagesHistoryAttachment struct
type MessagesHistoryAttachment struct {
	Attachment messagesHistoryMessageAttachment `json:"attachment"`
	MessageID  int                              `json:"message_id"` // Message ID
}

type messagesHistoryMessageAttachment struct {
	Audio  AudioAudioFull `json:"audio"`
	Doc    DocsDoc        `json:"doc"`
	Link   baseLink       `json:"link"`
	Market baseLink       `json:"market"`
	Photo  PhotosPhoto    `json:"photo"`
	Share  baseLink       `json:"share"`
	Type   string         `json:"type"`
	Video  VideoVideo     `json:"video"`
	Wall   baseLink       `json:"wall"`
}

// MessagesLastActivity struct
type MessagesLastActivity struct {
	Online int `json:"online"` // Information whether user is online
	Time   int `json:"time"`   // Time when user was online in Unixtime
}

// MessagesLongpollParams struct
type MessagesLongpollParams struct {
	Key    string `json:"key"`    // Key
	Pts    int    `json:"pts"`    // Persistent timestamp
	Server string `json:"server"` // Server URL
	Ts     int    `json:"ts"`     // Timestamp
}

type messagesMessageAction struct {
	ConversationMessageID int                        `json:"conversation_message_id"` // Message ID
	Email                 string                     `json:"email"`                   // Email address for chat_invite_user or chat_kick_user actions
	MemberID              int                        `json:"member_id"`               // User or email peer ID
	Message               string                     `json:"message"`                 // Message body of related message
	Photo                 messagesMessageActionPhoto `json:"photo"`
	Text                  string                     `json:"text"` // New chat title for chat_create and chat_title_update actions
	Type                  string                     `json:"type"`
}

type messagesMessageActionPhoto struct {
	Photo100 string `json:"photo_100"` // URL of the preview image with 100px in width
	Photo200 string `json:"photo_200"` // URL of the preview image with 200px in width
	Photo50  string `json:"photo_50"`  // URL of the preview image with 50px in width
}

type messagesMessageAttachment struct {
	Audio             AudioAudioFull       `json:"audio"`
	Doc               DocsDoc              `json:"doc"`
	Gift              GiftsLayout          `json:"gift"`
	Link              baseLink             `json:"link"`
	Market            MarketMarketItem     `json:"market"`
	MarketMarketAlbum MarketMarketAlbum    `json:"market_market_album"`
	Photo             PhotosPhoto          `json:"photo"`
	Sticker           baseSticker          `json:"sticker"`
	Type              string               `json:"type"`
	Video             VideoVideo           `json:"video"`
	Wall              wallWallpostAttached `json:"wall"`
	WallReply         WallWallComment      `json:"wall_reply"`
}

type messagesPinnedMessage struct {
	Attachments           []messagesMessageAttachment `json:"attachments"`
	ConversationMessageID int                         `json:"conversation_message_id"` // Unique auto-incremented number for all messages with this peer
	Date                  int                         `json:"date"`                    // Date when the message has been sent in Unixtime
	FromID                int                         `json:"fromНастроены_id"`        // Message author's ID
	FwdMessages           *MessagesMessage            `json:"fwd_messages"`
	Geo                   baseGeo                     `json:"geo"`
	ID                    int                         `json:"id"`      // Message ID
	PeerID                int                         `json:"peer_id"` // Peer ID
	ReplyMessage          *MessagesMessage            `json:"reply_message"`
	Text                  string                      `json:"text"` // Message text
}

type messagesUserXtrInvitedBy struct {
}
