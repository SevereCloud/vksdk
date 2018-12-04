package callback

import "github.com/severecloud/vksdk/object"

// MessageNewObject struct
type MessageNewObject object.MessagesMessage

// MessageReplyObject struct
type MessageReplyObject object.MessagesMessage

// MessagesAllowObject struct
type MessagesAllowObject struct {
	UserID int    `json:"user_id"`
	Key    string `json:"key"`
}

// MessagesDenyObject struct
type MessagesDenyObject struct {
	UserID int `json:"user_id"`
}

// PhotoNewObject struct
type PhotoNewObject object.PhotosPhoto

// PhotoCommentNewObject struct
type PhotoCommentNewObject struct{}

// PhotoCommentEditObject struct
type PhotoCommentEditObject struct{}

// PhotoCommentRestoreObject struct
type PhotoCommentRestoreObject struct{}

// PhotoCommentDeleteObject struct
type PhotoCommentDeleteObject struct {
	OwnerID   int `json:"owner_id"`
	ID        int `json:"id"`
	UserID    int `json:"user_id"`
	DeleterID int `json:"deleter_id"`
	PhotoID   int `json:"photo_id"`
}

// AudioNewObject struct
type AudioNewObject object.AudioAudioFull

// VideoNewObject struct
type VideoNewObject object.VideoVideo

// VideoCommentNewObject struct
type VideoCommentNewObject struct{}

// VideoCommentEditObject struct
type VideoCommentEditObject struct{}

// VideoCommentRestoreObject struct
type VideoCommentRestoreObject struct{}

// VideoCommentDeleteObject struct
type VideoCommentDeleteObject struct {
	OwnerID   int `json:"owner_id"`
	ID        int `json:"id"`
	UserID    int `json:"user_id"`
	DeleterID int `json:"deleter_id"`
	VideoID   int `json:"video_id"`
}

// WallPostNewObject struct
type WallPostNewObject struct{}

// WallRepostObject struct
type WallRepostObject struct{}

// WallReplyNewObject struct
type WallReplyNewObject struct{}

// WallReplyEditObject struct
type WallReplyEditObject struct{}

// WallReplyRestoreObject struct
type WallReplyRestoreObject struct{}

// WallReplyDeleteObject struct
type WallReplyDeleteObject struct {
	OwnerID   int `json:"owner_id"`
	ID        int `json:"id"`
	DeleterID int `json:"deleter_id"`
	PostID    int `json:"post_id"`
}

// BoardPostNewObject struct
type BoardPostNewObject struct{}

// BoardPostEditObject struct
type BoardPostEditObject struct{}

// BoardPostRestoreObject struct
type BoardPostRestoreObject struct{}

// BoardPostDeleteObject struct
type BoardPostDeleteObject struct {
	TopicOwnerID int `json:"topic_owner_id"`
	TopicID      int `json:"topic_id"`
	ID           int `json:"id"`
}

// MarketCommentNewObject struct
type MarketCommentNewObject struct{}

// MarketCommentEditObject struct
type MarketCommentEditObject struct{}

// MarketCommentRestoreObject struct
type MarketCommentRestoreObject struct{}

// MarketCommentDeleteObject struct
type MarketCommentDeleteObject struct {
	OwnerID   int `json:"owner_id"`
	ID        int `json:"id"`
	UserID    int `json:"user_id"`
	DeleterID int `json:"deleter_id"`
	ItemID    int `json:"item_id"`
}

// GroupLeaveObject struct
type GroupLeaveObject struct {
	UserID int `json:"user_id"`
	Self   int `json:"self"`
}

// GroupJoinObject struct
type GroupJoinObject struct {
	UserID   int    `json:"user_id"`
	JoinType string `json:"join_type"`
}

// UserBlockObject struct
type UserBlockObject struct {
	AdminID     int    `json:"admin_id"`
	UserID      int    `json:"user_id"`
	UnblockDate int    `json:"unblock_date"`
	Reason      int    `json:"reason"`
	Comment     string `json:"comment"`
}

// UserUnblockObject struct
type UserUnblockObject struct {
	AdminID   int `json:"admin_id"`
	UserID    int `json:"user_id"`
	ByEndDate int `json:"by_end_date"`
}

// PollVoteNewObject struct
// BUG при голосовании за несколько вариантов, возвращается только один
type PollVoteNewObject struct {
	OwnerID  int `json:"owner_id"`
	PollID   int `json:"poll_id"`
	OptionID int `json:"option_id"`
	UserID   int `json:"user_id"`
}

// GroupOfficersEditObject struct
type GroupOfficersEditObject struct {
	AdminID  int `json:"admin_id"`
	UserID   int `json:"user_id"`
	LevelOld int `json:"level_old"`
	LevelNew int `json:"level_new"`
}

type changes struct {
	OldValue string `json:"old_value"`
	NewValue string `json:"new_value"`
}

type changesInt struct {
	OldValue int `json:"old_value"`
	NewValue int `json:"new_value"`
}

// GroupChangeSettingsObject struct
// спасибо vk.com/eee
// BUG Phone https://vk.com/bugtracker?act=show&id=64240
type GroupChangeSettingsObject struct {
	UserID  int `json:"user_id"`
	Changes struct {
		Title                 changes    `json:"title"`
		Description           changes    `json:"description"`
		Access                changesInt `json:"access"`
		ScreenName            changes    `json:"screen_name"`
		PublicCategory        changesInt `json:"public_category"`
		PublicSubcategory     changesInt `json:"public_subcategory"`
		AgeLimits             changesInt `json:"age_limits"`
		Website               changes    `json:"website"`
		StatusDefault         changes    `json:"status_default"`
		Wall                  changesInt `json:"wall"`                    // на основе ответа
		Replies               changesInt `json:"replies"`                 // на основе ответа
		Topics                changesInt `json:"topics"`                  // на основе ответа
		Audio                 changesInt `json:"audio"`                   // на основе ответа
		Photos                changesInt `json:"photos"`                  // на основе ответа
		Video                 changesInt `json:"video"`                   // на основе ответа
		Market                changesInt `json:"market"`                  // на основе ответа
		Docs                  changesInt `json:"docs"`                    // на основе ответа
		Messages              changesInt `json:"messages"`                // на основе ответа
		EventGroupID          changesInt `json:"event_group_id"`          // на основе ответа
		Links                 changes    `json:"links"`                   // на основе ответа
		Email                 changes    `json:"email"`                   // на основе ответа
		EventStartDate        changesInt `json:"event_start_date::"`      // на основе ответа
		EventFinishDate       changesInt `json:"event_finish_date:"`      // на основе ответа
		Subject               changes    `json:"subject"`                 // на основе ответа
		MarketWiki            changes    `json:"market_wiki"`             // на основе ответа
		DisableMarketComments changesInt `json:"disable_market_comments"` // на основе ответа
		Phone                 changesInt `json:"phone"`                   // на основе ответа
		CountryID             changesInt `json:"country_id"`              // на основе ответа
		CityID                changesInt `json:"city_id"`                 // на основе ответа
	} `json:"changes"`
}

// GroupChangePhotoObject struct
type GroupChangePhotoObject struct {
	UserID int                `json:"user_id"`
	Photo  object.PhotosPhoto `json:"photo"`
}

// VkpayTransactionObject struct
type VkpayTransactionObject struct {
	FromID      int    `json:"from_id"`
	Amount      string `json:"amount"` // FIXME string or int?
	Description string `json:"description"`
	Date        string `json:"date"` // FIXME string or int?
}
