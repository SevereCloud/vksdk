package wrapper

import (
	"fmt"
	"time"

	"github.com/SevereCloud/vksdk/object"
)

// MessageFlagsSet struct for event with code 2
//
// Install message flags
type MessageFlagsSet struct {
	MessageID int
	Flags     MessageFlag
	PeerID    int
}

// Parse func for MessageFlagsSet
func (result *MessageFlagsSet) Parse(i []interface{}) error {
	if len(i) < 4 {
		return fmt.Errorf(errFmtTooShortArray, "MessageFlagsSet", 4, len(i))
	}

	if v, ok := i[1].(float64); ok {
		result.MessageID = int(v)
	}

	if v, ok := i[2].(float64); ok {
		result.Flags = MessageFlag(int(v))
	}

	if v, ok := i[3].(float64); ok {
		result.MessageID = int(v)
	}

	return nil
}

// MessageFlagsReset struct for event with code 3
//
// Reset message flags
type MessageFlagsReset struct {
	Message
}

// Parse func for MessageFlagsReset
func (result *MessageFlagsReset) Parse(i []interface{}) error {
	return result.parseMessage("MessageFlagsReset", i)
}

// NewMessage struct for event with code 4
//
// Add a new message
type NewMessage struct {
	Message
}

// Parse func for NewMessage
func (result *NewMessage) Parse(i []interface{}) error {
	return result.parseMessage("NewMessage", i)
}

// EditMessage struct for event with code 5
//
// Edit message.
type EditMessage struct {
	Message
}

// Parse func for EditMessage
func (result *EditMessage) Parse(i []interface{}) error {
	return result.parseMessage("EditMessage", i)
}

// ReadInMessages struct for event with code 6
//
// Read all incoming messages received before message with MessageID
type ReadInMessages struct {
	PeerID      int
	MessageID   int
	UnreadCount int
}

// Parse func for ReadInMessages
func (result *ReadInMessages) Parse(i []interface{}) error {
	if len(i) < 4 {
		return fmt.Errorf(errFmtTooShortArray, "ReadInMessages", 4, len(i))
	}

	if v, ok := i[1].(float64); ok {
		result.PeerID = int(v)
	}

	if v, ok := i[2].(float64); ok {
		result.MessageID = int(v)
	}

	if v, ok := i[3].(float64); ok {
		result.UnreadCount = int(v)
	}

	return nil
}

// ReadOutMessages struct for event with code 7
//
// Read all outgoing messages sent before message with MessageID
type ReadOutMessages struct {
	PeerID    int
	MessageID int
	// Note: Your unread messages count
	UnreadCount int
}

// Parse func for ReadOutMessages
func (result *ReadOutMessages) Parse(i []interface{}) error {
	if len(i) < 4 {
		return fmt.Errorf(errFmtTooShortArray, "ReadOutMessages", 4, len(i))
	}

	if v, ok := i[1].(float64); ok {
		result.PeerID = int(v)
	}

	if v, ok := i[2].(float64); ok {
		result.MessageID = int(v)
	}

	if v, ok := i[3].(float64); ok {
		result.UnreadCount = int(v)
	}

	return nil
}

// FriendBecameOnline struct for event with code 8
//
// A friend UserID is online
type FriendBecameOnline struct {
	UserID    int
	Platform  object.Platform
	Timestamp time.Time
	AppID     int
}

// Parse func for FriendBecameOnline
func (result *FriendBecameOnline) Parse(i []interface{}) error {
	if len(i) < 5 {
		return fmt.Errorf(errFmtTooShortArray, "FriendBecameOnline", 5, len(i))
	}

	if v, ok := i[1].(float64); ok {
		result.UserID = int(v)
	}

	if v, ok := i[2].(float64); ok {
		result.Platform = object.Platform(int(v))
	}

	if v, ok := i[3].(float64); ok {
		result.Timestamp = time.Unix(int64(v), 0)
	}

	if v, ok := i[4].(float64); ok {
		result.AppID = int(v)
	}

	return nil
}

// FriendBecameOffline struct for event with code 9
//
// A friend UserID is offline
type FriendBecameOffline struct {
	UserID int
	// true if offline is due to timing out, otherwise false
	ByTimeout bool
	Timestamp time.Time
	AppID     int
}

// Parse func for FriendBecameOffline
func (result *FriendBecameOffline) Parse(i []interface{}) error {
	if len(i) < 5 {
		return fmt.Errorf(errFmtTooShortArray, "FriendBecameOffline", 5, len(i))
	}

	if v, ok := i[1].(float64); ok {
		result.UserID = int(v)
	}

	if v, ok := i[3].(float64); ok {
		result.ByTimeout = int(v) > 0
	}

	if v, ok := i[3].(float64); ok {
		result.Timestamp = time.Unix(int64(v), 0)
	}

	if v, ok := i[4].(float64); ok {
		result.AppID = int(v)
	}

	return nil
}

// MentionViewed struct for event with code 10
//
// Current user mention marked as viewed
type MentionViewed struct {
	PeerID int
	// TODO specify field.
	// it is not a msg_id and not a DialogFlag/MessageFlag
	Flags MentionFlags
}

// Parse func for MentionViewed
func (result *MentionViewed) Parse(i []interface{}) error {
	if len(i) < 3 {
		return fmt.Errorf(errFmtTooShortArray, "MentionViewed", 3, len(i))
	}

	if v, ok := i[1].(float64); ok {
		result.PeerID = int(v)
	}

	if v, ok := i[2].(float64); ok {
		result.Flags = MentionFlags(int(v))
	}

	return nil
}

// UserMentioned struct for event with code 12
//
// Current user was mentioned in chat PeerID
type UserMentioned struct {
	PeerID int
	// TODO specify field.
	// it is not a msg_id and not a DialogFlag/MessageFlag
	Flags MentionFlags
}

// Parse func for UserMentioned
func (result *UserMentioned) Parse(i []interface{}) error {
	if len(i) < 3 {
		return fmt.Errorf(errFmtTooShortArray, "UserMentioned", 3, len(i))
	}

	if v, ok := i[1].(float64); ok {
		result.PeerID = int(v)
	}

	if v, ok := i[2].(float64); ok {
		result.Flags = MentionFlags(int(v))
	}

	return nil
}

// DeleteAllMessage struct for event with code 13
//
// Delete all message in chat PeerID
type DeleteAllMessage struct {
	PeerID        int
	LastMessageID int
}

// Parse func for DeleteAllMessage
func (result *DeleteAllMessage) Parse(i []interface{}) error {
	if len(i) < 3 {
		return fmt.Errorf(errFmtTooShortArray, "DeleteAllMessage", 3, len(i))
	}

	if v, ok := i[1].(float64); ok {
		result.PeerID = int(v)
	}

	if v, ok := i[2].(float64); ok {
		result.LastMessageID = int(v)
	}

	return nil
}

// LinkSnippetLoaded struct for event with code 18
//
// Link snippet loaded.
type LinkSnippetLoaded struct {
	Message
}

// Parse func for LinkSnippetLoaded
func (result *LinkSnippetLoaded) Parse(i []interface{}) error {
	return result.parseMessage("LinkSnippetLoaded", i)
}

// MessageCacheReset struct for event with code 19
//
// Message had changed
type MessageCacheReset struct {
	MessageID int
}

// Parse func for MessageCacheReset
func (result *MessageCacheReset) Parse(i []interface{}) error {
	if len(i) < 2 {
		return fmt.Errorf(errFmtTooShortArray, "MessageCacheReset", 2, len(i))
	}

	if v, ok := i[1].(float64); ok {
		result.MessageID = int(v)
	}

	return nil
}

// ChatParamsChange struct for event with code 51
//
// One of the parameters (content, topic) of the conversation ChatID was changed.
type ChatParamsChange struct {
	ChatID int
}

// Parse func for ChatParamsChange
func (result *ChatParamsChange) Parse(i []interface{}) error {
	if len(i) < 2 {
		return fmt.Errorf(errFmtTooShortArray, "ChatParamsChange", 2, len(i))
	}

	if v, ok := i[1].(float64); ok {
		result.ChatID = int(v)
	}

	return nil
}

// ChatInfoChange struct for event with code 52
//
// Chat info change
type ChatInfoChange struct {
	TypeID TypeID
	PeerID int
	Extra  int
}

// Parse func for ChatInfoChange
func (result *ChatInfoChange) Parse(i []interface{}) error {
	if len(i) < 4 {
		return fmt.Errorf(errFmtTooShortArray, "ChatInfoChange", 4, len(i))
	}

	if v, ok := i[1].(float64); ok {
		result.TypeID = TypeID(int(v))
	}

	if v, ok := i[2].(float64); ok {
		result.PeerID = int(v)
	}

	if v, ok := i[3].(float64); ok {
		result.Extra = int(v)
	}

	return nil
}

// UsersTyping struct for event with code 63
type UsersTyping struct {
	PeerID       int
	FromIDs      []int
	FromIDsCount int
	Ts           time.Time
}

// Parse func for UsersTyping
func (result *UsersTyping) Parse(i []interface{}) error {
	if len(i) < 5 {
		return fmt.Errorf(errFmtTooShortArray, "UsersTyping", 5, len(i))
	}

	if v, ok := i[1].(float64); ok {
		result.PeerID = int(v)
	}

	v, err := interfaceToIDSlice(i[2])
	if err != nil {
		return err
	}

	result.FromIDs = v

	if v, ok := i[3].(float64); ok {
		result.FromIDsCount = int(v)
	}

	if v, ok := i[4].(float64); ok {
		result.Ts = time.Unix(int64(v), 0)
	}

	return nil
}

// UsersRecordingAudioMessage struct for event with code 64
type UsersRecordingAudioMessage struct {
	PeerID       int
	FromIDs      []int
	FromIDsCount int
	Ts           time.Time
}

// Parse func for UsersRecordingAudioMessage
func (result *UsersRecordingAudioMessage) Parse(i []interface{}) error {
	if len(i) < 5 {
		return fmt.Errorf(errFmtTooShortArray, "UsersRecordingAudioMessage", 5, len(i))
	}

	if v, ok := i[1].(float64); ok {
		result.PeerID = int(v)
	}

	v, err := interfaceToIDSlice(i[2])
	if err != nil {
		return err
	}

	result.FromIDs = v

	if v, ok := i[3].(float64); ok {
		result.FromIDsCount = int(v)
	}

	if v, ok := i[4].(float64); ok {
		result.Ts = time.Unix(int64(v), 0)
	}

	return nil
}

// CounterChange struct for event with code 80
type CounterChange struct {
	Count                  int
	CountWithNotifications int
}

// Parse func for CounterChange
func (result *CounterChange) Parse(i []interface{}) error {
	if len(i) < 3 {
		return fmt.Errorf(errFmtTooShortArray, "CounterChange", 3, len(i))
	}

	if v, ok := i[1].(float64); ok {
		result.Count = int(v)
	}

	if v, ok := i[2].(float64); ok {
		result.CountWithNotifications = int(v)
	}

	return nil
}

// FriendInvisibilityChange struct for event with code 81
type FriendInvisibilityChange struct {
	UserID    int
	State     bool
	Timestamp time.Time
}

// Parse func for FriendInvisibilityChange
func (result *FriendInvisibilityChange) Parse(i []interface{}) error {
	if len(i) < 4 {
		return fmt.Errorf(errFmtTooShortArray, "FriendInvisibilityChange", 4, len(i))
	}

	if v, ok := i[1].(float64); ok {
		result.UserID = int(v)
	}

	if v, ok := i[2].(float64); ok {
		result.State = int(v) > 0
	}

	if v, ok := i[3].(float64); ok {
		result.Timestamp = time.Unix(int64(v), 0)
	}

	return nil
}

// NotificationSettingsChange struct for event with code 114
type NotificationSettingsChange struct {
	PeerID        int
	Sound         bool
	DisabledUntil int
}

// ParseMode8 should be called if ExtendedEvents flag set
func (result *NotificationSettingsChange) ParseMode8(i []interface{}) error {
	if len(i) < 2 {
		return fmt.Errorf(errFmtTooShortArray, "NotificationSettingsChange", 2, len(i))
	}

	v, err := interfaceToStringIntMap(i[1])
	if err != nil {
		return err
	}

	result.PeerID = v["peer_id"]
	result.Sound = v["sound"] > 0
	result.DisabledUntil = v["disabled_until"]

	return nil
}

// Parse func for NotificationSettingsChange
func (result *NotificationSettingsChange) Parse(i []interface{}) error {
	if len(i) < 3 {
		return fmt.Errorf(errFmtTooShortArray, "NotificationSettingsChange", 3, len(i))
	}

	if v, ok := i[1].(float64); ok {
		result.Sound = int(v) > 0
	}

	if v, ok := i[2].(float64); ok {
		result.DisabledUntil = int(v)
	}

	return nil
}
