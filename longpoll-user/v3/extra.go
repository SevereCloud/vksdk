package wrapper

import (
	"time"
)

type AdditionalData struct {
	Title             string
	RefSource         string
	From              string
	SourceAct         string
	SourceMid         string
	SourceText        string
	SourceOldText     string
	SourceMessage     string
	SourceChatLocalId string
	FromAdmin         string
	Emoji             string
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

	if sourceAct, ok := v["source_act"].(string); ok {
		result.SourceAct = sourceAct
	}

	switch result.SourceAct {
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

		if sourceChatLocalId, ok := v["source_chat_local_id"].(string); ok {
			result.SourceChatLocalId = sourceChatLocalId
		}
	}

	if fromAdmin, ok := v["from_admin"].(string); ok {
		result.FromAdmin = fromAdmin
	}

	if emoji, ok := v["emoji"].(string); ok {
		result.Emoji = emoji
	}
}

type LongPollAttachments map[string]interface{}

// https://vk.com/dev/using_longpoll_3, point 3.1
type ExtraFields struct {
	PeerID         int
	Timestamp      time.Time
	Text           string
	AdditionalData AdditionalData
	Attachments    LongPollAttachments
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

func convertSlice(v []float64) []int {
	r := make([]int, len(v))

	for i, v := range v {
		r[i] = int(v)
	}

	return r
}
