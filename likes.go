package vksdk

import "encoding/json"

// LikesAddResponse struct
type LikesAddResponse struct {
	Likes int `json:"likes,omitempty"`
}

// LikesAdd adds the specified object to the Likes list of the current user.
func (vk *VK) LikesAdd(params map[string]string) (response LikesAddResponse, err error) {
	rawResponse, err := vk.Request("likes.add", params)
	if err != nil {
		return
	}

	err = json.Unmarshal(rawResponse, &response)
	if err != nil {
		panic(err)
	}

	return
}

// LikesDeleteResponse struct
type LikesDeleteResponse struct {
	Likes int `json:"likes,omitempty"`
}

// LikesDelete deletes the specified object from the Likes list of the current user.
func (vk *VK) LikesDelete(params map[string]string) (response LikesDeleteResponse, err error) {
	rawResponse, err := vk.Request("likes.delete", params)
	if err != nil {
		return
	}

	err = json.Unmarshal(rawResponse, &response)
	if err != nil {
		panic(err)
	}

	return
}

// LikesGetListResponse struct
type LikesGetListResponse struct{}

// TODO likes.getList Returns a list of IDs of users who added the specified object to their Likes list.
// params["extended"] = "1"

// LikesIsLikedResponse struct
type LikesIsLikedResponse struct {
	Liked  int `json:"liked,omitempty"`
	Copied int `json:"copied,omitempty"`
}

// LikesIsLiked checks for the object in the Likes list of the specified user.
func (vk *VK) LikesIsLiked(params map[string]string) (response LikesIsLikedResponse, err error) {
	rawResponse, err := vk.Request("likes.isLiked", params)
	if err != nil {
		return
	}

	err = json.Unmarshal(rawResponse, &response)
	if err != nil {
		panic(err)
	}

	return
}
