package vksdk

import "encoding/json"

// StorageGetResponse struct
type StorageGetResponse []baseRequestParam

// StorageGet returns a value of variable with the name set by key parameter.
func (vk *VK) StorageGet(params map[string]string) (response StorageGetResponse, err error) {
	_, prs := params["keys"]
	if !prs {
		params["keys"] = params["key"]
	}
	rawResponse, err := vk.Request("storage.get", params)
	if err != nil {
		return
	}

	err = json.Unmarshal(rawResponse, &response)
	if err != nil {
		panic(err)
	}

	return
}

// StorageGetKeysResponse struct
type StorageGetKeysResponse []string

// StorageGetKeys returns the names of all variables.
func (vk *VK) StorageGetKeys(params map[string]string) (response StorageGetKeysResponse, err error) {
	rawResponse, err := vk.Request("storage.getKeys", params)
	if err != nil {
		return
	}

	err = json.Unmarshal(rawResponse, &response)
	if err != nil {
		panic(err)
	}

	return
}

// StorageSet saves a value of variable with the name set by key parameter.
func (vk *VK) StorageSet(params map[string]string) (err error) {
	_, err = vk.Request("storage.set", params)

	return
}
