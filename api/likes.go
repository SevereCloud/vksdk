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
func (vk *VK) LikesAdd(params map[string]string) (response LikesAddResponse, err error) {
	err = vk.RequestUnmarshal("likes.add", params, &response)
	return
}

// LikesDeleteResponse struct
type LikesDeleteResponse struct {
	Likes int `json:"likes"`
}

// LikesDelete deletes the specified object from the Likes list of the current user.
//
// https://vk.com/dev/likes.delete
func (vk *VK) LikesDelete(params map[string]string) (response LikesDeleteResponse, err error) {
	err = vk.RequestUnmarshal("likes.delete", params, &response)
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
func (vk *VK) LikesGetList(params map[string]string) (response LikesGetListResponse, err error) {
	params["extended"] = "0"
	err = vk.RequestUnmarshal("likes.getList", params, &response)

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
func (vk *VK) LikesGetListExtended(params map[string]string) (response LikesGetListExtendedResponse, err error) {
	params["extended"] = "1"
	err = vk.RequestUnmarshal("likes.getList", params, &response)

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
func (vk *VK) LikesIsLiked(params map[string]string) (response LikesIsLikedResponse, err error) {
	err = vk.RequestUnmarshal("likes.isLiked", params, &response)
	return
}
