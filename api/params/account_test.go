package params_test

import (
	"testing"

	"github.com/SevereCloud/vksdk/v3/api/params"
	"github.com/stretchr/testify/assert"
)

func TestAccountBanBuilder(t *testing.T) {
	t.Parallel()

	b := params.NewAccountBanBuilder()

	b.OwnerID(1)

	assert.Equal(t, 1, b.Params["owner_id"])
}

func TestAccountChangePasswordBuilder(t *testing.T) {
	t.Parallel()

	b := params.NewAccountChangePasswordBuilder()

	b.RestoreSID("text")
	b.ChangePasswordHash("text")
	b.OldPassword("text")
	b.NewPassword("text")

	assert.Equal(t, "text", b.Params["restore_sid"])
	assert.Equal(t, "text", b.Params["change_password_hash"])
	assert.Equal(t, "text", b.Params["old_password"])
	assert.Equal(t, "text", b.Params["new_password"])
}

func TestAccountGetActiveOffersBuilder(t *testing.T) {
	t.Parallel()

	b := params.NewAccountGetActiveOffersBuilder()

	b.Offset(1)
	b.Count(1)

	assert.Equal(t, 1, b.Params["offset"])
	assert.Equal(t, 1, b.Params["count"])
}

func TestAccountGetAppPermissionsBuilder(t *testing.T) {
	t.Parallel()

	b := params.NewAccountGetAppPermissionsBuilder()

	b.UserID(1)

	assert.Equal(t, 1, b.Params["user_id"])
}

func TestAccountGetBannedBuilder(t *testing.T) {
	t.Parallel()

	b := params.NewAccountGetBannedBuilder()

	b.Offset(1)
	b.Count(1)

	assert.Equal(t, 1, b.Params["offset"])
	assert.Equal(t, 1, b.Params["count"])
}

func TestAccountGetCountersBuilder(t *testing.T) {
	t.Parallel()

	b := params.NewAccountGetCountersBuilder()

	b.Filter([]string{"text"})
	b.UserID(1)

	assert.Equal(t, []string{"text"}, b.Params["filter"])
	assert.Equal(t, 1, b.Params["user_id"])
}

func TestAccountGetInfoBuilder(t *testing.T) {
	t.Parallel()

	b := params.NewAccountGetInfoBuilder()

	b.Fields([]string{"text"})

	assert.Equal(t, []string{"text"}, b.Params["fields"])
}

func TestAccountGetPushSettingsBuilder(t *testing.T) {
	t.Parallel()

	b := params.NewAccountGetPushSettingsBuilder()

	b.DeviceID("text")

	assert.Equal(t, "text", b.Params["device_id"])
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

	assert.Equal(t, "text", b.Params["token"])
	assert.Equal(t, "text", b.Params["device_model"])
	assert.Equal(t, 1, b.Params["device_year"])
	assert.Equal(t, "text", b.Params["device_id"])
	assert.Equal(t, "text", b.Params["system_version"])
	assert.Equal(t, "text", b.Params["settings"])
	assert.Equal(t, true, b.Params["sandbox"])
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

	assert.Equal(t, "text", b.Params["first_name"])
	assert.Equal(t, "text", b.Params["last_name"])
	assert.Equal(t, "text", b.Params["maiden_name"])
	assert.Equal(t, "text", b.Params["screen_name"])
	assert.Equal(t, 1, b.Params["cancel_request_id"])
	assert.Equal(t, 1, b.Params["sex"])
	assert.Equal(t, 1, b.Params["relation"])
	assert.Equal(t, 1, b.Params["relation_partner_id"])
	assert.Equal(t, "text", b.Params["bdate"])
	assert.Equal(t, 1, b.Params["bdate_visibility"])
	assert.Equal(t, "text", b.Params["home_town"])
	assert.Equal(t, 1, b.Params["country_id"])
	assert.Equal(t, 1, b.Params["city_id"])
	assert.Equal(t, "text", b.Params["status"])
}

func TestAccountSetInfoBuilder(t *testing.T) {
	t.Parallel()

	b := params.NewAccountSetInfoBuilder()

	b.Name("text")
	b.Value("text")

	assert.Equal(t, "text", b.Params["name"])
	assert.Equal(t, "text", b.Params["value"])
}

func TestAccountSetNameInMenuBuilder(t *testing.T) {
	t.Parallel()

	b := params.NewAccountSetNameInMenuBuilder()

	b.UserID(1)
	b.Name("text")

	assert.Equal(t, 1, b.Params["user_id"])
	assert.Equal(t, "text", b.Params["name"])
}

func TestAccountSetOnlineBuilder(t *testing.T) {
	t.Parallel()

	b := params.NewAccountSetOnlineBuilder()

	b.Voip(true)

	assert.Equal(t, true, b.Params["voip"])
}

func TestAccountSetPushSettingsBuilder(t *testing.T) {
	t.Parallel()

	b := params.NewAccountSetPushSettingsBuilder()

	b.DeviceID("text")
	b.Settings("text")
	b.Key("text")
	b.Value([]string{"text"})

	assert.Equal(t, "text", b.Params["device_id"])
	assert.Equal(t, "text", b.Params["settings"])
	assert.Equal(t, "text", b.Params["key"])
	assert.Equal(t, []string{"text"}, b.Params["value"])
}

func TestAccountSetSilenceModeBuilder(t *testing.T) {
	t.Parallel()

	b := params.NewAccountSetSilenceModeBuilder()

	b.DeviceID("text")
	b.Time(1)
	b.PeerID(1)
	b.Sound(1)

	assert.Equal(t, "text", b.Params["device_id"])
	assert.Equal(t, 1, b.Params["time"])
	assert.Equal(t, 1, b.Params["peer_id"])
	assert.Equal(t, 1, b.Params["sound"])
}

func TestAccountUnbanBuilder(t *testing.T) {
	t.Parallel()

	b := params.NewAccountUnbanBuilder()

	b.OwnerID(1)

	assert.Equal(t, 1, b.Params["owner_id"])
}

func TestAccountUnregisterDeviceBuilder(t *testing.T) {
	t.Parallel()

	b := params.NewAccountUnregisterDeviceBuilder()

	b.DeviceID("text")
	b.Sandbox(true)

	assert.Equal(t, "text", b.Params["device_id"])
	assert.Equal(t, true, b.Params["sandbox"])
}
