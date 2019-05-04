package object // import "github.com/SevereCloud/vksdk/5.92/object"

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

// MarketMarketItem struct
type MarketMarketItem struct {
	Availability int                  `json:"availability"`
	Category     marketMarketCategory `json:"category"`
	Date         int                  `json:"date"`
	Description  string               `json:"description"`
	ID           int                  `json:"id"`
	OwnerID      int                  `json:"owner_id"`
	Price        marketPrice          `json:"price"`
	ThumbPhoto   string               `json:"thumb_photo"`
	Title        string               `json:"title"`
	IsFavorite   bool                 `json:"is_favorite"`
	AlbumsIds    []int                `json:"albums_ids"`
	Photos       []PhotosPhoto        `json:"photos"`
	CanComment   int                  `json:"can_comment"`
	CanRepost    int                  `json:"can_repost"`
	Likes        baseLikesInfo        `json:"likes"`
	ViewsCount   int                  `json:"views_count"`
	URL          string               `json:"url"`
	ButtonTitle  string               `json:"button_title"`
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
