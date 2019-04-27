package object // import "github.com/severecloud/vksdk/5.92/object"

// UtilsDomainResolved struct
type UtilsDomainResolved struct {
	ObjectID int                     `json:"object_id"`
	Type     utilsDomainResolvedType `json:"type"`
}

type utilsDomainResolvedType string

// UtilsLastShortenedLink struct
type UtilsLastShortenedLink struct {
	AccessKey string `json:"access_key"`
	Key       string `json:"key"`
	ShortURL  string `json:"short_url"`
	Timestamp int    `json:"timestamp"`
	URL       string `json:"url"`
	Views     int    `json:"views"`
}

// UtilsLinkChecked struct
type UtilsLinkChecked struct {
	Link   string `json:"link"`
	Status string `json:"status"`
}

// UtilsLinkStats struct
type UtilsLinkStats struct {
	Key   string       `json:"key"`
	Stats []utilsStats `json:"stats"`
}

// UtilsLinkStatsExtended struct
type UtilsLinkStatsExtended struct {
	Key   string               `json:"key"`
	Stats []utilsStatsExtended `json:"stats"`
}

// UtilsShortLink struct
type UtilsShortLink struct {
	AccessKey string `json:"access_key"`
	Key       string `json:"key"`
	ShortURL  string `json:"short_url"`
	URL       string `json:"url"`
}

type utilsStats struct {
	Timestamp int `json:"timestamp"`
	Views     int `json:"views"`
}

type utilsStatsCity struct {
	CityID int `json:"city_id"`
	Views  int `json:"views"`
}

type utilsStatsCountry struct {
	CountryID int `json:"country_id"`
	Views     int `json:"views"`
}

type utilsStatsExtended struct {
	Cities    []utilsStatsCity    `json:"cities"`
	Countries []utilsStatsCountry `json:"countries"`
	SexAge    []utilsStatsSexAge  `json:"sex_age"`
	Timestamp int                 `json:"timestamp"`
	Views     int                 `json:"views"`
}

type utilsStatsSexAge struct {
	AgeRange string `json:"age_range"`
	Female   int    `json:"female"`
	Male     int    `json:"male"`
}
