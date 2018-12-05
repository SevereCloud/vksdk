package object

// AudioAudioFull struct
type AudioAudioFull struct {
	ID          int                `json:"id"`
	OwnerID     int                `json:"owner_id"`
	Artist      string             `json:"artist"`
	Title       string             `json:"title"`
	Duration    int                `json:"duration"`
	Date        int                `json:"date"`
	URL         string             `json:"url"`
	IsHq        bool               `json:"is_hq"`
	LyricsID    int                `json:"lyrics_id"`
	AlbumID     int                `json:"album_id"`
	GenreID     int                `json:"genre_id"`
	TrackCode   string             `json:"track_code"`
	IsExplicit  bool               `json:"is_explicit"`
	NoSearch    int                `json:"no_search"`
	MainArtists []audioAudioArtist `json:"main_artists"`
}

type audioAudioArtist struct {
	Name   string `json:"name"`
	ID     string `json:"id"`
	Domain string `json:"domain"`
}

type audioAudio struct {
	AccessKey string `json:"access_key"`
	Artist    string `json:"artist"`
	ID        int    `json:"id"`
	OwnerID   int    `json:"owner_id"`
	Title     string `json:"title"`
	URL       string `json:"url"`
}

type audioAudioUploadResponse struct {
	Audio    string `json:"audio"`
	Hash     string `json:"hash"`
	Redirect string `json:"redirect"`
	Server   int    `json:"server"`
}

type audioLyrics struct {
	LyricsID int    `json:"lyrics_id"`
	Text     string `json:"text"`
}
