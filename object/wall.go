package object

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
	Page              pagesWikipageFull `json:"page"`
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
}

type wallWallpost struct {
	AccessKey   string                   `json:"access_key"`
	Attachments []wallWallpostAttachment `json:"attachments"`
	Date        int                      `json:"date"`
	FromID      int                      `json:"from_id"`
	Geo         baseGeo                  `json:"geo"`
	ID          int                      `json:"id"`
	OwnerID     int                      `json:"owner_id"`
	PostSource  wallPostSource           `json:"post_source"`
	PostType    string                   `json:"post_type"`
	SignerID    int                      `json:"signer_id"`
	Text        string                   `json:"text"`
	Views       wallViews                `json:"views"`
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
	Page              pagesWikipageFull `json:"page"`
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
