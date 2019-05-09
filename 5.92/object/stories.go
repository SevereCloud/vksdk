package object // import "github.com/SevereCloud/vksdk/5.92/object"

// storiesStoryLink struct
type storiesStoryLink struct {

	// Link text
	Text string `json:"text"`

	// Link URL
	URL string `json:"url"`
}

// storiesStoryVideo struct
type storiesStoryVideo struct {
}

type storiesReplies struct {

	// Replies number.
	Count int `json:"count"`

	// New replies number.
	New int `json:"new,omitempty"`
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
	Count int    `json:"count"`
	State string `json:"state"`
}

// StoriesStory struct
type StoriesStory struct {
	// Access key for private object.
	AccessKey string `json:"access_key,omitempty"`

	// Information whether current user can comment the story (0 - no, 1 - yes).
	CanComment int `json:"can_comment,omitempty"`

	// Information whether current user can reply to the story (0 - no, 1 - yes).
	CanReply int `json:"can_reply,omitempty"`

	// Information whether current user can see the story (0 - no, 1 - yes).
	CanSee int `json:"can_see,omitempty"`

	// Information whether current user can share the story (0 - no, 1 - yes).
	CanShare int `json:"can_share,omitempty"`

	// Date when story has been added in Unixtime.
	Date int `json:"date,omitempty"`

	// Story ID.
	ID int `json:"id"`

	// Information whether the story is deleted (false - no, true - yes).
	IsDeleted bool `json:"is_deleted,omitempty"`

	// Information whether the story is expired (false - no, true - yes).
	IsExpired bool              `json:"is_expired,omitempty"`
	Link      *storiesStoryLink `json:"link,omitempty"`

	// Story owner's ID.
	OwnerID     int           `json:"owner_id"`
	ParentStory *StoriesStory `json:"parent_story,omitempty"`

	// Access key for private object.
	ParentStoryAccessKey string `json:"parent_story_access_key,omitempty"`

	// Parent story ID.
	ParentStoryID int `json:"parent_story_id,omitempty"`

	// Parent story owner's ID.
	ParentStoryOwnerID int          `json:"parent_story_owner_id,omitempty"`
	Photo              *PhotosPhoto `json:"photo,omitempty"`

	// Replies to current story.
	Replies []*storiesReplies `json:"replies,omitempty"`

	// Information whether current user has seen the story or not (0 - no, 1 - yes).
	Seen  int                `json:"seen,omitempty"`
	Type  string             `json:"type,omitempty"`
	Video *storiesStoryVideo `json:"video,omitempty"`

	// Views number.
	Views int `json:"views,omitempty"`
}
