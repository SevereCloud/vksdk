package object // import "github.com/severecloud/vksdk/object"

// DatabaseCity struct
type DatabaseCity struct {
	ID        int    `json:"id"`
	Title     string `json:"title"`
	Area      string `json:"area"`
	Region    string `json:"region"`
	Important int    `json:"important"`
}

// DatabaseMetroStation  struct
type DatabaseMetroStation struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Color string `json:"color"`
}
