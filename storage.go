package vksdk

import "encoding/json"

// StorageGetResponse struct
type StorageGetResponse []baseRequestParam

// StorageGet returns a value of variable with the name set by key parameter.
func (vk *VK) StorageGet(params map[string]string) (StorageGetResponse, error) {
	var response StorageGetResponse

	_, prs := params["keys"]
	if !prs {
		params["keys"] = params["key"]
	}
	rawResponse, err := vk.Request("storage.get", params)
	if err != nil {
		return response, err
	}

	err = json.Unmarshal(rawResponse, &response)
	if err != nil {
		panic(err)
	}

	return response, nil
}

// StorageGetKeysResponse struct
type StorageGetKeysResponse []string

// StorageGetKeys returns the names of all variables.
func (vk *VK) StorageGetKeys(params map[string]string) (StorageGetKeysResponse, error) {
	var response StorageGetKeysResponse

	rawResponse, err := vk.Request("storage.getKeys", params)
	if err != nil {
		return response, err
	}

	err = json.Unmarshal(rawResponse, &response)
	if err != nil {
		panic(err)
	}

	return response, nil
}

// StorageSet saves a value of variable with the name set by key parameter.
func (vk *VK) StorageSet(params map[string]string) error {
	_, err := vk.Request("storage.set", params)
	if err != nil {
		return err
	}

	return nil
}
