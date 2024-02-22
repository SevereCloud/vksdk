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

	assert.Equal(t, 1, b.Params["owner_id"])
	assert.Equal(t, 1, b.Params["offset"])
	assert.Equal(t, 1, b.Params["count"])
	assert.Equal(t, []string{"text"}, b.Params["fields"])
}

func TestDonutGetSubscriptionBuilder(t *testing.T) {
	t.Parallel()

	b := params.NewDonutGetSubscriptionBuilder()

	b.OwnerID(1)

	assert.Equal(t, 1, b.Params["owner_id"])
}

func TestDonutGetSubscriptionsBuilder(t *testing.T) {
	t.Parallel()

	b := params.NewDonutGetSubscriptionsBuilder()

	b.Offset(1)
	b.Count(1)
	b.Fields([]string{"text"})

	assert.Equal(t, 1, b.Params["offset"])
	assert.Equal(t, 1, b.Params["count"])
	assert.Equal(t, []string{"text"}, b.Params["fields"])
}

func TestDonutIsDonBuilder(t *testing.T) {
	t.Parallel()

	b := params.NewDonutIsDonBuilder()

	b.OwnerID(1)

	assert.Equal(t, 1, b.Params["owner_id"])
}
