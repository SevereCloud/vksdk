// Package oauth ...
package oauth // import "github.com/SevereCloud/vksdk/v2/api/oauth"

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"

	"github.com/SevereCloud/vksdk/v2/internal"
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
		return nil, err
	}

	if e.Type != "" {
		return nil, &e
	}

	var t UserToken
	err = json.Unmarshal(data, &t)

	return &t, err
}

// NewUserTokenFromURL ...
func NewUserTokenFromURL(u *url.URL) (*UserToken, error) {
	v, err := url.ParseQuery(u.Fragment)
	if err != nil {
		return nil, err
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
		q.Set("redirect_uri", defaultRedirectURI)
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
// https://vk.com/dev/implicit_flow_user
func ImplicitFlowUser(p UserParams) *url.URL {
	q := p.Values()
	q.Set("response_type", "token")

	u := &url.URL{
		Scheme:   scheme,
		Host:     oauthHost,
		Path:     "authorize",
		RawQuery: q.Encode(),
	}

	return u
}

// AuthCodeFlowUser need to run VK API methods from the server side of an
// application. Access token received this way is not bound to an ip address
// but set of permissions that can be granted is limited for security reasons.
//
// https://vk.com/dev/authcode_flow_user
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
		Host:     oauthHost,
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
		Host:     oauthHost,
		Path:     "access_token",
		RawQuery: q.Encode(),
	}

	req, _ := http.NewRequest("GET", uReq.String(), nil)
	req.Header.Set("User-Agent", a.UserAgent)

	return req
}

func (a AuthCodeFlowUser) request(code string) (*http.Response, error) {
	req := a.buildRequest(code)

	return a.Client.Do(req)
}

// Token ...
func (a AuthCodeFlowUser) Token(u *url.URL) (*UserToken, error) {
	code, err := parseCode(u)
	if err != nil {
		return nil, err
	}

	resp, err := a.request(code)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return NewUserTokenFromJSON(data)
}
