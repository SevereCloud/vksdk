package object // import "github.com/severecloud/vksdk/object"

type notesNote struct {
	CanComment int    `json:"can_comment"`
	Comments   int    `json:"comments"`
	Date       int    `json:"date"`
	ID         int    `json:"id"`
	OwnerID    int    `json:"owner_id"`
	Text       string `json:"text"`
	TextWiki   string `json:"text_wiki"`
	Title      string `json:"title"`
	ViewURL    string `json:"view_url"`
}

type notesNoteComment struct {
	Date    int    `json:"date"`
	ID      int    `json:"id"`
	Message string `json:"message"`
	Nid     int    `json:"nid"`
	Oid     int    `json:"oid"`
	ReplyTo int    `json:"reply_to"`
	UID     int    `json:"uid"`
}
