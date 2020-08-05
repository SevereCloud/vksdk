package events_test

import (
	"context"
	"testing"

	"github.com/SevereCloud/vksdk/events"
	"github.com/SevereCloud/vksdk/internal"
	"github.com/stretchr/testify/assert"
)

func TestGroupIDFromContext(t *testing.T) {
	const groupID = 123
	ctx := context.WithValue(context.Background(), internal.GroupIDKey, groupID)
	assert.Equal(t, groupID, events.GroupIDFromContext(ctx))
}

func TestEventIDFromContext(t *testing.T) {
	const eventID = "123"
	ctx := context.WithValue(context.Background(), internal.EventIDKey, eventID)
	assert.Equal(t, eventID, events.EventIDFromContext(ctx))
}
