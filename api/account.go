package api // import "github.com/SevereCloud/vksdk/v3/api"

import (
	"github.com/SevereCloud/vksdk/v3/object"
)

// AccountBan account.ban.
//
// https://dev.vk.com/method/account.ban
func (vk *VK) AccountBan(params Params) (response int, err error) {
	err = vk.RequestUnmarshal("account.ban", &response, params)
	return
}

// AccountChangePasswordResponse struct.
type AccountChangePasswordResponse struct {
	Token string `json:"token"`
}

// AccountChangePassword changes a user password after access is successfully restored with the auth.restore method.
//
// https://dev.vk.com/method/account.changePassword
func (vk *VK) AccountChangePassword(params Params) (response AccountChangePasswordResponse, err error) {
	err = vk.RequestUnmarshal("account.changePassword", &response, params)
	return
}

// AccountGetActiveOffersResponse struct.
type AccountGetActiveOffersResponse struct {
	Count int                   `json:"count"`
	Items []object.AccountOffer `json:"items"`
}

// AccountGetActiveOffers returns a list of active ads (offers).
// If the user fulfill their conditions, he will be able to get
// the appropriate number of votes to his balance.
//
// https://dev.vk.com/method/account.getActiveOffers
func (vk *VK) AccountGetActiveOffers(params Params) (response AccountGetActiveOffersResponse, err error) {
	err = vk.RequestUnmarshal("account.getActiveOffers", &response, params)
	return
}

// AccountGetAppPermissions gets settings of the user in this application.
//
// https://dev.vk.com/method/account.getAppPermissions
func (vk *VK) AccountGetAppPermissions(params Params) (response int, err error) {
	err = vk.RequestUnmarshal("account.getAppPermissions", &response, params)
	return
}

// AccountGetBannedResponse struct.
type AccountGetBannedResponse struct {
	Count int   `json:"count"`
	Items []int `json:"items"`
	object.ExtendedResponse
}

// AccountGetBanned returns a user's blacklist.
//
// https://dev.vk.com/method/account.getBanned
func (vk *VK) AccountGetBanned(params Params) (response AccountGetBannedResponse, err error) {
	err = vk.RequestUnmarshal("account.getBanned", &response, params)
	return
}

// AccountGetCountersResponse struct.
type AccountGetCountersResponse object.AccountAccountCounters

// AccountGetCounters returns non-null values of user counters.
//
// https://dev.vk.com/method/account.getCounters
func (vk *VK) AccountGetCounters(params Params) (response AccountGetCountersResponse, err error) {
	err = vk.RequestUnmarshal("account.getCounters", &response, params)
	return
}

// AccountGetInfoResponse struct.
type AccountGetInfoResponse object.AccountInfo

// AccountGetInfo returns current account info.
//
// https://dev.vk.com/method/account.getInfo
func (vk *VK) AccountGetInfo(params Params) (response AccountGetInfoResponse, err error) {
	err = vk.RequestUnmarshal("account.getInfo", &response, params)
	return
}

// AccountGetProfileInfoResponse struct.
type AccountGetProfileInfoResponse object.AccountUserSettings

// AccountGetProfileInfo returns the current account info.
//
// https://dev.vk.com/method/account.getProfileInfo
func (vk *VK) AccountGetProfileInfo(params Params) (response AccountGetProfileInfoResponse, err error) {
	err = vk.RequestUnmarshal("account.getProfileInfo", &response, params)
	return
}

// AccountGetPushSettingsResponse struct.
type AccountGetPushSettingsResponse object.AccountPushSettings

// AccountGetPushSettings account.getPushSettings Gets settings of push notifications.
//
// https://dev.vk.com/method/account.getPushSettings
func (vk *VK) AccountGetPushSettings(params Params) (response AccountGetPushSettingsResponse, err error) {
	err = vk.RequestUnmarshal("account.getPushSettings", &response, params)
	return
}

// AccountRegisterDevice subscribes an iOS/Android/Windows/Mac based device to receive push notifications.
//
// https://dev.vk.com/method/account.registerDevice
func (vk *VK) AccountRegisterDevice(params Params) (response int, err error) {
	err = vk.RequestUnmarshal("account.registerDevice", &response, params)
	return
}

// AccountSaveProfileInfoResponse struct.
type AccountSaveProfileInfoResponse struct {
	Changed     object.BaseBoolInt        `json:"changed"`
	NameRequest object.AccountNameRequest `json:"name_request"`
}

// AccountSaveProfileInfo edits current profile info.
//
// https://dev.vk.com/method/account.saveProfileInfo
func (vk *VK) AccountSaveProfileInfo(params Params) (response AccountSaveProfileInfoResponse, err error) {
	err = vk.RequestUnmarshal("account.saveProfileInfo", &response, params)
	return
}

// AccountSetInfo allows to edit the current account info.
//
// https://dev.vk.com/method/account.setInfo
func (vk *VK) AccountSetInfo(params Params) (response int, err error) {
	err = vk.RequestUnmarshal("account.setInfo", &response, params)
	return
}

// AccountSetNameInMenu sets an application screen name
// (up to 17 characters), that is shown to the user in the left menu.
//
// Deprecated: This method is deprecated and may be disabled soon, please avoid
//
// https://dev.vk.com/method/account.setNameInMenu
func (vk *VK) AccountSetNameInMenu(params Params) (response int, err error) {
	err = vk.RequestUnmarshal("account.setNameInMenu", &response, params)
	return
}

// AccountSetOffline marks a current user as offline.
//
// https://dev.vk.com/method/account.setOffline
func (vk *VK) AccountSetOffline(params Params) (response int, err error) {
	err = vk.RequestUnmarshal("account.setOffline", &response, params)
	return
}

// AccountSetOnline marks the current user as online for 5 minutes.
//
// https://dev.vk.com/method/account.setOnline
func (vk *VK) AccountSetOnline(params Params) (response int, err error) {
	err = vk.RequestUnmarshal("account.setOnline", &response, params)
	return
}

// AccountSetPushSettings change push settings.
//
// https://dev.vk.com/method/account.setPushSettings
func (vk *VK) AccountSetPushSettings(params Params) (response int, err error) {
	err = vk.RequestUnmarshal("account.setPushSettings", &response, params)
	return
}

// AccountSetSilenceMode mutes push notifications for the set period of time.
//
// https://dev.vk.com/method/account.setSilenceMode
func (vk *VK) AccountSetSilenceMode(params Params) (response int, err error) {
	err = vk.RequestUnmarshal("account.setSilenceMode", &response, params)
	return
}

// AccountUnban account.unban.
//
// https://dev.vk.com/method/account.unban
func (vk *VK) AccountUnban(params Params) (response int, err error) {
	err = vk.RequestUnmarshal("account.unban", &response, params)
	return
}

// AccountUnregisterDevice unsubscribes a device from push notifications.
//
// https://dev.vk.com/method/account.unregisterDevice
func (vk *VK) AccountUnregisterDevice(params Params) (response int, err error) {
	err = vk.RequestUnmarshal("account.unregisterDevice", &response, params)
	return
}
