package api_test

import (
	"testing"

	"github.com/SevereCloud/vksdk/api"
	"github.com/stretchr/testify/assert"
)

func TestVK_AccountBan(t *testing.T) {
	needUserToken(t)

	_, err := vkUser.AccountBan(api.Params{
		"owner_id": 1,
	})
	noError(t, err)

	_, err = vkUser.AccountBan(api.Params{
		"owner_id": -1,
	})
	noError(t, err)

	res, err := vkUser.AccountGetBanned(api.Params{})
	noError(t, err)
	assert.NotEmpty(t, res.Count)
	assert.NotEmpty(t, res.Items)
	assert.NotEmpty(t, res.Profiles)
	assert.NotEmpty(t, res.Groups)

	_, err = vkUser.AccountUnban(api.Params{
		"owner_id": 1,
	})
	noError(t, err)

	_, err = vkUser.AccountUnban(api.Params{
		"owner_id": -1,
	})
	noError(t, err)
}

func TestVK_AccountChangePassword(t *testing.T) {
	// TODO: Add test cases
}

func TestVK_AccountGetActiveOffers(t *testing.T) {
	needUserToken(t)

	_, err := vkUser.AccountGetActiveOffers(api.Params{})
	noError(t, err)
}

func TestVK_AccountGetAppPermissions(t *testing.T) {
	needUserToken(t)

	res, err := vkUser.AccountGetAppPermissions(api.Params{
		"user_id": vkUserID,
	})
	noError(t, err)
	assert.NotEmpty(t, res)
}

func TestVK_AccountGetCounters(t *testing.T) {
	needUserToken(t)

	_, err := vkUser.AccountGetCounters(api.Params{
		"user_id": vkUserID,
	})
	noError(t, err)
}

func TestVK_AccountGetInfo(t *testing.T) {
	needUserToken(t)

	_, err := vkUser.AccountGetInfo(api.Params{})
	noError(t, err)
}

func TestVK_AccountGetProfileInfo(t *testing.T) {
	needUserToken(t)

	info, err := vkUser.AccountGetProfileInfo(api.Params{})
	noError(t, err)
	assert.NotEmpty(t, info.FirstName)
	assert.NotEmpty(t, info.LastName)
	// assert.NotEmpty(t, info.Bdate)
	//assert.NotEmpty(t, info.BdateVisibility)
	//assert.NotEmpty(t, info.City)
	//assert.NotEmpty(t, info.Country)
	//assert.NotEmpty(t, info.HomeTown)
	//assert.NotEmpty(t, info.Relation)
	//assert.NotEmpty(t, info.Sex)
	//assert.NotEmpty(t, info.Status)
	//assert.NotEmpty(t, info.ScreenName)
	assert.NotEmpty(t, info.Phone)
}

func TestVK_AccountGetPushSettings(t *testing.T) {
	needUserToken(t)

	_, err := vkUser.AccountGetPushSettings(api.Params{
		"device_id": "aoa",
	})
	noError(t, err)
}

func TestVK_AccountRegisterDevice(t *testing.T) {
	// TODO: Add test cases
}

func TestVK_AccountSaveProfileInfo(t *testing.T) {
	needUserToken(t)

	_, err := vkUser.AccountSaveProfileInfo(api.Params{})
	noError(t, err)
}

func TestVK_AccountSetInfo(t *testing.T) {
	// TODO: Add test cases
}

func TestVK_AccountSetNameInMenu(t *testing.T) {
	// TODO: Add test cases
}

func TestVK_AccountSetOffline(t *testing.T) {
	needUserToken(t)

	_, err := vkUser.AccountSetOffline(api.Params{})
	noError(t, err)
}

func TestVK_AccountSetOnline(t *testing.T) {
	needUserToken(t)

	_, err := vkUser.AccountSetOnline(api.Params{})
	noError(t, err)
}

func TestVK_AccountSetPushSettings(t *testing.T) {
	// TODO: Add test cases
}

func TestVK_AccountSetSilenceMode(t *testing.T) {
	// TODO: Add test cases
}

func TestVK_AccountUnregisterDevice(t *testing.T) {
	// TODO: Add test cases
}
