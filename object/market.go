package object // import "github.com/SevereCloud/vksdk/object"

import (
	"fmt"
)

// Information whether the MarketMarketItem is available.
const (
	MarketItemAvailable = iota
	MarketItemRemoved
	MarketItemUnavailable
)

// MarketCurrency struct.
type MarketCurrency struct {
	ID   int    `json:"id"`   // Currency ID
	Name string `json:"name"` // Currency sign
}

// MarketMarketAlbum struct.
type MarketMarketAlbum struct {
	Count       int         `json:"count"`    // Items number
	ID          int         `json:"id"`       // Market album ID
	OwnerID     int         `json:"owner_id"` // Market album owner's ID
	Photo       PhotosPhoto `json:"photo"`
	Title       string      `json:"title"`        // Market album title
	UpdatedTime int         `json:"updated_time"` // Date when album has been updated last time in Unixtime
}

// ToAttachment return attachment format.
func (marketAlbum MarketMarketAlbum) ToAttachment() string {
	return fmt.Sprintf("market_album%d_%d", marketAlbum.OwnerID, marketAlbum.ID)
}

// MarketMarketCategory struct.
type MarketMarketCategory struct {
	ID      int           `json:"id"`   // Category ID
	Name    string        `json:"name"` // Category name
	Section MarketSection `json:"section"`
}

// MarketMarketItem struct.
type MarketMarketItem struct {
	AccessKey    string               `json:"access_key"`   // Access key for the market item
	Availability int                  `json:"availability"` // Information whether the item is available
	Category     MarketMarketCategory `json:"category"`
	Date         int                  `json:"date"`        // Date when the item has been created in Unixtime
	Description  string               `json:"description"` // Item description
	ID           int                  `json:"id"`          // Item ID
	OwnerID      int                  `json:"owner_id"`    // Item owner's ID
	Price        MarketPrice          `json:"price"`
	ThumbPhoto   string               `json:"thumb_photo"` // URL of the preview image
	Title        string               `json:"title"`       // Item title
	CanComment   BaseBoolInt          `json:"can_comment"`
	CanRepost    BaseBoolInt          `json:"can_repost"`
	IsFavorite   BaseBoolInt          `json:"is_favorite"`
	AlbumsIDs    []int                `json:"albums_ids"`
	Photos       []PhotosPhoto        `json:"photos"`
	Likes        BaseLikesInfo        `json:"likes"`
	Reposts      BaseRepostsInfo      `json:"reposts"`
	ViewsCount   int                  `json:"views_count"`
	URL          string               `json:"url"` // URL to item
	ButtonTitle  string               `json:"button_title"`
	ExternalID   string               `json:"external_id"`
}

// ToAttachment return attachment format.
func (market MarketMarketItem) ToAttachment() string {
	return fmt.Sprintf("market%d_%d", market.OwnerID, market.ID)
}

// MarketPrice struct.
type MarketPrice struct {
	Amount       string         `json:"amount"` // Amount
	Currency     MarketCurrency `json:"currency"`
	DiscountRate int            `json:"discount_rate"`
	OldAmount    string         `json:"old_amount"`
	Text         string         `json:"text"` // Text
}

// MarketSection struct.
type MarketSection struct {
	ID   int    `json:"id"`   // Section ID
	Name string `json:"name"` // Section name
}

// MarketOrder struct.
type MarketOrder struct {
	ID                int                 `json:"id"`
	GroupID           int                 `json:"group_id"`
	UserID            int                 `json:"user_id"`
	Date              int                 `json:"date"`
	Status            int                 `json:"status"`
	ItemsCount        int                 `json:"items_count"`
	TotalPrice        MarketPrice         `json:"total_price"`
	DisplayOrderID    string              `json:"display_order_id"`
	TrackNumber       string              `json:"track_number"`
	TrackLink         string              `json:"track_link"`
	Comment           string              `json:"comment"`
	Address           string              `json:"address"`
	PreviewOrderItems []MarketOrderItem   `json:"preview_order_items"`
	PriceDetails      []MarketPriceDetail `json:"price_details"`
}

// MarketOrderItem struct.
type MarketOrderItem struct {
	OwnerID  int              `json:"owner_id"`
	ItemID   int              `json:"item_id"`
	Price    MarketPrice      `json:"price"`
	Quantity int              `json:"quantity"`
	Item     MarketMarketItem `json:"item"`
	Title    string           `json:"title"`
	Photo    PhotosPhoto      `json:"photo"`
	Variants []string         `json:"variants"`
}

// MarketPriceDetail struct.
type MarketPriceDetail struct {
	Title    string      `json:"title"`
	Price    MarketPrice `json:"price"`
	IsAccent bool        `json:"is_accent,omitempty"`
}
