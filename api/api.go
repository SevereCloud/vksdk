package api // import "github.com/SevereCloud/vksdk/api"

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/url"
	"sync"
	"time"

	"github.com/SevereCloud/vksdk/errors"
	"github.com/SevereCloud/vksdk/object"
)

const (
	version      = "5.103"
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
func Init(token string) *VK {
	var vk VK
	vk.MethodURL = apiMethodURL
	vk.AccessToken = token
	vk.Version = version
	vk.Client = &http.Client{}
	return &vk
}

// Request provides access to VK API methods
func (vk *VK) Request(method string, params map[string]string) ([]byte, error) {
	var handler struct {
		Response json.RawMessage
		Error    object.Error `json:"error"`
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
	var beforeRps int
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
		beforeRps = vk.rps
		vk.mux.Unlock()
	}

	resp, err := vk.Client.Post(u, "application/x-www-form-urlencoded", rawBody)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if vk.Limit > 0 {
		vk.mux.Lock()

		if beforeRps != vk.rps {
			vk.rps++
		}
		sleepTime := time.Second - time.Since(vk.lastTime)
		if sleepTime < 0 {
			vk.lastTime = time.Now()
			vk.rps = 1
		}

		vk.mux.Unlock()
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(body, &handler)
	if err != nil {
		return nil, err
	}
	err = errors.New(handler.Error)

	return handler.Response, err
}

// RequestUnmarshal provides access to VK API methods
func (vk *VK) RequestUnmarshal(method string, params map[string]string, obj interface{}) error {
	rawResponse, err := vk.Request(method, params)
	if err != nil {
		return err
	}

	return json.Unmarshal(rawResponse, &obj)
}

// Execute a universal method for calling a sequence of other methods while saving and filtering interim results.
//
// https://vk.com/dev/Execute
func (vk *VK) Execute(code string, obj interface{}) error {
	params := map[string]string{"code": code}
	return vk.RequestUnmarshal("execute", params, &obj)
}
