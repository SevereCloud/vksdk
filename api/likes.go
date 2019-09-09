package api // import "github.com/SevereCloud/vksdk/api"

import (
	"github.com/SevereCloud/vksdk/object"
)

// LikesAddResponse struct
type LikesAddResponse struct {
	Likes int `json:"likes"`
}

// LikesAdd adds the specified object to the Likes list of the current user.
//
// https://vk.com/dev/likes.add
func (vk *VK) LikesAdd(params map[string]string) (response LikesAddResponse, vkErr Error) {
	vk.RequestUnmarshal("likes.add", params, &response, &vkErr)
	return
}

// LikesDeleteResponse struct
type LikesDeleteResponse struct {
	Likes int `json:"likes"`
}

// LikesDelete deletes the specified object from the Likes list of the current user.
//
// https://vk.com/dev/likes.delete
func (vk *VK) LikesDelete(params map[string]string) (response LikesDeleteResponse, vkErr Error) {
	vk.RequestUnmarshal("likes.delete", params, &response, &vkErr)
	return
}

// LikesGetListResponse struct
type LikesGetListResponse struct {
	Count int   `json:"count"`
	Items []int `json:"items"`
}

// LikesGetList likes.getList Returns a list of IDs of users who added the specified object to their Likes list.
//
// extended=0
//
// https://vk.com/dev/likes.getList
func (vk *VK) LikesGetList(params map[string]string) (response LikesGetListResponse, vkErr Error) {
	params["extended"] = "0"
	vk.RequestUnmarshal("likes.getList", params, &response, &vkErr)
	return
}

// LikesGetListExtendedResponse struct
type LikesGetListExtendedResponse struct {
	Count int                `json:"count"`
	Items []object.UsersUser `json:"items"`
}

// LikesGetListExtended likes.getList Returns a list of IDs of users who added the specified object to their Likes list.
//
// extended=1
//
// https://vk.com/dev/likes.getList
func (vk *VK) LikesGetListExtended(params map[string]string) (response LikesGetListExtendedResponse, vkErr Error) {
	params["extended"] = "1"
	vk.RequestUnmarshal("likes.getList", params, &response, &vkErr)
	return
}

// LikesIsLikedResponse struct
type LikesIsLikedResponse struct {
	Liked  int `json:"liked"`
	Copied int `json:"copied"`
}

// LikesIsLiked checks for the object in the Likes list of the specified user.
//
// https://vk.com/dev/likes.isLiked
func (vk *VK) LikesIsLiked(params map[string]string) (response LikesIsLikedResponse, vkErr Error) {
	vk.RequestUnmarshal("likes.isLiked", params, &response, &vkErr)
	return
}
