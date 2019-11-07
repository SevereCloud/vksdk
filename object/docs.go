package object // import "github.com/SevereCloud/vksdk/object"

import (
	"fmt"
)

// DocsDoc struct
type DocsDoc struct {
	AccessKey  string         `json:"access_key"` // Access key for the document
	Date       int            `json:"date"`       // Date when file has been uploaded in Unixtime
	Ext        string         `json:"ext"`        // File extension
	ID         int            `json:"id"`         // Document ID
	IsLicensed int            `json:"is_licensed"`
	OwnerID    int            `json:"owner_id"` // Document owner ID
	Preview    docsDocPreview `json:"preview"`
	Size       int            `json:"size"`  // File size in bites
	Title      string         `json:"title"` // Document title
	Type       int            `json:"type"`  // Document type
	URL        string         `json:"url"`   // File URL
	docsDocPreviewAudioMessage
}

// ToAttachment return attachment format
func (doc DocsDoc) ToAttachment() string {
	return fmt.Sprintf("doc%d_%d", doc.OwnerID, doc.ID)
}

type docsDocPreview struct {
	Photo        docsDocPreviewPhoto        `json:"photo"`
	Graffiti     docsDocPreviewGraffiti     `json:"graffiti"`
	Video        docsDocPreviewVideo        `json:"video"`
	AudioMessage docsDocPreviewAudioMessage `json:"audio_message"`
}

type docsDocPreviewPhoto struct {
	Sizes []docsDocPreviewPhotoSizes `json:"sizes"`
}

type docsDocPreviewPhotoSizes struct {
	// BUG(VK): json: cannot unmarshal number 162.000000 into Go struct field docsDocPreviewPhotoSizes.doc.preview.photo.sizes.height of type int
	Height float64 `json:"height"` // Height in px
	Src    string  `json:"src"`    // URL of the image
	Type   string  `json:"type"`
	Width  float64 `json:"width"` // Width in px
}

type docsDocPreviewGraffiti struct {
	Src    string `json:"src"`
	Width  int    `json:"width"`
	Height int    `json:"height"`
}

type docsDocPreviewVideo struct {
	FileSize int    `json:"file_size"` // Video file size in bites
	Height   int    `json:"height"`    // Video's height in pixels
	Src      string `json:"src"`       // Video URL
	Width    int    `json:"width"`     // Video's width in pixels
}

type docsDocPreviewAudioMessage struct {
	Duration int    `json:"duration"`
	Waveform []int  `json:"waveform"`
	LinkOgg  string `json:"link_ogg"`
	LinkMp3  string `json:"link_mp3"`
}

// DocsDocTypes struct
type DocsDocTypes struct {
	Count int    `json:"count"` // Number of docs
	ID    int    `json:"id"`    // Doc type ID
	Name  string `json:"name"`  // Doc type title
}

// DocsDocUploadResponse struct
type DocsDocUploadResponse struct {
	File string `json:"file"` // Uploaded file data
}
