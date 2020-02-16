package wrapper

import (
	"time"
)

type AdditionalData struct {
	Title     string
	RefSource string
	From      int64
}

func (result *AdditionalData) parse(v map[string]interface{}) {
	if title, ok := v["title"].(string); ok {
		result.Title = title
	}
	if refSource, ok := v["ref_source"].(string); ok {
		result.RefSource = refSource
	}
	if from, ok := v["from"].(float64); ok {
		result.From = int64(from)
	}
}

type LongPollAttachments map[string]interface{}

// https://vk.com/dev/using_longpoll_3, point 3.1
type ExtraFields struct {
	PeerId         int
	Timestamp      time.Time
	Text           string
	AdditionalData AdditionalData
	Attachments    LongPollAttachments
}

func (result *ExtraFields) parseExtraFields(i []interface{}) error {
	length := len(i)

	if length > 3 {
		if v, ok := i[3].(float64); ok {
			result.PeerId = int(v)
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
