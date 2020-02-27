package api

import (
	"bytes"
	"encoding/json"
	"net/http"

	"github.com/SevereCloud/vksdk/api/errors"
)

// Authenticator interface
type Authenticator interface {
	Authenticate(params Params) error
}

// AuthenticatorFunc func
type AuthenticatorFunc func(params Params) error

// Authenticate method
// Implements Authenticator for AuthenticatorFunc
func (f AuthenticatorFunc) Authenticate(params Params) error {
	return f(params)
}

// TokenAuthenticator func
// Creates new Authenticator which sets given access_token
func TokenAuthenticator(token string) Authenticator {
	return AuthenticatorFunc(func(params Params) error {
		if _, ok := params["access_token"]; !ok {
			params["access_token"] = token
		}

		return nil
	})
}

// AuthCodeFlowRequest struct
type AuthCodeFlowRequest struct {
	ClientID     int    `json:"client_id"`
	ClientSecret string `json:"client_secret"`
	RedirectURL  string `json:"redirect_url"`
	Code         string `json:"code"`
	Version      string `json:"v"`
}

// TokenResponse struct
type TokenResponse struct {
	AccessToken string `json:"access_token"`
	ExpiresIn   int    `json:"expires_in"`
	UserID      int    `json:"user_id"`
}

// gets token from given AuthCodeFlowRequest
func authCodeFlow(auth AuthCodeFlowRequest) (token TokenResponse, err error) {
	encoded, err := json.Marshal(auth)
	if err != nil {
		return
	}

	resp, err := http.Post(
		"https://oauth.vk.com/access_token",
		"application/json",
		bytes.NewReader(encoded),
	)
	if err != nil {
		return
	}

	if resp.Body != nil {
		defer resp.Body.Close()
	}

	var response Response

	err = json.NewDecoder(resp.Body).Decode(&response)
	if err != nil {
		return
	}

	err = errors.New(response.Error)
	if err != nil {
		return
	}

	err = json.Unmarshal(response.Response, &token)

	return token, err
}

// CodeAuthenticator func
// Authorization code flow implementation
func CodeAuthenticator(auth AuthCodeFlowRequest) (Authenticator, error) {
	r, err := authCodeFlow(auth)

	return AuthenticatorFunc(func(params Params) error {
		if _, ok := params["access_token"]; !ok {
			params["access_token"] = r.AccessToken
		}

		return nil
	}), err
}
