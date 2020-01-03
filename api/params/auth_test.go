package params_test

import (
	"testing"

	"github.com/SevereCloud/vksdk/api/params"
	"github.com/stretchr/testify/assert"
)

func TestAuthCheckPhoneBulder(t *testing.T) {
	b := params.NewAuthCheckPhoneBulder()

	b.Phone("text")
	b.ClientID(1)
	b.ClientSecret("text")
	b.AuthByPhone(true)

	assert.Equal(t, b.Params["phone"], "text")
	assert.Equal(t, b.Params["client_id"], 1)
	assert.Equal(t, b.Params["client_secret"], "text")
	assert.Equal(t, b.Params["auth_by_phone"], true)
}

func TestAuthRestoreBulder(t *testing.T) {
	b := params.NewAuthRestoreBulder()

	b.Phone("text")
	b.LastName("text")

	assert.Equal(t, b.Params["phone"], "text")
	assert.Equal(t, b.Params["last_name"], "text")
}
