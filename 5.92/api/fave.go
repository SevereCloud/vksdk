package api // import "github.com/SevereCloud/vksdk/5.92/api"

import (
	"github.com/SevereCloud/vksdk/5.92/object"
)

// FaveAddArticleResponse struct
type FaveAddArticleResponse int

// FaveAddArticle adds a link to user faves.
//
// https://vk.com/dev/fave.addArticle
func (vk *VK) FaveAddArticle(params map[string]string) (response FaveAddArticleResponse, vkErr Error) {
	vk.RequestUnmarshal("fave.addArticle", params, &response, &vkErr)
	return
}

// FaveAddLinkResponse struct
type FaveAddLinkResponse int

// FaveAddLink adds a link to user faves.
//
// https://vk.com/dev/fave.addLink
func (vk *VK) FaveAddLink(params map[string]string) (response FaveAddLinkResponse, vkErr Error) {
	vk.RequestUnmarshal("fave.addLink", params, &response, &vkErr)
	return
}

// FaveAddPageResponse struct
type FaveAddPageResponse int

// FaveAddPage
//
// https://vk.com/dev/fave.addPage
func (vk *VK) FaveAddPage(params map[string]string) (response FaveAddPageResponse, vkErr Error) {
	vk.RequestUnmarshal("fave.addPage", params, &response, &vkErr)
	return
}

// FaveAddPostResponse struct
type FaveAddPostResponse int

// FaveAddPost
//
// https://vk.com/dev/fave.addPost
func (vk *VK) FaveAddPost(params map[string]string) (response FaveAddPostResponse, vkErr Error) {
	vk.RequestUnmarshal("fave.addPost", params, &response, &vkErr)
	return
}

// FaveAddProductResponse struct
type FaveAddProductResponse int

// FaveAddProduct
//
// https://vk.com/dev/fave.addProduct
func (vk *VK) FaveAddProduct(params map[string]string) (response FaveAddProductResponse, vkErr Error) {
	vk.RequestUnmarshal("fave.addProduct", params, &response, &vkErr)
	return
}

// FaveAddTagResponse struct
type FaveAddTagResponse object.FaveTag

// FaveAddTag
//
// https://vk.com/dev/fave.addTag
func (vk *VK) FaveAddTag(params map[string]string) (response FaveAddTagResponse, vkErr Error) {
	vk.RequestUnmarshal("fave.addTag", params, &response, &vkErr)
	return
}

// FaveAddVideoResponse struct
type FaveAddVideoResponse int

// FaveAddVideo
//
// https://vk.com/dev/fave.addVideo
func (vk *VK) FaveAddVideo(params map[string]string) (response FaveAddVideoResponse, vkErr Error) {
	vk.RequestUnmarshal("fave.addVideo", params, &response, &vkErr)
	return
}

// FaveEditTagResponse struct
type FaveEditTagResponse int

// FaveEditTag
//
// https://vk.com/dev/fave.editTag
func (vk *VK) FaveEditTag(params map[string]string) (response FaveEditTagResponse, vkErr Error) {
	vk.RequestUnmarshal("fave.editTag", params, &response, &vkErr)
	return
}

// FaveGetResponse struct
type FaveGetResponse struct {
	Count int               `json:"count"`
	Items []object.FaveItem `json:"items"`
}

// FaveGet
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

// FaveGetExtended
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

// FaveGetPages
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

// FaveGetTags
//
// https://vk.com/dev/fave.getTags
func (vk *VK) FaveGetTags(params map[string]string) (response FaveGetTagsResponse, vkErr Error) {
	vk.RequestUnmarshal("fave.getTags", params, &response, &vkErr)
	return
}

// FaveMarkSeenResponse struct
type FaveMarkSeenResponse int

// FaveMarkSeen
//
// https://vk.com/dev/fave.markSeen
func (vk *VK) FaveMarkSeen(params map[string]string) (response FaveMarkSeenResponse, vkErr Error) {
	vk.RequestUnmarshal("fave.markSeen", params, &response, &vkErr)
	return
}

// FaveRemoveArticleResponse struct
type FaveRemoveArticleResponse int

// FaveRemoveArticle
//
// https://vk.com/dev/fave.removeArticle
func (vk *VK) FaveRemoveArticle(params map[string]string) (response FaveRemoveArticleResponse, vkErr Error) {
	vk.RequestUnmarshal("fave.removeArticle", params, &response, &vkErr)
	return
}

// FaveRemoveLinkResponse struct
type FaveRemoveLinkResponse int

// FaveRemoveLink removes link from the user's faves.
//
// https://vk.com/dev/fave.removeLink
func (vk *VK) FaveRemoveLink(params map[string]string) (response FaveRemoveLinkResponse, vkErr Error) {
	vk.RequestUnmarshal("fave.removeLink", params, &response, &vkErr)
	return
}

// FaveRemovePageResponse struct
type FaveRemovePageResponse int

// FaveRemovePage
//
// https://vk.com/dev/fave.removePage
func (vk *VK) FaveRemovePage(params map[string]string) (response FaveRemovePageResponse, vkErr Error) {
	vk.RequestUnmarshal("fave.removePage", params, &response, &vkErr)
	return
}

// FaveRemovePostResponse struct
type FaveRemovePostResponse int

// FaveRemovePost
//
// https://vk.com/dev/fave.removePost
func (vk *VK) FaveRemovePost(params map[string]string) (response FaveRemovePostResponse, vkErr Error) {
	vk.RequestUnmarshal("fave.removePost", params, &response, &vkErr)
	return
}

// FaveRemoveProductResponse struct
type FaveRemoveProductResponse int

// FaveRemoveProduct
//
// https://vk.com/dev/fave.removeProduct
func (vk *VK) FaveRemoveProduct(params map[string]string) (response FaveRemoveProductResponse, vkErr Error) {
	vk.RequestUnmarshal("fave.removeProduct", params, &response, &vkErr)
	return
}

// FaveRemoveTagResponse struct
type FaveRemoveTagResponse int

// FaveRemoveTag
//
// https://vk.com/dev/fave.removeTag
func (vk *VK) FaveRemoveTag(params map[string]string) (response FaveRemoveTagResponse, vkErr Error) {
	vk.RequestUnmarshal("fave.removeTag", params, &response, &vkErr)
	return
}

// FaveRemoveVideoResponse struct
type FaveRemoveVideoResponse int

// FaveRemoveVideo
//
// https://vk.com/dev/fave.removeVideo
func (vk *VK) FaveRemoveVideo(params map[string]string) (response FaveRemoveVideoResponse, vkErr Error) {
	vk.RequestUnmarshal("fave.removeVideo", params, &response, &vkErr)
	return
}

// FaveReorderTagsResponse struct
type FaveReorderTagsResponse int

// FaveReorderTags
//
// https://vk.com/dev/fave.reorderTags
func (vk *VK) FaveReorderTags(params map[string]string) (response FaveReorderTagsResponse, vkErr Error) {
	vk.RequestUnmarshal("fave.reorderTags", params, &response, &vkErr)
	return
}

// FaveSetPageTagsResponse struct
type FaveSetPageTagsResponse int

// FaveSetPageTags
//
// https://vk.com/dev/fave.setPageTags
func (vk *VK) FaveSetPageTags(params map[string]string) (response FaveSetPageTagsResponse, vkErr Error) {
	vk.RequestUnmarshal("fave.setPageTags", params, &response, &vkErr)
	return
}

// FaveSetTagsResponse struct
type FaveSetTagsResponse int

// FaveSetTags
//
// https://vk.com/dev/fave.setTags
func (vk *VK) FaveSetTags(params map[string]string) (response FaveSetTagsResponse, vkErr Error) {
	vk.RequestUnmarshal("fave.setTags", params, &response, &vkErr)
	return
}

// FaveTrackPageInteractionResponse struct
type FaveTrackPageInteractionResponse int

// FaveTrackPageInteraction
//
// https://vk.com/dev/fave.trackPageInteraction
func (vk *VK) FaveTrackPageInteraction(params map[string]string) (response FaveTrackPageInteractionResponse, vkErr Error) {
	vk.RequestUnmarshal("fave.trackPageInteraction", params, &response, &vkErr)
	return
}
