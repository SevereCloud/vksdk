package packer

import (
	"github.com/SevereCloud/vksdk/v2/api"
)

func iterateAll(iterFn func(key string, value interface{}), params ...api.Params) {
	for _, pmap := range params {
		for k, v := range pmap {
			iterFn(k, v)
		}
	}
}

func escape(s string) string {
	newstr := make([]rune, 0, len(s))
	escaped := false

	for _, r := range s {
		if r == '"' && !escaped {
			newstr = append(newstr, '\\')
		}

		newstr = append(newstr, r)
		escaped = r == '\\'
	}

	return string(newstr)
}
