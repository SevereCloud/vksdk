package vksdk

import "encoding/json"

// AccountBan account.ban
func (vk *VK) AccountBan(params map[string]string) (err error) {
	_, err = vk.Request("account.ban", params)

	return
}

// AccountChangePasswordResponse struct
type AccountChangePasswordResponse struct {
	Token string `json:"token"`
}

// AccountChangePassword changes a user password after access is successfully restored with the auth.restore method.
func (vk *VK) AccountChangePassword(params map[string]string) (response AccountChangePasswordResponse, err error) {
	rawResponse, err := vk.Request("account.changePassword", params)
	if err != nil {
		return
	}

	err = json.Unmarshal(rawResponse, &response)
	if err != nil {
		panic(err)
	}

	return
}

// AccountGetActiveOffersResponse struct
type AccountGetActiveOffersResponse struct {
	Count int            `json:"count,omitempty"`
	Items []accountOffer `json:"items,omitempty"`
}

// AccountGetActiveOffers returns a list of active ads (offers).If the user fulfill their conditions, he will be able to get the appropriate number of votes to his balance.
func (vk *VK) AccountGetActiveOffers(params map[string]string) (response AccountGetActiveOffersResponse, err error) {
	rawResponse, err := vk.Request("account.getActiveOffers", params)
	if err != nil {
		return
	}

	err = json.Unmarshal(rawResponse, &response)
	if err != nil {
		panic(err)
	}

	return
}

// AccountGetAppPermissions gets settings of the user in this application.
func (vk *VK) AccountGetAppPermissions(params map[string]string) (response int, err error) {
	rawResponse, err := vk.Request("account.getAppPermissions", params)
	if err != nil {
		return
	}

	err = json.Unmarshal(rawResponse, &response)
	if err != nil {
		panic(err)
	}

	return
}

// AccountGetBannedResponse struct
type AccountGetBannedResponse struct {
	Count    int           `json:"count,omitempty"`
	Items    []int         `json:"items,omitempty"`
	Profiles []usersUser   `json:"profiles,omitempty"`
	Groups   []groupsGroup `json:"groups,omitempty"`
}

// AccountGetBanned returns a user's blacklist.
func (vk *VK) AccountGetBanned(params map[string]string) (response AccountGetBannedResponse, err error) {
	rawResponse, err := vk.Request("account.getBanned", params)
	if err != nil {
		return
	}

	err = json.Unmarshal(rawResponse, &response)
	if err != nil {
		panic(err)
	}

	return
}

// AccountGetCountersResponse struct
type AccountGetCountersResponse struct {
	AppRequests        int `json:"app_requests,omitempty"`
	Events             int `json:"events,omitempty"`
	Friends            int `json:"friends,omitempty"`
	FriendsSuggestions int `json:"friends_suggestions,omitempty"`
	Gifts              int `json:"gifts,omitempty"`
	Groups             int `json:"groups,omitempty"`
	Messages           int `json:"messages,omitempty"`
	Notifications      int `json:"notifications,omitempty"`
	Photos             int `json:"photos,omitempty"`
	SDK                int `json:"sdk,omitempty"`
	Videos             int `json:"videos,omitempty"`
}

// AccountGetCounters returns non-null values of user counters.
func (vk *VK) AccountGetCounters(params map[string]string) (response AccountGetCountersResponse, err error) {
	rawResponse, err := vk.Request("account.getCounters", params)
	if err != nil {
		return
	}

	err = json.Unmarshal(rawResponse, &response)
	if err != nil {
		panic(err)
	}

	return
}

// AccountGetInfoResponse struct
type AccountGetInfoResponse struct {
	Country           string `json:"country,omitempty"`
	HTTPSRequired     int    `json:"https_required,omitempty"`
	Intro             int    `json:"intro,omitempty"`
	Lang              int    `json:"lang,omitempty"`
	NoWallReplies     int    `json:"no_wall_replies,omitempty"`
	OwnPostsDefault   int    `json:"own_posts_default,omitempty"`
	TwoFactorRequired int    `json:"2fa_required,omitempty"`
}

// AccountGetInfo returns current account info.
func (vk *VK) AccountGetInfo(params map[string]string) (response AccountGetInfoResponse, err error) {
	rawResponse, err := vk.Request("account.getInfo", params)
	if err != nil {
		return
	}

	err = json.Unmarshal(rawResponse, &response)
	if err != nil {
		panic(err)
	}

	return
}

// AccountGetProfileInfoResponse struct
type AccountGetProfileInfoResponse struct {
	Bdate            string             `json:"bdate,omitempty"`
	BdateVisibility  int                `json:"bdate_visibility,omitempty"`
	City             baseObject         `json:"city,omitempty"`
	Country          baseCountry        `json:"country,omitempty"`
	FirstName        string             `json:"first_name,omitempty"`
	HomeTown         string             `json:"home_town,omitempty"`
	LastName         string             `json:"last_name,omitempty"`
	MaidenName       string             `json:"maiden_name,omitempty"`
	NameRequest      accountNameRequest `json:"name_request,omitempty"`
	Phone            string             `json:"phone,omitempty"`
	Relation         int                `json:"relation,omitempty"`
	RelationPartner  usersUserMin       `json:"relation_partner,omitempty"`
	RelationPending  int                `json:"relation_pending,omitempty"`
	RelationRequests []usersUserMin     `json:"relation_requests,omitempty"`
	ScreenName       string             `json:"screen_name,omitempty"`
	Sex              baseSex            `json:"sex,omitempty"`
	Status           string             `json:"status,omitempty"`
}

// AccountGetProfileInfo returns the current account info.
func (vk *VK) AccountGetProfileInfo() (response AccountGetProfileInfoResponse, err error) {
	rawResponse, err := vk.Request("account.getProfileInfo", make(map[string]string))
	if err != nil {
		return
	}

	err = json.Unmarshal(rawResponse, &response)
	if err != nil {
		panic(err)
	}

	return
}

// AccountGetPushSettingsResponse struct
type AccountGetPushSettingsResponse struct {
	Conversations accountPushConversations `json:"conversations,omitempty"`
	Disabled      int                      `json:"disabled,omitempty"`
	DisabledUntil int                      `json:"disabled_until,omitempty"`
	Settings      accountPushParams        `json:"settings,omitempty"`
}

// AccountGetPushSettings account.getPushSettings Gets settings of push notifications.
func (vk *VK) AccountGetPushSettings(params map[string]string) (response AccountGetPushSettingsResponse, err error) {
	rawResponse, err := vk.Request("account.getPushSettings", params)
	if err != nil {
		return
	}

	err = json.Unmarshal(rawResponse, &response)
	if err != nil {
		panic(err)
	}

	return
}

// AccountRegisterDevice subscribes an iOS/Android/Windows/Mac based device to receive push notifications
func (vk *VK) AccountRegisterDevice(params map[string]string) (err error) {
	_, err = vk.Request("account.registerDevice", params)

	return
}

// AccountSaveProfileInfoResponse struct
type AccountSaveProfileInfoResponse struct {
	Changed     int                `json:"changed,omitempty"`
	NameRequest accountNameRequest `json:"name_request,omitempty"`
}

// AccountSaveProfileInfo edits current profile info.
func (vk *VK) AccountSaveProfileInfo(params map[string]string) (response AccountSaveProfileInfoResponse, err error) {
	rawResponse, err := vk.Request("account.saveProfileInfo", params)
	if err != nil {
		return
	}

	err = json.Unmarshal(rawResponse, &response)
	if err != nil {
		panic(err)
	}

	return
}

// AccountSetInfo allows to edit the current account info.
func (vk *VK) AccountSetInfo(params map[string]string) (err error) {
	_, err = vk.Request("account.setInfo", params)

	return
}

// AccountSetNameInMenu sets an application screen name (up to 17 characters), that is shown to the user in the left menu.
func (vk *VK) AccountSetNameInMenu(params map[string]string) (err error) {
	_, err = vk.Request("account.setNameInMenu", params)

	return
}

// AccountSetOffline marks a current user as offline.
func (vk *VK) AccountSetOffline() (err error) {
	_, err = vk.Request("account.setOffline", make(map[string]string))

	return
}

// AccountSetOnline marks the current user as online for 5 minutes.
func (vk *VK) AccountSetOnline(params map[string]string) (err error) {
	_, err = vk.Request("account.setOnline", params)

	return
}

// AccountSetPushSettings change push settings.
func (vk *VK) AccountSetPushSettings(params map[string]string) (err error) {
	_, err = vk.Request("account.setPushSettings", params)

	return
}

// AccountSetSilenceMode mutes push notifications for the set period of time.
func (vk *VK) AccountSetSilenceMode(params map[string]string) (err error) {
	_, err = vk.Request("account.setSilenceMode", params)

	return
}

// AccountUnban account.unban
func (vk *VK) AccountUnban(params map[string]string) (err error) {
	_, err = vk.Request("account.unban", params)

	return
}

// AccountUnregisterDevice unsubscribes a device from push notifications.
func (vk *VK) AccountUnregisterDevice(params map[string]string) (err error) {
	_, err = vk.Request("account.unregisterDevice", params)

	return
}
