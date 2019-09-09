package object // import "github.com/SevereCloud/vksdk/object"

// UtilsDomainResolved struct
type UtilsDomainResolved struct {
	ObjectID int    `json:"object_id"` // Object ID
	Type     string `json:"type"`
}

// UtilsLastShortenedLink struct
type UtilsLastShortenedLink struct {
	AccessKey string `json:"access_key"` // Access key for private stats
	Key       string `json:"key"`        // Link key (characters after vk.cc/)
	ShortURL  string `json:"short_url"`  // Short link URL
	Timestamp int    `json:"timestamp"`  // Creation time in Unixtime
	URL       string `json:"url"`        // Full URL
	Views     int    `json:"views"`      // Total views number
}

// UtilsLinkChecked struct
type UtilsLinkChecked struct {
	Link   string `json:"link"` // Link URL
	Status string `json:"status"`
}

// UtilsLinkStats struct
type UtilsLinkStats struct {
	Key   string       `json:"key"` // Link key (characters after vk.cc/)
	Stats []utilsStats `json:"stats"`
}

// UtilsLinkStatsExtended struct
type UtilsLinkStatsExtended struct {
	Key   string               `json:"key"` // Link key (characters after vk.cc/)
	Stats []utilsStatsExtended `json:"stats"`
}

// UtilsShortLink struct
type UtilsShortLink struct {
	AccessKey string `json:"access_key"` // Access key for private stats
	Key       string `json:"key"`        // Link key (characters after vk.cc/)
	ShortURL  string `json:"short_url"`  // Short link URL
	URL       string `json:"url"`        // Full URL
}

type utilsStats struct {
	Timestamp int `json:"timestamp"` // Start time
	Views     int `json:"views"`     // Total views number
}

type utilsStatsCity struct {
	CityID int `json:"city_id"` // City ID
	Views  int `json:"views"`   // Views number
}

type utilsStatsCountry struct {
	CountryID int `json:"country_id"` // Country ID
	Views     int `json:"views"`      // Views number
}

type utilsStatsExtended struct {
	Cities    []utilsStatsCity    `json:"cities"`
	Countries []utilsStatsCountry `json:"countries"`
	SexAge    []utilsStatsSexAge  `json:"sex_age"`
	Timestamp int                 `json:"timestamp"` // Start time
	Views     int                 `json:"views"`     // Total views number
}

type utilsStatsSexAge struct {
	AgeRange string `json:"age_range"` // Age denotation
	Female   int    `json:"female"`    //  Views by female users
	Male     int    `json:"male"`      //  Views by male users
}
