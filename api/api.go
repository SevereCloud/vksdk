/*
Package api implements VK API.

See more https://vk.com/dev/api_requests
*/
package api // import "github.com/SevereCloud/vksdk/api"

import (
	"bytes"
	"encoding/json"
	"fmt"
	"mime"
	"net/http"
	"net/url"
	"reflect"
	"sync"
	"time"

	"github.com/SevereCloud/vksdk/api/errors"
	"github.com/SevereCloud/vksdk/internal"
	"github.com/SevereCloud/vksdk/object"
)

// Api constants.
const (
	Version   = "5.122"
	MethodURL = "https://api.vk.com/method/"
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
// "Too many requests per second". (errors.TooMany).
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
// captcha_sid - the obtained identifier;
//
// captcha_key - text entered by the user;
//
// More info: https://vk.com/dev/api_requests
const (
	LimitUserToken  = 3
	LimitGroupToken = 20
)

// VK struct.
type VK struct {
	MethodURL    string
	AccessToken  string
	Version      string
	Client       *http.Client
	IsPoolClient bool
	Limit        int
	UserAgent    string
	Handler      func(method string, params Params) (Response, error)

	tokenPool internal.TokenPool
	mux       sync.Mutex
	lastTime  time.Time
	rps       int
}

// Error struct VK.
//
// Deprecated: use object.Error.
type Error struct {
	Code             int                       `json:"error_code"`
	Message          string                    `json:"error_msg"`
	Text             string                    `json:"error_text"`
	CaptchaSID       string                    `json:"captcha_sid"`
	CaptchaImg       string                    `json:"captcha_img"`
	ConfirmationText string                    `json:"confirmation_text"` //  text of the message to be shown in the default confirmation window.
	RequestParams    []object.BaseRequestParam `json:"request_params"`
}

// Response struct.
type Response struct {
	Response      json.RawMessage       `json:"response"`
	Error         object.Error          `json:"error"`
	ExecuteErrors []object.ExecuteError `json:"execute_errors"`
}

// NewVK returns a new VK.
//
// The VKSDK will use the http.DefaultClient.
// This means that if the http.DefaultClient is modified by other components
// of your application the modifications will be picked up by the SDK as well.
//
// In some cases this might be intended, but it is a better practice
// to create a custom HTTP Client to share explicitly through
// your application. You can configure the VKSDK to use the custom
// HTTP Client by setting the VK.Client value.
//
// This set limit 20 requests per second.
func NewVK(token string) *VK {
	var vk VK
	vk.AccessToken = token
	vk.Version = Version

	vk.Handler = vk.defaultHandler

	// TODO: remove in v2
	vk.MethodURL = MethodURL
	vk.Client = http.DefaultClient
	vk.Limit = LimitGroupToken
	vk.UserAgent = internal.UserAgent
	vk.IsPoolClient = false

	return &vk
}

// NewVKWithPool is similar to NewVK but uses token pool for api calls.
// Use this if you need to increase RPS limit.
func NewVKWithPool(tokens ...string) *VK {
	vk := NewVK("pool")
	vk.tokenPool = internal.NewTokenPool(tokens...)
	vk.Limit = LimitGroupToken * len(tokens)
	vk.IsPoolClient = true

	return vk
}

// Params type.
type Params map[string]interface{}

// Lang - determines the language for the data to be displayed on. For
// example country and city names. If you use a non-cyrillic language,
// cyrillic symbols will be transliterated automatically.
// Numeric format from account.getInfo is supported as well.
//
// 	p.Lang(object.LangRU)
//
// See all language code in module object.
func (p Params) Lang(v int) Params {
	p["lang"] = v
	return p
}

// TestMode allows to send requests from a native app without switching it on
// for all users.
func (p Params) TestMode(v bool) Params {
	p["test_mode"] = v
	return p
}

// CaptchaSID received ID.
//
// See https://vk.com/dev/captcha_error
func (p Params) CaptchaSID(v string) Params {
	p["captcha_sid"] = v
	return p
}

// CaptchaKey text input.
//
// See https://vk.com/dev/captcha_error
func (p Params) CaptchaKey(v string) Params {
	p["captcha_key"] = v
	return p
}

// Confirm parameter.
//
// See https://vk.com/dev/need_confirmation
func (p Params) Confirm(v bool) Params {
	p["confirm"] = v
	return p
}

// defaultHandler provides access to VK API methods.
func (vk *VK) defaultHandler(method string, params Params) (Response, error) {
	u := vk.MethodURL + method
	query := url.Values{}

	for key, value := range params {
		query.Set(key, FmtValue(value, 0))
	}

	attempt := 0

	for {
		var response Response

		attempt++

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

		rawBody := bytes.NewBufferString(query.Encode())

		req, err := http.NewRequest("POST", u, rawBody)
		if err != nil {
			return response, err
		}

		req.Header.Set("User-Agent", vk.UserAgent)
		req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

		resp, err := vk.Client.Do(req)
		if err != nil {
			return response, err
		}

		mediatype, _, _ := mime.ParseMediaType(resp.Header.Get("Content-Type"))
		if mediatype != "application/json" {
			_ = resp.Body.Close()
			return response, fmt.Errorf("invalid content-type")
		}

		err = json.NewDecoder(resp.Body).Decode(&response)
		if err != nil {
			_ = resp.Body.Close()
			return response, err
		}

		_ = resp.Body.Close()

		err = errors.New(response.Error)
		if err != nil {
			if errors.GetType(err) == errors.TooMany && attempt < vk.Limit {
				continue
			}

			return response, err
		}

		return response, nil
	}
}

// Request provides access to VK API methods.
//
// TODO: remove in v2.
func (vk *VK) Request(method string, params Params) ([]byte, error) {
	copyParams := make(Params)
	for key, value := range params {
		copyParams[key] = FmtValue(value, 0)
	}

	if _, ok := copyParams["access_token"]; !ok {
		token := vk.AccessToken
		if vk.IsPoolClient {
			token = vk.tokenPool.Get()
		}

		copyParams["access_token"] = token
	}

	if _, ok := copyParams["v"]; !ok {
		copyParams["v"] = vk.Version
	}

	resp, err := vk.Handler(method, copyParams)

	return resp.Response, err
}

// RequestUnmarshal provides access to VK API methods.
func (vk *VK) RequestUnmarshal(method string, params Params, obj interface{}) error {
	rawResponse, err := vk.Request(method, params)
	if err != nil {
		return err
	}

	return json.Unmarshal(rawResponse, &obj)
}

// ExecuteWithArgs a universal method for calling a sequence of other methods
// while saving and filtering interim results.
//
// The Args map variable allows you to retrieve the parameters passed during
// the request and avoids code formatting.
//
// 	return Args.code; // return parameter "code"
// 	return Args.v; // return parameter "v"
//
// https://vk.com/dev/execute
func (vk *VK) ExecuteWithArgs(code string, params Params, obj interface{}) error {
	token := vk.AccessToken
	if vk.IsPoolClient {
		token = vk.tokenPool.Get()
	}

	copyParams := make(Params)
	for key, value := range params {
		copyParams[key] = FmtValue(value, 0)
	}

	copyParams["code"] = code
	copyParams["access_token"] = token
	copyParams["v"] = vk.Version

	resp, err := vk.Handler("execute", copyParams)

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

// Execute a universal method for calling a sequence of other methods while
// saving and filtering interim results.
//
// https://vk.com/dev/execute
func (vk *VK) Execute(code string, obj interface{}) error {
	return vk.ExecuteWithArgs(code, Params{}, obj)
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

// FmtValue return vk format string.
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

// CaptchaForce api method.
func (vk *VK) CaptchaForce(params Params) (response int, err error) {
	err = vk.RequestUnmarshal("captcha.force", params, &response)
	return
}
