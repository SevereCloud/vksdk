package object // import "github.com/SevereCloud/vksdk/object"

import (
	"fmt"
)

// PhotosPhoto struct
type PhotosPhoto struct {
	AccessKey          string             `json:"access_key"` // Access key for the photo
	AlbumID            int                `json:"album_id"`   // Album ID
	Date               int                `json:"date"`       // Date when uploaded
	Height             int                `json:"height"`     // Original photo height
	ID                 int                `json:"id"`         // Photo ID
	Images             []PhotosImage      `json:"images"`
	Lat                float64            `json:"lat"`      // Latitude
	Long               float64            `json:"long"`     // Longitude
	OwnerID            int                `json:"owner_id"` // Photo owner's ID
	PostID             int                `json:"post_id"`  // Post ID
	Text               string             `json:"text"`     // Photo caption
	UserID             int                `json:"user_id"`  // ID of the user who have uploaded the photo
	Width              int                `json:"width"`    // Original photo width
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

// ToAttachment return attachment format
func (photo PhotosPhoto) ToAttachment() string {
	return fmt.Sprintf("photo%d_%d", photo.OwnerID, photo.ID)
}

// PhotosCommentXtrPid struct
type PhotosCommentXtrPid struct {
	Attachments    []wallCommentAttachment `json:"attachments"`
	Date           int                     `json:"date"`    // Date when the comment has been added in Unixtime
	FromID         int                     `json:"from_id"` // Author ID
	ID             int                     `json:"id"`      // Comment ID
	Likes          baseLikesInfo           `json:"likes"`
	ParentsStack   []int                   `json:"parents_stack"`
	Pid            int                     `json:"pid"`              // Photo ID
	ReplyToComment int                     `json:"reply_to_comment"` // Replied comment ID
	ReplyToUser    int                     `json:"reply_to_user"`    // Replied user ID
	Text           string                  `json:"text"`             // Comment text
	Thread         wallWallCommentThread   `json:"thread"`
}

// PhotosImage struct
type PhotosImage struct {
	Height int    `json:"height"` // Height of the photo in px.
	Type   string `json:"type"`
	URL    string `json:"url"`   // Photo URL.
	Width  int    `json:"width"` // Width of the photo in px.
}

// PhotosChatUploadResponse struct
type PhotosChatUploadResponse struct {
	Response string `json:"response"` // Uploaded photo data
}

// PhotosMarketAlbumUploadResponse struct
type PhotosMarketAlbumUploadResponse struct {
	GID    int    `json:"gid"`    // Community ID
	Hash   string `json:"hash"`   // Uploading hash
	Photo  string `json:"photo"`  // Uploaded photo data
	Server int    `json:"server"` // Upload server number
}

// PhotosMarketUploadResponse struct
type PhotosMarketUploadResponse struct {
	CropData string `json:"crop_data"` // Crop data
	CropHash string `json:"crop_hash"` // Crop hash
	GroupID  int    `json:"group_id"`  // Community ID
	Hash     string `json:"hash"`      // Uploading hash
	Photo    string `json:"photo"`     // Uploaded photo data
	Server   int    `json:"server"`    // Upload server number
}

// PhotosMessageUploadResponse struct
type PhotosMessageUploadResponse struct {
	Hash   string `json:"hash"`   // Uploading hash
	Photo  string `json:"photo"`  // Uploaded photo data
	Server int    `json:"server"` // Upload server number
}

// PhotosOwnerUploadResponse struct
type PhotosOwnerUploadResponse struct {
	Hash   string `json:"hash"`   // Uploading hash
	Photo  string `json:"photo"`  // Uploaded photo data
	Server int    `json:"server"` // Upload server number
}

type photosPhotoAlbum struct {
	Created     int         `json:"created"`     // Date when the album has been created in Unixtime
	Description string      `json:"description"` // Photo album description
	ID          string      `json:"id"`          // BUG(VK): Photo album ID
	OwnerID     int         `json:"owner_id"`    // Album owner's ID
	Size        int         `json:"size"`        // Photos number
	Thumb       PhotosPhoto `json:"thumb"`
	Title       string      `json:"title"`   // Photo album title
	Updated     int         `json:"updated"` // Date when the album has been updated last time in Unixtime
}

// ToAttachment return attachment format
func (album photosPhotoAlbum) ToAttachment() string {
	return fmt.Sprintf("album%d_%s", album.OwnerID, album.ID)
}

// PhotosPhotoAlbumFull struct
type PhotosPhotoAlbumFull struct {
	CanUpload        int    `json:"can_upload"`        // Information whether current user can upload photo to the album
	CommentsDisabled int    `json:"comments_disabled"` // Information whether album comments are disabled
	Created          int    `json:"created"`           // Date when the album has been created in Unixtime
	Description      string `json:"description"`       // Photo album description
	ID               int    `json:"id"`                // Photo album ID
	OwnerID          int    `json:"owner_id"`          // Album owner's ID
	Size             int    `json:"size"`              // Photos number
	// TODO: PrivacyComment     interface{}           `json:"privacy_comment"`
	// TODO: PrivacyView        interface{}           `json:"privacy_view"`
	Sizes              []photosPhotoSizes `json:"sizes"`
	ThumbID            int                `json:"thumb_id"`              // Thumb photo ID
	ThumbIsLast        int                `json:"thumb_is_last"`         // Information whether the album thumb is last photo
	ThumbSrc           string             `json:"thumb_src"`             // URL of the thumb image
	Title              string             `json:"title"`                 // Photo album title
	Updated            int                `json:"updated"`               // Date when the album has been updated last time in Unixtime
	UploadByAdminsOnly int                `json:"upload_by_admins_only"` // Information whether only community administrators can upload photos
}

// ToAttachment return attachment format
func (album PhotosPhotoAlbumFull) ToAttachment() string {
	return fmt.Sprintf("album%d_%d", album.OwnerID, album.ID)
}

// PhotosPhotoFull struct
type PhotosPhotoFull struct {
	AccessKey  string          `json:"access_key"`  // Access key for the photo
	AlbumID    int             `json:"album_id"`    // Album ID
	CanComment int             `json:"can_comment"` // Information whether current user can comment the photo
	Comments   BaseObjectCount `json:"comments"`
	Date       int             `json:"date"`   // Date when uploaded
	Height     int             `json:"height"` // Original photo height
	ID         int             `json:"id"`     // Photo ID
	Images     []PhotosImage   `json:"images"`
	Lat        float64         `json:"lat"` // Latitude
	Likes      baseLikes       `json:"likes"`
	Long       float64         `json:"long"`     // Longitude
	OwnerID    int             `json:"owner_id"` // Photo owner's ID
	PostID     int             `json:"post_id"`  // Post ID
	Reposts    BaseObjectCount `json:"reposts"`
	Tags       BaseObjectCount `json:"tags"`
	Text       string          `json:"text"`    // Photo caption
	UserID     int             `json:"user_id"` // ID of the user who have uploaded the photo
	Width      int             `json:"width"`   // Original photo width
}

// ToAttachment return attachment format
func (photo PhotosPhotoFull) ToAttachment() string {
	return fmt.Sprintf("photo%d_%d", photo.OwnerID, photo.ID)
}

// PhotosPhotoFullXtrRealOffset struct
type PhotosPhotoFullXtrRealOffset struct {
	AccessKey  string             `json:"access_key"` // Access key for the photo
	AlbumID    int                `json:"album_id"`   // Album ID
	CanComment int                `json:"can_comment"`
	Comments   BaseObjectCount    `json:"comments"`
	Date       int                `json:"date"`   // Date when uploaded
	Height     int                `json:"height"` // Original photo height
	Hidden     int                `json:"hidden"` // Returns if the photo is hidden above the wall
	ID         int                `json:"id"`     // Photo ID
	Lat        float64            `json:"lat"`    // Latitude
	Likes      baseLikes          `json:"likes"`
	Long       float64            `json:"long"`        // Longitude
	OwnerID    int                `json:"owner_id"`    // Photo owner's ID
	Photo1280  string             `json:"photo_1280"`  // URL of image with 1280 px width
	Photo130   string             `json:"photo_130"`   // URL of image with 130 px width
	Photo2560  string             `json:"photo_2560"`  // URL of image with 2560 px width
	Photo604   string             `json:"photo_604"`   // URL of image with 604 px width
	Photo75    string             `json:"photo_75"`    // URL of image with 75 px width
	Photo807   string             `json:"photo_807"`   // URL of image with 807 px width
	PostID     int                `json:"post_id"`     // Post ID
	RealOffset int                `json:"real_offset"` // Real position of the photo
	Reposts    BaseObjectCount    `json:"reposts"`
	Sizes      []photosPhotoSizes `json:"sizes"`
	Tags       BaseObjectCount    `json:"tags"`
	Text       string             `json:"text"`    // Photo caption
	UserID     int                `json:"user_id"` // ID of the user who have uploaded the photo
	Width      int                `json:"width"`   // Original photo width
}

// ToAttachment return attachment format
func (photo PhotosPhotoFullXtrRealOffset) ToAttachment() string {
	return fmt.Sprintf("photo%d_%d", photo.OwnerID, photo.ID)
}

type photosPhotoSizes struct {
	// BUG(VK): json: cannot unmarshal number 180.000000 into Go struct field photosPhotoSizes.height of type int
	Height float64 `json:"height"` // Height in px
	URL    string  `json:"url"`    // URL of the image
	Type   string  `json:"type"`
	Width  float64 `json:"width"` // Width in px
}

// PhotosPhotoTag struct
type PhotosPhotoTag struct {
	Date       int     `json:"date"`        // Date when tag has been added in Unixtime
	ID         int     `json:"id"`          // Tag ID
	PlacerID   int     `json:"placer_id"`   // ID of the tag creator
	TaggedName string  `json:"tagged_name"` // Tag description
	UserID     int     `json:"user_id"`     // Tagged user ID
	Viewed     int     `json:"viewed"`      // Information whether the tag is reviewed
	X          float64 `json:"x"`           // Coordinate X of the left upper corner
	X2         float64 `json:"x2"`          // Coordinate X of the right lower corner
	Y          float64 `json:"y"`           // Coordinate Y of the left upper corner
	Y2         float64 `json:"y2"`          // Coordinate Y of the right lower corner
}

// PhotosPhotoUpload struct
type PhotosPhotoUpload struct {
	AlbumID   int    `json:"album_id"`   // Album ID
	UploadURL string `json:"upload_url"` // URL to upload photo
	UserID    int    `json:"user_id"`    // User ID
}

// PhotosPhotoUploadResponse struct
type PhotosPhotoUploadResponse struct {
	AID        int    `json:"aid"`         // Album ID
	Hash       string `json:"hash"`        // Uploading hash
	PhotosList string `json:"photos_list"` // Uploaded photos data
	Server     int    `json:"server"`      // Upload server number
}

// PhotosPhotoXtrRealOffset struct
type PhotosPhotoXtrRealOffset struct {
	AccessKey  string             `json:"access_key"`  // Access key for the photo
	AlbumID    int                `json:"album_id"`    // Album ID
	Date       int                `json:"date"`        // Date when uploaded
	Height     int                `json:"height"`      // Original photo height
	Hidden     int                `json:"hidden"`      // Returns if the photo is hidden above the wall
	ID         int                `json:"id"`          // Photo ID
	Lat        float64            `json:"lat"`         // Latitude
	Long       float64            `json:"long"`        // Longitude
	OwnerID    int                `json:"owner_id"`    // Photo owner's ID
	Photo1280  string             `json:"photo_1280"`  // URL of image with 1280 px width
	Photo130   string             `json:"photo_130"`   // URL of image with 130 px width
	Photo2560  string             `json:"photo_2560"`  // URL of image with 2560 px width
	Photo604   string             `json:"photo_604"`   // URL of image with 604 px width
	Photo75    string             `json:"photo_75"`    // URL of image with 75 px width
	Photo807   string             `json:"photo_807"`   // URL of image with 807 px width
	PostID     int                `json:"post_id"`     // Post ID
	RealOffset int                `json:"real_offset"` // Real position of the photo
	Sizes      []photosPhotoSizes `json:"sizes"`
	Text       string             `json:"text"`    // Photo caption
	UserID     int                `json:"user_id"` // ID of the user who have uploaded the photo
	Width      int                `json:"width"`   // Original photo width
}

// ToAttachment return attachment format
func (photo PhotosPhotoXtrRealOffset) ToAttachment() string {
	return fmt.Sprintf("photo%d_%d", photo.OwnerID, photo.ID)
}

// PhotosPhotoXtrTagInfo struct
type PhotosPhotoXtrTagInfo struct {
	AccessKey  string             `json:"access_key"` // Access key for the photo
	AlbumID    int                `json:"album_id"`   // Album ID
	Date       int                `json:"date"`       // Date when uploaded
	Height     int                `json:"height"`     // Original photo height
	ID         int                `json:"id"`         // Photo ID
	Lat        float64            `json:"lat"`        // Latitude
	Long       float64            `json:"long"`       // Longitude
	OwnerID    int                `json:"owner_id"`   // Photo owner's ID
	Photo1280  string             `json:"photo_1280"` // URL of image with 1280 px width
	Photo130   string             `json:"photo_130"`  // URL of image with 130 px width
	Photo2560  string             `json:"photo_2560"` // URL of image with 2560 px width
	Photo604   string             `json:"photo_604"`  // URL of image with 604 px width
	Photo75    string             `json:"photo_75"`   // URL of image with 75 px width
	Photo807   string             `json:"photo_807"`  // URL of image with 807 px width
	PlacerID   int                `json:"placer_id"`  // ID of the tag creator
	PostID     int                `json:"post_id"`    // Post ID
	Sizes      []photosPhotoSizes `json:"sizes"`
	TagCreated int                `json:"tag_created"` // Date when tag has been added in Unixtime
	TagID      int                `json:"tag_id"`      // Tag ID
	Text       string             `json:"text"`        // Photo caption
	UserID     int                `json:"user_id"`     // ID of the user who have uploaded the photo
	Width      int                `json:"width"`       // Original photo width
}

// ToAttachment return attachment format
func (photo PhotosPhotoXtrTagInfo) ToAttachment() string {
	return fmt.Sprintf("photo%d_%d", photo.OwnerID, photo.ID)
}

// PhotosWallUploadResponse struct
type PhotosWallUploadResponse struct {
	Hash   string `json:"hash"`   // Uploading hash
	Photo  string `json:"photo"`  // Uploaded photo data
	Server int    `json:"server"` // Upload server number
}
