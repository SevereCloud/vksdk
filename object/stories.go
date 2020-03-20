package object // import "github.com/SevereCloud/vksdk/object"

// StoriesNarrativeInfo type
type StoriesNarrativeInfo struct {
	Author string `json:"author"`
	Title  string `json:"title"`
	Views  int    `json:"views"`
}

// StoriesPromoData struct
type StoriesPromoData struct {
	Name        string      `json:"name"`
	Photo50     string      `json:"photo_50"`
	Photo100    string      `json:"photo_100"`
	NotAnimated BaseBoolInt `json:"not_animated"`
}

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
	Likes       StoriesStoryStatsStat `json:"likes"`
}

// StoriesStoryStatsStat struct
type StoriesStoryStatsStat struct {
	Count int    `json:"count"` // Stat value
	State string `json:"state"`
}

// StoriesStory struct
type StoriesStory struct {
	AccessKey string      `json:"access_key"` // Access key for private object.
	ExpiresAt int         `json:"expires_at"` // Story expiration time. Unixtime.
	CanHide   BaseBoolInt `json:"can_hide"`
	// Information whether story has question sticker and current user can send question to the author
	CanAsk BaseBoolInt `json:"can_ask"`
	// Information whether story has question sticker and current user can send anonymous question to the author
	CanAskAnonymous      BaseBoolInt              `json:"can_ask_anonymous"`
	CanComment           BaseBoolInt              `json:"can_comment"`   // Information whether current user can comment the story (0 - no, 1 - yes).
	CanReply             BaseBoolInt              `json:"can_reply"`     // Information whether current user can reply to the story (0 - no, 1 - yes).
	CanSee               BaseBoolInt              `json:"can_see"`       // Information whether current user can see the story (0 - no, 1 - yes).
	CanShare             BaseBoolInt              `json:"can_share"`     // Information whether current user can share the story (0 - no, 1 - yes).
	Date                 int                      `json:"date"`          // Date when story has been added in Unixtime.
	ID                   int                      `json:"id"`            // Story ID.
	IsDeleted            BaseBoolInt              `json:"is_deleted"`    // Information whether the story is deleted (false - no, true - yes).
	IsExpired            BaseBoolInt              `json:"is_expired"`    // Information whether the story is expired (false - no, true - yes).
	NoSound              BaseBoolInt              `json:"no_sound"`      // Is video without sound
	IsRestricted         BaseBoolInt              `json:"is_restricted"` // Does author have stories privacy restrictions
	Seen                 BaseBoolInt              `json:"seen"`          // Information whether current user has seen the story or not (0 - no, 1 - yes).
	IsOwnerPinned        BaseBoolInt              `json:"is_owner_pinned"`
	Link                 StoriesStoryLink         `json:"link"`
	OwnerID              int                      `json:"owner_id"` // Story owner's ID.
	ParentStory          *StoriesStory            `json:"parent_story"`
	ParentStoryAccessKey string                   `json:"parent_story_access_key"` // Access key for private object.
	ParentStoryID        int                      `json:"parent_story_id"`         // Parent story ID.
	ParentStoryOwnerID   int                      `json:"parent_story_owner_id"`   // Parent story owner's ID.
	Photo                PhotosPhoto              `json:"photo"`
	Replies              StoriesReplies           `json:"replies"` // Replies to current story.
	Type                 string                   `json:"type"`
	Video                VideoVideo               `json:"video"`
	Views                int                      `json:"views"` // Views number.
	ClickableStickers    StoriesClickableStickers `json:"clickable_stickers"`
	TrackCode            string                   `json:"track_code"`
	LikesCount           int                      `json:"likes_count"`
	NarrativeID          int                      `json:"narrative_id"`
	NarrativeOwnerID     int                      `json:"narrative_owner_id"`
	NarrativeInfo        StoriesNarrativeInfo     `json:"narrative_info"`
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

	// type=post
	PostOwnerID int `json:"post_owner_id"`
	PostID      int `json:"post_id"`

	// type=sticker
	StickerID     int `json:"sticker_id"`
	StickerPackID int `json:"sticker_pack_id"`

	// type=place
	PlaceID int `json:"place_id"`

	// type=question
	Question               string      `json:"question"`
	QuestionButton         string      `json:"question_button"`
	QuestionDefaultPrivate BaseBoolInt `json:"question_default_private"`
	Color                  string      `json:"color"`

	// type=mention
	Mention string `json:"mention,omitempty"`

	// type=hashtag
	Hashtag string `json:"hashtag,omitempty"`

	// type=market_item
	MarketItem MarketMarketItem       `json:"market_item"`
	Product    BaseLinkProduct        `json:"product"`
	Button     MessagesKeyboardButton `json:"button"`
	Rating     BaseLinkRating         `json:"rating"`
	Subtype    string                 `json:"subtype"`

	// type=link
	LinkObject  BaseLink `json:"link_object"`
	TooltipText string   `json:"tooltip_text"`

	// type=story_reply
	OwnerID int `json:"owner_id"`
	StoryID int `json:"story_id"`

	// type=owner
	// OwnerID int `json:"owner_id"`

	// type=poll
	Poll PollsPoll `json:"poll"`

	// type=music
	Audio          AudioAudioFull `json:"audio"`
	AudioStartTime int            `json:"audio_start_time"`
}

// StoriesClickableStickerType type of clickable sticker
type StoriesClickableStickerType string

// ClickableSticker type
// FIXME: v2 StoriesClickableStickerType
const (
	StoriesClickableStickerPost       = "post"
	StoriesClickableStickerSticker    = "sticker"
	StoriesClickableStickerPlace      = "place"
	StoriesClickableStickerQuestion   = "question"
	StoriesClickableStickerMention    = "mention"
	StoriesClickableStickerHashtag    = "hashtag"
	StoriesClickableStickerMarketItem = "market_item"
	StoriesClickableStickerLink       = "link"
	StoriesClickableStickerStoryReply = "story_reply"
	StoriesClickableStickerOwner      = "owner"
	StoriesClickableStickerPoll       = "poll"
	StoriesClickableStickerMusic      = "music"
)

// StoriesClickablePoint struct
type StoriesClickablePoint struct {
	X int `json:"x"`
	Y int `json:"y"`
}
