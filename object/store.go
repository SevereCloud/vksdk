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
