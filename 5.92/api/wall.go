package api // import "github.com/SevereCloud/vksdk/5.92/api"

import (
	"github.com/SevereCloud/vksdk/5.92/object"
)

// WallCloseComments turn off post commenting
//
// https://vk.com/dev/wall.closeComments
func (vk *VK) WallCloseComments(params map[string]string) (response int, vkErr Error) {
	vk.RequestUnmarshal("wall.closeComments", params, &response, &vkErr)
	return
}

// WallCreateCommentResponse struct
type WallCreateCommentResponse struct {
	CommentID    int   `json:"comment_id"`
	ParentsStack []int `json:"parents_stack"`
}

// WallCreateComment Adds a comment to a post on a user wall or community wall.
//
// https://vk.com/dev/wall.createComment
func (vk *VK) WallCreateComment(params map[string]string) (response WallCreateCommentResponse, vkErr Error) {
	vk.RequestUnmarshal("wall.createComment", params, &response, &vkErr)
	return
}

// WallDelete deletes a post from a user wall or community wall.
//
// https://vk.com/dev/wall.delete
func (vk *VK) WallDelete(params map[string]string) (response int, vkErr Error) {
	vk.RequestUnmarshal("wall.delete", params, &response, &vkErr)
	return
}

// WallDeleteComment deletes a comment on a post on a user wall or community wall.
//
// https://vk.com/dev/wall.deleteComment
func (vk *VK) WallDeleteComment(params map[string]string) (response int, vkErr Error) {
	vk.RequestUnmarshal("wall.deleteComment", params, &response, &vkErr)
	return
}

// WallEdit edits a post on a user wall or community wall.
//
// TODO: update 5.100
//
// https://vk.com/dev/wall.edit
func (vk *VK) WallEdit(params map[string]string) (response int, vkErr Error) {
	vk.RequestUnmarshal("wall.edit", params, &response, &vkErr)
	return
}

// WallEditAdsStealth allows to edit hidden post.
//
// https://vk.com/dev/wall.editAdsStealth
func (vk *VK) WallEditAdsStealth(params map[string]string) (response int, vkErr Error) {
	vk.RequestUnmarshal("wall.editAdsStealth", params, &response, &vkErr)
	return
}

// WallEditComment edits a comment on a user wall or community wall.
//
// https://vk.com/dev/wall.editComment
func (vk *VK) WallEditComment(params map[string]string) (response int, vkErr Error) {
	vk.RequestUnmarshal("wall.editComment", params, &response, &vkErr)
	return
}

// WallGetResponse struct
type WallGetResponse struct {
	Count int                   `json:"count"`
	Items []object.WallWallpost `json:"items"`
}

// WallGet returns a list of posts on a user wall or community wall.
//
// extended=0
//
// https://vk.com/dev/wall.get
func (vk *VK) WallGet(params map[string]string) (response WallGetResponse, vkErr Error) {
	params["extended"] = "0"
	vk.RequestUnmarshal("wall.get", params, &response, &vkErr)
	return
}

// WallGetExtendedResponse struct
type WallGetExtendedResponse struct {
	Count    int                   `json:"count"`
	Items    []object.WallWallpost `json:"items"`
	Profiles []object.UsersUser    `json:"profiles"`
	Groups   []object.GroupsGroup  `json:"groups"`
}

// WallGetExtended returns a list of posts on a user wall or community wall.
//
// extended=1
//
// https://vk.com/dev/wall.get
func (vk *VK) WallGetExtended(params map[string]string) (response WallGetExtendedResponse, vkErr Error) {
	params["extended"] = "1"
	vk.RequestUnmarshal("wall.get", params, &response, &vkErr)
	return
}

// WallGetByIDResponse struct
type WallGetByIDResponse []object.WallWallpost

// WallGetByID returns a list of posts from user or community walls by their IDs.
//
// extended=0
//
// https://vk.com/dev/wall.getById
func (vk *VK) WallGetByID(params map[string]string) (response WallGetByIDResponse, vkErr Error) {
	params["extended"] = "0"
	vk.RequestUnmarshal("wall.getById", params, &response, &vkErr)
	return
}

// WallGetByIDExtendedResponse struct
type WallGetByIDExtendedResponse struct {
	Items    []object.WallWallpost `json:"items"`
	Profiles []object.UsersUser    `json:"profiles"`
	Groups   []object.GroupsGroup  `json:"groups"`
}

// WallGetByIDExtended returns a list of posts from user or community walls by their IDs.
//
// extended=1
//
// https://vk.com/dev/wall.getById
func (vk *VK) WallGetByIDExtended(params map[string]string) (response WallGetByIDExtendedResponse, vkErr Error) {
	params["extended"] = "1"
	vk.RequestUnmarshal("wall.getById", params, &response, &vkErr)
	return
}

// WallGetCommentResponse struct
type WallGetCommentResponse struct {
	Count             int                      `json:"count"`
	Items             []object.WallWallComment `json:"items"`
	CanPost           bool                     `json:"can_post"`
	ShowReplyButton   bool                     `json:"show_reply_button"`
	GroupsCanPost     bool                     `json:"groups_can_post"`
	CurrentLevelCount int                      `json:"current_level_count"`
}

// WallGetComment allows to obtain wall comment info.
//
// extended=0
//
// https://vk.com/dev/wall.getComment
func (vk *VK) WallGetComment(params map[string]string) (response WallGetCommentResponse, vkErr Error) {
	params["extended"] = "0"
	vk.RequestUnmarshal("wall.getComment", params, &response, &vkErr)
	return
}

// WallGetCommentExtendedResponse struct
type WallGetCommentExtendedResponse struct {
	Count             int                      `json:"count"`
	Items             []object.WallWallComment `json:"items"`
	CanPost           bool                     `json:"can_post"`
	ShowReplyButton   bool                     `json:"show_reply_button"`
	GroupsCanPost     bool                     `json:"groups_can_post"`
	CurrentLevelCount int                      `json:"current_level_count"`
	Profiles          []object.UsersUser       `json:"profiles"`
	Groups            []object.GroupsGroup     `json:"groups"`
}

// WallGetCommentExtended allows to obtain wall comment info.
//
// extended=1
//
// https://vk.com/dev/wall.getComment
func (vk *VK) WallGetCommentExtended(params map[string]string) (response WallGetCommentExtendedResponse, vkErr Error) {
	params["extended"] = "1"
	vk.RequestUnmarshal("wall.getComment", params, &response, &vkErr)
	return
}

// WallGetCommentsResponse struct
type WallGetCommentsResponse struct {
	Count int                      `json:"count"`
	Items []object.WallWallComment `json:"items"`
}

// WallGetComments returns a list of comments on a post on a user wall or community wall.
//
// extended=0
//
// https://vk.com/dev/wall.getComments
func (vk *VK) WallGetComments(params map[string]string) (response WallGetCommentsResponse, vkErr Error) {
	params["extended"] = "0"
	vk.RequestUnmarshal("wall.getComments", params, &response, &vkErr)
	return
}

// WallGetCommentsExtendedResponse struct
type WallGetCommentsExtendedResponse struct {
	Count    int                      `json:"count"`
	Items    []object.WallWallComment `json:"items"`
	Profiles []object.UsersUser       `json:"profiles"`
	Groups   []object.GroupsGroup     `json:"groups"`
}

// WallGetCommentsExtended returns a list of comments on a post on a user wall or community wall.
//
// extended=1
//
// https://vk.com/dev/wall.getComments
func (vk *VK) WallGetCommentsExtended(params map[string]string) (response WallGetCommentsExtendedResponse, vkErr Error) {
	params["extended"] = "1"
	vk.RequestUnmarshal("wall.getComments", params, &response, &vkErr)
	return
}

// WallGetRepostsResponse struct
type WallGetRepostsResponse struct {
	Items    []object.WallWallpost `json:"items"`
	Profiles []object.UsersUser    `json:"profiles"`
	Groups   []object.GroupsGroup  `json:"groups"`
}

// WallGetReposts returns information about reposts of a post on user wall or community wall.
//
// https://vk.com/dev/wall.getReposts
func (vk *VK) WallGetReposts(params map[string]string) (response WallGetRepostsResponse, vkErr Error) {
	vk.RequestUnmarshal("wall.getReposts", params, &response, &vkErr)
	return
}

// WallOpenComments includes posting comments.
//
// https://vk.com/dev/wall.openComments
func (vk *VK) WallOpenComments(params map[string]string) (response int, vkErr Error) {
	vk.RequestUnmarshal("wall.openComments", params, &response, &vkErr)
	return
}

// WallPin pins the post on wall.
//
// https://vk.com/dev/wall.pin
func (vk *VK) WallPin(params map[string]string) (response int, vkErr Error) {
	vk.RequestUnmarshal("wall.pin", params, &response, &vkErr)
	return
}

// WallPostResponse struct
type WallPostResponse struct {
	PostID int `json:"post_id"`
}

// WallPost adds a new post on a user wall or community wall.Can also be used to publish suggested or scheduled posts.
//
// https://vk.com/dev/wall.post
func (vk *VK) WallPost(params map[string]string) (response WallPostResponse, vkErr Error) {
	vk.RequestUnmarshal("wall.post", params, &response, &vkErr)
	return
}

// WallPostAdsStealthResponse struct
type WallPostAdsStealthResponse struct {
	PostID int `json:"post_id"`
}

// WallPostAdsStealth allows to create hidden post which will
// not be shown on the community's wall and can be used for creating
// an ad with type "Community post".
//
// https://vk.com/dev/wall.postAdsStealth
func (vk *VK) WallPostAdsStealth(params map[string]string) (response WallPostAdsStealthResponse, vkErr Error) {
	vk.RequestUnmarshal("wall.postAdsStealth", params, &response, &vkErr)
	return
}

// WallReportComment reports (submits a complaint about) a comment on a post on a user wall or community wall.
//
// https://vk.com/dev/wall.reportComment
func (vk *VK) WallReportComment(params map[string]string) (response int, vkErr Error) {
	vk.RequestUnmarshal("wall.reportComment", params, &response, &vkErr)
	return
}

// WallReportPost reports (submits a complaint about) a post on a user wall or community wall.
//
// https://vk.com/dev/wall.reportPost
func (vk *VK) WallReportPost(params map[string]string) (response int, vkErr Error) {
	vk.RequestUnmarshal("wall.reportPost", params, &response, &vkErr)
	return
}

// WallRepostResponse struct
type WallRepostResponse struct {
	Success      int `json:"success"`
	PostID       int `json:"post_id"`
	RepostsCount int `json:"reposts_count"`
	LikesCount   int `json:"likes_count"`
}

// WallRepost reposts ( copies) an object to a user wall or community wall.
//
// https://vk.com/dev/wall.repost
func (vk *VK) WallRepost(params map[string]string) (response WallRepostResponse, vkErr Error) {
	vk.RequestUnmarshal("wall.repost", params, &response, &vkErr)
	return
}

// WallRestore restores a post deleted from a user wall or community wall.
//
// https://vk.com/dev/wall.restore
func (vk *VK) WallRestore(params map[string]string) (response int, vkErr Error) {
	vk.RequestUnmarshal("wall.restore", params, &response, &vkErr)
	return
}

// WallRestoreComment restores a comment deleted from a user wall or community wall.
//
// https://vk.com/dev/wall.restoreComment
func (vk *VK) WallRestoreComment(params map[string]string) (response int, vkErr Error) {
	vk.RequestUnmarshal("wall.restoreComment", params, &response, &vkErr)
	return
}

// WallSearchResponse struct
type WallSearchResponse struct {
	Count int                   `json:"count"`
	Items []object.WallWallpost `json:"items"`
}

// WallSearch allows to search posts on user or community walls.
//
// extended=0
//
// https://vk.com/dev/wall.search
func (vk *VK) WallSearch(params map[string]string) (response WallSearchResponse, vkErr Error) {
	params["extended"] = "0"
	vk.RequestUnmarshal("wall.search", params, &response, &vkErr)
	return
}

// WallSearchExtendedResponse struct
type WallSearchExtendedResponse struct {
	Count    int                   `json:"count"`
	Items    []object.WallWallpost `json:"items"`
	Profiles []object.UsersUser    `json:"profiles"`
	Groups   []object.GroupsGroup  `json:"groups"`
}

// WallSearchExtended allows to search posts on user or community walls.
//
// extended=1
//
// https://vk.com/dev/wall.search
func (vk *VK) WallSearchExtended(params map[string]string) (response WallSearchExtendedResponse, vkErr Error) {
	params["extended"] = "1"
	vk.RequestUnmarshal("wall.search", params, &response, &vkErr)
	return
}

// WallUnpin unpins the post on wall.
//
// https://vk.com/dev/wall.unpin
func (vk *VK) WallUnpin(params map[string]string) (response int, vkErr Error) {
	vk.RequestUnmarshal("wall.unpin", params, &response, &vkErr)
	return
}
