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
	version = "5.92"
	apiURL  = "https://api.vk.com/method/"
)

// VK struct
type VK struct {
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

	u := apiURL + method

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
		handler.Error.Code = -1
		handler.Error.Message = err.Error()
	}

	return handler.Response, handler.Error
}

func (vk *VK) requestU(method string, params map[string]string, obj interface{}, vkErr *Error) {
	rawResponse, rawErr := vk.Request(method, params)
	*vkErr = rawErr
	if vkErr.Code != 0 {
		return
	}

	err := json.Unmarshal(rawResponse, &obj)
	if err != nil {
		vkErr.Code = -1
		vkErr.Message = err.Error()
	}
}

// Execute a universal method for calling a sequence of other methods while saving and filtering interim results.
func (vk *VK) Execute(code string) (response []byte, vkErr Error) {
	p := make(map[string]string)
	p["code"] = code
	response, vkErr = vk.Request("execute", p)

	return
}
