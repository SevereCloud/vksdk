package object // import "github.com/SevereCloud/vksdk/5.92/object"

// AccountNameRequest struct
type AccountNameRequest struct {
	FirstName string `json:"first_name"`
	ID        int    `json:"id"`
	LastName  string `json:"last_name"`
	Status    string `json:"status"`
}

// AccountPushConversations struct
type AccountPushConversations struct {
	Count int                            `json:"count"`
	Items []accountPushConversationsItem `json:"items"`
}

type accountPushConversationsItem struct {
	DisabledUntil int `json:"disabled_until"`
	PeerID        int `json:"peer_id"`
	Sound         int `json:"sound"`
}

// AccountPushParams struct
type AccountPushParams struct {
	AppRequest     []string `json:"app_request"`
	Birthday       []string `json:"birthday"`
	Chat           []string `json:"chat"`
	Comment        []string `json:"comment"`
	EventSoon      []string `json:"event_soon"`
	Friend         []string `json:"friend"`
	FriendAccepted []string `json:"friend_accepted"`
	FriendFound    []string `json:"friend_found"`
	GroupAccepted  []string `json:"group_accepted"`
	GroupInvite    []string `json:"group_invite"`
	Like           []string `json:"like"`
	Mention        []string `json:"mention"`
	Msg            []string `json:"msg"`
	NewPost        []string `json:"new_post"`
	PhotosTag      []string `json:"photos_tag"`
	Reply          []string `json:"reply"`
	Repost         []string `json:"repost"`
	SdkOpen        []string `json:"sdk_open"`
	WallPost       []string `json:"wall_post"`
	WallPublish    []string `json:"wall_publish"`
}

// AccountOffer struct
type AccountOffer struct {
	Description      string `json:"description"`
	ID               int    `json:"id"`
	Img              string `json:"img"`
	Instruction      string `json:"instruction"`
	InstructionHTML  string `json:"instruction_html"`
	Price            int    `json:"price"`
	ShortDescription string `json:"short_description"`
	Tag              string `json:"tag"`
	Title            string `json:"title"`
}
