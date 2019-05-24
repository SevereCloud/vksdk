package api // import "github.com/SevereCloud/vksdk/5.92/api"

import (
	"github.com/SevereCloud/vksdk/5.92/object"
)

// FaveAddGroupResponse struct
type FaveAddGroupResponse int

// FaveAddGroup adds a community to user faves.
//
// https://vk.com/dev/fave.addGroup
func (vk *VK) FaveAddGroup(params map[string]string) (response FaveAddGroupResponse, vkErr Error) {
	vk.RequestUnmarshal("fave.addGroup", params, &response, &vkErr)
	return
}

// FaveAddLinkResponse struct
type FaveAddLinkResponse int

// FaveAddLink adds a link to user faves.
//
// https://vk.com/dev/fave.addLink
func (vk *VK) FaveAddLink(params map[string]string) (response FaveAddLinkResponse, vkErr Error) {
	vk.RequestUnmarshal("fave.addLink", params, &response, &vkErr)
	return
}

// FaveAddUserResponse struct
type FaveAddUserResponse int

// FaveAddUser adds a profile to user faves.
//
// https://vk.com/dev/fave.addUser
func (vk *VK) FaveAddUser(params map[string]string) (response FaveAddUserResponse, vkErr Error) {
	vk.RequestUnmarshal("fave.addUser", params, &response, &vkErr)
	return
}

// FaveGetLinksResponse struct
type FaveGetLinksResponse struct {
	Count int                    `json:"count"`
	Items []object.FaveFavesLink `json:"items"`
}

// FaveGetLinks returns a list of links that the current user has bookmarked.
//
// https://vk.com/dev/fave.getLinks
func (vk *VK) FaveGetLinks(params map[string]string) (response FaveGetLinksResponse, vkErr Error) {
	vk.RequestUnmarshal("fave.getLinks", params, &response, &vkErr)
	return
}

// FaveGetMarketItemsResponse struct
type FaveGetMarketItemsResponse struct {
	Count int                       `json:"count"`
	Items []object.MarketMarketItem `json:"items"`
}

// FaveGetMarketItems returns market items bookmarked by current user.
//
// extended=0
//
// https://vk.com/dev/fave.getMarketItems
func (vk *VK) FaveGetMarketItems(params map[string]string) (response FaveGetMarketItemsResponse, vkErr Error) {
	params["extended"] = "0"
	vk.RequestUnmarshal("fave.getMarketItems", params, &response, &vkErr)
	return
}

// FaveGetMarketItemsExtended returns market items bookmarked by current user.
//
// extended=1
//
// https://vk.com/dev/fave.getMarketItems
func (vk *VK) FaveGetMarketItemsExtended(params map[string]string) (response FaveGetMarketItemsResponse, vkErr Error) {
	params["extended"] = "1"
	vk.RequestUnmarshal("fave.getMarketItems", params, &response, &vkErr)
	return
}

// FaveGetPhotosResponse struct
type FaveGetPhotosResponse struct {
	Count    int                  `json:"count"`
	Items    []object.PhotosPhoto `json:"items"`
	NextFrom string               `json:"next_from"`
}

// FaveGetPhotos returns a list of photos that the current user has liked.
//
// https://vk.com/dev/fave.getPhotos
func (vk *VK) FaveGetPhotos(params map[string]string) (response FaveGetPhotosResponse, vkErr Error) {
	vk.RequestUnmarshal("fave.getPhotos", params, &response, &vkErr)
	return
}

// FaveGetPostsResponse struct
type FaveGetPostsResponse struct {
	Count int                   `json:"count"`
	Items []object.WallWallpost `json:"items"`
}

// FaveGetPosts returns a list of wall posts that the current user has liked.
//
// extended=0
//
// https://vk.com/dev/fave.getPosts
func (vk *VK) FaveGetPosts(params map[string]string) (response FaveGetPostsResponse, vkErr Error) {
	params["extended"] = "0"
	vk.RequestUnmarshal("fave.getPosts", params, &response, &vkErr)
	return
}

// FaveGetPostsExtendedResponse struct
type FaveGetPostsExtendedResponse struct {
	Count    int                   `json:"count"`
	Items    []object.WallWallpost `json:"items"`
	Profiles []object.UsersUser    `json:"profiles"`
	Groups   []object.GroupsGroup  `json:"groups"`
}

// FaveGetPostsExtended returns a list of wall posts that the current user has liked.
//
// extended=1
//
// https://vk.com/dev/fave.getPosts
func (vk *VK) FaveGetPostsExtended(params map[string]string) (response FaveGetPostsExtendedResponse, vkErr Error) {
	params["extended"] = "1"
	vk.RequestUnmarshal("fave.getPosts", params, &response, &vkErr)
	return
}

// FaveGetUsersResponse struct
type FaveGetUsersResponse struct {
	Count int                `json:"count"`
	Items []object.UsersUser `json:"items"`
}

// FaveGetUsers returns a list of users whom the current user has bookmarked.
//
// https://vk.com/dev/fave.getUsers
func (vk *VK) FaveGetUsers(params map[string]string) (response FaveGetUsersResponse, vkErr Error) {
	vk.RequestUnmarshal("fave.getUsers", params, &response, &vkErr)
	return
}

// FaveGetVideosResponse struct
type FaveGetVideosResponse struct {
	Count int                 `json:"count"`
	Items []object.VideoVideo `json:"items"`
}

// FaveGetVideos returns a list of videos that the current user has liked.
//
// extended=0
//
// https://vk.com/dev/fave.getVideos
func (vk *VK) FaveGetVideos(params map[string]string) (response FaveGetVideosResponse, vkErr Error) {
	params["extended"] = "0"
	vk.RequestUnmarshal("fave.getVideos", params, &response, &vkErr)
	return
}

// FaveGetVideosExtendedResponse struct
type FaveGetVideosExtendedResponse struct {
	Count    int                  `json:"count"`
	Items    []object.VideoVideo  `json:"items"`
	Profiles []object.UsersUser   `json:"profiles"`
	Groups   []object.GroupsGroup `json:"groups"`
}

// FaveGetVideosExtended returns a list of videos that the current user has liked.
//
// extended=1
//
// https://vk.com/dev/fave.getVideos
func (vk *VK) FaveGetVideosExtended(params map[string]string) (response FaveGetVideosExtendedResponse, vkErr Error) {
	params["extended"] = "1"
	vk.RequestUnmarshal("fave.getVideos", params, &response, &vkErr)
	return
}

// FaveRemoveGroupResponse struct
type FaveRemoveGroupResponse int

// FaveRemoveGroup removes a community from user faves.
//
// https://vk.com/dev/fave.removeGroup
func (vk *VK) FaveRemoveGroup(params map[string]string) (response FaveRemoveGroupResponse, vkErr Error) {
	vk.RequestUnmarshal("fave.removeGroup", params, &response, &vkErr)
	return
}

// FaveRemoveLinkResponse struct
type FaveRemoveLinkResponse int

// FaveRemoveLink removes link from the user's faves.
//
// https://vk.com/dev/fave.removeLink
func (vk *VK) FaveRemoveLink(params map[string]string) (response FaveRemoveLinkResponse, vkErr Error) {
	vk.RequestUnmarshal("fave.removeLink", params, &response, &vkErr)
	return
}

// FaveRemoveUserResponse struct
type FaveRemoveUserResponse int

// FaveRemoveUser removes a profile from user faves.
//
// https://vk.com/dev/fave.removeUser
func (vk *VK) FaveRemoveUser(params map[string]string) (response FaveRemoveUserResponse, vkErr Error) {
	vk.RequestUnmarshal("fave.removeUser", params, &response, &vkErr)
	return
}
