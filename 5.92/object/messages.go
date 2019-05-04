package object // import "github.com/SevereCloud/vksdk/5.92/object"

// MessagesMessage struct
type MessagesMessage struct {
	AdminAuthorID         int                         `json:"admin_author_id"`
	Action                messagesMessageAction       `json:"action"`
	Attachments           []messagesMessageAttachment `json:"attachments"`
	ConversationMessageID int                         `json:"conversation_message_id"`
	Date                  int                         `json:"date"`
	FromID                int                         `json:"from_id"`
	FwdMessages           []MessagesMessage           `json:"fwd_messages"`
	ReplyMessage          *MessagesMessage            `json:"reply_message"`
	Geo                   baseGeo                     `json:"geo"`
	ID                    int                         `json:"id"`
	Important             bool                        `json:"important"`
	IsHidden              bool                        `json:"is_hidden"`
	Keyboard              MessagesKeyboard            `json:"keyboard"`
	Out                   int                         `json:"out"`
	Payload               string                      `json:"payload"`
	PeerID                int                         `json:"peer_id"`
	RandomID              int                         `json:"random_id"`
	Ref                   string                      `json:"ref"`
	RefSource             string                      `json:"ref_source"`
	Text                  string                      `json:"text"`
	UpdateTime            int                         `json:"update_time"`
	MembersСount          int                         `json:"members_count"` // FIXME: check members_count
}

// MessagesKeyboard struct
type MessagesKeyboard struct {
	Buttons [][]MessagesKeyboardButton `json:"buttons"`
	OneTime bool                       `json:"one_time"`
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
	Color  string                       `json:"color"`
}

// MessagesKeyboardButtonAction struct
type MessagesKeyboardButtonAction struct {
	Label   string `json:"label"`
	Payload string `json:"payload"`
	Type    string `json:"type"`
}

// MessagesChat struct
type MessagesChat struct {
	AdminID      int                      `json:"admin_id"`
	ID           int                      `json:"id"`
	Kicked       int                      `json:"kicked"`
	Left         int                      `json:"left"`
	Photo100     string                   `json:"photo_100"`
	Photo200     string                   `json:"photo_200"`
	Photo50      string                   `json:"photo_50"`
	PushSettings messagesChatPushSettings `json:"push_settings"`
	Title        string                   `json:"title"`
	Type         string                   `json:"type"`
	Users        []int                    `json:"users"`
	MembersCount int                      `json:"members_count"`
	Members      []int                    `json:"members"`
	Photo        struct {
		Photo50  string `json:"photo_50"`
		Photo100 string `json:"photo_100"`
		Photo200 string `json:"photo_200"`
	} `json:"photo"`
	Joined bool `json:"joined"`
}

type messagesChatFull struct {
	AdminID      int                        `json:"admin_id"`
	ID           int                        `json:"id"`
	Kicked       int                        `json:"kicked"`
	Left         int                        `json:"left"`
	Photo100     string                     `json:"photo_100"`
	Photo200     string                     `json:"photo_200"`
	Photo50      string                     `json:"photo_50"`
	PushSettings messagesChatPushSettings   `json:"push_settings"`
	Title        string                     `json:"title"`
	Type         string                     `json:"type"`
	Users        []messagesUserXtrInvitedBy `json:"users"`
}

type messagesChatPushSettings struct {
	DisabledUntil int `json:"disabled_until"`
	Sound         int `json:"sound"`
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
	Important       bool                             `json:"important"`
	InRead          int                              `json:"in_read"`
	OutRead         int                              `json:"out_read"`
	Peer            messagesConversationPeer         `json:"peer"`
	PushSettings    messagesConversationPushSettings `json:"push_settings"`
	Unanswered      bool                             `json:"unanswered"`
	UnreadCount     int                              `json:"unread_count"`
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
	DisabledForever bool `json:"disabled_forever"`
	DisabledUntil   int  `json:"disabled_until"`
	NoSound         bool `json:"no_sound"`
}

// MessagesConversationWithMessage struct
type MessagesConversationWithMessage struct {
	Conversation MessagesConversation `json:"conversation"`
	LastMessage  MessagesMessage      `json:"last_message"`
}

type messagesDialog struct {
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
	MessageID  int                              `json:"message_id"`
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
	Online int `json:"online"`
	Time   int `json:"time"`
}

type messagesLongpollMessages struct {
	Count int               `json:"count"`
	Items []MessagesMessage `json:"items"`
}

// MessagesLongpollParams struct
type MessagesLongpollParams struct {
	Key    string `json:"key"`
	Pts    int    `json:"pts"`
	Server string `json:"server"`
	Ts     int    `json:"ts"`
}

type messagesMessageAction struct {
	ConversationMessageID int                       `json:"conversation_message_id"`
	Email                 string                    `json:"email"`
	MemberID              int                       `json:"member_id"`
	Message               string                    `json:"message"`
	Photo                 messagesChatSettingsPhoto `json:"photo"`
	Text                  string                    `json:"text"`
	Type                  string                    `json:"type"`
}

type messagesMessageAttachment struct {
	Audio             AudioAudioFull       `json:"audio"`
	Doc               DocsDoc              `json:"doc"`
	Gift              GiftsLayout          `json:"gift"`
	Link              baseLink             `json:"link"`
	Market            MarketMarketItem     `json:"market"`
	MarketMarketAlbum marketMarketAlbum    `json:"market_market_album"`
	Photo             PhotosPhoto          `json:"photo"`
	Sticker           baseSticker          `json:"sticker"`
	Type              string               `json:"type"`
	Video             VideoVideo           `json:"video"`
	Wall              wallWallpostAttached `json:"wall"`
	WallReply         wallWallComment      `json:"wall_reply"`
}

type messagesPinnedMessage struct {
	Attachments           []messagesMessageAttachment `json:"attachments"`
	ConversationMessageID int                         `json:"conversation_message_id"`
	Date                  int                         `json:"date"`
	FromID                int                         `json:"from_id"`
	FwdMessages           []MessagesMessage           `json:"fwd_messages"`
	Geo                   baseGeo                     `json:"geo"`
	ID                    int                         `json:"id"`
	PeerID                int                         `json:"peer_id"`
	Text                  string                      `json:"text"`
}

type messagesUserXtrInvitedBy struct {
}
