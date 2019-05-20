package object // import "github.com/SevereCloud/vksdk/5.92/object"

type wallAppPost struct {
	ID       int    `json:"id"`        // Application ID
	Name     string `json:"name"`      // Application name
	Photo130 string `json:"photo_130"` // URL of the preview image with 130 px in width
	Photo604 string `json:"photo_604"` // URL of the preview image with 604 px in width
}

type wallAttachedNote struct {
	Comments     int    `json:"comments"`      // Comments number
	Date         int    `json:"date"`          // Date when the note has been created in Unixtime
	ID           int    `json:"id"`            // Note ID
	OwnerID      int    `json:"owner_id"`      // Note owner's ID
	ReadComments int    `json:"read_comments"` // Read comments number
	Title        string `json:"title"`         // Note title
	ViewURL      string `json:"view_url"`      // URL of the page with note preview
}

type wallCommentAttachment struct {
	Audio             AudioAudioFull    `json:"audio"`
	Doc               DocsDoc           `json:"doc"`
	Link              baseLink          `json:"link"`
	Market            MarketMarketItem  `json:"market"`
	MarketMarketAlbum MarketMarketAlbum `json:"market_market_album"`
	Note              wallAttachedNote  `json:"note"`
	Page              PagesWikipageFull `json:"page"`
	Photo             PhotosPhoto       `json:"photo"`
	Sticker           baseSticker       `json:"sticker"`
	Type              string            `json:"type"`
	Video             VideoVideo        `json:"video"`
}

type wallGraffiti struct {
	ID       int    `json:"id"`        // Graffiti ID
	OwnerID  int    `json:"owner_id"`  // Graffiti owner's ID
	Photo200 string `json:"photo_200"` // URL of the preview image with 200 px in width
	Photo586 string `json:"photo_586"` // URL of the preview image with 586 px in width
}

type wallPostSource struct {
	Data     string `json:"data"`     // Additional data
	Platform string `json:"platform"` // Platform name
	Type     string `json:"type"`
	URL      string `json:"url"` // URL to an external site used to publish the post
}

type wallPostedPhoto struct {
	ID       int    `json:"id"`        // Photo ID
	OwnerID  int    `json:"owner_id"`  // Photo owner's ID
	Photo130 string `json:"photo_130"` // URL of the preview image with 130 px in width
	Photo604 string `json:"photo_604"` // URL of the preview image with 604 px in width
}

type wallViews struct {
	Count int `json:"count"` // Count
}

type wallWallCommentThread struct {
	Count           int               `json:"count"` // Comments number
	Items           []WallWallComment `json:"items"`
	CanPost         bool              `json:"can_post"`        // Information whether current user can comment the post
	GroupsCanPost   bool              `json:"groups_can_post"` // Information whether groups can comment the post
	ShowReplyButton bool              `json:"show_reply_button"`
}

// WallWallComment struct
type WallWallComment struct {
	Attachments    []wallCommentAttachment `json:"attachments"`
	Date           int                     `json:"date"` // Date when the comment has been added in Unixtime
	Deleted        bool                    `json:"deleted"`
	FromID         int                     `json:"from_id"` // Author ID
	ID             int                     `json:"id"`      // Comment ID
	Likes          baseLikesInfo           `json:"likes"`
	RealOffset     int                     `json:"real_offset"`      // Real position of the comment
	ReplyToComment int                     `json:"reply_to_comment"` // Replied comment ID
	ReplyToUser    int                     `json:"reply_to_user"`    // Replied user ID
	Text           string                  `json:"text"`             // Comment text
	PostID         int                     `json:"post_id"`
	PostOwnerID    int                     `json:"post_owner_id"`
	PhotoID        int                     `json:"photo_id"`
	PhotoOwnerID   int                     `json:"photo_owner_id"`
	VideoID        int                     `json:"video_id"`
	VideoOwnerID   int                     `json:"video_owner_id"`
	ItemID         int                     `json:"item_id"`
	MarketOwnerID  int                     `json:"market_owner_id"`
	ParentsStack   []int                   `json:"parents_stack"`
	OwnerID        int                     `json:"owner_id"`
	Thread         wallWallCommentThread   `json:"thread"`
}

// WallWallpost  struct
type WallWallpost struct {
	AccessKey    string                   `json:"access_key"` // Access key to private object
	ID           int                      `json:"id"`         // Post ID
	OwnerID      int                      `json:"owner_id"`   // Wall owner's ID
	FromID       int                      `json:"from_id"`    // Post author ID
	CreatedBy    int                      `json:"created_by"`
	Date         int                      `json:"date"` // Date of publishing in Unixtime
	Text         string                   `json:"text"` // Post text
	ReplyOwnerID int                      `json:"reply_owner_id"`
	ReplyPostID  int                      `json:"reply_post_id"`
	FriendsOnly  int                      `json:"friends_only"`
	Comments     baseCommentsInfo         `json:"comments"`
	Likes        baseLikesInfo            `json:"likes"`   // Count of likes
	Reposts      baseRepostsInfo          `json:"reposts"` // Count of views
	Views        wallViews                `json:"views"`   // Count of views
	PostType     string                   `json:"post_type"`
	PostSource   wallPostSource           `json:"post_source"`
	Attachments  []wallWallpostAttachment `json:"attachments"`
	Geo          baseGeo                  `json:"geo"`
	SignerID     int                      `json:"signer_id"` // Post signer ID
	CopyHistory  []WallWallpost           `json:"copy_history"`
	CanPin       int                      `json:"can_pin"`
	CanDelete    int                      `json:"can_delete"`
	CanEdit      int                      `json:"can_edit"`
	IsPinned     int                      `json:"is_pinned"`
	MarkedAsAds  int                      `json:"marked_as_ads"`
	Edited       int                      `json:"edited"`      // Date of editing in Unixtime
	IsFavorite   bool                     `json:"is_favorite"` // Information whether the post in favorites list
	IsArchived   bool                     `json:"is_archived"` // Is post archived, only for post owners
}

type wallWallpostAttached struct {
	Attachments []wallWallpostAttachment `json:"attachments"`
	CanDelete   int                      `json:"can_delete"`
	Comments    baseCommentsInfo         `json:"comments"`
	CopyOwnerID int                      `json:"copy_owner_id"`
	CopyPostID  int                      `json:"copy_post_id"`
	CopyText    string                   `json:"copy_text"`
	Date        int                      `json:"date"`
	FromID      int                      `json:"from_id"`
	Geo         baseGeo                  `json:"geo"`
	ID          int                      `json:"id"`
	Likes       baseLikesInfo            `json:"likes"`
	PostSource  wallPostSource           `json:"post_source"`
	PostType    string                   `json:"post_type"`
	Reposts     baseRepostsInfo          `json:"reposts"`
	SignerID    int                      `json:"signer_id"`
	Text        string                   `json:"text"`
	ToID        int                      `json:"to_id"`
}

type wallWallpostAttachment struct {
	AccessKey         string            `json:"access_key"` // Access key for the audio
	Album             photosPhotoAlbum  `json:"album"`
	App               wallAppPost       `json:"app"`
	Audio             AudioAudioFull    `json:"audio"`
	Doc               DocsDoc           `json:"doc"`
	Event             eventsEventAttach `json:"event"`
	Graffiti          wallGraffiti      `json:"graffiti"`
	Link              baseLink          `json:"link"`
	Market            MarketMarketItem  `json:"market"`
	MarketMarketAlbum MarketMarketAlbum `json:"market_market_album"`
	Note              wallAttachedNote  `json:"note"`
	Page              PagesWikipageFull `json:"page"`
	Photo             PhotosPhoto       `json:"photo"`
	PhotosList        []string          `json:"photos_list"`
	Poll              PollsPoll         `json:"poll"`
	PostedPhoto       wallPostedPhoto   `json:"posted_photo"`
	Type              string            `json:"type"`
	Video             VideoVideo        `json:"video"`
}

// WallWallpostToID struct
type WallWallpostToID struct {
	Attachments []wallWallpostAttachment `json:"attachments"`
	Comments    baseCommentsInfo         `json:"comments"`
	CopyOwnerID int                      `json:"copy_owner_id"` // ID of the source post owner
	CopyPostID  int                      `json:"copy_post_id"`  // ID of the source post
	Date        int                      `json:"date"`          // Date of publishing in Unixtime
	FromID      int                      `json:"from_id"`       // Post author ID
	Geo         baseGeo                  `json:"geo"`
	ID          int                      `json:"id"` // Post ID
	Likes       baseLikesInfo            `json:"likes"`
	PostID      int                      `json:"post_id"` // wall post ID (if comment)
	PostSource  wallPostSource           `json:"post_source"`
	PostType    string                   `json:"post_type"`
	Reposts     baseRepostsInfo          `json:"reposts"`
	SignerID    int                      `json:"signer_id"` // Post signer ID
	Text        string                   `json:"text"`      // Post text
	ToID        int                      `json:"to_id"`     // Wall owner's ID
}
