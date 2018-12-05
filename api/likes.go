package api

import "encoding/json"

// LikesAddResponse struct
type LikesAddResponse struct {
	Likes int `json:"likes"`
}

// LikesAdd adds the specified object to the Likes list of the current user.
func (vk *VK) LikesAdd(params map[string]string) (response LikesAddResponse, vkErr Error) {
	rawResponse, vkErr := vk.Request("likes.add", params)
	if vkErr.Code != 0 {
		return
	}

	err := json.Unmarshal(rawResponse, &response)
	if err != nil {
		panic(err)
	}

	return
}

// LikesDeleteResponse struct
type LikesDeleteResponse struct {
	Likes int `json:"likes"`
}

// LikesDelete deletes the specified object from the Likes list of the current user.
func (vk *VK) LikesDelete(params map[string]string) (response LikesDeleteResponse, vkErr Error) {
	rawResponse, vkErr := vk.Request("likes.delete", params)
	if vkErr.Code != 0 {
		return
	}

	err := json.Unmarshal(rawResponse, &response)
	if err != nil {
		panic(err)
	}

	return
}

// LikesGetListResponse struct
type LikesGetListResponse struct{}

// LikesGetList likes.getList Returns a list of IDs of users who added the specified object to their Likes list.
// TODO: params["extended"] = "1"
// func (vk *VK) LikesGetList(params map[string]string) (response LikesGetListResponse, vkErr Error) {
// 	rawResponse, vkErr := vk.Request("likes.getList", params)
// 	if err != nil {
// 		return
// 	}

// 	err := json.Unmarshal(rawResponse, &response)
// 	if err != nil {
// 		panic(err)
// 	}

// 	return
// }

// LikesIsLikedResponse struct
type LikesIsLikedResponse struct {
	Liked  int `json:"liked"`
	Copied int `json:"copied"`
}

// LikesIsLiked checks for the object in the Likes list of the specified user.
func (vk *VK) LikesIsLiked(params map[string]string) (response LikesIsLikedResponse, vkErr Error) {
	rawResponse, vkErr := vk.Request("likes.isLiked", params)
	if vkErr.Code != 0 {
		return
	}

	err := json.Unmarshal(rawResponse, &response)
	if err != nil {
		panic(err)
	}

	return
}
