package api // import "github.com/severecloud/vksdk/api"

import (
	"encoding/json"

	"github.com/severecloud/vksdk/object"
)

// StorageGetResponse struct
type StorageGetResponse []object.BaseRequestParam

// StorageGet returns a value of variable with the name set by key parameter.
func (vk *VK) StorageGet(params map[string]string) (response StorageGetResponse, vkErr Error) {
	_, prs := params["keys"]
	if !prs {
		params["keys"] = params["key"]
	}
	rawResponse, vkErr := vk.Request("storage.get", params)
	if vkErr.Code != 0 {
		return
	}

	err := json.Unmarshal(rawResponse, &response)
	if err != nil {
		panic(err)
	}

	return
}

// StorageGetKeysResponse struct
type StorageGetKeysResponse []string

// StorageGetKeys returns the names of all variables.
func (vk *VK) StorageGetKeys(params map[string]string) (response StorageGetKeysResponse, vkErr Error) {
	rawResponse, vkErr := vk.Request("storage.getKeys", params)
	if vkErr.Code != 0 {
		return
	}

	err := json.Unmarshal(rawResponse, &response)
	if err != nil {
		panic(err)
	}

	return
}

// StorageSet saves a value of variable with the name set by key parameter.
func (vk *VK) StorageSet(params map[string]string) (vkErr Error) {
	_, vkErr = vk.Request("storage.set", params)

	return
}
