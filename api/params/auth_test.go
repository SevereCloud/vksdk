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

	assert.Equal(t, "text", b.Params["phone"])
	assert.Equal(t, 1, b.Params["client_id"])
	assert.Equal(t, "text", b.Params["client_secret"])
	assert.Equal(t, true, b.Params["auth_by_phone"])
}

func TestAuthRestoreBuilder(t *testing.T) {
	t.Parallel()

	b := params.NewAuthRestoreBuilder()

	b.Phone("text")
	b.LastName("text")

	assert.Equal(t, "text", b.Params["phone"])
	assert.Equal(t, "text", b.Params["last_name"])
}

func TestAuthGetProfileInfoBySilentTokenBuilder(t *testing.T) {
	t.Parallel()

	b := params.NewAuthGetProfileInfoBySilentTokenBuilder()

	b.Add("token", "uuid", "event")
	assert.Equal(t, "token", b.Params["token"])
	assert.Equal(t, "uuid", b.Params["uuid"])
	assert.Equal(t, "event", b.Params["event"])

	b.Add("token2", "uuid2", "event2")
	assert.Equal(t, "token,token2", b.Params["token"])
	assert.Equal(t, "uuid,uuid2", b.Params["uuid"])
	assert.Equal(t, "event,event2", b.Params["event"])
}

func TestAuthExchangeSilentAuthTokenBuilder(t *testing.T) {
	t.Parallel()

	b := params.NewAuthExchangeSilentAuthTokenBuilder()

	b.Token("token")
	b.UUID("uuid")

	assert.Equal(t, "token", b.Params["token"])
	assert.Equal(t, "uuid", b.Params["uuid"])
}
