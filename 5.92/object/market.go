package object // import "github.com/severecloud/vksdk/5.92/object"

type marketCurrency struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type marketMarketAlbum struct {
	Count       int         `json:"count"`
	ID          int         `json:"id"`
	OwnerID     int         `json:"owner_id"`
	Photo       PhotosPhoto `json:"photo"`
	Title       string      `json:"title"`
	UpdatedTime int         `json:"updated_time"`
}

type marketMarketCategory struct {
	ID      int           `json:"id"`
	Name    string        `json:"name"`
	Section marketSection `json:"section"`
}

type marketMarketItem struct {
	Availability int                  `json:"availability"`
	Category     marketMarketCategory `json:"category"`
	Date         int                  `json:"date"`
	Description  string               `json:"description"`
	ID           int                  `json:"id"`
	OwnerID      int                  `json:"owner_id"`
	Price        marketPrice          `json:"price"`
	ThumbPhoto   string               `json:"thumb_photo"`
	Title        string               `json:"title"`
}

type marketPrice struct {
	Amount   string         `json:"amount"`
	Currency marketCurrency `json:"currency"`
	Text     string         `json:"text"`
}

type marketSection struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}
