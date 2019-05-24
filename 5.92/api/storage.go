package api // import "github.com/SevereCloud/vksdk/5.92/api"

import (
	"github.com/SevereCloud/vksdk/5.92/object"
)

// StorageGetResponse struct
type StorageGetResponse []object.BaseRequestParam

// StorageGet returns a value of variable with the name set by key parameter.
//
// https://vk.com/dev/storage.get
func (vk *VK) StorageGet(params map[string]string) (response StorageGetResponse, vkErr Error) {
	_, prs := params["keys"]
	if !prs {
		params["keys"] = params["key"]
	}
	vk.RequestUnmarshal("storage.get", params, &response, &vkErr)
	return
}

// StorageGetKeysResponse struct
type StorageGetKeysResponse []string

// StorageGetKeys returns the names of all variables.
//
// https://vk.com/dev/storage.getKeys
func (vk *VK) StorageGetKeys(params map[string]string) (response StorageGetKeysResponse, vkErr Error) {
	vk.RequestUnmarshal("storage.getKeys", params, &response, &vkErr)
	return
}

// StorageSetResponse struct
type StorageSetResponse int

// StorageSet saves a value of variable with the name set by key parameter.
//
// https://vk.com/dev/storage.set
func (vk *VK) StorageSet(params map[string]string) (response StorageSetResponse, vkErr Error) {
	vk.RequestUnmarshal("storage.set", params, &response, &vkErr)
	return
}
