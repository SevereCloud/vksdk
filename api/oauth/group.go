// Package oauth ...
package oauth // import "github.com/SevereCloud/vksdk/v3/api/oauth"

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strconv"
	"strings"

	"github.com/SevereCloud/vksdk/v3/internal"
)

// GroupToken struct with access token.
type GroupToken struct {
	GroupID     int    `json:"group_id"`
	AccessToken string `json:"access_token"`
}

// GroupTokens struct with access tokens.
type GroupTokens struct {
	Groups    []GroupToken `json:"groups"`
	ExpiresIn int          `json:"expires_in"`
}

// NewGroupTokensFromJSON return group tokens.
func NewGroupTokensFromJSON(data []byte) (*GroupTokens, error) {
	var e Error

	err := json.Unmarshal(data, &e)
	if err != nil {
		return nil, fmt.Errorf("oauth: %w", err)
	}

	if e.Type != "" {
		return nil, &e
	}

	var t GroupTokens

	err = json.Unmarshal(data, &t)
	if err != nil {
		return nil, fmt.Errorf("oauth: %w", err)
	}

	return &t, nil
}

// NewGroupTokensFromURL  return group tokens.
func NewGroupTokensFromURL(u *url.URL) (*GroupTokens, error) {
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

	t := &GroupTokens{
		Groups: make([]GroupToken, 0),
	}

	for key, vs := range v {
		// NOTE: vs is not empty
		if strings.HasPrefix(key, "access_token_") {
			groupID, _ := strconv.Atoi(key[13:])
			accessToken := vs[0]

			t.Groups = append(t.Groups, GroupToken{
				GroupID:     groupID,
				AccessToken: accessToken,
			})
		}

		if key == "expires_in" {
			t.ExpiresIn, _ = strconv.Atoi(vs[0])
		}
	}

	return t, nil
}

// GroupParams struct.
type GroupParams struct {
	ClientID    int // required
	RedirectURI string
	GroupIDs    []int
	Display     Display // Default mobile
	Scope       int
	V           string // Default version module
	State       string
}

// Values ...
func (p GroupParams) Values() *url.Values {
	q := &url.Values{}

	q.Set("client_id", strconv.Itoa(p.ClientID))

	if p.RedirectURI == "" {
		q.Set("redirect_uri", DefaultRedirectURI)
	} else {
		q.Set("redirect_uri", p.RedirectURI)
	}

	var groupIDs string

	for i, id := range p.GroupIDs {
		if i != 0 {
			groupIDs += ","
		}

		groupIDs += strconv.Itoa(id)
	}

	q.Set("group_ids", groupIDs)
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

// ImplicitFlowGroup need to run methods directly from users devices. Access
// token received this way can not be used for server requests.
//
// https://dev.vk.com/ru/api/access-token/implicit-flow-community
func ImplicitFlowGroup(p GroupParams) *url.URL {
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

// AuthCodeFlowGroup need to run VK API methods from the server side of an
// application. Access token received this way is not bound to an ip address.
//
// https://dev.vk.com/ru/api/access-token/authcode-flow-community
type AuthCodeFlowGroup struct {
	params       GroupParams
	clientSecret string
	Client       *http.Client
	UserAgent    string
}

// NewAuthCodeFlowGroup returns a new AuthcodeFlowGroup.
func NewAuthCodeFlowGroup(p GroupParams, clientSecret string) *AuthCodeFlowGroup {
	return &AuthCodeFlowGroup{
		params:       p,
		clientSecret: clientSecret,
		Client:       http.DefaultClient,
		UserAgent:    internal.UserAgent,
	}
}

// URL authorization dialog.
func (a AuthCodeFlowGroup) URL() *url.URL {
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

func (a AuthCodeFlowGroup) buildRequest(code string) *http.Request {
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

func (a AuthCodeFlowGroup) request(code string) (*http.Response, error) {
	req := a.buildRequest(code)

	resp, err := a.Client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("oauth: %w", err)
	}

	return resp, nil
}

// Token ...
func (a AuthCodeFlowGroup) Token(u *url.URL) (*GroupTokens, error) {
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

	return NewGroupTokensFromJSON(data)
}
