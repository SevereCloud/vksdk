package api // import "github.com/SevereCloud/vksdk/api"

import (
	"github.com/SevereCloud/vksdk/object"
)

// SecureAddAppEventResponse struct
type SecureAddAppEventResponse int // FIXME: not found documentation. https://github.com/VKCOM/vk-api-schema/issues/98

// SecureAddAppEvent Adds user activity information to an application
//
// https://vk.com/dev/secure.addAppEvent
func (vk *VK) SecureAddAppEvent(params map[string]string) (response SecureAddAppEventResponse, vkErr Error) {
	vk.RequestUnmarshal("secure.addAppEvent", params, &response, &vkErr)
	return
}

// SecureCheckTokenResponse struct
type SecureCheckTokenResponse object.SecureTokenChecked

// SecureCheckToken Checks the user authentification in IFrame and Flash apps using the access_token parameter.
//
// https://vk.com/dev/secure.checkToken
func (vk *VK) SecureCheckToken(params map[string]string) (response SecureCheckTokenResponse, vkErr Error) {
	vk.RequestUnmarshal("secure.checkToken", params, &response, &vkErr)
	return
}

// SecureGetAppBalance Returns payment balance of the application in hundredth of a vote.
//
// https://vk.com/dev/secure.getAppBalance
func (vk *VK) SecureGetAppBalance(params map[string]string) (response int, vkErr Error) {
	vk.RequestUnmarshal("secure.getAppBalance", params, &response, &vkErr)
	return
}

// SecureGetSMSHistoryResponse struct
type SecureGetSMSHistoryResponse []object.SecureSmsNotification

// SecureGetSMSHistory Shows a list of SMS notifications sent by the application using secure.sendSMSNotification method.
//
// https://vk.com/dev/secure.getSMSHistory
func (vk *VK) SecureGetSMSHistory(params map[string]string) (response SecureGetSMSHistoryResponse, vkErr Error) {
	vk.RequestUnmarshal("secure.getSMSHistory", params, &response, &vkErr)
	return
}

// SecureGetTransactionsHistoryResponse struct
type SecureGetTransactionsHistoryResponse []object.SecureTransaction

// SecureGetTransactionsHistory Shows history of votes transaction between users and the application.
//
// https://vk.com/dev/secure.getTransactionsHistory
func (vk *VK) SecureGetTransactionsHistory(params map[string]string) (response SecureGetTransactionsHistoryResponse, vkErr Error) {
	vk.RequestUnmarshal("secure.getTransactionsHistory", params, &response, &vkErr)
	return
}

// SecureGetUserLevelResponse struct
type SecureGetUserLevelResponse []object.SecureLevel

// SecureGetUserLevel Returns one of the previously set game levels of one or more users in the application.
//
// https://vk.com/dev/secure.getUserLevel
func (vk *VK) SecureGetUserLevel(params map[string]string) (response SecureGetUserLevelResponse, vkErr Error) {
	vk.RequestUnmarshal("secure.getUserLevel", params, &response, &vkErr)
	return
}

// SecureGiveEventStickerResponse struct
type SecureGiveEventStickerResponse []struct {
	UserID int    `json:"user_id"`
	Status string `json:"status"`
}

// SecureGiveEventSticker method
//
// https://vk.com/dev/secure.giveEventSticker
func (vk *VK) SecureGiveEventSticker(params map[string]string) (response SecureGiveEventStickerResponse, vkErr Error) {
	vk.RequestUnmarshal("secure.giveEventSticker", params, &response, &vkErr)
	return
}

// SecureSendNotificationResponse struct
type SecureSendNotificationResponse []int //User ID

// SecureSendNotification Sends notification to the user.
//
// https://vk.com/dev/secure.sendNotification
func (vk *VK) SecureSendNotification(params map[string]string) (response SecureSendNotificationResponse, vkErr Error) {
	vk.RequestUnmarshal("secure.sendNotification", params, &response, &vkErr)
	return
}

// SecureSendSMSNotification Sends SMS notification to a user's mobile device.
//
// https://vk.com/dev/secure.sendSMSNotification
func (vk *VK) SecureSendSMSNotification(params map[string]string) (response int, vkErr Error) {
	vk.RequestUnmarshal("secure.sendSMSNotification", params, &response, &vkErr)
	return
}

// SecureSetCounter Sets a counter which is shown to the user in bold in the left menu.
//
// https://vk.com/dev/secure.setCounter
func (vk *VK) SecureSetCounter(params map[string]string) (response int, vkErr Error) {
	vk.RequestUnmarshal("secure.setCounter", params, &response, &vkErr)
	return
}
