package object

// AppWidgetsAppImageUploadResponse struct
type AppWidgetsAppImageUploadResponse struct {
	Image string `json:"image"`
	Hash  string `json:"hash"`
}

// AppWidgetsGroupImageUploadResponse struct
type AppWidgetsGroupImageUploadResponse struct {
	Image string `json:"image"`
	Hash  string `json:"hash"`
}

// AppWidgetsImagestruct
type AppWidgetsImage struct {
	ID     string    `json:"id"`
	Type   string    `json:"type"`
	Images BaseImage `json:"images"`
}
