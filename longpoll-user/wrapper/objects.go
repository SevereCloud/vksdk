package wrapper

import (
	"time"
)

// event with code 1
type MessageFlagsChange struct {
	MessageId int
	Flags     MessageFlag
	ExtraFields
}

func (result *MessageFlagsChange) Parse(i []interface{}) error {
	if len(i) < 3 {
		return ErrInvalidEvent
	}

	if v, ok := i[1].(float64); ok {
		result.MessageId = int(v)
	}
	if v, ok := i[2].(float64); ok {
		result.Flags = MessageFlag(int(v))
	}

	return result.ExtraFields.parseExtraFields(i)
}

// event with code 2
type MessageFlagsSet struct {
	MessageId int
	Mask      MessageFlag
	ExtraFields
}

func (result *MessageFlagsSet) Parse(i []interface{}) error {
	if len(i) < 3 {
		return ErrInvalidEvent
	}

	if v, ok := i[1].(float64); ok {
		result.MessageId = int(v)
	}
	if v, ok := i[2].(float64); ok {
		result.Mask = MessageFlag(int(v))
	}

	return result.ExtraFields.parseExtraFields(i)
}

// event with code 3
type MessageFlagsReset struct {
	MessageId int
	Mask      MessageFlag
	ExtraFields
}

func (result *MessageFlagsReset) Parse(i []interface{}) error {
	if len(i) < 3 {
		return ErrInvalidEvent
	}

	if v, ok := i[1].(float64); ok {
		result.MessageId = int(v)
	}
	if v, ok := i[2].(float64); ok {
		result.Mask = MessageFlag(int(v))
	}

	return result.ExtraFields.parseExtraFields(i)
}

// event with code 4
type NewMessage struct {
	MessageId int
	Flags     MessageFlag
	ExtraFields
}

func (result *NewMessage) Parse(i []interface{}) error {
	if len(i) < 3 {
		return ErrInvalidEvent
	}

	if v, ok := i[1].(float64); ok {
		result.MessageId = int(v)
	}
	if v, ok := i[2].(float64); ok {
		result.Flags = MessageFlag(int(v))
	}

	return result.ExtraFields.parseExtraFields(i)
}

// event with code 5
type EditMessage struct {
	MessageId      int
	Flags          MessageFlag
	PeerId         int
	Timestamp      time.Time
	NewText        string
	AdditionalData AdditionalData
	Attachments    LongPollAttachments
}

func (result *EditMessage) Parse(i []interface{}) error {
	if len(i) < 6 {
		return ErrInvalidEvent
	}

	if v, ok := i[1].(float64); ok {
		result.MessageId = int(v)
	}
	if v, ok := i[2].(float64); ok {
		result.Flags = MessageFlag(int(v))
	}
	if v, ok := i[3].(float64); ok {
		result.PeerId = int(v)
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

// event with code 6
type ReadInMessages struct {
	PeerId  int
	LocalId int
}

func (result *ReadInMessages) Parse(i []interface{}) error {
	if len(i) < 3 {
		return ErrInvalidEvent
	}
	if v, ok := i[1].(float64); ok {
		result.PeerId = int(v)
	}
	if v, ok := i[2].(float64); ok {
		result.LocalId = int(v)
	}
	return nil
}

// event with code 7
type ReadOutMessages struct {
	PeerId  int
	LocalId int
}

func (result *ReadOutMessages) Parse(i []interface{}) error {
	if len(i) < 3 {
		return ErrInvalidEvent
	}
	if v, ok := i[1].(float64); ok {
		result.PeerId = int(v)
	}
	if v, ok := i[2].(float64); ok {
		result.LocalId = int(v)
	}
	return nil
}

// event with code 8
type FriendBecameOnline struct {
	UserId    int
	Extra     int
	Timestamp time.Time
}

func (result *FriendBecameOnline) Parse(i []interface{}) error {
	if len(i) < 4 {
		return ErrInvalidEvent
	}
	if v, ok := i[1].(float64); ok {
		result.UserId = int(v)
	}
	if v, ok := i[2].(float64); ok {
		result.Extra = int(v)
	}
	if v, ok := i[3].(float64); ok {
		result.Timestamp = time.Unix(int64(v), 0)
	}
	return nil
}

// event with code 9
type FriendBecameOffline struct {
	UserId    int
	Flags     int
	Timestamp time.Time
}

func (result *FriendBecameOffline) Parse(i []interface{}) error {
	if len(i) < 4 {
		return ErrInvalidEvent
	}
	if v, ok := i[1].(float64); ok {
		result.UserId = int(v)
	}
	if v, ok := i[2].(float64); ok {
		result.Flags = int(v)
	}
	if v, ok := i[3].(float64); ok {
		result.Timestamp = time.Unix(int64(v), 0)
	}
	return nil
}

// event with code 10
type DialogFlagsReset struct {
	PeerId int
	Mask   DialogFlag
}

func (result *DialogFlagsReset) Parse(i []interface{}) error {
	if len(i) < 3 {
		return ErrInvalidEvent
	}
	if v, ok := i[1].(float64); ok {
		result.PeerId = int(v)
	}
	if v, ok := i[2].(float64); ok {
		result.Mask = DialogFlag(int(v))
	}
	return nil
}

// event with code 11
type DialogFlagsReplace struct {
	PeerId int
	Flags  DialogFlag
}

func (result *DialogFlagsReplace) Parse(i []interface{}) error {
	if len(i) < 3 {
		return ErrInvalidEvent
	}
	if v, ok := i[1].(float64); ok {
		result.PeerId = int(v)
	}
	if v, ok := i[2].(float64); ok {
		result.Flags = DialogFlag(int(v))
	}
	return nil
}

// event with code 12
type DialogsFlagsSet struct {
	PeerId int
	Mask   DialogFlag
}

func (result *DialogsFlagsSet) Parse(i []interface{}) error {
	if len(i) < 3 {
		return ErrInvalidEvent
	}
	if v, ok := i[1].(float64); ok {
		result.PeerId = int(v)
	}
	if v, ok := i[2].(float64); ok {
		result.Mask = DialogFlag(int(v))
	}
	return nil
}

// event with code 13
type DeleteMessages struct {
	PeerId  int
	LocalId int
}

func (result *DeleteMessages) Parse(i []interface{}) error {
	if len(i) < 3 {
		return ErrInvalidEvent
	}
	if v, ok := i[1].(float64); ok {
		result.PeerId = int(v)
	}
	if v, ok := i[2].(float64); ok {
		result.LocalId = int(v)
	}
	return nil
}

// event with code 14
type RestoreDeletedMessages struct {
	PeerId  int
	LocalId int
}

func (result *RestoreDeletedMessages) Parse(i []interface{}) error {
	if len(i) < 3 {
		return ErrInvalidEvent
	}
	if v, ok := i[1].(float64); ok {
		result.PeerId = int(v)
	}
	if v, ok := i[2].(float64); ok {
		result.LocalId = int(v)
	}
	return nil
}

// event with code 51
type ChatParamsChange struct {
	ChatId int
	Self   int
}

func (result *ChatParamsChange) Parse(i []interface{}) error {
	if len(i) < 3 {
		return ErrInvalidEvent
	}
	if v, ok := i[1].(float64); ok {
		result.ChatId = int(v)
	}
	if v, ok := i[2].(float64); ok {
		result.Self = int(v)
	}
	return nil
}

// event with code 52
type ChatInfoChange struct {
	TypeId TypeId
	PeerId int
	Info   int
}

func (result *ChatInfoChange) Parse(i []interface{}) error {
	if len(i) < 4 {
		return ErrInvalidEvent
	}
	if v, ok := i[1].(float64); ok {
		result.TypeId = TypeId(int(v))
	}
	if v, ok := i[2].(float64); ok {
		result.PeerId = int(v)
	}
	if v, ok := i[3].(float64); ok {
		result.Info = int(v)
	}
	return nil
}

// event with code 61
type UserTyping struct {
	UserId int
	Flags  int
}

func (result *UserTyping) Parse(i []interface{}) error {
	if len(i) < 3 {
		return ErrInvalidEvent
	}
	if v, ok := i[1].(float64); ok {
		result.UserId = int(v)
	}
	if v, ok := i[2].(float64); ok {
		result.Flags = int(v)
	}
	return nil
}

// event with code 62
type UserTypingChat struct {
	UserId int
	ChatId int
}

func (result *UserTypingChat) Parse(i []interface{}) error {
	if len(i) < 3 {
		return ErrInvalidEvent
	}
	if v, ok := i[1].(float64); ok {
		result.UserId = int(v)
	}
	if v, ok := i[2].(float64); ok {
		result.ChatId = int(v)
	}
	return nil
}

// event with code 63
type UsersTyping struct {
	UserIds    []int
	PeerId     int
	TotalCount int
	Ts         time.Time
}

func (result *UsersTyping) Parse(i []interface{}) error {
	if len(i) < 5 {
		return ErrInvalidEvent
	}
	if v, ok := i[1].([]float64); ok {
		result.UserIds = convertSlice(v)
	}
	if v, ok := i[2].(float64); ok {
		result.PeerId = int(v)
	}
	if v, ok := i[3].(float64); ok {
		result.TotalCount = int(v)
	}
	if v, ok := i[4].(float64); ok {
		result.Ts = time.Unix(int64(v), 0)
	}
	return nil
}

// event with code 64
type UsersRecordingAudioMessage struct {
	UserIds    []int
	PeerId     int
	TotalCount int
	Ts         time.Time
}

func (result *UsersRecordingAudioMessage) Parse(i []interface{}) error {
	if len(i) < 5 {
		return ErrInvalidEvent
	}
	if v, ok := i[1].([]float64); ok {
		result.UserIds = convertSlice(v)
	}
	if v, ok := i[2].(float64); ok {
		result.PeerId = int(v)
	}
	if v, ok := i[3].(float64); ok {
		result.TotalCount = int(v)
	}
	if v, ok := i[4].(float64); ok {
		result.Ts = time.Unix(int64(v), 0)
	}
	return nil
}

// event with code 70
type UserCall struct {
	UserId int
	CallId int
}

func (result *UserCall) Parse(i []interface{}) error {
	if len(i) < 3 {
		return ErrInvalidEvent
	}
	if v, ok := i[1].(float64); ok {
		result.UserId = int(v)
	}
	if v, ok := i[2].(float64); ok {
		result.CallId = int(v)
	}
	return nil
}

// event with code 80
type CounterChange struct {
	Count int
}

func (result *CounterChange) Parse(i []interface{}) error {
	if len(i) < 2 {
		return ErrInvalidEvent
	}
	if v, ok := i[1].(float64); ok {
		result.Count = int(v)
	}
	return nil
}

// event with code 114
type NotificationSettingsChange struct {
	PeerId        int
	Sound         bool
	DisabledUntil int
}

func (result *NotificationSettingsChange) Parse(i []interface{}) error {
	if len(i) < 4 {
		return ErrInvalidEvent
	}
	if v, ok := i[1].(float64); ok {
		result.PeerId = int(v)
	}
	if v, ok := i[2].(float64); ok {
		result.Sound = int(v) > 0
	}
	if v, ok := i[3].(float64); ok {
		result.DisabledUntil = int(v)
	}
	return nil
}
