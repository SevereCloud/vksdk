package api_test

import (
	"testing"

	"github.com/SevereCloud/vksdk/api"

	"github.com/stretchr/testify/assert"
)

func TestVK_GroupsAddAddress(t *testing.T) {
	t.Parallel()

	needGroupToken(t)

	res, err := vkGroup.GroupsAddAddress(api.Params{
		"group_id":           vkGroupID,
		"title":              "Точка встречи parkrun",
		"address":            "Сосновкий лесопарк",
		"additional_address": "Парковая дорожка между футбольным полем и спортивной площадкой",
		"country_id":         1,
		"city_id":            2,
		"metro_id":           189,
		"latitude":           60.0179405118554,
		"longitude":          30.365817365050702,
		"phone":              "88005553535",
		"work_info_status":   "always_opened",
	})
	noError(t, err)
	assert.NotEmpty(t, res.ID)
	assert.NotEmpty(t, res.Title)
	assert.NotEmpty(t, res.Address)
	assert.NotEmpty(t, res.AdditionalAddress)
	assert.NotEmpty(t, res.CountryID)
	assert.NotEmpty(t, res.CityID)
	assert.NotEmpty(t, res.MetroStationID)
	assert.NotEmpty(t, res.Latitude)
	assert.NotEmpty(t, res.Longitude)
	assert.NotEmpty(t, res.WorkInfoStatus)
	// assert.NotEmpty(t, res.TimeOffset)
	assert.NotEmpty(t, res.Phone)
	// if assert.NotEmpty(t, res.Timetable) {
	// 	assert.NotEmpty(t, res.Timetable.Sat.OpenTime)
	// 	assert.NotEmpty(t, res.Timetable.Sat.CloseTime)
	// 	assert.NotEmpty(t, res.Timetable.Sat.BreakOpenTime)
	// 	assert.NotEmpty(t, res.Timetable.Sat.BreakCloseTime)
	// }

	res2, err := vkGroup.GroupsEditAddress(api.Params{
		"group_id":           vkGroupID,
		"address_id":         res.ID,
		"title":              "Точка встречи parkrun",
		"address":            "Сосновкий лесопарк",
		"additional_address": "Парковая дорожка между футбольным полем и спортивной площадкой",
		"country_id":         1,
		"city_id":            2,
		"metro_id":           189,
		"latitude":           60.0179405118554,
		"longitude":          30.365817365050702,
		"phone":              "88005553535",
		"work_info_status":   "always_opened",
	})
	noError(t, err)
	assert.NotEmpty(t, res2.ID)
	assert.NotEmpty(t, res2.Title)
	assert.NotEmpty(t, res2.Address)
	assert.NotEmpty(t, res2.AdditionalAddress)
	assert.NotEmpty(t, res2.CountryID)
	assert.NotEmpty(t, res2.CityID)
	assert.NotEmpty(t, res2.MetroStationID)
	assert.NotEmpty(t, res2.Latitude)
	assert.NotEmpty(t, res2.Longitude)
	assert.NotEmpty(t, res2.WorkInfoStatus)
	// assert.NotEmpty(t, res2.TimeOffset)
	assert.NotEmpty(t, res2.Phone)
	// if assert.NotEmpty(t, res2.Timetable) {
	// 	assert.NotEmpty(t, res2.Timetable.Sat.OpenTime)
	// 	assert.NotEmpty(t, res2.Timetable.Sat.CloseTime)
	// 	assert.NotEmpty(t, res2.Timetable.Sat.BreakOpenTime)
	// 	assert.NotEmpty(t, res2.Timetable.Sat.BreakCloseTime)
	// }

	res3, err := vkGroup.GroupsDeleteAddress(api.Params{
		"group_id":   vkGroupID,
		"address_id": res.ID,
	})
	noError(t, err)
	assert.NotEmpty(t, res3)
}

func TestVK_GroupsAddCallbackServer(t *testing.T) {
	t.Parallel()

	needGroupToken(t)

	res, err := vkGroup.GroupsAddCallbackServer(api.Params{
		"group_id": vkGroupID,
		"url":      "https://example.com",
		"title":    "test",
	})
	noError(t, err)
	assert.NotEmpty(t, res.ServerID)

	res2, err := vkGroup.GroupsEditCallbackServer(api.Params{
		"group_id":   vkGroupID,
		"server_id":  res.ServerID,
		"url":        "https://example.com",
		"title":      "test edit",
		"secret_key": "secret_key",
	})
	noError(t, err)
	assert.NotEmpty(t, res2)

	res2, err = vkGroup.GroupsSetCallbackSettings(api.Params{
		"group_id":    vkGroupID,
		"server_id":   res.ServerID,
		"message_new": true,
	})
	noError(t, err)
	assert.NotEmpty(t, res2)

	servers, err := vkGroup.GroupsGetCallbackServers(api.Params{
		"group_id":   vkGroupID,
		"server_ids": res.ServerID,
	})
	noError(t, err)
	assert.NotEmpty(t, servers.Count)
	assert.NotEmpty(t, servers.Items[0].CreatorID)
	assert.NotEmpty(t, servers.Items[0].ID)
	assert.NotEmpty(t, servers.Items[0].Title)
	assert.NotEmpty(t, servers.Items[0].Status)
	assert.NotEmpty(t, servers.Items[0].URL)
	assert.NotEmpty(t, servers.Items[0].SecretKey)

	setting, err := vkGroup.GroupsGetCallbackSettings(api.Params{
		"group_id":  vkGroupID,
		"server_id": res.ServerID,
	})
	noError(t, err)
	assert.NotEmpty(t, setting.APIVersion)
	// assert.NotEmpty(t, setting.Events)

	code, err := vkGroup.GroupsGetCallbackConfirmationCode(api.Params{
		"group_id":  vkGroupID,
		"server_id": res.ServerID,
	})
	noError(t, err)
	assert.NotEmpty(t, code.Code)
	// assert.NotEmpty(t, setting.Events)

	res2, err = vkGroup.GroupsDeleteCallbackServer(api.Params{
		"group_id":  vkGroupID,
		"server_id": res.ServerID,
	})
	noError(t, err)
	assert.NotEmpty(t, res2)
}

func TestVK_GroupsAddLink(t *testing.T) {
	t.Parallel()

	needUserToken(t)
	needGroupToken(t)

	link, err := vkUser.GroupsAddLink(api.Params{
		"group_id": vkGroupID,
		"link":     "https://yandex.ru",
		"text":     "test",
	})
	noError(t, err)
	assert.NotEmpty(t, link.ID)
	assert.NotEmpty(t, link.URL)
	assert.NotEmpty(t, link.Name)
	assert.NotEmpty(t, link.Desc)

	_, err = vkUser.GroupsEditLink(api.Params{
		"group_id": vkGroupID,
		"link_id":  link.ID,
		"text":     "test edit",
	})
	noError(t, err)

	_, err = vkUser.GroupsReorderLink(api.Params{
		"group_id": vkGroupID,
		"link_id":  link.ID,
		"after":    0,
	})
	noError(t, err)

	_, err = vkUser.GroupsDeleteLink(api.Params{
		"group_id": vkGroupID,
		"link_id":  link.ID,
	})
	noError(t, err)
}

func TestVK_GroupsApproveRequest(t *testing.T) {
	// TODO: Add test cases.
}

func TestVK_GroupsBan(t *testing.T) {
	t.Parallel()

	needUserToken(t)
	needGroupToken(t)

	_, err := vkUser.GroupsBan(api.Params{
		"group_id": vkGroupID,
		"owner_id": 1,
		"reason":   2,
		"comment":  "test",
	})
	noError(t, err)

	res, err := vkGroup.GroupsGetBanned(api.Params{
		"group_id": vkGroupID,
	})
	noError(t, err)
	assert.NotEmpty(t, res.Count)

	if assert.NotEmpty(t, res.Items) {
		assert.NotEmpty(t, res.Items[0].Type)
		assert.NotEmpty(t, res.Items[0].BanInfo.AdminID)
		assert.NotEmpty(t, res.Items[0].BanInfo.Date)
	}

	_, err = vkUser.GroupsUnban(api.Params{
		"group_id": vkGroupID,
		"owner_id": 1,
	})
	noError(t, err)
}

func TestVK_GroupsCreate(t *testing.T) {
	t.Parallel()

	needUserToken(t)

	group, err := vkUser.GroupsCreate(api.Params{
		"title":       "My Community",
		"description": "Community description",
		"type":        "group",
	})
	noError(t, err)
	assert.NotEmpty(t, group.ID)
	// assert.NotEmpty(t, group.Name)
	// BUG (VK): https://vk.com/bug201823
	assert.NotEmpty(t, group.AdminLevel)
	assert.NotEmpty(t, group.IsAdmin)
	assert.NotEmpty(t, group.IsMember)
	assert.NotEmpty(t, group.Photo100)
	assert.NotEmpty(t, group.Photo200)
	assert.NotEmpty(t, group.Photo50)
	assert.NotEmpty(t, group.ScreenName)
	assert.NotEmpty(t, group.Type)

	_, err = vkUser.GroupsEdit(api.Params{
		"group_id":    group.ID,
		"description": "Community description edit",
	})
	noError(t, err)

	_, err = vkUser.GroupsSetSettings(api.Params{
		"group_id":          group.ID,
		"messages":          true,
		"bots_capabilities": true,
		"bots_start_button": true,
		"bots_add_to_chat":  true,
	})
	noError(t, err)

	_, err = vkUser.GroupsSetLongPollSettings(api.Params{
		"group_id":    group.ID,
		"enabled":     true,
		"api_version": api.Version,
	})
	noError(t, err)
}

func TestVK_GroupsEnableOnline(t *testing.T) {
	t.Parallel()

	needGroupToken(t)

	res, err := vkGroup.GroupsEnableOnline(api.Params{
		"group_id": vkGroupID,
	})
	noError(t, err)
	assert.NotEmpty(t, res)

	res, err = vkGroup.GroupsDisableOnline(api.Params{
		"group_id": vkGroupID,
	})
	noError(t, err)
	assert.NotEmpty(t, res)
}

func TestVK_GroupsEditManager(t *testing.T) {
	// TODO: Add test cases.
}

func TestVK_GroupsGet(t *testing.T) {
	t.Parallel()

	needUserToken(t)

	res, err := vkUser.GroupsGet(api.Params{
		"user_id": 117253521,
		"fields":  "market,member_status,is_favorite,is_subscribed,city,country,verified,description,wiki_page,members_count,counters,cover,can_post,can_see_all_posts,activity,fixed_post,can_create_topic,can_upload_video,has_photo,status,main_album_id,links,contacts,site,main_section,trending,can_message,is_market_cart_enabled,is_messages_blocked,can_send_notify,online_status,start_date,finish_date,age_limits,ban_info,action_button,author_id,phone,has_market_app,addresses,live_covers,is_adult,can_subscribe_posts,warning_notification,can_upload_doc,crop_photo,is_hidden_from_feed,place,public_date_label,wall",
	})
	noError(t, err)
	assert.NotEmpty(t, res.Count)
	assert.NotEmpty(t, res.Items)
}

func TestVK_GroupsGetExtended(t *testing.T) {
	t.Parallel()

	needUserToken(t)

	res, err := vkUser.GroupsGetExtended(api.Params{
		"user_id": 117253521,
		"fields":  "market,member_status,is_favorite,is_subscribed,city,country,verified,description,wiki_page,members_count,counters,cover,can_post,can_see_all_posts,activity,fixed_post,can_create_topic,can_upload_video,has_photo,status,main_album_id,links,contacts,site,main_section,trending,can_message,is_market_cart_enabled,is_messages_blocked,can_send_notify,online_status,start_date,finish_date,age_limits,ban_info,action_button,author_id,phone,has_market_app,addresses,live_covers,is_adult,can_subscribe_posts,warning_notification,can_upload_doc,crop_photo,is_hidden_from_feed,place,public_date_label,wall",
	})
	noError(t, err)
	assert.NotEmpty(t, res.Count)
	assert.NotEmpty(t, res.Items)
}

func TestVK_GroupsGetAddresses(t *testing.T) {
	t.Parallel()

	needServiceToken(t)

	res, err := vkService.GroupsGetAddresses(api.Params{
		"group_id": 167450351,
		"count":    1,
	})
	noError(t, err)
	assert.NotEmpty(t, res.Count)
	assert.NotEmpty(t, res.Items)

	for _, address := range res.Items {
		assert.NotEmpty(t, address.ID)
		assert.NotEmpty(t, address.Title)
		assert.NotEmpty(t, address.Address)
		// assert.NotEmpty(t, address.AdditionalAddress)
		assert.NotEmpty(t, address.CountryID)
		assert.NotEmpty(t, address.CityID)
		// assert.NotEmpty(t, address.MetroStationID)
		assert.NotEmpty(t, address.Latitude)
		assert.NotEmpty(t, address.Longitude)
		assert.NotEmpty(t, address.WorkInfoStatus)
		// assert.NotEmpty(t, address.TimeOffset)
		assert.NotEmpty(t, address.Phone)

		if assert.NotEmpty(t, address.Timetable) {
			assert.NotEmpty(t, address.Timetable.Sat.OpenTime)
			// assert.NotEmpty(t, address.Timetable.Sat.BreakOpenTime)
			// assert.NotEmpty(t, address.Timetable.Sat.BreakCloseTime)
			assert.NotEmpty(t, address.Timetable.Sat.CloseTime)
		}
	}
}

func TestVK_GroupsGetByID(t *testing.T) {
	t.Parallel()

	needGroupToken(t)

	res, err := vkGroup.GroupsGetByID(api.Params{
		"group_ids": "apiclub",
		"fields":    "market,member_status,is_favorite,is_subscribed,city,country,verified,description,wiki_page,members_count,counters,cover,can_post,can_see_all_posts,activity,fixed_post,can_create_topic,can_upload_video,has_photo,status,main_album_id,links,contacts,site,main_section,trending,can_message,is_market_cart_enabled,is_messages_blocked,can_send_notify,online_status,start_date,finish_date,age_limits,ban_info,action_button,author_id,phone,has_market_app,addresses,live_covers,is_adult,can_subscribe_posts,warning_notification,can_upload_doc,crop_photo,is_hidden_from_feed,place,public_date_label,wall",
	})

	noError(t, err)

	if assert.NotEmpty(t, res[0]) {
		group := res[0]
		assert.NotEmpty(t, group.ID)
		assert.NotEmpty(t, group.Name)
		assert.NotEmpty(t, group.ScreenName)
		// assert.NotEmpty(t, group.IsClosed)
		assert.NotEmpty(t, group.Type)
		// assert.NotEmpty(t, group.IsAdmin)
		// assert.NotEmpty(t, group.IsMember)
		// assert.NotEmpty(t, group.IsAdvertiser)
		// assert.NotEmpty(t, group.Market)
		// assert.NotEmpty(t, group.MemberStatus)
		// assert.NotEmpty(t, group.IsSubscribed)
		assert.NotEmpty(t, group.City)
		assert.NotEmpty(t, group.Country)
		assert.NotEmpty(t, group.Verified)
		assert.NotEmpty(t, group.Description)
		assert.NotEmpty(t, group.WikiPage)
		assert.NotEmpty(t, group.MembersCount)
		assert.NotEmpty(t, group.Counters)
		assert.NotEmpty(t, group.Counters.Topics)
		assert.NotEmpty(t, group.Counters.Videos)
		assert.NotEmpty(t, group.Counters.Addresses)
		// assert.NotEmpty(t, group.CanPost)
		// assert.NotEmpty(t, group.CanSeeAllPosts)
		assert.NotEmpty(t, group.Activity)
		// assert.NotEmpty(t, group.CanCreateTopic)
		// assert.NotEmpty(t, group.CanUploadVideo)
		assert.NotEmpty(t, group.Status)

		if assert.NotEmpty(t, group.Links) {
			assert.NotEmpty(t, group.Links[0].ID)
			assert.NotEmpty(t, group.Links[0].URL)
			assert.NotEmpty(t, group.Links[0].Name)
			assert.NotEmpty(t, group.Links[0].Desc)
			assert.NotEmpty(t, group.Links[0].Photo50)
			assert.NotEmpty(t, group.Links[0].Photo100)
		}
		// assert.NotEmpty(t, group.Contacts)
		assert.NotEmpty(t, group.Site)
		// assert.NotEmpty(t, group.MainSection)
		// assert.NotEmpty(t, group.Trending)
		// assert.NotEmpty(t, group.IsMessagesBlocked)
		assert.NotEmpty(t, group.OnlineStatus)
		assert.NotEmpty(t, group.AgeLimits)
		// assert.NotEmpty(t, group.HasMarketApp)
		if assert.NotEmpty(t, group.Addresses) {
			assert.NotEmpty(t, group.Addresses.IsEnabled)
			assert.NotEmpty(t, group.Addresses.MainAddressID)
		}
		// assert.NotEmpty(t, group.LiveCovers)
		// assert.NotEmpty(t, group.IsAdult)
		// assert.NotEmpty(t, group.CanSubscribePosts)
		// assert.NotEmpty(t, group.CanUploadDoc)
		assert.NotEmpty(t, group.CropPhoto)
		// assert.NotEmpty(t, group.IsHiddenFromFeed)
		assert.NotEmpty(t, group.Wall)

		if assert.NotEmpty(t, group.Cover) {
			assert.NotEmpty(t, group.Cover.Enabled)
			assert.NotEmpty(t, group.Cover.Images)
		}

		assert.NotEmpty(t, group.HasPhoto)
		// assert.NotEmpty(t, group.CanMessage)
		assert.NotEmpty(t, group.Photo50)
		assert.NotEmpty(t, group.Photo100)
		assert.NotEmpty(t, group.Photo200)
	}
}

func TestVK_GroupsGetByID_private(t *testing.T) {
	t.Parallel()

	needGroupToken(t)

	res, err := vkGroup.GroupsGetByID(api.Params{
		"group_ids": "184580855",
		"fields":    "market,member_status,is_favorite,is_subscribed,city,country,verified,description,wiki_page,members_count,counters,cover,can_post,can_see_all_posts,activity,fixed_post,can_create_topic,can_upload_video,has_photo,status,main_album_id,links,contacts,site,main_section,trending,can_message,is_market_cart_enabled,is_messages_blocked,can_send_notify,online_status,start_date,finish_date,age_limits,ban_info,action_button,author_id,phone,has_market_app,addresses,live_covers,is_adult,can_subscribe_posts,warning_notification,can_upload_doc,crop_photo,is_hidden_from_feed,place,public_date_label,wall",
	})

	noError(t, err)

	if assert.NotEmpty(t, res[0]) {
		group := res[0]
		assert.NotEmpty(t, group.ID)
		assert.NotEmpty(t, group.Name)
		assert.NotEmpty(t, group.ScreenName)
		assert.NotEmpty(t, group.IsClosed)
		assert.NotEmpty(t, group.Type)
		assert.NotEmpty(t, group.Photo50)
		assert.NotEmpty(t, group.Photo100)
		assert.NotEmpty(t, group.Photo200)
	}
}

func TestVK_GroupsGetCatalog(t *testing.T) {
	t.Parallel()

	needUserToken(t)

	res, err := vkUser.GroupsGetCatalog(api.Params{
		"category_id": 0,
	})
	noError(t, err)
	assert.NotEmpty(t, res.Count)
	assert.NotEmpty(t, res.Items)
}

func TestVK_GroupsGetCatalogInfo(t *testing.T) {
	t.Parallel()

	needUserToken(t)

	res, err := vkUser.GroupsGetCatalogInfo(api.Params{})
	noError(t, err)
	// assert.NotEmpty(t, res.Categories[0].ID)
	assert.NotEmpty(t, res.Enabled)
}

func TestVK_GroupsGetCatalogInfoExtended(t *testing.T) {
	t.Parallel()

	needUserToken(t)

	res, err := vkUser.GroupsGetCatalogInfoExtended(api.Params{})
	noError(t, err)
	// if assert.NotEmpty(t, res.Categories) {
	// 	assert.NotEmpty(t, res.Categories[0].Name)
	// 	assert.NotEmpty(t, res.Categories[0].PageCount)
	// 	assert.NotEmpty(t, res.Categories[0].PagePreviews)
	// }
	assert.NotEmpty(t, res.Enabled)
}

func TestVK_GroupsGetInvitedUsers(t *testing.T) {
	t.Parallel()

	needGroupToken(t)
	needUserToken(t)

	_, err := vkUser.GroupsGetInvitedUsers(api.Params{
		"group_id": vkGroupID,
	})
	noError(t, err)
}

func TestVK_GroupsGetInvites(t *testing.T) {
	t.Parallel()

	needUserToken(t)

	_, err := vkUser.GroupsGetInvites(api.Params{})
	noError(t, err)
}

func TestVK_GroupsGetInvitesExtended(t *testing.T) {
	t.Parallel()

	needUserToken(t)

	_, err := vkUser.GroupsGetInvitesExtended(api.Params{})
	noError(t, err)
}

func TestVK_GroupsGetLongPollServer(t *testing.T) {
	t.Parallel()

	needGroupToken(t)

	res, err := vkGroup.GroupsGetLongPollServer(api.Params{
		"group_id": vkGroupID,
	})
	noError(t, err)
	assert.NotEmpty(t, res.Key)
	assert.NotEmpty(t, res.Server)
	assert.NotEmpty(t, res.Ts)
}

func TestVK_GroupsGetLongPollSettings(t *testing.T) {
	t.Parallel()

	needGroupToken(t)

	res, err := vkGroup.GroupsGetLongPollSettings(api.Params{
		"group_id": vkGroupID,
	})
	noError(t, err)
	assert.NotEmpty(t, res.Events)
	assert.NotEmpty(t, res.APIVersion)
}

func TestVK_GroupsGetMembers(t *testing.T) {
	t.Parallel()

	needServiceToken(t)

	res, err := vkService.GroupsGetMembers(api.Params{
		"group_id": 1,
	})
	noError(t, err)
	assert.NotEmpty(t, res.Count)
	assert.NotEmpty(t, res.Items)
}

func TestVK_GroupsGetMembersFields(t *testing.T) {
	t.Parallel()

	needServiceToken(t)

	res, err := vkService.GroupsGetMembersFields(api.Params{
		"group_id": 1,
	})
	noError(t, err)
	assert.NotEmpty(t, res.Count)
	assert.NotEmpty(t, res.Items)
}

func TestVK_GroupsGetMembersFilterManagers(t *testing.T) {
	t.Parallel()

	needUserToken(t)
	needGroupToken(t)

	res, err := vkGroup.GroupsGetMembersFilterManagers(api.Params{
		"group_id": vkGroupID,
	})
	noError(t, err)
	assert.NotEmpty(t, res.Count)

	if assert.NotEmpty(t, res.Items) {
		assert.NotEmpty(t, res.Items[0].Role)
	}
}

func TestVK_GroupsGetOnlineStatus(t *testing.T) {
	t.Parallel()

	needGroupToken(t)

	res, err := vkGroup.GroupsGetOnlineStatus(api.Params{
		"group_id": 126054662,
	})
	noError(t, err)
	assert.NotEmpty(t, res.Status)
}

func TestVK_GroupsGetRequests(t *testing.T) {
	t.Parallel()

	needUserToken(t)
	needGroupToken(t)

	_, err := vkUser.GroupsGetRequests(api.Params{
		"group_id": vkGroupID,
	})
	noError(t, err)
}

func TestVK_GroupsGetRequestsFields(t *testing.T) {
	t.Parallel()

	needUserToken(t)
	needGroupToken(t)

	_, err := vkUser.GroupsGetRequestsFields(api.Params{
		"group_id": vkGroupID,
	})
	noError(t, err)
}

func TestVK_GroupsGetSettings(t *testing.T) {
	t.Parallel()

	needUserToken(t)
	needGroupToken(t)

	res, err := vkUser.GroupsGetSettings(api.Params{
		"group_id": vkGroupID,
	})
	noError(t, err)
	assert.NotEmpty(t, res.Title)
}

func TestVK_GroupsGetTokenPermissions(t *testing.T) {
	t.Parallel()

	needGroupToken(t)

	res, err := vkGroup.GroupsGetTokenPermissions(api.Params{})
	noError(t, err)
	assert.NotEmpty(t, res.Mask)
	assert.NotEmpty(t, res.Permissions[0].Name)
	assert.NotEmpty(t, res.Permissions[0].Setting)
}

func TestVK_GroupsInvite(t *testing.T) {
	// TODO: Add test cases.
}

func TestVK_GroupsIsMember(t *testing.T) {
	t.Parallel()

	needGroupToken(t)

	res, err := vkGroup.GroupsIsMember(api.Params{
		"group_id": 1,
		"user_id":  117253521,
	})
	noError(t, err)
	assert.NotEmpty(t, res)
}

func TestVK_GroupsIsMemberExtended(t *testing.T) {
	t.Parallel()

	needGroupToken(t)

	res, err := vkGroup.GroupsIsMemberExtended(api.Params{
		"group_id": 1,
		"user_id":  117253521,
	})
	noError(t, err)
	assert.NotEmpty(t, res.Member)
}

func TestVK_GroupsIsMemberUserIDsExtended(t *testing.T) {
	t.Parallel()

	needGroupToken(t)

	res, err := vkGroup.GroupsIsMemberUserIDsExtended(api.Params{
		"group_id": 1,
		"user_ids": 117253521,
	})
	noError(t, err)

	if assert.NotEmpty(t, res) {
		assert.NotEmpty(t, res[0].Member)
		assert.NotEmpty(t, res[0].UserID)
	}
}

func TestVK_GroupsIsMemberUserIDs(t *testing.T) {
	t.Parallel()

	needGroupToken(t)

	res, err := vkGroup.GroupsIsMemberUserIDs(api.Params{
		"group_id": 1,
		"user_ids": 117253521,
	})
	noError(t, err)

	if assert.NotEmpty(t, res) {
		assert.NotEmpty(t, res[0].Member)
		assert.NotEmpty(t, res[0].UserID)
	}
}

func TestVK_GroupsJoin(t *testing.T) {
	t.Parallel()

	needUserToken(t)

	res, err := vkUser.GroupsJoin(api.Params{
		"group_id": 1,
	})
	noError(t, err)
	assert.NotEmpty(t, res)
}

func TestVK_GroupsLeave(t *testing.T) {
	t.Parallel()

	needUserToken(t)

	res, err := vkUser.GroupsLeave(api.Params{
		"group_id": 1,
	})
	noError(t, err)
	assert.NotEmpty(t, res)
}

func TestVK_GroupsRemoveUser(t *testing.T) {
	// TODO: Add test cases.
}

func TestVK_GroupsSearch(t *testing.T) {
	t.Parallel()

	needUserToken(t)

	res, err := vkUser.GroupsSearch(api.Params{
		"q": "Golang",
	})
	noError(t, err)
	assert.NotEmpty(t, res.Count)
	assert.NotEmpty(t, res.Items)
}
