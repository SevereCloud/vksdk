package api // import "github.com/SevereCloud/vksdk/5.92/api"

import (
	"github.com/SevereCloud/vksdk/5.92/object"
)

// NewsfeedAddBanResponse struct
type NewsfeedAddBanResponse int

// NewsfeedAddBan prevents news from specified users and communities
// from appearing in the current user's newsfeed.
//
// https://vk.com/dev/newsfeed.addBan
func (vk *VK) NewsfeedAddBan(params map[string]string) (response NewsfeedAddBanResponse, vkErr Error) {
	vk.RequestUnmarshal("newsfeed.addBan", params, &response, &vkErr)
	return
}

// NewsfeedDeleteBanResponse struct
type NewsfeedDeleteBanResponse int

// NewsfeedDeleteBan allows news from previously banned users and
// communities to be shown in the current user's newsfeed.
//
// https://vk.com/dev/newsfeed.deleteBan
func (vk *VK) NewsfeedDeleteBan(params map[string]string) (response NewsfeedDeleteBanResponse, vkErr Error) {
	vk.RequestUnmarshal("newsfeed.deleteBan", params, &response, &vkErr)
	return
}

// NewsfeedDeleteListResponse struct
type NewsfeedDeleteListResponse int

// NewsfeedDeleteList The method allows you to delete a custom news list.
//
// https://vk.com/dev/newsfeed.deleteList
func (vk *VK) NewsfeedDeleteList(params map[string]string) (response NewsfeedDeleteListResponse, vkErr Error) {
	vk.RequestUnmarshal("newsfeed.deleteList", params, &response, &vkErr)
	return
}

// NewsfeedGetResponse struct
type NewsfeedGetResponse struct {
	Items    []object.NewsfeedNewsfeedItem `json:"items"`
	Profiles []object.UsersUser            `json:"profiles"`
	Groups   []object.GroupsGroup          `json:"groups"`
	NextFrom string                        `json:"next_from"`
}

// NewsfeedGet returns data required to show newsfeed for the current user.
//
// https://vk.com/dev/newsfeed.get
func (vk *VK) NewsfeedGet(params map[string]string) (response NewsfeedGetResponse, vkErr Error) {
	vk.RequestUnmarshal("newsfeed.get", params, &response, &vkErr)
	return
}

// NewsfeedGetBannedResponse struct
type NewsfeedGetBannedResponse struct{
	Members []int              `json:"members"`
	Groups   []int             `json:"groups"`
}

// NewsfeedGetBanned returns a list of users and communities banned from the current user's newsfeed.
//
// extended=0
//
// https://vk.com/dev/newsfeed.getBanned
func (vk *VK) NewsfeedGetBanned(params map[string]string) (response NewsfeedGetBannedResponse, vkErr Error) {
	params["extended"] = "0"
	vk.RequestUnmarshal("newsfeed.getBanned", params, &response, &vkErr)
	return
}

// NewsfeedGetBannedExtendedResponse struct
type NewsfeedGetBannedExtendedResponse struct{
	Profiles []object.UsersUser               `json:"profiles"`
	Groups   []object.GroupsGroup             `json:"groups"`
}

// NewsfeedGetBannedExtended returns a list of users and communities banned from the current user's newsfeed.
//
// extended=1
//
// https://vk.com/dev/newsfeed.getBanned
func (vk *VK) NewsfeedGetBannedExtended(params map[string]string) (response NewsfeedGetBannedExtendedResponse, vkErr Error) {
	params["extended"] = "1"
	vk.RequestUnmarshal("newsfeed.getBanned", params, &response, &vkErr)
	return
}

// NewsfeedGetCommentsResponse struct
type NewsfeedGetCommentsResponse struct{
	Items    []object.NewsfeedNewsfeedItem `json:"items"`
	Profiles []object.UsersUser            `json:"profiles"`
	Groups   []object.GroupsGroup          `json:"groups"`
	NextFrom string                        `json:"next_from"`
}

// NewsfeedGetComments returns a list of comments in the current user's newsfeed.
//
// https://vk.com/dev/newsfeed.getComments
func (vk *VK) NewsfeedGetComments(params map[string]string) (response NewsfeedGetCommentsResponse, vkErr Error) {
	vk.RequestUnmarshal("newsfeed.getComments", params, &response, &vkErr)
	return
}

type NewsfeedGetListsResponse struct {
	Count int `json:"count"`
	Items []struct {
		ID        int    `json:"id"`
		Title     string `json:"title"`
		NoReposts int    `json:"no_reposts"`
		SourceIds []int  `json:"source_ids"`
	} `json:"items"`
}

// NewsfeedGetLists returns a list of newsfeeds followed by the current user.
//
// https://vk.com/dev/newsfeed.getLists
func (vk *VK) NewsfeedGetLists(params map[string]string) (response NewsfeedGetListsResponse, vkErr Error) {
	vk.RequestUnmarshal("newsfeed.getLists", params, &response, &vkErr)
	return
}

// NewsfeedGetMentionsResponse struct
type NewsfeedGetMentionsResponse struct {
	Count int                       `json:"count"`
	Items []object.WallWallpostToID `json:"items"`
}

// NewsfeedGetMentions returns a list of posts on user walls in which the current user is mentioned.
//
// https://vk.com/dev/newsfeed.getMentions
func (vk *VK) NewsfeedGetMentions(params map[string]string) (response NewsfeedGetMentionsResponse, vkErr Error) {
	vk.RequestUnmarshal("newsfeed.getMentions", params, &response, &vkErr)
	return
}

// NewsfeedGetRecommendedResponse struct
type NewsfeedGetRecommendedResponse struct {
	Items      []object.NewsfeedNewsfeedItem `json:"items"`
	Profiles   []object.UsersUser            `json:"profiles"`
	Groups     []object.GroupsGroup          `json:"groups"`
	NextOffset string                        `json:"next_offset"`
	NextFrom   string                        `json:"next_from"`
}

// NewsfeedGetRecommended returns a list of newsfeeds recommended to the current user.
//
// https://vk.com/dev/newsfeed.getRecommended
func (vk *VK) NewsfeedGetRecommended(params map[string]string) (response NewsfeedGetRecommendedResponse, vkErr Error) {
	vk.RequestUnmarshal("newsfeed.getRecommended", params, &response, &vkErr)
	return
}

// NewsfeedGetSuggestedSourcesResponse struct
type NewsfeedGetSuggestedSourcesResponse struct {
	Count int                  `json:"count"`
	Items []object.GroupsGroup `json:"items"` // FIXME: GroupsGroup + UsersUser
}

// NewsfeedGetSuggestedSources returns communities and users that current user is suggested to follow.
//
// https://vk.com/dev/newsfeed.getSuggestedSources
func (vk *VK) NewsfeedGetSuggestedSources(params map[string]string) (response NewsfeedGetSuggestedSourcesResponse, vkErr Error) {
	vk.RequestUnmarshal("newsfeed.getSuggestedSources", params, &response, &vkErr)
	return
}

// NewsfeedIgnoreItemHidesResponse struct
type NewsfeedIgnoreItemHidesResponse int

// NewsfeedIgnoreItemHides an item from the newsfeed.
//
// https://vk.com/dev/newsfeed.ignoreItemHides
func (vk *VK) NewsfeedIgnoreItemHides(params map[string]string) (response NewsfeedIgnoreItemHidesResponse, vkErr Error) {
	vk.RequestUnmarshal("newsfeed.ignoreItemHides", params, &response, &vkErr)
	return
}

// NewsfeedSaveListResponse struct
type NewsfeedSaveListResponse int

// NewsfeedSaveList creates and edits user newsfeed lists
//
// https://vk.com/dev/newsfeed.saveList
func (vk *VK) NewsfeedSaveList(params map[string]string) (response NewsfeedSaveListResponse, vkErr Error) {
	vk.RequestUnmarshal("newsfeed.saveList", params, &response, &vkErr)
	return
}

// NewsfeedSearchResponse struct
type NewsfeedSearchResponse struct {
	Items      []object.WallWallpost `json:"items"`
	Count      int                   `json:"count"`
	TotalCount int                   `json:"total_count"`
	NextFrom   string                `json:"next_from"`
}

// NewsfeedSearch returns search results by statuses.
//
// extended=0
//
// https://vk.com/dev/newsfeed.search
func (vk *VK) NewsfeedSearch(params map[string]string) (response NewsfeedSearchResponse, vkErr Error) {
	params["extended"] = "0"
	vk.RequestUnmarshal("newsfeed.search", params, &response, &vkErr)
	return
}

// NewsfeedSearchExtendedResponse struct
type NewsfeedSearchExtendedResponse struct {
	Items      []object.WallWallpost `json:"items"`
	Count      int                   `json:"count"`
	TotalCount int                   `json:"total_count"`
	Profiles   []object.UsersUser    `json:"profiles"`
	Groups     []object.GroupsGroup  `json:"groups"`
	NextFrom   string                `json:"next_from"`
}

// NewsfeedSearchExtended returns search results by statuses.
//
// extended=1
//
// https://vk.com/dev/newsfeed.search
func (vk *VK) NewsfeedSearchExtended(params map[string]string) (response NewsfeedSearchExtendedResponse, vkErr Error) {
	params["extended"] = "1"
	vk.RequestUnmarshal("newsfeed.search", params, &response, &vkErr)
	return
}

// NewsfeedUnignoreItemResponse struct
type NewsfeedUnignoreItemResponse int

// NewsfeedUnignoreItem returns a hidden item to the newsfeed.
//
// https://vk.com/dev/newsfeed.unignoreItem
func (vk *VK) NewsfeedUnignoreItem(params map[string]string) (response NewsfeedUnignoreItemResponse, vkErr Error) {
	vk.RequestUnmarshal("newsfeed.unignoreItem", params, &response, &vkErr)
	return
}

// NewsfeedUnsubscribeResponse struct
type NewsfeedUnsubscribeResponse int

// NewsfeedUnsubscribe unsubscribes the current user from specified newsfeeds.
//
// https://vk.com/dev/newsfeed.unsubscribe
func (vk *VK) NewsfeedUnsubscribe(params map[string]string) (response NewsfeedUnsubscribeResponse, vkErr Error) {
	vk.RequestUnmarshal("newsfeed.unsubscribe", params, &response, &vkErr)
	return
}
