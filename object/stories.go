package object // import "github.com/SevereCloud/vksdk/object"

// storiesStoryLink struct
type storiesStoryLink struct {
	Text string `json:"text"` // Link text
	URL  string `json:"url"`  // Link URL
}

type storiesReplies struct {
	Count int `json:"count"` // Replies number.
	New   int `json:"new"`   // New replies number.
}

// StoriesStoryStats struct
type StoriesStoryStats struct {
	Answer      storiesStoryStatsStat `json:"answer"`
	Bans        storiesStoryStatsStat `json:"bans"`
	OpenLink    storiesStoryStatsStat `json:"open_link"`
	Replies     storiesStoryStatsStat `json:"replies"`
	Shares      storiesStoryStatsStat `json:"shares"`
	Subscribers storiesStoryStatsStat `json:"subscribers"`
	Views       storiesStoryStatsStat `json:"views"`
}

type storiesStoryStatsStat struct {
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
	CanAskAnonymous      int              `json:"can_ask_anonymous"`
	CanComment           int              `json:"can_comment"`   // Information whether current user can comment the story (0 - no, 1 - yes).
	CanReply             int              `json:"can_reply"`     // Information whether current user can reply to the story (0 - no, 1 - yes).
	CanSee               int              `json:"can_see"`       // Information whether current user can see the story (0 - no, 1 - yes).
	CanShare             int              `json:"can_share"`     // Information whether current user can share the story (0 - no, 1 - yes).
	Date                 int              `json:"date"`          // Date when story has been added in Unixtime.
	ID                   int              `json:"id"`            // Story ID.
	IsDeleted            bool             `json:"is_deleted"`    // Information whether the story is deleted (false - no, true - yes).
	IsExpired            bool             `json:"is_expired"`    // Information whether the story is expired (false - no, true - yes).
	NoSound              bool             `json:"no_sound"`      // Is video without sound
	IsRestricted         bool             `json:"is_restricted"` // Does author have stories privacy restrictions
	Link                 storiesStoryLink `json:"link"`
	OwnerID              int              `json:"owner_id"` // Story owner's ID.
	ParentStory          *StoriesStory    `json:"parent_story"`
	ParentStoryAccessKey string           `json:"parent_story_access_key"` // Access key for private object.
	ParentStoryID        int              `json:"parent_story_id"`         // Parent story ID.
	ParentStoryOwnerID   int              `json:"parent_story_owner_id"`   // Parent story owner's ID.
	Photo                PhotosPhoto      `json:"photo"`
	Replies              storiesReplies   `json:"replies"` // Replies to current story.
	Seen                 int              `json:"seen"`    // Information whether current user has seen the story or not (0 - no, 1 - yes).
	Type                 string           `json:"type"`
	Video                VideoVideo       `json:"video"`
	Views                int              `json:"views"` // Views number.
}
