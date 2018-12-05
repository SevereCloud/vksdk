package object // import "github.com/severecloud/vksdk/object"

type friendsFriendStatus struct {
	FriendStatus   int    `json:"friend_status"`
	ReadState      int    `json:"read_state"`
	RequestMessage string `json:"request_message"`
	Sign           string `json:"sign"`
	UserID         int    `json:"user_id"`
}

type friendsFriendsList struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type friendsMutualFriend struct {
	CommonCount   int   `json:"common_count"`
	CommonFriends []int `json:"common_friends"`
	ID            int   `json:"id"`
}

type friendsRequests struct {
	From   string                `json:"from"`
	Mutual friendsRequestsMutual `json:"mutual"`
	UserID int                   `json:"user_id"`
}

type friendsRequestsMutual struct {
	Count int   `json:"count"`
	Users []int `json:"users"`
}

type friendsRequestsXtrMessage struct {
	From    string                `json:"from"`
	Message string                `json:"message"`
	Mutual  friendsRequestsMutual `json:"mutual"`
	UserID  int                   `json:"user_id"`
}
