package params // import "github.com/SevereCloud/vksdk/api/params"

import (
	"github.com/SevereCloud/vksdk/api"
)

// FaveAddArticleBulder builder
//
// https://vk.com/dev/fave.addArticle
type FaveAddArticleBulder struct {
	api.Params
}

// NewFaveAddArticleBulder func
func NewFaveAddArticleBulder() *FaveAddArticleBulder {
	return &FaveAddArticleBulder{api.Params{}}
}

// URL parameter
func (b *FaveAddArticleBulder) URL(v string) {
	b.Params["url"] = v
}

// FaveAddLinkBulder builder
//
// Adds a link to user faves.
//
// https://vk.com/dev/fave.addLink
type FaveAddLinkBulder struct {
	api.Params
}

// NewFaveAddLinkBulder func
func NewFaveAddLinkBulder() *FaveAddLinkBulder {
	return &FaveAddLinkBulder{api.Params{}}
}

// Link Link URL.
func (b *FaveAddLinkBulder) Link(v string) {
	b.Params["link"] = v
}

// FaveAddPageBulder builder
//
// https://vk.com/dev/fave.addPage
type FaveAddPageBulder struct {
	api.Params
}

// NewFaveAddPageBulder func
func NewFaveAddPageBulder() *FaveAddPageBulder {
	return &FaveAddPageBulder{api.Params{}}
}

// UserID parameter
func (b *FaveAddPageBulder) UserID(v int) {
	b.Params["user_id"] = v
}

// GroupID parameter
func (b *FaveAddPageBulder) GroupID(v int) {
	b.Params["group_id"] = v
}

// FaveAddPostBulder builder
//
// https://vk.com/dev/fave.addPost
type FaveAddPostBulder struct {
	api.Params
}

// NewFaveAddPostBulder func
func NewFaveAddPostBulder() *FaveAddPostBulder {
	return &FaveAddPostBulder{api.Params{}}
}

// OwnerID parameter
func (b *FaveAddPostBulder) OwnerID(v int) {
	b.Params["owner_id"] = v
}

// ID parameter
func (b *FaveAddPostBulder) ID(v int) {
	b.Params["id"] = v
}

// AccessKey parameter
func (b *FaveAddPostBulder) AccessKey(v string) {
	b.Params["access_key"] = v
}

// FaveAddProductBulder builder
//
// https://vk.com/dev/fave.addProduct
type FaveAddProductBulder struct {
	api.Params
}

// NewFaveAddProductBulder func
func NewFaveAddProductBulder() *FaveAddProductBulder {
	return &FaveAddProductBulder{api.Params{}}
}

// OwnerID parameter
func (b *FaveAddProductBulder) OwnerID(v int) {
	b.Params["owner_id"] = v
}

// ID parameter
func (b *FaveAddProductBulder) ID(v int) {
	b.Params["id"] = v
}

// AccessKey parameter
func (b *FaveAddProductBulder) AccessKey(v string) {
	b.Params["access_key"] = v
}

// FaveAddTagBulder builder
//
// https://vk.com/dev/fave.addTag
type FaveAddTagBulder struct {
	api.Params
}

// NewFaveAddTagBulder func
func NewFaveAddTagBulder() *FaveAddTagBulder {
	return &FaveAddTagBulder{api.Params{}}
}

// Name parameter
func (b *FaveAddTagBulder) Name(v string) {
	b.Params["name"] = v
}

// FaveAddVideoBulder builder
//
// https://vk.com/dev/fave.addVideo
type FaveAddVideoBulder struct {
	api.Params
}

// NewFaveAddVideoBulder func
func NewFaveAddVideoBulder() *FaveAddVideoBulder {
	return &FaveAddVideoBulder{api.Params{}}
}

// OwnerID parameter
func (b *FaveAddVideoBulder) OwnerID(v int) {
	b.Params["owner_id"] = v
}

// ID parameter
func (b *FaveAddVideoBulder) ID(v int) {
	b.Params["id"] = v
}

// AccessKey parameter
func (b *FaveAddVideoBulder) AccessKey(v string) {
	b.Params["access_key"] = v
}

// FaveEditTagBulder builder
//
// https://vk.com/dev/fave.editTag
type FaveEditTagBulder struct {
	api.Params
}

// NewFaveEditTagBulder func
func NewFaveEditTagBulder() *FaveEditTagBulder {
	return &FaveEditTagBulder{api.Params{}}
}

// ID parameter
func (b *FaveEditTagBulder) ID(v int) {
	b.Params["id"] = v
}

// Name parameter
func (b *FaveEditTagBulder) Name(v string) {
	b.Params["name"] = v
}

// FaveGetBulder builder
//
// https://vk.com/dev/fave.get
type FaveGetBulder struct {
	api.Params
}

// NewFaveGetBulder func
func NewFaveGetBulder() *FaveGetBulder {
	return &FaveGetBulder{api.Params{}}
}

// Extended '1' — to return additional 'wall', 'profiles', and 'groups' fields. By default: '0'.
func (b *FaveGetBulder) Extended(v bool) {
	b.Params["extended"] = v
}

// ItemType parameter
func (b *FaveGetBulder) ItemType(v string) {
	b.Params["item_type"] = v
}

// TagID Tag ID.
func (b *FaveGetBulder) TagID(v int) {
	b.Params["tag_id"] = v
}

// Offset Offset needed to return a specific subset of users.
func (b *FaveGetBulder) Offset(v int) {
	b.Params["offset"] = v
}

// Count Number of users to return.
func (b *FaveGetBulder) Count(v int) {
	b.Params["count"] = v
}

// Fields parameter
func (b *FaveGetBulder) Fields(v string) {
	b.Params["fields"] = v
}

// IsFromSnackbar parameter
func (b *FaveGetBulder) IsFromSnackbar(v bool) {
	b.Params["is_from_snackbar"] = v
}

// FaveGetPagesBulder builder
//
// https://vk.com/dev/fave.getPages
type FaveGetPagesBulder struct {
	api.Params
}

// NewFaveGetPagesBulder func
func NewFaveGetPagesBulder() *FaveGetPagesBulder {
	return &FaveGetPagesBulder{api.Params{}}
}

// Offset parameter
func (b *FaveGetPagesBulder) Offset(v int) {
	b.Params["offset"] = v
}

// Count parameter
func (b *FaveGetPagesBulder) Count(v int) {
	b.Params["count"] = v
}

// Type parameter
func (b *FaveGetPagesBulder) Type(v string) {
	b.Params["type"] = v
}

// Fields parameter
func (b *FaveGetPagesBulder) Fields(v []string) {
	b.Params["fields"] = v
}

// TagID parameter
func (b *FaveGetPagesBulder) TagID(v int) {
	b.Params["tag_id"] = v
}

// FaveRemoveArticleBulder builder
//
// https://vk.com/dev/fave.removeArticle
type FaveRemoveArticleBulder struct {
	api.Params
}

// NewFaveRemoveArticleBulder func
func NewFaveRemoveArticleBulder() *FaveRemoveArticleBulder {
	return &FaveRemoveArticleBulder{api.Params{}}
}

// OwnerID parameter
func (b *FaveRemoveArticleBulder) OwnerID(v int) {
	b.Params["owner_id"] = v
}

// ArticleID parameter
func (b *FaveRemoveArticleBulder) ArticleID(v int) {
	b.Params["article_id"] = v
}

// FaveRemoveLinkBulder builder
//
// Removes link from the user's faves.
//
// https://vk.com/dev/fave.removeLink
type FaveRemoveLinkBulder struct {
	api.Params
}

// NewFaveRemoveLinkBulder func
func NewFaveRemoveLinkBulder() *FaveRemoveLinkBulder {
	return &FaveRemoveLinkBulder{api.Params{}}
}

// LinkID Link ID (can be obtained by [vk.com/dev/faves.getLinks|faves.getLinks] method).
func (b *FaveRemoveLinkBulder) LinkID(v string) {
	b.Params["link_id"] = v
}

// Link Link URL
func (b *FaveRemoveLinkBulder) Link(v string) {
	b.Params["link"] = v
}

// FaveRemovePageBulder builder
//
// https://vk.com/dev/fave.removePage
type FaveRemovePageBulder struct {
	api.Params
}

// NewFaveRemovePageBulder func
func NewFaveRemovePageBulder() *FaveRemovePageBulder {
	return &FaveRemovePageBulder{api.Params{}}
}

// UserID parameter
func (b *FaveRemovePageBulder) UserID(v int) {
	b.Params["user_id"] = v
}

// GroupID parameter
func (b *FaveRemovePageBulder) GroupID(v int) {
	b.Params["group_id"] = v
}

// FaveRemovePostBulder builder
//
// https://vk.com/dev/fave.removePost
type FaveRemovePostBulder struct {
	api.Params
}

// NewFaveRemovePostBulder func
func NewFaveRemovePostBulder() *FaveRemovePostBulder {
	return &FaveRemovePostBulder{api.Params{}}
}

// OwnerID parameter
func (b *FaveRemovePostBulder) OwnerID(v int) {
	b.Params["owner_id"] = v
}

// ID parameter
func (b *FaveRemovePostBulder) ID(v int) {
	b.Params["id"] = v
}

// FaveRemoveProductBulder builder
//
// https://vk.com/dev/fave.removeProduct
type FaveRemoveProductBulder struct {
	api.Params
}

// NewFaveRemoveProductBulder func
func NewFaveRemoveProductBulder() *FaveRemoveProductBulder {
	return &FaveRemoveProductBulder{api.Params{}}
}

// OwnerID parameter
func (b *FaveRemoveProductBulder) OwnerID(v int) {
	b.Params["owner_id"] = v
}

// ID parameter
func (b *FaveRemoveProductBulder) ID(v int) {
	b.Params["id"] = v
}

// FaveRemoveTagBulder builder
//
// https://vk.com/dev/fave.removeTag
type FaveRemoveTagBulder struct {
	api.Params
}

// NewFaveRemoveTagBulder func
func NewFaveRemoveTagBulder() *FaveRemoveTagBulder {
	return &FaveRemoveTagBulder{api.Params{}}
}

// ID parameter
func (b *FaveRemoveTagBulder) ID(v int) {
	b.Params["id"] = v
}

// FaveReorderTagsBulder builder
//
// https://vk.com/dev/fave.reorderTags
type FaveReorderTagsBulder struct {
	api.Params
}

// NewFaveReorderTagsBulder func
func NewFaveReorderTagsBulder() *FaveReorderTagsBulder {
	return &FaveReorderTagsBulder{api.Params{}}
}

// IDs parameter
func (b *FaveReorderTagsBulder) IDs(v []int) {
	b.Params["ids"] = v
}

// FaveSetPageTagsBulder builder
//
// https://vk.com/dev/fave.setPageTags
type FaveSetPageTagsBulder struct {
	api.Params
}

// NewFaveSetPageTagsBulder func
func NewFaveSetPageTagsBulder() *FaveSetPageTagsBulder {
	return &FaveSetPageTagsBulder{api.Params{}}
}

// UserID parameter
func (b *FaveSetPageTagsBulder) UserID(v int) {
	b.Params["user_id"] = v
}

// GroupID parameter
func (b *FaveSetPageTagsBulder) GroupID(v int) {
	b.Params["group_id"] = v
}

// TagIDs parameter
func (b *FaveSetPageTagsBulder) TagIDs(v []int) {
	b.Params["tag_ids"] = v
}

// FaveSetTagsBulder builder
//
// https://vk.com/dev/fave.setTags
type FaveSetTagsBulder struct {
	api.Params
}

// NewFaveSetTagsBulder func
func NewFaveSetTagsBulder() *FaveSetTagsBulder {
	return &FaveSetTagsBulder{api.Params{}}
}

// ItemType parameter
func (b *FaveSetTagsBulder) ItemType(v string) {
	b.Params["item_type"] = v
}

// ItemOwnerID parameter
func (b *FaveSetTagsBulder) ItemOwnerID(v int) {
	b.Params["item_owner_id"] = v
}

// ItemID parameter
func (b *FaveSetTagsBulder) ItemID(v int) {
	b.Params["item_id"] = v
}

// TagIDs parameter
func (b *FaveSetTagsBulder) TagIDs(v []int) {
	b.Params["tag_ids"] = v
}

// LinkID parameter
func (b *FaveSetTagsBulder) LinkID(v string) {
	b.Params["link_id"] = v
}

// LinkURL parameter
func (b *FaveSetTagsBulder) LinkURL(v string) {
	b.Params["link_url"] = v
}

// FaveTrackPageInteractionBulder builder
//
// https://vk.com/dev/fave.trackPageInteraction
type FaveTrackPageInteractionBulder struct {
	api.Params
}

// NewFaveTrackPageInteractionBulder func
func NewFaveTrackPageInteractionBulder() *FaveTrackPageInteractionBulder {
	return &FaveTrackPageInteractionBulder{api.Params{}}
}

// UserID parameter
func (b *FaveTrackPageInteractionBulder) UserID(v int) {
	b.Params["user_id"] = v
}

// GroupID parameter
func (b *FaveTrackPageInteractionBulder) GroupID(v int) {
	b.Params["group_id"] = v
}
