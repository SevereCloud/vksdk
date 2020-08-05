package events // import "github.com/SevereCloud/vksdk/events"

import (
	"context"

	"github.com/SevereCloud/vksdk/object"
)

// MessageNewFunc func.
type MessageNewFunc func(context.Context, MessageNewObject)

// MessageNewObject struct.
type MessageNewObject struct {
	Message    object.MessagesMessage `json:"message"`
	ClientInfo object.ClientInfo      `json:"client_info"`
}

// MessageReplyFunc func.
type MessageReplyFunc func(context.Context, MessageReplyObject)

// MessageReplyObject struct.
type MessageReplyObject object.MessagesMessage

// MessageEditFunc func.
type MessageEditFunc func(context.Context, MessageEditObject)

// MessageEditObject struct.
type MessageEditObject object.MessagesMessage

// MessageAllowFunc func.
type MessageAllowFunc func(context.Context, MessageAllowObject)

// MessageAllowObject struct.
type MessageAllowObject struct {
	UserID int    `json:"user_id"`
	Key    string `json:"key"`
}

// MessageDenyFunc func.
type MessageDenyFunc func(context.Context, MessageDenyObject)

// MessageDenyObject struct.
type MessageDenyObject struct {
	UserID int `json:"user_id"`
}

// MessageTypingStateFunc func.
type MessageTypingStateFunc func(context.Context, MessageTypingStateObject)

// MessageTypingStateObject struct.
type MessageTypingStateObject struct {
	State  string `json:"state"`
	FromID int    `json:"from_id"`
	ToID   int    `json:"to_id"`
}

// MessageEventFunc func.
type MessageEventFunc func(context.Context, MessageEventObject)

// MessageEventObject struct.
type MessageEventObject struct {
	UserID                int    `json:"user_id"`
	PeerID                int    `json:"peer_id"`
	EventID               string `json:"event_id"`
	Payload               string `json:"payload"`
	ConversationMessageID int    `json:"conversation_message_id"`
}

// PhotoNewFunc func.
type PhotoNewFunc func(context.Context, PhotoNewObject)

// PhotoNewObject struct.
type PhotoNewObject object.PhotosPhoto

// PhotoCommentNewFunc func.
type PhotoCommentNewFunc func(context.Context, PhotoCommentNewObject)

// PhotoCommentNewObject struct.
type PhotoCommentNewObject object.WallWallComment

// PhotoCommentEditFunc func.
type PhotoCommentEditFunc func(context.Context, PhotoCommentEditObject)

// PhotoCommentEditObject struct.
type PhotoCommentEditObject object.WallWallComment

// PhotoCommentRestoreFunc func.
type PhotoCommentRestoreFunc func(context.Context, PhotoCommentRestoreObject)

// PhotoCommentRestoreObject struct.
type PhotoCommentRestoreObject object.WallWallComment

// PhotoCommentDeleteFunc func.
type PhotoCommentDeleteFunc func(context.Context, PhotoCommentDeleteObject)

// PhotoCommentDeleteObject struct.
type PhotoCommentDeleteObject struct {
	OwnerID   int `json:"owner_id"`
	ID        int `json:"id"`
	UserID    int `json:"user_id"`
	DeleterID int `json:"deleter_id"`
	PhotoID   int `json:"photo_id"`
}

// AudioNewFunc func.
type AudioNewFunc func(context.Context, AudioNewObject)

// AudioNewObject struct.
type AudioNewObject object.AudioAudio

// VideoNewFunc func.
type VideoNewFunc func(context.Context, VideoNewObject)

// VideoNewObject struct.
type VideoNewObject object.VideoVideo

// VideoCommentNewFunc func.
type VideoCommentNewFunc func(context.Context, VideoCommentNewObject)

// VideoCommentNewObject struct.
type VideoCommentNewObject object.WallWallComment

// VideoCommentEditFunc func.
type VideoCommentEditFunc func(context.Context, VideoCommentEditObject)

// VideoCommentEditObject struct.
type VideoCommentEditObject object.WallWallComment

// VideoCommentRestoreFunc func.
type VideoCommentRestoreFunc func(context.Context, VideoCommentRestoreObject)

// VideoCommentRestoreObject struct.
type VideoCommentRestoreObject object.WallWallComment

// VideoCommentDeleteFunc func.
type VideoCommentDeleteFunc func(context.Context, VideoCommentDeleteObject)

// VideoCommentDeleteObject struct.
type VideoCommentDeleteObject struct {
	OwnerID   int `json:"owner_id"`
	ID        int `json:"id"`
	UserID    int `json:"user_id"`
	DeleterID int `json:"deleter_id"`
	VideoID   int `json:"video_id"`
}

// WallPostNewFunc func.
type WallPostNewFunc func(context.Context, WallPostNewObject)

// WallPostNewObject struct.
type WallPostNewObject object.WallWallpost

// WallRepostFunc func.
type WallRepostFunc func(context.Context, WallRepostObject)

// WallRepostObject struct.
type WallRepostObject object.WallWallpost

// WallReplyNewFunc func.
type WallReplyNewFunc func(context.Context, WallReplyNewObject)

// WallReplyNewObject struct.
type WallReplyNewObject object.WallWallComment

// WallReplyEditFunc func.
type WallReplyEditFunc func(context.Context, WallReplyEditObject)

// WallReplyEditObject struct.
type WallReplyEditObject object.WallWallComment

// WallReplyRestoreFunc func.
type WallReplyRestoreFunc func(context.Context, WallReplyRestoreObject)

// WallReplyRestoreObject struct.
type WallReplyRestoreObject object.WallWallComment

// WallReplyDeleteFunc func.
type WallReplyDeleteFunc func(context.Context, WallReplyDeleteObject)

// WallReplyDeleteObject struct.
type WallReplyDeleteObject struct {
	OwnerID   int `json:"owner_id"`
	ID        int `json:"id"`
	DeleterID int `json:"deleter_id"`
	PostID    int `json:"post_id"`
}

// BoardPostNewFunc func.
type BoardPostNewFunc func(context.Context, BoardPostNewObject)

// BoardPostNewObject struct.
type BoardPostNewObject object.BoardTopicComment

// BoardPostEditFunc func.
type BoardPostEditFunc func(context.Context, BoardPostEditObject)

// BoardPostEditObject struct.
type BoardPostEditObject object.BoardTopicComment

// BoardPostRestoreFunc func.
type BoardPostRestoreFunc func(context.Context, BoardPostRestoreObject)

// BoardPostRestoreObject struct.
type BoardPostRestoreObject object.BoardTopicComment

// BoardPostDeleteFunc func.
type BoardPostDeleteFunc func(context.Context, BoardPostDeleteObject)

// BoardPostDeleteObject struct.
type BoardPostDeleteObject struct {
	TopicOwnerID int `json:"topic_owner_id"`
	TopicID      int `json:"topic_id"`
	ID           int `json:"id"`
}

// MarketCommentNewFunc func.
type MarketCommentNewFunc func(context.Context, MarketCommentNewObject)

// MarketCommentNewObject struct.
type MarketCommentNewObject object.WallWallComment

// MarketCommentEditFunc func.
type MarketCommentEditFunc func(context.Context, MarketCommentEditObject)

// MarketCommentEditObject struct.
type MarketCommentEditObject object.WallWallComment

// MarketCommentRestoreFunc func.
type MarketCommentRestoreFunc func(context.Context, MarketCommentRestoreObject)

// MarketCommentRestoreObject struct.
type MarketCommentRestoreObject object.WallWallComment

// MarketCommentDeleteFunc func.
type MarketCommentDeleteFunc func(context.Context, MarketCommentDeleteObject)

// MarketCommentDeleteObject struct.
type MarketCommentDeleteObject struct {
	OwnerID   int `json:"owner_id"`
	ID        int `json:"id"`
	UserID    int `json:"user_id"`
	DeleterID int `json:"deleter_id"`
	ItemID    int `json:"item_id"`
}

// MarketOrderNewFunc func.
type MarketOrderNewFunc func(context.Context, MarketOrderNewObject)

// MarketOrderNewObject struct.
type MarketOrderNewObject object.MarketOrder

// MarketOrderEditFunc func.
type MarketOrderEditFunc func(context.Context, MarketOrderEditObject)

// MarketOrderEditObject struct.
type MarketOrderEditObject object.MarketOrder

// GroupLeaveFunc func.
type GroupLeaveFunc func(context.Context, GroupLeaveObject)

// GroupLeaveObject struct.
type GroupLeaveObject struct {
	UserID int                `json:"user_id"`
	Self   object.BaseBoolInt `json:"self"`
}

// GroupJoinFunc func.
type GroupJoinFunc func(context.Context, GroupJoinObject)

// GroupJoinObject struct.
type GroupJoinObject struct {
	UserID   int    `json:"user_id"`
	JoinType string `json:"join_type"`
}

// UserBlockFunc func.
type UserBlockFunc func(context.Context, UserBlockObject)

// UserBlockObject struct.
type UserBlockObject struct {
	AdminID     int    `json:"admin_id"`
	UserID      int    `json:"user_id"`
	UnblockDate int    `json:"unblock_date"`
	Reason      int    `json:"reason"`
	Comment     string `json:"comment"`
}

// UserUnblockFunc func.
type UserUnblockFunc func(context.Context, UserUnblockObject)

// UserUnblockObject struct.
type UserUnblockObject struct {
	AdminID   int `json:"admin_id"`
	UserID    int `json:"user_id"`
	ByEndDate int `json:"by_end_date"`
}

// PollVoteNewFunc func.
type PollVoteNewFunc func(context.Context, PollVoteNewObject)

// PollVoteNewObject struct.
//
// BUG(VK): при голосовании за несколько вариантов, возвращается только один.
type PollVoteNewObject struct {
	OwnerID  int `json:"owner_id"`
	PollID   int `json:"poll_id"`
	OptionID int `json:"option_id"`
	UserID   int `json:"user_id"`
}

// GroupOfficersEditFunc func.
type GroupOfficersEditFunc func(context.Context, GroupOfficersEditObject)

// GroupOfficersEditObject struct.
type GroupOfficersEditObject struct {
	AdminID  int `json:"admin_id"`
	UserID   int `json:"user_id"`
	LevelOld int `json:"level_old"`
	LevelNew int `json:"level_new"`
}

// Changes struct.
type Changes struct {
	OldValue string `json:"old_value"`
	NewValue string `json:"new_value"`
}

// ChangesInt struct.
type ChangesInt struct {
	OldValue int `json:"old_value"`
	NewValue int `json:"new_value"`
}

// GroupChangeSettingsFunc func.
type GroupChangeSettingsFunc func(context.Context, GroupChangeSettingsObject)

// GroupChangeSettingsObject struct
//
// BUG(VK): Phone https://vk.com/bugtracker?act=show&id=64240
//
// BUG(VK): Email https://vk.com/bugtracker?act=show&id=86650
type GroupChangeSettingsObject struct {
	UserID  int `json:"user_id"`
	Changes struct {
		Title                 Changes    `json:"title"`
		Description           Changes    `json:"description"`
		Access                ChangesInt `json:"access"`
		ScreenName            Changes    `json:"screen_name"`
		PublicCategory        ChangesInt `json:"public_category"`
		PublicSubcategory     ChangesInt `json:"public_subcategory"`
		AgeLimits             ChangesInt `json:"age_limits"`
		Website               Changes    `json:"website"`
		StatusDefault         Changes    `json:"status_default"`
		Wall                  ChangesInt `json:"wall"`                    // на основе ответа
		Replies               ChangesInt `json:"replies"`                 // на основе ответа
		Topics                ChangesInt `json:"topics"`                  // на основе ответа
		Audio                 ChangesInt `json:"audio"`                   // на основе ответа
		Photos                ChangesInt `json:"photos"`                  // на основе ответа
		Video                 ChangesInt `json:"video"`                   // на основе ответа
		Market                ChangesInt `json:"market"`                  // на основе ответа
		Docs                  ChangesInt `json:"docs"`                    // на основе ответа
		Messages              ChangesInt `json:"messages"`                // на основе ответа
		EventGroupID          ChangesInt `json:"event_group_id"`          // на основе ответа
		Links                 Changes    `json:"links"`                   // на основе ответа
		Email                 Changes    `json:"email"`                   // на основе ответа
		EventStartDate        ChangesInt `json:"event_start_date::"`      // на основе ответа
		EventFinishDate       ChangesInt `json:"event_finish_date:"`      // на основе ответа
		Subject               Changes    `json:"subject"`                 // на основе ответа
		MarketWiki            Changes    `json:"market_wiki"`             // на основе ответа
		DisableMarketComments ChangesInt `json:"disable_market_comments"` // на основе ответа
		Phone                 ChangesInt `json:"phone"`                   // на основе ответа
		CountryID             ChangesInt `json:"country_id"`              // на основе ответа
		CityID                ChangesInt `json:"city_id"`                 // на основе ответа
	} `json:"Changes"`
}

// GroupChangePhotoFunc func.
type GroupChangePhotoFunc func(context.Context, GroupChangePhotoObject)

// GroupChangePhotoObject struct.
type GroupChangePhotoObject struct {
	UserID int                `json:"user_id"`
	Photo  object.PhotosPhoto `json:"photo"`
}

// VkpayTransactionFunc func.
type VkpayTransactionFunc func(context.Context, VkpayTransactionObject)

// VkpayTransactionObject struct.
type VkpayTransactionObject struct {
	FromID      int    `json:"from_id"`
	Amount      int    `json:"amount"`
	Description string `json:"description"`
	Date        int    `json:"date"`
}

// LeadFormsNewFunc func.
type LeadFormsNewFunc func(context.Context, LeadFormsNewObject)

// LeadFormsNewObject struct.
type LeadFormsNewObject struct {
	LeadID   int    `json:"lead_id"`
	GroupID  int    `json:"group_id"`
	UserID   int    `json:"user_id"`
	FormID   int    `json:"form_id"`
	FormName string `json:"form_name"`
	AdID     int    `json:"ad_id"`
	Answers  []struct {
		Key      string `json:"key"`
		Question string `json:"question"`
		Answer   string `json:"answer"`
	} `json:"answers"`
}

// AppPayloadFunc func.
type AppPayloadFunc func(context.Context, AppPayloadObject)

// AppPayloadObject struct.
type AppPayloadObject struct {
	UserID  int    `json:"user_id"`
	AppID   int    `json:"app_id"`
	Payload string `json:"payload"`
}

// MessageReadFunc func.
type MessageReadFunc func(context.Context, MessageReadObject)

// MessageReadObject struct.
type MessageReadObject struct {
	FromID        int `json:"from_id"`
	PeerID        int `json:"peer_id"`
	ReadMessageID int `json:"read_message_id"`
}

// LikeAddFunc func.
type LikeAddFunc func(context.Context, LikeAddObject)

// LikeAddObject struct.
type LikeAddObject struct {
	LikerID       int    `json:"liker_id"`
	ObjectType    string `json:"object_type"`
	ObjectOwnerID int    `json:"object_owner_id"`
	ObjectID      int    `json:"object_id"`
	ThreadReplyID int    `json:"thread_reply_id"`
	PostID        int    `json:"post_id"` // for comment
}

// LikeRemoveFunc func.
type LikeRemoveFunc func(context.Context, LikeRemoveObject)

// LikeRemoveObject struct.
type LikeRemoveObject struct {
	LikerID       int    `json:"liker_id"`
	ObjectType    string `json:"object_type"`
	ObjectOwnerID int    `json:"object_owner_id"`
	ObjectID      int    `json:"object_id"`
	ThreadReplyID int    `json:"thread_reply_id"`
	PostID        int    `json:"post_id"` // for comment
}
