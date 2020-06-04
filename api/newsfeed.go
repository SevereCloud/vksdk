package api // import "github.com/SevereCloud/vksdk/api"

import (
	"github.com/SevereCloud/vksdk/object"
)

// NewsfeedAddBan prevents news from specified users and communities
// from appearing in the current user's newsfeed.
//
// https://vk.com/dev/newsfeed.addBan
func (vk *VK) NewsfeedAddBan(params Params) (response int, err error) {
	err = vk.RequestUnmarshal("newsfeed.addBan", params, &response)
	return
}

// NewsfeedDeleteBan allows news from previously banned users and
// communities to be shown in the current user's newsfeed.
//
// https://vk.com/dev/newsfeed.deleteBan
func (vk *VK) NewsfeedDeleteBan(params Params) (response int, err error) {
	err = vk.RequestUnmarshal("newsfeed.deleteBan", params, &response)
	return
}

// NewsfeedDeleteList the method allows you to delete a custom news list.
//
// https://vk.com/dev/newsfeed.deleteList
func (vk *VK) NewsfeedDeleteList(params Params) (response int, err error) {
	err = vk.RequestUnmarshal("newsfeed.deleteList", params, &response)
	return
}

// NewsfeedGetResponse struct.
type NewsfeedGetResponse struct {
	Items []object.NewsfeedNewsfeedItem `json:"items"`
	object.ExtendedResponse
	NextFrom string `json:"next_from"`
}

// NewsfeedGet returns data required to show newsfeed for the current user.
//
// https://vk.com/dev/newsfeed.get
func (vk *VK) NewsfeedGet(params Params) (response NewsfeedGetResponse, err error) {
	err = vk.RequestUnmarshal("newsfeed.get", params, &response)
	return
}

// NewsfeedGetBannedResponse struct.
type NewsfeedGetBannedResponse struct {
	Members []int `json:"members"`
	Groups  []int `json:"groups"`
}

// NewsfeedGetBanned returns a list of users and communities banned from the current user's newsfeed.
//
// extended=0
//
// https://vk.com/dev/newsfeed.getBanned
func (vk *VK) NewsfeedGetBanned(params Params) (response NewsfeedGetBannedResponse, err error) {
	params["extended"] = false
	err = vk.RequestUnmarshal("newsfeed.getBanned", params, &response)

	return
}

// NewsfeedGetBannedExtendedResponse struct.
type NewsfeedGetBannedExtendedResponse struct {
	object.ExtendedResponse
}

// NewsfeedGetBannedExtended returns a list of users and communities banned from the current user's newsfeed.
//
// extended=1
//
// https://vk.com/dev/newsfeed.getBanned
func (vk *VK) NewsfeedGetBannedExtended(params Params) (response NewsfeedGetBannedExtendedResponse, err error) {
	params["extended"] = true
	err = vk.RequestUnmarshal("newsfeed.getBanned", params, &response)

	return
}

// NewsfeedGetCommentsResponse struct.
type NewsfeedGetCommentsResponse struct {
	Items []object.NewsfeedNewsfeedItem `json:"items"`
	object.ExtendedResponse
	NextFrom string `json:"next_from"`
}

// NewsfeedGetComments returns a list of comments in the current user's newsfeed.
//
// https://vk.com/dev/newsfeed.getComments
func (vk *VK) NewsfeedGetComments(params Params) (response NewsfeedGetCommentsResponse, err error) {
	err = vk.RequestUnmarshal("newsfeed.getComments", params, &response)
	return
}

// NewsfeedGetListsResponse struct.
type NewsfeedGetListsResponse struct {
	Count int `json:"count"`
	Items []struct {
		ID        int    `json:"id"`
		Title     string `json:"title"`
		NoReposts int    `json:"no_reposts"`
		SourceIDS []int  `json:"source_ids"` // FIXME: v2 SourceIDs
	} `json:"items"`
}

// NewsfeedGetLists returns a list of newsfeeds followed by the current user.
//
// https://vk.com/dev/newsfeed.getLists
func (vk *VK) NewsfeedGetLists(params Params) (response NewsfeedGetListsResponse, err error) {
	err = vk.RequestUnmarshal("newsfeed.getLists", params, &response)
	return
}

// NewsfeedGetMentionsResponse struct.
type NewsfeedGetMentionsResponse struct {
	Count int                       `json:"count"`
	Items []object.WallWallpostToID `json:"items"`
}

// NewsfeedGetMentions returns a list of posts on user walls in which the current user is mentioned.
//
// https://vk.com/dev/newsfeed.getMentions
func (vk *VK) NewsfeedGetMentions(params Params) (response NewsfeedGetMentionsResponse, err error) {
	err = vk.RequestUnmarshal("newsfeed.getMentions", params, &response)
	return
}

// NewsfeedGetRecommendedResponse struct.
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
func (vk *VK) NewsfeedGetRecommended(params Params) (response NewsfeedGetRecommendedResponse, err error) {
	err = vk.RequestUnmarshal("newsfeed.getRecommended", params, &response)
	return
}

// NewsfeedGetSuggestedSourcesResponse struct.
type NewsfeedGetSuggestedSourcesResponse struct {
	Count int                  `json:"count"`
	Items []object.GroupsGroup `json:"items"` // FIXME: GroupsGroup + UsersUser
}

// NewsfeedGetSuggestedSources returns communities and users that current user is suggested to follow.
//
// https://vk.com/dev/newsfeed.getSuggestedSources
func (vk *VK) NewsfeedGetSuggestedSources(params Params) (response NewsfeedGetSuggestedSourcesResponse, err error) {
	err = vk.RequestUnmarshal("newsfeed.getSuggestedSources", params, &response)
	return
}

// NewsfeedIgnoreItem hides an item from the newsfeed.
//
// https://vk.com/dev/newsfeed.ignoreItem
func (vk *VK) NewsfeedIgnoreItem(params Params) (response int, err error) {
	err = vk.RequestUnmarshal("newsfeed.ignoreItem", params, &response)
	return
}

// NewsfeedSaveList creates and edits user newsfeed lists.
//
// https://vk.com/dev/newsfeed.saveList
func (vk *VK) NewsfeedSaveList(params Params) (response int, err error) {
	err = vk.RequestUnmarshal("newsfeed.saveList", params, &response)
	return
}

// NewsfeedSearchResponse struct.
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
func (vk *VK) NewsfeedSearch(params Params) (response NewsfeedSearchResponse, err error) {
	params["extended"] = false
	err = vk.RequestUnmarshal("newsfeed.search", params, &response)

	return
}

// NewsfeedSearchExtendedResponse struct.
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
func (vk *VK) NewsfeedSearchExtended(params Params) (response NewsfeedSearchExtendedResponse, err error) {
	params["extended"] = true
	err = vk.RequestUnmarshal("newsfeed.search", params, &response)

	return
}

// NewsfeedUnignoreItem returns a hidden item to the newsfeed.
//
// https://vk.com/dev/newsfeed.unignoreItem
func (vk *VK) NewsfeedUnignoreItem(params Params) (response int, err error) {
	err = vk.RequestUnmarshal("newsfeed.unignoreItem", params, &response)
	return
}

// NewsfeedUnsubscribe unsubscribes the current user from specified newsfeeds.
//
// https://vk.com/dev/newsfeed.unsubscribe
func (vk *VK) NewsfeedUnsubscribe(params Params) (response int, err error) {
	err = vk.RequestUnmarshal("newsfeed.unsubscribe", params, &response)
	return
}
