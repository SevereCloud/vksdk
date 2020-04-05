package params_test

import (
	"testing"

	"github.com/SevereCloud/vksdk/api/params"
	"github.com/stretchr/testify/assert"
)

func TestAccountBanBuilder(t *testing.T) {
	t.Parallel()

	b := params.NewAccountBanBuilder()

	b.OwnerID(1)

	assert.Equal(t, b.Params["owner_id"], 1)
}

func TestAccountChangePasswordBuilder(t *testing.T) {
	t.Parallel()

	b := params.NewAccountChangePasswordBuilder()

	b.RestoreSID("text")
	b.ChangePasswordHash("text")
	b.OldPassword("text")
	b.NewPassword("text")

	assert.Equal(t, b.Params["restore_sid"], "text")
	assert.Equal(t, b.Params["change_password_hash"], "text")
	assert.Equal(t, b.Params["old_password"], "text")
	assert.Equal(t, b.Params["new_password"], "text")
}

func TestAccountGetActiveOffersBuilder(t *testing.T) {
	t.Parallel()

	b := params.NewAccountGetActiveOffersBuilder()

	b.Offset(1)
	b.Count(1)

	assert.Equal(t, b.Params["offset"], 1)
	assert.Equal(t, b.Params["count"], 1)
}

func TestAccountGetAppPermissionsBuilder(t *testing.T) {
	t.Parallel()

	b := params.NewAccountGetAppPermissionsBuilder()

	b.UserID(1)

	assert.Equal(t, b.Params["user_id"], 1)
}

func TestAccountGetBannedBuilder(t *testing.T) {
	t.Parallel()

	b := params.NewAccountGetBannedBuilder()

	b.Offset(1)
	b.Count(1)

	assert.Equal(t, b.Params["offset"], 1)
	assert.Equal(t, b.Params["count"], 1)
}

func TestAccountGetCountersBuilder(t *testing.T) {
	t.Parallel()

	b := params.NewAccountGetCountersBuilder()

	b.Filter([]string{"text"})

	assert.Equal(t, b.Params["filter"], []string{"text"})
}

func TestAccountGetInfoBuilder(t *testing.T) {
	t.Parallel()

	b := params.NewAccountGetInfoBuilder()

	b.Fields([]string{"text"})

	assert.Equal(t, b.Params["fields"], []string{"text"})
}

func TestAccountGetPushSettingsBuilder(t *testing.T) {
	t.Parallel()

	b := params.NewAccountGetPushSettingsBuilder()

	b.DeviceID("text")

	assert.Equal(t, b.Params["device_id"], "text")
}

func TestAccountRegisterDeviceBuilder(t *testing.T) {
	t.Parallel()

	b := params.NewAccountRegisterDeviceBuilder()

	b.Token("text")
	b.DeviceModel("text")
	b.DeviceYear(1)
	b.DeviceID("text")
	b.SystemVersion("text")
	b.Settings("text")
	b.Sandbox(true)

	assert.Equal(t, b.Params["token"], "text")
	assert.Equal(t, b.Params["device_model"], "text")
	assert.Equal(t, b.Params["device_year"], 1)
	assert.Equal(t, b.Params["device_id"], "text")
	assert.Equal(t, b.Params["system_version"], "text")
	assert.Equal(t, b.Params["settings"], "text")
	assert.Equal(t, b.Params["sandbox"], true)
}

func TestAccountSaveProfileInfoBuilder(t *testing.T) {
	t.Parallel()

	b := params.NewAccountSaveProfileInfoBuilder()

	b.FirstName("text")
	b.LastName("text")
	b.MaidenName("text")
	b.ScreenName("text")
	b.CancelRequestID(1)
	b.Sex(1)
	b.Relation(1)
	b.RelationPartnerID(1)
	b.Bdate("text")
	b.BdateVisibility(1)
	b.HomeTown("text")
	b.CountryID(1)
	b.CityID(1)
	b.Status("text")

	assert.Equal(t, b.Params["first_name"], "text")
	assert.Equal(t, b.Params["last_name"], "text")
	assert.Equal(t, b.Params["maiden_name"], "text")
	assert.Equal(t, b.Params["screen_name"], "text")
	assert.Equal(t, b.Params["cancel_request_id"], 1)
	assert.Equal(t, b.Params["sex"], 1)
	assert.Equal(t, b.Params["relation"], 1)
	assert.Equal(t, b.Params["relation_partner_id"], 1)
	assert.Equal(t, b.Params["bdate"], "text")
	assert.Equal(t, b.Params["bdate_visibility"], 1)
	assert.Equal(t, b.Params["home_town"], "text")
	assert.Equal(t, b.Params["country_id"], 1)
	assert.Equal(t, b.Params["city_id"], 1)
	assert.Equal(t, b.Params["status"], "text")
}

func TestAccountSetInfoBuilder(t *testing.T) {
	t.Parallel()

	b := params.NewAccountSetInfoBuilder()

	b.Name("text")
	b.Value("text")

	assert.Equal(t, b.Params["name"], "text")
	assert.Equal(t, b.Params["value"], "text")
}

func TestAccountSetNameInMenuBuilder(t *testing.T) {
	t.Parallel()

	b := params.NewAccountSetNameInMenuBuilder()

	b.UserID(1)
	b.Name("text")

	assert.Equal(t, b.Params["user_id"], 1)
	assert.Equal(t, b.Params["name"], "text")
}

func TestAccountSetOnlineBuilder(t *testing.T) {
	t.Parallel()

	b := params.NewAccountSetOnlineBuilder()

	b.Voip(true)

	assert.Equal(t, b.Params["voip"], true)
}

func TestAccountSetPushSettingsBuilder(t *testing.T) {
	t.Parallel()

	b := params.NewAccountSetPushSettingsBuilder()

	b.DeviceID("text")
	b.Settings("text")
	b.Key("text")
	b.Value([]string{"text"})

	assert.Equal(t, b.Params["device_id"], "text")
	assert.Equal(t, b.Params["settings"], "text")
	assert.Equal(t, b.Params["key"], "text")
	assert.Equal(t, b.Params["value"], []string{"text"})
}

func TestAccountSetSilenceModeBuilder(t *testing.T) {
	t.Parallel()

	b := params.NewAccountSetSilenceModeBuilder()

	b.DeviceID("text")
	b.Time(1)
	b.PeerID(1)
	b.Sound(1)

	assert.Equal(t, b.Params["device_id"], "text")
	assert.Equal(t, b.Params["time"], 1)
	assert.Equal(t, b.Params["peer_id"], 1)
	assert.Equal(t, b.Params["sound"], 1)
}

func TestAccountUnbanBuilder(t *testing.T) {
	t.Parallel()

	b := params.NewAccountUnbanBuilder()

	b.OwnerID(1)

	assert.Equal(t, b.Params["owner_id"], 1)
}

func TestAccountUnregisterDeviceBuilder(t *testing.T) {
	t.Parallel()

	b := params.NewAccountUnregisterDeviceBuilder()

	b.DeviceID("text")
	b.Sandbox(true)

	assert.Equal(t, b.Params["device_id"], "text")
	assert.Equal(t, b.Params["sandbox"], true)
}
