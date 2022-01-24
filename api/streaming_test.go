package api_test

import (
	"errors"
	"testing"

	"github.com/SevereCloud/vksdk/v2/api"
	"github.com/stretchr/testify/assert"
)

func TestVK_StreamingGetServerURL(t *testing.T) {
	t.Parallel()

	needServiceToken(t)

	response, err := vkService.StreamingGetServerURL(nil)
	if err != nil {
		t.Errorf("%v", err)
	}

	if response.Endpoint == "" {
		t.Error("Empty endpoint field")
	}

	if response.Key == "" {
		t.Error("Empty key field")
	}
}

func TestVK_StreamingGetSettings(t *testing.T) {
	t.Parallel()

	needServiceToken(t)

	response, err := vkService.StreamingGetSettings(nil)
	if err != nil {
		t.Errorf("%v", err)
	}

	if response.MonthlyLimit == "" {
		t.Error("Empty monthly_limit field")
	}
}

func TestVK_StreamingGetStats(t *testing.T) {
	t.Parallel()

	needServiceToken(t)

	params := api.Params{
		"type":     "received",
		"interval": "1h",
	}

	// TODO: check it
	_, err := vkService.StreamingGetStats(params)
	if err != nil {
		t.Errorf("%v", err)
	}
}

func TestVK_StreamingGetStem(t *testing.T) {
	t.Parallel()

	needServiceToken(t)

	params := api.Params{
		"word": "собака",
	}

	response, err := vkService.StreamingGetStem(params)
	if err != nil {
		t.Errorf("%v", err)
	}

	assert.Equal(t, "собак", response.Stem)
}

func TestVK_StreamingSetSettings(t *testing.T) {
	t.Parallel()

	needServiceToken(t)

	params := api.Params{
		"monthly_tier": "unlimited",
	}

	// TODO: check StreamingSetSettings test
	_, err := vkService.StreamingSetSettings(params)
	if err != nil {
		t.Errorf("%v", err)
	}
}

func TestVK_StreamingError(t *testing.T) {
	t.Parallel()

	vk := api.NewVK("")

	_, err := vk.StreamingGetServerURL(nil)
	if !errors.Is(err, api.ErrAuth) {
		t.Errorf("StreamingGetServerURL error bad %v", err)
	}

	_, err = vk.StreamingGetSettings(nil)
	if !errors.Is(err, api.ErrAuth) {
		t.Errorf("StreamingGetSettings error bad %v", err)
	}

	_, err = vk.StreamingGetStats(nil)
	if !errors.Is(err, api.ErrAuth) {
		t.Errorf("StreamingGetStats error bad %v", err)
	}

	_, err = vk.StreamingGetStem(nil)
	if !errors.Is(err, api.ErrAuth) {
		t.Errorf("StreamingGetStem error bad %v", err)
	}

	_, err = vk.StreamingSetSettings(nil)
	if !errors.Is(err, api.ErrAuth) {
		t.Errorf("StreamingSetSettings error bad %v", err)
	}
}
