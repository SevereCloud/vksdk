/*
Package api implements VK API

See more https://vk.com/dev/api_requests
*/
package api // import "github.com/SevereCloud/vksdk/api"

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"reflect"
	"sync"
	"time"

	"github.com/SevereCloud/vksdk/api/errors"
	"github.com/SevereCloud/vksdk/object"
)

// Api constants
const (
	Version      = "5.103"
	APIMethodURL = "https://api.vk.com/method/"
)

// VKontakte API methods (except for methods from secure and ads sections)
// with user access key or service access key can be accessed
// no more than 3 times per second. The community access key is limited
// to 20 requests per second.
//
// Maximum amount of calls to the secure section methods depends
// on the app's users amount. If an app has less than
// 10 000 users, 5 requests per second,
// up to 100 000 – 8 requests,
// up to 1 000 000 – 20 requests,
// 1 000 000+ – 35 requests.
//
// The ads section methods are subject to their own limitations,
// you can read them on this page - https://vk.com/dev/ads_limits
//
// If one of this limits is exceeded, the server will return following error:
// "Too many requests per second". (errors.TooMany)
//
// If your app's logic implies many requests in a row, check the execute method.
// It allows for up to 25 requests for different methods in a single request.
//
// In addition to restrictions on the frequency of calls, there are also
// quantitative restrictions on calling the same type of methods.
//
// After exceeding the quantitative limit, access to a particular method may
// require entering a captcha (see https://vk.com/dev/captcha_error),
// and may also be temporarily restricted (in this case, the server does
// not return a response to the call of a particular method, but handles
// any other requests without problems).
//
// If this error occurs, the following parameters are also passed in
// the error message:
//
// CaptchaSID - identifier captcha.
//
// CaptchaImg - a link to the image that you want to show the user
// to enter text from that image.
//
// In this case, you should ask the user to enter text from
// the CaptchaImg image and repeat the request by adding parameters to it:
//
// captcha_sid - the obtained identifier
//
// captcha_key - text entered by the user
//
// More info: https://vk.com/dev/api_requests
const (
	LimitUserToken  = 3
	LimitGroupToken = 20
)

// VK struct
type VK struct {
	MethodURL   string
	AccessToken string
	Version     string
	Client      *http.Client
	Limit       int
	Handler     func(method string, params Params) (Response, error)

	mux         sync.Mutex
	lastTime    time.Time
	rps         int
}

// Error struct VK
//
// Deprecated: use object.Error
type Error struct {
	Code             int                       `json:"error_code"`
	Message          string                    `json:"error_msg"`
	Text             string                    `json:"error_text"`
	CaptchaSID       string                    `json:"captcha_sid"`
	CaptchaImg       string                    `json:"captcha_img"`
	ConfirmationText string                    `json:"confirmation_text"` //  text of the message to be shown in the default confirmation window.
	RequestParams    []object.BaseRequestParam `json:"request_params"`
}

// Response struct
type Response struct {
	Response      json.RawMessage       `json:"response"`
	Error         object.Error          `json:"error"`
	ExecuteErrors []object.ExecuteError `json:"execute_errors"`
}

// Init VK API
//
// The VKSDK will use the http.DefaultClient.
// This means that if the http.DefaultClient is modified by other components
// of your application the modifications will be picked up by the SDK as well.
//
// In some cases this might be intended, but it is a better practice
// to create a custom HTTP Client to share explicitly through
// your application. You can configure the VKSDK to use  the custom
// HTTP Client by setting the VK.Client value.
//
// This set limit 20 requests per second.
func Init(token string) *VK {
	var vk VK
	vk.AccessToken = token
	vk.Version = Version
	
	vk.Handler = vk.defaultHandler

	// TODO: remove in v2
	vk.MethodURL = APIMethodURL
	vk.Client = http.DefaultClient
	vk.Limit = LimitGroupToken

	return &vk
}

// Params type
type Params map[string]interface{}

// defaultHandler provides access to VK API methods
func (vk *VK) defaultHandler(method string, params Params) (response Response, err error) {
	u := vk.MethodURL + method
	query := url.Values{}

	for key, value := range params {
		query.Set(key, FmtValue(value, 0))
	}

	// query.Set("access_token", vk.AccessToken)
	// query.Set("v", vk.Version)

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
		return
	}
	defer resp.Body.Close()

	err = json.NewDecoder(resp.Body).Decode(&response)
	if err != nil {
		return
	}

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

	err = errors.New(response.Error)

	return response, err
}

// Request provides access to VK API methods
//
// TODO: remove in v2
func (vk *VK) Request(method string, params Params) ([]byte, error) {
	if _, ok := params["access_token"]; !ok {
		params["access_token"] = vk.AccessToken
	}

	if _, ok := params["v"]; !ok {
		params["v"] = vk.Version
	}

	resp, err := vk.Handler(method, params)

	return resp.Response, err
}

// RequestUnmarshal provides access to VK API methods
func (vk *VK) RequestUnmarshal(method string, params Params, obj interface{}) error {
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
	params := Params{
		"code":         code,
		"access_token": vk.AccessToken,
		"v":            vk.Version,
	}

	resp, err := vk.Handler("execute", params)

	// Add execute errors
	for _, executeError := range resp.ExecuteErrors {
		context := object.Error{
			Code:    executeError.ErrorCode,
			Message: executeError.ErrorMsg,
			RequestParams: []object.BaseRequestParam{
				{
					Key:   "method",
					Value: executeError.Method,
				},
			},
		}

		err = errors.AddErrorContext(err, context)
	}

	jsonErr := json.Unmarshal(resp.Response, &obj)
	if jsonErr != nil {
		return jsonErr
	}

	return err
}

func fmtBool(value bool) string {
	if value {
		return "1"
	}

	return "0"
}

func fmtReflectValue(value reflect.Value, depth int) string {
	switch f := value; value.Kind() {
	case reflect.Invalid:
		return ""
	case reflect.Bool:
		return fmtBool(f.Bool())
	case reflect.Array, reflect.Slice:
		s := ""

		for i := 0; i < f.Len(); i++ {
			if i > 0 {
				s += ","
			}

			s += FmtValue(f.Index(i).Interface(), depth)
		}

		return s
	case reflect.Ptr:
		// pointer to array or slice or struct? ok at top level
		// but not embedded (avoid loops)
		if depth == 0 && f.Pointer() != 0 {
			switch a := f.Elem(); a.Kind() {
			case reflect.Array, reflect.Slice, reflect.Struct, reflect.Map:
				return FmtValue(a.Interface(), depth+1)
			}
		}
	}

	return fmt.Sprint(value)
}

// FmtValue return vk format string
func FmtValue(value interface{}, depth int) string {
	if value == nil {
		return ""
	}

	switch f := value.(type) {
	case bool:
		return fmtBool(f)
	case object.Attachment:
		return f.ToAttachment()
	case object.JSONObject:
		return f.ToJSON()
	case reflect.Value:
		return fmtReflectValue(f, depth)
	}

	return fmtReflectValue(reflect.ValueOf(value), depth)
}

// CaptchaForce api method
func (vk *VK) CaptchaForce(params Params) (response int, err error) {
	err = vk.RequestUnmarshal("captcha.force", params, &response)
	return
}
