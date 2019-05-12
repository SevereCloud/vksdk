package object // import "github.com/SevereCloud/vksdk/5.92/object"

// FriendsFriendStatus struct
type FriendsFriendStatus struct {
	FriendStatus   int    `json:"friend_status"`
	ReadState      int    `json:"read_state"`      // Information whether request is unviewed
	RequestMessage string `json:"request_message"` // Message sent with request
	Sign           string `json:"sign"`            // MD5 hash for the result validation
	UserID         int    `json:"user_id"`         // User ID
}

type friendsFriendsList struct {
	ID   int    `json:"id"`   // List ID
	Name string `json:"name"` // List title
}

type friendsMutualFriend struct {
	CommonCount   int   `json:"common_count"` // Total mutual friends number
	CommonFriends []int `json:"common_friends"`
	ID            int   `json:"id"` // User ID
}

type friendsRequests struct {
	From   string                `json:"from"` // ID of the user by whom friend has been suggested
	Mutual friendsRequestsMutual `json:"mutual"`
	UserID int                   `json:"user_id"` // User ID
}

type friendsRequestsMutual struct {
	Count int   `json:"count"` // Total mutual friends number
	Users []int `json:"users"`
}

type friendsRequestsXtrMessage struct {
	From    string                `json:"from"`    // ID of the user by whom friend has been suggested
	Message string                `json:"message"` // Message sent with a request
	Mutual  friendsRequestsMutual `json:"mutual"`
	UserID  int                   `json:"user_id"` // User ID
}
