// Package callback implements Callback API.
package callback // import "github.com/SevereCloud/vksdk/callback"

import (
	"crypto/rand"
	"errors"

	"github.com/SevereCloud/vksdk"
	"github.com/SevereCloud/vksdk/api"
)

// generateRandomBytes returns securely generated random bytes.
// It will return an error if the system's secure random
// number generator fails to function correctly, in which
// case the caller should not continue.
func generateRandomBytes(n int) ([]byte, error) {
	b := make([]byte, n)
	_, err := rand.Read(b)
	// Note that err == nil only if we read len(b) bytes.
	if err != nil {
		return nil, err
	}

	return b, nil
}

// generateRandomString returns a securely generated random string.
// It will return an error if the system's secure random
// number generator fails to function correctly, in which
// case the caller should not continue.
func generateRandomString(n int) (string, error) {
	const letters = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"

	bytes, err := generateRandomBytes(n)
	if err != nil {
		return "", err
	}

	for i, b := range bytes {
		bytes[i] = letters[b%byte(len(letters))]
	}

	return string(bytes), nil
}

// ErrNeedGroupToken for AutoSetting.
var ErrNeedGroupToken = errors.New("callback: need group access token")

// AutoSetting automatically configures callback.
//
// Need *api.VK with group access token, access setting: community management.
func (cb *Callback) AutoSetting(
	vk *api.VK,
	urlCallback string,
) error {
	callbackServerID := 0

	g, err := vk.GroupsGetByID(api.Params{})
	if err != nil {
		if errors.Is(err, api.ErrParam) {
			return ErrNeedGroupToken
		}

		return err
	}

	groupID := g[0].ID

	callbackServers, err := vk.GroupsGetCallbackServers(api.Params{
		"group_id": groupID,
	})
	if err != nil {
		return err
	}

	for _, cbServer := range callbackServers.Items {
		if cbServer.URL == urlCallback {
			// Checking the server status
			if cbServer.Status == "ok" {
				// Find confirmed Callback server
				callbackServerID = cbServer.ID
				cb.SecretKeys[groupID] = cbServer.SecretKey

				break
			} else {
				// Broken Callback server
				_, _ = vk.GroupsDeleteCallbackServer(api.Params{
					"group_id":  groupID,
					"server_id": cbServer.ID,
				})
			}
		}
	}

	// If we did not find a server, create a new one
	if callbackServerID == 0 {
		// Generate secret key
		secretKey, err := generateRandomString(24)
		if err != nil {
			return err
		}

		cb.SecretKeys[groupID] = secretKey

		// Get confirmation code
		confirmationCodeResponse, err := vk.GroupsGetCallbackConfirmationCode(api.Params{
			"group_id": groupID,
		})
		if err != nil {
			return err
		}

		cb.ConfirmationKeys[groupID] = confirmationCodeResponse.Code

		// Web server must be running
		addCallbackResponse, err := vk.GroupsAddCallbackServer(api.Params{
			"group_id":   groupID,
			"url":        urlCallback,
			"title":      cb.Title,
			"secret_key": secretKey,
		})
		if err != nil {
			return err
		}

		callbackServerID = addCallbackResponse.ServerID
	}

	params := api.Params{
		"group_id":    groupID,
		"server_id":   callbackServerID,
		"api_version": vksdk.API,
	}
	for _, event := range cb.ListEvents() {
		params[string(event)] = true
	}

	// Updating Callback settings
	_, err = vk.GroupsSetCallbackSettings(params)

	return err
}
