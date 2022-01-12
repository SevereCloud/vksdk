package api_test

import (
	"testing"

	"github.com/SevereCloud/vksdk/v2/api"

	"github.com/stretchr/testify/assert"
)

func TestVK_DatabaseGetChairs(t *testing.T) {
	t.Parallel()

	needServiceToken(t)

	res, err := vkService.DatabaseGetChairs(api.Params{
		"faculty_id": 15,
	})
	noError(t, err)
	assert.NotEmpty(t, res.Count)

	if assert.NotEmpty(t, res.Items) {
		assert.NotEmpty(t, res.Items[0].ID)
		assert.NotEmpty(t, res.Items[0].Title)
	}
}

func TestVK_DatabaseGetCities(t *testing.T) {
	t.Parallel()

	needServiceToken(t)

	res, err := vkService.DatabaseGetCities(api.Params{
		"country_id": 1,
		"need_all":   true,
	})
	noError(t, err)
	assert.NotEmpty(t, res.Count)

	if assert.NotEmpty(t, res.Items) {
		assert.NotEmpty(t, res.Items[0].ID)
		assert.NotEmpty(t, res.Items[0].Title)
	}
}

func TestVK_DatabaseGetCitiesByID(t *testing.T) {
	t.Parallel()

	needServiceToken(t)

	res, err := vkService.DatabaseGetCitiesByID(api.Params{
		"city_ids": "1,5,192",
	})
	noError(t, err)

	if assert.NotEmpty(t, res) {
		assert.NotEmpty(t, res[0].ID)
		assert.NotEmpty(t, res[0].Title)
	}
}

func TestVK_DatabaseGetCountries(t *testing.T) {
	t.Parallel()

	needServiceToken(t)

	res, err := vkService.DatabaseGetCountries(nil)
	noError(t, err)
	assert.NotEmpty(t, res.Count)

	if assert.NotEmpty(t, res.Items) {
		assert.NotEmpty(t, res.Items[0].ID)
		assert.NotEmpty(t, res.Items[0].Title)
	}
}

func TestVK_DatabaseGetCountriesByID(t *testing.T) {
	t.Parallel()

	needServiceToken(t)

	res, err := vkService.DatabaseGetCountriesByID(api.Params{
		"country_ids": 1,
	})
	noError(t, err)

	if assert.NotEmpty(t, res) {
		assert.NotEmpty(t, res[0].ID)
		assert.NotEmpty(t, res[0].Title)
	}
}

func TestVK_DatabaseGetFaculties(t *testing.T) {
	t.Parallel()

	needServiceToken(t)

	res, err := vkService.DatabaseGetFaculties(api.Params{
		"university_id": 1,
	})
	noError(t, err)
	assert.NotEmpty(t, res.Count)

	if assert.NotEmpty(t, res.Items) {
		assert.NotEmpty(t, res.Items[0].ID)
		assert.NotEmpty(t, res.Items[0].Title)
	}
}

func TestVK_DatabaseGetMetroStations(t *testing.T) {
	t.Parallel()

	needServiceToken(t)

	res, err := vkService.DatabaseGetMetroStations(api.Params{
		"city_id":  1,
		"extended": 1,
	})
	noError(t, err)
	assert.NotEmpty(t, res.Count)

	if assert.NotEmpty(t, res.Items) {
		assert.NotEmpty(t, res.Items[0].ID)
		assert.NotEmpty(t, res.Items[0].Name)
		assert.NotEmpty(t, res.Items[0].Color)
	}
}

func TestVK_DatabaseGetMetroStationsByID(t *testing.T) {
	t.Parallel()

	needServiceToken(t)

	res, err := vkService.DatabaseGetMetroStationsByID(api.Params{
		"station_ids": "189, 181",
	})
	noError(t, err)

	if assert.NotEmpty(t, res) {
		assert.NotEmpty(t, res[0].ID)
		assert.NotEmpty(t, res[0].Name)
		assert.NotEmpty(t, res[0].Color)
	}
}

func TestVK_DatabaseGetRegions(t *testing.T) {
	t.Parallel()

	needServiceToken(t)

	res, err := vkService.DatabaseGetRegions(api.Params{
		"country_id": 1,
	})
	noError(t, err)
	assert.NotEmpty(t, res.Count)

	if assert.NotEmpty(t, res.Items) {
		assert.NotEmpty(t, res.Items[0].ID)
		assert.NotEmpty(t, res.Items[0].Title)
	}
}

func TestVK_DatabaseGetSchoolClasses(t *testing.T) {
	t.Parallel()

	needServiceToken(t)

	res, err := vkService.DatabaseGetSchoolClasses(api.Params{
		"country_id": 1,
	})
	noError(t, err)
	assert.NotEmpty(t, res)
}

func TestVK_DatabaseGetSchools(t *testing.T) {
	t.Parallel()

	needServiceToken(t)

	res, err := vkService.DatabaseGetSchools(api.Params{
		"q":       56,
		"city_id": 2,
	})
	noError(t, err)
	assert.NotEmpty(t, res.Count)

	if assert.NotEmpty(t, res.Items) {
		assert.NotEmpty(t, res.Items[0].ID)
		assert.NotEmpty(t, res.Items[0].Title)
	}
}

func TestVK_DatabaseGetUniversities(t *testing.T) {
	t.Parallel()

	needServiceToken(t)

	res, err := vkService.DatabaseGetUniversities(api.Params{
		"q": "СПб",
	})
	noError(t, err)
	assert.NotEmpty(t, res.Count)

	if assert.NotEmpty(t, res.Items) {
		assert.NotEmpty(t, res.Items[0].ID)
		assert.NotEmpty(t, res.Items[0].Title)
	}
}
