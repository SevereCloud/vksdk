package vksdk

import "encoding/json"

// AccountBan account.ban
func (vk *VK) AccountBan(params map[string]string) (vkErr Error) {
	_, vkErr = vk.Request("account.ban", params)
	return
}

// AccountChangePasswordResponse struct
type AccountChangePasswordResponse struct {
	Token string `json:"token"`
}

// AccountChangePassword changes a user password after access is successfully restored with the auth.restore method.
func (vk *VK) AccountChangePassword(params map[string]string) (response AccountChangePasswordResponse, vkErr Error) {
	rawResponse, vkErr := vk.Request("account.changePassword", params)
	if vkErr.Code != 0 {
		return
	}

	err := json.Unmarshal(rawResponse, &response)
	if err != nil {
		panic(err)
	}

	return
}

// AccountGetActiveOffersResponse struct
type AccountGetActiveOffersResponse struct {
	Count int            `json:"count"`
	Items []accountOffer `json:"items"`
}

// AccountGetActiveOffers returns a list of active ads (offers).If the user fulfill their conditions, he will be able to get the appropriate number of votes to his balance.
func (vk *VK) AccountGetActiveOffers(params map[string]string) (response AccountGetActiveOffersResponse, vkErr Error) {
	rawResponse, vkErr := vk.Request("account.getActiveOffers", params)
	if vkErr.Code != 0 {
		return
	}

	err := json.Unmarshal(rawResponse, &response)
	if err != nil {
		panic(err)
	}

	return
}

// AccountGetAppPermissions gets settings of the user in this application.
func (vk *VK) AccountGetAppPermissions(params map[string]string) (response int, vkErr Error) {
	rawResponse, vkErr := vk.Request("account.getAppPermissions", params)
	if vkErr.Code != 0 {
		return
	}

	err := json.Unmarshal(rawResponse, &response)
	if err != nil {
		panic(err)
	}

	return
}

// AccountGetBannedResponse struct
type AccountGetBannedResponse struct {
	Count    int           `json:"count"`
	Items    []int         `json:"items"`
	Profiles []usersUser   `json:"profiles"`
	Groups   []groupsGroup `json:"groups"`
}

// AccountGetBanned returns a user's blacklist.
func (vk *VK) AccountGetBanned(params map[string]string) (response AccountGetBannedResponse, vkErr Error) {
	rawResponse, vkErr := vk.Request("account.getBanned", params)
	if vkErr.Code != 0 {
		return
	}

	err := json.Unmarshal(rawResponse, &response)
	if err != nil {
		panic(err)
	}

	return
}

// AccountGetCountersResponse struct
type AccountGetCountersResponse struct {
	AppRequests        int `json:"app_requests"`
	Events             int `json:"events"`
	Friends            int `json:"friends"`
	FriendsSuggestions int `json:"friends_suggestions"`
	Gifts              int `json:"gifts"`
	Groups             int `json:"groups"`
	Messages           int `json:"messages"`
	Notifications      int `json:"notifications"`
	Photos             int `json:"photos"`
	SDK                int `json:"sdk"`
	Videos             int `json:"videos"`
}

// AccountGetCounters returns non-null values of user counters.
func (vk *VK) AccountGetCounters(params map[string]string) (response AccountGetCountersResponse, vkErr Error) {
	rawResponse, vkErr := vk.Request("account.getCounters", params)
	if vkErr.Code != 0 {
		return
	}

	err := json.Unmarshal(rawResponse, &response)
	if err != nil {
		panic(err)
	}

	return
}

// AccountGetInfoResponse struct
type AccountGetInfoResponse struct {
	Country           string `json:"country"`
	HTTPSRequired     int    `json:"https_required"`
	Intro             int    `json:"intro"`
	Lang              int    `json:"lang"`
	NoWallReplies     int    `json:"no_wall_replies"`
	OwnPostsDefault   int    `json:"own_posts_default"`
	TwoFactorRequired int    `json:"2fa_required"`
}

// AccountGetInfo returns current account info.
func (vk *VK) AccountGetInfo(params map[string]string) (response AccountGetInfoResponse, vkErr Error) {
	rawResponse, vkErr := vk.Request("account.getInfo", params)
	if vkErr.Code != 0 {
		return
	}

	err := json.Unmarshal(rawResponse, &response)
	if err != nil {
		panic(err)
	}

	return
}

// AccountGetProfileInfoResponse struct
type AccountGetProfileInfoResponse struct {
	Bdate            string             `json:"bdate"`
	BdateVisibility  int                `json:"bdate_visibility"`
	City             baseObject         `json:"city"`
	Country          baseCountry        `json:"country"`
	FirstName        string             `json:"first_name"`
	HomeTown         string             `json:"home_town"`
	LastName         string             `json:"last_name"`
	MaidenName       string             `json:"maiden_name"`
	NameRequest      accountNameRequest `json:"name_request"`
	Phone            string             `json:"phone"`
	Relation         int                `json:"relation"`
	RelationPartner  usersUserMin       `json:"relation_partner"`
	RelationPending  int                `json:"relation_pending"`
	RelationRequests []usersUserMin     `json:"relation_requests"`
	ScreenName       string             `json:"screen_name"`
	Sex              baseSex            `json:"sex"`
	Status           string             `json:"status"`
}

// AccountGetProfileInfo returns the current account info.
func (vk *VK) AccountGetProfileInfo() (response AccountGetProfileInfoResponse, vkErr Error) {
	rawResponse, vkErr := vk.Request("account.getProfileInfo", make(map[string]string))
	if vkErr.Code != 0 {
		return
	}

	err := json.Unmarshal(rawResponse, &response)
	if err != nil {
		panic(err)
	}

	return
}

// AccountGetPushSettingsResponse struct
type AccountGetPushSettingsResponse struct {
	Conversations accountPushConversations `json:"conversations"`
	Disabled      int                      `json:"disabled"`
	DisabledUntil int                      `json:"disabled_until"`
	Settings      accountPushParams        `json:"settings"`
}

// AccountGetPushSettings account.getPushSettings Gets settings of push notifications.
func (vk *VK) AccountGetPushSettings(params map[string]string) (response AccountGetPushSettingsResponse, vkErr Error) {
	rawResponse, vkErr := vk.Request("account.getPushSettings", params)
	if vkErr.Code != 0 {
		return
	}

	err := json.Unmarshal(rawResponse, &response)
	if err != nil {
		panic(err)
	}

	return
}

// AccountRegisterDevice subscribes an iOS/Android/Windows/Mac based device to receive push notifications
func (vk *VK) AccountRegisterDevice(params map[string]string) (vkErr Error) {
	_, vkErr = vk.Request("account.registerDevice", params)
	return
}

// AccountSaveProfileInfoResponse struct
type AccountSaveProfileInfoResponse struct {
	Changed     int                `json:"changed"`
	NameRequest accountNameRequest `json:"name_request"`
}

// AccountSaveProfileInfo edits current profile info.
func (vk *VK) AccountSaveProfileInfo(params map[string]string) (response AccountSaveProfileInfoResponse, vkErr Error) {
	rawResponse, vkErr := vk.Request("account.saveProfileInfo", params)
	if vkErr.Code != 0 {
		return
	}

	err := json.Unmarshal(rawResponse, &response)
	if err != nil {
		panic(err)
	}

	return
}

// AccountSetInfo allows to edit the current account info.
func (vk *VK) AccountSetInfo(params map[string]string) (vkErr Error) {
	_, vkErr = vk.Request("account.setInfo", params)
	return
}

// AccountSetNameInMenu sets an application screen name (up to 17 characters), that is shown to the user in the left menu.
func (vk *VK) AccountSetNameInMenu(params map[string]string) (vkErr Error) {
	_, vkErr = vk.Request("account.setNameInMenu", params)
	return
}

// AccountSetOffline marks a current user as offline.
func (vk *VK) AccountSetOffline() (vkErr Error) {
	_, vkErr = vk.Request("account.setOffline", make(map[string]string))
	return
}

// AccountSetOnline marks the current user as online for 5 minutes.
func (vk *VK) AccountSetOnline(params map[string]string) (vkErr Error) {
	_, vkErr = vk.Request("account.setOnline", params)
	return
}

// AccountSetPushSettings change push settings.
func (vk *VK) AccountSetPushSettings(params map[string]string) (vkErr Error) {
	_, vkErr = vk.Request("account.setPushSettings", params)
	return
}

// AccountSetSilenceMode mutes push notifications for the set period of time.
func (vk *VK) AccountSetSilenceMode(params map[string]string) (vkErr Error) {
	_, vkErr = vk.Request("account.setSilenceMode", params)
	return
}

// AccountUnban account.unban
func (vk *VK) AccountUnban(params map[string]string) (vkErr Error) {
	_, vkErr = vk.Request("account.unban", params)
	return
}

// AccountUnregisterDevice unsubscribes a device from push notifications.
func (vk *VK) AccountUnregisterDevice(params map[string]string) (vkErr Error) {
	_, vkErr = vk.Request("account.unregisterDevice", params)
	return
}
