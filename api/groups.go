package api // import "github.com/SevereCloud/vksdk/api"

import (
	"github.com/SevereCloud/vksdk/object"
)

// GroupsAddAddressResponse struct
type GroupsAddAddressResponse object.GroupsAddress

// GroupsAddAddress groups.addAddress
//
// https://vk.com/dev/groups.addAddress
func (vk *VK) GroupsAddAddress(params map[string]string) (response GroupsAddAddressResponse, vkErr Error) {
	vk.RequestUnmarshal("groups.addAddress", params, &response, &vkErr)
	return
}

// GroupsAddCallbackServerResponse struct
type GroupsAddCallbackServerResponse struct {
	ServerID int `json:"server_id"`
}

// GroupsAddCallbackServer callback API server to the community.
//
// https://vk.com/dev/groups.addCallbackServer
func (vk *VK) GroupsAddCallbackServer(params map[string]string) (response GroupsAddCallbackServerResponse, vkErr Error) {
	vk.RequestUnmarshal("groups.addCallbackServer", params, &response, &vkErr)
	return
}

// GroupsAddLinkResponse struct
type GroupsAddLinkResponse object.GroupsGroupLink

// GroupsAddLink allows to add a link to the community.
//
// https://vk.com/dev/groups.addLink
func (vk *VK) GroupsAddLink(params map[string]string) (response GroupsAddLinkResponse, vkErr Error) {
	vk.RequestUnmarshal("groups.addLink", params, &response, &vkErr)
	return
}

// GroupsApproveRequest allows to approve join request to the community.
//
// https://vk.com/dev/groups.approveRequest
func (vk *VK) GroupsApproveRequest(params map[string]string) (response int, vkErr Error) {
	vk.RequestUnmarshal("groups.approveRequest", params, &response, &vkErr)
	return
}

// GroupsBan adds a user or a group to the community blacklist.
//
// https://vk.com/dev/groups.ban
func (vk *VK) GroupsBan(params map[string]string) (response int, vkErr Error) {
	vk.RequestUnmarshal("groups.ban", params, &response, &vkErr)
	return
}

// GroupsCreateResponse struct
type GroupsCreateResponse object.GroupsGroup

// GroupsCreate creates a new community.
//
// https://vk.com/dev/groups.create
func (vk *VK) GroupsCreate(params map[string]string) (response GroupsCreateResponse, vkErr Error) {
	vk.RequestUnmarshal("groups.create", params, &response, &vkErr)
	return
}

// GroupsDeleteAddress groups.deleteAddress
//
// https://vk.com/dev/groups.deleteAddress
func (vk *VK) GroupsDeleteAddress(params map[string]string) (response int, vkErr Error) {
	vk.RequestUnmarshal("groups.deleteAddress", params, &response, &vkErr)
	return
}

// GroupsDeleteCallbackServer callback API server from the community.
//
// https://vk.com/dev/groups.deleteCallbackServer
func (vk *VK) GroupsDeleteCallbackServer(params map[string]string) (response int, vkErr Error) {
	vk.RequestUnmarshal("groups.deleteCallbackServer", params, &response, &vkErr)
	return
}

// GroupsDeleteLink allows to delete a link from the community.
//
// https://vk.com/dev/groups.deleteLink
func (vk *VK) GroupsDeleteLink(params map[string]string) (response int, vkErr Error) {
	vk.RequestUnmarshal("groups.deleteLink", params, &response, &vkErr)
	return
}

// GroupsDisableOnline disables "online" status in the community.
//
// https://vk.com/dev/groups.disableOnline
func (vk *VK) GroupsDisableOnline(params map[string]string) (response int, vkErr Error) {
	vk.RequestUnmarshal("groups.disableOnline", params, &response, &vkErr)
	return
}

// GroupsEdit edits a community.
//
// https://vk.com/dev/groups.edit
func (vk *VK) GroupsEdit(params map[string]string) (response int, vkErr Error) {
	vk.RequestUnmarshal("groups.edit", params, &response, &vkErr)
	return
}

// GroupsEditAddressResponse struct
type GroupsEditAddressResponse object.GroupsAddress

// GroupsEditAddress groups.editAddress
//
// https://vk.com/dev/groups.editAddress
func (vk *VK) GroupsEditAddress(params map[string]string) (response GroupsEditAddressResponse, vkErr Error) {
	vk.RequestUnmarshal("groups.editAddress", params, &response, &vkErr)
	return
}

// GroupsEditCallbackServer edits Callback API server in the community.
//
// https://vk.com/dev/groups.editCallbackServer
func (vk *VK) GroupsEditCallbackServer(params map[string]string) (response int, vkErr Error) {
	vk.RequestUnmarshal("groups.editCallbackServer", params, &response, &vkErr)
	return
}

// GroupsEditLink allows to edit a link in the community.
//
// https://vk.com/dev/groups.editLink
func (vk *VK) GroupsEditLink(params map[string]string) (response int, vkErr Error) {
	vk.RequestUnmarshal("groups.editLink", params, &response, &vkErr)
	return
}

// GroupsEditManager allows to add, remove or edit the community manager .
//
// https://vk.com/dev/groups.editManager
func (vk *VK) GroupsEditManager(params map[string]string) (response int, vkErr Error) {
	vk.RequestUnmarshal("groups.editManager", params, &response, &vkErr)
	return
}

// GroupsEnableOnline enables "online" status in the community.
//
// https://vk.com/dev/groups.enableOnline
func (vk *VK) GroupsEnableOnline(params map[string]string) (response int, vkErr Error) {
	vk.RequestUnmarshal("groups.enableOnline", params, &response, &vkErr)
	return
}

// GroupsGetResponse struct
type GroupsGetResponse struct {
	Count int   `json:"count"`
	Items []int `json:"items"`
}

// GroupsGet returns a list of the communities to which a user belongs.
//
// extended=0
//
// https://vk.com/dev/groups.get
func (vk *VK) GroupsGet(params map[string]string) (response GroupsGetResponse, vkErr Error) {
	params["extended"] = "0"
	vk.RequestUnmarshal("groups.get", params, &response, &vkErr)
	return
}

// GroupsGetExtendedResponse struct
type GroupsGetExtendedResponse struct {
	Count int                  `json:"count"`
	Items []object.GroupsGroup `json:"items"`
}

// GroupsGetExtended returns a list of the communities to which a user belongs.
//
// extended=1
//
// https://vk.com/dev/groups.get
func (vk *VK) GroupsGetExtended(params map[string]string) (response GroupsGetExtendedResponse, vkErr Error) {
	params["extended"] = "1"
	vk.RequestUnmarshal("groups.get", params, &response, &vkErr)
	return
}

// GroupsGetAddressesResponse struct
type GroupsGetAddressesResponse struct {
	Count int                    `json:"count"`
	Items []object.GroupsAddress `json:"items"`
}

// GroupsGetAddresses groups.getAddresses
//
// https://vk.com/dev/groups.getAddresses
func (vk *VK) GroupsGetAddresses(params map[string]string) (response GroupsGetAddressesResponse, vkErr Error) {
	vk.RequestUnmarshal("groups.getAddresses", params, &response, &vkErr)
	return
}

// GroupsGetBannedResponse struct
type GroupsGetBannedResponse struct {
	Count int                            `json:"count"`
	Items []object.GroupsOwnerXtrBanInfo `json:"items"`
}

// GroupsGetBanned returns a list of users on a community blacklist.
//
// https://vk.com/dev/groups.getBanned
func (vk *VK) GroupsGetBanned(params map[string]string) (response GroupsGetBannedResponse, vkErr Error) {
	vk.RequestUnmarshal("groups.getBanned", params, &response, &vkErr)
	return
}

// GroupsGetByIDResponse struct
type GroupsGetByIDResponse []object.GroupsGroup

// GroupsGetByID returns information about communities by their IDs.
//
// https://vk.com/dev/groups.getById
func (vk *VK) GroupsGetByID(params map[string]string) (response GroupsGetByIDResponse, vkErr Error) {
	vk.RequestUnmarshal("groups.getById", params, &response, &vkErr)
	return
}

// GroupsGetCallbackConfirmationCodeResponse struct
type GroupsGetCallbackConfirmationCodeResponse struct {
	Code string `json:"code"`
}

// GroupsGetCallbackConfirmationCode returns Callback API confirmation code for the community.
//
// https://vk.com/dev/groups.getCallbackConfirmationCode
func (vk *VK) GroupsGetCallbackConfirmationCode(params map[string]string) (response GroupsGetCallbackConfirmationCodeResponse, vkErr Error) {
	vk.RequestUnmarshal("groups.getCallbackConfirmationCode", params, &response, &vkErr)
	return
}

// GroupsGetCallbackServersResponse struct
type GroupsGetCallbackServersResponse struct {
	Count int                           `json:"count"`
	Items []object.GroupsCallbackServer `json:"items"`
}

// GroupsGetCallbackServers receives a list of Callback API servers from the community.
//
// https://vk.com/dev/groups.getCallbackServers
func (vk *VK) GroupsGetCallbackServers(params map[string]string) (response GroupsGetCallbackServersResponse, vkErr Error) {
	vk.RequestUnmarshal("groups.getCallbackServers", params, &response, &vkErr)
	return
}

// GroupsGetCallbackSettingsResponse struct
type GroupsGetCallbackSettingsResponse object.GroupsCallbackSettings

// GroupsGetCallbackSettings returns Callback API notifications settings.
// BUG(VK): MessageEdit always 0 https://vk.com/bugtracker?act=show&id=86762
//
// https://vk.com/dev/groups.getCallbackSettings
func (vk *VK) GroupsGetCallbackSettings(params map[string]string) (response GroupsGetCallbackSettingsResponse, vkErr Error) {
	vk.RequestUnmarshal("groups.getCallbackSettings", params, &response, &vkErr)
	return
}

// GroupsGetCatalogResponse struct
type GroupsGetCatalogResponse struct {
	Count int                  `json:"count"`
	Items []object.GroupsGroup `json:"items"`
}

// GroupsGetCatalog returns communities list for a catalog category.
//
// https://vk.com/dev/groups.getCatalog
func (vk *VK) GroupsGetCatalog(params map[string]string) (response GroupsGetCatalogResponse, vkErr Error) {
	vk.RequestUnmarshal("groups.getCatalog", params, &response, &vkErr)
	return
}

// GroupsGetCatalogInfoResponse struct
type GroupsGetCatalogInfoResponse object.GroupsGroupCategory

// GroupsGetCatalogInfo returns categories list for communities catalog
//
// extended=0
//
// https://vk.com/dev/groups.getCatalogInfo
func (vk *VK) GroupsGetCatalogInfo(params map[string]string) (response GroupsGetCatalogInfoResponse, vkErr Error) {
	params["extended"] = "0"
	vk.RequestUnmarshal("groups.getCatalogInfo", params, &response, &vkErr)
	return
}

// GroupsGetCatalogInfoExtendedResponse struct
type GroupsGetCatalogInfoExtendedResponse object.GroupsGroupCategoryFull

// GroupsGetCatalogInfoExtended returns categories list for communities catalog
//
// extended=1
//
// https://vk.com/dev/groups.getCatalogInfo
func (vk *VK) GroupsGetCatalogInfoExtended(params map[string]string) (response GroupsGetCatalogInfoExtendedResponse, vkErr Error) {
	params["extended"] = "1"
	vk.RequestUnmarshal("groups.getCatalogInfo", params, &response, &vkErr)
	return
}

// GroupsGetInvitedUsersResponse struct
type GroupsGetInvitedUsersResponse struct {
	Count int                `json:"count"`
	Items []object.UsersUser `json:"items"`
}

// GroupsGetInvitedUsers returns invited users list of a community
//
// https://vk.com/dev/groups.getInvitedUsers
func (vk *VK) GroupsGetInvitedUsers(params map[string]string) (response GroupsGetInvitedUsersResponse, vkErr Error) {
	vk.RequestUnmarshal("groups.getInvitedUsers", params, &response, &vkErr)
	return
}

// GroupsGetInvitesResponse struct
type GroupsGetInvitesResponse struct {
	Count int                              `json:"count"`
	Items []object.GroupsGroupXtrInvitedBy `json:"items"`
}

// GroupsGetInvites returns a list of invitations to join communities and events.
//
// https://vk.com/dev/groups.getInvites
func (vk *VK) GroupsGetInvites(params map[string]string) (response GroupsGetInvitesResponse, vkErr Error) {
	vk.RequestUnmarshal("groups.getInvites", params, &response, &vkErr)
	return
}

// GroupsGetInvitesExtendedResponse struct
type GroupsGetInvitesExtendedResponse struct {
	Count    int                              `json:"count"`
	Items    []object.GroupsGroupXtrInvitedBy `json:"items"`
	Profiles []object.UsersUser               `json:"profiles"`
	Groups   []object.GroupsGroup             `json:"groups"`
}

// GroupsGetInvitesExtended returns a list of invitations to join communities and events.
//
// https://vk.com/dev/groups.getInvites
func (vk *VK) GroupsGetInvitesExtended(params map[string]string) (response GroupsGetInvitesExtendedResponse, vkErr Error) {
	vk.RequestUnmarshal("groups.getInvites", params, &response, &vkErr)
	return
}

// GroupsGetLongPollServerResponse struct
type GroupsGetLongPollServerResponse object.GroupsLongPollServer

// GroupsGetLongPollServer returns data for Bots Long Poll API connection.
//
// https://vk.com/dev/groups.getLongPollServer
func (vk *VK) GroupsGetLongPollServer(params map[string]string) (response GroupsGetLongPollServerResponse, vkErr Error) {
	vk.RequestUnmarshal("groups.getLongPollServer", params, &response, &vkErr)
	return
}

// GroupsGetLongPollSettingsResponse struct
type GroupsGetLongPollSettingsResponse object.GroupsLongPollSettings

// GroupsGetLongPollSettings returns Bots Long Poll API settings.
//
// https://vk.com/dev/groups.getLongPollSettings
func (vk *VK) GroupsGetLongPollSettings(params map[string]string) (response GroupsGetLongPollSettingsResponse, vkErr Error) {
	vk.RequestUnmarshal("groups.getLongPollSettings", params, &response, &vkErr)
	return
}

// GroupsGetMembersResponse struct
type GroupsGetMembersResponse struct {
	Count int   `json:"count"`
	Items []int `json:"items"`
}

// GroupsGetMembers returns a list of community members
//
// https://vk.com/dev/groups.getMembers
func (vk *VK) GroupsGetMembers(params map[string]string) (response GroupsGetMembersResponse, vkErr Error) {
	params["fields"] = ""
	vk.RequestUnmarshal("groups.getMembers", params, &response, &vkErr)
	return
}

// GroupsGetMembersFieldsResponse struct
type GroupsGetMembersFieldsResponse struct {
	Count int                `json:"count"`
	Items []object.UsersUser `json:"items"`
}

// GroupsGetMembersFields returns a list of community members
//
// https://vk.com/dev/groups.getMembers
func (vk *VK) GroupsGetMembersFields(params map[string]string) (response GroupsGetMembersFieldsResponse, vkErr Error) {
	vk.RequestUnmarshal("groups.getMembers", params, &response, &vkErr)
	return
}

// GroupsGetMembersFilterManagersResponse struct
type GroupsGetMembersFilterManagersResponse struct {
	Count int                         `json:"count"`
	Items []object.GroupsMemberStatus `json:"items"`
}

// GroupsGetMembersFilterManagers returns a list of community members
// filter=managers
//
// https://vk.com/dev/groups.getMembers
func (vk *VK) GroupsGetMembersFilterManagers(params map[string]string) (response GroupsGetMembersFilterManagersResponse, vkErr Error) {
	params["filter"] = "managers"
	vk.RequestUnmarshal("groups.getMembers", params, &response, &vkErr)
	return
}

// GroupsGetOnlineStatusResponse struct
type GroupsGetOnlineStatusResponse object.GroupsOnlineStatus

// GroupsGetOnlineStatus returns a community's online status.
//
// https://vk.com/dev/groups.getOnlineStatus
func (vk *VK) GroupsGetOnlineStatus(params map[string]string) (response GroupsGetOnlineStatusResponse, vkErr Error) {
	vk.RequestUnmarshal("groups.getOnlineStatus", params, &response, &vkErr)
	return
}

// GroupsGetRequestsResponse struct
type GroupsGetRequestsResponse struct {
	Count int                `json:"count"`
	Items []object.UsersUser `json:"items"`
}

// GroupsGetRequests returns a list of requests to the community.
//
// https://vk.com/dev/groups.getRequests
func (vk *VK) GroupsGetRequests(params map[string]string) (response GroupsGetRequestsResponse, vkErr Error) {
	vk.RequestUnmarshal("groups.getRequests", params, &response, &vkErr)
	return
}

// GroupsGetSettingsResponse struct
type GroupsGetSettingsResponse object.GroupsGroupSettings

// GroupsGetSettings returns community settings.
//
// https://vk.com/dev/groups.getSettings
func (vk *VK) GroupsGetSettings(params map[string]string) (response GroupsGetSettingsResponse, vkErr Error) {
	vk.RequestUnmarshal("groups.getSettings", params, &response, &vkErr)
	return
}

// GroupsGetTokenPermissionsResponse struct
type GroupsGetTokenPermissionsResponse object.GroupsTokenPermissions

// GroupsGetTokenPermissions returns permissions scope for the community's access_token.
//
// https://vk.com/dev/groups.getTokenPermissions
func (vk *VK) GroupsGetTokenPermissions(params map[string]string) (response GroupsGetTokenPermissionsResponse, vkErr Error) {
	vk.RequestUnmarshal("groups.getTokenPermissions", params, &response, &vkErr)
	return
}

// GroupsInvite allows to invite friends to the community.
//
// https://vk.com/dev/groups.invite
func (vk *VK) GroupsInvite(params map[string]string) (response int, vkErr Error) {
	vk.RequestUnmarshal("groups.invite", params, &response, &vkErr)
	return
}

// GroupsIsMember returns information specifying whether a user is a member of a community.
//
// extended=0
//
// https://vk.com/dev/groups.isMember
func (vk *VK) GroupsIsMember(params map[string]string) (response int, vkErr Error) {
	params["extended"] = "0"
	vk.RequestUnmarshal("groups.isMember", params, &response, &vkErr)
	return
}

// GroupsIsMemberExtendedResponse struct
type GroupsIsMemberExtendedResponse struct {
	Invitation int `json:"invitation"` // Information whether user has been invited to the group
	Member     int `json:"member"`     // Information whether user is a member of the group
	Request    int `json:"request"`    // Information whether user has send request to the group
	CanInvite  int `json:"can_invite"` // Information whether user can be invite
	CanRecall  int `json:"can_recall"` // Information whether user's invite to the group can be recalled
}

// GroupsIsMemberExtended returns information specifying whether a user is a member of a community.
//
// extended=1
//
// https://vk.com/dev/groups.isMember
func (vk *VK) GroupsIsMemberExtended(params map[string]string) (response GroupsIsMemberExtendedResponse, vkErr Error) {
	params["extended"] = "1"
	vk.RequestUnmarshal("groups.isMember", params, &response, &vkErr)
	return
}

// GroupsIsMemberUserIDsExtendedResponse struct
type GroupsIsMemberUserIDsExtendedResponse []object.GroupsMemberStatusFull

// GroupsIsMemberUserIDsExtended returns information specifying whether a user is a member of a community.
//
// extended=1
// need user_ids
//
// https://vk.com/dev/groups.isMember
func (vk *VK) GroupsIsMemberUserIDsExtended(params map[string]string) (response GroupsIsMemberUserIDsExtendedResponse, vkErr Error) {
	params["extended"] = "1"
	vk.RequestUnmarshal("groups.isMember", params, &response, &vkErr)
	return
}

// GroupsIsMemberUserIDsResponse struct
type GroupsIsMemberUserIDsResponse []object.GroupsMemberStatus

// GroupsIsMemberUserIDs returns information specifying whether a user is a member of a community.
//
// extended=0
// need user_ids
//
// https://vk.com/dev/groups.isMember
func (vk *VK) GroupsIsMemberUserIDs(params map[string]string) (response GroupsIsMemberUserIDsResponse, vkErr Error) {
	params["extended"] = "0"
	vk.RequestUnmarshal("groups.isMember", params, &response, &vkErr)
	return
}

// GroupsJoin with this method you can join the group or public page, and also confirm your participation in an event.
//
// https://vk.com/dev/groups.join
func (vk *VK) GroupsJoin(params map[string]string) (response int, vkErr Error) {
	vk.RequestUnmarshal("groups.join", params, &response, &vkErr)
	return
}

// GroupsLeave with this method you can leave a group, public page, or event.
//
// https://vk.com/dev/groups.leave
func (vk *VK) GroupsLeave(params map[string]string) (response int, vkErr Error) {
	vk.RequestUnmarshal("groups.leave", params, &response, &vkErr)
	return
}

// GroupsRemoveUser removes a user from the community.
//
// https://vk.com/dev/groups.removeUser
func (vk *VK) GroupsRemoveUser(params map[string]string) (response int, vkErr Error) {
	vk.RequestUnmarshal("groups.removeUser", params, &response, &vkErr)
	return
}

// GroupsReorderLink allows to reorder links in the community.
//
// https://vk.com/dev/groups.reorderLink
func (vk *VK) GroupsReorderLink(params map[string]string) (response int, vkErr Error) {
	vk.RequestUnmarshal("groups.reorderLink", params, &response, &vkErr)
	return
}

// GroupsSearchResponse struct
type GroupsSearchResponse struct {
	Count int                  `json:"count"`
	Items []object.GroupsGroup `json:"items"`
}

// GroupsSearch returns a list of communities matching the search criteria.
//
// https://vk.com/dev/groups.search
func (vk *VK) GroupsSearch(params map[string]string) (response GroupsSearchResponse, vkErr Error) {
	vk.RequestUnmarshal("groups.search", params, &response, &vkErr)
	return
}

// GroupsSetCallbackSettings allow to set notifications settings for Callback API.
//
// https://vk.com/dev/groups.setCallbackSettings
func (vk *VK) GroupsSetCallbackSettings(params map[string]string) (response int, vkErr Error) {
	vk.RequestUnmarshal("groups.setCallbackSettings", params, &response, &vkErr)
	return
}

// GroupsSetLongPollSettings allows to set Bots Long Poll API settings in the community.
//
// https://vk.com/dev/groups.setLongPollSettings
func (vk *VK) GroupsSetLongPollSettings(params map[string]string) (response int, vkErr Error) {
	vk.RequestUnmarshal("groups.setLongPollSettings", params, &response, &vkErr)
	return
}

// GroupsSetSettings sets community settings
//
// https://vk.com/dev/groups.setSettings
func (vk *VK) GroupsSetSettings(params map[string]string) (response int, vkErr Error) {
	vk.RequestUnmarshal("groups.setSettings", params, &response, &vkErr)
	return
}

// GroupsUnban groups.unban
//
// https://vk.com/dev/groups.unban
func (vk *VK) GroupsUnban(params map[string]string) (response int, vkErr Error) {
	vk.RequestUnmarshal("groups.unban", params, &response, &vkErr)
	return
}
