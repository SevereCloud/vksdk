package callback_test

import (
	"context"
	"os"
	"testing"

	"github.com/SevereCloud/vksdk/v2/api"
	"github.com/SevereCloud/vksdk/v2/callback"
	"github.com/SevereCloud/vksdk/v2/events"
	"github.com/stretchr/testify/assert"
)

func TestAutoSetting_ErrAuth(t *testing.T) {
	t.Parallel()

	vk := api.NewVK("")

	cb := callback.NewCallback()
	cb.MessageNew(func(_ context.Context, _ events.MessageNewObject) {})
	err := cb.AutoSetting(vk, "https://example.com")
	assert.ErrorIs(t, err, api.ErrAuth)
}

func TestAutoSetting_ErrNeedGroupToken(t *testing.T) {
	t.Parallel()

	userToken := os.Getenv("USER_TOKEN")
	if userToken == "" {
		t.Skip("USER_TOKEN empty")
	}

	vk := api.NewVK(userToken)

	cb := callback.NewCallback()
	cb.MessageNew(func(_ context.Context, _ events.MessageNewObject) {})
	err := cb.AutoSetting(vk, "https://example.com")
	assert.ErrorIs(t, err, callback.ErrNeedGroupToken)
}

func TestAutoSetting_Err(t *testing.T) {
	t.Parallel()

	groupToken := os.Getenv("GROUP_TOKEN")
	if groupToken == "" {
		t.Skip("GROUP_TOKEN empty")
	}

	vk := api.NewVK(groupToken)

	cb := callback.NewCallback()
	cb.MessageNew(func(_ context.Context, _ events.MessageNewObject) {})
	err := cb.AutoSetting(vk, "https://example.com")
	assert.NoError(t, err)
}
