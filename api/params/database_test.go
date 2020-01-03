package params_test

import (
	"testing"

	"github.com/SevereCloud/vksdk/api/params"
	"github.com/stretchr/testify/assert"
)

func TestDatabaseGetChairsBulder(t *testing.T) {
	b := params.NewDatabaseGetChairsBulder()

	b.FacultyID(1)
	b.Offset(1)
	b.Count(1)

	assert.Equal(t, b.Params["faculty_id"], 1)
	assert.Equal(t, b.Params["offset"], 1)
	assert.Equal(t, b.Params["count"], 1)
}

func TestDatabaseGetCitiesBulder(t *testing.T) {
	b := params.NewDatabaseGetCitiesBulder()

	b.CountryID(1)
	b.RegionID(1)
	b.Q("text")
	b.NeedAll(true)
	b.Offset(1)
	b.Count(1)

	assert.Equal(t, b.Params["country_id"], 1)
	assert.Equal(t, b.Params["region_id"], 1)
	assert.Equal(t, b.Params["q"], "text")
	assert.Equal(t, b.Params["need_all"], true)
	assert.Equal(t, b.Params["offset"], 1)
	assert.Equal(t, b.Params["count"], 1)
}

func TestDatabaseGetCitiesByIDBulder(t *testing.T) {
	b := params.NewDatabaseGetCitiesByIDBulder()

	b.CityIDs([]int{1})

	assert.Equal(t, b.Params["city_ids"], []int{1})
}

func TestDatabaseGetCountriesBulder(t *testing.T) {
	b := params.NewDatabaseGetCountriesBulder()

	b.NeedAll(true)
	b.Code("text")
	b.Offset(1)
	b.Count(1)

	assert.Equal(t, b.Params["need_all"], true)
	assert.Equal(t, b.Params["code"], "text")
	assert.Equal(t, b.Params["offset"], 1)
	assert.Equal(t, b.Params["count"], 1)
}

func TestDatabaseGetCountriesByIDBulder(t *testing.T) {
	b := params.NewDatabaseGetCountriesByIDBulder()

	b.CountryIDs([]int{1})

	assert.Equal(t, b.Params["country_ids"], []int{1})
}

func TestDatabaseGetFacultiesBulder(t *testing.T) {
	b := params.NewDatabaseGetFacultiesBulder()

	b.UniversityID(1)
	b.Offset(1)
	b.Count(1)

	assert.Equal(t, b.Params["university_id"], 1)
	assert.Equal(t, b.Params["offset"], 1)
	assert.Equal(t, b.Params["count"], 1)
}

func TestDatabaseGetMetroStationsBulder(t *testing.T) {
	b := params.NewDatabaseGetMetroStationsBulder()

	b.CityID(1)
	b.Offset(1)
	b.Count(1)
	b.Extended(true)

	assert.Equal(t, b.Params["city_id"], 1)
	assert.Equal(t, b.Params["offset"], 1)
	assert.Equal(t, b.Params["count"], 1)
	assert.Equal(t, b.Params["extended"], true)
}

func TestDatabaseGetMetroStationsByIDBulder(t *testing.T) {
	b := params.NewDatabaseGetMetroStationsByIDBulder()

	b.StationIDs([]int{1})

	assert.Equal(t, b.Params["station_ids"], []int{1})
}

func TestDatabaseGetRegionsBulder(t *testing.T) {
	b := params.NewDatabaseGetRegionsBulder()

	b.CountryID(1)
	b.Q("text")
	b.Offset(1)
	b.Count(1)

	assert.Equal(t, b.Params["country_id"], 1)
	assert.Equal(t, b.Params["q"], "text")
	assert.Equal(t, b.Params["offset"], 1)
	assert.Equal(t, b.Params["count"], 1)
}

func TestDatabaseGetSchoolClassesBulder(t *testing.T) {
	b := params.NewDatabaseGetSchoolClassesBulder()

	b.CountryID(1)

	assert.Equal(t, b.Params["country_id"], 1)
}

func TestDatabaseGetSchoolsBulder(t *testing.T) {
	b := params.NewDatabaseGetSchoolsBulder()

	b.Q("text")
	b.CityID(1)
	b.Offset(1)
	b.Count(1)

	assert.Equal(t, b.Params["q"], "text")
	assert.Equal(t, b.Params["city_id"], 1)
	assert.Equal(t, b.Params["offset"], 1)
	assert.Equal(t, b.Params["count"], 1)
}

func TestDatabaseGetUniversitiesBulder(t *testing.T) {
	b := params.NewDatabaseGetUniversitiesBulder()

	b.Q("text")
	b.CountryID(1)
	b.CityID(1)
	b.Offset(1)
	b.Count(1)

	assert.Equal(t, b.Params["q"], "text")
	assert.Equal(t, b.Params["country_id"], 1)
	assert.Equal(t, b.Params["city_id"], 1)
	assert.Equal(t, b.Params["offset"], 1)
	assert.Equal(t, b.Params["count"], 1)
}
