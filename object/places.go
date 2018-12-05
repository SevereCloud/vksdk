package object // import "github.com/severecloud/vksdk/object"

type placesCheckin struct {
	Date         int     `json:"date"`
	Distance     int     `json:"distance"`
	ID           int     `json:"id"`
	Latitude     float64 `json:"latitude"`
	Longitude    float64 `json:"longitude"`
	PlaceCity    int     `json:"place_city"`
	PlaceCountry int     `json:"place_country"`
	PlaceID      int     `json:"place_id"`
	PlaceIcon    string  `json:"place_icon"`
	PlaceTitle   string  `json:"place_title"`
	PlaceType    string  `json:"place_type"`
	Text         string  `json:"text"`
	UserID       int     `json:"user_id"`
}

type placesPlaceFull struct {
	Address    string  `json:"address"`
	Checkins   int     `json:"checkins"`
	City       int     `json:"city"`
	Country    int     `json:"country"`
	Created    int     `json:"created"`
	Distance   int     `json:"distance"`
	GroupID    int     `json:"group_id"`
	GroupPhoto string  `json:"group_photo"`
	ID         int     `json:"id"`
	Icon       string  `json:"icon"`
	Latitude   float64 `json:"latitude"`
	Longitude  float64 `json:"longitude"`
	Title      string  `json:"title"`
	Type       string  `json:"type"`
}

type placesPlaceMin struct {
	Address   string  `json:"address"`
	Checkins  int     `json:"checkins"`
	City      int     `json:"city"`
	Country   int     `json:"country"`
	Created   int     `json:"created"`
	ID        int     `json:"id"`
	Icon      string  `json:"icon"`
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
	Title     string  `json:"title"`
	Type      string  `json:"type"`
}

type placesTypes struct {
	ID    int    `json:"id"`
	Icon  string `json:"icon"`
	Title string `json:"title"`
}
