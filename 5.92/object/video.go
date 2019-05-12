package object // import "github.com/SevereCloud/vksdk/5.92/object"

// VideoVideo struct
type VideoVideo struct {
	AccessKey     string          `json:"access_key"`  // Video access key
	AddingDate    int             `json:"adding_date"` // Date when the video has been added in Unixtime
	CanAdd        int             `json:"can_add"`     // Information whether current user can add the video
	CanComment    int             `json:"can_comment"` // Information whether current user can comment the video
	CanEdit       int             `json:"can_edit"`    // Information whether current user can edit the video
	CanLike       int             `json:"can_like"`    // Information whether current user can like the video
	CanRepost     int             `json:"can_repost"`  // Information whether current user can repost this video
	Comments      int             `json:"comments"`    // Number of comments
	Date          int             `json:"date"`        // Date when video has been uploaded in Unixtime
	Description   string          `json:"description"` // Video description
	Duration      int             `json:"duration"`    // Video duration in seconds
	Files         videoVideoFiles `json:"files"`
	FirstFrame130 string          `json:"first_frame_130"` // URL of the first frame for the corresponding width.
	FirstFrame160 string          `json:"first_frame_160"` // URL of the first frame for the corresponding width.
	FirstFrame320 string          `json:"first_frame_320"` // URL of the first frame for the corresponding width.
	FirstFrame800 string          `json:"first_frame_800"` // URL of the first frame for the corresponding width.
	Height        int             `json:"height"`          // Video height
	ID            int             `json:"id"`              // Video ID
	IsFavorite    bool            `json:"is_favorite"`
	Live          int             `json:"live"`       // Returns if the video is a live stream
	OwnerID       int             `json:"owner_id"`   // Video owner ID
	Photo130      string          `json:"photo_130"`  // URL of the preview image with 130 px in width
	Photo320      string          `json:"photo_320"`  // URL of the preview image with 320 px in width
	Photo800      string          `json:"photo_800"`  // URL of the preview image with 800 px in width
	Player        string          `json:"player"`     // URL of the page with a player that can be used to play the video in the browser.
	Processing    int             `json:"processing"` // Returns if the video is processing
	Title         string          `json:"title"`      // Video title
	Type          string          `json:"type"`
	Views         int             `json:"views"` // Number of views
	Width         int             `json:"width"` // Video width
	Platform      string          `json:"platform"`
	LocalViews    int             `json:"local_views"`
	IsPrivate     int             `json:"is_private"`
}

// videoVideoFiles struct
type videoVideoFiles struct {
	External string `json:"external"` // URL of the external player
	Mp1080   string `json:"mp_1080"`  // URL of the mpeg4 file with 1080p quality
	Mp240    string `json:"mp_240"`   // URL of the mpeg4 file with 240p quality
	Mp360    string `json:"mp_360"`   // URL of the mpeg4 file with 360p quality
	Mp480    string `json:"mp_480"`   // URL of the mpeg4 file with 480p quality
	Mp720    string `json:"mp_720"`   // URL of the mpeg4 file with 720p quality
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
	Description string `json:"description"` // Video description
	OwnerID     int    `json:"owner_id"`    // Video owner ID
	Title       string `json:"title"`       // Video title
	UploadURL   string `json:"upload_url"`  // URL for the video uploading
	VideoID     int    `json:"video_id"`    // Video ID
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
	Count       int               `json:"count"`        // Total number of videos in album
	ID          int               `json:"id"`           // Album ID
	Image       []videoVideoImage `json:"image"`        // Album cover image in different sizes
	IsSystem    int               `json:"is_system"`    // Information whether album is system
	OwnerID     int               `json:"owner_id"`     // Album owner's ID
	Photo160    string            `json:"photo_160"`    // URL of the preview image with 160px in width
	Photo320    string            `json:"photo_320"`    // URL of the preview image with 320px in width
	Title       string            `json:"title"`        // Album title
	UpdatedTime int               `json:"updated_time"` // Date when the album has been updated last time in Unixtime
}

type videoVideoFull struct {
	AccessKey     string          `json:"access_key"`  // Video access key
	AddingDate    int             `json:"adding_date"` // Date when the video has been added in Unixtime
	CanAdd        int             `json:"can_add"`     // Information whether current user can add the video
	CanComment    int             `json:"can_comment"` // Information whether current user can comment the video
	CanEdit       int             `json:"can_edit"`    // Information whether current user can edit the video
	CanRepost     int             `json:"can_repost"`  // Information whether current user can comment the video
	Comments      int             `json:"comments"`    // Number of comments
	Date          int             `json:"date"`        // Date when video has been uploaded in Unixtime
	Description   string          `json:"description"` // Video description
	Duration      int             `json:"duration"`    // Video duration in seconds
	Files         videoVideoFiles `json:"files"`
	FirstFrame130 string          `json:"first_frame_130"` // URL of the first frame for the corresponding width.
	FirstFrame160 string          `json:"first_frame_160"` // URL of the first frame for the corresponding width.
	FirstFrame320 string          `json:"first_frame_320"` // URL of the first frame for the corresponding width.
	FirstFrame800 string          `json:"first_frame_800"` // URL of the first frame for the corresponding width.
	ID            int             `json:"id"`              // Video ID
	Likes         baseLikes       `json:"likes"`
	Live          int             `json:"live"`       // Returns if the video is live translation
	OwnerID       int             `json:"owner_id"`   // Video owner ID
	Photo130      string          `json:"photo_130"`  // URL of the preview image with 130 px in width
	Photo320      string          `json:"photo_320"`  // URL of the preview image with 320 px in width
	Photo800      string          `json:"photo_800"`  // URL of the preview image with 800 px in width
	Player        string          `json:"player"`     // URL of the page with a player that can be used to play the video in the browser.
	Processing    int             `json:"processing"` // Returns if the video is processing
	Repeat        int             `json:"repeat"`     // Information whether the video is repeated
	Title         string          `json:"title"`      // Video title
	Views         int             `json:"views"`      // Number of views
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

type videoVideoImage struct {
}
