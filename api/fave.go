package api // import "github.com/SevereCloud/vksdk/api"

import (
	"github.com/SevereCloud/vksdk/object"
)

// FaveAddArticle adds a link to user faves.
//
// https://vk.com/dev/fave.addArticle
func (vk *VK) FaveAddArticle(params Params) (response int, err error) {
	err = vk.RequestUnmarshal("fave.addArticle", params, &response)
	return
}

// FaveAddLink adds a link to user faves.
//
// https://vk.com/dev/fave.addLink
func (vk *VK) FaveAddLink(params Params) (response int, err error) {
	err = vk.RequestUnmarshal("fave.addLink", params, &response)
	return
}

// FaveAddPage method
//
// https://vk.com/dev/fave.addPage
func (vk *VK) FaveAddPage(params Params) (response int, err error) {
	err = vk.RequestUnmarshal("fave.addPage", params, &response)
	return
}

// FaveAddPost method
//
// https://vk.com/dev/fave.addPost
func (vk *VK) FaveAddPost(params Params) (response int, err error) {
	err = vk.RequestUnmarshal("fave.addPost", params, &response)
	return
}

// FaveAddProduct method
//
// https://vk.com/dev/fave.addProduct
func (vk *VK) FaveAddProduct(params Params) (response int, err error) {
	err = vk.RequestUnmarshal("fave.addProduct", params, &response)
	return
}

// FaveAddTagResponse struct
type FaveAddTagResponse object.FaveTag

// FaveAddTag method
//
// https://vk.com/dev/fave.addTag
func (vk *VK) FaveAddTag(params Params) (response FaveAddTagResponse, err error) {
	err = vk.RequestUnmarshal("fave.addTag", params, &response)
	return
}

// FaveAddVideo method
//
// https://vk.com/dev/fave.addVideo
func (vk *VK) FaveAddVideo(params Params) (response int, err error) {
	err = vk.RequestUnmarshal("fave.addVideo", params, &response)
	return
}

// FaveEditTag method
//
// https://vk.com/dev/fave.editTag
func (vk *VK) FaveEditTag(params Params) (response int, err error) {
	err = vk.RequestUnmarshal("fave.editTag", params, &response)
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
func (vk *VK) FaveGet(params Params) (response FaveGetResponse, err error) {
	params["extended"] = "0"
	err = vk.RequestUnmarshal("fave.get", params, &response)

	return
}

// FaveGetExtendedResponse struct
type FaveGetExtendedResponse struct {
	Count int               `json:"count"`
	Items []object.FaveItem `json:"items"`
	object.ExtendedResponse
}

// FaveGetExtended method
//
// extended=1
//
// https://vk.com/dev/fave.get
func (vk *VK) FaveGetExtended(params Params) (response FaveGetExtendedResponse, err error) {
	params["extended"] = true
	err = vk.RequestUnmarshal("fave.get", params, &response)

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
func (vk *VK) FaveGetPages(params Params) (response FaveGetPagesResponse, err error) {
	err = vk.RequestUnmarshal("fave.getPages", params, &response)
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
func (vk *VK) FaveGetTags(params Params) (response FaveGetTagsResponse, err error) {
	err = vk.RequestUnmarshal("fave.getTags", params, &response)
	return
}

// FaveMarkSeen method
//
// https://vk.com/dev/fave.markSeen
func (vk *VK) FaveMarkSeen(params Params) (response int, err error) {
	err = vk.RequestUnmarshal("fave.markSeen", params, &response)
	return
}

// FaveRemoveArticle method
//
// https://vk.com/dev/fave.removeArticle
func (vk *VK) FaveRemoveArticle(params Params) (response int, err error) {
	err = vk.RequestUnmarshal("fave.removeArticle", params, &response)
	return
}

// FaveRemoveLink removes link from the user's faves.
//
// https://vk.com/dev/fave.removeLink
func (vk *VK) FaveRemoveLink(params Params) (response int, err error) {
	err = vk.RequestUnmarshal("fave.removeLink", params, &response)
	return
}

// FaveRemovePage method
//
// https://vk.com/dev/fave.removePage
func (vk *VK) FaveRemovePage(params Params) (response int, err error) {
	err = vk.RequestUnmarshal("fave.removePage", params, &response)
	return
}

// FaveRemovePost method
//
// https://vk.com/dev/fave.removePost
func (vk *VK) FaveRemovePost(params Params) (response int, err error) {
	err = vk.RequestUnmarshal("fave.removePost", params, &response)
	return
}

// FaveRemoveProduct method
//
// https://vk.com/dev/fave.removeProduct
func (vk *VK) FaveRemoveProduct(params Params) (response int, err error) {
	err = vk.RequestUnmarshal("fave.removeProduct", params, &response)
	return
}

// FaveRemoveTag method
//
// https://vk.com/dev/fave.removeTag
func (vk *VK) FaveRemoveTag(params Params) (response int, err error) {
	err = vk.RequestUnmarshal("fave.removeTag", params, &response)
	return
}

// FaveRemoveVideo method
//
// https://vk.com/dev/fave.removeVideo
func (vk *VK) FaveRemoveVideo(params Params) (response int, err error) {
	err = vk.RequestUnmarshal("fave.removeVideo", params, &response)
	return
}

// FaveReorderTags method
//
// https://vk.com/dev/fave.reorderTags
func (vk *VK) FaveReorderTags(params Params) (response int, err error) {
	err = vk.RequestUnmarshal("fave.reorderTags", params, &response)
	return
}

// FaveSetPageTags method
//
// https://vk.com/dev/fave.setPageTags
func (vk *VK) FaveSetPageTags(params Params) (response int, err error) {
	err = vk.RequestUnmarshal("fave.setPageTags", params, &response)
	return
}

// FaveSetTags method
//
// https://vk.com/dev/fave.setTags
func (vk *VK) FaveSetTags(params Params) (response int, err error) {
	err = vk.RequestUnmarshal("fave.setTags", params, &response)
	return
}

// FaveTrackPageInteraction method
//
// https://vk.com/dev/fave.trackPageInteraction
func (vk *VK) FaveTrackPageInteraction(params Params) (response int, err error) {
	err = vk.RequestUnmarshal("fave.trackPageInteraction", params, &response)
	return
}
