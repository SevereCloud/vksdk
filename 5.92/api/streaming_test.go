package api

import (
	"os"
	"testing"
)

func TestVK_StreamingGetServerURL(t *testing.T) {
	serviceToken := os.Getenv("SERVICE_TOKEN")
	if serviceToken == "" {
		t.Skip("SERVICE_TOKEN empty")
	}
	vk := Init(serviceToken)

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
	serviceToken := os.Getenv("SERVICE_TOKEN")
	if serviceToken == "" {
		t.Skip("SERVICE_TOKEN empty")
	}
	vk := Init(serviceToken)

	response, vkErr := vk.StreamingGetSettings()
	if vkErr.Code != 0 {
		t.Errorf("%d %s", vkErr.Code, vkErr.Message)
	}

	if response.MonthlyLimit == "" {
		t.Error("Empty monthly_limit field")
	}
}

func TestVK_StreamingGetStats(t *testing.T) {
	serviceToken := os.Getenv("SERVICE_TOKEN")
	if serviceToken == "" {
		t.Skip("SERVICE_TOKEN empty")
	}
	vk := Init(serviceToken)

	params := make(map[string]string)
	params["type"] = "received"
	params["interval"] = "1h"
	_, vkErr := vk.StreamingGetStats(params)
	if vkErr.Code != 0 {
		t.Errorf("%d %s", vkErr.Code, vkErr.Message)
	}

	// if len(response) == 0 {
	// 	t.Error("Empty array")
	// }
}

func TestVK_StreamingGetStem(t *testing.T) {
	serviceToken := os.Getenv("SERVICE_TOKEN")
	if serviceToken == "" {
		t.Skip("SERVICE_TOKEN empty")
	}
	vk := Init(serviceToken)

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
	serviceToken := os.Getenv("SERVICE_TOKEN")
	if serviceToken == "" {
		t.Skip("SERVICE_TOKEN empty")
	}
	vk := Init(serviceToken)

	params := make(map[string]string)
	params["monthly_tier"] = "unlimited"
	vkErr := vk.StreamingSetSettings(params)
	if vkErr.Code != 0 {
		t.Errorf("%d %s", vkErr.Code, vkErr.Message)
	}
}

func TestVK_StreamingError(t *testing.T) {
	vk := Init("")

	_, vkErr := vk.StreamingGetServerURL()
	if vkErr.Code != 5 {
		t.Errorf("StreamingGetServerURL error bad %d %s", vkErr.Code, vkErr.Message)
	}

	_, vkErr = vk.StreamingGetSettings()
	if vkErr.Code != 5 {
		t.Errorf("StreamingGetServerURL error bad %d %s", vkErr.Code, vkErr.Message)
	}

	_, vkErr = vk.StreamingGetStats(map[string]string{})
	if vkErr.Code != 5 {
		t.Errorf("StreamingGetServerURL error bad %d %s", vkErr.Code, vkErr.Message)
	}

	_, vkErr = vk.StreamingGetStem(map[string]string{})
	if vkErr.Code != 5 {
		t.Errorf("StreamingGetServerURL error bad %d %s", vkErr.Code, vkErr.Message)
	}

	vkErr = vk.StreamingSetSettings(map[string]string{})
	if vkErr.Code != 5 {
		t.Errorf("StreamingGetServerURL error bad %d %s", vkErr.Code, vkErr.Message)
	}
}
