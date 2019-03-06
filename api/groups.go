package api // import "github.com/severecloud/vksdk/api"

import (
	"encoding/json"

	"github.com/severecloud/vksdk/object"
)

// GroupsAddAddressResponse struct
type GroupsAddAddressResponse struct{}

// TODO: groups.addAddress
// https://vk.com/dev/groups.addAddress

// GroupsAddCallbackServerAddsResponse struct
type GroupsAddCallbackServerAddsResponse struct{}

// TODO: groups.addCallbackServerAdds callback API server to the community.
// https://vk.com/dev/groups.addCallbackServerAdds

// GroupsAddLinkResponse struct
type GroupsAddLinkResponse struct{}

// TODO: groups.addLink allows to add a link to the community.
// https://vk.com/dev/groups.addLink

// GroupsApproveRequestResponse struct
type GroupsApproveRequestResponse struct{}

// TODO: groups.approveRequest allows to approve join request to the community.
// https://vk.com/dev/groups.approveRequest

// GroupsBanResponse struct
type GroupsBanResponse struct{}

// TODO: groups.ban adds a user or a group to the community blacklist.
// https://vk.com/dev/groups.ban

// GroupsCreateResponse struct
type GroupsCreateResponse struct{}

// TODO: groups.create creates a new community.
// https://vk.com/dev/groups.create

// GroupsDeleteAddress groups.deleteAddress
// TODO: information for groups.deleteAddress
// https://vk.com/dev/groups.deleteAddress
func (vk VK) GroupsDeleteAddress(params map[string]string) (vkErr Error) {
	_, vkErr = vk.Request("groups.deleteAddress", params)

	return
}

// GroupsDeleteCallbackServerDeletesResponse struct
type GroupsDeleteCallbackServerDeletesResponse struct{}

// TODO: groups.deleteCallbackServerDeletes callback API server from the community.
// https://vk.com/dev/groups.deleteCallbackServerDeletes

// GroupsDeleteLinkResponse struct
type GroupsDeleteLinkResponse struct{}

// TODO: groups.deleteLink allows to delete a link from the community.
// https://vk.com/dev/groups.deleteLink

// GroupsDisableOnlineResponse struct
type GroupsDisableOnlineResponse struct{}

// TODO: groups.disableOnline disables "online" status in the community.
// https://vk.com/dev/groups.disableOnline

// GroupsEditAddressResponse struct
type GroupsEditAddressResponse struct{}

// TODO: groups.editAddress
// https://vk.com/dev/groups.editAddress

// GroupsEditResponse struct
type GroupsEditResponse struct{}

// TODO: groups.edit edits a community.
// https://vk.com/dev/groups.edit

// GroupsEditCallbackServerResponse struct
type GroupsEditCallbackServerResponse struct{}

// TODO: groups.editCallbackServer edits Callback API server in the community.
// https://vk.com/dev/groups.editCallbackServer

// GroupsEditLinkResponse struct
type GroupsEditLinkResponse struct{}

// TODO: groups.editLink allows to edit a link in the community.
// https://vk.com/dev/groups.editLink

// GroupsEditManagerResponse struct
type GroupsEditManagerResponse struct{}

// TODO: groups.editManager allows to add, remove or edit the community manager .
// https://vk.com/dev/groups.editManager

// GroupsEditPlaceResponse struct
type GroupsEditPlaceResponse struct{}

// TODO: groups.editPlace edits the place in community.
// https://vk.com/dev/groups.editPlace

// GroupsEnableOnlineResponse struct
type GroupsEnableOnlineResponse struct{}

// TODO: groups.enableOnline enables "online" status in the community.
// https://vk.com/dev/groups.enableOnline

// GroupsGetResponse struct
type GroupsGetResponse struct{}

// TODO: groups.get returns a list of the communities to which a user belongs.
// https://vk.com/dev/groups.get

// GroupsGetAddressesResponse struct
type GroupsGetAddressesResponse struct{}

// TODO: groups.getAddresses
// https://vk.com/dev/groups.getAddresses

// GroupsGetBannedResponse struct
type GroupsGetBannedResponse struct{}

// TODO: groups.getBanned returns a list of users on a community blacklist.
// https://vk.com/dev/groups.getBanned

// GroupsGetByIDResponse struct
type GroupsGetByIDResponse struct{}

// TODO: groups.getById returns information about communities by their IDs.
// https://vk.com/dev/groups.getById

// GroupsGetCallbackConfirmationCodeResponse struct
type GroupsGetCallbackConfirmationCodeResponse struct{}

// TODO: groups.getCallbackConfirmationCode returns Callback API confirmation code for the community.
// https://vk.com/dev/groups.getCallbackConfirmationCode

// GroupsGetCallbackServersResponse struct
type GroupsGetCallbackServersResponse struct{}

// TODO: groups.getCallbackServers receives a list of Callback API servers from the community.
// https://vk.com/dev/groups.getCallbackServers

// GroupsGetCallbackSettingsResponse struct
type GroupsGetCallbackSettingsResponse object.GroupsLongPollEvents

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
type GroupsGetCatalogResponse struct{}

// TODO: groups.getCatalog returns communities list for a catalog category.
// https://vk.com/dev/groups.getCatalog

// GroupsGetCatalogInfoResponse struct
type GroupsGetCatalogInfoResponse struct{}

// TODO: groups.getCatalogInfo returns categories list for communities catalog
// https://vk.com/dev/groups.getCatalogInfo

// GroupsGetInvitedUsersResponse struct
type GroupsGetInvitedUsersResponse struct{}

// TODO: groups.getInvitedUsers returns invited users list of a community
// https://vk.com/dev/groups.getInvitedUsers

// GroupsGetInvitesResponse struct
type GroupsGetInvitesResponse struct{}

// TODO: groups.getInvites returns a list of invitations to join communities and events.
// https://vk.com/dev/groups.getInvites

// GroupsGetLongPollServerResponse struct
type GroupsGetLongPollServerResponse struct{}

// TODO: groups.getLongPollServer returns data for Bots Long Poll API connection.
// https://vk.com/dev/groups.getLongPollServer

// GroupsGetLongPollSettingsResponse struct
type GroupsGetLongPollSettingsResponse struct{}

// TODO: groups.getLongPollSettings returns Bots Long Poll API settings.
// https://vk.com/dev/groups.getLongPollSettings

// GroupsGetMembersResponse struct
type GroupsGetMembersResponse struct {
	Count int                `json:"count"`
	Items []object.UsersUser `json:"items"`
}

// GroupsGetMembers returns a list of community members
// TODO: groups.getMembers
// https://vk.com/dev/groups.getMembers
func (vk VK) GroupsGetMembers(params map[string]string) (response GroupsGetMembersResponse, vkErr Error) {
	rawResponse, vkErr := vk.Request("groups.getMembers", params)
	if vkErr.Code != 0 {
		return
	}
	// FIXME list if no filter
	err := json.Unmarshal(rawResponse, &response)
	if err != nil {
		panic(err)
	}

	return
}

// GroupsGetOnlineStatusResponse struct
type GroupsGetOnlineStatusResponse struct{}

// TODO: groups.getOnlineStatus returns a community's online status.
// https://vk.com/dev/groups.getOnlineStatus

// GroupsGetRequestsResponse struct
type GroupsGetRequestsResponse struct{}

// TODO: groups.getRequests returns a list of requests to the community.
// https://vk.com/dev/groups.getRequests

// GroupsGetSettingsResponse struct
type GroupsGetSettingsResponse struct{}

// TODO: groups.getSettings returns community settings.
// https://vk.com/dev/groups.getSettings

// GroupsGetTokenPermissionsResponse struct
type GroupsGetTokenPermissionsResponse struct{}

// TODO: groups.getTokenPermissions returns permissions scope for the community's access_token.
// https://vk.com/dev/groups.getTokenPermissions

// GroupsInviteResponse struct
type GroupsInviteResponse struct{}

// TODO: groups.invite allows to invite friends to the community.
// https://vk.com/dev/groups.invite

// GroupsIsMemberResponse struct
type GroupsIsMemberResponse struct{}

// TODO: groups.isMember returns information specifying whether a user is a member of a community.
// https://vk.com/dev/groups.isMember

// GroupsJoinResponse struct
type GroupsJoinResponse struct{}

// TODO: groups.join with this method you can join the group or public page, and also confirm your participation in an event.
// https://vk.com/dev/groups.join

// GroupsLeaveResponse struct
type GroupsLeaveResponse struct{}

// TODO: groups.leave with this method you can leave a group, public page, or event.
// https://vk.com/dev/groups.leave

// GroupsRemoveUserResponse struct
type GroupsRemoveUserResponse struct{}

// TODO: groups.removeUser removes a user from the community.
// https://vk.com/dev/groups.removeUser

// GroupsReorderLinkResponse struct
type GroupsReorderLinkResponse struct{}

// TODO: groups.reorderLink allows to reorder links in the community.
// https://vk.com/dev/groups.reorderLink

// GroupsSearchResponse struct
type GroupsSearchResponse struct{}

// TODO: groups.search returns a list of communities matching the search criteria.
// https://vk.com/dev/groups.search

// GroupsSetCallbackSettingsResponse struct
type GroupsSetCallbackSettingsResponse struct{}

// TODO: groups.setCallbackSettings allow to set notifications settings for Callback API.
// https://vk.com/dev/groups.setCallbackSettings

// GroupsSetLongPollSettingsResponse struct
type GroupsSetLongPollSettingsResponse struct{}

// TODO: groups.setLongPollSettings allows to set Bots Long Poll API settings in the community.
// https://vk.com/dev/groups.setLongPollSettings

// GroupsUnbanResponse struct
type GroupsUnbanResponse struct{}

// TODO: groups.unban
// https://vk.com/dev/groups.unban
