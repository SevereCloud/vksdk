package object // import "github.com/severecloud/vksdk/5.92/object"

type notificationsFeedback struct {
	Attachments []wallWallpostAttachment `json:"attachments"`
	FromID      int                      `json:"from_id"`
	Geo         baseGeo                  `json:"geo"`
	ID          int                      `json:"id"`
	Likes       baseLikesInfo            `json:"likes"`
	Text        string                   `json:"text"`
	ToID        int                      `json:"to_id"`
}

type notificationsNotification struct {
	Date     int                             `json:"date"`
	Feedback notificationsFeedback           `json:"feedback"`
	Parent   notificationsNotificationParent `json:"parent"`
	Reply    notificationsReply              `json:"reply"`
	Type     string                          `json:"type"`
}

type notificationsNotificationParent struct {
}

type notificationsNotificationsComment struct {
	Date    int          `json:"date"`
	ID      int          `json:"id"`
	OwnerID int          `json:"owner_id"`
	Photo   PhotosPhoto  `json:"photo"`
	Post    WallWallpost `json:"post"`
	Text    string       `json:"text"`
	Topic   boardTopic   `json:"topic"`
	Video   VideoVideo   `json:"video"`
}

type notificationsReply struct {
	Date int `json:"date"`
	ID   int `json:"id"`
	Text int `json:"text"`
}
