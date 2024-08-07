package api // import "github.com/SevereCloud/vksdk/v3/api"

import (
	"github.com/SevereCloud/vksdk/v3/object"
)

// NewsfeedAddBan prevents news from specified users and communities
// from appearing in the current user's newsfeed.
//
// https://dev.vk.com/method/newsfeed.addBan
func (vk *VK) NewsfeedAddBan(params Params) (response int, err error) {
	err = vk.RequestUnmarshal("newsfeed.addBan", &response, params)
	return
}

// NewsfeedDeleteBan allows news from previously banned users and
// communities to be shown in the current user's newsfeed.
//
// https://dev.vk.com/method/newsfeed.deleteBan
func (vk *VK) NewsfeedDeleteBan(params Params) (response int, err error) {
	err = vk.RequestUnmarshal("newsfeed.deleteBan", &response, params)
	return
}

// NewsfeedDeleteList the method allows you to delete a custom news list.
//
// https://dev.vk.com/method/newsfeed.deleteList
func (vk *VK) NewsfeedDeleteList(params Params) (response int, err error) {
	err = vk.RequestUnmarshal("newsfeed.deleteList", &response, params)
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
// https://dev.vk.com/method/newsfeed.get
func (vk *VK) NewsfeedGet(params Params) (response NewsfeedGetResponse, err error) {
	err = vk.RequestUnmarshal("newsfeed.get", &response, params)
	return
}

// NewsfeedGetBannedResponse struct.
type NewsfeedGetBannedResponse struct {
	Members []int `json:"members"`
	Groups  []int `json:"groups"`
}

// NewsfeedGetBanned returns a list of users and communities banned from the current user's newsfeed.
//
//	extended=0
//
// https://dev.vk.com/method/newsfeed.getBanned
func (vk *VK) NewsfeedGetBanned(params Params) (response NewsfeedGetBannedResponse, err error) {
	err = vk.RequestUnmarshal("newsfeed.getBanned", &response, params, Params{"extended": false})

	return
}

// NewsfeedGetBannedExtendedResponse struct.
type NewsfeedGetBannedExtendedResponse struct {
	object.ExtendedResponse
}

// NewsfeedGetBannedExtended returns a list of users and communities banned from the current user's newsfeed.
//
//	extended=1
//
// https://dev.vk.com/method/newsfeed.getBanned
func (vk *VK) NewsfeedGetBannedExtended(params Params) (response NewsfeedGetBannedExtendedResponse, err error) {
	err = vk.RequestUnmarshal("newsfeed.getBanned", &response, params, Params{"extended": true})

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
// https://dev.vk.com/method/newsfeed.getComments
func (vk *VK) NewsfeedGetComments(params Params) (response NewsfeedGetCommentsResponse, err error) {
	err = vk.RequestUnmarshal("newsfeed.getComments", &response, params)
	return
}

// NewsfeedGetListsResponse struct.
type NewsfeedGetListsResponse struct {
	Count int `json:"count"`
	Items []struct {
		ID        int    `json:"id"`
		Title     string `json:"title"`
		NoReposts int    `json:"no_reposts"`
		SourceIDs []int  `json:"source_ids"`
	} `json:"items"`
}

// NewsfeedGetLists returns a list of newsfeeds followed by the current user.
//
// https://dev.vk.com/method/newsfeed.getLists
func (vk *VK) NewsfeedGetLists(params Params) (response NewsfeedGetListsResponse, err error) {
	err = vk.RequestUnmarshal("newsfeed.getLists", &response, params)
	return
}

// NewsfeedGetMentionsResponse struct.
type NewsfeedGetMentionsResponse struct {
	Count int                       `json:"count"`
	Items []object.WallWallpostToID `json:"items"`
}

// NewsfeedGetMentions returns a list of posts on user walls in which the current user is mentioned.
//
// https://dev.vk.com/method/newsfeed.getMentions
func (vk *VK) NewsfeedGetMentions(params Params) (response NewsfeedGetMentionsResponse, err error) {
	err = vk.RequestUnmarshal("newsfeed.getMentions", &response, params)
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
// https://dev.vk.com/method/newsfeed.getRecommended
func (vk *VK) NewsfeedGetRecommended(params Params) (response NewsfeedGetRecommendedResponse, err error) {
	err = vk.RequestUnmarshal("newsfeed.getRecommended", &response, params)
	return
}

// NewsfeedGetSuggestedSourcesResponse struct.
type NewsfeedGetSuggestedSourcesResponse struct {
	Count int                  `json:"count"`
	Items []object.GroupsGroup `json:"items"` // FIXME: GroupsGroup + UsersUser
}

// NewsfeedGetSuggestedSources returns communities and users that current user is suggested to follow.
//
// https://dev.vk.com/method/newsfeed.getSuggestedSources
func (vk *VK) NewsfeedGetSuggestedSources(params Params) (response NewsfeedGetSuggestedSourcesResponse, err error) {
	err = vk.RequestUnmarshal("newsfeed.getSuggestedSources", &response, params)
	return
}

// NewsfeedIgnoreItemResponse struct.
type NewsfeedIgnoreItemResponse struct {
	Status  bool    `json:"status"`
	Message *string `json:"message"`
}

// NewsfeedIgnoreItem hides an item from the newsfeed.
//
// https://dev.vk.com/method/newsfeed.ignoreItem
func (vk *VK) NewsfeedIgnoreItem(params Params) (response NewsfeedIgnoreItemResponse, err error) {
	err = vk.RequestUnmarshal("newsfeed.ignoreItem", &response, params)
	return
}

// NewsfeedSaveList creates and edits user newsfeed lists.
//
// https://dev.vk.com/method/newsfeed.saveList
func (vk *VK) NewsfeedSaveList(params Params) (response int, err error) {
	err = vk.RequestUnmarshal("newsfeed.saveList", &response, params)
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
//	extended=0
//
// https://dev.vk.com/method/newsfeed.search
func (vk *VK) NewsfeedSearch(params Params) (response NewsfeedSearchResponse, err error) {
	err = vk.RequestUnmarshal("newsfeed.search", &response, params, Params{"extended": false})

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
//	extended=1
//
// https://dev.vk.com/method/newsfeed.search
func (vk *VK) NewsfeedSearchExtended(params Params) (response NewsfeedSearchExtendedResponse, err error) {
	err = vk.RequestUnmarshal("newsfeed.search", &response, params, Params{"extended": true})

	return
}

// NewsfeedUnignoreItem returns a hidden item to the newsfeed.
//
// https://dev.vk.com/method/newsfeed.unignoreItem
func (vk *VK) NewsfeedUnignoreItem(params Params) (response int, err error) {
	err = vk.RequestUnmarshal("newsfeed.unignoreItem", &response, params)
	return
}

// NewsfeedUnsubscribe unsubscribes the current user from specified newsfeeds.
//
// https://dev.vk.com/method/newsfeed.unsubscribe
func (vk *VK) NewsfeedUnsubscribe(params Params) (response int, err error) {
	err = vk.RequestUnmarshal("newsfeed.unsubscribe", &response, params)
	return
}
