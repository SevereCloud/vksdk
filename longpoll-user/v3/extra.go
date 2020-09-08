package wrapper // import "github.com/SevereCloud/vksdk/v2/longpoll-user/v3"

import (
	"reflect"
	"time"
)

// Action struct for AdditionalData.
type Action struct {
	SourceAct         string // Service action name with multiple dialogs
	SourceMid         string // User ID to whom the service action concerns
	SourceText        string
	SourceOldText     string
	SourceMessage     string
	SourceChatLocalID string
}

// Possible values for Action.SourceAct.
const (
	ChatCreate      = "chat_create"       // create chat
	ChatTitleUpdate = "chat_title_update" // change chat name
	ChatPhotoUpdate = "chat_photo_update" // change chat photo
	ChatInviteUser  = "chat_invite_user"  // invite user to chat
	ChatKickUser    = "chat_kick_user"    // kick out user from chat
)

func (result *Action) parse(v map[string]interface{}) {
	if sourceAct, exist := v["source_act"].(string); exist {
		result.SourceAct = sourceAct

		switch sourceAct {
		case "chat_create":
			if sourceText, ok := v["source_text"].(string); ok {
				result.SourceText = sourceText
			}

		case "chat_title_update":
			if sourceText, ok := v["source_text"].(string); ok {
				result.SourceText = sourceText
			}

			if sourceOldText, ok := v["source_old_text"].(string); ok {
				result.SourceOldText = sourceOldText
			}

		case "chat_kick_user", "chat_invite_user":
			if sourceMid, ok := v["source_mid"].(string); ok {
				result.SourceMid = sourceMid
			}

		case "chat_pin_message", "chat_unpin_message":
			if sourceMid, ok := v["source_mid"].(string); ok {
				result.SourceMid = sourceMid
			}

			if sourceMessage, ok := v["source_message"].(string); ok {
				result.SourceMessage = sourceMessage
			}

			if sourceChatLocalID, ok := v["source_chat_local_id"].(string); ok {
				result.SourceChatLocalID = sourceChatLocalID
			}
		}
	}
}

// AdditionalData struct. If mode contains the flag 2 along with text and the
// message topic, a JSON-object may be passed. This object contains media
// attachments or other additional information. Descriptions of the object
// fields are listed below.
type AdditionalData struct {
	Title     string // Message's subject.
	RefSource string
	From      string // User ID of who sent the message if the message is from a chat
	ExpireTTL string
	IsExpired string
	// FromAdmin ID of the administrator who sent the message. It is returned for
	// messages sent from a community (only for community administrators).
	FromAdmin string
	Emoji     string // Message contains emoji.
	Action
}

func (result *AdditionalData) parse(v map[string]interface{}) {
	if title, ok := v["title"].(string); ok {
		result.Title = title
	}

	if refSource, ok := v["ref_source"].(string); ok {
		result.RefSource = refSource
	}

	if from, ok := v["from"].(string); ok {
		result.From = from
	}

	if fromAdmin, ok := v["from_admin"].(string); ok {
		result.FromAdmin = fromAdmin
	}

	if expireTTL, ok := v["expire_ttl"].(string); ok {
		result.ExpireTTL = expireTTL
	}

	if isExpired, ok := v["is_expired"].(string); ok {
		result.IsExpired = isExpired
	}

	if emoji, ok := v["emoji"].(string); ok {
		result.Emoji = emoji
	}

	result.Action.parse(v)
}

// LongPollAttachments type.
type LongPollAttachments map[string]interface{}

// ExtraFields for a message object.
//
// https://vk.com/dev/using_longpoll_3, point 3.1
type ExtraFields struct {
	PeerID         int       // destination ID
	Timestamp      time.Time // message sent time
	Text           string    // message text
	AdditionalData AdditionalData
	Attachments    LongPollAttachments // attachments, if mode = 2 was chosen
}

func (result *ExtraFields) parseExtraFields(i []interface{}) error {
	length := len(i)

	if length > 3 {
		if v, ok := i[3].(float64); ok {
			result.PeerID = int(v)
		}
	}

	if length > 4 {
		if v, ok := i[4].(float64); ok {
			result.Timestamp = time.Unix(int64(v), 0)
		}
	}

	if length > 5 {
		if v, ok := i[5].(string); ok {
			result.Text = v
		}
	}

	if length > 6 {
		if v, ok := i[6].(map[string]interface{}); ok {
			result.AdditionalData.parse(v)
		}
	}

	if length > 7 {
		if v, ok := i[7].(map[string]interface{}); ok {
			result.Attachments = v
		}
	}

	return nil
}

func interfaceToStringIntMap(m interface{}) (map[string]int, error) {
	reflectedMap := reflect.ValueOf(m)
	if reflectedMap.Kind() != reflect.Map {
		return nil, &expectedSlice{m}
	}

	result := make(map[string]int, reflectedMap.Len())

	for _, key := range reflectedMap.MapKeys() {
		v, ok := reflectedMap.MapIndex(key).Interface().(float64)
		if !ok {
			return nil, &failedCast{reflectedMap.MapIndex(key).Interface()}
		}

		result[key.String()] = int(v)
	}

	return result, nil
}

func interfaceToIDSlice(slice interface{}) ([]int, error) {
	reflectedSlice := reflect.ValueOf(slice)
	if reflectedSlice.Kind() != reflect.Slice {
		return nil, &expectedSlice{slice}
	}

	result := make([]int, reflectedSlice.Len())

	for i := 0; i < reflectedSlice.Len(); i++ {
		v, ok := reflectedSlice.Index(i).Interface().(float64)
		if !ok {
			return nil, &failedCast{reflectedSlice.Index(i).Interface()}
		}

		result[i] = int(v)
	}

	return result, nil
}
