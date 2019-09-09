package api

import (
	"testing"
)

func TestVK_DatabaseGetChairs(t *testing.T) {
	if vkService.AccessToken == "" {
		t.Skip("SERVICE_TOKEN empty")
	}

	_, gotVkErr := vkService.DatabaseGetChairs(map[string]string{
		"faculty_id": "15",
	})
	if gotVkErr.Code != 0 {
		t.Errorf("VK.DatabaseGetChairs() gotVkErr = %v", gotVkErr)
	}
}

func TestVK_DatabaseGetCities(t *testing.T) {
	if vkService.AccessToken == "" {
		t.Skip("SERVICE_TOKEN empty")
	}

	_, gotVkErr := vkService.DatabaseGetCities(map[string]string{
		"country_id": "1",
	})
	if gotVkErr.Code != 0 {
		t.Errorf("VK.DatabaseGetCities() gotVkErr = %v", gotVkErr)
	}
}

func TestVK_DatabaseGetCitiesByID(t *testing.T) {
	if vkService.AccessToken == "" {
		t.Skip("SERVICE_TOKEN empty")
	}

	_, gotVkErr := vkService.DatabaseGetCitiesByID(map[string]string{
		"city_ids": "1,5,192",
	})
	if gotVkErr.Code != 0 {
		t.Errorf("VK.DatabaseGetCitiesByID() gotVkErr = %v", gotVkErr)
	}
}

func TestVK_DatabaseGetCountries(t *testing.T) {
	if vkService.AccessToken == "" {
		t.Skip("SERVICE_TOKEN empty")
	}

	_, gotVkErr := vkService.DatabaseGetCountries(map[string]string{})
	if gotVkErr.Code != 0 {
		t.Errorf("VK.DatabaseGetCountries() gotVkErr = %v", gotVkErr)
	}
}

func TestVK_DatabaseGetCountriesByID(t *testing.T) {
	if vkService.AccessToken == "" {
		t.Skip("SERVICE_TOKEN empty")
	}

	_, gotVkErr := vkService.DatabaseGetCountriesByID(map[string]string{
		"country_ids": "1",
	})
	if gotVkErr.Code != 0 {
		t.Errorf("VK.DatabaseGetCountriesByID() gotVkErr = %v", gotVkErr)
	}
}

func TestVK_DatabaseGetFaculties(t *testing.T) {
	if vkService.AccessToken == "" {
		t.Skip("SERVICE_TOKEN empty")
	}

	_, gotVkErr := vkService.DatabaseGetFaculties(map[string]string{
		"university_id": "1",
	})
	if gotVkErr.Code != 0 {
		t.Errorf("VK.DatabaseGetFaculties() gotVkErr = %v", gotVkErr)
	}
}

func TestVK_DatabaseGetMetroStations(t *testing.T) {
	if vkService.AccessToken == "" {
		t.Skip("SERVICE_TOKEN empty")
	}

	_, gotVkErr := vkService.DatabaseGetMetroStations(map[string]string{
		"city_id":  "1",
		"extended": "1",
	})
	if gotVkErr.Code != 0 {
		t.Errorf("VK.DatabaseGetMetroStations() gotVkErr = %v", gotVkErr)
	}
}

func TestVK_DatabaseGetMetroStationsByID(t *testing.T) {
	if vkService.AccessToken == "" {
		t.Skip("SERVICE_TOKEN empty")
	}

	_, gotVkErr := vkService.DatabaseGetMetroStationsByID(map[string]string{
		"station_ids": "189, 181",
	})
	if gotVkErr.Code != 0 {
		t.Errorf("VK.DatabaseGetMetroStationsByID() gotVkErr = %v", gotVkErr)
	}
}

func TestVK_DatabaseGetRegions(t *testing.T) {
	if vkService.AccessToken == "" {
		t.Skip("SERVICE_TOKEN empty")
	}

	_, gotVkErr := vkService.DatabaseGetRegions(map[string]string{
		"country_id": "1",
	})
	if gotVkErr.Code != 0 {
		t.Errorf("VK.DatabaseGetRegions() gotVkErr = %v", gotVkErr)
	}
}

func TestVK_DatabaseGetSchoolClasses(t *testing.T) {
	if vkService.AccessToken == "" {
		t.Skip("SERVICE_TOKEN empty")
	}

	_, gotVkErr := vkService.DatabaseGetSchoolClasses(map[string]string{
		"country_id": "1",
	})
	if gotVkErr.Code != 0 {
		t.Errorf("VK.DatabaseGetSchoolClasses() gotVkErr = %v", gotVkErr)
	}
}

func TestVK_DatabaseGetSchools(t *testing.T) {
	if vkService.AccessToken == "" {
		t.Skip("SERVICE_TOKEN empty")
	}

	_, gotVkErr := vkService.DatabaseGetSchools(map[string]string{
		"q":       "56",
		"city_id": "2",
	})
	if gotVkErr.Code != 0 {
		t.Errorf("VK.DatabaseGetSchools() gotVkErr = %v", gotVkErr)
	}
}

func TestVK_DatabaseGetUniversities(t *testing.T) {
	if vkService.AccessToken == "" {
		t.Skip("SERVICE_TOKEN empty")
	}

	_, gotVkErr := vkService.DatabaseGetUniversities(map[string]string{
		"q": "СПб",
	})
	if gotVkErr.Code != 0 {
		t.Errorf("VK.DatabaseGetUniversities() gotVkErr = %v", gotVkErr)
	}
}
