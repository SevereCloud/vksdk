package object // import "github.com/SevereCloud/vksdk/v2/object"

import (
	"encoding/json"
)

// MarusiaPicture struct.
type MarusiaPicture struct {
	ID      int `json:"id"`
	OwnerID int `json:"owner_id"`
}

// MarusiaPictureUploadResponse struct.
type MarusiaPictureUploadResponse struct {
	Hash        string          `json:"hash"`   // Uploading hash
	Photo       json.RawMessage `json:"photo"`  // Uploaded photo data
	Server      int             `json:"server"` // Upload server number
	AID         int             `json:"aid"`
	MessageCode int             `json:"message_code"`
}
