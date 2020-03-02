package params // import "github.com/SevereCloud/vksdk/api/params"

import (
	"github.com/SevereCloud/vksdk/api"
)

// AccountBanBuilder builder
//
// https://vk.com/dev/account.ban
type AccountBanBuilder struct {
	api.Params
}

// NewAccountBanBuilder func
func NewAccountBanBuilder() *AccountBanBuilder {
	return &AccountBanBuilder{api.Params{}}
}

// OwnerID parameter
func (b *AccountBanBuilder) OwnerID(v int) *AccountBanBuilder {
	b.Params["owner_id"] = v
	return b
}

// AccountChangePasswordBuilder builder
//
// Changes a user password after access is successfully restored with the [vk.com/dev/auth.restore|auth.restore] method.
//
// https://vk.com/dev/account.changePassword
type AccountChangePasswordBuilder struct {
	api.Params
}

// NewAccountChangePasswordBuilder func
func NewAccountChangePasswordBuilder() *AccountChangePasswordBuilder {
	return &AccountChangePasswordBuilder{api.Params{}}
}

// RestoreSID Session id received after the [vk.com/dev/auth.restore|auth.restore] method is executed.
// (If the password is changed right after the access was restored)
func (b *AccountChangePasswordBuilder) RestoreSID(v string) *AccountChangePasswordBuilder {
	b.Params["restore_sid"] = v
	return b
}

// ChangePasswordHash Hash received after a successful OAuth authorization with a code got by SMS.
// (If the password is changed right after the access was restored)
func (b *AccountChangePasswordBuilder) ChangePasswordHash(v string) *AccountChangePasswordBuilder {
	b.Params["change_password_hash"] = v
	return b
}

// OldPassword Current user password.
func (b *AccountChangePasswordBuilder) OldPassword(v string) *AccountChangePasswordBuilder {
	b.Params["old_password"] = v
	return b
}

// NewPassword New password that will be set as a current
func (b *AccountChangePasswordBuilder) NewPassword(v string) *AccountChangePasswordBuilder {
	b.Params["new_password"] = v
	return b
}

// AccountGetActiveOffersBuilder builder
//
// Returns a list of active ads (offers) which executed by the user will bring him/her respective number of votes to his balance in the application.
//
// https://vk.com/dev/account.getActiveOffers
type AccountGetActiveOffersBuilder struct {
	api.Params
}

// NewAccountGetActiveOffersBuilder func
func NewAccountGetActiveOffersBuilder() *AccountGetActiveOffersBuilder {
	return &AccountGetActiveOffersBuilder{api.Params{}}
}

// Offset parameter
func (b *AccountGetActiveOffersBuilder) Offset(v int) *AccountGetActiveOffersBuilder {
	b.Params["offset"] = v
	return b
}

// Count Number of results to return.
func (b *AccountGetActiveOffersBuilder) Count(v int) *AccountGetActiveOffersBuilder {
	b.Params["count"] = v
	return b
}

// AccountGetAppPermissionsBuilder builder
//
// Gets settings of the user in this application.
//
// https://vk.com/dev/account.getAppPermissions
type AccountGetAppPermissionsBuilder struct {
	api.Params
}

// NewAccountGetAppPermissionsBuilder func
func NewAccountGetAppPermissionsBuilder() *AccountGetAppPermissionsBuilder {
	return &AccountGetAppPermissionsBuilder{api.Params{}}
}

// UserID User ID whose settings information shall be got.
//
// By default: current user.
func (b *AccountGetAppPermissionsBuilder) UserID(v int) *AccountGetAppPermissionsBuilder {
	b.Params["user_id"] = v
	return b
}

// AccountGetBannedBuilder builder
//
// Returns a user's blacklist.
//
// https://vk.com/dev/account.getBanned
type AccountGetBannedBuilder struct {
	api.Params
}

// NewAccountGetBannedBuilder func
func NewAccountGetBannedBuilder() *AccountGetBannedBuilder {
	return &AccountGetBannedBuilder{api.Params{}}
}

// Offset Offset needed to return a specific subset of results.
func (b *AccountGetBannedBuilder) Offset(v int) *AccountGetBannedBuilder {
	b.Params["offset"] = v
	return b
}

// Count Number of results to return.
func (b *AccountGetBannedBuilder) Count(v int) *AccountGetBannedBuilder {
	b.Params["count"] = v
	return b
}

// AccountGetCountersBuilder builder
//
// Returns non-null values of user counters.
//
// https://vk.com/dev/account.getCounters
type AccountGetCountersBuilder struct {
	api.Params
}

// NewAccountGetCountersBuilder func
func NewAccountGetCountersBuilder() *AccountGetCountersBuilder {
	return &AccountGetCountersBuilder{api.Params{}}
}

// Filter Counters to be returned.
func (b *AccountGetCountersBuilder) Filter(v []string) *AccountGetCountersBuilder {
	b.Params["filter"] = v
	return b
}

// AccountGetInfoBuilder builder
//
// Returns current account info.
//
// https://vk.com/dev/account.getInfo
type AccountGetInfoBuilder struct {
	api.Params
}

// NewAccountGetInfoBuilder func
func NewAccountGetInfoBuilder() *AccountGetInfoBuilder {
	return &AccountGetInfoBuilder{api.Params{}}
}

// Fields Fields to return. Possible values:
//
// * country — user country
//
// * https_required — is HTTPS only option enabled
//
// * own_posts_default — is Show my posts only option is enabled
//
// * no_wall_replies — are wall replies disabled or not
//
// * intro — is intro passed by user or not
//
// * lang — user language.
//
// By default: all.
func (b *AccountGetInfoBuilder) Fields(v []string) *AccountGetInfoBuilder {
	b.Params["fields"] = v
	return b
}

// AccountGetPushSettingsBuilder builder
//
// Gets settings of push notifications.
//
// https://vk.com/dev/account.getPushSettings
type AccountGetPushSettingsBuilder struct {
	api.Params
}

// NewAccountGetPushSettingsBuilder func
func NewAccountGetPushSettingsBuilder() *AccountGetPushSettingsBuilder {
	return &AccountGetPushSettingsBuilder{api.Params{}}
}

// DeviceID Unique device ID.
func (b *AccountGetPushSettingsBuilder) DeviceID(v string) *AccountGetPushSettingsBuilder {
	b.Params["device_id"] = v
	return b
}

// AccountRegisterDeviceBuilder builder
//
// Subscribes an iOS/Android/Windows Phone-based device to receive push notifications
//
// https://vk.com/dev/account.registerDevice
type AccountRegisterDeviceBuilder struct {
	api.Params
}

// NewAccountRegisterDeviceBuilder func
func NewAccountRegisterDeviceBuilder() *AccountRegisterDeviceBuilder {
	return &AccountRegisterDeviceBuilder{api.Params{}}
}

// Token Device token used to send notifications.
// (for mpns, the token shall be URL for sending of notifications)
func (b *AccountRegisterDeviceBuilder) Token(v string) *AccountRegisterDeviceBuilder {
	b.Params["token"] = v
	return b
}

// DeviceModel String name of device model.
func (b *AccountRegisterDeviceBuilder) DeviceModel(v string) *AccountRegisterDeviceBuilder {
	b.Params["device_model"] = v
	return b
}

// DeviceYear Device year.
func (b *AccountRegisterDeviceBuilder) DeviceYear(v int) *AccountRegisterDeviceBuilder {
	b.Params["device_year"] = v
	return b
}

// DeviceID Unique device ID.
func (b *AccountRegisterDeviceBuilder) DeviceID(v string) *AccountRegisterDeviceBuilder {
	b.Params["device_id"] = v
	return b
}

// SystemVersion String version of device operating system.
func (b *AccountRegisterDeviceBuilder) SystemVersion(v string) *AccountRegisterDeviceBuilder {
	b.Params["system_version"] = v
	return b
}

// Settings Push settings in a [vk.com/dev/push_settings|special format].
func (b *AccountRegisterDeviceBuilder) Settings(v string) *AccountRegisterDeviceBuilder {
	b.Params["settings"] = v
	return b
}

// Sandbox parameter
func (b *AccountRegisterDeviceBuilder) Sandbox(v bool) *AccountRegisterDeviceBuilder {
	b.Params["sandbox"] = v
	return b
}

// AccountSaveProfileInfoBuilder builder
//
// Edits current profile info.
//
// https://vk.com/dev/account.saveProfileInfo
type AccountSaveProfileInfoBuilder struct {
	api.Params
}

// NewAccountSaveProfileInfoBuilder func
func NewAccountSaveProfileInfoBuilder() *AccountSaveProfileInfoBuilder {
	return &AccountSaveProfileInfoBuilder{api.Params{}}
}

// FirstName User first name.
func (b *AccountSaveProfileInfoBuilder) FirstName(v string) *AccountSaveProfileInfoBuilder {
	b.Params["first_name"] = v
	return b
}

// LastName User last name.
func (b *AccountSaveProfileInfoBuilder) LastName(v string) *AccountSaveProfileInfoBuilder {
	b.Params["last_name"] = v
	return b
}

// MaidenName User maiden name (female only)
func (b *AccountSaveProfileInfoBuilder) MaidenName(v string) *AccountSaveProfileInfoBuilder {
	b.Params["maiden_name"] = v
	return b
}

// ScreenName User screen name.
func (b *AccountSaveProfileInfoBuilder) ScreenName(v string) *AccountSaveProfileInfoBuilder {
	b.Params["screen_name"] = v
	return b
}

// CancelRequestID ID of the name change request to be canceled.
// If this parameter is sent, all the others are ignored.
func (b *AccountSaveProfileInfoBuilder) CancelRequestID(v int) *AccountSaveProfileInfoBuilder {
	b.Params["cancel_request_id"] = v
	return b
}

// Sex User sex. Possible values:
//
// * 1 – female
//
// * 2 – male.
func (b *AccountSaveProfileInfoBuilder) Sex(v int) *AccountSaveProfileInfoBuilder {
	b.Params["sex"] = v
	return b
}

// Relation User relationship status. Possible values:
//
// * 1 – single
//
// * 2 – in a relationship
//
// * 3 – engaged
//
// * 4 – married
//
// * 5 – its complicated
//
// * 6 – actively searching
//
// * 7 – in love
//
// * 0 – not specified.
func (b *AccountSaveProfileInfoBuilder) Relation(v int) *AccountSaveProfileInfoBuilder {
	b.Params["relation"] = v
	return b
}

// RelationPartnerID ID of the relationship partner.
func (b *AccountSaveProfileInfoBuilder) RelationPartnerID(v int) *AccountSaveProfileInfoBuilder {
	b.Params["relation_partner_id"] = v
	return b
}

// Bdate User birth date, format: DD.MM.YYYY.
func (b *AccountSaveProfileInfoBuilder) Bdate(v string) *AccountSaveProfileInfoBuilder {
	b.Params["bdate"] = v
	return b
}

// BdateVisibility Birth date visibility. Returned values:
//
// * 1 – show birth date
//
// * 2 – show only month and day
//
// * 0 – hide birth date.
func (b *AccountSaveProfileInfoBuilder) BdateVisibility(v int) *AccountSaveProfileInfoBuilder {
	b.Params["bdate_visibility"] = v
	return b
}

// HomeTown User home town.
func (b *AccountSaveProfileInfoBuilder) HomeTown(v string) *AccountSaveProfileInfoBuilder {
	b.Params["home_town"] = v
	return b
}

// CountryID User country.
func (b *AccountSaveProfileInfoBuilder) CountryID(v int) *AccountSaveProfileInfoBuilder {
	b.Params["country_id"] = v
	return b
}

// CityID User city.
func (b *AccountSaveProfileInfoBuilder) CityID(v int) *AccountSaveProfileInfoBuilder {
	b.Params["city_id"] = v
	return b
}

// Status Status text.
func (b *AccountSaveProfileInfoBuilder) Status(v string) *AccountSaveProfileInfoBuilder {
	b.Params["status"] = v
	return b
}

// AccountSetInfoBuilder builder
//
// Allows to edit the current account info.
//
// https://vk.com/dev/account.setInfo
type AccountSetInfoBuilder struct {
	api.Params
}

// NewAccountSetInfoBuilder func
func NewAccountSetInfoBuilder() *AccountSetInfoBuilder {
	return &AccountSetInfoBuilder{api.Params{}}
}

// Name Setting name.
func (b *AccountSetInfoBuilder) Name(v string) *AccountSetInfoBuilder {
	b.Params["name"] = v
	return b
}

// Value Setting value.
func (b *AccountSetInfoBuilder) Value(v string) *AccountSetInfoBuilder {
	b.Params["value"] = v
	return b
}

// AccountSetNameInMenuBuilder builder
//
// Sets an application screen name (up to 17 characters), that is shown to the user in the left menu.
//
// https://vk.com/dev/account.setNameInMenu
type AccountSetNameInMenuBuilder struct {
	api.Params
}

// NewAccountSetNameInMenuBuilder func
func NewAccountSetNameInMenuBuilder() *AccountSetNameInMenuBuilder {
	return &AccountSetNameInMenuBuilder{api.Params{}}
}

// UserID User ID.
func (b *AccountSetNameInMenuBuilder) UserID(v int) *AccountSetNameInMenuBuilder {
	b.Params["user_id"] = v
	return b
}

// Name Application screen name.
func (b *AccountSetNameInMenuBuilder) Name(v string) *AccountSetNameInMenuBuilder {
	b.Params["name"] = v
	return b
}

// AccountSetOnlineBuilder builder
//
// Marks the current user as online for 15 minutes.
//
// https://vk.com/dev/account.setOnline
type AccountSetOnlineBuilder struct {
	api.Params
}

// NewAccountSetOnlineBuilder func
func NewAccountSetOnlineBuilder() *AccountSetOnlineBuilder {
	return &AccountSetOnlineBuilder{api.Params{}}
}

// Voip parameter '1' if videocalls are available for current device.
func (b *AccountSetOnlineBuilder) Voip(v bool) *AccountSetOnlineBuilder {
	b.Params["voip"] = v
	return b
}

// AccountSetPushSettingsBuilder builder
//
// Change push settings.
//
// https://vk.com/dev/account.setPushSettings
type AccountSetPushSettingsBuilder struct {
	api.Params
}

// NewAccountSetPushSettingsBuilder func
func NewAccountSetPushSettingsBuilder() *AccountSetPushSettingsBuilder {
	return &AccountSetPushSettingsBuilder{api.Params{}}
}

// DeviceID Unique device ID.
func (b *AccountSetPushSettingsBuilder) DeviceID(v string) *AccountSetPushSettingsBuilder {
	b.Params["device_id"] = v
	return b
}

// Settings Push settings in a [vk.com/dev/push_settings|special format].
func (b *AccountSetPushSettingsBuilder) Settings(v string) *AccountSetPushSettingsBuilder {
	b.Params["settings"] = v
	return b
}

// Key Notification key.
func (b *AccountSetPushSettingsBuilder) Key(v string) *AccountSetPushSettingsBuilder {
	b.Params["key"] = v
	return b
}

// Value New value for the key in a [vk.com/dev/push_settings|special format].
func (b *AccountSetPushSettingsBuilder) Value(v []string) *AccountSetPushSettingsBuilder {
	b.Params["value"] = v
	return b
}

// AccountSetSilenceModeBuilder builder
//
// Mutes push notifications for the set period of time.
//
// https://vk.com/dev/account.setSilenceMode
type AccountSetSilenceModeBuilder struct {
	api.Params
}

// NewAccountSetSilenceModeBuilder func
func NewAccountSetSilenceModeBuilder() *AccountSetSilenceModeBuilder {
	return &AccountSetSilenceModeBuilder{api.Params{}}
}

// DeviceID Unique device ID.
func (b *AccountSetSilenceModeBuilder) DeviceID(v string) *AccountSetSilenceModeBuilder {
	b.Params["device_id"] = v
	return b
}

// Time Time in seconds for what notifications should be disabled.
// '-1' to disable forever.
func (b *AccountSetSilenceModeBuilder) Time(v int) *AccountSetSilenceModeBuilder {
	b.Params["time"] = v
	return b
}

// PeerID Destination ID.
// For user 'User ID', e.g. '12345'.
// For chat '2000000000' + 'Chat ID', e.g. '2000000001'.
// For community '- Community ID', e.g. '-12345'.
func (b *AccountSetSilenceModeBuilder) PeerID(v int) *AccountSetSilenceModeBuilder {
	b.Params["peer_id"] = v
	return b
}

// Sound parameter
//
// * '1' — to enable sound in this dialog
//
// * '0' — to disable sound.
// Only if 'peer_id' contains user or community ID.
func (b *AccountSetSilenceModeBuilder) Sound(v int) *AccountSetSilenceModeBuilder {
	b.Params["sound"] = v
	return b
}

// AccountUnbanBuilder builder
//
// https://vk.com/dev/account.unban
type AccountUnbanBuilder struct {
	api.Params
}

// NewAccountUnbanBuilder func
func NewAccountUnbanBuilder() *AccountUnbanBuilder {
	return &AccountUnbanBuilder{api.Params{}}
}

// OwnerID parameter
func (b *AccountUnbanBuilder) OwnerID(v int) *AccountUnbanBuilder {
	b.Params["owner_id"] = v
	return b
}

// AccountUnregisterDeviceBuilder builder
//
// Unsubscribes a device from push notifications.
//
// https://vk.com/dev/account.unregisterDevice
type AccountUnregisterDeviceBuilder struct {
	api.Params
}

// NewAccountUnregisterDeviceBuilder func
func NewAccountUnregisterDeviceBuilder() *AccountUnregisterDeviceBuilder {
	return &AccountUnregisterDeviceBuilder{api.Params{}}
}

// DeviceID Unique device ID.
func (b *AccountUnregisterDeviceBuilder) DeviceID(v string) *AccountUnregisterDeviceBuilder {
	b.Params["device_id"] = v
	return b
}

// Sandbox parameter
func (b *AccountUnregisterDeviceBuilder) Sandbox(v bool) *AccountUnregisterDeviceBuilder {
	b.Params["sandbox"] = v
	return b
}
