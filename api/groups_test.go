package api

import (
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestVK_GroupsAddAddress(t *testing.T) {
	needGroupToken(t)

	res, err := vkGroup.GroupsAddAddress(map[string]string{
		"group_id":           strconv.Itoa(vkGroupID),
		"title":              "Точка встречи parkrun",
		"address":            "Сосновкий лесопарк",
		"additional_address": "Парковая дорожка между футбольным полем и спортивной площадкой",
		"country_id":         "1",
		"city_id":            "2",
		"metro_id":           "189",
		"latitude":           "60.0179405118554",
		"longitude":          "30.365817365050702",
		"phone":              "88005553535",
		"work_info_status":   "always_opened",
	})
	assert.NoError(t, err)
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
	assert.NotEmpty(t, res.TimeOffset)
	assert.NotEmpty(t, res.Phone)
	// if assert.NotEmpty(t, res.Timetable) {
	// 	assert.NotEmpty(t, res.Timetable.Sat.OpenTime)
	// 	assert.NotEmpty(t, res.Timetable.Sat.CloseTime)
	// 	assert.NotEmpty(t, res.Timetable.Sat.BreakOpenTime)
	// 	assert.NotEmpty(t, res.Timetable.Sat.BreakCloseTime)
	// }

	res2, err := vkGroup.GroupsEditAddress(map[string]string{
		"group_id":            strconv.Itoa(vkGroupID),
		"address_id":          strconv.Itoa(res.ID),
		"title":               "Точка встречи parkrun",
		"addres2s":            "Сосновкий лесопарк",
		"additional_addres2s": "Парковая дорожка между футбольным полем и спортивной площадкой",
		"country_id":          "1",
		"city_id":             "2",
		"metro_id":            "189",
		"latitude":            "60.0179405118554",
		"longitude":           "30.365817365050702",
		"phone":               "88005553535",
		"work_info_status":    "always_opened",
	})
	assert.NoError(t, err)
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
	assert.NotEmpty(t, res2.TimeOffset)
	assert.NotEmpty(t, res2.Phone)
	// if assert.NotEmpty(t, res2.Timetable) {
	// 	assert.NotEmpty(t, res2.Timetable.Sat.OpenTime)
	// 	assert.NotEmpty(t, res2.Timetable.Sat.CloseTime)
	// 	assert.NotEmpty(t, res2.Timetable.Sat.BreakOpenTime)
	// 	assert.NotEmpty(t, res2.Timetable.Sat.BreakCloseTime)
	// }

	res3, err := vkGroup.GroupsDeleteAddress(map[string]string{
		"group_id":   strconv.Itoa(vkGroupID),
		"address_id": strconv.Itoa(res.ID),
	})
	assert.NoError(t, err)
	assert.NotEmpty(t, res3)
}

func TestVK_GroupsAddCallbackServer(t *testing.T) {
	needGroupToken(t)

	res, err := vkGroup.GroupsAddCallbackServer(map[string]string{
		"group_id": strconv.Itoa(vkGroupID),
		"url":      "https://example.com",
		"title":    "test",
	})
	assert.NoError(t, err)
	assert.NotEmpty(t, res.ServerID)

	res2, err := vkGroup.GroupsEditCallbackServer(map[string]string{
		"group_id":   strconv.Itoa(vkGroupID),
		"server_id":  strconv.Itoa(res.ServerID),
		"url":        "https://example.com",
		"title":      "test edit",
		"secret_key": "secret_key",
	})
	assert.NoError(t, err)
	assert.NotEmpty(t, res2)

	servers, err := vkGroup.GroupsGetCallbackServers(map[string]string{
		"group_id":   strconv.Itoa(vkGroupID),
		"server_ids": strconv.Itoa(res.ServerID),
	})
	assert.NoError(t, err)
	assert.NotEmpty(t, servers.Count)
	assert.NotEmpty(t, servers.Items[0].CreatorID)
	assert.NotEmpty(t, servers.Items[0].ID)
	assert.NotEmpty(t, servers.Items[0].Title)
	assert.NotEmpty(t, servers.Items[0].Status)
	assert.NotEmpty(t, servers.Items[0].URL)
	assert.NotEmpty(t, servers.Items[0].SecretKey)

	res2, err = vkGroup.GroupsDeleteCallbackServer(map[string]string{
		"group_id":  strconv.Itoa(vkGroupID),
		"server_id": strconv.Itoa(res.ServerID),
	})
	assert.NoError(t, err)
	assert.NotEmpty(t, res2)
}

func TestVK_GroupsAddLink(t *testing.T) {
	// TODO: Add test cases.
}

func TestVK_GroupsApproveRequest(t *testing.T) {
	// TODO: Add test cases.
}

func TestVK_GroupsBan(t *testing.T) {
	// TODO: Add test cases.
}

func TestVK_GroupsCreate(t *testing.T) {
	// TODO: Add test cases.
}

func TestVK_GroupsDeleteLink(t *testing.T) {
	// TODO: Add test cases.
}

func TestVK_GroupsDisableOnline(t *testing.T) {
	// TODO: Add test cases.
}

func TestVK_GroupsEdit(t *testing.T) {
	// TODO: Add test cases.
}

func TestVK_GroupsEditLink(t *testing.T) {
	// TODO: Add test cases.
}

func TestVK_GroupsEditManager(t *testing.T) {
	// TODO: Add test cases.
}

func TestVK_GroupsEnableOnline(t *testing.T) {
	// TODO: Add test cases.
}

func TestVK_GroupsGet(t *testing.T) {
	// TODO: Add test cases.
}

func TestVK_GroupsGetExtended(t *testing.T) {
	// TODO: Add test cases.
}

func TestVK_GroupsGetAddresses(t *testing.T) {
	needServiceToken(t)

	res, err := vkService.GroupsGetAddresses(map[string]string{
		"group_id": "167450351",
		"count":    "1",
	})
	assert.NoError(t, err)
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
		assert.NotEmpty(t, address.TimeOffset)
		assert.NotEmpty(t, address.Phone)
		if assert.NotEmpty(t, address.Timetable) {
			assert.NotEmpty(t, address.Timetable.Sat.OpenTime)
			assert.NotEmpty(t, address.Timetable.Sat.CloseTime)
			// assert.NotEmpty(t, address.Timetable.Sat.BreakOpenTime)
			// assert.NotEmpty(t, address.Timetable.Sat.BreakCloseTime)
		}
	}
}

func TestVK_GroupsGetBanned(t *testing.T) {
	// TODO: Add test cases.
}

func TestVK_GroupsGetByID(t *testing.T) {
	// TODO: Add test cases.
}

func TestVK_GroupsGetCallbackConfirmationCode(t *testing.T) {
	// TODO: Add test cases.
}

func TestVK_GroupsGetCallbackSettings(t *testing.T) {
	// TODO: Add test cases.
}

func TestVK_GroupsGetCatalog(t *testing.T) {
	// TODO: Add test cases.
}

func TestVK_GroupsGetCatalogInfo(t *testing.T) {
	// TODO: Add test cases.
}

func TestVK_GroupsGetCatalogInfoExtended(t *testing.T) {
	// TODO: Add test cases.
}

func TestVK_GroupsGetInvitedUsers(t *testing.T) {
	// TODO: Add test cases.
}

func TestVK_GroupsGetInvites(t *testing.T) {
	// TODO: Add test cases.
}

func TestVK_GroupsGetInvitesExtended(t *testing.T) {
	// TODO: Add test cases.
}

func TestVK_GroupsGetLongPollServer(t *testing.T) {
	// TODO: Add test cases.
}

func TestVK_GroupsGetLongPollSettings(t *testing.T) {
	// TODO: Add test cases.
}

func TestVK_GroupsGetMembers(t *testing.T) {
	// TODO: Add test cases.
}

func TestVK_GroupsGetMembersFields(t *testing.T) {
	// TODO: Add test cases.
}

func TestVK_GroupsGetMembersFilterManagers(t *testing.T) {
	// TODO: Add test cases.
}

func TestVK_GroupsGetOnlineStatus(t *testing.T) {
	// TODO: Add test cases.
}

func TestVK_GroupsGetRequests(t *testing.T) {
	// TODO: Add test cases.
}

func TestVK_GroupsGetSettings(t *testing.T) {
	// TODO: Add test cases.
}

func TestVK_GroupsGetTokenPermissions(t *testing.T) {
	needGroupToken(t)

	res, err := vkGroup.GroupsGetTokenPermissions(map[string]string{})
	assert.NoError(t, err)
	assert.NotEmpty(t, res.Mask)
	assert.NotEmpty(t, res.Permissions[0].Name)
	assert.NotEmpty(t, res.Permissions[0].Setting)
}

func TestVK_GroupsInvite(t *testing.T) {
	// TODO: Add test cases.
}

func TestVK_GroupsIsMember(t *testing.T) {
	needGroupToken(t)

	res, err := vkGroup.GroupsIsMember(map[string]string{
		"group_id": "134304772",
		"user_id":  "216916273",
	})
	assert.NoError(t, err)
	assert.NotEmpty(t, res)
}

func TestVK_GroupsIsMemberExtended(t *testing.T) {
	needGroupToken(t)

	res, err := vkGroup.GroupsIsMemberExtended(map[string]string{
		"group_id": "134304772",
		"user_id":  "216916273",
	})
	assert.NoError(t, err)
	assert.NotEmpty(t, res.Member)
}

func TestVK_GroupsIsMemberUserIDsExtended(t *testing.T) {
	needGroupToken(t)

	res, err := vkGroup.GroupsIsMemberUserIDsExtended(map[string]string{
		"group_id": "134304772",
		"user_ids": "216916273",
	})
	assert.NoError(t, err)
	assert.NotEmpty(t, res[0].Member)
	assert.NotEmpty(t, res[0].UserID)
}

func TestVK_GroupsIsMemberUserIDs(t *testing.T) {
	res, err := vkGroup.GroupsIsMemberUserIDs(map[string]string{
		"group_id": "134304772",
		"user_ids": "216916273",
	})
	assert.NoError(t, err)
	assert.NotEmpty(t, res[0].Member)
	assert.NotEmpty(t, res[0].UserID)
}

func TestVK_GroupsJoin(t *testing.T) {
	// TODO: Add test cases.
}

func TestVK_GroupsLeave(t *testing.T) {
	// TODO: Add test cases.
}

func TestVK_GroupsRemoveUser(t *testing.T) {
	// TODO: Add test cases.
}

func TestVK_GroupsReorderLink(t *testing.T) {
	// TODO: Add test cases.
}

func TestVK_GroupsSearch(t *testing.T) {
	// TODO: Add test cases.
}

func TestVK_GroupsSetCallbackSettings(t *testing.T) {
	// TODO: Add test cases.
}

func TestVK_GroupsSetLongPollSettings(t *testing.T) {
	// TODO: Add test cases.
}

func TestVK_GroupsSetSettings(t *testing.T) {
	// TODO: Add test cases.
}

func TestVK_GroupsUnban(t *testing.T) {
	// TODO: Add test cases.
}
