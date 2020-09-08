package params // import "github.com/SevereCloud/vksdk/v2/api/params"

import (
	"github.com/SevereCloud/vksdk/v2/api"
)

// NotificationsGetBuilder builder.
//
// Returns a list of notifications about other users' feedback to the current user's wall posts.
//
// https://vk.com/dev/notifications.get
type NotificationsGetBuilder struct {
	api.Params
}

// NewNotificationsGetBuilder func.
func NewNotificationsGetBuilder() *NotificationsGetBuilder {
	return &NotificationsGetBuilder{api.Params{}}
}

// Count number of notifications to return.
func (b *NotificationsGetBuilder) Count(v int) *NotificationsGetBuilder {
	b.Params["count"] = v
	return b
}

// StartFrom parameter.
func (b *NotificationsGetBuilder) StartFrom(v string) *NotificationsGetBuilder {
	b.Params["start_from"] = v
	return b
}

// Filters Type of notifications to return:
//
// * wall — wall posts,
//
// * mentions — mentions in wall posts, comments, or topics,
//
// * comments — comments to wall posts, photos, and videos,
//
// * likes — likes,
//
// * reposted — wall posts that are copied from the current user's wall,
//
// * followers — new followers, 'friends' — accepted friend requests.
func (b *NotificationsGetBuilder) Filters(v []string) *NotificationsGetBuilder {
	b.Params["filters"] = v
	return b
}

// StartTime earliest timestamp (in Unix time) of a notification to return. By default, 24 hours ago.
func (b *NotificationsGetBuilder) StartTime(v int) *NotificationsGetBuilder {
	b.Params["start_time"] = v
	return b
}

// EndTime latest timestamp (in Unix time) of a notification to return. By default, the current time.
func (b *NotificationsGetBuilder) EndTime(v int) *NotificationsGetBuilder {
	b.Params["end_time"] = v
	return b
}

// NotificationsSendMessageBuilder builder.
//
// https://vk.com/dev/notifications.sendMessage
type NotificationsSendMessageBuilder struct {
	api.Params
}

// NewNotificationsSendMessageBuilder func.
func NewNotificationsSendMessageBuilder() *NotificationsSendMessageBuilder {
	return &NotificationsSendMessageBuilder{api.Params{}}
}

// UserIDs parameter.
func (b *NotificationsSendMessageBuilder) UserIDs(v []int) *NotificationsSendMessageBuilder {
	b.Params["user_ids"] = v
	return b
}

// Message parameter.
func (b *NotificationsSendMessageBuilder) Message(v string) *NotificationsSendMessageBuilder {
	b.Params["message"] = v
	return b
}

// Fragment parameter.
func (b *NotificationsSendMessageBuilder) Fragment(v string) *NotificationsSendMessageBuilder {
	b.Params["fragment"] = v
	return b
}

// GroupID parameter.
func (b *NotificationsSendMessageBuilder) GroupID(v int) *NotificationsSendMessageBuilder {
	b.Params["group_id"] = v
	return b
}

// RandomID a unique (API_ID and Sender ID) identifier designed to prevent
// sending the same message again.
//
// The specified random_id is used to check uniqueness of the notification within an hour after sending.
//
// Accessible for versions from 5.107.
func (b *NotificationsSendMessageBuilder) RandomID(v int64) *NotificationsSendMessageBuilder {
	b.Params["random_id"] = v
	return b
}
