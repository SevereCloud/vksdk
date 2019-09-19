package object // import "github.com/SevereCloud/vksdk/object"

import (
	"fmt"
)

type marketCurrency struct {
	ID   int    `json:"id"`   // Currency ID
	Name string `json:"name"` // Currency sign
}

// MarketMarketAlbum struct
type MarketMarketAlbum struct {
	Count       int         `json:"count"`    // Items number
	ID          int         `json:"id"`       // Market album ID
	OwnerID     int         `json:"owner_id"` // Market album owner's ID
	Photo       PhotosPhoto `json:"photo"`
	Title       string      `json:"title"`        // Market album title
	UpdatedTime int         `json:"updated_time"` // Date when album has been updated last time in Unixtime
}

// ToAttachment return attachment format
func (marketAlbum MarketMarketAlbum) ToAttachment() string {
	return fmt.Sprintf("market_album%d_%d", marketAlbum.OwnerID, marketAlbum.ID)
}

// MarketMarketCategory struct
type MarketMarketCategory struct {
	ID      int           `json:"id"`   // Category ID
	Name    string        `json:"name"` // Category name
	Section marketSection `json:"section"`
}

// MarketMarketItem struct
type MarketMarketItem struct {
	AccessKey    string               `json:"access_key"` // Access key for the market item
	Availability int                  `json:"availability"`
	Category     MarketMarketCategory `json:"category"`
	Date         int                  `json:"date"`        // Date when the item has been created in Unixtime
	Description  string               `json:"description"` // Item description
	ID           int                  `json:"id"`          // Item ID
	OwnerID      int                  `json:"owner_id"`    // Item owner's ID
	Price        marketPrice          `json:"price"`
	ThumbPhoto   string               `json:"thumb_photo"` // URL of the preview image
	Title        string               `json:"title"`       // Item title
	IsFavorite   bool                 `json:"is_favorite"`
	AlbumsIDs    []int                `json:"albums_ids"`
	Photos       []PhotosPhoto        `json:"photos"`
	CanComment   int                  `json:"can_comment"`
	CanRepost    int                  `json:"can_repost"`
	Likes        baseLikesInfo        `json:"likes"`
	Reposts      baseRepostsInfo      `json:"reposts"`
	ViewsCount   int                  `json:"views_count"`
	URL          string               `json:"url"` // URL to item
	ButtonTitle  string               `json:"button_title"`
	ExternalID   string               `json:"external_id"`
}

// ToAttachment return attachment format
func (market MarketMarketItem) ToAttachment() string {
	return fmt.Sprintf("market%d_%d", market.OwnerID, market.ID)
}

type marketPrice struct {
	Amount       string         `json:"amount"` // Amount
	Currency     marketCurrency `json:"currency"`
	DiscountRate int            `json:"discount_rate"`
	OldAmount    string         `json:"old_amount"`
	Text         string         `json:"text"` // Text
}

type marketSection struct {
	ID   int    `json:"id"`   // Section ID
	Name string `json:"name"` // Section name
}
