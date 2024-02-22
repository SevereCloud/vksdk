package params_test

import (
	"testing"

	"github.com/SevereCloud/vksdk/v2/api/params"
	"github.com/stretchr/testify/assert"
)

func TestLeadsCheckUserBuilder(t *testing.T) {
	t.Parallel()

	b := params.NewLeadsCheckUserBuilder()

	b.LeadID(1)
	b.TestResult(1)
	b.TestMode(true)
	b.AutoStart(true)
	b.Age(1)
	b.Country("text")

	assert.Equal(t, 1, b.Params["lead_id"])
	assert.Equal(t, 1, b.Params["test_result"])
	assert.Equal(t, true, b.Params["test_mode"])
	assert.Equal(t, true, b.Params["auto_start"])
	assert.Equal(t, 1, b.Params["age"])
	assert.Equal(t, "text", b.Params["country"])
}

func TestLeadsCompleteBuilder(t *testing.T) {
	t.Parallel()

	b := params.NewLeadsCompleteBuilder()

	b.VkSID("text")
	b.Secret("text")
	b.Comment("text")

	assert.Equal(t, "text", b.Params["vk_sid"])
	assert.Equal(t, "text", b.Params["secret"])
	assert.Equal(t, "text", b.Params["comment"])
}

func TestLeadsGetStatsBuilder(t *testing.T) {
	t.Parallel()

	b := params.NewLeadsGetStatsBuilder()

	b.LeadID(1)
	b.Secret("text")
	b.DateStart("text")
	b.DateEnd("text")

	assert.Equal(t, 1, b.Params["lead_id"])
	assert.Equal(t, "text", b.Params["secret"])
	assert.Equal(t, "text", b.Params["date_start"])
	assert.Equal(t, "text", b.Params["date_end"])
}

func TestLeadsGetUsersBuilder(t *testing.T) {
	t.Parallel()

	b := params.NewLeadsGetUsersBuilder()

	b.OfferID(1)
	b.Secret("text")
	b.Offset(1)
	b.Count(1)
	b.Status(1)
	b.Reverse(true)

	assert.Equal(t, 1, b.Params["offer_id"])
	assert.Equal(t, "text", b.Params["secret"])
	assert.Equal(t, 1, b.Params["offset"])
	assert.Equal(t, 1, b.Params["count"])
	assert.Equal(t, 1, b.Params["status"])
	assert.Equal(t, true, b.Params["reverse"])
}

func TestLeadsMetricHitBuilder(t *testing.T) {
	t.Parallel()

	b := params.NewLeadsMetricHitBuilder()

	b.Data("text")

	assert.Equal(t, "text", b.Params["data"])
}

func TestLeadsStartBuilder(t *testing.T) {
	t.Parallel()

	b := params.NewLeadsStartBuilder()

	b.LeadID(1)
	b.Secret("text")
	b.UID(1)
	b.AID(1)
	b.TestMode(true)
	b.Force(true)

	assert.Equal(t, 1, b.Params["lead_id"])
	assert.Equal(t, "text", b.Params["secret"])
	assert.Equal(t, 1, b.Params["uid"])
	assert.Equal(t, 1, b.Params["aid"])
	assert.Equal(t, true, b.Params["test_mode"])
	assert.Equal(t, true, b.Params["force"])
}
