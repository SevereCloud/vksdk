package params // import "github.com/SevereCloud/vksdk/api/params"

import (
	"github.com/SevereCloud/vksdk/api"
)

// SecureAddAppEventBuilder builder
//
// Adds user activity information to an application
//
// https://vk.com/dev/secure.addAppEvent
type SecureAddAppEventBuilder struct {
	api.Params
}

// NewSecureAddAppEventBuilder func
func NewSecureAddAppEventBuilder() *SecureAddAppEventBuilder {
	return &SecureAddAppEventBuilder{api.Params{}}
}

// UserID ID of a user to save the data
func (b *SecureAddAppEventBuilder) UserID(v int) {
	b.Params["user_id"] = v
}

// ActivityID there are 2 default activities:
//
// * 1 – level. Works similar to;
//
// * 2 – points, saves points amount/
//
// Any other value is for saving completed missions
func (b *SecureAddAppEventBuilder) ActivityID(v int) {
	b.Params["activity_id"] = v
}

// Value depends on activity_id:
//
// * 1 – number, current level number.
//
// * 2 – number, current user's points amount.
//
// Any other value is ignored
func (b *SecureAddAppEventBuilder) Value(v int) {
	b.Params["value"] = v
}

// SecureCheckTokenBuilder builder
//
// Checks the user authentication in 'IFrame' and 'Flash' apps using the 'access_token' parameter.
//
// https://vk.com/dev/secure.checkToken
type SecureCheckTokenBuilder struct {
	api.Params
}

// NewSecureCheckTokenBuilder func
func NewSecureCheckTokenBuilder() *SecureCheckTokenBuilder {
	return &SecureCheckTokenBuilder{api.Params{}}
}

// Token client 'access_token'
func (b *SecureCheckTokenBuilder) Token(v string) {
	b.Params["token"] = v
}

// IP user 'ip address'.
//
// Note that user may access using the 'ipv6' address, in this case it is required to transmit the 'ipv6' address.
// If not transmitted, the address will not be checked.
func (b *SecureCheckTokenBuilder) IP(v string) {
	b.Params["ip"] = v
}

// SecureGetSMSHistoryBuilder builder
//
// Shows a list of SMS notifications sent by the application using
// [vk.com/dev/secure.sendSMSNotification|secure.sendSMSNotification] method.
//
// https://vk.com/dev/secure.getSMSHistory
type SecureGetSMSHistoryBuilder struct {
	api.Params
}

// NewSecureGetSMSHistoryBuilder func
func NewSecureGetSMSHistoryBuilder() *SecureGetSMSHistoryBuilder {
	return &SecureGetSMSHistoryBuilder{api.Params{}}
}

// UserID parameter
func (b *SecureGetSMSHistoryBuilder) UserID(v int) {
	b.Params["user_id"] = v
}

// DateFrom filter by start date. It is set as UNIX-time.
func (b *SecureGetSMSHistoryBuilder) DateFrom(v int) {
	b.Params["date_from"] = v
}

// DateTo filter by end date. It is set as UNIX-time.
func (b *SecureGetSMSHistoryBuilder) DateTo(v int) {
	b.Params["date_to"] = v
}

// Limit number of returned posts. By default — 1000.
func (b *SecureGetSMSHistoryBuilder) Limit(v int) {
	b.Params["limit"] = v
}

// SecureGetTransactionsHistoryBuilder builder
//
// Shows history of votes transaction between users and the application.
//
// https://vk.com/dev/secure.getTransactionsHistory
type SecureGetTransactionsHistoryBuilder struct {
	api.Params
}

// NewSecureGetTransactionsHistoryBuilder func
func NewSecureGetTransactionsHistoryBuilder() *SecureGetTransactionsHistoryBuilder {
	return &SecureGetTransactionsHistoryBuilder{api.Params{}}
}

// Type parameter
func (b *SecureGetTransactionsHistoryBuilder) Type(v int) {
	b.Params["type"] = v
}

// UIDFrom parameter
func (b *SecureGetTransactionsHistoryBuilder) UIDFrom(v int) {
	b.Params["uid_from"] = v
}

// UIDTo parameter
func (b *SecureGetTransactionsHistoryBuilder) UIDTo(v int) {
	b.Params["uid_to"] = v
}

// DateFrom parameter
func (b *SecureGetTransactionsHistoryBuilder) DateFrom(v int) {
	b.Params["date_from"] = v
}

// DateTo parameter
func (b *SecureGetTransactionsHistoryBuilder) DateTo(v int) {
	b.Params["date_to"] = v
}

// Limit parameter
func (b *SecureGetTransactionsHistoryBuilder) Limit(v int) {
	b.Params["limit"] = v
}

// SecureGetUserLevelBuilder builder
//
// Returns one of the previously set game levels of one or more users in the application.
//
// https://vk.com/dev/secure.getUserLevel
type SecureGetUserLevelBuilder struct {
	api.Params
}

// NewSecureGetUserLevelBuilder func
func NewSecureGetUserLevelBuilder() *SecureGetUserLevelBuilder {
	return &SecureGetUserLevelBuilder{api.Params{}}
}

// UserIDs parameter
func (b *SecureGetUserLevelBuilder) UserIDs(v []int) {
	b.Params["user_ids"] = v
}

// SecureGiveEventStickerBuilder builder
//
// Opens the game achievement and gives the user a sticker
//
// https://vk.com/dev/secure.giveEventSticker
type SecureGiveEventStickerBuilder struct {
	api.Params
}

// NewSecureGiveEventStickerBuilder func
func NewSecureGiveEventStickerBuilder() *SecureGiveEventStickerBuilder {
	return &SecureGiveEventStickerBuilder{api.Params{}}
}

// UserIDs parameter
func (b *SecureGiveEventStickerBuilder) UserIDs(v []int) {
	b.Params["user_ids"] = v
}

// AchievementID parameter
func (b *SecureGiveEventStickerBuilder) AchievementID(v int) {
	b.Params["achievement_id"] = v
}

// SecureSendNotificationBuilder builder
//
// Sends notification to the user.
//
// https://vk.com/dev/secure.sendNotification
type SecureSendNotificationBuilder struct {
	api.Params
}

// NewSecureSendNotificationBuilder func
func NewSecureSendNotificationBuilder() *SecureSendNotificationBuilder {
	return &SecureSendNotificationBuilder{api.Params{}}
}

// UserIDs parameter
func (b *SecureSendNotificationBuilder) UserIDs(v []int) {
	b.Params["user_ids"] = v
}

// UserID parameter
func (b *SecureSendNotificationBuilder) UserID(v int) {
	b.Params["user_id"] = v
}

// Message notification text which should be sent in 'UTF-8' encoding ('254' characters maximum).
func (b *SecureSendNotificationBuilder) Message(v string) {
	b.Params["message"] = v
}

// SecureSendSMSNotificationBuilder builder
//
// Sends 'SMS' notification to a user's mobile device.
//
// https://vk.com/dev/secure.sendSMSNotification
type SecureSendSMSNotificationBuilder struct {
	api.Params
}

// NewSecureSendSMSNotificationBuilder func
func NewSecureSendSMSNotificationBuilder() *SecureSendSMSNotificationBuilder {
	return &SecureSendSMSNotificationBuilder{api.Params{}}
}

// UserID ID of the user to whom SMS notification is sent. The user shall allow the application to send him/her
// notifications (, +1).
func (b *SecureSendSMSNotificationBuilder) UserID(v int) {
	b.Params["user_id"] = v
}

// Message 'SMS' text to be sent in 'UTF-8' encoding. Only Latin letters and numbers are allowed.
// Maximum size is '160' characters.
func (b *SecureSendSMSNotificationBuilder) Message(v string) {
	b.Params["message"] = v
}

// SecureSetCounterBuilder builder
//
// Sets a counter which is shown to the user in bold in the left menu.
//
// https://vk.com/dev/secure.setCounter
type SecureSetCounterBuilder struct {
	api.Params
}

// NewSecureSetCounterBuilder func
func NewSecureSetCounterBuilder() *SecureSetCounterBuilder {
	return &SecureSetCounterBuilder{api.Params{}}
}

// Counters parameter
func (b *SecureSetCounterBuilder) Counters(v []string) {
	b.Params["counters"] = v
}

// UserID parameter
func (b *SecureSetCounterBuilder) UserID(v int) {
	b.Params["user_id"] = v
}

// Counter counter value.
func (b *SecureSetCounterBuilder) Counter(v int) {
	b.Params["counter"] = v
}

// Increment parameter
func (b *SecureSetCounterBuilder) Increment(v bool) {
	b.Params["increment"] = v
}
