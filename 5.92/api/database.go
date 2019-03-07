package api // import "github.com/severecloud/vksdk/5.92/api"

import (
	"encoding/json"

	"github.com/severecloud/vksdk/5.92/object"
)

// DatabaseGetChairsResponse struct
type DatabaseGetChairsResponse struct {
	Count int                 `json:"count"`
	Items []object.BaseObject `json:"items"`
}

// DatabaseGetChairs returns list of chairs on a specified faculty.
// https://vk.com/dev/database.getChairs
func (vk VK) DatabaseGetChairs(params map[string]string) (response DatabaseGetChairsResponse, vkErr Error) {
	rawResponse, vkErr := vk.Request("database.getChairs", params)
	if vkErr.Code != 0 {
		return
	}

	err := json.Unmarshal(rawResponse, &response)
	if err != nil {
		panic(err)
	}

	return
}

// DatabaseGetCitiesResponse struct
type DatabaseGetCitiesResponse struct {
	Count int                   `json:"count"`
	Items []object.DatabaseCity `json:"items"`
}

// DatabaseGetCities returns a list of cities
// https://vk.com/dev/database.getCities
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
type DatabaseGetCitiesByIDResponse []object.BaseObject

// DatabaseGetCitiesByID returns information about cities by their IDs.
// https://vk.com/dev/database.getCitiesByID
func (vk VK) DatabaseGetCitiesByID(params map[string]string) (response DatabaseGetCitiesByIDResponse, vkErr Error) {
	rawResponse, vkErr := vk.Request("database.getCitiesByID", params)
	if vkErr.Code != 0 {
		return
	}

	err := json.Unmarshal(rawResponse, &response)
	if err != nil {
		panic(err)
	}

	return
}

// DatabaseGetCountriesResponse struct
type DatabaseGetCountriesResponse struct {
	Count int                 `json:"count"`
	Items []object.BaseObject `json:"items"`
}

// DatabaseGetCountries returns a list of countries.
// https://vk.com/dev/database.getCountries
func (vk VK) DatabaseGetCountries(params map[string]string) (response DatabaseGetCountriesResponse, vkErr Error) {
	rawResponse, vkErr := vk.Request("database.getCountries", params)
	if vkErr.Code != 0 {
		return
	}

	err := json.Unmarshal(rawResponse, &response)
	if err != nil {
		panic(err)
	}

	return
}

// DatabaseGetCountriesByIDResponse struct
type DatabaseGetCountriesByIDResponse []object.BaseObject

// DatabaseGetCountriesByID returns information about countries by their IDs.
// https://vk.com/dev/database.getCountriesByID
func (vk VK) DatabaseGetCountriesByID(params map[string]string) (response DatabaseGetCountriesByIDResponse, vkErr Error) {
	rawResponse, vkErr := vk.Request("database.getCountriesByID", params)
	if vkErr.Code != 0 {
		return
	}

	err := json.Unmarshal(rawResponse, &response)
	if err != nil {
		panic(err)
	}

	return
}

// DatabaseGetFacultiesResponse struct
type DatabaseGetFacultiesResponse struct {
	Count int                 `json:"count"`
	Items []object.BaseObject `json:"items"`
}

// DatabaseGetFaculties returns a list of faculties (i.e., university departments).
// https://vk.com/dev/database.getFaculties
func (vk VK) DatabaseGetFaculties(params map[string]string) (response DatabaseGetFacultiesResponse, vkErr Error) {
	rawResponse, vkErr := vk.Request("database.getFaculties", params)
	if vkErr.Code != 0 {
		return
	}

	err := json.Unmarshal(rawResponse, &response)
	if err != nil {
		panic(err)
	}

	return
}

// DatabaseGetMetroStationsResponse struct
type DatabaseGetMetroStationsResponse struct {
	Count int                           `json:"count"`
	Items []object.DatabaseMetroStation `json:"items"`
}

// DatabaseGetMetroStations returns the list of metro stations.
// https://vk.com/dev/database.getMetroStations
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
// https://vk.com/dev/database.getMetroStationsById
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
type DatabaseGetRegionsResponse struct {
	Count int                 `json:"count"`
	Items []object.BaseObject `json:"items"`
}

// DatabaseGetRegions returns a list of regions.
// https://vk.com/dev/database.getRegions
func (vk VK) DatabaseGetRegions(params map[string]string) (response DatabaseGetRegionsResponse, vkErr Error) {
	rawResponse, vkErr := vk.Request("database.getRegions", params)
	if vkErr.Code != 0 {
		return
	}

	err := json.Unmarshal(rawResponse, &response)
	if err != nil {
		panic(err)
	}

	return
}

// DatabaseGetSchoolClassesResponse struct
type DatabaseGetSchoolClassesResponse struct{}

// TODO: database.getSchoolClasses returns a list of school classes specified for the country.
// BUG(VK): database.getSchoolClasses  bad return
// https://vk.com/dev/database.getSchoolClasses

// DatabaseGetSchoolsResponse struct
type DatabaseGetSchoolsResponse struct {
	Count int                 `json:"count"`
	Items []object.BaseObject `json:"items"`
}

// DatabaseGetSchools returns a list of schools.
// https://vk.com/dev/database.getSchools
func (vk VK) DatabaseGetSchools(params map[string]string) (response DatabaseGetSchoolsResponse, vkErr Error) {
	rawResponse, vkErr := vk.Request("database.getSchools", params)
	if vkErr.Code != 0 {
		return
	}

	err := json.Unmarshal(rawResponse, &response)
	if err != nil {
		panic(err)
	}

	return
}

// DatabaseGetUniversitiesResponse struct
type DatabaseGetUniversitiesResponse struct {
	Count int                 `json:"count"`
	Items []object.BaseObject `json:"items"`
}

// DatabaseGetUniversities returns a list of higher education institutions.
// https://vk.com/dev/database.getUniversities
func (vk VK) DatabaseGetUniversities(params map[string]string) (response DatabaseGetUniversitiesResponse, vkErr Error) {
	rawResponse, vkErr := vk.Request("database.getUniversities", params)
	if vkErr.Code != 0 {
		return
	}

	err := json.Unmarshal(rawResponse, &response)
	if err != nil {
		panic(err)
	}

	return
}
