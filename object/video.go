package object // import "github.com/severecloud/vksdk/object"

// VideoVideo struct
type VideoVideo struct {
	AccessKey   string          `json:"access_key"`
	AddingDate  int             `json:"adding_date"`
	CanAdd      int             `json:"can_add"`
	CanEdit     int             `json:"can_edit"`
	Comments    int             `json:"comments"`
	Date        int             `json:"date"`
	Description string          `json:"description"`
	Duration    int             `json:"duration"`
	Files       videoVideoFiles `json:"files"`
	ID          int             `json:"id"`
	Live        int             `json:"live"`
	OwnerID     int             `json:"owner_id"`
	Photo130    string          `json:"photo_130"`
	Photo320    string          `json:"photo_320"`
	Photo800    string          `json:"photo_800"`
	Player      string          `json:"player"`
	Processing  int             `json:"processing"`
	Title       string          `json:"title"`
	Views       int             `json:"views"`
}

type videoCatBlock struct {
	CanHide int               `json:"can_hide"`
	ID      int               `json:"id"`
	Items   []videoCatElement `json:"items"`
	Name    string            `json:"name"`
	Next    string            `json:"next"`
	Type    string            `json:"type"`
	View    string            `json:"view"`
}

type videoCatElement struct {
	CanAdd      int    `json:"can_add"`
	CanEdit     int    `json:"can_edit"`
	Comments    int    `json:"comments"`
	Count       int    `json:"count"`
	Date        int    `json:"date"`
	Description string `json:"description"`
	Duration    int    `json:"duration"`
	ID          int    `json:"id"`
	IsPrivate   int    `json:"is_private"`
	OwnerID     int    `json:"owner_id"`
	Photo130    string `json:"photo_130"`
	Photo160    string `json:"photo_160"`
	Photo320    string `json:"photo_320"`
	Photo640    string `json:"photo_640"`
	Photo800    string `json:"photo_800"`
	Title       string `json:"title"`
	Type        string `json:"type"`
	UpdatedTime int    `json:"updated_time"`
	Views       int    `json:"views"`
}

type videoSaveResult struct {
	Description string `json:"description"`
	OwnerID     int    `json:"owner_id"`
	Title       string `json:"title"`
	UploadURL   string `json:"upload_url"`
	VideoID     int    `json:"video_id"`
}

type videoUploadResponse struct {
	Size    int `json:"size"`
	VideoID int `json:"video_id"`
}

type videoVideoAlbum struct {
	ID      int    `json:"id"`
	OwnerID int    `json:"owner_id"`
	Title   string `json:"title"`
}

type videoVideoAlbumFull struct {
	Count       int    `json:"count"`
	ID          int    `json:"id"`
	IsSystem    int    `json:"is_system"`
	OwnerID     int    `json:"owner_id"`
	Photo160    string `json:"photo_160"`
	Photo320    string `json:"photo_320"`
	Title       string `json:"title"`
	UpdatedTime int    `json:"updated_time"`
}

type videoVideoFiles struct {
	External string `json:"external"`
	Mp1080   string `json:"mp_1080"`
	Mp240    string `json:"mp_240"`
	Mp360    string `json:"mp_360"`
	Mp480    string `json:"mp_480"`
	Mp720    string `json:"mp_720"`
}

type videoVideoFull struct {
	AccessKey      string          `json:"access_key"`
	AddingDate     int             `json:"adding_date"`
	CanAdd         int             `json:"can_add"`
	CanComment     int             `json:"can_comment"`
	CanEdit        int             `json:"can_edit"`
	CanRepost      int             `json:"can_repost"`
	Comments       int             `json:"comments"`
	Date           int             `json:"date"`
	Description    string          `json:"description"`
	Duration       int             `json:"duration"`
	Files          videoVideoFiles `json:"files"`
	ID             int             `json:"id"`
	Likes          baseLikes       `json:"likes"`
	Live           int             `json:"live"`
	OwnerID        int             `json:"owner_id"`
	Photo130       string          `json:"photo_130"`
	Photo320       string          `json:"photo_320"`
	Photo800       string          `json:"photo_800"`
	Player         string          `json:"player"`
	PrivacyComment []string        `json:"privacy_comment"`
	PrivacyView    []string        `json:"privacy_view"`
	Processing     int             `json:"processing"`
	Repeat         int             `json:"repeat"`
	Title          string          `json:"title"`
	Views          int             `json:"views"`
}

type videoVideoTag struct {
	Date       int    `json:"date"`
	ID         int    `json:"id"`
	PlacerID   int    `json:"placer_id"`
	TaggedName string `json:"tagged_name"`
	UserID     int    `json:"user_id"`
	Viewed     int    `json:"viewed"`
}

type videoVideoTagInfo struct {
	AccessKey   string          `json:"access_key"`
	AddingDate  int             `json:"adding_date"`
	CanAdd      int             `json:"can_add"`
	CanEdit     int             `json:"can_edit"`
	Comments    int             `json:"comments"`
	Date        int             `json:"date"`
	Description string          `json:"description"`
	Duration    int             `json:"duration"`
	Files       videoVideoFiles `json:"files"`
	ID          int             `json:"id"`
	Live        int             `json:"live"`
	OwnerID     int             `json:"owner_id"`
	Photo130    string          `json:"photo_130"`
	Photo320    string          `json:"photo_320"`
	Photo800    string          `json:"photo_800"`
	PlacerID    int             `json:"placer_id"`
	Player      string          `json:"player"`
	Processing  int             `json:"processing"`
	TagCreated  int             `json:"tag_created"`
	TagID       int             `json:"tag_id"`
	Title       string          `json:"title"`
	Views       int             `json:"views"`
}
