package api_test

import (
	"testing"

	"github.com/SevereCloud/vksdk/api"

	"github.com/SevereCloud/vksdk/api/errors"
)

func TestVK_StreamingGetServerURL(t *testing.T) {
	needServiceToken(t)

	t.Run("StreamingGetServerURL not empty", func(t *testing.T) {
		response, err := vkService.StreamingGetServerURL(api.Params{})
		if err != nil {
			t.Errorf("%v", err)
		}

		if response.Endpoint == "" {
			t.Error("Empty endpoint field")
		}
		if response.Key == "" {
			t.Error("Empty key field")
		}
	})
}

func TestVK_StreamingGetSettings(t *testing.T) {
	needServiceToken(t)

	t.Run("StreamingGetSettings not empty", func(t *testing.T) {
		response, err := vkService.StreamingGetSettings(api.Params{})
		if err != nil {
			t.Errorf("%v", err)
		}

		if response.MonthlyLimit == "" {
			t.Error("Empty monthly_limit field")
		}
	})
}

func TestVK_StreamingGetStats(t *testing.T) {
	needServiceToken(t)

	params := api.Params{
		"type":     "received",
		"interval": "1h",
	}

	t.Run("StreamingGetStats not empty", func(t *testing.T) {
		// TODO: check it
		_, err := vkService.StreamingGetStats(params)
		if err != nil {
			t.Errorf("%v", err)
		}

		// if len(response) == 0 {
		// 	t.Error("Empty array")
		// }
	})
}

func TestVK_StreamingGetStem(t *testing.T) {
	needServiceToken(t)

	params := api.Params{
		"word": "собака",
	}

	t.Run("StreamingGetStem not empty", func(t *testing.T) {
		response, err := vkService.StreamingGetStem(params)
		if err != nil {
			t.Errorf("%v", err)
		}

		if response.Stem != "собак" {
			t.Error("Sterm wrong")
		}
	})
}

func TestVK_StreamingSetSettings(t *testing.T) {
	needServiceToken(t)

	params := api.Params{
		"monthly_tier": "unlimited",
	}

	t.Run("StreamingSetSettings not empty", func(t *testing.T) {
		// TODO: check StreamingSetSettings test
		_, err := vkService.StreamingSetSettings(params)
		if err != nil {
			t.Errorf("%v", err)
		}
	})
}

func TestVK_StreamingError(t *testing.T) {
	vk := api.Init("")

	t.Run("StreamingGetServerURL error", func(t *testing.T) {
		_, err := vk.StreamingGetServerURL(api.Params{})
		if errors.GetType(err) != errors.Auth {
			t.Errorf("StreamingGetServerURL error bad %v", err)
		}
	})

	t.Run("StreamingGetSettings error", func(t *testing.T) {
		_, err := vk.StreamingGetSettings(api.Params{})
		if errors.GetType(err) != errors.Auth {
			t.Errorf("StreamingGetSettings error bad %v", err)
		}
	})

	t.Run("StreamingGetStats error", func(t *testing.T) {
		_, err := vk.StreamingGetStats(api.Params{})
		if errors.GetType(err) != errors.Auth {
			t.Errorf("StreamingGetStats error bad %v", err)
		}
	})

	t.Run("StreamingGetStem error", func(t *testing.T) {
		_, err := vk.StreamingGetStem(api.Params{})
		if errors.GetType(err) != errors.Auth {
			t.Errorf("StreamingGetStem error bad %v", err)
		}
	})

	t.Run("StreamingSetSettings error", func(t *testing.T) {
		_, err := vk.StreamingSetSettings(api.Params{})
		if errors.GetType(err) != errors.Auth {
			t.Errorf("StreamingSetSettings error bad %v", err)
		}
	})
}
