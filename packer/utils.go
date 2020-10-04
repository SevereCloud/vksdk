package packer

import "github.com/SevereCloud/vksdk/v2/api"

func getTokenFromParams(params ...api.Params) (interface{}, bool) {
	for _, pmap := range params {
		for k, v := range pmap {
			if k == "access_token" {
				return v, true
			}
		}
	}

	return nil, false
}

func iterateAll(iterFn func(key string, value interface{}), params ...api.Params) {
	for _, pmap := range params {
		for k, v := range pmap {
			iterFn(k, v)
		}
	}
}
