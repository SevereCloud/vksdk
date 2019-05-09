package object // import "github.com/SevereCloud/vksdk/5.92/object"

// VideoVideo struct
type VideoVideo struct {
	// Video access key
	AccessKey string `json:"access_key,omitempty"`

	// Date when the video has been added in Unixtime
	AddingDate int `json:"adding_date,omitempty"`

	// Information whether current user can add the video
	CanAdd int `json:"can_add,omitempty"`

	// Information whether current user can comment the video
	CanComment int `json:"can_comment,omitempty"`

	// Information whether current user can edit the video
	CanEdit int `json:"can_edit,omitempty"`

	// Information whether current user can like the video
	CanLike int `json:"can_like,omitempty"`

	// Information whether current user can repost this video
	CanRepost int `json:"can_repost,omitempty"`

	// Number of comments
	Comments int `json:"comments,omitempty"`

	// Date when video has been uploaded in Unixtime
	Date int `json:"date,omitempty"`

	// Video description
	Description string `json:"description,omitempty"`

	// Video duration in seconds
	Duration int              `json:"duration,omitempty"`
	Files    *videoVideoFiles `json:"files,omitempty"`

	// URL of the first frame for the corresponding width.
	FirstFrame130 string `json:"first_frame_130,omitempty"`

	// URL of the first frame for the corresponding width.
	FirstFrame160 string `json:"first_frame_160,omitempty"`

	// URL of the first frame for the corresponding width.
	FirstFrame320 string `json:"first_frame_320,omitempty"`

	// URL of the first frame for the corresponding width.
	FirstFrame800 string `json:"first_frame_800,omitempty"`

	// Video height
	Height int `json:"height,omitempty"`

	// Video ID
	ID         int  `json:"id,omitempty"`
	IsFavorite bool `json:"is_favorite,omitempty"`

	// Returns if the video is a live stream
	Live int `json:"live,omitempty"`

	// Video owner ID
	OwnerID int `json:"owner_id,omitempty"`

	// URL of the preview image with 130 px in width
	Photo130 string `json:"photo_130,omitempty"`

	// URL of the preview image with 320 px in width
	Photo320 string `json:"photo_320,omitempty"`

	// URL of the preview image with 800 px in width
	Photo800 string `json:"photo_800,omitempty"`

	// URL of the page with a player that can be used to play the video in the browser.
	Player string `json:"player,omitempty"`

	// Returns if the video is processing
	Processing int `json:"processing,omitempty"`

	// Video title
	Title string `json:"title,omitempty"`
	Type  string `json:"type,omitempty"`

	// Number of views
	Views int `json:"views,omitempty"`

	// Video width
	Width int `json:"width,omitempty"`

	Platform   string `json:"platform"`
	LocalViews int    `json:"local_views"`
	IsPrivate  int    `json:"is_private"`
}

// videoVideoFiles struct
type videoVideoFiles struct {
	// URL of the external player
	External string `json:"external,omitempty"`

	// URL of the mpeg4 file with 1080p quality
	Mp1080 string `json:"mp_1080,omitempty"`

	// URL of the mpeg4 file with 240p quality
	Mp240 string `json:"mp_240,omitempty"`

	// URL of the mpeg4 file with 360p quality
	Mp360 string `json:"mp_360,omitempty"`

	// URL of the mpeg4 file with 480p quality
	Mp480 string `json:"mp_480,omitempty"`

	// URL of the mpeg4 file with 720p quality
	Mp720 string `json:"mp_720,omitempty"`
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
