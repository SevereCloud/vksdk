package api // import "github.com/SevereCloud/vksdk/5.92/api"

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/url"
	"sync"
	"time"

	"github.com/SevereCloud/vksdk/5.92/object"
)

const (
	version      = "5.92"
	apiMethodURL = "https://api.vk.com/method/"
)

// VK struct
type VK struct {
	MethodURL   string
	AccessToken string
	Version     string
	Client      *http.Client
	Limit       int

	mux      sync.Mutex
	lastTime time.Time
	rps      int
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
	vk.MethodURL = apiMethodURL
	vk.AccessToken = token
	vk.Version = version
	vk.Client = &http.Client{}
	return
}

// Request provides access to VK API methods
func (vk *VK) Request(method string, params map[string]string) ([]byte, Error) {
	var handler struct {
		Response json.RawMessage
		Error    Error `json:"error"`
	}

	u := vk.MethodURL + method

	query := url.Values{}
	for key, value := range params {
		query.Set(key, value)
	}
	query.Set("access_token", vk.AccessToken)
	query.Set("v", vk.Version)

	rawBody := bytes.NewBufferString(query.Encode())

	// Rate limiting
	if vk.Limit > 0 {
		vk.mux.Lock()
		sleepTime := time.Second - time.Since(vk.lastTime)
		if sleepTime < 0 {
			vk.lastTime = time.Now()
			vk.rps = 0
		} else if vk.rps == vk.Limit {
			time.Sleep(sleepTime)
			vk.lastTime = time.Now()
			vk.rps = 0
		}
		vk.rps++
		vk.mux.Unlock()
	}

	resp, err := vk.Client.Post(u, "application/x-www-form-urlencoded", rawBody)
	if err != nil {
		handler.Error = NewError(-1, err.Error(), method, params)
		return handler.Response, handler.Error
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	err = json.Unmarshal(body, &handler)
	if err != nil {
		handler.Error = NewError(-1, err.Error(), method, params)
	}

	return handler.Response, handler.Error
}

// RequestUnmarshal provides access to VK API methods
func (vk *VK) RequestUnmarshal(method string, params map[string]string, obj interface{}, vkErr *Error) {
	rawResponse, rawErr := vk.Request(method, params)
	*vkErr = rawErr
	if vkErr.Code != 0 {
		return
	}

	err := json.Unmarshal(rawResponse, &obj)
	if err != nil {
		*vkErr = NewError(-1, err.Error(), method, params)
	}
}

// NewError return new Error
func NewError(code int, message, method string, params map[string]string) (vkErr Error) {
	vkErr.Code = code
	vkErr.Message = message
	vkErr.RequestParams = append(
		vkErr.RequestParams,
		object.BaseRequestParam{
			Key:   "method",
			Value: method,
		},
	)
	for key, value := range params {
		vkErr.RequestParams = append(
			vkErr.RequestParams,
			object.BaseRequestParam{
				Key:   key,
				Value: value,
			},
		)
	}
	vkErr.RequestParams = append(
		vkErr.RequestParams,
		object.BaseRequestParam{
			Key:   "v",
			Value: version,
		},
	)
	return
}

// Execute a universal method for calling a sequence of other methods while saving and filtering interim results.
//
// https://vk.com/dev/Execute
func (vk *VK) Execute(code string, obj interface{}, vkErr *Error) {
	params := map[string]string{"code": code}
	vk.RequestUnmarshal("execute", params, &obj, vkErr)
}
