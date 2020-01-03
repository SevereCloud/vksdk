package params_test

import (
	"testing"

	"github.com/SevereCloud/vksdk/api/params"
	"github.com/stretchr/testify/assert"
)

func TestLeadsCheckUserBulder(t *testing.T) {
	b := params.NewLeadsCheckUserBulder()

	b.LeadID(1)
	b.TestResult(1)
	b.TestMode(true)
	b.AutoStart(true)
	b.Age(1)
	b.Country("text")

	assert.Equal(t, b.Params["lead_id"], 1)
	assert.Equal(t, b.Params["test_result"], 1)
	assert.Equal(t, b.Params["test_mode"], true)
	assert.Equal(t, b.Params["auto_start"], true)
	assert.Equal(t, b.Params["age"], 1)
	assert.Equal(t, b.Params["country"], "text")
}

func TestLeadsCompleteBulder(t *testing.T) {
	b := params.NewLeadsCompleteBulder()

	b.VkSID("text")
	b.Secret("text")
	b.Comment("text")

	assert.Equal(t, b.Params["vk_sid"], "text")
	assert.Equal(t, b.Params["secret"], "text")
	assert.Equal(t, b.Params["comment"], "text")
}

func TestLeadsGetStatsBulder(t *testing.T) {
	b := params.NewLeadsGetStatsBulder()

	b.LeadID(1)
	b.Secret("text")
	b.DateStart("text")
	b.DateEnd("text")

	assert.Equal(t, b.Params["lead_id"], 1)
	assert.Equal(t, b.Params["secret"], "text")
	assert.Equal(t, b.Params["date_start"], "text")
	assert.Equal(t, b.Params["date_end"], "text")
}

func TestLeadsGetUsersBulder(t *testing.T) {
	b := params.NewLeadsGetUsersBulder()

	b.OfferID(1)
	b.Secret("text")
	b.Offset(1)
	b.Count(1)
	b.Status(1)
	b.Reverse(true)

	assert.Equal(t, b.Params["offer_id"], 1)
	assert.Equal(t, b.Params["secret"], "text")
	assert.Equal(t, b.Params["offset"], 1)
	assert.Equal(t, b.Params["count"], 1)
	assert.Equal(t, b.Params["status"], 1)
	assert.Equal(t, b.Params["reverse"], true)
}

func TestLeadsMetricHitBulder(t *testing.T) {
	b := params.NewLeadsMetricHitBulder()

	b.Data("text")

	assert.Equal(t, b.Params["data"], "text")
}

func TestLeadsStartBulder(t *testing.T) {
	b := params.NewLeadsStartBulder()

	b.LeadID(1)
	b.Secret("text")
	b.UID(1)
	b.AID(1)
	b.TestMode(true)
	b.Force(true)

	assert.Equal(t, b.Params["lead_id"], 1)
	assert.Equal(t, b.Params["secret"], "text")
	assert.Equal(t, b.Params["uid"], 1)
	assert.Equal(t, b.Params["aid"], 1)
	assert.Equal(t, b.Params["test_mode"], true)
	assert.Equal(t, b.Params["force"], true)
}
