package object // import "github.com/SevereCloud/vksdk/5.92/object"

type secureLevel struct {
	Level int `json:"level"`
	UID   int `json:"uid"`
}

type secureSmsNotification struct {
	AppID   int    `json:"app_id"`
	Date    int    `json:"date"`
	ID      int    `json:"id"`
	Message string `json:"message"`
	UserID  int    `json:"user_id"`
}

type secureTokenChecked struct {
	Date    int `json:"date"`
	Expire  int `json:"expire"`
	Success int `json:"success"`
	UserID  int `json:"user_id"`
}

type secureTransaction struct {
	Date    int `json:"date"`
	ID      int `json:"id"`
	UIDFrom int `json:"uid_from"`
	UIDTo   int `json:"uid_to"`
	Votes   int `json:"votes"`
}
