package object // import "github.com/SevereCloud/vksdk/5.92/object"

// PhotosPhoto struct
type PhotosPhoto struct {
	AccessKey          string             `json:"access_key"`
	AlbumID            int                `json:"album_id"`
	Date               int                `json:"date"`
	Height             int                `json:"height"`
	ID                 int                `json:"id"`
	Images             []photosImage      `json:"images"`
	Lat                float64            `json:"lat"`
	Long               float64            `json:"long"`
	OwnerID            int                `json:"owner_id"`
	PostID             int                `json:"post_id"`
	Text               string             `json:"text"`
	UserID             int                `json:"user_id"`
	Width              int                `json:"width"`
	CanUpload          int                `json:"can_upload"`
	CommentsDisabled   int                `json:"comments_disabled"`
	Created            int                `json:"created"`
	Description        string             `json:"description"`
	PrivacyComment     []string           `json:"privacy_comment"`
	PrivacyView        []string           `json:"privacy_view"`
	Size               int                `json:"size"`
	Sizes              []photosPhotoSizes `json:"sizes"`
	ThumbID            int                `json:"thumb_id"`
	ThumbIsLast        int                `json:"thumb_is_last"`
	ThumbSrc           string             `json:"thumb_src"`
	Title              string             `json:"title"`
	Updated            int                `json:"updated"`
	UploadByAdminsOnly int                `json:"upload_by_admins_only"`
}

type photosCommentXtrPid struct {
	Attachments    []wallCommentAttachment `json:"attachments"`
	Date           int                     `json:"date"`
	FromID         int                     `json:"from_id"`
	ID             int                     `json:"id"`
	Likes          baseLikesInfo           `json:"likes"`
	Pid            int                     `json:"pid"`
	ReplyToComment int                     `json:"reply_to_comment"`
	ReplyToUser    int                     `json:"reply_to_user"`
	Text           string                  `json:"text"`
}

type photosImage struct {
	Height int    `json:"height"`
	Type   string `json:"type"`
	URL    string `json:"url"`
	Width  int    `json:"width"`
}

type photosMarketAlbumUploadResponse struct {
	Gid    int    `json:"gid"`
	Hash   string `json:"hash"`
	Photo  string `json:"photo"`
	Server int    `json:"server"`
}

type photosMarketUploadResponse struct {
	CropData string `json:"crop_data"`
	CropHash string `json:"crop_hash"`
	GroupID  int    `json:"group_id"`
	Hash     string `json:"hash"`
	Photo    string `json:"photo"`
	Server   int    `json:"server"`
}

type photosMessageUploadResponse struct {
	Hash   string `json:"hash"`
	Photo  string `json:"photo"`
	Server int    `json:"server"`
}

type photosOwnerUploadResponse struct {
	Hash   string `json:"hash"`
	Photo  string `json:"photo"`
	Server int    `json:"server"`
}

type photosPhotoAlbum struct {
	Created     int         `json:"created"`
	Description string      `json:"description"`
	ID          int         `json:"id"`
	OwnerID     int         `json:"owner_id"`
	Size        int         `json:"size"`
	Thumb       PhotosPhoto `json:"thumb"`
	Title       string      `json:"title"`
	Updated     int         `json:"updated"`
}

type photosPhotoAlbumFull struct {
	CanUpload          int                `json:"can_upload"`
	CommentsDisabled   int                `json:"comments_disabled"`
	Created            int                `json:"created"`
	Description        string             `json:"description"`
	ID                 int                `json:"id"`
	OwnerID            int                `json:"owner_id"`
	PrivacyComment     []string           `json:"privacy_comment"`
	PrivacyView        []string           `json:"privacy_view"`
	Size               int                `json:"size"`
	Sizes              []photosPhotoSizes `json:"sizes"`
	ThumbID            int                `json:"thumb_id"`
	ThumbIsLast        int                `json:"thumb_is_last"`
	ThumbSrc           string             `json:"thumb_src"`
	Title              string             `json:"title"`
	Updated            int                `json:"updated"`
	UploadByAdminsOnly int                `json:"upload_by_admins_only"`
}

type photosPhotoFull struct {
	AccessKey  string          `json:"access_key"`
	AlbumID    int             `json:"album_id"`
	CanComment int             `json:"can_comment"`
	Comments   BaseObjectCount `json:"comments"`
	Date       int             `json:"date"`
	Height     int             `json:"height"`
	ID         int             `json:"id"`
	Images     []photosImage   `json:"images"`
	Lat        float64         `json:"lat"`
	Likes      baseLikes       `json:"likes"`
	Long       float64         `json:"long"`
	OwnerID    int             `json:"owner_id"`
	PostID     int             `json:"post_id"`
	Reposts    BaseObjectCount `json:"reposts"`
	Tags       BaseObjectCount `json:"tags"`
	Text       string          `json:"text"`
	UserID     int             `json:"user_id"`
	Width      int             `json:"width"`
}

type photosPhotoFullXtrRealOffset struct {
	AccessKey  string             `json:"access_key"`
	AlbumID    int                `json:"album_id"`
	CanComment int                `json:"can_comment"`
	Comments   BaseObjectCount    `json:"comments"`
	Date       int                `json:"date"`
	Height     int                `json:"height"`
	Hidden     int                `json:"hidden"`
	ID         int                `json:"id"`
	Lat        float64            `json:"lat"`
	Likes      baseLikes          `json:"likes"`
	Long       float64            `json:"long"`
	OwnerID    int                `json:"owner_id"`
	Photo1280  string             `json:"photo_1280"`
	Photo130   string             `json:"photo_130"`
	Photo2560  string             `json:"photo_2560"`
	Photo604   string             `json:"photo_604"`
	Photo75    string             `json:"photo_75"`
	Photo807   string             `json:"photo_807"`
	PostID     int                `json:"post_id"`
	RealOffset int                `json:"real_offset"`
	Reposts    BaseObjectCount    `json:"reposts"`
	Sizes      []photosPhotoSizes `json:"sizes"`
	Tags       BaseObjectCount    `json:"tags"`
	Text       string             `json:"text"`
	UserID     int                `json:"user_id"`
	Width      int                `json:"width"`
}

type photosPhotoSizes struct {
	Height int    `json:"height"`
	Src    string `json:"src"`
	Type   string `json:"type"`
	Width  int    `json:"width"`
}

type photosPhotoTag struct {
	Date       int     `json:"date"`
	ID         int     `json:"id"`
	PlacerID   int     `json:"placer_id"`
	TaggedName string  `json:"tagged_name"`
	UserID     int     `json:"user_id"`
	Viewed     int     `json:"viewed"`
	X          float64 `json:"x"`
	X2         float64 `json:"x2"`
	Y          float64 `json:"y"`
	Y2         float64 `json:"y2"`
}

type photosPhotoUpload struct {
	AlbumID   int    `json:"album_id"`
	UploadURL string `json:"upload_url"`
	UserID    int    `json:"user_id"`
}

type photosPhotoUploadResponse struct {
	Aid        int    `json:"aid"`
	Hash       string `json:"hash"`
	PhotosList string `json:"photos_list"`
	Server     int    `json:"server"`
}

type photosPhotoXtrRealOffset struct {
	AccessKey  string             `json:"access_key"`
	AlbumID    int                `json:"album_id"`
	Date       int                `json:"date"`
	Height     int                `json:"height"`
	Hidden     int                `json:"hidden"`
	ID         int                `json:"id"`
	Lat        float64            `json:"lat"`
	Long       float64            `json:"long"`
	OwnerID    int                `json:"owner_id"`
	Photo1280  string             `json:"photo_1280"`
	Photo130   string             `json:"photo_130"`
	Photo2560  string             `json:"photo_2560"`
	Photo604   string             `json:"photo_604"`
	Photo75    string             `json:"photo_75"`
	Photo807   string             `json:"photo_807"`
	PostID     int                `json:"post_id"`
	RealOffset int                `json:"real_offset"`
	Sizes      []photosPhotoSizes `json:"sizes"`
	Text       string             `json:"text"`
	UserID     int                `json:"user_id"`
	Width      int                `json:"width"`
}

type photosPhotoXtrTagInfo struct {
	AccessKey  string             `json:"access_key"`
	AlbumID    int                `json:"album_id"`
	Date       int                `json:"date"`
	Height     int                `json:"height"`
	ID         int                `json:"id"`
	Lat        float64            `json:"lat"`
	Long       float64            `json:"long"`
	OwnerID    int                `json:"owner_id"`
	Photo1280  string             `json:"photo_1280"`
	Photo130   string             `json:"photo_130"`
	Photo2560  string             `json:"photo_2560"`
	Photo604   string             `json:"photo_604"`
	Photo75    string             `json:"photo_75"`
	Photo807   string             `json:"photo_807"`
	PlacerID   int                `json:"placer_id"`
	PostID     int                `json:"post_id"`
	Sizes      []photosPhotoSizes `json:"sizes"`
	TagCreated int                `json:"tag_created"`
	TagID      int                `json:"tag_id"`
	Text       string             `json:"text"`
	UserID     int                `json:"user_id"`
	Width      int                `json:"width"`
}

// PhotosWallUploadResponse struct
type PhotosWallUploadResponse struct {
	Hash   string `json:"hash"`
	Photo  string `json:"photo"`
	Server int    `json:"server"`
}
