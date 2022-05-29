package api_test

import (
	"testing"

	"github.com/Derad6709/vksdk/v2/api"
	"github.com/stretchr/testify/assert"
)

func TestVK_Execute_error(t *testing.T) {
	t.Parallel()

	needGroupToken(t)

	var response int

	err := vkGroup.Execute(`API.unknown.method();return 1;`, &response)
	assert.Error(t, err)
	assert.Equal(t, 1, response)
}

func TestVK_Execute_object(t *testing.T) {
	t.Parallel()

	needGroupToken(t)

	var response struct {
		Text string `json:"text"`
	}

	err := vkGroup.Execute(`return {text: "hello"};`, &response)
	assert.NoError(t, err)
	assert.Equal(t, "hello", response.Text)
}

func TestVK_ExecuteWithArgs_error(t *testing.T) {
	t.Parallel()

	needGroupToken(t)

	var response int

	err := vkGroup.ExecuteWithArgs(
		`API.unknown.method();return parseInt(Args.user_id);`,
		api.Params{"user_id": 1},
		&response,
	)
	assert.Error(t, err)
	assert.Equal(t, 1, response)
}

func TestVK_ExecuteWithArgs_object(t *testing.T) {
	t.Parallel()

	needGroupToken(t)

	var response struct {
		Text string `json:"text"`
	}

	err := vkGroup.ExecuteWithArgs(`return {text: Args.text};`, api.Params{"text": "hello"}, &response)
	assert.NoError(t, err)
	assert.Equal(t, "hello", response.Text)
}
