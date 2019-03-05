package api

import (
	"os"
	"testing"
)

func TestVK_StreamingGetServerURL(t *testing.T) {
	vk := Init(os.Getenv("SERVICE_TOKEN"))

	response, vkErr := vk.StreamingGetServerURL()
	if vkErr.Code != 0 {
		t.Errorf("%d %s", vkErr.Code, vkErr.Message)
	}

	if response.Endpoint == "" {
		t.Error("Empty endpoint field")
	}
	if response.Key == "" {
		t.Error("Empty key field")
	}
}

func TestVK_StreamingGetSettings(t *testing.T) {
	vk := Init(os.Getenv("SERVICE_TOKEN"))

	response, vkErr := vk.StreamingGetSettings()
	if vkErr.Code != 0 {
		t.Errorf("%d %s", vkErr.Code, vkErr.Message)
	}

	if response.MonthlyLimit == "" {
		t.Error("Empty monthly_limit field")
	}
}

func TestVK_StreamingGetStats(t *testing.T) {
	vk := Init(os.Getenv("SERVICE_TOKEN"))

	params := make(map[string]string)
	params["type"] = "received"
	params["interval"] = "1h"
	response, vkErr := vk.StreamingGetStats(params)
	if vkErr.Code != 0 {
		t.Errorf("%d %s", vkErr.Code, vkErr.Message)
	}

	if len(response) == 0 {
		t.Error("Empty array")
	}
}

func TestVK_StreamingGetStem(t *testing.T) {
	vk := Init(os.Getenv("SERVICE_TOKEN"))

	params := make(map[string]string)
	params["word"] = "собака"
	response, vkErr := vk.StreamingGetStem(params)
	if vkErr.Code != 0 {
		t.Errorf("%d %s", vkErr.Code, vkErr.Message)
	}

	if response.Stem != "собак" {
		t.Error("Sterm wrong")
	}
}

func TestVK_StreamingSetSettings(t *testing.T) {
	// TODO: Add test cases.
}
