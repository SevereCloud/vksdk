package object

type FaveTag struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type FavePage struct {
	Type        string      `json:"type"`
	Description string      `json:"description"`
	Tags        []FaveTag   `json:"tags"`
	UpdatedDate int         `json:"updated_date"`
	User        UsersUser   `json:"user"`
	Group       GroupsGroup `json:"group"`
}

type FaveFavesLink struct {
	URL         string      `json:"url"`
	Title       string      `json:"title"`
	Caption     string      `json:"caption"`
	Description string      `json:"description"`
	Photo       PhotosPhoto `json:"photo"`
	IsFavorite  bool        `json:"is_favorite"`
	ID          string      `json:"id"`
}

type FaveItem struct {
	Type      string           `json:"type"`
	Seen      bool             `json:"seen"`
	AddedDate int              `json:"added_date"`
	Tags      []FaveTag        `json:"tags"`
	Link      FaveFavesLink    `json:"link,omitempty"`
	Post      WallWallpost     `json:"post,omitempty"`
	Video     VideoVideoFull   `json:"video,omitempty"`
	Product   MarketMarketItem `json:"product,omitempty"`
	Article   Article          `json:"article,omitempty"`
}
