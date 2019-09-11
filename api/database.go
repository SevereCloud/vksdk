package api // import "github.com/SevereCloud/vksdk/api"

import (
	"github.com/SevereCloud/vksdk/object"
)

// DatabaseGetChairsResponse struct
type DatabaseGetChairsResponse struct {
	Count int                 `json:"count"`
	Items []object.BaseObject `json:"items"`
}

// DatabaseGetChairs returns list of chairs on a specified faculty.
//
// https://vk.com/dev/database.getChairs
func (vk *VK) DatabaseGetChairs(params map[string]string) (response DatabaseGetChairsResponse, err error) {
	err = vk.RequestUnmarshal("database.getChairs", params, &response)
	return
}

// DatabaseGetCitiesResponse struct
type DatabaseGetCitiesResponse struct {
	Count int                   `json:"count"`
	Items []object.DatabaseCity `json:"items"`
}

// DatabaseGetCities returns a list of cities
//
// https://vk.com/dev/database.getCities
func (vk *VK) DatabaseGetCities(params map[string]string) (response DatabaseGetCitiesResponse, err error) {
	err = vk.RequestUnmarshal("database.getCities", params, &response)
	return
}

// DatabaseGetCitiesByIDResponse struct
type DatabaseGetCitiesByIDResponse []object.DatabaseCity

// DatabaseGetCitiesByID returns information about cities by their IDs.
//
// https://vk.com/dev/database.getCitiesByID
func (vk *VK) DatabaseGetCitiesByID(params map[string]string) (response DatabaseGetCitiesByIDResponse, err error) {
	err = vk.RequestUnmarshal("database.getCitiesById", params, &response)
	return
}

// DatabaseGetCountriesResponse struct
type DatabaseGetCountriesResponse struct {
	Count int                 `json:"count"`
	Items []object.BaseObject `json:"items"`
}

// DatabaseGetCountries returns a list of countries.
//
// https://vk.com/dev/database.getCountries
func (vk *VK) DatabaseGetCountries(params map[string]string) (response DatabaseGetCountriesResponse, err error) {
	err = vk.RequestUnmarshal("database.getCountries", params, &response)
	return
}

// DatabaseGetCountriesByIDResponse struct
type DatabaseGetCountriesByIDResponse []object.BaseObject

// DatabaseGetCountriesByID returns information about countries by their IDs.
//
// https://vk.com/dev/database.getCountriesByID
func (vk *VK) DatabaseGetCountriesByID(params map[string]string) (response DatabaseGetCountriesByIDResponse, err error) {
	err = vk.RequestUnmarshal("database.getCountriesById", params, &response)
	return
}

// DatabaseGetFacultiesResponse struct
type DatabaseGetFacultiesResponse struct {
	Count int                      `json:"count"`
	Items []object.DatabaseFaculty `json:"items"`
}

// DatabaseGetFaculties returns a list of faculties (i.e., university departments).
//
// https://vk.com/dev/database.getFaculties
func (vk *VK) DatabaseGetFaculties(params map[string]string) (response DatabaseGetFacultiesResponse, err error) {
	err = vk.RequestUnmarshal("database.getFaculties", params, &response)
	return
}

// DatabaseGetMetroStationsResponse struct
type DatabaseGetMetroStationsResponse struct {
	Count int                           `json:"count"`
	Items []object.DatabaseMetroStation `json:"items"`
}

// DatabaseGetMetroStations returns the list of metro stations.
//
// https://vk.com/dev/database.getMetroStations
func (vk *VK) DatabaseGetMetroStations(params map[string]string) (response DatabaseGetMetroStationsResponse, err error) {
	err = vk.RequestUnmarshal("database.getMetroStations", params, &response)
	return
}

// DatabaseGetMetroStationsByIDResponse struct
type DatabaseGetMetroStationsByIDResponse []object.DatabaseMetroStation

// DatabaseGetMetroStationsByID returns information about one or several metro stations by their identifiers.
//
// https://vk.com/dev/database.getMetroStationsById
func (vk *VK) DatabaseGetMetroStationsByID(params map[string]string) (response DatabaseGetMetroStationsByIDResponse, err error) {
	err = vk.RequestUnmarshal("database.getMetroStationsById", params, &response)
	return
}

// DatabaseGetRegionsResponse struct
type DatabaseGetRegionsResponse struct {
	Count int                     `json:"count"`
	Items []object.DatabaseRegion `json:"items"`
}

// DatabaseGetRegions returns a list of regions.
//
// https://vk.com/dev/database.getRegions
func (vk *VK) DatabaseGetRegions(params map[string]string) (response DatabaseGetRegionsResponse, err error) {
	err = vk.RequestUnmarshal("database.getRegions", params, &response)
	return
}

// DatabaseGetSchoolClassesResponse struct
type DatabaseGetSchoolClassesResponse [][]interface{}

// DatabaseGetSchoolClasses returns a list of school classes specified for the country.
// BUG(VK): database.getSchoolClasses  bad return
//
// https://vk.com/dev/database.getSchoolClasses
func (vk *VK) DatabaseGetSchoolClasses(params map[string]string) (response DatabaseGetSchoolClassesResponse, err error) {
	err = vk.RequestUnmarshal("database.getSchoolClasses", params, &response)
	return
}

// DatabaseGetSchoolsResponse struct
type DatabaseGetSchoolsResponse struct {
	Count int                     `json:"count"`
	Items []object.DatabaseSchool `json:"items"`
}

// DatabaseGetSchools returns a list of schools.
//
// https://vk.com/dev/database.getSchools
func (vk *VK) DatabaseGetSchools(params map[string]string) (response DatabaseGetSchoolsResponse, err error) {
	err = vk.RequestUnmarshal("database.getSchools", params, &response)
	return
}

// DatabaseGetUniversitiesResponse struct
type DatabaseGetUniversitiesResponse struct {
	Count int                         `json:"count"`
	Items []object.DatabaseUniversity `json:"items"`
}

// DatabaseGetUniversities returns a list of higher education institutions.
//
// https://vk.com/dev/database.getUniversities
func (vk *VK) DatabaseGetUniversities(params map[string]string) (response DatabaseGetUniversitiesResponse, err error) {
	err = vk.RequestUnmarshal("database.getUniversities", params, &response)
	return
}
