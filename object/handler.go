package object // import "github.com/SevereCloud/vksdk/object"

import (
	"bytes"
	"encoding/json"
)

// MessageNewFunc func.
type MessageNewFunc func(MessageNewObject, int)

// MessageNewObject struct.
type MessageNewObject struct {
	Message    MessagesMessage `json:"message"`
	ClientInfo ClientInfo      `json:"client_info"`
}

// UnmarshalJSON need for support api version < 5.103.
//
// To unmarshal JSON into a value implementing the Unmarshaler interface,
// Unmarshal calls that value's UnmarshalJSON method.
// See more https://golang.org/pkg/encoding/json/#Unmarshal
func (obj *MessageNewObject) UnmarshalJSON(data []byte) (err error) {
	// The renamed type is necessary to avoid loops
	type renamedMessageNewObject MessageNewObject

	var renamedObj renamedMessageNewObject

	if bytes.Contains(data, []byte(`"message":`)) {
		// v >= 5.103
		err = json.Unmarshal(data, &renamedObj)
	} else {
		// Support v < 5.103
		err = json.Unmarshal(data, &renamedObj.Message)
	}

	obj.Message = renamedObj.Message
	obj.ClientInfo = renamedObj.ClientInfo

	return
}

// MessageReplyFunc func.
type MessageReplyFunc func(MessageReplyObject, int)

// MessageReplyObject struct.
type MessageReplyObject MessagesMessage

// MessageEditFunc func.
type MessageEditFunc func(MessageEditObject, int)

// MessageEditObject struct.
type MessageEditObject MessagesMessage

// MessageAllowFunc func.
type MessageAllowFunc func(MessageAllowObject, int)

// MessageAllowObject struct.
type MessageAllowObject struct {
	UserID int    `json:"user_id"`
	Key    string `json:"key"`
}

// MessageDenyFunc func.
type MessageDenyFunc func(MessageDenyObject, int)

// MessageDenyObject struct.
type MessageDenyObject struct {
	UserID int `json:"user_id"`
}

// MessageTypingStateFunc func.
type MessageTypingStateFunc func(MessageTypingStateObject, int)

// MessageTypingStateObject struct.
type MessageTypingStateObject struct {
	State  string `json:"state"`
	FromID int    `json:"from_id"`
	ToID   int    `json:"to_id"`
}

// PhotoNewFunc func.
type PhotoNewFunc func(PhotoNewObject, int)

// PhotoNewObject struct.
type PhotoNewObject PhotosPhoto

// PhotoCommentNewFunc func.
type PhotoCommentNewFunc func(PhotoCommentNewObject, int)

// PhotoCommentNewObject struct.
type PhotoCommentNewObject WallWallComment

// PhotoCommentEditFunc func.
type PhotoCommentEditFunc func(PhotoCommentEditObject, int)

// PhotoCommentEditObject struct.
type PhotoCommentEditObject WallWallComment

// PhotoCommentRestoreFunc func.
type PhotoCommentRestoreFunc func(PhotoCommentRestoreObject, int)

// PhotoCommentRestoreObject struct.
type PhotoCommentRestoreObject WallWallComment

// PhotoCommentDeleteFunc func.
type PhotoCommentDeleteFunc func(PhotoCommentDeleteObject, int)

// PhotoCommentDeleteObject struct.
type PhotoCommentDeleteObject struct {
	OwnerID   int `json:"owner_id"`
	ID        int `json:"id"`
	UserID    int `json:"user_id"`
	DeleterID int `json:"deleter_id"`
	PhotoID   int `json:"photo_id"`
}

// AudioNewFunc func.
type AudioNewFunc func(AudioNewObject, int)

// AudioNewObject struct.
type AudioNewObject AudioAudioFull

// VideoNewFunc func.
type VideoNewFunc func(VideoNewObject, int)

// VideoNewObject struct.
type VideoNewObject VideoVideo

// VideoCommentNewFunc func.
type VideoCommentNewFunc func(VideoCommentNewObject, int)

// VideoCommentNewObject struct.
type VideoCommentNewObject WallWallComment

// VideoCommentEditFunc func.
type VideoCommentEditFunc func(VideoCommentEditObject, int)

// VideoCommentEditObject struct.
type VideoCommentEditObject WallWallComment

// VideoCommentRestoreFunc func.
type VideoCommentRestoreFunc func(VideoCommentRestoreObject, int)

// VideoCommentRestoreObject struct.
type VideoCommentRestoreObject WallWallComment

// VideoCommentDeleteFunc func.
type VideoCommentDeleteFunc func(VideoCommentDeleteObject, int)

// VideoCommentDeleteObject struct.
type VideoCommentDeleteObject struct {
	OwnerID   int `json:"owner_id"`
	ID        int `json:"id"`
	UserID    int `json:"user_id"`
	DeleterID int `json:"deleter_id"`
	VideoID   int `json:"video_id"`
}

// WallPostNewFunc func.
type WallPostNewFunc func(WallPostNewObject, int)

// WallPostNewObject struct.
type WallPostNewObject WallWallpost

// WallRepostFunc func.
type WallRepostFunc func(WallRepostObject, int)

// WallRepostObject struct.
type WallRepostObject WallWallpost

// WallReplyNewFunc func.
type WallReplyNewFunc func(WallReplyNewObject, int)

// WallReplyNewObject struct.
type WallReplyNewObject WallWallComment

// WallReplyEditFunc func.
type WallReplyEditFunc func(WallReplyEditObject, int)

// WallReplyEditObject struct.
type WallReplyEditObject WallWallComment

// WallReplyRestoreFunc func.
type WallReplyRestoreFunc func(WallReplyRestoreObject, int)

// WallReplyRestoreObject struct.
type WallReplyRestoreObject WallWallComment

// WallReplyDeleteFunc func.
type WallReplyDeleteFunc func(WallReplyDeleteObject, int)

// WallReplyDeleteObject struct.
type WallReplyDeleteObject struct {
	OwnerID   int `json:"owner_id"`
	ID        int `json:"id"`
	DeleterID int `json:"deleter_id"`
	PostID    int `json:"post_id"`
}

// BoardPostNewFunc func.
type BoardPostNewFunc func(BoardPostNewObject, int)

// BoardPostNewObject struct.
type BoardPostNewObject BoardTopicComment

// BoardPostEditFunc func.
type BoardPostEditFunc func(BoardPostEditObject, int)

// BoardPostEditObject struct.
type BoardPostEditObject BoardTopicComment

// BoardPostRestoreFunc func.
type BoardPostRestoreFunc func(BoardPostRestoreObject, int)

// BoardPostRestoreObject struct.
type BoardPostRestoreObject BoardTopicComment

// BoardPostDeleteFunc func.
type BoardPostDeleteFunc func(BoardPostDeleteObject, int)

// BoardPostDeleteObject struct.
type BoardPostDeleteObject struct {
	TopicOwnerID int `json:"topic_owner_id"`
	TopicID      int `json:"topic_id"`
	ID           int `json:"id"`
}

// MarketCommentNewFunc func.
type MarketCommentNewFunc func(MarketCommentNewObject, int)

// MarketCommentNewObject struct.
type MarketCommentNewObject WallWallComment

// MarketCommentEditFunc func.
type MarketCommentEditFunc func(MarketCommentEditObject, int)

// MarketCommentEditObject struct.
type MarketCommentEditObject WallWallComment

// MarketCommentRestoreFunc func.
type MarketCommentRestoreFunc func(MarketCommentRestoreObject, int)

// MarketCommentRestoreObject struct.
type MarketCommentRestoreObject WallWallComment

// MarketCommentDeleteFunc func.
type MarketCommentDeleteFunc func(MarketCommentDeleteObject, int)

// MarketCommentDeleteObject struct.
type MarketCommentDeleteObject struct {
	OwnerID   int `json:"owner_id"`
	ID        int `json:"id"`
	UserID    int `json:"user_id"`
	DeleterID int `json:"deleter_id"`
	ItemID    int `json:"item_id"`
}

// MarketOrderNewFunc func.
type MarketOrderNewFunc func(MarketOrderNewObject, int)

// MarketOrderNewObject struct.
type MarketOrderNewObject MarketOrder

// MarketOrderEditFunc func.
type MarketOrderEditFunc func(MarketOrderEditObject, int)

// MarketOrderEditObject struct.
type MarketOrderEditObject MarketOrder

// GroupLeaveFunc func.
type GroupLeaveFunc func(GroupLeaveObject, int)

// GroupLeaveObject struct.
type GroupLeaveObject struct {
	UserID int         `json:"user_id"`
	Self   BaseBoolInt `json:"self"`
}

// GroupJoinFunc func.
type GroupJoinFunc func(GroupJoinObject, int)

// GroupJoinObject struct.
type GroupJoinObject struct {
	UserID   int    `json:"user_id"`
	JoinType string `json:"join_type"`
}

// UserBlockFunc func.
type UserBlockFunc func(UserBlockObject, int)

// UserBlockObject struct.
type UserBlockObject struct {
	AdminID     int    `json:"admin_id"`
	UserID      int    `json:"user_id"`
	UnblockDate int    `json:"unblock_date"`
	Reason      int    `json:"reason"`
	Comment     string `json:"comment"`
}

// UserUnblockFunc func.
type UserUnblockFunc func(UserUnblockObject, int)

// UserUnblockObject struct.
type UserUnblockObject struct {
	AdminID   int `json:"admin_id"`
	UserID    int `json:"user_id"`
	ByEndDate int `json:"by_end_date"`
}

// PollVoteNewFunc func.
type PollVoteNewFunc func(PollVoteNewObject, int)

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
type GroupOfficersEditFunc func(GroupOfficersEditObject, int)

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
type GroupChangeSettingsFunc func(GroupChangeSettingsObject, int)

// GroupChangeSettingsObject struct
// спасибо vk.com/eee
// BUG(VK): Phone https://vk.com/bugtracker?act=show&id=64240
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
type GroupChangePhotoFunc func(GroupChangePhotoObject, int)

// GroupChangePhotoObject struct.
type GroupChangePhotoObject struct {
	UserID int         `json:"user_id"`
	Photo  PhotosPhoto `json:"photo"`
}

// VkpayTransactionFunc func.
type VkpayTransactionFunc func(VkpayTransactionObject, int)

// VkpayTransactionObject struct.
type VkpayTransactionObject struct {
	FromID      int    `json:"from_id"`
	Amount      int    `json:"amount"`
	Description string `json:"description"`
	Date        int    `json:"date"`
}

// LeadFormsNewFunc func.
type LeadFormsNewFunc func(LeadFormsNewObject, int)

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
type AppPayloadFunc func(AppPayloadObject, int)

// AppPayloadObject struct.
type AppPayloadObject struct {
	UserID  int    `json:"user_id"`
	AppID   int    `json:"app_id"`
	Payload string `json:"payload"`
}

// MessageReadFunc func.
type MessageReadFunc func(MessageReadObject, int)

// MessageReadObject struct.
type MessageReadObject struct {
	FromID        int `json:"from_id"`
	PeerID        int `json:"peer_id"`
	ReadMessageID int `json:"read_message_id"`
}

// LikeAddFunc func.
type LikeAddFunc func(LikeAddObject, int)

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
type LikeRemoveFunc func(LikeRemoveObject, int)

// LikeRemoveObject struct.
type LikeRemoveObject struct {
	LikerID       int    `json:"liker_id"`
	ObjectType    string `json:"object_type"`
	ObjectOwnerID int    `json:"object_owner_id"`
	ObjectID      int    `json:"object_id"`
	ThreadReplyID int    `json:"thread_reply_id"`
	PostID        int    `json:"post_id"` // for comment
}
