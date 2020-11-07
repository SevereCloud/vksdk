package params_test

import (
	"testing"

	"github.com/SevereCloud/vksdk/v2/api/params"
	"github.com/stretchr/testify/assert"
)

func TestDonutGetFriendsBuilder(t *testing.T) {
	t.Parallel()

	b := params.NewDonutGetFriendsBuilder()

	b.OwnerID(1)
	b.Offset(1)
	b.Count(1)
	b.Fields([]string{"text"})

	assert.Equal(t, b.Params["owner_id"], 1)
	assert.Equal(t, b.Params["offset"], 1)
	assert.Equal(t, b.Params["count"], 1)
	assert.Equal(t, b.Params["fields"], []string{"text"})
}

func TestDonutGetSubscriptionBuilder(t *testing.T) {
	t.Parallel()

	b := params.NewDonutGetSubscriptionBuilder()

	b.OwnerID(1)

	assert.Equal(t, b.Params["owner_id"], 1)
}

func TestDonutGetSubscriptionsBuilder(t *testing.T) {
	t.Parallel()

	b := params.NewDonutGetSubscriptionsBuilder()

	b.Offset(1)
	b.Count(1)
	b.Fields([]string{"text"})

	assert.Equal(t, b.Params["offset"], 1)
	assert.Equal(t, b.Params["count"], 1)
	assert.Equal(t, b.Params["fields"], []string{"text"})
}

func TestDonutIsDonBuilder(t *testing.T) {
	t.Parallel()

	b := params.NewDonutIsDonBuilder()

	b.OwnerID(1)

	assert.Equal(t, b.Params["owner_id"], 1)
}
