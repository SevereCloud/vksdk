package object // import "github.com/SevereCloud/vksdk/5.92/object"

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
	Audio  newsfeedItemAudioAudio `json:"audio"`
	PostID int                    `json:"post_id"` // Post ID
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
	Type        string         `json:"type"` // type of digest
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
	PostID int                     `json:"post_id"` // Post ID
}

type newsfeedItemPhotoPhotos struct {
	Count int                     `json:"count"` // Photos number
	Items []newsfeedNewsfeedPhoto `json:"items"`
}

// NewsfeedItemPhotoTag struct
type NewsfeedItemPhotoTag struct {
	PhotoTags newsfeedItemPhotoTagPhotoTags `json:"photo_tags"`
	PostID    int                           `json:"post_id"` // Post ID
}

type newsfeedItemPhotoTagPhotoTags struct {
	Count int                     `json:"count"` // Tags number
	Items []newsfeedNewsfeedPhoto `json:"items"`
}

// NewsfeedItemStoriesBlock struct
type NewsfeedItemStoriesBlock struct {
	BlockType string         `json:"block_type"`
	Stories   []StoriesStory `json:"stories"`
	Title     string         `json:"title"`
	TrackCode string         `json:"track_code"`
	Type      string         `json:"type"`
}

// NewsfeedItemTopic struct
type NewsfeedItemTopic struct {
	Comments baseCommentsInfo `json:"comments"`
	Likes    baseLikesInfo    `json:"likes"`
	PostID   int              `json:"post_id"` // Topic post ID
	Text     string           `json:"text"`    // Post text
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
	PostID      int                      `json:"post_id"` // Post ID
	PostSource  wallPostSource           `json:"post_source"`
	PostType    string                   `json:"post_type"`
	Reposts     baseRepostsInfo          `json:"reposts"`
	Text        string                   `json:"text"` // Post text
}

// NewsfeedList struct
type NewsfeedList struct {
	ID    int    `json:"id"`    // List ID
	Title string `json:"title"` // List title
}

// type newsfeedListFull struct {
// }

// type newsfeedNewsfeedItem struct {
// }

type newsfeedNewsfeedNote struct {
	Comments int    `json:"comments"` // Comments Number
	ID       int    `json:"id"`       // Note ID
	OwnerID  int    `json:"owner_id"` // integer
	Title    string `json:"title"`    // Note title
}

type newsfeedNewsfeedPhoto struct {
}
