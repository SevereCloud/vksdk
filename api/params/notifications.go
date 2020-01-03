package params // import "github.com/SevereCloud/vksdk/api/params"

import (
	"github.com/SevereCloud/vksdk/api"
)

// NotificationsGetBulder builder
//
// Returns a list of notifications about other users' feedback to the current user's wall posts.
//
// https://vk.com/dev/notifications.get
type NotificationsGetBulder struct {
	api.Params
}

// NewNotificationsGetBulder func
func NewNotificationsGetBulder() *NotificationsGetBulder {
	return &NotificationsGetBulder{api.Params{}}
}

// Count Number of notifications to return.
func (b *NotificationsGetBulder) Count(v int) {
	b.Params["count"] = v
}

// StartFrom parameter
func (b *NotificationsGetBulder) StartFrom(v string) {
	b.Params["start_from"] = v
}

// Filters Type of notifications to return: 'wall' — wall posts, 'mentions' — mentions in wall posts, comments, or topics, 'comments' — comments to wall posts, photos, and videos, 'likes' — likes, 'reposted' — wall posts that are copied from the current user's wall, 'followers' — new followers, 'friends' — accepted friend requests
func (b *NotificationsGetBulder) Filters(v []string) {
	b.Params["filters"] = v
}

// StartTime Earliest timestamp (in Unix time) of a notification to return. By default, 24 hours ago.
func (b *NotificationsGetBulder) StartTime(v int) {
	b.Params["start_time"] = v
}

// EndTime Latest timestamp (in Unix time) of a notification to return. By default, the current time.
func (b *NotificationsGetBulder) EndTime(v int) {
	b.Params["end_time"] = v
}

// NotificationsSendMessageBulder builder
//
// https://vk.com/dev/notifications.sendMessage
type NotificationsSendMessageBulder struct {
	api.Params
}

// NewNotificationsSendMessageBulder func
func NewNotificationsSendMessageBulder() *NotificationsSendMessageBulder {
	return &NotificationsSendMessageBulder{api.Params{}}
}

// UserIDs parameter
func (b *NotificationsSendMessageBulder) UserIDs(v []int) {
	b.Params["user_ids"] = v
}

// Message parameter
func (b *NotificationsSendMessageBulder) Message(v string) {
	b.Params["message"] = v
}

// Fragment parameter
func (b *NotificationsSendMessageBulder) Fragment(v string) {
	b.Params["fragment"] = v
}

// GroupID parameter
func (b *NotificationsSendMessageBulder) GroupID(v int) {
	b.Params["group_id"] = v
}
