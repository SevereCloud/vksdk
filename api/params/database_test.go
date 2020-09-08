package params_test

import (
	"testing"

	"github.com/SevereCloud/vksdk/v2/api/params"
	"github.com/stretchr/testify/assert"
)

func TestDatabaseGetChairsBuilder(t *testing.T) {
	t.Parallel()

	b := params.NewDatabaseGetChairsBuilder()

	b.FacultyID(1)
	b.Offset(1)
	b.Count(1)

	assert.Equal(t, b.Params["faculty_id"], 1)
	assert.Equal(t, b.Params["offset"], 1)
	assert.Equal(t, b.Params["count"], 1)
}

func TestDatabaseGetCitiesBuilder(t *testing.T) {
	t.Parallel()

	b := params.NewDatabaseGetCitiesBuilder()

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

func TestDatabaseGetCitiesByIDBuilder(t *testing.T) {
	t.Parallel()

	b := params.NewDatabaseGetCitiesByIDBuilder()

	b.CityIDs([]int{1})

	assert.Equal(t, b.Params["city_ids"], []int{1})
}

func TestDatabaseGetCountriesBuilder(t *testing.T) {
	t.Parallel()

	b := params.NewDatabaseGetCountriesBuilder()

	b.NeedAll(true)
	b.Code("text")
	b.Offset(1)
	b.Count(1)

	assert.Equal(t, b.Params["need_all"], true)
	assert.Equal(t, b.Params["code"], "text")
	assert.Equal(t, b.Params["offset"], 1)
	assert.Equal(t, b.Params["count"], 1)
}

func TestDatabaseGetCountriesByIDBuilder(t *testing.T) {
	t.Parallel()

	b := params.NewDatabaseGetCountriesByIDBuilder()

	b.CountryIDs([]int{1})

	assert.Equal(t, b.Params["country_ids"], []int{1})
}

func TestDatabaseGetFacultiesBuilder(t *testing.T) {
	t.Parallel()

	b := params.NewDatabaseGetFacultiesBuilder()

	b.UniversityID(1)
	b.Offset(1)
	b.Count(1)

	assert.Equal(t, b.Params["university_id"], 1)
	assert.Equal(t, b.Params["offset"], 1)
	assert.Equal(t, b.Params["count"], 1)
}

func TestDatabaseGetMetroStationsBuilder(t *testing.T) {
	t.Parallel()

	b := params.NewDatabaseGetMetroStationsBuilder()

	b.CityID(1)
	b.Offset(1)
	b.Count(1)
	b.Extended(true)

	assert.Equal(t, b.Params["city_id"], 1)
	assert.Equal(t, b.Params["offset"], 1)
	assert.Equal(t, b.Params["count"], 1)
	assert.Equal(t, b.Params["extended"], true)
}

func TestDatabaseGetMetroStationsByIDBuilder(t *testing.T) {
	t.Parallel()

	b := params.NewDatabaseGetMetroStationsByIDBuilder()

	b.StationIDs([]int{1})

	assert.Equal(t, b.Params["station_ids"], []int{1})
}

func TestDatabaseGetRegionsBuilder(t *testing.T) {
	t.Parallel()

	b := params.NewDatabaseGetRegionsBuilder()

	b.CountryID(1)
	b.Q("text")
	b.Offset(1)
	b.Count(1)

	assert.Equal(t, b.Params["country_id"], 1)
	assert.Equal(t, b.Params["q"], "text")
	assert.Equal(t, b.Params["offset"], 1)
	assert.Equal(t, b.Params["count"], 1)
}

func TestDatabaseGetSchoolClassesBuilder(t *testing.T) {
	t.Parallel()

	b := params.NewDatabaseGetSchoolClassesBuilder()

	b.CountryID(1)

	assert.Equal(t, b.Params["country_id"], 1)
}

func TestDatabaseGetSchoolsBuilder(t *testing.T) {
	t.Parallel()

	b := params.NewDatabaseGetSchoolsBuilder()

	b.Q("text")
	b.CityID(1)
	b.Offset(1)
	b.Count(1)

	assert.Equal(t, b.Params["q"], "text")
	assert.Equal(t, b.Params["city_id"], 1)
	assert.Equal(t, b.Params["offset"], 1)
	assert.Equal(t, b.Params["count"], 1)
}

func TestDatabaseGetUniversitiesBuilder(t *testing.T) {
	t.Parallel()

	b := params.NewDatabaseGetUniversitiesBuilder()

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
