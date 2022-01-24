package params_test

import (
	"testing"

	"github.com/SevereCloud/vksdk/v2/api/params"
	"github.com/stretchr/testify/assert"
)

func TestAuthCheckPhoneBuilder(t *testing.T) {
	t.Parallel()

	b := params.NewAuthCheckPhoneBuilder()

	b.Phone("text")
	b.ClientID(1)
	b.ClientSecret("text")
	b.AuthByPhone(true)

	assert.Equal(t, b.Params["phone"], "text")
	assert.Equal(t, b.Params["client_id"], 1)
	assert.Equal(t, b.Params["client_secret"], "text")
	assert.Equal(t, b.Params["auth_by_phone"], true)
}

func TestAuthRestoreBuilder(t *testing.T) {
	t.Parallel()

	b := params.NewAuthRestoreBuilder()

	b.Phone("text")
	b.LastName("text")

	assert.Equal(t, b.Params["phone"], "text")
	assert.Equal(t, b.Params["last_name"], "text")
}

func TestAuthGetProfileInfoBySilentTokenBuilder(t *testing.T) {
	t.Parallel()

	b := params.NewAuthGetProfileInfoBySilentTokenBuilder()

	b.Add("token", "uuid", "event")
	assert.Equal(t, b.Params["token"], "token")
	assert.Equal(t, b.Params["uuid"], "uuid")
	assert.Equal(t, b.Params["event"], "event")

	b.Add("token2", "uuid2", "event2")
	assert.Equal(t, b.Params["token"], "token,token2")
	assert.Equal(t, b.Params["uuid"], "uuid,uuid2")
	assert.Equal(t, b.Params["event"], "event,event2")
}

func TestAuthExchangeSilentAuthTokenBuilder(t *testing.T) {
	t.Parallel()

	b := params.NewAuthExchangeSilentAuthTokenBuilder()

	b.Token("token")
	b.UUID("uuid")

	assert.Equal(t, b.Params["token"], "token")
	assert.Equal(t, b.Params["uuid"], "uuid")
}
