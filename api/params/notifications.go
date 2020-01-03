package params // import "github.com/SevereCloud/vksdk/api/params"

import (
	"github.com/SevereCloud/vksdk/api"
)

// NotificationsGetBuilder builder
//
// Returns a list of notifications about other users' feedback to the current user's wall posts.
//
// https://vk.com/dev/notifications.get
type NotificationsGetBuilder struct {
	api.Params
}

// NewNotificationsGetBuilder func
func NewNotificationsGetBuilder() *NotificationsGetBuilder {
	return &NotificationsGetBuilder{api.Params{}}
}

// Count Number of notifications to return.
func (b *NotificationsGetBuilder) Count(v int) {
	b.Params["count"] = v
}

// StartFrom parameter
func (b *NotificationsGetBuilder) StartFrom(v string) {
	b.Params["start_from"] = v
}

// Filters Type of notifications to return: 'wall' — wall posts, 'mentions' — mentions in wall posts, comments, or topics, 'comments' — comments to wall posts, photos, and videos, 'likes' — likes, 'reposted' — wall posts that are copied from the current user's wall, 'followers' — new followers, 'friends' — accepted friend requests
func (b *NotificationsGetBuilder) Filters(v []string) {
	b.Params["filters"] = v
}

// StartTime Earliest timestamp (in Unix time) of a notification to return. By default, 24 hours ago.
func (b *NotificationsGetBuilder) StartTime(v int) {
	b.Params["start_time"] = v
}

// EndTime Latest timestamp (in Unix time) of a notification to return. By default, the current time.
func (b *NotificationsGetBuilder) EndTime(v int) {
	b.Params["end_time"] = v
}

// NotificationsSendMessageBuilder builder
//
// https://vk.com/dev/notifications.sendMessage
type NotificationsSendMessageBuilder struct {
	api.Params
}

// NewNotificationsSendMessageBuilder func
func NewNotificationsSendMessageBuilder() *NotificationsSendMessageBuilder {
	return &NotificationsSendMessageBuilder{api.Params{}}
}

// UserIDs parameter
func (b *NotificationsSendMessageBuilder) UserIDs(v []int) {
	b.Params["user_ids"] = v
}

// Message parameter
func (b *NotificationsSendMessageBuilder) Message(v string) {
	b.Params["message"] = v
}

// Fragment parameter
func (b *NotificationsSendMessageBuilder) Fragment(v string) {
	b.Params["fragment"] = v
}

// GroupID parameter
func (b *NotificationsSendMessageBuilder) GroupID(v int) {
	b.Params["group_id"] = v
}
