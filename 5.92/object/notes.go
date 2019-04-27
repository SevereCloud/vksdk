package object // import "github.com/severecloud/vksdk/5.92/object"

// NotesNote struct
type NotesNote struct {
	CanComment     int           `json:"can_comment"`
	Comments       int           `json:"comments"`
	Date           int           `json:"date"`
	ID             int           `json:"id"`
	OwnerID        int           `json:"owner_id"`
	Text           string        `json:"text"`
	TextWiki       string        `json:"text_wiki"`
	Title          string        `json:"title"`
	ViewURL        string        `json:"view_url"`
	ReadComments   int           `json:"read_comments"`
	PrivacyView    []interface{} `json:"privacy_view"`
	PrivacyComment []interface{} `json:"privacy_comment"`
}

// NotesNoteComment struct
type NotesNoteComment struct {
	Date    int    `json:"date"`
	ID      int    `json:"id"`
	Message string `json:"message"`
	NID     int    `json:"nid"`
	OID     int    `json:"oid"`
	ReplyTo int    `json:"reply_to"`
	UID     int    `json:"uid"`
}
