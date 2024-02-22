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

	assert.Equal(t, 1, b.Params["faculty_id"])
	assert.Equal(t, 1, b.Params["offset"])
	assert.Equal(t, 1, b.Params["count"])
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

	assert.Equal(t, 1, b.Params["country_id"])
	assert.Equal(t, 1, b.Params["region_id"])
	assert.Equal(t, "text", b.Params["q"])
	assert.Equal(t, true, b.Params["need_all"])
	assert.Equal(t, 1, b.Params["offset"])
	assert.Equal(t, 1, b.Params["count"])
}

func TestDatabaseGetCitiesByIDBuilder(t *testing.T) {
	t.Parallel()

	b := params.NewDatabaseGetCitiesByIDBuilder()

	b.CityIDs([]int{1})

	assert.Equal(t, []int{1}, b.Params["city_ids"])
}

func TestDatabaseGetCountriesBuilder(t *testing.T) {
	t.Parallel()

	b := params.NewDatabaseGetCountriesBuilder()

	b.NeedAll(true)
	b.Code("text")
	b.Offset(1)
	b.Count(1)

	assert.Equal(t, true, b.Params["need_all"])
	assert.Equal(t, "text", b.Params["code"])
	assert.Equal(t, 1, b.Params["offset"])
	assert.Equal(t, 1, b.Params["count"])
}

func TestDatabaseGetCountriesByIDBuilder(t *testing.T) {
	t.Parallel()

	b := params.NewDatabaseGetCountriesByIDBuilder()

	b.CountryIDs([]int{1})

	assert.Equal(t, []int{1}, b.Params["country_ids"])
}

func TestDatabaseGetFacultiesBuilder(t *testing.T) {
	t.Parallel()

	b := params.NewDatabaseGetFacultiesBuilder()

	b.UniversityID(1)
	b.Offset(1)
	b.Count(1)

	assert.Equal(t, 1, b.Params["university_id"])
	assert.Equal(t, 1, b.Params["offset"])
	assert.Equal(t, 1, b.Params["count"])
}

func TestDatabaseGetMetroStationsBuilder(t *testing.T) {
	t.Parallel()

	b := params.NewDatabaseGetMetroStationsBuilder()

	b.CityID(1)
	b.Offset(1)
	b.Count(1)
	b.Extended(true)

	assert.Equal(t, 1, b.Params["city_id"])
	assert.Equal(t, 1, b.Params["offset"])
	assert.Equal(t, 1, b.Params["count"])
	assert.Equal(t, true, b.Params["extended"])
}

func TestDatabaseGetMetroStationsByIDBuilder(t *testing.T) {
	t.Parallel()

	b := params.NewDatabaseGetMetroStationsByIDBuilder()

	b.StationIDs([]int{1})

	assert.Equal(t, []int{1}, b.Params["station_ids"])
}

func TestDatabaseGetRegionsBuilder(t *testing.T) {
	t.Parallel()

	b := params.NewDatabaseGetRegionsBuilder()

	b.CountryID(1)
	b.Q("text")
	b.Offset(1)
	b.Count(1)

	assert.Equal(t, 1, b.Params["country_id"])
	assert.Equal(t, "text", b.Params["q"])
	assert.Equal(t, 1, b.Params["offset"])
	assert.Equal(t, 1, b.Params["count"])
}

func TestDatabaseGetSchoolClassesBuilder(t *testing.T) {
	t.Parallel()

	b := params.NewDatabaseGetSchoolClassesBuilder()

	b.CountryID(1)

	assert.Equal(t, 1, b.Params["country_id"])
}

func TestDatabaseGetSchoolsBuilder(t *testing.T) {
	t.Parallel()

	b := params.NewDatabaseGetSchoolsBuilder()

	b.Q("text")
	b.CityID(1)
	b.Offset(1)
	b.Count(1)

	assert.Equal(t, "text", b.Params["q"])
	assert.Equal(t, 1, b.Params["city_id"])
	assert.Equal(t, 1, b.Params["offset"])
	assert.Equal(t, 1, b.Params["count"])
}

func TestDatabaseGetUniversitiesBuilder(t *testing.T) {
	t.Parallel()

	b := params.NewDatabaseGetUniversitiesBuilder()

	b.Q("text")
	b.CountryID(1)
	b.CityID(1)
	b.Offset(1)
	b.Count(1)

	assert.Equal(t, "text", b.Params["q"])
	assert.Equal(t, 1, b.Params["country_id"])
	assert.Equal(t, 1, b.Params["city_id"])
	assert.Equal(t, 1, b.Params["offset"])
	assert.Equal(t, 1, b.Params["count"])
}
