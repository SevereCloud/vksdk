package events_test

import (
	"context"
	"testing"

	"github.com/SevereCloud/vksdk/v2/events"
	"github.com/SevereCloud/vksdk/v2/internal"
	"github.com/stretchr/testify/assert"
)

func TestGroupIDFromContext(t *testing.T) {
	t.Parallel()

	const groupID = 123
	ctx := context.WithValue(context.Background(), internal.GroupIDKey, groupID)
	assert.Equal(t, groupID, events.GroupIDFromContext(ctx))
}

func TestEventIDFromContext(t *testing.T) {
	t.Parallel()

	const eventID = "123"
	ctx := context.WithValue(context.Background(), internal.EventIDKey, eventID)
	assert.Equal(t, eventID, events.EventIDFromContext(ctx))
}

func TestVersionFromContext(t *testing.T) {
	t.Parallel()

	const version = "5.131"
	ctx := context.WithValue(context.Background(), internal.EventVersionKey, version)
	assert.Equal(t, version, events.VersionFromContext(ctx))
}
