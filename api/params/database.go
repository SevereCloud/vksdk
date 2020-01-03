package params // import "github.com/SevereCloud/vksdk/api/params"

import (
	"github.com/SevereCloud/vksdk/api"
)

// DatabaseGetChairsBulder builder
//
// Returns list of chairs on a specified faculty.
//
// https://vk.com/dev/database.getChairs
type DatabaseGetChairsBulder struct {
	api.Params
}

// NewDatabaseGetChairsBulder func
func NewDatabaseGetChairsBulder() *DatabaseGetChairsBulder {
	return &DatabaseGetChairsBulder{api.Params{}}
}

// FacultyID id of the faculty to get chairs from
func (b *DatabaseGetChairsBulder) FacultyID(v int) {
	b.Params["faculty_id"] = v
}

// Offset offset required to get a certain subset of chairs
func (b *DatabaseGetChairsBulder) Offset(v int) {
	b.Params["offset"] = v
}

// Count amount of chairs to get
func (b *DatabaseGetChairsBulder) Count(v int) {
	b.Params["count"] = v
}

// DatabaseGetCitiesBulder builder
//
// Returns a list of cities.
//
// https://vk.com/dev/database.getCities
type DatabaseGetCitiesBulder struct {
	api.Params
}

// NewDatabaseGetCitiesBulder func
func NewDatabaseGetCitiesBulder() *DatabaseGetCitiesBulder {
	return &DatabaseGetCitiesBulder{api.Params{}}
}

// CountryID Country ID.
func (b *DatabaseGetCitiesBulder) CountryID(v int) {
	b.Params["country_id"] = v
}

// RegionID Region ID.
func (b *DatabaseGetCitiesBulder) RegionID(v int) {
	b.Params["region_id"] = v
}

// Q Search query.
func (b *DatabaseGetCitiesBulder) Q(v string) {
	b.Params["q"] = v
}

// NeedAll '1' — to return all cities in the country, '0' — to return major cities in the country (default),
func (b *DatabaseGetCitiesBulder) NeedAll(v bool) {
	b.Params["need_all"] = v
}

// Offset Offset needed to return a specific subset of cities.
func (b *DatabaseGetCitiesBulder) Offset(v int) {
	b.Params["offset"] = v
}

// Count Number of cities to return.
func (b *DatabaseGetCitiesBulder) Count(v int) {
	b.Params["count"] = v
}

// DatabaseGetCitiesByIDBulder builder
//
// Returns information about cities by their IDs.
//
// https://vk.com/dev/database.getCitiesById
type DatabaseGetCitiesByIDBulder struct {
	api.Params
}

// NewDatabaseGetCitiesByIDBulder func
func NewDatabaseGetCitiesByIDBulder() *DatabaseGetCitiesByIDBulder {
	return &DatabaseGetCitiesByIDBulder{api.Params{}}
}

// CityIDs City IDs.
func (b *DatabaseGetCitiesByIDBulder) CityIDs(v []int) {
	b.Params["city_ids"] = v
}

// DatabaseGetCountriesBulder builder
//
// Returns a list of countries.
//
// https://vk.com/dev/database.getCountries
type DatabaseGetCountriesBulder struct {
	api.Params
}

// NewDatabaseGetCountriesBulder func
func NewDatabaseGetCountriesBulder() *DatabaseGetCountriesBulder {
	return &DatabaseGetCountriesBulder{api.Params{}}
}

// NeedAll '1' — to return a full list of all countries, '0' — to return a list of countries near the current user's country (default).
func (b *DatabaseGetCountriesBulder) NeedAll(v bool) {
	b.Params["need_all"] = v
}

// Code Country codes in [vk.com/dev/country_codes|ISO 3166-1 alpha-2] standard.
func (b *DatabaseGetCountriesBulder) Code(v string) {
	b.Params["code"] = v
}

// Offset Offset needed to return a specific subset of countries.
func (b *DatabaseGetCountriesBulder) Offset(v int) {
	b.Params["offset"] = v
}

// Count Number of countries to return.
func (b *DatabaseGetCountriesBulder) Count(v int) {
	b.Params["count"] = v
}

// DatabaseGetCountriesByIDBulder builder
//
// Returns information about countries by their IDs.
//
// https://vk.com/dev/database.getCountriesById
type DatabaseGetCountriesByIDBulder struct {
	api.Params
}

// NewDatabaseGetCountriesByIDBulder func
func NewDatabaseGetCountriesByIDBulder() *DatabaseGetCountriesByIDBulder {
	return &DatabaseGetCountriesByIDBulder{api.Params{}}
}

// CountryIDs Country IDs.
func (b *DatabaseGetCountriesByIDBulder) CountryIDs(v []int) {
	b.Params["country_ids"] = v
}

// DatabaseGetFacultiesBulder builder
//
// Returns a list of faculties (i.e., university departments).
//
// https://vk.com/dev/database.getFaculties
type DatabaseGetFacultiesBulder struct {
	api.Params
}

// NewDatabaseGetFacultiesBulder func
func NewDatabaseGetFacultiesBulder() *DatabaseGetFacultiesBulder {
	return &DatabaseGetFacultiesBulder{api.Params{}}
}

// UniversityID University ID.
func (b *DatabaseGetFacultiesBulder) UniversityID(v int) {
	b.Params["university_id"] = v
}

// Offset Offset needed to return a specific subset of faculties.
func (b *DatabaseGetFacultiesBulder) Offset(v int) {
	b.Params["offset"] = v
}

// Count Number of faculties to return.
func (b *DatabaseGetFacultiesBulder) Count(v int) {
	b.Params["count"] = v
}

// DatabaseGetMetroStationsBulder builder
//
// Get metro stations by city
//
// https://vk.com/dev/database.getMetroStations
type DatabaseGetMetroStationsBulder struct {
	api.Params
}

// NewDatabaseGetMetroStationsBulder func
func NewDatabaseGetMetroStationsBulder() *DatabaseGetMetroStationsBulder {
	return &DatabaseGetMetroStationsBulder{api.Params{}}
}

// CityID parameter
func (b *DatabaseGetMetroStationsBulder) CityID(v int) {
	b.Params["city_id"] = v
}

// Offset parameter
func (b *DatabaseGetMetroStationsBulder) Offset(v int) {
	b.Params["offset"] = v
}

// Count parameter
func (b *DatabaseGetMetroStationsBulder) Count(v int) {
	b.Params["count"] = v
}

// Extended parameter
func (b *DatabaseGetMetroStationsBulder) Extended(v bool) {
	b.Params["extended"] = v
}

// DatabaseGetMetroStationsByIDBulder builder
//
// Get metro station by his id
//
// https://vk.com/dev/database.getMetroStationsById
type DatabaseGetMetroStationsByIDBulder struct {
	api.Params
}

// NewDatabaseGetMetroStationsByIDBulder func
func NewDatabaseGetMetroStationsByIDBulder() *DatabaseGetMetroStationsByIDBulder {
	return &DatabaseGetMetroStationsByIDBulder{api.Params{}}
}

// StationIDs parameter
func (b *DatabaseGetMetroStationsByIDBulder) StationIDs(v []int) {
	b.Params["station_ids"] = v
}

// DatabaseGetRegionsBulder builder
//
// Returns a list of regions.
//
// https://vk.com/dev/database.getRegions
type DatabaseGetRegionsBulder struct {
	api.Params
}

// NewDatabaseGetRegionsBulder func
func NewDatabaseGetRegionsBulder() *DatabaseGetRegionsBulder {
	return &DatabaseGetRegionsBulder{api.Params{}}
}

// CountryID Country ID, received in [vk.com/dev/database.getCountries|database.getCountries] method.
func (b *DatabaseGetRegionsBulder) CountryID(v int) {
	b.Params["country_id"] = v
}

// Q Search query.
func (b *DatabaseGetRegionsBulder) Q(v string) {
	b.Params["q"] = v
}

// Offset Offset needed to return specific subset of regions.
func (b *DatabaseGetRegionsBulder) Offset(v int) {
	b.Params["offset"] = v
}

// Count Number of regions to return.
func (b *DatabaseGetRegionsBulder) Count(v int) {
	b.Params["count"] = v
}

// DatabaseGetSchoolClassesBulder builder
//
// Returns a list of school classes specified for the country.
//
// https://vk.com/dev/database.getSchoolClasses
type DatabaseGetSchoolClassesBulder struct {
	api.Params
}

// NewDatabaseGetSchoolClassesBulder func
func NewDatabaseGetSchoolClassesBulder() *DatabaseGetSchoolClassesBulder {
	return &DatabaseGetSchoolClassesBulder{api.Params{}}
}

// CountryID Country ID.
func (b *DatabaseGetSchoolClassesBulder) CountryID(v int) {
	b.Params["country_id"] = v
}

// DatabaseGetSchoolsBulder builder
//
// Returns a list of schools.
//
// https://vk.com/dev/database.getSchools
type DatabaseGetSchoolsBulder struct {
	api.Params
}

// NewDatabaseGetSchoolsBulder func
func NewDatabaseGetSchoolsBulder() *DatabaseGetSchoolsBulder {
	return &DatabaseGetSchoolsBulder{api.Params{}}
}

// Q Search query.
func (b *DatabaseGetSchoolsBulder) Q(v string) {
	b.Params["q"] = v
}

// CityID City ID.
func (b *DatabaseGetSchoolsBulder) CityID(v int) {
	b.Params["city_id"] = v
}

// Offset Offset needed to return a specific subset of schools.
func (b *DatabaseGetSchoolsBulder) Offset(v int) {
	b.Params["offset"] = v
}

// Count Number of schools to return.
func (b *DatabaseGetSchoolsBulder) Count(v int) {
	b.Params["count"] = v
}

// DatabaseGetUniversitiesBulder builder
//
// Returns a list of higher education institutions.
//
// https://vk.com/dev/database.getUniversities
type DatabaseGetUniversitiesBulder struct {
	api.Params
}

// NewDatabaseGetUniversitiesBulder func
func NewDatabaseGetUniversitiesBulder() *DatabaseGetUniversitiesBulder {
	return &DatabaseGetUniversitiesBulder{api.Params{}}
}

// Q Search query.
func (b *DatabaseGetUniversitiesBulder) Q(v string) {
	b.Params["q"] = v
}

// CountryID Country ID.
func (b *DatabaseGetUniversitiesBulder) CountryID(v int) {
	b.Params["country_id"] = v
}

// CityID City ID.
func (b *DatabaseGetUniversitiesBulder) CityID(v int) {
	b.Params["city_id"] = v
}

// Offset Offset needed to return a specific subset of universities.
func (b *DatabaseGetUniversitiesBulder) Offset(v int) {
	b.Params["offset"] = v
}

// Count Number of universities to return.
func (b *DatabaseGetUniversitiesBulder) Count(v int) {
	b.Params["count"] = v
}
