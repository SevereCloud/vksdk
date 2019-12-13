package api

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestVK_DatabaseGetChairs(t *testing.T) {
	needServiceToken(t)

	res, err := vkService.DatabaseGetChairs(map[string]string{
		"faculty_id": "15",
	})
	assert.NoError(t, err)
	assert.NotEmpty(t, res.Count)

	if assert.NotEmpty(t, res.Items) {
		assert.NotEmpty(t, res.Items[0].ID)
		assert.NotEmpty(t, res.Items[0].Title)
	}
}

func TestVK_DatabaseGetCities(t *testing.T) {
	needServiceToken(t)

	res, err := vkService.DatabaseGetCities(map[string]string{
		"country_id": "1",
		"need_all":   "1",
	})
	assert.NoError(t, err)
	assert.NotEmpty(t, res.Count)

	if assert.NotEmpty(t, res.Items) {
		assert.NotEmpty(t, res.Items[0].ID)
		assert.NotEmpty(t, res.Items[0].Title)
		assert.NotEmpty(t, res.Items[0].Area)
		// assert.NotEmpty(t, res.Items[0].Important)
		assert.NotEmpty(t, res.Items[0].Region)
	}
}

func TestVK_DatabaseGetCitiesByID(t *testing.T) {
	needServiceToken(t)

	res, err := vkService.DatabaseGetCitiesByID(map[string]string{
		"city_ids": "1,5,192",
	})
	assert.NoError(t, err)

	if assert.NotEmpty(t, res) {
		assert.NotEmpty(t, res[0].ID)
		assert.NotEmpty(t, res[0].Title)
	}
}

func TestVK_DatabaseGetCountries(t *testing.T) {
	needServiceToken(t)

	res, err := vkService.DatabaseGetCountries(map[string]string{})
	assert.NoError(t, err)
	assert.NotEmpty(t, res.Count)

	if assert.NotEmpty(t, res.Items) {
		assert.NotEmpty(t, res.Items[0].ID)
		assert.NotEmpty(t, res.Items[0].Title)
	}
}

func TestVK_DatabaseGetCountriesByID(t *testing.T) {
	needServiceToken(t)

	res, err := vkService.DatabaseGetCountriesByID(map[string]string{
		"country_ids": "1",
	})
	assert.NoError(t, err)

	if assert.NotEmpty(t, res) {
		assert.NotEmpty(t, res[0].ID)
		assert.NotEmpty(t, res[0].Title)
	}
}

func TestVK_DatabaseGetFaculties(t *testing.T) {
	needServiceToken(t)

	res, err := vkService.DatabaseGetFaculties(map[string]string{
		"university_id": "1",
	})
	assert.NoError(t, err)
	assert.NotEmpty(t, res.Count)

	if assert.NotEmpty(t, res.Items) {
		assert.NotEmpty(t, res.Items[0].ID)
		assert.NotEmpty(t, res.Items[0].Title)
	}
}

func TestVK_DatabaseGetMetroStations(t *testing.T) {
	needServiceToken(t)

	res, err := vkService.DatabaseGetMetroStations(map[string]string{
		"city_id":  "1",
		"extended": "1",
	})
	assert.NoError(t, err)
	assert.NotEmpty(t, res.Count)

	if assert.NotEmpty(t, res.Items) {
		assert.NotEmpty(t, res.Items[0].ID)
		assert.NotEmpty(t, res.Items[0].Name)
		assert.NotEmpty(t, res.Items[0].Color)
	}
}

func TestVK_DatabaseGetMetroStationsByID(t *testing.T) {
	needServiceToken(t)

	res, err := vkService.DatabaseGetMetroStationsByID(map[string]string{
		"station_ids": "189, 181",
	})
	assert.NoError(t, err)

	if assert.NotEmpty(t, res) {
		assert.NotEmpty(t, res[0].ID)
		assert.NotEmpty(t, res[0].Name)
		assert.NotEmpty(t, res[0].Color)
	}
}

func TestVK_DatabaseGetRegions(t *testing.T) {
	needServiceToken(t)

	res, err := vkService.DatabaseGetRegions(map[string]string{
		"country_id": "1",
	})
	assert.NoError(t, err)
	assert.NotEmpty(t, res.Count)

	if assert.NotEmpty(t, res.Items) {
		assert.NotEmpty(t, res.Items[0].ID)
		assert.NotEmpty(t, res.Items[0].Title)
	}
}

func TestVK_DatabaseGetSchoolClasses(t *testing.T) {
	needServiceToken(t)

	res, err := vkService.DatabaseGetSchoolClasses(map[string]string{
		"country_id": "1",
	})
	assert.NoError(t, err)
	assert.NotEmpty(t, res)
}

func TestVK_DatabaseGetSchools(t *testing.T) {
	needServiceToken(t)

	res, err := vkService.DatabaseGetSchools(map[string]string{
		"q":       "56",
		"city_id": "2",
	})
	assert.NoError(t, err)
	assert.NotEmpty(t, res.Count)

	if assert.NotEmpty(t, res.Items) {
		assert.NotEmpty(t, res.Items[0].ID)
		assert.NotEmpty(t, res.Items[0].Title)
	}
}

func TestVK_DatabaseGetUniversities(t *testing.T) {
	needServiceToken(t)

	res, err := vkService.DatabaseGetUniversities(map[string]string{
		"q": "СПб",
	})
	assert.NoError(t, err)
	assert.NotEmpty(t, res.Count)

	if assert.NotEmpty(t, res.Items) {
		assert.NotEmpty(t, res.Items[0].ID)
		assert.NotEmpty(t, res.Items[0].Title)
	}
}
