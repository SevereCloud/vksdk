package api // import "github.com/severecloud/vksdk/5.92/api"

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/url"

	"github.com/severecloud/vksdk/5.92/object"
	"golang.org/x/net/proxy"
)

const (
	version = "5.92"
	apiURL  = "https://api.vk.com/method/"
)

// VK struct
type VK struct {
	AccessToken  string
	Version      string
	ProxyAddress string
}

// Error struct VK
type Error struct {
	Code          int                       `json:"error_code"`
	Message       string                    `json:"error_msg"`
	Text          string                    `json:"error_text"`
	CaptchaSID    string                    `json:"captcha_sid"`
	CaptchaImg    string                    `json:"captcha_img"`
	RequestParams []object.BaseRequestParam `json:"request_params"`
}

// Init VK API
func Init(token string) (vk VK) {
	vk.AccessToken = token
	vk.Version = version
	return
}

// Request provides access to VK API methods
func (vk VK) Request(method string, params map[string]string) ([]byte, Error) {
	// TODO: ограничитель на запросы
	var handler struct {
		Response json.RawMessage
		Error    Error `json:"error"`
	}

	myClient := &http.Client{}
	if vk.ProxyAddress != "" {
		dialer, err := proxy.SOCKS5("tcp", vk.ProxyAddress, nil, proxy.Direct)
		if err != nil {
			handler.Error.Code = -1
			handler.Error.Message = err.Error()
			return handler.Response, handler.Error
		}
		httpTransport := &http.Transport{}
		httpTransport.Dial = dialer.Dial
		myClient.Transport = httpTransport
	}

	u := apiURL + method

	query := url.Values{}
	for key, value := range params {
		query.Set(key, value)
	}
	query.Set("access_token", vk.AccessToken)
	query.Set("v", vk.Version)

	rawBody := bytes.NewBufferString(query.Encode())
	resp, err := myClient.Post(u, "application/x-www-form-urlencoded", rawBody)
	if err != nil {
		handler.Error.Code = -1
		handler.Error.Message = err.Error()
		return handler.Response, handler.Error
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	err = json.Unmarshal(body, &handler)
	if err != nil {
		panic(err)
	}

	return handler.Response, handler.Error
}

// Execute a universal method for calling a sequence of other methods while saving and filtering interim results.
func (vk VK) Execute(Code string) (response []byte, vkErr Error) {
	p := make(map[string]string)
	p["code"] = Code
	response, vkErr = vk.Request("execute", p)

	return
}
