package api // import "github.com/severecloud/vksdk/5.92/api"

import (
	"encoding/json"

	"github.com/severecloud/vksdk/5.92/object"
)

// GroupsAddAddressResponse struct
type GroupsAddAddressResponse struct{}

// TODO: GroupsAddAddress
// https://vk.com/dev/groups.addAddress

// GroupsAddCallbackServerResponse struct
type GroupsAddCallbackServerResponse struct {
	ServerID int `json:"server_id"`
}

// GroupsAddCallbackServer callback API server to the community.
// https://vk.com/dev/groups.addCallbackServer
func (vk VK) GroupsAddCallbackServer(params map[string]string) (response GroupsAddCallbackServerResponse, vkErr Error) {
	rawResponse, vkErr := vk.Request("groups.addCallbackServer", params)
	if vkErr.Code != 0 {
		return
	}

	err := json.Unmarshal(rawResponse, &response)
	if err != nil {
		panic(err)
	}

	return
}

// GroupsAddLinkResponse struct
type GroupsAddLinkResponse object.GroupsGroupLink

// GroupsAddLink allows to add a link to the community.
// https://vk.com/dev/groups.addLink
func (vk VK) GroupsAddLink(params map[string]string) (response GroupsAddLinkResponse, vkErr Error) {
	rawResponse, vkErr := vk.Request("groups.addLink", params)
	if vkErr.Code != 0 {
		return
	}

	err := json.Unmarshal(rawResponse, &response)
	if err != nil {
		panic(err)
	}

	return
}

// GroupsApproveRequest allows to approve join request to the community.
// https://vk.com/dev/groups.approveRequest
func (vk VK) GroupsApproveRequest(params map[string]string) (vkErr Error) {
	_, vkErr = vk.Request("groups.approveRequest", params)

	return
}

// GroupsBan adds a user or a group to the community blacklist.
// https://vk.com/dev/groups.ban
func (vk VK) GroupsBan(params map[string]string) (vkErr Error) {
	_, vkErr = vk.Request("groups.ban", params)

	return
}

// GroupsCreateResponse struct
type GroupsCreateResponse object.GroupsGroup

// GroupsCreate creates a new community.
// https://vk.com/dev/groups.create
func (vk VK) GroupsCreate(params map[string]string) (response GroupsCreateResponse, vkErr Error) {
	rawResponse, vkErr := vk.Request("groups.create", params)
	if vkErr.Code != 0 {
		return
	}

	err := json.Unmarshal(rawResponse, &response)
	if err != nil {
		panic(err)
	}

	return
}

// GroupsDeleteAddress groups.deleteAddress
// https://vk.com/dev/groups.deleteAddress
func (vk VK) GroupsDeleteAddress(params map[string]string) (vkErr Error) {
	_, vkErr = vk.Request("groups.deleteAddress", params)

	return
}

// GroupsDeleteCallbackServer callback API server from the community.
// https://vk.com/dev/groups.deleteCallbackServer
func (vk VK) GroupsDeleteCallbackServer(params map[string]string) (vkErr Error) {
	_, vkErr = vk.Request("groups.deleteCallbackServer", params)

	return
}

// GroupsDeleteLink allows to delete a link from the community.
// https://vk.com/dev/groups.deleteLink
func (vk VK) GroupsDeleteLink(params map[string]string) (vkErr Error) {
	_, vkErr = vk.Request("groups.deleteLink", params)

	return
}

// GroupsDisableOnline disables "online" status in the community.
// https://vk.com/dev/groups.disableOnline
func (vk VK) GroupsDisableOnline(params map[string]string) (vkErr Error) {
	_, vkErr = vk.Request("groups.disableOnline", params)

	return
}

// GroupsEdit edits a community.
// https://vk.com/dev/groups.edit
func (vk VK) GroupsEdit(params map[string]string) (vkErr Error) {
	_, vkErr = vk.Request("groups.edit", params)

	return
}

// GroupsEditAddressResponse struct
type GroupsEditAddressResponse struct{}

// TODO: GroupsEditAddress
// https://vk.com/dev/groups.editAddress

// GroupsEditCallbackServer edits Callback API server in the community.
// https://vk.com/dev/groups.editCallbackServer
func (vk VK) GroupsEditCallbackServer(params map[string]string) (vkErr Error) {
	_, vkErr = vk.Request("groups.editCallbackServer", params)

	return
}

// GroupsEditLink allows to edit a link in the community.
// https://vk.com/dev/groups.editLink
func (vk VK) GroupsEditLink(params map[string]string) (vkErr Error) {
	_, vkErr = vk.Request("groups.editLink", params)

	return
}

// GroupsEditManager allows to add, remove or edit the community manager .
// https://vk.com/dev/groups.editManager
func (vk VK) GroupsEditManager(params map[string]string) (vkErr Error) {
	_, vkErr = vk.Request("groups.editManager", params)

	return
}

// GroupsEnableOnline enables "online" status in the community.
// https://vk.com/dev/groups.enableOnline
func (vk VK) GroupsEnableOnline(params map[string]string) (vkErr Error) {
	_, vkErr = vk.Request("groups.enableOnline", params)

	return
}

// GroupsGetResponse struct
type GroupsGetResponse struct {
	Count int   `json:"count"`
	Items []int `json:"items"`
}

// GroupsGet returns a list of the communities to which a user belongs.
// https://vk.com/dev/groups.get
func (vk VK) GroupsGet(params map[string]string) (response GroupsGetResponse, vkErr Error) {
	params["extended"] = "0"
	rawResponse, vkErr := vk.Request("groups.get", params)
	if vkErr.Code != 0 {
		return
	}

	err := json.Unmarshal(rawResponse, &response)
	if err != nil {
		panic(err)
	}

	return
}

// GroupsGetExtendedResponse struct
type GroupsGetExtendedResponse struct {
	Count int                  `json:"count"`
	Items []object.GroupsGroup `json:"items"`
}

// GroupsGetExtended returns a list of the communities to which a user belongs.
// https://vk.com/dev/groups.get
func (vk VK) GroupsGetExtended(params map[string]string) (response GroupsGetExtendedResponse, vkErr Error) {
	params["extended"] = "1"
	rawResponse, vkErr := vk.Request("groups.get", params)
	if vkErr.Code != 0 {
		return
	}

	err := json.Unmarshal(rawResponse, &response)
	if err != nil {
		panic(err)
	}

	return
}

// GroupsGetAddressesResponse struct
type GroupsGetAddressesResponse struct{}

// TODO: GroupsGetAddresses
// https://vk.com/dev/groups.getAddresses

// GroupsGetBannedResponse struct
type GroupsGetBannedResponse struct {
	Count int                            `json:"count"`
	Items []object.GroupsOwnerXtrBanInfo `json:"items"`
}

// GroupsGetBanned returns a list of users on a community blacklist.
// https://vk.com/dev/groups.getBanned
func (vk VK) GroupsGetBanned(params map[string]string) (response GroupsGetBannedResponse, vkErr Error) {
	rawResponse, vkErr := vk.Request("groups.getBanned", params)
	if vkErr.Code != 0 {
		return
	}

	err := json.Unmarshal(rawResponse, &response)
	if err != nil {
		panic(err)
	}

	return
}

// GroupsGetByIDResponse struct
type GroupsGetByIDResponse struct {
	Count int                  `json:"count"`
	Items []object.GroupsGroup `json:"items"`
}

// GroupsGetByID returns information about communities by their IDs.
// https://vk.com/dev/groups.getById
func (vk VK) GroupsGetByID(params map[string]string) (response GroupsGetByIDResponse, vkErr Error) {
	rawResponse, vkErr := vk.Request("groups.getById", params)
	if vkErr.Code != 0 {
		return
	}

	err := json.Unmarshal(rawResponse, &response)
	if err != nil {
		panic(err)
	}

	return
}

// GroupsGetCallbackConfirmationCodeResponse struct
type GroupsGetCallbackConfirmationCodeResponse struct {
	Code string `json:"code"`
}

// GroupsGetCallbackConfirmationCode returns Callback API confirmation code for the community.
// https://vk.com/dev/groups.getCallbackConfirmationCode
func (vk VK) GroupsGetCallbackConfirmationCode(params map[string]string) (response GroupsGetCallbackConfirmationCodeResponse, vkErr Error) {
	rawResponse, vkErr := vk.Request("groups.getCallbackConfirmationCode", params)
	if vkErr.Code != 0 {
		return
	}

	err := json.Unmarshal(rawResponse, &response)
	if err != nil {
		panic(err)
	}

	return
}

// GroupsGetCallbackServersResponse struct
type GroupsGetCallbackServersResponse struct {
	Count int `json:"count"`
	Items []struct {
		ID        int    `json:"id"`
		Title     string `json:"title"`
		CreatorID int    `json:"creator_id"`
		URL       string `json:"url"`
		SecretKey string `json:"secret_key"`
		Status    string `json:"status"`
	} `json:"items"`
}

// GroupsGetCallbackServers receives a list of Callback API servers from the community.
// https://vk.com/dev/groups.getCallbackServers
func (vk VK) GroupsGetCallbackServers(params map[string]string) (response GroupsGetCallbackServersResponse, vkErr Error) {
	rawResponse, vkErr := vk.Request("groups.getCallbackServers", params)
	if vkErr.Code != 0 {
		return
	}

	err := json.Unmarshal(rawResponse, &response)
	if err != nil {
		panic(err)
	}

	return
}

// GroupsGetCallbackSettingsResponse struct
type GroupsGetCallbackSettingsResponse object.GroupsCallbackSettings

// GroupsGetCallbackSettings returns Callback API notifications settings.
// BUG(VK): MessageEdit always 0 https://vk.com/bugtracker?act=show&id=86762
// https://vk.com/dev/groups.getCallbackSettings
func (vk VK) GroupsGetCallbackSettings(params map[string]string) (response GroupsGetCallbackSettingsResponse, vkErr Error) {
	rawResponse, vkErr := vk.Request("groups.getCallbackSettings", params)
	if vkErr.Code != 0 {
		return
	}

	err := json.Unmarshal(rawResponse, &response)
	if err != nil {
		panic(err)
	}

	return
}

// GroupsGetCatalogResponse struct
type GroupsGetCatalogResponse struct {
	Count int                  `json:"count"`
	Items []object.GroupsGroup `json:"items"`
}

// GroupsGetCatalog returns communities list for a catalog category.
// https://vk.com/dev/groups.getCatalog
func (vk VK) GroupsGetCatalog(params map[string]string) (response GroupsGetCatalogResponse, vkErr Error) {
	rawResponse, vkErr := vk.Request("groups.getCatalog", params)
	if vkErr.Code != 0 {
		return
	}

	err := json.Unmarshal(rawResponse, &response)
	if err != nil {
		panic(err)
	}

	return
}

// GroupsGetCatalogInfoResponse struct
type GroupsGetCatalogInfoResponse object.GroupsGroupCategory

// GroupsGetCatalogInfo returns categories list for communities catalog
// https://vk.com/dev/groups.getCatalogInfo
func (vk VK) GroupsGetCatalogInfo(params map[string]string) (response GroupsGetCatalogInfoResponse, vkErr Error) {
	params["extended"] = "0"
	rawResponse, vkErr := vk.Request("groups.getCatalogInfo", params)
	if vkErr.Code != 0 {
		return
	}

	err := json.Unmarshal(rawResponse, &response)
	if err != nil {
		panic(err)
	}

	return
}

// GroupsGetCatalogInfoExtendedResponse struct
type GroupsGetCatalogInfoExtendedResponse object.GroupsGroupCategoryFull

// GroupsGetCatalogInfoExtended returns categories list for communities catalog
// https://vk.com/dev/groups.getCatalogInfo
func (vk VK) GroupsGetCatalogInfoExtended(params map[string]string) (response GroupsGetCatalogInfoExtendedResponse, vkErr Error) {
	params["extended"] = "1"
	rawResponse, vkErr := vk.Request("groups.getCatalogInfo", params)
	if vkErr.Code != 0 {
		return
	}

	err := json.Unmarshal(rawResponse, &response)
	if err != nil {
		panic(err)
	}

	return
}

// GroupsGetInvitedUsersResponse struct
type GroupsGetInvitedUsersResponse struct {
	Count int                `json:"count"`
	Items []object.UsersUser `json:"items"`
}

// GroupsGetInvitedUsers returns invited users list of a community
// https://vk.com/dev/groups.getInvitedUsers
func (vk VK) GroupsGetInvitedUsers(params map[string]string) (response GroupsGetInvitedUsersResponse, vkErr Error) {
	rawResponse, vkErr := vk.Request("groups.getInvitedUsers", params)
	if vkErr.Code != 0 {
		return
	}

	err := json.Unmarshal(rawResponse, &response)
	if err != nil {
		panic(err)
	}

	return
}

// GroupsGetInvitesResponse struct
type GroupsGetInvitesResponse struct{}

// GroupsGetInvites returns a list of invitations to join communities and events.
// https://vk.com/dev/groups.getInvites
func (vk VK) GroupsGetInvites(params map[string]string) (response GroupsGetInvitesResponse, vkErr Error) {
	rawResponse, vkErr := vk.Request("groups.getInvites", params)
	if vkErr.Code != 0 {
		return
	}

	err := json.Unmarshal(rawResponse, &response)
	if err != nil {
		panic(err)
	}

	return
}

// GroupsGetLongPollServerResponse struct
type GroupsGetLongPollServerResponse object.GroupsLongPollServer

// GroupsGetLongPollServer returns data for Bots Long Poll API connection.
// https://vk.com/dev/groups.getLongPollServer
func (vk VK) GroupsGetLongPollServer(params map[string]string) (response GroupsGetLongPollServerResponse, vkErr Error) {
	rawResponse, vkErr := vk.Request("groups.getLongPollServer", params)
	if vkErr.Code != 0 {
		return
	}

	err := json.Unmarshal(rawResponse, &response)
	if err != nil {
		panic(err)
	}

	return
}

// GroupsGetLongPollSettingsResponse struct
type GroupsGetLongPollSettingsResponse object.GroupsLongPollSettings

// GroupsGetLongPollSettings returns Bots Long Poll API settings.
// https://vk.com/dev/groups.getLongPollSettings
func (vk VK) GroupsGetLongPollSettings(params map[string]string) (response GroupsGetLongPollSettingsResponse, vkErr Error) {
	rawResponse, vkErr := vk.Request("groups.getLongPollSettings", params)
	if vkErr.Code != 0 {
		return
	}

	err := json.Unmarshal(rawResponse, &response)
	if err != nil {
		panic(err)
	}

	return
}

// GroupsGetMembersResponse struct
type GroupsGetMembersResponse struct {
	Count int                `json:"count"`
	Items []object.UsersUser `json:"items"`
}

// GroupsGetMembers returns a list of community members
// https://vk.com/dev/groups.getMembers
func (vk VK) GroupsGetMembers(params map[string]string) (response GroupsGetMembersResponse, vkErr Error) {
	rawResponse, vkErr := vk.Request("groups.getMembers", params)
	if vkErr.Code != 0 {
		return
	}

	err := json.Unmarshal(rawResponse, &response)
	if err != nil {
		panic(err)
	}

	return
}

// GroupsGetMembersFilterManagersResponse struct
type GroupsGetMembersFilterManagersResponse struct {
	Count int `json:"count"`
	Items []struct {
		ID          int      `json:"id"`
		Role        string   `json:"role"`
		Permissions []string `json:"permissions,omitempty"`
	} `json:"items"`
}

// GroupsGetMembersFilterManagers returns a list of community members
// filter=managers
// https://vk.com/dev/groups.getMembers
func (vk VK) GroupsGetMembersFilterManagers(params map[string]string) (response GroupsGetMembersFilterManagersResponse, vkErr Error) {
	params["filter"] = "managers"
	rawResponse, vkErr := vk.Request("groups.getMembers", params)
	if vkErr.Code != 0 {
		return
	}

	err := json.Unmarshal(rawResponse, &response)
	if err != nil {
		panic(err)
	}

	return
}

// GroupsGetOnlineStatusResponse struct
type GroupsGetOnlineStatusResponse object.GroupsOnlineStatus

// GroupsGetOnlineStatus returns a community's online status.
// https://vk.com/dev/groups.getOnlineStatus
func (vk VK) GroupsGetOnlineStatus(params map[string]string) (response GroupsGetOnlineStatusResponse, vkErr Error) {
	rawResponse, vkErr := vk.Request("groups.getOnlineStatus", params)
	if vkErr.Code != 0 {
		return
	}

	err := json.Unmarshal(rawResponse, &response)
	if err != nil {
		panic(err)
	}

	return
}

// GroupsGetRequestsResponse struct
type GroupsGetRequestsResponse struct {
	Count int                `json:"count"`
	Items []object.UsersUser `json:"items"`
}

// GroupsGetRequests returns a list of requests to the community.
// https://vk.com/dev/groups.getRequests
func (vk VK) GroupsGetRequests(params map[string]string) (response GroupsGetRequestsResponse, vkErr Error) {
	rawResponse, vkErr := vk.Request("groups.getRequests", params)
	if vkErr.Code != 0 {
		return
	}

	err := json.Unmarshal(rawResponse, &response)
	if err != nil {
		panic(err)
	}

	return
}

// GroupsGetSettingsResponse struct
type GroupsGetSettingsResponse object.GroupsGroupSettings

// GroupsGetSettings returns community settings.
// https://vk.com/dev/groups.getSettings
func (vk VK) GroupsGetSettings(params map[string]string) (response GroupsGetSettingsResponse, vkErr Error) {
	rawResponse, vkErr := vk.Request("groups.getSettings", params)
	if vkErr.Code != 0 {
		return
	}

	err := json.Unmarshal(rawResponse, &response)
	if err != nil {
		panic(err)
	}

	return
}

// GroupsGetTokenPermissionsResponse struct
type GroupsGetTokenPermissionsResponse object.GroupsTokenPermissions

// GroupsGetTokenPermissions returns permissions scope for the community's access_token.
// https://vk.com/dev/groups.getTokenPermissions
func (vk VK) GroupsGetTokenPermissions() (response GroupsGetTokenPermissionsResponse, vkErr Error) {
	rawResponse, vkErr := vk.Request("groups.getTokenPermissions", map[string]string{})
	if vkErr.Code != 0 {
		return
	}

	err := json.Unmarshal(rawResponse, &response)
	if err != nil {
		panic(err)
	}

	return
}

// GroupsInvite allows to invite friends to the community.
// https://vk.com/dev/groups.invite
func (vk VK) GroupsInvite(params map[string]string) (vkErr Error) {
	_, vkErr = vk.Request("groups.invite", params)

	return
}

// GroupsIsMemberResponse struct
type GroupsIsMemberResponse struct{}

// TODO: GroupsIsMember returns information specifying whether a user is a member of a community.
// https://vk.com/dev/groups.isMember

// GroupsJoin with this method you can join the group or public page, and also confirm your participation in an event.
// https://vk.com/dev/groups.join
func (vk VK) GroupsJoin(params map[string]string) (vkErr Error) {
	_, vkErr = vk.Request("groups.join", params)

	return
}

// GroupsLeave with this method you can leave a group, public page, or event.
// https://vk.com/dev/groups.leave
func (vk VK) GroupsLeave(params map[string]string) (vkErr Error) {
	_, vkErr = vk.Request("groups.leave", params)

	return
}

// GroupsRemoveUser removes a user from the community.
// https://vk.com/dev/groups.removeUser
func (vk VK) GroupsRemoveUser(params map[string]string) (vkErr Error) {
	_, vkErr = vk.Request("groups.removeUser", params)

	return
}

// GroupsReorderLink allows to reorder links in the community.
// https://vk.com/dev/groups.reorderLink
func (vk VK) GroupsReorderLink(params map[string]string) (vkErr Error) {
	_, vkErr = vk.Request("groups.reorderLink", params)

	return
}

// GroupsSearchResponse struct
type GroupsSearchResponse struct {
	Count int                  `json:"count"`
	Items []object.GroupsGroup `json:"items"`
}

// GroupsSearch returns a list of communities matching the search criteria.
// https://vk.com/dev/groups.search
func (vk VK) GroupsSearch(params map[string]string) (response GroupsSearchResponse, vkErr Error) {
	rawResponse, vkErr := vk.Request("groups.search", params)
	if vkErr.Code != 0 {
		return
	}

	err := json.Unmarshal(rawResponse, &response)
	if err != nil {
		panic(err)
	}

	return
}

// GroupsSetCallbackSettings allow to set notifications settings for Callback API.
// https://vk.com/dev/groups.setCallbackSettings
func (vk VK) GroupsSetCallbackSettings(params map[string]string) (vkErr Error) {
	_, vkErr = vk.Request("groups.setCallbackSettings", params)

	return
}

// GroupsSetLongPollSettings allows to set Bots Long Poll API settings in the community.
// https://vk.com/dev/groups.setLongPollSettings
func (vk VK) GroupsSetLongPollSettings(params map[string]string) (vkErr Error) {
	_, vkErr = vk.Request("groups.setLongPollSettings", params)

	return
}

// GroupsUnban groups.unban
// https://vk.com/dev/groups.unban
func (vk VK) GroupsUnban(params map[string]string) (vkErr Error) {
	_, vkErr = vk.Request("groups.unban", params)

	return
}
