package params // import "github.com/SevereCloud/vksdk/v2/api/params"

import (
	"github.com/SevereCloud/vksdk/v2/api"
)

// SecureAddAppEventBuilder builder.
//
// Adds user activity information to an application.
//
// https://vk.com/dev/secure.addAppEvent
type SecureAddAppEventBuilder struct {
	api.Params
}

// NewSecureAddAppEventBuilder func.
func NewSecureAddAppEventBuilder() *SecureAddAppEventBuilder {
	return &SecureAddAppEventBuilder{api.Params{}}
}

// UserID ID of a user to save the data.
func (b *SecureAddAppEventBuilder) UserID(v int) *SecureAddAppEventBuilder {
	b.Params["user_id"] = v
	return b
}

// ActivityID there are 2 default activities:
//
// * 1 – level. Works similar to;
//
// * 2 – points, saves points amount.
//
// Any other value is for saving completed missions.
func (b *SecureAddAppEventBuilder) ActivityID(v int) *SecureAddAppEventBuilder {
	b.Params["activity_id"] = v
	return b
}

// Value depends on activity_id:
//
// * 1 – number, current level number.
//
// * 2 – number, current user's points amount.
//
// Any other value is ignored.
func (b *SecureAddAppEventBuilder) Value(v int) *SecureAddAppEventBuilder {
	b.Params["value"] = v
	return b
}

// SecureCheckTokenBuilder builder.
//
// Checks the user authentication in 'IFrame' and 'Flash' apps using the 'access_token' parameter.
//
// https://vk.com/dev/secure.checkToken
type SecureCheckTokenBuilder struct {
	api.Params
}

// NewSecureCheckTokenBuilder func.
func NewSecureCheckTokenBuilder() *SecureCheckTokenBuilder {
	return &SecureCheckTokenBuilder{api.Params{}}
}

// Token client 'access_token'.
func (b *SecureCheckTokenBuilder) Token(v string) *SecureCheckTokenBuilder {
	b.Params["token"] = v
	return b
}

// IP user 'ip address'.
//
// Note that user may access using the 'ipv6' address, in this case it is required to transmit the 'ipv6' address.
// If not transmitted, the address will not be checked.
func (b *SecureCheckTokenBuilder) IP(v string) *SecureCheckTokenBuilder {
	b.Params["ip"] = v
	return b
}

// SecureGetSMSHistoryBuilder builder.
//
// Shows a list of SMS notifications sent by the application using
// [vk.com/dev/secure.sendSMSNotification|secure.sendSMSNotification] method.
//
// https://vk.com/dev/secure.getSMSHistory
type SecureGetSMSHistoryBuilder struct {
	api.Params
}

// NewSecureGetSMSHistoryBuilder func.
func NewSecureGetSMSHistoryBuilder() *SecureGetSMSHistoryBuilder {
	return &SecureGetSMSHistoryBuilder{api.Params{}}
}

// UserID parameter.
func (b *SecureGetSMSHistoryBuilder) UserID(v int) *SecureGetSMSHistoryBuilder {
	b.Params["user_id"] = v
	return b
}

// DateFrom filter by start date. It is set as UNIX-time.
func (b *SecureGetSMSHistoryBuilder) DateFrom(v int) *SecureGetSMSHistoryBuilder {
	b.Params["date_from"] = v
	return b
}

// DateTo filter by end date. It is set as UNIX-time.
func (b *SecureGetSMSHistoryBuilder) DateTo(v int) *SecureGetSMSHistoryBuilder {
	b.Params["date_to"] = v
	return b
}

// Limit number of returned posts. By default — 1000.
func (b *SecureGetSMSHistoryBuilder) Limit(v int) *SecureGetSMSHistoryBuilder {
	b.Params["limit"] = v
	return b
}

// SecureGetTransactionsHistoryBuilder builder.
//
// Shows history of votes transaction between users and the application.
//
// https://vk.com/dev/secure.getTransactionsHistory
type SecureGetTransactionsHistoryBuilder struct {
	api.Params
}

// NewSecureGetTransactionsHistoryBuilder func.
func NewSecureGetTransactionsHistoryBuilder() *SecureGetTransactionsHistoryBuilder {
	return &SecureGetTransactionsHistoryBuilder{api.Params{}}
}

// Type parameter.
func (b *SecureGetTransactionsHistoryBuilder) Type(v int) *SecureGetTransactionsHistoryBuilder {
	b.Params["type"] = v
	return b
}

// UIDFrom parameter.
func (b *SecureGetTransactionsHistoryBuilder) UIDFrom(v int) *SecureGetTransactionsHistoryBuilder {
	b.Params["uid_from"] = v
	return b
}

// UIDTo parameter.
func (b *SecureGetTransactionsHistoryBuilder) UIDTo(v int) *SecureGetTransactionsHistoryBuilder {
	b.Params["uid_to"] = v
	return b
}

// DateFrom parameter.
func (b *SecureGetTransactionsHistoryBuilder) DateFrom(v int) *SecureGetTransactionsHistoryBuilder {
	b.Params["date_from"] = v
	return b
}

// DateTo parameter.
func (b *SecureGetTransactionsHistoryBuilder) DateTo(v int) *SecureGetTransactionsHistoryBuilder {
	b.Params["date_to"] = v
	return b
}

// Limit parameter.
func (b *SecureGetTransactionsHistoryBuilder) Limit(v int) *SecureGetTransactionsHistoryBuilder {
	b.Params["limit"] = v
	return b
}

// SecureGetUserLevelBuilder builder.
//
// Returns one of the previously set game levels of one or more users in the application.
//
// https://vk.com/dev/secure.getUserLevel
type SecureGetUserLevelBuilder struct {
	api.Params
}

// NewSecureGetUserLevelBuilder func.
func NewSecureGetUserLevelBuilder() *SecureGetUserLevelBuilder {
	return &SecureGetUserLevelBuilder{api.Params{}}
}

// UserIDs parameter.
func (b *SecureGetUserLevelBuilder) UserIDs(v []int) *SecureGetUserLevelBuilder {
	b.Params["user_ids"] = v
	return b
}

// SecureGiveEventStickerBuilder builder.
//
// Opens the game achievement and gives the user a sticker.
//
// https://vk.com/dev/secure.giveEventSticker
type SecureGiveEventStickerBuilder struct {
	api.Params
}

// NewSecureGiveEventStickerBuilder func.
func NewSecureGiveEventStickerBuilder() *SecureGiveEventStickerBuilder {
	return &SecureGiveEventStickerBuilder{api.Params{}}
}

// UserIDs parameter.
func (b *SecureGiveEventStickerBuilder) UserIDs(v []int) *SecureGiveEventStickerBuilder {
	b.Params["user_ids"] = v
	return b
}

// AchievementID parameter.
func (b *SecureGiveEventStickerBuilder) AchievementID(v int) *SecureGiveEventStickerBuilder {
	b.Params["achievement_id"] = v
	return b
}

// SecureSendNotificationBuilder builder.
//
// Sends notification to the user.
//
// https://vk.com/dev/secure.sendNotification
type SecureSendNotificationBuilder struct {
	api.Params
}

// NewSecureSendNotificationBuilder func.
func NewSecureSendNotificationBuilder() *SecureSendNotificationBuilder {
	return &SecureSendNotificationBuilder{api.Params{}}
}

// UserIDs parameter.
func (b *SecureSendNotificationBuilder) UserIDs(v []int) *SecureSendNotificationBuilder {
	b.Params["user_ids"] = v
	return b
}

// UserID parameter.
func (b *SecureSendNotificationBuilder) UserID(v int) *SecureSendNotificationBuilder {
	b.Params["user_id"] = v
	return b
}

// Message notification text which should be sent in 'UTF-8' encoding ('254' characters maximum).
func (b *SecureSendNotificationBuilder) Message(v string) *SecureSendNotificationBuilder {
	b.Params["message"] = v
	return b
}

// SecureSendSMSNotificationBuilder builder.
//
// Sends 'SMS' notification to a user's mobile device.
//
// https://vk.com/dev/secure.sendSMSNotification
type SecureSendSMSNotificationBuilder struct {
	api.Params
}

// NewSecureSendSMSNotificationBuilder func.
func NewSecureSendSMSNotificationBuilder() *SecureSendSMSNotificationBuilder {
	return &SecureSendSMSNotificationBuilder{api.Params{}}
}

// UserID ID of the user to whom SMS notification is sent. The user shall allow the application to send him/her
// notifications (, +1).
func (b *SecureSendSMSNotificationBuilder) UserID(v int) *SecureSendSMSNotificationBuilder {
	b.Params["user_id"] = v
	return b
}

// Message 'SMS' text to be sent in 'UTF-8' encoding. Only Latin letters and numbers are allowed.
// Maximum size is '160' characters.
func (b *SecureSendSMSNotificationBuilder) Message(v string) *SecureSendSMSNotificationBuilder {
	b.Params["message"] = v
	return b
}

// SecureSetCounterBuilder builder.
//
// Sets a counter which is shown to the user in bold in the left menu.
//
// https://vk.com/dev/secure.setCounter
type SecureSetCounterBuilder struct {
	api.Params
}

// NewSecureSetCounterBuilder func.
func NewSecureSetCounterBuilder() *SecureSetCounterBuilder {
	return &SecureSetCounterBuilder{api.Params{}}
}

// Counters parameter.
func (b *SecureSetCounterBuilder) Counters(v []string) *SecureSetCounterBuilder {
	b.Params["counters"] = v
	return b
}

// UserID parameter.
func (b *SecureSetCounterBuilder) UserID(v int) *SecureSetCounterBuilder {
	b.Params["user_id"] = v
	return b
}

// Counter counter value.
func (b *SecureSetCounterBuilder) Counter(v int) *SecureSetCounterBuilder {
	b.Params["counter"] = v
	return b
}

// Increment parameter.
func (b *SecureSetCounterBuilder) Increment(v bool) *SecureSetCounterBuilder {
	b.Params["increment"] = v
	return b
}
