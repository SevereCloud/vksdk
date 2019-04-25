package object // import "github.com/severecloud/vksdk/5.92/object"

type wallAppPost struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Photo130 string `json:"photo_130"`
	Photo604 string `json:"photo_604"`
}

type wallAttachedNote struct {
	Comments     int    `json:"comments"`
	Date         int    `json:"date"`
	ID           int    `json:"id"`
	OwnerID      int    `json:"owner_id"`
	ReadComments int    `json:"read_comments"`
	Title        string `json:"title"`
	ViewURL      string `json:"view_url"`
}

type wallCommentAttachment struct {
	Audio             AudioAudioFull    `json:"audio"`
	Doc               docsDoc           `json:"doc"`
	Link              baseLink          `json:"link"`
	Market            marketMarketItem  `json:"market"`
	MarketMarketAlbum marketMarketAlbum `json:"market_market_album"`
	Note              wallAttachedNote  `json:"note"`
	Page              PagesWikipageFull `json:"page"`
	Photo             PhotosPhoto       `json:"photo"`
	Sticker           baseSticker       `json:"sticker"`
	Type              string            `json:"type"`
	Video             VideoVideo        `json:"video"`
}

type wallGraffiti struct {
	ID       int    `json:"id"`
	OwnerID  int    `json:"owner_id"`
	Photo200 string `json:"photo_200"`
	Photo586 string `json:"photo_586"`
}

type wallPostSource struct {
	Data     string `json:"data"`
	Platform string `json:"platform"`
	Type     string `json:"type"`
	URL      string `json:"url"`
}

type wallPostedPhoto struct {
	ID       int    `json:"id"`
	OwnerID  int    `json:"owner_id"`
	Photo130 string `json:"photo_130"`
	Photo604 string `json:"photo_604"`
}

type wallViews struct {
	Count int `json:"count"`
}

type wallWallCommentThread struct {
	Count           int               `json:"count"`
	Items           []wallWallComment `json:"items"`
	CanPost         bool              `json:"can_post"`
	ShowReplyButton bool              `json:"show_reply_button"`
	GroupsCanPost   bool              `json:"groups_can_post"`
}

type wallWallComment struct {
	Attachments    []wallCommentAttachment `json:"attachments"`
	Date           int                     `json:"date"`
	FromID         int                     `json:"from_id"`
	ID             int                     `json:"id"`
	Likes          baseLikesInfo           `json:"likes"`
	RealOffset     int                     `json:"real_offset"`
	ReplyToComment int                     `json:"reply_to_comment"`
	ReplyToUser    int                     `json:"reply_to_user"`
	Text           string                  `json:"text"`
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
	ID           int                      `json:"id"`
	OwnerID      int                      `json:"owner_id"`
	FromID       int                      `json:"from_id"`
	CreatedBy    int                      `json:"created_by"`
	Date         int                      `json:"date"`
	Text         string                   `json:"text"`
	ReplyOwnerID int                      `json:"reply_owner_id"`
	ReplyPostID  int                      `json:"reply_post_id"`
	FriendsOnly  int                      `json:"friends_only"`
	Comments     baseCommentsInfo         `json:"comments"`
	Likes        baseLikesInfo            `json:"likes"`
	Reposts      baseRepostsInfo          `json:"reposts"`
	Views        wallViews                `json:"views"`
	PostType     string                   `json:"post_type"`
	PostSource   wallPostSource           `json:"post_source"`
	Attachments  []wallWallpostAttachment `json:"attachments"`
	Geo          baseGeo                  `json:"geo"`
	SignerID     int                      `json:"signer_id"`
	CopyHistory  []WallWallpost           `json:"copy_history"`
	CanPin       int                      `json:"can_pin"`
	CanDelete    int                      `json:"can_delete"`
	CanEdit      int                      `json:"can_edit"`
	IsPinned     int                      `json:"is_pinned"`
	MarkedAsAds  int                      `json:"marked_as_ads"`
	IsFavorite   bool                     `json:"is_favorite"`
	AccessKey    string                   `json:"access_key"`
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
	Album             photosPhotoAlbum  `json:"album"`
	App               wallAppPost       `json:"app"`
	Audio             AudioAudioFull    `json:"audio"`
	Doc               docsDoc           `json:"doc"`
	Graffiti          wallGraffiti      `json:"graffiti"`
	Link              baseLink          `json:"link"`
	Market            marketMarketItem  `json:"market"`
	MarketMarketAlbum marketMarketAlbum `json:"market_market_album"`
	Note              wallAttachedNote  `json:"note"`
	Page              PagesWikipageFull `json:"page"`
	Photo             PhotosPhoto       `json:"photo"`
	PhotosList        []string          `json:"photos_list"`
	Poll              pollsPoll         `json:"poll"`
	PostedPhoto       wallPostedPhoto   `json:"posted_photo"`
	Type              string            `json:"type"`
	Video             VideoVideo        `json:"video"`
}

type wallWallpostToID struct {
	Attachments []wallWallpostAttachment `json:"attachments"`
	Comments    baseCommentsInfo         `json:"comments"`
	CopyOwnerID int                      `json:"copy_owner_id"`
	CopyPostID  int                      `json:"copy_post_id"`
	Date        int                      `json:"date"`
	FromID      int                      `json:"from_id"`
	Geo         baseGeo                  `json:"geo"`
	ID          int                      `json:"id"`
	Likes       baseLikesInfo            `json:"likes"`
	PostID      int                      `json:"post_id"`
	PostSource  wallPostSource           `json:"post_source"`
	PostType    string                   `json:"post_type"`
	Reposts     baseRepostsInfo          `json:"reposts"`
	SignerID    int                      `json:"signer_id"`
	Text        string                   `json:"text"`
	ToID        int                      `json:"to_id"`
}
