package params // import "github.com/SevereCloud/vksdk/api/params"

import (
	"github.com/SevereCloud/vksdk/api"
)

// FaveAddArticleBuilder builder
//
// https://vk.com/dev/fave.addArticle
type FaveAddArticleBuilder struct {
	api.Params
}

// NewFaveAddArticleBuilder func
func NewFaveAddArticleBuilder() *FaveAddArticleBuilder {
	return &FaveAddArticleBuilder{api.Params{}}
}

// URL parameter
func (b *FaveAddArticleBuilder) URL(v string) {
	b.Params["url"] = v
}

// FaveAddLinkBuilder builder
//
// Adds a link to user faves.
//
// https://vk.com/dev/fave.addLink
type FaveAddLinkBuilder struct {
	api.Params
}

// NewFaveAddLinkBuilder func
func NewFaveAddLinkBuilder() *FaveAddLinkBuilder {
	return &FaveAddLinkBuilder{api.Params{}}
}

// Link Link URL.
func (b *FaveAddLinkBuilder) Link(v string) {
	b.Params["link"] = v
}

// FaveAddPageBuilder builder
//
// https://vk.com/dev/fave.addPage
type FaveAddPageBuilder struct {
	api.Params
}

// NewFaveAddPageBuilder func
func NewFaveAddPageBuilder() *FaveAddPageBuilder {
	return &FaveAddPageBuilder{api.Params{}}
}

// UserID parameter
func (b *FaveAddPageBuilder) UserID(v int) {
	b.Params["user_id"] = v
}

// GroupID parameter
func (b *FaveAddPageBuilder) GroupID(v int) {
	b.Params["group_id"] = v
}

// FaveAddPostBuilder builder
//
// https://vk.com/dev/fave.addPost
type FaveAddPostBuilder struct {
	api.Params
}

// NewFaveAddPostBuilder func
func NewFaveAddPostBuilder() *FaveAddPostBuilder {
	return &FaveAddPostBuilder{api.Params{}}
}

// OwnerID parameter
func (b *FaveAddPostBuilder) OwnerID(v int) {
	b.Params["owner_id"] = v
}

// ID parameter
func (b *FaveAddPostBuilder) ID(v int) {
	b.Params["id"] = v
}

// AccessKey parameter
func (b *FaveAddPostBuilder) AccessKey(v string) {
	b.Params["access_key"] = v
}

// FaveAddProductBuilder builder
//
// https://vk.com/dev/fave.addProduct
type FaveAddProductBuilder struct {
	api.Params
}

// NewFaveAddProductBuilder func
func NewFaveAddProductBuilder() *FaveAddProductBuilder {
	return &FaveAddProductBuilder{api.Params{}}
}

// OwnerID parameter
func (b *FaveAddProductBuilder) OwnerID(v int) {
	b.Params["owner_id"] = v
}

// ID parameter
func (b *FaveAddProductBuilder) ID(v int) {
	b.Params["id"] = v
}

// AccessKey parameter
func (b *FaveAddProductBuilder) AccessKey(v string) {
	b.Params["access_key"] = v
}

// FaveAddTagBuilder builder
//
// https://vk.com/dev/fave.addTag
type FaveAddTagBuilder struct {
	api.Params
}

// NewFaveAddTagBuilder func
func NewFaveAddTagBuilder() *FaveAddTagBuilder {
	return &FaveAddTagBuilder{api.Params{}}
}

// Name parameter
func (b *FaveAddTagBuilder) Name(v string) {
	b.Params["name"] = v
}

// FaveAddVideoBuilder builder
//
// https://vk.com/dev/fave.addVideo
type FaveAddVideoBuilder struct {
	api.Params
}

// NewFaveAddVideoBuilder func
func NewFaveAddVideoBuilder() *FaveAddVideoBuilder {
	return &FaveAddVideoBuilder{api.Params{}}
}

// OwnerID parameter
func (b *FaveAddVideoBuilder) OwnerID(v int) {
	b.Params["owner_id"] = v
}

// ID parameter
func (b *FaveAddVideoBuilder) ID(v int) {
	b.Params["id"] = v
}

// AccessKey parameter
func (b *FaveAddVideoBuilder) AccessKey(v string) {
	b.Params["access_key"] = v
}

// FaveEditTagBuilder builder
//
// https://vk.com/dev/fave.editTag
type FaveEditTagBuilder struct {
	api.Params
}

// NewFaveEditTagBuilder func
func NewFaveEditTagBuilder() *FaveEditTagBuilder {
	return &FaveEditTagBuilder{api.Params{}}
}

// ID parameter
func (b *FaveEditTagBuilder) ID(v int) {
	b.Params["id"] = v
}

// Name parameter
func (b *FaveEditTagBuilder) Name(v string) {
	b.Params["name"] = v
}

// FaveGetBuilder builder
//
// https://vk.com/dev/fave.get
type FaveGetBuilder struct {
	api.Params
}

// NewFaveGetBuilder func
func NewFaveGetBuilder() *FaveGetBuilder {
	return &FaveGetBuilder{api.Params{}}
}

// Extended '1' — to return additional 'wall', 'profiles', and 'groups' fields. By default: '0'.
func (b *FaveGetBuilder) Extended(v bool) {
	b.Params["extended"] = v
}

// ItemType parameter
func (b *FaveGetBuilder) ItemType(v string) {
	b.Params["item_type"] = v
}

// TagID Tag ID.
func (b *FaveGetBuilder) TagID(v int) {
	b.Params["tag_id"] = v
}

// Offset Offset needed to return a specific subset of users.
func (b *FaveGetBuilder) Offset(v int) {
	b.Params["offset"] = v
}

// Count Number of users to return.
func (b *FaveGetBuilder) Count(v int) {
	b.Params["count"] = v
}

// Fields parameter
func (b *FaveGetBuilder) Fields(v string) {
	b.Params["fields"] = v
}

// IsFromSnackbar parameter
func (b *FaveGetBuilder) IsFromSnackbar(v bool) {
	b.Params["is_from_snackbar"] = v
}

// FaveGetPagesBuilder builder
//
// https://vk.com/dev/fave.getPages
type FaveGetPagesBuilder struct {
	api.Params
}

// NewFaveGetPagesBuilder func
func NewFaveGetPagesBuilder() *FaveGetPagesBuilder {
	return &FaveGetPagesBuilder{api.Params{}}
}

// Offset parameter
func (b *FaveGetPagesBuilder) Offset(v int) {
	b.Params["offset"] = v
}

// Count parameter
func (b *FaveGetPagesBuilder) Count(v int) {
	b.Params["count"] = v
}

// Type parameter
func (b *FaveGetPagesBuilder) Type(v string) {
	b.Params["type"] = v
}

// Fields parameter
func (b *FaveGetPagesBuilder) Fields(v []string) {
	b.Params["fields"] = v
}

// TagID parameter
func (b *FaveGetPagesBuilder) TagID(v int) {
	b.Params["tag_id"] = v
}

// FaveRemoveArticleBuilder builder
//
// https://vk.com/dev/fave.removeArticle
type FaveRemoveArticleBuilder struct {
	api.Params
}

// NewFaveRemoveArticleBuilder func
func NewFaveRemoveArticleBuilder() *FaveRemoveArticleBuilder {
	return &FaveRemoveArticleBuilder{api.Params{}}
}

// OwnerID parameter
func (b *FaveRemoveArticleBuilder) OwnerID(v int) {
	b.Params["owner_id"] = v
}

// ArticleID parameter
func (b *FaveRemoveArticleBuilder) ArticleID(v int) {
	b.Params["article_id"] = v
}

// FaveRemoveLinkBuilder builder
//
// Removes link from the user's faves.
//
// https://vk.com/dev/fave.removeLink
type FaveRemoveLinkBuilder struct {
	api.Params
}

// NewFaveRemoveLinkBuilder func
func NewFaveRemoveLinkBuilder() *FaveRemoveLinkBuilder {
	return &FaveRemoveLinkBuilder{api.Params{}}
}

// LinkID Link ID (can be obtained by [vk.com/dev/faves.getLinks|faves.getLinks] method).
func (b *FaveRemoveLinkBuilder) LinkID(v string) {
	b.Params["link_id"] = v
}

// Link Link URL
func (b *FaveRemoveLinkBuilder) Link(v string) {
	b.Params["link"] = v
}

// FaveRemovePageBuilder builder
//
// https://vk.com/dev/fave.removePage
type FaveRemovePageBuilder struct {
	api.Params
}

// NewFaveRemovePageBuilder func
func NewFaveRemovePageBuilder() *FaveRemovePageBuilder {
	return &FaveRemovePageBuilder{api.Params{}}
}

// UserID parameter
func (b *FaveRemovePageBuilder) UserID(v int) {
	b.Params["user_id"] = v
}

// GroupID parameter
func (b *FaveRemovePageBuilder) GroupID(v int) {
	b.Params["group_id"] = v
}

// FaveRemovePostBuilder builder
//
// https://vk.com/dev/fave.removePost
type FaveRemovePostBuilder struct {
	api.Params
}

// NewFaveRemovePostBuilder func
func NewFaveRemovePostBuilder() *FaveRemovePostBuilder {
	return &FaveRemovePostBuilder{api.Params{}}
}

// OwnerID parameter
func (b *FaveRemovePostBuilder) OwnerID(v int) {
	b.Params["owner_id"] = v
}

// ID parameter
func (b *FaveRemovePostBuilder) ID(v int) {
	b.Params["id"] = v
}

// FaveRemoveProductBuilder builder
//
// https://vk.com/dev/fave.removeProduct
type FaveRemoveProductBuilder struct {
	api.Params
}

// NewFaveRemoveProductBuilder func
func NewFaveRemoveProductBuilder() *FaveRemoveProductBuilder {
	return &FaveRemoveProductBuilder{api.Params{}}
}

// OwnerID parameter
func (b *FaveRemoveProductBuilder) OwnerID(v int) {
	b.Params["owner_id"] = v
}

// ID parameter
func (b *FaveRemoveProductBuilder) ID(v int) {
	b.Params["id"] = v
}

// FaveRemoveTagBuilder builder
//
// https://vk.com/dev/fave.removeTag
type FaveRemoveTagBuilder struct {
	api.Params
}

// NewFaveRemoveTagBuilder func
func NewFaveRemoveTagBuilder() *FaveRemoveTagBuilder {
	return &FaveRemoveTagBuilder{api.Params{}}
}

// ID parameter
func (b *FaveRemoveTagBuilder) ID(v int) {
	b.Params["id"] = v
}

// FaveReorderTagsBuilder builder
//
// https://vk.com/dev/fave.reorderTags
type FaveReorderTagsBuilder struct {
	api.Params
}

// NewFaveReorderTagsBuilder func
func NewFaveReorderTagsBuilder() *FaveReorderTagsBuilder {
	return &FaveReorderTagsBuilder{api.Params{}}
}

// IDs parameter
func (b *FaveReorderTagsBuilder) IDs(v []int) {
	b.Params["ids"] = v
}

// FaveSetPageTagsBuilder builder
//
// https://vk.com/dev/fave.setPageTags
type FaveSetPageTagsBuilder struct {
	api.Params
}

// NewFaveSetPageTagsBuilder func
func NewFaveSetPageTagsBuilder() *FaveSetPageTagsBuilder {
	return &FaveSetPageTagsBuilder{api.Params{}}
}

// UserID parameter
func (b *FaveSetPageTagsBuilder) UserID(v int) {
	b.Params["user_id"] = v
}

// GroupID parameter
func (b *FaveSetPageTagsBuilder) GroupID(v int) {
	b.Params["group_id"] = v
}

// TagIDs parameter
func (b *FaveSetPageTagsBuilder) TagIDs(v []int) {
	b.Params["tag_ids"] = v
}

// FaveSetTagsBuilder builder
//
// https://vk.com/dev/fave.setTags
type FaveSetTagsBuilder struct {
	api.Params
}

// NewFaveSetTagsBuilder func
func NewFaveSetTagsBuilder() *FaveSetTagsBuilder {
	return &FaveSetTagsBuilder{api.Params{}}
}

// ItemType parameter
func (b *FaveSetTagsBuilder) ItemType(v string) {
	b.Params["item_type"] = v
}

// ItemOwnerID parameter
func (b *FaveSetTagsBuilder) ItemOwnerID(v int) {
	b.Params["item_owner_id"] = v
}

// ItemID parameter
func (b *FaveSetTagsBuilder) ItemID(v int) {
	b.Params["item_id"] = v
}

// TagIDs parameter
func (b *FaveSetTagsBuilder) TagIDs(v []int) {
	b.Params["tag_ids"] = v
}

// LinkID parameter
func (b *FaveSetTagsBuilder) LinkID(v string) {
	b.Params["link_id"] = v
}

// LinkURL parameter
func (b *FaveSetTagsBuilder) LinkURL(v string) {
	b.Params["link_url"] = v
}

// FaveTrackPageInteractionBuilder builder
//
// https://vk.com/dev/fave.trackPageInteraction
type FaveTrackPageInteractionBuilder struct {
	api.Params
}

// NewFaveTrackPageInteractionBuilder func
func NewFaveTrackPageInteractionBuilder() *FaveTrackPageInteractionBuilder {
	return &FaveTrackPageInteractionBuilder{api.Params{}}
}

// UserID parameter
func (b *FaveTrackPageInteractionBuilder) UserID(v int) {
	b.Params["user_id"] = v
}

// GroupID parameter
func (b *FaveTrackPageInteractionBuilder) GroupID(v int) {
	b.Params["group_id"] = v
}
