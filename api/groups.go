package api // import "github.com/SevereCloud/vksdk/api"

import (
	"github.com/SevereCloud/vksdk/object"
)

// GroupsAddAddressResponse struct
type GroupsAddAddressResponse object.GroupsAddress

// GroupsAddAddress groups.addAddress
//
// https://vk.com/dev/groups.addAddress
func (vk *VK) GroupsAddAddress(params Params) (response GroupsAddAddressResponse, err error) {
	err = vk.RequestUnmarshal("groups.addAddress", params, &response)
	return
}

// GroupsAddCallbackServerResponse struct
type GroupsAddCallbackServerResponse struct {
	ServerID int `json:"server_id"`
}

// GroupsAddCallbackServer callback API server to the community.
//
// https://vk.com/dev/groups.addCallbackServer
func (vk *VK) GroupsAddCallbackServer(params Params) (response GroupsAddCallbackServerResponse, err error) {
	err = vk.RequestUnmarshal("groups.addCallbackServer", params, &response)
	return
}

// GroupsAddLinkResponse struct
type GroupsAddLinkResponse object.GroupsGroupLink

// GroupsAddLink allows to add a link to the community.
//
// https://vk.com/dev/groups.addLink
func (vk *VK) GroupsAddLink(params Params) (response GroupsAddLinkResponse, err error) {
	err = vk.RequestUnmarshal("groups.addLink", params, &response)
	return
}

// GroupsApproveRequest allows to approve join request to the community.
//
// https://vk.com/dev/groups.approveRequest
func (vk *VK) GroupsApproveRequest(params Params) (response int, err error) {
	err = vk.RequestUnmarshal("groups.approveRequest", params, &response)
	return
}

// GroupsBan adds a user or a group to the community blacklist.
//
// https://vk.com/dev/groups.ban
func (vk *VK) GroupsBan(params Params) (response int, err error) {
	err = vk.RequestUnmarshal("groups.ban", params, &response)
	return
}

// GroupsCreateResponse struct
type GroupsCreateResponse object.GroupsGroup

// GroupsCreate creates a new community.
//
// https://vk.com/dev/groups.create
func (vk *VK) GroupsCreate(params Params) (response GroupsCreateResponse, err error) {
	err = vk.RequestUnmarshal("groups.create", params, &response)
	return
}

// GroupsDeleteAddress groups.deleteAddress
//
// https://vk.com/dev/groups.deleteAddress
func (vk *VK) GroupsDeleteAddress(params Params) (response int, err error) {
	err = vk.RequestUnmarshal("groups.deleteAddress", params, &response)
	return
}

// GroupsDeleteCallbackServer callback API server from the community.
//
// https://vk.com/dev/groups.deleteCallbackServer
func (vk *VK) GroupsDeleteCallbackServer(params Params) (response int, err error) {
	err = vk.RequestUnmarshal("groups.deleteCallbackServer", params, &response)
	return
}

// GroupsDeleteLink allows to delete a link from the community.
//
// https://vk.com/dev/groups.deleteLink
func (vk *VK) GroupsDeleteLink(params Params) (response int, err error) {
	err = vk.RequestUnmarshal("groups.deleteLink", params, &response)
	return
}

// GroupsDisableOnline disables "online" status in the community.
//
// https://vk.com/dev/groups.disableOnline
func (vk *VK) GroupsDisableOnline(params Params) (response int, err error) {
	err = vk.RequestUnmarshal("groups.disableOnline", params, &response)
	return
}

// GroupsEdit edits a community.
//
// https://vk.com/dev/groups.edit
func (vk *VK) GroupsEdit(params Params) (response int, err error) {
	err = vk.RequestUnmarshal("groups.edit", params, &response)
	return
}

// GroupsEditAddressResponse struct
type GroupsEditAddressResponse object.GroupsAddress

// GroupsEditAddress groups.editAddress
//
// https://vk.com/dev/groups.editAddress
func (vk *VK) GroupsEditAddress(params Params) (response GroupsEditAddressResponse, err error) {
	err = vk.RequestUnmarshal("groups.editAddress", params, &response)
	return
}

// GroupsEditCallbackServer edits Callback API server in the community.
//
// https://vk.com/dev/groups.editCallbackServer
func (vk *VK) GroupsEditCallbackServer(params Params) (response int, err error) {
	err = vk.RequestUnmarshal("groups.editCallbackServer", params, &response)
	return
}

// GroupsEditLink allows to edit a link in the community.
//
// https://vk.com/dev/groups.editLink
func (vk *VK) GroupsEditLink(params Params) (response int, err error) {
	err = vk.RequestUnmarshal("groups.editLink", params, &response)
	return
}

// GroupsEditManager allows to add, remove or edit the community manager .
//
// https://vk.com/dev/groups.editManager
func (vk *VK) GroupsEditManager(params Params) (response int, err error) {
	err = vk.RequestUnmarshal("groups.editManager", params, &response)
	return
}

// GroupsEnableOnline enables "online" status in the community.
//
// https://vk.com/dev/groups.enableOnline
func (vk *VK) GroupsEnableOnline(params Params) (response int, err error) {
	err = vk.RequestUnmarshal("groups.enableOnline", params, &response)
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
func (vk *VK) GroupsGet(params Params) (response GroupsGetResponse, err error) {
	params["extended"] = false
	err = vk.RequestUnmarshal("groups.get", params, &response)

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
func (vk *VK) GroupsGetExtended(params Params) (response GroupsGetExtendedResponse, err error) {
	params["extended"] = true
	err = vk.RequestUnmarshal("groups.get", params, &response)

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
func (vk *VK) GroupsGetAddresses(params Params) (response GroupsGetAddressesResponse, err error) {
	err = vk.RequestUnmarshal("groups.getAddresses", params, &response)
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
func (vk *VK) GroupsGetBanned(params Params) (response GroupsGetBannedResponse, err error) {
	err = vk.RequestUnmarshal("groups.getBanned", params, &response)
	return
}

// GroupsGetByIDResponse struct
type GroupsGetByIDResponse []object.GroupsGroup

// GroupsGetByID returns information about communities by their IDs.
//
// https://vk.com/dev/groups.getById
func (vk *VK) GroupsGetByID(params Params) (response GroupsGetByIDResponse, err error) {
	err = vk.RequestUnmarshal("groups.getById", params, &response)
	return
}

// GroupsGetCallbackConfirmationCodeResponse struct
type GroupsGetCallbackConfirmationCodeResponse struct {
	Code string `json:"code"`
}

// GroupsGetCallbackConfirmationCode returns Callback API confirmation code for the community.
//
// https://vk.com/dev/groups.getCallbackConfirmationCode
func (vk *VK) GroupsGetCallbackConfirmationCode(params Params) (response GroupsGetCallbackConfirmationCodeResponse, err error) {
	err = vk.RequestUnmarshal("groups.getCallbackConfirmationCode", params, &response)
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
func (vk *VK) GroupsGetCallbackServers(params Params) (response GroupsGetCallbackServersResponse, err error) {
	err = vk.RequestUnmarshal("groups.getCallbackServers", params, &response)
	return
}

// GroupsGetCallbackSettingsResponse struct
type GroupsGetCallbackSettingsResponse object.GroupsCallbackSettings

// GroupsGetCallbackSettings returns Callback API notifications settings.
// BUG(VK): MessageEdit always 0 https://vk.com/bugtracker?act=show&id=86762
//
// https://vk.com/dev/groups.getCallbackSettings
func (vk *VK) GroupsGetCallbackSettings(params Params) (response GroupsGetCallbackSettingsResponse, err error) {
	err = vk.RequestUnmarshal("groups.getCallbackSettings", params, &response)
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
func (vk *VK) GroupsGetCatalog(params Params) (response GroupsGetCatalogResponse, err error) {
	err = vk.RequestUnmarshal("groups.getCatalog", params, &response)
	return
}

// GroupsGetCatalogInfoResponse struct
type GroupsGetCatalogInfoResponse struct {
	Enabled    object.BaseBoolInt           `json:"enabled"`
	Categories []object.GroupsGroupCategory `json:"categories"`
}

// GroupsGetCatalogInfo returns categories list for communities catalog
//
// extended=0
//
// https://vk.com/dev/groups.getCatalogInfo
func (vk *VK) GroupsGetCatalogInfo(params Params) (response GroupsGetCatalogInfoResponse, err error) {
	params["extended"] = false
	err = vk.RequestUnmarshal("groups.getCatalogInfo", params, &response)

	return
}

// GroupsGetCatalogInfoExtendedResponse struct
type GroupsGetCatalogInfoExtendedResponse struct {
	Enabled    object.BaseBoolInt               `json:"enabled"`
	Categories []object.GroupsGroupCategoryFull `json:"categories"`
}

// GroupsGetCatalogInfoExtended returns categories list for communities catalog
//
// extended=1
//
// https://vk.com/dev/groups.getCatalogInfo
func (vk *VK) GroupsGetCatalogInfoExtended(params Params) (response GroupsGetCatalogInfoExtendedResponse, err error) {
	params["extended"] = true
	err = vk.RequestUnmarshal("groups.getCatalogInfo", params, &response)

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
func (vk *VK) GroupsGetInvitedUsers(params Params) (response GroupsGetInvitedUsersResponse, err error) {
	err = vk.RequestUnmarshal("groups.getInvitedUsers", params, &response)
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
func (vk *VK) GroupsGetInvites(params Params) (response GroupsGetInvitesResponse, err error) {
	err = vk.RequestUnmarshal("groups.getInvites", params, &response)
	return
}

// GroupsGetInvitesExtendedResponse struct
type GroupsGetInvitesExtendedResponse struct {
	Count int                              `json:"count"`
	Items []object.GroupsGroupXtrInvitedBy `json:"items"`
	object.ExtendedResponse
}

// GroupsGetInvitesExtended returns a list of invitations to join communities and events.
//
// https://vk.com/dev/groups.getInvites
func (vk *VK) GroupsGetInvitesExtended(params Params) (response GroupsGetInvitesExtendedResponse, err error) {
	err = vk.RequestUnmarshal("groups.getInvites", params, &response)
	return
}

// GroupsGetLongPollServerResponse struct
type GroupsGetLongPollServerResponse object.GroupsLongPollServer

// GroupsGetLongPollServer returns data for Bots Long Poll API connection.
//
// https://vk.com/dev/groups.getLongPollServer
func (vk *VK) GroupsGetLongPollServer(params Params) (response GroupsGetLongPollServerResponse, err error) {
	err = vk.RequestUnmarshal("groups.getLongPollServer", params, &response)
	return
}

// GroupsGetLongPollSettingsResponse struct
type GroupsGetLongPollSettingsResponse object.GroupsLongPollSettings

// GroupsGetLongPollSettings returns Bots Long Poll API settings.
//
// https://vk.com/dev/groups.getLongPollSettings
func (vk *VK) GroupsGetLongPollSettings(params Params) (response GroupsGetLongPollSettingsResponse, err error) {
	err = vk.RequestUnmarshal("groups.getLongPollSettings", params, &response)
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
func (vk *VK) GroupsGetMembers(params Params) (response GroupsGetMembersResponse, err error) {
	params["fields"] = ""
	params["filter"] = ""
	err = vk.RequestUnmarshal("groups.getMembers", params, &response)

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
func (vk *VK) GroupsGetMembersFields(params Params) (response GroupsGetMembersFieldsResponse, err error) {
	if v, prs := params["fields"]; v == "" || !prs {
		params["fields"] = "id"
	}

	err = vk.RequestUnmarshal("groups.getMembers", params, &response)

	return
}

// GroupsGetMembersFilterManagersResponse struct
type GroupsGetMembersFilterManagersResponse struct {
	Count int                                   `json:"count"`
	Items []object.GroupsMemberRoleXtrUsersUser `json:"items"`
}

// GroupsGetMembersFilterManagers returns a list of community members
// filter=managers
//
// https://vk.com/dev/groups.getMembers
func (vk *VK) GroupsGetMembersFilterManagers(params Params) (response GroupsGetMembersFilterManagersResponse, err error) {
	params["filter"] = "managers"
	err = vk.RequestUnmarshal("groups.getMembers", params, &response)

	return
}

// GroupsGetOnlineStatusResponse struct
type GroupsGetOnlineStatusResponse object.GroupsOnlineStatus

// GroupsGetOnlineStatus returns a community's online status.
//
// https://vk.com/dev/groups.getOnlineStatus
func (vk *VK) GroupsGetOnlineStatus(params Params) (response GroupsGetOnlineStatusResponse, err error) {
	err = vk.RequestUnmarshal("groups.getOnlineStatus", params, &response)
	return
}

// GroupsGetRequestsResponse struct
type GroupsGetRequestsResponse struct {
	Count int   `json:"count"`
	Items []int `json:"items"`
}

// GroupsGetRequests returns a list of requests to the community.
//
// https://vk.com/dev/groups.getRequests
func (vk *VK) GroupsGetRequests(params Params) (response GroupsGetRequestsResponse, err error) {
	params["fields"] = ""
	err = vk.RequestUnmarshal("groups.getRequests", params, &response)

	return
}

// GroupsGetRequestsFieldsResponse struct
type GroupsGetRequestsFieldsResponse struct {
	Count int                `json:"count"`
	Items []object.UsersUser `json:"items"`
}

// GroupsGetRequestsFields returns a list of requests to the community.
//
// https://vk.com/dev/groups.getRequests
func (vk *VK) GroupsGetRequestsFields(params Params) (response GroupsGetRequestsFieldsResponse, err error) {
	if v, prs := params["fields"]; v == "" || !prs {
		params["fields"] = "id"
	}

	err = vk.RequestUnmarshal("groups.getRequests", params, &response)

	return
}

// GroupsGetSettingsResponse struct
type GroupsGetSettingsResponse object.GroupsGroupSettings

// GroupsGetSettings returns community settings.
//
// https://vk.com/dev/groups.getSettings
func (vk *VK) GroupsGetSettings(params Params) (response GroupsGetSettingsResponse, err error) {
	err = vk.RequestUnmarshal("groups.getSettings", params, &response)
	return
}

// GroupsGetTokenPermissionsResponse struct
type GroupsGetTokenPermissionsResponse object.GroupsTokenPermissions

// GroupsGetTokenPermissions returns permissions scope for the community's access_token.
//
// https://vk.com/dev/groups.getTokenPermissions
func (vk *VK) GroupsGetTokenPermissions(params Params) (response GroupsGetTokenPermissionsResponse, err error) {
	err = vk.RequestUnmarshal("groups.getTokenPermissions", params, &response)
	return
}

// GroupsInvite allows to invite friends to the community.
//
// https://vk.com/dev/groups.invite
func (vk *VK) GroupsInvite(params Params) (response int, err error) {
	err = vk.RequestUnmarshal("groups.invite", params, &response)
	return
}

// GroupsIsMember returns information specifying whether a user is a member of a community.
//
// extended=0
//
// https://vk.com/dev/groups.isMember
func (vk *VK) GroupsIsMember(params Params) (response int, err error) {
	params["extended"] = false
	err = vk.RequestUnmarshal("groups.isMember", params, &response)

	return
}

// GroupsIsMemberExtendedResponse struct
type GroupsIsMemberExtendedResponse struct {
	Invitation object.BaseBoolInt `json:"invitation"` // Information whether user has been invited to the group
	Member     object.BaseBoolInt `json:"member"`     // Information whether user is a member of the group
	Request    object.BaseBoolInt `json:"request"`    // Information whether user has send request to the group
	CanInvite  object.BaseBoolInt `json:"can_invite"` // Information whether user can be invite
	CanRecall  object.BaseBoolInt `json:"can_recall"` // Information whether user's invite to the group can be recalled
}

// GroupsIsMemberExtended returns information specifying whether a user is a member of a community.
//
// extended=1
//
// https://vk.com/dev/groups.isMember
func (vk *VK) GroupsIsMemberExtended(params Params) (response GroupsIsMemberExtendedResponse, err error) {
	params["extended"] = true
	err = vk.RequestUnmarshal("groups.isMember", params, &response)

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
func (vk *VK) GroupsIsMemberUserIDsExtended(params Params) (response GroupsIsMemberUserIDsExtendedResponse, err error) {
	params["extended"] = true
	err = vk.RequestUnmarshal("groups.isMember", params, &response)

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
func (vk *VK) GroupsIsMemberUserIDs(params Params) (response GroupsIsMemberUserIDsResponse, err error) {
	params["extended"] = false
	err = vk.RequestUnmarshal("groups.isMember", params, &response)

	return
}

// GroupsJoin with this method you can join the group or public page, and also confirm your participation in an event.
//
// https://vk.com/dev/groups.join
func (vk *VK) GroupsJoin(params Params) (response int, err error) {
	err = vk.RequestUnmarshal("groups.join", params, &response)
	return
}

// GroupsLeave with this method you can leave a group, public page, or event.
//
// https://vk.com/dev/groups.leave
func (vk *VK) GroupsLeave(params Params) (response int, err error) {
	err = vk.RequestUnmarshal("groups.leave", params, &response)
	return
}

// GroupsRemoveUser removes a user from the community.
//
// https://vk.com/dev/groups.removeUser
func (vk *VK) GroupsRemoveUser(params Params) (response int, err error) {
	err = vk.RequestUnmarshal("groups.removeUser", params, &response)
	return
}

// GroupsReorderLink allows to reorder links in the community.
//
// https://vk.com/dev/groups.reorderLink
func (vk *VK) GroupsReorderLink(params Params) (response int, err error) {
	err = vk.RequestUnmarshal("groups.reorderLink", params, &response)
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
func (vk *VK) GroupsSearch(params Params) (response GroupsSearchResponse, err error) {
	err = vk.RequestUnmarshal("groups.search", params, &response)
	return
}

// GroupsSetCallbackSettings allow to set notifications settings for Callback API.
//
// https://vk.com/dev/groups.setCallbackSettings
func (vk *VK) GroupsSetCallbackSettings(params Params) (response int, err error) {
	err = vk.RequestUnmarshal("groups.setCallbackSettings", params, &response)
	return
}

// GroupsSetLongPollSettings allows to set Bots Long Poll API settings in the community.
//
// https://vk.com/dev/groups.setLongPollSettings
func (vk *VK) GroupsSetLongPollSettings(params Params) (response int, err error) {
	err = vk.RequestUnmarshal("groups.setLongPollSettings", params, &response)
	return
}

// GroupsSetSettings sets community settings
//
// https://vk.com/dev/groups.setSettings
func (vk *VK) GroupsSetSettings(params Params) (response int, err error) {
	err = vk.RequestUnmarshal("groups.setSettings", params, &response)
	return
}

// GroupsUnban groups.unban
//
// https://vk.com/dev/groups.unban
func (vk *VK) GroupsUnban(params Params) (response int, err error) {
	err = vk.RequestUnmarshal("groups.unban", params, &response)
	return
}
