package object // import "github.com/SevereCloud/vksdk/object"

type newsfeedEventActivity struct {
	Address      string `json:"address"`       // address of event
	ButtonText   string `json:"button_text"`   // text of attach
	Friends      []int  `json:"friends"`       // array of friends ids
	MemberStatus int    `json:"member_status"` // Current user's member status
	Text         string `json:"text"`          // text of attach
	Time         int    `json:"time"`          // event start time
}

// NewsfeedItemAudio struct
type NewsfeedItemAudio struct {
	Audio newsfeedItemAudioAudio `json:"audio"`
}

type newsfeedItemAudioAudio struct {
	Count int              `json:"count"` // Audios number
	Items []AudioAudioFull `json:"items"`
}

// NewsfeedItemDigest struct
type NewsfeedItemDigest struct {
	ButtonText  string         `json:"button_text"`
	FeedID      string         `json:"feed_id"` // id of feed in digest
	Items       []WallWallpost `json:"items"`
	MainPostIDs []string       `json:"main_post_ids"`
	Template    string         `json:"template"` // type of digest
	Title       string         `json:"title"`
	TrackCode   string         `json:"track_code"`
	// Type        string         `json:"type"`
}

// NewsfeedItemFriend struct
type NewsfeedItemFriend struct {
	Friends newsfeedItemFriendFriends `json:"friends"`
}

type newsfeedItemFriendFriends struct {
	Count int          `json:"count"` // Number of friends has been added
	Items []baseUserID `json:"items"`
}

// NewsfeedItemNote struct
type NewsfeedItemNote struct {
	Notes newsfeedItemNoteNotes `json:"notes"`
}

type newsfeedItemNoteNotes struct {
	Count int                    `json:"count"` // Notes number
	Items []newsfeedNewsfeedNote `json:"items"`
}

// NewsfeedItemPhoto struct
type NewsfeedItemPhoto struct {
	Photos newsfeedItemPhotoPhotos `json:"photos"`
}

type newsfeedItemPhotoPhotos struct {
	Count int               `json:"count"` // Photos number
	Items []PhotosPhotoFull `json:"items"`
}

// NewsfeedItemPhotoTag struct
type NewsfeedItemPhotoTag struct {
	PhotoTags newsfeedItemPhotoTagPhotoTags `json:"photo_tags"`
}

type newsfeedItemPhotoTagPhotoTags struct {
	Count int               `json:"count"` // Tags number
	Items []PhotosPhotoFull `json:"items"`
}

// NewsfeedItemStoriesBlock struct
type NewsfeedItemStoriesBlock struct {
	BlockType string         `json:"block_type"`
	Stories   []StoriesStory `json:"stories"`
	// Title     string         `json:"title"`
	// TrackCode string `json:"track_code"`
	// Type      string         `json:"type"`
}

// NewsfeedItemTopic struct
type NewsfeedItemTopic struct {
	// Comments baseCommentsInfo `json:"comments"`
	// Likes baseLikesInfo `json:"likes"`
	// Text  string        `json:"text"` // Post text
}

// NewsfeedItemVideo struct
type NewsfeedItemVideo struct {
	Video newsfeedItemVideoVideo `json:"video"`
}

type newsfeedItemVideoVideo struct {
	Count int          `json:"count"` // Tags number
	Items []VideoVideo `json:"items"`
}

// NewsfeedItemWallpost struct
type NewsfeedItemWallpost struct {
	Activity    newsfeedEventActivity    `json:"activity"`
	Attachments []wallWallpostAttachment `json:"attachments"`
	Comments    baseCommentsInfo         `json:"comments"`
	CopyHistory []WallWallpost           `json:"copy_history"`
	Geo         baseGeo                  `json:"geo"`
	Likes       baseLikesInfo            `json:"likes"`
	PostSource  wallPostSource           `json:"post_source"`
	PostType    string                   `json:"post_type"`
	Reposts     baseRepostsInfo          `json:"reposts"`
	MarkedAsAds int                      `json:"marked_as_ads,omitempty"`
	Views       int                      `json:"views,omitempty"`
	IsFavorite  bool                     `json:"is_favorite,omitempty"`
	SignerID    int                      `json:"signer_id,omitempty"`
	Text        string                   `json:"text"` // Post text
}

// NewsfeedList struct
type NewsfeedList struct {
	ID    int    `json:"id"`    // List ID
	Title string `json:"title"` // List title
}

// type newsfeedListFull struct {
// }

// NewsfeedNewsfeedItem struct
type NewsfeedNewsfeedItem struct {
	Type     string `json:"type"`
	SourceID int    `json:"source_id"`
	Date     int    `json:"date"`

	PostID int `json:"post_id,omitempty"`

	NewsfeedItemWallpost
	NewsfeedItemPhoto
	NewsfeedItemPhotoTag
	NewsfeedItemFriend
	NewsfeedItemNote
	NewsfeedItemAudio
	NewsfeedItemTopic
	NewsfeedItemVideo
	NewsfeedItemDigest
	NewsfeedItemStoriesBlock

	CanEdit   int `json:"can_edit,omitempty"`
	CreatedBy int `json:"created_by,omitempty"`
	CanDelete int `json:"can_delete,omitempty"`
	// TODO: Need more fields
}

type newsfeedNewsfeedNote struct {
	Comments int    `json:"comments"` // Comments Number
	ID       int    `json:"id"`       // Note ID
	OwnerID  int    `json:"owner_id"` // integer
	Title    string `json:"title"`    // Note title
}
