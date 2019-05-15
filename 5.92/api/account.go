package api // import "github.com/SevereCloud/vksdk/5.92/api"

import (
	"github.com/SevereCloud/vksdk/5.92/object"
)

// AccountBan account.ban
// https://vk.com/dev/account.ban
func (vk *VK) AccountBan(params map[string]string) (vkErr Error) {
	_, vkErr = vk.Request("account.ban", params)
	return
}

// AccountChangePasswordResponse struct
type AccountChangePasswordResponse struct {
	Token string `json:"token"`
}

// AccountChangePassword changes a user password after access is successfully restored with the auth.restore method.
// https://vk.com/dev/account.changePassword
func (vk *VK) AccountChangePassword(params map[string]string) (response AccountChangePasswordResponse, vkErr Error) {
	vk.requestU("account.changePassword", params, &response, &vkErr)
	return
}

// AccountGetActiveOffersResponse struct
type AccountGetActiveOffersResponse struct {
	Count int                   `json:"count"`
	Items []object.AccountOffer `json:"items"`
}

// AccountGetActiveOffers returns a list of active ads (offers).If the user fulfill their conditions, he will be able to get the appropriate number of votes to his balance.
// https://vk.com/dev/account.getActiveOffers
func (vk *VK) AccountGetActiveOffers(params map[string]string) (response AccountGetActiveOffersResponse, vkErr Error) {
	vk.requestU("account.getActiveOffers", params, &response, &vkErr)
	return
}

// AccountGetAppPermissions gets settings of the user in this application.
// https://vk.com/dev/account.getAppPermissions
func (vk *VK) AccountGetAppPermissions(params map[string]string) (response int, vkErr Error) {
	vk.requestU("account.getAppPermissions", params, &response, &vkErr)
	return
}

// AccountGetBannedResponse struct
type AccountGetBannedResponse struct {
	Count    int                  `json:"count"`
	Items    []int                `json:"items"`
	Profiles []object.UsersUser   `json:"profiles"`
	Groups   []object.GroupsGroup `json:"groups"`
}

// AccountGetBanned returns a user's blacklist.
// https://vk.com/dev/account.getBanned
func (vk *VK) AccountGetBanned(params map[string]string) (response AccountGetBannedResponse, vkErr Error) {
	vk.requestU("account.getBanned", params, &response, &vkErr)
	return
}

// AccountGetCountersResponse struct
type AccountGetCountersResponse object.AccountAccountCounters

// AccountGetCounters returns non-null values of user counters.
// https://vk.com/dev/account.getCounters
func (vk *VK) AccountGetCounters(params map[string]string) (response AccountGetCountersResponse, vkErr Error) {
	vk.requestU("account.getCounters", params, &response, &vkErr)
	return
}

// AccountGetInfoResponse struct
type AccountGetInfoResponse object.AccountInfo

// AccountGetInfo returns current account info.
// https://vk.com/dev/account.getInfo
func (vk *VK) AccountGetInfo(params map[string]string) (response AccountGetInfoResponse, vkErr Error) {
	vk.requestU("account.getInfo", params, &response, &vkErr)
	return
}

// AccountGetProfileInfoResponse struct
type AccountGetProfileInfoResponse object.AccountUserSettings

// AccountGetProfileInfo returns the current account info.
// https://vk.com/dev/account.getProfileInfo
func (vk *VK) AccountGetProfileInfo() (response AccountGetProfileInfoResponse, vkErr Error) {
	vk.requestU("account.getProfileInfo", map[string]string{}, &response, &vkErr)
	return
}

// AccountGetPushSettingsResponse struct
type AccountGetPushSettingsResponse object.AccountPushSettings

// AccountGetPushSettings account.getPushSettings Gets settings of push notifications.
// https://vk.com/dev/account.getPushSettings
func (vk *VK) AccountGetPushSettings(params map[string]string) (response AccountGetPushSettingsResponse, vkErr Error) {
	vk.requestU("account.getPushSettings", params, &response, &vkErr)
	return
}

// AccountRegisterDevice subscribes an iOS/Android/Windows/Mac based device to receive push notifications
// https://vk.com/dev/account.registerDevice
func (vk *VK) AccountRegisterDevice(params map[string]string) (vkErr Error) {
	_, vkErr = vk.Request("account.registerDevice", params)
	return
}

// AccountSaveProfileInfoResponse struct
type AccountSaveProfileInfoResponse struct {
	Changed     int                       `json:"changed"`
	NameRequest object.AccountNameRequest `json:"name_request"`
}

// AccountSaveProfileInfo edits current profile info.
// https://vk.com/dev/account.saveProfileInfo
func (vk *VK) AccountSaveProfileInfo(params map[string]string) (response AccountSaveProfileInfoResponse, vkErr Error) {
	vk.requestU("account.saveProfileInfo", params, &response, &vkErr)
	return
}

// AccountSetInfo allows to edit the current account info.
// https://vk.com/dev/account.setInfo
func (vk *VK) AccountSetInfo(params map[string]string) (vkErr Error) {
	_, vkErr = vk.Request("account.setInfo", params)
	return
}

// AccountSetNameInMenu sets an application screen name (up to 17 characters), that is shown to the user in the left menu.
// https://vk.com/dev/account.setNameInMenu
func (vk *VK) AccountSetNameInMenu(params map[string]string) (vkErr Error) {
	_, vkErr = vk.Request("account.setNameInMenu", params)
	return
}

// AccountSetOffline marks a current user as offline.
// https://vk.com/dev/account.setOffline
func (vk *VK) AccountSetOffline() (vkErr Error) {
	_, vkErr = vk.Request("account.setOffline", make(map[string]string))
	return
}

// AccountSetOnline marks the current user as online for 5 minutes.
// https://vk.com/dev/account.setOnline
func (vk *VK) AccountSetOnline(params map[string]string) (vkErr Error) {
	_, vkErr = vk.Request("account.setOnline", params)
	return
}

// AccountSetPushSettings change push settings.
// https://vk.com/dev/account.setPushSettings
func (vk *VK) AccountSetPushSettings(params map[string]string) (vkErr Error) {
	_, vkErr = vk.Request("account.setPushSettings", params)
	return
}

// AccountSetSilenceMode mutes push notifications for the set period of time.
// https://vk.com/dev/account.setSilenceMode
func (vk *VK) AccountSetSilenceMode(params map[string]string) (vkErr Error) {
	_, vkErr = vk.Request("account.setSilenceMode", params)
	return
}

// AccountUnban account.unban
// https://vk.com/dev/account.unban
func (vk *VK) AccountUnban(params map[string]string) (vkErr Error) {
	_, vkErr = vk.Request("account.unban", params)
	return
}

// AccountUnregisterDevice unsubscribes a device from push notifications.
// https://vk.com/dev/account.unregisterDevice
func (vk *VK) AccountUnregisterDevice(params map[string]string) (vkErr Error) {
	_, vkErr = vk.Request("account.unregisterDevice", params)
	return
}
