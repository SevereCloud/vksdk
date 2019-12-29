package object // import "github.com/SevereCloud/vksdk/object"

import (
	"fmt"
)

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
	IsExplicit  bool               `json:"is_explicit"`
	LyricsID    int                `json:"lyrics_id"`
	AlbumID     int                `json:"album_id"`
	GenreID     int                `json:"genre_id"`
	TrackCode   string             `json:"track_code"`
	NoSearch    int                `json:"no_search"`
	MainArtists []AudioAudioArtist `json:"main_artists"`
}

// ToAttachment return attachment format
func (audio AudioAudioFull) ToAttachment() string {
	return fmt.Sprintf("audio%d_%d", audio.OwnerID, audio.ID)
}

// AudioAudioArtist struct
type AudioAudioArtist struct {
	Name   string `json:"name"`
	ID     string `json:"id"`
	Domain string `json:"domain"`
}

// AudioAudio struct
type AudioAudio struct {
	AccessKey    string `json:"access_key"` // Access key for the audio
	Artist       string `json:"artist"`     // Artist name
	ID           int    `json:"id"`         // Audio ID
	IsExplicit   bool   `json:"is_explicit"`
	IsFocusTrack bool   `json:"is_focus_track"`
	IsLicensed   bool   `json:"is_licensed"`
	OwnerID      int    `json:"owner_id"` // Audio owner's ID
	Title        string `json:"title"`    // Title
	URL          string `json:"url"`      // URL of mp3 file
}

// ToAttachment return attachment format
func (audio AudioAudio) ToAttachment() string {
	return fmt.Sprintf("audio%d_%d", audio.OwnerID, audio.ID)
}

// AudioAudioUploadResponse struct
type AudioAudioUploadResponse struct {
	Audio    string `json:"audio"`
	Hash     string `json:"hash"`
	Redirect string `json:"redirect"`
	Server   int    `json:"server"`
}

// AudioLyrics struct
type AudioLyrics struct {
	LyricsID int    `json:"lyrics_id"`
	Text     string `json:"text"`
}
