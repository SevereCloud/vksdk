package object

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
}

type baseCity struct {
	ID    int    `json:"id"`
	Title string `json:"title"`
}

type baseCommentsInfo struct {
	CanPost       int `json:"can_post"`
	Count         int `json:"count"`
	GroupsCanPost int `json:"groups_can_post"`
}

type baseError struct {
	ErrorCode     int                `json:"error_code"`
	ErrorMsg      string             `json:"error_msg"`
	RequestParams []BaseRequestParam `json:"request_params"`
}

type baseGeo struct {
	Coordinates baseGeoCoordinates `json:"coordinates"`
	Place       basePlace          `json:"place"`
	Showmap     int                `json:"showmap"`
	Type        string             `json:"type"`
}

type baseGeoCoordinates struct {
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
}

type baseImage struct {
	Height int    `json:"height"`
	URL    string `json:"url"`
	Width  int    `json:"width"`
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

type baseUploadServer struct {
	UploadURL string `json:"upload_url"`
}

type baseUserID struct {
	UserID int `json:"user_id"`
}

type faveFavesLink struct {
	Description string `json:"description"`
	ID          int    `json:"id"`
	Photo100    string `json:"photo_100"`
	Photo200    string `json:"photo_200"`
	Photo50     string `json:"photo_50"`
	Title       string `json:"title"`
	URL         string `json:"url"`
}

type oauthError struct {
	Error            string `json:"error"`
	ErrorDescription string `json:"error_description"`
	RedirectURI      string `json:"redirect_uri"`
}

type searchHint struct {
	Description string       `json:"description"`
	Global      int          `json:"global"`
	Group       GroupsGroup  `json:"group"`
	Profile     UsersUserMin `json:"profile"`
	Section     string       `json:"section"`
	Type        string       `json:"type"`
}
