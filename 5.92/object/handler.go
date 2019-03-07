package object // import "github.com/severecloud/vksdk/5.92/object"

// MessageNewFunc func
type MessageNewFunc func(MessageNewObject, int)

// MessageNewObject struct
type MessageNewObject MessagesMessage

// MessageReplyFunc func
type MessageReplyFunc func(MessageReplyObject, int)

// MessageReplyObject struct
type MessageReplyObject MessagesMessage

// MessageEditFunc func
type MessageEditFunc func(MessageEditObject, int)

// MessageEditObject struct
type MessageEditObject MessagesMessage

// MessageAllowFunc func
type MessageAllowFunc func(MessageAllowObject, int)

// MessageAllowObject struct
type MessageAllowObject struct {
	UserID int    `json:"user_id"`
	Key    string `json:"key"`
}

// MessageDenyFunc func
type MessageDenyFunc func(MessageDenyObject, int)

// MessageDenyObject struct
type MessageDenyObject struct {
	UserID int `json:"user_id"`
}

// MessageTypingStateFunc func
type MessageTypingStateFunc func(MessageTypingStateObject, int)

// MessageTypingStateObject struct
type MessageTypingStateObject struct {
	State  string `json:"state"`
	FromID int    `json:"from_id"`
	ToID   int    `json:"to_id"`
}

// PhotoNewFunc func
type PhotoNewFunc func(PhotoNewObject, int)

// PhotoNewObject struct
type PhotoNewObject PhotosPhoto

// PhotoCommentNewFunc func
type PhotoCommentNewFunc func(PhotoCommentNewObject, int)

// PhotoCommentNewObject struct
type PhotoCommentNewObject wallWallComment

// PhotoCommentEditFunc func
type PhotoCommentEditFunc func(PhotoCommentEditObject, int)

// PhotoCommentEditObject struct
type PhotoCommentEditObject wallWallComment

// PhotoCommentRestoreFunc func
type PhotoCommentRestoreFunc func(PhotoCommentRestoreObject, int)

// PhotoCommentRestoreObject struct
type PhotoCommentRestoreObject wallWallComment

// PhotoCommentDeleteFunc func
type PhotoCommentDeleteFunc func(PhotoCommentDeleteObject, int)

// PhotoCommentDeleteObject struct
type PhotoCommentDeleteObject struct {
	OwnerID   int `json:"owner_id"`
	ID        int `json:"id"`
	UserID    int `json:"user_id"`
	DeleterID int `json:"deleter_id"`
	PhotoID   int `json:"photo_id"`
}

// AudioNewFunc func
type AudioNewFunc func(AudioNewObject, int)

// AudioNewObject struct
type AudioNewObject AudioAudioFull

// VideoNewFunc func
type VideoNewFunc func(VideoNewObject, int)

// VideoNewObject struct
type VideoNewObject VideoVideo

// VideoCommentNewFunc func
type VideoCommentNewFunc func(VideoCommentNewObject, int)

// VideoCommentNewObject struct
type VideoCommentNewObject wallWallComment

// VideoCommentEditFunc func
type VideoCommentEditFunc func(VideoCommentEditObject, int)

// VideoCommentEditObject struct
type VideoCommentEditObject wallWallComment

// VideoCommentRestoreFunc func
type VideoCommentRestoreFunc func(VideoCommentRestoreObject, int)

// VideoCommentRestoreObject struct
type VideoCommentRestoreObject wallWallComment

// VideoCommentDeleteFunc func
type VideoCommentDeleteFunc func(VideoCommentDeleteObject, int)

// VideoCommentDeleteObject struct
type VideoCommentDeleteObject struct {
	OwnerID   int `json:"owner_id"`
	ID        int `json:"id"`
	UserID    int `json:"user_id"`
	DeleterID int `json:"deleter_id"`
	VideoID   int `json:"video_id"`
}

// WallPostNewFunc func
type WallPostNewFunc func(WallPostNewObject, int)

// WallPostNewObject struct
type WallPostNewObject WallWallpost

// WallRepostFunc func
type WallRepostFunc func(WallRepostObject, int)

// WallRepostObject struct
type WallRepostObject WallWallpost

// WallReplyNewFunc func
type WallReplyNewFunc func(WallReplyNewObject, int)

// WallReplyNewObject struct
type WallReplyNewObject wallWallComment

// WallReplyEditFunc func
type WallReplyEditFunc func(WallReplyEditObject, int)

// WallReplyEditObject struct
type WallReplyEditObject wallWallComment

// WallReplyRestoreFunc func
type WallReplyRestoreFunc func(WallReplyRestoreObject, int)

// WallReplyRestoreObject struct
type WallReplyRestoreObject wallWallComment

// WallReplyDeleteFunc func
type WallReplyDeleteFunc func(WallReplyDeleteObject, int)

// WallReplyDeleteObject struct
type WallReplyDeleteObject struct {
	OwnerID   int `json:"owner_id"`
	ID        int `json:"id"`
	DeleterID int `json:"deleter_id"`
	PostID    int `json:"post_id"`
}

// BoardPostNewFunc func
type BoardPostNewFunc func(BoardPostNewObject, int)

// BoardPostNewObject struct
type BoardPostNewObject boardTopicComment

// BoardPostEditFunc func
type BoardPostEditFunc func(BoardPostEditObject, int)

// BoardPostEditObject struct
type BoardPostEditObject boardTopicComment

// BoardPostRestoreFunc func
type BoardPostRestoreFunc func(BoardPostRestoreObject, int)

// BoardPostRestoreObject struct
type BoardPostRestoreObject boardTopicComment

// BoardPostDeleteFunc func
type BoardPostDeleteFunc func(BoardPostDeleteObject, int)

// BoardPostDeleteObject struct
type BoardPostDeleteObject struct {
	TopicOwnerID int `json:"topic_owner_id"`
	TopicID      int `json:"topic_id"`
	ID           int `json:"id"`
}

// MarketCommentNewFunc func
type MarketCommentNewFunc func(MarketCommentNewObject, int)

// MarketCommentNewObject struct
type MarketCommentNewObject wallWallComment

// MarketCommentEditFunc func
type MarketCommentEditFunc func(MarketCommentEditObject, int)

// MarketCommentEditObject struct
type MarketCommentEditObject wallWallComment

// MarketCommentRestoreFunc func
type MarketCommentRestoreFunc func(MarketCommentRestoreObject, int)

// MarketCommentRestoreObject struct
type MarketCommentRestoreObject wallWallComment

// MarketCommentDeleteFunc func
type MarketCommentDeleteFunc func(MarketCommentDeleteObject, int)

// MarketCommentDeleteObject struct
type MarketCommentDeleteObject struct {
	OwnerID   int `json:"owner_id"`
	ID        int `json:"id"`
	UserID    int `json:"user_id"`
	DeleterID int `json:"deleter_id"`
	ItemID    int `json:"item_id"`
}

// GroupLeaveFunc func
type GroupLeaveFunc func(GroupLeaveObject, int)

// GroupLeaveObject struct
type GroupLeaveObject struct {
	UserID int `json:"user_id"`
	Self   int `json:"self"`
}

// GroupJoinFunc func
type GroupJoinFunc func(GroupJoinObject, int)

// GroupJoinObject struct
type GroupJoinObject struct {
	UserID   int    `json:"user_id"`
	JoinType string `json:"join_type"`
}

// UserBlockFunc func
type UserBlockFunc func(UserBlockObject, int)

// UserBlockObject struct
type UserBlockObject struct {
	AdminID     int    `json:"admin_id"`
	UserID      int    `json:"user_id"`
	UnblockDate int    `json:"unblock_date"`
	Reason      int    `json:"reason"`
	Comment     string `json:"comment"`
}

// UserUnblockFunc func
type UserUnblockFunc func(UserUnblockObject, int)

// UserUnblockObject struct
type UserUnblockObject struct {
	AdminID   int `json:"admin_id"`
	UserID    int `json:"user_id"`
	ByEndDate int `json:"by_end_date"`
}

// PollVoteNewFunc func
type PollVoteNewFunc func(PollVoteNewObject, int)

// PollVoteNewObject struct
// BUG(VK): при голосовании за несколько вариантов, возвращается только один
type PollVoteNewObject struct {
	OwnerID  int `json:"owner_id"`
	PollID   int `json:"poll_id"`
	OptionID int `json:"option_id"`
	UserID   int `json:"user_id"`
}

// GroupOfficersEditFunc func
type GroupOfficersEditFunc func(GroupOfficersEditObject, int)

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

// GroupChangeSettingsFunc func
type GroupChangeSettingsFunc func(GroupChangeSettingsObject, int)

// GroupChangeSettingsObject struct
// спасибо vk.com/eee
// BUG(VK): Phone https://vk.com/bugtracker?act=show&id=64240
// BUG(VK): Email https://vk.com/bugtracker?act=show&id=86650
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

// GroupChangePhotoFunc func
type GroupChangePhotoFunc func(GroupChangePhotoObject, int)

// GroupChangePhotoObject struct
type GroupChangePhotoObject struct {
	UserID int         `json:"user_id"`
	Photo  PhotosPhoto `json:"photo"`
}

// VkpayTransactionFunc func
type VkpayTransactionFunc func(VkpayTransactionObject, int)

// VkpayTransactionObject struct
type VkpayTransactionObject struct {
	FromID      int    `json:"from_id"`
	Amount      string `json:"amount"` // FIXME string or int?
	Description string `json:"description"`
	Date        string `json:"date"` // FIXME string or int?
}
