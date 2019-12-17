package api // import "github.com/SevereCloud/vksdk/api"

import (
	"github.com/SevereCloud/vksdk/object"
)

// WallCloseComments turn off post commenting
//
// https://vk.com/dev/wall.closeComments
func (vk *VK) WallCloseComments(params Params) (response int, err error) {
	err = vk.RequestUnmarshal("wall.closeComments", params, &response)
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
func (vk *VK) WallCreateComment(params Params) (response WallCreateCommentResponse, err error) {
	err = vk.RequestUnmarshal("wall.createComment", params, &response)
	return
}

// WallDelete deletes a post from a user wall or community wall.
//
// https://vk.com/dev/wall.delete
func (vk *VK) WallDelete(params Params) (response int, err error) {
	err = vk.RequestUnmarshal("wall.delete", params, &response)
	return
}

// WallDeleteComment deletes a comment on a post on a user wall or community wall.
//
// https://vk.com/dev/wall.deleteComment
func (vk *VK) WallDeleteComment(params Params) (response int, err error) {
	err = vk.RequestUnmarshal("wall.deleteComment", params, &response)
	return
}

type WallEditResponse struct {
	PostID int `json:"post_id"`
}

// WallEdit edits a post on a user wall or community wall.
//
// https://vk.com/dev/wall.edit
func (vk *VK) WallEdit(params Params) (response WallEditResponse, err error) {
	err = vk.RequestUnmarshal("wall.edit", params, &response)
	return
}

// WallEditAdsStealth allows to edit hidden post.
//
// https://vk.com/dev/wall.editAdsStealth
func (vk *VK) WallEditAdsStealth(params Params) (response int, err error) {
	err = vk.RequestUnmarshal("wall.editAdsStealth", params, &response)
	return
}

// WallEditComment edits a comment on a user wall or community wall.
//
// https://vk.com/dev/wall.editComment
func (vk *VK) WallEditComment(params Params) (response int, err error) {
	err = vk.RequestUnmarshal("wall.editComment", params, &response)
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
func (vk *VK) WallGet(params Params) (response WallGetResponse, err error) {
	params["extended"] = "0"
	err = vk.RequestUnmarshal("wall.get", params, &response)

	return
}

// WallGetExtendedResponse struct
type WallGetExtendedResponse struct {
	Count int                   `json:"count"`
	Items []object.WallWallpost `json:"items"`
	object.ExtendedResponse
}

// WallGetExtended returns a list of posts on a user wall or community wall.
//
// extended=1
//
// https://vk.com/dev/wall.get
func (vk *VK) WallGetExtended(params Params) (response WallGetExtendedResponse, err error) {
	params["extended"] = true
	err = vk.RequestUnmarshal("wall.get", params, &response)

	return
}

// WallGetByIDResponse struct
type WallGetByIDResponse []object.WallWallpost

// WallGetByID returns a list of posts from user or community walls by their IDs.
//
// extended=0
//
// https://vk.com/dev/wall.getById
func (vk *VK) WallGetByID(params Params) (response WallGetByIDResponse, err error) {
	params["extended"] = "0"
	err = vk.RequestUnmarshal("wall.getById", params, &response)

	return
}

// WallGetByIDExtendedResponse struct
type WallGetByIDExtendedResponse struct {
	Items []object.WallWallpost `json:"items"`
	object.ExtendedResponse
}

// WallGetByIDExtended returns a list of posts from user or community walls by their IDs.
//
// extended=1
//
// https://vk.com/dev/wall.getById
func (vk *VK) WallGetByIDExtended(params Params) (response WallGetByIDExtendedResponse, err error) {
	params["extended"] = true
	err = vk.RequestUnmarshal("wall.getById", params, &response)

	return
}

// WallGetCommentResponse struct
type WallGetCommentResponse struct {
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
func (vk *VK) WallGetComment(params Params) (response WallGetCommentResponse, err error) {
	params["extended"] = "0"
	err = vk.RequestUnmarshal("wall.getComment", params, &response)

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
func (vk *VK) WallGetCommentExtended(params Params) (response WallGetCommentExtendedResponse, err error) {
	params["extended"] = true
	err = vk.RequestUnmarshal("wall.getComment", params, &response)

	return
}

// WallGetCommentsResponse struct
type WallGetCommentsResponse struct {
	CanPost           bool                     `json:"can_post"`
	ShowReplyButton   bool                     `json:"show_reply_button"`
	GroupsCanPost     bool                     `json:"groups_can_post"`
	CurrentLevelCount int                      `json:"current_level_count"`
	Count             int                      `json:"count"`
	Items             []object.WallWallComment `json:"items"`
}

// WallGetComments returns a list of comments on a post on a user wall or community wall.
//
// extended=0
//
// https://vk.com/dev/wall.getComments
func (vk *VK) WallGetComments(params Params) (response WallGetCommentsResponse, err error) {
	params["extended"] = "0"
	err = vk.RequestUnmarshal("wall.getComments", params, &response)

	return
}

// WallGetCommentsExtendedResponse struct
type WallGetCommentsExtendedResponse struct {
	CanPost           bool                     `json:"can_post"`
	ShowReplyButton   bool                     `json:"show_reply_button"`
	GroupsCanPost     bool                     `json:"groups_can_post"`
	CurrentLevelCount int                      `json:"current_level_count"`
	Count             int                      `json:"count"`
	Items             []object.WallWallComment `json:"items"`
	object.ExtendedResponse
}

// WallGetCommentsExtended returns a list of comments on a post on a user wall or community wall.
//
// extended=1
//
// https://vk.com/dev/wall.getComments
func (vk *VK) WallGetCommentsExtended(params Params) (response WallGetCommentsExtendedResponse, err error) {
	params["extended"] = true
	err = vk.RequestUnmarshal("wall.getComments", params, &response)

	return
}

// WallGetRepostsResponse struct
type WallGetRepostsResponse struct {
	Items []object.WallWallpost `json:"items"`
	object.ExtendedResponse
}

// WallGetReposts returns information about reposts of a post on user wall or community wall.
//
// https://vk.com/dev/wall.getReposts
func (vk *VK) WallGetReposts(params Params) (response WallGetRepostsResponse, err error) {
	err = vk.RequestUnmarshal("wall.getReposts", params, &response)
	return
}

// WallOpenComments includes posting comments.
//
// https://vk.com/dev/wall.openComments
func (vk *VK) WallOpenComments(params Params) (response int, err error) {
	err = vk.RequestUnmarshal("wall.openComments", params, &response)
	return
}

// WallPin pins the post on wall.
//
// https://vk.com/dev/wall.pin
func (vk *VK) WallPin(params Params) (response int, err error) {
	err = vk.RequestUnmarshal("wall.pin", params, &response)
	return
}

// WallPostResponse struct
type WallPostResponse struct {
	PostID int `json:"post_id"`
}

// WallPost adds a new post on a user wall or community wall.Can also be used to publish suggested or scheduled posts.
//
// https://vk.com/dev/wall.post
func (vk *VK) WallPost(params Params) (response WallPostResponse, err error) {
	err = vk.RequestUnmarshal("wall.post", params, &response)
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
func (vk *VK) WallPostAdsStealth(params Params) (response WallPostAdsStealthResponse, err error) {
	err = vk.RequestUnmarshal("wall.postAdsStealth", params, &response)
	return
}

// WallReportComment reports (submits a complaint about) a comment on a post on a user wall or community wall.
//
// https://vk.com/dev/wall.reportComment
func (vk *VK) WallReportComment(params Params) (response int, err error) {
	err = vk.RequestUnmarshal("wall.reportComment", params, &response)
	return
}

// WallReportPost reports (submits a complaint about) a post on a user wall or community wall.
//
// https://vk.com/dev/wall.reportPost
func (vk *VK) WallReportPost(params Params) (response int, err error) {
	err = vk.RequestUnmarshal("wall.reportPost", params, &response)
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
func (vk *VK) WallRepost(params Params) (response WallRepostResponse, err error) {
	err = vk.RequestUnmarshal("wall.repost", params, &response)
	return
}

// WallRestore restores a post deleted from a user wall or community wall.
//
// https://vk.com/dev/wall.restore
func (vk *VK) WallRestore(params Params) (response int, err error) {
	err = vk.RequestUnmarshal("wall.restore", params, &response)
	return
}

// WallRestoreComment restores a comment deleted from a user wall or community wall.
//
// https://vk.com/dev/wall.restoreComment
func (vk *VK) WallRestoreComment(params Params) (response int, err error) {
	err = vk.RequestUnmarshal("wall.restoreComment", params, &response)
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
func (vk *VK) WallSearch(params Params) (response WallSearchResponse, err error) {
	params["extended"] = "0"
	err = vk.RequestUnmarshal("wall.search", params, &response)

	return
}

// WallSearchExtendedResponse struct
type WallSearchExtendedResponse struct {
	Count int                   `json:"count"`
	Items []object.WallWallpost `json:"items"`
	object.ExtendedResponse
}

// WallSearchExtended allows to search posts on user or community walls.
//
// extended=1
//
// https://vk.com/dev/wall.search
func (vk *VK) WallSearchExtended(params Params) (response WallSearchExtendedResponse, err error) {
	params["extended"] = true
	err = vk.RequestUnmarshal("wall.search", params, &response)

	return
}

// WallUnpin unpins the post on wall.
//
// https://vk.com/dev/wall.unpin
func (vk *VK) WallUnpin(params Params) (response int, err error) {
	err = vk.RequestUnmarshal("wall.unpin", params, &response)
	return
}
