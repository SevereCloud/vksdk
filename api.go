package vksdk

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

const (
	version = "5.92"
	apiURL  = "https://api.vk.com/method/"
)

// VK struct
type VK struct {
	AccessToken string
	Version     string
}

// Error struct VK
type Error struct {
	Code          int    `json:"error_code"`
	Message       string `json:"error_msg"`
	RequestParams []struct {
		Key   string `json:"key"`
		Value string `json:"value"`
	} `json:"request_params"`
}

// Error function
func (e *Error) Error() string {
	// TODO fix return
	return fmt.Sprintf("code %d: %s", e.Code, e.Message)
}

// Init VK API
func Init(token string) (vk VK) {
	vk.AccessToken = token
	vk.Version = version
	return
}

// Request provides access to VK API methods
func (vk *VK) Request(method string, params map[string]string) ([]byte, error) {
	// TODO ограничитель на запросы
	u := apiURL + method

	query := url.Values{}
	for key, value := range params {
		query.Set(key, value)
	}
	query.Set("access_token", vk.AccessToken)
	query.Set("v", vk.Version)

	rawBody := bytes.NewBufferString(query.Encode())
	resp, err := http.Post(u, "application/x-www-form-urlencoded", rawBody)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var handler struct {
		Error    *Error
		Response json.RawMessage
	}
	err = json.Unmarshal(body, &handler)

	if handler.Error != nil {
		return nil, handler.Error
	}

	return handler.Response, nil
}

// TODO execute
