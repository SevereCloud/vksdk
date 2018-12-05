package object // import "github.com/severecloud/vksdk/object"

type docsDoc struct {
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
	Photo docsDocPreviewPhoto `json:"photo"`
	Video docsDocPreviewVideo `json:"video"`
}

type docsDocPreviewPhoto struct {
	Sizes []photosPhotoSizes `json:"sizes"`
}

type docsDocPreviewVideo struct {
	Filesize int    `json:"filesize"`
	Height   int    `json:"height"`
	Src      string `json:"src"`
	Width    int    `json:"width"`
}

type docsDocTypes struct {
	Count int    `json:"count"`
	ID    int    `json:"id"`
	Title string `json:"title"`
}

type docsDocUploadResponse struct {
	File string `json:"file"`
}
