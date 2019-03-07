package object // import "github.com/severecloud/vksdk/5.92/object"

// GiftsGift struct
type GiftsGift struct {
	Date     int         `json:"date"`
	FromID   int         `json:"from_id"`
	Gift     GiftsLayout `json:"gift"`
	GiftHash string      `json:"gift_hash"`
	ID       int         `json:"id"`
	Message  string      `json:"message"`
	Privacy  int         `json:"privacy"`
}

// GiftsLayout struct
type GiftsLayout struct {
	ID       int    `json:"id"`
	Thumb256 string `json:"thumb_256"`
	Thumb48  string `json:"thumb_48"`
	Thumb96  string `json:"thumb_96"`
}
