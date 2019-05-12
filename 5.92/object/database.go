package object // import "github.com/SevereCloud/vksdk/5.92/object"

// DatabaseCity struct
type DatabaseCity struct {
	ID        int    `json:"id"`    // City ID
	Title     string `json:"title"` // City title
	Area      string `json:"area"`
	Region    string `json:"region"`
	Important int    `json:"important"`
}

// DatabaseMetroStation  struct
type DatabaseMetroStation struct {
	ID    int    `json:"id"`    // Metro station ID
	Name  string `json:"name"`  // Metro station name
	Color string `json:"color"` // Metro station color
}

// databaseFaculty
type databaseFaculty struct {
	ID    int    `json:"id"`    // Faculty ID
	Title string `json:"title"` // Faculty title
}

type databaseRegion struct {
	ID    int    `json:"id"`    // Region ID
	Title string `json:"title"` // Region title
}

type databaseSchool struct {
	ID    int    `json:"id"`    // School ID
	Title string `json:"title"` // School title
}

type databaseStation struct {
	CityID int    `json:"city_id"` // City ID
	Color  string `json:"color"`   // Hex color code without #
	ID     int    `json:"id"`      // Station ID
	Name   string `json:"name"`    // Station name
}

type databaseUniversity struct {
	ID    int    `json:"id"`    // University ID
	Title string `json:"title"` // University title
}
