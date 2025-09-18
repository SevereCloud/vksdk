// Package oauth ...
package oauth // import "github.com/SevereCloud/vksdk/v3/api/oauth"

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strconv"

	"github.com/SevereCloud/vksdk/v3/internal"
)

// UserToken ...
type UserToken struct {
	AccessToken string `json:"access_token"`
	ExpiresIn   int    `json:"expires_in"`
	UserID      int    `json:"user_id"`
	Email       string `json:"email,omitempty"`
	State       string `json:"state,omitempty"`
}

// NewUserTokenFromJSON ...
func NewUserTokenFromJSON(data []byte) (*UserToken, error) {
	var e Error

	err := json.Unmarshal(data, &e)
	if err != nil {
		return nil, fmt.Errorf("oauth: %w", err)
	}

	if e.Type != "" {
		return nil, &e
	}

	var t UserToken

	err = json.Unmarshal(data, &t)
	if err != nil {
		return nil, fmt.Errorf("oauth: %w", err)
	}

	return &t, nil
}

// NewUserTokenFromURL ...
func NewUserTokenFromURL(u *url.URL) (*UserToken, error) {
	v, err := url.ParseQuery(u.Fragment)
	if err != nil {
		return nil, fmt.Errorf("oauth: %w", err)
	}

	if errType := v.Get("error"); errType != "" {
		return nil, &Error{
			Type:        ErrorType(errType),
			Reason:      ErrorReason(v.Get("error_reason")),
			Description: v.Get("error_description"),
		}
	}

	expiresIn, _ := strconv.Atoi(v.Get("expires_in"))
	userID, _ := strconv.Atoi(v.Get("user_id"))
	t := &UserToken{
		AccessToken: v.Get("access_token"),
		ExpiresIn:   expiresIn,
		State:       v.Get("state"),
		UserID:      userID,
		Email:       v.Get("email"),
	}

	return t, nil
}

// UserParams struct.
type UserParams struct {
	ClientID    int // required
	RedirectURI string
	Display     Display // Default mobile
	Scope       int
	V           string // Default version module
	State       string
}

// Values ...
func (p UserParams) Values() *url.Values {
	q := &url.Values{}

	q.Set("client_id", strconv.Itoa(p.ClientID))

	if p.RedirectURI == "" {
		q.Set("redirect_uri", DefaultRedirectURI)
	} else {
		q.Set("redirect_uri", p.RedirectURI)
	}

	q.Set("display", string(p.Display))
	q.Set("scope", strconv.Itoa(p.Scope))
	q.Set("state", p.State)

	if p.V == "" {
		q.Set("v", version)
	} else {
		q.Set("v", p.V)
	}

	return q
}

// ImplicitFlowUser need to run methods directly from users devices. Access
// token received this way can not be used for server requests.
//
// https://dev.vk.ru/ru/api/access-token/implicit-flow-user
func ImplicitFlowUser(p UserParams) *url.URL {
	q := p.Values()
	q.Set("response_type", "token")

	u := &url.URL{
		Scheme:   scheme,
		Host:     OAuthHost,
		Path:     "authorize",
		RawQuery: q.Encode(),
	}

	return u
}

// AuthCodeFlowUser need to run VK API methods from the server side of an
// application. Access token received this way is not bound to an ip address
// but set of permissions that can be granted is limited for security reasons.
//
// https://dev.vk.ru/ru/api/access-token/authcode-flow-user
type AuthCodeFlowUser struct {
	params       UserParams
	clientSecret string
	Client       *http.Client
	UserAgent    string
}

// NewAuthCodeFlowUser returns a new AuthcodeFlowUser.
func NewAuthCodeFlowUser(p UserParams, clientSecret string) *AuthCodeFlowUser {
	return &AuthCodeFlowUser{
		params:       p,
		clientSecret: clientSecret,
		Client:       http.DefaultClient,
		UserAgent:    internal.UserAgent,
	}
}

// URL authorization dialog.
func (a AuthCodeFlowUser) URL() *url.URL {
	q := a.params.Values()
	q.Set("response_type", "code")

	u := &url.URL{
		Scheme:   scheme,
		Host:     OAuthHost,
		Path:     "authorize",
		RawQuery: q.Encode(),
	}

	return u
}

func (a AuthCodeFlowUser) buildRequest(code string) *http.Request {
	q := &url.Values{}

	q.Set("client_id", strconv.Itoa(a.params.ClientID))
	q.Set("client_secret", a.clientSecret)
	q.Set("redirect_uri", a.params.RedirectURI)
	q.Set("code", code)

	uReq := &url.URL{
		Scheme:   scheme,
		Host:     OAuthHost,
		Path:     "access_token",
		RawQuery: q.Encode(),
	}

	req, _ := http.NewRequest(http.MethodGet, uReq.String(), nil)
	req.Header.Set("User-Agent", a.UserAgent)

	return req
}

func (a AuthCodeFlowUser) request(code string) (*http.Response, error) {
	req := a.buildRequest(code)

	resp, err := a.Client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("oauth: %w", err)
	}

	return resp, nil
}

// Token ...
func (a AuthCodeFlowUser) Token(u *url.URL) (*UserToken, error) {
	code, err := parseCode(u)
	if err != nil {
		return nil, err
	}

	resp, err := a.request(code)
	if err != nil {
		return nil, fmt.Errorf("oauth: %w", err)
	}
	defer resp.Body.Close()

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("oauth: %w", err)
	}

	return NewUserTokenFromJSON(data)
}

// DirectAuthParams parameters.
type DirectAuthParams struct {
	ClientSecret       string
	ClientID           int
	Username, Password string

	Scope int
	V     string

	TwoFactorSupported bool
	ForceSMS           bool
	Code               string

	CaptchaSID string
	CaptchaKey string

	TestRedirectURI bool

	Client    *http.Client
	UserAgent string
}

func buildDirectAuthRequest(p DirectAuthParams) *http.Request {
	q := &url.Values{}

	q.Set("grant_type", "password")
	q.Set("client_id", strconv.Itoa(p.ClientID))
	q.Set("client_secret", p.ClientSecret)
	q.Set("username", p.Username)
	q.Set("password", p.Password)

	q.Set("scope", strconv.Itoa(p.Scope))

	if p.TestRedirectURI {
		q.Set("test_redirect_uri", "1")
	}

	if p.V == "" {
		p.V = version
	}

	q.Set("v", p.V)

	if p.TwoFactorSupported {
		q.Set("2fa_supported", "1")
	}

	if p.ForceSMS {
		q.Set("force_sms", "1")
	}

	if p.Code != "" {
		q.Set("code", p.Code)
	}

	if p.CaptchaSID != "" && p.CaptchaKey != "" {
		q.Set("captcha_sid", p.CaptchaSID)
		q.Set("captcha_key", p.CaptchaKey)
	}

	uReq := &url.URL{
		Scheme:   scheme,
		Host:     OAuthHost,
		Path:     "token",
		RawQuery: q.Encode(),
	}

	req, _ := http.NewRequest(http.MethodGet, uReq.String(), nil)

	if p.UserAgent == "" {
		p.UserAgent = internal.UserAgent
	}

	req.Header.Set("User-Agent", p.UserAgent)

	return req
}

// DirectAuth type of authorization.
//
// Please note! This type of authorization is available only after preliminary
// approval of VK administration.
//
// To apply for access you need to contact our support service at
// https://vk.ru/support and specify you application ID.
//
// Currently, this functionality is available only for the following categories:
//
// - Fully functional clients. At the time you apply for access your
// application shall be functional (using standard authorization); and you
// shall provide the download link in the request.
//
// - Applications for platforms which do not support standard authorization.
// Provide a short description of application functionality in the request.
//
// Trusted applications can get time-unlimited access_token to access API by
// passing user's login and password.
//
// Note that application shall not store user's password.
// Issued access_token is not bound to user's IP address, that is why it is
// sufficient for using API in the future without repeating authorization
// procedure.
//
// See https://dev.vk.ru/ru/api/direct-auth
func DirectAuth(p DirectAuthParams) (*UserToken, error) {
	req := buildDirectAuthRequest(p)

	if p.Client == nil {
		p.Client = http.DefaultClient
	}

	resp, err := p.Client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("oauth: %w", err)
	}
	defer resp.Body.Close()

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("oauth: %w", err)
	}

	return NewUserTokenFromJSON(data)
}
