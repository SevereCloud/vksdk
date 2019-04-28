package object // import "github.com/severecloud/vksdk/5.92/object"

// DocsDoc struct
type DocsDoc struct {
	AccessKey string         `json:"access_key"`
	Date      int            `json:"date"`
	Ext       string         `json:"ext"`
	ID        int            `json:"id"`
	OwnerID   int            `json:"owner_id"`
	Preview   docsDocPreview `json:"preview"`
	Size      int            `json:"size"`
	Title     string         `json:"title"`
	Type      int            `json:"type"`
	URL       string         `json:"url"`
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
	Filesize int    `json:"filesize"`
	Height   int    `json:"height"`
	Src      string `json:"src"`
	Width    int    `json:"width"`
}

type docsDocPreviewAudioMessage struct {
	Duration int    `json:"duration"`
	Waveform []int  `json:"waveform"`
	LinkOgg  string `json:"link_ogg"`
	LinkMp3  string `json:"link_mp3"`
}

// DocsDocTypes struct
type DocsDocTypes struct {
	Count int    `json:"count"`
	ID    int    `json:"id"`
	Title string `json:"title"`
}

type docsDocUploadResponse struct {
	File string `json:"file"`
}
