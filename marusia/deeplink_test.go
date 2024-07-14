package marusia_test

import (
	"testing"

	"github.com/SevereCloud/vksdk/v3/marusia"
	"github.com/stretchr/testify/assert"
)

func TestDeeplink(t *testing.T) {
	t.Parallel()

	f := func(marusiaID string, params map[string]string, actual string) {
		t.Helper()

		assert.Equal(t, marusia.CreateDeepLink(marusiaID, params), actual)
	}

	f("Marusya_ID", nil, "https://vk.me/marusia?event_name=external.Marusya_ID.start")
	f(
		"Marusya_ID",
		map[string]string{
			"key1": "value1",
			"key2": "value2",
		},
		"https://vk.me/marusia?event_data=eyJrZXkxIjoidmFsdWUxIiwia2V5MiI6InZhbHVlMiJ9&event_name=external.Marusya_ID.start",
	)
}
