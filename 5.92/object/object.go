package object // import "github.com/SevereCloud/vksdk/5.92/object"

import "encoding/json"

// BaseCountry struct
type BaseCountry struct {
	ID    int    `json:"id"`
	Title string `json:"title"`
}

// BaseObject struct
type BaseObject struct {
	ID    int    `json:"id"`
	Title string `json:"title"`
}

// BaseObjectCount struct
type BaseObjectCount struct {
	Count int `json:"count"`
}

type baseObjectWithName struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

// BaseRequestParam struct
type BaseRequestParam struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

// GroupEvent struct
type GroupEvent struct {
	Type    string          `json:"type"`
	Object  json.RawMessage `json:"object"`
	GroupID int             `json:"group_id"`
	Secret  string          `json:"secret"`
}

// LongpollResponse struct
type LongpollResponse struct {
	Ts      int             `json:"ts"`
	Updates [][]interface{} `json:"updates"`
	Failed  int             `json:"failed"`
}

// LongpollBotResponse struct
type LongpollBotResponse struct {
	Ts      string       `json:"ts"`
	Updates []GroupEvent `json:"updates"`
	Failed  int          `json:"failed"`
}

type baseCommentsInfo struct {
	CanPost       int  `json:"can_post"`
	Count         int  `json:"count"`
	GroupsCanPost bool `json:"groups_can_post"`
	CanClose      bool `json:"can_close"`
	CanOpen       bool `json:"can_open"`
}

type baseGeoCoordinates struct {
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
}

type baseGeo struct {
	Coordinates baseGeoCoordinates `json:"coordinates"`
	Place       basePlace          `json:"place"`
	Showmap     int                `json:"showmap"`
	Type        string             `json:"type"`
}

type baseImage struct {
	Height int    `json:"height"`
	URL    string `json:"url"`
	Width  int    `json:"width"`
	Type   string `json:"type"`
}

type baseLikes struct {
	Count     int `json:"count"`
	UserLikes int `json:"user_likes"`
}

type baseLikesInfo struct {
	CanLike    int `json:"can_like"`
	CanPublish int `json:"can_publish"`
	Count      int `json:"count"`
	UserLikes  int `json:"user_likes"`
}

type baseLink struct {
	Application baseLinkApplication `json:"application"`
	Button      baseLinkButton      `json:"button"`
	Caption     string              `json:"caption"`
	Description string              `json:"description"`
	Photo       PhotosPhoto         `json:"photo"`
	PreviewPage string              `json:"preview_page"`
	PreviewURL  string              `json:"preview_url"`
	Product     baseLinkProduct     `json:"product"`
	Rating      baseLinkRating      `json:"rating"`
	Title       string              `json:"title"`
	URL         string              `json:"url"`
}

type baseLinkApplication struct {
	AppID float64                  `json:"app_id"`
	Store baseLinkApplicationStore `json:"store"`
}

type baseLinkApplicationStore struct {
	ID   float64 `json:"id"`
	Name string  `json:"name"`
}

type baseLinkButton struct {
	Action baseLinkButtonAction `json:"action"`
	Title  string               `json:"title"`
}

type baseLinkButtonAction struct {
	Type string `json:"type"`
	URL  string `json:"url"`
}

type baseLinkProduct struct {
	Price marketPrice `json:"price"`
}

type baseLinkRating struct {
	ReviewsCount int     `json:"reviews_count"`
	Stars        float64 `json:"stars"`
}

type basePlace struct {
	Address   string  `json:"address"`
	Checkins  int     `json:"checkins"`
	City      string  `json:"city"`
	Country   string  `json:"country"`
	Created   int     `json:"created"`
	ID        int     `json:"id"`
	Icon      string  `json:"icon"`
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
	Title     string  `json:"title"`
	Type      string  `json:"type"`
}

type baseRepostsInfo struct {
	Count        int `json:"count"`
	UserReposted int `json:"user_reposted"`
}

type baseSticker struct {
	Images               []baseImage `json:"images"`
	ImagesWithBackground []baseImage `json:"images_with_background"`
	ProductID            int         `json:"product_id"`
	StickerID            int         `json:"sticker_id"`
}

type baseUserID struct {
	UserID int `json:"user_id"`
}

type eventsEventAttach struct {
	Address      string `json:"address,omitempty"`       // address of event
	ButtonText   string `json:"button_text"`             // text of attach
	Friends      []int  `json:"friends"`                 // array of friends ids
	ID           int    `json:"id"`                      // event ID
	IsFavorite   bool   `json:"is_favorite"`             // is favorite
	MemberStatus int    `json:"member_status,omitempty"` // Current user's member status
	Text         string `json:"text"`                    // text of attach
	Time         int    `json:"time,omitempty"`          // event start time
}

// OauthError struct
type OauthError struct {
	Error            string `json:"error"`
	ErrorDescription string `json:"error_description"`
	RedirectURI      string `json:"redirect_uri"`
}

// Article struct
type Article struct {
	ID            int         `json:"id"`
	OwnerID       int         `json:"owner_id"`
	OwnerName     string      `json:"owner_name"`
	OwnerPhoto    string      `json:"owner_photo"`
	State         string      `json:"state"`
	CanReport     bool        `json:"can_report"`
	IsFavorite    bool        `json:"is_favorite"`
	Title         string      `json:"title"`
	Subtitle      string      `json:"subtitle"`
	Views         int         `json:"views"`
	Shares        int         `json:"shares"`
	URL           string      `json:"url"`
	ViewURL       string      `json:"view_url"`
	AccessKey     string      `json:"access_key"`
	PublishedDate int         `json:"published_date"`
	Photo         PhotosPhoto `json:"photo"`
}
