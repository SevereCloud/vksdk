package params // import "github.com/SevereCloud/vksdk/api/params"

import (
	"github.com/SevereCloud/vksdk/api"
)

// AccountBanBulder builder
//
// https://vk.com/dev/account.ban
type AccountBanBulder struct {
	api.Params
}

// NewAccountBanBulder func
func NewAccountBanBulder() *AccountBanBulder {
	return &AccountBanBulder{api.Params{}}
}

// OwnerID parameter
func (b *AccountBanBulder) OwnerID(v int) {
	b.Params["owner_id"] = v
}

// AccountChangePasswordBulder builder
//
// Changes a user password after access is successfully restored with the [vk.com/dev/auth.restore|auth.restore] method.
//
// https://vk.com/dev/account.changePassword
type AccountChangePasswordBulder struct {
	api.Params
}

// NewAccountChangePasswordBulder func
func NewAccountChangePasswordBulder() *AccountChangePasswordBulder {
	return &AccountChangePasswordBulder{api.Params{}}
}

// RestoreSID Session id received after the [vk.com/dev/auth.restore|auth.restore] method is executed. (If the password is changed right after the access was restored)
func (b *AccountChangePasswordBulder) RestoreSID(v string) {
	b.Params["restore_sid"] = v
}

// ChangePasswordHash Hash received after a successful OAuth authorization with a code got by SMS. (If the password is changed right after the access was restored)
func (b *AccountChangePasswordBulder) ChangePasswordHash(v string) {
	b.Params["change_password_hash"] = v
}

// OldPassword Current user password.
func (b *AccountChangePasswordBulder) OldPassword(v string) {
	b.Params["old_password"] = v
}

// NewPassword New password that will be set as a current
func (b *AccountChangePasswordBulder) NewPassword(v string) {
	b.Params["new_password"] = v
}

// AccountGetActiveOffersBulder builder
//
// Returns a list of active ads (offers) which executed by the user will bring him/her respective number of votes to his balance in the application.
//
// https://vk.com/dev/account.getActiveOffers
type AccountGetActiveOffersBulder struct {
	api.Params
}

// NewAccountGetActiveOffersBulder func
func NewAccountGetActiveOffersBulder() *AccountGetActiveOffersBulder {
	return &AccountGetActiveOffersBulder{api.Params{}}
}

// Offset parameter
func (b *AccountGetActiveOffersBulder) Offset(v int) {
	b.Params["offset"] = v
}

// Count Number of results to return.
func (b *AccountGetActiveOffersBulder) Count(v int) {
	b.Params["count"] = v
}

// AccountGetAppPermissionsBulder builder
//
// Gets settings of the user in this application.
//
// https://vk.com/dev/account.getAppPermissions
type AccountGetAppPermissionsBulder struct {
	api.Params
}

// NewAccountGetAppPermissionsBulder func
func NewAccountGetAppPermissionsBulder() *AccountGetAppPermissionsBulder {
	return &AccountGetAppPermissionsBulder{api.Params{}}
}

// UserID User ID whose settings information shall be got. By default: current user.
func (b *AccountGetAppPermissionsBulder) UserID(v int) {
	b.Params["user_id"] = v
}

// AccountGetBannedBulder builder
//
// Returns a user's blacklist.
//
// https://vk.com/dev/account.getBanned
type AccountGetBannedBulder struct {
	api.Params
}

// NewAccountGetBannedBulder func
func NewAccountGetBannedBulder() *AccountGetBannedBulder {
	return &AccountGetBannedBulder{api.Params{}}
}

// Offset Offset needed to return a specific subset of results.
func (b *AccountGetBannedBulder) Offset(v int) {
	b.Params["offset"] = v
}

// Count Number of results to return.
func (b *AccountGetBannedBulder) Count(v int) {
	b.Params["count"] = v
}

// AccountGetCountersBulder builder
//
// Returns non-null values of user counters.
//
// https://vk.com/dev/account.getCounters
type AccountGetCountersBulder struct {
	api.Params
}

// NewAccountGetCountersBulder func
func NewAccountGetCountersBulder() *AccountGetCountersBulder {
	return &AccountGetCountersBulder{api.Params{}}
}

// Filter Counters to be returned.
func (b *AccountGetCountersBulder) Filter(v []string) {
	b.Params["filter"] = v
}

// AccountGetInfoBulder builder
//
// Returns current account info.
//
// https://vk.com/dev/account.getInfo
type AccountGetInfoBulder struct {
	api.Params
}

// NewAccountGetInfoBulder func
func NewAccountGetInfoBulder() *AccountGetInfoBulder {
	return &AccountGetInfoBulder{api.Params{}}
}

// Fields Fields to return. Possible values: *'country' — user country,, *'https_required' — is "HTTPS only" option enabled,, *'own_posts_default' — is "Show my posts only" option is enabled,, *'no_wall_replies' — are wall replies disabled or not,, *'intro' — is intro passed by user or not,, *'lang' — user language. By default: all.
func (b *AccountGetInfoBulder) Fields(v []string) {
	b.Params["fields"] = v
}

// AccountGetPushSettingsBulder builder
//
// Gets settings of push notifications.
//
// https://vk.com/dev/account.getPushSettings
type AccountGetPushSettingsBulder struct {
	api.Params
}

// NewAccountGetPushSettingsBulder func
func NewAccountGetPushSettingsBulder() *AccountGetPushSettingsBulder {
	return &AccountGetPushSettingsBulder{api.Params{}}
}

// DeviceID Unique device ID.
func (b *AccountGetPushSettingsBulder) DeviceID(v string) {
	b.Params["device_id"] = v
}

// AccountRegisterDeviceBulder builder
//
// Subscribes an iOS/Android/Windows Phone-based device to receive push notifications
//
// https://vk.com/dev/account.registerDevice
type AccountRegisterDeviceBulder struct {
	api.Params
}

// NewAccountRegisterDeviceBulder func
func NewAccountRegisterDeviceBulder() *AccountRegisterDeviceBulder {
	return &AccountRegisterDeviceBulder{api.Params{}}
}

// Token Device token used to send notifications. (for mpns, the token shall be URL for sending of notifications)
func (b *AccountRegisterDeviceBulder) Token(v string) {
	b.Params["token"] = v
}

// DeviceModel String name of device model.
func (b *AccountRegisterDeviceBulder) DeviceModel(v string) {
	b.Params["device_model"] = v
}

// DeviceYear Device year.
func (b *AccountRegisterDeviceBulder) DeviceYear(v int) {
	b.Params["device_year"] = v
}

// DeviceID Unique device ID.
func (b *AccountRegisterDeviceBulder) DeviceID(v string) {
	b.Params["device_id"] = v
}

// SystemVersion String version of device operating system.
func (b *AccountRegisterDeviceBulder) SystemVersion(v string) {
	b.Params["system_version"] = v
}

// Settings Push settings in a [vk.com/dev/push_settings|special format].
func (b *AccountRegisterDeviceBulder) Settings(v string) {
	b.Params["settings"] = v
}

// Sandbox parameter
func (b *AccountRegisterDeviceBulder) Sandbox(v bool) {
	b.Params["sandbox"] = v
}

// AccountSaveProfileInfoBulder builder
//
// Edits current profile info.
//
// https://vk.com/dev/account.saveProfileInfo
type AccountSaveProfileInfoBulder struct {
	api.Params
}

// NewAccountSaveProfileInfoBulder func
func NewAccountSaveProfileInfoBulder() *AccountSaveProfileInfoBulder {
	return &AccountSaveProfileInfoBulder{api.Params{}}
}

// FirstName User first name.
func (b *AccountSaveProfileInfoBulder) FirstName(v string) {
	b.Params["first_name"] = v
}

// LastName User last name.
func (b *AccountSaveProfileInfoBulder) LastName(v string) {
	b.Params["last_name"] = v
}

// MaidenName User maiden name (female only)
func (b *AccountSaveProfileInfoBulder) MaidenName(v string) {
	b.Params["maiden_name"] = v
}

// ScreenName User screen name.
func (b *AccountSaveProfileInfoBulder) ScreenName(v string) {
	b.Params["screen_name"] = v
}

// CancelRequestID ID of the name change request to be canceled. If this parameter is sent, all the others are ignored.
func (b *AccountSaveProfileInfoBulder) CancelRequestID(v int) {
	b.Params["cancel_request_id"] = v
}

// Sex User sex. Possible values: , * '1' – female,, * '2' – male.
func (b *AccountSaveProfileInfoBulder) Sex(v int) {
	b.Params["sex"] = v
}

// Relation User relationship status. Possible values: , * '1' – single,, * '2' – in a relationship,, * '3' – engaged,, * '4' – married,, * '5' – it's complicated,, * '6' – actively searching,, * '7' – in love,, * '0' – not specified.
func (b *AccountSaveProfileInfoBulder) Relation(v int) {
	b.Params["relation"] = v
}

// RelationPartnerID ID of the relationship partner.
func (b *AccountSaveProfileInfoBulder) RelationPartnerID(v int) {
	b.Params["relation_partner_id"] = v
}

// Bdate User birth date, format: DD.MM.YYYY.
func (b *AccountSaveProfileInfoBulder) Bdate(v string) {
	b.Params["bdate"] = v
}

// BdateVisibility Birth date visibility. Returned values: , * '1' – show birth date,, * '2' – show only month and day,, * '0' – hide birth date.
func (b *AccountSaveProfileInfoBulder) BdateVisibility(v int) {
	b.Params["bdate_visibility"] = v
}

// HomeTown User home town.
func (b *AccountSaveProfileInfoBulder) HomeTown(v string) {
	b.Params["home_town"] = v
}

// CountryID User country.
func (b *AccountSaveProfileInfoBulder) CountryID(v int) {
	b.Params["country_id"] = v
}

// CityID User city.
func (b *AccountSaveProfileInfoBulder) CityID(v int) {
	b.Params["city_id"] = v
}

// Status Status text.
func (b *AccountSaveProfileInfoBulder) Status(v string) {
	b.Params["status"] = v
}

// AccountSetInfoBulder builder
//
// Allows to edit the current account info.
//
// https://vk.com/dev/account.setInfo
type AccountSetInfoBulder struct {
	api.Params
}

// NewAccountSetInfoBulder func
func NewAccountSetInfoBulder() *AccountSetInfoBulder {
	return &AccountSetInfoBulder{api.Params{}}
}

// Name Setting name.
func (b *AccountSetInfoBulder) Name(v string) {
	b.Params["name"] = v
}

// Value Setting value.
func (b *AccountSetInfoBulder) Value(v string) {
	b.Params["value"] = v
}

// AccountSetNameInMenuBulder builder
//
// Sets an application screen name (up to 17 characters), that is shown to the user in the left menu.
//
// https://vk.com/dev/account.setNameInMenu
type AccountSetNameInMenuBulder struct {
	api.Params
}

// NewAccountSetNameInMenuBulder func
func NewAccountSetNameInMenuBulder() *AccountSetNameInMenuBulder {
	return &AccountSetNameInMenuBulder{api.Params{}}
}

// UserID User ID.
func (b *AccountSetNameInMenuBulder) UserID(v int) {
	b.Params["user_id"] = v
}

// Name Application screen name.
func (b *AccountSetNameInMenuBulder) Name(v string) {
	b.Params["name"] = v
}

// AccountSetOnlineBulder builder
//
// Marks the current user as online for 15 minutes.
//
// https://vk.com/dev/account.setOnline
type AccountSetOnlineBulder struct {
	api.Params
}

// NewAccountSetOnlineBulder func
func NewAccountSetOnlineBulder() *AccountSetOnlineBulder {
	return &AccountSetOnlineBulder{api.Params{}}
}

// Voip '1' if videocalls are available for current device.
func (b *AccountSetOnlineBulder) Voip(v bool) {
	b.Params["voip"] = v
}

// AccountSetPushSettingsBulder builder
//
// Change push settings.
//
// https://vk.com/dev/account.setPushSettings
type AccountSetPushSettingsBulder struct {
	api.Params
}

// NewAccountSetPushSettingsBulder func
func NewAccountSetPushSettingsBulder() *AccountSetPushSettingsBulder {
	return &AccountSetPushSettingsBulder{api.Params{}}
}

// DeviceID Unique device ID.
func (b *AccountSetPushSettingsBulder) DeviceID(v string) {
	b.Params["device_id"] = v
}

// Settings Push settings in a [vk.com/dev/push_settings|special format].
func (b *AccountSetPushSettingsBulder) Settings(v string) {
	b.Params["settings"] = v
}

// Key Notification key.
func (b *AccountSetPushSettingsBulder) Key(v string) {
	b.Params["key"] = v
}

// Value New value for the key in a [vk.com/dev/push_settings|special format].
func (b *AccountSetPushSettingsBulder) Value(v []string) {
	b.Params["value"] = v
}

// AccountSetSilenceModeBulder builder
//
// Mutes push notifications for the set period of time.
//
// https://vk.com/dev/account.setSilenceMode
type AccountSetSilenceModeBulder struct {
	api.Params
}

// NewAccountSetSilenceModeBulder func
func NewAccountSetSilenceModeBulder() *AccountSetSilenceModeBulder {
	return &AccountSetSilenceModeBulder{api.Params{}}
}

// DeviceID Unique device ID.
func (b *AccountSetSilenceModeBulder) DeviceID(v string) {
	b.Params["device_id"] = v
}

// Time Time in seconds for what notifications should be disabled. '-1' to disable forever.
func (b *AccountSetSilenceModeBulder) Time(v int) {
	b.Params["time"] = v
}

// PeerID Destination ID. "For user: 'User ID', e.g. '12345'. For chat: '2000000000' + 'Chat ID', e.g. '2000000001'. For community: '- Community ID', e.g. '-12345'. "
func (b *AccountSetSilenceModeBulder) PeerID(v int) {
	b.Params["peer_id"] = v
}

// Sound '1' — to enable sound in this dialog, '0' — to disable sound. Only if 'peer_id' contains user or community ID.
func (b *AccountSetSilenceModeBulder) Sound(v int) {
	b.Params["sound"] = v
}

// AccountUnbanBulder builder
//
// https://vk.com/dev/account.unban
type AccountUnbanBulder struct {
	api.Params
}

// NewAccountUnbanBulder func
func NewAccountUnbanBulder() *AccountUnbanBulder {
	return &AccountUnbanBulder{api.Params{}}
}

// OwnerID parameter
func (b *AccountUnbanBulder) OwnerID(v int) {
	b.Params["owner_id"] = v
}

// AccountUnregisterDeviceBulder builder
//
// Unsubscribes a device from push notifications.
//
// https://vk.com/dev/account.unregisterDevice
type AccountUnregisterDeviceBulder struct {
	api.Params
}

// NewAccountUnregisterDeviceBulder func
func NewAccountUnregisterDeviceBulder() *AccountUnregisterDeviceBulder {
	return &AccountUnregisterDeviceBulder{api.Params{}}
}

// DeviceID Unique device ID.
func (b *AccountUnregisterDeviceBulder) DeviceID(v string) {
	b.Params["device_id"] = v
}

// Sandbox parameter
func (b *AccountUnregisterDeviceBulder) Sandbox(v bool) {
	b.Params["sandbox"] = v
}
