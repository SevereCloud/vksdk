package api // import "github.com/severecloud/vksdk/api"

import (
	"encoding/json"

	"github.com/severecloud/vksdk/object"
)

// DatabaseGetChairsResponse struct
type DatabaseGetChairsResponse struct{}

// TODO: database.getChairs returns list of chairs on a specified faculty.

// DatabaseGetCitiesResponse struct
type DatabaseGetCitiesResponse struct {
	Count int `json:"count"`
	Items []struct {
		ID     int    `json:"id"`
		Title  string `json:"title"`
		Region string `json:"region"`
	} `json:"items"`
}

// DatabaseGetCities returns a list of cities
// TODO: testing dataxase.getCities
func (vk VK) DatabaseGetCities(params map[string]string) (response DatabaseGetCitiesResponse, vkErr Error) {
	rawResponse, vkErr := vk.Request("database.getCities", params)
	if vkErr.Code != 0 {
		return
	}

	err := json.Unmarshal(rawResponse, &response)
	if err != nil {
		panic(err)
	}

	return
}

// DatabaseGetCitiesByIDResponse struct
type DatabaseGetCitiesByIDResponse struct{}

// TODO: database.getCitiesByID returns information about cities by their IDs.

// DatabaseGetCountriesResponse struct
type DatabaseGetCountriesResponse struct{}

// TODO: database.getCountries returns a list of countries.

// DatabaseGetCountriesByIDResponse struct
type DatabaseGetCountriesByIDResponse struct{}

// TODO: database.getCountriesByID returns information about countries by their IDs.

// DatabaseGetFacultiesResponse struct
type DatabaseGetFacultiesResponse struct{}

// TODO: database.getFaculties returns a list of faculties (i.e., university departments).

// DatabaseGetMetroStationsResponse struct
type DatabaseGetMetroStationsResponse struct {
	Count int                           `json:"count"`
	Items []object.DatabaseMetroStation `json:"items"`
}

// DatabaseGetMetroStations returns the list of metro stations.
func (vk VK) DatabaseGetMetroStations(params map[string]string) (response DatabaseGetMetroStationsResponse, vkErr Error) {
	rawResponse, vkErr := vk.Request("database.getMetroStations", params)
	if vkErr.Code != 0 {
		return
	}

	err := json.Unmarshal(rawResponse, &response)
	if err != nil {
		panic(err)
	}

	return
}

// DatabaseGetMetroStationsByIDResponse struct
type DatabaseGetMetroStationsByIDResponse []object.DatabaseMetroStation

// DatabaseGetMetroStationsByID returns information about one or several metro stations by their identifiers.
func (vk VK) DatabaseGetMetroStationsByID(params map[string]string) (response DatabaseGetMetroStationsByIDResponse, vkErr Error) {
	rawResponse, vkErr := vk.Request("database.getMetroStationsById", params)
	if vkErr.Code != 0 {
		return
	}

	err := json.Unmarshal(rawResponse, &response)
	if err != nil {
		panic(err)
	}

	return
}

// DatabaseGetRegionsResponse struct
type DatabaseGetRegionsResponse struct{}

// TODO: database.getRegions returns a list of regions.

// DatabaseGetSchoolClassesResponse struct
type DatabaseGetSchoolClassesResponse struct{}

// TODO: database.getSchoolClasses returns a list of school classes specified for the country.

// DatabaseGetSchoolsResponse struct
type DatabaseGetSchoolsResponse struct{}

// TODO: database.getSchools returns a list of schools.

// DatabaseGetUniversitiesResponse struct
type DatabaseGetUniversitiesResponse struct{}

// TODO: database.getUniversities returns a list of higher education institutions.
