package object

type appsApp struct {
	AuthorGroup     int           `json:"author_group"`
	AuthorID        int           `json:"author_id"`
	AuthorURL       string        `json:"author_url"`
	Banner1120      string        `json:"banner_1120"`
	Banner560       string        `json:"banner_560"`
	CatalogPosition int           `json:"catalog_position"`
	Description     string        `json:"description"`
	Genre           string        `json:"genre"`
	GenreID         int           `json:"genre_id"`
	ID              int           `json:"id"`
	Icon139         string        `json:"icon_139"`
	Icon150         string        `json:"icon_150"`
	Icon278         string        `json:"icon_278"`
	Icon75          string        `json:"icon_75"`
	International   int           `json:"international"`
	IsInCatalog     int           `json:"is_in_catalog"`
	LeaderboardType int           `json:"leaderboard_type"`
	MembersCount    int           `json:"members_count"`
	PlatformID      int           `json:"platform_id"`
	PublishedDate   int           `json:"published_date"`
	ScreenName      string        `json:"screen_name"`
	Screenshots     []PhotosPhoto `json:"screenshots"`
	Section         string        `json:"section"`
	Title           string        `json:"title"`
	Type            string        `json:"type"`
}

type appsLeaderboard struct {
	Level  int `json:"level"`
	Points int `json:"points"`
	Score  int `json:"score"`
	UserID int `json:"user_id"`
}
