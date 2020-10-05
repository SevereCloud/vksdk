package packer

import (
	"strings"

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
	return strings.ReplaceAll(s, "\"", "\\\"")
}
