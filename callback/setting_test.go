package callback_test

import (
	"context"
	"errors"
	"os"
	"testing"

	"github.com/Derad6709/vksdk/v2/api"
	"github.com/Derad6709/vksdk/v2/callback"
	"github.com/Derad6709/vksdk/v2/events"
	"github.com/stretchr/testify/assert"
)

func TestAutoSetting_ErrAuth(t *testing.T) {
	t.Parallel()

	vk := api.NewVK("")

	cb := callback.NewCallback()
	cb.MessageNew(func(_ context.Context, obj events.MessageNewObject) {})
	err := cb.AutoSetting(vk, "https://example.com")
	assert.Equal(t, true, errors.Is(err, api.ErrAuth), err)
}

func TestAutoSetting_ErrNeedGroupToken(t *testing.T) {
	t.Parallel()

	userToken := os.Getenv("USER_TOKEN")
	if userToken == "" {
		t.Skip("USER_TOKEN empty")
	}

	vk := api.NewVK(userToken)

	cb := callback.NewCallback()
	cb.MessageNew(func(_ context.Context, obj events.MessageNewObject) {})
	err := cb.AutoSetting(vk, "https://example.com")
	assert.Equal(t, true, errors.Is(err, callback.ErrNeedGroupToken), err)
}

func TestAutoSetting_Err(t *testing.T) {
	t.Parallel()

	groupToken := os.Getenv("GROUP_TOKEN")
	if groupToken == "" {
		t.Skip("GROUP_TOKEN empty")
	}

	vk := api.NewVK(groupToken)

	cb := callback.NewCallback()
	cb.MessageNew(func(_ context.Context, obj events.MessageNewObject) {})
	err := cb.AutoSetting(vk, "https://example.com")
	assert.NoError(t, err)
}
