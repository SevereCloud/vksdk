package api // import "github.com/SevereCloud/vksdk/api"

import (
	"github.com/SevereCloud/vksdk/object"
)

// FaveAddArticle adds a link to user faves.
//
// https://vk.com/dev/fave.addArticle
func (vk *VK) FaveAddArticle(params map[string]string) (response int, vkErr Error) {
	vk.RequestUnmarshal("fave.addArticle", params, &response, &vkErr)
	return
}

// FaveAddLink adds a link to user faves.
//
// https://vk.com/dev/fave.addLink
func (vk *VK) FaveAddLink(params map[string]string) (response int, vkErr Error) {
	vk.RequestUnmarshal("fave.addLink", params, &response, &vkErr)
	return
}

// FaveAddPage method
//
// https://vk.com/dev/fave.addPage
func (vk *VK) FaveAddPage(params map[string]string) (response int, vkErr Error) {
	vk.RequestUnmarshal("fave.addPage", params, &response, &vkErr)
	return
}

// FaveAddPost method
//
// https://vk.com/dev/fave.addPost
func (vk *VK) FaveAddPost(params map[string]string) (response int, vkErr Error) {
	vk.RequestUnmarshal("fave.addPost", params, &response, &vkErr)
	return
}

// FaveAddProduct method
//
// https://vk.com/dev/fave.addProduct
func (vk *VK) FaveAddProduct(params map[string]string) (response int, vkErr Error) {
	vk.RequestUnmarshal("fave.addProduct", params, &response, &vkErr)
	return
}

// FaveAddTagResponse struct
type FaveAddTagResponse object.FaveTag

// FaveAddTag method
//
// https://vk.com/dev/fave.addTag
func (vk *VK) FaveAddTag(params map[string]string) (response FaveAddTagResponse, vkErr Error) {
	vk.RequestUnmarshal("fave.addTag", params, &response, &vkErr)
	return
}

// FaveAddVideo method
//
// https://vk.com/dev/fave.addVideo
func (vk *VK) FaveAddVideo(params map[string]string) (response int, vkErr Error) {
	vk.RequestUnmarshal("fave.addVideo", params, &response, &vkErr)
	return
}

// FaveEditTag method
//
// https://vk.com/dev/fave.editTag
func (vk *VK) FaveEditTag(params map[string]string) (response int, vkErr Error) {
	vk.RequestUnmarshal("fave.editTag", params, &response, &vkErr)
	return
}

// FaveGetResponse struct
type FaveGetResponse struct {
	Count int               `json:"count"`
	Items []object.FaveItem `json:"items"`
}

// FaveGet method
//
// extended=0
//
// https://vk.com/dev/fave.get
func (vk *VK) FaveGet(params map[string]string) (response FaveGetResponse, vkErr Error) {
	params["extended"] = "0"
	vk.RequestUnmarshal("fave.get", params, &response, &vkErr)
	return
}

// FaveGetExtendedResponse struct
type FaveGetExtendedResponse struct {
	Count    int                  `json:"count"`
	Items    []object.FaveItem    `json:"items"`
	Profiles []object.UsersUser   `json:"profiles"`
	Groups   []object.GroupsGroup `json:"groups"`
}

// FaveGetExtended method
//
// extended=1
//
// https://vk.com/dev/fave.get
func (vk *VK) FaveGetExtended(params map[string]string) (response FaveGetExtendedResponse, vkErr Error) {
	params["extended"] = "1"
	vk.RequestUnmarshal("fave.get", params, &response, &vkErr)
	return
}

// FaveGetPagesResponse struct
type FaveGetPagesResponse struct {
	Count int               `json:"count"`
	Items []object.FavePage `json:"items"`
}

// FaveGetPages method
//
// https://vk.com/dev/fave.getPages
func (vk *VK) FaveGetPages(params map[string]string) (response FaveGetPagesResponse, vkErr Error) {
	vk.RequestUnmarshal("fave.getPages", params, &response, &vkErr)
	return
}

// FaveGetTagsResponse struct
type FaveGetTagsResponse struct {
	Count int              `json:"count"`
	Items []object.FaveTag `json:"items"`
}

// FaveGetTags method
//
// https://vk.com/dev/fave.getTags
func (vk *VK) FaveGetTags(params map[string]string) (response FaveGetTagsResponse, vkErr Error) {
	vk.RequestUnmarshal("fave.getTags", params, &response, &vkErr)
	return
}

// FaveMarkSeen method
//
// https://vk.com/dev/fave.markSeen
func (vk *VK) FaveMarkSeen(params map[string]string) (response int, vkErr Error) {
	vk.RequestUnmarshal("fave.markSeen", params, &response, &vkErr)
	return
}

// FaveRemoveArticle method
//
// https://vk.com/dev/fave.removeArticle
func (vk *VK) FaveRemoveArticle(params map[string]string) (response int, vkErr Error) {
	vk.RequestUnmarshal("fave.removeArticle", params, &response, &vkErr)
	return
}

// FaveRemoveLink removes link from the user's faves.
//
// https://vk.com/dev/fave.removeLink
func (vk *VK) FaveRemoveLink(params map[string]string) (response int, vkErr Error) {
	vk.RequestUnmarshal("fave.removeLink", params, &response, &vkErr)
	return
}

// FaveRemovePage method
//
// https://vk.com/dev/fave.removePage
func (vk *VK) FaveRemovePage(params map[string]string) (response int, vkErr Error) {
	vk.RequestUnmarshal("fave.removePage", params, &response, &vkErr)
	return
}

// FaveRemovePost method
//
// https://vk.com/dev/fave.removePost
func (vk *VK) FaveRemovePost(params map[string]string) (response int, vkErr Error) {
	vk.RequestUnmarshal("fave.removePost", params, &response, &vkErr)
	return
}

// FaveRemoveProduct method
//
// https://vk.com/dev/fave.removeProduct
func (vk *VK) FaveRemoveProduct(params map[string]string) (response int, vkErr Error) {
	vk.RequestUnmarshal("fave.removeProduct", params, &response, &vkErr)
	return
}

// FaveRemoveTag method
//
// https://vk.com/dev/fave.removeTag
func (vk *VK) FaveRemoveTag(params map[string]string) (response int, vkErr Error) {
	vk.RequestUnmarshal("fave.removeTag", params, &response, &vkErr)
	return
}

// FaveRemoveVideo method
//
// https://vk.com/dev/fave.removeVideo
func (vk *VK) FaveRemoveVideo(params map[string]string) (response int, vkErr Error) {
	vk.RequestUnmarshal("fave.removeVideo", params, &response, &vkErr)
	return
}

// FaveReorderTags method
//
// https://vk.com/dev/fave.reorderTags
func (vk *VK) FaveReorderTags(params map[string]string) (response int, vkErr Error) {
	vk.RequestUnmarshal("fave.reorderTags", params, &response, &vkErr)
	return
}

// FaveSetPageTags method
//
// https://vk.com/dev/fave.setPageTags
func (vk *VK) FaveSetPageTags(params map[string]string) (response int, vkErr Error) {
	vk.RequestUnmarshal("fave.setPageTags", params, &response, &vkErr)
	return
}

// FaveSetTags method
//
// https://vk.com/dev/fave.setTags
func (vk *VK) FaveSetTags(params map[string]string) (response int, vkErr Error) {
	vk.RequestUnmarshal("fave.setTags", params, &response, &vkErr)
	return
}

// FaveTrackPageInteraction method
//
// https://vk.com/dev/fave.trackPageInteraction
func (vk *VK) FaveTrackPageInteraction(params map[string]string) (response int, vkErr Error) {
	vk.RequestUnmarshal("fave.trackPageInteraction", params, &response, &vkErr)
	return
}
