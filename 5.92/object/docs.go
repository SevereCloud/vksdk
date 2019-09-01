package object // import "github.com/SevereCloud/vksdk/5.92/object"

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
}

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
	Sizes []photosPhotoSizes `json:"sizes"`
}

type docsDocPreviewGraffiti struct {
	Src    string `json:"src"`
	Width  int    `json:"width"`
	Height int    `json:"height"`
}

type docsDocPreviewVideo struct {
	Filesize int    `json:"filesize"` // Video file size in bites
	Height   int    `json:"height"`   // Video's height in pixels
	Src      string `json:"src"`      // Video URL
	Width    int    `json:"width"`    // Video's width in pixels
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
	Title string `json:"title"` // Doc type title
}

// DocsDocUploadResponse struct
type DocsDocUploadResponse struct {
	File string `json:"file"` // Uploaded file data
}
