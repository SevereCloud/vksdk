package object // import "github.com/SevereCloud/vksdk/object"

// StoriesStoryLink struct
type StoriesStoryLink struct {
	Text string `json:"text"` // Link text
	URL  string `json:"url"`  // Link URL
}

// StoriesReplies struct
type StoriesReplies struct {
	Count int `json:"count"` // Replies number.
	New   int `json:"new"`   // New replies number.
}

// StoriesStoryStats struct
type StoriesStoryStats struct {
	Answer      StoriesStoryStatsStat `json:"answer"`
	Bans        StoriesStoryStatsStat `json:"bans"`
	OpenLink    StoriesStoryStatsStat `json:"open_link"`
	Replies     StoriesStoryStatsStat `json:"replies"`
	Shares      StoriesStoryStatsStat `json:"shares"`
	Subscribers StoriesStoryStatsStat `json:"subscribers"`
	Views       StoriesStoryStatsStat `json:"views"`
}

// StoriesStoryStatsStat struct
type StoriesStoryStatsStat struct {
	Count int    `json:"count"` // Stat value
	State string `json:"state"`
}

// StoriesStory struct
type StoriesStory struct {
	AccessKey string `json:"access_key"` // Access key for private object.
	ExpiresAt int    `json:"expires_at"` // Story expiration time. Unixtime.
	CanHide   int    `json:"can_hide"`
	// Information whether story has question sticker and current user can send question to the author
	CanAsk int `json:"can_ask"`
	// Information whether story has question sticker and current user can send anonymous question to the author
	CanAskAnonymous      int                      `json:"can_ask_anonymous"`
	CanComment           int                      `json:"can_comment"`   // Information whether current user can comment the story (0 - no, 1 - yes).
	CanReply             int                      `json:"can_reply"`     // Information whether current user can reply to the story (0 - no, 1 - yes).
	CanSee               int                      `json:"can_see"`       // Information whether current user can see the story (0 - no, 1 - yes).
	CanShare             int                      `json:"can_share"`     // Information whether current user can share the story (0 - no, 1 - yes).
	Date                 int                      `json:"date"`          // Date when story has been added in Unixtime.
	ID                   int                      `json:"id"`            // Story ID.
	IsDeleted            bool                     `json:"is_deleted"`    // Information whether the story is deleted (false - no, true - yes).
	IsExpired            bool                     `json:"is_expired"`    // Information whether the story is expired (false - no, true - yes).
	NoSound              bool                     `json:"no_sound"`      // Is video without sound
	IsRestricted         bool                     `json:"is_restricted"` // Does author have stories privacy restrictions
	Link                 StoriesStoryLink         `json:"link"`
	OwnerID              int                      `json:"owner_id"` // Story owner's ID.
	ParentStory          *StoriesStory            `json:"parent_story"`
	ParentStoryAccessKey string                   `json:"parent_story_access_key"` // Access key for private object.
	ParentStoryID        int                      `json:"parent_story_id"`         // Parent story ID.
	ParentStoryOwnerID   int                      `json:"parent_story_owner_id"`   // Parent story owner's ID.
	Photo                PhotosPhoto              `json:"photo"`
	Replies              StoriesReplies           `json:"replies"` // Replies to current story.
	Seen                 int                      `json:"seen"`    // Information whether current user has seen the story or not (0 - no, 1 - yes).
	Type                 string                   `json:"type"`
	Video                VideoVideo               `json:"video"`
	Views                int                      `json:"views"` // Views number.
	ClickableStickers    StoriesClickableStickers `json:"clickable_stickers"`
}

// StoriesClickableStickers struct
//
// https://vk.com/dev/objects/clickable_stickers
type StoriesClickableStickers struct {
	OriginalHeight    int                       `json:"original_height"`
	OriginalWidth     int                       `json:"original_width"`
	ClickableStickers []StoriesClickableSticker `json:"clickable_stickers"`
}

// StoriesClickableSticker struct
type StoriesClickableSticker struct {
	ID            int                     `json:"id"`
	Type          string                  `json:"type"`
	ClickableArea []StoriesClickablePoint `json:"clickable_area"`
	Style         string                  `json:"style"`
	Subtype       string                  `json:"subtype"`

	// type=post
	PostOwnerID int `json:"post_owner_id"`
	PostID      int `json:"post_id"`

	// type=sticker
	StickerID     int `json:"sticker_id"`
	StickerPackID int `json:"sticker_pack_id"`

	// type=place
	PlaceID int `json:"place_id"`

	// type=question
	Question               string `json:"question"`
	QuestionButton         string `json:"question_button"`
	QuestionDefaultPrivate bool   `json:"question_default_private"`
	Color                  string `json:"color"`

	// type=mention
	Mention string `json:"mention,omitempty"`

	// type=hashtag
	Hashtag string `json:"hashtag,omitempty"`

	// type=market_item
	MarketItem MarketMarketItem `json:"market_item"`
}

// StoriesClickablePoint struct
type StoriesClickablePoint struct {
	X int `json:"x"`
	Y int `json:"y"`
}
