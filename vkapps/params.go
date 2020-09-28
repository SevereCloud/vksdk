package vkapps

import (
	"net/url"
	"regexp"
	"strconv"
	"strings"

	"github.com/gorilla/schema"
)

// Referral source.
type Referral string

// Possible values.
const (
	Catalog           Referral = "catalog"
	FeaturingDiscover Referral = "featuring_discover"
	FeaturingMenu     Referral = "featuring_menu"
	FeaturingNew      Referral = "featuring_new"
	ImChat            Referral = "im_chat"
	Notifications     Referral = "notifications"
	SnippetPost       Referral = "snippet_post"
	SnippetIm         Referral = "snippet_im"
	SuperApp          Referral = "super_app"
	GroupMenu         Referral = "group_menu"
	HomeScreen        Referral = "home_screen"
	LeftNav           Referral = "left_nav"
	QuickSearch       Referral = "quick_search"
	Other             Referral = "other"
	Feed              Referral = "feed"
	Im                Referral = "im"
	Menu              Referral = "menu"
	SnippedPost       Referral = "snipped_post"
	Story             Referral = "story"
	StoryReply        Referral = "story_reply"
	StoryViewer       Referral = "story_viewer"
)

// Catalog transition from the catalog. Return empty string if not Catalog.
func (r Referral) Catalog() (categoryName string) {
	if strings.HasPrefix(string(r), "catalog_") {
		return strings.TrimPrefix(string(r), "catalog_")
	}

	return
}

// Story transition from the story sticker. Return empty string if not Story.
func (r Referral) Story() (userID int, params string) {
	find := regexp.MustCompile(`story(\d+)_(.+?)$`).FindStringSubmatch(string(r))
	if len(find) > 0 {
		userID, _ = strconv.Atoi(find[1])
		params = find[2]
	}

	return
}

// Role in the community from which the application is launched.
type Role string

// Possible values.
const (
	RoleNone   Role = "none"
	RoleMember Role = "member"
	RoleModer  Role = "moder"
	RoleEditor Role = "editor"
	RoleAdmin  Role = "admin"
)

// Platform from which the service is launched.
type Platform string

// Possible values.
const (
	MobileAndroid          Platform = "mobile_android"
	MobileIPhone           Platform = "mobile_iphone"
	MobileWeb              Platform = "mobile_web"
	DesktopWeb             Platform = "desktop_web"
	MobileAndroidMessenger Platform = "mobile_android_messenger"
	MobileIPhoneMessenger  Platform = "mobile_iphone_messenger"
)

// Params service launch parameters.
type Params struct {
	VkUserID                  int      `schema:"vk_user_id"`
	VkAppID                   int      `schema:"vk_app_id"`
	VkIsAppUser               bool     `schema:"vk_is_app_user"`
	VkAreNotificationsEnabled bool     `schema:"vk_are_notifications_enabled"`
	VkIsFavorite              bool     `schema:"vk_is_favorite"`
	VkLanguage                string   `schema:"vk_language"`
	VkRef                     Referral `schema:"vk_ref"`
	VkAccessTokenSettings     string   `schema:"vk_access_token_settings"`
	VkGroupID                 int      `schema:"vk_group_id"`
	VkViewerGroupRole         Role     `schema:"vk_viewer_group_role"`
	VkPlatform                Platform `schema:"vk_platform"`
	VkTs                      string   `schema:"vk_ts"`
	Sign                      string   `schema:"sign"`
}

// NewParams returns Params from url.
func NewParams(u *url.URL) (*Params, error) {
	params := new(Params)
	decoder := schema.NewDecoder()
	decoder.IgnoreUnknownKeys(true)

	err := decoder.Decode(params, u.Query())

	return params, err
}
