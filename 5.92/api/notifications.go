package api // import "github.com/SevereCloud/vksdk/5.92/api"

// NotificationsGetResponse struct
type NotificationsGetResponse struct{}

// TODO: notifications.get returns a list of notifications about other users' feedback to the current user's wall posts.
// https://vk.com/dev/notifications.get

// NotificationsMarkAsViewedResponse struct
type NotificationsMarkAsViewedResponse int

// NotificationsMarkAsViewed resets the counter of new notifications about other users' feedback to the current user's wall posts.
// https://vk.com/dev/notifications.markAsViewed
func (vk VK) NotificationsMarkAsViewed(params map[string]string) (response NotificationsMarkAsViewedResponse, vkErr Error) {
	vk.requestU("notifications.markAsViewed", params, &response, &vkErr)
	return
}

// NotificationsSendMessageResponse struct
type NotificationsSendMessageResponse []struct {
	UserID int  `json:"user_id"`
	Status bool `json:"status"`
	Error  struct {
		Code        int    `json:"code"`
		Description string `json:"description"`
	} `json:"error"`
}

// NotificationsSendMessage sends notification to the VK Apps user.
// https://vk.com/dev/notifications.sendMessage
func (vk VK) NotificationsSendMessage(params map[string]string) (response NotificationsSendMessageResponse, vkErr Error) {
	vk.requestU("notifications.sendMessage", params, &response, &vkErr)
	return
}
