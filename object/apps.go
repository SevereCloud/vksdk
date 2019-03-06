package object // import "github.com/severecloud/vksdk/object"

// AppsApp struct
type AppsApp struct {
	ID              int           `json:"id"`
	Title           string        `json:"title"`
	Icon278         string        `json:"icon_278"`
	Icon139         string        `json:"icon_139"`
	Icon150         string        `json:"icon_150"`
	Icon75          string        `json:"icon_75"`
	Icon16          string        `json:"icon_16"`
	Banner560       string        `json:"banner_560"`
	Banner1120      string        `json:"banner_1120"`
	Description     string        `json:"description"`
	ScreenName      string        `json:"screen_name"`
	Section         string        `json:"section"`
	Type            string        `json:"type"`
	AuthorURL       string        `json:"author_url"`
	AuthorGroup     int           `json:"author_group"`
	MembersCount    int           `json:"members_count"`
	PublishedDate   int           `json:"published_date"`
	GenreID         int           `json:"genre_id"`
	Genre           string        `json:"genre"`
	PushEnabled     int           `json:"push_enabled"`
	LeaderboardType int           `json:"leaderboard_type"`
	IsInCatalog     int           `json:"is_in_catalog"`
	CatalogPosition int           `json:"catalog_position"`
	AuthorID        int           `json:"author_id"`
	International   int           `json:"international"`
	PlatformID      int           `json:"platform_id"`
	Screenshots     []PhotosPhoto `json:"screenshots"`
	Friends         []int         `json:"friends"`
}

type appsLeaderboard struct {
	Level  int `json:"level"`
	Points int `json:"points"`
	Score  int `json:"score"`
	UserID int `json:"user_id"`
}
