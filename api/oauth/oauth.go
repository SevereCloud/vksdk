// Package oauth ...
package oauth // import "github.com/Derad6709/vksdk/v2/api/oauth"

import (
	"net/url"

	"github.com/Derad6709/vksdk/v2/api"
)

// nolint:gochecknoglobals
var (
	OAuthHost          = "oauth.vk.com"
	DefaultRedirectURI = "https://oauth.vk.com/blank.html"
)

const (
	scheme  = "https"
	version = api.Version
)

// Access Permissions for User Token.
const (
	ScopeUserNotify        = 1 << 0
	ScopeUserFriends       = 1 << 1
	ScopeUserPhotos        = 1 << 2
	ScopeUserAudio         = 1 << 3
	ScopeUserVideo         = 1 << 4
	ScopeUserStories       = 1 << 6
	ScopeUserPages         = 1 << 7
	ScopeUserMenu          = 1 << 8
	ScopeUserWallmenu      = 1 << 9
	ScopeUserStatus        = 1 << 10
	ScopeUserNotes         = 1 << 11
	ScopeUserMessages      = 1 << 12
	ScopeUserWall          = 1 << 13
	ScopeUserAds           = 1 << 15
	ScopeUserOffline       = 1 << 16
	ScopeUserDocs          = 1 << 17
	ScopeUserGroups        = 1 << 18
	ScopeUserNotifications = 1 << 19
	ScopeUserStats         = 1 << 20
	ScopeUserEmail         = 1 << 22
	ScopeUserAdsweb        = 1 << 23
	ScopeUserLeads         = 1 << 24
	ScopeUserGroupMessages = 1 << 25
	ScopeUserExchange      = 1 << 26
	ScopeUserMarket        = 1 << 27
	ScopeUserPhone         = 1 << 28
)

// Access Permissions for Community Token.
const (
	ScopeGroupStories   = 1 << 0
	ScopeGroupPhotos    = 1 << 2
	ScopeGroupAppWidget = 1 << 6
	ScopeGroupMessages  = 1 << 12
	ScopeGroupDocs      = 1 << 17
	ScopeGroupManage    = 1 << 18
)

// CheckScope ...
func CheckScope(scope int, permissions ...int) bool {
	for i := 0; i < len(permissions); i++ {
		if scope&permissions[i] != permissions[i] {
			return false
		}
	}

	return true
}

// Display sets authorization page appearance.
type Display string

// The supported values.
const (
	Page   Display = "page"   // authorization form in a separate window
	Popup  Display = "popup"  // a pop-up window
	Mobile Display = "mobile" // authorization for mobile devices (uses no Javascript).
)

func parseCode(u *url.URL) (string, error) {
	v, err := url.ParseQuery(u.Fragment)

	if errType := v.Get("error"); errType != "" {
		err = &Error{
			Type:        ErrorType(errType),
			Reason:      ErrorReason(v.Get("error_reason")),
			Description: v.Get("error_description"),
		}
	}

	return v.Get("code"), err
}
