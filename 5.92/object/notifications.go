package object // import "github.com/SevereCloud/vksdk/5.92/object"

import "encoding/json"

type notificationsFeedback struct {
	Attachments []wallWallpostAttachment `json:"attachments"`
	FromID      int                      `json:"from_id"`
	Geo         baseGeo                  `json:"geo"`
	ID          int                      `json:"id"`
	Likes       baseLikesInfo            `json:"likes"`
	Text        string                   `json:"text"`
	ToID        int                      `json:"to_id"`
}

// NotificationsNotification struct
type NotificationsNotification struct {
	Date     int                `json:"date"`
	Feedback json.RawMessage    `json:"feedback"`
	Parent   json.RawMessage    `json:"parent"`
	Reply    notificationsReply `json:"reply"`
	Type     string             `json:"type"`
}

type notificationsNotificationsComment struct {
	Date    int          `json:"date"`
	ID      int          `json:"id"`
	OwnerID int          `json:"owner_id"`
	Photo   PhotosPhoto  `json:"photo"`
	Post    WallWallpost `json:"post"`
	Text    string       `json:"text"`
	Topic   BoardTopic   `json:"topic"`
	Video   VideoVideo   `json:"video"`
}

type notificationsReply struct {
	Date string `json:"date"`
	ID   int    `json:"id"`
	Text string `json:"text"`
}
