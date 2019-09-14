package api // import "github.com/SevereCloud/vksdk/api"

import (
	"github.com/SevereCloud/vksdk/object"
)

// StorageGetResponse struct
type StorageGetResponse []object.BaseRequestParam

// StorageGet returns a value of variable with the name set by key parameter.
//
// StorageGet always return array!
// https://vk.com/dev/storage.get
func (vk *VK) StorageGet(params map[string]string) (response StorageGetResponse, err error) {
	_, prs := params["keys"]
	if !prs {
		params["keys"] = params["key"]
		params["key"] = ""
	}
	err = vk.RequestUnmarshal("storage.get", params, &response)
	return
}

// StorageGetKeysResponse struct
type StorageGetKeysResponse []string

// StorageGetKeys returns the names of all variables.
//
// https://vk.com/dev/storage.getKeys
func (vk *VK) StorageGetKeys(params map[string]string) (response StorageGetKeysResponse, err error) {
	err = vk.RequestUnmarshal("storage.getKeys", params, &response)
	return
}

// StorageSet saves a value of variable with the name set by key parameter.
//
// https://vk.com/dev/storage.set
func (vk *VK) StorageSet(params map[string]string) (response int, err error) {
	err = vk.RequestUnmarshal("storage.set", params, &response)
	return
}
