package object

// StoreProductType type.
type StoreProductType string

// Possible values.
const (
	StoreProductStickers StoreProductType = "stickers"
)

// StoreProduct struct.
type StoreProduct struct {
	ID           int         `json:"id"`
	Type         string      `json:"type"`
	Purchased    BaseBoolInt `json:"purchased"`
	Active       BaseBoolInt `json:"active"`
	PurchaseDate int         `json:"purchase_date,omitempty"`
}

// StoreProductExtended struct.
type StoreProductExtended struct {
	StoreProduct
	Title        string        `json:"title"`
	Stickers     []BaseSticker `json:"stickers"`
	Icon         []BaseImage   `json:"icon"`
	Previews     []BaseImage   `json:"previews"`
	HasAnimation BaseBoolInt   `json:"has_animation"`
}

// StoreStickersDictionary struct.
type StoreStickersDictionary struct {
	Words            []string       `json:"words"`
	UserStickers     []BaseSticker  `json:"user_stickers,omitempty"`
	PromotedStickers []BaseSticker  `json:"promoted_stickers,omitempty"` // all_products=1
	Stickers         []StoreSticker `json:"sticker,omitempty"`           // need_stickers=0
}

// StoreSticker struct.
type StoreSticker struct {
	PackID    int `json:"pack_id"`
	StickerID int `json:"sticker_id"`
}
