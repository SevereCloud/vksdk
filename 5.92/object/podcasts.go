package object

type PodcastsItem struct {
	OwnerID int `json:"owner_id"`
}

type PodcastsCategory struct {
	ID    int         `json:"id"`
	Title string      `json:"title"`
	Cover []baseImage `json:"cover"`
}

type PodcastsEpisode struct {
	ID           int                 `json:"id"`
	OwnerID      int                 `json:"owner_id"`
	Artist       string              `json:"artist"`
	Title        string              `json:"title"`
	Duration     int                 `json:"duration"`
	Date         int                 `json:"date"`
	URL          string              `json:"url"`
	LyricsID     int                 `json:"lyrics_id"`
	NoSearch     int                 `json:"no_search"`
	TrackCode    string              `json:"track_code"`
	IsHq         bool                `json:"is_hq"`
	IsFocusTrack bool                `json:"is_focus_track"`
	IsExplicit   bool                `json:"is_explicit"`
	PodcastInfo  PodcastsPodcastInfo `json:"podcast_info"`
}

type PodcastsPodcastInfo struct {
	Cover struct {
		Sizes []baseImage `json:"cover"`
	}
	Plays       int    `json:"plays"`
	IsFavorite  bool   `json:"is_favorite"`
	Description string `json:"description"`
	Position    int    `json:"position"`
}
