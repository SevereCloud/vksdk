package params // import "github.com/SevereCloud/vksdk/api/params"

import (
	"github.com/SevereCloud/vksdk/api"
)

// SecureAddAppEventBulder builder
//
// Adds user activity information to an application
//
// https://vk.com/dev/secure.addAppEvent
type SecureAddAppEventBulder struct {
	api.Params
}

// NewSecureAddAppEventBulder func
func NewSecureAddAppEventBulder() *SecureAddAppEventBulder {
	return &SecureAddAppEventBulder{api.Params{}}
}

// UserID ID of a user to save the data
func (b *SecureAddAppEventBulder) UserID(v int) {
	b.Params["user_id"] = v
}

// ActivityID there are 2 default activities: , * 1 – level. Works similar to ,, * 2 – points, saves points amount, Any other value is for saving completed missions
func (b *SecureAddAppEventBulder) ActivityID(v int) {
	b.Params["activity_id"] = v
}

// Value depends on activity_id: * 1 – number, current level number,, * 2 – number, current user's points amount, , Any other value is ignored
func (b *SecureAddAppEventBulder) Value(v int) {
	b.Params["value"] = v
}

// SecureCheckTokenBulder builder
//
// Checks the user authentication in 'IFrame' and 'Flash' apps using the 'access_token' parameter.
//
// https://vk.com/dev/secure.checkToken
type SecureCheckTokenBulder struct {
	api.Params
}

// NewSecureCheckTokenBulder func
func NewSecureCheckTokenBulder() *SecureCheckTokenBulder {
	return &SecureCheckTokenBulder{api.Params{}}
}

// Token client 'access_token'
func (b *SecureCheckTokenBulder) Token(v string) {
	b.Params["token"] = v
}

// IP user 'ip address'. Note that user may access using the 'ipv6' address, in this case it is required to transmit the 'ipv6' address. If not transmitted, the address will not be checked.
func (b *SecureCheckTokenBulder) IP(v string) {
	b.Params["ip"] = v
}

// SecureGetSMSHistoryBulder builder
//
// Shows a list of SMS notifications sent by the application using [vk.com/dev/secure.sendSMSNotification|secure.sendSMSNotification] method.
//
// https://vk.com/dev/secure.getSMSHistory
type SecureGetSMSHistoryBulder struct {
	api.Params
}

// NewSecureGetSMSHistoryBulder func
func NewSecureGetSMSHistoryBulder() *SecureGetSMSHistoryBulder {
	return &SecureGetSMSHistoryBulder{api.Params{}}
}

// UserID parameter
func (b *SecureGetSMSHistoryBulder) UserID(v int) {
	b.Params["user_id"] = v
}

// DateFrom filter by start date. It is set as UNIX-time.
func (b *SecureGetSMSHistoryBulder) DateFrom(v int) {
	b.Params["date_from"] = v
}

// DateTo filter by end date. It is set as UNIX-time.
func (b *SecureGetSMSHistoryBulder) DateTo(v int) {
	b.Params["date_to"] = v
}

// Limit number of returned posts. By default — 1000.
func (b *SecureGetSMSHistoryBulder) Limit(v int) {
	b.Params["limit"] = v
}

// SecureGetTransactionsHistoryBulder builder
//
// Shows history of votes transaction between users and the application.
//
// https://vk.com/dev/secure.getTransactionsHistory
type SecureGetTransactionsHistoryBulder struct {
	api.Params
}

// NewSecureGetTransactionsHistoryBulder func
func NewSecureGetTransactionsHistoryBulder() *SecureGetTransactionsHistoryBulder {
	return &SecureGetTransactionsHistoryBulder{api.Params{}}
}

// Type parameter
func (b *SecureGetTransactionsHistoryBulder) Type(v int) {
	b.Params["type"] = v
}

// UIDFrom parameter
func (b *SecureGetTransactionsHistoryBulder) UIDFrom(v int) {
	b.Params["uid_from"] = v
}

// UIDTo parameter
func (b *SecureGetTransactionsHistoryBulder) UIDTo(v int) {
	b.Params["uid_to"] = v
}

// DateFrom parameter
func (b *SecureGetTransactionsHistoryBulder) DateFrom(v int) {
	b.Params["date_from"] = v
}

// DateTo parameter
func (b *SecureGetTransactionsHistoryBulder) DateTo(v int) {
	b.Params["date_to"] = v
}

// Limit parameter
func (b *SecureGetTransactionsHistoryBulder) Limit(v int) {
	b.Params["limit"] = v
}

// SecureGetUserLevelBulder builder
//
// Returns one of the previously set game levels of one or more users in the application.
//
// https://vk.com/dev/secure.getUserLevel
type SecureGetUserLevelBulder struct {
	api.Params
}

// NewSecureGetUserLevelBulder func
func NewSecureGetUserLevelBulder() *SecureGetUserLevelBulder {
	return &SecureGetUserLevelBulder{api.Params{}}
}

// UserIDs parameter
func (b *SecureGetUserLevelBulder) UserIDs(v []int) {
	b.Params["user_ids"] = v
}

// SecureGiveEventStickerBulder builder
//
// Opens the game achievement and gives the user a sticker
//
// https://vk.com/dev/secure.giveEventSticker
type SecureGiveEventStickerBulder struct {
	api.Params
}

// NewSecureGiveEventStickerBulder func
func NewSecureGiveEventStickerBulder() *SecureGiveEventStickerBulder {
	return &SecureGiveEventStickerBulder{api.Params{}}
}

// UserIDs parameter
func (b *SecureGiveEventStickerBulder) UserIDs(v []int) {
	b.Params["user_ids"] = v
}

// AchievementID parameter
func (b *SecureGiveEventStickerBulder) AchievementID(v int) {
	b.Params["achievement_id"] = v
}

// SecureSendNotificationBulder builder
//
// Sends notification to the user.
//
// https://vk.com/dev/secure.sendNotification
type SecureSendNotificationBulder struct {
	api.Params
}

// NewSecureSendNotificationBulder func
func NewSecureSendNotificationBulder() *SecureSendNotificationBulder {
	return &SecureSendNotificationBulder{api.Params{}}
}

// UserIDs parameter
func (b *SecureSendNotificationBulder) UserIDs(v []int) {
	b.Params["user_ids"] = v
}

// UserID parameter
func (b *SecureSendNotificationBulder) UserID(v int) {
	b.Params["user_id"] = v
}

// Message notification text which should be sent in 'UTF-8' encoding ('254' characters maximum).
func (b *SecureSendNotificationBulder) Message(v string) {
	b.Params["message"] = v
}

// SecureSendSMSNotificationBulder builder
//
// Sends 'SMS' notification to a user's mobile device.
//
// https://vk.com/dev/secure.sendSMSNotification
type SecureSendSMSNotificationBulder struct {
	api.Params
}

// NewSecureSendSMSNotificationBulder func
func NewSecureSendSMSNotificationBulder() *SecureSendSMSNotificationBulder {
	return &SecureSendSMSNotificationBulder{api.Params{}}
}

// UserID ID of the user to whom SMS notification is sent. The user shall allow the application to send him/her notifications (, +1).
func (b *SecureSendSMSNotificationBulder) UserID(v int) {
	b.Params["user_id"] = v
}

// Message 'SMS' text to be sent in 'UTF-8' encoding. Only Latin letters and numbers are allowed. Maximum size is '160' characters.
func (b *SecureSendSMSNotificationBulder) Message(v string) {
	b.Params["message"] = v
}

// SecureSetCounterBulder builder
//
// Sets a counter which is shown to the user in bold in the left menu.
//
// https://vk.com/dev/secure.setCounter
type SecureSetCounterBulder struct {
	api.Params
}

// NewSecureSetCounterBulder func
func NewSecureSetCounterBulder() *SecureSetCounterBulder {
	return &SecureSetCounterBulder{api.Params{}}
}

// Counters parameter
func (b *SecureSetCounterBulder) Counters(v []string) {
	b.Params["counters"] = v
}

// UserID parameter
func (b *SecureSetCounterBulder) UserID(v int) {
	b.Params["user_id"] = v
}

// Counter counter value.
func (b *SecureSetCounterBulder) Counter(v int) {
	b.Params["counter"] = v
}

// Increment parameter
func (b *SecureSetCounterBulder) Increment(v bool) {
	b.Params["increment"] = v
}
