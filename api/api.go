package api // import "github.com/severecloud/vksdk/api"

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/url"

	"github.com/severecloud/vksdk/object"
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
	Code          int                     `json:"error_code"`
	Message       string                  `json:"error_msg"`
	Text          string                  `json:"error_text"`
	CaptchaSID    string                  `json:"captcha_sid"`
	CaptchaImg    string                  `json:"captcha_img"`
	RequestParams object.BaseRequestParam `json:"request_params"`
}

// Init VK API
func Init(token string) (vk VK) {
	vk.AccessToken = token
	vk.Version = version
	return
}

// Request provides access to VK API methods
func (vk *VK) Request(method string, params map[string]string) ([]byte, Error) {
	// TODO: ограничитель на запросы
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
		panic(err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	var handler struct {
		Response json.RawMessage
		Error    Error `json:"error"`
	}
	err = json.Unmarshal(body, &handler)
	if err != nil {
		panic(err)
	}

	return handler.Response, handler.Error
}

// Execute a universal method for calling a sequence of other methods while saving and filtering interim results.
func (vk *VK) Execute(Code string) (response []byte, vkErr Error) {
	p := make(map[string]string)
	p["code"] = Code
	response, vkErr = vk.Request("execute", p)

	return
}
