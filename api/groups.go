package api // import "github.com/severecloud/vksdk/api"

import (
	"encoding/json"

	"github.com/severecloud/vksdk/object"
)

// GroupsAddCallbackServerAddsResponse struct
type GroupsAddCallbackServerAddsResponse struct{}

// TODO: groups.addCallbackServerAdds Callback API server to the community.

// GroupsAddLinkResponse struct
type GroupsAddLinkResponse struct{}

// TODO: groups.addLink Allows to add a link to the community.

// GroupsApproveRequestResponse struct
type GroupsApproveRequestResponse struct{}

// TODO: groups.approveRequest Allows to approve join request to the community.

// GroupsBanResponse struct
type GroupsBanResponse struct{}

// TODO: groups.ban Adds a user or a group to the community blacklist.

// GroupsCreateResponse struct
type GroupsCreateResponse struct{}

// TODO: groups.create Creates a new community.

// GroupsDeleteCallbackServerDeletesResponse struct
type GroupsDeleteCallbackServerDeletesResponse struct{}

// TODO: groups.deleteCallbackServerDeletes Callback API server from the community.

// GroupsDeleteLinkResponse struct
type GroupsDeleteLinkResponse struct{}

// TODO: groups.deleteLink Allows to delete a link from the community.

// GroupsDisableOnlineResponse struct
type GroupsDisableOnlineResponse struct{}

// TODO: groups.disableOnline Disables "online" status in the community.

// GroupsEditResponse struct
type GroupsEditResponse struct{}

// TODO: groups.edit Edits a community.

// GroupsEditCallbackServerResponse struct
type GroupsEditCallbackServerResponse struct{}

// TODO: groups.editCallbackServer Edits Callback API server in the community.

// GroupsEditLinkResponse struct
type GroupsEditLinkResponse struct{}

// TODO: groups.editLink Allows to edit a link in the community.

// GroupsEditManagerResponse struct
type GroupsEditManagerResponse struct{}

// TODO: groups.editManager Allows to add, remove or edit the community manager .

// GroupsEditPlaceResponse struct
type GroupsEditPlaceResponse struct{}

// TODO: groups.editPlace Edits the place in community.

// GroupsEnableOnlineResponse struct
type GroupsEnableOnlineResponse struct{}

// TODO: groups.enableOnline Enables "online" status in the community.

// GroupsGetResponse struct
type GroupsGetResponse struct{}

// TODO: groups.get Returns a list of the communities to which a user belongs.

// GroupsGetAddressesResponse struct
type GroupsGetAddressesResponse struct{}

// TODO: groups.getAddresses

// GroupsGetBannedResponse struct
type GroupsGetBannedResponse struct{}

// TODO: groups.getBanned Returns a list of users on a community blacklist.

// GroupsGetByIDResponse struct
type GroupsGetByIDResponse struct{}

// TODO: groups.getById Returns information about communities by their IDs.

// GroupsGetCallbackConfirmationCodeResponse struct
type GroupsGetCallbackConfirmationCodeResponse struct{}

// TODO: groups.getCallbackConfirmationCode Returns Callback API confirmation code for the community.

// GroupsGetCallbackServersResponse struct
type GroupsGetCallbackServersResponse struct{}

// TODO: groups.getCallbackServers Receives a list of Callback API servers from the community.

// GroupsGetCallbackSettingsResponse struct
type GroupsGetCallbackSettingsResponse object.GroupsLongPollEvents

// GroupsGetCallbackSettings returns Callback API notifications settings.
// BUG(VK): MessageEdit always 0 https://vk.com/bugtracker?act=show&id=86762
func (vk *VK) GroupsGetCallbackSettings(params map[string]string) (response GroupsGetCallbackSettingsResponse, vkErr Error) {
	rawResponse, vkErr := vk.Request("storage.getKeys", params)
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

// TODO: groups.getCatalog Returns communities list for a catalog category.

// GroupsGetCatalogInfoResponse struct
type GroupsGetCatalogInfoResponse struct{}

// TODO: groups.getCatalogInfo Returns categories list for communities catalog

// GroupsGetInvitedUsersResponse struct
type GroupsGetInvitedUsersResponse struct{}

// TODO: groups.getInvitedUsers Returns invited users list of a community

// GroupsGetInvitesResponse struct
type GroupsGetInvitesResponse struct{}

// TODO: groups.getInvites Returns a list of invitations to join communities and events.

// GroupsGetLongPollServerResponse struct
type GroupsGetLongPollServerResponse struct{}

// TODO: groups.getLongPollServer Returns data for Bots Long Poll API connection.

// GroupsGetLongPollSettingsResponse struct
type GroupsGetLongPollSettingsResponse struct{}

// TODO: groups.getLongPollSettings Returns Bots Long Poll API settings.

// GroupsGetMembersResponse struct
type GroupsGetMembersResponse struct {
	Count int                `json:"count"`
	Items []object.UsersUser `json:"items"`
}

// GroupsGetMembers returns a list of community members
// TODO: groupx.getMembers
func (vk *VK) GroupsGetMembers(params map[string]string) (response GroupsGetMembersResponse, vkErr Error) {
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

// TODO: groups.getOnlineStatus Returns a community's online status.

// GroupsGetRequestsResponse struct
type GroupsGetRequestsResponse struct{}

// TODO: groups.getRequests Returns a list of requests to the community.

// GroupsGetSettingsResponse struct
type GroupsGetSettingsResponse struct{}

// TODO: groups.getSettings Returns community settings.

// GroupsGetTokenPermissionsResponse struct
type GroupsGetTokenPermissionsResponse struct{}

// TODO: groups.getTokenPermissions Returns permissions scope for the community's access_token.

// GroupsInviteResponse struct
type GroupsInviteResponse struct{}

// TODO: groups.invite Allows to invite friends to the community.

// GroupsIsMemberResponse struct
type GroupsIsMemberResponse struct{}

// TODO: groups.isMember Returns information specifying whether a user is a member of a community.

// GroupsJoinResponse struct
type GroupsJoinResponse struct{}

// TODO: groups.join With this method you can join the group or public page, and also confirm your participation in an event.

// GroupsLeaveResponse struct
type GroupsLeaveResponse struct{}

// TODO: groups.leave With this method you can leave a group, public page, or event.

// GroupsRemoveUserResponse struct
type GroupsRemoveUserResponse struct{}

// TODO: groups.removeUser Removes a user from the community.

// GroupsReorderLinkResponse struct
type GroupsReorderLinkResponse struct{}

// TODO: groups.reorderLink Allows to reorder links in the community.

// GroupsSearchResponse struct
type GroupsSearchResponse struct{}

// TODO: groups.search Returns a list of communities matching the search criteria.

// GroupsSetCallbackSettingsResponse struct
type GroupsSetCallbackSettingsResponse struct{}

// TODO: groups.setCallbackSettings Allow to set notifications settings for Callback API.

// GroupsSetLongPollSettingsResponse struct
type GroupsSetLongPollSettingsResponse struct{}

// TODO: groups.setLongPollSettings Allows to set Bots Long Poll API settings in the community.

// GroupsUnbanResponse struct
type GroupsUnbanResponse struct{}

// TODO: groups.unban
