package api

import (
	"testing"
)

func TestVK_DatabaseGetChairs(t *testing.T) {
	needServiceToken(t)

	_, err := vkService.DatabaseGetChairs(map[string]string{
		"faculty_id": "15",
	})
	if err != nil {
		t.Errorf("VK.DatabaseGetChairs() err = %v", err)
	}
}

func TestVK_DatabaseGetCities(t *testing.T) {
	needServiceToken(t)

	_, err := vkService.DatabaseGetCities(map[string]string{
		"country_id": "1",
	})
	if err != nil {
		t.Errorf("VK.DatabaseGetCities() err = %v", err)
	}
}

func TestVK_DatabaseGetCitiesByID(t *testing.T) {
	needServiceToken(t)

	_, err := vkService.DatabaseGetCitiesByID(map[string]string{
		"city_ids": "1,5,192",
	})
	if err != nil {
		t.Errorf("VK.DatabaseGetCitiesByID() err = %v", err)
	}
}

func TestVK_DatabaseGetCountries(t *testing.T) {
	needServiceToken(t)

	_, err := vkService.DatabaseGetCountries(map[string]string{})
	if err != nil {
		t.Errorf("VK.DatabaseGetCountries() err = %v", err)
	}
}

func TestVK_DatabaseGetCountriesByID(t *testing.T) {
	needServiceToken(t)

	_, err := vkService.DatabaseGetCountriesByID(map[string]string{
		"country_ids": "1",
	})
	if err != nil {
		t.Errorf("VK.DatabaseGetCountriesByID() err = %v", err)
	}
}

func TestVK_DatabaseGetFaculties(t *testing.T) {
	needServiceToken(t)

	_, err := vkService.DatabaseGetFaculties(map[string]string{
		"university_id": "1",
	})
	if err != nil {
		t.Errorf("VK.DatabaseGetFaculties() err = %v", err)
	}
}

func TestVK_DatabaseGetMetroStations(t *testing.T) {
	needServiceToken(t)

	_, err := vkService.DatabaseGetMetroStations(map[string]string{
		"city_id":  "1",
		"extended": "1",
	})
	if err != nil {
		t.Errorf("VK.DatabaseGetMetroStations() err = %v", err)
	}
}

func TestVK_DatabaseGetMetroStationsByID(t *testing.T) {
	needServiceToken(t)

	_, err := vkService.DatabaseGetMetroStationsByID(map[string]string{
		"station_ids": "189, 181",
	})
	if err != nil {
		t.Errorf("VK.DatabaseGetMetroStationsByID() err = %v", err)
	}
}

func TestVK_DatabaseGetRegions(t *testing.T) {
	needServiceToken(t)

	_, err := vkService.DatabaseGetRegions(map[string]string{
		"country_id": "1",
	})
	if err != nil {
		t.Errorf("VK.DatabaseGetRegions() err = %v", err)
	}
}

func TestVK_DatabaseGetSchoolClasses(t *testing.T) {
	needServiceToken(t)

	_, err := vkService.DatabaseGetSchoolClasses(map[string]string{
		"country_id": "1",
	})
	if err != nil {
		t.Errorf("VK.DatabaseGetSchoolClasses() err = %v", err)
	}
}

func TestVK_DatabaseGetSchools(t *testing.T) {
	needServiceToken(t)

	_, err := vkService.DatabaseGetSchools(map[string]string{
		"q":       "56",
		"city_id": "2",
	})
	if err != nil {
		t.Errorf("VK.DatabaseGetSchools() err = %v", err)
	}
}

func TestVK_DatabaseGetUniversities(t *testing.T) {
	needServiceToken(t)

	_, err := vkService.DatabaseGetUniversities(map[string]string{
		"q": "СПб",
	})
	if err != nil {
		t.Errorf("VK.DatabaseGetUniversities() err = %v", err)
	}
}
