package params // import "github.com/SevereCloud/vksdk/v2/api/params"

import (
	"github.com/SevereCloud/vksdk/v2/api"
)

// DatabaseGetChairsBuilder builder.
//
// Returns list of chairs on a specified faculty.
//
// https://vk.com/dev/database.getChairs
type DatabaseGetChairsBuilder struct {
	api.Params
}

// NewDatabaseGetChairsBuilder func.
func NewDatabaseGetChairsBuilder() *DatabaseGetChairsBuilder {
	return &DatabaseGetChairsBuilder{api.Params{}}
}

// FacultyID id of the faculty to get chairs from.
func (b *DatabaseGetChairsBuilder) FacultyID(v int) *DatabaseGetChairsBuilder {
	b.Params["faculty_id"] = v
	return b
}

// Offset offset required to get a certain subset of chairs.
func (b *DatabaseGetChairsBuilder) Offset(v int) *DatabaseGetChairsBuilder {
	b.Params["offset"] = v
	return b
}

// Count amount of chairs to get.
func (b *DatabaseGetChairsBuilder) Count(v int) *DatabaseGetChairsBuilder {
	b.Params["count"] = v
	return b
}

// DatabaseGetCitiesBuilder builder.
//
// Returns a list of cities.
//
// https://vk.com/dev/database.getCities
type DatabaseGetCitiesBuilder struct {
	api.Params
}

// NewDatabaseGetCitiesBuilder func.
func NewDatabaseGetCitiesBuilder() *DatabaseGetCitiesBuilder {
	return &DatabaseGetCitiesBuilder{api.Params{}}
}

// CountryID parameter.
func (b *DatabaseGetCitiesBuilder) CountryID(v int) *DatabaseGetCitiesBuilder {
	b.Params["country_id"] = v
	return b
}

// RegionID parameter.
func (b *DatabaseGetCitiesBuilder) RegionID(v int) *DatabaseGetCitiesBuilder {
	b.Params["region_id"] = v
	return b
}

// Q search query.
func (b *DatabaseGetCitiesBuilder) Q(v string) *DatabaseGetCitiesBuilder {
	b.Params["q"] = v
	return b
}

// NeedAll parameter.
//
// * 1 — to return all cities in the country,
//
// * 0 — to return major cities in the country (default).
func (b *DatabaseGetCitiesBuilder) NeedAll(v bool) *DatabaseGetCitiesBuilder {
	b.Params["need_all"] = v
	return b
}

// Offset needed to return a specific subset of cities.
func (b *DatabaseGetCitiesBuilder) Offset(v int) *DatabaseGetCitiesBuilder {
	b.Params["offset"] = v
	return b
}

// Count number of cities to return.
func (b *DatabaseGetCitiesBuilder) Count(v int) *DatabaseGetCitiesBuilder {
	b.Params["count"] = v
	return b
}

// DatabaseGetCitiesByIDBuilder builder.
//
// Returns information about cities by their IDs.
//
// https://vk.com/dev/database.getCitiesById
type DatabaseGetCitiesByIDBuilder struct {
	api.Params
}

// NewDatabaseGetCitiesByIDBuilder func.
func NewDatabaseGetCitiesByIDBuilder() *DatabaseGetCitiesByIDBuilder {
	return &DatabaseGetCitiesByIDBuilder{api.Params{}}
}

// CityIDs parameter.
func (b *DatabaseGetCitiesByIDBuilder) CityIDs(v []int) *DatabaseGetCitiesByIDBuilder {
	b.Params["city_ids"] = v
	return b
}

// DatabaseGetCountriesBuilder builder.
//
// Returns a list of countries.
//
// https://vk.com/dev/database.getCountries
type DatabaseGetCountriesBuilder struct {
	api.Params
}

// NewDatabaseGetCountriesBuilder func.
func NewDatabaseGetCountriesBuilder() *DatabaseGetCountriesBuilder {
	return &DatabaseGetCountriesBuilder{api.Params{}}
}

// NeedAll parameter.
//
// * 1 — to return a full list of all countries.
//
// * 0 — to return a list of countries near the current user's country (default).
func (b *DatabaseGetCountriesBuilder) NeedAll(v bool) *DatabaseGetCountriesBuilder {
	b.Params["need_all"] = v
	return b
}

// Code - country codes in [vk.com/dev/country_codes|ISO 3166-1 alpha-2] standard.
func (b *DatabaseGetCountriesBuilder) Code(v string) *DatabaseGetCountriesBuilder {
	b.Params["code"] = v
	return b
}

// Offset needed to return a specific subset of countries.
func (b *DatabaseGetCountriesBuilder) Offset(v int) *DatabaseGetCountriesBuilder {
	b.Params["offset"] = v
	return b
}

// Count number of countries to return.
func (b *DatabaseGetCountriesBuilder) Count(v int) *DatabaseGetCountriesBuilder {
	b.Params["count"] = v
	return b
}

// DatabaseGetCountriesByIDBuilder builder.
//
// Returns information about countries by their IDs.
//
// https://vk.com/dev/database.getCountriesById
type DatabaseGetCountriesByIDBuilder struct {
	api.Params
}

// NewDatabaseGetCountriesByIDBuilder func.
func NewDatabaseGetCountriesByIDBuilder() *DatabaseGetCountriesByIDBuilder {
	return &DatabaseGetCountriesByIDBuilder{api.Params{}}
}

// CountryIDs parameter.
func (b *DatabaseGetCountriesByIDBuilder) CountryIDs(v []int) *DatabaseGetCountriesByIDBuilder {
	b.Params["country_ids"] = v
	return b
}

// DatabaseGetFacultiesBuilder builder.
//
// Returns a list of faculties (i.e., university departments).
//
// https://vk.com/dev/database.getFaculties
type DatabaseGetFacultiesBuilder struct {
	api.Params
}

// NewDatabaseGetFacultiesBuilder func.
func NewDatabaseGetFacultiesBuilder() *DatabaseGetFacultiesBuilder {
	return &DatabaseGetFacultiesBuilder{api.Params{}}
}

// UniversityID parameter.
func (b *DatabaseGetFacultiesBuilder) UniversityID(v int) *DatabaseGetFacultiesBuilder {
	b.Params["university_id"] = v
	return b
}

// Offset needed to return a specific subset of faculties.
func (b *DatabaseGetFacultiesBuilder) Offset(v int) *DatabaseGetFacultiesBuilder {
	b.Params["offset"] = v
	return b
}

// Count number of faculties to return.
func (b *DatabaseGetFacultiesBuilder) Count(v int) *DatabaseGetFacultiesBuilder {
	b.Params["count"] = v
	return b
}

// DatabaseGetMetroStationsBuilder builder.
//
// Get metro stations by city.
//
// https://vk.com/dev/database.getMetroStations
type DatabaseGetMetroStationsBuilder struct {
	api.Params
}

// NewDatabaseGetMetroStationsBuilder func.
func NewDatabaseGetMetroStationsBuilder() *DatabaseGetMetroStationsBuilder {
	return &DatabaseGetMetroStationsBuilder{api.Params{}}
}

// CityID parameter.
func (b *DatabaseGetMetroStationsBuilder) CityID(v int) *DatabaseGetMetroStationsBuilder {
	b.Params["city_id"] = v
	return b
}

// Offset parameter.
func (b *DatabaseGetMetroStationsBuilder) Offset(v int) *DatabaseGetMetroStationsBuilder {
	b.Params["offset"] = v
	return b
}

// Count parameter.
func (b *DatabaseGetMetroStationsBuilder) Count(v int) *DatabaseGetMetroStationsBuilder {
	b.Params["count"] = v
	return b
}

// Extended parameter.
func (b *DatabaseGetMetroStationsBuilder) Extended(v bool) *DatabaseGetMetroStationsBuilder {
	b.Params["extended"] = v
	return b
}

// DatabaseGetMetroStationsByIDBuilder builder.
//
// Get metro station by his id.
//
// https://vk.com/dev/database.getMetroStationsById
type DatabaseGetMetroStationsByIDBuilder struct {
	api.Params
}

// NewDatabaseGetMetroStationsByIDBuilder func.
func NewDatabaseGetMetroStationsByIDBuilder() *DatabaseGetMetroStationsByIDBuilder {
	return &DatabaseGetMetroStationsByIDBuilder{api.Params{}}
}

// StationIDs parameter.
func (b *DatabaseGetMetroStationsByIDBuilder) StationIDs(v []int) *DatabaseGetMetroStationsByIDBuilder {
	b.Params["station_ids"] = v
	return b
}

// DatabaseGetRegionsBuilder builder.
//
// Returns a list of regions.
//
// https://vk.com/dev/database.getRegions
type DatabaseGetRegionsBuilder struct {
	api.Params
}

// NewDatabaseGetRegionsBuilder func.
func NewDatabaseGetRegionsBuilder() *DatabaseGetRegionsBuilder {
	return &DatabaseGetRegionsBuilder{api.Params{}}
}

// CountryID parameter, received in [vk.com/dev/database.getCountries|database.getCountries] method.
func (b *DatabaseGetRegionsBuilder) CountryID(v int) *DatabaseGetRegionsBuilder {
	b.Params["country_id"] = v
	return b
}

// Q search query.
func (b *DatabaseGetRegionsBuilder) Q(v string) *DatabaseGetRegionsBuilder {
	b.Params["q"] = v
	return b
}

// Offset needed to return specific subset of regions.
func (b *DatabaseGetRegionsBuilder) Offset(v int) *DatabaseGetRegionsBuilder {
	b.Params["offset"] = v
	return b
}

// Count number of regions to return.
func (b *DatabaseGetRegionsBuilder) Count(v int) *DatabaseGetRegionsBuilder {
	b.Params["count"] = v
	return b
}

// DatabaseGetSchoolClassesBuilder builder.
//
// Returns a list of school classes specified for the country.
//
// https://vk.com/dev/database.getSchoolClasses
type DatabaseGetSchoolClassesBuilder struct {
	api.Params
}

// NewDatabaseGetSchoolClassesBuilder func.
func NewDatabaseGetSchoolClassesBuilder() *DatabaseGetSchoolClassesBuilder {
	return &DatabaseGetSchoolClassesBuilder{api.Params{}}
}

// CountryID parameter.
func (b *DatabaseGetSchoolClassesBuilder) CountryID(v int) *DatabaseGetSchoolClassesBuilder {
	b.Params["country_id"] = v
	return b
}

// DatabaseGetSchoolsBuilder builder.
//
// Returns a list of schools.
//
// https://vk.com/dev/database.getSchools
type DatabaseGetSchoolsBuilder struct {
	api.Params
}

// NewDatabaseGetSchoolsBuilder func.
func NewDatabaseGetSchoolsBuilder() *DatabaseGetSchoolsBuilder {
	return &DatabaseGetSchoolsBuilder{api.Params{}}
}

// Q search query.
func (b *DatabaseGetSchoolsBuilder) Q(v string) *DatabaseGetSchoolsBuilder {
	b.Params["q"] = v
	return b
}

// CityID parameter.
func (b *DatabaseGetSchoolsBuilder) CityID(v int) *DatabaseGetSchoolsBuilder {
	b.Params["city_id"] = v
	return b
}

// Offset needed to return a specific subset of schools.
func (b *DatabaseGetSchoolsBuilder) Offset(v int) *DatabaseGetSchoolsBuilder {
	b.Params["offset"] = v
	return b
}

// Count number of schools to return.
func (b *DatabaseGetSchoolsBuilder) Count(v int) *DatabaseGetSchoolsBuilder {
	b.Params["count"] = v
	return b
}

// DatabaseGetUniversitiesBuilder builder.
//
// Returns a list of higher education institutions.
//
// https://vk.com/dev/database.getUniversities
type DatabaseGetUniversitiesBuilder struct {
	api.Params
}

// NewDatabaseGetUniversitiesBuilder func.
func NewDatabaseGetUniversitiesBuilder() *DatabaseGetUniversitiesBuilder {
	return &DatabaseGetUniversitiesBuilder{api.Params{}}
}

// Q search query.
func (b *DatabaseGetUniversitiesBuilder) Q(v string) *DatabaseGetUniversitiesBuilder {
	b.Params["q"] = v
	return b
}

// CountryID parameter.
func (b *DatabaseGetUniversitiesBuilder) CountryID(v int) *DatabaseGetUniversitiesBuilder {
	b.Params["country_id"] = v
	return b
}

// CityID parameter.
func (b *DatabaseGetUniversitiesBuilder) CityID(v int) *DatabaseGetUniversitiesBuilder {
	b.Params["city_id"] = v
	return b
}

// Offset needed to return a specific subset of universities.
func (b *DatabaseGetUniversitiesBuilder) Offset(v int) *DatabaseGetUniversitiesBuilder {
	b.Params["offset"] = v
	return b
}

// Count number of universities to return.
func (b *DatabaseGetUniversitiesBuilder) Count(v int) *DatabaseGetUniversitiesBuilder {
	b.Params["count"] = v
	return b
}
