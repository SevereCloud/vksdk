package wrapper // import "github.com/SevereCloud/vksdk/longpoll-user/v3"

import (
	"fmt"
	"time"
)

// MessageFlagsChange struct for event with code 1
//
// Replace message flags
type MessageFlagsChange struct {
	MessageID int
	Flags     MessageFlag
	ExtraFields
}

// Parse func for MessageFlagsChange
func (result *MessageFlagsChange) Parse(i []interface{}) error {
	if len(i) < 3 {
		return fmt.Errorf(errFmtTooShortArray, "MessageFlagsChange", 3, len(i))
	}

	if v, ok := i[1].(float64); ok {
		result.MessageID = int(v)
	}

	if v, ok := i[2].(float64); ok {
		result.Flags = MessageFlag(int(v))
	}

	return result.ExtraFields.parseExtraFields(i)
}

// MessageFlagsSet struct for event with code 2
//
// Install message flags
type MessageFlagsSet struct {
	MessageID int
	Mask      MessageFlag
	ExtraFields
}

// Parse func for MessageFlagsSet
func (result *MessageFlagsSet) Parse(i []interface{}) error {
	if len(i) < 3 {
		return fmt.Errorf(errFmtTooShortArray, "MessageFlagsSet", 3, len(i))
	}

	if v, ok := i[1].(float64); ok {
		result.MessageID = int(v)
	}

	if v, ok := i[2].(float64); ok {
		result.Mask = MessageFlag(int(v))
	}

	return result.ExtraFields.parseExtraFields(i)
}

// MessageFlagsReset struct for event with code 3
//
// Reset message flags
type MessageFlagsReset struct {
	MessageID int
	Mask      MessageFlag
	ExtraFields
}

// Parse func for MessageFlagsReset
func (result *MessageFlagsReset) Parse(i []interface{}) error {
	if len(i) < 3 {
		return fmt.Errorf(errFmtTooShortArray, "MessageFlagsReset", 3, len(i))
	}

	if v, ok := i[1].(float64); ok {
		result.MessageID = int(v)
	}

	if v, ok := i[2].(float64); ok {
		result.Mask = MessageFlag(int(v))
	}

	return result.ExtraFields.parseExtraFields(i)
}

// NewMessage struct for event with code 4
//
// Add a new message
type NewMessage struct {
	MessageID int
	Flags     MessageFlag
	ExtraFields
}

// Parse func for NewMessage
func (result *NewMessage) Parse(i []interface{}) error {
	if len(i) < 3 {
		return fmt.Errorf(errFmtTooShortArray, "NewMessage", 3, len(i))
	}

	if v, ok := i[1].(float64); ok {
		result.MessageID = int(v)
	}

	if v, ok := i[2].(float64); ok {
		result.Flags = MessageFlag(int(v))
	}

	return result.ExtraFields.parseExtraFields(i)
}

// EditMessage struct for event with code 5
//
// Edit message.
type EditMessage struct {
	MessageID      int
	Flags          MessageFlag
	PeerID         int
	Timestamp      time.Time
	NewText        string
	AdditionalData AdditionalData
	Attachments    LongPollAttachments
}

// Parse func for EditMessage
func (result *EditMessage) Parse(i []interface{}) error {
	if len(i) < 6 {
		return fmt.Errorf(errFmtTooShortArray, "EditMessage", 6, len(i))
	}

	if v, ok := i[1].(float64); ok {
		result.MessageID = int(v)
	}

	if v, ok := i[2].(float64); ok {
		result.Flags = MessageFlag(int(v))
	}

	if v, ok := i[3].(float64); ok {
		result.PeerID = int(v)
	}

	if v, ok := i[4].(float64); ok {
		result.Timestamp = time.Unix(int64(v), 0)
	}

	if v, ok := i[5].(string); ok {
		result.NewText = v
	}

	if len(i) > 6 {
		if v, ok := i[6].(map[string]interface{}); ok {
			result.AdditionalData.parse(v)
		}
	}

	if len(i) > 7 {
		if v, ok := i[7].(map[string]interface{}); ok {
			result.Attachments = v
		}
	}

	return nil
}

// ReadInMessages struct for event with code 6
//
// Read all incoming messages received before message with LocalID
type ReadInMessages struct {
	PeerID  int
	LocalID int
}

// Parse func for ReadInMessages
func (result *ReadInMessages) Parse(i []interface{}) error {
	if len(i) < 3 {
		return fmt.Errorf(errFmtTooShortArray, "ReadInMessages", 3, len(i))
	}

	if v, ok := i[1].(float64); ok {
		result.PeerID = int(v)
	}

	if v, ok := i[2].(float64); ok {
		result.LocalID = int(v)
	}

	return nil
}

// ReadOutMessages struct for event with code 7
//
// Read all outgoing messages sent before message with LocalID
type ReadOutMessages struct {
	PeerID  int
	LocalID int
}

// Parse func for ReadOutMessages
func (result *ReadOutMessages) Parse(i []interface{}) error {
	if len(i) < 3 {
		return fmt.Errorf(errFmtTooShortArray, "ReadOutMessages", 3, len(i))
	}

	if v, ok := i[1].(float64); ok {
		result.PeerID = int(v)
	}

	if v, ok := i[2].(float64); ok {
		result.LocalID = int(v)
	}

	return nil
}

// FriendBecameOnline struct for event with code 8
//
// A friend UserID is online
type FriendBecameOnline struct {
	UserID    int
	Extra     int
	Timestamp time.Time
}

// Parse func for FriendBecameOnline
func (result *FriendBecameOnline) Parse(i []interface{}) error {
	if len(i) < 4 {
		return fmt.Errorf(errFmtTooShortArray, "FriendBecameOnline", 4, len(i))
	}

	if v, ok := i[1].(float64); ok {
		result.UserID = int(v)
	}

	if v, ok := i[2].(float64); ok {
		result.Extra = int(v)
	}

	if v, ok := i[3].(float64); ok {
		result.Timestamp = time.Unix(int64(v), 0)
	}

	return nil
}

// FriendBecameOffline struct for event with code 9
//
// A friend UserID is offline
type FriendBecameOffline struct {
	UserID    int
	Flags     int
	Timestamp time.Time
}

// Parse func for FriendBecameOffline
func (result *FriendBecameOffline) Parse(i []interface{}) error {
	if len(i) < 4 {
		return fmt.Errorf(errFmtTooShortArray, "FriendBecameOffline", 4, len(i))
	}

	if v, ok := i[1].(float64); ok {
		result.UserID = int(v)
	}

	if v, ok := i[2].(float64); ok {
		result.Flags = int(v)
	}

	if v, ok := i[3].(float64); ok {
		result.Timestamp = time.Unix(int64(v), 0)
	}

	return nil
}

// DialogFlagsReset struct for event with code 10
//
// Reset dialog flags
type DialogFlagsReset struct {
	PeerID int
	Mask   DialogFlag
}

// Parse func for DialogFlagsReset
func (result *DialogFlagsReset) Parse(i []interface{}) error {
	if len(i) < 3 {
		return fmt.Errorf(errFmtTooShortArray, "DialogFlagsReset", 3, len(i))
	}

	if v, ok := i[1].(float64); ok {
		result.PeerID = int(v)
	}

	if v, ok := i[2].(float64); ok {
		result.Mask = DialogFlag(int(v))
	}

	return nil
}

// DialogFlagsReplace struct for event with code 11
//
// Replace dialog flags
type DialogFlagsReplace struct {
	PeerID int
	Flags  DialogFlag
}

// Parse func for DialogFlagsReplace
func (result *DialogFlagsReplace) Parse(i []interface{}) error {
	if len(i) < 3 {
		return fmt.Errorf(errFmtTooShortArray, "DialogFlagsReplace", 3, len(i))
	}

	if v, ok := i[1].(float64); ok {
		result.PeerID = int(v)
	}

	if v, ok := i[2].(float64); ok {
		result.Flags = DialogFlag(int(v))
	}

	return nil
}

// DialogsFlagsSet struct for event with code 12
//
// Install dialog flags
type DialogsFlagsSet struct {
	PeerID int
	Mask   DialogFlag
}

// Parse func for DialogsFlagsSet
func (result *DialogsFlagsSet) Parse(i []interface{}) error {
	if len(i) < 3 {
		return fmt.Errorf(errFmtTooShortArray, "DialogsFlagsSet", 3, len(i))
	}

	if v, ok := i[1].(float64); ok {
		result.PeerID = int(v)
	}

	if v, ok := i[2].(float64); ok {
		result.Mask = DialogFlag(int(v))
	}

	return nil
}

// DeleteMessages struct for event with code 13
//
// Deletings all messages in PeerID dialogs with IDs up to LocalID
type DeleteMessages struct {
	PeerID  int
	LocalID int
}

// Parse func for DeleteMessages
func (result *DeleteMessages) Parse(i []interface{}) error {
	if len(i) < 3 {
		return fmt.Errorf(errFmtTooShortArray, "DeleteMessages", 3, len(i))
	}

	if v, ok := i[1].(float64); ok {
		result.PeerID = int(v)
	}

	if v, ok := i[2].(float64); ok {
		result.LocalID = int(v)
	}

	return nil
}

// RestoreDeletedMessages struct for event with code 14
//
// Restore message
type RestoreDeletedMessages struct {
	PeerID  int
	LocalID int
}

// Parse func for RestoreDeletedMessages
func (result *RestoreDeletedMessages) Parse(i []interface{}) error {
	if len(i) < 3 {
		return fmt.Errorf(errFmtTooShortArray, "RestoreDeletedMessages", 3, len(i))
	}

	if v, ok := i[1].(float64); ok {
		result.PeerID = int(v)
	}

	if v, ok := i[2].(float64); ok {
		result.LocalID = int(v)
	}

	return nil
}

// ChatParamsChange struct for event with code 51
//
// One of the parameters (content, topic) of the conversation ChatID was
// changed. Self — 1 or 0 (whether the change was caused by the user).
type ChatParamsChange struct {
	ChatID int
	Self   int
}

// Parse func for ChatParamsChange
func (result *ChatParamsChange) Parse(i []interface{}) error {
	if len(i) < 2 {
		return fmt.Errorf(errFmtTooShortArray, "ChatParamsChange", 2, len(i))
	}

	if v, ok := i[1].(float64); ok {
		result.ChatID = int(v)
	}

	if len(i) > 2 {
		if v, ok := i[2].(float64); ok {
			result.Self = int(v)
		}
	}

	return nil
}

// ChatInfoChange struct for event with code 52
//
// Chat info change
type ChatInfoChange struct {
	TypeID TypeID
	PeerID int
	Info   int
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
		result.Info = int(v)
	}

	return nil
}

// UserTyping struct for event with code 61
//
// User UserID began typing in the conversation ChatID
type UserTyping struct {
	UserID int
	Flags  int
}

// Parse func for UserTyping
func (result *UserTyping) Parse(i []interface{}) error {
	if len(i) < 3 {
		return fmt.Errorf(errFmtTooShortArray, "UserTyping", 3, len(i))
	}

	if v, ok := i[1].(float64); ok {
		result.UserID = int(v)
	}

	if v, ok := i[2].(float64); ok {
		result.Flags = int(v)
	}

	return nil
}

// UserTypingChat struct for event with code 62
//
// User UserID began typing in the conversation ChatID
type UserTypingChat struct {
	UserID int
	ChatID int
}

// Parse func for UserTypingChat
func (result *UserTypingChat) Parse(i []interface{}) error {
	if len(i) < 3 {
		return fmt.Errorf(errFmtTooShortArray, "UserTypingChat", 3, len(i))
	}

	if v, ok := i[1].(float64); ok {
		result.UserID = int(v)
	}

	if v, ok := i[2].(float64); ok {
		result.ChatID = int(v)
	}

	return nil
}

// UsersTyping struct for event with code 63
type UsersTyping struct {
	UserIDs    []int
	PeerID     int
	TotalCount int
	Ts         time.Time
}

// Parse func for UsersTyping
func (result *UsersTyping) Parse(i []interface{}) error {
	if len(i) < 5 {
		return fmt.Errorf(errFmtTooShortArray, "UsersTyping", 5, len(i))
	}

	v, err := interfaceToIDSlice(i[1])
	if err != nil {
		return err
	}

	result.UserIDs = v

	if v, ok := i[2].(float64); ok {
		result.PeerID = int(v)
	}

	if v, ok := i[3].(float64); ok {
		result.TotalCount = int(v)
	}

	if v, ok := i[4].(float64); ok {
		result.Ts = time.Unix(int64(v), 0)
	}

	return nil
}

// UsersRecordingAudioMessage struct for event with code 64
type UsersRecordingAudioMessage struct {
	PeerID     int
	UserIDs    []int
	TotalCount int
	Ts         time.Time
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

	result.UserIDs = v

	if v, ok := i[3].(float64); ok {
		result.TotalCount = int(v)
	}

	if v, ok := i[4].(float64); ok {
		result.Ts = time.Unix(int64(v), 0)
	}

	return nil
}

// UserCall struct for event with code 70
type UserCall struct {
	UserID int
	CallID int
}

// Parse func for UserCall
func (result *UserCall) Parse(i []interface{}) error {
	if len(i) < 3 {
		return fmt.Errorf(errFmtTooShortArray, "UserCall", 3, len(i))
	}

	if v, ok := i[1].(float64); ok {
		result.UserID = int(v)
	}

	if v, ok := i[2].(float64); ok {
		result.CallID = int(v)
	}

	return nil
}

// CounterChange struct for event with code 80
type CounterChange struct {
	Count int
}

// Parse func for CounterChange
func (result *CounterChange) Parse(i []interface{}) error {
	if len(i) < 2 {
		return fmt.Errorf(errFmtTooShortArray, "CounterChange", 2, len(i))
	}

	if v, ok := i[1].(float64); ok {
		result.Count = int(v)
	}

	return nil
}

// NotificationSettingsChange struct for event with code 114
type NotificationSettingsChange struct {
	PeerID        int
	Sound         bool
	DisabledUntil int
}

// should be called if ExtendedEvents flag set
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
