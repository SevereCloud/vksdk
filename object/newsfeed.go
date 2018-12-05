package object

type newsfeedItemAudio struct {
	Audio  newsfeedItemAudioAudio `json:"audio"`
	PostID int                    `json:"post_id"`
}

type newsfeedItemAudioAudio struct {
	Count int              `json:"count"`
	Items []AudioAudioFull `json:"items"`
}

type newsfeedItemFriend struct {
	Friends newsfeedItemFriendFriends `json:"friends"`
}

type newsfeedItemFriendFriends struct {
	Count int          `json:"count"`
	Items []baseUserID `json:"items"`
}

type newsfeedItemNote struct {
	Notes newsfeedItemNoteNotes `json:"notes"`
}

type newsfeedItemNoteNotes struct {
	Count int                    `json:"count"`
	Items []newsfeedNewsfeedNote `json:"items"`
}

type newsfeedItemPhoto struct {
	Photos newsfeedItemPhotoPhotos `json:"photos"`
	PostID int                     `json:"post_id"`
}

type newsfeedItemPhotoPhotos struct {
	Count int                     `json:"count"`
	Items []newsfeedNewsfeedPhoto `json:"items"`
}

type newsfeedItemPhotoTag struct {
	PhotoTags newsfeedItemPhotoTagPhotoTags `json:"photo_tags"`
	PostID    int                           `json:"post_id"`
}

type newsfeedItemPhotoTagPhotoTags struct {
	Count int                     `json:"count"`
	Items []newsfeedNewsfeedPhoto `json:"items"`
}

type newsfeedItemTopic struct {
	Comments baseCommentsInfo `json:"comments"`
	Likes    baseLikesInfo    `json:"likes"`
	PostID   int              `json:"post_id"`
	Text     string           `json:"text"`
}

type newsfeedItemVideo struct {
	Video newsfeedItemVideoVideo `json:"video"`
}

type newsfeedItemVideoVideo struct {
	Count int          `json:"count"`
	Items []VideoVideo `json:"items"`
}

type newsfeedItemWallpost struct {
	Attachments []wallWallpostAttachment `json:"attachments"`
	Comments    baseCommentsInfo         `json:"comments"`
	CopyHistory []wallWallpost           `json:"copy_history"`
	Geo         baseGeo                  `json:"geo"`
	Likes       baseLikesInfo            `json:"likes"`
	PostID      int                      `json:"post_id"`
	PostSource  wallPostSource           `json:"post_source"`
	PostType    string                   `json:"post_type"`
	Reposts     baseRepostsInfo          `json:"reposts"`
	Text        string                   `json:"text"`
}

type newsfeedList struct {
	ID    int    `json:"id"`
	Title string `json:"title"`
}

type newsfeedNewsfeedItemType string

type newsfeedNewsfeedNote struct {
	Comments int    `json:"comments"`
	ID       int    `json:"id"`
	OwnerID  int    `json:"owner_id"`
	Title    string `json:"title"`
}

type newsfeedNewsfeedPhoto struct {
}
