package object // import "github.com/SevereCloud/vksdk/object"

// AccountNameRequest struct
type AccountNameRequest struct {
	FirstName string `json:"first_name"` // First name in request
	ID        int    `json:"id"`         // Request ID needed to cancel the request
	LastName  string `json:"last_name"`  // Last name in request
	Status    string `json:"status"`
}

// AccountPushConversations struct
type AccountPushConversations struct {
	Count int                             `json:"count"` // Items count
	Items []*accountPushConversationsItem `json:"items"`
}

type accountPushConversationsItem struct {
	DisabledUntil int `json:"disabled_until"` // Time until that notifications are disabled in seconds
	PeerID        int `json:"peer_id"`        // Peer ID
	Sound         int `json:"sound"`          // Information whether the sound are enabled
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
	Description      string `json:"description"`       // Offer description
	ID               int    `json:"id"`                // Offer ID
	Img              string `json:"img"`               // URL of the preview image
	Instruction      string `json:"instruction"`       // Instruction how to process the offer
	InstructionHTML  string `json:"instruction_html"`  // Instruction how to process the offer (HTML format)
	Price            int    `json:"price"`             // Offer price
	ShortDescription string `json:"short_description"` // Offer short description
	Tag              string `json:"tag"`               // Offer tag
	Title            string `json:"title"`             // Offer title
}

// AccountAccountCounters struct
type AccountAccountCounters struct {
	AppRequests            int `json:"app_requests"`            // New app requests number
	Events                 int `json:"events"`                  // New events number
	Friends                int `json:"friends"`                 // New friends requests number
	FriendsRecommendations int `json:"friends_recommendations"` // New friends recommendations number
	FriendsSuggestions     int `json:"friends_suggestions"`     // New friends suggestions number
	Gifts                  int `json:"gifts"`                   // New gifts number
	Groups                 int `json:"groups"`                  // New groups number
	Messages               int `json:"messages"`                // New messages number
	Notifications          int `json:"notifications"`           // New notifications number
	Photos                 int `json:"photos"`                  // New photo tags number
	SDK                    int `json:"sdk"`                     // New SDK number
	MenuDiscoverBadge      int `json:"menu_discover_badge"`     // New menu discover badge number
	Videos                 int `json:"videos"`                  // New video tags number
}

// AccountInfo struct
type AccountInfo struct {
	Country           string `json:"country"`           // Country code
	HTTPSRequired     int    `json:"https_required"`    // Information whether HTTPS-only is enabled
	Intro             int    `json:"intro"`             // Information whether user has been processed intro
	Lang              int    `json:"lang"`              // Language ID
	NoWallReplies     int    `json:"no_wall_replies"`   // Information whether wall comments should be hidden
	OwnPostsDefault   int    `json:"own_posts_default"` // Information whether only owners posts should be shown
	TwoFactorRequired int    `json:"2fa_required"`      // Two factor authentication is enabled
}

// AccountPushSettings struct
type AccountPushSettings struct {
	Conversations AccountPushConversations `json:"conversations"`
	Disabled      int                      `json:"disabled"`       // Information whether notifications are disabled
	DisabledUntil int                      `json:"disabled_until"` // Time until that notifications are disabled in Unixtime
	Settings      AccountPushParams        `json:"settings"`
}

// AccountUserSettings struct
type AccountUserSettings struct {
	Bdate            string             `json:"bdate"`            // User's date of birth
	BdateVisibility  int                `json:"bdate_visibility"` // Information whether user's birthdate are hidden
	City             BaseObject         `json:"city"`
	Country          BaseCountry        `json:"country"`
	FirstName        string             `json:"first_name"`  // User first name
	HomeTown         string             `json:"home_town"`   // User's hometown
	LastName         string             `json:"last_name"`   // User last name
	MaidenName       string             `json:"maiden_name"` // User maiden name
	NameRequest      AccountNameRequest `json:"name_request"`
	Phone            string             `json:"phone"`    // User phone number with some hidden digits
	Relation         int                `json:"relation"` // User relationship status
	RelationPartner  UsersUserMin       `json:"relation_partner"`
	RelationPending  int                `json:"relation_pending"` // Information whether relation status is pending
	RelationRequests []UsersUserMin     `json:"relation_requests"`
	ScreenName       string             `json:"screen_name"` // Domain name of the user's page
	Sex              int                `json:"sex"`         // User sex
	Status           string             `json:"status"`      // User status
}
